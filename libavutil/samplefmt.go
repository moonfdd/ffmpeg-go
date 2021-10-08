package libavutil

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"syscall"
	"unsafe"
)

/**
 * Return the name of sample_fmt, or NULL if sample_fmt is not
 * recognized.
 */
//const char *av_get_sample_fmt_name(enum AVSampleFormat sample_fmt);
//未测试
func AvGetSampleFmtName(sample_fmt ffconstant.AVSampleFormat) (res ffcommon.FConstCharP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_get_sample_fmt_name").Call(
		uintptr(sample_fmt),
	)
	if err != nil {
		//return
	}
	res = ffcommon.GoAStr(t)
	return
}

/**
 * Return a sample format corresponding to name, or AV_SAMPLE_FMT_NONE
 * on error.
 */
//enum AVSampleFormat av_get_sample_fmt(const char *name);
//未测试
func AvGetSampleFmt(name ffcommon.FConstCharP) (res ffconstant.AVSampleFormat, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_get_sample_fmt").Call(
		uintptr(unsafe.Pointer(namep)),
	)
	if err != nil {
		//return
	}
	res = ffconstant.AVSampleFormat(t)
	return
}

/**
 * Return the planar<->packed alternative form of the given sample format, or
 * AV_SAMPLE_FMT_NONE on error. If the passed sample_fmt is already in the
 * requested planar/packed format, the format returned is the same as the
 * input.
 */
//enum AVSampleFormat av_get_alt_sample_fmt(enum AVSampleFormat sample_fmt, int planar);
//未测试
func AvGetAltSampleFmt(sample_fmt ffconstant.AVSampleFormat, planar ffcommon.FInt) (res ffconstant.AVSampleFormat, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_get_alt_sample_fmt").Call(
		uintptr(sample_fmt),
		uintptr(planar),
	)
	if err != nil {
		//return
	}
	res = ffconstant.AVSampleFormat(t)
	return
}

/**
* Get the packed alternative form of the given sample format.
*
* If the passed sample_fmt is already in packed format, the format returned is
* the same as the input.
*
* @return  the packed alternative form of the given sample format or
           AV_SAMPLE_FMT_NONE on error.
*/
//enum AVSampleFormat av_get_packed_sample_fmt(enum AVSampleFormat sample_fmt);
//未测试
func AvGetPackedSampleFmt(sample_fmt ffconstant.AVSampleFormat) (res ffconstant.AVSampleFormat, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_get_packed_sample_fmt").Call(
		uintptr(sample_fmt),
	)
	if err != nil {
		//return
	}
	res = ffconstant.AVSampleFormat(t)
	return
}

/**
* Get the planar alternative form of the given sample format.
*
* If the passed sample_fmt is already in planar format, the format returned is
* the same as the input.
*
* @return  the planar alternative form of the given sample format or
           AV_SAMPLE_FMT_NONE on error.
*/
//enum AVSampleFormat av_get_planar_sample_fmt(enum AVSampleFormat sample_fmt);
//未测试
func AvGetPlanarSampleFmt(sample_fmt ffconstant.AVSampleFormat) (res ffconstant.AVSampleFormat, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_get_planar_sample_fmt").Call(
		uintptr(sample_fmt),
	)
	if err != nil {
		//return
	}
	res = ffconstant.AVSampleFormat(t)
	return
}

/**
 * Generate a string corresponding to the sample format with
 * sample_fmt, or a header if sample_fmt is negative.
 *
 * @param buf the buffer where to write the string
 * @param buf_size the size of buf
 * @param sample_fmt the number of the sample format to print the
 * corresponding info string, or a negative value to print the
 * corresponding header.
 * @return the pointer to the filled buffer or NULL if sample_fmt is
 * unknown or in case of other errors
 */
//char *av_get_sample_fmt_string(char *buf, int buf_size, enum AVSampleFormat sample_fmt);
//未测试
func AvGetSampleFmtString(buf ffcommon.FBuf, buf_size ffcommon.FInt, sample_fmt ffconstant.AVSampleFormat) (res ffcommon.FConstCharP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_get_sample_fmt_string").Call(
		uintptr(unsafe.Pointer(buf)),
		uintptr(buf_size),
		uintptr(sample_fmt),
	)
	if err != nil {
		//return
	}
	res = ffcommon.GoAStr(t)
	return
}

/**
 * Return number of bytes per sample.
 *
 * @param sample_fmt the sample format
 * @return number of bytes per sample or zero if unknown for the given
 * sample format
 */
