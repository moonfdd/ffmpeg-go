package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/ffmpeg-go/libavcodec"
	"github.com/moonfdd/ffmpeg-go/libavformat"
	"github.com/moonfdd/ffmpeg-go/libavutil"
	"github.com/moonfdd/ffmpeg-go/libswresample"
)

func main0() (ret ffcommon.FInt) {
	var input_format_context, output_format_context *libavformat.AVFormatContext
	var input_codec_context, output_codec_context *libavcodec.AVCodecContext
	var resample_context *libswresample.SwrContext
	var fifo *libavutil.AVAudioFifo
	ret = libavutil.AVERROR_EXIT

	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <input file> <output file>\n", os.Args[0])
		os.Exit(1)
	}

	/* Open the input file for reading. */
	if open_input_file(os.Args[1], &input_format_context,
		&input_codec_context) != 0 {
		goto cleanup
	}
	/* Open the output file for writing. */
	if open_output_file(os.Args[2], input_codec_context,
		&output_format_context, &output_codec_context) != 0 {
		goto cleanup
	}
	/* Initialize the resampler to be able to convert audio sample formats. */
	if init_resampler(input_codec_context, output_codec_context,
		&resample_context) != 0 {
		goto cleanup
	}
	/* Initialize the FIFO buffer to store audio samples to be encoded. */
	if init_fifo(&fifo, output_codec_context) != 0 {
		goto cleanup
	}
	/* Write the header of the output file container. */
	if write_output_file_header(output_format_context) != 0 {
		goto cleanup
	}

	/* Loop as long as we have input samples to read or output samples
	 * to write; abort as soon as we have neither. */
	for {
		/* Use the encoder's desired frame size for processing. */
		var output_frame_size ffcommon.FInt = output_codec_context.FrameSize
		var finished ffcommon.FInt = 0

		/* Make sure that there is one frame worth of samples in the FIFO
		 * buffer so that the encoder can do its work.
		 * Since the decoder's and the encoder's frame size may differ, we
		 * need to FIFO buffer to store as many frames worth of input samples
		 * that they make up at least one frame worth of output samples. */
		for fifo.AvAudioFifoSize() < output_frame_size {
			/* Decode one frame worth of audio samples, convert it to the
			 * output sample format and put it into the FIFO buffer. */
			if read_decode_convert_and_store(fifo, input_format_context,
				input_codec_context,
				output_codec_context,
				resample_context, &finished) != 0 {
				goto cleanup

			}
			/* If we are at the end of the input file, we continue
			 * encoding the remaining audio samples to the output file. */
			if finished != 0 {
				break
			}
		}

		/* If we have enough samples for the encoder, we encode them.
		 * At the end of the file, we pass the remaining samples to
		 * the encoder. */
		for fifo.AvAudioFifoSize() >= output_frame_size ||
			(finished != 0 && fifo.AvAudioFifoSize() > 0) {
			/* Take one frame worth of audio samples from the FIFO buffer,
			 * encode it and write it to the output file. */
			if load_encode_and_write(fifo, output_format_context,
				output_codec_context) != 0 {
				goto cleanup
			}
		}

		/* If we are at the end of the input file and have encoded
		 * all remaining samples, we can exit this loop and finish. */
		if finished != 0 {
			var data_written ffcommon.FInt
			/* Flush the encoder as it may have delayed frames. */
			for {
				data_written = 0
				if encode_audio_frame(nil, output_format_context,
					output_codec_context, &data_written) != 0 {
					goto cleanup
				}
				if data_written == 0 {
					break
				}
			}
			break
		}
	}

	/* Write the trailer of the output file container. */
	if write_output_file_trailer(output_format_context) != 0 {
		goto cleanup
	}
	ret = 0

cleanup:
	if fifo != nil {
		fifo.AvAudioFifoFree()
	}
	libswresample.SwrFree(&resample_context)
	if output_codec_context != nil {
		libavcodec.AvcodecFreeContext(&output_codec_context)
	}
	if output_format_context != nil {
		libavformat.AvioClosep(&output_format_context.Pb)
		output_format_context.AvformatFreeContext()
	}
	if input_codec_context != nil {
		libavcodec.AvcodecFreeContext(&input_codec_context)
	}
	if input_format_context != nil {
		libavformat.AvformatCloseInput(&input_format_context)
	}

	return ret
}

