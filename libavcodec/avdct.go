package libavcodec

import (
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/ffmpeg-go/libavutil"
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

//#ifndef AVCODEC_AVDCT_H
//#define AVCODEC_AVDCT_H
//
//#include "../libavutil/opt.h"

/**
 * AVDCT context.
 * @note function pointers can be NULL if the specific features have been
 *       disabled at build time.
 */
type AVDCT struct {
	AvClass *libavutil.AVClass

	//void (*idct)(int16_t *block /* align 16 */);
	Idct uintptr
	/**
	 * IDCT input permutation.
	 * Several optimized IDCTs need a permutated input (relative to the
	 * normal order of the reference IDCT).
	 * This permutation must be performed before the idct_put/add.
	 * Note, normally this can be merged with the zigzag/alternate scan<br>
	 * An example to avoid confusion:
	 * - (->decode coeffs -> zigzag reorder -> dequant -> reference IDCT -> ...)
	 * - (x -> reference DCT -> reference IDCT -> x)
	 * - (x -> reference DCT -> simple_mmx_perm = idct_permutation
	 *    -> simple_idct_mmx -> x)
	 * - (-> decode coeffs -> zigzag reorder -> simple_mmx_perm -> dequant
	 *    -> simple_idct_mmx -> ...)
	 */
	IdctPermutation [64]ffcommon.FUint8T

	//void (*fdct)(int16_t *block /* align 16 */);
	Fdct uintptr

	/**
	 * DCT algorithm.
	 * must use AVOptions to set this field.
	 */
	DctAlgo ffcommon.FInt

	/**
	 * IDCT algorithm.
	 * must use AVOptions to set this field.
	 */
	IdctAlgo ffcommon.FInt

	//void (*get_pixels)(int16_t *block /* align 16 */,
	//const uint8_t *pixels /* align 8 */,
	//ptrdiff_t line_size);
	GetPixels     uintptr
	BitsPerSample ffcommon.FInt

	//void (*get_pixels_unaligned)(int16_t *block /* align 16 */,
	//const uint8_t *pixels,
	//ptrdiff_t line_size);
	GetPixelsUnaligned uintptr
}

/**
 * Allocates a AVDCT context.
 * This needs to be initialized with avcodec_dct_init() after optionally
 * configuring it with AVOptions.
 *
 * To free it use av_free()
 */
//AVDCT *avcodec_dct_alloc(void);
func AvcodecDctAlloc() (res *AVDCT) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_dct_alloc").Call()
	res = (*AVDCT)(unsafe.Pointer(t))
	return
}

//int avcodec_dct_init(AVDCT *);
func (a *AVDCT) AvcodecDctInit() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_dct_init").Call(
		uintptr(unsafe.Pointer(a)),
	)
	res = ffcommon.FInt(t)
	return
}

//const AVClass *avcodec_dct_get_class(void);
func AvcodecDctGetClass() (res *AVClass) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_dct_get_class").Call()
	res = (*AVClass)(unsafe.Pointer(t))
	return
}

//#endif /* AVCODEC_AVDCT_H */
