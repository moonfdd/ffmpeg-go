package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/ffmpeg-go/libavcodec"
	"github.com/moonfdd/ffmpeg-go/libavfilter"
	"github.com/moonfdd/ffmpeg-go/libavformat"
	"github.com/moonfdd/ffmpeg-go/libavutil"
)

func main0() (ret ffcommon.FInt) {
	var packet libavcodec.AVPacket
	frame := libavutil.AvFrameAlloc()
	filt_frame := libavutil.AvFrameAlloc()

	if frame == nil || filt_frame == nil {
		fmt.Println("Could not allocate frame")
		os.Exit(1)
	}
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s file | %s\n", os.Args[0], player)
		os.Exit(1)
	}

	ret = open_input_file(os.Args[1])
	if ret < 0 {
		goto end
	}

	ret = init_filters(filter_descr)
	if ret < 0 {
		goto end
	}

	/* read all packets */
	for {
		ret = fmt_ctx.AvReadFrame(&packet)
		if ret < 0 {
			break
		}

		if int32(packet.StreamIndex) == audio_stream_index {
			ret = dec_ctx.AvcodecSendPacket(&packet)
			if ret < 0 {
				libavutil.AvLog(uintptr(0), libavutil.AV_LOG_ERROR, "Error while sending a packet to the decoder\n")
				break
			}

			for ret >= 0 {
				ret = dec_ctx.AvcodecReceiveFrame(frame)
				if ret == -libavutil.EAGAIN || ret == libavutil.AVERROR_EOF {
					break
				} else if ret < 0 {
					libavutil.AvLog(uintptr(0), libavutil.AV_LOG_ERROR, "Error while receiving a frame from the decoder\n")
					goto end
				}

				if ret >= 0 {
					/* push the audio data from decoded frame into the filtergraph */
					if buffersrc_ctx.AvBuffersrcAddFrameFlags(frame, libavfilter.AV_BUFFERSRC_FLAG_KEEP_REF) < 0 {
						libavutil.AvLog(uintptr(0), libavutil.AV_LOG_ERROR, "Error while feeding the audio filtergraph\n")
						break
					}

					/* pull filtered audio from the filtergraph */
					for {
						ret = buffersink_ctx.AvBuffersinkGetFrame(filt_frame)
						if ret == -libavutil.EAGAIN || ret == libavutil.AVERROR_EOF {
							break
						}
						if ret < 0 {
							goto end
						}
						print_frame(filt_frame)
						filt_frame.AvFrameUnref()
					}
					frame.AvFrameUnref()
				}
			}
		}
		packet.AvPacketUnref()
	}
end:
	libavfilter.AvfilterGraphFree(&filter_graph)
	libavcodec.AvcodecFreeContext(&dec_ctx)
	libavformat.AvformatCloseInput(&fmt_ctx)
	libavutil.AvFrameFree(&frame)
	libavutil.AvFrameFree(&filt_frame)
	if ret < 0 && ret != libavutil.AVERROR_EOF {
		fmt.Printf("Error occurred: %s\n", libavutil.AvErr2str(ret))
		os.Exit(1)
	}

	os.Exit(0)
	return 0
}

var filter_descr = "aresample=8000,aformat=sample_fmts=s16:channel_layouts=mono"
var player = "ffplay -f s16le -ar 8000 -ac 1 -"

var fmt_ctx *libavformat.AVFormatContext
var dec_ctx *libavcodec.AVCodecContext
var buffersink_ctx *libavfilter.AVFilterContext
var buffersrc_ctx *libavfilter.AVFilterContext
var filter_graph *libavfilter.AVFilterGraph
var audio_stream_index ffcommon.FInt = -1

