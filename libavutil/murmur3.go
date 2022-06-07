package libavutil

import (
	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"unsafe"
)

/*
 * Copyright (C) 2013 Reimar Döffinger <Reimar.Doeffinger@gmx.de>
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
 * @ingroup lavu_murmur3
 * Public header for MurmurHash3 hash function implementation.
 */

//#ifndef AVUTIL_MURMUR3_H
//#define AVUTIL_MURMUR3_H
//
//#include <stddef.h>
//#include <stdint.h>
//
//#include "version.h"

/**
 * @defgroup lavu_murmur3 Murmur3
 * @ingroup lavu_hash
 * MurmurHash3 hash function implementation.
 *
 * MurmurHash3 is a non-cryptographic hash function, of which three
 * incompatible versions were created by its inventor Austin Appleby:
 *
 * - 32-bit output
 * - 128-bit output for 32-bit platforms
 * - 128-bit output for 64-bit platforms
 *
 * FFmpeg only implements the last variant: 128-bit output designed for 64-bit
 * platforms. Even though the hash function was designed for 64-bit platforms,
 * the function in reality works on 32-bit systems too, only with reduced
 * performance.
 *
 * @anchor lavu_murmur3_seedinfo
 * By design, MurmurHash3 requires a seed to operate. In response to this,
 * libavutil provides two functions for hash initiation, one that requires a
 * seed (av_murmur3_init_seeded()) and one that uses a fixed arbitrary integer
 * as the seed, and therefore does not (av_murmur3_init()).
 *
 * To make hashes comparable, you should provide the same seed for all calls to
 * this hash function -- if you are supplying one yourself, that is.
 *
 * @{
 */

/**
 * Allocate an AVMurMur3 hash context.
 *
 * @return Uninitialized hash context or `NULL` in case of error
 */
//struct AVMurMur3 *av_murmur3_alloc(void);
type AVMurMur3 struct {
}

func AvMurmur3Alloc() (res *AVMurMur3) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_murmur3_alloc").Call()
	if t == 0 {

	}
	res = (*AVMurMur3)(unsafe.Pointer(t))
	return
}

/**
 * Initialize or reinitialize an AVMurMur3 hash context with a seed.
 *
 * @param[out] c    Hash context
 * @param[in]  seed Random seed
 *
 * @see av_murmur3_init()
 * @see @ref lavu_murmur3_seedinfo "Detailed description" on a discussion of
 * seeds for MurmurHash3.
 */
//void av_murmur3_init_seeded(struct AVMurMur3 *c, uint64_t seed);
func (c *AVMurMur3) AvMurmur3InitSeeded(seed ffcommon.FUint64T) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_murmur3_init_seeded").Call(
		uintptr(unsafe.Pointer(c)),
		uintptr(seed),
	)
	if t == 0 {

	}
	return
}

/**
 * Initialize or reinitialize an AVMurMur3 hash context.
 *
 * Equivalent to av_murmur3_init_seeded() with a built-in seed.
 *
 * @param[out] c    Hash context
 *
 * @see av_murmur3_init_seeded()
 * @see @ref lavu_murmur3_seedinfo "Detailed description" on a discussion of
 * seeds for MurmurHash3.
 */
//void av_murmur3_init(struct AVMurMur3 *c);
func (c *AVMurMur3) AvMurmur3Init() {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_murmur3_init").Call(
		uintptr(unsafe.Pointer(c)),
	)
	if t == 0 {

	}
	return
}

/**
 * Update hash context with new data.
 *
 * @param[out] c    Hash context
 * @param[in]  src  Input data to update hash with
 * @param[in]  len  Number of bytes to read from `src`
 */
//#if FF_API_CRYPTO_SIZE_T
//void av_murmur3_update(struct AVMurMur3 *c, const uint8_t *src, int len);
//#else
//void av_murmur3_update(struct AVMurMur3 *c, const uint8_t *src, size_t len);
//#endif
func (c *AVMurMur3) AvMurmur3Update(src *ffcommon.FUint8T, len0 ffcommon.FIntOrSizeT) (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_murmur3_update").Call(
		uintptr(unsafe.Pointer(c)),
		uintptr(unsafe.Pointer(src)),
		uintptr(len0),
	)
	if t == 0 {

	}
	res = ffcommon.StringFromPtr(t)
	return
}

/**
 * Finish hashing and output digest value.
 *
 * @param[in,out] c    Hash context
 * @param[out]    dst  Buffer where output digest value is stored
 */
//void av_murmur3_final(struct AVMurMur3 *c, uint8_t dst[16]);
func (c *AVMurMur3) AvMurmur3Final(dst [16]ffcommon.FUint8T) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_murmur3_final").Call(
		uintptr(unsafe.Pointer(c)),
		uintptr(unsafe.Pointer(&dst)),
	)
	if t == 0 {

	}
	return
}

/**
 * @}
 */

//#endif /* AVUTIL_MURMUR3_H */
