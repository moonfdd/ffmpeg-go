package main

import (
	"fmt"
	"math"
	"os"
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/ffmpeg-go/libavcodec"
	"github.com/moonfdd/ffmpeg-go/libavformat"
	"github.com/moonfdd/ffmpeg-go/libavutil"
	"github.com/moonfdd/ffmpeg-go/libswresample"
	"github.com/moonfdd/ffmpeg-go/libswscale"
)

func main() {
	// 示例本程序会生成一个合成的音频和视频流，并将它们编码和封装输出到输出文件，输出格式是根据文件扩展名自动猜测的。
	// https://www.lmlphp.com/user/4129/article/item/31675/

	os.Setenv("Path", os.Getenv("Path")+";./lib")
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

func main0() (ret ffcommon.FInt) {
	var video_st, audio_st OutputStream
	var filename string
	var fmt0 *libavformat.AVOutputFormat
	var oc *libavformat.AVFormatContext
	var audio_codec, video_codec *libavcodec.AVCodec
	var have_video, have_audio ffcommon.FInt
	var encode_video, encode_audio ffcommon.FInt
	var opt *libavutil.AVDictionary
	var i ffcommon.FInt

	if len(os.Args) < 2 {
		fmt.Printf("usage: %s output_file\nAPI example program to output a media file with libavformat.\nThis program generates a synthetic audio and video stream, encodes and\nmuxes them into a file named output_file.\nThe output format is automatically guessed according to the file extension.\nRaw images can also be output by using '%%d' in the filename.\n\n", os.Args[0])
		return 1
	}

	filename = os.Args[1]
	for i = 2; i+1 < ffcommon.FInt(len(os.Args)); i += 2 {
		if os.Args[i] == "-flags" || os.Args[i] == "-fflags" {
			libavutil.AvDictSet(&opt, os.Args[i+1][1:], os.Args[i+1], 0)
		}
	}

	/* allocate the output media context */
	libavformat.AvformatAllocOutputContext2(&oc, nil, "", filename)
	if oc == nil {
		fmt.Printf("Could not deduce output format from file extension: using MPEG.\n")
		libavformat.AvformatAllocOutputContext2(&oc, nil, "mpeg", filename)
	}
	if oc == nil {
		return 1
	}

	fmt0 = oc.Oformat

	/* Add the audio and video streams using the default format codecs
	 * and initialize the codecs. */
	if fmt0.VideoCodec != libavcodec.AV_CODEC_ID_NONE {
		add_stream(&video_st, oc, &video_codec, fmt0.VideoCodec)
		have_video = 1
		encode_video = 1
	}
	if fmt0.AudioCodec != libavcodec.AV_CODEC_ID_NONE {
		add_stream(&audio_st, oc, &audio_codec, fmt0.AudioCodec)
		have_audio = 1
		encode_audio = 1
	}

	// /* Now that all the parameters are set, we can open the audio and
	//  * video codecs and allocate the necessary encode buffers. */
	if have_video != 0 {
		open_video(oc, video_codec, &video_st, opt)
	}

	if have_audio != 0 {
		open_audio(oc, audio_codec, &audio_st, opt)
	}

	oc.AvDumpFormat(0, filename, 1)

	/* open the output file, if needed */
	if fmt0.Flags&libavformat.AVFMT_NOFILE == 0 {
		ret = libavformat.AvioOpen(&oc.Pb, filename, libavformat.AVIO_FLAG_WRITE)
		if ret < 0 {
			fmt.Printf("Could not open '%s': %s\n", filename,
				libavutil.AvErr2str(ret))
			return 1
		}
	}

	/* Write the stream header, if any. */
	ret = oc.AvformatWriteHeader(&opt)
	if ret < 0 {
		fmt.Printf("Error occurred when opening output file: %s\n",
			libavutil.AvErr2str(ret))
		return 1
	}

	for encode_video != 0 || encode_audio != 0 {
		/* select the stream to encode */
		if encode_video != 0 &&
			(encode_audio == 0 || libavutil.AvCompareTs(video_st.next_pts, video_st.enc.TimeBase,
				audio_st.next_pts, audio_st.enc.TimeBase) <= 0) {
			if write_video_frame(oc, &video_st) == 0 {
				encode_video = 1
			} else {
				encode_video = 0
			}
		} else {
			if write_audio_frame(oc, &audio_st) == 0 {
				encode_audio = 1
			} else {
				encode_audio = 0
			}
		}
	}

	/* Write the trailer, if any. The trailer must be written before you
	 * close the CodecContexts open when you wrote the header; otherwise
	 * av_write_trailer() may try to use memory that was freed on
	 * av_codec_close(). */
	oc.AvWriteTrailer()

	// /* Close each codec. */
	if have_video != 0 {
		close_stream(oc, &video_st)
	}
	if have_audio != 0 {
		close_stream(oc, &audio_st)
	}

	if fmt0.Flags&libavformat.AVFMT_NOFILE == 0 {
		/* Close the output file. */
		libavformat.AvioClosep(&oc.Pb)
	}

	/* free the stream */
	oc.AvformatFreeContext()

	return 0
}

const STREAM_DURATION = 10.0
const STREAM_FRAME_RATE = 25                        /* 25 images/s */
const STREAM_PIX_FMT = libavutil.AV_PIX_FMT_YUV420P /* default pix_fmt */

const SCALE_FLAGS = libswscale.SWS_BICUBIC

// a wrapper around a single output AVStream
type OutputStream struct {
	st  *libavformat.AVStream
	enc *libavcodec.AVCodecContext

	/* pts of the next frame that will be generated */
	next_pts      ffcommon.FInt64T
	samples_count ffcommon.FInt

	frame     *libavutil.AVFrame
	tmp_frame *libavutil.AVFrame

	t, tincr, tincr2 ffcommon.FFloat

	sws_ctx *libswscale.SwsContext
	swr_ctx *libswresample.SwrContext
}

func log_packet(fmt_ctx *libavformat.AVFormatContext, pkt *libavcodec.AVPacket) {
	time_base := &fmt_ctx.GetStream(pkt.StreamIndex).TimeBase

	fmt.Printf("pts:%s pts_time:%s dts:%s dts_time:%s duration:%s duration_time:%s stream_index:%d\n",
		libavutil.AvTs2str(pkt.Pts), libavutil.AvTs2timestr(pkt.Pts, time_base),
		libavutil.AvTs2str(pkt.Dts), libavutil.AvTs2timestr(pkt.Dts, time_base),
		libavutil.AvTs2str(pkt.Duration), libavutil.AvTs2timestr(pkt.Duration, time_base),
		pkt.StreamIndex)
}

func write_frame(fmt_ctx *libavformat.AVFormatContext, c *libavcodec.AVCodecContext, st *libavformat.AVStream, frame *libavutil.AVFrame) ffcommon.FInt {
	var ret ffcommon.FInt

	// send the frame to the encoder
	ret = c.AvcodecSendFrame(frame)
	if ret < 0 {
		fmt.Printf("Error sending a frame to the encoder: %s\n",
			libavutil.AvErr2str(ret))
		os.Exit(1)
	}

	for ret >= 0 {
		var pkt libavcodec.AVPacket

		ret = c.AvcodecReceivePacket(&pkt)
		if ret == -libavutil.EAGAIN || ret == libavutil.AVERROR_EOF {
			break
		} else if ret < 0 {
			fmt.Printf("Error encoding a frame: %s\n", libavutil.AvErr2str(ret))
			os.Exit(1)
		}

		/* rescale output packet timestamp values from codec to stream timebase */
		pkt.AvPacketRescaleTs(c.TimeBase, st.TimeBase)
		pkt.StreamIndex = uint32(st.Index)

		/* Write the compressed frame to the media file. */
		log_packet(fmt_ctx, &pkt)
		ret = fmt_ctx.AvInterleavedWriteFrame(&pkt)
		pkt.AvPacketUnref()
		if ret < 0 {
			fmt.Printf("Error while writing output packet: %s\n", libavutil.AvErr2str(ret))
			os.Exit(1)
		}
	}

	if ret == libavutil.AVERROR_EOF {
		return 1
	} else {
		return 0
	}
}

/* Add an output stream. */
func add_stream(ost *OutputStream, oc *libavformat.AVFormatContext, codec **libavcodec.AVCodec, codec_id libavcodec.AVCodecID) {
	var c *libavcodec.AVCodecContext
	var i ffcommon.FInt

	/* find the encoder */
	*codec = libavcodec.AvcodecFindEncoder(codec_id)
	if *codec == nil {
		fmt.Printf("Could not find encoder for '%s'\n",
			libavcodec.AvcodecGetName(codec_id))
		os.Exit(1)
	}

	ost.st = oc.AvformatNewStream(nil)
	if ost.st == nil {
		fmt.Printf("Could not allocate stream\n")
		os.Exit(1)
	}
	ost.st.Id = int32(oc.NbStreams) - 1
	c = (*codec).AvcodecAllocContext3()
	if c == nil {
		fmt.Printf("Could not alloc an encoding context\n")
		os.Exit(1)
	}
	ost.enc = c

	switch (*codec).Type {
	case libavutil.AVMEDIA_TYPE_AUDIO:
		if (*codec).SampleFmts != nil {
			c.SampleFmt = (*codec).GetSampleFmt(0)
		} else {
			c.SampleFmt = libavutil.AV_SAMPLE_FMT_FLTP
		}
		c.BitRate = 64000
		c.SampleRate = 44100
		if (*codec).SupportedSamplerates != nil {
			c.SampleRate = (*codec).GetSupportedSamplerate(0)
			for i = 0; (*codec).GetSupportedSamplerate(uint32(i)) != 0; i++ {
				if (*codec).GetSupportedSamplerate(uint32(i)) == 44100 {
					c.SampleRate = 44100
				}
			}
		}
		c.Channels = libavutil.AvGetChannelLayoutNbChannels(c.ChannelLayout)
		c.ChannelLayout = libavutil.AV_CH_LAYOUT_STEREO
		if (*codec).ChannelLayouts != nil {
			c.ChannelLayout = (*codec).GetChannelLayout(0)
			for i = 0; (*codec).GetChannelLayout(uint32(i)) != 0; i++ {
				if (*codec).GetChannelLayout(uint32(i)) == libavutil.AV_CH_LAYOUT_STEREO {
					c.ChannelLayout = libavutil.AV_CH_LAYOUT_STEREO
				}
			}
		}
		c.Channels = libavutil.AvGetChannelLayoutNbChannels(c.ChannelLayout)
		ost.st.TimeBase = libavutil.AVRational{1, c.SampleRate}
		break

	case libavutil.AVMEDIA_TYPE_VIDEO:
		c.CodecId = codec_id

		c.BitRate = 400000
		/* Resolution must be a multiple of two. */
		c.Width = 352
		c.Height = 288
		//     /* timebase: This is the fundamental unit of time (in seconds) in terms
		//      * of which frame timestamps are represented. For fixed-fps content,
		//      * timebase should be 1/framerate and timestamp increments should be
		//      * identical to 1. */
		ost.st.TimeBase = libavutil.AVRational{1, STREAM_FRAME_RATE}
		c.TimeBase = ost.st.TimeBase

		c.GopSize = 12 /* emit one intra frame every twelve frames at most */
		c.PixFmt = STREAM_PIX_FMT
		if c.CodecId == libavcodec.AV_CODEC_ID_MPEG2VIDEO {
			/* just for testing, we also add B-frames */
			c.MaxBFrames = 2
		}
		if c.CodecId == libavcodec.AV_CODEC_ID_MPEG1VIDEO {
			/* Needed to avoid using macroblocks in which some coeffs overflow.
			 * This does not happen with normal video, it just happens here as
			 * the motion of the chroma plane does not match the luma plane. */
			c.MbDecision = 2
		}
		break

	default:
		break
	}

	/* Some formats want stream headers to be separate. */
	if oc.Oformat.Flags&libavformat.AVFMT_GLOBALHEADER != 0 {
		c.Flags |= libavcodec.AV_CODEC_FLAG_GLOBAL_HEADER
	}
}

/**************************************************************/
/* audio output */

func alloc_audio_frame(sample_fmt libavutil.AVSampleFormat,
	channel_layout ffcommon.FUint64T,
	sample_rate, nb_samples ffcommon.FInt) *libavutil.AVFrame {
	frame := libavutil.AvFrameAlloc()
	var ret ffcommon.FInt

	if frame == nil {
		fmt.Printf("Error allocating an audio frame\n")
		os.Exit(1)
	}

	frame.Format = int32(sample_fmt)
	frame.ChannelLayout = channel_layout
	frame.SampleRate = sample_rate
	frame.NbSamples = nb_samples

	if nb_samples != 0 {
		ret = frame.AvFrameGetBuffer(0)
		if ret < 0 {
			fmt.Printf("Error allocating an audio buffer\n")
			os.Exit(1)
		}
	}

	return frame
}

func open_audio(oc *libavformat.AVFormatContext, codec *libavcodec.AVCodec, ost *OutputStream, opt_arg *libavutil.AVDictionary) {
	var c *libavcodec.AVCodecContext
	var nb_samples ffcommon.FInt
	var ret ffcommon.FInt
	var opt *libavutil.AVDictionary

	c = ost.enc

	/* open it */
	libavutil.AvDictCopy(&opt, opt_arg, 0)
	ret = c.AvcodecOpen2(codec, &opt)
	libavutil.AvDictFree(&opt)
	if ret < 0 {
		fmt.Printf("Could not open audio codec: %s\n", libavutil.AvErr2str(ret))
		os.Exit(1)
	}

	/* init signal generator */
	ost.t = 0
	ost.tincr = float32(2 * libavutil.M_PI * 110.0 / float64(c.SampleRate))
	// /* increment frequency by 110 Hz per second */
	ost.tincr2 = float32(2 * libavutil.M_PI * 110.0 / float64(c.SampleRate) / float64(c.SampleRate))

	if c.Codec.Capabilities&libavcodec.AV_CODEC_CAP_VARIABLE_FRAME_SIZE != 0 {
		nb_samples = 10000
	} else {
		nb_samples = c.FrameSize
	}

	ost.frame = alloc_audio_frame(c.SampleFmt, c.ChannelLayout,
		c.SampleRate, nb_samples)
	ost.tmp_frame = alloc_audio_frame(libavutil.AV_SAMPLE_FMT_S16, c.ChannelLayout,
		c.SampleRate, nb_samples)

	/* copy the stream parameters to the muxer */
	ret = ost.st.Codecpar.AvcodecParametersFromContext(c)
	if ret < 0 {
		fmt.Printf("Could not copy the stream parameters\n")
		os.Exit(1)
	}

	/* create resampler context */
	ost.swr_ctx = libswresample.SwrAlloc()
	if ost.swr_ctx == nil {
		fmt.Printf("Could not allocate resampler context\n")
		os.Exit(1)
	}

	// /* set options */
	libavutil.AvOptSetInt(uintptr(unsafe.Pointer(ost.swr_ctx)), "in_channel_count", int64(c.Channels), 0)
	libavutil.AvOptSetInt(uintptr(unsafe.Pointer(ost.swr_ctx)), "in_sample_rate", int64(c.SampleRate), 0)
	libavutil.AvOptSetSampleFmt(uintptr(unsafe.Pointer(ost.swr_ctx)), "in_sample_fmt", libavutil.AV_SAMPLE_FMT_S16, 0)
	libavutil.AvOptSetInt(uintptr(unsafe.Pointer(ost.swr_ctx)), "out_channel_count", int64(c.Channels), 0)
	libavutil.AvOptSetInt(uintptr(unsafe.Pointer(ost.swr_ctx)), "out_sample_rate", int64(c.SampleRate), 0)
	libavutil.AvOptSetSampleFmt(uintptr(unsafe.Pointer(ost.swr_ctx)), "out_sample_fmt", c.SampleFmt, 0)
	/* initialize the resampling context */
	ret = ost.swr_ctx.SwrInit()
	if ret < 0 {
		fmt.Printf("Failed to initialize the resampling context\n")
		os.Exit(1)
	}
}

/* Prepare a 16 bit dummy audio frame of 'frame_size' samples and
 * 'nb_channels' channels. */
func get_audio_frame(ost *OutputStream) *libavutil.AVFrame {
	frame := ost.tmp_frame
	var j, i, v ffcommon.FInt
	q := (*ffcommon.FInt16T)(unsafe.Pointer(frame.Data[0]))

	// /* check if we want to generate more frames */
	if libavutil.AvCompareTs(ost.next_pts, ost.enc.TimeBase,
		STREAM_DURATION, libavutil.AVRational{1, 1}) > 0 {
		return nil
	}

	for j = 0; j < frame.NbSamples; j++ {
		v = ffcommon.FInt(math.Sin(float64(ost.t)) * 10000)
		for i = 0; i < ost.enc.Channels; i++ {
			*q = ffcommon.FInt16T(v)
			q = (*ffcommon.FInt16T)(unsafe.Pointer(uintptr(unsafe.Pointer(q)) + 2))
		}
		ost.t += ost.tincr
		ost.tincr += ost.tincr2
	}

	frame.Pts = ost.next_pts
	ost.next_pts += int64(frame.NbSamples)

	return frame
}

/*
 * encode one audio frame and send it to the muxer
 * return 1 when encoding is finished, 0 otherwise
 */
func write_audio_frame(oc *libavformat.AVFormatContext, ost *OutputStream) ffcommon.FInt {
	var c *libavcodec.AVCodecContext
	var frame *libavutil.AVFrame
	var ret ffcommon.FInt
	var dst_nb_samples ffcommon.FInt

	c = ost.enc

	frame = get_audio_frame(ost)

	if frame != nil {
		/* convert samples from native format to destination codec format, using the resampler */
		/* compute destination number of samples */
		dst_nb_samples = int32(libavutil.AvRescaleRnd(ost.swr_ctx.SwrGetDelay(int64(c.SampleRate))+int64(frame.NbSamples),
			int64(c.SampleRate), int64(c.SampleRate), libavutil.AV_ROUND_UP))
		//     av_assert0(dst_nb_samples == frame->nb_samples);

		/* when we pass a frame to the encoder, it may keep a reference to it
		 * internally;
		 * make sure we do not overwrite it here
		 */
		ret = ost.frame.AvFrameMakeWritable()
		if ret < 0 {
			os.Exit(1)
		}

		/* convert to destination format */
		ret = ost.swr_ctx.SwrConvert((**byte)(unsafe.Pointer(&ost.frame.Data)), dst_nb_samples,
			(**byte)(unsafe.Pointer(&frame.Data)), frame.NbSamples)
		if ret < 0 {
			fmt.Printf("Error while converting\n")
			os.Exit(1)
		}
		frame = ost.frame

		frame.Pts = libavutil.AvRescaleQ(int64(ost.samples_count), libavutil.AVRational{1, c.SampleRate}, c.TimeBase)
		ost.samples_count += dst_nb_samples
	}

	return write_frame(oc, c, ost.st, frame)
}

// /**************************************************************/
// /* video output */

func alloc_picture(pix_fmt libavutil.AVPixelFormat, width, height ffcommon.FInt) *libavutil.AVFrame {
	var picture *libavutil.AVFrame
	var ret ffcommon.FInt

	picture = libavutil.AvFrameAlloc()
	if picture == nil {
		return nil
	}

	picture.Format = pix_fmt
	picture.Width = width
	picture.Height = height

	// /* allocate the buffers for the frame data */
	ret = picture.AvFrameGetBuffer(0)
	if ret < 0 {
		fmt.Printf("Could not allocate frame data.\n")
		os.Exit(1)
	}

	return picture
}

func open_video(oc *libavformat.AVFormatContext, codec *libavcodec.AVCodec, ost *OutputStream, opt_arg *libavutil.AVDictionary) {
	var ret ffcommon.FInt
	c := ost.enc
	var opt *libavutil.AVDictionary

	libavutil.AvDictCopy(&opt, opt_arg, 0)

	/* open the codec */
	ret = c.AvcodecOpen2(codec, &opt)
	libavutil.AvDictFree(&opt)
	if ret < 0 {
		fmt.Printf("Could not open video codec: %s\n", libavutil.AvErr2str(ret))
		os.Exit(1)
	}

	/* allocate and init a re-usable frame */
	ost.frame = alloc_picture(c.PixFmt, c.Width, c.Height)
	if ost.frame == nil {
		fmt.Printf("Could not allocate video frame\n")
		os.Exit(1)
	}

	/* If the output format is not YUV420P, then a temporary YUV420P
	 * picture is needed too. It is then converted to the required
	 * output format. */
	ost.tmp_frame = nil
	if c.PixFmt != libavutil.AV_PIX_FMT_YUV420P {
		ost.tmp_frame = alloc_picture(libavutil.AV_PIX_FMT_YUV420P, c.Width, c.Height)
		if ost.tmp_frame == nil {
			fmt.Printf("Could not allocate temporary picture\n")
			os.Exit(1)
		}
	}

	/* copy the stream parameters to the muxer */
	ret = ost.st.Codecpar.AvcodecParametersFromContext(c)
	if ret < 0 {
		fmt.Printf("Could not copy the stream parameters\n")
		os.Exit(1)
	}
}

/* Prepare a dummy image. */
func fill_yuv_image(pict *libavutil.AVFrame, frame_index,
	width, height ffcommon.FInt) {
	var x, y, i ffcommon.FInt

	i = frame_index

	/* Y */
	for y = 0; y < height; y++ {
		for x = 0; x < width; x++ {
			*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(pict.Data[0])) + uintptr(y*pict.Linesize[0]+x))) = byte((x + y + i*3) % 256)
		}
	}

	// /* Cb and Cr */
	for y = 0; y < height/2; y++ {
		for x = 0; x < width/2; x++ {
			*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(pict.Data[1])) + uintptr(y*pict.Linesize[1]+x))) = byte((128 + y + i*2) % 256)
			*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(pict.Data[2])) + uintptr(y*pict.Linesize[2]+x))) = byte((64 + x + i*5) % 256)
		}
	}
}

