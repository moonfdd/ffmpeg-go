package ffconstant

type AVTXType int32

const (
	/**
	 * Standard complex to complex FFT with sample data type AVComplexFloat.
	 * Output is not 1/len normalized. Scaling currently unsupported.
	 * The stride parameter is ignored.
	 */
	AV_TX_FLOAT_FFT = 0

	/**
	 * Standard MDCT with sample data type of float and a scale type of
	 * float. Length is the frame size, not the window size (which is 2x frame)
	 * For forward transforms, the stride specifies the spacing between each
	 * sample in the output array in bytes. The input must be a flat array.
	 * For inverse transforms, the stride specifies the spacing between each
	 * sample in the input array in bytes. The output will be a flat array.
	 * Stride must be a non-zero multiple of sizeof(float).
	 * NOTE: the inverse transform is half-length, meaning the output will not
	 * contain redundant data. This is what most codecs work with.
	 */
	AV_TX_FLOAT_MDCT = 1

	/**
	 * Same as AV_TX_FLOAT_FFT with a data type of AVComplexDouble.
	 */
	AV_TX_DOUBLE_FFT = 2

	/**
	 * Same as AV_TX_FLOAT_MDCT with data and scale type of double.
	 * Stride must be a non-zero multiple of sizeof(double).
	 */
	AV_TX_DOUBLE_MDCT = 3

	/**
	 * Same as AV_TX_FLOAT_FFT with a data type of AVComplexInt32.
	 */
	AV_TX_INT32_FFT = 4

	/**
	 * Same as AV_TX_FLOAT_MDCT with data type of int32_t and scale type of float.
	 * Only scale values less than or equal to 1.0 are supported.
	 * Stride must be a non-zero multiple of sizeof(int32_t).
	 */
	AV_TX_INT32_MDCT = 5
)

/**
 * Flags for av_tx_init()
 */
type AVTXFlags uint64

const (
	/**
	 * Performs an in-place transformation on the input. The output argument
	 * of av_tn_fn() MUST match the input. May be unsupported or slower for some
	 * transform types.
	 */
	AV_TX_INPLACE = 1 << 0
)
