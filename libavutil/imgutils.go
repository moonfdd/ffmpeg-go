package libavutil

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"unsafe"
)

/**
 * Compute the max pixel step for each plane of an image with a
 * format described by pixdesc.
 *
 * The pixel step is the distance in bytes between the first byte of
 * the group of bytes which describe a pixel component and the first
 * byte of the successive group in the same plane for the same
 * component.
 *
 * @param max_pixsteps an array which is filled with the max pixel step
 * for each plane. Since a plane may contain different pixel
 * components, the computed max_pixsteps[plane] is relative to the
 * component in the plane with the max pixel step.
 * @param max_pixstep_comps an array which is filled with the component
 * for each plane which has the max pixel step. May be NULL.
 */
//void av_image_fill_max_pixsteps(int max_pixsteps[4], int max_pixstep_comps[4],
//const AVPixFmtDescriptor *pixdesc);
//未测试
func AvImageFillMaxPixsteps(max_pixsteps [4]ffcommon.FInt, max_pixstep_comps [4]ffcommon.FInt,
	pixdesc *AVPixFmtDescriptor) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_image_fill_max_pixsteps").Call(
		uintptr(unsafe.Pointer(&max_pixsteps)),
		uintptr(unsafe.Pointer(&max_pixstep_comps)),
		uintptr(unsafe.Pointer(pixdesc)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Compute the size of an image line with format pix_fmt and width
 * width for the plane plane.
 *
 * @return the computed size in bytes
 */
//int av_image_get_linesize(enum AVPixelFormat pix_fmt, int width, int plane);
//未测试
func AvImageGetLinesize(pix_fmt ffconstant.AVPixelFormat, width ffcommon.FInt, planer ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_image_get_linesize").Call(
		uintptr(pix_fmt),
		uintptr(width),
		uintptr(planer),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Fill plane linesizes for an image with pixel format pix_fmt and
 * width width.
 *
 * @param linesizes array to be filled with the linesize for each plane
 * @return >= 0 in case of success, a negative error code otherwise
 */
//int av_image_fill_linesizes(int linesizes[4], enum AVPixelFormat pix_fmt, int width);
//未测试
func AvImageFillLinesizes(linesizes [4]ffcommon.FInt, pix_fmt ffconstant.AVPixelFormat, width ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_image_fill_linesizes").Call(
		uintptr(unsafe.Pointer(&linesizes)),
		uintptr(pix_fmt),
		uintptr(width),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Fill plane sizes for an image with pixel format pix_fmt and height height.
 *
 * @param size the array to be filled with the size of each image plane
 * @param linesizes the array containing the linesize for each
 *        plane, should be filled by av_image_fill_linesizes()
 * @return >= 0 in case of success, a negative error code otherwise
 *
 * @note The linesize parameters have the type ptrdiff_t here, while they are
 *       int for av_image_fill_linesizes().
 */
//int av_image_fill_plane_sizes(size_t size[4], enum AVPixelFormat pix_fmt,
//int height, const ptrdiff_t linesizes[4]);
//未测试
func AvImageFillPlaneSizes(size [4]ffcommon.FSizeT, pix_fmt ffconstant.AVPixelFormat,
	height ffcommon.FInt, linesizes [4]ffcommon.FPtrdiffT) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_image_fill_plane_sizes").Call(
		uintptr(unsafe.Pointer(&size)),
		uintptr(pix_fmt),
		uintptr(height),
		uintptr(unsafe.Pointer(&linesizes)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Fill plane data pointers for an image with pixel format pix_fmt and
 * height height.
 *
 * @param data pointers array to be filled with the pointer for each image plane
 * @param ptr the pointer to a buffer which will contain the image
 * @param linesizes the array containing the linesize for each
 * plane, should be filled by av_image_fill_linesizes()
 * @return the size in bytes required for the image buffer, a negative
 * error code in case of failure
 */
//int av_image_fill_pointers(uint8_t *data[4], enum AVPixelFormat pix_fmt, int height,
//uint8_t *ptr, const int linesizes[4]);
//未测试
func AvImageFillPointers(data [4]*ffcommon.FUint8T, pix_fmt ffconstant.AVPixelFormat, height ffcommon.FInt,
	ptr *ffcommon.FUint8T, linesizes [4]ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_image_fill_pointers").Call(
		uintptr(unsafe.Pointer(&data)),
		uintptr(pix_fmt),
		uintptr(height),
		uintptr(unsafe.Pointer(ptr)),
		uintptr(unsafe.Pointer(&linesizes)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Allocate an image with size w and h and pixel format pix_fmt, and
 * fill pointers and linesizes accordingly.
 * The allocated image buffer has to be freed by using
 * av_freep(&pointers[0]).
 *
 * @param align the value to use for buffer size alignment
 * @return the size in bytes required for the image buffer, a negative
 * error code in case of failure
 */
//int av_image_alloc(uint8_t *pointers[4], int linesizes[4],
//int w, int h, enum AVPixelFormat pix_fmt, int align);
//未测试
func AvImageAlloc(pointers [4]*ffcommon.FUint8T, linesizes [4]ffcommon.FInt,
	w ffcommon.FInt, h ffcommon.FInt, pix_fmt ffconstant.AVPixelFormat, align ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_image_alloc").Call(
		uintptr(unsafe.Pointer(&pointers)),
		uintptr(unsafe.Pointer(&linesizes)),
		uintptr(w),
		uintptr(h),
		uintptr(pix_fmt),
		uintptr(align),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Copy image plane from src to dst.
 * That is, copy "height" number of lines of "bytewidth" bytes each.
 * The first byte of each successive line is separated by *_linesize
 * bytes.
 *
 * bytewidth must be contained by both absolute values of dst_linesize
 * and src_linesize, otherwise the function behavior is undefined.
 *
 * @param dst_linesize linesize for the image plane in dst
 * @param src_linesize linesize for the image plane in src
 */
//void av_image_copy_plane(uint8_t       *dst, int dst_linesize,
//const uint8_t *src, int src_linesize,
//int bytewidth, int height);
//未测试
func AvImageCopyPlane(dst *ffcommon.FUint8T, dst_linesize ffcommon.FInt,
	src *ffcommon.FUint8T, src_linesize ffcommon.FInt,
	bytewidth ffcommon.FInt, height ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_image_copy_plane").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(dst_linesize),
		uintptr(unsafe.Pointer(src)),
		uintptr(src_linesize),
		uintptr(bytewidth),
		uintptr(height),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Copy image in src_data to dst_data.
 *
 * @param dst_linesizes linesizes for the image in dst_data
 * @param src_linesizes linesizes for the image in src_data
 */
//void av_image_copy(uint8_t *dst_data[4], int dst_linesizes[4],
//const uint8_t *src_data[4], const int src_linesizes[4],
//enum AVPixelFormat pix_fmt, int width, int height);
//未测试
func AvImageCopy(dst_data [4]*ffcommon.FUint8T, dst_linesizes [4]ffcommon.FInt,
	src_data [4]*ffcommon.FUint8T, src_linesizes [4]ffcommon.FInt,
	pix_fmt ffconstant.AVPixelFormat, width ffcommon.FInt, height ffcommon.FInt) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_image_copy").Call(
		uintptr(unsafe.Pointer(&dst_data)),
		uintptr(unsafe.Pointer(&dst_linesizes)),
		uintptr(unsafe.Pointer(&src_data)),
		uintptr(unsafe.Pointer(&src_linesizes)),
		uintptr(pix_fmt),
		uintptr(width),
		uintptr(height),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Copy image data located in uncacheable (e.g. GPU mapped) memory. Where
 * available, this function will use special functionality for reading from such
 * memory, which may result in greatly improved performance compared to plain
 * av_image_copy().
 *
 * The data pointers and the linesizes must be aligned to the maximum required
 * by the CPU architecture.
 *
 * @note The linesize parameters have the type ptrdiff_t here, while they are
 *       int for av_image_copy().
 * @note On x86, the linesizes currently need to be aligned to the cacheline
 *       size (i.e. 64) to get improved performance.
 */
//void av_image_copy_uc_from(uint8_t *dst_data[4],       const ptrdiff_t dst_linesizes[4],
//const uint8_t *src_data[4], const ptrdiff_t src_linesizes[4],
//enum AVPixelFormat pix_fmt, int width, int height);
//未测试
func AvImageCopyUcFrom(dst_data [4]*ffcommon.FUint8T, dst_linesizes [4]ffcommon.FPtrdiffT,
	src_data [4]*ffcommon.FUint8T, src_linesizes [4]ffcommon.FPtrdiffT,
	pix_fmt ffconstant.AVPixelFormat, width ffcommon.FInt, height ffcommon.FInt) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_image_copy_uc_from").Call(
		uintptr(unsafe.Pointer(&dst_data)),
		uintptr(unsafe.Pointer(&dst_linesizes)),
		uintptr(unsafe.Pointer(&src_data)),
		uintptr(unsafe.Pointer(&src_linesizes)),
		uintptr(pix_fmt),
		uintptr(width),
		uintptr(height),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Setup the data pointers and linesizes based on the specified image
 * parameters and the provided array.
 *
 * The fields of the given image are filled in by using the src
 * address which points to the image data buffer. Depending on the
 * specified pixel format, one or multiple image data pointers and
 * line sizes will be set.  If a planar format is specified, several
 * pointers will be set pointing to the different picture planes and
 * the line sizes of the different planes will be stored in the
 * lines_sizes array. Call with src == NULL to get the required
 * size for the src buffer.
 *
 * To allocate the buffer and fill in the dst_data and dst_linesize in
 * one call, use av_image_alloc().
 *
 * @param dst_data      data pointers to be filled in
 * @param dst_linesize  linesizes for the image in dst_data to be filled in
 * @param src           buffer which will contain or contains the actual image data, can be NULL
 * @param pix_fmt       the pixel format of the image
 * @param width         the width of the image in pixels
 * @param height        the height of the image in pixels
 * @param align         the value used in src for linesize alignment
 * @return the size in bytes required for src, a negative error code
 * in case of failure
 */
//int av_image_fill_arrays(uint8_t *dst_data[4], int dst_linesize[4],
//const uint8_t *src,
//enum AVPixelFormat pix_fmt, int width, int height, int align);
//未测试
func AvImageFillArrays(dst_data [4]*ffcommon.FUint8T, dst_linesizes [4]ffcommon.FPtrdiffT,
	src *ffcommon.FUint8T,
	src_linesizes [4]ffcommon.FPtrdiffT,
	pix_fmt ffconstant.AVPixelFormat, width ffcommon.FInt, height ffcommon.FInt, align ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_image_fill_arrays").Call(
		uintptr(unsafe.Pointer(&dst_data)),
		uintptr(unsafe.Pointer(&dst_linesizes)),
		uintptr(unsafe.Pointer(src)),
		uintptr(unsafe.Pointer(&src_linesizes)),
		uintptr(pix_fmt),
		uintptr(width),
		uintptr(height),
		uintptr(align),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Return the size in bytes of the amount of data required to store an
 * image with the given parameters.
 *
 * @param pix_fmt  the pixel format of the image
 * @param width    the width of the image in pixels
 * @param height   the height of the image in pixels
 * @param align    the assumed linesize alignment
 * @return the buffer size in bytes, a negative error code in case of failure
 */
//int av_image_get_buffer_size(enum AVPixelFormat pix_fmt, int width, int height, int align);
//未测试
func AvImageGetBufferSize(pix_fmt ffconstant.AVPixelFormat, width ffcommon.FInt, height ffcommon.FInt, align ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_image_get_buffer_size").Call(
		uintptr(pix_fmt),
		uintptr(width),
		uintptr(height),
		uintptr(align),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Copy image data from an image into a buffer.
 *
 * av_image_get_buffer_size() can be used to compute the required size
 * for the buffer to fill.
 *
 * @param dst           a buffer into which picture data will be copied
 * @param dst_size      the size in bytes of dst
 * @param src_data      pointers containing the source image data
 * @param src_linesize  linesizes for the image in src_data
 * @param pix_fmt       the pixel format of the source image
 * @param width         the width of the source image in pixels
 * @param height        the height of the source image in pixels
 * @param align         the assumed linesize alignment for dst
 * @return the number of bytes written to dst, or a negative value
 * (error code) on error
 */
//int av_image_copy_to_buffer(uint8_t *dst, int dst_size,
//const uint8_t * const src_data[4], const int src_linesize[4],
//enum AVPixelFormat pix_fmt, int width, int height, int align);
//未测试
func AvImageCopyToBuffer(dst *ffcommon.FUint8T, dst_size ffcommon.FInt,
	src_data [4]*ffcommon.FUint8T, src_linesize [4]ffcommon.FInt,
	pix_fmt ffconstant.AVPixelFormat, width ffcommon.FInt, height ffcommon.FInt, align ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_image_copy_to_buffer").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(dst_size),
		uintptr(unsafe.Pointer(&src_data)),
		uintptr(unsafe.Pointer(&src_linesize)),
		uintptr(pix_fmt),
		uintptr(width),
		uintptr(height),
		uintptr(align),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Check if the given dimension of an image is valid, meaning that all
 * bytes of the image can be addressed with a signed int.
 *
 * @param w the width of the picture
 * @param h the height of the picture
 * @param log_offset the offset to sum to the log level for logging with log_ctx
 * @param log_ctx the parent logging context, it may be NULL
 * @return >= 0 if valid, a negative error code otherwise
 */
//int av_image_check_size(unsigned int w, unsigned int h, int log_offset, void *log_ctx);
//未测试
func AvImageCheckSize(w ffcommon.FUnsignedInt, h ffcommon.FUnsignedInt, log_offset ffcommon.FUnsignedInt, log_ctx ffcommon.FVoidP) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_image_check_size").Call(
		uintptr(w),
		uintptr(h),
		uintptr(log_offset),
		uintptr(unsafe.Pointer(log_ctx)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Check if the given dimension of an image is valid, meaning that all
 * bytes of a plane of an image with the specified pix_fmt can be addressed
 * with a signed int.
 *
 * @param w the width of the picture
 * @param h the height of the picture
 * @param max_pixels the maximum number of pixels the user wants to accept
 * @param pix_fmt the pixel format, can be AV_PIX_FMT_NONE if unknown.
 * @param log_offset the offset to sum to the log level for logging with log_ctx
 * @param log_ctx the parent logging context, it may be NULL
 * @return >= 0 if valid, a negative error code otherwise
 */
//int av_image_check_size2(unsigned int w, unsigned int h, int64_t max_pixels, enum AVPixelFormat pix_fmt, int log_offset, void *log_ctx);
//未测试
func AvImageCheckSize2(w ffcommon.FUnsignedInt, h ffcommon.FUnsignedInt, max_pixels ffcommon.FUint64T, pix_fmt ffconstant.AVPixelFormat, log_offset ffcommon.FInt, log_ctx ffcommon.FVoidP) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_image_check_size2").Call(
		uintptr(w),
		uintptr(h),
		uintptr(max_pixels),
		uintptr(pix_fmt),
		uintptr(log_offset),
		uintptr(unsafe.Pointer(log_ctx)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Check if the given sample aspect ratio of an image is valid.
 *
 * It is considered invalid if the denominator is 0 or if applying the ratio
 * to the image size would make the smaller dimension less than 1. If the
 * sar numerator is 0, it is considered unknown and will return as valid.
 *
 * @param w width of the image
 * @param h height of the image
 * @param sar sample aspect ratio of the image
 * @return 0 if valid, a negative AVERROR code otherwise
 */
//int av_image_check_sar(unsigned int w, unsigned int h, AVRational sar);
//未测试
func AvImageCheckSar(w ffcommon.FUnsignedInt, h ffcommon.FUnsignedInt, sar AVRational) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_image_check_sar").Call(
		uintptr(w),
		uintptr(h),
		uintptr(unsafe.Pointer(&sar)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Overwrite the image data with black. This is suitable for filling a
 * sub-rectangle of an image, meaning the padding between the right most pixel
 * and the left most pixel on the next line will not be overwritten. For some
 * formats, the image size might be rounded up due to inherent alignment.
 *
 * If the pixel format has alpha, the alpha is cleared to opaque.
 *
 * This can return an error if the pixel format is not supported. Normally, all
 * non-hwaccel pixel formats should be supported.
 *
 * Passing NULL for dst_data is allowed. Then the function returns whether the
 * operation would have succeeded. (It can return an error if the pix_fmt is
 * not supported.)
 *
 * @param dst_data      data pointers to destination image
 * @param dst_linesize  linesizes for the destination image
 * @param pix_fmt       the pixel format of the image
 * @param range         the color range of the image (important for colorspaces such as YUV)
 * @param width         the width of the image in pixels
 * @param height        the height of the image in pixels
 * @return 0 if the image data was cleared, a negative AVERROR code otherwise
 */
//int av_image_fill_black(uint8_t *dst_data[4], const ptrdiff_t dst_linesize[4],
//enum AVPixelFormat pix_fmt, enum AVColorRange range,
//int width, int height);
//未测试
func AvImageFillBlack(dst_data [4]*ffcommon.FUint8T, dst_linesize [4]ffcommon.FPtrdiffT,
	pix_fmt ffconstant.AVPixelFormat, range0 ffconstant.AVColorRange,
	width ffcommon.FInt, height ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_image_fill_black").Call(
		uintptr(unsafe.Pointer(&dst_data)),
		uintptr(unsafe.Pointer(&dst_linesize)),
		uintptr(pix_fmt),
		uintptr(range0),
		uintptr(pix_fmt),
		uintptr(width),
		uintptr(height),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}
