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

type S struct {
	AA int32
	BB int32
	CC int
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

	inVFileName := "./out/test.mp3"
	outFileName := "./out/test.pcm"

	// ./lib/ffmpeg -i ./resources/big_buck_bunny.mp4 -acodec libmp3lame -vn ./out/test.mp3
	//是否存在mp3文件
	_, err = os.Stat(inVFileName)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("create mp3 file")
			exec.Command("./lib/ffmpeg", "-i", "./resources/big_buck_bunny.mp4", "-acodec", "libmp3lame", "-vn", inVFileName, "-y").CombinedOutput()
		}
	}

	os.Remove(outFileName)
	f, err := os.OpenFile(outFileName, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}

	fmtCtx := libavformat.AvformatAllocContext()
	var codecCtx *libavcodec.AVCodecContext
	pkt := libavcodec.AvPacketAlloc()
	frame := libavutil.AvFrameAlloc()

	aStreamIndex := -1

	for {
		if libavformat.AvformatOpenInput(&fmtCtx, inVFileName, nil, nil) < 0 {
			fmt.Printf("Cannot open input file.\n")
			break
		}

		if fmtCtx.AvformatFindStreamInfo(nil) < 0 {
			fmt.Printf("Cannot find stream info in input file.\n")
			break
		}

		fmtCtx.AvDumpFormat(0, inVFileName, 0)

		//查找视频流在文件中的位置
		for i := uint32(0); i < fmtCtx.NbStreams; i++ {
			if fmtCtx.GetStream(i).Codecpar.CodecType == libavutil.AVMEDIA_TYPE_AUDIO {
				aStreamIndex = int(i)
				break
			}
		}

		if aStreamIndex == -1 {
			fmt.Printf("Cannot find audio stream.\n")
			return
		}

		aCodecPara := fmtCtx.GetStream(uint32(aStreamIndex)).Codecpar
		codec := libavcodec.AvcodecFindDecoder(aCodecPara.CodecId)
		if codec == nil {
			fmt.Printf("Cannot find any codec for audio.\n")
			return
		}

		codecCtx = codec.AvcodecAllocContext3()

		if codecCtx.AvcodecParametersToContext(aCodecPara) < 0 {
			fmt.Printf("Cannot alloc codec context.\n")
			return
		}

		codecCtx.PktTimebase = fmtCtx.GetStream(uint32(aStreamIndex)).TimeBase

		if codecCtx.AvcodecOpen2(codec, nil) < 0 {
			fmt.Printf("Cannot open audio codec.\n")
			return
		}

		for (fmtCtx.AvReadFrame(pkt)) >= 0 {
			if pkt.StreamIndex == uint32(aStreamIndex) {
				if codecCtx.AvcodecSendPacket(pkt) >= 0 {
					for codecCtx.AvcodecReceiveFrame(frame) >= 0 {
						/*
						   Planar（平面），其数据格式排列方式为 (特别记住，该处是以点nb_samples采样点来交错，不是以字节交错）:
						   LLLLLLRRRRRRLLLLLLRRRRRRLLLLLLRRRRRRL...（每个LLLLLLRRRRRR为一个音频帧）
						   而不带P的数据格式（即交错排列）排列方式为：
						   LRLRLRLRLRLRLRLRLRLRLRLRLRLRLRLRLRLRL...（每个LR为一个音频样本）
						*/
						if libavutil.AvSampleFmtIsPlanar(codecCtx.SampleFmt) != 0 {
							numBytes := libavutil.AvGetBytesPerSample(codecCtx.SampleFmt)
							//pcm播放时是LRLRLR格式，所以要交错保存数据
							bytes := []byte{}
							for i := int32(0); i < frame.NbSamples; i++ {
								for ch := int32(0); ch < codecCtx.Channels; ch++ {
									ptr := uintptr(unsafe.Pointer(uintptr(unsafe.Pointer(frame.Data[ch])) + uintptr(numBytes*i)))
									for k := int32(0); k < numBytes; k++ {
										bytes = append(bytes, *(*byte)(unsafe.Pointer(ptr)))
										ptr++
									}

								}
							}
							f.Write(bytes)
						}
					}
				}
			}
			pkt.AvPacketUnref()
		}

		break
	}

	libavutil.AvFrameFree(&frame)
	libavcodec.AvPacketFree(&pkt)
	codecCtx.AvcodecClose()
	libavcodec.AvcodecFreeContext(&codecCtx)
	fmtCtx.AvformatFreeContext()
	f.Close()

	fmt.Println("-----------------------------------------")
	// ./lib/ffplay -ar 22050 -ac 2 -f f32le -i ./out/test.pcm
	_, err = exec.Command("./lib/ffplay.exe", "-ar", "22050", "-ac", "2", "-f", "f32le", "-i", "./out/test.pcm").Output()
	if err != nil {
		fmt.Println("play err = ", err)
	}
}
