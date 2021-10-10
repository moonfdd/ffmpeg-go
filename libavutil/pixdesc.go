package libavutil

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"syscall"
	"unsafe"
)

/*
 * pixel format descriptor
 * Copyright (c) 2009 Michael Niedermayer <michaelni@gmx.at>
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

type AVComponentDescriptor struct {
	/**
	 * Which of the 4 planes contains the component.
	 */
	Plane ffcommon.FInt

	/**
	 * Number of elements between 2 horizontally consecutive pixels.
	 * Elements are bits for bitstream formats, bytes otherwise.
	 */
	Step ffcommon.FInt

	/**
	 * Number of elements before the component of the first pixel.
	 * Elements are bits for bitstream formats, bytes otherwise.
	 */
	Offset ffcommon.FInt

	/**
	 * Number of least significant bits that must be shifted away
	 * to get the value.
	 */
	Shift ffcommon.FInt

	/**
	 * Number of bits in the component.
	 */
	Depth ffcommon.FInt

	//#if FF_API_PLUS1_MINUS1
	/** deprecated, use step instead */
	StepMinus1 ffcommon.FInt

	/** deprecated, use depth instead */
	DepthMinus1 ffcommon.FInt

	/** deprecated, use offset instead */
	OffsetPlus1 ffcommon.FInt
	//#endif
}

/**
* Descriptor that unambiguously describes how the bits of a pixel are
* stored in the up to 4 data planes of an image. It also stores the
* subsampling factors and number of components.
*
* @note This is separate of the colorspace (RGB, YCbCr, YPbPr, JPEG-style YUV
*       and all the YUV variants) AVPixFmtDescriptor just stores how values
*       are stored not what these values represent.
 */
type AVPixFmtDescriptor struct {
	name          ffcommon.FBuf
	nb_components ffcommon.FUint8T ///< The number of components each pixel has, (1-4)

	/**
	 * Amount to shift the luma width right to find the chroma width.
	 * For YV12 this is 1 for example.
	 * chroma_width = AV_CEIL_RSHIFT(luma_width, log2_chroma_w)
	 * The note above is needed to ensure rounding up.
	 * This value only refers to the chroma components.
	 */
	log2_chroma_w ffcommon.FUint8T

	/**
	 * Amount to shift the luma height right to find the chroma height.
	 * For YV12 this is 1 for example.
	 * chroma_height= AV_CEIL_RSHIFT(luma_height, log2_chroma_h)
	 * The note above is needed to ensure rounding up.
	 * This value only refers to the chroma components.
	 */
	log2_chroma_h ffcommon.FUint8T

	/**
	 * Combination of AV_PIX_FMT_FLAG_... flags.
	 */
	flags ffcommon.FUint64T

	/**
	 * Parameters that describe how pixels are packed.
	 * If the format has 1 or 2 components, then luma is 0.
	 * If the format has 3 or 4 components:
	 *   if the RGB flag is set then 0 is red, 1 is green and 2 is blue;
	 *   otherwise 0 is luma, 1 is chroma-U and 2 is chroma-V.
	 *
	 * If present, the Alpha channel is always the last component.
	 */
	comp [4]AVComponentDescriptor

	/**
	 * Alternative comma-separated names.
	 */
	alias ffcommon.FBuf
}

/**
* Return the number of bits per pixel used by the pixel format
* described by pixdesc. Note that this is not the same as the number
* of bits per sample.
*
* The returned number of bits refers to the number of bits actually
* used for storing the pixel information, that is padding bits are
* not counted.
 */
