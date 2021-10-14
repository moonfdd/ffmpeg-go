package libavutil

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

/**
 * @defgroup lavu_base64 Base64
 * @ingroup lavu_crypto
 * @{
 */

/**
 * Decode a base64-encoded string.
 *
 * @param out      buffer for decoded data
 * @param in       null-terminated input string
 * @param out_size size in bytes of the out buffer, must be at
 *                 least 3/4 of the length of in, that is AV_BASE64_DECODE_SIZE(strlen(in))
 * @return         number of bytes written, or a negative value in case of
 *                 invalid input
 */
//int av_base64_decode(uint8_t *out, const char *in, int out_size);
//未测试
func AvBase64Decode(out *ffcommon.FUint8T, in *ffcommon.FUint8T, out_size ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_base64_decode").Call(
		uintptr(unsafe.Pointer(out)),
		uintptr(unsafe.Pointer(in)),
		uintptr(out_size),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Encode data to base64 and null-terminate.
 *
 * @param out      buffer for encoded data
 * @param out_size size in bytes of the out buffer (including the
 *                 null terminator), must be at least AV_BASE64_SIZE(in_size)
 * @param in       input buffer containing the data to encode
 * @param in_size  size in bytes of the in buffer
 * @return         out or NULL in case of error
 */
//char *av_base64_encode(char *out, int out_size, const uint8_t *in, int in_size);
//未测试
func AvBase64Encode(out *ffcommon.FUint8T, out_size ffcommon.FInt, in *ffcommon.FUint8T, in_size ffcommon.FInt) (res ffcommon.FCharP, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_base64_encode").Call(
		uintptr(unsafe.Pointer(out)),
		uintptr(out_size),
		uintptr(unsafe.Pointer(in)),
		uintptr(in_size),
	)
	if err != nil {
		//return
	}
	res = ffcommon.GoAStr(t)
	return
}
