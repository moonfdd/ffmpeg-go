package libavutil

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

/**
 * @brief Decodes LZO 1x compressed data.
 * @param out output buffer
 * @param outlen size of output buffer, number of bytes left are returned here
 * @param in input buffer
 * @param inlen size of input buffer, number of bytes left are returned here
 * @return 0 on success, otherwise a combination of the error flags above
 *
 * Make sure all buffers are appropriately padded, in must provide
 * AV_LZO_INPUT_PADDING, out must provide AV_LZO_OUTPUT_PADDING additional bytes.
 */
//int av_lzo1x_decode(void *out, int *outlen, const void *in, int *inlen);
//未测试
func AvLzo1xDecode(out ffcommon.FVoidP, outlen *ffcommon.FInt, in ffcommon.FVoidP, inlen *ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_lzo1x_decode").Call(
		out,
		uintptr(unsafe.Pointer(outlen)),
		in,
		uintptr(unsafe.Pointer(inlen)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}