/* The output bit rate in bit/s */
const OUTPUT_BIT_RATE = 96000

/* The number of output channels */
const OUTPUT_CHANNELS = 2

/**
 * Open an input file and the required decoder.
 * @param      filename             File to be opened
 * @param[out] input_format_context Format context of opened file
 * @param[out] input_codec_context  Codec context of opened file
 * @return Error code (0 if successful)
 */
func open_input_file(filename string,
	input_format_context **libavformat.AVFormatContext,
	input_codec_context **libavcodec.AVCodecContext) ffcommon.FInt {
	var avctx *libavcodec.AVCodecContext
	var input_codec *libavcodec.AVCodec
	var err ffcommon.FInt

	/* Open the input file to read from it. */
	err = libavformat.AvformatOpenInput(input_format_context, filename, nil,
		nil)
	if err < 0 {
		fmt.Printf("Could not open input file '%s' (error '%s')\n",
			filename, libavutil.AvErr2str(err))
		*input_format_context = nil
		return err
	}

	/* Get information on the input file (number of streams etc.). */
	err = (*input_format_context).AvformatFindStreamInfo(nil)
	if err < 0 {
		fmt.Printf("Could not open find stream info (error '%s')\n",
			libavutil.AvErr2str(err))
		libavformat.AvformatCloseInput(input_format_context)
		return err
	}

	/* Make sure that there is only one stream in the input file. */
	if (*input_format_context).NbStreams != 1 {
		fmt.Printf("Expected one audio input stream, but found %d\n",
			(*input_format_context).NbStreams)
		libavformat.AvformatCloseInput(input_format_context)
		return libavutil.AVERROR_EXIT
	}

	/* Find a decoder for the audio stream. */
	input_codec = libavcodec.AvcodecFindDecoder((*input_format_context).GetStream(0).Codecpar.CodecId)
	if input_codec == nil {
		fmt.Printf("Could not find input codec\n")
		libavformat.AvformatCloseInput(input_format_context)
		return libavutil.AVERROR_EXIT
	}

	/* Allocate a new decoding context. */
	avctx = input_codec.AvcodecAllocContext3()
	if avctx == nil {
		fmt.Printf("Could not allocate a decoding context\n")
		libavformat.AvformatCloseInput(input_format_context)
		return -libavutil.ENOMEM
	}

	/* Initialize the stream parameters with demuxer information. */
	err = avctx.AvcodecParametersToContext((*input_format_context).GetStream(0).Codecpar)
	if err < 0 {
		libavformat.AvformatCloseInput(input_format_context)
		libavcodec.AvcodecFreeContext(&avctx)
		return err
	}

	/* Open the decoder for the audio stream to use it later. */
	err = avctx.AvcodecOpen2(input_codec, nil)
	if err < 0 {
		fmt.Printf("Could not open input codec (error '%s')\n",
			libavutil.AvErr2str(err))
		libavcodec.AvcodecFreeContext(&avctx)
		libavformat.AvformatCloseInput(input_format_context)
		return err
	}

	/* Save the decoder context for easier access later. */
	*input_codec_context = avctx

	return 0
}

/**
 * Open an output file and the required encoder.
 * Also set some basic encoder parameters.
 * Some of these parameters are based on the input file's parameters.
 * @param      filename              File to be opened
 * @param      input_codec_context   Codec context of input file
 * @param[out] output_format_context Format context of output file
 * @param[out] output_codec_context  Codec context of output file
 * @return Error code (0 if successful)
 */
