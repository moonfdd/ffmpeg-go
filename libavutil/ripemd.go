package libavutil

import (
	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"unsafe"
)

/*
 * Copyright (C) 2007 Michael Niedermayer <michaelni@gmx.at>
 * Copyright (C) 2013 James Almer <jamrial@gmail.com>
 *
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

/**
 * @file
 * @ingroup lavu_ripemd
 * Public header for RIPEMD hash function implementation.
 */

//#ifndef AVUTIL_RIPEMD_H
//#define AVUTIL_RIPEMD_H
//
//#include <stddef.h>
//#include <stdint.h>
//
//#include "attributes.h"
//#include "version.h"

/**
 * @defgroup lavu_ripemd RIPEMD
 * @ingroup lavu_hash
 * RIPEMD hash function implementation.
 *
 * @{
 */

//extern const int av_ripemd_size;

//struct AVRIPEMD;
type AVRIPEMD struct {
}

/**
 * Allocate an AVRIPEMD context.
 */
//struct AVRIPEMD *av_ripemd_alloc(void);
func AvRipemdAlloc() (res *AVRIPEMD) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_ripemd_alloc").Call()
	if t == 0 {

	}
	res = (*AVRIPEMD)(unsafe.Pointer(t))
	return
}

/**
 * Initialize RIPEMD hashing.
 *
 * @param context pointer to the function context (of size av_ripemd_size)
 * @param bits    number of bits in digest (128, 160, 256 or 320 bits)
 * @return        zero if initialization succeeded, -1 otherwise
 */
//int av_ripemd_init(struct AVRIPEMD* context, int bits);
func (context *AVRIPEMD) AvRipemdInit(bits ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_ripemd_init").Call(
		uintptr(unsafe.Pointer(context)),
		uintptr(bits),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Update hash value.
 *
 * @param context hash function context
 * @param data    input data to update hash with
 * @param len     input data length
 */
//#if FF_API_CRYPTO_SIZE_T
//void av_ripemd_update(struct AVRIPEMD* context, const uint8_t* data, unsigned int len);
//#else
//void av_ripemd_update(struct AVRIPEMD* context, const uint8_t* data, size_t len);
//#endif
func (context *AVRIPEMD) AvRipemdUpdate(data *ffcommon.FUint8T, len0 ffcommon.FUnsignedIntOrSizeT) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_ripemd_update").Call(
		uintptr(unsafe.Pointer(context)),
		uintptr(unsafe.Pointer(data)),
		uintptr(len0),
	)
	if t == 0 {

	}
	return
}

/**
 * Finish hashing and output digest value.
 *
 * @param context hash function context
 * @param digest  buffer where output digest value is stored
 */
//void av_ripemd_final(struct AVRIPEMD* context, uint8_t *digest);
func (context *AVRIPEMD) AvRipemdFinal(digest *ffcommon.FUint8T) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_ripemd_final").Call(
		uintptr(unsafe.Pointer(context)),
		uintptr(unsafe.Pointer(digest)),
	)
	if t == 0 {

	}
	return
}

/**
 * @}
 */

//#endif /* AVUTIL_RIPEMD_H */
