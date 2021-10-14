package libavutil

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"syscall"
	"unsafe"
)

/**
 * Buffer to print data progressively
 *
 * The string buffer grows as necessary and is always 0-terminated.
 * The content of the string is never accessed, and thus is
 * encoding-agnostic and can even hold binary data.
 *
 * Small buffers are kept in the structure itself, and thus require no
 * memory allocation at all (unless the contents of the buffer is needed
 * after the structure goes out of scope). This is almost as lightweight as
 * declaring a local "char buf[512]".
 *
 * The length of the string can go beyond the allocated size: the buffer is
 * then truncated, but the functions still keep account of the actual total
 * length.
 *
 * In other words, buf->len can be greater than buf->size and records the
 * total length of what would have been to the buffer if there had been
 * enough memory.
 *
 * Append operations do not need to be tested for failure: if a memory
 * allocation fails, data stop being appended to the buffer, but the length
 * is still updated. This situation can be tested with
 * av_bprint_is_complete().
 *
 * The size_max field determines several possible behaviours:
 *
 * size_max = -1 (= UINT_MAX) or any large value will let the buffer be
 * reallocated as necessary, with an amortized linear cost.
 *
 * size_max = 0 prevents writing anything to the buffer: only the total
 * length is computed. The write operations can then possibly be repeated in
 * a buffer with exactly the necessary size
 * (using size_init = size_max = len + 1).
 *
 * size_max = 1 is automatically replaced by the exact size available in the
 * structure itself, thus ensuring no dynamic memory allocation. The
 * internal buffer is large enough to hold a reasonable paragraph of text,
 * such as the current paragraph.
 */

//FF_PAD_STRUCTURE(AVBPrint, 1024,
//char *str;         /**< string so far */
//unsigned len;      /**< length so far */
//unsigned size;     /**< allocated memory */
//unsigned size_max; /**< maximum allocated memory */
//char reserved_internal_buffer[1];
//)
type AVBPrint struct {
	str                      ffcommon.FBuf      /**< string so far */
	Len                      ffcommon.FUnsigned /**< length so far */
	size                     ffcommon.FUnsigned /**< allocated memory */
	size_max                 ffcommon.FUnsigned /**< maximum allocated memory */
	reserved_internal_buffer [1]byte
}

/**
 * Init a print buffer.
 *
 * @param buf        buffer to init
 * @param size_init  initial size (including the final 0)
 * @param size_max   maximum size;
 *                   0 means do not write anything, just count the length;
 *                   1 is replaced by the maximum value for automatic storage;
 *                   any large value means that the internal buffer will be
 *                   reallocated as needed up to that limit; -1 is converted to
 *                   UINT_MAX, the largest limit possible.
 *                   Check also AV_BPRINT_SIZE_* macros.
 */
//void av_bprint_init(AVBPrint *buf, unsigned size_init, unsigned size_max);
//未测试
func (buf *AVBPrint) AvBprintInit(size_init ffcommon.FUnsigned, size_max ffcommon.FUnsigned) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_bprint_init").Call(
		uintptr(unsafe.Pointer(buf)),
		uintptr(size_init),
		uintptr(size_max),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Init a print buffer using a pre-existing buffer.
 *
 * The buffer will not be reallocated.
 *
 * @param buf     buffer structure to init
 * @param buffer  byte buffer to use for the string data
 * @param size    size of buffer
 */
//void av_bprint_init_for_buffer(AVBPrint *buf, char *buffer, unsigned size);
//未测试
func (buf *AVBPrint) AvBprintInitForBuffer(buffer ffcommon.FBuf, size ffcommon.FUnsigned) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_bprint_init_for_buffer").Call(
		uintptr(unsafe.Pointer(buf)),
		uintptr(unsafe.Pointer(buffer)),
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
 * Append a formatted string to a print buffer.
 */
