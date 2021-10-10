package libavutil

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

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
//未测试
func AvFifoAlloc(size ffcommon.FUnsignedInt) (res *AVFifoBuffer, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_fifo_alloc").Call(
		uintptr(size),
	)
	if err != nil {
		//return
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
//未测试
func AvFifoAllocArray(nmemb ffcommon.FSizeT, size ffcommon.FSizeT) (res *AVFifoBuffer, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_fifo_alloc_array").Call(
		uintptr(nmemb),
		uintptr(size),
	)
	if err != nil {
		//return
	}
	res = (*AVFifoBuffer)(unsafe.Pointer(t))
	return
}

/**
* Free an AVFifoBuffer.
* @param f AVFifoBuffer to free
 */
//void av_fifo_free(AVFifoBuffer *f);
//未测试
func (f *AVFifoBuffer) AvFifoFree() (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_fifo_free").Call()
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
* Free an AVFifoBuffer and reset pointer to NULL.
* @param f AVFifoBuffer to free
 */
//void av_fifo_freep(AVFifoBuffer **f);
//未测试
func AvFifoFreep(f **AVFifoBuffer) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_fifo_freeav_fifo_freep").Call(
		uintptr(unsafe.Pointer(&f)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
* Reset the AVFifoBuffer to the state right after av_fifo_alloc, in particular it is emptied.
* @param f AVFifoBuffer to reset
 */
//void av_fifo_reset(AVFifoBuffer *f);
//未测试
func (f *AVFifoBuffer) AvFifoReset() (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_fifo_reset").Call()
	if err != nil {
		//return
	}
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
//未测试
func (f *AVFifoBuffer) AvFifoSize() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_fifo_size").Call()
	if err != nil {
		//return
	}
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
//未测试
func (f *AVFifoBuffer) AvFifoSpace() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_fifo_space").Call()
	if err != nil {
		//return
	}
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
//未测试
func (f *AVFifoBuffer) av_fifo_generic_peek_at(dest ffcommon.FVoidP, offset ffcommon.FInt, buf_size ffcommon.FInt, ff func(ffcommon.FVoidP, ffcommon.FVoidP, ffcommon.FInt)) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_fifo_generic_peek_at").Call(
		uintptr(unsafe.Pointer(f)),
		uintptr(unsafe.Pointer(dest)),
		uintptr(buf_size),
		uintptr(unsafe.Pointer(&ff)),
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
* Feed data from an AVFifoBuffer to a user-supplied callback.
* Similar as av_fifo_gereric_read but without discarding data.
* @param f AVFifoBuffer to read from
* @param buf_size number of bytes to read
* @param func generic read function
* @param dest data destination
 */
//int av_fifo_generic_peek(AVFifoBuffer *f, void *dest, int buf_size, void (*func)(void*, void*, int));
//未测试
func (f *AVFifoBuffer) AvFifoGenericPeek(dest ffcommon.FVoidP, buf_size ffcommon.FInt, ff func(ffcommon.FVoidP, ffcommon.FVoidP, ffcommon.FInt)) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_fifo_generic_peek").Call(
		uintptr(unsafe.Pointer(f)),
		uintptr(unsafe.Pointer(dest)),
		uintptr(buf_size),
		uintptr(unsafe.Pointer(&ff)),
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
* Feed data from an AVFifoBuffer to a user-supplied callback.
* @param f AVFifoBuffer to read from
* @param buf_size number of bytes to read
* @param func generic read function
* @param dest data destination
 */
//int av_fifo_generic_read(AVFifoBuffer *f, void *dest, int buf_size, void (*func)(void*, void*, int));
//未测试
func (f *AVFifoBuffer) AvFifoGenericRead(dest ffcommon.FVoidP, buf_size ffcommon.FInt, ff func(ffcommon.FVoidP, ffcommon.FVoidP, ffcommon.FInt)) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_fifo_generic_read").Call(
		uintptr(unsafe.Pointer(f)),
		uintptr(unsafe.Pointer(dest)),
		uintptr(buf_size),
		uintptr(unsafe.Pointer(&ff)),
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
//未测试
func (f *AVFifoBuffer) AvFifoGenericWrite(dest ffcommon.FVoidP, buf_size ffcommon.FInt, ff func(ffcommon.FVoidP, ffcommon.FVoidP, ffcommon.FInt)) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_fifo_generic_write").Call(
		uintptr(unsafe.Pointer(f)),
		uintptr(unsafe.Pointer(dest)),
		uintptr(buf_size),
		uintptr(unsafe.Pointer(&ff)),
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
* Resize an AVFifoBuffer.
* In case of reallocation failure, the old FIFO is kept unchanged.
*
* @param f AVFifoBuffer to resize
* @param size new AVFifoBuffer size in bytes
* @return <0 for failure, >=0 otherwise
 */
//int av_fifo_realloc2(AVFifoBuffer *f, unsigned int size);
//未测试
func (f *AVFifoBuffer) AvFifoRealloc2(size ffcommon.FUnsignedInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_fifo_realloc2").Call(
		uintptr(unsafe.Pointer(f)),
		uintptr(size),
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
* Enlarge an AVFifoBuffer.
* In case of reallocation failure, the old FIFO is kept unchanged.
* The new fifo size may be larger than the requested size.
*
* @param f AVFifoBuffer to resize
* @param additional_space the amount of space in bytes to allocate in addition to av_fifo_size()
* @return <0 for failure, >=0 otherwise
 */
//int av_fifo_grow(AVFifoBuffer *f, unsigned int additional_space);
//未测试
func (f *AVFifoBuffer) av_fifo_grow(additional_space ffcommon.FUnsignedInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_fifo_grow").Call(
		uintptr(unsafe.Pointer(f)),
		uintptr(additional_space),
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
* Read and discard the specified amount of data from an AVFifoBuffer.
* @param f AVFifoBuffer to read from
* @param size amount of data to read in bytes
 */
//void av_fifo_drain(AVFifoBuffer *f, int size);
//未测试
func (f *AVFifoBuffer) av_fifo_drain(size ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_fifo_drain").Call(
		uintptr(unsafe.Pointer(f)),
		uintptr(size),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}
