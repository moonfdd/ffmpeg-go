package libavfilter

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/libavutil"
	"unsafe"
)

/**
 * Get the number of failed requests.
 *
 * A failed request is when the request_frame method is called while no
 * frame is present in the buffer.
 * The number is reset when a frame is added.
 */
//unsigned av_buffersrc_get_nb_failed_requests(AVFilterContext *buffer_src);
//未测试
func (buffer_src *AVFilterContext) AvBuffersrcGetNbFailedRequests() (res ffcommon.FUnsigned, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_buffersrc_get_nb_failed_requests").Call(
		uintptr(unsafe.Pointer(buffer_src)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FUnsigned(t)
	return
}

/**
 * This structure contains the parameters describing the frames that will be
 * passed to this filter.
 *
 * It should be allocated with av_buffersrc_parameters_alloc() and freed with
 * av_free(). All the allocated fields in it remain owned by the caller.
 */
type AVBufferSrcParameters struct {
	///**
	// * video: the pixel format, value corresponds to enum AVPixelFormat
	// * audio: the sample format, value corresponds to enum AVSampleFormat
	// */
	//int format;
	///**
	// * The timebase to be used for the timestamps on the input frames.
	// */
	//AVRational time_base;
	//
	///**
	// * Video only, the display dimensions of the input frames.
	// */
	//int width, height;
	//
	///**
	// * Video only, the sample (pixel) aspect ratio.
	// */
	//AVRational sample_aspect_ratio;
	//
	///**
	// * Video only, the frame rate of the input video. This field must only be
	// * set to a non-zero value if input stream has a known ffconstant framerate
	// * and should be left at its initial value if the framerate is variable or
	// * unknown.
	// */
	//AVRational frame_rate;
	//
	///**
	// * Video with a hwaccel pixel format only. This should be a reference to an
	// * AVHWFramesContext instance describing the input frames.
	// */
	//AVBufferRef *hw_frames_ctx;
	//
	///**
	// * Audio only, the audio sampling rate in samples per second.
	// */
	//int sample_rate;
	//
	///**
	// * Audio only, the audio channel layout
	// */
	//uint64_t channel_layout;
}

/**
 * Allocate a new AVBufferSrcParameters instance. It should be freed by the
 * caller with av_free().
 */
//AVBufferSrcParameters *av_buffersrc_parameters_alloc(void);
//未测试
func AvBuffersrcParametersAlloc() (res *AVBufferSrcParameters, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_buffersrc_parameters_alloc").Call()
	if err != nil {
		//return
	}
	res = (*AVBufferSrcParameters)(unsafe.Pointer(t))
	return
}

/**
 * Initialize the buffersrc or abuffersrc filter with the provided parameters.
 * This function may be called multiple times, the later calls override the
 * previous ones. Some of the parameters may also be set through AVOptions, then
 * whatever method is used last takes precedence.
 *
 * @param ctx an instance of the buffersrc or abuffersrc filter
 * @param param the stream parameters. The frames later passed to this filter
 *              must conform to those parameters. All the allocated fields in
 *              param remain owned by the caller, libavfilter will make internal
 *              copies or references when necessary.
 * @return 0 on success, a negative AVERROR code on failure.
 */
//int av_buffersrc_parameters_set(AVFilterContext *ctx, AVBufferSrcParameters *param);
//未测试
func (ctx *AVFilterContext) AvBuffersrcParametersSet(param *AVBufferSrcParameters) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_buffersrc_parameters_set").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(param)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Add a frame to the buffer source.
 *
 * @param ctx   an instance of the buffersrc filter
 * @param frame frame to be added. If the frame is reference counted, this
 * function will make a new reference to it. Otherwise the frame data will be
 * copied.
 *
 * @return 0 on success, a negative AVERROR on error
 *
 * This function is equivalent to av_buffersrc_add_frame_flags() with the
 * AV_BUFFERSRC_FLAG_KEEP_REF flag.
 */
//av_warn_unused_result
//int av_buffersrc_write_frame(AVFilterContext *ctx, const AVFrame *frame);
//未测试
func (ctx *AVFilterContext) AvBuffersrcWriteFrame(frame *libavutil.AVFrame) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_buffersrc_write_frame").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(frame)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Add a frame to the buffer source.
 *
 * @param ctx   an instance of the buffersrc filter
 * @param frame frame to be added. If the frame is reference counted, this
 * function will take ownership of the reference(s) and reset the frame.
 * Otherwise the frame data will be copied. If this function returns an error,
 * the input frame is not touched.
 *
 * @return 0 on success, a negative AVERROR on error.
 *
 * @note the difference between this function and av_buffersrc_write_frame() is
 * that av_buffersrc_write_frame() creates a new reference to the input frame,
 * while this function takes ownership of the reference passed to it.
 *
 * This function is equivalent to av_buffersrc_add_frame_flags() without the
 * AV_BUFFERSRC_FLAG_KEEP_REF flag.
 */
//av_warn_unused_result
//int av_buffersrc_add_frame(AVFilterContext *ctx, AVFrame *frame);
//未测试
func (ctx *AVFilterContext) AvBuffersrcAddFrame(frame *libavutil.AVFrame) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_buffersrc_add_frame").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(frame)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Add a frame to the buffer source.
 *
 * By default, if the frame is reference-counted, this function will take
 * ownership of the reference(s) and reset the frame. This can be controlled
 * using the flags.
 *
 * If this function returns an error, the input frame is not touched.
 *
 * @param buffer_src  pointer to a buffer source context
 * @param frame       a frame, or NULL to mark EOF
 * @param flags       a combination of AV_BUFFERSRC_FLAG_*
 * @return            >= 0 in case of success, a negative AVERROR code
 *                    in case of failure
 */
//av_warn_unused_result
//int av_buffersrc_add_frame_flags(AVFilterContext *buffer_src,
//AVFrame *frame, int flags);
//未测试
func (buffer_src *AVFilterContext) AvBuffersrcAddFrameFlags(frame *libavutil.AVFrame, flags ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_buffersrc_add_frame_flags").Call(
		uintptr(unsafe.Pointer(buffer_src)),
		uintptr(unsafe.Pointer(frame)),
		uintptr(flags),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Close the buffer source after EOF.
 *
 * This is similar to passing NULL to av_buffersrc_add_frame_flags()
 * except it takes the timestamp of the EOF, i.e. the timestamp of the end
 * of the last frame.
 */
//int av_buffersrc_close(AVFilterContext *ctx, int64_t pts, unsigned flags);
//未测试
func (ctx *AVFilterContext) AvBuffersrcClose(pts ffcommon.FUint64T, flags ffcommon.FUnsigned) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_buffersrc_close").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(pts),
		uintptr(flags),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}
