package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/ffmpeg-go/libavcodec"
	"github.com/moonfdd/ffmpeg-go/libavutil"
)

func main0() (ret ffcommon.FInt) {
	var outfilename, filename string
	var codec *libavcodec.AVCodec
	var c *libavcodec.AVCodecContext
	var parser *libavcodec.AVCodecParserContext
	var len0 ffcommon.FInt
	var f, outfile *os.File
	var inbuf [AUDIO_INBUF_SIZE + libavcodec.AV_INPUT_BUFFER_PADDING_SIZE]ffcommon.FUint8T
	var data *ffcommon.FUint8T
	var data_size ffcommon.FSizeT
	var pkt *libavcodec.AVPacket
	var decoded_frame *libavutil.AVFrame
	var sfmt libavutil.AVSampleFormat
	var n_channels ffcommon.FInt = 0
	var fmt0 string

	if len(os.Args) <= 2 {
		fmt.Errorf("Usage: %s <input file> <output file>\n", os.Args[0])
		os.Exit(0)
	}
	filename = os.Args[1]
	outfilename = os.Args[2]

	pkt = libavcodec.AvPacketAlloc()

	/* find the MPEG audio decoder */
	codec = libavcodec.AvcodecFindDecoder(libavcodec.AV_CODEC_ID_MP2)
	if codec == nil {
		fmt.Errorf("Codec not found\n")
		os.Exit(1)
	}

	parser = libavcodec.AvParserInit(int32(codec.Id))
	if parser == nil {
		fmt.Errorf("Parser not found\n")
		os.Exit(1)
	}

	c = codec.AvcodecAllocContext3()
	if c == nil {
		fmt.Errorf("Could not allocate audio codec context\n")
		os.Exit(1)
	}

	/* open it */
	if c.AvcodecOpen2(codec, nil) < 0 {
		fmt.Errorf("Could not open codec\n")
		os.Exit(1)
	}

	var err error
	f, err = os.Open(filename)
	if err != nil {
		fmt.Errorf("Could not open %s\n", filename)
		os.Exit(1)
	}

	outfile, err = os.Create(outfilename)
	if err != nil {
		libavutil.AvFree(uintptr(unsafe.Pointer(c)))
		os.Exit(1)
	}

	/* decode until eof */
	data = (*byte)(unsafe.Pointer(&inbuf))
	var n int
	n, _ = f.Read(inbuf[0:AUDIO_INBUF_SIZE])
	data_size = uint64(n)

	for data_size > 0 {
		if decoded_frame == nil {
			decoded_frame = libavutil.AvFrameAlloc()
			if decoded_frame == nil {
				fmt.Errorf("Could not allocate audio frame\n")
				os.Exit(1)
			}
		}

		ret = parser.AvParserParse2(c, &pkt.Data, (*int32)(unsafe.Pointer(&pkt.Size)),
			data, int32(data_size),
			libavutil.AV_NOPTS_VALUE, libavutil.AV_NOPTS_VALUE, 0)
		if ret < 0 {
			fmt.Errorf("Error while parsing\n")
			os.Exit(1)
		}
		data = (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(ret)))
		data_size -= uint64(ret)

		if pkt.Size != 0 {
			decode(c, pkt, decoded_frame, outfile)
		}

		if data_size < AUDIO_REFILL_THRESH {
			for i := uint64(0); i < data_size; i++ {
				inbuf[i] = *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(i)))
			}
			data = (*byte)(unsafe.Pointer(&inbuf))
			n, _ = f.Read(inbuf[data_size:AUDIO_INBUF_SIZE])
			len0 = int32(n)
			if len0 > 0 {
				data_size += uint64(len0)
			}
		}
	}

	/* flush the decoder */
	pkt.Data = nil
	pkt.Size = 0
	decode(c, pkt, decoded_frame, outfile)

	/* print output pcm infomations, because there have no metadata of pcm */
	sfmt = c.SampleFmt

	if libavutil.AvSampleFmtIsPlanar(sfmt) != 0 {
		packed := libavutil.AvGetSampleFmtName(sfmt)
		pa := ""
		if packed == "" {
			pa = "?"
		} else {
			pa = packed
		}
		fmt.Printf("Warning: the sample format the decoder produced is planar (%s). This example will output the first channel only.\n", pa)
		sfmt = libavutil.AvGetPackedSampleFmt(sfmt)
	}

	n_channels = c.Channels
	for {
		ret = get_format_from_sample_fmt(&fmt0, sfmt)
		if ret < 0 {
			break
		}

		fmt.Printf("Play the output audio file with the command:\nffplay -f %s -ac %d -ar %d %s\n",
			fmt0, n_channels, c.SampleRate,
			outfilename)
		break
	}
	// end:
	outfile.Close()
	f.Close()

	libavcodec.AvcodecFreeContext(&c)
	parser.AvParserClose()
	libavutil.AvFrameFree(&decoded_frame)
	libavcodec.AvPacketFree(&pkt)

	return 0
}

const AUDIO_INBUF_SIZE = 20480
const AUDIO_REFILL_THRESH = 4096

func get_format_from_sample_fmt(fmt0 *string, sample_fmt libavutil.AVSampleFormat) (ret ffcommon.FInt) {
	switch sample_fmt {
	case libavutil.AV_SAMPLE_FMT_U8:
		*fmt0 = "u8"
	case libavutil.AV_SAMPLE_FMT_S16:
		*fmt0 = "s16le"
	case libavutil.AV_SAMPLE_FMT_S32:
		*fmt0 = "s32le"
	case libavutil.AV_SAMPLE_FMT_FLT:
		*fmt0 = "f32le"
	case libavutil.AV_SAMPLE_FMT_DBL:
		*fmt0 = "f64le"
	default:
		fmt.Printf("sample format %s is not supported as output format\n",
			libavutil.AvGetSampleFmtName(sample_fmt))
		ret = -1
	}
	return
}

func decode(dec_ctx *libavcodec.AVCodecContext, pkt *libavcodec.AVPacket, frame *libavutil.AVFrame, outfile *os.File) {
	var i, ch ffcommon.FInt
	var ret, data_size ffcommon.FInt

	/* send the packet with the compressed data to the decoder */
	ret = dec_ctx.AvcodecSendPacket(pkt)
	if ret < 0 {
		fmt.Printf("Error submitting the packet to the decoder\n")
		os.Exit(1)
	}

	/* read all the output frames (in general there may be any number of them */
	for ret >= 0 {
		ret = dec_ctx.AvcodecReceiveFrame(frame)
		if ret == -libavutil.EAGAIN || ret == libavutil.AVERROR_EOF {
			return
		} else if ret < 0 {
			fmt.Printf("Error during decoding\n")
			os.Exit(1)
		}
		data_size = libavutil.AvGetBytesPerSample(dec_ctx.SampleFmt)
		if data_size < 0 {
			/* This should not occur, checking just for paranoia */
			fmt.Printf("Failed to calculate data size\n")
			os.Exit(1)
		}
		bytes := []byte{}
		for i = 0; i < frame.NbSamples; i++ {
			for ch = 0; ch < dec_ctx.Channels; ch++ {
				ptr := uintptr(unsafe.Pointer(frame.Data[ch])) + uintptr(data_size*i)
				for k := int32(0); k < data_size; k++ {
					bytes = append(bytes, *(*byte)(unsafe.Pointer(ptr)))
					ptr++
				}
			}
		}
		outfile.Write(bytes)
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
