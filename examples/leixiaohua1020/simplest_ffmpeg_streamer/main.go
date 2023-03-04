package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/ffmpeg-go/libavcodec"
	"github.com/moonfdd/ffmpeg-go/libavformat"
	"github.com/moonfdd/ffmpeg-go/libavutil"
)

func main0() (ret ffcommon.FInt) {
	var ofmt *libavformat.AVOutputFormat
	//Input AVFormatContext and Output AVFormatContext
	var ifmt_ctx, ofmt_ctx *libavformat.AVFormatContext
	var pkt libavcodec.AVPacket
	var in_filename, out_filename string
	var i ffcommon.FInt
	var videoindex ffcommon.FInt = -1
	var frame_index ffcommon.FInt = 0
	var start_time ffcommon.FInt64T = 0
	var err error
	//in_filename  = "cuc_ieschool.mov";
	//in_filename  = "cuc_ieschool.mkv";
	//in_filename  = "cuc_ieschool.ts";
	//in_filename  = "cuc_ieschool.mp4";
	//in_filename  = "cuc_ieschool.h264";
	in_filename = "./out/cuc_ieschool.flv" //输入URL（Input file URL）
	//in_filename  = "shanghai03_p.h264";
	_, err = os.Stat(in_filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("create flv file")
			exec.Command("./lib/ffmpeg", "-i", "./resources/big_buck_bunny.mp4", "-vcodec", "copy", "-acodec", "copy", in_filename).Output()
		}
	}

	out_filename = "rtmp://localhost/publishlive/livestream" //输出 URL（Output URL）[RTMP]
	//out_filename = "rtp://233.233.233.233:6666";//输出 URL（Output URL）[UDP]

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

	libavformat.AvformatAllocOutputContext2(&ofmt_ctx, nil, "flv", out_filename) //RTMP
	//avformat_alloc_output_context2(&ofmt_ctx, NULL, "mpegts", out_filename);//UDP

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

	start_time = libavutil.AvGettime()
	for {
		var in_stream, out_stream *libavformat.AVStream
		//Get an AVPacket
		ret = ifmt_ctx.AvReadFrame(&pkt)
		if ret < 0 {
			break
		}
		//FIX：No PTS (Example: Raw H.264)
		//Simple Write PTS
		if pkt.Pts == libavutil.AV_NOPTS_VALUE {
			//Write PTS
			time_base1 := ifmt_ctx.GetStream(uint32(videoindex)).TimeBase
			//Duration between 2 frames (us)
			calc_duration := int64(libavutil.AV_TIME_BASE / libavutil.AvQ2d(ifmt_ctx.GetStream(uint32(videoindex)).RFrameRate))
			//Parameters
			pkt.Pts = int64(float64(frame_index) * float64(calc_duration) / (libavutil.AvQ2d(time_base1) * libavutil.AV_TIME_BASE))
			pkt.Dts = pkt.Pts
			pkt.Duration = int64(float64(calc_duration) / (libavutil.AvQ2d(time_base1) * libavutil.AV_TIME_BASE))
		}
		//Important:Delay
		if pkt.StreamIndex == uint32(videoindex) {
			time_base := ifmt_ctx.GetStream(uint32(videoindex)).TimeBase
			time_base_q := libavutil.AVRational{1, libavutil.AV_TIME_BASE}
			pts_time := libavutil.AvRescaleQ(pkt.Dts, time_base, time_base_q)
			now_time := libavutil.AvGettime() - start_time
			if pts_time > now_time {
				libavutil.AvUsleep(uint32(pts_time - now_time))
			}

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
			fmt.Printf("Send %8d video frames to output URL\n", frame_index)
			frame_index++
		}
		//ret = av_write_frame(ofmt_ctx, &pkt);
		ret = ofmt_ctx.AvInterleavedWriteFrame(&pkt)

		if ret < 0 {
			fmt.Printf("Error muxing packet\n")
			break
		}

		pkt.AvFreePacket()

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

	go func() {
		time.Sleep(1000)
		exec.Command("./lib/ffplay.exe", "rtmp://localhost/publishlive/livestream").Output()
		if err != nil {
			fmt.Println("play err = ", err)
		}
	}()

	main0()
}