func open_input_file(filename string) ffcommon.FInt {
	var ret ffcommon.FInt
	var dec *libavcodec.AVCodec

	ret = libavformat.AvformatOpenInput(&fmt_ctx, filename, nil, nil)
	if ret < 0 {
		libavutil.AvLog(uintptr(0), libavutil.AV_LOG_ERROR, "Cannot open input file\n")
		return ret
	}

	ret = fmt_ctx.AvformatFindStreamInfo(nil)
	if ret < 0 {
		libavutil.AvLog(uintptr(0), libavutil.AV_LOG_ERROR, "Cannot find stream information\n")
		return ret
	}

	/* select the audio stream */
	ret = fmt_ctx.AvFindBestStream(libavutil.AVMEDIA_TYPE_AUDIO, -1, -1, &dec, 0)
	if ret < 0 {
		libavutil.AvLog(uintptr(0), libavutil.AV_LOG_ERROR, "Cannot find an audio stream in the input file\n")
		return ret
	}
	audio_stream_index = ret

	/* create decoding context */
	dec_ctx = dec.AvcodecAllocContext3()
	if dec_ctx == nil {
		return -libavutil.ENOMEM
	}
	dec_ctx.AvcodecParametersToContext(fmt_ctx.GetStream(uint32(audio_stream_index)).Codecpar)

	/* init the audio decoder */
	ret = dec_ctx.AvcodecOpen2(dec, nil)
	if ret < 0 {
		libavutil.AvLog(uintptr(0), libavutil.AV_LOG_ERROR, "Cannot open audio decoder\n")
		return ret
	}

	return 0
}

