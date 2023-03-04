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

/*
FIX: H.264 in some container format (FLV, MP4, MKV etc.) need
"h264_mp4toannexb" bitstream filter (BSF)
  *Add SPS,PPS in front of IDR frame
  *Add start code ("0,0,0,1") in front of NALU
H.264 in some container (MPEG2TS) don't need this BSF.
*/
//'1': Use H.264 Bitstream Filter
const USE_H264BSF = 0

/*
FIX:AAC in some container format (FLV, MP4, MKV etc.) need
"aac_adtstoasc" bitstream filter (BSF)
*/
//'1': Use AAC Bitstream Filter
const USE_AACBSF = 0

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

	inFilenameVideo := "./out/a24.h264"
	inFilenameAudio := "./out/a24.aac"
	outFilename := "./out/a24.mp4"

	_, err = os.Stat(inFilenameVideo)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("create h264 file")
			exec.Command("./lib/ffmpeg", "-i", "./resources/big_buck_bunny.mp4", "-vcodec", "copy", "-an", inFilenameVideo).Output()
		}
	}

	_, err = os.Stat(inFilenameAudio)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("create aac file")
			exec.Command("./lib/ffmpeg", "-i", "./resources/big_buck_bunny.mp4", "-acodec", "copy", "-vn", inFilenameAudio).Output()
		}
	}

	var ifmtCtxVideo, ifmtCtxAudio, ofmtCtx *libavformat.AVFormatContext
	var packet libavcodec.AVPacket

	var inVideoIndex, inAudioIndex ffcommon.FInt = -1, -1
	var outVideoIndex, outAudioIndex ffcommon.FInt = -1, -1
	var frameIndex ffcommon.FInt = 0

	var curPstVideo, curPstAudio ffcommon.FInt64T = 0, 0

	var ret ffcommon.FInt = 0
	var i ffcommon.FUnsignedInt = 0

	var h264bsfc, aacbsfc *libavcodec.AVBitStreamFilterContext

	//打开输入视频文件
	ret = libavformat.AvformatOpenInput(&ifmtCtxVideo, inFilenameVideo, nil, nil)
	if ret < 0 {
		fmt.Printf("can't open input video file\n")
		goto end
	}

	//查找输入流
	ret = ifmtCtxVideo.AvformatFindStreamInfo(nil)
	if ret < 0 {
		fmt.Printf("failed to retrieve input video stream information\n")
		goto end
	}

	//打开输入音频文件
	ret = libavformat.AvformatOpenInput(&ifmtCtxAudio, inFilenameAudio, nil, nil)
	if ret < 0 {
		fmt.Printf("can't open input audio file\n")
		goto end
	}

	//查找输入流
	ret = ifmtCtxAudio.AvformatFindStreamInfo(nil)
	if ret < 0 {
		fmt.Printf("failed to retrieve input audio stream information\n")
		goto end
	}

	fmt.Printf("===========Input Information==========\n")
	ifmtCtxVideo.AvDumpFormat(0, inFilenameVideo, 0)
	ifmtCtxAudio.AvDumpFormat(0, inFilenameAudio, 0)
	fmt.Printf("======================================\n")

	//新建输出上下文
	libavformat.AvformatAllocOutputContext2(&ofmtCtx, nil, "", outFilename)
	if ofmtCtx == nil {
		fmt.Printf("can't create output context\n")
		goto end
	}

	//视频输入流
	for i = 0; i < ifmtCtxVideo.NbStreams; i++ {
		if ifmtCtxVideo.GetStream(i).Codecpar.CodecType == libavutil.AVMEDIA_TYPE_VIDEO {
			inStream := ifmtCtxVideo.GetStream(i)
			outStream := ofmtCtx.AvformatNewStream(nil)
			inVideoIndex = int32(i)

			if outStream == nil {
				fmt.Printf("failed to allocate output stream\n")
				goto end
			}

			outVideoIndex = outStream.Index

			if libavcodec.AvcodecParametersCopy(outStream.Codecpar, inStream.Codecpar) < 0 {
				fmt.Printf("faild to copy context from input to output stream")
				goto end
			}

			outStream.Codecpar.CodecTag = 0

			break
		}
	}

	//音频输入流
	for i = 0; i < ifmtCtxAudio.NbStreams; i++ {
		if ifmtCtxAudio.GetStream(i).Codecpar.CodecType == libavutil.AVMEDIA_TYPE_AUDIO {
			inStream := ifmtCtxAudio.GetStream(i)
			outStream := ofmtCtx.AvformatNewStream(nil)
			inAudioIndex = int32(i)

			if outStream == nil {
				fmt.Printf("failed to allocate output stream\n")
				goto end
			}

			outAudioIndex = outStream.Index

			if libavcodec.AvcodecParametersCopy(outStream.Codecpar, inStream.Codecpar) < 0 {
				fmt.Printf("faild to copy context from input to output stream")
				goto end
			}

			outStream.Codecpar.CodecTag = 0

			break
		}
	}

	fmt.Printf("==========Output Information==========\n")
	ofmtCtx.AvDumpFormat(0, outFilename, 1)
	fmt.Printf("======================================\n")

	//打开输入文件
	if ofmtCtx.Oformat.Flags&libavformat.AVFMT_NOFILE == 0 {
		if libavformat.AvioOpen(&ofmtCtx.Pb, outFilename, libavformat.AVIO_FLAG_WRITE) < 0 {
			fmt.Printf("can't open out file\n")
			goto end
		}
	}

	//写文件头
	if ofmtCtx.AvformatWriteHeader(nil) < 0 {
		fmt.Printf("Error occurred when opening output file\n")
		goto end
	}

	if USE_H264BSF != 0 {
		h264bsfc = libavcodec.AvBitstreamFilterInit("h264_mp4toannexb")
	}
	if USE_AACBSF != 0 {
		aacbsfc = libavcodec.AvBitstreamFilterInit("aac_adtstoasc")
	}
	for {
		var ifmtCtx *libavformat.AVFormatContext
		var inStream, outStream *libavformat.AVStream
		var streamIndex ffcommon.FInt = 0

		if libavutil.AvCompareTs(curPstVideo, ifmtCtxVideo.GetStream(uint32(inVideoIndex)).TimeBase, curPstAudio, ifmtCtxAudio.GetStream(uint32(inAudioIndex)).TimeBase) < 0 {
			ifmtCtx = ifmtCtxVideo
			streamIndex = outVideoIndex

			if ifmtCtx.AvReadFrame(&packet) >= 0 {
				for {
					inStream = ifmtCtx.GetStream(packet.StreamIndex)
					outStream = ofmtCtx.GetStream(uint32(streamIndex))

					if packet.StreamIndex == uint32(inVideoIndex) {
						//Fix: No PTS(Example: Raw H.264
						//Simple Write PTS
						if packet.Pts == libavutil.AV_NOPTS_VALUE {
							//write PTS
							timeBase1 := inStream.TimeBase
							//Duration between 2 frames
							calcDuration := int64(libavutil.AV_TIME_BASE / libavutil.AvQ2d(inStream.RFrameRate))
							//Parameters
							packet.Pts = int64((float64(frameIndex) * float64(calcDuration)) / (libavutil.AvQ2d(timeBase1) * libavutil.AV_TIME_BASE))
							packet.Dts = packet.Pts
							packet.Duration = int64(float64(calcDuration) / (libavutil.AvQ2d(timeBase1) * libavutil.AV_TIME_BASE))
							frameIndex++
						}

						curPstVideo = packet.Pts
						break
					}
					if ifmtCtx.AvReadFrame(&packet) >= 0 {

					} else {
						break
					}
				}
			} else {
				break
			}
		} else {
			ifmtCtx = ifmtCtxAudio
			streamIndex = outAudioIndex

			if ifmtCtx.AvReadFrame(&packet) >= 0 {
				for {
					inStream = ifmtCtx.GetStream(packet.StreamIndex)
					outStream = ofmtCtx.GetStream(uint32(streamIndex))

					if packet.StreamIndex == uint32(inAudioIndex) {
						//Fix: No PTS(Example: Raw H.264
						//Simple Write PTS
						if packet.Pts == libavutil.AV_NOPTS_VALUE {
							//write PTS
							timeBase1 := inStream.TimeBase
							//Duration between 2 frames
							calcDuration := int64(libavutil.AV_TIME_BASE / libavutil.AvQ2d(inStream.RFrameRate))
							//Parameters
							packet.Pts = int64((float64(frameIndex) * float64(calcDuration)) / (libavutil.AvQ2d(timeBase1) * libavutil.AV_TIME_BASE))
							packet.Dts = packet.Pts
							packet.Duration = int64(float64(calcDuration) / (libavutil.AvQ2d(timeBase1) * libavutil.AV_TIME_BASE))
							frameIndex++
						}

						curPstAudio = packet.Pts
						break
					}
					if ifmtCtx.AvReadFrame(&packet) >= 0 {

					} else {
						break
					}
				}
			} else {
				break
			}
		}

		//FIX:Bitstream Filter
		if USE_H264BSF != 0 {
			h264bsfc.AvBitstreamFilterFilter(inStream.Codec, "", &packet.Data, (*int32)(unsafe.Pointer(&packet.Size)), packet.Data, int32(packet.Size), 0)
		}
		if USE_AACBSF != 0 {
			aacbsfc.AvBitstreamFilterFilter(outStream.Codec, "", &packet.Data, (*int32)(unsafe.Pointer(&packet.Size)), packet.Data, int32(packet.Size), 0)
		}

		//Convert PTS/DTS
		packet.Pts = libavutil.AvRescaleQRnd(packet.Pts, inStream.TimeBase, outStream.TimeBase, libavutil.AV_ROUND_NEAR_INF|libavutil.AV_ROUND_PASS_MINMAX)
		packet.Dts = libavutil.AvRescaleQRnd(packet.Dts, inStream.TimeBase, outStream.TimeBase, libavutil.AV_ROUND_NEAR_INF|libavutil.AV_ROUND_PASS_MINMAX)
		packet.Duration = libavutil.AvRescaleQ(packet.Duration, inStream.TimeBase, outStream.TimeBase)
		packet.Pos = -1
		packet.StreamIndex = uint32(streamIndex)

		//write
		if ofmtCtx.AvInterleavedWriteFrame(&packet) < 0 {
			fmt.Printf("error muxing packet")
			break
		}

		packet.AvPacketUnref()
	}

	//Write file trailer
	ofmtCtx.AvWriteTrailer()

	if USE_H264BSF != 0 {
		h264bsfc.AvBitstreamFilterClose()
	}
	if USE_AACBSF != 0 {
		aacbsfc.AvBitstreamFilterClose()
	}

end:
	libavformat.AvformatCloseInput(&ifmtCtxVideo)
	libavformat.AvformatCloseInput(&ifmtCtxAudio)
	if ofmtCtx != nil && ofmtCtx.Oformat.Flags&libavformat.AVFMT_NOFILE == 0 {
		ofmtCtx.Pb.AvioClose()
	}

	ofmtCtx.AvformatFreeContext()
	fmt.Println("-----------------------------------------")
	_, err = exec.Command("./lib/ffplay.exe", outFilename).Output()
	if err != nil {
		fmt.Println("play err = ", err)
	}
}
