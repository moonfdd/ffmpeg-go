package libavutil

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"unsafe"
)

/**
 * Represents the percentile at a specific percentage in
 * a distribution.
 */
type AVHDRPlusPercentile struct {
	/**
	 * The percentage value corresponding to a specific percentile linearized
	 * RGB value in the processing window in the scene. The value shall be in
	 * the range of 0 to100, inclusive.
	 */
	percentage ffcommon.FUint8T

	/**
	 * The linearized maxRGB value at a specific percentile in the processing
	 * window in the scene. The value shall be in the range of 0 to 1, inclusive
	 * and in multiples of 0.00001.
	 */
	percentile AVRational
}

/**
 * Color transform parameters at a processing window in a dynamic metadata for
 * SMPTE 2094-40.
 */
type AVHDRPlusColorTransformParams struct {

	/**
	 * The relative x coordinate of the top left pixel of the processing
	 * window. The value shall be in the range of 0 and 1, inclusive and
	 * in multiples of 1/(width of Picture - 1). The value 1 corresponds
	 * to the absolute coordinate of width of Picture - 1. The value for
	 * first processing window shall be 0.
	 */
	window_upper_left_corner_x AVRational

	/**
	 * The relative y coordinate of the top left pixel of the processing
	 * window. The value shall be in the range of 0 and 1, inclusive and
	 * in multiples of 1/(height of Picture - 1). The value 1 corresponds
	 * to the absolute coordinate of height of Picture - 1. The value for
	 * first processing window shall be 0.
	 */
	window_upper_left_corner_y AVRational

	/**
	 * The relative x coordinate of the bottom right pixel of the processing
	 * window. The value shall be in the range of 0 and 1, inclusive and
	 * in multiples of 1/(width of Picture - 1). The value 1 corresponds
	 * to the absolute coordinate of width of Picture - 1. The value for
	 * first processing window shall be 1.
	 */
	window_lower_right_corner_x AVRational

	/**
	 * The relative y coordinate of the bottom right pixel of the processing
	 * window. The value shall be in the range of 0 and 1, inclusive and
	 * in multiples of 1/(height of Picture - 1). The value 1 corresponds
	 * to the absolute coordinate of height of Picture - 1. The value for
	 * first processing window shall be 1.
	 */
	window_lower_right_corner_y AVRational

	/**
	 * The x coordinate of the center position of the concentric internal and
	 * external ellipses of the elliptical pixel selector in the processing
	 * window. The value shall be in the range of 0 to (width of Picture - 1),
	 * inclusive and in multiples of 1 pixel.
	 */
	center_of_ellipse_x ffcommon.FUint16T

	/**
	 * The y coordinate of the center position of the concentric internal and
	 * external ellipses of the elliptical pixel selector in the processing
	 * window. The value shall be in the range of 0 to (height of Picture - 1),
	 * inclusive and in multiples of 1 pixel.
	 */
	center_of_ellipse_y ffcommon.FUint16T

	/**
	 * The clockwise rotation angle in degree of arc with respect to the
	 * positive direction of the x-axis of the concentric internal and external
	 * ellipses of the elliptical pixel selector in the processing window. The
	 * value shall be in the range of 0 to 180, inclusive and in multiples of 1.
	 */
	rotation_angle ffcommon.FUint16T

	/**
	 * The semi-major axis value of the internal ellipse of the elliptical pixel
	 * selector in amount of pixels in the processing window. The value shall be
	 * in the range of 1 to 65535, inclusive and in multiples of 1 pixel.
	 */
	semimajor_axis_internal_ellipse ffcommon.FUint16T

	/**
	 * The semi-major axis value of the external ellipse of the elliptical pixel
	 * selector in amount of pixels in the processing window. The value
	 * shall not be less than semimajor_axis_internal_ellipse of the current
	 * processing window. The value shall be in the range of 1 to 65535,
	 * inclusive and in multiples of 1 pixel.
	 */
	semimajor_axis_external_ellipse ffcommon.FUint16T

	/**
	 * The semi-minor axis value of the external ellipse of the elliptical pixel
	 * selector in amount of pixels in the processing window. The value shall be
	 * in the range of 1 to 65535, inclusive and in multiples of 1 pixel.
	 */
	semiminor_axis_external_ellipse ffcommon.FUint16T

	/**
	 * Overlap process option indicates one of the two methods of combining
	 * rendered pixels in the processing window in an image with at least one
	 * elliptical pixel selector. For overlapping elliptical pixel selectors
	 * in an image, overlap_process_option shall have the same value.
	 */
	overlap_process_option ffconstant.AVHDRPlusOverlapProcessOption

	/**
	 * The maximum of the color components of linearized RGB values in the
	 * processing window in the scene. The values should be in the range of 0 to
	 * 1, inclusive and in multiples of 0.00001. maxscl[ 0 ], maxscl[ 1 ], and
	 * maxscl[ 2 ] are corresponding to R, G, B color components respectively.
	 */
	maxscl [3]AVRational

	/**
	 * The average of linearized maxRGB values in the processing window in the
	 * scene. The value should be in the range of 0 to 1, inclusive and in
	 * multiples of 0.00001.
	 */
	average_maxrgb AVRational

	/**
	 * The number of linearized maxRGB values at given percentiles in the
	 * processing window in the scene. The maximum value shall be 15.
	 */
	num_distribution_maxrgb_percentiles ffcommon.FUint8T

	/**
	 * The linearized maxRGB values at given percentiles in the
	 * processing window in the scene.
	 */
	distribution_maxrgb [15]AVHDRPlusPercentile

	/**
	 * The fraction of selected pixels in the image that contains the brightest
	 * pixel in the scene. The value shall be in the range of 0 to 1, inclusive
	 * and in multiples of 0.001.
	 */
	fraction_bright_pixels AVRational

	/**
	 * This flag indicates that the metadata for the tone mapping function in
	 * the processing window is present (for value of 1).
	 */
	tone_mapping_flag ffcommon.FUint8T

	/**
	 * The x coordinate of the separation point between the linear part and the
	 * curved part of the tone mapping function. The value shall be in the range
	 * of 0 to 1, excluding 0 and in multiples of 1/4095.
	 */
	knee_point_x AVRational

	/**
	 * The y coordinate of the separation point between the linear part and the
	 * curved part of the tone mapping function. The value shall be in the range
	 * of 0 to 1, excluding 0 and in multiples of 1/4095.
	 */
	knee_point_y AVRational

	/**
	 * The number of the intermediate anchor parameters of the tone mapping
	 * function in the processing window. The maximum value shall be 15.
	 */
	num_bezier_curve_anchors ffcommon.FUint8T

	/**
	 * The intermediate anchor parameters of the tone mapping function in the
	 * processing window in the scene. The values should be in the range of 0
	 * to 1, inclusive and in multiples of 1/1023.
	 */
	bezier_curve_anchors [15]AVRational

	/**
	 * This flag shall be equal to 0 in bitstreams conforming to this version of
	 * this Specification. Other values are reserved for future use.
	 */
	color_saturation_mapping_flag ffcommon.FUint8T

	/**
	 * The color saturation gain in the processing window in the scene. The
	 * value shall be in the range of 0 to 63/8, inclusive and in multiples of
	 * 1/8. The default value shall be 1.
	 */
	color_saturation_weight AVRational
}

