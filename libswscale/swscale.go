package libswscale

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
)

/**
 * @defgroup libsws libswscale
 * Color conversion and scaling library.
 *
 * @{
 *
 * Return the LIBSWSCALE_VERSION_INT ffconstant.
 */
//unsigned swscale_version(void);
func SwscaleVersion() (res ffcommon.FUnsigned) {
	return
}

/**
 * Return the libswscale build-time configuration.
 */
//const char *swscale_configuration(void);
func SwscaleConfiguration() (res ffcommon.FConstCharP) {
	return
}

/**
 * Return the libswscale license.
 */
//const char *swscale_license(void);
func SwscaleLicense() (res ffcommon.FConstCharP) {
	return
}

/**
 * Return a pointer to yuv<->rgb coefficients for the given colorspace
 * suitable for sws_setColorspaceDetails().
 *
 * @param colorspace One of the SWS_CS_* macros. If invalid,
 * SWS_CS_DEFAULT is used.
 */
//const int *sws_getCoefficients(int colorspace);
func SwsGetCoefficients(colorspace ffcommon.FInt) (res ffcommon.FConstIntP) {
	return
}

// when used for filters they must have an odd number of elements
// coeffs cannot be shared between vectors
type SwsVector struct {
}

// vectors can be shared
type SwsFilter struct {
}

type SwsContext struct {
}

/**
 * Return a positive value if pix_fmt is a supported input format, 0
 * otherwise.
 */
//int sws_isSupportedInput(enum AVPixelFormat pix_fmt);
func SwsIsSupportedInput(pix_fmt ffconstant.AVPixelFormat) (res ffcommon.FInt) {
	return
}

/**
 * Return a positive value if pix_fmt is a supported output format, 0
 * otherwise.
 */
//int sws_isSupportedOutput(enum AVPixelFormat pix_fmt);
func SwsIsSupportedOutput(pix_fmt ffconstant.AVPixelFormat) (res ffcommon.FInt) {
	return
}

/**
 * @param[in]  pix_fmt the pixel format
 * @return a positive value if an endianness conversion for pix_fmt is
 * supported, 0 otherwise.
 */
//int sws_isSupportedEndiannessConversion(enum AVPixelFormat pix_fmt);
func SwsIsSupportedEndiannessConversion(pix_fmt ffconstant.AVPixelFormat) (res ffcommon.FInt) {
	return
}

/**
 * Allocate an empty SwsContext. This must be filled and passed to
 * sws_init_context(). For filling see AVOptions, options.c and
 * sws_setColorspaceDetails().
 */
//struct SwsContext *sws_alloc_context(void);

