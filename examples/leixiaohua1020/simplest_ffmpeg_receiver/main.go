// https://github.com/leixiaohua1020/simplest_ffmpeg_streamer/blob/master/simplest_ffmpeg_receiver/simplest_ffmpeg_receiver.cpp
package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/ffmpeg-go/libavcodec"
	"github.com/moonfdd/ffmpeg-go/libavformat"
	"github.com/moonfdd/ffmpeg-go/libavutil"
)

// '1': Use H.264 Bitstream Filter
const USE_H264BSF = 0

func main0() (ret ffcommon.FInt) {
	var ofmt *libavformat.AVOutputFormat
	//Input AVFormatContext and Output AVFormatContext
	var ifmt_ctx, ofmt_ctx *libavformat.AVFormatContext
	var pkt libavcodec.AVPacket
	var in_filename, out_filename string
	var i ffcommon.FInt
	var videoindex ffcommon.FInt = -1
	var frame_index ffcommon.FInt = 0
	var h264bsfc *libavcodec.AVBitStreamFilterContext
	in_filename = "rtmp://localhost/publishlive/livestream"
	//in_filename  = "rtp://233.233.233.233:6666";
	//out_filename = "receive.ts";
	//out_filename = "receive.mkv";
	out_filename = "./out/receive.flv"

	libavformat.AvRegisterAll()
	//Network
	libavformat.AvformatNetworkInit()
	//Input
	ret = libavformat.AvformatOpenInput(&ifmt_ctx, in_filename, nil, nil)
	if ret < 0 {
		fmt.Printf("Could not open input file.")
		goto end
	}
	ret = ifmt_ctx.AvformatFindStreamInfo(nil)
	if ret < 0 {
		fmt.Printf("Failed to retrieve input stream information")
		goto end
	}

	for i = 0; i < int32(ifmt_ctx.NbStreams); i++ {
		if ifmt_ctx.GetStream(uint32(i)).Codec.CodecType == libavutil.AVMEDIA_TYPE_VIDEO {
			videoindex = i
			break
		}
	}

	ifmt_ctx.AvDumpFormat(0, in_filename, 0)

	//Output
	libavformat.AvformatAllocOutputContext2(&ofmt_ctx, nil, "", out_filename) //RTMP

	if ofmt_ctx == nil {
		fmt.Printf("Could not create output context\n")
		ret = libavutil.AVERROR_UNKNOWN
		goto end
	}
	ofmt = ofmt_ctx.Oformat
	for i = 0; i < int32(ifmt_ctx.NbStreams); i++ {
		//Create output AVStream according to input AVStream
		in_stream := ifmt_ctx.GetStream(uint32(i))
		out_stream := ofmt_ctx.AvformatNewStream(in_stream.Codec.Codec)
		if out_stream == nil {
			fmt.Printf("Failed allocating output stream\n")
			ret = libavutil.AVERROR_UNKNOWN
			goto end
		}
		//Copy the settings of AVCodecContext
		ret = libavcodec.AvcodecCopyContext(out_stream.Codec, in_stream.Codec)
		if ret < 0 {
			fmt.Printf("Failed to copy context from input to output stream codec context\n")
			goto end
		}
		out_stream.Codec.CodecTag = 0
		if ofmt_ctx.Oformat.Flags&libavformat.AVFMT_GLOBALHEADER != 0 {
			out_stream.Codec.Flags |= libavcodec.AV_CODEC_FLAG_GLOBAL_HEADER
		}
	}
	//Dump Format------------------
	ofmt_ctx.AvDumpFormat(0, out_filename, 1)
	//Open output URL
	if ofmt.Flags&libavformat.AVFMT_NOFILE == 0 {
		ret = libavformat.AvioOpen(&ofmt_ctx.Pb, out_filename, libavformat.AVIO_FLAG_WRITE)
		if ret < 0 {
			fmt.Printf("Could not open output URL '%s'", out_filename)
			goto end
		}
	}
	//Write file header
	ret = ofmt_ctx.AvformatWriteHeader(nil)
	if ret < 0 {
		fmt.Printf("Error occurred when opening output URL\n")
		goto end
	}

	if USE_H264BSF != 0 {
		h264bsfc = libavcodec.AvBitstreamFilterInit("h264_mp4toannexb")
	}

	for {
		var in_stream, out_stream *libavformat.AVStream
		//Get an AVPacket
		ret = ifmt_ctx.AvReadFrame(&pkt)
		if ret < 0 {
			break
		}

		in_stream = ifmt_ctx.GetStream(pkt.StreamIndex)
		out_stream = ofmt_ctx.GetStream(pkt.StreamIndex)
		/* copy packet */
		//Convert PTS/DTS
		pkt.Pts = libavutil.AvRescaleQRnd(pkt.Pts, in_stream.TimeBase, out_stream.TimeBase, libavutil.AV_ROUND_NEAR_INF|libavutil.AV_ROUND_PASS_MINMAX)
		pkt.Dts = libavutil.AvRescaleQRnd(pkt.Dts, in_stream.TimeBase, out_stream.TimeBase, libavutil.AV_ROUND_NEAR_INF|libavutil.AV_ROUND_PASS_MINMAX)
		pkt.Duration = libavutil.AvRescaleQ(pkt.Duration, in_stream.TimeBase, out_stream.TimeBase)
		pkt.Pos = -1
		//Print to Screen
		if pkt.StreamIndex == uint32(videoindex) {
			fmt.Printf("Receive %8d video frames from input URL\n", frame_index)
			frame_index++

			if USE_H264BSF != 0 {
				h264bsfc.AvBitstreamFilterFilter(in_stream.Codec, "", &pkt.Data, (*int32)(unsafe.Pointer(&pkt.Size)), pkt.Data, int32(pkt.Size), 0)
			}
		}
		//ret = av_write_frame(ofmt_ctx, &pkt);
		ret = ofmt_ctx.AvInterleavedWriteFrame(&pkt)

		if ret < 0 {
			fmt.Printf("Error muxing packet\n")
			break
		}

		pkt.AvFreePacket()

	}

	if USE_H264BSF != 0 {
		h264bsfc.AvBitstreamFilterClose()
	}

	//Write file trailer
	ofmt_ctx.AvWriteTrailer()
end:
	libavformat.AvformatCloseInput(&ifmt_ctx)
	/* close output */
	if ofmt_ctx != nil && ofmt.Flags&libavformat.AVFMT_NOFILE == 0 {
		ofmt_ctx.Pb.AvioClose()
	}
	ofmt_ctx.AvformatFreeContext()
	if ret < 0 && ret != libavutil.AVERROR_EOF {
		fmt.Printf("Error occurred.\n")
		return -1
	}
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

	// go func() {
	// 	time.Sleep(1000)
	// 	exec.Command("./lib/ffplay.exe", "rtmp://localhost/publishlive/livestream").Output()
	// 	if err != nil {
	// 		fmt.Println("play err = ", err)
	// 	}
	// }()

	main0()
}
