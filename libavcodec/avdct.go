package libavcodec

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/libavutil"
	"unsafe"
)

/**
 * AVDCT context.
 * @note function pointers can be NULL if the specific features have been
 *       disabled at build time.
 */
type AVDCT struct {
	//const AVClass *av_class;
	//
	//void (*idct)(int16_t *block /* align 16 */);
	//
	///**
	// * IDCT input permutation.
	// * Several optimized IDCTs need a permutated input (relative to the
	// * normal order of the reference IDCT).
	// * This permutation must be performed before the idct_put/add.
	// * Note, normally this can be merged with the zigzag/alternate scan<br>
	// * An example to avoid confusion:
	// * - (->decode coeffs -> zigzag reorder -> dequant -> reference IDCT -> ...)
	// * - (x -> reference DCT -> reference IDCT -> x)
	// * - (x -> reference DCT -> simple_mmx_perm = idct_permutation
	// *    -> simple_idct_mmx -> x)
	// * - (-> decode coeffs -> zigzag reorder -> simple_mmx_perm -> dequant
	// *    -> simple_idct_mmx -> ...)
	// */
	//uint8_t idct_permutation[64];
	//
	//void (*fdct)(int16_t *block /* align 16 */);
	//
	//
	///**
	// * DCT algorithm.
	// * must use AVOptions to set this field.
	// */
	//int dct_algo;
	//
	///**
	// * IDCT algorithm.
	// * must use AVOptions to set this field.
	// */
	//int idct_algo;
	//
	//void (*get_pixels)(int16_t *block /* align 16 */,
	//const uint8_t *pixels /* align 8 */,
	//ptrdiff_t line_size);
	//
	//int bits_per_sample;
	//
	//void (*get_pixels_unaligned)(int16_t *block /* align 16 */,
	//const uint8_t *pixels,
	//ptrdiff_t line_size);
}

/**
 * Allocates a AVDCT context.
 * This needs to be initialized with avcodec_dct_init() after optionally
 * configuring it with AVOptions.
 *
 * To free it use av_free()
 */
//AVDCT *avcodec_dct_alloc(void);
//未测试
func AvcodecDctAlloc() (res *AVDCT, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("avcodec_dct_alloc").Call()
	if err != nil {
		//return
	}
	res = (*AVDCT)(unsafe.Pointer(t))
	return
}

//int avcodec_dct_init(AVDCT *);
//未测试
func (a *AVDCT) AvcodecDctInit() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("avcodec_dct_init").Call(
		uintptr(unsafe.Pointer(a)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//const AVClass *avcodec_dct_get_class(void);
//未测试
func AvcodecDctGetClass() (res *libavutil.AVClass, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("avcodec_dct_get_class").Call()
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = (*libavutil.AVClass)(unsafe.Pointer(t))
	return
}