func init_filters(filters_descr string) ffcommon.FInt {
	var args string
	var args0 [512]byte
	var ret ffcommon.FInt = 0
	abuffersrc := libavfilter.AvfilterGetByName("abuffer")
	abuffersink := libavfilter.AvfilterGetByName("abuffersink")
	outputs := libavfilter.AvfilterInoutAlloc()
	inputs := libavfilter.AvfilterInoutAlloc()
	out_sample_fmts := [...]libavutil.AVSampleFormat{libavutil.AV_SAMPLE_FMT_S16, -1}
	out_channel_layouts := [...]ffcommon.FInt64T{libavutil.AV_CH_LAYOUT_MONO, -1}
	out_sample_rates := [...]ffcommon.FInt{8000, -1}
	var outlink *libavfilter.AVFilterLink
	time_base := fmt_ctx.GetStream(uint32(audio_stream_index)).TimeBase
	f := ""
	ii := int64(-1)

	filter_graph = libavfilter.AvfilterGraphAlloc()
	if outputs == nil || inputs == nil || filter_graph == nil {
		ret = -libavutil.ENOMEM
		goto end
	}

	/* buffer audio source: the decoded frames from the decoder will be inserted here. */
	if dec_ctx.ChannelLayout == 0 {
		dec_ctx.ChannelLayout = uint64(libavutil.AvGetDefaultChannelLayout(dec_ctx.Channels))
	}
	args = fmt.Sprintf("time_base=%d/%d:sample_rate=%d:sample_fmt=%s:channel_layout=0x0%x",
		time_base.Num, time_base.Den, dec_ctx.SampleRate,
		libavutil.AvGetSampleFmtName(dec_ctx.SampleFmt), dec_ctx.ChannelLayout)
	//fmt.Println("args = ", args)
	copy(args0[:], args[:])
	ret = libavfilter.AvfilterGraphCreateFilter(&buffersrc_ctx, abuffersrc, "in",
		args, uintptr(0), filter_graph)
	if ret < 0 {
		libavutil.AvLog(uintptr(0), libavutil.AV_LOG_ERROR, "Cannot create audio buffer source\n")
		goto end
	}

	/* buffer audio sink: to terminate the filter chain. */
	ret = libavfilter.AvfilterGraphCreateFilter(&buffersink_ctx, abuffersink, "out",
		"", uintptr(0), filter_graph)
	if ret < 0 {
		libavutil.AvLog(uintptr(0), libavutil.AV_LOG_ERROR, "Cannot create audio buffer sink\n")
		goto end
	}

	ret = libavutil.AvOptSetIntList(uintptr(unsafe.Pointer(buffersink_ctx)), "sample_fmts", uintptr(unsafe.Pointer(&out_sample_fmts)), 4, uint64(uintptr(ii)),
		libavutil.AV_OPT_SEARCH_CHILDREN)
	if ret < 0 {
		libavutil.AvLog(uintptr(0), libavutil.AV_LOG_ERROR, "Cannot set output sample format\n")
		goto end
	}

	ret = libavutil.AvOptSetIntList(uintptr(unsafe.Pointer(buffersink_ctx)), "channel_layouts", uintptr(unsafe.Pointer(&out_channel_layouts)), 8, uint64(uintptr(ii)),
		libavutil.AV_OPT_SEARCH_CHILDREN)
	if ret < 0 {
		libavutil.AvLog(uintptr(0), libavutil.AV_LOG_ERROR, "Cannot set output channel layout\n")
		goto end
	}

	ret = libavutil.AvOptSetIntList(uintptr(unsafe.Pointer(buffersink_ctx)), "sample_rates", uintptr(unsafe.Pointer(&out_sample_rates)), 4, uint64(uintptr(ii)),
		libavutil.AV_OPT_SEARCH_CHILDREN)
	if ret < 0 {
		libavutil.AvLog(uintptr(0), libavutil.AV_LOG_ERROR, "Cannot set output sample rate\n")
		goto end
	}

	/*
	 * Set the endpoints for the filter graph. The filter_graph will
	 * be linked to the graph described by filters_descr.
	 */

	/*
	 * The buffer source output must be connected to the input pad of
	 * the first filter described by filters_descr; since the first
	 * filter input label is not specified, it is set to "in" by
	 * default.
	 */
	outputs.Name = ffcommon.UintPtrFromString(libavutil.AvStrdup("in"))
	// outputs.Name = ffcommon.UintPtrFromString("in")
	outputs.FilterCtx = buffersrc_ctx
	outputs.PadIdx = 0
	outputs.Next = nil

	/*
	 * The buffer sink input must be connected to the output pad of
	 * the last filter described by filters_descr; since the last
	 * filter output label is not specified, it is set to "out" by
	 * default.
	 */
	inputs.Name = ffcommon.UintPtrFromString(libavutil.AvStrdup("out"))
	// inputs.Name = ffcommon.UintPtrFromString("out")
	inputs.FilterCtx = buffersink_ctx
	inputs.PadIdx = 0
	inputs.Next = nil

	// ret = filter_graph.AvfilterGraphParsePtr(filters_descr,
	// 	&inputs, &outputs, uintptr(0))
	fmt.Println("filters_descr = ", filters_descr)
	ret = filter_graph.AvfilterGraphParsePtr(filters_descr,
		&inputs, &outputs, uintptr(0))
	if ret < 0 {
		goto end
	}
	ret = filter_graph.AvfilterGraphConfig(uintptr(0))
	if ret < 0 {
		goto end
	}

	/* Print summary of the sink buffer
	 * Note: args buffer is reused to store channel layout string */
	outlink = buffersink_ctx.GetInput(0)
	libavutil.AvGetChannelLayoutString((*byte)(unsafe.Pointer(&args0)), int32(len(args0)), -1, outlink.ChannelLayout)
	f = libavutil.AvGetSampleFmtName(libavutil.AVSampleFormat(outlink.Format))
	if f == "" {
		f = "?"
	}
	libavutil.AvLog(uintptr(0), libavutil.AV_LOG_INFO, "Output: srate:%sHz fmt:%s chlayout:%s\n",
		fmt.Sprint(outlink.SampleRate),
		f,
		ffcommon.StringFromPtr(uintptr(unsafe.Pointer(&args0))))

end:
	libavfilter.AvfilterInoutFree(&inputs)
	libavfilter.AvfilterInoutFree(&outputs)

	return ret
}

func print_frame(frame *libavutil.AVFrame) {
	n := frame.NbSamples * libavutil.AvGetChannelLayoutNbChannels(frame.ChannelLayout)
	p := uintptr(unsafe.Pointer(frame.Data[0]))
	p_end := p + uintptr(2*n)

	for p < p_end {
		fmt.Print(string([]byte{byte(*(*ffcommon.FInt16T)(unsafe.Pointer(p)) & 0xff), byte(*(*ffcommon.FInt16T)(unsafe.Pointer(p)) >> 8 & 0xff)}))
		p += 2
	}
	// fflush(stdout);
}

func main() {
	os.Setenv("Path", os.Getenv("Path")+";./lib")
	ffcommon.SetAvutilPath("./lib/avutil-56.dll")
	ffcommon.SetAvcodecPath("./lib/avcodec-58.dll")
	ffcommon.SetAvdevicePath("./lib/avdevice-58.dll")
	ffcommon.SetAvfilterPath("./lib/avfilter-7.dll")
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
