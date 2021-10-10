package libavutil

import (
	"ffmpeg-go/ffcommon"
	"syscall"
	"unsafe"
)

/**
 * @file
 * misc parsing utilities
 */

/**
 * Parse str and store the parsed ratio in q.
 *
 * Note that a ratio with infinite (1/0) or negative value is
 * considered valid, so you should check on the returned value if you
 * want to exclude those values.
 *
 * The undefined value can be expressed using the "0:0" string.
 *
 * @param[in,out] q pointer to the AVRational which will contain the ratio
 * @param[in] str the string to parse: it has to be a string in the format
 * num:den, a float number or an expression
 * @param[in] max the maximum allowed numerator and denominator
 * @param[in] log_offset log level offset which is applied to the log
 * level of log_ctx
 * @param[in] log_ctx parent logging context
 * @return >= 0 on success, a negative error code otherwise
 */
//int av_parse_ratio(AVRational *q, const char *str, int max,
//int log_offset, void *log_ctx);
//未测试
func AvParseRatio(q *AVRational, str ffcommon.FConstCharP, max ffcommon.FInt,
	log_offset ffcommon.FInt, log_ctx ffcommon.FVoidP) (res ffcommon.FInt, err error) {
	var t uintptr
	var strp *byte
	strp, err = syscall.BytePtrFromString(str)
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_parse_ratio").Call(
		uintptr(unsafe.Pointer(q)),
		uintptr(unsafe.Pointer(strp)),
		uintptr(max),
		uintptr(log_offset),
		uintptr(log_ctx),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Parse str and put in width_ptr and height_ptr the detected values.
 *
 * @param[in,out] width_ptr pointer to the variable which will contain the detected
 * width value
 * @param[in,out] height_ptr pointer to the variable which will contain the detected
 * height value
 * @param[in] str the string to parse: it has to be a string in the format
 * width x height or a valid video size abbreviation.
 * @return >= 0 on success, a negative error code otherwise
 */
//int av_parse_video_size(int *width_ptr, int *height_ptr, const char *str);
//未测试
func AvParseVideoSize(width_ptr *ffcommon.FInt, height_ptr *ffcommon.FInt, str ffcommon.FConstCharP) (res ffcommon.FInt, err error) {
	var t uintptr
	var strp *byte
	strp, err = syscall.BytePtrFromString(str)
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_parse_video_size").Call(
		uintptr(unsafe.Pointer(width_ptr)),
		uintptr(unsafe.Pointer(height_ptr)),
		uintptr(unsafe.Pointer(strp)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Parse str and store the detected values in *rate.
 *
 * @param[in,out] rate pointer to the AVRational which will contain the detected
 * frame rate
 * @param[in] str the string to parse: it has to be a string in the format
 * rate_num / rate_den, a float number or a valid video rate abbreviation
 * @return >= 0 on success, a negative error code otherwise
 */
//int av_parse_video_rate(AVRational *rate, const char *str);
//未测试
func AvParseVideoRate(rate *AVRational, str ffcommon.FConstCharP) (res ffcommon.FInt, err error) {
	var t uintptr
	var strp *byte
	strp, err = syscall.BytePtrFromString(str)
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_parse_video_rate").Call(
		uintptr(unsafe.Pointer(rate)),
		uintptr(unsafe.Pointer(strp)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Put the RGBA values that correspond to color_string in rgba_color.
 *
 * @param color_string a string specifying a color. It can be the name of
 * a color (case insensitive match) or a [0x|#]RRGGBB[AA] sequence,
 * possibly followed by "@" and a string representing the alpha
 * component.
 * The alpha component may be a string composed by "0x" followed by an
 * hexadecimal number or a decimal number between 0.0 and 1.0, which
 * represents the opacity value (0x00/0.0 means completely transparent,
 * 0xff/1.0 completely opaque).
 * If the alpha component is not specified then 0xff is assumed.
 * The string "random" will result in a random color.
 * @param slen length of the initial part of color_string containing the
 * color. It can be set to -1 if color_string is a null terminated string
 * containing nothing else than the color.
 * @return >= 0 in case of success, a negative value in case of
 * failure (for example if color_string cannot be parsed).
 */
//int av_parse_color(uint8_t *rgba_color, const char *color_string, int slen,
//void *log_ctx);
//未测试
func AvParseColor(rgba_color *ffcommon.FUint8T, color_string ffcommon.FConstCharP, slen ffcommon.FInt,
	log_ctx ffcommon.FVoidP) (res ffcommon.FInt, err error) {
	var t uintptr
	var color_stringp *byte
	color_stringp, err = syscall.BytePtrFromString(color_string)
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_parse_color").Call(
		uintptr(unsafe.Pointer(rgba_color)),
		uintptr(unsafe.Pointer(color_stringp)),
		uintptr(slen),
		uintptr(log_ctx),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Get the name of a color from the internal table of hard-coded named
 * colors.
 *
 * This function is meant to enumerate the color names recognized by
 * av_parse_color().
 *
 * @param color_idx index of the requested color, starting from 0
 * @param rgbp      if not NULL, will point to a 3-elements array with the color value in RGB
 * @return the color name string or NULL if color_idx is not in the array
 */
//const char *av_get_known_color_name(int color_idx, const uint8_t **rgb);
//未测试
func AvGetKnownColorName(color_idx ffcommon.FInt, rgb **ffcommon.FUint8T) (res ffcommon.FConstCharP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_get_known_color_name").Call(
		uintptr(color_idx),
		uintptr(unsafe.Pointer(&rgb)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.GoAStr(t)
	return
}

/**
 * Parse timestr and return in *time a corresponding number of
 * microseconds.
 *
 * @param timeval puts here the number of microseconds corresponding
 * to the string in timestr. If the string represents a duration, it
 * is the number of microseconds contained in the time interval.  If
 * the string is a date, is the number of microseconds since 1st of
 * January, 1970 up to the time of the parsed date.  If timestr cannot
 * be successfully parsed, set *time to INT64_MIN.

 * @param timestr a string representing a date or a duration.
 * - If a date the syntax is:
 * @code
 * [{YYYY-MM-DD|YYYYMMDD}[T|t| ]]{{HH:MM:SS[.m...]]]}|{HHMMSS[.m...]]]}}[Z]
 * now
 * @endcode
 * If the value is "now" it takes the current time.
 * Time is local time unless Z is appended, in which case it is
 * interpreted as UTC.
 * If the year-month-day part is not specified it takes the current
 * year-month-day.
 * - If a duration the syntax is:
 * @code
 * [-][HH:]MM:SS[.m...]
 * [-]S+[.m...]
 * @endcode
 * @param duration flag which tells how to interpret timestr, if not
 * zero timestr is interpreted as a duration, otherwise as a date
 * @return >= 0 in case of success, a negative value corresponding to an
 * AVERROR code otherwise
 */
//int av_parse_time(int64_t *timeval, const char *timestr, int duration);
//未测试
func AvParseTime(timeval *ffcommon.FUint64T, timestr ffcommon.FConstCharP, duration ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	var timestrp *byte
	timestrp, err = syscall.BytePtrFromString(timestr)
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_parse_time").Call(
		uintptr(unsafe.Pointer(timeval)),
		uintptr(unsafe.Pointer(&timestrp)),
		uintptr(duration),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Attempt to find a specific tag in a URL.
 *
 * syntax: '?tag1=val1&tag2=val2...'. Little URL decoding is done.
 * Return 1 if found.
 */
//int av_find_info_tag(char *arg, int arg_size, const char *tag1, const char *info);
//未测试
func AvFindInfoTag(arg ffcommon.FConstCharP, arg_size ffcommon.FInt, tag1 ffcommon.FConstCharP, info ffcommon.FConstCharP) (res ffcommon.FInt, err error) {
	var t uintptr
	var argp *byte
	argp, err = syscall.BytePtrFromString(arg)
	var tag1p *byte
	tag1p, err = syscall.BytePtrFromString(tag1)
	var infop *byte
	infop, err = syscall.BytePtrFromString(info)
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_find_info_tag").Call(
		uintptr(unsafe.Pointer(argp)),
		uintptr(arg_size),
		uintptr(unsafe.Pointer(tag1p)),
		uintptr(unsafe.Pointer(infop)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Simplified version of strptime
 *
 * Parse the input string p according to the format string fmt and
 * store its results in the structure dt.
 * This implementation supports only a subset of the formats supported
 * by the standard strptime().
 *
 * The supported input field descriptors are listed below.
 * - %H: the hour as a decimal number, using a 24-hour clock, in the
 *   range '00' through '23'
 * - %J: hours as a decimal number, in the range '0' through INT_MAX
 * - %M: the minute as a decimal number, using a 24-hour clock, in the
 *   range '00' through '59'
 * - %S: the second as a decimal number, using a 24-hour clock, in the
 *   range '00' through '59'
 * - %Y: the year as a decimal number, using the Gregorian calendar
 * - %m: the month as a decimal number, in the range '1' through '12'
 * - %d: the day of the month as a decimal number, in the range '1'
 *   through '31'
 * - %T: alias for '%H:%M:%S'
 * - %%: a literal '%'
 *
 * @return a pointer to the first character not processed in this function
 *         call. In case the input string contains more characters than
 *         required by the format string the return value points right after
 *         the last consumed input character. In case the whole input string
 *         is consumed the return value points to the null byte at the end of
 *         the string. On failure NULL is returned.
 */
//char *av_small_strptime(const char *p, const char *fmt, struct tm *dt);
//未测试
func AvSmallStrptime(p ffcommon.FConstCharP, fmt0 ffcommon.FConstCharP, dt *Tm) (res ffcommon.FCharP, err error) {
	var t uintptr
	var pp *byte
	pp, err = syscall.BytePtrFromString(p)
	var fmt0p *byte
	fmt0p, err = syscall.BytePtrFromString(fmt0)
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_small_strptime").Call(
		uintptr(unsafe.Pointer(pp)),
		uintptr(unsafe.Pointer(fmt0p)),
		uintptr(unsafe.Pointer(dt)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.GoAStr(t)
	return
}

/**
 * Convert the decomposed UTC time in tm to a time_t value.
 */
//time_t av_timegm(struct tm *tm);
//未测试
func AvTimegm(dt *Tm) (res ffcommon.FTimeT, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_timegm").Call(
		uintptr(unsafe.Pointer(dt)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FTimeT(t)
	return
}

type Tm struct {
	TmSec   ffcommon.FInt // seconds after the minute - [0, 60] including leap second
	TmMin   ffcommon.FInt // minutes after the hour - [0, 59]
	TmHour  ffcommon.FInt // hours since midnight - [0, 23]
	TmMday  ffcommon.FInt // day of the month - [1, 31]
	TmMon   ffcommon.FInt // months since January - [0, 11]
	TmYear  ffcommon.FInt // years since 1900
	TmWday  ffcommon.FInt // days since Sunday - [0, 6]
	TmYday  ffcommon.FInt // days since January 1 - [0, 365]
	TmIsdst ffcommon.FInt // daylight savings time flag
}
