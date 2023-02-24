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
	var filename, outfilename string
	var codec *libavcodec.AVCodec
	var parser *libavcodec.AVCodecParserContext
	var c *libavcodec.AVCodecContext
	var f *os.File
	var frame *libavutil.AVFrame
	var inbuf [INBUF_SIZE + libavcodec.AV_INPUT_BUFFER_PADDING_SIZE]ffcommon.FUint8T
	var data *ffcommon.FUint8T
	var data_size ffcommon.FSizeT
	var pkt *libavcodec.AVPacket

	if len(os.Args) <= 2 {
		fmt.Printf("Usage: %s <input file> <output file>\nAnd check your input file is encoded by mpeg1video please.\n", os.Args[0])
		os.Exit(0)
	}
	filename = os.Args[1]
	outfilename = os.Args[2]

	pkt = libavcodec.AvPacketAlloc()
	if pkt == nil {
		os.Exit(1)
	}

	/* set end of buffer to 0 (this ensures that no overreading happens for damaged MPEG streams) */
	//memset(inbuf + INBUF_SIZE, 0, AV_INPUT_BUFFER_PADDING_SIZE);

	/* find the MPEG-1 video decoder */
	codec = libavcodec.AvcodecFindDecoder(libavcodec.AV_CODEC_ID_MPEG1VIDEO)
	if codec == nil {
		fmt.Printf("Codec not found\n")
		os.Exit(1)
	}

	parser = libavcodec.AvParserInit(int32(codec.Id))
	if parser == nil {
		fmt.Printf("parser not found\n")
		os.Exit(1)
	}

	c = codec.AvcodecAllocContext3()
	if c == nil {
		fmt.Printf("Could not allocate video codec context\n")
		os.Exit(1)
	}

	/* For some codecs, such as msmpeg4 and mpeg4, width and height
	   MUST be initialized there because this information is not
	   available in the bitstream. */

	/* open it */
	if c.AvcodecOpen2(codec, nil) < 0 {
		fmt.Printf("Could not open codec\n")
		os.Exit(1)
	}

	var err error
	f, err = os.Open(filename)
	if err != nil {
		fmt.Printf("Could not open %s,err = %s\n", filename, err)
		os.Exit(1)
	}

	frame = libavutil.AvFrameAlloc()
	if frame == nil {
		fmt.Printf("Could not allocate video frame\n")
		os.Exit(1)
	}

	for {
		/* read raw data from the input file */
		var n int
		n, err = f.Read(inbuf[:INBUF_SIZE])
		if err != nil {
			break
		}
		data_size = uint64(n)
		if data_size == 0 {
			break
		}

		/* use the parser to split the data into frames */
		data = (*byte)(unsafe.Pointer(&inbuf))
		for data_size > 0 {
			ret = parser.AvParserParse2(c, &pkt.Data, (*int32)(unsafe.Pointer(&pkt.Size)),
				data, int32(data_size), libavutil.AV_NOPTS_VALUE, libavutil.AV_NOPTS_VALUE, 0)
			if ret < 0 {
				fmt.Printf("Error while parsing\n")
				os.Exit(1)
			}
			data = (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(ret)))
			data_size -= uint64(ret)

			if pkt.Size != 0 {
				decode(c, frame, pkt, outfilename)
			}
		}
	}

	/* flush the decoder */
	decode(c, frame, nil, outfilename)

	f.Close()

	parser.AvParserClose()
	libavcodec.AvcodecFreeContext(&c)
	libavutil.AvFrameFree(&frame)
	libavcodec.AvPacketFree(&pkt)

	return 0
}

const INBUF_SIZE = 4096

func pgm_save(buf ffcommon.FBuf, wrap, xsize, ysize ffcommon.FInt, filename string) {
	var f *os.File
	var i ffcommon.FInt

	var err error
	f, err = os.Create(filename)
	if err != nil {
		return
	}
	f.WriteString(fmt.Sprintf("P5\n%d %d\n%d\n", xsize, ysize, 255))
	bytes := []byte{}
	for i = 0; i < ysize; i++ {
		for j := int32(0); j < xsize; j++ {
			bytes = append(bytes, *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(i*wrap+j))))
		}
	}
	f.Write(bytes)
	f.Close()
}

func decode(dec_ctx *libavcodec.AVCodecContext, frame *libavutil.AVFrame, pkt *libavcodec.AVPacket, filename string) {
	// var buf [1024]byte
	var ret ffcommon.FInt

	ret = dec_ctx.AvcodecSendPacket(pkt)
	if ret < 0 {
		fmt.Printf("Error sending a packet for decoding\n")
		os.Exit(1)
	}

	for ret >= 0 {
		ret = dec_ctx.AvcodecReceiveFrame(frame)
		if ret == -libavutil.EAGAIN || ret == libavutil.AVERROR_EOF {
			return
		} else if ret < 0 {
			fmt.Printf("Error during decoding %d\n", ret)
			os.Exit(1)
		}

		fmt.Printf("saving frame %3d\n", dec_ctx.FrameNumber)
		//fflush(stdout)

		/* the picture is allocated by the decoder. no need to
		   free it */
		pgm_save(frame.Data[0], frame.Linesize[0],
			frame.Width, frame.Height, fmt.Sprintf("%s-%d.ppm", filename, dec_ctx.FrameNumber))
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
