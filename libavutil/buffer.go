package libavutil

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

/**
 * @defgroup lavu_buffer AVBuffer
 * @ingroup lavu_data
 *
 * @{
 * AVBuffer is an API for reference-counted data buffers.
 *
 * There are two core objects in this API -- AVBuffer and AVBufferRef. AVBuffer
 * represents the data buffer itself; it is opaque and not meant to be accessed
 * by the caller directly, but only through AVBufferRef. However, the caller may
 * e.g. compare two AVBuffer pointers to check whether two different references
 * are describing the same data buffer. AVBufferRef represents a single
 * reference to an AVBuffer and it is the object that may be manipulated by the
 * caller directly.
 *
 * There are two functions provided for creating a new AVBuffer with a single
 * reference -- av_buffer_alloc() to just allocate a new buffer, and
 * av_buffer_create() to wrap an existing array in an AVBuffer. From an existing
 * reference, additional references may be created with av_buffer_ref().
 * Use av_buffer_unref() to free a reference (this will automatically free the
 * data once all the references are freed).
 *
 * The convention throughout this API and the rest of FFmpeg is such that the
 * buffer is considered writable if there exists only one reference to it (and
 * it has not been marked as read-only). The av_buffer_is_writable() function is
 * provided to check whether this is true and av_buffer_make_writable() will
 * automatically create a new writable buffer when necessary.
 * Of course nothing prevents the calling code from violating this convention,
 * however that is safe only when all the existing references are under its
 * control.
 *
 * @note Referencing and unreferencing the buffers is thread-safe and thus
 * may be done from multiple threads simultaneously without any need for
 * additional locking.
 *
 * @note Two different references to the same buffer can point to different
 * parts of the buffer (i.e. their AVBufferRef.data will not be equal).
 */

/**
 * A reference counted buffer type. It is opaque and is meant to be used through
 * references (AVBufferRef).
 */
type AVBuffer struct {
}

/**
 * A reference to a data buffer.
 *
 * The size of this struct is not a part of the public ABI and it is not meant
 * to be allocated directly.
 */
type AVBufferRef struct {
	buffer *AVBuffer

	/**
	 * The data buffer. It is considered writable if and only if
	 * this is the only reference to the buffer, in which case
	 * av_buffer_is_writable() returns 1.
	 */
	data *ffcommon.FUint8T
	/**
	 * Size of data in bytes.
	 */
	//#if FF_API_BUFFER_SIZE_T
	//int      size;
	//#else
	//size_t   size;
	//#endif
	size ffcommon.FSizeT
}

/**
 * Allocate an AVBuffer of the given size using av_malloc().
 *
 * @return an AVBufferRef of given size or NULL when out of memory
 */
//#if FF_API_BUFFER_SIZE_T
//AVBufferRef *av_buffer_alloc(int size);
//#else
//AVBufferRef *av_buffer_alloc(size_t size);
//#endif
//未测试
func AvBufferAlloc(size ffcommon.FSizeT) (res *AVBufferRef, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_buffer_alloc").Call(
		uintptr(size),
	)
	if err != nil {
		//return
	}
	res = (*AVBufferRef)(unsafe.Pointer(t))
	return
}

/**
 * Same as av_buffer_alloc(), except the returned buffer will be initialized
 * to zero.
 */
//#if FF_API_BUFFER_SIZE_T
//AVBufferRef *av_buffer_allocz(int size);
//#else
//AVBufferRef *av_buffer_allocz(size_t size);
//#endif
//未测试
func AvBufferAllocz(size ffcommon.FSizeT) (res *AVBufferRef, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_buffer_allocz").Call(
		uintptr(size),
	)
	if err != nil {
		//return
	}
	res = (*AVBufferRef)(unsafe.Pointer(t))
	return
}

