package libavutil

import (
	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"unsafe"
)

/*
 * rational numbers
 * Copyright (c) 2003 Michael Niedermayer <michaelni@gmx.at>
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

/**
 * @file
 * @ingroup lavu_math_rational
 * Utilties for rational number calculation.
 * @author Michael Niedermayer <michaelni@gmx.at>
 */

//#ifndef AVUTIL_RATIONAL_H
//#define AVUTIL_RATIONAL_H
//
//#include <stdint.h>
//#include <limits.h>
//#include "attributes.h"

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
 * Create an AVRational.
 *
 * Useful for compilers that do not support compound literals.
 *
 * @note The return value is not reduced.
 * @see av_reduce()
 */
//static inline AVRational av_make_q(int num, int den)
//{
//AVRational r = { num, den };
//return r;
//}
func AvMakeQ(num, den ffcommon.FInt) (res AVRational) {
	res = AVRational{Num: num, Den: den}
	return
}

/**
 * Compare two rationals.
 *
 * @param a First rational
 * @param b Second rational
 *
 * @return One of the following values:
 *         - 0 if `a == b`
 *         - 1 if `a > b`
 *         - -1 if `a < b`
 *         - `INT_MIN` if one of the values is of the form `0 / 0`
 */
//static inline int av_cmp_q(AVRational a, AVRational b){
//const int64_t tmp= a.num * (int64_t)b.den - b.num * (int64_t)a.den;
//
//if(tmp) return (int)((tmp ^ a.den ^ b.den)>>63)|1;
//else if(b.den && a.den) return 0;
//else if(a.num && b.num) return (a.num>>31) - (b.num>>31);
//else                    return INT_MIN;
//}
func AvCmpQ(a, b AVRational) (res ffcommon.FInt) {
	tmp := a.Num*b.Den - b.Num*a.Den
	//todo
	//if(tmp) return (int)((tmp ^ a.den ^ b.den)>>63)|1;
	//else if(b.den && a.den) return 0;
	//else if(a.num && b.num) return (a.num>>31) - (b.num>>31);
	//else                    return INT_MIN;
	if tmp == 0 {

	}
	//t, _, _ := ffcommon.GetAvutilDll().NewProc("av_cmp_q").Call()
	//if t == 0 {
	//
	//}
	//res = ffcommon.StringFromPtr(t)
	return
}

/**
 * Convert an AVRational to a `double`.
 * @param a AVRational to convert
 * @return `a` in floating-point form
 * @see av_d2q()
 */
//static inline double av_q2d(AVRational a){
//return a.num / (double) a.den;
//}
func AvQ2d(a AVRational) (res ffcommon.FDouble) {
	res = float64(a.Num) / float64(a.Den)
	return
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
func AvReduce(dst_num, dst_den ffcommon.FInt, num, den, max ffcommon.FInt64T) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_reduce").Call(
		uintptr(dst_num),
		uintptr(dst_den),
		uintptr(num),
		uintptr(den),
		uintptr(max),
	)
	if t == 0 {

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
func AvMulQ(b, c AVRational) (res AVRational) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_mul_q").Call(
		uintptr(unsafe.Pointer(&b)),
		uintptr(unsafe.Pointer(&c)),
	)
	if t == 0 {

	}
	res = *(*AVRational)(unsafe.Pointer(&t))
	return
}

/**
 * Divide one rational by another.
 * @param b First rational
 * @param c Second rational
 * @return b/c
 */
//AVRational av_div_q(AVRational b, AVRational c) av_const;
func AvDivQ(b, c AVRational) (res AVRational) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_div_q").Call(
		uintptr(unsafe.Pointer(&b)),
		uintptr(unsafe.Pointer(&c)),
	)
	if t == 0 {

	}
	res = *(*AVRational)(unsafe.Pointer(&t))
	return
}

/**
 * Add two rationals.
 * @param b First rational
 * @param c Second rational
 * @return b+c
 */
//AVRational av_add_q(AVRational b, AVRational c) av_const;
func AvAddQ(b, c AVRational) (res AVRational) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_add_q").Call(
		uintptr(unsafe.Pointer(&b)),
		uintptr(unsafe.Pointer(&c)),
	)
	if t == 0 {

	}
	res = *(*AVRational)(unsafe.Pointer(&t))
	return
}

/**
 * Subtract one rational from another.
 * @param b First rational
 * @param c Second rational
 * @return b-c
 */
//AVRational av_sub_q(AVRational b, AVRational c) av_const;
func AvSubQ(b, c AVRational) (res AVRational) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_sub_q").Call(
		uintptr(unsafe.Pointer(&b)),
		uintptr(unsafe.Pointer(&c)),
	)
	if t == 0 {

	}
	res = *(*AVRational)(unsafe.Pointer(&t))
	return
}

/**
 * Invert a rational.
 * @param q value
 * @return 1 / q
 */
//static av_always_inline AVRational av_inv_q(AVRational q)
//{
//AVRational r = { q.den, q.num };
//return r;
//}
func AvInvQ(q AVRational) (res AVRational) {
	res = AVRational{Num: q.Den, Den: q.Num}
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
func AvD2q(d ffcommon.FDouble, max ffcommon.FInt) (res AVRational) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_d2q").Call(
		uintptr(unsafe.Pointer(&d)),
		uintptr(max),
	)
	if t == 0 {

	}
	res = *(*AVRational)(unsafe.Pointer(&t))
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
func AvNearerQ(q, q1, q2 AVRational) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_nearer_q").Call(
		uintptr(unsafe.Pointer(&q)),
		uintptr(unsafe.Pointer(&q1)),
		uintptr(unsafe.Pointer(&q2)),
	)
	if t == 0 {

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
func AvFindNearestQIdx(q AVRational, q_list *AVRational) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_find_nearest_q_idx").Call(
		uintptr(unsafe.Pointer(&q)),
		uintptr(unsafe.Pointer(q_list)),
	)
	if t == 0 {

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
func AvQ2intfloat(q AVRational) (res ffcommon.FUint32T) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_q2intfloat").Call(
		uintptr(unsafe.Pointer(&q)),
	)
	if t == 0 {

	}
	res = ffcommon.FUint32T(t)
	return
}

/**
 * Return the best rational so that a and b are multiple of it.
 * If the resulting denominator is larger than max_den, return def.
 */
//AVRational av_gcd_q(AVRational a, AVRational b, int max_den, AVRational def);
func AvGcdQ(a, b AVRational, max_den ffcommon.FInt, def AVRational) (res AVRational) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_gcd_q").Call(
		uintptr(unsafe.Pointer(&a)),
		uintptr(unsafe.Pointer(&b)),
		uintptr(max_den),
		uintptr(unsafe.Pointer(&def)),
	)
	if t == 0 {

	}
	res = *(*AVRational)(unsafe.Pointer(&t))
	return
}

/**
 * @}
 */

//#endif /* AVUTIL_RATIONAL_H */
