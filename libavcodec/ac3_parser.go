package libavcodec

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

/**
 * Extract the bitstream ID and the frame size from AC-3 data.
 */
//int av_ac3_parse_header(const uint8_t *buf, size_t size,
//uint8_t *bitstream_id, uint16_t *frame_size);
//未测试
func AvAc3ParseHeader(buf *ffcommon.FUint8T, size ffcommon.FSizeT,
	bitstream_id *ffcommon.FUint8T, frame_size *ffcommon.FUint16T) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_ac3_parse_header").Call(
		uintptr(unsafe.Pointer(buf)),
		uintptr(size),
		uintptr(unsafe.Pointer(bitstream_id)),
		uintptr(unsafe.Pointer(frame_size)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}
