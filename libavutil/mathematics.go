package libavutil

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"unsafe"
)

/**
 * Compute the greatest common divisor of two integer operands.
 *
 * @param a,b Operands
 * @return GCD of a and b up to sign; if a >= 0 and b >= 0, return value is >= 0;
 * if a == 0 and b == 0, returns 0.
 */
//int64_t av_const av_gcd(int64_t a, int64_t b);
//未测试
func AvGcd(a ffcommon.FInt64T, b ffcommon.FInt64T) (res ffcommon.FInt64T, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_gcd").Call(
		uintptr(a),
		uintptr(b),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt64T(t)
	return
}

/**
 * Rescale a 64-bit integer with rounding to nearest.
 *
 * The operation is mathematically equivalent to `a * b / c`, but writing that
 * directly can overflow.
 *
 * This function is equivalent to av_rescale_rnd() with #AV_ROUND_NEAR_INF.
 *
 * @see av_rescale_rnd(), av_rescale_q(), av_rescale_q_rnd()
 */
//int64_t av_rescale(int64_t a, int64_t b, int64_t c) av_const;
//未测试
func AvRescale(a ffcommon.FInt64T, b ffcommon.FInt64T) (res ffcommon.FInt64T, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_rescale").Call(
		uintptr(a),
		uintptr(b),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt64T(t)
	return
}

/**
 * Rescale a 64-bit integer with specified rounding.
 *
 * The operation is mathematically equivalent to `a * b / c`, but writing that
 * directly can overflow, and does not support different rounding methods.
 *
 * @see av_rescale(), av_rescale_q(), av_rescale_q_rnd()
 */
//int64_t av_rescale_rnd(int64_t a, int64_t b, int64_t c, enum AVRounding rnd) av_const;
//未测试
func AvRescaleRnd(a ffcommon.FInt64T, b ffcommon.FInt64T, t0 ffcommon.FInt64T, rnd ffconstant.AVRounding) (res ffcommon.FInt64T, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_rescale_rnd").Call(
		uintptr(a),
		uintptr(b),
		uintptr(t0),
		uintptr(rnd),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt64T(t)
	return
}

/**
 * Rescale a 64-bit integer by 2 rational numbers.
 *
 * The operation is mathematically equivalent to `a * bq / cq`.
 *
 * This function is equivalent to av_rescale_q_rnd() with #AV_ROUND_NEAR_INF.
 *
 * @see av_rescale(), av_rescale_rnd(), av_rescale_q_rnd()
 */
