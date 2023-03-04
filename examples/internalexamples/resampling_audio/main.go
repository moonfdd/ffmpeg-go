package main

import (
	"fmt"
	"math"
	"os"
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/ffmpeg-go/libavutil"
	"github.com/moonfdd/ffmpeg-go/libswresample"
)

func main0() (ret ffcommon.FInt) {
	var src_ch_layout ffcommon.FInt64T = libavutil.AV_CH_LAYOUT_STEREO
	var dst_ch_layout ffcommon.FInt64T = libavutil.AV_CH_LAYOUT_SURROUND
	var src_rate ffcommon.FInt = 48000
	var dst_rate ffcommon.FInt = 44100
	var src_data, dst_data **ffcommon.FUint8T
	var src_nb_channels, dst_nb_channels ffcommon.FInt
	var src_linesize, dst_linesize ffcommon.FInt
	var src_nb_samples ffcommon.FInt = 1024
	var dst_nb_samples ffcommon.FInt
	var max_dst_nb_samples ffcommon.FInt
	var src_sample_fmt libavutil.AVSampleFormat = libavutil.AV_SAMPLE_FMT_DBL
	var dst_sample_fmt libavutil.AVSampleFormat = libavutil.AV_SAMPLE_FMT_S16
	var dst_filename string
	var dst_file *os.File
	var dst_bufsize ffcommon.FInt
	var fmt0 string
	var swr_ctx *libswresample.SwrContext
	var t ffcommon.FDouble

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s output_file\nAPI example program to show how to resample an audio stream with libswresample.\nThis program generates a series of audio frames, resamples them to a specified output format and rate and saves them to an output file named output_file.\n",
			os.Args[0])
		os.Exit(1)
	}
	dst_filename = os.Args[1]
	dst_file, _ = os.Create(dst_filename)
	if dst_file == nil {
		fmt.Printf("Could not open destination file %s\n", dst_filename)
		os.Exit(1)
	}

	/* create resampler context */
	swr_ctx = libswresample.SwrAlloc()
	if swr_ctx == nil {
		fmt.Printf("Could not allocate resampler context\n")
		ret = -libavutil.ENOMEM
		goto end
	}

	/* set options */
	libavutil.AvOptSetInt(uintptr(unsafe.Pointer(swr_ctx)), "in_channel_layout", src_ch_layout, 0)
	libavutil.AvOptSetInt(uintptr(unsafe.Pointer(swr_ctx)), "in_sample_rate", int64(src_rate), 0)
	libavutil.AvOptSetSampleFmt(uintptr(unsafe.Pointer(swr_ctx)), "in_sample_fmt", src_sample_fmt, 0)

	libavutil.AvOptSetInt(uintptr(unsafe.Pointer(swr_ctx)), "out_channel_layout", dst_ch_layout, 0)
	libavutil.AvOptSetInt(uintptr(unsafe.Pointer(swr_ctx)), "out_sample_rate", int64(src_rate), 0)
	libavutil.AvOptSetSampleFmt(uintptr(unsafe.Pointer(swr_ctx)), "out_sample_fmt", dst_sample_fmt, 0)

	/* initialize the resampling context */
	ret = swr_ctx.SwrInit()
	if ret < 0 {
		fmt.Printf("Failed to initialize the resampling context\n")
		goto end
	}

	/* allocate source and destination samples buffers */

	src_nb_channels = libavutil.AvGetChannelLayoutNbChannels(uint64(src_ch_layout))
	ret = libavutil.AvSamplesAllocArrayAndSamples(&src_data, &src_linesize, src_nb_channels,
		src_nb_samples, src_sample_fmt, 0)
	if ret < 0 {
		fmt.Printf("Could not allocate source samples\n")
		goto end
	}

	/* compute the number of converted samples: buffering is avoided
	 * ensuring that the output buffer will contain at least all the
	 * converted input samples */
	dst_nb_samples = int32(libavutil.AvRescaleRnd(int64(src_nb_samples), int64(dst_rate), int64(src_rate), libavutil.AV_ROUND_UP))
	max_dst_nb_samples = dst_nb_samples

	/* buffer is going to be directly written to a rawaudio file, no alignment */
	dst_nb_channels = libavutil.AvGetChannelLayoutNbChannels(uint64(dst_ch_layout))
	ret = libavutil.AvSamplesAllocArrayAndSamples(&dst_data, &dst_linesize, dst_nb_channels,
		dst_nb_samples, dst_sample_fmt, 0)
	if ret < 0 {
		fmt.Printf("Could not allocate destination samples\n")
		goto end
	}

	t = 0
	for {
		/* generate synthetic audio */
		fill_samples((*float64)(unsafe.Pointer(*src_data)), src_nb_samples, src_nb_channels, src_rate, &t)

		/* compute destination number of samples */
		dst_nb_samples = int32(libavutil.AvRescaleRnd(swr_ctx.SwrGetDelay(int64(src_rate))+
			int64(src_nb_samples), int64(dst_rate), int64(src_rate), libavutil.AV_ROUND_UP))
		if dst_nb_samples > max_dst_nb_samples {
			libavutil.AvFreep(uintptr(unsafe.Pointer(dst_data)))
			ret = libavutil.AvSamplesAlloc(dst_data, &dst_linesize, dst_nb_channels,
				dst_nb_samples, dst_sample_fmt, 1)
			if ret < 0 {
				break
			}
			max_dst_nb_samples = dst_nb_samples
		}

		/* convert to destination format */
		ret = swr_ctx.SwrConvert(dst_data, dst_nb_samples, src_data, src_nb_samples)
		if ret < 0 {
			fmt.Printf("Error while converting\n")
			goto end
		}
		dst_bufsize = libavutil.AvSamplesGetBufferSize(&dst_linesize, dst_nb_channels,
			ret, dst_sample_fmt, 1)
		if dst_bufsize < 0 {
			fmt.Printf("Could not get sample buffer size\n")
			goto end
		}
		fmt.Printf("t:%f in:%d out:%d\n", t, src_nb_samples, ret)
		dst_file.Write(ffcommon.ByteSliceFromByteP(*dst_data, int(dst_bufsize)))
		if t < 10 {

		} else {
			break
		}
	}

	ret = get_format_from_sample_fmt(&fmt0, dst_sample_fmt)
	if ret < 0 {
		goto end
	}
	fmt.Printf("Resampling succeeded. Play the output file with the command:\nffplay -f %s -channel_layout %d -channels %d -ar %d %s\n",
		fmt0, dst_ch_layout, dst_nb_channels, dst_rate, dst_filename)