//int av_get_bytes_per_sample(enum AVSampleFormat sample_fmt);
//未测试
func AvGetBytesPerSample(sample_fmt ffconstant.AVSampleFormat) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_get_bytes_per_sample").Call(
		uintptr(sample_fmt),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Check if the sample format is planar.
 *
 * @param sample_fmt the sample format to inspect
 * @return 1 if the sample format is planar, 0 if it is interleaved
 */
//int av_sample_fmt_is_planar(enum AVSampleFormat sample_fmt);
//未测试
func AvSampleFmtIsPlanar(sample_fmt ffconstant.AVSampleFormat) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_sample_fmt_is_planar").Call(
		uintptr(sample_fmt),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Get the required buffer size for the given audio parameters.
 *
 * @param[out] linesize calculated linesize, may be NULL
 * @param nb_channels   the number of channels
 * @param nb_samples    the number of samples in a single channel
 * @param sample_fmt    the sample format
 * @param align         buffer size alignment (0 = default, 1 = no alignment)
 * @return              required buffer size, or negative error code on failure
 */
//int av_samples_get_buffer_size(int *linesize, int nb_channels, int nb_samples,
//enum AVSampleFormat sample_fmt, int align);
//未测试
func AvSamplesGetBufferSize(linesize *ffcommon.FInt, nb_channels ffcommon.FInt, nb_samples ffcommon.FInt,
	sample_fmt ffconstant.AVSampleFormat, align ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_samples_get_buffer_size").Call(
		uintptr(unsafe.Pointer(linesize)),
		uintptr(nb_channels),
		uintptr(nb_samples),
		uintptr(sample_fmt),
		uintptr(align),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * @}
 *
 * @defgroup lavu_sampmanip Samples manipulation
 *
 * Functions that manipulate audio samples
 * @{
 */

/**
 * Fill plane data pointers and linesize for samples with sample
 * format sample_fmt.
 *
 * The audio_data array is filled with the pointers to the samples data planes:
 * for planar, set the start point of each channel's data within the buffer,
 * for packed, set the start point of the entire buffer only.
 *
 * The value pointed to by linesize is set to the aligned size of each
 * channel's data buffer for planar layout, or to the aligned size of the
 * buffer for all channels for packed layout.
 *
 * The buffer in buf must be big enough to contain all the samples
 * (use av_samples_get_buffer_size() to compute its minimum size),
 * otherwise the audio_data pointers will point to invalid data.
 *
 * @see enum AVSampleFormat
 * The documentation for AVSampleFormat describes the data layout.
 *
 * @param[out] audio_data  array to be filled with the pointer for each channel
 * @param[out] linesize    calculated linesize, may be NULL
 * @param buf              the pointer to a buffer containing the samples
 * @param nb_channels      the number of channels
 * @param nb_samples       the number of samples in a single channel
 * @param sample_fmt       the sample format
 * @param align            buffer size alignment (0 = default, 1 = no alignment)
 * @return                 >=0 on success or a negative error code on failure
 * @todo return minimum size in bytes required for the buffer in case
 * of success at the next bump
 */
//int av_samples_fill_arrays(uint8_t **audio_data, int *linesize,
//const uint8_t *buf,
//int nb_channels, int nb_samples,
//enum AVSampleFormat sample_fmt, int align);
//未测试
func AvSamplesFillArrays(audio_data **ffcommon.FUint8T, linesize *ffcommon.FInt,
	buf ffcommon.FBuf,
	nb_channels ffcommon.FInt, nb_samples ffcommon.FInt,
	sample_fmt ffconstant.AVSampleFormat, align ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_samples_fill_arrays").Call(
		uintptr(unsafe.Pointer(&audio_data)),
		uintptr(unsafe.Pointer(linesize)),
		uintptr(unsafe.Pointer(buf)),
		uintptr(nb_channels),
		uintptr(nb_samples),
		uintptr(sample_fmt),
		uintptr(align),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Allocate a samples buffer for nb_samples samples, and fill data pointers and
 * linesize accordingly.
 * The allocated samples buffer can be freed by using av_freep(&audio_data[0])
 * Allocated data will be initialized to silence.
 *
 * @see enum AVSampleFormat
 * The documentation for AVSampleFormat describes the data layout.
 *
 * @param[out] audio_data  array to be filled with the pointer for each channel
 * @param[out] linesize    aligned size for audio buffer(s), may be NULL
 * @param nb_channels      number of audio channels
 * @param nb_samples       number of samples per channel
 * @param align            buffer size alignment (0 = default, 1 = no alignment)
 * @return                 >=0 on success or a negative error code on failure
 * @todo return the size of the allocated buffer in case of success at the next bump
 * @see av_samples_fill_arrays()
 * @see av_samples_alloc_array_and_samples()
 */
//int av_samples_alloc(uint8_t **audio_data, int *linesize, int nb_channels,
//int nb_samples, enum AVSampleFormat sample_fmt, int align);
//未测试
func AvSamplesAlloc(audio_data **ffcommon.FUint8T, linesize *ffcommon.FInt, buf ffcommon.FBuf, nb_channels ffcommon.FInt,
	nb_samples ffcommon.FInt, sample_fmt ffconstant.AVSampleFormat, align ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_samples_alloc").Call(
		uintptr(unsafe.Pointer(&audio_data)),
		uintptr(unsafe.Pointer(linesize)),
		uintptr(unsafe.Pointer(buf)),
		uintptr(nb_channels),
		uintptr(nb_samples),
		uintptr(sample_fmt),
		uintptr(align),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Allocate a data pointers array, samples buffer for nb_samples
 * samples, and fill data pointers and linesize accordingly.
 *
 * This is the same as av_samples_alloc(), but also allocates the data
 * pointers array.
 *
 * @see av_samples_alloc()
 */
//int av_samples_alloc_array_and_samples(uint8_t ***audio_data, int *linesize, int nb_channels,
//int nb_samples, enum AVSampleFormat sample_fmt, int align);
//未测试
func AvSamplesAllocArrayAndSamples(audio_data **ffcommon.FUint8T, linesize *ffcommon.FInt, buf ffcommon.FBuf, nb_channels ffcommon.FInt,
	nb_samples ffcommon.FInt, sample_fmt ffconstant.AVSampleFormat, align ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_samples_alloc_array_and_samples").Call(
		uintptr(unsafe.Pointer(&audio_data)),
		uintptr(unsafe.Pointer(linesize)),
		uintptr(unsafe.Pointer(buf)),
		uintptr(nb_channels),
		uintptr(nb_samples),
		uintptr(sample_fmt),
		uintptr(align),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Copy samples from src to dst.
 *
 * @param dst destination array of pointers to data planes
 * @param src source array of pointers to data planes
 * @param dst_offset offset in samples at which the data will be written to dst
 * @param src_offset offset in samples at which the data will be read from src
 * @param nb_samples number of samples to be copied
 * @param nb_channels number of audio channels
 * @param sample_fmt audio sample format
 */
//int av_samples_copy(uint8_t **dst, uint8_t * const *src, int dst_offset,
//int src_offset, int nb_samples, int nb_channels,
//enum AVSampleFormat sample_fmt);
//未测试
func AvSamplesCopy(dst **ffcommon.FUint8T, src **ffcommon.FUint8T, dst_offset ffcommon.FInt,
	src_offset ffcommon.FInt, nb_samples ffcommon.FInt, nb_channels ffcommon.FInt,
	sample_fmt ffconstant.AVSampleFormat) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_samples_copy").Call(
		uintptr(unsafe.Pointer(&dst)),
		uintptr(unsafe.Pointer(&src)),
		uintptr(dst_offset),
		uintptr(src_offset),
		uintptr(nb_samples),
		uintptr(nb_channels),
		uintptr(sample_fmt),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Fill an audio buffer with silence.
 *
 * @param audio_data  array of pointers to data planes
 * @param offset      offset in samples at which to start filling
 * @param nb_samples  number of samples to fill
 * @param nb_channels number of audio channels
 * @param sample_fmt  audio sample format
 */
//int av_samples_set_silence(uint8_t **audio_data, int offset, int nb_samples,
//int nb_channels, enum AVSampleFormat sample_fmt);
//未测试
func AvSamplesSetSilence(audio_data **ffcommon.FUint8T, offset ffcommon.FInt, nb_samples ffcommon.FInt,
	nb_channels ffcommon.FInt, sample_fmt ffconstant.AVSampleFormat) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_samples_set_silence").Call(
		uintptr(unsafe.Pointer(&audio_data)),
		uintptr(offset),
		uintptr(nb_samples),
		uintptr(nb_channels),
		uintptr(sample_fmt),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}