func open_output_file(filename string,
	input_codec_context *libavcodec.AVCodecContext,
	output_format_context **libavformat.AVFormatContext,
	output_codec_context **libavcodec.AVCodecContext) ffcommon.FInt {
	var avctx *libavcodec.AVCodecContext
	var output_io_context *libavformat.AVIOContext
	var stream *libavformat.AVStream
	var output_codec *libavcodec.AVCodec
	var err ffcommon.FInt

	/* Open the output file to write to it. */
	err = libavformat.AvioOpen(&output_io_context, filename, libavformat.AVIO_FLAG_WRITE)
	if err < 0 {
		fmt.Printf("Could not open output file '%s' (error '%s')\n",
			filename, libavutil.AvErr2str(err))
		return err
	}

	/* Create a new format context for the output container format. */
	*output_format_context = libavformat.AvformatAllocContext()
	if *output_format_context == nil {
		fmt.Printf("Could not allocate output format context\n")
		return -libavutil.ENOMEM
	}

	/* Associate the output file (pointer) with the container format context. */
	(*output_format_context).Pb = output_io_context

	/* Guess the desired container format based on the file extension. */
	(*output_format_context).Oformat = libavformat.AvGuessFormat("", filename,
		"")
	if (*output_format_context).Oformat == nil {
		fmt.Printf("Could not find output file format\n")
		goto cleanup
	}

	(*output_format_context).Url = ffcommon.UintPtrFromString(libavutil.AvStrdup(filename))
	if (*output_format_context).Url == 0 {
		fmt.Printf("Could not allocate url.\n")
		err = -libavutil.ENOMEM
		goto cleanup
	}

	/* Find the encoder to be used by its name. */
	output_codec = libavcodec.AvcodecFindEncoder(libavcodec.AV_CODEC_ID_AAC)
	if output_codec == nil {
		fmt.Printf("Could not find an AAC encoder.\n")
		goto cleanup
	}

	/* Create a new audio stream in the output file container. */
	stream = libavformat.AvformatAllocContext().AvformatNewStream(nil)
	if stream == nil {
		fmt.Printf("Could not create new stream\n")
		err = -libavutil.ENOMEM
		goto cleanup
	}

	avctx = output_codec.AvcodecAllocContext3()
	if avctx == nil {
		fmt.Printf("Could not allocate an encoding context\n")
		err = -libavutil.ENOMEM
		goto cleanup
	}

	/* Set the basic encoder parameters.
	 * The input file's sample rate is used to avoid a sample rate conversion. */
	avctx.Channels = OUTPUT_CHANNELS
	avctx.ChannelLayout = uint64(libavutil.AvGetDefaultChannelLayout(OUTPUT_CHANNELS))
	avctx.SampleRate = input_codec_context.SampleRate
	avctx.SampleFmt = output_codec.GetSampleFmt(0)
	avctx.BitRate = OUTPUT_BIT_RATE

	/* Allow the use of the experimental AAC encoder. */
	avctx.StrictStdCompliance = libavcodec.FF_COMPLIANCE_EXPERIMENTAL

	/* Set the sample rate for the container. */
	stream.TimeBase.Den = input_codec_context.SampleRate
	stream.TimeBase.Num = 1

	/* Some container formats (like MP4) require global headers to be present.
	 * Mark the encoder so that it behaves accordingly. */
	if (*output_format_context).Oformat.Flags&libavformat.AVFMT_GLOBALHEADER != 0 {
		avctx.Flags |= libavcodec.AV_CODEC_FLAG_GLOBAL_HEADER
	}

	/* Open the encoder for the audio stream to use it later. */
	err = avctx.AvcodecOpen2(output_codec, nil)
	if err < 0 {
		fmt.Printf("Could not open output codec (error '%s')\n",
			libavutil.AvErr2str(err))
		goto cleanup
	}

	err = stream.Codecpar.AvcodecParametersFromContext(avctx)
	if err < 0 {
		fmt.Printf("Could not initialize stream parameters\n")
		goto cleanup
	}

	/* Save the encoder context for easier access later. */
	*output_codec_context = avctx

	return 0

cleanup:
	libavcodec.AvcodecFreeContext(&avctx)
	libavformat.AvioClosep(&(*output_format_context).Pb)
	(*output_format_context).AvformatFreeContext()
	*output_format_context = nil
	if err < 0 {
		return err
	} else {
		return libavutil.AVERROR_EXIT
	}
}

/**
 * Initialize one data packet for reading or writing.
 * @param[out] packet Packet to be initialized
 * @return Error code (0 if successful)
 */
