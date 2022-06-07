package libavutil

import (
	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"unsafe"
)

/*
 * copyright (c) 2006 Michael Niedermayer <michaelni@gmx.at>
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
 * @ingroup lavu_crc32
 * Public header for CRC hash function implementation.
 */

//#ifndef AVUTIL_CRC_H
//#define AVUTIL_CRC_H
//
//#include <stdint.h>
//#include <stddef.h>
//#include "attributes.h"
//#include "version.h"

/**
 * @defgroup lavu_crc32 CRC
 * @ingroup lavu_hash
 * CRC (Cyclic Redundancy Check) hash function implementation.
 *
 * This module supports numerous CRC polynomials, in addition to the most
 * widely used CRC-32-IEEE. See @ref AVCRCId for a list of available
 * polynomials.
 *
 * @{
 */

//typedef uint32_t AVCRC;
type AVCRC ffcommon.FUint32T
type AVCRCId = int32

const (
	AV_CRC_8_ATM = iota
	AV_CRC_16_ANSI
	AV_CRC_16_CCITT
	AV_CRC_32_IEEE
	AV_CRC_32_IEEE_LE /*< reversed bitorder version of AV_CRC_32_IEEE */
	AV_CRC_16_ANSI_LE /*< reversed bitorder version of AV_CRC_16_ANSI */
	AV_CRC_24_IEEE
	AV_CRC_8_EBU
	AV_CRC_MAX /*< Not part of public API! Do not use outside libavutil. */
)

/**
 * Initialize a CRC table.
 * @param ctx must be an array of size sizeof(AVCRC)*257 or sizeof(AVCRC)*1024
 * @param le If 1, the lowest bit represents the coefficient for the highest
 *           exponent of the corresponding polynomial (both for poly and
 *           actual CRC).
 *           If 0, you must swap the CRC parameter and the result of av_crc
 *           if you need the standard representation (can be simplified in
 *           most cases to e.g. bswap16):
 *           av_bswap32(crc << (32-bits))
 * @param bits number of bits for the CRC
 * @param poly generator polynomial without the x**bits coefficient, in the
 *             representation as specified by le
 * @param ctx_size size of ctx in bytes
 * @return <0 on failure
 */
//int av_crc_init(AVCRC *ctx, int le, int bits, uint32_t poly, int ctx_size);
func (ctx *AVCRC) AvCrcInit(le, bits ffcommon.FInt, poly ffcommon.FUint32T, ctx_size ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_crc_init").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(le),
		uintptr(bits),
		uintptr(poly),
		uintptr(ctx_size),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Get an initialized standard CRC table.
 * @param crc_id ID of a standard CRC
 * @return a pointer to the CRC table or NULL on failure
 */
//const AVCRC *av_crc_get_table(AVCRCId crc_id);
func AvCrcGetTable(crc_id AVCRCId) (res *AVCRC) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_crc_get_table").Call(
		uintptr(crc_id),
	)
	if t == 0 {

	}
	res = (*AVCRC)(unsafe.Pointer(t))
	return
}

/**
 * Calculate the CRC of a block.
 * @param crc CRC of previous blocks if any or initial value for CRC
 * @return CRC updated with the data from the given block
 *
 * @see av_crc_init() "le" parameter
 */
//uint32_t av_crc(const AVCRC *ctx, uint32_t crc,
//const uint8_t *buffer, size_t length) av_pure;
func (ctx *AVCRC) AvCrc(crc ffcommon.FUint32T,
	buffer *ffcommon.FUint8T, length ffcommon.FSizeT) (res ffcommon.FUint32T) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_crc").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(crc),
		uintptr(unsafe.Pointer(buffer)),
		uintptr(length),
	)
	if t == 0 {

	}
	res = ffcommon.FUint32T(t)
	return
}

/**
 * @}
 */

//#endif /* AVUTIL_CRC_H */
