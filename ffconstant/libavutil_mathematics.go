package ffconstant

const M_E = 2.7182818284590452354        /* e */
const M_LN2 = 0.69314718055994530942     /* log_e 2 */
const M_LN10 = 2.30258509299404568402    /* log_e 10 */
const M_LOG2_10 = 3.32192809488736234787 /* log_2 10 */
const M_PHI = 1.61803398874989484820     /* phi / golden ratio */
const M_PI = 3.14159265358979323846      /* pi */
const M_PI_2 = 1.57079632679489661923    /* pi/2 */
const M_SQRT1_2 = 0.70710678118654752440 /* 1/sqrt(2) */
const M_SQRT2 = 1.41421356237309504880   /* sqrt(2) */
//#ifndef NAN
//#define NAN            av_int2float(0x7fc00000)
//#endif
//#ifndef INFINITY
//#define INFINITY       av_int2float(0x7f800000)
//#endif

/**
 * @addtogroup lavu_math
 *
 * @{
 */

/**
 * Rounding methods.
 */
type AVRounding int32

const (
	AV_ROUND_ZERO     = 0 ///< Round toward zero.
	AV_ROUND_INF      = 1 ///< Round away from zero.
	AV_ROUND_DOWN     = 2 ///< Round toward -infinity.
	AV_ROUND_UP       = 3 ///< Round toward +infinity.
	AV_ROUND_NEAR_INF = 5 ///< Round to nearest and halfway cases away from zero.
	/**
	 * Flag telling rescaling functions to pass `INT64_MIN`/`MAX` through
	 * unchanged, avoiding special cases for #AV_NOPTS_VALUE.
	 *
	 * Unlike other values of the enumeration AVRounding, this value is a
	 * bitmask that must be used in conjunction with another value of the
	 * enumeration through a bitwise OR, in order to set behavior for normal
	 * cases.
	 *
	 * @code{.c}
	 * av_rescale_rnd(3, 1, 2, AV_ROUND_UP | AV_ROUND_PASS_MINMAX);
	 * // Rescaling 3:
	 * //     Calculating 3 * 1 / 2
	 * //     3 / 2 is rounded up to 2
	 * //     => 2
	 *
	 * av_rescale_rnd(AV_NOPTS_VALUE, 1, 2, AV_ROUND_UP | AV_ROUND_PASS_MINMAX);
	 * // Rescaling AV_NOPTS_VALUE:
	 * //     AV_NOPTS_VALUE == INT64_MIN
	 * //     AV_NOPTS_VALUE is passed through
	 * //     => AV_NOPTS_VALUE
	 * @endcode
	 */
	AV_ROUND_PASS_MINMAX = 8192
)
