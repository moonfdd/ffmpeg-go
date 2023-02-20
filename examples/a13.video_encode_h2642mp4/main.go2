package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/ffmpeg-go/libavcodec"
	"github.com/moonfdd/ffmpeg-go/libavformat"
	"github.com/moonfdd/ffmpeg-go/libavutil"
)

func main() {
	// https://blog.51cto.com/fengyuzaitu/2467100
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

	videoIndex := -1
	pszFile := "./out/result.h264"
	// pszFile := "./resources/big_buck_bunny.mp4"
	pszRTMPURL := "./out/result.mp4"

	// //是否存在yuv文件
	// _, err = os.Stat("./out/result.yuv")
	// if err != nil {
	// 	if os.IsNotExist(err) {
	// 		fmt.Println("create yuv file")
	// 		exec.Command("./lib/ffmpeg", "-i", "./resources/big_buck_bunny.mp4", "-pix_fmt", "yuv420p", "./out/result.yuv", "-y").CombinedOutput()
	// 	}
	// }

	var pInputAVFormatContext *libavformat.AVFormatContext
	var pAVOutputFormat *libavformat.AVOutputFormat

	//======================输入部分============================//
	//打开输入文件
	if libavformat.AvformatOpenInput(&pInputAVFormatContext, pszFile, nil, nil) < 0 {
		fmt.Printf("Cannot open input file.\n")
		return
	}

	//查找输入文件中的流
	if pInputAVFormatContext.AvformatFindStreamInfo(nil) < 0 {
		fmt.Printf("Cannot find stream info in input file.\n")
		return
	}

	fmt.Printf("===============Input information========>\n")
	pInputAVFormatContext.AvDumpFormat(0, pszFile, 0)
	fmt.Printf("===============Input information========<\n")

	var pOutputAVFormatContext *libavformat.AVFormatContext

	//=====================输出部分=========================//
	//打开输出文件并填充格式数据
	if libavformat.AvformatAllocOutputContext2(&pOutputAVFormatContext, nil, "mp4", pszRTMPURL) < 0 {
		fmt.Printf("Cannot alloc output file context.\n")
		return
	}

	pAVOutputFormat = pOutputAVFormatContext.Oformat
	if pAVOutputFormat == nil {

	}

	//查找视频流在文件中的位置
	for i := uint32(0); i < pInputAVFormatContext.NbStreams; i++ {
		pInputAVStream := pInputAVFormatContext.GetStream(i)
		pOutputAVStream := pOutputAVFormatContext.AvformatNewStream(nil)
		if libavcodec.AvcodecParametersCopy(pOutputAVStream.Codecpar, pInputAVStream.Codecpar) < 0 {
			fmt.Println("AvcodecParametersCopy err")
			return
		}
		pOutputAVStream.Codecpar.CodecTag = 0
	}
	for i := uint32(0); i < pInputAVFormatContext.NbStreams; i++ {
		if pInputAVFormatContext.GetStream(i).Codecpar.CodecType == libavutil.AVMEDIA_TYPE_VIDEO {
			videoIndex = int(i)
			break
		}
	}
	pOutputAVFormatContext.AvDumpFormat(0, pszRTMPURL, 1)

	//打开输出文件并填充数据
	if libavformat.AvioOpen(&pOutputAVFormatContext.Pb, pszRTMPURL, libavformat.AVIO_FLAG_READ_WRITE) < 0 {
		fmt.Printf("output file open failed.\n")
		return
	}

	//写入文件头
	if pOutputAVFormatContext.AvformatWriteHeader(nil) < 0 {
		fmt.Printf("Cannot write header to file.\n")
		return
	}

	var pkt libavcodec.AVPacket
	// llStartTime := libavutil.AvGettime()
	llFrameIndex := 0
	for {
		var pInputStream *libavformat.AVStream
		var pOutputStream *libavformat.AVStream
		if pInputAVFormatContext.AvReadFrame(&pkt) < 0 {
			break
		}

		if pkt.Pts == libavutil.AV_NOPTS_VALUE {
			time_base1 := pInputAVFormatContext.GetStream(uint32(videoIndex)).TimeBase
			llCalcDuration := int64(libavutil.AV_TIME_BASE / libavutil.AvQ2d(pInputAVFormatContext.GetStream(uint32(videoIndex)).RFrameRate))
			pkt.Pts = int64(float64(llFrameIndex) * float64(llCalcDuration) / (libavutil.AvQ2d(time_base1) * libavutil.AV_TIME_BASE))
			pkt.Dts = pkt.Pts
			pkt.Duration = int64(float64(llCalcDuration) / (libavutil.AvQ2d(time_base1) * libavutil.AV_TIME_BASE))
		}
		// if pkt.StreamIndex == uint32(videoIndex) {
		// 	time_base := pInputAVFormatContext.GetStream(uint32(videoIndex)).TimeBase
		// 	ime_base_q := libavutil.AV_TIME_BASE_Q
		// 	pts_time := av_rescale_q(pkt.dts, time_base, time_base_q)
		// 	now_time := av_gettime() - llStartTime
		// }

		pInputStream = pInputAVFormatContext.GetStream(pkt.StreamIndex)
		pOutputStream = pOutputAVFormatContext.GetStream(pkt.StreamIndex)
		pkt.Pts = libavutil.AvRescaleQRnd(pkt.Pts, pInputStream.TimeBase, pOutputStream.TimeBase, libavutil.AV_ROUND_NEAR_INF|libavutil.AV_ROUND_PASS_MINMAX)
		pkt.Dts = libavutil.AvRescaleQRnd(pkt.Dts, pInputStream.TimeBase, pOutputStream.TimeBase, libavutil.AV_ROUND_NEAR_INF|libavutil.AV_ROUND_PASS_MINMAX)
		pkt.Duration = libavutil.AvRescaleQ(pkt.Duration, pInputStream.TimeBase, pOutputStream.TimeBase)
		fmt.Println("Pts = ", pkt.Pts)
		fmt.Println("Dts = ", pkt.Dts)
		fmt.Println("Duration = ", pkt.Duration)
		fmt.Println("--------------------")
		pkt.Pos = -1
		if pkt.StreamIndex == uint32(videoIndex) {
			llFrameIndex++
		}
		if pOutputAVFormatContext.AvInterleavedWriteFrame(&pkt) < 0 {
			fmt.Printf("发送数据包出错\n")
			break
		}
		pkt.AvPacketUnref()
	}
	pOutputAVFormatContext.AvWriteTrailer()
	if pOutputAVFormatContext.Oformat.Flags&libavformat.AVFMT_NOFILE == 0 {
		pOutputAVFormatContext.Pb.AvioClose()
	}
	pOutputAVFormatContext.AvformatFreeContext()
	libavformat.AvformatCloseInput(&pInputAVFormatContext)

	return

	_, err = exec.Command("./lib/ffplay.exe", "./out/result.h264").Output()
	if err != nil {
		fmt.Println("play err = ", err)
	}
}