func get_video_frame(ost *OutputStream) *libavutil.AVFrame {
	c := ost.enc

	/* check if we want to generate more frames */
	if libavutil.AvCompareTs(ost.next_pts, c.TimeBase,
		STREAM_DURATION, libavutil.AVRational{1, 1}) > 0 {
		return nil
	}

	/* when we pass a frame to the encoder, it may keep a reference to it
	 * internally; make sure we do not overwrite it here */
	if ost.frame.AvFrameMakeWritable() < 0 {
		os.Exit(1)
	}

	if c.PixFmt != libavutil.AV_PIX_FMT_YUV420P {
		/* as we only generate a YUV420P picture, we must convert it
		 * to the codec pixel format if needed */
		if ost.sws_ctx == nil {
			ost.sws_ctx = libswscale.SwsGetContext(c.Width, c.Height,
				libavutil.AV_PIX_FMT_YUV420P,
				c.Width, c.Height,
				c.PixFmt,
				SCALE_FLAGS, nil, nil, nil)
			if ost.sws_ctx == nil {
				fmt.Printf("Could not initialize the conversion context\n")
				os.Exit(1)
			}
		}
		fill_yuv_image(ost.tmp_frame, int32(ost.next_pts), c.Width, c.Height)
		ost.sws_ctx.SwsScale((**byte)(unsafe.Pointer(&ost.tmp_frame.Data)),
			(*int32)(unsafe.Pointer(&ost.tmp_frame.Linesize)), 0, uint32(c.Height), (**byte)(unsafe.Pointer(&ost.frame.Data)),
			(*int32)(unsafe.Pointer(&ost.frame.Linesize)))
	} else {
		fill_yuv_image(ost.frame, int32(ost.next_pts), c.Width, c.Height)
	}

	ost.frame.Pts = ost.next_pts
	ost.next_pts++

	return ost.frame
}

/*
 * encode one video frame and send it to the muxer
 * return 1 when encoding is finished, 0 otherwise
 */
func write_video_frame(oc *libavformat.AVFormatContext, ost *OutputStream) ffcommon.FInt {
	return write_frame(oc, ost.enc, ost.st, get_video_frame(ost))
}

func close_stream(oc *libavformat.AVFormatContext, ost *OutputStream) {
	libavcodec.AvcodecFreeContext(&ost.enc)
	libavutil.AvFrameFree(&ost.frame)
	libavutil.AvFrameFree(&ost.tmp_frame)
	ost.sws_ctx.SwsFreeContext()
	libswresample.SwrFree(&ost.swr_ctx)
}
