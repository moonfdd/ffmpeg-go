package libavutil

import "ffmpeg-go/ffcommon"

//#ifndef av_log2
//av_const int av_log2(unsigned v);
//未测试
func AvLog2(v ffcommon.FUnsigned) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_log2").Call(
		uintptr(v),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//#endif

//#ifndef av_log2_16bit
//av_const int av_log2_16bit(unsigned v);
//未测试
func AvLog2B16bit(v ffcommon.FUnsigned) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_log2_16bit").Call(
		uintptr(v),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//#endif
