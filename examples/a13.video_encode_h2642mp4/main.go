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

	frame_index := 0 //统计帧数
	inVStreamIndex := -1
	outVStreamIndex := -1 //输入输出视频流在文件中的索引位置
	inVFileName := "./out/result.h264"
	outFileName := "./out/result.mp4"

	//是否存在h264文件
	_, err = os.Stat(inVFileName)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("create h264 file")
			exec.Command("./lib/ffmpeg", "-i", "./resources/big_buck_bunny.mp4", "-vcodec", "copy", "-an", inVFileName, "-y").CombinedOutput()
		}
	}

	var inVFmtCtx *libavformat.AVFormatContext
	var outFmtCtx *libavformat.AVFormatContext
	var codecPara *libavcodec.AVCodecParameters
	var outVStream *libavformat.AVStream
	var outCodec *libavcodec.AVCodec
	var outCodecCtx *libavcodec.AVCodecContext
	var outCodecPara *libavcodec.AVCodecParameters
	var inVStream *libavformat.AVStream
	pkt := libavcodec.AvPacketAlloc()

	for {
		//======================输入部分============================//
		//打开输入文件
		if libavformat.AvformatOpenInput(&inVFmtCtx, inVFileName, nil, nil) < 0 {
			fmt.Printf("Cannot open input file.\n")
			break
		}

		//查找输入文件中的流
		if inVFmtCtx.AvformatFindStreamInfo(nil) < 0 {
			fmt.Printf("Cannot find stream info in input file.\n")
			break
		}

		//查找视频流在文件中的位置
		for i := uint32(0); i < inVFmtCtx.NbStreams; i++ {
			if inVFmtCtx.GetStream(i).Codecpar.CodecType == libavutil.AVMEDIA_TYPE_VIDEO {
				inVStreamIndex = int(i)
				break
			}
		}

		codecPara = inVFmtCtx.GetStream(uint32(inVStreamIndex)).Codecpar //输入视频流的编码参数

		fmt.Printf("===============Input information========>\n")
		inVFmtCtx.AvDumpFormat(0, inVFileName, 0)
		fmt.Printf("===============Input information========<\n")

		//=====================输出部分=========================//
		//打开输出文件并填充格式数据
		if libavformat.AvformatAllocOutputContext2(&outFmtCtx, nil, "", outFileName) < 0 {
			fmt.Printf("Cannot alloc output file context.\n")
			break
		}

		//打开输出文件并填充数据
		if libavformat.AvioOpen(&outFmtCtx.Pb, outFileName, libavformat.AVIO_FLAG_READ_WRITE) < 0 {
			fmt.Printf("output file open failed.\n")
			break
		}

		//在输出的mp4文件中创建一条视频流
		outVStream = outFmtCtx.AvformatNewStream(nil)
		if outVStream == nil {
			fmt.Printf("Failed allocating output stream.\n")
			break
		}
		outVStream.TimeBase.Den = 25
		outVStream.TimeBase.Num = 1
		outVStreamIndex = int(outVStream.Index)

		//查找编码器
		outCodec = libavcodec.AvcodecFindEncoder(codecPara.CodecId)
		if outCodec == nil {
			fmt.Printf("Cannot find any encoder.\n")
			break
		}

		//从输入的h264编码器数据复制一份到输出文件的编码器中
		outCodecCtx = outCodec.AvcodecAllocContext3()
		outCodecPara = outFmtCtx.GetStream(uint32(outVStream.Index)).Codecpar
		if libavcodec.AvcodecParametersCopy(outCodecPara, codecPara) < 0 {
			fmt.Printf("Cannot copy codec para.\n")
			break
		}
		if outCodecCtx.AvcodecParametersToContext(outCodecPara) < 0 {
			fmt.Printf("Cannot alloc codec ctx from para.\n")
			break
		}
		outCodecCtx.TimeBase.Den = 25
		outCodecCtx.TimeBase.Num = 1

		//打开输出文件需要的编码器
		if outCodecCtx.AvcodecOpen2(outCodec, nil) < 0 {
			fmt.Printf("Cannot open output codec.\n")
			break
		}

		fmt.Printf("============Output Information=============>\n")
		outFmtCtx.AvDumpFormat(0, outFileName, 1)
		fmt.Printf("============Output Information=============<\n")

		//写入文件头
		if outFmtCtx.AvformatWriteHeader(nil) < 0 {
			fmt.Printf("Cannot write header to file.\n")
			return
		}
		//===============编码部分===============//

		inVStream = inVFmtCtx.GetStream(uint32(inVStreamIndex))
		for inVFmtCtx.AvReadFrame(pkt) >= 0 { //循环读取每一帧直到读完
			if pkt.StreamIndex == uint32(inVStreamIndex) { //确保处理的是视频流
				//FIXME：No PTS (Example: Raw H.264)
				//Simple Write PTS
				//如果当前处理帧的显示时间戳为0或者没有等等不是正常值
				if pkt.Pts == libavutil.AV_NOPTS_VALUE {
					fmt.Printf("frame_index:%d\n", frame_index)
					//Write PTS
					time_base1 := inVStream.TimeBase
					//Duration between 2 frames (us)
					calc_duration := libavutil.AV_TIME_BASE / libavutil.AvQ2d(inVStream.RFrameRate)

					//Parameters
					pkt.Pts = int64((float64(frame_index) * calc_duration) / (libavutil.AvQ2d(time_base1) * float64(libavutil.AV_TIME_BASE)))
					pkt.Dts = pkt.Pts
					pkt.Duration = int64(calc_duration / (libavutil.AvQ2d(time_base1) * float64(libavutil.AV_TIME_BASE)))
					frame_index++
				}
				//Convert PTS/DTS
				pkt.Pts = libavutil.AvRescaleQRnd(pkt.Pts, inVStream.TimeBase, outVStream.TimeBase, libavutil.AV_ROUND_NEAR_INF|libavutil.AV_ROUND_PASS_MINMAX)
				pkt.Dts = libavutil.AvRescaleQRnd(pkt.Dts, inVStream.TimeBase, outVStream.TimeBase, libavutil.AV_ROUND_NEAR_INF|libavutil.AV_ROUND_PASS_MINMAX)
				pkt.Duration = libavutil.AvRescaleQ(pkt.Duration, inVStream.TimeBase, outVStream.TimeBase)
				pkt.Pos = -1
				pkt.StreamIndex = uint32(outVStreamIndex)
				fmt.Printf("Write 1 Packet. size:%5d\tpts:%d\tduration:%d\n", pkt.Size, pkt.Pts, pkt.Duration)

				//Write
				if outFmtCtx.AvInterleavedWriteFrame(pkt) < 0 {
					fmt.Printf("Error muxing packet\n")
					break
				}
				pkt.AvPacketUnref()
			}
		}
		outFmtCtx.AvWriteTrailer()
		break
	}

	//===========================释放所有指针===============================//
	libavcodec.AvPacketFree(&pkt)
	libavformat.AvformatCloseInput(&outFmtCtx)
	outCodecCtx.AvcodecClose()
	libavcodec.AvcodecFreeContext(&outCodecCtx)
	libavformat.AvformatCloseInput(&inVFmtCtx)
	inVFmtCtx.AvformatFreeContext()
	// outFmtCtx.Pb.AvioClose()//案例里面有，但个人感觉不对

	fmt.Println("-----------------------------------------")
	_, err = exec.Command("./lib/ffplay.exe", "./out/result.mp4").Output()
	if err != nil {
		fmt.Println("play err = ", err)
	}
}
