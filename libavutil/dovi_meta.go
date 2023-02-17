package libavutil

import (
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
)

/*
 * Copyright (c) 2020 Vacing Fang <vacingfang@tencent.com>
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
 * DOVI configuration
 */

//#ifndef AVUTIL_DOVI_META_H
//#define AVUTIL_DOVI_META_H
//
//#include <stdint.h>
//#include <stddef.h>

/*
   * DOVI configuration
   * ref: dolby-vision-bitstreams-within-the-iso-base-media-file-format-v2.1.2
          dolby-vision-bitstreams-in-mpeg-2-transport-stream-multiplex-v1.2
   * @code
   * uint8_t  dv_version_major, the major version number that the stream complies with
   * uint8_t  dv_version_minor, the minor version number that the stream complies with
   * uint8_t  dv_profile, the Dolby Vision profile
   * uint8_t  dv_level, the Dolby Vision level
   * uint8_t  rpu_present_flag
   * uint8_t  el_present_flag
   * uint8_t  bl_present_flag
   * uint8_t  dv_bl_signal_compatibility_id
   * @endcode
   *
   * @note The struct must be allocated with av_dovi_alloc() and
   *       its size is not a part of the public ABI.
*/
type AVDOVIDecoderConfigurationRecord struct {
	DvVersionMajor            ffcommon.FUint8T
	DvVersionMinor            ffcommon.FUint8T
	DvProfile                 ffcommon.FUint8T
	DvLevel                   ffcommon.FUint8T
	RpuPresentFlag            ffcommon.FUint8T
	ElPresentFlag             ffcommon.FUint8T
	BlPresentFlag             ffcommon.FUint8T
	DvBlSignalCompatibilityId ffcommon.FUint8T
}

/**
 * Allocate a AVDOVIDecoderConfigurationRecord structure and initialize its
 * fields to default values.
 *
 * @return the newly allocated struct or NULL on failure
 */
//AVDOVIDecoderConfigurationRecord *av_dovi_alloc(size_t *size);
func AvDoviAlloc() (res *AVDOVIDecoderConfigurationRecord) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_dovi_alloc").Call()
	res = (*AVDOVIDecoderConfigurationRecord)(unsafe.Pointer(t))
	return
}

//#endif /* AVUTIL_DOVI_META_H */