func init_packet(packet **libavcodec.AVPacket) ffcommon.FInt {
	*packet = libavcodec.AvPacketAlloc()
	if *packet == nil {
		fmt.Printf("Could not allocate packet\n")
		return -libavutil.ENOMEM
	}
	return 0
}

/**
 * Initialize one audio frame for reading from the input file.
 * @param[out] frame Frame to be initialized
 * @return Error code (0 if successful)
 */
func init_input_frame(frame **libavutil.AVFrame) ffcommon.FInt {
	*frame = libavutil.AvFrameAlloc()
	if *frame == nil {
		fmt.Printf("Could not allocate input frame\n")
		return -libavutil.ENOMEM
	}
	return 0
}

/**
 * Initialize the audio resampler based on the input and output codec settings.
 * If the input and output sample formats differ, a conversion is required
 * libswresample takes care of this, but requires initialization.
 * @param      input_codec_context  Codec context of the input file
 * @param      output_codec_context Codec context of the output file
 * @param[out] resample_context     Resample context for the required conversion
 * @return Error code (0 if successful)
 */
func init_resampler(input_codec_context *libavcodec.AVCodecContext,
	output_codec_context *libavcodec.AVCodecContext,
	resample_context **libswresample.SwrContext) ffcommon.FInt {
	var err ffcommon.FInt

	/*
	 * Create a resampler context for the conversion.
	 * Set the conversion parameters.
	 * Default channel layouts based on the number of channels
	 * are assumed for simplicity (they are sometimes not detected
	 * properly by the demuxer and/or decoder).
	 */
	var s *libswresample.SwrContext
	*resample_context = s.SwrAllocSetOpts(
		libavutil.AvGetDefaultChannelLayout(output_codec_context.Channels),
		output_codec_context.SampleFmt,
		output_codec_context.SampleRate,
		libavutil.AvGetDefaultChannelLayout(input_codec_context.Channels),
		input_codec_context.SampleFmt,
		input_codec_context.SampleRate,
		0, uintptr(0))
	if *resample_context == nil {
		fmt.Printf("Could not allocate resample context\n")
		return -libavutil.ENOMEM
	}
	// /*
	// * Perform a sanity check so that the number of converted samples is
	// * not greater than the number of samples to be converted.
	// * If the sample rates differ, this case has to be handled differently
	// */
	// av_assert0(output_codec_context->sample_rate == input_codec_context->sample_rate);

	/* Open the resampler with the specified parameters. */
	err = (*resample_context).SwrInit()
	if err < 0 {
		fmt.Printf("Could not open resample context\n")
		libswresample.SwrFree(resample_context)
		return err
	}
	return 0
}

/**
 * Initialize a FIFO buffer for the audio samples to be encoded.
 * @param[out] fifo                 Sample buffer
 * @param      output_codec_context Codec context of the output file
 * @return Error code (0 if successful)
 */
func init_fifo(fifo **libavutil.AVAudioFifo, output_codec_context *libavcodec.AVCodecContext) ffcommon.FInt {
	/* Create the FIFO buffer based on the specified output sample format. */
	*fifo = libavutil.AvAudioFifoAlloc(output_codec_context.SampleFmt,
		output_codec_context.Channels, 1)
	if *fifo == nil {
		fmt.Printf("Could not allocate FIFO\n")
		return -libavutil.ENOMEM
	}
	return 0
}

/**
 * Write the header of the output file container.
 * @param output_format_context Format context of the output file
 * @return Error code (0 if successful)
 */
func write_output_file_header(output_format_context *libavformat.AVFormatContext) ffcommon.FInt {
	var err ffcommon.FInt
	err = output_format_context.AvformatWriteHeader(nil)
	if err < 0 {
		fmt.Printf("Could not write output file header (error '%s')\n",
			libavutil.AvErr2str(err))
		return err
	}
	return 0
}

/**
 * Decode one audio frame from the input file.
 * @param      frame                Audio frame to be decoded
 * @param      input_format_context Format context of the input file
 * @param      input_codec_context  Codec context of the input file
 * @param[out] data_present         Indicates whether data has been decoded
 * @param[out] finished             Indicates whether the end of file has
 *                                  been reached and all data has been
 *                                  decoded. If this flag is false, there
 *                                  is more data to be decoded, i.e., this
 *                                  function has to be called again.
 * @return Error code (0 if successful)
 */
