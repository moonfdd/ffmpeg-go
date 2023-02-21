package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/ffmpeg-go/libavcodec"
	"github.com/moonfdd/ffmpeg-go/libavdevice"
	"github.com/moonfdd/ffmpeg-go/libavformat"
	"github.com/moonfdd/ffmpeg-go/libavutil"
	"github.com/moonfdd/ffmpeg-go/libswscale"
)

func main() {
	// ./lib/ffmpeg -list_devices true -f dshow -i dummy
	// ./lib/ffplay -f dshow -i video="Full HD webcam"
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

	// frame_index := 0 //统计帧数
	// inVStreamIndex := -1
	// outVStreamIndex := -1 //输入输出视频流在文件中的索引位置
	// inVFileName := "./out/result.h264"
	// outFileName := "./out/result.mp4"

	//是否存在h264文件
	// _, err = os.Stat(inVFileName)
	// if err != nil {
	// 	if os.IsNotExist(err) {
	// 		fmt.Println("create h264 file")
	// 		exec.Command("./lib/ffmpeg", "-i", "./resources/big_buck_bunny.mp4", "-vcodec", "copy", "-an", inVFileName, "-y").CombinedOutput()
	// 	}
	// }
	ret := int32(0)
	libavdevice.AvdeviceRegisterAll()
	inFmtCtx := libavformat.AvformatAllocContext()
	var inCodecCtx *libavcodec.AVCodecContext
	var inCodec *libavcodec.AVCodec
	inPkt := libavcodec.AvPacketAlloc()
	srcFrame := libavutil.AvFrameAlloc()
	yuvFrame := libavutil.AvFrameAlloc()

	//打开输出文件，并填充fmtCtx数据
	outFmtCtx := libavformat.AvformatAllocContext()
	var outFmt *libavformat.AVOutputFormat
	var outCodecCtx *libavcodec.AVCodecContext
	var outCodec *libavcodec.AVCodec
	var outVStream *libavformat.AVStream
	outPkt := libavcodec.AvPacketAlloc()
	var img_ctx *libswscale.SwsContext
	inVideoStreamIndex := -1

	for {
		/////////////解码器部分//////////////////////
		//打开摄像头
		inFmt := libavformat.AvFindInputFormat("dshow")
		if libavformat.AvformatOpenInput(&inFmtCtx, "video=Full HD webcam", inFmt, nil) < 0 {
			fmt.Printf("Cannot open camera.\n")
			return
		}
		if inFmtCtx.AvformatFindStreamInfo(nil) < 0 {
			fmt.Printf("Cannot find any stream in file.\n")
			return
		}

		for i := 0; i < int(inFmtCtx.NbStreams); i++ {
			if inFmtCtx.GetStream(uint32(i)).Codecpar.CodecType == libavutil.AVMEDIA_TYPE_VIDEO {
				inVideoStreamIndex = i
				break
			}
		}

		if inVideoStreamIndex == -1 {
			fmt.Printf("Cannot find video stream in file.\n")
			return
		}

		inVideoCodecPara := inFmtCtx.GetStream(uint32(inVideoStreamIndex)).Codecpar
		inCodec = libavcodec.AvcodecFindDecoder(inVideoCodecPara.CodecId)
		if inCodec == nil {
			fmt.Printf("Cannot find valid video decoder.\n")
			return
		}
		inCodecCtx = inCodec.AvcodecAllocContext3()
		if inCodecCtx == nil {
			fmt.Printf("Cannot alloc valid decode codec context.\n")
			return
		}
		if inCodecCtx.AvcodecParametersToContext(inVideoCodecPara) < 0 {
			fmt.Printf("Cannot initialize parameters.\n")
			return
		}

		if inCodecCtx.AvcodecOpen2(inCodec, nil) < 0 {
			fmt.Printf("Cannot open codec.\n")
			return
		}

		img_ctx = libswscale.SwsGetContext(inCodecCtx.Width,
			inCodecCtx.Height,
			inCodecCtx.PixFmt,
			inCodecCtx.Width,
			inCodecCtx.Height,
			libavutil.AV_PIX_FMT_YUV420P,
			libswscale.SWS_BICUBIC,
			nil, nil, nil)

		numBytes := libavutil.AvImageGetBufferSize(libavutil.AV_PIX_FMT_YUV420P,
			inCodecCtx.Width,
			inCodecCtx.Height, 1)

		out_buffer := libavutil.AvMalloc(uint64(numBytes))

		ret = libavutil.AvImageFillArrays((*[4]*ffcommon.FUint8T)(unsafe.Pointer(&yuvFrame.Data)),
			(*[4]ffcommon.FInt)(unsafe.Pointer(&yuvFrame.Linesize)),
			(*ffcommon.FUint8T)(unsafe.Pointer(out_buffer)),
			libavutil.AV_PIX_FMT_YUV420P,
			inCodecCtx.Width,
			inCodecCtx.Height,
			1)
		if ret < 0 {
			fmt.Printf("Fill arrays failed.\n")
			return
		}
		//////////////解码器部分结束/////////////////////

		//////////////编码器部分开始/////////////////////
		outFile := "./out/result14.h264"
		if libavformat.AvformatAllocOutputContext2(&outFmtCtx, nil, "", outFile) < 0 {
			fmt.Printf("Cannot alloc output file context.\n")
			return
		}
		outFmt = outFmtCtx.Oformat

		//打开输出文件
		if libavformat.AvioOpen(&outFmtCtx.Pb, outFile, libavformat.AVIO_FLAG_READ_WRITE) < 0 {
			fmt.Printf("output file open failed.\n")
			return
		}

		//创建h264视频流，并设置参数
		outVStream = outFmtCtx.AvformatNewStream(outCodec)
		if outVStream == nil {
			fmt.Printf("create new video stream fialed.\n")
			return
		}
		outVStream.TimeBase.Den = 30
		outVStream.TimeBase.Num = 1

		//编码参数相关
		outCodecPara := outFmtCtx.GetStream(uint32(outVStream.Index)).Codecpar
		outCodecPara.CodecType = libavutil.AVMEDIA_TYPE_VIDEO
		outCodecPara.CodecId = outFmt.VideoCodec
		outCodecPara.Width = 480
		outCodecPara.Height = 360
		outCodecPara.BitRate = 110000

		//查找编码器
		outCodec = libavcodec.AvcodecFindEncoder(outFmt.VideoCodec)
		if outCodec == nil {
			fmt.Printf("Cannot find any encoder.\n")
			return
		}

		//设置编码器内容
		outCodecCtx = outCodec.AvcodecAllocContext3()
		outCodecCtx.AvcodecParametersToContext(outCodecPara)
		if outCodecCtx == nil {
			fmt.Printf("Cannot alloc output codec content.\n")
			return
		}
		outCodecCtx.CodecId = outFmt.VideoCodec
		outCodecCtx.CodecType = libavutil.AVMEDIA_TYPE_VIDEO
		outCodecCtx.PixFmt = libavutil.AV_PIX_FMT_YUV420P
		outCodecCtx.Width = inCodecCtx.Width
		outCodecCtx.Height = inCodecCtx.Height
		outCodecCtx.TimeBase.Num = 1
		outCodecCtx.TimeBase.Den = 30
		outCodecCtx.BitRate = 110000
		outCodecCtx.GopSize = 10

		if outCodecCtx.CodecId == libavcodec.AV_CODEC_ID_H264 {
			outCodecCtx.Qmin = 10
			outCodecCtx.Qmax = 51
			outCodecCtx.Qcompress = 0.6
		} else if outCodecCtx.CodecId == libavcodec.AV_CODEC_ID_MPEG2VIDEO {
			outCodecCtx.MaxBFrames = 2
		} else if outCodecCtx.CodecId == libavcodec.AV_CODEC_ID_MPEG1VIDEO {
			outCodecCtx.MbDecision = 2
		}

		//打开编码器
		if outCodecCtx.AvcodecOpen2(outCodec, nil) < 0 {
			fmt.Printf("Open encoder failed.\n")
			return
		}
		///////////////编码器部分结束////////////////////

		///////////////编解码部分//////////////////////
		yuvFrame.Format = outCodecCtx.PixFmt
		yuvFrame.Width = outCodecCtx.Width
		yuvFrame.Height = outCodecCtx.Height

		ret = outFmtCtx.AvformatWriteHeader(nil)

		count := 0
		for inFmtCtx.AvReadFrame(inPkt) >= 0 && count < 50 {
			if inPkt.StreamIndex == uint32(inVideoStreamIndex) {
				if inCodecCtx.AvcodecSendPacket(inPkt) >= 0 {
					ret = inCodecCtx.AvcodecReceiveFrame(srcFrame)
					for ret >= 0 {
						if ret == -libavutil.EAGAIN || ret == libavutil.AVERROR_EOF {
							break
						} else if ret < 0 {
							fmt.Printf("Error during decoding\n")
							return
						}
						img_ctx.SwsScale((**byte)(unsafe.Pointer(&srcFrame.Data)),
							(*int32)(unsafe.Pointer(&srcFrame.Linesize)),
							0, uint32(inCodecCtx.Height),
							(**byte)(unsafe.Pointer(&yuvFrame.Data)), (*int32)(unsafe.Pointer(&yuvFrame.Linesize)))

						yuvFrame.Pts = srcFrame.Pts
						//encode
						if outCodecCtx.AvcodecSendFrame(yuvFrame) >= 0 {
							if outCodecCtx.AvcodecReceivePacket(outPkt) >= 0 {
								fmt.Printf("encode one frame.\n")
								count++
								outPkt.StreamIndex = uint32(outVStream.Index)
								outPkt.AvPacketRescaleTs(outCodecCtx.TimeBase,
									outVStream.TimeBase)
								outPkt.Pos = -1
								outFmtCtx.AvInterleavedWriteFrame(outPkt)
								outPkt.AvPacketUnref()
							}
						}
						// usleep(1000*24);
						time.Sleep(time.Millisecond * 24)
						ret = inCodecCtx.AvcodecReceiveFrame(srcFrame)
					}
				}
				inPkt.AvPacketUnref()
			}
		}
		ret = flush_encoder(outFmtCtx, outCodecCtx, int(outVStream.Index))
		if ret < 0 {
			fmt.Printf("flushing encoder failed.\n")
			return
		}
		outFmtCtx.AvWriteTrailer()
		////////////////编解码部分结束////////////////
		break
	}
	///////////内存释放部分/////////////////////////
	libavcodec.AvPacketFree(&inPkt)
	libavcodec.AvcodecFreeContext(&inCodecCtx)
	inCodecCtx.AvcodecClose()
	libavformat.AvformatCloseInput(&inFmtCtx)
	libavutil.AvFrameFree(&srcFrame)
	libavutil.AvFrameFree(&yuvFrame)

	libavcodec.AvPacketFree(&outPkt)
	libavcodec.AvcodecFreeContext(&outCodecCtx)
	outCodecCtx.AvcodecClose()
	libavformat.AvformatCloseInput(&outFmtCtx)

	fmt.Println("-----------------------------------------")
	_, err = exec.Command("./lib/ffplay.exe", "./out/result14.h264").Output()
	if err != nil {
		fmt.Println("play err = ", err)
	}
}

func flush_encoder(fmtCtx *libavformat.AVFormatContext, codecCtx *libavcodec.AVCodecContext, vStreamIndex int) int32 {
	ret := int32(0)
	enc_pkt := libavcodec.AvPacketAlloc()
	enc_pkt.Data = nil
	enc_pkt.Size = 0

	if (codecCtx.Codec.Capabilities & libavcodec.AV_CODEC_CAP_DELAY) == 0 {
		return 0
	}

	fmt.Printf("Flushing stream #%d encoder\n", vStreamIndex)
	if codecCtx.AvcodecSendFrame(nil) >= 0 {
		for codecCtx.AvcodecReceivePacket(enc_pkt) >= 0 {
			fmt.Printf("success encoder 1 frame.\n")

			// parpare packet for muxing
			enc_pkt.StreamIndex = uint32(vStreamIndex)
			enc_pkt.AvPacketRescaleTs(codecCtx.TimeBase,
				fmtCtx.GetStream(uint32(vStreamIndex)).TimeBase)
			ret = fmtCtx.AvInterleavedWriteFrame(enc_pkt)
			if ret < 0 {
				break
			}
		}
	}

	enc_pkt.AvPacketUnref()

	return ret
}
