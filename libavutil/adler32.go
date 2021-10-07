package libavutil

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

/**
 * Calculate the Adler32 checksum of a buffer.
 *
 * Passing the return value to a subsequent av_adler32_update() call
 * allows the checksum of multiple buffers to be calculated as though
 * they were concatenated.
 *
 * @param adler initial checksum value
 * @param buf   pointer to input buffer
 * @param len   size of input buffer
 * @return      updated checksum
 */
//AVAdler av_adler32_update(AVAdler adler, const uint8_t *buf,
//#if FF_API_CRYPTO_SIZE_T
//unsigned int len) av_pure;
//#else
//size_t len) av_pure;
//#endif
func AvAdler32Update(adler ffcommon.FAVAdler, buf *ffcommon.FUint8T,
	len0 ffcommon.FUnsignedInt) (res ffcommon.FAVAdler, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_adler32_update").Call(
		uintptr(adler),
		uintptr(unsafe.Pointer(buf)),
		uintptr(len0),
	)
	res = ffcommon.FAVAdler(t)
	return
}
