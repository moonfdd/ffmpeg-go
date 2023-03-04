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

func main0() (ret ffcommon.FInt) {
	var ofmt *libavformat.AVOutputFormat
	var ifmt_ctx, ofmt_ctx *libavformat.AVFormatContext
	var pkt libavcodec.AVPacket
	var in_filename, out_filename string
	var i ffcommon.FInt
	var stream_index ffcommon.FInt = 0
	var stream_mapping *ffcommon.FInt
	var stream_mapping_size ffcommon.FInt = 0

	if len(os.Args) < 3 {
		fmt.Printf("usage: %s input output\nAPI example program to remux a media file with libavformat and libavcodec.\nThe output format is guessed according to the file extension.\n\n", os.Args[0])
		return 1
	}

	in_filename = os.Args[1]
	out_filename = os.Args[2]

	ret = libavformat.AvformatOpenInput(&ifmt_ctx, in_filename, nil, nil)
	if ret < 0 {
		fmt.Printf("Could not open input file '%s'", in_filename)
		goto end
	}

	ret = ifmt_ctx.AvformatFindStreamInfo(nil)
	if ret < 0 {
		fmt.Printf("Failed to retrieve input stream information")
		goto end
	}

	ifmt_ctx.AvDumpFormat(0, in_filename, 0)

	libavformat.AvformatAllocOutputContext2(&ofmt_ctx, nil, "", out_filename)
	if ofmt_ctx == nil {
		fmt.Printf("Could not create output context\n")
		ret = libavutil.AVERROR_UNKNOWN
		goto end
	}

	stream_mapping_size = int32(ifmt_ctx.NbStreams)
	stream_mapping = (*int32)(unsafe.Pointer(libavutil.AvMalloczArray(uint64(stream_mapping_size), 4)))
	if stream_mapping == nil {
		ret = -libavutil.ENOMEM
		goto end
	}

	ofmt = ofmt_ctx.Oformat

	for i = 0; i < int32(ifmt_ctx.NbStreams); i++ {
		var out_stream *libavformat.AVStream
		in_stream := ifmt_ctx.GetStream(uint32(i))
		in_codecpar := in_stream.Codecpar

		if in_codecpar.CodecType != libavutil.AVMEDIA_TYPE_AUDIO &&
			in_codecpar.CodecType != libavutil.AVMEDIA_TYPE_VIDEO &&
			in_codecpar.CodecType != libavutil.AVMEDIA_TYPE_SUBTITLE {
			*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(stream_mapping)) + uintptr(4*i))) = -1
			continue
		}

		*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(stream_mapping)) + uintptr(4*i))) = stream_index
		stream_index++

		out_stream = ofmt_ctx.AvformatNewStream(nil)
		if out_stream == nil {
			fmt.Printf("Failed allocating output stream\n")
			ret = libavutil.AVERROR_UNKNOWN
			goto end
		}

		ret = libavcodec.AvcodecParametersCopy(out_stream.Codecpar, in_codecpar)
		if ret < 0 {
			fmt.Printf("Failed to copy codec parameters\n")
			goto end
		}
		out_stream.Codecpar.CodecTag = 0
	}
	ofmt_ctx.AvDumpFormat(0, out_filename, 1)

	if ofmt.Flags&libavformat.AVFMT_NOFILE == 0 {
		ret = libavformat.AvioOpen(&ofmt_ctx.Pb, out_filename, libavformat.AVIO_FLAG_WRITE)
		if ret < 0 {
			fmt.Printf("Could not open output file '%s'", out_filename)
			goto end
		}
	}

	ret = ofmt_ctx.AvformatWriteHeader(nil)
	if ret < 0 {
		fmt.Printf("Error occurred when opening output file\n")
		goto end
	}

	for {
		var in_stream, out_stream *libavformat.AVStream

		ret = ifmt_ctx.AvReadFrame(&pkt)
		if ret < 0 {
			break
		}

		in_stream = ifmt_ctx.GetStream(pkt.StreamIndex)
		if pkt.StreamIndex >= uint32(stream_mapping_size) ||
			*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(stream_mapping)) + uintptr(4*pkt.StreamIndex))) < 0 {
			pkt.AvPacketUnref()
			continue
		}

		pkt.StreamIndex = uint32(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(stream_mapping)) + uintptr(4*pkt.StreamIndex))))
		out_stream = ofmt_ctx.GetStream(pkt.StreamIndex)
		log_packet(ifmt_ctx, &pkt, "in")

		/* copy packet */
		pkt.Pts = libavutil.AvRescaleQRnd(pkt.Pts, in_stream.TimeBase, out_stream.TimeBase, libavutil.AV_ROUND_NEAR_INF|libavutil.AV_ROUND_PASS_MINMAX)
		pkt.Dts = libavutil.AvRescaleQRnd(pkt.Dts, in_stream.TimeBase, out_stream.TimeBase, libavutil.AV_ROUND_NEAR_INF|libavutil.AV_ROUND_PASS_MINMAX)
		pkt.Duration = libavutil.AvRescaleQ(pkt.Duration, in_stream.TimeBase, out_stream.TimeBase)
		pkt.Pos = -1
		log_packet(ofmt_ctx, &pkt, "out")

		ret = ofmt_ctx.AvInterleavedWriteFrame(&pkt)
		if ret < 0 {
			fmt.Printf("Error muxing packet\n")
			break
		}
		pkt.AvPacketUnref()
	}

	ofmt_ctx.AvWriteTrailer()
end:

	libavformat.AvformatCloseInput(&ifmt_ctx)

	/* close output */
	if ofmt_ctx != nil && ofmt.Flags&libavformat.AVFMT_NOFILE == 0 {
		libavformat.AvioClosep(&ofmt_ctx.Pb)
	}
	ofmt_ctx.AvformatFreeContext()

	libavutil.AvFreep(uintptr(unsafe.Pointer(&stream_mapping)))

	if ret < 0 && ret != libavutil.AVERROR_EOF {
		fmt.Printf("Error occurred: %s\n", libavutil.AvErr2str(ret))
		return 1
	}

	return 0
}

func log_packet(fmt_ctx *libavformat.AVFormatContext, pkt *libavcodec.AVPacket, tag string) {
	time_base := &fmt_ctx.GetStream(pkt.StreamIndex).TimeBase

	fmt.Printf("%s: pts:%s pts_time:%s dts:%s dts_time:%s duration:%s duration_time:%s stream_index:%d\n",
		tag,
		libavutil.AvTs2str(pkt.Pts), libavutil.AvTs2timestr(pkt.Pts, time_base),
		libavutil.AvTs2str(pkt.Dts), libavutil.AvTs2timestr(pkt.Dts, time_base),
		libavutil.AvTs2str(pkt.Duration), libavutil.AvTs2timestr(pkt.Duration, time_base),
		pkt.StreamIndex)
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
