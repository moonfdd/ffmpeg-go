package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/ffmpeg-go/libavutil"
	"github.com/moonfdd/ffmpeg-go/libswscale"
)

func main0() (ret ffcommon.FInt) {
	var src_data, dst_data [4]*ffcommon.FUint8T
	var src_linesize, dst_linesize [4]ffcommon.FInt
	var src_w ffcommon.FInt = 320
	var src_h ffcommon.FInt = 240
	var dst_w ffcommon.FInt
	var dst_h ffcommon.FInt
	var src_pix_fmt libavutil.AVPixelFormat = libavutil.AV_PIX_FMT_YUV420P
	var dst_pix_fmt libavutil.AVPixelFormat = libavutil.AV_PIX_FMT_RGB24
	var dst_size string
	var dst_filename string
	var dst_file *os.File
	var dst_bufsize ffcommon.FInt
	var sws_ctx *libswscale.SwsContext
	var i ffcommon.FInt

	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s output_file output_size\nAPI example program to show how to scale an image with libswscale.\nThis program generates a series of pictures, rescales them to the given output_size and saves them to an output file named output_file\n.\n", os.Args[0])
		os.Exit(1)
	}
	dst_filename = os.Args[1]
	dst_size = os.Args[2]

	if libavutil.AvParseVideoSize(&dst_w, &dst_h, dst_size) < 0 {
		fmt.Printf("Invalid size '%s', must be in the form WxH or a valid size abbreviation\n",
			dst_size)
		os.Exit(1)
	}

	dst_file, _ = os.Create(dst_filename)
	if dst_file == nil {
		fmt.Printf("Could not open destination file %s\n", dst_filename)
		os.Exit(1)
	}

	/* create scaling context */
	sws_ctx = libswscale.SwsGetContext(src_w, src_h, src_pix_fmt,
		dst_w, dst_h, dst_pix_fmt,
		libswscale.SWS_BILINEAR, nil, nil, nil)
	if sws_ctx == nil {
		fmt.Printf(
			"Impossible to create scale context for the conversion fmt:%s s:%dx%d -> fmt:%s s:%dx%d\n",
			libavutil.AvGetPixFmtName(src_pix_fmt), src_w, src_h,
			libavutil.AvGetPixFmtName(dst_pix_fmt), dst_w, dst_h)
		ret = -libavutil.EINVAL
		goto end
	}

	/* allocate source and destination image buffers */
	ret = libavutil.AvImageAlloc(&src_data, &src_linesize,
		src_w, src_h, src_pix_fmt, 16)
	if ret < 0 {
		fmt.Printf("Could not allocate source image\n")
		goto end
	}

	/* buffer is going to be written to rawvideo file, no alignment */
	ret = libavutil.AvImageAlloc(&dst_data, &dst_linesize,
		dst_w, dst_h, dst_pix_fmt, 1)
	if ret < 0 {
		fmt.Printf("Could not allocate destination image\n")
		goto end
	}
	dst_bufsize = ret

	for i = 0; i < 100; i++ {
		// /* generate synthetic video */
		fill_yuv_image(&src_data, &src_linesize, src_w, src_h, i)

		/* convert to destination format */
		sws_ctx.SwsScale((**byte)(unsafe.Pointer(&src_data)),
			(*int32)(unsafe.Pointer(&src_linesize)), 0, uint32(src_h), (**byte)(unsafe.Pointer(&dst_data)), (*int32)(unsafe.Pointer(&dst_linesize)))

		// /* write scaled image to file */
		dst_file.Write(ffcommon.ByteSliceFromByteP(dst_data[0], int(dst_bufsize)))
	}

	fmt.Printf("Scaling succeeded. Play the output file with the command:\nffplay -f rawvideo -pix_fmt %s -video_size %dx%d %s\n",
		libavutil.AvGetPixFmtName(dst_pix_fmt), dst_w, dst_h, dst_filename)

end:
	dst_file.Close()
	libavutil.AvFreep(uintptr(unsafe.Pointer(&src_data[0])))
	libavutil.AvFreep(uintptr(unsafe.Pointer(&dst_data[0])))
	sws_ctx.SwsFreeContext()
	if ret < 0 {
		return 1
	} else {
		return 0
	}
}

func fill_yuv_image(data *[4]*ffcommon.FUint8T, linesize *[4]ffcommon.FInt, width, height, frame_index ffcommon.FInt) {
	var x, y ffcommon.FInt

	/* Y */
	for y = 0; y < height; y++ {
		for x = 0; x < width; x++ {
			//data[0][y*linesize[0]+x] = x + y + frame_index*3
			*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(data[0])) + uintptr(y*linesize[0]+x))) = byte((x + y + frame_index*3) % 256)
		}
	}

	/* Cb and Cr */
	for y = 0; y < height/2; y++ {
		for x = 0; x < width/2; x++ {
			// data[1][y * linesize[1] + x] = 128 + y + frame_index * 2;
			// data[2][y * linesize[2] + x] = 64 + x + frame_index * 5;
			*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(data[1])) + uintptr(y*linesize[1]+x))) = byte((128 + y + frame_index*2) % 256)
			*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(data[2])) + uintptr(y*linesize[2]+x))) = byte((64 + x + frame_index*5) % 256)
		}
	}
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
