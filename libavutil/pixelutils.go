package libavutil

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

/**
 * Sum of abs(src1[x] - src2[x])
 */
//typedef int (*av_pixelutils_sad_fn)(const uint8_t *src1, ptrdiff_t stride1,
//const uint8_t *src2, ptrdiff_t stride2);
type AvPixelutilsSadFn = func(src1 *ffcommon.FUint8T, stride1 ffcommon.FPtrdiffT,
	src2 *ffcommon.FUint8T, stride2 ffcommon.FPtrdiffT) ffcommon.FInt

/**
 * Get a potentially optimized pointer to a Sum-of-absolute-differences
 * function (see the av_pixelutils_sad_fn prototype).
 *
 * @param w_bits  1<<w_bits is the requested width of the block size
 * @param h_bits  1<<h_bits is the requested height of the block size
 * @param aligned If set to 2, the returned sad function will assume src1 and
 *                src2 addresses are aligned on the block size.
 *                If set to 1, the returned sad function will assume src1 is
 *                aligned on the block size.
 *                If set to 0, the returned sad function assume no particular
 *                alignment.
 * @param log_ctx context used for logging, can be NULL
 *
 * @return a pointer to the SAD function or NULL in case of error (because of
 *         invalid parameters)
 */
//av_pixelutils_sad_fn av_pixelutils_get_sad_fn(int w_bits, int h_bits,
//int aligned, void *log_ctx);
//未测试
func AvPixelutilsGetSadFn(w_bits ffcommon.FInt, h_bits ffcommon.FInt,
	aligned ffcommon.FInt, log_ctx ffcommon.FVoidP) (res AvPixelutilsSadFn, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_pixelutils_get_sad_fn").Call(
		uintptr(w_bits),
		uintptr(h_bits),
		uintptr(aligned),
		uintptr(log_ctx),
	)
	if err != nil {
		//return
	}
	res = *(*AvPixelutilsSadFn)(unsafe.Pointer(t))
	return
}