//void av_bprintf(AVBPrint *buf, const char *fmt, ...) av_printf_format(2, 3);
//未测试
func (buf *AVBPrint) AvBprintf(fmt0 ffcommon.FConstCharP, a []ffcommon.FVoidP) (err error) {
	var t uintptr
	var fmt0p *byte
	fmt0p, err = syscall.BytePtrFromString(fmt0)
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_bprintf").Call(
		uintptr(unsafe.Pointer(buf)),
		uintptr(unsafe.Pointer(fmt0p)),
		uintptr(unsafe.Pointer(&a)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Append a formatted string to a print buffer.
 */
//void av_vbprintf(AVBPrint *buf, const char *fmt, va_list vl_arg);
//未测试
func (buf *AVBPrint) AvVbprintf(fmt0 ffcommon.FConstCharP, vl_arg ffcommon.FVaList) (err error) {
	var t uintptr
	var fmt0p *byte
	fmt0p, err = syscall.BytePtrFromString(fmt0)
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_vbprintf").Call(
		uintptr(unsafe.Pointer(buf)),
		uintptr(unsafe.Pointer(fmt0p)),
		uintptr(unsafe.Pointer(vl_arg)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Append char c n times to a print buffer.
 */
//void av_bprint_chars(AVBPrint *buf, char c, unsigned n);
//未测试
func (buf *AVBPrint) AvBprintChars(c byte, n ffcommon.FUnsigned) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_bprint_chars").Call(
		uintptr(unsafe.Pointer(buf)),
		uintptr(c),
		uintptr(n),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Append data to a print buffer.
 *
 * param buf  bprint buffer to use
 * param data pointer to data
 * param size size of data
 */
//void av_bprint_append_data(AVBPrint *buf, const char *data, unsigned size);
//未测试
func (buf *AVBPrint) AvBprintAppendData(data ffcommon.FBuf, size ffcommon.FUnsigned) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_bprint_append_data").Call(
		uintptr(unsafe.Pointer(buf)),
		uintptr(unsafe.Pointer(data)),
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
 * Append a formatted date and time to a print buffer.
 *
 * param buf  bprint buffer to use
 * param fmt  date and time format string, see strftime()
 * param tm   broken-down time structure to translate
 *
 * @note due to poor design of the standard strftime function, it may
 * produce poor results if the format string expands to a very long text and
 * the bprint buffer is near the limit stated by the size_max option.
 */
//void av_bprint_strftime(AVBPrint *buf, const char *fmt, const struct tm *tm);
//未测试
func (buf *AVBPrint) AvBprintStrftime(fmt0 ffcommon.FConstCharP, tm *Tm) (err error) {
	var t uintptr
	var fmt0p *byte
	fmt0p, err = syscall.BytePtrFromString(fmt0)
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_bprint_strftime").Call(
		uintptr(unsafe.Pointer(buf)),
		uintptr(unsafe.Pointer(fmt0p)),
		uintptr(unsafe.Pointer(tm)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Allocate bytes in the buffer for external use.
 *
 * @param[in]  buf          buffer structure
 * @param[in]  size         required size
 * @param[out] mem          pointer to the memory area
 * @param[out] actual_size  size of the memory area after allocation;
 *                          can be larger or smaller than size
 */
//void av_bprint_get_buffer(AVBPrint *buf, unsigned size,
//unsigned char **mem, unsigned *actual_size);
//未测试
func (buf *AVBPrint) AvBprintGetBuffer(size ffcommon.FUnsigned,
	mem *ffcommon.FBuf, actual_size *ffcommon.FUnsigned) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_bprint_get_buffer").Call(
		uintptr(unsafe.Pointer(buf)),
		uintptr(size),
		uintptr(unsafe.Pointer(&mem)),
		uintptr(unsafe.Pointer(actual_size)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Reset the string to "" but keep internal allocated data.
 */
//void av_bprint_clear(AVBPrint *buf);
//未测试
func (buf *AVBPrint) AvBprintClear() (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_bprint_clear").Call(
		uintptr(unsafe.Pointer(buf)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Finalize a print buffer.
 *
 * The print buffer can no longer be used afterwards,
 * but the len and size fields are still valid.
 *
 * @arg[out] ret_str  if not NULL, used to return a permanent copy of the
 *                    buffer contents, or NULL if memory allocation fails;
 *                    if NULL, the buffer is discarded and freed
 * @return  0 for success or error code (probably AVERROR(ENOMEM))
 */
//int av_bprint_finalize(AVBPrint *buf, char **ret_str);
//未测试
func (buf *AVBPrint) AvBprintFinalize(ret_str *ffcommon.FBuf) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_bprint_finalize").Call(
		uintptr(unsafe.Pointer(buf)),
		uintptr(unsafe.Pointer(&ret_str)),
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
 * Escape the content in src and append it to dstbuf.
 *
 * @param dstbuf        already inited destination bprint buffer
 * @param src           string containing the text to escape
 * @param special_chars string containing the special characters which
 *                      need to be escaped, can be NULL
 * @param mode          escape mode to employ, see AV_ESCAPE_MODE_* macros.
 *                      Any unknown value for mode will be considered equivalent to
 *                      AV_ESCAPE_MODE_BACKSLASH, but this behaviour can change without
 *                      notice.
 * @param flags         flags which control how to escape, see AV_ESCAPE_FLAG_* macros
 */
//void av_bprint_escape(AVBPrint *dstbuf, const char *src, const char *special_chars,
//enum AVEscapeMode mode, int flags);
//未测试
func (buf *AVBPrint) AvBprintEscape(src ffcommon.FConstCharP, special_chars ffcommon.FConstCharP,
	mode ffconstant.AVEscapeMode, flags ffcommon.FInt) (err error) {
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
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_bprint_escape").Call(
		uintptr(unsafe.Pointer(buf)),
		uintptr(unsafe.Pointer(srcp)),
		uintptr(unsafe.Pointer(special_charsp)),
		uintptr(mode),
		uintptr(flags),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}
