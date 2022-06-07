package libpostproc

import "github.com/moonfdd/ffmpeg-go/ffcommon"

/*
 * Copyright (C) 2001-2003 Michael Niedermayer (michaelni@gmx.at)
 *
 * This file is part of FFmpeg.
 *
 * FFmpeg is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * FFmpeg is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with FFmpeg; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA
 */

//#ifndef POSTPROC_POSTPROCESS_H
//const POSTPROC_POSTPROCESS_H

/**
 * @file
 * @ingroup lpp
 * external API header
 */

/**
 * @defgroup lpp libpostproc
 * Video postprocessing library.
 *
 * @{
 */

//#include "../libpostproc/version.h"

/**
 * Return the LIBPOSTPROC_VERSION_INT constant.
 */
//unsigned postproc_version(void);
func PostprocVersion() (res ffcommon.FUnsigned) {
	t, _, _ := ffcommon.GetAvpostprocDll().NewProc("postproc_version").Call()
	if t == 0 {

	}
	res = ffcommon.FUnsigned(t)
	return
}

/**
 * Return the libpostproc build-time configuration.
 */
//const char *postproc_configuration(void);
func PostprocConfiguration() (res ffcommon.FConstCharP) {
	t, _, _ := ffcommon.GetAvpostprocDll().NewProc("postproc_configuration").Call()
	if t == 0 {

	}
	res = ffcommon.StringFromPtr(t)
	return
}

/**
 * Return the libpostproc license.
 */
//const char *postproc_license(void);
func PostprocLicense() (res ffcommon.FConstCharP) {
	t, _, _ := ffcommon.GetAvpostprocDll().NewProc("postproc_license").Call()
	if t == 0 {

	}
	res = ffcommon.StringFromPtr(t)
	return
}

const PP_QUALITY_MAX = 6

//#include <inttypes.h>
//
//typedef void pp_context;
//typedef void pp_mode;
//
//#if LIBPOSTPROC_VERSION_INT < (52<<16)
//typedef pp_context pp_context_t;
//typedef pp_mode pp_mode_t;
//extern const char *const pp_help; ///< a simple help text
//#else
//extern const char pp_help[]; ///< a simple help text
//#endif
//
//void  pp_postprocess(const uint8_t * src[3], const int srcStride[3],
//uint8_t * dst[3], const int dstStride[3],
//int horizontalSize, int verticalSize,
//const int8_t *QP_store,  int QP_stride,
//pp_mode *mode, pp_context *ppContext, int pict_type);

/**
 * Return a pp_mode or NULL if an error occurred.
 *
 * @param name    the string after "-pp" on the command line
 * @param quality a number from 0 to PP_QUALITY_MAX
 */
//pp_mode *pp_get_mode_by_name_and_quality(const char *name, int quality);
func PpGetModeByNameAndQuality(name ffcommon.FConstCharP, quality ffcommon.FInt) (res ffcommon.FVoidP) {
	t, _, _ := ffcommon.GetAvpostprocDll().NewProc("pp_get_mode_by_name_and_quality").Call(
		ffcommon.UintPtrFromString(name),
		uintptr(quality),
	)
	if t == 0 {

	}
	res = t
	return
}

//void pp_free_mode(pp_mode *mode);
func PpFreeMode(mode ffcommon.FVoidP) {
	t, _, _ := ffcommon.GetAvpostprocDll().NewProc("pp_free_mode").Call(
		mode,
	)
	if t == 0 {

	}
	return
}

//pp_context *pp_get_context(int width, int height, int flags);
func PpGetContext(width, height, flags ffcommon.FInt) (res ffcommon.FVoidP) {
	t, _, _ := ffcommon.GetAvpostprocDll().NewProc("pp_get_context").Call(
		uintptr(width),
		uintptr(height),
		uintptr(flags),
	)
	if t == 0 {

	}
	res = t
	return
}

//void pp_free_context(pp_context *ppContext);
func PpFreeContext(ppContext ffcommon.FVoidP) {
	t, _, _ := ffcommon.GetAvpostprocDll().NewProc("pp_free_context").Call(
		ppContext,
	)
	if t == 0 {

	}
	return
}

const PP_CPU_CAPS_MMX = 0x80000000
const PP_CPU_CAPS_MMX2 = 0x20000000
const PP_CPU_CAPS_3DNOW = 0x40000000
const PP_CPU_CAPS_ALTIVEC = 0x10000000
const PP_CPU_CAPS_AUTO = 0x00080000

const PP_FORMAT = 0x00000008
const PP_FORMAT_420 = (0x00000011 | PP_FORMAT)
const PP_FORMAT_422 = (0x00000001 | PP_FORMAT)
const PP_FORMAT_411 = (0x00000002 | PP_FORMAT)
const PP_FORMAT_444 = (0x00000000 | PP_FORMAT)
const PP_FORMAT_440 = (0x00000010 | PP_FORMAT)

const PP_PICT_TYPE_QP2 = 0x00000010 ///< MPEG2 style QScale

/**
 * @}
 */

//#endif /* POSTPROC_POSTPROCESS_H */
