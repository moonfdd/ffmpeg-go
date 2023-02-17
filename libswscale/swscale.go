package libswscale

import (
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/ffmpeg-go/libavutil"
)

/*
 * Copyright (C) 2001-2011 Michael Niedermayer <michaelni@gmx.at>
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

//#ifndef SWSCALE_SWSCALE_H
//const SWSCALE_SWSCALE_H
//
///**
// * @file
// * @ingroup libsws
// * external API header
// */
//
//#include <stdint.h>
//
//#include "../libavutil/avutil.h"
//#include "../libavutil/log.h"
//#include "../libavutil/pixfmt.h"
//#include "version.h"

/**
 * @defgroup libsws libswscale
 * Color conversion and scaling library.
 *
 * @{
 *
 * Return the LIBSWSCALE_VERSION_INT constant.
 */
//unsigned swscale_version(void);
//todo
func swscale_version() (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvswscaleDll().NewProc("swscale_version").Call()
	if t == 0 {

	}
	res = ffcommon.StringFromPtr(t)
	return
}

/**
 * Return the libswscale build-time configuration.
 */
//const char *swscale_configuration(void);
func SwscaleConfiguration() (res ffcommon.FConstCharP) {
	t, _, _ := ffcommon.GetAvswscaleDll().NewProc("swscale_configuration").Call()
	if t == 0 {

	}
	res = ffcommon.StringFromPtr(t)
	return
}

/**
 * Return the libswscale license.
 */
//const char *swscale_license(void);
func SwscaleLicense() (res ffcommon.FConstCharP) {
	t, _, _ := ffcommon.GetAvswscaleDll().NewProc("swscale_license").Call()
	if t == 0 {

	}
	res = ffcommon.StringFromPtr(t)
	return
}

/* values for the flags, the stuff on the command line is different */
const SWS_FAST_BILINEAR = 1
const SWS_BILINEAR = 2
const SWS_BICUBIC = 4
const SWS_X = 8
const SWS_POINT = 0x10
const SWS_AREA = 0x20
const SWS_BICUBLIN = 0x40
const SWS_GAUSS = 0x80
const SWS_SINC = 0x100
const SWS_LANCZOS = 0x200
const SWS_SPLINE = 0x400

const SWS_SRC_V_CHR_DROP_MASK = 0x30000
const SWS_SRC_V_CHR_DROP_SHIFT = 16

const SWS_PARAM_DEFAULT = 123456

const SWS_PRINT_INFO = 0x1000

//the following 3 flags are not completely implemented
//internal chrominance subsampling info
const SWS_FULL_CHR_H_INT = 0x2000

//input subsampling info
const SWS_FULL_CHR_H_INP = 0x4000
const SWS_DIRECT_BGR = 0x8000
const SWS_ACCURATE_RND = 0x40000
const SWS_BITEXACT = 0x80000
const SWS_ERROR_DIFFUSION = 0x800000

const SWS_MAX_REDUCE_CUTOFF = 0.002

const SWS_CS_ITU709 = 1
const SWS_CS_FCC = 4
const SWS_CS_ITU601 = 5
const SWS_CS_ITU624 = 5
const SWS_CS_SMPTE170M = 5
const SWS_CS_SMPTE240M = 7
const SWS_CS_DEFAULT = 5
const SWS_CS_BT2020 = 9

/**
 * Return a pointer to yuv<->rgb coefficients for the given colorspace
 * suitable for sws_setColorspaceDetails().
 *
 * @param colorspace One of the SWS_CS_* macros. If invalid,
 * SWS_CS_DEFAULT is used.
 */
//const int *sws_getCoefficients(int colorspace);
func SwsGetCoefficients(colorspace ffcommon.FInt) (res *ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvswscaleDll().NewProc("sws_getCoefficients").Call(
		uintptr(colorspace),
	)
	if t == 0 {

	}
	res = (*ffcommon.FInt)(unsafe.Pointer(t))
	return
}

