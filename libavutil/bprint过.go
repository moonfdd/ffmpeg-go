package libavutil

import (
	"github.com/moonfdd/ffmpeg-go/ffcommon"
)

/*
 * Copyright (c) 2012 Nicolas George
 *
 * This file is part of FFmpeg.
 *
 * FFmpeg is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * FFmpeg is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with FFmpeg; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA
 */

//#ifndef AVUTIL_BPRINT_H
//#define AVUTIL_BPRINT_H
//
//#include <stdarg.h>
//
//#include "attributes.h"
//#include "avstring.h"
//
///**
// * Define a structure with extra padding to a fixed size
// * This helps ensuring binary compatibility with future versions.
// */
//
//#define FF_PAD_STRUCTURE(name, size, ...) \
//struct ff_pad_helper_##name { __VA_ARGS__ }; \
//typedef struct name { \
//__VA_ARGS__ \
//char reserved_padding[size - sizeof(struct ff_pad_helper_##name)]; \
//} name;

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
	_ [128]uintptr
}

/**
 * Convenience macros for special values for av_bprint_init() size_max
 * parameter.
 */
//#define AV_BPRINT_SIZE_UNLIMITED  ((unsigned)-1)
const AV_BPRINT_SIZE_AUTOMATIC = 1
const AV_BPRINT_SIZE_COUNT_ONLY = 0

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
//todo
func av_bprint_init() (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_blowfish_crypt_ecb").Call()
	if t == 0 {

	}
	res = ffcommon.StringFromPtr(t)
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
//todo
func av_bprint_init_for_buffer() (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_bprint_init_for_buffer").Call()
	if t == 0 {

	}
	res = ffcommon.StringFromPtr(t)
	return
}

/**
 * Append a formatted string to a print buffer.
 */
//void av_bprintf(AVBPrint *buf, const char *fmt, ...) av_printf_format(2, 3);
//todo
func av_bprintf() (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_bprintf").Call()
	if t == 0 {

	}
	res = ffcommon.StringFromPtr(t)
	return
}

/**
 * Append a formatted string to a print buffer.
 */
//void av_vbprintf(AVBPrint *buf, const char *fmt, va_list vl_arg);
//todo
func av_vbprintf() (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_vbprintf").Call()
	if t == 0 {

	}
	res = ffcommon.StringFromPtr(t)
	return
}

/**
 * Append char c n times to a print buffer.
 */
//void av_bprint_chars(AVBPrint *buf, char c, unsigned n);
//todo
func av_bprint_chars() (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_bprint_chars").Call()
	if t == 0 {

	}
	res = ffcommon.StringFromPtr(t)
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
//todo
func av_bprint_append_data() (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_bprint_append_data").Call()
	if t == 0 {

	}
	res = ffcommon.StringFromPtr(t)
	return
}

//struct tm;
type Tm struct {
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
//todo
func av_bprint_strftime() (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_bprint_strftime").Call()
	if t == 0 {

	}
	res = ffcommon.StringFromPtr(t)
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
//todo
func av_bprint_get_buffer() (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_bprint_get_buffer").Call()
	if t == 0 {

	}
	res = ffcommon.StringFromPtr(t)
	return
}

/**
 * Reset the string to "" but keep internal allocated data.
 */
//void av_bprint_clear(AVBPrint *buf);
//todo
func av_bprint_clear() (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_bprint_clear").Call()
	if t == 0 {

	}
	res = ffcommon.StringFromPtr(t)
	return
}

/**
 * Test if the print buffer is complete (not truncated).
 *
 * It may have been truncated due to a memory allocation failure
 * or the size_max limit (compare size and size_max if necessary).
 */
//static inline int av_bprint_is_complete(const AVBPrint *buf)
//{
//return buf->len < buf->size;
//}
//todo
func av_bprint_is_complete() (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_bprint_is_complete").Call()
	if t == 0 {

	}
	res = ffcommon.StringFromPtr(t)
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
//todo
func av_bprint_finalize() (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_bprint_finalize").Call()
	if t == 0 {

	}
	res = ffcommon.StringFromPtr(t)
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
//todo
func av_bprint_escape() (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_bprint_escape").Call()
	if t == 0 {

	}
	res = ffcommon.StringFromPtr(t)
	return
}

//#endif /* AVUTIL_BPRINT_H */