/**
* This struct represents dynamic metadata for color volume transform -
* application 4 of SMPTE 2094-40:2016 standard.
*
* To be used as payload of a AVFrameSideData or AVPacketSideData with the
* appropriate type.
*
* @note The struct should be allocated with
* av_dynamic_hdr_plus_alloc() and its size is not a part of
* the public ABI.
 */
type AVDynamicHDRPlus struct {
	/**
	 * Country code by Rec. ITU-T T.35 Annex A. The value shall be 0xB5.
	 */
	itu_t_t35_country_code ffcommon.FUint8T

	/**
	 * Application version in the application defining document in ST-2094
	 * suite. The value shall be set to 0.
	 */
	application_version ffcommon.FUint8T

	/**
	 * The number of processing windows. The value shall be in the range
	 * of 1 to 3, inclusive.
	 */
	num_windows ffcommon.FUint8T

	/**
	 * The color transform parameters for every processing window.
	 */
	params [3]AVHDRPlusColorTransformParams

	/**
	 * The nominal maximum display luminance of the targeted system display,
	 * in units of 0.0001 candelas per square metre. The value shall be in
	 * the range of 0 to 10000, inclusive.
	 */
	targeted_system_display_maximum_luminance AVRational

	/**
	 * This flag shall be equal to 0 in bit streams conforming to this version
	 * of this Specification. The value 1 is reserved for future use.
	 */
	targeted_system_display_actual_peak_luminance_flag ffcommon.FUint8T

	/**
	 * The number of rows in the targeted system_display_actual_peak_luminance
	 * array. The value shall be in the range of 2 to 25, inclusive.
	 */
	num_rows_targeted_system_display_actual_peak_luminance ffcommon.FUint8T

	/**
	 * The number of columns in the
	 * targeted_system_display_actual_peak_luminance array. The value shall be
	 * in the range of 2 to 25, inclusive.
	 */
	num_cols_targeted_system_display_actual_peak_luminance ffcommon.FUint8T

	/**
	 * The normalized actual peak luminance of the targeted system display. The
	 * values should be in the range of 0 to 1, inclusive and in multiples of
	 * 1/15.
	 */
	targeted_system_display_actual_peak_luminance [25][25]AVRational

	/**
	 * This flag shall be equal to 0 in bitstreams conforming to this version of
	 * this Specification. The value 1 is reserved for future use.
	 */
	mastering_display_actual_peak_luminance_flag ffcommon.FUint8T

	/**
	 * The number of rows in the mastering_display_actual_peak_luminance array.
	 * The value shall be in the range of 2 to 25, inclusive.
	 */
	num_rows_mastering_display_actual_peak_luminance ffcommon.FUint8T

	/**
	 * The number of columns in the mastering_display_actual_peak_luminance
	 * array. The value shall be in the range of 2 to 25, inclusive.
	 */
	num_cols_mastering_display_actual_peak_luminance ffcommon.FUint8T

	/**
	 * The normalized actual peak luminance of the mastering display used for
	 * mastering the image essence. The values should be in the range of 0 to 1,
	 * inclusive and in multiples of 1/15.
	 */
	mastering_display_actual_peak_luminance [25][25]AVRational
}

