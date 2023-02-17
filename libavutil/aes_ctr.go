package libavutil

import (
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
)

/*
 * AES-CTR cipher
 * Copyright (c) 2015 Eran Kornblau <erankor at gmail dot com>
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

//#ifndef AVUTIL_AES_CTR_H
//#define AVUTIL_AES_CTR_H
//
//#include <stdint.h>
//
//#include "attributes.h"
//#include "version.h"

const AES_CTR_KEY_SIZE = (16)
const AES_CTR_IV_SIZE = (8)

//struct AVAESCTR;
type AVAESCTR struct {
}

/**
 * Allocate an AVAESCTR context.
 */
//struct AVAESCTR *av_aes_ctr_alloc(void);
func AvAesCtrAlloc() (res *AVAESCTR) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_aes_ctr_alloc").Call()
	res = (*AVAESCTR)(unsafe.Pointer(t))
	return
}

/**
 * Initialize an AVAESCTR context.
 * @param key encryption key, must have a length of AES_CTR_KEY_SIZE
 */
//int av_aes_ctr_init(struct AVAESCTR *a, const uint8_t *key);
func (a *AVAESCTR) AvAesCtrInit(key *ffcommon.FUint8T) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_aes_ctr_init").Call(
		uintptr(unsafe.Pointer(a)),
		uintptr(unsafe.Pointer(key)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Release an AVAESCTR context.
 */
//void av_aes_ctr_free(struct AVAESCTR *a);
func (a *AVAESCTR) AvAesCtrFree() {
	ffcommon.GetAvutilDll().NewProc("av_aes_ctr_free").Call(
		uintptr(unsafe.Pointer(a)),
	)
}

/**
 * Process a buffer using a previously initialized context.
 * @param dst destination array, can be equal to src
 * @param src source array, can be equal to dst
 * @param size the size of src and dst
 */
//void av_aes_ctr_crypt(struct AVAESCTR *a, uint8_t *dst, const uint8_t *src, int size);
func (a *AVAESCTR) AvAesCtrCrypt(dst, src *ffcommon.FUint8T, size ffcommon.FUint) {
	ffcommon.GetAvutilDll().NewProc("av_aes_ctr_crypt").Call(
		uintptr(unsafe.Pointer(a)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
		uintptr(size),
	)
}

/**
 * Get the current iv
 */
//const uint8_t* av_aes_ctr_get_iv(struct AVAESCTR *a);
func (a *AVAESCTR) AvAesCtrGetIv() (res *ffcommon.FUint8T) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_aes_ctr_get_iv").Call(
		uintptr(unsafe.Pointer(a)),
	)
	res = (*ffcommon.FUint8T)(unsafe.Pointer(t))
	return
}

/**
 * Generate a random iv
 */
//void av_aes_ctr_set_random_iv(struct AVAESCTR *a);
func (a *AVAESCTR) AvAesCtrSetRandomIv() {
	ffcommon.GetAvutilDll().NewProc("av_aes_ctr_set_random_iv").Call(
		uintptr(unsafe.Pointer(a)),
	)
}

/**
 * Forcefully change the 8-byte iv
 */
//void av_aes_ctr_set_iv(struct AVAESCTR *a, const uint8_t* iv);
func (a *AVAESCTR) AvAesCtrSetIv(iv *ffcommon.FUint8T) {
	ffcommon.GetAvutilDll().NewProc("av_aes_ctr_set_iv").Call(
		uintptr(unsafe.Pointer(a)),
		uintptr(unsafe.Pointer(iv)),
	)
}

/**
 * Forcefully change the "full" 16-byte iv, including the counter
 */
//void av_aes_ctr_set_full_iv(struct AVAESCTR *a, const uint8_t* iv);
func (a *AVAESCTR) AvAesCtrSetFullIv(iv *ffcommon.FUint8T) {
	ffcommon.GetAvutilDll().NewProc("av_aes_ctr_set_full_iv").Call(
		uintptr(unsafe.Pointer(a)),
		uintptr(unsafe.Pointer(iv)),
	)
}

/**
 * Increment the top 64 bit of the iv (performed after each frame)
 */
//void av_aes_ctr_increment_iv(struct AVAESCTR *a);
func (a *AVAESCTR) AvAesCtrIncrementIv() {
	ffcommon.GetAvutilDll().NewProc("av_aes_ctr_increment_iv").Call(
		uintptr(unsafe.Pointer(a)),
	)
}

/**
 * @}
 */

//#endif /* AVUTIL_AES_CTR_H */
