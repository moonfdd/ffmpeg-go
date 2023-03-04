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
	var input_ctx *libavformat.AVFormatContext
	var video_stream ffcommon.FInt
	var video *libavformat.AVStream
	var decoder_ctx *libavcodec.AVCodecContext
	var decoder *libavcodec.AVCodec
	var packet libavformat.AVPacket
	var type0 libavutil.AVHWDeviceType
	var i ffcommon.FInt

	if len(os.Args) < 4 {
		fmt.Printf("Usage: %s <device type> <input file> <output file>\n", os.Args[0])
		return -1
	}

	type0 = libavutil.AvHwdeviceFindTypeByName(os.Args[1])
	if type0 == libavutil.AV_HWDEVICE_TYPE_NONE {
		fmt.Printf("Device type %s is not supported.\n", os.Args[1])
		fmt.Printf("Available device types:")
		type0 = libavutil.AvHwdeviceIterateTypes(type0)
		for type0 != libavutil.AV_HWDEVICE_TYPE_NONE {
			fmt.Printf(" %s", libavutil.AvHwdeviceGetTypeName(type0))
			type0 = libavutil.AvHwdeviceIterateTypes(type0)
		}
		fmt.Printf("\n")
		return -1
	}

	/* open the input file */
	if libavformat.AvformatOpenInput(&input_ctx, os.Args[2], nil, nil) != 0 {
		fmt.Printf("Cannot open input file '%s'\n", os.Args[2])
		return -1
	}

	if input_ctx.AvformatFindStreamInfo(nil) < 0 {
		fmt.Printf("Cannot find input stream information.\n")
		return -1
	}

	/* find the video stream information */
	ret = input_ctx.AvFindBestStream(libavutil.AVMEDIA_TYPE_VIDEO, -1, -1, &decoder, 0)
	if ret < 0 {
		fmt.Printf("Cannot find a video stream in the input file\n")
		return -1
	}
	video_stream = ret

	for i = 0; ; i++ {
		config := decoder.AvcodecGetHwConfig(i)
		if config == nil {
			fmt.Printf("Decoder %s does not support device type %s.\n",
				ffcommon.StringFromPtr(decoder.Name), libavutil.AvHwdeviceGetTypeName(type0))
			return -1
		}
		if config.Methods&libavcodec.AV_CODEC_HW_CONFIG_METHOD_HW_DEVICE_CTX != 0 && config.DeviceType == type0 {
			hw_pix_fmt = config.PixFmt
			break
		}
	}

	decoder_ctx = decoder.AvcodecAllocContext3()
	if decoder_ctx == nil {
		return -libavutil.ENOMEM
	}

	video = input_ctx.GetStream(uint32(video_stream))
	if decoder_ctx.AvcodecParametersToContext(video.Codecpar) < 0 {
		return -1
	}

	decoder_ctx.GetFormat = ffcommon.NewCallback(get_hw_format)

	if hw_decoder_init(decoder_ctx, type0) < 0 {
		return -1
	}

	ret = decoder_ctx.AvcodecOpen2(decoder, nil)
	if ret < 0 {
		fmt.Printf("Failed to open codec for stream #%d\n", video_stream)
		return -1
	}

	/* open the file to dump raw data */
	output_file, _ = os.Create(os.Args[3])

	/* actual decoding and dump the raw data */
	for ret >= 0 {
		ret = input_ctx.AvReadFrame(&packet)
		if ret < 0 {
			break
		}

		if uint32(video_stream) == packet.StreamIndex {
			ret = decode_write(decoder_ctx, &packet)
		}

		packet.AvPacketUnref()
	}

	/* flush the decoder */
	packet.Data = nil
	packet.Size = 0
	ret = decode_write(decoder_ctx, &packet)
	packet.AvPacketUnref()

	if output_file != nil {
		output_file.Close()
	}
	libavcodec.AvcodecFreeContext(&decoder_ctx)
	libavformat.AvformatCloseInput(&input_ctx)
	libavutil.AvBufferUnref(&hw_device_ctx)

	return 0
}

var hw_device_ctx *libavutil.AVBufferRef
var hw_pix_fmt libavutil.AVPixelFormat
var output_file *os.File

