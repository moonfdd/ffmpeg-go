package libavcodec

import (
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
)

/*
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

//#ifndef AVCODEC_ADTS_PARSER_H
//#define AVCODEC_ADTS_PARSER_H
//
//#include <stddef.h>
//#include <stdint.h>

const AV_AAC_ADTS_HEADER_SIZE = 7

/**
 * Extract the number of samples and frames from AAC data.
 * @param[in]  buf     pointer to AAC data buffer
 * @param[out] samples Pointer to where number of samples is written
 * @param[out] frames  Pointer to where number of frames is written
 * @return Returns 0 on success, error code on failure.
 */
//int av_adts_header_parse(const uint8_t *buf, uint32_t *samples,
//uint8_t *frames);
func AvAdtsHeaderParse(buf *ffcommon.FUint8T, samples *ffcommon.FUint32T,
	frames *ffcommon.FUint8T) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_adts_header_parse").Call(
		uintptr(unsafe.Pointer(buf)),
		uintptr(unsafe.Pointer(samples)),
		uintptr(unsafe.Pointer(frames)),
	)
	res = ffcommon.FInt(t)
	return
}

//#endif /* AVCODEC_ADTS_PARSER_H */