func decode_audio_frame(frame *libavutil.AVFrame,
	input_format_context *libavformat.AVFormatContext,
	input_codec_context *libavcodec.AVCodecContext,
	data_present, finished *ffcommon.FInt) ffcommon.FInt {
	/* Packet used for temporary storage. */
	var input_packet *libavcodec.AVPacket
	var err ffcommon.FInt

	err = init_packet(&input_packet)
	if err < 0 {
		return err
	}

	/* Read one audio frame from the input file into a temporary packet. */
	err = input_format_context.AvReadFrame(input_packet)
	if err < 0 {
		/* If we are at the end of the file, flush the decoder below. */
		if err == libavutil.AVERROR_EOF {
			*finished = 1
		} else {
			fmt.Printf("Could not read frame (error '%s')\n",
				libavutil.AvErr2str(err))
			goto cleanup
		}
	}

	/* Send the audio frame stored in the temporary packet to the decoder.
	 * The input audio stream decoder is used to do this. */
	err = input_codec_context.AvcodecSendPacket(input_packet)
	if err < 0 {
		fmt.Printf("Could not send packet for decoding (error '%s')\n",
			libavutil.AvErr2str(err))
		goto cleanup
	}

	/* Receive one frame from the decoder. */
	err = input_codec_context.AvcodecReceiveFrame(frame)
	/* If the decoder asks for more data to be able to decode a frame,
	 * return indicating that no data is present. */
	if err == -libavutil.EAGAIN {
		err = 0
		goto cleanup
		/* If the end of the input file is reached, stop decoding. */
	} else if err == libavutil.AVERROR_EOF {
		*finished = 1
		err = 0
		goto cleanup
	} else if err < 0 {
		fmt.Printf("Could not decode frame (error '%s')\n",
			libavutil.AvErr2str(err))
		goto cleanup
		/* Default case: Return decoded data. */
	} else {
		*data_present = 1
		goto cleanup
	}

cleanup:
	libavcodec.AvPacketFree(&input_packet)

	return err
}

/**
 * Initialize a temporary storage for the specified number of audio samples.
 * The conversion requires temporary storage due to the different format.
 * The number of audio samples to be allocated is specified in frame_size.
 * @param[out] converted_input_samples Array of converted samples. The
 *                                     dimensions are reference, channel
 *                                     (for multi-channel audio), sample.
 * @param      output_codec_context    Codec context of the output file
 * @param      frame_size              Number of samples to be converted in
 *                                     each round
 * @return Error code (0 if successful)
 */
func init_converted_samples(converted_input_samples ***ffcommon.FUint8T,
	output_codec_context *libavcodec.AVCodecContext,
	frame_size ffcommon.FInt) ffcommon.FInt {
	var err ffcommon.FInt

	/* Allocate as many pointers as there are audio channels.
	 * Each pointer will later point to the audio samples of the corresponding
	 * channels (although it may be NULL for interleaved formats).
	 */

	*converted_input_samples = (**byte)(unsafe.Pointer(libavutil.AvCalloc(uint64(output_codec_context.Channels), 8)))
	if *converted_input_samples == nil {
		fmt.Printf("Could not allocate converted input sample pointers\n")
		return -libavutil.ENOMEM
	}

	/* Allocate memory for the samples of all channels in one consecutive
	 * block for convenience. */
	err = libavutil.AvSamplesAlloc(*converted_input_samples, nil,
		output_codec_context.Channels,
		frame_size,
		output_codec_context.SampleFmt, 0)
	if err < 0 {
		fmt.Printf("Could not allocate converted input samples (error '%s')\n",
			libavutil.AvErr2str(err))
		libavutil.AvFreep(uintptr(unsafe.Pointer(*converted_input_samples)))
		libavutil.AvFree(uintptr(unsafe.Pointer(*converted_input_samples)))
		return err
	}
	return 0
}