/**
* Allocate an AVDynamicHDRPlus structure and set its fields to
* default values. The resulting struct can be freed using av_freep().
*
* @return An AVDynamicHDRPlus filled with default values or NULL
*         on failure.
 */
//AVDynamicHDRPlus *av_dynamic_hdr_plus_alloc(size_t *size);
//未测试
func AvDynamicHdrPlusAlloc(size *ffcommon.FSizeT) (res *AVDynamicHDRPlus, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_dynamic_hdr_plus_alloc").Call(
		uintptr(unsafe.Pointer(size)),
	)
	if err != nil {
		//return
	}
	res = (*AVDynamicHDRPlus)(unsafe.Pointer(t))
	return
}

/**
* Allocate a complete AVDynamicHDRPlus and add it to the frame.
* @param frame The frame which side data is added to.
*
* @return The AVDynamicHDRPlus structure to be filled by caller or NULL
*         on failure.
 */
//AVDynamicHDRPlus *av_dynamic_hdr_plus_create_side_data(AVFrame *frame);
//未测试
func (frame *AVFrame) AvDynamicHdrPlusCreateSideData() (res *AVDynamicHDRPlus, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_dynamic_hdr_plus_create_side_data").Call(
		uintptr(unsafe.Pointer(frame)),
	)
	if err != nil {
		//return
	}
	res = (*AVDynamicHDRPlus)(unsafe.Pointer(t))
	return
}
