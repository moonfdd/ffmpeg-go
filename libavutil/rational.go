package libavutil

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

/**
 * @defgroup lavu_math_rational AVRational
 * @ingroup lavu_math
 * Rational number calculation.
 *
 * While rational numbers can be expressed as floating-point numbers, the
 * conversion process is a lossy one, so are floating-point operations. On the
 * other hand, the nature of FFmpeg demands highly accurate calculation of
 * timestamps. This set of rational number utilities serves as a generic
 * interface for manipulating rational numbers as pairs of numerators and
 * denominators.
 *
 * Many of the functions that operate on AVRational's have the suffix `_q`, in
 * reference to the mathematical symbol "ℚ" (Q) which denotes the set of all
 * rational numbers.
 *
 * @{
 */

/**
 * Rational number (pair of numerator and denominator).
 */
type AVRational struct {
	Num ffcommon.FInt ///< Numerator
	Den ffcommon.FInt ///< Denominator
}

/**
 * Reduce a fraction.
 *
 * This is useful for framerate calculations.
 *
 * @param[out] dst_num Destination numerator
 * @param[out] dst_den Destination denominator
 * @param[in]      num Source numerator
 * @param[in]      den Source denominator
 * @param[in]      max Maximum allowed values for `dst_num` & `dst_den`
 * @return 1 if the operation is exact, 0 otherwise
 */
//int av_reduce(int *dst_num, int *dst_den, int64_t num, int64_t den, int64_t max);
//未测试
func AvReduce(dst_num *ffcommon.FInt, dst_den *ffcommon.FInt, num ffcommon.FUint64T, den ffcommon.FUint64T, max ffcommon.FUint64T) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_reduce").Call(
		uintptr(unsafe.Pointer(dst_num)),
		uintptr(unsafe.Pointer(dst_den)),
		uintptr(num),
		uintptr(den),
		uintptr(max),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Multiply two rationals.
 * @param b First rational
 * @param c Second rational
 * @return b*c
 */
//AVRational av_mul_q(AVRational b, AVRational c) av_const;
//未测试
func AvMulQ(a AVRational, b AVRational) (res AVRational, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_mul_q").Call(
		uintptr(unsafe.Pointer(&a)),
		uintptr(unsafe.Pointer(&b)),
	)
	if err != nil {
		//return
	}
	res = *(*AVRational)(unsafe.Pointer(t))
	return
}

/**
* Divide one rational by another.
* @param b First rational
* @param c Second rational
* @return b/c
 */
//AVRational av_div_q(AVRational b, AVRational c) av_const;
//未测试
func AvDivQ(a AVRational, b AVRational) (res AVRational, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_div_q").Call(
		uintptr(unsafe.Pointer(&a)),
		uintptr(unsafe.Pointer(&b)),
	)
	if err != nil {
		//return
	}
	res = *(*AVRational)(unsafe.Pointer(t))
	return
}

/**
* Add two rationals.
* @param b First rational
* @param c Second rational
* @return b+c
 */
//AVRational av_add_q(AVRational b, AVRational c) av_const;
//未测试
func AvAddQ(a AVRational, b AVRational) (res AVRational, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_add_q").Call(
		uintptr(unsafe.Pointer(&a)),
		uintptr(unsafe.Pointer(&b)),
	)
	if err != nil {
		//return
	}
	res = *(*AVRational)(unsafe.Pointer(t))
	return
}

/**
* Subtract one rational from another.
* @param b First rational
* @param c Second rational
* @return b-c
 */
//AVRational av_sub_q(AVRational b, AVRational c) av_const;
//未测试
func AvSubQ(a AVRational, b AVRational) (res AVRational, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_sub_q").Call(
		uintptr(unsafe.Pointer(&a)),
		uintptr(unsafe.Pointer(&b)),
	)
	if err != nil {
		//return
	}
	res = *(*AVRational)(unsafe.Pointer(t))
	return
}

/**
* Convert a double precision floating point number to a rational.
*
* In case of infinity, the returned value is expressed as `{1, 0}` or
* `{-1, 0}` depending on the sign.
*
* @param d   `double` to convert
* @param max Maximum allowed numerator and denominator
* @return `d` in AVRational form
* @see av_q2d()
 */
//AVRational av_d2q(double d, int max) av_const;
//未测试
func AvD2q(a ffcommon.FDouble, b ffcommon.FDouble) (res AVRational, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_d2q").Call(
		uintptr(a),
		uintptr(b),
	)
	if err != nil {
		//return
	}
	res = *(*AVRational)(unsafe.Pointer(t))
	return
}

/**
* Find which of the two rationals is closer to another rational.
*
* @param q     Rational to be compared against
* @param q1,q2 Rationals to be tested
* @return One of the following values:
*         - 1 if `q1` is nearer to `q` than `q2`
*         - -1 if `q2` is nearer to `q` than `q1`
*         - 0 if they have the same distance
 */
//int av_nearer_q(AVRational q, AVRational q1, AVRational q2);
//未测试
func AvNearerQ(q AVRational, q1 AVRational, q2 AVRational) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_nearer_q").Call(
		uintptr(unsafe.Pointer(&q)),
		uintptr(unsafe.Pointer(&q1)),
		uintptr(unsafe.Pointer(&q2)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
* Find the value in a list of rationals nearest a given reference rational.
*
* @param q      Reference rational
* @param q_list Array of rationals terminated by `{0, 0}`
* @return Index of the nearest value found in the array
 */
//int av_find_nearest_q_idx(AVRational q, const AVRational* q_list);
//未测试
func AvFindNearestQIdx(q AVRational, q_list *AVRational) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_find_nearest_q_idx").Call(
		uintptr(unsafe.Pointer(&q)),
		uintptr(unsafe.Pointer(q_list)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
* Convert an AVRational to a IEEE 32-bit `float` expressed in fixed-point
* format.
*
* @param q Rational to be converted
* @return Equivalent floating-point value, expressed as an unsigned 32-bit
*         integer.
* @note The returned value is platform-indepedant.
 */
//uint32_t av_q2intfloat(AVRational q);
//未测试
func AvQ2intfloat(q AVRational) (res ffcommon.FUint32T, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_q2intfloat").Call(
		uintptr(unsafe.Pointer(&q)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FUint32T(t)
	return
}

/**
* Return the best rational so that a and b are multiple of it.
* If the resulting denominator is larger than max_den, return def.
 */
//AVRational av_gcd_q(AVRational a, AVRational b, int max_den, AVRational def);
//未测试
func AvGcdQ(a AVRational, b AVRational, max_den ffcommon.FInt, def AVRational) (res AVRational, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_gcd_q").Call(
		uintptr(unsafe.Pointer(&a)),
		uintptr(unsafe.Pointer(&b)),
		uintptr(max_den),
		uintptr(unsafe.Pointer(&b)),
	)
	if err != nil {
		//return
	}
	res = *(*AVRational)(unsafe.Pointer(t))
	return
}
