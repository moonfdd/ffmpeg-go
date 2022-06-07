package libavutil

import (
	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"unsafe"
)

/*
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

/**
 * @file
 * a very simple circular buffer FIFO implementation
 */

//#ifndef AVUTIL_FIFO_H
//#define AVUTIL_FIFO_H
//
//#include <stdint.h>
//#include "avutil.h"
//#include "attributes.h"

type AVFifoBuffer struct {
	Buffer          *ffcommon.FUint8T
	Rptr, Wptr, End *ffcommon.FUint8T
	Rndx, Wndx      ffcommon.FUint32T
}

/**
 * Initialize an AVFifoBuffer.
 * @param size of FIFO
 * @return AVFifoBuffer or NULL in case of memory allocation failure
 */
//AVFifoBuffer *av_fifo_alloc(unsigned int size);
func AvFifoAlloc(size ffcommon.FUnsignedInt) (res *AVFifoBuffer) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_fifo_alloc").Call(
		uintptr(size),
	)
	if t == 0 {

	}
	res = (*AVFifoBuffer)(unsafe.Pointer(t))
	return
}

/**
 * Initialize an AVFifoBuffer.
 * @param nmemb number of elements
 * @param size  size of the single element
 * @return AVFifoBuffer or NULL in case of memory allocation failure
 */
//AVFifoBuffer *av_fifo_alloc_array(size_t nmemb, size_t size);
func AvFifoAllocArray(nmemb, size ffcommon.FSizeT) (res *AVFifoBuffer) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_fifo_alloc_array").Call(
		uintptr(nmemb),
		uintptr(size),
	)
	if t == 0 {

	}
	res = (*AVFifoBuffer)(unsafe.Pointer(t))
	return
}

/**
 * Free an AVFifoBuffer.
 * @param f AVFifoBuffer to free
 */
//void av_fifo_free(AVFifoBuffer *f);
func (f *AVFifoBuffer) AvFifoFree() {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_fifo_free").Call(
		uintptr(unsafe.Pointer(f)),
	)
	if t == 0 {

	}
	return
}

/**
 * Free an AVFifoBuffer and reset pointer to NULL.
 * @param f AVFifoBuffer to free
 */
//void av_fifo_freep(AVFifoBuffer **f);
func AvFifoFreep(f **AVFifoBuffer) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_fifo_freep").Call(
		uintptr(unsafe.Pointer(f)),
	)
	if t == 0 {

	}
	return
}

/**
 * Reset the AVFifoBuffer to the state right after av_fifo_alloc, in particular it is emptied.
 * @param f AVFifoBuffer to reset
 */
//void av_fifo_reset(AVFifoBuffer *f);
func (f *AVFifoBuffer) AvFifoReset() {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_fifo_reset").Call(
		uintptr(unsafe.Pointer(f)),
	)
	if t == 0 {

	}
	return
}

/**
 * Return the amount of data in bytes in the AVFifoBuffer, that is the
 * amount of data you can read from it.
 * @param f AVFifoBuffer to read from
 * @return size
 */
//int av_fifo_size(const AVFifoBuffer *f);
func (f *AVFifoBuffer) AvFifoSize() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_fifo_size").Call(
		uintptr(unsafe.Pointer(f)),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Return the amount of space in bytes in the AVFifoBuffer, that is the
 * amount of data you can write into it.
 * @param f AVFifoBuffer to write into
 * @return size
 */
//int av_fifo_space(const AVFifoBuffer *f);
func (f *AVFifoBuffer) AvFifoSpace() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_fifo_space").Call(
		uintptr(unsafe.Pointer(f)),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Feed data at specific position from an AVFifoBuffer to a user-supplied callback.
 * Similar as av_fifo_gereric_read but without discarding data.
 * @param f AVFifoBuffer to read from
 * @param offset offset from current read position
 * @param buf_size number of bytes to read
 * @param func generic read function
 * @param dest data destination
 */
//int av_fifo_generic_peek_at(AVFifoBuffer *f, void *dest, int offset, int buf_size, void (*func)(void*, void*, int));
func (f *AVFifoBuffer) AvFifoGenericPeekAt(dest ffcommon.FVoidP, offset, buf_size ffcommon.FInt, func0 func(ffcommon.FVoidP, ffcommon.FVoidP, ffcommon.FInt) uintptr) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_fifo_generic_peek_at").Call(
		uintptr(unsafe.Pointer(f)),
		dest,
		uintptr(offset),
		uintptr(buf_size),
		ffcommon.NewCallback(func0),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Feed data from an AVFifoBuffer to a user-supplied callback.
 * Similar as av_fifo_gereric_read but without discarding data.
 * @param f AVFifoBuffer to read from
 * @param buf_size number of bytes to read
 * @param func generic read function
 * @param dest data destination
 */
