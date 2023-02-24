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
	var filename, codec_name string
	var codec *libavcodec.AVCodec
	var c *libavcodec.AVCodecContext
	var i, x, y ffcommon.FInt
	var f *os.File
	var frame *libavutil.AVFrame
	var pkt *libavcodec.AVPacket
	endcode := [...]ffcommon.FUint8T{0, 0, 1, 0xb7}

	if len(os.Args) <= 2 {
		fmt.Printf("Usage: %s <output file> <codec name>\n", os.Args[0])
		return 0
	}
	filename = os.Args[1]
	codec_name = os.Args[2]

	/* find the mpeg1video encoder */
	codec = libavcodec.AvcodecFindEncoderByName(codec_name)
	if codec == nil {
		fmt.Printf("Codec '%s' not found\n", codec_name)
		os.Exit(1)
	}

	c = codec.AvcodecAllocContext3()
	if c == nil {
		fmt.Printf("Could not allocate video codec context\n")
		os.Exit(1)
	}

	pkt = libavcodec.AvPacketAlloc()
	if pkt == nil {
		os.Exit(1)
	}

	/* put sample parameters */
	c.BitRate = 400000
	/* resolution must be a multiple of two */
	c.Width = 352
	c.Height = 288
	/* frames per second */
	c.TimeBase = libavutil.AVRational{1, 25}
	c.Framerate = libavutil.AVRational{25, 1}

	/* emit one intra frame every ten frames
	 * check frame pict_type before passing frame
	 * to encoder, if frame->pict_type is AV_PICTURE_TYPE_I
	 * then gop_size is ignored and the output of encoder
	 * will always be I frame irrespective to gop_size
	 */
	c.GopSize = 10
	c.MaxBFrames = 1
	c.PixFmt = libavutil.AV_PIX_FMT_YUV420P

	if codec.Id == libavcodec.AV_CODEC_ID_H264 {
		libavutil.AvOptSet(c.PrivData, "preset", "slow", 0)
	}

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

	frame = libavutil.AvFrameAlloc()
	if frame == nil {
		fmt.Printf("Could not allocate video frame\n")
		os.Exit(1)
	}
	frame.Format = c.PixFmt
	frame.Width = c.Width
	frame.Height = c.Height

	ret = frame.AvFrameGetBuffer(0)
	if ret < 0 {
		fmt.Printf("Could not allocate the video frame data\n")
		os.Exit(1)
	}
	/* encode 1 second of video */
	for i = 0; i < 25; i++ {
		// fflush(stdout);

		/* make sure the frame data is writable */
		ret = frame.AvFrameMakeWritable()
		if ret < 0 {
			os.Exit(1)
		}

		/* prepare a dummy image */
		/* Y */
		for y = 0; y < c.Height; y++ {
			for x = 0; x < c.Width; x++ {
				*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(frame.Data[0])) + uintptr(y*frame.Linesize[0]+x))) = byte((x + y + i*3) % 256)
			}
		}

		/* Cb and Cr */
		for y = 0; y < c.Height/2; y++ {
			for x = 0; x < c.Width/2; x++ {
				*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(frame.Data[1])) + uintptr(y*frame.Linesize[1]+x))) = byte((128 + y + i*2) % 256)
				*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(frame.Data[2])) + uintptr(y*frame.Linesize[2]+x))) = byte((64 + x + i*5) % 256)
			}
		}

		frame.Pts = int64(i)

		/* encode the image */
		encode(c, frame, pkt, f)
	}

	/* flush the encoder */
	encode(c, nil, pkt, f)

	/* add sequence end code to have a real MPEG file */
	if codec.Id == libavcodec.AV_CODEC_ID_MPEG1VIDEO || codec.Id == libavcodec.AV_CODEC_ID_MPEG2VIDEO {
		f.Write(endcode[:])
	}
	f.Close()

	libavutil.AvFrameFree(&frame)
	libavcodec.AvPacketFree(&pkt)
	libavcodec.AvcodecFreeContext(&c)

	return 0
}

func encode(enc_ctx *libavcodec.AVCodecContext, frame *libavutil.AVFrame, pkt *libavcodec.AVPacket, output *os.File) {
	var ret ffcommon.FInt

	/* send the frame to the encoder */
	if frame != nil {
		fmt.Printf("Send frame %3d\n", frame.Pts)
	}

	ret = enc_ctx.AvcodecSendFrame(frame)
	if ret < 0 {
		fmt.Printf("rror sending a frame for encoding\n")
		os.Exit(1)
	}

	for ret >= 0 {
		ret = enc_ctx.AvcodecReceivePacket(pkt)
		if ret == -libavutil.EAGAIN || ret == libavutil.AVERROR_EOF {
			return
		} else if ret < 0 {
			fmt.Printf("Error during encoding\n")
			os.Exit(1)
		}
		fmt.Printf("Write packet %3d (size=%5d)\n", pkt.Pts, pkt.Size)
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
