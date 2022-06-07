package libavcodec

import (
	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"unsafe"
)

/*
 * This file is part of FFmpeg.
 *
 * FFmpeg is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * FFmpeg is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with FFmpeg; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA
 */

//#ifndef AVCODEC_AVFFT_H
//#define AVCODEC_AVFFT_H

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

//type  FFTSample=ffcommon.FFloat

type FFTComplex struct {
	Re, Im ffcommon.FFloat
}

//typedef struct FFTContext FFTContext;
type FFTContext struct {
}

/**
 * Set up a complex FFT.
 * @param nbits           log2 of the length of the input array
 * @param inverse         if 0 perform the forward transform, if 1 perform the inverse
 */
//FFTContext *av_fft_init(int nbits, int inverse);
func AvFftInit(nbits, inverse ffcommon.FInt) (res *FFTContext) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_fft_init").Call(
		uintptr(nbits),
		uintptr(inverse),
	)
	if t == 0 {

	}
	res = (*FFTContext)(unsafe.Pointer(t))
	return
}

/**
 * Do the permutation needed BEFORE calling ff_fft_calc().
 */
//void av_fft_permute(FFTContext *s, FFTComplex *z);
func (s *FFTContext) AvFftPermute(z *FFTComplex) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_fft_permute").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(z)),
	)
	if t == 0 {

	}
	return
}

/**
 * Do a complex FFT with the parameters defined in av_fft_init(). The
 * input data must be permuted before. No 1.0/sqrt(n) normalization is done.
 */
//void av_fft_calc(FFTContext *s, FFTComplex *z);
func (s *FFTContext) AvFftCalc(z *FFTComplex) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_fft_calc").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(z)),
	)
	if t == 0 {

	}
	return
}

//void av_fft_end(FFTContext *s);
func (s *FFTContext) AvFftEnd() {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_fft_end").Call(
		uintptr(unsafe.Pointer(s)),
	)
	if t == 0 {

	}
	return
}

//FFTContext *av_mdct_init(int nbits, int inverse, double scale);
func AvMdctInit(nbits, inverse ffcommon.FInt, scale ffcommon.FDouble) (res *FFTContext) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_mdct_init").Call(
		uintptr(nbits),
		uintptr(inverse),
		uintptr(unsafe.Pointer(&scale)),
	)
	if t == 0 {

	}
	res = (*FFTContext)(unsafe.Pointer(t))
	return
}

//void av_imdct_calc(FFTContext *s, FFTSample *output, const FFTSample *input);
func (s *FFTContext) AvImdctCalc(output, input *ffcommon.FFTSample) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_imdct_calc").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(output)),
		uintptr(unsafe.Pointer(input)),
	)
	if t == 0 {

	}
	return
}

//void av_imdct_half(FFTContext *s, FFTSample *output, const FFTSample *input);
func (s *FFTContext) AvImdctHalf(output, input *ffcommon.FFTSample) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_imdct_half").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(output)),
		uintptr(unsafe.Pointer(input)),
	)
	if t == 0 {

	}
	return
}

//void av_mdct_calc(FFTContext *s, FFTSample *output, const FFTSample *input);
func (s *FFTContext) AvMdctCalc(output, input *ffcommon.FFTSample) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_mdct_calc").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(output)),
		uintptr(unsafe.Pointer(input)),
	)
	if t == 0 {

	}
	return
}

//void av_mdct_end(FFTContext *s);
func (s *FFTContext) AvMdctEnd(output, input *ffcommon.FFTSample) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_mdct_end").Call(
		uintptr(unsafe.Pointer(s)),
	)
	if t == 0 {

	}
	return
}

/* Real Discrete Fourier Transform */
type RDFTransformType = int32

const (
	DFT_R2C = iota
	IDFT_C2R
	IDFT_R2C
	DFT_C2R
)

//typedef struct RDFTContext RDFTContext;
type RDFTContext struct {
}

/**
 * Set up a real FFT.
 * @param nbits           log2 of the length of the input array
 * @param trans           the type of transform
 */
//RDFTContext *av_rdft_init(int nbits, enum RDFTransformType trans);
func AvRdftInit(nbits ffcommon.FInt, trans RDFTransformType) (res *RDFTContext) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_rdft_init").Call(
		uintptr(nbits),
		uintptr(trans),
	)
	if t == 0 {

	}
	res = (*RDFTContext)(unsafe.Pointer(t))
	return
}

//void av_rdft_calc(RDFTContext *s, FFTSample *data);
func (s *RDFTContext) AvRdftCalc(data *ffcommon.FFTSample) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_rdft_calc").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(data)),
	)
	if t == 0 {

	}
	return
}

//void av_rdft_end(RDFTContext *s);
func (s *RDFTContext) AvRdftEnd() {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_rdft_end").Call(
		uintptr(unsafe.Pointer(s)),
	)
	if t == 0 {

	}
	return
}

/* Discrete Cosine Transform */

//typedef struct DCTContext DCTContext;
type DCTContext struct {
}
type DCTTransformType = int32

const (
	DCT_II = iota
	DCT_III
	DCT_I
	DST_I
)

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
func AvDctInit(nbits ffcommon.FInt, type0 DCTTransformType) (res *DCTContext) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_dct_init").Call(
		uintptr(nbits),
		uintptr(type0),
	)
	if t == 0 {

	}
	res = (*DCTContext)(unsafe.Pointer(t))
	return
}

//void av_dct_calc(DCTContext *s, FFTSample *data);
func (s *DCTContext) AvDctCalc(data *ffcommon.FFTSample) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_dct_calc").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(data)),
	)
	if t == 0 {

	}
	return
}

//void av_dct_end (DCTContext *s);
func (s *DCTContext) AvDctEnd() {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_dct_end").Call(
		uintptr(unsafe.Pointer(s)),
	)
	if t == 0 {

	}
	return
}

/**
 * @}
 */

//#endif /* AVCODEC_AVFFT_H */
