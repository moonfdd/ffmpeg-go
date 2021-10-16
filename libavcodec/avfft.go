package libavcodec

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"unsafe"
)

/**
 * @file
 * @ingroup lavc_fft
 * FFT functions
 */

/**
 * @defgroup lavc_fft FFT functions
 * @ingroup lavc_misc
 *
 * @{
 */

type FFTComplex struct {
	re, im ffcommon.FFTSample
}

type FFTContext struct {
}

/**
 * Set up a complex FFT.
 * @param nbits           log2 of the length of the input array
 * @param inverse         if 0 perform the forward transform, if 1 perform the inverse
 */
//FFTContext *av_fft_init(int nbits, int inverse);
//未测试
func AvFftInit(nbits ffcommon.FInt, invers ffcommon.FInt) (res *FFTContext, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_fft_init").Call(
		uintptr(nbits),
		uintptr(invers),
	)
	if err != nil {
		//return
	}
	res = (*FFTContext)(unsafe.Pointer(t))
	return
}

/**
 * Do the permutation needed BEFORE calling ff_fft_calc().
 */
//void av_fft_permute(FFTContext *s, FFTComplex *z);
//未测试
func (s *FFTContext) AvFftPermute(z *FFTComplex) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_fft_permute").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(z)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Do a complex FFT with the parameters defined in av_fft_init(). The
 * input data must be permuted before. No 1.0/sqrt(n) normalization is done.
 */
//void av_fft_calc(FFTContext *s, FFTComplex *z);
//未测试
func (s *FFTContext) AvFftCalc(z *FFTComplex) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_fft_calc").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(z)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

//void av_fft_end(FFTContext *s);
//未测试
func (s *FFTContext) AvFftEnd() (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_fft_end").Call(
		uintptr(unsafe.Pointer(s)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

//FFTContext *av_mdct_init(int nbits, int inverse, double scale);
//未测试
func AvMdctInit() (res *FFTContext, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_mdct_init").Call()
	if err != nil {
		//return
	}
	res = (*FFTContext)(unsafe.Pointer(t))
	return
}

//void av_imdct_calc(FFTContext *s, FFTSample *output, const FFTSample *input);
//未测试
func (s *FFTContext) AvImdctCalc(output *FFTSample, input *FFTSample) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_imdct_calc").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(output)),
		uintptr(unsafe.Pointer(input)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

//void av_imdct_half(FFTContext *s, FFTSample *output, const FFTSample *input);
//未测试
func (s *FFTContext) AvImdctHalf(output *FFTSample, input *FFTSample) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_imdct_half").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(output)),
		uintptr(unsafe.Pointer(input)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

//void av_mdct_calc(FFTContext *s, FFTSample *output, const FFTSample *input);
//未测试
func (s *FFTContext) AvMdctCalc(output *FFTSample, input *FFTSample) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_mdct_calc").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(output)),
		uintptr(unsafe.Pointer(input)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

//void av_mdct_end(FFTContext *s);
//未测试
func (s *FFTContext) AvMdctEnd() (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_mdct_end").Call(
		uintptr(unsafe.Pointer(s)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

type RDFTContext struct {
}

/**
 * Set up a real FFT.
 * @param nbits           log2 of the length of the input array
 * @param trans           the type of transform
 */
//RDFTContext *av_rdft_init(int nbits, enum RDFTransformType trans);
//未测试
func AvRdftInit(nbits ffcommon.FInt, trans ffconstant.RDFTransformType) (res *RDFTContext, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_rdft_init").Call(
		uintptr(nbits),
		uintptr(trans),
	)
	if err != nil {
		//return
	}
	res = (*RDFTContext)(unsafe.Pointer(t))
	return
}

//void av_rdft_calc(RDFTContext *s, FFTSample *data);
//未测试
func (s *RDFTContext) AvMdctCalc(data *FFTSample) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_mdct_calc").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(data)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

//void av_rdft_end(RDFTContext *s);
//未测试
func (s *FFTContext) AvRdftEnd() (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_rdft_end").Call(
		uintptr(unsafe.Pointer(s)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/* Discrete Cosine Transform */

type DCTContext struct {
}

/**
 * Set up DCT.
 *
 * @param nbits           size of the input array:
 *                        (1 << nbits)     for DCT-II, DCT-III and DST-I
 *                        (1 << nbits) + 1 for DCT-I
 * @param type            the type of transform
 *
 * @note the first element of the input of DST-I is ignored
 */
//DCTContext *av_dct_init(int nbits, enum DCTTransformType type);
//未测试
func AvDctInit(nbits ffcommon.FInt, trans ffconstant.RDFTransformType) (res *RDFTContext, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_dct_init").Call(
		uintptr(nbits),
		uintptr(trans),
	)
	if err != nil {
		//return
	}
	res = (*RDFTContext)(unsafe.Pointer(t))
	return
}

//void av_dct_calc(DCTContext *s, FFTSample *data);
//未测试
func (s *DCTContext) AvDctCalc(data *FFTSample) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_dct_calc").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(data)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

//void av_dct_end (DCTContext *s);
//未测试
func (s *DCTContext) AvDctEnd() (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_dct_end").Call(
		uintptr(unsafe.Pointer(s)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}
