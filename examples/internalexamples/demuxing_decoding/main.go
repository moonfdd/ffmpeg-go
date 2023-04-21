package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/ffmpeg-go/libavcodec"
	"github.com/moonfdd/ffmpeg-go/libavformat"
	"github.com/moonfdd/ffmpeg-go/libavutil"
)

// go run ./examples/internalexamples/demuxing_decoding/main.go ./resources/big_buck_bunny.mp4 ./out/big_buck_bunny.yuv ./out/big_buck_bunny.pcm
// ./lib/ffplay -f rawvideo -pix_fmt yuv420p -video_size 640x360 ./out/big_buck_bunny.yuv
// ./lib/ffplay -f f32le -ac 1 -ar 22050 .\out\big_buck_bunny.pcm

func main0() (ret ffcommon.FInt) {

	if len(os.Args) != 4 {
		fmt.Printf("usage: %s  input_file video_output_file audio_output_file\nAPI example program to show how to read frames from an input file.\nThis program reads frames from a file, decodes them, and writes decoded\nvideo frames to a rawvideo file named video_output_file, and decoded\naudio frames to a rawaudio file named audio_output_file.\n",
			os.Args[0])
		os.Exit(1)
	}
	src_filename = os.Args[1]
	video_dst_filename = os.Args[2]
	audio_dst_filename = os.Args[3]

	/* open input file, and allocate format context */
	if libavformat.AvformatOpenInput(&fmt_ctx, src_filename, nil, nil) < 0 {
		fmt.Printf("Could not open source file %s\n", src_filename)
		os.Exit(1)
	}

	/* retrieve stream information */
	if fmt_ctx.AvformatFindStreamInfo(nil) < 0 {
		fmt.Printf("Could not find stream information\n")
		os.Exit(1)
	}

	for {
		if open_codec_context(&video_stream_idx, &video_dec_ctx, fmt_ctx, libavutil.AVMEDIA_TYPE_VIDEO) >= 0 {
			video_stream = fmt_ctx.GetStream(uint32(video_stream_idx))

			video_dst_file, _ = os.Create(video_dst_filename)
			if video_dst_file == nil {
				fmt.Printf("Could not open destination file %s\n", video_dst_filename)
				ret = 1
				break
			}

			/* allocate image where the decoded image will be put */
			width = video_dec_ctx.Width
			height = video_dec_ctx.Height
			pix_fmt = video_dec_ctx.PixFmt
			ret = libavutil.AvImageAlloc(&video_dst_data, &video_dst_linesize,
				width, height, pix_fmt, 1)
			if ret < 0 {
				fmt.Printf("Could not allocate raw video buffer\n")
				break
			}
			video_dst_bufsize = ret
		}

		if open_codec_context(&audio_stream_idx, &audio_dec_ctx, fmt_ctx, libavutil.AVMEDIA_TYPE_AUDIO) >= 0 {
			audio_stream = fmt_ctx.GetStream(uint32(audio_stream_idx))
			audio_dst_file, _ = os.Create(audio_dst_filename)
			if audio_dst_file == nil {
				fmt.Printf("Could not open destination file %s\n", audio_dst_filename)
				ret = 1
				break
			}
		}

		/* dump input information to stderr */
		fmt_ctx.AvDumpFormat(0, src_filename, 0)

		if audio_stream == nil && video_stream == nil {
			fmt.Printf("Could not find audio or video stream in the input, aborting\n")
			ret = 1
			break
		}

		frame = libavutil.AvFrameAlloc()
		if frame == nil {
			fmt.Printf("Could not allocate frame\n")
			ret = -libavutil.ENOMEM
			break
		}

		pkt = libavcodec.AvPacketAlloc()
		if pkt == nil {
			fmt.Printf("Could not allocate packet\n")
			ret = -libavutil.ENOMEM
			break
		}

		if video_stream != nil {
			fmt.Printf("Demuxing video from file '%s' into '%s'\n", src_filename, video_dst_filename)
		}
		if audio_stream != nil {
			fmt.Printf("Demuxing audio from file '%s' into '%s'\n", src_filename, audio_dst_filename)
		}

		/* read frames from the file */
		for fmt_ctx.AvReadFrame(pkt) >= 0 {
			// check if the packet belongs to a stream we are interested in, otherwise
			// skip it
			if pkt.StreamIndex == uint32(video_stream_idx) {
				ret = decode_packet(video_dec_ctx, pkt)
			} else if pkt.StreamIndex == uint32(audio_stream_idx) {
				ret = decode_packet(audio_dec_ctx, pkt)
			}
			pkt.AvPacketUnref()
			if ret < 0 {
				break
			}
		}

		/* flush the decoders */
		if video_dec_ctx != nil {
			decode_packet(video_dec_ctx, nil)
		}
		if audio_dec_ctx != nil {
			decode_packet(audio_dec_ctx, nil)
		}

		fmt.Printf("Demuxing succeeded.\n")

		if video_stream != nil {
			fmt.Printf("Play the output video file with the command:\nffplay -f rawvideo -pix_fmt %s -video_size %dx%d %s\n",
				libavutil.AvGetPixFmtName(pix_fmt), width, height,
				video_dst_filename)
		}

		if audio_stream != nil {
			sfmt := audio_dec_ctx.SampleFmt
			n_channels := audio_dec_ctx.Channels
			var fmt0 string

			if libavutil.AvSampleFmtIsPlanar(sfmt) != 0 {
				packed := libavutil.AvGetSampleFmtName(sfmt)
				if packed == "" {
					packed = "?"
				}
				fmt.Printf("Warning: the sample format the decoder produced is planar (%s). This example will output the first channel only.\n",
					packed)
				sfmt = libavutil.AvGetPackedSampleFmt(sfmt)
				n_channels = 1
			}

			ret = get_format_from_sample_fmt(&fmt0, sfmt)
			if ret < 0 {
				break
			}

			fmt.Printf("Play the output audio file with the command:\nffplay -f %s -ac %d -ar %d %s\n",
				fmt0, n_channels, audio_dec_ctx.SampleRate,
				audio_dst_filename)
		}
		break
	}
	// end:
	libavcodec.AvcodecFreeContext(&video_dec_ctx)
	libavcodec.AvcodecFreeContext(&audio_dec_ctx)
	libavformat.AvformatCloseInput(&fmt_ctx)
	if video_dst_file != nil {
		video_dst_file.Close()
	}
	if audio_dst_file != nil {
		audio_dst_file.Close()
	}
	libavcodec.AvPacketFree(&pkt)
	libavutil.AvFrameFree(&frame)
	libavutil.AvFree(uintptr(unsafe.Pointer(video_dst_data[0])))

	if ret < 0 {
		return 1
	} else {
		return 0
	}
}

