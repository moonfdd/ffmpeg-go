package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/ffmpeg-go/libavcodec"
	"github.com/moonfdd/ffmpeg-go/libavdevice"
	"github.com/moonfdd/ffmpeg-go/libavformat"
	"github.com/moonfdd/ffmpeg-go/libavutil"
)

func open_codec_context(streamIndex *ffcommon.FInt, ofmtCtx **libavformat.AVFormatContext, ifmtCtx *libavformat.AVFormatContext, type0 libavutil.AVMediaType) ffcommon.FInt {
	var outStream, inStream *libavformat.AVStream
	// int ret = -1, index = -1;
	var ret ffcommon.FInt = -1
	var index ffcommon.FInt = -1

	index = ifmtCtx.AvFindBestStream(type0, -1, -1, nil, 0)
	if index < 0 {
		fmt.Printf("can't find %s stream in input file\n", libavutil.AvGetMediaTypeString(type0))
		return ret
	}

	inStream = ifmtCtx.GetStream(uint32(index))

	outStream = (*ofmtCtx).AvformatNewStream(nil)
	if outStream == nil {
		fmt.Printf("failed to allocate output stream\n")
		return ret
	}

	ret = libavcodec.AvcodecParametersCopy(outStream.Codecpar, inStream.Codecpar)
	if ret < 0 {
		fmt.Printf("failed to copy codec parametes\n")
		return ret
	}

	outStream.Codecpar.CodecTag = 0

	*streamIndex = index

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

	inFileName := "./resources/big_buck_bunny.mp4"
	outFilenameAudio := "./out/a22.aac"
	outFilenameVideo := "./out/a22.h264"

	var ifmtCtx, ofmtCtxAudio, ofmtCtxVideo *libavformat.AVFormatContext
	var packet libavcodec.AVPacket

	var videoIndex ffcommon.FInt = -1
	var audioIndex ffcommon.FInt = -1
	var ret ffcommon.FInt = 0

	//注册设备
	libavdevice.AvdeviceRegisterAll()

	for {
		//打开输入流
		if libavformat.AvformatOpenInput(&ifmtCtx, inFileName, nil, nil) < 0 {
			fmt.Printf("Cannot open input file.\n")
			break
		}

		//获取流信息
		if ifmtCtx.AvformatFindStreamInfo(nil) < 0 {
			fmt.Printf("Cannot find stream info in input file.\n")
			break
		}

		//创建输出上下文:视频
		libavformat.AvformatAllocOutputContext2(&ofmtCtxVideo, nil, "", outFilenameVideo)
		if ofmtCtxVideo == nil {
			fmt.Printf("can't create video output context")
			break
		}

		//创建输出上下文：音频
		libavformat.AvformatAllocOutputContext2(&ofmtCtxAudio, nil, "", outFilenameAudio)
		if ofmtCtxAudio == nil {
			fmt.Printf("can't create audio output context")
			break
		}

		ret = open_codec_context(&videoIndex, &ofmtCtxVideo, ifmtCtx, libavutil.AVMEDIA_TYPE_VIDEO)
		if ret < 0 {
			fmt.Printf("can't decode video context\n")
			break
		}

		ret = open_codec_context(&audioIndex, &ofmtCtxAudio, ifmtCtx, libavutil.AVMEDIA_TYPE_AUDIO)
		if ret < 0 {
			fmt.Printf("can't decode video context\n")
			break
		}

		//Dump Format------------------
		fmt.Printf("\n==============Input Video=============\n")
		ifmtCtx.AvDumpFormat(0, inFileName, 0)
		fmt.Printf("\n==============Output Video============\n")
		ofmtCtxVideo.AvDumpFormat(0, outFilenameVideo, 1)
		fmt.Printf("\n==============Output Audio============\n")
		ofmtCtxAudio.AvDumpFormat(0, outFilenameAudio, 1)
		fmt.Printf("\n======================================\n")

		//打开输出文件:视频
		if ofmtCtxVideo.Oformat.Flags&libavformat.AVFMT_NOFILE == 0 {
			if libavformat.AvioOpen(&ofmtCtxVideo.Pb, outFilenameVideo, libavformat.AVIO_FLAG_WRITE) < 0 {
				fmt.Printf("can't open output file: %s\n", outFilenameVideo)
				break
			}
		}

		//打开输出文件：音频
		if ofmtCtxAudio.Oformat.Flags&libavformat.AVFMT_NOFILE == 0 {
			if libavformat.AvioOpen(&ofmtCtxAudio.Pb, outFilenameAudio, libavformat.AVIO_FLAG_WRITE) < 0 {
				fmt.Printf("can't open output file: %s\n", outFilenameVideo)
				break
			}
		}

		//写文件头
		if ofmtCtxVideo.AvformatWriteHeader(nil) < 0 {
			fmt.Printf("Error occurred when opening video output file\n")
			break
		}

		if ofmtCtxAudio.AvformatWriteHeader(nil) < 0 {
			fmt.Printf("Error occurred when opening audio output file\n")
			break
		}

		for {
			var ofmtCtx *libavformat.AVFormatContext
			var inStream, outStream *libavformat.AVStream

			if ifmtCtx.AvReadFrame(&packet) < 0 {
				break
			}

			inStream = ifmtCtx.GetStream(packet.StreamIndex)

			if packet.StreamIndex == uint32(videoIndex) {
				outStream = ofmtCtxVideo.GetStream(0)
				ofmtCtx = ofmtCtxVideo
			} else if packet.StreamIndex == uint32(audioIndex) {
				outStream = ofmtCtxAudio.GetStream(0)
				ofmtCtx = ofmtCtxAudio
			} else {
				continue
			}

			//convert PTS/DTS
			packet.Pts = libavutil.AvRescaleQRnd(packet.Pts, inStream.TimeBase, outStream.TimeBase,
				libavutil.AV_ROUND_NEAR_INF|libavutil.AV_ROUND_PASS_MINMAX)
			packet.Dts = libavutil.AvRescaleQRnd(packet.Dts, inStream.TimeBase, outStream.TimeBase,
				libavutil.AV_ROUND_NEAR_INF|libavutil.AV_ROUND_PASS_MINMAX)
			packet.Duration = libavutil.AvRescaleQ(packet.Duration, inStream.TimeBase, outStream.TimeBase)
			packet.Pos = -1
			packet.StreamIndex = 0

			//write
			if ofmtCtx.AvInterleavedWriteFrame(&packet) < 0 {
				fmt.Printf("Error muxing packet\n")
				break
			}

			packet.AvPacketUnref()
		}

		//write file trailer
		ofmtCtxVideo.AvWriteTrailer()
		ofmtCtxAudio.AvWriteTrailer()

		break
	}

	libavformat.AvformatCloseInput(&ifmtCtx)

	if ofmtCtxVideo != nil && (ofmtCtxVideo.Oformat.Flags&libavformat.AVFMT_NOFILE) == 0 {
		ofmtCtxVideo.Pb.AvioClose()
	}

	if ofmtCtxAudio != nil && (ofmtCtxAudio.Oformat.Flags&libavformat.AVFMT_NOFILE) == 0 {
		ofmtCtxAudio.Pb.AvioClose()
	}

	ofmtCtxVideo.AvformatFreeContext()
	ofmtCtxAudio.AvformatFreeContext()
	fmt.Println("-----------------------------------------")
	go func() {
		_, err = exec.Command("./lib/ffplay.exe", outFilenameAudio).Output()
		if err != nil {
			fmt.Println("play err = ", err)
		}
	}()
	_, err = exec.Command("./lib/ffplay.exe", outFilenameVideo).Output()
	if err != nil {
		fmt.Println("play err = ", err)
	}
}
