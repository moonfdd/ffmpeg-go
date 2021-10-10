package libavutil

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"unsafe"
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
	num_y_points ffcommon.FInt
	y_points     [14][2] /* value, scaling */ ffcommon.FUint8T

	/**
	 * Signals whether to derive the chroma scaling function from the luma.
	 * Not equivalent to copying the luma values and scales.
	 */
	chroma_scaling_from_luma ffcommon.FInt

	/**
	 * If chroma_scaling_from_luma is set to 0, signals the chroma scaling
	 * function parameters.
	 */
	num_uv_points [2] /* cb, cr */ ffcommon.FInt
	uv_points     [2] /* cb, cr */ [10][2] /* value, scaling */ ffcommon.FUint8T

	/**
	 * Specifies the shift applied to the chroma components. For AV1, its within
	 * [8; 11] and determines the range and quantization of the film grain.
	 */
	scaling_shift ffcommon.FInt

	/**
	 * Specifies the auto-regression lag.
	 */
	ar_coeff_lag ffcommon.FInt

	/**
	 * Luma auto-regression coefficients. The number of coefficients is given by
	 * 2 * ar_coeff_lag * (ar_coeff_lag + 1).
	 */
	ar_coeffs_y [24]ffcommon.FUint8T

	/**
	 * Chroma auto-regression coefficients. The number of coefficients is given by
	 * 2 * ar_coeff_lag * (ar_coeff_lag + 1) + !!num_y_points.
	 */
	ar_coeffs_uv [2] /* cb, cr */ [25]ffcommon.FUint8T

	/**
	 * Specifies the range of the auto-regressive coefficients. Values of 6,
	 * 7, 8 and so on represent a range of [-2, 2), [-1, 1), [-0.5, 0.5) and
	 * so on. For AV1 must be between 6 and 9.
	 */
	ar_coeff_shift ffcommon.FInt

	/**
	 * Signals the down shift applied to the generated gaussian numbers during
	 * synthesis.
	 */
	grain_scale_shift ffcommon.FInt

	/**
	 * Specifies the luma/chroma multipliers for the index to the component
	 * scaling function.
	 */
	uv_mult      [2] /* cb, cr */ ffcommon.FInt
	uv_mult_luma [2] /* cb, cr */ ffcommon.FInt

	/**
	 * Offset used for component scaling function. For AV1 its a 9-bit value
	 * with a range [-256, 255]
	 */
	uv_offset [2] /* cb, cr */ ffcommon.FInt

	/**
	 * Signals whether to overlap film grain blocks.
	 */
	overlap_flag ffcommon.FInt

	/**
	 * Signals to clip to limited color levels after film grain application.
	 */
	limit_output_range ffcommon.FInt
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
	type0 ffconstant.AVFilmGrainParamsType

	/**
	 * Seed to use for the synthesis process, if the codec allows for it.
	 */
	seed ffcommon.FUint64T

	/**
	 * Additional fields may be added both here and in any structure included.
	 * If a codec's film grain structure differs slightly over another
	 * codec's, fields within may change meaning depending on the type.
	 */
	//union {
	//AVFilmGrainAOMParams aom;
	//} codec;
	codec AVFilmGrainAOMParams
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
//未测试
func AvFilmGrainParamsAlloc(size *ffcommon.FSizeT) (res *AVFilmGrainParams, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_film_grain_params_alloc").Call(
		uintptr(unsafe.Pointer(size)),
	)
	if err != nil {
		//return
	}
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
//未测试
func (frame *AVFrame) AvFilmGrainParamsCreateSideData() (res *AVFilmGrainParams, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_film_grain_params_create_side_data").Call(
		uintptr(unsafe.Pointer(frame)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = (*AVFilmGrainParams)(unsafe.Pointer(t))
	return
}