//int av_get_bits_per_pixel(const AVPixFmtDescriptor *pixdesc);
//未测试
func (pixdesc *AVPixFmtDescriptor) AvGetBitsPerPixel() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_get_bits_per_pixel").Call(
		uintptr(unsafe.Pointer(pixdesc)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
* Return the number of bits per pixel for the pixel format
* described by pixdesc, including any padding or unused bits.
 */
//int av_get_padded_bits_per_pixel(const AVPixFmtDescriptor *pixdesc);
//未测试
func (pixdesc *AVPixFmtDescriptor) AvGetPaddedBitsPerPixel() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_get_padded_bits_per_pixel").Call(
		uintptr(unsafe.Pointer(pixdesc)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
* @return a pixel format descriptor for provided pixel format or NULL if
* this pixel format is unknown.
 */
//const AVPixFmtDescriptor *av_pix_fmt_desc_get(enum AVPixelFormat pix_fmt);
//未测试
func AvPixFmtDescGet(pix_fmt ffconstant.AVPixelFormat) (res *AVPixFmtDescriptor, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_pix_fmt_desc_get").Call(
		uintptr(pix_fmt),
	)
	if err != nil {
		//return
	}
	res = (*AVPixFmtDescriptor)(unsafe.Pointer(t))
	return
}

/**
* Iterate over all pixel format descriptors known to libavutil.
*
* @param prev previous descriptor. NULL to get the first descriptor.
*
* @return next descriptor or NULL after the last descriptor
 */
//const AVPixFmtDescriptor *av_pix_fmt_desc_next(const AVPixFmtDescriptor *prev);
//未测试
func (prev *AVPixFmtDescriptor) AvPixFmtDescNext() (res *AVPixFmtDescriptor, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_pix_fmt_desc_next").Call(
		uintptr(unsafe.Pointer(prev)),
	)
	if err != nil {
		//return
	}
	res = (*AVPixFmtDescriptor)(unsafe.Pointer(t))
	return
}

//
///**
//* @return an AVPixelFormat id described by desc, or AV_PIX_FMT_NONE if desc
//* is not a valid pointer to a pixel format descriptor.
//*/
//enum AVPixelFormat av_pix_fmt_desc_get_id(const AVPixFmtDescriptor *desc);
//未测试
func (desc *AVPixFmtDescriptor) AvPixFmtDescGetId() (res ffconstant.AVPixelFormat, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_pix_fmt_desc_get_id").Call(
		uintptr(unsafe.Pointer(desc)),
	)
	if err != nil {
		//return
	}
	res = ffconstant.AVPixelFormat(t)
	return
}

/**
* Utility function to access log2_chroma_w log2_chroma_h from
* the pixel format AVPixFmtDescriptor.
*
* @param[in]  pix_fmt the pixel format
* @param[out] h_shift store log2_chroma_w (horizontal/width shift)
* @param[out] v_shift store log2_chroma_h (vertical/height shift)
*
* @return 0 on success, AVERROR(ENOSYS) on invalid or unknown pixel format
 */
