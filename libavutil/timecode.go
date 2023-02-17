package libavutil

import (
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
)

/*
 * Copyright (c) 2006 Smartjog S.A.S, Baptiste Coudurier <baptiste.coudurier@gmail.com>
 * Copyright (c) 2011-2012 Smartjog S.A.S, Clément Bœsch <clement.boesch@smartjog.com>
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
 * Timecode helpers header
 */

//#ifndef AVUTIL_TIMECODE_H
//#define AVUTIL_TIMECODE_H
//
//#include <stdint.h>
//#include "rational.h"

const AV_TIMECODE_STR_SIZE = 23

type AVTimecodeFlag int32

const (
	AV_TIMECODE_FLAG_DROPFRAME     = 1 << 0 ///< timecode is drop frame
	AV_TIMECODE_FLAG_24HOURSMAX    = 1 << 1 ///< timecode wraps after 24 hours
	AV_TIMECODE_FLAG_ALLOWNEGATIVE = 1 << 2 ///< negative time values are allowed
)

type AVTimecode struct {
	Start ffcommon.FInt      ///< timecode frame start (first base frame number)
	Flags ffcommon.FUint32T  ///< flags such as drop frame, +24 hours support, ...
	Rate  AVRational         ///< frame rate in rational form
	Fps   ffcommon.FUnsigned ///< frame per second; must be consistent with the rate field
}

/**
 * Adjust frame number for NTSC drop frame time code.
 *
 * @param framenum frame number to adjust
 * @param fps      frame per second, multiples of 30
 * @return         adjusted frame number
 * @warning        adjustment is only valid for multiples of NTSC 29.97
 */
//int av_timecode_adjust_ntsc_framenum2(int framenum, int fps);
func AvTimecodeAdjustNtscFramenum2(framenum, fps ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_timecode_adjust_ntsc_framenum2").Call(
		uintptr(framenum),
		uintptr(fps),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Convert frame number to SMPTE 12M binary representation.
 *
 * @param tc       timecode data correctly initialized
 * @param framenum frame number
 * @return         the SMPTE binary representation
 *
 * See SMPTE ST 314M-2005 Sec 4.4.2.2.1 "Time code pack (TC)"
 * the format description as follows:
 * bits 0-5:   hours, in BCD(6bits)
 * bits 6:     BGF1
 * bits 7:     BGF2 (NTSC) or FIELD (PAL)
 * bits 8-14:  minutes, in BCD(7bits)
 * bits 15:    BGF0 (NTSC) or BGF2 (PAL)
 * bits 16-22: seconds, in BCD(7bits)
 * bits 23:    FIELD (NTSC) or BGF0 (PAL)
 * bits 24-29: frames, in BCD(6bits)
 * bits 30:    drop  frame flag (0: non drop,    1: drop)
 * bits 31:    color frame flag (0: unsync mode, 1: sync mode)
 * @note BCD numbers (6 or 7 bits): 4 or 5 lower bits for units, 2 higher bits for tens.
 * @note Frame number adjustment is automatically done in case of drop timecode,
 *       you do NOT have to call av_timecode_adjust_ntsc_framenum2().
 * @note The frame number is relative to tc->start.
 * @note Color frame (CF) and binary group flags (BGF) bits are set to zero.
 */
//uint32_t av_timecode_get_smpte_from_framenum(const AVTimecode *tc, int framenum);
func (tc *AVTimecode) AvTimecodeGetSmpteFromFramenum(framenum ffcommon.FInt) (res ffcommon.FUint32T) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_timecode_get_smpte_from_framenum").Call(
		uintptr(unsafe.Pointer(tc)),
		uintptr(framenum),
	)
	res = ffcommon.FUint32T(t)
	return
}

/**
 * Convert sei info to SMPTE 12M binary representation.
 *
 * @param rate     frame rate in rational form
 * @param drop     drop flag
 * @param hh       hour
 * @param mm       minute
 * @param ss       second
 * @param ff       frame number
 * @return         the SMPTE binary representation
 */