/**
 * Create an AVBuffer from an existing array.
 *
 * If this function is successful, data is owned by the AVBuffer. The caller may
 * only access data through the returned AVBufferRef and references derived from
 * it.
 * If this function fails, data is left untouched.
 * @param data   data array
 * @param size   size of data in bytes
 * @param free   a callback for freeing this buffer's data
 * @param opaque parameter to be got for processing or passed to free
 * @param flags  a combination of AV_BUFFER_FLAG_*
 *
 * @return an AVBufferRef referring to data on success, NULL on failure.
 */
//#if FF_API_BUFFER_SIZE_T
//AVBufferRef *av_buffer_create(uint8_t *data, int size,
//#else
//AVBufferRef *av_buffer_create(uint8_t *data, size_t size,
//#endif
//void (*free)(void *opaque, uint8_t *data),
//void *opaque, int flags);
//未测试
func AvBufferCreate(data *ffcommon.FUint8T, size ffcommon.FInt,
	free func(opaque ffcommon.FVoidP, data ffcommon.FUint8T), opaque ffcommon.FVoidP, flags ffcommon.FInt) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_buffer_create").Call(
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
 * Default free callback, which calls av_free() on the buffer data.
 * This function is meant to be passed to av_buffer_create(), not called
 * directly.
 */
