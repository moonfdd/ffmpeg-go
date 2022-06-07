package libavutil

import (
	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"unsafe"
)

/*
 * copyright (c) 2007 Michael Niedermayer <michaelni@gmx.at>
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

//#ifndef AVUTIL_AES_H
//#define AVUTIL_AES_H
//
//#include <stdint.h>
//
//#include "attributes.h"
//#include "version.h"
//
///**
// * @defgroup lavu_aes AES
// * @ingroup lavu_crypto
// * @{
// */
//
//extern const int av_aes_size;

//struct AVAES;
type AVAES struct {
}

/**
 * Allocate an AVAES context.
 */
//struct AVAES *av_aes_alloc(void);
func AvAesAlloc() (res *AVAES) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_aes_alloc").Call()
	if t == 0 {

	}
	res = (*AVAES)(unsafe.Pointer(t))
	return
}

/**
 * Initialize an AVAES context.
 * @param key_bits 128, 192 or 256
 * @param decrypt 0 for encryption, 1 for decryption
 */
//int av_aes_init(struct AVAES *a, const uint8_t *key, int key_bits, int decrypt);
func (a *AVAES) AvAesInit(key *ffcommon.FUint8T, key_bits, decrypt ffcommon.FUint) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_aes_init").Call(
		uintptr(unsafe.Pointer(a)),
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
 * Encrypt or decrypt a buffer using a previously initialized context.
 * @param count number of 16 byte blocks
 * @param dst destination array, can be equal to src
 * @param src source array, can be equal to dst
 * @param iv initialization vector for CBC mode, if NULL then ECB will be used
 * @param decrypt 0 for encryption, 1 for decryption
 */
//void av_aes_crypt(struct AVAES *a, uint8_t *dst, const uint8_t *src, int count, uint8_t *iv, int decrypt);
func (a *AVAES) AvAesCrypt(dst, src *ffcommon.FUint8T, count ffcommon.FInt, iv *ffcommon.FUint8T, decrypt ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_aes_crypt").Call(
		uintptr(unsafe.Pointer(a)),
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

//#endif /* AVUTIL_AES_H */
