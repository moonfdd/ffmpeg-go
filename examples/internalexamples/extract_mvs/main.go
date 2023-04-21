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

func main0() (ret ffcommon.FInt) {
	var pkt libavformat.AVPacket

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <input video>\n", os.Args[0])
		os.Exit(1)
	}
	src_filename = os.Args[1]

	if libavformat.AvformatOpenInput(&fmt_ctx, src_filename, nil, nil) < 0 {
		fmt.Printf("Could not open source file %s\n", src_filename)
		os.Exit(1)
	}

	if fmt_ctx.AvformatFindStreamInfo(nil) < 0 {
		fmt.Printf("Could not find stream information\n")
		os.Exit(1)
	}

	open_codec_context(fmt_ctx, libavutil.AVMEDIA_TYPE_VIDEO)

	fmt_ctx.AvDumpFormat(0, src_filename, 0)
	for {
		if video_stream == nil {
			fmt.Printf("Could not find video stream in the input, aborting\n")
			ret = 1
			break
		}

		frame = libavutil.AvFrameAlloc()
		if frame == nil {
			fmt.Printf("Could not allocate frame\n")
			ret = -libavutil.ENOMEM
			break
		}

		fmt.Printf("framenum,source,blockw,blockh,srcx,srcy,dstx,dsty,flags\n")

		/* read frames from the file */
		for fmt_ctx.AvReadFrame(&pkt) >= 0 {
			if pkt.StreamIndex == uint32(video_stream_idx) {
				ret = decode_packet(&pkt)
			}
			pkt.AvPacketUnref()
			if ret < 0 {
				break
			}
		}

		/* flush cached frames */
		decode_packet(nil)
		break
	}
	// end:
	libavcodec.AvcodecFreeContext(&video_dec_ctx)
	libavformat.AvformatCloseInput(&fmt_ctx)
	libavutil.AvFrameFree(&frame)
	if ret < 0 {
		return 1
	} else {
		return 0
	}
}

var fmt_ctx *libavformat.AVFormatContext
var video_dec_ctx *libavcodec.AVCodecContext
var video_stream *libavformat.AVStream
var src_filename string

var video_stream_idx ffcommon.FInt = -1
var frame *libavutil.AVFrame
var video_frame_count ffcommon.FInt

func decode_packet(pkt *libavcodec.AVPacket) ffcommon.FInt {
	ret := video_dec_ctx.AvcodecSendPacket(pkt)
	if ret < 0 {
		fmt.Printf("Error while sending a packet to the decoder: %s\n", libavutil.AvErr2str(ret))
		return ret
	}

	for ret >= 0 {
		ret = video_dec_ctx.AvcodecReceiveFrame(frame)
		if ret == -libavutil.EAGAIN || ret == libavutil.AVERROR_EOF {
			break
		} else if ret < 0 {
			fmt.Printf("Error while receiving a frame from the decoder: %s\n", libavutil.AvErr2str(ret))
			return ret
		}

		if ret >= 0 {
			var i ffcommon.FInt
			var sd *libavutil.AVFrameSideData

			video_frame_count++
			sd = frame.AvFrameGetSideData(libavutil.AV_FRAME_DATA_MOTION_VECTORS)
			if sd != nil {
				//const AVMotionVector
				// mvs := (*libavutil.AVMotionVector)(unsafe.Pointer(sd.Data))
				var a [2]libavutil.AVMotionVector
				len0 := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
				for i = 0; i < sd.Size/int32(len0); i++ {
					mv := (*libavutil.AVMotionVector)(unsafe.Pointer(uintptr(unsafe.Pointer(sd.Data)) + len0*uintptr(i)))
					fmt.Printf("%d,%2d,%2d,%2d,%4d,%4d,%4d,%4d,0x%d\n",
						video_frame_count, mv.Source,
						mv.W, mv.H, mv.SrcX, mv.SrcY,
						mv.DstX, mv.DstY, mv.Flags)
				}
			}
			frame.AvFrameUnref()
		}
	}

	return 0
}

func open_codec_context(fmt_ctx *libavformat.AVFormatContext, type0 libavutil.AVMediaType) ffcommon.FInt {
	var ret ffcommon.FInt
	var st *libavformat.AVStream
	var dec_ctx *libavcodec.AVCodecContext
	var dec *libavcodec.AVCodec
	var opts *libavutil.AVDictionary

	ret = fmt_ctx.AvFindBestStream(type0, -1, -1, &dec, 0)
	if ret < 0 {
		fmt.Printf("Could not find %s stream in input file '%s'\n",
			libavutil.AvGetMediaTypeString(type0), src_filename)
		return ret
	} else {
		stream_idx := ret
		st = fmt_ctx.GetStream(uint32(stream_idx))

		dec_ctx = dec.AvcodecAllocContext3()
		if dec_ctx == nil {
			fmt.Printf("Failed to allocate codec\n")
			return -libavutil.EINVAL
		}

		ret = dec_ctx.AvcodecParametersToContext(st.Codecpar)
		if ret < 0 {
			fmt.Printf("Failed to copy codec parameters to codec context\n")
			return ret
		}

		/* Init the video decoder */
		libavutil.AvDictSet(&opts, "flags2", "+export_mvs", 0)
		ret = dec_ctx.AvcodecOpen2(dec, &opts)
		if ret < 0 {
			fmt.Printf("Failed to open %s codec\n",
				libavutil.AvGetMediaTypeString(type0))
			return ret
		}

		video_stream_idx = stream_idx
		video_stream = fmt_ctx.GetStream(uint32(video_stream_idx))
		video_dec_ctx = dec_ctx
	}

	return 0
}

func main() {
	os.Setenv("Path", os.Getenv("Path")+";./lib")
	ffcommon.SetAvutilPath("./lib/avutil-56.dll")
	ffcommon.SetAvcodecPath("./lib/avcodec-58.dll")
	ffcommon.SetAvdevicePath("./lib/avdevice-58.dll")
	ffcommon.SetAvfilterPath("./lib/avfilter-7.dll")
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
