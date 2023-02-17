package libavutil

import (
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
)

/*
 * Copyright (C) 2007 Michael Niedermayer <michaelni@gmx.at>
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
 * @ingroup lavu_sha
 * Public header for SHA-1 & SHA-256 hash function implementations.
 */

//#ifndef AVUTIL_SHA_H
//#define AVUTIL_SHA_H
//
//#include <stddef.h>
//#include <stdint.h>
//
//#include "attributes.h"
//#include "version.h"

/**
 * @defgroup lavu_sha SHA
 * @ingroup lavu_hash
 * SHA-1 and SHA-256 (Secure Hash Algorithm) hash function implementations.
 *
 * This module supports the following SHA hash functions:
 *
 * - SHA-1: 160 bits
 * - SHA-224: 224 bits, as a variant of SHA-2
 * - SHA-256: 256 bits, as a variant of SHA-2
 *
 * @see For SHA-384, SHA-512, and variants thereof, see @ref lavu_sha512.
 *
 * @{
 */

//extern const int av_sha_size;

//struct AVSHA;
type AVSHA struct {
}

/**
 * Allocate an AVSHA context.
 */
//struct AVSHA *av_sha_alloc(void);
func AvShaSlloc() (res *AVRC4) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_sha_alloc").Call()
	res = (*AVRC4)(unsafe.Pointer(t))
	return
}

/**
 * Initialize SHA-1 or SHA-2 hashing.
 *
 * @param context pointer to the function context (of size av_sha_size)
 * @param bits    number of bits in digest (SHA-1 - 160 bits, SHA-2 224 or 256 bits)
 * @return        zero if initialization succeeded, -1 otherwise
 */
//int av_sha_init(struct AVSHA* context, int bits);
func (context *AVSHA) AvShaInit(key *ffcommon.FUint8T, bits ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_sha_init").Call(
		uintptr(unsafe.Pointer(context)),
		uintptr(unsafe.Pointer(key)),
		uintptr(bits),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Update hash value.
 *
 * @param ctx     hash function context
 * @param data    input data to update hash with
 * @param len     input data length
 */
//#if FF_API_CRYPTO_SIZE_T
//void av_sha_update(struct AVSHA *ctx, const uint8_t *data, unsigned int len);
//#else
//void av_sha_update(struct AVSHA *ctx, const uint8_t *data, size_t len);
//#endif
func (context *AVSHA) AvShaUpdate(data *ffcommon.FUint8T, len0 ffcommon.FUnsignedIntOrSizeT) {
	ffcommon.GetAvutilDll().NewProc("av_sha_update").Call(
		uintptr(unsafe.Pointer(context)),
		uintptr(unsafe.Pointer(data)),
		uintptr(len0),
	)
}

/**
 * Finish hashing and output digest value.
 *
 * @param context hash function context
 * @param digest  buffer where output digest value is stored
 */
//void av_sha_final(struct AVSHA* context, uint8_t *digest);
func (context *AVRIPEMD) AvShaFinal(digest *ffcommon.FUint8T) {
	ffcommon.GetAvutilDll().NewProc("av_sha_final").Call(
		uintptr(unsafe.Pointer(context)),
		uintptr(unsafe.Pointer(digest)),
	)
}

/**
 * @}
 */

//#endif /* AVUTIL_SHA_H */