// when used for filters they must have an odd number of elements
// coeffs cannot be shared between vectors
type SwsVector struct {
	Coeff  *ffcommon.FDouble ///< pointer to the list of coefficients
	Length ffcommon.FInt     ///< number of coefficients in the vector
}

// vectors can be shared
type SwsFilter struct {
	LumH *SwsVector
	LumV *SwsVector
	ChrH *SwsVector
	ChrV *SwsVector
}

//struct SwsContext;
type SwsContext struct {
}

/**
 * Return a positive value if pix_fmt is a supported input format, 0
 * otherwise.
 */
//int sws_isSupportedInput(enum AVPixelFormat pix_fmt);
type AVPixelFormat = libavutil.AVPixelFormat

func SwsIsSupportedInput(pix_fmt AVPixelFormat) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvswscaleDll().NewProc("sws_isSupportedInput").Call(
		uintptr(pix_fmt),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Return a positive value if pix_fmt is a supported output format, 0
 * otherwise.
 */
//int sws_isSupportedOutput(enum AVPixelFormat pix_fmt);
func SwsIsSupportedOutput(pix_fmt AVPixelFormat) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvswscaleDll().NewProc("sws_isSupportedOutput").Call(
		uintptr(pix_fmt),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * @param[in]  pix_fmt the pixel format
 * @return a positive value if an endianness conversion for pix_fmt is
 * supported, 0 otherwise.
 */
//int sws_isSupportedEndiannessConversion(enum AVPixelFormat pix_fmt);
func SwsIsSupportedEndiannessConversion(pix_fmt AVPixelFormat) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvswscaleDll().NewProc("sws_isSupportedEndiannessConversion").Call(
		uintptr(pix_fmt),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Allocate an empty SwsContext. This must be filled and passed to
 * sws_init_context(). For filling see AVOptions, options.c and
 * sws_setColorspaceDetails().
 */
//struct SwsContext *sws_alloc_context(void);
func SwsAllocContext() (res *SwsContext) {
	t, _, _ := ffcommon.GetAvswscaleDll().NewProc("sws_alloc_context").Call()
	if t == 0 {

	}
	res = (*SwsContext)(unsafe.Pointer(t))
	return
}

/**
 * Initialize the swscaler context sws_context.
 *
 * @return zero or positive value on success, a negative value on
 * error
 */
//av_warn_unused_result
//int sws_init_context(struct SwsContext *sws_context, SwsFilter *srcFilter, SwsFilter *dstFilter);
func (sws_context *SwsContext) SwsInitContext(srcFilter, dstFilter *SwsFilter) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvswscaleDll().NewProc("sws_init_context").Call(
		uintptr(unsafe.Pointer(sws_context)),
		uintptr(unsafe.Pointer(srcFilter)),
		uintptr(unsafe.Pointer(dstFilter)),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Free the swscaler context swsContext.
 * If swsContext is NULL, then does nothing.
 */
//void sws_freeContext(struct SwsContext *swsContext);
func (swsContext *SwsContext) SwsFreeContext() {
	t, _, _ := ffcommon.GetAvswscaleDll().NewProc("sws_freeContext").Call(
		uintptr(unsafe.Pointer(swsContext)),
	)
	if t == 0 {

	}
	return
}

/**
 * Allocate and return an SwsContext. You need it to perform
 * scaling/conversion operations using sws_scale().
 *
 * @param srcW the width of the source image
 * @param srcH the height of the source image
 * @param srcFormat the source image format
 * @param dstW the width of the destination image
 * @param dstH the height of the destination image
 * @param dstFormat the destination image format
 * @param flags specify which algorithm and options to use for rescaling
 * @param param extra parameters to tune the used scaler
 *              For SWS_BICUBIC param[0] and [1] tune the shape of the basis
 *              function, param[0] tunes f(1) and param[1] f´(1)
 *              For SWS_GAUSS param[0] tunes the exponent and thus cutoff
 *              frequency
 *              For SWS_LANCZOS param[0] tunes the width of the window function
 * @return a pointer to an allocated context, or NULL in case of error
 * @note this function is to be removed after a saner alternative is
 *       written
 */
//struct SwsContext *sws_getContext(int srcW, int srcH, enum AVPixelFormat srcFormat,
//int dstW, int dstH, enum AVPixelFormat dstFormat,
//int flags, SwsFilter *srcFilter,
//SwsFilter *dstFilter, const double *param);
func SwsGetContext(srcW, srcH ffcommon.FInt, srcFormat AVPixelFormat,
	dstW, dstH ffcommon.FInt, dstFormat AVPixelFormat,
	flags ffcommon.FInt, srcFilter,
	dstFilter *SwsFilter, param *ffcommon.FDouble) (res *SwsContext) {
	t, _, _ := ffcommon.GetAvswscaleDll().NewProc("sws_getContext").Call(
		uintptr(srcW),
		uintptr(srcH),
		uintptr(srcFormat),
		uintptr(dstW),
		uintptr(dstH),
		uintptr(dstFormat),
		uintptr(flags),
		uintptr(unsafe.Pointer(srcFilter)),
		uintptr(unsafe.Pointer(dstFilter)),
		uintptr(unsafe.Pointer(param)),
	)
	res = (*SwsContext)(unsafe.Pointer(t))
	return
}

/**
 * Scale the image slice in srcSlice and put the resulting scaled
 * slice in the image in dst. A slice is a sequence of consecutive
 * rows in an image.
 *
 * Slices have to be provided in sequential order, either in
 * top-bottom or bottom-top order. If slices are provided in
 * non-sequential order the behavior of the function is undefined.
 *
 * @param c         the scaling context previously created with
 *                  sws_getContext()
 * @param srcSlice  the array containing the pointers to the planes of
 *                  the source slice
 * @param srcStride the array containing the strides for each plane of
 *                  the source image
 * @param srcSliceY the position in the source image of the slice to
 *                  process, that is the number (counted starting from
 *                  zero) in the image of the first row of the slice
 * @param srcSliceH the height of the source slice, that is the number
 *                  of rows in the slice
 * @param dst       the array containing the pointers to the planes of
 *                  the destination image
 * @param dstStride the array containing the strides for each plane of
 *                  the destination image
 * @return          the height of the output slice
 */
//int sws_scale(struct SwsContext *c, const uint8_t *const srcSlice[],
//const int srcStride[], int srcSliceY, int srcSliceH,
//uint8_t *const dst[], const int dstStride[]);
func (c *SwsContext) SwsScale(srcSlice **ffcommon.FUint8T,
	srcStride *ffcommon.FInt, srcSliceY, srcSliceH ffcommon.FUint,
	dst **ffcommon.FUint8T, dstStride *ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvswscaleDll().NewProc("sws_scale").Call(
		uintptr(unsafe.Pointer(c)),
		uintptr(unsafe.Pointer(srcSlice)),
		uintptr(unsafe.Pointer(srcStride)),
		uintptr(srcSliceY),
		uintptr(srcSliceH),
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(dstStride)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * @param dstRange flag indicating the while-black range of the output (1=jpeg / 0=mpeg)
 * @param srcRange flag indicating the while-black range of the input (1=jpeg / 0=mpeg)
 * @param table the yuv2rgb coefficients describing the output yuv space, normally ff_yuv2rgb_coeffs[x]
 * @param inv_table the yuv2rgb coefficients describing the input yuv space, normally ff_yuv2rgb_coeffs[x]
 * @param brightness 16.16 fixed point brightness correction
 * @param contrast 16.16 fixed point contrast correction
 * @param saturation 16.16 fixed point saturation correction
 * @return -1 if not supported
 */
//int sws_setColorspaceDetails(struct SwsContext *c, const int inv_table[4],
//int srcRange, const int table[4], int dstRange,
//int brightness, int contrast, int saturation);
func (c *SwsContext) SwsSetColorspaceDetails(inv_table [4]*ffcommon.FInt,
	srcRange ffcommon.FInt, table [4]ffcommon.FInt, dstRange,
	brightness, contrast, saturation ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvswscaleDll().NewProc("sws_setColorspaceDetails").Call(
		uintptr(unsafe.Pointer(c)),
		uintptr(unsafe.Pointer(&inv_table)),
		uintptr(srcRange),
		uintptr(unsafe.Pointer(&table)),
		uintptr(dstRange),
		uintptr(brightness),
		uintptr(contrast),
		uintptr(saturation),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * @return -1 if not supported
 */
//int sws_getColorspaceDetails(struct SwsContext *c, int **inv_table,
//int *srcRange, int **table, int *dstRange,
//int *brightness, int *contrast, int *saturation);
func (c *SwsContext) SwsGetColorspaceDetails(inv_table **ffcommon.FInt,
	srcRange *ffcommon.FInt, table **ffcommon.FInt, dstRange,
	brightness, contrast, saturation *ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvswscaleDll().NewProc("sws_getColorspaceDetails").Call(
		uintptr(unsafe.Pointer(c)),
		uintptr(unsafe.Pointer(inv_table)),
		uintptr(unsafe.Pointer(srcRange)),
		uintptr(unsafe.Pointer(&table)),
		uintptr(unsafe.Pointer(dstRange)),
		uintptr(unsafe.Pointer(brightness)),
		uintptr(unsafe.Pointer(contrast)),
		uintptr(unsafe.Pointer(saturation)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Allocate and return an uninitialized vector with length coefficients.
 */
//SwsVector *sws_allocVec(int length);
func SwsAllocVec(length ffcommon.FInt) (res *SwsVector) {
	t, _, _ := ffcommon.GetAvswscaleDll().NewProc("sws_allocVec").Call(
		uintptr(length),
	)
	res = (*SwsVector)(unsafe.Pointer(t))
	return
}

/**
 * Return a normalized Gaussian curve used to filter stuff
 * quality = 3 is high quality, lower is lower quality.
 */
//SwsVector *sws_getGaussianVec(double variance, double quality);
func SwsGetGaussianVec(variance, quality ffcommon.FDouble) (res *SwsVector) {
	t, _, _ := ffcommon.GetAvswscaleDll().NewProc("sws_getGaussianVec").Call(
		uintptr(unsafe.Pointer(&variance)),
		uintptr(unsafe.Pointer(&quality)),
	)
	res = (*SwsVector)(unsafe.Pointer(t))
	return
}

/**
 * Scale all the coefficients of a by the scalar value.
 */
//void sws_scaleVec(SwsVector *a, double scalar);
func (a *SwsVector) SwsScaleVec(scalar ffcommon.FDouble) {
	ffcommon.GetAvswscaleDll().NewProc("sws_scaleVec").Call(
		uintptr(unsafe.Pointer(a)),
		uintptr(unsafe.Pointer(&scalar)),
	)
}

/**
 * Scale all the coefficients of a so that their sum equals height.
 */
//void sws_normalizeVec(SwsVector *a, double height);
func (a *SwsVector) SwsNormalizeVec(height ffcommon.FDouble) {
	ffcommon.GetAvswscaleDll().NewProc("sws_normalizeVec").Call(
		uintptr(unsafe.Pointer(a)),
		uintptr(unsafe.Pointer(&height)),
	)
}

//#if FF_API_SWS_VECTOR
//attribute_deprecated SwsVector *sws_getConstVec(double c, int length);
func SwsGetConstVec(c ffcommon.FDouble, length ffcommon.FInt) (res *SwsVector) {
	t, _, _ := ffcommon.GetAvswscaleDll().NewProc("sws_getConstVec").Call(
		uintptr(unsafe.Pointer(&c)),
		uintptr(length),
	)
	res = (*SwsVector)(unsafe.Pointer(t))
	return
}

//attribute_deprecated SwsVector *sws_getIdentityVec(void);
func SwsGetIdentityVec() (res *SwsVector) {
	t, _, _ := ffcommon.GetAvswscaleDll().NewProc("sws_getIdentityVec").Call()
	res = (*SwsVector)(unsafe.Pointer(t))
	return
}

//attribute_deprecated void sws_convVec(SwsVector *a, SwsVector *b);
func (a *SwsVector) SwsConvVec(b *SwsVector) {
	ffcommon.GetAvswscaleDll().NewProc("sws_convVec").Call(
		uintptr(unsafe.Pointer(a)),
		uintptr(unsafe.Pointer(b)),
	)
}

//attribute_deprecated void sws_addVec(SwsVector *a, SwsVector *b);
func (a *SwsVector) SwsAddVec(b *SwsVector) {
	ffcommon.GetAvswscaleDll().NewProc("sws_addVec").Call(
		uintptr(unsafe.Pointer(a)),
		uintptr(unsafe.Pointer(b)),
	)
}

//attribute_deprecated void sws_subVec(SwsVector *a, SwsVector *b);
func (a *SwsVector) SwsSubVec(b *SwsVector) {
	ffcommon.GetAvswscaleDll().NewProc("sws_subVec").Call(
		uintptr(unsafe.Pointer(a)),
		uintptr(unsafe.Pointer(b)),
	)
}

//attribute_deprecated void sws_shiftVec(SwsVector *a, int shift);
func (a *SwsVector) SwsShiftVec(shift ffcommon.FInt) {
	ffcommon.GetAvswscaleDll().NewProc("sws_shiftVec").Call(
		uintptr(unsafe.Pointer(a)),
		uintptr(shift),
	)
}

//attribute_deprecated SwsVector *sws_cloneVec(SwsVector *a);
func (a *SwsVector) SwsCloneVec() (res *SwsVector) {
	t, _, _ := ffcommon.GetAvswscaleDll().NewProc("sws_cloneVec").Call(
		uintptr(unsafe.Pointer(a)),
	)
	res = (*SwsVector)(unsafe.Pointer(t))
	return
}

//attribute_deprecated void sws_printVec2(SwsVector *a, AVClass *log_ctx, int log_level);
func (a *SwsVector) SwsPrintVec2(log_ctx *AVClass, log_level ffcommon.FInt) {
	ffcommon.GetAvswscaleDll().NewProc("sws_printVec2").Call(
		uintptr(unsafe.Pointer(a)),
		uintptr(unsafe.Pointer(log_ctx)),
		uintptr(log_level),
	)
}

//#endif

//void sws_freeVec(SwsVector *a);
func (a *SwsVector) SwsFreeVec() {
	ffcommon.GetAvswscaleDll().NewProc("sws_freeVec").Call(
		uintptr(unsafe.Pointer(a)),
	)
}

//SwsFilter *sws_getDefaultFilter(float lumaGBlur, float chromaGBlur,
//float lumaSharpen, float chromaSharpen,
//float chromaHShift, float chromaVShift,
//int verbose);
func SwsGetDefaultFilter(lumaGBlur, chromaGBlur,
	lumaSharpen, chromaSharpen,
	chromaHShift, chromaVShift ffcommon.FFloat,
	verbose ffcommon.FInt) (res *SwsFilter) {
	t, _, _ := ffcommon.GetAvswscaleDll().NewProc("sws_getDefaultFilter").Call(
		uintptr(unsafe.Pointer(&lumaGBlur)),
		uintptr(unsafe.Pointer(&chromaGBlur)),
		uintptr(unsafe.Pointer(&lumaSharpen)),
		uintptr(unsafe.Pointer(&chromaSharpen)),
		uintptr(unsafe.Pointer(&chromaHShift)),
		uintptr(unsafe.Pointer(&chromaVShift)),
		uintptr(verbose),
	)
	res = (*SwsFilter)(unsafe.Pointer(t))
	return
}

//void sws_freeFilter(SwsFilter *filter);
func (filter *SwsFilter) SwsFreeFilter() {
	ffcommon.GetAvswscaleDll().NewProc("sws_freeFilter").Call(
		uintptr(unsafe.Pointer(filter)),
	)
}

/**
 * Check if context can be reused, otherwise reallocate a new one.
 *
 * If context is NULL, just calls sws_getContext() to get a new
 * context. Otherwise, checks if the parameters are the ones already
 * saved in context. If that is the case, returns the current
 * context. Otherwise, frees context and gets a new context with
 * the new parameters.
 *
 * Be warned that srcFilter and dstFilter are not checked, they
 * are assumed to remain the same.
 */
//struct SwsContext *sws_getCachedContext(struct SwsContext *context,
//int srcW, int srcH, enum AVPixelFormat srcFormat,
//int dstW, int dstH, enum AVPixelFormat dstFormat,
//int flags, SwsFilter *srcFilter,
//SwsFilter *dstFilter, const double *param);
func (context *SwsContext) SwsGetCachedContext(srcW, srcH ffcommon.FInt, srcFormat AVPixelFormat,
	dstW, dstH ffcommon.FInt, dstFormat AVPixelFormat,
	flags ffcommon.FInt, srcFilter,
	dstFilter *SwsFilter, param *ffcommon.FDouble) (res *SwsContext) {
	t, _, _ := ffcommon.GetAvswscaleDll().NewProc("sws_getCachedContext").Call(
		uintptr(unsafe.Pointer(context)),
		uintptr(srcW),
		uintptr(srcH),
		uintptr(srcFormat),
		uintptr(dstW),
		uintptr(dstH),
		uintptr(dstFormat),
		uintptr(flags),
		uintptr(unsafe.Pointer(srcFilter)),
		uintptr(unsafe.Pointer(dstFilter)),
		uintptr(unsafe.Pointer(param)),
	)
	res = (*SwsContext)(unsafe.Pointer(t))
	return
}

/**
 * Convert an 8-bit paletted frame into a frame with a color depth of 32 bits.
 *
 * The output frame will have the same packed format as the palette.
 *
 * @param src        source frame buffer
 * @param dst        destination frame buffer
 * @param num_pixels number of pixels to convert
 * @param palette    array with [256] entries, which must match color arrangement (RGB or BGR) of src
 */
//void sws_convertPalette8ToPacked32(const uint8_t *src, uint8_t *dst, int num_pixels, const uint8_t *palette);
func SwsConvertPalette8ToPacked32(src, dst *ffcommon.FUint8T, num_pixels ffcommon.FInt, palette *ffcommon.FUint8T) {
	ffcommon.GetAvswscaleDll().NewProc("sws_convertPalette8ToPacked32").Call(
		uintptr(unsafe.Pointer(src)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(num_pixels),
		uintptr(unsafe.Pointer(palette)),
	)
}

/**
 * Convert an 8-bit paletted frame into a frame with a color depth of 24 bits.
 *
 * With the palette format "ABCD", the destination frame ends up with the format "ABC".
 *
 * @param src        source frame buffer
 * @param dst        destination frame buffer
 * @param num_pixels number of pixels to convert
 * @param palette    array with [256] entries, which must match color arrangement (RGB or BGR) of src
 */
//void sws_convertPalette8ToPacked24(const uint8_t *src, uint8_t *dst, int num_pixels, const uint8_t *palette);
func SwsConvertPalette8ToPacked24(src, dst *ffcommon.FUint8T, num_pixels ffcommon.FInt, palette *ffcommon.FUint8T) {
	ffcommon.GetAvswscaleDll().NewProc("sws_convertPalette8ToPacked24").Call(
		uintptr(unsafe.Pointer(src)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(num_pixels),
		uintptr(unsafe.Pointer(palette)),
	)
}

/**
 * Get the AVClass for swsContext. It can be used in combination with
 * AV_OPT_SEARCH_FAKE_OBJ for examining options.
 *
 * @see av_opt_find().
 */
//const AVClass *sws_get_class(void);
type AVClass = libavutil.AVClass

func SwsGetClass() (res *AVClass) {
	t, _, _ := ffcommon.GetAvswscaleDll().NewProc("sws_get_class").Call()
	res = (*AVClass)(unsafe.Pointer(t))
	return
}

/**
 * @}
 */

//#endif /* SWSCALE_SWSCALE_H */
