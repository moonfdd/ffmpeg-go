package libavutil

import (
	"ffmpeg-go/ffcommon"
)

/**
 * Assert that floating point operations can be executed.
 *
 * This will av_assert0() that the cpu is not in MMX state on X86
 */
//void av_assert0_fpu(void);
//未测试
func AvAssert0Fpu() (err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_assert0_fpu").Call()
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}
