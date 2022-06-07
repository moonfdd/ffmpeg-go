package libavutil

import (
	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"unsafe"
)

/*
 * DES encryption/decryption
 * Copyright (c) 2007 Reimar Doeffinger
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

//#ifndef AVUTIL_DES_H
//#define AVUTIL_DES_H
//
//#include <stdint.h>

/**
 * @defgroup lavu_des DES
 * @ingroup lavu_crypto
 * @{
 */

type AVDES struct {
	RoundKeys [3][16]ffcommon.FUint64T
	TripleDes ffcommon.FInt
}

/**
 * Allocate an AVDES context.
 */
//AVDES *av_des_alloc(void);
func AvDesAlloc() (res *AVDES) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_des_alloc").Call()
	if t == 0 {

	}
	res = (*AVDES)(unsafe.Pointer(t))
	return
}

/**
 * @brief Initializes an AVDES context.
 *
 * @param key_bits must be 64 or 192
 * @param decrypt 0 for encryption/CBC-MAC, 1 for decryption
 * @return zero on success, negative value otherwise
 */
//int av_des_init(struct AVDES *d, const uint8_t *key, int key_bits, int decrypt);
func (d *AVDES) AvDesInit(key *ffcommon.FUint8T, key_bits, decrypt ffcommon.FUint) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_des_init").Call(
		uintptr(unsafe.Pointer(d)),
		uintptr(unsafe.Pointer(key)),
		uintptr(key_bits),
		uintptr(decrypt),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * @brief Encrypts / decrypts using the DES algorithm.
 *
 * @param count number of 8 byte blocks
 * @param dst destination array, can be equal to src, must be 8-byte aligned
 * @param src source array, can be equal to dst, must be 8-byte aligned, may be NULL
 * @param iv initialization vector for CBC mode, if NULL then ECB will be used,
 *           must be 8-byte aligned
 * @param decrypt 0 for encryption, 1 for decryption
 */
//void av_des_crypt(struct AVDES *d, uint8_t *dst, const uint8_t *src, int count, uint8_t *iv, int decrypt);
func (d *AVDES) AvDesCrypt(dst, src *ffcommon.FUint8T, count ffcommon.FInt, iv *ffcommon.FUint8T, decrypt ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_des_crypt").Call(
		uintptr(unsafe.Pointer(d)),
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
 * @brief Calculates CBC-MAC using the DES algorithm.
 *
 * @param count number of 8 byte blocks
 * @param dst destination array, can be equal to src, must be 8-byte aligned
 * @param src source array, can be equal to dst, must be 8-byte aligned, may be NULL
 */
//void av_des_mac(struct AVDES *d, uint8_t *dst, const uint8_t *src, int count);
func (d *AVDES) AvDesMac(dst, src *ffcommon.FUint8T, count ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_des_mac").Call(
		uintptr(unsafe.Pointer(d)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
		uintptr(count),
	)
	if t == 0 {

	}
	return
}

/**
 * @}
 */

//#endif /* AVUTIL_DES_H */
