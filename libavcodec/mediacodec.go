package libavcodec

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

/**
 * This structure holds a reference to a android/view/Surface object that will
 * be used as output by the decoder.
 *
 */
type AVMediaCodecContext struct {

	/**
	 * android/view/Surface object reference.
	 */
	surface ffcommon.FVoidP
}

/**
 * Allocate and initialize a MediaCodec context.
 *
 * When decoding with MediaCodec is finished, the caller must free the
 * MediaCodec context with av_mediacodec_default_free.
 *
 * @return a pointer to a newly allocated AVMediaCodecContext on success, NULL otherwise
 */
//AVMediaCodecContext *av_mediacodec_alloc_context(void);
//未测试
func AvMediacodecAllocContext() (res *AVMediaCodecContext, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_mediacodec_alloc_context").Call()
	if err != nil {
		//return
	}
	res = (*AVMediaCodecContext)(unsafe.Pointer(t))
	return
}

/**
 * Convenience function that sets up the MediaCodec context.
 *
 * @param avctx codec context
 * @param ctx MediaCodec context to initialize
 * @param surface reference to an android/view/Surface
 * @return 0 on success, < 0 otherwise
 */
//int av_mediacodec_default_init(AVCodecContext *avctx, AVMediaCodecContext *ctx, void *surface);
//未测试
func (avctx *AVCodecContext) AvMediacodecDefaultInit(ctx *AVMediaCodecContext, surface ffcommon.FVoidP) (res *AVMediaCodecContext, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_mediacodec_default_init").Call(
		uintptr(unsafe.Pointer(avctx)),
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(surface)),
	)
	if err != nil {
		//return
	}
	res = (*AVMediaCodecContext)(unsafe.Pointer(t))
	return
}

/**
 * This function must be called to free the MediaCodec context initialized with
 * av_mediacodec_default_init().
 *
 * @param avctx codec context
 */
//void av_mediacodec_default_free(AVCodecContext *avctx);
//未测试
func (avctx *AVCodecContext) av_mediacodec_default_free() (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_mediacodec_default_free").Call(
		uintptr(unsafe.Pointer(avctx)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Opaque structure representing a MediaCodec buffer to render.
 */
type MediaCodecBuffer struct {
}

/**
 * Release a MediaCodec buffer and render it to the surface that is associated
 * with the decoder. This function should only be called once on a given
 * buffer, once released the underlying buffer returns to the codec, thus
 * subsequent calls to this function will have no effect.
 *
 * @param buffer the buffer to render
 * @param render 1 to release and render the buffer to the surface or 0 to
 * discard the buffer
 * @return 0 on success, < 0 otherwise
 */
//int av_mediacodec_release_buffer(AVMediaCodecBuffer *buffer, int render);
//未测试
func (buffer *AVMediaCodecBuffer) AvMediacodecReleaseBuffer(render ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_mediacodec_release_buffer").Call(
		uintptr(unsafe.Pointer(buffer)),
		uintptr(render),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Release a MediaCodec buffer and render it at the given time to the surface
 * that is associated with the decoder. The timestamp must be within one second
 * of the current java/lang/System#nanoTime() (which is implemented using
 * CLOCK_MONOTONIC on Android). See the Android MediaCodec documentation
 * of android/media/MediaCodec#releaseOutputBuffer(int,long) for more details.
 *
 * @param buffer the buffer to render
 * @param time timestamp in nanoseconds of when to render the buffer
 * @return 0 on success, < 0 otherwise
 */
//int av_mediacodec_render_buffer_at_time(AVMediaCodecBuffer *buffer, int64_t time);
//未测试
func (buffer *AVMediaCodecBuffer) AvMediacodecRenderBufferAtTime(time ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_mediacodec_render_buffer_at_time").Call(
		uintptr(unsafe.Pointer(buffer)),
		uintptr(time),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}
