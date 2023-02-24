package main

import (
	"fmt"
	"math"
	"os"
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/ffmpeg-go/libavcodec"
	"github.com/moonfdd/ffmpeg-go/libavutil"
)

func main0() (ret ffcommon.FInt) {
	var filename string
	var codec *libavcodec.AVCodec
	var c *libavcodec.AVCodecContext
	var frame *libavutil.AVFrame
	var pkt *libavcodec.AVPacket
	var i, j, k ffcommon.FInt
	var f *os.File
	var samples *ffcommon.FUint16T
	var t, tincr ffcommon.FFloat

	if len(os.Args) <= 1 {
		fmt.Printf("Usage: %s <output file>\n", os.Args[0])
		return 0
	}
	filename = os.Args[1]

	/* find the MP2 encoder */
	codec = libavcodec.AvcodecFindEncoder(libavcodec.AV_CODEC_ID_MP2)
	if codec == nil {
		fmt.Printf("Codec not found\n")
		os.Exit(1)
	}

	c = codec.AvcodecAllocContext3()
	if c == nil {
		fmt.Printf("Could not allocate audio codec context\n")
		os.Exit(1)
	}

	/* put sample parameters */
	c.BitRate = 64000

	/* check that the encoder supports s16 pcm input */
	c.SampleFmt = libavutil.AV_SAMPLE_FMT_S16
	if check_sample_fmt(codec, c.SampleFmt) == 0 {
		fmt.Printf("Encoder does not support sample format %s",
			libavutil.AvGetSampleFmtName(c.SampleFmt))
		os.Exit(1)
	}

	/* select other audio parameters supported by the encoder */
	c.SampleRate = select_sample_rate(codec)
	c.ChannelLayout = uint64(select_channel_layout(codec))
	c.Channels = libavutil.AvGetChannelLayoutNbChannels(c.ChannelLayout)

	/* open it */
	if c.AvcodecOpen2(codec, nil) < 0 {
		fmt.Printf("Could not open codec\n")
		os.Exit(1)
	}

	f, _ = os.Create(filename)
	if f == nil {
		fmt.Printf("Could not open %s\n", filename)
		os.Exit(1)
	}

	/* packet for holding encoded output */
	pkt = libavcodec.AvPacketAlloc()
	if pkt == nil {
		fmt.Printf("could not allocate the packet\n")
		os.Exit(1)
	}

	/* frame containing input raw audio */
	frame = libavutil.AvFrameAlloc()
	if frame == nil {
		fmt.Printf("Could not allocate audio frame\n")
		os.Exit(1)
	}

	frame.NbSamples = c.FrameSize
	frame.Format = int32(c.SampleFmt)
	frame.ChannelLayout = c.ChannelLayout

	/* allocate the data buffers */
	ret = frame.AvFrameGetBuffer(0)
	if ret < 0 {
		fmt.Printf("Could not allocate audio data buffers\n")
		os.Exit(1)
	}

	/* encode a single tone sound */
	t = 0
	tincr = float32(2 * libavutil.M_PI * 440.0 / float64(c.SampleRate))
	for i = 0; i < 200; i++ {
		/* make sure the frame is writable -- makes a copy if the encoder
		 * kept a reference internally */
		ret = frame.AvFrameMakeWritable()
		if ret < 0 {
			os.Exit(1)
		}
		samples = (*ffcommon.FUint16T)(unsafe.Pointer(frame.Data[0]))

		for j = 0; j < c.FrameSize; j++ {
			*(*ffcommon.FUint16T)(unsafe.Pointer(uintptr(unsafe.Pointer(samples)) + uintptr(2*j*2))) = ffcommon.FUint16T(math.Sin(float64(t)) * 10000)

			for k = 1; k < c.Channels; k++ {
				*(*ffcommon.FUint16T)(unsafe.Pointer(uintptr(unsafe.Pointer(samples)) + uintptr((2*j+k)*2))) = *(*ffcommon.FUint16T)(unsafe.Pointer(uintptr(unsafe.Pointer(samples)) + uintptr(2*j*2)))
			}
			t += tincr
		}
		encode(c, frame, pkt, f)
	}

	/* flush the encoder */
	encode(c, nil, pkt, f)

	f.Close()

	libavutil.AvFrameFree(&frame)
	libavcodec.AvPacketFree(&pkt)
	libavcodec.AvcodecFreeContext(&c)

	return 0
}

/* check that a given sample format is supported by the encoder */
func check_sample_fmt(codec *libavcodec.AVCodec, sample_fmt libavutil.AVSampleFormat) ffcommon.FInt {
	p := codec.SampleFmts

	for *p != libavutil.AV_SAMPLE_FMT_NONE {
		if *p == sample_fmt {
			return 1
		}
		p = (*libavutil.AVSampleFormat)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(8)))
	}
	return 0
}

/* just pick the highest supported samplerate */
func select_sample_rate(codec *libavcodec.AVCodec) ffcommon.FInt {
	var p *ffcommon.FInt
	var best_samplerate ffcommon.FInt

	if codec.SupportedSamplerates == nil {
		return 44100
	}

	p = codec.SupportedSamplerates
	for *p != 0 {
		if best_samplerate == 0 || int32(math.Abs(float64(44100-*p))) < int32(math.Abs(float64(44100-best_samplerate))) {
			best_samplerate = *p
		}
		p = (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(4)))
	}
	return best_samplerate
}

/* select layout with the highest channel count */
func select_channel_layout(codec *libavcodec.AVCodec) ffcommon.FInt {

	var p *ffcommon.FUint64T
	var best_ch_layout ffcommon.FUint64T
	var best_nb_channels ffcommon.FInt

	if codec.ChannelLayouts == nil {
		return libavutil.AV_CH_LAYOUT_STEREO
	}

	p = codec.ChannelLayouts
	for *p != 0 {
		nb_channels := libavutil.AvGetChannelLayoutNbChannels(*p)

		if nb_channels > best_nb_channels {
			best_ch_layout = *p
			best_nb_channels = nb_channels
		}
		p = (*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(8)))
	}
	return ffcommon.FInt(best_ch_layout)
}

func encode(ctx *libavcodec.AVCodecContext, frame *libavutil.AVFrame, pkt *libavcodec.AVPacket, output *os.File) {
	var ret ffcommon.FInt

	/* send the frame for encoding */
	ret = ctx.AvcodecSendFrame(frame)
	if ret < 0 {
		fmt.Printf("Error sending the frame to the encoder\n")
		os.Exit(1)
	}

	/* read all the available output packets (in general there may be any
	 * number of them */
	for ret >= 0 {
		ret = ctx.AvcodecReceivePacket(pkt)
		if ret == -libavutil.EAGAIN || ret == libavutil.AVERROR_EOF {
			return
		} else if ret < 0 {
			fmt.Printf("Error encoding audio frame\n")
			os.Exit(1)
		}

		output.Write(ffcommon.ByteSliceFromByteP(pkt.Data, int(pkt.Size)))
		pkt.AvPacketUnref()
	}
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

	main0()
}
