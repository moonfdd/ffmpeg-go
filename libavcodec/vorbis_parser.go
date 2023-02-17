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

/**
 * @file
 * A public API for Vorbis parsing
 *
 * Determines the duration for each packet.
 */

//#ifndef AVCODEC_VORBIS_PARSER_H
//#define AVCODEC_VORBIS_PARSER_H
//
//#include <stdint.h>

//typedef struct AVVorbisParseContext AVVorbisParseContext;
type AVVorbisParseContext struct {
}

/**
 * Allocate and initialize the Vorbis parser using headers in the extradata.
 */
//AVVorbisParseContext *av_vorbis_parse_init(const uint8_t *extradata,
//                                           int extradata_size);
func AvVorbisParseInit(extradata *ffcommon.FUint8T, extradata_size ffcommon.FInt) (res *AVVorbisParseContext) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_vorbis_parse_init").Call(
		uintptr(unsafe.Pointer(extradata)),
		uintptr(extradata_size),
	)
	res = (*AVVorbisParseContext)(unsafe.Pointer(t))
	return
}

/**
 * Free the parser and everything associated with it.
 */
//void av_vorbis_parse_free(AVVorbisParseContext **s);
func AvVorbisParseFree(s **AVVorbisParseContext) {
	ffcommon.GetAvcodecDll().NewProc("av_vorbis_parse_free").Call(
		uintptr(unsafe.Pointer(s)),
	)
}

const VORBIS_FLAG_HEADER = 0x00000001
const VORBIS_FLAG_COMMENT = 0x00000002
const VORBIS_FLAG_SETUP = 0x00000004

/**
 * Get the duration for a Vorbis packet.
 *
 * If @p flags is @c NULL,
 * special frames are considered invalid.
 *
 * @param s        Vorbis parser context
 * @param buf      buffer containing a Vorbis frame
 * @param buf_size size of the buffer
 * @param flags    flags for special frames
 */
//int av_vorbis_parse_frame_flags(AVVorbisParseContext *s, const uint8_t *buf,
//                                int buf_size, int *flags);
func (s *AVVorbisParseContext) AvVorbisParseFrameFlags(buf *ffcommon.FUint8T, buf_size ffcommon.FInt, flags *ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_vorbis_parse_frame_flags").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(buf)),
		uintptr(buf_size),
		uintptr(unsafe.Pointer(flags)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Get the duration for a Vorbis packet.
 *
 * @param s        Vorbis parser context
 * @param buf      buffer containing a Vorbis frame
 * @param buf_size size of the buffer
 */
//int av_vorbis_parse_frame(AVVorbisParseContext *s, const uint8_t *buf,
//                          int buf_size);
func (s *AVVorbisParseContext) AvVorbisParseFrame(buf *ffcommon.FUint8T, buf_size ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_vorbis_parse_frame").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(buf)),
		uintptr(buf_size),
	)
	res = ffcommon.FInt(t)
	return
}

//void av_vorbis_parse_reset(AVVorbisParseContext *s);
func (s *AVVorbisParseContext) AvVorbisParseReset() {
	ffcommon.GetAvcodecDll().NewProc("av_vorbis_parse_reset").Call(
		uintptr(unsafe.Pointer(s)),
	)
}

//#endif /* AVCODEC_VORBIS_PARSER_H */
