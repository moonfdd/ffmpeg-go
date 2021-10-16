package libavutil

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"syscall"
	"unsafe"
)

/**
 * @addtogroup lavu_string
 * @{
 */

/**
 * Return non-zero if pfx is a prefix of str. If it is, *ptr is set to
 * the address of the first character in str after the prefix.
 *
 * @param str input string
 * @param pfx prefix to test
 * @param ptr updated if the prefix is matched inside str
 * @return non-zero if the prefix matches, zero otherwise
 */
//int av_strstart(const char *str, const char *pfx, const char **ptr);
//未测试
func AvStrstart(str ffcommon.FConstCharP, pfx ffcommon.FConstCharP, ptr *ffcommon.FBuf) (res ffcommon.FInt, err error) {
	var t uintptr
	var strp *byte
	strp, err = syscall.BytePtrFromString(str)
	if err != nil {
		return
	}
	var pfxp *byte
	pfxp, err = syscall.BytePtrFromString(pfx)
	if err != nil {
		return
	}
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_strstart").Call(
		uintptr(unsafe.Pointer(strp)),
		uintptr(unsafe.Pointer(pfxp)),
		uintptr(unsafe.Pointer(ptr)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Return non-zero if pfx is a prefix of str independent of case. If
 * it is, *ptr is set to the address of the first character in str
 * after the prefix.
 *
 * @param str input string
 * @param pfx prefix to test
 * @param ptr updated if the prefix is matched inside str
 * @return non-zero if the prefix matches, zero otherwise
 */
//int av_stristart(const char *str, const char *pfx, const char **ptr);
//未测试
func AvStristart(str ffcommon.FConstCharP, pfx ffcommon.FConstCharP, ptr *ffcommon.FBuf) (res ffcommon.FInt, err error) {
	var t uintptr
	var strp *byte
	strp, err = syscall.BytePtrFromString(str)
	if err != nil {
		return
	}
	var pfxp *byte
	pfxp, err = syscall.BytePtrFromString(pfx)
	if err != nil {
		return
	}
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_stristart").Call(
		uintptr(unsafe.Pointer(strp)),
		uintptr(unsafe.Pointer(pfxp)),
		uintptr(unsafe.Pointer(ptr)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Locate the first case-independent occurrence in the string haystack
 * of the string needle.  A zero-length string needle is considered to
 * match at the start of haystack.
 *
 * This function is a case-insensitive version of the standard strstr().
 *
 * @param haystack string to search in
 * @param needle   string to search for
 * @return         pointer to the located match within haystack
 *                 or a null pointer if no match
 */
//char *av_stristr(const char *haystack, const char *needle);
//未测试
func AvStristr(haystack ffcommon.FConstCharP, needle ffcommon.FConstCharP) (res ffcommon.FCharP, err error) {
	var t uintptr
	var haystackp *byte
	haystackp, err = syscall.BytePtrFromString(haystack)
	if err != nil {
		return
	}
	var needlep *byte
	needlep, err = syscall.BytePtrFromString(needle)
	if err != nil {
		return
	}
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_stristr").Call(
		uintptr(unsafe.Pointer(haystackp)),
		uintptr(unsafe.Pointer(needlep)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.GoAStr(t)
	return
}

/**
 * Locate the first occurrence of the string needle in the string haystack
 * where not more than hay_length characters are searched. A zero-length
 * string needle is considered to match at the start of haystack.
 *
 * This function is a length-limited version of the standard strstr().
 *
 * @param haystack   string to search in
 * @param needle     string to search for
 * @param hay_length length of string to search in
 * @return           pointer to the located match within haystack
 *                   or a null pointer if no match
 */
//char *av_strnstr(const char *haystack, const char *needle, size_t hay_length);
//未测试
func AvStrnstr(haystack ffcommon.FConstCharP, needle ffcommon.FConstCharP, hay_length ffcommon.FSizeT) (res ffcommon.FCharP, err error) {
	var t uintptr
	var haystackp *byte
	haystackp, err = syscall.BytePtrFromString(haystack)
	if err != nil {
		return
	}
	var needlep *byte
	needlep, err = syscall.BytePtrFromString(needle)
	if err != nil {
		return
	}
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_strnstr").Call(
		uintptr(unsafe.Pointer(haystackp)),
		uintptr(unsafe.Pointer(needlep)),
		uintptr(hay_length),
	)
	if err != nil {
		//return
	}
	res = ffcommon.GoAStr(t)
	return
}

/**
 * Copy the string src to dst, but no more than size - 1 bytes, and
 * null-terminate dst.
 *
 * This function is the same as BSD strlcpy().
 *
 * @param dst destination buffer
 * @param src source string
 * @param size size of destination buffer
 * @return the length of src
 *
 * @warning since the return value is the length of src, src absolutely
 * _must_ be a properly 0-terminated string, otherwise this will read beyond
 * the end of the buffer and possibly crash.
 */
//size_t av_strlcpy(char *dst, const char *src, size_t size);
//未测试
func AvStrlcpy(dst ffcommon.FBuf, src ffcommon.FBuf, size ffcommon.FSizeT) (res ffcommon.FSizeT, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_strlcpy").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
		uintptr(size),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FSizeT(t)
	return
}

/**
 * Append the string src to the string dst, but to a total length of
 * no more than size - 1 bytes, and null-terminate dst.
 *
 * This function is similar to BSD strlcat(), but differs when
 * size <= strlen(dst).
 *
 * @param dst destination buffer
 * @param src source string
 * @param size size of destination buffer
 * @return the total length of src and dst
 *
 * @warning since the return value use the length of src and dst, these
 * absolutely _must_ be a properly 0-terminated strings, otherwise this
 * will read beyond the end of the buffer and possibly crash.
 */
//size_t av_strlcat(char *dst, const char *src, size_t size);
//未测试
func AvStrlcat(dst ffcommon.FBuf, src ffcommon.FBuf, size ffcommon.FSizeT) (res ffcommon.FSizeT, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_strlcat").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
		uintptr(size),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FSizeT(t)
	return
}

/**
 * Append output to a string, according to a format. Never write out of
 * the destination buffer, and always put a terminating 0 within
 * the buffer.
 * @param dst destination buffer (string to which the output is
 *  appended)
 * @param size total size of the destination buffer
 * @param fmt printf-compatible format string, specifying how the
 *  following parameters are used
 * @return the length of the string that would have been generated
 *  if enough space had been available
 */
//size_t av_strlcatf(char *dst, size_t size, const char *fmt, ...) av_printf_format(3, 4);
//未测试
func AvStrlcatf(dst ffcommon.FBuf, size ffcommon.FSizeT, fmt0 ffcommon.FConstCharP, a ...[]uintptr) (res ffcommon.FSizeT, err error) {
	var t uintptr
	var fmt0p *byte
	fmt0p, err = syscall.BytePtrFromString(fmt0)
	if err != nil {
		return
	}
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_strlcatf").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(size),
		uintptr(unsafe.Pointer(fmt0p)),
		uintptr(unsafe.Pointer(&a)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FSizeT(t)
	return
}

/**
 * Print arguments following specified format into a large enough auto
 * allocated buffer. It is similar to GNU asprintf().
 * @param fmt printf-compatible format string, specifying how the
 *            following parameters are used.
 * @return the allocated string
 * @note You have to free the string yourself with av_free().
 */
//char *av_asprintf(const char *fmt, ...) av_printf_format(1, 2);
//未测试
func AvAsprintf(fmt0 ffcommon.FConstCharP) (res ffcommon.FCharP, err error) {
	var t uintptr
	var fmt0p *byte
	fmt0p, err = syscall.BytePtrFromString(fmt0)
	if err != nil {
		return
	}
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_asprintf").Call(
		uintptr(unsafe.Pointer(fmt0p)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.GoAStr(t)
	return
}

//#if FF_API_D2STR
/**
 * Convert a number to an av_malloced string.
 * @deprecated  use av_asprintf() with "%f" or a more specific format
 */
//attribute_deprecated
//char *av_d2str(double d);
//未测试
func AvD2str(d ffcommon.FDouble) (res ffcommon.FCharP, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_d2str").Call(
		uintptr(d),
	)
	if err != nil {
		//return
	}
	res = ffcommon.GoAStr(t)
	return
}

//#endif

/**
 * Unescape the given string until a non escaped terminating char,
 * and return the token corresponding to the unescaped string.
 *
 * The normal \ and ' escaping is supported. Leading and trailing
 * whitespaces are removed, unless they are escaped with '\' or are
 * enclosed between ''.
 *
 * @param buf the buffer to parse, buf will be updated to point to the
 * terminating char
 * @param term a 0-terminated list of terminating chars
 * @return the malloced unescaped string, which must be av_freed by
 * the user, NULL in case of allocation failure
 */
//char *av_get_token(const char **buf, const char *term);
//未测试
func AvGetToken(buf *ffcommon.FBuf, term ffcommon.FConstCharP) (res ffcommon.FCharP, err error) {
	var t uintptr
	var termp *byte
	termp, err = syscall.BytePtrFromString(term)
	if err != nil {
		return
	}
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_get_token").Call(
		uintptr(unsafe.Pointer(&buf)),
		uintptr(unsafe.Pointer(termp)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.GoAStr(t)
	return
}

/**
 * Split the string into several tokens which can be accessed by
 * successive calls to av_strtok().
 *
 * A token is defined as a sequence of characters not belonging to the
 * set specified in delim.
 *
 * On the first call to av_strtok(), s should point to the string to
 * parse, and the value of saveptr is ignored. In subsequent calls, s
 * should be NULL, and saveptr should be unchanged since the previous
 * call.
 *
 * This function is similar to strtok_r() defined in POSIX.1.
 *
 * @param s the string to parse, may be NULL
 * @param delim 0-terminated list of token delimiters, must be non-NULL
 * @param saveptr user-provided pointer which points to stored
 * information necessary for av_strtok() to continue scanning the same
 * string. saveptr is updated to point to the next character after the
 * first delimiter found, or to NULL if the string was terminated
 * @return the found token, or NULL when no token is found
 */
//char *av_strtok(char *s, const char *delim, char **saveptr);
//未测试
func AvStrtok(s ffcommon.FConstCharP, delim ffcommon.FConstCharP, saveptr *ffcommon.FBuf) (res ffcommon.FCharP, err error) {
	var t uintptr
	var sp *byte
	sp, err = syscall.BytePtrFromString(s)
	if err != nil {
		return
	}
	var delimp *byte
	delimp, err = syscall.BytePtrFromString(delim)
	if err != nil {
		return
	}
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_strtok").Call(
		uintptr(unsafe.Pointer(sp)),
		uintptr(unsafe.Pointer(delimp)),
		uintptr(unsafe.Pointer(&saveptr)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.GoAStr(t)
	return
}

/**
 * Locale-independent case-insensitive compare.
 * @note This means only ASCII-range characters are case-insensitive
 */
//int av_strcasecmp(const char *a, const char *b);
//未测试
func AvStrcasecmp(a ffcommon.FConstCharP, b ffcommon.FConstCharP) (res ffcommon.FInt, err error) {
	var t uintptr
	var ap *byte
	ap, err = syscall.BytePtrFromString(a)
	if err != nil {
		return
	}
	var bp *byte
	bp, err = syscall.BytePtrFromString(b)
	if err != nil {
		return
	}
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_strcasecmp").Call(
		uintptr(unsafe.Pointer(ap)),
		uintptr(unsafe.Pointer(bp)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Locale-independent case-insensitive compare.
 * @note This means only ASCII-range characters are case-insensitive
 */
//int av_strncasecmp(const char *a, const char *b, size_t n);
//未测试
func AvStrncasecmp(a ffcommon.FConstCharP, b ffcommon.FConstCharP, n ffcommon.FSizeT) (res ffcommon.FInt, err error) {
	var t uintptr
	var ap *byte
	ap, err = syscall.BytePtrFromString(a)
	if err != nil {
		return
	}
	var bp *byte
	bp, err = syscall.BytePtrFromString(b)
	if err != nil {
		return
	}
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_strncasecmp").Call(
		uintptr(unsafe.Pointer(ap)),
		uintptr(unsafe.Pointer(bp)),
		uintptr(n),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Locale-independent strings replace.
 * @note This means only ASCII-range characters are replace
 */
//char *av_strireplace(const char *str, const char *from, const char *to);
//未测试
func AvStrireplace(str ffcommon.FConstCharP, from ffcommon.FConstCharP, to ffcommon.FConstCharP) (res ffcommon.FCharP, err error) {
	var t uintptr
	var strp *byte
	strp, err = syscall.BytePtrFromString(str)
	if err != nil {
		return
	}
	var fromp *byte
	fromp, err = syscall.BytePtrFromString(from)
	if err != nil {
		return
	}
	var top *byte
	top, err = syscall.BytePtrFromString(to)
	if err != nil {
		return
	}
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_strireplace").Call(
		uintptr(unsafe.Pointer(strp)),
		uintptr(unsafe.Pointer(fromp)),
		uintptr(unsafe.Pointer(top)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.GoAStr(t)
	return
}

/**
 * Thread safe basename.
 * @param path the string to parse, on DOS both \ and / are considered separators.
 * @return pointer to the basename substring.
 * If path does not contain a slash, the function returns a copy of path.
 * If path is a NULL pointer or points to an empty string, a pointer
 * to a string "." is returned.
 */
//const char *av_basename(const char *path);
//未测试
func AvBasename(path0 ffcommon.FConstCharP) (res ffcommon.FCharP, err error) {
	var t uintptr
	var path0p *byte
	path0p, err = syscall.BytePtrFromString(path0)
	if err != nil {
		return
	}
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_basename").Call(
		uintptr(unsafe.Pointer(path0p)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.GoAStr(t)
	return
}

/**
 * Thread safe dirname.
 * @param path the string to parse, on DOS both \ and / are considered separators.
 * @return A pointer to a string that's the parent directory of path.
 * If path is a NULL pointer or points to an empty string, a pointer
 * to a string "." is returned.
 * @note the function may modify the contents of the path, so copies should be passed.
 */
//const char *av_dirname(char *path);
//未测试
func AvDirname(path0 ffcommon.FConstCharP) (res ffcommon.FCharP, err error) {
	var t uintptr
	var path0p *byte
	path0p, err = syscall.BytePtrFromString(path0)
	if err != nil {
		return
	}
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_dirname").Call(
		uintptr(unsafe.Pointer(path0p)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.GoAStr(t)
	return
}

/**
 * Match instances of a name in a comma-separated list of names.
 * List entries are checked from the start to the end of the names list,
 * the first match ends further processing. If an entry prefixed with '-'
 * matches, then 0 is returned. The "ALL" list entry is considered to
 * match all names.
 *
 * @param name  Name to look for.
 * @param names List of names.
 * @return 1 on match, 0 otherwise.
 */
//int av_match_name(const char *name, const char *names);
//未测试
func AvMatchName(name ffcommon.FConstCharP, names ffcommon.FConstCharP) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	var namesp *byte
	namesp, err = syscall.BytePtrFromString(names)
	if err != nil {
		return
	}
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_match_name").Call(
		uintptr(unsafe.Pointer(namep)),
		uintptr(unsafe.Pointer(namesp)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Append path component to the existing path.
 * Path separator '/' is placed between when needed.
 * Resulting string have to be freed with av_free().
 * @param path      base path
 * @param component component to be appended
 * @return new path or NULL on error.
 */
//char *av_append_path_component(const char *path, const char *component);
//未测试
func AvAppendPathComponent(name ffcommon.FConstCharP, names ffcommon.FConstCharP) (res ffcommon.FCharP, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	var namesp *byte
	namesp, err = syscall.BytePtrFromString(names)
	if err != nil {
		return
	}
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_append_path_component").Call(
		uintptr(unsafe.Pointer(namep)),
		uintptr(unsafe.Pointer(namesp)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.GoAStr(t)
	return
}

/**
 * Escape string in src, and put the escaped string in an allocated
 * string in *dst, which must be freed with av_free().
 *
 * @param dst           pointer where an allocated string is put
 * @param src           string to escape, must be non-NULL
 * @param special_chars string containing the special characters which
 *                      need to be escaped, can be NULL
 * @param mode          escape mode to employ, see AV_ESCAPE_MODE_* macros.
 *                      Any unknown value for mode will be considered equivalent to
 *                      AV_ESCAPE_MODE_BACKSLASH, but this behaviour can change without
 *                      notice.
 * @param flags         flags which control how to escape, see AV_ESCAPE_FLAG_ macros
 * @return the length of the allocated string, or a negative error code in case of error
 * @see av_bprint_escape()
 */
//av_warn_unused_result
//int av_escape(char **dst, const char *src, const char *special_chars,
//enum AVEscapeMode mode, int flags);
//未测试
func AvEscape(dst *ffcommon.FBuf, src ffcommon.FConstCharP, special_chars ffcommon.FConstCharP,
	mode ffconstant.AVEscapeMode, flags ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	var srcp *byte
	srcp, err = syscall.BytePtrFromString(src)
	if err != nil {
		return
	}
	var special_charsp *byte
	special_charsp, err = syscall.BytePtrFromString(special_chars)
	if err != nil {
		return
	}
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_escape").Call(
		uintptr(unsafe.Pointer(&dst)),
		uintptr(unsafe.Pointer(srcp)),
		uintptr(unsafe.Pointer(special_charsp)),
		uintptr(mode),
		uintptr(flags),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Read and decode a single UTF-8 code point (character) from the
 * buffer in *buf, and update *buf to point to the next byte to
 * decode.
 *
 * In case of an invalid byte sequence, the pointer will be updated to
 * the next byte after the invalid sequence and the function will
 * return an error code.
 *
 * Depending on the specified flags, the function will also fail in
 * case the decoded code point does not belong to a valid range.
 *
 * @note For speed-relevant code a carefully implemented use of
 * GET_UTF8() may be preferred.
 *
 * @param codep   pointer used to return the parsed code in case of success.
 *                The value in *codep is set even in case the range check fails.
 * @param bufp    pointer to the address the first byte of the sequence
 *                to decode, updated by the function to point to the
 *                byte next after the decoded sequence
 * @param buf_end pointer to the end of the buffer, points to the next
 *                byte past the last in the buffer. This is used to
 *                avoid buffer overreads (in case of an unfinished
 *                UTF-8 sequence towards the end of the buffer).
 * @param flags   a collection of AV_UTF8_FLAG_* flags
 * @return >= 0 in case a sequence was successfully read, a negative
 * value in case of invalid sequence
 */
//av_warn_unused_result
//int av_utf8_decode(int32_t *codep, const uint8_t **bufp, const uint8_t *buf_end,
//unsigned int flags);
//未测试
func AvUtf8Decode(codep *ffcommon.FInt32T, bufp **ffcommon.FUint8T, buf_end *ffcommon.FUint8T,
	flags ffcommon.FUnsignedInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_utf8_decode").Call(
		uintptr(unsafe.Pointer(codep)),
		uintptr(unsafe.Pointer(&bufp)),
		uintptr(unsafe.Pointer(buf_end)),
		uintptr(flags),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Check if a name is in a list.
 * @returns 0 if not found, or the 1 based index where it has been found in the
 *            list.
 */
//int av_match_list(const char *name, const char *list, char separator);
//未测试
func AvMatchList(name ffcommon.FConstCharP, list0 ffcommon.FConstCharP, separator ffcommon.FUint8T) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	var list0p *byte
	list0p, err = syscall.BytePtrFromString(list0)
	if err != nil {
		return
	}
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_match_list").Call(
		uintptr(unsafe.Pointer(namep)),
		uintptr(unsafe.Pointer(list0p)),
		uintptr(separator),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * See libc sscanf manual for more information.
 * Locale-independent sscanf implementation.
 */
//int av_sscanf(const char *string, const char *format, ...);
//未测试
func AvSscanf(string0 ffcommon.FConstCharP, format0 ffcommon.FConstCharP, a ...[]uintptr) (res ffcommon.FInt, err error) {
	var t uintptr
	var string0p *byte
	string0p, err = syscall.BytePtrFromString(string0)
	if err != nil {
		return
	}
	var format0p *byte
	format0p, err = syscall.BytePtrFromString(format0)
	if err != nil {
		return
	}
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_sscanf").Call(
		uintptr(unsafe.Pointer(string0p)),
		uintptr(unsafe.Pointer(format0p)),
		uintptr(unsafe.Pointer(&a)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}
