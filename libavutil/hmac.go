package libavutil

import (
	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"unsafe"
)

/*
 * Copyright (C) 2012 Martin Storsjo
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

//#ifndef AVUTIL_HMAC_H
//#define AVUTIL_HMAC_H
//
//#include <stdint.h>
//
//#include "version.h"
/**
 * @defgroup lavu_hmac HMAC
 * @ingroup lavu_crypto
 * @{
 */
type AVHMACType = int32

const (
	AV_HMAC_MD5 = iota
	AV_HMAC_SHA1
	AV_HMAC_SHA224
	AV_HMAC_SHA256
	AV_HMAC_SHA384
	AV_HMAC_SHA512
)

//typedef struct AVHMAC AVHMAC;
type AVHMAC struct {
}

/**
 * Allocate an AVHMAC context.
 * @param type The hash function used for the HMAC.
 */
//AVHMAC *av_hmac_alloc(enum AVHMACType type);
func AvHmacAlloc(type0 AVHMACType) (res *AVHMAC) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_hmac_alloc").Call(
		uintptr(type0),
	)
	if t == 0 {

	}
	res = (*AVHMAC)(unsafe.Pointer(t))
	return
}

/**
 * Free an AVHMAC context.
 * @param ctx The context to free, may be NULL
 */
//void av_hmac_free(AVHMAC *ctx);
func (ctx *AVHMAC) AvHmacFree() {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_hmac_free").Call(
		uintptr(unsafe.Pointer(ctx)),
	)
	if t == 0 {

	}
	return
}

/**
 * Initialize an AVHMAC context with an authentication key.
 * @param ctx    The HMAC context
 * @param key    The authentication key
 * @param keylen The length of the key, in bytes
 */
//void av_hmac_init(AVHMAC *ctx, const uint8_t *key, unsigned int keylen);
func (ctx *AVHMAC) AvHmacInit(key *ffcommon.FUint8T, keylen ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_hmac_init").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(key)),
		uintptr(keylen),
	)
	if t == 0 {

	}
	return
}

/**
 * Hash data with the HMAC.
 * @param ctx  The HMAC context
 * @param data The data to hash
 * @param len  The length of the data, in bytes
 */
//void av_hmac_update(AVHMAC *ctx, const uint8_t *data, unsigned int len);
func (ctx *AVHMAC) AvHmacUpdate(data *ffcommon.FUint8T, keylen ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_hmac_update").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(data)),
		uintptr(keylen),
	)
	if t == 0 {

	}
	return
}

/**
 * Finish hashing and output the HMAC digest.
 * @param ctx    The HMAC context
 * @param out    The output buffer to write the digest into
 * @param outlen The length of the out buffer, in bytes
 * @return       The number of bytes written to out, or a negative error code.
 */
//int av_hmac_final(AVHMAC *ctx, uint8_t *out, unsigned int outlen);
func (ctx *AVHMAC) AvHmacFinal(out *ffcommon.FUint8T, outlen ffcommon.FUnsignedInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_hmac_final").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(out)),
		uintptr(outlen),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Hash an array of data with a key.
 * @param ctx    The HMAC context
 * @param data   The data to hash
 * @param len    The length of the data, in bytes
 * @param key    The authentication key
 * @param keylen The length of the key, in bytes
 * @param out    The output buffer to write the digest into
 * @param outlen The length of the out buffer, in bytes
 * @return       The number of bytes written to out, or a negative error code.
 */
//int av_hmac_calc(AVHMAC *ctx, const uint8_t *data, unsigned int len,
//const uint8_t *key, unsigned int keylen,
//uint8_t *out, unsigned int outlen);
func (ctx *AVHMAC) AvHmacCalc(data *ffcommon.FUint8T, len0 ffcommon.FUnsignedInt,
	key *ffcommon.FUint8T, keylen ffcommon.FUnsignedInt,
	out *ffcommon.FUint8T, outlen ffcommon.FUnsignedInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_hmac_calc").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(data)),
		uintptr(len0),
		uintptr(unsafe.Pointer(key)),
		uintptr(keylen),
		uintptr(unsafe.Pointer(out)),
		uintptr(outlen),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * @}
 */

//#endif /* AVUTIL_HMAC_H */
