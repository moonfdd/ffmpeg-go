package libavutil

import "github.com/moonfdd/ffmpeg-go/ffcommon"

/*
 * Copyright (c) 2000-2003 Fabrice Bellard
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

//#ifndef AVUTIL_TIME_H
//#define AVUTIL_TIME_H
//
//#include <stdint.h>

/**
 * Get the current time in microseconds.
 */
//int64_t av_gettime(void);
func AvGettime() (res ffcommon.FInt64T) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_gettime").Call()
	if t == 0 {

	}
	res = ffcommon.FInt64T(t)
	return
}

/**
 * Get the current time in microseconds since some unspecified starting point.
 * On platforms that support it, the time comes from a monotonic clock
 * This property makes this time source ideal for measuring relative time.
 * The returned values may not be monotonic on platforms where a monotonic
 * clock is not available.
 */
//int64_t av_gettime_relative(void);
func AvGettimeRelative() (res ffcommon.FInt64T) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_gettime_relative").Call()
	if t == 0 {

	}
	res = ffcommon.FInt64T(t)
	return
}

/**
 * Indicates with a boolean result if the av_gettime_relative() time source
 * is monotonic.
 */
//int av_gettime_relative_is_monotonic(void);
func AvGettimeRelativeIsMonotonic() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_gettime_relative_is_monotonic").Call()
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Sleep for a period of time.  Although the duration is expressed in
 * microseconds, the actual delay may be rounded to the precision of the
 * system timer.
 *
 * @param  usec Number of microseconds to sleep.
 * @return zero on success or (negative) error code.
 */
//int av_usleep(unsigned usec);
func AvUsleep(usec ffcommon.FUnsigned) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_usleep").Call(
		uintptr(usec),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//#endif /* AVUTIL_TIME_H */