//int av_fifo_generic_peek(AVFifoBuffer *f, void *dest, int buf_size, void (*func)(void*, void*, int));
func (f *AVFifoBuffer) AvFifoGenericPeek(dest ffcommon.FVoidP, buf_size ffcommon.FInt, func0 func(ffcommon.FVoidP, ffcommon.FVoidP, ffcommon.FInt) uintptr) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_fifo_generic_peek").Call(
		uintptr(unsafe.Pointer(f)),
		dest,
		uintptr(buf_size),
		ffcommon.NewCallback(func0),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Feed data from an AVFifoBuffer to a user-supplied callback.
 * @param f AVFifoBuffer to read from
 * @param buf_size number of bytes to read
 * @param func generic read function
 * @param dest data destination
 */
//int av_fifo_generic_read(AVFifoBuffer *f, void *dest, int buf_size, void (*func)(void*, void*, int));
func (f *AVFifoBuffer) AvFifoGenericRead(dest ffcommon.FVoidP, buf_size ffcommon.FInt, func0 func(ffcommon.FVoidP, ffcommon.FVoidP, ffcommon.FInt) uintptr) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_fifo_generic_read").Call(
		uintptr(unsafe.Pointer(f)),
		dest,
		uintptr(buf_size),
		ffcommon.NewCallback(func0),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Feed data from a user-supplied callback to an AVFifoBuffer.
 * @param f AVFifoBuffer to write to
 * @param src data source; non-const since it may be used as a
 * modifiable context by the function defined in func
 * @param size number of bytes to write
 * @param func generic write function; the first parameter is src,
 * the second is dest_buf, the third is dest_buf_size.
 * func must return the number of bytes written to dest_buf, or <= 0 to
 * indicate no more data available to write.
 * If func is NULL, src is interpreted as a simple byte array for source data.
 * @return the number of bytes written to the FIFO
 */
//int av_fifo_generic_write(AVFifoBuffer *f, void *src, int size, int (*func)(void*, void*, int));
func (f *AVFifoBuffer) AvFifoGenericWrite(src ffcommon.FVoidP, size ffcommon.FInt, func0 func(ffcommon.FVoidP, ffcommon.FVoidP, ffcommon.FInt) uintptr) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_fifo_generic_write").Call(
		uintptr(unsafe.Pointer(f)),
		src,
		uintptr(size),
		ffcommon.NewCallback(func0),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Resize an AVFifoBuffer.
 * In case of reallocation failure, the old FIFO is kept unchanged.
 *
 * @param f AVFifoBuffer to resize
 * @param size new AVFifoBuffer size in bytes
 * @return <0 for failure, >=0 otherwise
 */
//int av_fifo_realloc2(AVFifoBuffer *f, unsigned int size);
func (f *AVFifoBuffer) AvFifoRealloc2(size ffcommon.FUnsignedInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_fifo_realloc2").Call(
		uintptr(unsafe.Pointer(f)),
		uintptr(size),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Enlarge an AVFifoBuffer.
 * In case of reallocation failure, the old FIFO is kept unchanged.
 * The new fifo size may be larger than the requested size.
 *
 * @param f AVFifoBuffer to resize
 * @param additional_space the amount of space in bytes to allocate in addition to av_fifo_size()
 * @return <0 for failure, >=0 otherwise
 */
//int av_fifo_grow(AVFifoBuffer *f, unsigned int additional_space);
func (f *AVFifoBuffer) AvFifoGrow(additional_space ffcommon.FUnsignedInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_fifo_grow").Call(
		uintptr(unsafe.Pointer(f)),
		uintptr(additional_space),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Read and discard the specified amount of data from an AVFifoBuffer.
 * @param f AVFifoBuffer to read from
 * @param size amount of data to read in bytes
 */
//void av_fifo_drain(AVFifoBuffer *f, int size);
func (f *AVFifoBuffer) AvFifoDrain(size ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_fifo_drain").Call(
		uintptr(unsafe.Pointer(f)),
		uintptr(size),
	)
	if t == 0 {

	}
	return
}

/**
 * Return a pointer to the data stored in a FIFO buffer at a certain offset.
 * The FIFO buffer is not modified.
 *
 * @param f    AVFifoBuffer to peek at, f must be non-NULL
 * @param offs an offset in bytes, its absolute value must be less
 *             than the used buffer size or the returned pointer will
 *             point outside to the buffer data.
 *             The used buffer size can be checked with av_fifo_size().
 */
//static inline uint8_t *av_fifo_peek2(const AVFifoBuffer *f, int offs)
//{
//uint8_t *ptr = f->rptr + offs;
//if (ptr >= f->end)
//ptr = f->buffer + (ptr - f->end);
//else if (ptr < f->buffer)
//ptr = f->end - (f->buffer - ptr);
//return ptr;
//}
//todo
func av_fifo_peek2() (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_fifo_peek2").Call()
	if t == 0 {

	}
	res = ffcommon.StringFromPtr(t)
	return
}

//#endif /* AVUTIL_FIFO_H */
