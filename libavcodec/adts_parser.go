package libavcodec

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

/**
 * Extract the number of samples and frames from AAC data.
 * @param[in]  buf     pointer to AAC data buffer
 * @param[out] samples Pointer to where number of samples is written
 * @param[out] frames  Pointer to where number of frames is written
 * @return Returns 0 on success, error code on failure.
 */
//int av_adts_header_parse(const uint8_t *buf, uint32_t *samples,
//uint8_t *frames);

//未测试
func AvAdtsHeaderParse(buf *ffcommon.FUint8T, samples *ffcommon.FUint64T,
	frames *ffcommon.FUint8T) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_adts_header_parse").Call(
		uintptr(unsafe.Pointer(buf)),
		uintptr(unsafe.Pointer(samples)),
		uintptr(unsafe.Pointer(frames)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}
