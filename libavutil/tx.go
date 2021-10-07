package libavutil

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"unsafe"
)

type AVTXContext struct {
}

type AVComplexFloat struct {
	Re, Im ffcommon.FFloat
}

type AVComplexDouble struct {
	Re, Im ffcommon.FDouble
}

type AVComplexInt32 struct {
	Re, Im ffcommon.FInt32T
}

///**
// * Function pointer to a function to perform the transform.
// *
// * @note Using a different context than the one allocated during av_tx_init()
// * is not allowed.
// *
// * @param s the transform context
// * @param out the output array
// * @param in the input array
// * @param stride the input or output stride in bytes
// *
// * The out and in arrays must be aligned to the maximum required by the CPU
// * architecture.
// * The stride must follow the constraints the transform type has specified.
// */
//typedef void (*av_tx_fn)(AVTXContext *s, void *out, void *in, ptrdiff_t stride);
//func AvTxFn(s *AVTXContext, out ffcommon.FVoidP, in ffcommon.FVoidP, stride ffcommon.FPtrdiffT) {}
type AvTxFn = func(s *AVTXContext, out ffcommon.FVoidP, in ffcommon.FVoidP, stride ffcommon.FPtrdiffT)

/**
* Initialize a transform context with the given configuration
* (i)MDCTs with an odd length are currently not supported.
*
* @param ctx the context to allocate, will be NULL on error
* @param tx pointer to the transform function pointer to set
* @param type type the type of transform
* @param inv whether to do an inverse or a forward transform
* @param len the size of the transform in samples
* @param scale pointer to the value to scale the output if supported by type
* @param flags a bitmask of AVTXFlags or 0
*
* @return 0 on success, negative error code on failure
 */
//int av_tx_init(AVTXContext **ctx, av_tx_fn *tx, enum AVTXType type,
//int inv, int len, const void *scale, uint64_t flags);
//未测试
func AvTxInit(ctx **AVTXContext, tx AvTxFn, type0 ffconstant.AVTXType,
	inv ffcommon.FInt, len0 ffcommon.FInt, scale ffcommon.FVoidP, flags ffcommon.FUint64T) (res *ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_tx_init").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(&tx)),
		uintptr(type0),
		uintptr(inv),
		uintptr(len0),
		uintptr(scale),
		uintptr(flags),
	)
	if err != nil {
		//return
	}
	res = (*ffcommon.FInt)(unsafe.Pointer(t))
	return
}

/**
* Frees a context and sets ctx to NULL, does nothing when ctx == NULL
 */
//void av_tx_uninit(AVTXContext **ctx);
//未测试
func AvTxUninit(ctx **AVTXContext) (err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_tx_uninit").Call(
		uintptr(unsafe.Pointer(ctx)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}