//int av_pix_fmt_get_chroma_sub_sample(enum AVPixelFormat pix_fmt,
//int *h_shift, int *v_shift);
//未测试
func AvPixFmtGetChromaSubSample(pix_fmt ffconstant.AVPixelFormat,
	h_shift *ffcommon.FInt, v_shift *ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_pix_fmt_desc_get_id").Call(
		uintptr(pix_fmt),
		uintptr(unsafe.Pointer(h_shift)),
		uintptr(unsafe.Pointer(v_shift)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
* @return number of planes in pix_fmt, a negative AVERROR if pix_fmt is not a
* valid pixel format.
 */
//int av_pix_fmt_count_planes(enum AVPixelFormat pix_fmt);
//未测试
func AvPixFmtCountPlanes(pix_fmt ffconstant.AVPixelFormat) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_pix_fmt_count_planes").Call(
		uintptr(pix_fmt),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
* @return the name for provided color range or NULL if unknown.
 */
//const char *av_color_range_name(enum AVColorRange range);
//未测试
func AvColorRangeName(range0 ffconstant.AVColorRange) (res ffcommon.FConstCharP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_color_range_name").Call(
		uintptr(range0),
	)
	if err != nil {
		//return
	}
	res = ffcommon.GoAStr(t)
	return
}

/**
* @return the AVColorRange value for name or an AVError if not found.
 */
//int av_color_range_from_name(const char *name);
//未测试
func AvColorRangeFromName(name ffcommon.FConstCharP) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_color_range_from_name").Call(
		uintptr(unsafe.Pointer(namep)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
* @return the name for provided color primaries or NULL if unknown.
 */
//const char *av_color_primaries_name(enum AVColorPrimaries primaries);
//未测试
func AvColorPrimariesName(name ffconstant.AVColorPrimaries) (res ffcommon.FConstCharP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_color_primaries_name").Call(
		uintptr(name),
	)
	if err != nil {
		//return
	}
	res = ffcommon.GoAStr(t)
	return
}

/**
* @return the AVColorPrimaries value for name or an AVError if not found.
 */
//int av_color_primaries_from_name(const char *name);
//未测试
func AvColorPrimariesFromName(name ffcommon.FConstCharP) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_color_primaries_from_name").Call(
		uintptr(unsafe.Pointer(namep)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

//
///**
//* @return the name for provided color transfer or NULL if unknown.
//*/
//const char *av_color_transfer_name(enum AVColorTransferCharacteristic transfer);
//未测试
func AvColorTransferName(transfer ffconstant.AVColorTransferCharacteristic) (res ffcommon.FConstCharP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_color_primaries_from_name").Call(
		uintptr(transfer),
	)
	if err != nil {
		//return
	}
	res = ffcommon.GoAStr(t)
	return
}

//
///**
//* @return the AVColorTransferCharacteristic value for name or an AVError if not found.
//*/
//int av_color_transfer_from_name(const char *name);
//未测试
func AvColorTransferFromName(name ffcommon.FConstCharP) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_color_transfer_from_name").Call(
		uintptr(unsafe.Pointer(namep)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

//
///**
//* @return the name for provided color space or NULL if unknown.
//*/
//const char *av_color_space_name(enum AVColorSpace space);
//未测试
func AvColorSpaceName(space ffconstant.AVColorSpace) (res ffcommon.FConstCharP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_color_space_name").Call(
		uintptr(space),
	)
	if err != nil {
		//return
	}
	res = ffcommon.GoAStr(t)
	return
}

//
///**
//* @return the AVColorSpace value for name or an AVError if not found.
//*/
//int av_color_space_from_name(const char *name);
//未测试
func AvColorSpaceFromName(name ffcommon.FConstCharP) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_color_space_from_name").Call(
		uintptr(unsafe.Pointer(namep)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

//
///**
//* @return the name for provided chroma location or NULL if unknown.
//*/
//const char *av_chroma_location_name(enum AVChromaLocation location);
//未测试
func AvChromaLocationName(location ffconstant.AVChromaLocation) (res ffcommon.FConstCharP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_chroma_location_name").Call(
		uintptr(location),
	)
	if err != nil {
		//return
	}
	res = ffcommon.GoAStr(t)
	return
}

//
///**
//* @return the AVChromaLocation value for name or an AVError if not found.
//*/
//int av_chroma_location_from_name(const char *name);
//未测试
func AvChromaLocationFromName(name ffcommon.FConstCharP) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_chroma_location_from_name").Call(
		uintptr(unsafe.Pointer(namep)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

//
///**
//* Return the pixel format corresponding to name.
//*
//* If there is no pixel format with name name, then looks for a
//* pixel format with the name corresponding to the native endian
//* format of name.
//* For example in a little-endian system, first looks for "gray16",
//* then for "gray16le".
//*
//* Finally if no pixel format has been found, returns AV_PIX_FMT_NONE.
//*/
//enum AVPixelFormat av_get_pix_fmt(const char *name);
//未测试
func AvGetPixFmt(name ffcommon.FConstCharP) (res ffconstant.AVPixelFormat, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_get_pix_fmt").Call(
		uintptr(unsafe.Pointer(namep)),
	)
	if err != nil {
		//return
	}
	res = ffconstant.AVPixelFormat(t)
	return
}

//
///**
//* Return the short name for a pixel format, NULL in case pix_fmt is
//* unknown.
//*
//* @see av_get_pix_fmt(), av_get_pix_fmt_string()
//*/
//const char *av_get_pix_fmt_name(enum AVPixelFormat pix_fmt);
//未测试
func AvGetPixFmtName(pix_fmt ffconstant.AVPixelFormat) (res ffcommon.FConstCharP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_get_pix_fmt_name").Call(
		uintptr(pix_fmt),
	)
	if err != nil {
		//return
	}
	res = ffcommon.GoAStr(t)
	return
}

//
///**
//* Print in buf the string corresponding to the pixel format with
//* number pix_fmt, or a header if pix_fmt is negative.
//*
//* @param buf the buffer where to write the string
//* @param buf_size the size of buf
//* @param pix_fmt the number of the pixel format to print the
//* corresponding info string, or a negative value to print the
//* corresponding header.
//*/
//char *av_get_pix_fmt_string(char *buf, int buf_size,
//enum AVPixelFormat pix_fmt);
//未测试
func AvGetPixFmtString(buf ffcommon.FBuf, buf_size ffcommon.FInt,
	pix_fmt ffconstant.AVPixelFormat) (res ffcommon.FConstCharP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_get_pix_fmt_string").Call(
		uintptr(unsafe.Pointer(buf)),
		uintptr(buf_size),
		uintptr(pix_fmt),
	)
	if err != nil {
		//return
	}
	res = ffcommon.GoAStr(t)
	return
}

