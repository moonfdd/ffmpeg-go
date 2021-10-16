package libavcodec

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

type AVCodecContext struct {
}
type AVFrame struct {
}

//typedef int (*AVVDPAU_Render2)(struct AVCodecContext *, struct AVFrame *,
//const VdpPictureInfo *, uint32_t,
//const VdpBitstreamBuffer *);

/**
* This structure is used to share data between the libavcodec library and
* the client video application.
* The user shall allocate the structure via the av_alloc_vdpau_hwaccel
* function and make it available as
* AVCodecContext.hwaccel_context. Members can be set by the user once
* during initialization or through each AVCodecContext.get_buffer()
* function call. In any case, they must be valid prior to calling
* decoding functions.
*
* The size of this structure is not a part of the public ABI and must not
* be used outside of libavcodec. Use av_vdpau_alloc_context() to allocate an
* AVVDPAUContext.
 */
type AVVDPAUContext struct {
	//
	///**
	//* VDPAU decoder handle
	//*
	//* Set by user.
	//*/
	// decoder ffconstant.VdpDecoder
	//
	///**
	//* VDPAU decoder render callback
	//*
	//* Set by the user.
	//*/
	//VdpDecoderRender *render;
	//
	//AVVDPAU_Render2 render2;
}

/**
 * @brief allocation function for AVVDPAUContext
 *
 * Allows extending the struct without breaking API/ABI
 */
//AVVDPAUContext *av_alloc_vdpaucontext(void);
//未测试
func AvAllocVdpaucontext() (res *AVVDPAUContext, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_alloc_vdpaucontext").Call()
	if err != nil {
		//return
	}
	//res = &AVAES{}
	res = (*AVVDPAUContext)(unsafe.Pointer(t))
	//res.instance = t
	//res.ptr = unsafe.Pointer(t)
	return
}

//AVVDPAU_Render2 av_vdpau_hwaccel_get_render2(const AVVDPAUContext *);
//未测试
func (c *AVVDPAUContext) av_vdpau_hwaccel_get_render2() (res *AVVDPAU_Render2, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_vdpau_hwaccel_get_render2").Call(
		uintptr(unsafe.Pointer(c)),
	)
	if err != nil {
		//return
	}
	//res = &AVAES{}
	res = (*AVVDPAU_Render2)(unsafe.Pointer(t))
	//res.instance = t
	//res.ptr = unsafe.Pointer(t)
	return
}

//void av_vdpau_hwaccel_set_render2(AVVDPAUContext *, AVVDPAU_Render2);
//未测试
func (c *AVVDPAUContext) AvVdpauHwaccelSetRender2(r *AVVDPAU_Render2) (err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_vdpau_hwaccel_set_render2").Call(
		uintptr(unsafe.Pointer(c)),
		uintptr(unsafe.Pointer(r)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
* Associate a VDPAU device with a codec context for hardware acceleration.
* This function is meant to be called from the get_format() codec callback,
* or earlier. It can also be called after avcodec_flush_buffers() to change
* the underlying VDPAU device mid-stream (e.g. to recover from non-transparent
* display preemption).
*
* @note get_format() must return AV_PIX_FMT_VDPAU if this function completes
* successfully.
*
* @param avctx decoding context whose get_format() callback is invoked
* @param device VDPAU device handle to use for hardware acceleration
* @param get_proc_address VDPAU device driver
* @param flags zero of more OR'd AV_HWACCEL_FLAG_* flags
*
* @return 0 on success, an AVERROR code on failure.
 */
//int av_vdpau_bind_context(AVCodecContext *avctx, VdpDevice device,
//VdpGetProcAddress *get_proc_address, unsigned flags);
//未测试
func (avctx *AVCodecContext) AvVdpauBindContext(device VdpDevice,
	get_proc_address *VdpGetProcAddress, flags ffcommon.FUnsigned) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_vdpau_bind_context").Call(
		uintptr(unsafe.Pointer(avctx)),
		uintptr(unsafe.Pointer(&device)),
		uintptr(unsafe.Pointer(get_proc_address)),
		uintptr(flags),
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
* Gets the parameters to create an adequate VDPAU video surface for the codec
* context using VDPAU hardware decoding acceleration.
*
* @note Behavior is undefined if the context was not successfully bound to a
* VDPAU device using av_vdpau_bind_context().
*
* @param avctx the codec context being used for decoding the stream
* @param type storage space for the VDPAU video surface chroma type
*              (or NULL to ignore)
* @param width storage space for the VDPAU video surface pixel width
*              (or NULL to ignore)
* @param height storage space for the VDPAU video surface pixel height
*              (or NULL to ignore)
*
* @return 0 on success, a negative AVERROR code on failure.
 */
//int av_vdpau_get_surface_parameters(AVCodecContext *avctx, VdpChromaType *type,
//uint32_t *width, uint32_t *height);
//未测试
func (avctx *AVCodecContext) AvVdpauGetSurfaceParameters(type0 *VdpChromaType,
	width *ffcommon.FUint32T, height *ffcommon.FUint32T) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_vdpau_get_surface_parameters").Call(
		uintptr(unsafe.Pointer(avctx)),
		uintptr(unsafe.Pointer(type0)),
		uintptr(unsafe.Pointer(width)),
		uintptr(unsafe.Pointer(height)),
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
* Allocate an AVVDPAUContext.
*
* @return Newly-allocated AVVDPAUContext or NULL on failure.
 */
//AVVDPAUContext *av_vdpau_alloc_context(void);
//未测试
func AvVdpauAllocContext() (res *AVVDPAUContext, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_vdpau_alloc_context").Call()
	if err != nil {
		//return
	}
	//res = &AVAES{}
	res = (*AVVDPAUContext)(unsafe.Pointer(t))
	//res.instance = t
	//res.ptr = unsafe.Pointer(t)
	return
}

//#if FF_API_VDPAU_PROFILE
/**
* Get a decoder profile that should be used for initializing a VDPAU decoder.
* Should be called from the AVCodecContext.get_format() callback.
*
* @deprecated Use av_vdpau_bind_context() instead.
*
* @param avctx the codec context being used for decoding the stream
* @param profile a pointer into which the result will be written on success.
*                The contents of profile are undefined if this function returns
*                an error.
*
* @return 0 on success (non-negative), a negative AVERROR on failure.
 */
//attribute_deprecated
//int av_vdpau_get_profile(AVCodecContext *avctx, VdpDecoderProfile *profile);
//未测试
func (avctx *AVCodecContext) AvVdpauGetProfile(profile *VdpDecoderProfile) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_vdpau_get_profile").Call(
		uintptr(unsafe.Pointer(avctx)),
		uintptr(unsafe.Pointer(profile)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//#endif
