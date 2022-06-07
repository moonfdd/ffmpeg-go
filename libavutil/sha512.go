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
 * @ingroup lavu_sha512
 * Public header for SHA-512 implementation.
 */

//#ifndef AVUTIL_SHA512_H
//#define AVUTIL_SHA512_H
//
//#include <stddef.h>
//#include <stdint.h>
//
//#include "attributes.h"
//#include "version.h"

/**
 * @defgroup lavu_sha512 SHA-512
 * @ingroup lavu_hash
 * SHA-512 (Secure Hash Algorithm) hash function implementations.
 *
 * This module supports the following SHA-2 hash functions:
 *
 * - SHA-512/224: 224 bits
 * - SHA-512/256: 256 bits
 * - SHA-384: 384 bits
 * - SHA-512: 512 bits
 *
 * @see For SHA-1, SHA-256, and variants thereof, see @ref lavu_sha.
 *
 * @{
 */

//extern const int av_sha512_size;

//struct AVSHA512;
type AVSHA512 struct {
}

/**
 * Allocate an AVSHA512 context.
 */
//struct AVSHA512 *av_sha512_alloc(void);
func AvSha512Alloc() (res *AVSHA512) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_sha512_alloc").Call()
	if t == 0 {

	}
	res = (*AVSHA512)(unsafe.Pointer(t))
	return
}

/**
 * Initialize SHA-2 512 hashing.
 *
 * @param context pointer to the function context (of size av_sha512_size)
 * @param bits    number of bits in digest (224, 256, 384 or 512 bits)
 * @return        zero if initialization succeeded, -1 otherwise
 */
//int av_sha512_init(struct AVSHA512* context, int bits);
func (context *AVSHA512) AvSha512Init(key *ffcommon.FUint8T, bits ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_sha512_init").Call(
		uintptr(unsafe.Pointer(context)),
		uintptr(unsafe.Pointer(key)),
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
//void av_sha512_update(struct AVSHA512* context, const uint8_t* data, unsigned int len);
//#else
//void av_sha512_update(struct AVSHA512* context, const uint8_t* data, size_t len);
//#endif
func (context *AVSHA) AvSha512Update(data *ffcommon.FUint8T, len0 ffcommon.FUnsignedIntOrSizeT) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_sha512_update").Call(
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
//void av_sha512_final(struct AVSHA512* context, uint8_t *digest);
func (context *AVSHA512) AvSha512Final(digest *ffcommon.FUint8T) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_sha512_final").Call(
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

//#endif /* AVUTIL_SHA512_H */
