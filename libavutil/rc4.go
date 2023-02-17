package libavutil

import (
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
)

/*
 * RC4 encryption/decryption/pseudo-random number generator
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

//#ifndef AVUTIL_RC4_H
//#define AVUTIL_RC4_H
//
//#include <stdint.h>

/**
 * @defgroup lavu_rc4 RC4
 * @ingroup lavu_crypto
 * @{
 */

type AVRC4 struct {
	State [256]ffcommon.FUint8T
	X, Y  ffcommon.FInt
}

/**
 * Allocate an AVRC4 context.
 */
//AVRC4 *av_rc4_alloc(void);
func AvRc4Alloc() (res *AVRC4) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_rc4_alloc").Call()
	res = (*AVRC4)(unsafe.Pointer(t))
	return
}

/**
 * @brief Initializes an AVRC4 context.
 *
 * @param key_bits must be a multiple of 8
 * @param decrypt 0 for encryption, 1 for decryption, currently has no effect
 * @return zero on success, negative value otherwise
 */
//int av_rc4_init(struct AVRC4 *d, const uint8_t *key, int key_bits, int decrypt);
func (d *AVRC4) AvRc4Init(key *ffcommon.FUint8T, key_bits, decrypt ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_rc4_init").Call(
		uintptr(unsafe.Pointer(d)),
		uintptr(unsafe.Pointer(key)),
		uintptr(key_bits),
		uintptr(decrypt),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * @brief Encrypts / decrypts using the RC4 algorithm.
 *
 * @param count number of bytes
 * @param dst destination array, can be equal to src
 * @param src source array, can be equal to dst, may be NULL
 * @param iv not (yet) used for RC4, should be NULL
 * @param decrypt 0 for encryption, 1 for decryption, not (yet) used
 */
//void av_rc4_crypt(struct AVRC4 *d, uint8_t *dst, const uint8_t *src, int count, uint8_t *iv, int decrypt);
func (d *AVRC4) AvRc4Crypt(dst, src *ffcommon.FUint8T, count ffcommon.FInt, iv *ffcommon.FUint8T, decrypt ffcommon.FInt) {
	ffcommon.GetAvutilDll().NewProc("av_rc4_crypt").Call(
		uintptr(unsafe.Pointer(d)),
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

//#endif /* AVUTIL_RC4_H */
