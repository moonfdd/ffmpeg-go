package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/ffmpeg-go/libavfilter"
	"github.com/moonfdd/ffmpeg-go/libavutil"
)

func main0() (ret ffcommon.FInt) {
	var md5 *libavutil.AVMD5
	var graph *libavfilter.AVFilterGraph
	var src, sink *libavfilter.AVFilterContext
	var frame *libavutil.AVFrame
	var errstr [1024]ffcommon.FUint8T
	var duration ffcommon.FFloat
	var err, nb_frames, i ffcommon.FInt

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <duration>\n", os.Args[0])
		return 1
	}

	f, err2 := strconv.ParseFloat(os.Args[1], 32)
	if err2 != nil {
		// handle error
		return 1
	}
	duration = float32(f)
	nb_frames = int32(float64(duration) * INPUT_SAMPLERATE / FRAME_SIZE)
	if nb_frames <= 0 {
		fmt.Printf("Invalid duration: %s\n", os.Args[1])
		return 1
	}

	/* Allocate the frame we will be using to store the data. */
	frame = libavutil.AvFrameAlloc()
	if frame == nil {
		fmt.Printf("Error allocating the frame\n")
		return 1
	}

	md5 = libavutil.AvMd5Alloc()
	if md5 == nil {
		fmt.Printf("Error allocating the MD5 context\n")
		return 1
	}

	/* Set up the filtergraph. */
	err = init_filter_graph(&graph, &src, &sink)
	if err < 0 {
		fmt.Printf("Unable to init filter graph:")
		goto fail
	}

	/* the main filtering loop */
	for i = 0; i < nb_frames; i++ {
		// /* get an input frame to be filtered */
		err = get_input(frame, i)
		if err < 0 {
			fmt.Printf("Error generating input frame:")
			goto fail
		}

		/* Send the frame to the input of the filtergraph. */
		err = src.AvBuffersrcAddFrame(frame)
		if err < 0 {
			frame.AvFrameUnref()
			fmt.Printf("Error submitting the frame to the filtergraph:")
			goto fail
		}

		/* Get all the filtered output that is available. */
		err = sink.AvBuffersinkGetFrame(frame)
		for err >= 0 {
			/* now do something with our filtered frame */
			err = process_output(md5, frame)
			if err < 0 {
				fmt.Printf("Error processing the filtered frame:")
				goto fail
			}
			frame.AvFrameUnref()
			err = sink.AvBuffersinkGetFrame(frame)
		}

		if err == -libavutil.EAGAIN {
			/* Need to feed more frames in. */
			continue
		} else if err == libavutil.AVERROR_EOF {
			/* Nothing more to do, finish. */
			break
		} else if err < 0 {
			/* An error occurred. */
			fmt.Printf("Error filtering the data:")
			goto fail
		}
	}

	libavfilter.AvfilterGraphFree(&graph)
	libavutil.AvFrameFree(&frame)
	libavutil.AvFreep(uintptr(unsafe.Pointer(&md5)))

	return 0

fail:
	libavutil.AvStrerror(err, (*byte)(unsafe.Pointer(&errstr)), uint64(len(errstr)))
	fmt.Printf("%s\n", errstr)
	return 1
}

const INPUT_SAMPLERATE = 48000
const INPUT_FORMAT = libavutil.AV_SAMPLE_FMT_FLTP
const INPUT_CHANNEL_LAYOUT = libavutil.AV_CH_LAYOUT_5POINT0

const VOLUME_VAL = 0.90

