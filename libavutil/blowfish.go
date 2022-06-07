package libavutil

import (
	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"unsafe"
)

/*
 * Blowfish algorithm
 * Copyright (c) 2012 Samuel Pitoiset
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

//#ifndef AVUTIL_BLOWFISH_H
//#define AVUTIL_BLOWFISH_H
//
//#include <stdint.h>

/**
 * @defgroup lavu_blowfish Blowfish
 * @ingroup lavu_crypto
 * @{
 */

const AV_BF_ROUNDS = 16

type AVBlowfish struct {
	P [AV_BF_ROUNDS + 2]ffcommon.FUint32T
	S [4][256]ffcommon.FUint32T
}

/**
 * Allocate an AVBlowfish context.
 */
//AVBlowfish *av_blowfish_alloc(void);
//todo
func AvBlowfishAlloc() (res *AVBlowfish) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_blowfish_alloc").Call()
	if t == 0 {

	}
	res = (*AVBlowfish)(unsafe.Pointer(t))
	return
}

/**
 * Initialize an AVBlowfish context.
 *
 * @param ctx an AVBlowfish context
 * @param key a key
 * @param key_len length of the key
 */
//void av_blowfish_init(struct AVBlowfish *ctx, const uint8_t *key, int key_len);
func (ctx *AVBlowfish) AvBlowfishInit(key *ffcommon.FUint8T, key_len ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_blowfish_init").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(key)),
		uintptr(key_len),
	)
	if t == 0 {

	}

	return
}

/**
 * Encrypt or decrypt a buffer using a previously initialized context.
 *
 * @param ctx an AVBlowfish context
 * @param xl left four bytes halves of input to be encrypted
 * @param xr right four bytes halves of input to be encrypted
 * @param decrypt 0 for encryption, 1 for decryption
 */
//void av_blowfish_crypt_ecb(struct AVBlowfish *ctx, uint32_t *xl, uint32_t *xr,
//int decrypt);
func (ctx *AVBlowfish) AvBlowfishCryptEcb(xl, xr *ffcommon.FUint32T, decrypt ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_blowfish_crypt_ecb").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(xl)),
		uintptr(unsafe.Pointer(xr)),
		uintptr(decrypt),
	)
	if t == 0 {

	}
	return
}

/**
 * Encrypt or decrypt a buffer using a previously initialized context.
 *
 * @param ctx an AVBlowfish context
 * @param dst destination array, can be equal to src
 * @param src source array, can be equal to dst
 * @param count number of 8 byte blocks
 * @param iv initialization vector for CBC mode, if NULL ECB will be used
 * @param decrypt 0 for encryption, 1 for decryption
 */
//void av_blowfish_crypt(struct AVBlowfish *ctx, uint8_t *dst, const uint8_t *src,
//int count, uint8_t *iv, int decrypt);
func (ctx *AVBlowfish) AvBlowfishCrypt(dst, src *ffcommon.FUint8T,
	count ffcommon.FInt, iv *ffcommon.FUint8T, decrypt ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_blowfish_crypt").Call(
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

//#endif /* AVUTIL_BLOWFISH_H */
