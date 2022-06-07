package libavutil

import (
	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"unsafe"
)

/*
 * Copyright (c) 2018 Mohammad Izadi <moh.izadi at gmail.com>
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

//#ifndef AVUTIL_HDR_DYNAMIC_METADATA_H
//#define AVUTIL_HDR_DYNAMIC_METADATA_H
//
//#include "frame.h"
//#include "rational.h"

/**
 * Option for overlapping elliptical pixel selectors in an image.
 */
type AVHDRPlusOverlapProcessOption = int32

const (
	AV_HDR_PLUS_OVERLAP_PROCESS_WEIGHTED_AVERAGING = 0
	AV_HDR_PLUS_OVERLAP_PROCESS_LAYERING           = 1
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
	Percentage ffcommon.FUint8T

	/**
	 * The linearized maxRGB value at a specific percentile in the processing
	 * window in the scene. The value shall be in the range of 0 to 1, inclusive
	 * and in multiples of 0.00001.
	 */
	Percentile AVRational
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
	WindowUpperLeftCornerX AVRational

	/**
	 * The relative y coordinate of the top left pixel of the processing
	 * window. The value shall be in the range of 0 and 1, inclusive and
	 * in multiples of 1/(height of Picture - 1). The value 1 corresponds
	 * to the absolute coordinate of height of Picture - 1. The value for
	 * first processing window shall be 0.
	 */
	WindowUpperLeftCornerY AVRational

	/**
	 * The relative x coordinate of the bottom right pixel of the processing
	 * window. The value shall be in the range of 0 and 1, inclusive and
	 * in multiples of 1/(width of Picture - 1). The value 1 corresponds
	 * to the absolute coordinate of width of Picture - 1. The value for
	 * first processing window shall be 1.
	 */
	WindowLowerRightCornerX AVRational

	/**
	 * The relative y coordinate of the bottom right pixel of the processing
	 * window. The value shall be in the range of 0 and 1, inclusive and
	 * in multiples of 1/(height of Picture - 1). The value 1 corresponds
	 * to the absolute coordinate of height of Picture - 1. The value for
	 * first processing window shall be 1.
	 */
	WindowLowerRightCornerY AVRational

	/**
	 * The x coordinate of the center position of the concentric internal and
	 * external ellipses of the elliptical pixel selector in the processing
	 * window. The value shall be in the range of 0 to (width of Picture - 1),
	 * inclusive and in multiples of 1 pixel.
	 */
	CenterOfEllipseX ffcommon.FUint16T

	/**
	 * The y coordinate of the center position of the concentric internal and
	 * external ellipses of the elliptical pixel selector in the processing
	 * window. The value shall be in the range of 0 to (height of Picture - 1),
	 * inclusive and in multiples of 1 pixel.
	 */
	CenterOfEllipseY ffcommon.FUint16T

	/**
	 * The clockwise rotation angle in degree of arc with respect to the
	 * positive direction of the x-axis of the concentric internal and external
	 * ellipses of the elliptical pixel selector in the processing window. The
	 * value shall be in the range of 0 to 180, inclusive and in multiples of 1.
	 */
	RotationAngle ffcommon.FUint8T

	/**
	 * The semi-major axis value of the internal ellipse of the elliptical pixel
	 * selector in amount of pixels in the processing window. The value shall be
	 * in the range of 1 to 65535, inclusive and in multiples of 1 pixel.
	 */
	SemimajorAxisInternalEllipse ffcommon.FUint16T

	/**
	 * The semi-major axis value of the external ellipse of the elliptical pixel
	 * selector in amount of pixels in the processing window. The value
	 * shall not be less than semimajor_axis_internal_ellipse of the current
	 * processing window. The value shall be in the range of 1 to 65535,
	 * inclusive and in multiples of 1 pixel.
	 */
	SemimajorAxisExternalEllipse ffcommon.FUint16T

	/**
	 * The semi-minor axis value of the external ellipse of the elliptical pixel
	 * selector in amount of pixels in the processing window. The value shall be
	 * in the range of 1 to 65535, inclusive and in multiples of 1 pixel.
	 */
	SemiminorAxisExternalEllipse ffcommon.FUint16T

	/**
	 * Overlap process option indicates one of the two methods of combining
	 * rendered pixels in the processing window in an image with at least one
	 * elliptical pixel selector. For overlapping elliptical pixel selectors
	 * in an image, overlap_process_option shall have the same value.
	 */
	OverlapProcessOption AVHDRPlusOverlapProcessOption

	/**
	 * The maximum of the color components of linearized RGB values in the
	 * processing window in the scene. The values should be in the range of 0 to
	 * 1, inclusive and in multiples of 0.00001. maxscl[ 0 ], maxscl[ 1 ], and
	 * maxscl[ 2 ] are corresponding to R, G, B color components respectively.
	 */
	Maxscl [3]AVRational

	/**
	 * The average of linearized maxRGB values in the processing window in the
	 * scene. The value should be in the range of 0 to 1, inclusive and in
	 * multiples of 0.00001.
	 */
	AverageMaxrgb AVRational

	/**
	 * The number of linearized maxRGB values at given percentiles in the
	 * processing window in the scene. The maximum value shall be 15.
	 */
	NumDistributionMaxrgbPercentiles ffcommon.FUint8T

	/**
	 * The linearized maxRGB values at given percentiles in the
	 * processing window in the scene.
	 */
	DistributionMaxrgb [15]AVHDRPlusPercentile

	/**
	 * The fraction of selected pixels in the image that contains the brightest
	 * pixel in the scene. The value shall be in the range of 0 to 1, inclusive
	 * and in multiples of 0.001.
	 */
	FractionBrightPixels AVRational

	/**
	 * This flag indicates that the metadata for the tone mapping function in
	 * the processing window is present (for value of 1).
	 */
	ToneMappingFlag ffcommon.FUint8T

	/**
	 * The x coordinate of the separation point between the linear part and the
	 * curved part of the tone mapping function. The value shall be in the range
	 * of 0 to 1, excluding 0 and in multiples of 1/4095.
	 */
	KneePointX AVRational

	/**
	 * The y coordinate of the separation point between the linear part and the
	 * curved part of the tone mapping function. The value shall be in the range
	 * of 0 to 1, excluding 0 and in multiples of 1/4095.
	 */
	KneePointY AVRational

	/**
	 * The number of the intermediate anchor parameters of the tone mapping
	 * function in the processing window. The maximum value shall be 15.
	 */
	NumBezierCurveAnchors ffcommon.FUint8T

	/**
	 * The intermediate anchor parameters of the tone mapping function in the
	 * processing window in the scene. The values should be in the range of 0
	 * to 1, inclusive and in multiples of 1/1023.
	 */
	BezierCurveAnchors [15]AVRational

	/**
	 * This flag shall be equal to 0 in bitstreams conforming to this version of
	 * this Specification. Other values are reserved for future use.
	 */
	ColorSaturationMappingFlag ffcommon.FUint8T

	/**
	 * The color saturation gain in the processing window in the scene. The
	 * value shall be in the range of 0 to 63/8, inclusive and in multiples of
	 * 1/8. The default value shall be 1.
	 */
	ColorSaturationWeight AVRational
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
	ItuTT35CountryCode ffcommon.FUint8T

	/**
	 * Application version in the application defining document in ST-2094
	 * suite. The value shall be set to 0.
	 */
	ApplicationVersion ffcommon.FUint8T

	/**
	 * The number of processing windows. The value shall be in the range
	 * of 1 to 3, inclusive.
	 */
	NumWindows ffcommon.FUint8T

	/**
	 * The color transform parameters for every processing window.
	 */
	Params [3]AVHDRPlusColorTransformParams

	/**
	 * The nominal maximum display luminance of the targeted system display,
	 * in units of 0.0001 candelas per square metre. The value shall be in
	 * the range of 0 to 10000, inclusive.
	 */
	TargetedSystemDisplayMaximumLuminance AVRational

	/**
	 * This flag shall be equal to 0 in bit streams conforming to this version
	 * of this Specification. The value 1 is reserved for future use.
	 */
	TargetedSystemDisplayActualPeakLuminanceFlag ffcommon.FUint8T

	/**
	 * The number of rows in the targeted system_display_actual_peak_luminance
	 * array. The value shall be in the range of 2 to 25, inclusive.
	 */
	NumRowsTargetedSystemDisplayActualPeakLuminance ffcommon.FUint8T

	/**
	 * The number of columns in the
	 * targeted_system_display_actual_peak_luminance array. The value shall be
	 * in the range of 2 to 25, inclusive.
	 */
	NumColsTargetedSystemDisplayActualPeakLuminance ffcommon.FUint8T

	/**
	 * The normalized actual peak luminance of the targeted system display. The
	 * values should be in the range of 0 to 1, inclusive and in multiples of
	 * 1/15.
	 */
	TargetedSystemDisplayActualPeakLuminance [25][25]AVRational

	/**
	 * This flag shall be equal to 0 in bitstreams conforming to this version of
	 * this Specification. The value 1 is reserved for future use.
	 */
	MasteringDisplayActualPeakLuminanceFlag ffcommon.FUint8T

	/**
	 * The number of rows in the mastering_display_actual_peak_luminance array.
	 * The value shall be in the range of 2 to 25, inclusive.
	 */
	NumRowsMasteringDisplayActualPeakLuminance ffcommon.FUint8T

	/**
	 * The number of columns in the mastering_display_actual_peak_luminance
	 * array. The value shall be in the range of 2 to 25, inclusive.
	 */
	NumColsMasteringDisplayActualPeakLuminance ffcommon.FUint8T

	/**
	 * The normalized actual peak luminance of the mastering display used for
	 * mastering the image essence. The values should be in the range of 0 to 1,
	 * inclusive and in multiples of 1/15.
	 */
	MasteringDisplayActualPeakLuminance [25][25]AVRational
}

/**
 * Allocate an AVDynamicHDRPlus structure and set its fields to
 * default values. The resulting struct can be freed using av_freep().
 *
 * @return An AVDynamicHDRPlus filled with default values or NULL
 *         on failure.
 */
//AVDynamicHDRPlus *av_dynamic_hdr_plus_alloc(size_t *size);
//todo
func av_dynamic_hdr_plus_alloc() (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_dynamic_hdr_plus_alloc").Call()
	if t == 0 {

	}
	res = ffcommon.StringFromPtr(t)
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
func (frame *AVFrame) AvDynamicHdrPlusCreateSideData() (res *AVDynamicHDRPlus) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_dynamic_hdr_plus_create_side_data").Call(
		uintptr(unsafe.Pointer(frame)),
	)
	if t == 0 {

	}
	res = (*AVDynamicHDRPlus)(unsafe.Pointer(t))
	return
}

//#endif /* AVUTIL_HDR_DYNAMIC_METADATA_H */