func init_filter_graph(graph **libavfilter.AVFilterGraph, src **libavfilter.AVFilterContext,
	sink **libavfilter.AVFilterContext) ffcommon.FInt {
	var filter_graph *libavfilter.AVFilterGraph
	var abuffer_ctx *libavfilter.AVFilterContext
	var abuffer *libavfilter.AVFilter
	var volume_ctx *libavfilter.AVFilterContext
	var volume *libavfilter.AVFilter
	var aformat_ctx *libavfilter.AVFilterContext
	var aformat *libavfilter.AVFilter
	var abuffersink_ctx *libavfilter.AVFilterContext
	var abuffersink *libavfilter.AVFilter

	var options_dict *libavutil.AVDictionary
	var options_str string
	var ch_layout [64]ffcommon.FUint8T

	var err ffcommon.FInt

	/* Create a new filtergraph, which will contain all the filters. */
	filter_graph = libavfilter.AvfilterGraphAlloc()
	if filter_graph == nil {
		fmt.Printf("Unable to create filter graph.\n")
		return libavutil.ENOMEM
	}

	/* Create the abuffer filter;
	 * it will be used for feeding the data into the graph. */
	abuffer = libavfilter.AvfilterGetByName("abuffer")
	if abuffer == nil {
		fmt.Printf("Could not find the abuffer filter.\n")
		return libavutil.AVERROR_FILTER_NOT_FOUND
	}

	abuffer_ctx = filter_graph.AvfilterGraphAllocFilter(abuffer, "src")
	if abuffer_ctx == nil {
		fmt.Printf("Could not allocate the abuffer instance.\n")
		return -libavutil.ENOMEM
	}

	/* Set the filter options through the AVOptions API. */
	libavutil.AvGetChannelLayoutString((*byte)(unsafe.Pointer(&ch_layout)), int32(len(ch_layout)), 0, INPUT_CHANNEL_LAYOUT)
	libavutil.AvOptSet(uintptr(unsafe.Pointer(abuffer_ctx)), "channel_layout", ffcommon.StringFromPtr(uintptr(unsafe.Pointer(&ch_layout))), libavutil.AV_OPT_SEARCH_CHILDREN)
	libavutil.AvOptSet(uintptr(unsafe.Pointer(abuffer_ctx)), "sample_fmt", libavutil.AvGetSampleFmtName(INPUT_FORMAT), libavutil.AV_OPT_SEARCH_CHILDREN)
	libavutil.AvOptSetQ(uintptr(unsafe.Pointer(abuffer_ctx)), "time_base", libavutil.AVRational{1, INPUT_SAMPLERATE}, libavutil.AV_OPT_SEARCH_CHILDREN)
	libavutil.AvOptSetInt(uintptr(unsafe.Pointer(abuffer_ctx)), "sample_rate", INPUT_SAMPLERATE, libavutil.AV_OPT_SEARCH_CHILDREN)

	/* Now initialize the filter; we pass NULL options, since we have already
	 * set all the options above. */
	err = abuffer_ctx.AvfilterInitStr("")
	if err < 0 {
		fmt.Printf("Could not initialize the abuffer filter.\n")
		return err
	}

	/* Create volume filter. */
	volume = libavfilter.AvfilterGetByName("volume")
	if volume == nil {
		fmt.Printf("Could not find the volume filter.\n")
		return libavutil.AVERROR_FILTER_NOT_FOUND
	}

	volume_ctx = filter_graph.AvfilterGraphAllocFilter(volume, "volume")
	if volume_ctx == nil {
		fmt.Printf("Could not allocate the volume instance.\n")
		return -libavutil.ENOMEM
	}

	/* A different way of passing the options is as key/value pairs in a
	 * dictionary. */
	libavutil.AvDictSet(&options_dict, "volume", fmt.Sprint(VOLUME_VAL), 0)
	err = volume_ctx.AvfilterInitDict(&options_dict)
	libavutil.AvDictFree(&options_dict)
	if err < 0 {
		fmt.Printf("Could not initialize the volume filter.\n")
		return err
	}

	/* Create the aformat filter;
	 * it ensures that the output is of the format we want. */
	aformat = libavfilter.AvfilterGetByName("aformat")
	if aformat == nil {
		fmt.Printf("Could not find the aformat filter.\n")
		return libavutil.AVERROR_FILTER_NOT_FOUND
	}

	aformat_ctx = filter_graph.AvfilterGraphAllocFilter(aformat, "aformat")
	if aformat_ctx == nil {
		fmt.Printf("Could not allocate the aformat instance.\n")
		return -libavutil.ENOMEM
	}

	/* A third way of passing the options is in a string of the form
	 * key1=value1:key2=value2.... */
	// snprintf(options_str, sizeof(options_str),
	//          "sample_fmts=%s:sample_rates=%d:channel_layouts=0x%"PRIx64,
	//          av_get_sample_fmt_name(AV_SAMPLE_FMT_S16), 44100,
	//          (uint64_t)AV_CH_LAYOUT_STEREO);
	options_str = fmt.Sprintf("sample_fmts=%s:sample_rates=%d:channel_layouts=0x%x",
		libavutil.AvGetSampleFmtName(libavutil.AV_SAMPLE_FMT_S16), 44100,
		libavutil.AV_CH_LAYOUT_STEREO)
	fmt.Println(options_str)
	err = aformat_ctx.AvfilterInitStr(options_str)
	if err < 0 {
		libavutil.AvLog(uintptr(0), libavutil.AV_LOG_ERROR, "Could not initialize the aformat filter.\n")
		return err
	}

	/* Finally create the abuffersink filter;
	 * it will be used to get the filtered data out of the graph. */
	abuffersink = libavfilter.AvfilterGetByName("abuffersink")
	if abuffersink == nil {
		fmt.Printf("Could not find the abuffersink filter.\n")
		return libavutil.AVERROR_FILTER_NOT_FOUND
	}

	abuffersink_ctx = filter_graph.AvfilterGraphAllocFilter(abuffersink, "sink")
	if abuffersink_ctx == nil {
		fmt.Printf("Could not allocate the abuffersink instance.\n")
		return -libavutil.ENOMEM
	}

	/* This filter takes no options. */
	err = abuffersink_ctx.AvfilterInitStr("")
	if err < 0 {
		fmt.Printf("Could not initialize the abuffersink instance.\n")
		return err
	}

	/* Connect the filters;
	 * in this simple case the filters just form a linear chain. */
	err = abuffer_ctx.AvfilterLink(0, volume_ctx, 0)
	if err >= 0 {
		err = volume_ctx.AvfilterLink(0, aformat_ctx, 0)
	}
	if err >= 0 {
		err = aformat_ctx.AvfilterLink(0, abuffersink_ctx, 0)
	}
	if err < 0 {
		fmt.Printf("Error connecting filters\n")
		return err
	}

	/* Configure the graph. */
	err = filter_graph.AvfilterGraphConfig(uintptr(0))
	if err < 0 {
		libavutil.AvLog(uintptr(0), libavutil.AV_LOG_ERROR, "Error configuring the filter graph\n")
		return err
	}

	*graph = filter_graph
	*src = abuffer_ctx
	*sink = abuffersink_ctx

	return 0
}