var fmt_ctx *libavformat.AVFormatContext
var video_dec_ctx, audio_dec_ctx *libavcodec.AVCodecContext
var width, height ffcommon.FInt
var pix_fmt libavutil.AVPixelFormat
var video_stream, audio_stream *libavformat.AVStream
var src_filename string
var video_dst_filename string
var audio_dst_filename string
var video_dst_file *os.File
var audio_dst_file *os.File
var video_dst_data [4]*ffcommon.FUint8T
var video_dst_linesize [4]ffcommon.FInt
var video_dst_bufsize ffcommon.FInt
var video_stream_idx, audio_stream_idx ffcommon.FInt = -1, -1
var frame *libavutil.AVFrame
var pkt *libavcodec.AVPacket
var video_frame_count ffcommon.FInt
var audio_frame_count ffcommon.FInt

func output_video_frame(frame *libavutil.AVFrame) ffcommon.FInt {
	if frame.Width != width || frame.Height != height ||
		frame.Format != pix_fmt {
		/* To handle this change, one could call av_image_alloc again and
		 * decode the following frames into another rawvideo file. */
		fmt.Printf("Error: Width, height and pixel format have to be constant in a rawvideo file, but the width, height or pixel format of the input video changed:\nold: width = %d, height = %d, format = %s\nnew: width = %d, height = %d, format = %s\n",
			width, height, libavutil.AvGetPixFmtName(pix_fmt),
			frame.Width, frame.Height,
			libavutil.AvGetPixFmtName(frame.Format))
		return -1
	}

	fmt.Printf("video_frame n:%d coded_n:%d\n",
		video_frame_count, frame.CodedPictureNumber)
	video_frame_count++

	/* copy decoded frame to destination buffer:
	 * this is required since rawvideo expects non aligned data */
	libavutil.AvImageCopy(&video_dst_data, &video_dst_linesize,
		(*[4]*byte)(unsafe.Pointer(&frame.Data)), (*[4]int32)(unsafe.Pointer(&frame.Linesize)),
		pix_fmt, width, height)

	/* write to rawvideo file */
	video_dst_file.Write(ffcommon.ByteSliceFromByteP(video_dst_data[0], int(video_dst_bufsize)))
	return 0
}

