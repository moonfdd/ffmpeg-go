package main

import (
	"fmt"
	"os"
	"os/exec"
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/ffmpeg-go/libavcodec"
	"github.com/moonfdd/ffmpeg-go/libavdevice"
	"github.com/moonfdd/ffmpeg-go/libavformat"
	"github.com/moonfdd/ffmpeg-go/libavutil"
	"github.com/moonfdd/ffmpeg-go/libswresample"
	"github.com/moonfdd/ffmpeg-go/libswscale"
)

/*
 * 音频播放命令：ffplay -ar 44100 -ac 1 -f s16le -i out.pcm
 */
const i44100 = 22050

func avcodec_save_audio_file(pFormatCtx *libavformat.AVFormatContext, streamIndex ffcommon.FInt, fileName string) ffcommon.FInt {
	var pCodec *libavcodec.AVCodec
	var pCodecCtx *libavcodec.AVCodecContext
	codecpar := pFormatCtx.GetStream(uint32(streamIndex)).Codecpar

	//4.获取解码器（一）:音频
	//根据索引拿到对应的流
	pCodec = libavcodec.AvcodecFindDecoder(codecpar.CodecId)
	if pCodec == nil {
		fmt.Printf("can't decoder audio\n")
		return -1
	}
	//申请一个解码上下文
	pCodecCtx = pCodec.AvcodecAllocContext3()
	if pCodecCtx == nil {
		fmt.Printf("can't allocate a audio decoding context\n")
		return -1
	}

	//用流解码信息初始化编码参数
	pCodecCtx.AvcodecParametersToContext(codecpar)

	//没有此句会出现：Could not update timestamps for skipped samples
	pCodecCtx.PktTimebase = pFormatCtx.GetStream(uint32(streamIndex)).TimeBase

	//5.打开解码器
	if pCodecCtx.AvcodecOpen2(pCodec, nil) < 0 {
		fmt.Printf("can't open codec\n")
		return -1
	}

	//	printf("--------------- File Information ----------------\n");
	//	av_dump_format(pFormatCtx, 0, fileName, 0);
	//	printf("-------------------------------------------------\n");

	//编码数据
	packet := new(libavcodec.AVPacket)
	//解压缩数据
	frame := libavutil.AvFrameAlloc()

	//frame->16bit 44100 PCM统一音频采样格式与采样率
	swrCtx := libswresample.SwrAlloc()
	//重采样设置选项-----------------------------------------------------------start
	//输入的采样格式
	inSampleFmt := pCodecCtx.SampleFmt
	//输出的采样格式
	var outSampleFmt libswresample.AVSampleFormat = libavutil.AV_SAMPLE_FMT_S16
	//输入的采样率
	inSampleRate := pCodecCtx.SampleRate
	//输出的采样率
	var outSampleRate ffcommon.FInt = i44100
	//输入的声道布局
	var inChannelLayout ffcommon.FUint64T = pCodecCtx.ChannelLayout
	//输出的声道布局：CHANNEL_IN_MONO为单声道，CHANNEL_IN_STEREO为双声道
	var outChannelLayout ffcommon.FUint64T = libavutil.AV_CH_LAYOUT_MONO

	fmt.Printf("inSampleFmt = %d, inSampleRate = %d, inChannelLayout = %d， name = %s\n", inSampleFmt, inSampleRate,
		inChannelLayout, ffcommon.StringFromPtr(pCodec.Name))

	swrCtx.SwrAllocSetOpts(int64(outChannelLayout), outSampleFmt, outSampleRate,
		int64(inChannelLayout), inSampleFmt, inSampleRate, 0, uintptr(0))
	swrCtx.SwrInit()
	//重采样设置选项-----------------------------------------------------------end

	//获取输出的声道个数
	outChannelNb := libavutil.AvGetChannelLayoutNbChannels(outChannelLayout)
	fmt.Printf("outChannelNb = %d\n", outChannelNb)

	// 	 //存储PCM数据
	outBuffer := (*byte)(unsafe.Pointer(libavutil.AvMalloc(2 * i44100)))

	// 	 FILE *fp = fopen(fileName, "wb");
	fp, _ := os.Create(fileName)

	//回到流的初始位置
	pFormatCtx.AvSeekFrame(streamIndex, 0, libavformat.AVSEEK_FLAG_BACKWARD)

	//6.一帧一帧读取压缩的音频数据AVPacket
	for pFormatCtx.AvReadFrame(packet) >= 0 {
		if packet.StreamIndex == uint32(streamIndex) {
			//解码AVPacket --> AVFrame
			ret := pCodecCtx.AvcodecSendPacket(packet)
			if ret < 0 {
				fmt.Printf("Decode error\n")
				break
			}

			if pCodecCtx.AvcodecReceiveFrame(frame) >= 0 {
				swrCtx.SwrConvert(&outBuffer, 2*i44100, (**byte)(unsafe.Pointer(&frame.Data)), frame.NbSamples)
				//获取sample的size
				outBufferSize := libavutil.AvSamplesGetBufferSize(nil, outChannelNb, frame.NbSamples, outSampleFmt, 1)
				//写入文件
				fp.Write(ffcommon.ByteSliceFromByteP(outBuffer, int(outBufferSize)))
			}
		}

		packet.AvPacketUnref()
	}

	fp.Close()
	libavutil.AvFrameFree(&frame)
	libavutil.AvFree(uintptr(unsafe.Pointer(outBuffer)))
	libswresample.SwrFree(&swrCtx)
	pCodecCtx.AvcodecClose()

	return 0
}

