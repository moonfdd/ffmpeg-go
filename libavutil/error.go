package libavutil

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

/**
* Put a description of the AVERROR code errnum in errbuf.
* In case of failure the global variable errno is set to indicate the
* error. Even in case of failure av_strerror() will print a generic
* error message indicating the errnum provided to errbuf.
*
* @param errnum      error code to describe
* @param errbuf      buffer to which description is written
* @param errbuf_size the size in bytes of errbuf
* @return 0 on success, a negative value if a description for errnum
* cannot be found
 */
//int av_strerror(int errnum, char *errbuf, size_t errbuf_size);
//未测试
func AvStrerror(errnum ffcommon.FInt, errbuf ffcommon.FBuf, errbuf_size ffcommon.FSizeT) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_strerror").Call(
		uintptr(errnum),
		uintptr(unsafe.Pointer(errbuf)),
		uintptr(errbuf_size),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}
