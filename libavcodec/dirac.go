package libavcodec

import (
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
)

/*
 * Copyright (C) 2007 Marco Gerards <marco@gnu.org>
 * Copyright (C) 2009 David Conrad
 * Copyright (C) 2011 Jordi Ortiz
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

//#ifndef AVCODEC_DIRAC_H
//#define AVCODEC_DIRAC_H
//
///**
// * @file
// * Interface to Dirac Decoder/Encoder
// * @author Marco Gerards <marco@gnu.org>
// * @author David Conrad
// * @author Jordi Ortiz
// */
//
//#include "avcodec.h"

/**
 * The spec limits the number of wavelet decompositions to 4 for both
 * level 1 (VC-2) and 128 (long-gop default).
 * 5 decompositions is the maximum before >16-bit buffers are needed.
 * Schroedinger allows this for DD 9,7 and 13,7 wavelets only, limiting
 * the others to 4 decompositions (or 3 for the fidelity filter).
 *
 * We use this instead of MAX_DECOMPOSITIONS to save some memory.
 */
const MAX_DWT_LEVELS = 5

/**
 * Parse code values:
 *
 * Dirac Specification ->
 * 9.6.1  Table 9.1
 *
 * VC-2 Specification  ->
 * 10.4.1 Table 10.1
 */
type DiracParseCodes int32

const (
	DIRAC_PCODE_SEQ_HEADER      = 0x00
	DIRAC_PCODE_END_SEQ         = 0x10
	DIRAC_PCODE_AUX             = 0x20
	DIRAC_PCODE_PAD             = 0x30
	DIRAC_PCODE_PICTURE_CODED   = 0x08
	DIRAC_PCODE_PICTURE_RAW     = 0x48
	DIRAC_PCODE_PICTURE_LOW_DEL = 0xC8
	DIRAC_PCODE_PICTURE_HQ      = 0xE8
	DIRAC_PCODE_INTER_NOREF_CO1 = 0x0A
	DIRAC_PCODE_INTER_NOREF_CO2 = 0x09
	DIRAC_PCODE_INTER_REF_CO1   = 0x0D
	DIRAC_PCODE_INTER_REF_CO2   = 0x0E
	DIRAC_PCODE_INTRA_REF_CO    = 0x0C
	DIRAC_PCODE_INTRA_REF_RAW   = 0x4C
	DIRAC_PCODE_INTRA_REF_PICT  = 0xCC
	DIRAC_PCODE_MAGIC           = 0x42424344
)

type DiracVersionInfo struct {
	Major ffcommon.FInt
	Minor ffcommon.FInt
}

type AVDiracSeqHeader struct {
	Width        ffcommon.FUnsigned
	Height       ffcommon.FUnsigned
	ChromaFormat ffcommon.FUint8T ///< 0: 444  1: 422  2: 420

	Interlaced    ffcommon.FUint8T
	TopFieldFirst ffcommon.FUint8T

	FrameRateIndex   ffcommon.FUint8T ///< index into dirac_frame_rate[]
	AspectRatioIndex ffcommon.FUint8T ///< index into dirac_aspect_ratio[]

	CleanWidth       ffcommon.FUint16T
	CleanHeight      ffcommon.FUint16T
	CleanLeftOffset  ffcommon.FUint16T
	CleanRightOffset ffcommon.FUint16T

	PixelRangeIndex ffcommon.FUint8T ///< index into dirac_pixel_range_presets[]
	ColorSpecIndex  ffcommon.FUint8T ///< index into dirac_color_spec_presets[]

	Profile ffcommon.FInt
	Level   ffcommon.FInt

	Framerate         AVRational
	SampleAspectRatio AVRational

	PixFmt         AVPixelFormat
	ColorRange     AVColorRange
	ColorPrimaries AVColorPrimaries
	ColorTrc       AVColorTransferCharacteristic
	Colorspace     AVColorSpace

	Version  DiracVersionInfo
	BitDepth ffcommon.FInt
}

/**
 * Parse a Dirac sequence header.
 *
 * @param dsh this function will allocate and fill an AVDiracSeqHeader struct
 *            and write it into this pointer. The caller must free it with
 *            av_free().
 * @param buf the data buffer
 * @param buf_size the size of the data buffer in bytes
 * @param log_ctx if non-NULL, this function will log errors here
 * @return 0 on success, a negative AVERROR code on failure
 */
//int av_dirac_parse_sequence_header(AVDiracSeqHeader **dsh,
//const uint8_t *buf, size_t buf_size,
//void *log_ctx);
func AvDiracParseSequenceHeader(dsh **AVDiracSeqHeader,
	buf *ffcommon.FUint8T, buf_size ffcommon.FSizeT,
	log_ctx ffcommon.FVoidP) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_dirac_parse_sequence_header").Call(
		uintptr(unsafe.Pointer(dsh)),
		uintptr(unsafe.Pointer(buf)),
		uintptr(buf_size),
		log_ctx,
	)
	res = ffcommon.FInt(t)
	return
}

//#endif /* AVCODEC_DIRAC_H */