/*
 * 视频播放命令：ffplay -video_size 654x368 -i out.yuv
 */
func avcodec_save_video_file(pFormatCtx *libavformat.AVFormatContext, streamIndex ffcommon.FInt, fileName string) ffcommon.FInt {
	var pCodec *libavcodec.AVCodec
	var pCodecCtx *libavcodec.AVCodecContext
	codecpar := pFormatCtx.GetStream(uint32(streamIndex)).Codecpar

	//4.获取解码器（一）:音频
	//根据索引拿到对应的流
	pCodec = libavcodec.AvcodecFindDecoder(codecpar.CodecId)
	if pCodec == nil {
		fmt.Printf("can't decoder audio\n")
		return -1
	}
	//申请一个解码上下文
	pCodecCtx = pCodec.AvcodecAllocContext3()
	if pCodecCtx == nil {
		fmt.Printf("can't allocate a audio decoding context\n")
		return -1
	}

	//用流解码信息初始化编码参数
	pCodecCtx.AvcodecParametersToContext(codecpar)

	//没有此句会出现：Could not update timestamps for skipped samples
	pCodecCtx.PktTimebase = pFormatCtx.GetStream(uint32(streamIndex)).TimeBase

	//5.打开解码器
	if pCodecCtx.AvcodecOpen2(pCodec, nil) < 0 {
		fmt.Printf("can't open codec\n")
		return -1
	}

	//	printf("--------------- File Information ----------------\n");
	//	av_dump_format(pFormatCtx, 0, fileName, 0);
	//	printf("-------------------------------------------------\n");

	//编码数据
	pPacket := new(libavcodec.AVPacket)
	//解压缩数据
	pFrame := libavutil.AvFrameAlloc()
	pFrameYUV := libavutil.AvFrameAlloc()

	outBuffer := (*byte)(unsafe.Pointer(libavutil.AvMalloc(
		uint64(libavutil.AvImageGetBufferSize(libavutil.AV_PIX_FMT_YUV420P, pCodecCtx.Width, pCodecCtx.Height, 1)))))
	libavutil.AvImageFillArrays((*[4]*byte)(unsafe.Pointer(&pFrameYUV.Data)), (*[4]int32)(unsafe.Pointer(&pFrameYUV.Linesize)), outBuffer,
		libavutil.AV_PIX_FMT_YUV420P, pCodecCtx.Width,
		pCodecCtx.Height, 1)

	pImgConvertCtx := libswscale.SwsGetContext(pCodecCtx.Width, pCodecCtx.Height, pCodecCtx.PixFmt,
		pCodecCtx.Width, pCodecCtx.Height, libavutil.AV_PIX_FMT_YUV420P, libswscale.SWS_BICUBIC, nil, nil, nil)

	fmt.Printf("width = %d, height = %d, name = %s\n", pCodecCtx.Width, pCodecCtx.Height, ffcommon.StringFromPtr(pCodec.Name))

	fp, _ := os.Create(fileName)

	//回到流的初始位置
	pFormatCtx.AvSeekFrame(streamIndex, 0, libavformat.AVSEEK_FLAG_BACKWARD)

	//6.一帧一帧读取压缩的视频数据AVPacket
	for pFormatCtx.AvReadFrame(pPacket) >= 0 {
		if pPacket.StreamIndex == uint32(streamIndex) {
			//解码AVPacket --> AVFrame
			ret := pCodecCtx.AvcodecSendPacket(pPacket)
			if ret < 0 {
				fmt.Printf("Decode error\n")
				break
			}

			if pCodecCtx.AvcodecReceiveFrame(pFrame) >= 0 {
				pImgConvertCtx.SwsScale((**byte)(unsafe.Pointer(&pFrame.Data)), (*int32)(unsafe.Pointer(&pFrame.Linesize)), 0,
					uint32(pCodecCtx.Height), (**byte)(unsafe.Pointer(&pFrameYUV.Data)), (*int32)(unsafe.Pointer(&pFrameYUV.Linesize)))

				y_size := pCodecCtx.Width * pCodecCtx.Height
				fp.Write(ffcommon.ByteSliceFromByteP(pFrameYUV.Data[0], int(y_size)))   //Y
				fp.Write(ffcommon.ByteSliceFromByteP(pFrameYUV.Data[1], int(y_size/4))) //U
				fp.Write(ffcommon.ByteSliceFromByteP(pFrameYUV.Data[2], int(y_size/4))) //V
			}
		}

		pPacket.AvPacketUnref()
	}

	fp.Close()
	libavutil.AvFrameFree(&pFrame)
	libavutil.AvFrameFree(&pFrameYUV)
	libavutil.AvFree(uintptr(unsafe.Pointer(outBuffer)))
	pCodecCtx.AvcodecClose()

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

	inputFile := "./resources/big_buck_bunny.mp4"
	outAudioFile := "./out/a23.pcm"
	outVideoFile := "./out/a23.yuv"

	var videoStreamIndex ffcommon.FInt = -1
	var audioStreamIndex ffcommon.FInt = -1
	var i ffcommon.FUnsignedInt = 0
	var pFormatCtx *libavformat.AVFormatContext

	//1.注册组件
	libavdevice.AvdeviceRegisterAll()

	//封装格式上下文
	pFormatCtx = libavformat.AvformatAllocContext()

	//2.打开输入文件
	if libavformat.AvformatOpenInput(&pFormatCtx, inputFile, nil, nil) != 0 {
		fmt.Printf("can't open input file\n")
		return
	}

	//3.获取音视频信息
	if pFormatCtx.AvformatFindStreamInfo(nil) < 0 {
		fmt.Printf("can't find stream info\n")
		return
	}

	//音视频编码，找到对应的音视频流的索引位置
	//找到音频流的索引
	for i = 0; i < pFormatCtx.NbStreams; i++ {
		if pFormatCtx.GetStream(i).Codecpar.CodecType == libavutil.AVMEDIA_TYPE_AUDIO {
			audioStreamIndex = int32(i)
			break
		}
	}

	//找到视频流的索引
	for i = 0; i < pFormatCtx.NbStreams; i++ {
		if pFormatCtx.GetStream(i).Codecpar.CodecType == libavutil.AVMEDIA_TYPE_VIDEO {
			videoStreamIndex = int32(i)
			break
		}
	}

	fmt.Printf("audioStreamIndex = %d, videoStreamIndex = %d\n", audioStreamIndex, videoStreamIndex)

	if audioStreamIndex == -1 {
		fmt.Printf("can't find a audio stream\n")
	} else {
		fmt.Printf("try to save audio stream\n")
		avcodec_save_audio_file(pFormatCtx, audioStreamIndex, outAudioFile)
	}

	if videoStreamIndex == -1 {
		fmt.Printf("can't find a video stream\n")
	} else {
		fmt.Printf("try to save video stream\n")
		avcodec_save_video_file(pFormatCtx, videoStreamIndex, outVideoFile)
	}

	libavformat.AvformatCloseInput(&pFormatCtx)

	fmt.Println("-----------------------------------------")
	go func() {
		_, err = exec.Command("./lib/ffplay.exe", "-ar", "22050", "-ac", "1", "-f", "s16le", "-i", outAudioFile).Output()
		if err != nil {
			fmt.Println("play err = ", err)
		}
	}()
	_, err = exec.Command("./lib/ffplay.exe", "-video_size", "640*360", "-i", outVideoFile).Output()
	if err != nil {
		fmt.Println("play err = ", err)
	}
}
