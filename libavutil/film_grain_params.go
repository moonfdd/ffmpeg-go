package libavutil

import (
	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"unsafe"
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

//#ifndef AVUTIL_FILM_GRAIN_PARAMS_H
//#define AVUTIL_FILM_GRAIN_PARAMS_H
//
//#include "frame.h"

type AVFilmGrainParamsType = int32

const (
	AV_FILM_GRAIN_PARAMS_NONE = iota

	/**
	 * The union is valid when interpreted as AVFilmGrainAOMParams (codec.aom)
	 */
	AV_FILM_GRAIN_PARAMS_AV1
)

/**
 * This structure describes how to handle film grain synthesis for AOM codecs.
 *
 * @note The struct must be allocated as part of AVFilmGrainParams using
 *       av_film_grain_params_alloc(). Its size is not a part of the public ABI.
 */
type AVFilmGrainAOMParams struct {

	/**
	 * Number of points, and the scale and value for each point of the
	 * piecewise linear scaling function for the uma plane.
	 */
	NumYPoints ffcommon.FInt
	YPoints    [14][2] /* value, scaling */ ffcommon.FUint8T

	/**
	 * Signals whether to derive the chroma scaling function from the luma.
	 * Not equivalent to copying the luma values and scales.
	 */
	ChromaScalingFromLuma ffcommon.FInt

	/**
	 * If chroma_scaling_from_luma is set to 0, signals the chroma scaling
	 * function parameters.
	 */
	NumUvPoints [2] /* cb, cr */ ffcommon.FInt
	UvPoints    [2] /* cb, cr */ [10][2] /* value, scaling */ ffcommon.FUint8T

	/**
	 * Specifies the shift applied to the chroma components. For AV1, its within
	 * [8; 11] and determines the range and quantization of the film grain.
	 */
	ScalingShift ffcommon.FInt

	/**
	 * Specifies the auto-regression lag.
	 */
	ArCoeffLag ffcommon.FInt

	/**
	 * Luma auto-regression coefficients. The number of coefficients is given by
	 * 2 * ar_coeff_lag * (ar_coeff_lag + 1).
	 */
	ArCoeffsY [24]ffcommon.FInt8T

	/**
	 * Chroma auto-regression coefficients. The number of coefficients is given by
	 * 2 * ar_coeff_lag * (ar_coeff_lag + 1) + !!num_y_points.
	 */
	ArCoeffsUv [2] /* cb, cr */ [25]ffcommon.FInt8T

	/**
	 * Specifies the range of the auto-regressive coefficients. Values of 6,
	 * 7, 8 and so on represent a range of [-2, 2), [-1, 1), [-0.5, 0.5) and
	 * so on. For AV1 must be between 6 and 9.
	 */
	ArCoeffShift ffcommon.FInt

	/**
	 * Signals the down shift applied to the generated gaussian numbers during
	 * synthesis.
	 */
	GrainScaleShift ffcommon.FInt

	/**
	 * Specifies the luma/chroma multipliers for the index to the component
	 * scaling function.
	 */
	UvMult     [2] /* cb, cr */ ffcommon.FInt
	UvMultLuma [2] /* cb, cr */ ffcommon.FInt

	/**
	 * Offset used for component scaling function. For AV1 its a 9-bit value
	 * with a range [-256, 255]
	 */
	UvOffset [2] /* cb, cr */ ffcommon.FInt

	/**
	 * Signals whether to overlap film grain blocks.
	 */
	OverlapFlag ffcommon.FInt

	/**
	 * Signals to clip to limited color levels after film grain application.
	 */
	LimitOutputRange ffcommon.FInt
}

/**
 * This structure describes how to handle film grain synthesis in video
 * for specific codecs. Must be present on every frame where film grain is
 * meant to be synthesised for correct presentation.
 *
 * @note The struct must be allocated with av_film_grain_params_alloc() and
 *       its size is not a part of the public ABI.
 */
type AVFilmGrainParams struct {

	/**
	 * Specifies the codec for which this structure is valid.
	 */
	Type AVFilmGrainParamsType

	/**
	 * Seed to use for the synthesis process, if the codec allows for it.
	 */
	Seed ffcommon.FUint64T

	/**
	 * Additional fields may be added both here and in any structure included.
	 * If a codec's film grain structure differs slightly over another
	 * codec's, fields within may change meaning depending on the type.
	 */
	//union {
	// aom AVFilmGrainAOMParams
	//} codec;
	AomOrCodec AVFilmGrainAOMParams
}

/**
 * Allocate an AVFilmGrainParams structure and set its fields to
 * default values. The resulting struct can be freed using av_freep().
 * If size is not NULL it will be set to the number of bytes allocated.
 *
 * @return An AVFilmGrainParams filled with default values or NULL
 *         on failure.
 */
//AVFilmGrainParams *av_film_grain_params_alloc(size_t *size);
func AvFilmGrainParamsAlloc(size *ffcommon.FSizeT) (res *AVFilmGrainParams) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_film_grain_params_alloc").Call(
		uintptr(unsafe.Pointer(size)),
	)
	if t == 0 {

	}
	res = (*AVFilmGrainParams)(unsafe.Pointer(t))
	return
}

/**
 * Allocate a complete AVFilmGrainParams and add it to the frame.
 *
 * @param frame The frame which side data is added to.
 *
 * @return The AVFilmGrainParams structure to be filled by caller.
 */
//AVFilmGrainParams *av_film_grain_params_create_side_data(AVFrame *frame);
func (frame *AVFrame) AvFilmGrainParamsCreateSideData() (res *AVFilmGrainParams) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_film_grain_params_create_side_data").Call(
		uintptr(unsafe.Pointer(frame)),
	)
	if t == 0 {

	}
	res = (*AVFilmGrainParams)(unsafe.Pointer(t))
	return
}

//#endif /* AVUTIL_FILM_GRAIN_PARAMS_H */
