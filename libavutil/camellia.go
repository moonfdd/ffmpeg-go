package libavutil

import (
	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"unsafe"
)

/*
 * An implementation of the CAMELLIA algorithm as mentioned in RFC3713
 * Copyright (c) 2014 Supraja Meedinti
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

//#ifndef AVUTIL_CAMELLIA_H
//#define AVUTIL_CAMELLIA_H
//
//#include <stdint.h>

/**
 * @file
 * @brief Public header for libavutil CAMELLIA algorithm
 * @defgroup lavu_camellia CAMELLIA
 * @ingroup lavu_crypto
 * @{
 */

//extern const int av_camellia_size;

//struct AVCAMELLIA;
type AVCAMELLIA struct {
}

/**
 * Allocate an AVCAMELLIA context
 * To free the struct: av_free(ptr)
 */
//struct AVCAMELLIA *av_camellia_alloc(void);
func AvCamelliaAlloc() (res *AVCAMELLIA) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_camellia_alloc").Call()
	if t == 0 {

	}
	res = (*AVCAMELLIA)(unsafe.Pointer(t))
	return
}

/**
 * Initialize an AVCAMELLIA context.
 *
 * @param ctx an AVCAMELLIA context
 * @param key a key of 16, 24, 32 bytes used for encryption/decryption
 * @param key_bits number of keybits: possible are 128, 192, 256
 */
//int av_camellia_init(struct AVCAMELLIA *ctx, const uint8_t *key, int key_bits);
func (ctx *AVCAMELLIA) AvCamelliaInit(key *ffcommon.FUint8T, key_bits ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_camellia_init").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(key)),
		uintptr(key_bits),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Encrypt or decrypt a buffer using a previously initialized context
 *
 * @param ctx an AVCAMELLIA context
 * @param dst destination array, can be equal to src
 * @param src source array, can be equal to dst
 * @param count number of 16 byte blocks
 * @paran iv initialization vector for CBC mode, NULL for ECB mode
 * @param decrypt 0 for encryption, 1 for decryption
 */
//void av_camellia_crypt(struct AVCAMELLIA *ctx, uint8_t *dst, const uint8_t *src, int count, uint8_t* iv, int decrypt);
func (ctx *AVCAMELLIA) AvCamelliaCrypt(dst, src *ffcommon.FUint8T, count ffcommon.FInt, iv *ffcommon.FUint8T, decrypt ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_camellia_crypt").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
		uintptr(count),
		uintptr(unsafe.Pointer(iv)),
		uintptr(decrypt),
	)
	if t == 0 {

	}
	return
}

/**
 * @}
 */
//#endif /* AVUTIL_CAMELLIA_H */