/**
 * Convert the input audio samples into the output sample format.
 * The conversion happens on a per-frame basis, the size of which is
 * specified by frame_size.
 * @param      input_data       Samples to be decoded. The dimensions are
 *                              channel (for multi-channel audio), sample.
 * @param[out] converted_data   Converted samples. The dimensions are channel
 *                              (for multi-channel audio), sample.
 * @param      frame_size       Number of samples to be converted
 * @param      resample_context Resample context for the conversion
 * @return Error code (0 if successful)
 */
func convert_samples(input_data, converted_data **ffcommon.FUint8T, frame_size ffcommon.FInt,
	resample_context *libswresample.SwrContext) ffcommon.FInt {
	var err ffcommon.FInt

	/* Convert the samples using the resampler. */
	err = resample_context.SwrConvert(converted_data, frame_size,
		input_data, frame_size)
	if err < 0 {
		fmt.Printf("Could not convert input samples (error '%s')\n",
			libavutil.AvErr2str(err))
		return err
	}

	return 0
}

/**
 * Add converted input audio samples to the FIFO buffer for later processing.
 * @param fifo                    Buffer to add the samples to
 * @param converted_input_samples Samples to be added. The dimensions are channel
 *                                (for multi-channel audio), sample.
 * @param frame_size              Number of samples to be converted
 * @return Error code (0 if successful)
 */
func add_samples_to_fifo(fifo *libavutil.AVAudioFifo,
	converted_input_samples **ffcommon.FUint8T,
	frame_size ffcommon.FInt) ffcommon.FInt {
	var err ffcommon.FInt

	/* Make the FIFO as large as it needs to be to hold both,
	 * the old and the new samples. */
	err = fifo.AvAudioFifoRealloc(fifo.AvAudioFifoSize() + frame_size)
	if err < 0 {
		fmt.Printf("Could not reallocate FIFO\n")
		return err
	}

	/* Store the new samples in the FIFO buffer. */
	if fifo.AvAudioFifoWrite((*uintptr)(unsafe.Pointer(converted_input_samples)), frame_size) < frame_size {
		fmt.Printf("Could not write data to FIFO\n")
		return libavutil.AVERROR_EXIT
	}
	return 0
}

/**
 * Read one audio frame from the input file, decode, convert and store
 * it in the FIFO buffer.
 * @param      fifo                 Buffer used for temporary storage
 * @param      input_format_context Format context of the input file
 * @param      input_codec_context  Codec context of the input file
 * @param      output_codec_context Codec context of the output file
 * @param      resampler_context    Resample context for the conversion
 * @param[out] finished             Indicates whether the end of file has
 *                                  been reached and all data has been
 *                                  decoded. If this flag is false,
 *                                  there is more data to be decoded,
 *                                  i.e., this function has to be called
 *                                  again.
 * @return Error code (0 if successful)
 */
func read_decode_convert_and_store(fifo *libavutil.AVAudioFifo,
	input_format_context *libavformat.AVFormatContext,
	input_codec_context *libavcodec.AVCodecContext,
	output_codec_context *libavcodec.AVCodecContext,
	resampler_context *libswresample.SwrContext,
	finished *ffcommon.FInt) ffcommon.FInt {
	/* Temporary storage of the input samples of the frame read from the file. */
	var input_frame *libavutil.AVFrame
	/* Temporary storage for the converted input samples. */
	var converted_input_samples **ffcommon.FUint8T
	var data_present ffcommon.FInt
	var ret ffcommon.FInt = libavutil.AVERROR_EXIT

	/* Initialize temporary storage for one input frame. */
	if init_input_frame(&input_frame) != 0 {
		goto cleanup
	}
	/* Decode one frame worth of audio samples. */
	if decode_audio_frame(input_frame, input_format_context,
		input_codec_context, &data_present, finished) != 0 {
		goto cleanup
	}
	/* If we are at the end of the file and there are no more samples
	 * in the decoder which are delayed, we are actually finished.
	 * This must not be treated as an error. */
	if (*finished) != 0 {
		ret = 0
		goto cleanup
	}
	/* If there is decoded data, convert and store it. */
	if data_present != 0 {
		/* Initialize the temporary storage for the converted input samples. */
		if init_converted_samples(&converted_input_samples, output_codec_context,
			input_frame.NbSamples) != 0 {
			goto cleanup
		}

		/* Convert the input samples to the desired output sample format.
		 * This requires a temporary storage provided by converted_input_samples. */
		if convert_samples(input_frame.ExtendedData, converted_input_samples,
			input_frame.NbSamples, resampler_context) != 0 {
			goto cleanup
		}

		/* Add the converted input samples to the FIFO buffer for later processing. */
		if add_samples_to_fifo(fifo, converted_input_samples,
			input_frame.NbSamples) != 0 {
			goto cleanup
		}
		ret = 0
	}
	ret = 0

cleanup:
	if converted_input_samples != nil {
		libavutil.AvFreep(uintptr(unsafe.Pointer(converted_input_samples)))
		libavutil.AvFree(uintptr(unsafe.Pointer(converted_input_samples)))
		// free(converted_input_samples);
	}
	libavutil.AvFrameFree(&input_frame)

	return ret
}

