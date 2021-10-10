package libavutil

import (
	"ffmpeg-go/ffcommon"
	"syscall"
	"unsafe"
)

/**
 * @file
 * Misc file utilities.
 */

/**
 * Read the file with name filename, and put its content in a newly
 * allocated buffer or map it with mmap() when available.
 * In case of success set *bufptr to the read or mmapped buffer, and
 * *size to the size in bytes of the buffer in *bufptr.
 * Unlike mmap this function succeeds with zero sized files, in this
 * case *bufptr will be set to NULL and *size will be set to 0.
 * The returned buffer must be released with av_file_unmap().
 *
 * @param log_offset loglevel offset used for logging
 * @param log_ctx context used for logging
 * @return a non negative number in case of success, a negative value
 * corresponding to an AVERROR error code in case of failure
 */
//av_warn_unused_result
//int av_file_map(const char *filename, uint8_t **bufptr, size_t *size,
//int log_offset, void *log_ctx);
//未测试
func AvFileMap(filename ffcommon.FConstCharP, bufptr **ffcommon.FUint8T, size *ffcommon.FSizeT,
	log_offset ffcommon.FInt, log_ctx ffcommon.FVoidP) (res ffcommon.FInt, err error) {
	var t uintptr
	var filenamep *byte
	filenamep, err = syscall.BytePtrFromString(filename)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_file_map").Call(
		uintptr(unsafe.Pointer(filenamep)),
		uintptr(unsafe.Pointer(&bufptr)),
		uintptr(unsafe.Pointer(size)),
		uintptr(log_offset),
		log_ctx,
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Unmap or free the buffer bufptr created by av_file_map().
 *
 * @param size size in bytes of bufptr, must be the same as returned
 * by av_file_map()
 */
//void av_file_unmap(uint8_t *bufptr, size_t size);
//未测试
func AvFileUnmap(bufptr *ffcommon.FUint8T, size ffcommon.FSizeT) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_file_unmap").Call(
		uintptr(unsafe.Pointer(bufptr)),
		uintptr(size),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Wrapper to work around the lack of mkstemp() on mingw.
 * Also, tries to create file in /tmp first, if possible.
 * *prefix can be a character ffconstant; *filename will be allocated internally.
 * @return file descriptor of opened file (or negative value corresponding to an
 * AVERROR code on error)
 * and opened file name in **filename.
 * @note On very old libcs it is necessary to set a secure umask before
 *       calling this, av_tempfile() can't call umask itself as it is used in
 *       libraries and could interfere with the calling application.
 * @deprecated as fd numbers cannot be passed saftely between libs on some platforms
 */
//int av_tempfile(const char *prefix, char **filename, int log_offset, void *log_ctx);
//未测试
func av_tempfile(prefix ffcommon.FConstCharP, filename **ffcommon.FUint8T, log_offset ffcommon.FInt, log_ctx ffcommon.FVoidP) (res ffcommon.FInt, err error) {
	var t uintptr
	var prefixp *byte
	prefixp, err = syscall.BytePtrFromString(prefix)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_tempfile").Call(
		uintptr(unsafe.Pointer(prefixp)),
		uintptr(unsafe.Pointer(&filename)),
		uintptr(log_offset),
		log_ctx,
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}