func output_audio_frame(frame *libavutil.AVFrame) ffcommon.FInt {
	unpadded_linesize := ffcommon.FSizeT(frame.NbSamples * libavutil.AvGetBytesPerSample(libavutil.AVSampleFormat(frame.Format)))
	fmt.Printf("audio_frame n:%d nb_samples:%d pts:%s\n",
		audio_frame_count, frame.NbSamples,
		libavutil.AvTs2timestr(frame.Pts, &audio_dec_ctx.TimeBase))
	audio_frame_count++
	/* Write the raw audio data samples of the first plane. This works
	 * fine for packed formats (e.g. AV_SAMPLE_FMT_S16). However,
	 * most audio decoders output planar audio, which uses a separate
	 * plane of audio samples for each channel (e.g. AV_SAMPLE_FMT_S16P).
	 * In other words, this code will write only the first audio channel
	 * in these cases.
	 * You should use libswresample or libavfilter to convert the frame
	 * to packed data. */
	audio_dst_file.Write(ffcommon.ByteSliceFromByteP(*frame.ExtendedData, int(unpadded_linesize)))

	return 0
}

func decode_packet(dec *libavcodec.AVCodecContext, pkt *libavcodec.AVPacket) ffcommon.FInt {
	ret := ffcommon.FInt(0)

	// submit the packet to the decoder
	ret = dec.AvcodecSendPacket(pkt)
	if ret < 0 {
		fmt.Printf("Error submitting a packet for decoding (%s)\n", libavutil.AvErr2str(ret))
		return ret
	}

	// get all the available frames from the decoder
	for ret >= 0 {
		ret = dec.AvcodecReceiveFrame(frame)
		if ret < 0 {
			// those two return values are special and mean there is no output
			// frame available, but there were no errors during decoding
			if ret == libavutil.AVERROR_EOF || ret == -libavutil.EAGAIN {
				return 0
			}

			fmt.Printf("Error during decoding (%s)%d\n", libavutil.AvErr2str(ret), ret)
			return ret
		}

		// write the frame data to output file
		if dec.Codec.Type == libavutil.AVMEDIA_TYPE_VIDEO {
			ret = output_video_frame(frame)
		} else {
			ret = output_audio_frame(frame)
		}

		frame.AvFrameUnref()
		if ret < 0 {
			return ret
		}
	}

	return 0
}

func open_codec_context(stream_idx *ffcommon.FInt,
	dec_ctx **libavcodec.AVCodecContext, fmt_ctx *libavformat.AVFormatContext, type0 libavutil.AVMediaType) ffcommon.FInt {
	var ret, stream_index ffcommon.FInt
	var st *libavformat.AVStream
	var dec *libavcodec.AVCodec
	var opts *libavutil.AVDictionary

	ret = fmt_ctx.AvFindBestStream(type0, -1, -1, nil, 0)
	if ret < 0 {
		fmt.Printf("Could not find %s stream in input file '%s'\n",
			libavutil.AvGetMediaTypeString(type0), src_filename)
		return ret
	} else {
		stream_index = ret
		st = fmt_ctx.GetStream(uint32(stream_index))

		/* find decoder for the stream */
		dec = libavcodec.AvcodecFindDecoder(st.Codecpar.CodecId)
		if dec == nil {
			fmt.Printf("Failed to find %s codec\n",
				libavutil.AvGetMediaTypeString(type0))
			return -libavutil.EINVAL
		}

		/* Allocate a codec context for the decoder */
		*dec_ctx = dec.AvcodecAllocContext3()
		if *dec_ctx == nil {
			fmt.Printf("Failed to allocate the %s codec context\n",
				libavutil.AvGetMediaTypeString(type0))
			return -libavutil.ENOMEM
		}

		/* Copy codec parameters from input stream to output codec context */
		ret = (*dec_ctx).AvcodecParametersToContext(st.Codecpar)
		if ret < 0 {
			fmt.Printf("Failed to copy %s codec parameters to decoder context\n",
				libavutil.AvGetMediaTypeString(type0))
			return ret
		}

		/* Init the decoders */
		ret = (*dec_ctx).AvcodecOpen2(dec, &opts)
		if ret < 0 {
			fmt.Printf("Failed to open %s codec\n",
				libavutil.AvGetMediaTypeString(type0))
			return ret
		}
		*stream_idx = stream_index
	}

	return 0
}

func get_format_from_sample_fmt(fmt0 *string, sample_fmt libavutil.AVSampleFormat) (ret ffcommon.FInt) {
	switch sample_fmt {
	case libavutil.AV_SAMPLE_FMT_U8:
		*fmt0 = "u8"
	case libavutil.AV_SAMPLE_FMT_S16:
		*fmt0 = "s16le"
	case libavutil.AV_SAMPLE_FMT_S32:
		*fmt0 = "s32le"
	case libavutil.AV_SAMPLE_FMT_FLT:
		*fmt0 = "f32le"
	case libavutil.AV_SAMPLE_FMT_DBL:
		*fmt0 = "f64le"
	default:
		fmt.Printf("sample format %s is not supported as output format\n",
			libavutil.AvGetSampleFmtName(sample_fmt))
		ret = -1
	}
	return
}

func main() {
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