/**
 * Initialize one input frame for writing to the output file.
 * The frame will be exactly frame_size samples large.
 * @param[out] frame                Frame to be initialized
 * @param      output_codec_context Codec context of the output file
 * @param      frame_size           Size of the frame
 * @return Error code (0 if successful)
 */
func init_output_frame(frame **libavutil.AVFrame,
	output_codec_context *libavcodec.AVCodecContext,
	frame_size ffcommon.FInt) ffcommon.FInt {
	var err ffcommon.FInt

	/* Create a new frame to store the audio samples. */
	*frame = libavutil.AvFrameAlloc()
	if *frame == nil {
		fmt.Printf("Could not allocate output frame\n")
		return libavutil.AVERROR_EXIT
	}

	/* Set the frame's parameters, especially its size and format.
	 * av_frame_get_buffer needs this to allocate memory for the
	 * audio samples of the frame.
	 * Default channel layouts based on the number of channels
	 * are assumed for simplicity. */
	(*frame).NbSamples = frame_size
	(*frame).ChannelLayout = output_codec_context.ChannelLayout
	(*frame).Format = int32(output_codec_context.SampleFmt)
	(*frame).SampleRate = output_codec_context.SampleRate

	/* Allocate the samples of the created frame. This call will make
	 * sure that the audio frame can hold as many samples as specified. */
	err = (*frame).AvFrameGetBuffer(0)
	if err < 0 {
		fmt.Printf("Could not allocate output frame samples (error '%s')\n",
			libavutil.AvErr2str(err))
		libavutil.AvFrameFree(frame)
		return err
	}

	return 0
}

/* Global timestamp for the audio frames. */
var pts ffcommon.FInt64T = 0

/**
 * Encode one frame worth of audio to the output file.
 * @param      frame                 Samples to be encoded
 * @param      output_format_context Format context of the output file
 * @param      output_codec_context  Codec context of the output file
 * @param[out] data_present          Indicates whether data has been
 *                                   encoded
 * @return Error code (0 if successful)
 */
func encode_audio_frame(frame *libavutil.AVFrame,
	output_format_context *libavformat.AVFormatContext,
	output_codec_context *libavcodec.AVCodecContext,
	data_present *ffcommon.FInt) ffcommon.FInt {
	/* Packet used for temporary storage. */
	var output_packet *libavcodec.AVPacket
	var err ffcommon.FInt

	err = init_packet(&output_packet)
	if err < 0 {
		return err
	}

	/* Set a timestamp based on the sample rate for the container. */
	if frame != nil {
		frame.Pts = pts
		pts += int64(frame.NbSamples)
	}

	/* Send the audio frame stored in the temporary packet to the encoder.
	 * The output audio stream encoder is used to do this. */
	err = output_codec_context.AvcodecSendFrame(frame)
	/* The encoder signals that it has nothing more to encode. */
	if err == libavutil.AVERROR_EOF {
		err = 0
		goto cleanup
	} else if err < 0 {
		fmt.Printf("Could not send packet for encoding (error '%s')\n",
			libavutil.AvErr2str(err))
		goto cleanup
	}

	/* Receive one encoded frame from the encoder. */
	err = output_codec_context.AvcodecReceivePacket(output_packet)
	/* If the encoder asks for more data to be able to provide an
	 * encoded frame, return indicating that no data is present. */
	if err == -libavutil.EAGAIN {
		err = 0
		goto cleanup
		/* If the last frame has been encoded, stop encoding. */
	} else if err == libavutil.AVERROR_EOF {
		err = 0
		goto cleanup
	} else if err < 0 {
		fmt.Printf("Could not encode frame (error '%s')\n",
			libavutil.AvErr2str(err))
		goto cleanup
		/* Default case: Return encoded data. */
	} else {
		*data_present = 1
	}

	/* Write one audio frame from the temporary packet to the output file. */
	if *data_present != 0 {
		err = output_format_context.AvWriteFrame(output_packet)
		if err < 0 {
			fmt.Printf("Could not write frame (error '%s')\n",
				libavutil.AvErr2str(err))
			goto cleanup
		}
	}

cleanup:
	libavcodec.AvPacketFree(&output_packet)
	return err
}
func getMin(a, b int32) int32 {
	if a < b {
		return a
	} else {
		return b
	}
}

