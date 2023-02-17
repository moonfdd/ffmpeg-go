package libavutil

import (
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
)

/*
 * A 32-bit implementation of the TEA algorithm
 * Copyright (c) 2015 Vesselin Bontchev
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

//#ifndef AVUTIL_TEA_H
//#define AVUTIL_TEA_H
//
//#include <stdint.h>

/**
 * @file
 * @brief Public header for libavutil TEA algorithm
 * @defgroup lavu_tea TEA
 * @ingroup lavu_crypto
 * @{
 */

//extern const int av_tea_size;

//struct AVTEA;
type AVTEA struct {
}

/**
 * Allocate an AVTEA context
 * To free the struct: av_free(ptr)
 */
//struct AVTEA *av_tea_alloc(void);
func AvTeaAlloc() (res *AVTEA) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_tea_alloc").Call()
	res = (*AVTEA)(unsafe.Pointer(t))
	return
}

/**
 * Initialize an AVTEA context.
 *
 * @param ctx an AVTEA context
 * @param key a key of 16 bytes used for encryption/decryption
 * @param rounds the number of rounds in TEA (64 is the "standard")
 */
//void av_tea_init(struct AVTEA *ctx, const uint8_t key[16], int rounds);
func (ctx *AVTEA) AvTeaInit(key [16]ffcommon.FUint8T, rounds ffcommon.FInt) {
	ffcommon.GetAvutilDll().NewProc("av_tea_init").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(&key)),
		uintptr(rounds),
	)
}

/**
 * Encrypt or decrypt a buffer using a previously initialized context.
 *
 * @param ctx an AVTEA context
 * @param dst destination array, can be equal to src
 * @param src source array, can be equal to dst
 * @param count number of 8 byte blocks
 * @param iv initialization vector for CBC mode, if NULL then ECB will be used
 * @param decrypt 0 for encryption, 1 for decryption
 */
//void av_tea_crypt(struct AVTEA *ctx, uint8_t *dst, const uint8_t *src,
//int count, uint8_t *iv, int decrypt);
func (ctx *AVTEA) AvTeaCrypt(dst, src *ffcommon.FUint8T,
	count ffcommon.FInt, iv *ffcommon.FUint8T, decrypt ffcommon.FInt) {
	ffcommon.GetAvutilDll().NewProc("av_tea_crypt").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
		uintptr(count),
		uintptr(unsafe.Pointer(iv)),
		uintptr(decrypt),
	)
}

/**
 * @}
 */

//#endif /* AVUTIL_TEA_H */