//uint32_t av_timecode_get_smpte(AVRational rate, int drop, int hh, int mm, int ss, int ff);
func AvTimecodeGetSmpte(rate AVRational, drop, hh, mm, ss, ff ffcommon.FInt) (res ffcommon.FUint32T) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_timecode_get_smpte").Call(
		uintptr(unsafe.Pointer(&rate)),
		uintptr(drop),
		uintptr(hh),
		uintptr(mm),
		uintptr(ss),
		uintptr(ff),
	)
	res = ffcommon.FUint32T(t)
	return
}

/**
 * Load timecode string in buf.
 *
 * @param buf      destination buffer, must be at least AV_TIMECODE_STR_SIZE long
 * @param tc       timecode data correctly initialized
 * @param framenum frame number
 * @return         the buf parameter
 *
 * @note Timecode representation can be a negative timecode and have more than
 *       24 hours, but will only be honored if the flags are correctly set.
 * @note The frame number is relative to tc->start.
 */
//char *av_timecode_make_string(const AVTimecode *tc, char *buf, int framenum);
func (tc *AVTimecode) AvTimecodeMakeString(buf ffcommon.FCharP, framenum ffcommon.FInt) (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_timecode_make_string").Call(
		uintptr(unsafe.Pointer(tc)),
		ffcommon.UintPtrFromString(buf),
		uintptr(framenum),
	)
	res = ffcommon.StringFromPtr(t)
	return
}

/**
 * Get the timecode string from the SMPTE timecode format.
 *
 * In contrast to av_timecode_make_smpte_tc_string this function supports 50/60
 * fps timecodes by using the field bit.
 *
 * @param buf        destination buffer, must be at least AV_TIMECODE_STR_SIZE long
 * @param rate       frame rate of the timecode
 * @param tcsmpte    the 32-bit SMPTE timecode
 * @param prevent_df prevent the use of a drop flag when it is known the DF bit
 *                   is arbitrary
 * @param skip_field prevent the use of a field flag when it is known the field
 *                   bit is arbitrary (e.g. because it is used as PC flag)
 * @return           the buf parameter
 */
//char *av_timecode_make_smpte_tc_string2(char *buf, AVRational rate, uint32_t tcsmpte, int prevent_df, int skip_field);
func AvTimecodeMakeSmpteTcString2(buf ffcommon.FCharP, rate AVRational, tcsmpte ffcommon.FUint32T, prevent_df, skip_field ffcommon.FInt) (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_timecode_make_smpte_tc_string2").Call(
		ffcommon.UintPtrFromString(buf),
		uintptr(unsafe.Pointer(&rate)),
		uintptr(tcsmpte),
		uintptr(prevent_df),
		uintptr(skip_field),
	)
	res = ffcommon.StringFromPtr(t)
	return
}

/**
 * Get the timecode string from the SMPTE timecode format.
 *
 * @param buf        destination buffer, must be at least AV_TIMECODE_STR_SIZE long
 * @param tcsmpte    the 32-bit SMPTE timecode
 * @param prevent_df prevent the use of a drop flag when it is known the DF bit
 *                   is arbitrary
 * @return           the buf parameter
 */
//char *av_timecode_make_smpte_tc_string(char *buf, uint32_t tcsmpte, int prevent_df);
func AvTimecodeMakeSmpteTcString(buf ffcommon.FCharP, tcsmpte ffcommon.FUint32T, prevent_df ffcommon.FInt) (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_timecode_make_smpte_tc_string").Call(
		ffcommon.UintPtrFromString(buf),
		uintptr(tcsmpte),
		uintptr(prevent_df),
	)
	res = ffcommon.StringFromPtr(t)
	return
}

/**
 * Get the timecode string from the 25-bit timecode format (MPEG GOP format).
 *
 * @param buf     destination buffer, must be at least AV_TIMECODE_STR_SIZE long
 * @param tc25bit the 25-bits timecode
 * @return        the buf parameter
 */