/* Do something useful with the filtered data: this simple
 * example just prints the MD5 checksum of each plane to stdout. */
func process_output(md5 *libavutil.AVMD5, frame *libavutil.AVFrame) ffcommon.FInt {
	planar := libavutil.AvSampleFmtIsPlanar(libavutil.AVSampleFormat(frame.Format))
	channels := libavutil.AvGetChannelLayoutNbChannels(frame.ChannelLayout)
	planes := channels
	if planar == 0 {
		planes = 1
	}
	bps := libavutil.AvGetBytesPerSample(libavutil.AVSampleFormat(frame.Format))
	plane_size := bps * frame.NbSamples
	if planar == 0 {
		plane_size = plane_size * channels
	}
	var i, j ffcommon.FInt

	for i = 0; i < planes; i++ {
		var checksum [16]ffcommon.FUint8T

		md5.AvMd5Init()
		ptr := *(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(frame.ExtendedData)) + uintptr(i*8)))
		libavutil.AvMd5Sum((*byte)(unsafe.Pointer(&checksum)), (*byte)(unsafe.Pointer(ptr)), plane_size)

		fmt.Printf("plane %d: 0x", i)
		for j = 0; j < int32(len(checksum)); j++ {
			fmt.Printf("%02X", checksum[j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")

	return 0
}

const FRAME_SIZE = 1024

/* Construct a frame of audio data to be filtered;
 * this simple example just synthesizes a sine wave. */
func get_input(frame *libavutil.AVFrame, frame_num ffcommon.FInt) ffcommon.FInt {
	var err, i, j ffcommon.FInt

	// #define FRAME_SIZE 1024

	/* Set up the frame properties and allocate the buffer for the data. */
	frame.SampleRate = INPUT_SAMPLERATE
	frame.Format = INPUT_FORMAT
	frame.ChannelLayout = INPUT_CHANNEL_LAYOUT
	frame.NbSamples = FRAME_SIZE
	frame.Pts = int64(frame_num) * FRAME_SIZE

	err = frame.AvFrameGetBuffer(0)
	if err < 0 {
		return err
	}

	/* Fill the data for each channel. */
	for i = 0; i < 5; i++ {
		// float *data = (float*)frame->extended_data[i];
		ptr := *(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(frame.ExtendedData)) + uintptr(i*8)))
		data := (*ffcommon.FFloat)(unsafe.Pointer(ptr))

		for j = 0; j < frame.NbSamples; j++ {
			*(*ffcommon.FFloat)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(4*j))) = float32(math.Sin(2 * libavutil.M_PI * (float64(frame_num + j)) * float64((i+1)/FRAME_SIZE)))
		}
	}

	return 0
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
