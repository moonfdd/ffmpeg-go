package libavutil

import (
	"ffmpeg-go/ffcommon"
)

/**
 * Get a seed to use in conjunction with random functions.
 * This function tries to provide a good seed at a best effort bases.
 * Its possible to call this function multiple times if more bits are needed.
 * It can be quite slow, which is why it should only be used as seed for a faster
 * PRNG. The quality of the seed depends on the platform.
 */
//uint32_t av_get_random_seed(void);
//未测试
func AvGetRandomSeed() (res ffcommon.FUint32T, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_get_random_seed").Call()
	if err != nil {
		//return
	}
	res = ffcommon.FUint32T(t)
	return
}
