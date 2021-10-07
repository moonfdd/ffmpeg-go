package libavutil

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"fmt"
	"syscall"
	"unsafe"
)

/**
 * @file
 * @ingroup lavu
 * Convenience header that includes @ref lavu "libavutil"'s core.
 */

/**
 * @mainpage
 *
 * @section ffmpeg_intro Introduction
 *
 * This document describes the usage of the different libraries
 * provided by FFmpeg.
 *
 * @li @ref libavc "libavcodec" encoding/decoding library
 * @li @ref lavfi "libavfilter" graph-based frame editing library
 * @li @ref libavf "libavformat" I/O and muxing/demuxing library
 * @li @ref lavd "libavdevice" special devices muxing/demuxing library
 * @li @ref lavu "libavutil" common utility library
 * @li @ref lswr "libswresample" audio resampling, format conversion and mixing
 * @li @ref lpp  "libpostproc" post processing library
 * @li @ref libsws "libswscale" color conversion and scaling library
 *
 * @section ffmpeg_versioning Versioning and compatibility
 *
 * Each of the FFmpeg libraries contains a version.h header, which defines a
 * major, minor and micro version number with the
 * <em>LIBRARYNAME_VERSION_{MAJOR,MINOR,MICRO}</em> macros. The major version
 * number is incremented with backward incompatible changes - e.g. removing
 * parts of the public API, reordering public struct members, etc. The minor
 * version number is incremented for backward compatible API changes or major
 * new features - e.g. adding a new public function or a new decoder. The micro
 * version number is incremented for smaller changes that a calling program
 * might still want to check for - e.g. changing behavior in a previously
 * unspecified situation.
 *
 * FFmpeg guarantees backward API and ABI compatibility for each library as long
 * as its major version number is unchanged. This means that no public symbols
 * will be removed or renamed. Types and names of the public struct members and
 * values of public macros and enums will remain the same (unless they were
 * explicitly declared as not part of the public API). Documented behavior will
 * not change.
 *
 * In other words, any correct program that works with a given FFmpeg snapshot
 * should work just as well without any changes with any later snapshot with the
 * same major versions. This applies to both rebuilding the program against new
 * FFmpeg versions or to replacing the dynamic FFmpeg libraries that a program
 * links against.
 *
 * However, new public symbols may be added and new members may be appended to
 * public structs whose size is not part of public ABI (most public structs in
 * FFmpeg). New macros and enum values may be added. Behavior in undocumented
 * situations may change slightly (and be documented). All those are accompanied
 * by an entry in doc/APIchanges and incrementing either the minor or micro
 * version number.
 */

/**
 * @defgroup lavu libavutil
 * Common code shared across all FFmpeg libraries.
 *
 * @note
 * libavutil is designed to be modular. In most cases, in order to use the
 * functions provided by one component of libavutil you must explicitly include
 * the specific header containing that feature. If you are only using
 * media-related components, you could simply include libavutil/avutil.h, which
 * brings in most of the "core" components.
 *
 * @{
 *
 * @defgroup lavu_crypto Crypto and Hashing
 *
 * @{
 * @}
 *
 * @defgroup lavu_math Mathematics
 * @{
 *
 * @}
 *
 * @defgroup lavu_string String Manipulation
 *
 * @{
 *
 * @}
 *
 * @defgroup lavu_mem Memory Management
 *
 * @{
 *
 * @}
 *
 * @defgroup lavu_data Data Structures
 * @{
 *
 * @}
 *
 * @defgroup lavu_video Video related
 *
 * @{
 *
 * @}
 *
 * @defgroup lavu_audio Audio related
 *
 * @{
 *
 * @}
 *
 * @defgroup lavu_error Error Codes
 *
 * @{
 *
 * @}
 *
 * @defgroup lavu_log Logging Facility
 *
 * @{
 *
 * @}
 *
 * @defgroup lavu_misc Other
 *
 * @{
 *
 * @defgroup preproc_misc Preprocessor String Macros
 *
 * @{
 *
 * @}
 *
 * @defgroup version_utils Library Version Macros
 *
 * @{
 *
 * @}
 */

/**
 * @addtogroup lavu_ver
 * @{
 */

/**
 * Return the LIBAVUTIL_VERSION_INT ffconstant.
 */
func AvutilVersion() (res ffcommon.FUint, err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("avutil_version").Call()
	res = ffcommon.FUint(t)
	return
}

/**
 * Return an informative version string. This usually is the actual release
 * version number or a git commit description. This string has no fixed format
 * and can change any time. It should never be parsed by code.
 */
//const char *av_version_info(void);
func AvVersionInfo() (res ffcommon.FConstCharP, err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_version_info").Call()
	res = ffcommon.GoAStr(t)
	return
}

/**
* Return the libavutil build-time configuration.
 */
//const char *avutil_configuration(void);
func AvutilConfiguration() (res ffcommon.FConstCharP, err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("avutil_configuration").Call()
	res = ffcommon.GoAStr(t)
	return
}

/**
* Return the libavutil license.
 */
//const char *avutil_license(void);
func AvutilLicense() (res ffcommon.FConstCharP, err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("avutil_license").Call()
	res = ffcommon.GoAStr(t)
	return
}

/**
* Return a string describing the media_type enum, NULL if media_type
* is unknown.
 */
//const char *av_get_media_type_string(enum AVMediaType media_type);
func AvGetMediaTypeString(media_type ffconstant.AVMediaType) (res ffcommon.FConstCharP, err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_get_media_type_string").Call(uintptr(media_type))
	res = ffcommon.GoAStr(t)
	return
}

/**
* Return a single letter to describe the given picture type
* pict_type.
*
* @param[in] pict_type the picture type @return a single character
* representing the picture type, '?' if pict_type is unknown
 */