/**
 * Load one audio frame from the FIFO buffer, encode and write it to the
 * output file.
 * @param fifo                  Buffer used for temporary storage
 * @param output_format_context Format context of the output file
 * @param output_codec_context  Codec context of the output file
 * @return Error code (0 if successful)
 */
func load_encode_and_write(fifo *libavutil.AVAudioFifo,
	output_format_context *libavformat.AVFormatContext,
	output_codec_context *libavcodec.AVCodecContext) ffcommon.FInt {
	/* Temporary storage of the output samples of the frame written to the file. */
	var output_frame *libavutil.AVFrame
	/* Use the maximum number of possible samples per frame.
	 * If there is less than the maximum possible frame size in the FIFO
	 * buffer use this number. Otherwise, use the maximum possible frame size. */
	frame_size := getMin(fifo.AvAudioFifoSize(),
		output_codec_context.FrameSize)
	var data_written ffcommon.FInt

	/* Initialize temporary storage for one output frame. */
	if init_output_frame(&output_frame, output_codec_context, frame_size) != 0 {
		return libavutil.AVERROR_EXIT
	}

	/* Read as many samples from the FIFO buffer as required to fill the frame.
	 * The samples are stored in the frame temporarily. */
	if fifo.AvAudioFifoRead((*uintptr)(unsafe.Pointer(&output_frame.Data)), frame_size) < frame_size {
		fmt.Printf("Could not read data from FIFO\n")
		libavutil.AvFrameFree(&output_frame)
		return libavutil.AVERROR_EXIT
	}

	/* Encode one frame worth of audio samples. */
	if encode_audio_frame(output_frame, output_format_context,
		output_codec_context, &data_written) != 0 {
		libavutil.AvFrameFree(&output_frame)
		return libavutil.AVERROR_EXIT
	}
	libavutil.AvFrameFree(&output_frame)
	return 0
}

/**
 * Write the trailer of the output file container.
 * @param output_format_context Format context of the output file
 * @return Error code (0 if successful)
 */
func write_output_file_trailer(output_format_context *libavformat.AVFormatContext) ffcommon.FInt {
	var err ffcommon.FInt
	err = output_format_context.AvWriteTrailer()
	if err < 0 {
		fmt.Printf("Could not write output file trailer (error '%s')\n",
			libavutil.AvErr2str(err))
		return err
	}
	return 0
}

func main() {

	os.Setenv("Path", "./lib;"+os.Getenv("Path"))
	ffcommon.SetAvutilPath("./lib/avutil-56.dll")
	ffcommon.SetAvcodecPath("./lib/avcodec-58.dll")
	ffcommon.SetAvdevicePath("./lib/avdevice-58.dll")
	ffcommon.SetAvfilterPath("./lib/avfilter-56.dll")
	ffcommon.SetAvformatPath("./lib/avformat-58.dll")
	ffcommon.SetAvpostprocPath("./lib/postproc-55.dll")
	ffcommon.SetAvswresamplePath("./lib/swresample-3.dll")
	ffcommon.SetAvswscalePath("./lib/swscale-5.dll")

	genDir := "./out"
	_, err := os.Stat(genDir)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(genDir, 0777) //  Everyone can read write and execute
		}
	}

	main0()
}