/**
* Read a line from an image, and write the values of the
* pixel format component c to dst.
*
* @param data the array containing the pointers to the planes of the image
* @param linesize the array containing the linesizes of the image
* @param desc the pixel format descriptor for the image
* @param x the horizontal coordinate of the first pixel to read
* @param y the vertical coordinate of the first pixel to read
* @param w the width of the line to read, that is the number of
* values to write to dst
* @param read_pal_component if not zero and the format is a paletted
* format writes the values corresponding to the palette
* component c in data[1] to dst, rather than the palette indexes in
* data[0]. The behavior is undefined if the format is not paletted.
* @param dst_element_size size of elements in dst array (2 or 4 byte)
 */
//void av_read_image_line2(void *dst, const uint8_t *data[4],
//const int linesize[4], const AVPixFmtDescriptor *desc,
//int x, int y, int c, int w, int read_pal_component,
//int dst_element_size);
//未测试
func av_read_image_line2(dst ffcommon.FVoidP, data [4]*ffcommon.FUint8T,
	linesize [4]ffcommon.FInt, desc *AVPixFmtDescriptor,
	x ffcommon.FInt, y ffcommon.FInt, c ffcommon.FInt, w ffcommon.FInt, read_pal_component ffcommon.FInt,
	dst_element_size ffcommon.FInt) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_read_image_line2").Call(
		uintptr(dst),
		uintptr(unsafe.Pointer(&data)),
		uintptr(unsafe.Pointer(&linesize)),
		uintptr(unsafe.Pointer(desc)),
		uintptr(x),
		uintptr(y),
		uintptr(c),
		uintptr(w),
		uintptr(read_pal_component),
		uintptr(dst_element_size),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

//void av_read_image_line(uint16_t *dst, const uint8_t *data[4],
//const int linesize[4], const AVPixFmtDescriptor *desc,
//int x, int y, int c, int w, int read_pal_component);
//未测试
func AvReadImageLine(dst *ffcommon.FUint16T, data [4]*ffcommon.FUint8T,
	linesize [4]ffcommon.FInt, desc *AVPixFmtDescriptor,
	x ffcommon.FInt, y ffcommon.FInt, c ffcommon.FInt, w ffcommon.FInt, read_pal_component ffcommon.FInt) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_read_image_line").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(&data)),
		uintptr(unsafe.Pointer(&linesize)),
		uintptr(unsafe.Pointer(desc)),
		uintptr(x),
		uintptr(y),
		uintptr(c),
		uintptr(w),
		uintptr(read_pal_component),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
* Write the values from src to the pixel format component c of an
* image line.
*
* @param src array containing the values to write
* @param data the array containing the pointers to the planes of the
* image to write into. It is supposed to be zeroed.
* @param linesize the array containing the linesizes of the image
* @param desc the pixel format descriptor for the image
* @param x the horizontal coordinate of the first pixel to write
* @param y the vertical coordinate of the first pixel to write
* @param w the width of the line to write, that is the number of
* values to write to the image line
* @param src_element_size size of elements in src array (2 or 4 byte)
 */
//void av_write_image_line2(const void *src, uint8_t *data[4],
//const int linesize[4], const AVPixFmtDescriptor *desc,
//int x, int y, int c, int w, int src_element_size);
//未测试
func AvWriteImageLine2(dst ffcommon.FConstVoidP, data [4]*ffcommon.FUint8T,
	linesize [4]ffcommon.FInt, desc *AVPixFmtDescriptor,
	x ffcommon.FInt, y ffcommon.FInt, c ffcommon.FInt, w ffcommon.FInt, src_element_size ffcommon.FInt) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_write_image_line2").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(&data)),
		uintptr(unsafe.Pointer(&linesize)),
		uintptr(unsafe.Pointer(desc)),
		uintptr(x),
		uintptr(y),
		uintptr(c),
		uintptr(w),
		uintptr(src_element_size),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

