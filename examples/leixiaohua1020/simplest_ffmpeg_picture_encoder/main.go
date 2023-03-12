// https://github.com/leixiaohua1020/simplest_ffmpeg_picture_encoder/blob/master/simplest_ffmpeg_picture_encoder/simplest_ffmpeg_picture_encoder.cpp
package main

import (
	"fmt"
	"os"
	"os/exec"
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/ffmpeg-go/libavcodec"
	"github.com/moonfdd/ffmpeg-go/libavformat"
	"github.com/moonfdd/ffmpeg-go/libavutil"
)

func main0() (ret ffcommon.FInt) {
	var pFormatCtx *libavformat.AVFormatContext
	var fmt0 *libavformat.AVOutputFormat
	var video_st *libavformat.AVStream
	var pCodecCtx *libavcodec.AVCodecContext
	var pCodec *libavcodec.AVCodec

	var picture_buf *ffcommon.FUint8T
	var picture *libavutil.AVFrame
	var pkt libavcodec.AVPacket
	var y_size ffcommon.FInt
	var got_picture ffcommon.FInt = 0
	var size ffcommon.FInt

	var in_file *os.File                    //YUV source
	var in_w, in_h ffcommon.FInt = 640, 360 //YUV's width and height
	var out_file = "./out/pic.jpg"          //Output file
	in := "./out/pic.yuv"

	//是否存在yuv文件
	_, err := os.Stat(in)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("create yuv file")
			exec.Command("./lib/ffmpeg", "-i", "./resources/big_buck_bunny.mp4", "-pix_fmt", "yuv420p", in, "-y").CombinedOutput()
		}
	}

	in_file, _ = os.Open(in)
	if in_file == nil {
		return -1
	}

	libavformat.AvRegisterAll()

	//Method 1
	pFormatCtx = libavformat.AvformatAllocContext()
	//Guess format
	fmt0 = libavformat.AvGuessFormat("mjpeg", "", "")
	pFormatCtx.Oformat = fmt0
	//Output URL
	if libavformat.AvioOpen(&pFormatCtx.Pb, out_file, libavformat.AVIO_FLAG_READ_WRITE) < 0 {
		fmt.Printf("Couldn't open output file.")
		return -1
	}

	//Method 2. More simple
	//avformat_alloc_output_context2(&pFormatCtx, NULL, NULL, out_file);
	//fmt = pFormatCtx->oformat;

	video_st = pFormatCtx.AvformatNewStream(nil)
	if video_st == nil {
		return -1
	}
	pCodecCtx = video_st.Codec
	pCodecCtx.CodecId = fmt0.VideoCodec
	pCodecCtx.CodecType = libavutil.AVMEDIA_TYPE_VIDEO
	pCodecCtx.PixFmt = libavutil.AV_PIX_FMT_YUVJ420P

	pCodecCtx.Width = in_w
	pCodecCtx.Height = in_h

	pCodecCtx.TimeBase.Num = 1
	pCodecCtx.TimeBase.Den = 25
	//Output some information
	pFormatCtx.AvDumpFormat(0, out_file, 1)

	pCodec = libavcodec.AvcodecFindEncoder(pCodecCtx.CodecId)
	if pCodec == nil {
		fmt.Printf("Codec not found.")
		return -1
	}
	if pCodecCtx.AvcodecOpen2(pCodec, nil) < 0 {
		fmt.Printf("Could not open codec.")
		return -1
	}
	picture = libavutil.AvFrameAlloc()
	picture.Width = pCodecCtx.Width
	picture.Height = pCodecCtx.Height
	picture.Format = pCodecCtx.PixFmt
	size = libavcodec.AvpictureGetSize(pCodecCtx.PixFmt, pCodecCtx.Width, pCodecCtx.Height)
	picture_buf = (*byte)(unsafe.Pointer(libavutil.AvMalloc(uint64(size))))
	if picture_buf == nil {
		return -1
	}
	((*libavcodec.AVPicture)(unsafe.Pointer(picture))).AvpictureFill(picture_buf, pCodecCtx.PixFmt, pCodecCtx.Width, pCodecCtx.Height)

	//Write Header
	pFormatCtx.AvformatWriteHeader(nil)

	y_size = pCodecCtx.Width * pCodecCtx.Height
	pkt.AvNewPacket(y_size * 3)
	//Read YUV
	_, err = in_file.Read(ffcommon.ByteSliceFromByteP(picture_buf, int(y_size*3/2)))
	if err != nil {
		fmt.Printf("Could not read input file.%s", err)
		return -1
	}
	picture.Data[0] = picture_buf                                                                         // Y
	picture.Data[1] = (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(picture_buf)) + uintptr(y_size)))     // U
	picture.Data[2] = (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(picture_buf)) + uintptr(y_size*5/4))) // V

	//Encode
	ret = pCodecCtx.AvcodecEncodeVideo2(&pkt, picture, &got_picture)
	if ret < 0 {
		fmt.Printf("Encode Error.\n")
		return -1
	}
	if got_picture == 1 {
		pkt.StreamIndex = uint32(video_st.Index)
		ret = pFormatCtx.AvWriteFrame(&pkt)
	}

	pkt.AvFreePacket()
	//Write Trailer
	pFormatCtx.AvWriteTrailer()

	fmt.Printf("Encode Successful.\n")

	if video_st != nil {
		video_st.Codec.AvcodecClose()
		libavutil.AvFree(uintptr(unsafe.Pointer(picture)))
		libavutil.AvFree(uintptr(unsafe.Pointer(picture_buf)))
	}
	pFormatCtx.Pb.AvioClose()
	pFormatCtx.AvformatFreeContext()

	in_file.Close()

	exec.Command("./lib/ffplay.exe", out_file).Output()
	if err != nil {
		fmt.Println("play err = ", err)
	}

	return 0
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

	// go func() {
	// 	time.Sleep(1000)
	// 	exec.Command("./lib/ffplay.exe", "rtmp://localhost/publishlive/livestream").Output()
	// 	if err != nil {
	// 		fmt.Println("play err = ", err)
	// 	}
	// }()

	main0()
}
