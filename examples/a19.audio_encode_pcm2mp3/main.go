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
	"github.com/moonfdd/ffmpeg-go/libswresample"
)

func main() {
	// https://blog.csdn.net/guoyunfei123/article/details/105643255
	// 时长没误差
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

	//./lib/ffmpeg -i .\resources\big_buck_bunny.mp4 -f s16le -ar 44100 -ac 2 -acodec pcm_s16le -vn ./out/s16le.pcm
	// ./lib/ffmpeg -y -f s16le -ac 2 -ar 44100 -acodec pcm_s16le -vn -i ./out/s16le.pcm ./out/s16le.mp3
	inFileName := "./out/s16le.pcm"
	// inFileName := "./out/test16.pcm"
	outFileName := "./out/out19.mp3"

	// ./lib/ffmpeg -i ./resources/big_buck_bunny.mp4 -acodec libmp3lame -vn ./out/test.mp3
	//是否存在mp3文件
	// _, err = os.Stat(inVFileName)
	// if err != nil {
	// 	if os.IsNotExist(err) {
	// 		fmt.Println("create mp3 file")
	// 		exec.Command("./lib/ffmpeg", "-i", "./resources/big_buck_bunny.mp4", "-acodec", "libmp3lame", "-vn", inVFileName, "-y").CombinedOutput()
	// 	}
	// }

	// os.Remove(outFileName)
	// f, err := os.OpenFile(outFileName, os.O_CREATE|os.O_RDWR, 0777)
	// if err != nil {
	// 	fmt.Println("open file failed,err:", err)
	// 	return
	// }

	var pFormatCtx *libavformat.AVFormatContext
	var pCodecCtx *libavcodec.AVCodecContext
	var pCodec *libavcodec.AVCodec
	var pkt libavcodec.AVPacket
	var pFrame *libavutil.AVFrame

	//libavdevice.AvdeviceRegisterAll()

	for {

		libavformat.AvformatAllocOutputContext2(&pFormatCtx, nil, "", outFileName)

		if libavformat.AvioOpen(&pFormatCtx.Pb, outFileName, libavformat.AVIO_FLAG_READ_WRITE) < 0 {
			fmt.Printf("Cannot open output file.\n")
			return
		}

		stream := pFormatCtx.AvformatNewStream(nil)
		if stream == nil {
			fmt.Printf("Cannot create a new stream to output file.\n")
			return
		}

		//设置参数
		pCodecCtx = stream.Codec
		pCodecCtx.CodecType = libavutil.AVMEDIA_TYPE_AUDIO
		pCodecCtx.CodecId = pFormatCtx.Oformat.AudioCodec
		pCodecCtx.SampleFmt = libavutil.AV_SAMPLE_FMT_FLTP
		pCodecCtx.SampleRate = 44100
		pCodecCtx.ChannelLayout = libavutil.AV_CH_LAYOUT_STEREO
		pCodecCtx.BitRate = 128000
		pCodecCtx.Channels = libavutil.AvGetChannelLayoutNbChannels(pCodecCtx.ChannelLayout)

		//查找编码器
		pCodec = libavcodec.AvcodecFindEncoder(pCodecCtx.CodecId)
		if pCodec == nil {
			fmt.Printf("Cannot find audio encoder.\n")
			return
		}

		//打开编码器
		if pCodecCtx.AvcodecOpen2(pCodec, nil) < 0 {
			fmt.Printf("Cannot open encoder.\n")
			return
		}

		//fmtCtx.AvDumpFormat(0, outFileName, 1)
		pFrame = libavutil.AvFrameAlloc()
		if pFrame == nil {
			fmt.Printf("can't alloc frame\n")
			return
		}

		//===========
		pFrame.NbSamples = pCodecCtx.FrameSize
		pFrame.Format = int32(pCodecCtx.SampleFmt)
		pFrame.Channels = 2

		// PCM重采样
		var swr_ctx *libswresample.SwrContext = libswresample.SwrAlloc()
		swr_ctx.SwrAllocSetOpts(libavutil.AvGetDefaultChannelLayout(pCodecCtx.Channels),
			pCodecCtx.SampleFmt,
			pCodecCtx.SampleRate,
			libavutil.AvGetDefaultChannelLayout(pFrame.Channels),
			libavutil.AV_SAMPLE_FMT_S16, // PCM源文件的采样格式
			44100,
			0, uintptr(0))
		swr_ctx.SwrInit()

		/* 分配空间 */
		// uint8_t **convert_data = (uint8_t**)calloc(codecCtx->channels,sizeof(*convert_data));
		convert_data := (**byte)(unsafe.Pointer(libavutil.AvCalloc(uint64(pCodecCtx.Channels), 8)))
		libavutil.AvSamplesAlloc(convert_data, nil, pCodecCtx.Channels, pCodecCtx.FrameSize,
			pCodecCtx.SampleFmt, 0)

		size := libavutil.AvSamplesGetBufferSize(nil, pCodecCtx.Channels,
			pCodecCtx.FrameSize, pCodecCtx.SampleFmt, 1)
		frameBuf := libavutil.AvMalloc(uint64(size))
		libavcodec.AvcodecFillAudioFrame(pFrame, pCodecCtx.Channels, pCodecCtx.SampleFmt,
			(*byte)(unsafe.Pointer(frameBuf)), size, 1)

		//写帧头
		pFormatCtx.AvformatWriteHeader(nil)

		inFile, err := os.Open(inFileName)
		if err != nil {
			fmt.Printf("annot open input file.\n")
			return
		}

		pkt.AvInitPacket()
		pkt.Data = nil
		pkt.Size = 0

		for i := 0; ; i++ {
			//输入一帧数据的长度
			length := pFrame.NbSamples * libavutil.AvGetBytesPerSample(libavutil.AV_SAMPLE_FMT_S16) * pFrame.Channels
			//读PCM：特意注意读取的长度，否则可能出现转码之后声音变快或者变慢
			buf := make([]byte, length)
			n, err := inFile.Read(buf)
			if err != nil {
				fmt.Println("read end")
				break
			}
			if n <= 0 {
				break
			}

			for j := 0; j < n; j++ {
				*(*byte)(unsafe.Pointer(frameBuf + uintptr(j))) = buf[j]
			}

			swr_ctx.SwrConvert(convert_data, pCodecCtx.FrameSize,
				(**byte)(unsafe.Pointer(&pFrame.Data)),
				pFrame.NbSamples)

			//输出一帧数据的长度
			length = pCodecCtx.FrameSize * libavutil.AvGetBytesPerSample(pCodecCtx.SampleFmt)
			//双通道赋值（输出的AAC为双通道）
			// memcpy(frame->data[0],convert_data[0],length);
			// memcpy(frame->data[1],convert_data[1],length);
			c := *(*[2]uintptr)(unsafe.Pointer(convert_data))
			fd0 := uintptr(unsafe.Pointer(pFrame.Data[0]))
			cd0 := uintptr(unsafe.Pointer(c[0]))
			fd1 := uintptr(unsafe.Pointer(pFrame.Data[1]))
			cd1 := uintptr(unsafe.Pointer(c[1]))
			for j := int32(0); j < length; j++ {
				*(*byte)(unsafe.Pointer(fd0)) = *(*byte)(unsafe.Pointer(cd0))
				*(*byte)(unsafe.Pointer(fd1)) = *(*byte)(unsafe.Pointer(cd1))
				fd0++
				cd0++
				fd1++
				cd1++
			}

			pFrame.Pts = int64(i * 100)
			if pCodecCtx.AvcodecSendFrame(pFrame) < 0 {
				fmt.Printf("can't send frame for encoding\n")
				break

				// for codecCtx.AvcodecReceivePacket(pkt) >= 0 {
				// 	pkt.StreamIndex = uint32(outStream.Index)
				// 	fmt.Printf("write %4d frame, size=%d, length=%d\n", i, size, length)
				// 	fmtCtx.AvWriteFrame(pkt)
				// }
			}
			if pCodecCtx.AvcodecReceivePacket(&pkt) >= 0 {
				pkt.StreamIndex = uint32(stream.Index)
				fmt.Printf("write %4d frame, size = %d, length = %d\n", i, size, length)
				pFormatCtx.AvWriteFrame(&pkt)

			}
			pkt.AvPacketUnref()
		}

		// flush encoder
		// if flush_encoder(fmtCtx, codecCtx, int(outStream.Index)) < 0 {
		// 	fmt.Printf("Cannot flush encoder.\n")
		// 	return
		// }
		if flush_encoder(pFormatCtx, 0) < 0 {
			fmt.Printf("flushing encoder failed\n")
			return
		}

		// write trailer
		pFormatCtx.AvWriteTrailer()

		inFile.Close()
		stream.Codec.AvcodecClose()
		libavutil.AvFree(uintptr(unsafe.Pointer(pFrame)))
		libavutil.AvFree(frameBuf)
		pFormatCtx.Pb.AvioClose()
		pFormatCtx.AvformatFreeContext()
		break
	}

	// codecCtx.AvcodecClose()
	// libavutil.AvFree(uintptr(unsafe.Pointer(frame)))
	// fmtCtx.Pb.AvioClose()
	// fmtCtx.AvformatFreeContext()
	return
	fmt.Println("-----------------------------------------")
	// ./lib/ffplay -ar 44100 -ac 2 -f s16le -i ./out/test.pcm
	_, err = exec.Command("./lib/ffplay.exe", "-ar", "44100", "-ac", "2", "-f", "s16le", "-i", "./out/test16.pcm").Output()
	if err != nil {
		fmt.Println("play err = ", err)
	}
}
func flush_encoder(fmt_ctx *libavformat.AVFormatContext, stream_index int) int32 {
	ret := int32(0)
	var got_frame int32
	var enc_pkt libavcodec.AVPacket
	if fmt_ctx.GetStream(uint32(stream_index)).Codec.Codec.Capabilities&libavcodec.AV_CODEC_CAP_DELAY == 0 {
		return 0
	}
	for {
		enc_pkt.Data = nil
		enc_pkt.Size = 0
		enc_pkt.AvInitPacket()
		ret = fmt_ctx.GetStream(uint32(stream_index)).Codec.AvcodecEncodeAudio2(&enc_pkt,
			nil, &got_frame)
		//av_frame_free(NULL)
		if ret < 0 {
			break
		}
		if got_frame == 0 {
			ret = 0
			break
		}
		fmt.Printf("Flush Encoder: Succeed to encode 1 frame!\tsize:%5d\n", enc_pkt.Size)
		/* mux encoded frame */
		ret = fmt_ctx.AvWriteFrame(&enc_pkt)
		if ret < 0 {
			break
		}
	}

	// fmt.Printf("Flushing stream #%d encoder\n", aStreamIndex)
	// ret = codecCtx.AvcodecSendFrame(nil)
	// if ret >= 0 {
	// 	for codecCtx.AvcodecReceivePacket(enc_pkt) >= 0 {
	// 		fmt.Printf("success encoder 1 frame.\n")
	// 		/* mux encoded frame */
	// 		ret = fmtCtx.AvWriteFrame(enc_pkt)
	// 		if ret < 0 {
	// 			break
	// 		}
	// 	}
	// }

	return ret
}