//void av_buffer_default_free(void *opaque, uint8_t *data);
//未测试
func AvBufferDefaultFree(opaque ffcommon.FVoidP, data *ffcommon.FUint8T) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_buffer_default_free").Call(
		uintptr(unsafe.Pointer(opaque)),
		uintptr(unsafe.Pointer(data)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Create a new reference to an AVBuffer.
 *
 * @return a new AVBufferRef referring to the same AVBuffer as buf or NULL on
 * failure.
 */
//AVBufferRef *av_buffer_ref(AVBufferRef *buf);
//未测试
func (buf *AVBufferRef) AvBufferRef() (res *AVBufferRef, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_buffer_ref").Call(
		uintptr(unsafe.Pointer(buf)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = (*AVBufferRef)(unsafe.Pointer(t))
	return
}

/**
 * Free a given reference and automatically free the buffer if there are no more
 * references to it.
 *
 * @param buf the reference to be freed. The pointer is set to NULL on return.
 */
//void av_buffer_unref(AVBufferRef **buf);
//未测试
func AvBufferUnref(buf **AVBufferRef) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_buffer_unref").Call(
		uintptr(unsafe.Pointer(&buf)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * @return 1 if the caller may write to the data referred to by buf (which is
 * true if and only if buf is the only reference to the underlying AVBuffer).
 * Return 0 otherwise.
 * A positive answer is valid until av_buffer_ref() is called on buf.
 */
//int av_buffer_is_writable(const AVBufferRef *buf);
//未测试
func (buf *AVBufferRef) AvBufferIsWritable() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_buffer_is_writable").Call(
		uintptr(unsafe.Pointer(buf)),
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
 * @return the opaque parameter set by av_buffer_create.
 */
//void *av_buffer_get_opaque(const AVBufferRef *buf);
//未测试
func (buf *AVBufferRef) AvBufferGetOpaque() (res ffcommon.FVoidP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_buffer_get_opaque").Call(
		uintptr(unsafe.Pointer(buf)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FVoidP(t)
	return
}

//int av_buffer_get_ref_count(const AVBufferRef *buf);
//未测试
func (buf *AVBufferRef) AvBufferGetRefCount() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_buffer_get_ref_count").Call(
		uintptr(unsafe.Pointer(buf)),
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
 * Create a writable reference from a given buffer reference, avoiding data copy
 * if possible.
 *
 * @param buf buffer reference to make writable. On success, buf is either left
 *            untouched, or it is unreferenced and a new writable AVBufferRef is
 *            written in its place. On failure, buf is left untouched.
 * @return 0 on success, a negative AVERROR on failure.
 */
//int av_buffer_make_writable(AVBufferRef **buf);
//未测试
func AvBufferMakeWritable(buf **AVBufferRef) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_buffer_make_writable").Call(
		uintptr(unsafe.Pointer(buf)),
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
 * Reallocate a given buffer.
 *
 * @param buf  a buffer reference to reallocate. On success, buf will be
 *             unreferenced and a new reference with the required size will be
 *             written in its place. On failure buf will be left untouched. *buf
 *             may be NULL, then a new buffer is allocated.
 * @param size required new buffer size.
 * @return 0 on success, a negative AVERROR on failure.
 *
 * @note the buffer is actually reallocated with av_realloc() only if it was
 * initially allocated through av_buffer_realloc(NULL) and there is only one
 * reference to it (i.e. the one passed to this function). In all other cases
 * a new buffer is allocated and the data is copied.
 */
//#if FF_API_BUFFER_SIZE_T
//int av_buffer_realloc(AVBufferRef **buf, int size);
//#else
//int av_buffer_realloc(AVBufferRef **buf, size_t size);
//#endif
//未测试
func AvBufferRealloc(buf **AVBufferRef, size ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_buffer_realloc").Call(
		uintptr(unsafe.Pointer(buf)),
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
 * Ensure dst refers to the same data as src.
 *
 * When *dst is already equivalent to src, do nothing. Otherwise unreference dst
 * and replace it with a new reference to src.
 *
 * @param dst Pointer to either a valid buffer reference or NULL. On success,
 *            this will point to a buffer reference equivalent to src. On
 *            failure, dst will be left untouched.
 * @param src A buffer reference to replace dst with. May be NULL, then this
 *            function is equivalent to av_buffer_unref(dst).
 * @return 0 on success
 *         AVERROR(ENOMEM) on memory allocation failure.
 */
//int av_buffer_replace(AVBufferRef **dst, AVBufferRef *src);
//未测试
func AvBufferReplace(dst **AVBufferRef, src **AVBufferRef) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_buffer_replace").Call(
		uintptr(unsafe.Pointer(&dst)),
		uintptr(unsafe.Pointer(&src)),
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
 * @}
 */

/**
 * @defgroup lavu_bufferpool AVBufferPool
 * @ingroup lavu_data
 *
 * @{
 * AVBufferPool is an API for a lock-free thread-safe pool of AVBuffers.
 *
 * Frequently allocating and freeing large buffers may be slow. AVBufferPool is
 * meant to solve this in cases when the caller needs a set of buffers of the
 * same size (the most obvious use case being buffers for raw video or audio
 * frames).
 *
 * At the beginning, the user must call av_buffer_pool_init() to create the
 * buffer pool. Then whenever a buffer is needed, call av_buffer_pool_get() to
 * get a reference to a new buffer, similar to av_buffer_alloc(). This new
 * reference works in all aspects the same way as the one created by
 * av_buffer_alloc(). However, when the last reference to this buffer is
 * unreferenced, it is returned to the pool instead of being freed and will be
 * reused for subsequent av_buffer_pool_get() calls.
 *
 * When the caller is done with the pool and no longer needs to allocate any new
 * buffers, av_buffer_pool_uninit() must be called to mark the pool as freeable.
 * Once all the buffers are released, it will automatically be freed.
 *
 * Allocating and releasing buffers with this API is thread-safe as long as
 * either the default alloc callback is used, or the user-supplied one is
 * thread-safe.
 */

/**
 * The buffer pool. This structure is opaque and not meant to be accessed
 * directly. It is allocated with av_buffer_pool_init() and freed with
 * av_buffer_pool_uninit().
 */
type AVBufferPool struct {
}

/**
 * Allocate and initialize a buffer pool.
 *
 * @param size size of each buffer in this pool
 * @param alloc a function that will be used to allocate new buffers when the
 * pool is empty. May be NULL, then the default allocator will be used
 * (av_buffer_alloc()).
 * @return newly created buffer pool on success, NULL on error.
 */
//#if FF_API_BUFFER_SIZE_T
//AVBufferPool *av_buffer_pool_init(int size, AVBufferRef* (*alloc)(int size));
//#else
//AVBufferPool *av_buffer_pool_init(size_t size, AVBufferRef* (*alloc)(size_t size));
//#endif
//未测试
func AvBufferPoolInit(size ffcommon.FInt, alloc func(size ffcommon.FInt) *AVBufferRef) (res *AVBufferPool, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_buffer_pool_init").Call(
		uintptr(size),
		uintptr(unsafe.Pointer(&alloc)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = (*AVBufferPool)(unsafe.Pointer(t))
	return
}

/**
 * Allocate and initialize a buffer pool with a more complex allocator.
 *
 * @param size size of each buffer in this pool
 * @param opaque arbitrary user data used by the allocator
 * @param alloc a function that will be used to allocate new buffers when the
 *              pool is empty. May be NULL, then the default allocator will be
 *              used (av_buffer_alloc()).
 * @param pool_free a function that will be called immediately before the pool
 *                  is freed. I.e. after av_buffer_pool_uninit() is called
 *                  by the caller and all the frames are returned to the pool
 *                  and freed. It is intended to uninitialize the user opaque
 *                  data. May be NULL.
 * @return newly created buffer pool on success, NULL on error.
 */
//#if FF_API_BUFFER_SIZE_T
//AVBufferPool *av_buffer_pool_init2(int size, void *opaque,
//AVBufferRef* (*alloc)(void *opaque, int size),
//#else
//AVBufferPool *av_buffer_pool_init2(size_t size, void *opaque,
//AVBufferRef* (*alloc)(void *opaque, size_t size),
//#endif
//void (*pool_free)(void *opaque));
//未测试
func AvBufferPoolInit2(size ffcommon.FInt, opaque ffcommon.FVoidP,
	alloc func(opaque ffcommon.FVoidP, size ffcommon.FInt) *AVBufferRef,
	pool_free func(opaque ffcommon.FVoidP)) (res *AVBufferPool, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_buffer_pool_init2").Call(
		uintptr(size),
		uintptr(unsafe.Pointer(opaque)),
		uintptr(unsafe.Pointer(&alloc)),
		uintptr(unsafe.Pointer(&pool_free)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = (*AVBufferPool)(unsafe.Pointer(t))
	return
}

/**
 * Mark the pool as being available for freeing. It will actually be freed only
 * once all the allocated buffers associated with the pool are released. Thus it
 * is safe to call this function while some of the allocated buffers are still
 * in use.
 *
 * @param pool pointer to the pool to be freed. It will be set to NULL.
 */
//void av_buffer_pool_uninit(AVBufferPool **pool);
//未测试
func AvBufferPoolUninit(pool **AVBufferPool) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_buffer_pool_uninit").Call(
		uintptr(unsafe.Pointer(&pool)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Allocate a new AVBuffer, reusing an old buffer from the pool when available.
 * This function may be called simultaneously from multiple threads.
 *
 * @return a reference to the new buffer on success, NULL on error.
 */
//AVBufferRef *av_buffer_pool_get(AVBufferPool *pool);
//未测试
func (pool *AVBufferPool) AvBufferPoolGet() (res *AVBufferRef, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_buffer_pool_get").Call(
		uintptr(unsafe.Pointer(pool)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = (*AVBufferRef)(unsafe.Pointer(t))
	return
}

/**
* Query the original opaque parameter of an allocated buffer in the pool.
*
* @param ref a buffer reference to a buffer returned by av_buffer_pool_get.
* @return the opaque parameter set by the buffer allocator function of the
*         buffer pool.
*
* @note the opaque parameter of ref is used by the buffer pool implementation,
* therefore you have to use this function to access the original opaque
* parameter of an allocated buffer.
 */
//void *av_buffer_pool_buffer_get_opaque(AVBufferRef *ref);
//未测试
func (pool *AVBufferPool) AvBufferPoolBufferGetOpaque() (res ffcommon.FVoidP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_buffer_pool_buffer_get_opaque").Call(
		uintptr(unsafe.Pointer(pool)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = t
	return
}