func SwsAllocContext() (res *SwsContext) {
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
func (sws_context *SwsContext) SwsInitContext(srcFilter *SwsFilter, dstFilter *SwsFilter) (res ffcommon.FInt) {
	return
}

/**
 * Free the swscaler context swsContext.
 * If swsContext is NULL, then does nothing.
 */
//void sws_freeContext(struct SwsContext *swsContext);
func (swsContext *SwsContext) SwsFreeContext() {
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
func SwsGetContext(srcW ffcommon.FInt, srcH ffcommon.FInt, srcFormat ffconstant.AVPixelFormat,
	dstW ffcommon.FInt, dstH ffcommon.FInt, dstFormat ffconstant.AVPixelFormat,
	flags ffcommon.FInt, srcFilter *SwsFilter,
	dstFilter *SwsFilter, param ffcommon.FConstDoubleP) (ans *SwsContext) {
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
func (c *SwsContext) SwsScale(srcSlice []ffcommon.FUint8T,
	srcStride []ffcommon.FInt, srcSliceY ffcommon.FInt, srcSliceH ffcommon.FInt,
	dst []ffcommon.FUint8T, dstStride []ffcommon.FConstInt) (res ffcommon.FInt) {
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
func (c *SwsContext) SwsSetColorspaceDetails(inv_table [4]ffcommon.FInt,
	srcRange ffcommon.FInt, table [4]ffcommon.FInt, dstRange ffcommon.FInt,
	brightness ffcommon.FInt, contrast ffcommon.FInt, saturation ffcommon.FInt) (res ffcommon.FInt) {
	return
}

/**
 * @return -1 if not supported
 */
//int sws_getColorspaceDetails(struct SwsContext *c, int **inv_table,
//int *srcRange, int **table, int *dstRange,
//int *brightness, int *contrast, int *saturation);
func (c *SwsContext) SwsGetColorspaceDetails(inv_table **ffcommon.FInt,
	srcRange *ffcommon.FInt, table **ffcommon.FInt, dstRange *ffcommon.FInt,
	brightness *ffcommon.FInt, ontrast *ffcommon.FInt, saturation *ffcommon.FInt) (res ffcommon.FInt) {
	return
}

/**
 * Allocate and return an uninitialized vector with length coefficients.
 */
//SwsVector *sws_allocVec(int length);
func SwsAllocVec(length ffcommon.FInt) (res *SwsVector) {
	return
}

/**
 * Return a normalized Gaussian curve used to filter stuff
 * quality = 3 is high quality, lower is lower quality.
 */
//SwsVector *sws_getGaussianVec(double variance, double quality);
func SwsGetGaussianVec(variance ffcommon.FConstDouble, quality ffcommon.FConstDouble) (res *SwsVector) {
	return
}

/**
 * Scale all the coefficients of a by the scalar value.
 */
//void sws_scaleVec(SwsVector *a, double scalar);
func (a *SwsVector) SwsScaleVec(scalar ffcommon.FConstDouble) {
	return
}

/**
 * Scale all the coefficients of a so that their sum equals height.
 */
//void sws_normalizeVec(SwsVector *a, double height);
func (a *SwsVector) SwsNormalizeVec(scalar ffcommon.FConstDouble) {
	return
}

//#if FF_API_SWS_VECTOR
//attribute_deprecated SwsVector *sws_getConstVec(double c, int length);
func SwsGetConstVec(c ffcommon.FDouble, length ffcommon.FInt) (res *SwsVector) {
	return
}

//attribute_deprecated SwsVector *sws_getIdentityVec(void);
func SwsGetIdentityVec() (res *SwsVector) {
	return
}

//attribute_deprecated void sws_convVec(SwsVector *a, SwsVector *b);
func SwsConvVec(a *SwsVector, b *SwsVector) {
	return
}

//attribute_deprecated void sws_addVec(SwsVector *a, SwsVector *b);
func SwsAddVec(a *SwsVector, b *SwsVector) {
	return
}

//attribute_deprecated void sws_subVec(SwsVector *a, SwsVector *b);
func SwsSubVec(a *SwsVector, b *SwsVector) {
	return
}

//attribute_deprecated void sws_shiftVec(SwsVector *a, int shift);
func SwsShiftVec(a *SwsVector, shift ffcommon.FInt) {
	return
}

//attribute_deprecated SwsVector *sws_cloneVec(SwsVector *a);
func (a *SwsVector) SwsCloneVec() (ans *SwsVector) {
	return
}

//attribute_deprecated void sws_printVec2(SwsVector *a, AVClass *log_ctx, int log_level);
func (a *SwsVector) SwsPrintVec2(log_ctx *AVClass, log_level ffcommon.FInt) {
	return
}

//#endif

//void sws_freeVec(SwsVector *a);
func (a *SwsVector) SwsFreeVec() {
	return
}

//SwsFilter *sws_getDefaultFilter(float lumaGBlur, float chromaGBlur,
//float lumaSharpen, float chromaSharpen,
//float chromaHShift, float chromaVShift,
//int verbose);
func SwsGetDefaultFilter(lumaGBlur ffcommon.FFloat, chromaGBlur ffcommon.FFloat,
	lumaSharpen ffcommon.FFloat, chromaSharpen ffcommon.FFloat,
	chromaHShift ffcommon.FFloat, chromaVShift ffcommon.FFloat,
	verbose ffcommon.FInt) (res *SwsFilter) {
	return
}

//void sws_freeFilter(SwsFilter *filter);
func (filter *SwsFilter) SwsFreeFilter() {
	return
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

func (context *SwsContext) SwsGetCachedContext(srcW ffcommon.FInt, srcH ffcommon.FInt, srcFormat ffconstant.AVPixelFormat,
	dstW ffcommon.FInt, dstH ffcommon.FInt, dstFormat ffconstant.AVPixelFormat,
	flags ffcommon.FInt, srcFilter *SwsFilter,
	dstFilter *SwsFilter, param ffcommon.FConstDoubleP) (res *SwsContext) {
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
func SwsConvertPalette8ToPacked32(src *ffcommon.FUint8T, dst *ffcommon.FUint8T, num_pixels ffcommon.FUint8T, palette *ffcommon.FUint8T) {
	return
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
func SwsConvertPalette8ToPacked24(src *ffcommon.FUint8T, dst *ffcommon.FUint8T, num_pixels ffcommon.FInt, palette *ffcommon.FUint8T) {
	return
}

/**
 * Get the AVClass for swsContext. It can be used in combination with
 * AV_OPT_SEARCH_FAKE_OBJ for examining options.
 *
 * @see av_opt_find().
 */
//const AVClass *sws_get_class(void);
func SwsGetClass() (res *AVClass) {
	return
}
