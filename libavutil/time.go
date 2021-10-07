package libavutil

import (
	"ffmpeg-go/ffcommon"
)

/**
 * Get the current time in microseconds.
 */
//int64_t av_gettime(void);
//未测试
func AvGettime() (res ffcommon.FInt64T, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_gettime").Call()
	if err != nil {
		return
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
//未测试
func AvGettimeRelative() (res ffcommon.FInt64T, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_gettime_relative").Call()
	if err != nil {
		return
	}
	res = ffcommon.FInt64T(t)
	return
}

/**
 * Indicates with a boolean result if the av_gettime_relative() time source
 * is monotonic.
 */
//int av_gettime_relative_is_monotonic(void);
//未测试
func AvGettimeRelativeIsMonotonic() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_gettime_relative_is_monotonic").Call()
	if err != nil {
		return
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
//未测试
func AvUsleep(usec ffcommon.FUnsigned) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_usleep").Call(
		uintptr(usec),
	)
	if err != nil {
		return
	}
	res = ffcommon.FInt(t)
	return
}