//
//void av_write_image_line(const uint16_t *src, uint8_t *data[4],
//const int linesize[4], const AVPixFmtDescriptor *desc,
//int x, int y, int c, int w);
//未测试
func AvWriteImageLine(src *ffcommon.FUint16T, data [4]*ffcommon.FUint8T,
	linesize [4]ffcommon.FInt, desc *AVPixFmtDescriptor,
	x ffcommon.FInt, y ffcommon.FInt, c ffcommon.FInt, w ffcommon.FInt) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_write_image_line").Call(
		uintptr(unsafe.Pointer(src)),
		uintptr(unsafe.Pointer(&data)),
		uintptr(unsafe.Pointer(&linesize)),
		uintptr(unsafe.Pointer(desc)),
		uintptr(x),
		uintptr(y),
		uintptr(c),
		uintptr(w),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
* Utility function to swap the endianness of a pixel format.
*
* @param[in]  pix_fmt the pixel format
*
* @return pixel format with swapped endianness if it exists,
* otherwise AV_PIX_FMT_NONE
 */
//enum AVPixelFormat av_pix_fmt_swap_endianness(enum AVPixelFormat pix_fmt);
//未测试
func AvPixFmtSwapEndianness(pix_fmt ffconstant.AVPixelFormat) (res ffconstant.AVPixelFormat, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_pix_fmt_swap_endianness").Call(
		uintptr(pix_fmt),
	)
	if err != nil {
		//return
	}
	res = ffconstant.AVPixelFormat(t)
	return
}

/**
* Compute what kind of losses will occur when converting from one specific
* pixel format to another.
* When converting from one pixel format to another, information loss may occur.
* For example, when converting from RGB24 to GRAY, the color information will
* be lost. Similarly, other losses occur when converting from some formats to
* other formats. These losses can involve loss of chroma, but also loss of
* resolution, loss of color depth, loss due to the color space conversion, loss
* of the alpha bits or loss due to color quantization.
* av_get_fix_fmt_loss() informs you about the various types of losses
* which will occur when converting from one pixel format to another.
*
* @param[in] dst_pix_fmt destination pixel format
* @param[in] src_pix_fmt source pixel format
* @param[in] has_alpha Whether the source pixel format alpha channel is used.
* @return Combination of flags informing you what kind of losses will occur
* (maximum loss for an invalid dst_pix_fmt).
 */
//int av_get_pix_fmt_loss(enum AVPixelFormat dst_pix_fmt,
//enum AVPixelFormat src_pix_fmt,
//int has_alpha);
//未测试
func AvGetPixFmtLoss(dst_pix_fmt ffconstant.AVPixelFormat,
	src_pix_fmt ffconstant.AVPixelFormat,
	has_alpha ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_get_pix_fmt_loss").Call(
		uintptr(dst_pix_fmt),
		uintptr(src_pix_fmt),
		uintptr(has_alpha),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
* Compute what kind of losses will occur when converting from one specific
* pixel format to another.
* When converting from one pixel format to another, information loss may occur.
* For example, when converting from RGB24 to GRAY, the color information will
* be lost. Similarly, other losses occur when converting from some formats to
* other formats. These losses can involve loss of chroma, but also loss of
* resolution, loss of color depth, loss due to the color space conversion, loss
* of the alpha bits or loss due to color quantization.
* av_get_fix_fmt_loss() informs you about the various types of losses
* which will occur when converting from one pixel format to another.
*
* @param[in] dst_pix_fmt destination pixel format
* @param[in] src_pix_fmt source pixel format
* @param[in] has_alpha Whether the source pixel format alpha channel is used.
* @return Combination of flags informing you what kind of losses will occur
* (maximum loss for an invalid dst_pix_fmt).
 */
//enum AVPixelFormat av_find_best_pix_fmt_of_2(enum AVPixelFormat dst_pix_fmt1, enum AVPixelFormat dst_pix_fmt2,
//enum AVPixelFormat src_pix_fmt, int has_alpha, int *loss_ptr);
//未测试
func AvFindBestPixFmtOf2(dst_pix_fmt1 ffconstant.AVPixelFormat, dst_pix_fmt2 ffconstant.AVPixelFormat,
	src_pix_fmt ffconstant.AVPixelFormat, has_alpha ffcommon.FInt, loss_ptr *ffcommon.FInt) (res ffconstant.AVPixelFormat, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_find_best_pix_fmt_of_2").Call(
		uintptr(dst_pix_fmt1),
		uintptr(dst_pix_fmt2),
		uintptr(src_pix_fmt),
		uintptr(has_alpha),
		uintptr(unsafe.Pointer(loss_ptr)),
	)
	if err != nil {
		//return
	}
	res = ffconstant.AVPixelFormat(t)
	return
}