//char *av_timecode_make_mpeg_tc_string(char *buf, uint32_t tc25bit);
func AvTimecodeMakeMpegTcString(buf ffcommon.FCharP, tc25bit ffcommon.FUint32T) (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_timecode_make_mpeg_tc_string").Call(
		ffcommon.UintPtrFromString(buf),
		uintptr(tc25bit),
	)
	res = ffcommon.StringFromPtr(t)
	return
}

/**
 * Init a timecode struct with the passed parameters.
 *
 * @param log_ctx     a pointer to an arbitrary struct of which the first field
 *                    is a pointer to an AVClass struct (used for av_log)
 * @param tc          pointer to an allocated AVTimecode
 * @param rate        frame rate in rational form
 * @param flags       miscellaneous flags such as drop frame, +24 hours, ...
 *                    (see AVTimecodeFlag)
 * @param frame_start the first frame number
 * @return            0 on success, AVERROR otherwise
 */
//int av_timecode_init(AVTimecode *tc, AVRational rate, int flags, int frame_start, void *log_ctx);
func (tc *AVTimecode) AvTimecodeInit(rate AVRational, flags, frame_start ffcommon.FInt, log_ctx ffcommon.FVoidP) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_timecode_init").Call(
		uintptr(unsafe.Pointer(tc)),
		uintptr(unsafe.Pointer(&rate)),
		uintptr(flags),
		uintptr(frame_start),
		log_ctx,
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Init a timecode struct from the passed timecode components.
 *
 * @param log_ctx     a pointer to an arbitrary struct of which the first field
 *                    is a pointer to an AVClass struct (used for av_log)
 * @param tc          pointer to an allocated AVTimecode
 * @param rate        frame rate in rational form
 * @param flags       miscellaneous flags such as drop frame, +24 hours, ...
 *                    (see AVTimecodeFlag)
 * @param hh          hours
 * @param mm          minutes
 * @param ss          seconds
 * @param ff          frames
 * @return            0 on success, AVERROR otherwise
 */
//int av_timecode_init_from_components(AVTimecode *tc, AVRational rate, int flags, int hh, int mm, int ss, int ff, void *log_ctx);
func (tc *AVTimecode) AvTimecodeInitFromComponents(rate AVRational, flags, hh, mm, ss, ff ffcommon.FInt, log_ctx ffcommon.FVoidP) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_timecode_init_from_components").Call(
		uintptr(unsafe.Pointer(tc)),
		uintptr(unsafe.Pointer(&rate)),
		uintptr(flags),
		uintptr(hh),
		uintptr(mm),
		uintptr(ss),
		uintptr(ff),
		log_ctx,
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Parse timecode representation (hh:mm:ss[:;.]ff).
 *
 * @param log_ctx a pointer to an arbitrary struct of which the first field is a
 *                pointer to an AVClass struct (used for av_log).
 * @param tc      pointer to an allocated AVTimecode
 * @param rate    frame rate in rational form
 * @param str     timecode string which will determine the frame start
 * @return        0 on success, AVERROR otherwise
 */
//int av_timecode_init_from_string(AVTimecode *tc, AVRational rate, const char *str, void *log_ctx);
func (tc *AVTimecode) AvTimecodeInitFromString(rate AVRational, str ffcommon.FConstCharP, log_ctx ffcommon.FVoidP) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_timecode_init_from_string").Call(
		uintptr(unsafe.Pointer(tc)),
		uintptr(unsafe.Pointer(&rate)),
		ffcommon.UintPtrFromString(str),
		log_ctx,
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Check if the timecode feature is available for the given frame rate
 *
 * @return 0 if supported, <0 otherwise
 */
//int av_timecode_check_frame_rate(AVRational rate);
func AvTimecodeCheckFrameRate(rate AVRational) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_timecode_check_frame_rate").Call(
		uintptr(unsafe.Pointer(&rate)),
	)
	res = ffcommon.FInt(t)
	return
}

//#endif /* AVUTIL_TIMECODE_H */
