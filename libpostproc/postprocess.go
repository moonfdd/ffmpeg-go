package libpostproc

import "ffmpeg-go/ffcommon"

/**
 * Return the LIBPOSTPROC_VERSION_INT ffconstant.
 */
//unsigned postproc_version(void);
func PostprocVersion() (res ffcommon.FUnsigned) {
	return
}

/**
 * Return the libpostproc build-time configuration.
 */
//const char *postproc_configuration(void);
func PostprocConfiguration() (res ffcommon.FConstCharP) {
	return
}

/**
 * Return the libpostproc license.
 */
//const char *postproc_license(void);
func PostprocLicense() (res ffcommon.FConstCharP) {
	return
}

//void  pp_postprocess(const uint8_t * src[3], const int srcStride[3],
//uint8_t * dst[3], const int dstStride[3],
//int horizontalSize, int verticalSize,
//const int8_t *QP_store,  int QP_stride,
//pp_mode *mode, pp_context *ppContext, int pict_type);
func PpPostprocess(src [3]*ffcommon.FUint8T, srcStride [3]ffcommon.FInt,
	dst [3]*ffcommon.FUint8T, dstStride [3]ffcommon.FInt,
	horizontalSize ffcommon.FInt, verticalSize ffcommon.FInt,
	QP_store *ffcommon.FUint8T, QP_stride ffcommon.FInt,
	mode ffcommon.FVoidP, ppContext ffcommon.FVoidP, pict_type ffcommon.FInt) {
	return
}

/**
 * Return a pp_mode or NULL if an error occurred.
 *
 * @param name    the string after "-pp" on the command line
 * @param quality a number from 0 to PP_QUALITY_MAX
 */
//pp_mode *pp_get_mode_by_name_and_quality(const char *name, int quality);
func PpGetModeByNameAndQuality(name ffcommon.FConstCharP, quality ffcommon.FInt) (res ffcommon.FVoidP) {
	return
}

//void pp_free_mode(pp_mode *mode);
func PpFreeMode(mode ffcommon.FVoidP) {
	return
}

//pp_context *pp_get_context(int width, int height, int flags);
func pp_get_context(width ffcommon.FInt, height ffcommon.FInt, flags ffcommon.FInt) (res ffcommon.FVoidP) {
	return
}

//void pp_free_context(pp_context *ppContext);
func pp_free_context(ppContext ffcommon.FVoidP) {
	return
}