//char av_get_picture_type_char(enum AVPictureType pict_type);
func AvGetPictureTypeChar(pict_type ffconstant.AVPictureType) (res ffcommon.FConstCharP, err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_get_picture_type_char").Call(uintptr(pict_type))
	res = ffcommon.GoAStr(t)
	return
}

/**
* Compute the length of an integer list.
*
* @param elsize  size in bytes of each list element (only 1, 2, 4 or 8)
* @param term    list terminator (usually 0 or -1)
* @param list    pointer to the list
* @return  length of the list, in elements, not counting the terminator
 */
//unsigned av_int_list_length_for_size(unsigned elsize,
//const void *list, uint64_t term) av_pure;
func AvIntListLengthForSize(elsize ffcommon.FUnsigned,
	list ffcommon.FConstVoidP, term ffcommon.FUint64T) (res ffcommon.FUnsigned, err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_int_list_length_for_size").Call(uintptr(elsize), uintptr(list), uintptr(term))
	res = ffcommon.FUnsigned(t)
	return
}

/**
* Open a file using a UTF-8 filename.
* The API of this function matches POSIX fopen(), errors are returned through
* errno.
 */
//FILE *av_fopen_utf8(const char *path, const char *mode);
func AvFopenUtf8(path0 ffcommon.FConstCharP, mode ffcommon.FConstCharP) (res ffcommon.FFileP, err error) {

	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_fopen_utf8").Call(
		uintptr(unsafe.Pointer(&UTF8ToANSI(path0)[0])),
		uintptr(unsafe.Pointer(&UTF8ToANSI(mode)[0])),
	)
	//t, _, err := ffcommon.GetAvutilDll().NewProc("av_fopen_utf8").Call(
	//	uintptr(unsafe.Pointer(&path0)),
	//	uintptr(unsafe.Pointer(&mode)),
	//)
	fmt.Println("err = ", err)
	res = ffcommon.FFileP(t)
	return
}

const CP_UTF8 = 65001
const CP_ACP = 0

// UTF8ToANSI 将UTF-8字符转为ANSI格式
func UTF8ToANSI(str string) []uint8 {
	if str == "" {
		return nil
	}
	utf8StrPtr := uintptr(unsafe.Pointer(&([]byte(str + "\x00")[0])))
	nLen := MultiByteToWideChar(CP_UTF8, 0, utf8StrPtr, -1, 0, 0)
	wCharBuffer := make([]uint16, nLen+1)
	wCharBufferPtr := uintptr(unsafe.Pointer(&wCharBuffer[0]))
	MultiByteToWideChar(CP_UTF8, 0, utf8StrPtr, -1, wCharBufferPtr, nLen)

	nLen = WideCharToMultiByte(CP_ACP, 0, wCharBufferPtr, -1, 0, 0, 0, nil)
	aCharBuffer := make([]uint8, nLen) // +1
	aCharBufferPtr := uintptr(unsafe.Pointer(&aCharBuffer[0]))
	WideCharToMultiByte(CP_ACP, 0, wCharBufferPtr, -1, aCharBufferPtr, nLen, 0, nil)

	return aCharBuffer
}

var kernel32dll = syscall.NewLazyDLL("kernel32.dll")
var _MultiByteToWideChar = kernel32dll.NewProc("MultiByteToWideChar")

// MultiByteToWideChar
func MultiByteToWideChar(CodePage, dwFlags uint32, lpMultiByteStr uintptr, cchMultiByte int, lpWideCharStr uintptr, cchWideChar int) int {
	r, _, _ := _MultiByteToWideChar.Call(uintptr(CodePage), uintptr(dwFlags), lpMultiByteStr, uintptr(cchMultiByte), lpWideCharStr, uintptr(cchWideChar))
	return int(r)
}

// WideCharToMultiByte
func WideCharToMultiByte(CodePage, dwFlags uint32, lpWideCharStr uintptr, cchWideChar int, lpMultiByteStr uintptr, cchMultiByte int, lpDefaultChar uintptr,
	lpUsedDefaultChar *int32) int {
	r, _, _ := _WideCharToMultiByte.Call(uintptr(CodePage), uintptr(dwFlags), lpWideCharStr, uintptr(cchWideChar), lpMultiByteStr, uintptr(cchMultiByte), lpDefaultChar,
		uintptr(unsafe.Pointer(lpUsedDefaultChar)))
	return int(r)
}

var _WideCharToMultiByte = kernel32dll.NewProc("WideCharToMultiByte")

/**
* Return the fractional representation of the internal time base.
 */
//AVRational av_get_time_base_q(void);
func AvGetTimeBaseQ() (res AVRational, err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_get_time_base_q").Call()
	res = *(*AVRational)(unsafe.Pointer(&t))
	return
}

/**
* Fill the provided buffer with a string containing a FourCC (four-character
* code) representation.
*
* @param buf    a buffer with size in bytes of at least AV_FOURCC_MAX_STRING_SIZE
* @param fourcc the fourcc to represent
* @return the buffer in input
 */
//char *av_fourcc_make_string(char *buf, uint32_t fourcc);
//测试失败
func AvFourccMakeString(buf ffcommon.FBuf, fourcc ffcommon.FUint32T) (res ffcommon.FCharP, err error) {
	//t, _, _ := ffcommon.GetAvutilDll().NewProc("av_fourcc_make_string").Call(uintptr(unsafe.Pointer(&buf[0])), uintptr(fourcc))
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_fourcc_make_string").Call(uintptr(unsafe.Pointer(uintptr(unsafe.Pointer(buf)))), uintptr(fourcc))
	res = ffcommon.GoAStr(t)
	return
}