end:
	dst_file.Close()

	if src_data != nil {
		libavutil.AvFreep(uintptr(unsafe.Pointer(src_data)))
	}
	libavutil.AvFreep(uintptr(unsafe.Pointer(&src_data)))

	if dst_data != nil {
		libavutil.AvFreep(uintptr(unsafe.Pointer(dst_data)))
	}
	libavutil.AvFreep(uintptr(unsafe.Pointer(&dst_data)))

	libswresample.SwrFree(&swr_ctx)
	if ret < 0 {
		return 1
	} else {
		return 0
	}
}

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

/**
* Fill dst buffer with nb_samples, generated starting from t.
 */
func fill_samples(dst *ffcommon.FDouble, nb_samples, nb_channels, sample_rate ffcommon.FInt, t *ffcommon.FDouble) {
	var i, j ffcommon.FInt
	tincr := 1.0 / float64(sample_rate)
	dstp := dst
	c := 2 * libavutil.M_PI * 440.0

	/* generate sin tone with 440Hz frequency and duplicated channels */
	for i = 0; i < nb_samples; i++ {
		*dstp = math.Sin(c * *t)
		for j = 1; j < nb_channels; j++ {
			*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(dstp)) + uintptr(8*j))) = *dstp
		}
		dstp = (*ffcommon.FDouble)(unsafe.Pointer(uintptr(unsafe.Pointer(dstp)) + uintptr(8*nb_channels)))
		*t += tincr
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
