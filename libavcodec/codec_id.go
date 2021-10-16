package libavcodec

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
)

/**
 * Get the type of the given codec.
 */
//enum AVMediaType avcodec_get_type(enum AVCodecID codec_id);
//未测试
func AvcodecGetType(codec_id ffconstant.AVCodecID) (res ffconstant.AVMediaType, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("avcodec_get_type").Call(
		uintptr(codec_id),
	)
	if err != nil {
		//return
	}
	res = ffconstant.AVMediaType(t)
	return
}

/**
 * Get the name of a codec.
 * @return  a static string identifying the codec; never NULL
 */
//const char *avcodec_get_name(enum AVCodecID id);
//未测试
func AvcodecGetName(id ffconstant.AVCodecID) (res ffcommon.FConstCharP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("avcodec_get_name").Call(
		uintptr(id),
	)
	if err != nil {
		//return
	}
	res = ffcommon.GoAStr(t)
	return
}
