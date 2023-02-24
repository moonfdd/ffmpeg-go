package libavutil

import (
	"fmt"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
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

/**
 * @file
 * timestamp utils, mostly useful for debugging/logging purposes
 */

//#ifndef AVUTIL_TIMESTAMP_H
//#define AVUTIL_TIMESTAMP_H
//
//#include "common.h"
//
//#if defined(__cplusplus) && !defined(__STDC_FORMAT_MACROS) && !defined(PRId64)
//#error missing -D__STDC_FORMAT_MACROS / #define __STDC_FORMAT_MACROS
//#endif

const AV_TS_MAX_STRING_SIZE = 32

/**
 * Fill the provided buffer with a string containing a timestamp
 * representation.
 *
 * @param buf a buffer with size in bytes of at least AV_TS_MAX_STRING_SIZE
 * @param ts the timestamp to represent
 * @return the buffer in input
 */
//static inline char *av_ts_make_string(char *buf, int64_t ts)
//{
//if (ts == AV_NOPTS_VALUE) snprintf(buf, AV_TS_MAX_STRING_SIZE, "NOPTS");
//else                      snprintf(buf, AV_TS_MAX_STRING_SIZE, "%" PRId64, ts);
//return buf;
//}
//todo
// func AvTsMakeString(buf ffcommon.FBuf, ts ffcommon.FInt64T) (res ffcommon.FCharP) {
// 	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_ts_make_string").Call()
// 	if t == 0 {

// 	}
// 	res = ffcommon.StringFromPtr(t)
// 	return
// }

/**
 * Convenience macro, the return value should be used only directly in
 * function arguments but never stand-alone.
 */
//#define av_ts2str(ts) av_ts_make_string((char[AV_TS_MAX_STRING_SIZE]){0}, ts)
func AvTs2str(ts ffcommon.FInt64T) (res ffcommon.FCharP) {
	if ts == AV_NOPTS_VALUE {
		res = "NOPTS"
	} else {
		res = fmt.Sprint(ts)
	}
	return
}

/**
 * Fill the provided buffer with a string containing a timestamp time
 * representation.
 *
 * @param buf a buffer with size in bytes of at least AV_TS_MAX_STRING_SIZE
 * @param ts the timestamp to represent
 * @param tb the timebase of the timestamp
 * @return the buffer in input
 */
//static inline char *av_ts_make_time_string(char *buf, int64_t ts, AVRational *tb)
//{
//if (ts == AV_NOPTS_VALUE) snprintf(buf, AV_TS_MAX_STRING_SIZE, "NOPTS");
//else                      snprintf(buf, AV_TS_MAX_STRING_SIZE, "%.6g", av_q2d(*tb) * ts);
//return buf;
//}
//todo
// func AvTsMakeTimeString() (res ffcommon.FCharP) {
// 	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_ts_make_time_string").Call()
// 	if t == 0 {

// 	}
// 	res = ffcommon.StringFromPtr(t)
// 	return
// }

/**
 * Convenience macro, the return value should be used only directly in
 * function arguments but never stand-alone.
 */
//#define av_ts2timestr(ts, tb) av_ts_make_time_string((char[AV_TS_MAX_STRING_SIZE]){0}, ts, tb)
func AvTs2timestr(ts ffcommon.FInt64T, tb *AVRational) (res ffcommon.FCharP) {
	if ts == AV_NOPTS_VALUE {
		res = "NOPTS"
	} else {
		res = fmt.Sprintf("%.6g", AvQ2d(*tb)*float64(ts))
	}
	return
}

//#endif /* AVUTIL_TIMESTAMP_H */