func hw_decoder_init(ctx *libavcodec.AVCodecContext, type0 libavutil.AVHWDeviceType) ffcommon.FInt {
	var err ffcommon.FInt = 0

	err = libavutil.AvHwdeviceCtxCreate(&hw_device_ctx, type0, "", nil, 0)
	if err < 0 {
		fmt.Printf("Failed to create specified HW device.\n")
		return err
	}
	ctx.HwDeviceCtx = hw_device_ctx.AvBufferRef()

	return err
}

func get_hw_format(ctx *libavcodec.AVCodecContext, pix_fmts *libavutil.AVPixelFormat) uintptr {
	var p *libavutil.AVPixelFormat

	for p = pix_fmts; *p != -1; p = (*libavutil.AVPixelFormat)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(4))) {
		if *p == hw_pix_fmt {
			return uintptr(*p)
		}
	}

	fmt.Printf("Failed to get HW surface format.\n")
	r := libavutil.AVPixelFormat(libavutil.AV_PIX_FMT_NONE)
	return uintptr(r)
}

func decode_write(avctx *libavcodec.AVCodecContext, packet *libavcodec.AVPacket) ffcommon.FInt {
	var frame, sw_frame *libavutil.AVFrame
	var tmp_frame *libavutil.AVFrame
	var buffer *ffcommon.FUint8T
	var size ffcommon.FInt
	var ret ffcommon.FInt = 0
	var e error

	ret = avctx.AvcodecSendPacket(packet)
	if ret < 0 {
		fmt.Printf("Error during decoding\n")
		return ret
	}

	for {
		frame = libavutil.AvFrameAlloc()
		sw_frame = libavutil.AvFrameAlloc()
		if frame == nil || sw_frame == nil {
			fmt.Printf("Can not alloc frame\n")
			ret = -libavutil.ENOMEM
			goto fail
		}

		ret = avctx.AvcodecReceiveFrame(frame)
		if ret == -libavutil.EAGAIN || ret == libavutil.AVERROR_EOF {
			libavutil.AvFrameFree(&frame)
			libavutil.AvFrameFree(&sw_frame)
			return 0
		} else if ret < 0 {
			fmt.Printf("Error while decoding\n")
			goto fail
		}

		if frame.Format == hw_pix_fmt {
			/* retrieve data from GPU to CPU */
			ret = libavutil.AvHwframeTransferData(sw_frame, frame, 0)
			if ret < 0 {
				fmt.Printf("Error transferring the data to system memory\n")
				goto fail
			}
			tmp_frame = sw_frame
		} else {
			tmp_frame = frame
		}

		size = libavutil.AvImageGetBufferSize(tmp_frame.Format, tmp_frame.Width,
			tmp_frame.Height, 1)
		buffer = (*byte)(unsafe.Pointer(libavutil.AvMalloc(uint64(size))))
		if buffer == nil {
			fmt.Printf("Can not alloc buffer\n")
			ret = -libavutil.ENOMEM
			goto fail
		}
		ret = libavutil.AvImageCopyToBuffer(buffer, size,
			(*[4]*byte)(unsafe.Pointer(&tmp_frame.Data)),
			(*[4]int32)(unsafe.Pointer(&tmp_frame.Linesize)), tmp_frame.Format,
			tmp_frame.Width, tmp_frame.Height, 1)
		if ret < 0 {
			fmt.Printf("Can not copy image to buffer\n")
			goto fail
		}

		_, e = output_file.Write(ffcommon.ByteSliceFromByteP(buffer, int(size)))

		if e != nil {
			fmt.Printf("Failed to dump raw data.\n")
			goto fail
		}

	fail:
		libavutil.AvFrameFree(&frame)
		libavutil.AvFrameFree(&sw_frame)
		libavutil.AvFreep(uintptr(unsafe.Pointer(&buffer)))
		if ret < 0 {
			return ret
		}
	}
}

func main() {
	// go run .\examples\internalexamples\hw_decode\main.go cuda .\resources\big_buck_bunny.mp4 ./out/hw.yuv
	// ./lib/ffplay -pixel_format yuv420p -video_size 640x360 ./out/hw.yuv

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
