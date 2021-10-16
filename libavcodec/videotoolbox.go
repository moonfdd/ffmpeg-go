package libavcodec

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

/**
 * This struct holds all the information that needs to be passed
 * between the caller and libavcodec for initializing Videotoolbox decoding.
 * Its size is not a part of the public ABI, it must be allocated with
 * av_videotoolbox_alloc_context() and freed with av_free().
 */
type AVVideotoolboxContext struct {
	///**
	// * Videotoolbox decompression session object.
	// * Created and freed the caller.
	// */
	//VTDecompressionSessionRef session;
	//
	///**
	// * The output callback that must be passed to the session.
	// * Set by av_videottoolbox_default_init()
	// */
	//VTDecompressionOutputCallback output_callback;
	//
	///**
	// * CVPixelBuffer Format Type that Videotoolbox will use for decoded frames.
	// * set by the caller. If this is set to 0, then no specific format is
	// * requested from the decoder, and its native format is output.
	// */
	//OSType cv_pix_fmt_type;
	//
	///**
	// * CoreMedia Format Description that Videotoolbox will use to create the decompression session.
	// * Set by the caller.
	// */
	//CMVideoFormatDescriptionRef cm_fmt_desc;
	//
	///**
	// * CoreMedia codec type that Videotoolbox will use to create the decompression session.
	// * Set by the caller.
	// */
	//int cm_codec_type;
}

/**
 * Allocate and initialize a Videotoolbox context.
 *
 * This function should be called from the get_format() callback when the caller
 * selects the AV_PIX_FMT_VIDETOOLBOX format. The caller must then create
 * the decoder object (using the output callback provided by libavcodec) that
 * will be used for Videotoolbox-accelerated decoding.
 *
 * When decoding with Videotoolbox is finished, the caller must destroy the decoder
 * object and free the Videotoolbox context using av_free().
 *
 * @return the newly allocated context or NULL on failure
 */
//AVVideotoolboxContext *av_videotoolbox_alloc_context(void);
//未测试
func AvDesAlloc() (res *AVVideotoolboxContext, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_des_alloc").Call()
	if err != nil {
		//return
	}
	//res = &AVAES{}
	res = (*AVVideotoolboxContext)(unsafe.Pointer(t))
	//res.instance = t
	//res.ptr = unsafe.Pointer(t)
	return
}

/**
 * This is a convenience function that creates and sets up the Videotoolbox context using
 * an internal implementation.
 *
 * @param avctx the corresponding codec context
 *
 * @return >= 0 on success, a negative AVERROR code on failure
 */
//int av_videotoolbox_default_init(AVCodecContext *avctx);
//未测试
func (avctx *AVCodecContext) AvVideotoolboxDefaultInit() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_videotoolbox_default_init").Call(
		uintptr(unsafe.Pointer(avctx)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * This is a convenience function that creates and sets up the Videotoolbox context using
 * an internal implementation.
 *
 * @param avctx the corresponding codec context
 * @param vtctx the Videotoolbox context to use
 *
 * @return >= 0 on success, a negative AVERROR code on failure
 */
//int av_videotoolbox_default_init2(AVCodecContext *avctx, AVVideotoolboxContext *vtctx);
//未测试
func (avctx *AVCodecContext) AvVideotoolboxDefaultInit2(vtctx *AVVideotoolboxContext) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_videotoolbox_default_init2").Call(
		uintptr(unsafe.Pointer(avctx)),
		uintptr(unsafe.Pointer(vtctx)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * This function must be called to free the Videotoolbox context initialized with
 * av_videotoolbox_default_init().
 *
 * @param avctx the corresponding codec context
 */
//void av_videotoolbox_default_free(AVCodecContext *avctx);
//未测试
func (avctx *AVCodecContext) av_videotoolbox_default_free() (err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_videotoolbox_default_free").Call(
		uintptr(unsafe.Pointer(avctx)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}
