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

func main() {
	os.Setenv("Path", os.Getenv("Path")+";./lib")
	ffcommon.SetAvutilPath("./lib/avutil-56.dll")
	ffcommon.SetAvcodecPath("./lib/avcodec-58.dll")
	ffcommon.SetAvdevicePath("./lib/avdevice-56.dll")
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

	//是否存在yuv文件
	_, err = os.Stat("./out/result.yuv")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("create yuv file")
			exec.Command("./lib/ffmpeg", "-i", "./resources/big_buck_bunny.mp4", "-pix_fmt", "yuv420p", "./out/result.yuv", "-y").CombinedOutput()
		}
	}

	ret := int32(0) //默认返回值
	//需要的变量名并初始化
	var fmtCtx *libavformat.AVFormatContext
	var outFmt *libavformat.AVOutputFormat
	var vStream *libavformat.AVStream
	pkt := libavcodec.AvPacketAlloc()
	var codecCtx *libavcodec.AVCodecContext
	var codec *libavcodec.AVCodec
	var picture_buf uintptr
	var picFrame *libavutil.AVFrame
	var size ffcommon.FSizeT

	//[1]!打开视频文件
	in_file, err := os.Open("./out/result.yuv")
	if err != nil {
		fmt.Printf("can not open file!\n")
		return
	}
	defer in_file.Close()

	for {
		//[2]!打开输出文件，并填充fmtCtx数据
		in_w := int32(640)
		in_h := int32(360)
		frameCnt := 1440
		outFile := "./out/result.h264"
		os.Remove(outFile)
		if libavformat.AvformatAllocOutputContext2(&fmtCtx, nil, "", outFile) < 0 {
			fmt.Printf("Cannot alloc output file context.\n")
			break
		}
		outFmt = fmtCtx.Oformat
		//[2]!

		//[3]!打开输出文件
		if libavformat.AvioOpen(&fmtCtx.Pb, outFile, libavformat.AVIO_FLAG_READ_WRITE) < 0 {
			fmt.Printf("output file open failed.\n")
			break
		}

		//[3]!

		//[4]!创建h264视频流，并设置参数
		vStream = fmtCtx.AvformatNewStream(codec)
		if vStream == nil {
			fmt.Printf("failed create new video stream.\n")
			break
		}
		vStream.TimeBase.Den = 25
		vStream.TimeBase.Num = 1
		//[4]!

		//[5]!编码参数相关
		codecPara := fmtCtx.GetStream(uint32(vStream.Index)).Codecpar
		codecPara.CodecType = libavutil.AVMEDIA_TYPE_VIDEO
		codecPara.Width = in_w
		codecPara.Height = in_h
		//[5]!

		//[6]!查找编码器
		codec = libavcodec.AvcodecFindEncoder(outFmt.VideoCodec)
		if codec == nil {
			fmt.Printf("Cannot find any endcoder.\n")
			break
		}
		//[6]!

		//[7]!设置编码器内容
		codecCtx = codec.AvcodecAllocContext3()
		codecCtx.AvcodecParametersToContext(codecPara)
		if codecCtx == nil {
			fmt.Printf("Cannot alloc context.")
			break
		}
		codecCtx.CodecId = outFmt.VideoCodec
		codecCtx.CodecType = libavutil.AVMEDIA_TYPE_VIDEO
		codecCtx.PixFmt = libavutil.AV_PIX_FMT_YUV420P
		codecCtx.Width = in_w
		codecCtx.Height = in_h
		codecCtx.TimeBase.Num = 1
		codecCtx.TimeBase.Den = 25
		codecCtx.BitRate = 400000
		codecCtx.GopSize = 12

		if codecCtx.CodecId == libavcodec.AV_CODEC_ID_H264 {
			codecCtx.Qmin = 10
			codecCtx.Qmax = 51
			codecCtx.Qcompress = 0.6
		}
		if codecCtx.CodecId == libavcodec.AV_CODEC_ID_MPEG2VIDEO {
			codecCtx.MaxBFrames = 2
		}
		if codecCtx.CodecId == libavcodec.AV_CODEC_ID_MPEG1VIDEO {
			codecCtx.MbDecision = 2
		}
		//[7]!

		//[8]!打开编码器
		if codecCtx.AvcodecOpen2(codec, nil) < 0 {
			fmt.Printf("Open encoder failed.\n")
			break
		}
		//[8]!

		fmtCtx.AvDumpFormat(0, outFile, 1) //输出 输出文件流信息

		//初始化帧
		picFrame = libavutil.AvFrameAlloc()
		picFrame.Width = codecCtx.Width
		picFrame.Height = codecCtx.Height
		picFrame.Format = codecCtx.PixFmt
		size = uint64(libavutil.AvImageGetBufferSize(codecCtx.PixFmt, codecCtx.Width, codecCtx.Height, 1))
		picture_buf = libavutil.AvMalloc(size)
		libavutil.AvImageFillArrays((*[4]*byte)(unsafe.Pointer(&picFrame.Data)), (*[4]int32)(unsafe.Pointer(&picFrame.Linesize)), (*byte)(unsafe.Pointer(picture_buf)), codecCtx.PixFmt,
			codecCtx.Width, codecCtx.Height, 1)

		//[9] --写头文件
		ret = fmtCtx.AvformatWriteHeader(nil)
		//[9]

		y_size := codecCtx.Width * codecCtx.Height
		pkt.AvNewPacket(int32(size * 3))
		buf := make([]byte, size)
		//picture_buf = uintptr(unsafe.Pointer(&buf[0]))

		//[10] --循环编码每一帧
		for i := 0; i < frameCnt; i++ {
			//读入YUV
			n, err := in_file.Read(buf)
			if err != nil {
				fmt.Println("read end")
				break
			}
			if n <= 0 {
				break
			}

			for i := 0; i < n; i++ {
				*(*byte)(unsafe.Pointer(picture_buf + uintptr(i))) = buf[i]
			}

			picFrame.Data[0] = (*byte)(unsafe.Pointer(picture_buf))                       //亮度Y
			picFrame.Data[1] = (*byte)(unsafe.Pointer(picture_buf + uintptr(y_size)))     // U
			picFrame.Data[2] = (*byte)(unsafe.Pointer(picture_buf + uintptr(y_size*5/4))) // V
			// AVFrame PTS
			picFrame.Pts = int64(i)

			//编码
			if codecCtx.AvcodecSendFrame(picFrame) >= 0 {
				for codecCtx.AvcodecReceivePacket(pkt) >= 0 {
					fmt.Printf("encoder success!\n")
					// parpare packet for muxing
					pkt.StreamIndex = uint32(vStream.Index)
					pkt.AvPacketRescaleTs(codecCtx.TimeBase, vStream.TimeBase)
					pkt.Pos = -1
					ret = fmtCtx.AvInterleavedWriteFrame(pkt)
					if ret < 0 {
						fmt.Printf("error is: %s.\n", libavutil.AvErr2str(ret))
					}
					pkt.AvPacketUnref() //刷新缓存
				}
			}
		}
		//[10]

		//[11] --Flush encoder
		ret = flush_encoder(fmtCtx, codecCtx, vStream.Index)
		if ret < 0 {
			fmt.Printf("flushing encoder failed!\n")
			break
		}
		//[11]

		//[12] --写文件尾
		fmtCtx.AvWriteTrailer()
		//[12]

		break
	}

	// //===========================释放所有指针===============================//
	libavcodec.AvPacketFree(&pkt)
	codecCtx.AvcodecClose()
	libavutil.AvFree(uintptr(unsafe.Pointer(picFrame)))
	libavutil.AvFree(picture_buf)

	if fmtCtx != nil {
		fmtCtx.Pb.AvioClose()
		fmtCtx.AvformatFreeContext()
	}

	_, err = exec.Command("./lib/ffplay.exe", "./out/result.h264").Output()
	if err != nil {
		fmt.Println("play err = ", err)
	}
}

//刷新编码器
func flush_encoder(fmtCtx *libavformat.AVFormatContext, codecCtx *libavcodec.AVCodecContext, vStreamIndex int32) int32 {
	var ret int32
	enc_pkt := libavcodec.AvPacketAlloc()
	enc_pkt.Data = nil
	enc_pkt.Size = 0

	if codecCtx.Codec.Capabilities&libavcodec.AV_CODEC_CAP_DELAY == 0 {
		return 0
	}

	fmt.Printf("Flushing stream #%d encoder\n", vStreamIndex)
	ret = codecCtx.AvcodecSendFrame(nil)
	if ret >= 0 {
		for codecCtx.AvcodecReceivePacket(enc_pkt) >= 0 {
			fmt.Printf("success encoder 1 frame.\n")

			// parpare packet for muxing
			enc_pkt.StreamIndex = uint32(vStreamIndex)
			enc_pkt.AvPacketRescaleTs(codecCtx.TimeBase, fmtCtx.GetStream(uint32(vStreamIndex)).TimeBase)
			ret = fmtCtx.AvInterleavedWriteFrame(enc_pkt)
			if ret < 0 {
				break
			}
		}
	}

	return ret
}
