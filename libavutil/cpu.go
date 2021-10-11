package libavutil

import (
	"ffmpeg-go/ffcommon"
	"syscall"
	"unsafe"
)

/**
 * Return the flags which specify extensions supported by the CPU.
 * The returned value is affected by av_force_cpu_flags() if that was used
 * before. So av_get_cpu_flags() can easily be used in an application to
 * detect the enabled cpu flags.
 */
//int av_get_cpu_flags(void);
//未测试
func AvGetCpuFlags() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_get_cpu_flags").Call()
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Disables cpu detection and forces the specified flags.
 * -1 is a special case that disables forcing of specific flags.
 */
//void av_force_cpu_flags(int flags);
//未测试
func AvForceCpuFlags(flags ffcommon.FInt) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_force_cpu_flags").Call(
		uintptr(flags),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Set a mask on flags returned by av_get_cpu_flags().
 * This function is mainly useful for testing.
 * Please use av_force_cpu_flags() and av_get_cpu_flags() instead which are more flexible
 */
//attribute_deprecated void av_set_cpu_flags_mask(int mask);
//未测试
func AvSetCpuFlagsMask(flags ffcommon.FInt) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_set_cpu_flags_mask").Call(
		uintptr(flags),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Parse CPU flags from a string.
 *
 * The returned flags contain the specified flags as well as related unspecified flags.
 *
 * This function exists only for compatibility with libav.
 * Please use av_parse_cpu_caps() when possible.
 * @return a combination of AV_CPU_* flags, negative on error.
 */
//attribute_deprecated
//int av_parse_cpu_flags(const char *s);
//未测试
func AvParseCpuFlags(s ffcommon.FConstCharP) (res ffcommon.FInt, err error) {
	var t uintptr
	var sp *byte
	sp, err = syscall.BytePtrFromString(s)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_parse_cpu_flags").Call(
		uintptr(unsafe.Pointer(sp)),
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
 * Parse CPU caps from a string and update the given AV_CPU_* flags based on that.
 *
 * @return negative on error.
 */
//int av_parse_cpu_caps(unsigned *flags, const char *s);
//未测试
func AvParseCpuCaps(flags *ffcommon.FUnsigned, s ffcommon.FConstCharP) (res ffcommon.FInt, err error) {
	var t uintptr
	var sp *byte
	sp, err = syscall.BytePtrFromString(s)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_parse_cpu_caps").Call(
		uintptr(unsafe.Pointer(flags)),
		uintptr(unsafe.Pointer(sp)),
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
 * @return the number of logical CPU cores present.
 */
//int av_cpu_count(void);
//未测试
func AvCpuCount() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_cpu_count").Call()
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Get the maximum data alignment that may be required by FFmpeg.
 *
 * Note that this is affected by the build configuration and the CPU flags mask,
 * so e.g. if the CPU supports AVX, but libavutil has been built with
 * --disable-avx or the AV_CPU_FLAG_AVX flag has been disabled through
 *  av_set_cpu_flags_mask(), then this function will behave as if AVX is not
 *  present.
 */
//size_t av_cpu_max_align(void);
//未测试
func AvCpuMaxAlign() (res ffcommon.FSizeT, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_cpu_max_align").Call()
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FSizeT(t)
	return
}