//int64_t av_rescale_q(int64_t a, AVRational bq, AVRational cq) av_const;
//未测试
func AvRescaleQ(a ffcommon.FInt64T, bq AVRational, cq AVRational) (res ffcommon.FInt64T, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_rescale_q").Call(
		uintptr(a),
		uintptr(unsafe.Pointer(&bq)),
		uintptr(unsafe.Pointer(&cq)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt64T(t)
	return
}

/**
 * Rescale a 64-bit integer by 2 rational numbers with specified rounding.
 *
 * The operation is mathematically equivalent to `a * bq / cq`.
 *
 * @see av_rescale(), av_rescale_rnd(), av_rescale_q()
 */
//int64_t av_rescale_q_rnd(int64_t a, AVRational bq, AVRational cq,
//enum AVRounding rnd) av_const;
//未测试
func AvRescaleQRnd(a ffcommon.FInt64T, bq AVRational, cq AVRational, rnd ffconstant.AVRounding) (res ffcommon.FInt64T, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_rescale_q_rnd").Call(
		uintptr(a),
		uintptr(unsafe.Pointer(&bq)),
		uintptr(unsafe.Pointer(&cq)),
		uintptr(rnd),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt64T(t)
	return
}

/**
 * Compare two timestamps each in its own time base.
 *
 * @return One of the following values:
 *         - -1 if `ts_a` is before `ts_b`
 *         - 1 if `ts_a` is after `ts_b`
 *         - 0 if they represent the same position
 *
 * @warning
 * The result of the function is undefined if one of the timestamps is outside
 * the `int64_t` range when represented in the other's timebase.
 */
//int av_compare_ts(int64_t ts_a, AVRational tb_a, int64_t ts_b, AVRational tb_b);
//未测试
func AvCompareTs(ts_a ffcommon.FInt64T, tb_a AVRational, ts_b ffcommon.FInt64T, tb_b AVRational) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_compare_ts").Call(
		uintptr(ts_a),
		uintptr(unsafe.Pointer(&tb_a)),
		uintptr(ts_b),
		uintptr(unsafe.Pointer(&tb_b)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Compare the remainders of two integer operands divided by a common divisor.
 *
 * In other words, compare the least significant `log2(mod)` bits of integers
 * `a` and `b`.
 *
 * @code{.c}
 * av_compare_mod(0x11, 0x02, 0x10) < 0 // since 0x11 % 0x10  (0x1) < 0x02 % 0x10  (0x2)
 * av_compare_mod(0x11, 0x02, 0x20) > 0 // since 0x11 % 0x20 (0x11) > 0x02 % 0x20 (0x02)
 * @endcode
 *
 * @param a,b Operands
 * @param mod Divisor; must be a power of 2
 * @return
 *         - a negative value if `a % mod < b % mod`
 *         - a positive value if `a % mod > b % mod`
 *         - zero             if `a % mod == b % mod`
 */
//int64_t av_compare_mod(uint64_t a, uint64_t b, uint64_t mod);
//未测试
func AvCompareMod(a, b, mod ffcommon.FUint64T) (res ffcommon.FInt64T, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_compare_mod").Call(
		uintptr(a),
		uintptr(b),
		uintptr(mod),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt64T(t)
	return
}

/**
 * Rescale a timestamp while preserving known durations.
 *
 * This function is designed to be called per audio packet to scale the input
 * timestamp to a different time base. Compared to a simple av_rescale_q()
 * call, this function is robust against possible inconsistent frame durations.
 *
 * The `last` parameter is a state variable that must be preserved for all
 * subsequent calls for the same stream. For the first call, `*last` should be
 * initialized to #AV_NOPTS_VALUE.
 *
 * @param[in]     in_tb    Input time base
 * @param[in]     in_ts    Input timestamp
 * @param[in]     fs_tb    Duration time base; typically this is finer-grained
 *                         (greater) than `in_tb` and `out_tb`
 * @param[in]     duration Duration till the next call to this function (i.e.
 *                         duration of the current packet/frame)
 * @param[in,out] last     Pointer to a timestamp expressed in terms of
 *                         `fs_tb`, acting as a state variable
 * @param[in]     out_tb   Output timebase
 * @return        Timestamp expressed in terms of `out_tb`
 *
 * @note In the context of this function, "duration" is in term of samples, not
 *       seconds.
 */
//int64_t av_rescale_delta(AVRational in_tb, int64_t in_ts,  AVRational fs_tb, int duration, int64_t *last, AVRational out_tb);
//未测试
func AvRescaleDelta(in_tb AVRational, in_ts ffcommon.FInt64T, fs_tb AVRational, duration ffcommon.FInt, last *ffcommon.FInt64T, out_tb AVRational) (res ffcommon.FInt64T, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_rescale_delta").Call(
		uintptr(unsafe.Pointer(&in_tb)),
		uintptr(in_ts),
		uintptr(unsafe.Pointer(&fs_tb)),
		uintptr(duration),
		uintptr(unsafe.Pointer(last)),
		uintptr(unsafe.Pointer(&out_tb)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt64T(t)
	return
}

/**
 * Add a value to a timestamp.
 *
 * This function guarantees that when the same value is repeatly added that
 * no accumulation of rounding errors occurs.
 *
 * @param[in] ts     Input timestamp
 * @param[in] ts_tb  Input timestamp time base
 * @param[in] inc    Value to be added
 * @param[in] inc_tb Time base of `inc`
 */
//int64_t av_add_stable(AVRational ts_tb, int64_t ts, AVRational inc_tb, int64_t inc);
//未测试
func AvAddStable(ts_tb AVRational, ts ffcommon.FInt64T, inc_tb AVRational, inc ffcommon.FInt64T) (res ffcommon.FInt64T, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_add_stable").Call(
		uintptr(unsafe.Pointer(&ts_tb)),
		uintptr(ts),
		uintptr(unsafe.Pointer(&inc_tb)),
		uintptr(inc),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt64T(t)
	return
}
