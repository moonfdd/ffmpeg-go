package libavutil

import (
	"ffmpeg-go/ffcommon"
	"syscall"
	"unsafe"
)

/**
 * @defgroup lavu_mem_funcs Heap Management
 * Functions responsible for allocating, freeing, and copying memory.
 *
 * All memory allocation functions have a built-in upper limit of `INT_MAX`
 * bytes. This may be changed with av_max_alloc(), although exercise extreme
 * caution when doing so.
 *
 * @{
 */

/**
 * Allocate a memory block with alignment suitable for all memory accesses
 * (including vectors if available on the CPU).
 *
 * @param size Size in bytes for the memory block to be allocated
 * @return Pointer to the allocated block, or `NULL` if the block cannot
 *         be allocated
 * @see av_mallocz()
 */
//void *av_malloc(size_t size) av_malloc_attrib av_alloc_size(1);
//未测试
func AvMalloc(size ffcommon.FSizeT) (res ffcommon.FVoidP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_malloc").Call(
		uintptr(size),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = t
	return
}

/**
 * Allocate a memory block with alignment suitable for all memory accesses
 * (including vectors if available on the CPU) and zero all the bytes of the
 * block.
 *
 * @param size Size in bytes for the memory block to be allocated
 * @return Pointer to the allocated block, or `NULL` if it cannot be allocated
 * @see av_malloc()
 */
//void *av_mallocz(size_t size) av_malloc_attrib av_alloc_size(1);
//未测试
func AvMallocz(size ffcommon.FSizeT) (res ffcommon.FVoidP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_mallocz").Call(
		uintptr(size),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = t
	return
}

/**
 * Allocate a memory block for an array with av_malloc().
 *
 * The allocated memory will have size `size * nmemb` bytes.
 *
 * @param nmemb Number of element
 * @param size  Size of a single element
 * @return Pointer to the allocated block, or `NULL` if the block cannot
 *         be allocated
 * @see av_malloc()
 */
//av_alloc_size(1, 2) void *av_malloc_array(size_t nmemb, size_t size);
//未测试
func AvAllocSize(nmemb, size ffcommon.FSizeT) (res ffcommon.FVoidP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_alloc_size").Call(
		uintptr(nmemb),
		uintptr(size),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = t
	return
}

/**
 * Allocate a memory block for an array with av_mallocz().
 *
 * The allocated memory will have size `size * nmemb` bytes.
 *
 * @param nmemb Number of elements
 * @param size  Size of the single element
 * @return Pointer to the allocated block, or `NULL` if the block cannot
 *         be allocated
 *
 * @see av_mallocz()
 * @see av_malloc_array()
 */
//av_alloc_size(1, 2) void *av_mallocz_array(size_t nmemb, size_t size);
//未测试
func AvMalloczArray(nmemb, size ffcommon.FSizeT) (res ffcommon.FVoidP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_mallocz_array").Call(
		uintptr(nmemb),
		uintptr(size),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = t
	return
}

/**
 * Non-inlined equivalent of av_mallocz_array().
 *
 * Created for symmetry with the calloc() C function.
 */
//void *av_calloc(size_t nmemb, size_t size) av_malloc_attrib;
//未测试
func AvCalloc(nmemb, size ffcommon.FSizeT) (res ffcommon.FVoidP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_calloc").Call(
		uintptr(nmemb),
		uintptr(size),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = t
	return
}

/**
 * Allocate, reallocate, or free a block of memory.
 *
 * If `ptr` is `NULL` and `size` > 0, allocate a new block. If `size` is
 * zero, free the memory block pointed to by `ptr`. Otherwise, expand or
 * shrink that block of memory according to `size`.
 *
 * @param ptr  Pointer to a memory block already allocated with
 *             av_realloc() or `NULL`
 * @param size Size in bytes of the memory block to be allocated or
 *             reallocated
 *
 * @return Pointer to a newly-reallocated block or `NULL` if the block
 *         cannot be reallocated or the function is used to free the memory block
 *
 * @warning Unlike av_malloc(), the returned pointer is not guaranteed to be
 *          correctly aligned.
 * @see av_fast_realloc()
 * @see av_reallocp()
 */
//void *av_realloc(void *ptr, size_t size) av_alloc_size(2);
//未测试
func AvRealloc(ptr ffcommon.FVoidP, size ffcommon.FSizeT) (res ffcommon.FVoidP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_realloc").Call(
		ptr,
		uintptr(size),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = t
	return
}

/**
 * Allocate, reallocate, or free a block of memory through a pointer to a
 * pointer.
 *
 * If `*ptr` is `NULL` and `size` > 0, allocate a new block. If `size` is
 * zero, free the memory block pointed to by `*ptr`. Otherwise, expand or
 * shrink that block of memory according to `size`.
 *
 * @param[in,out] ptr  Pointer to a pointer to a memory block already allocated
 *                     with av_realloc(), or a pointer to `NULL`. The pointer
 *                     is updated on success, or freed on failure.
 * @param[in]     size Size in bytes for the memory block to be allocated or
 *                     reallocated
 *
 * @return Zero on success, an AVERROR error code on failure
 *
 * @warning Unlike av_malloc(), the allocated memory is not guaranteed to be
 *          correctly aligned.
 */
//av_warn_unused_result
//int av_reallocp(void *ptr, size_t size);
//未测试
func AvReallocp(ptr ffcommon.FVoidP, size ffcommon.FSizeT) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_reallocp").Call(
		ptr,
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
 * Allocate, reallocate, or free a block of memory.
 *
 * This function does the same thing as av_realloc(), except:
 * - It takes two size arguments and allocates `nelem * elsize` bytes,
 *   after checking the result of the multiplication for integer overflow.
 * - It frees the input block in case of failure, thus avoiding the memory
 *   leak with the classic
 *   @code{.c}
 *   buf = realloc(buf);
 *   if (!buf)
 *       return -1;
 *   @endcode
 *   pattern.
 */
//void *av_realloc_f(void *ptr, size_t nelem, size_t elsize);
//未测试
func AvReallocF(ptr ffcommon.FVoidP, nelem ffcommon.FSizeT, elsize ffcommon.FSizeT) (res ffcommon.FVoidP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_realloc_f").Call(
		ptr,
		uintptr(nelem),
		uintptr(elsize),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = t
	return
}

/**
 * Allocate, reallocate, or free an array.
 *
 * If `ptr` is `NULL` and `nmemb` > 0, allocate a new block. If
 * `nmemb` is zero, free the memory block pointed to by `ptr`.
 *
 * @param ptr   Pointer to a memory block already allocated with
 *              av_realloc() or `NULL`
 * @param nmemb Number of elements in the array
 * @param size  Size of the single element of the array
 *
 * @return Pointer to a newly-reallocated block or NULL if the block
 *         cannot be reallocated or the function is used to free the memory block
 *
 * @warning Unlike av_malloc(), the allocated memory is not guaranteed to be
 *          correctly aligned.
 * @see av_reallocp_array()
 */
//av_alloc_size(2, 3) void *av_realloc_array(void *ptr, size_t nmemb, size_t size);
//未测试
func AvReallocArray(ptr ffcommon.FVoidP, nmemb ffcommon.FSizeT, size ffcommon.FSizeT) (res ffcommon.FVoidP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_realloc_array").Call(
		ptr,
		uintptr(nmemb),
		uintptr(size),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = t
	return
}

/**
 * Allocate, reallocate, or free an array through a pointer to a pointer.
 *
 * If `*ptr` is `NULL` and `nmemb` > 0, allocate a new block. If `nmemb` is
 * zero, free the memory block pointed to by `*ptr`.
 *
 * @param[in,out] ptr   Pointer to a pointer to a memory block already
 *                      allocated with av_realloc(), or a pointer to `NULL`.
 *                      The pointer is updated on success, or freed on failure.
 * @param[in]     nmemb Number of elements
 * @param[in]     size  Size of the single element
 *
 * @return Zero on success, an AVERROR error code on failure
 *
 * @warning Unlike av_malloc(), the allocated memory is not guaranteed to be
 *          correctly aligned.
 */
//int av_reallocp_array(void *ptr, size_t nmemb, size_t size);
//未测试
func AvReallocpArray(ptr ffcommon.FVoidP, nmemb ffcommon.FSizeT, size ffcommon.FSizeT) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_reallocp_array").Call(
		ptr,
		uintptr(nmemb),
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
 * Reallocate the given buffer if it is not large enough, otherwise do nothing.
 *
 * If the given buffer is `NULL`, then a new uninitialized buffer is allocated.
 *
 * If the given buffer is not large enough, and reallocation fails, `NULL` is
 * returned and `*size` is set to 0, but the original buffer is not changed or
 * freed.
 *
 * A typical use pattern follows:
 *
 * @code{.c}
 * uint8_t *buf = ...;
 * uint8_t *new_buf = av_fast_realloc(buf, &current_size, size_needed);
 * if (!new_buf) {
 *     // Allocation failed; clean up original buffer
 *     av_freep(&buf);
 *     return AVERROR(ENOMEM);
 * }
 * @endcode
 *
 * @param[in,out] ptr      Already allocated buffer, or `NULL`
 * @param[in,out] size     Pointer to the size of buffer `ptr`. `*size` is
 *                         updated to the new allocated size, in particular 0
 *                         in case of failure.
 * @param[in]     min_size Desired minimal size of buffer `ptr`
 * @return `ptr` if the buffer is large enough, a pointer to newly reallocated
 *         buffer if the buffer was not large enough, or `NULL` in case of
 *         error
 * @see av_realloc()
 * @see av_fast_malloc()
 */
//void *av_fast_realloc(void *ptr, unsigned int *size, size_t min_size);
//未测试
func AvFastRealloc(ptr ffcommon.FVoidP, size *ffcommon.FInt, min_size ffcommon.FSizeT) (res ffcommon.FVoidP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_fast_realloc").Call(
		ptr,
		uintptr(unsafe.Pointer(size)),
		uintptr(min_size),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FVoidP(t)
	return
}

/**
 * Allocate a buffer, reusing the given one if large enough.
 *
 * Contrary to av_fast_realloc(), the current buffer contents might not be
 * preserved and on error the old buffer is freed, thus no special handling to
 * avoid memleaks is necessary.
 *
 * `*ptr` is allowed to be `NULL`, in which case allocation always happens if
 * `size_needed` is greater than 0.
 *
 * @code{.c}
 * uint8_t *buf = ...;
 * av_fast_malloc(&buf, &current_size, size_needed);
 * if (!buf) {
 *     // Allocation failed; buf already freed
 *     return AVERROR(ENOMEM);
 * }
 * @endcode
 *
 * @param[in,out] ptr      Pointer to pointer to an already allocated buffer.
 *                         `*ptr` will be overwritten with pointer to new
 *                         buffer on success or `NULL` on failure
 * @param[in,out] size     Pointer to the size of buffer `*ptr`. `*size` is
 *                         updated to the new allocated size, in particular 0
 *                         in case of failure.
 * @param[in]     min_size Desired minimal size of buffer `*ptr`
 * @see av_realloc()
 * @see av_fast_mallocz()
 */
//void av_fast_malloc(void *ptr, unsigned int *size, size_t min_size);
//未测试
func AvFastMalloc(ptr ffcommon.FVoidP, size *ffcommon.FInt, min_size ffcommon.FSizeT) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_fast_malloc").Call(
		ptr,
		uintptr(unsafe.Pointer(size)),
		uintptr(min_size),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Allocate and clear a buffer, reusing the given one if large enough.
 *
 * Like av_fast_malloc(), but all newly allocated space is initially cleared.
 * Reused buffer is not cleared.
 *
 * `*ptr` is allowed to be `NULL`, in which case allocation always happens if
 * `size_needed` is greater than 0.
 *
 * @param[in,out] ptr      Pointer to pointer to an already allocated buffer.
 *                         `*ptr` will be overwritten with pointer to new
 *                         buffer on success or `NULL` on failure
 * @param[in,out] size     Pointer to the size of buffer `*ptr`. `*size` is
 *                         updated to the new allocated size, in particular 0
 *                         in case of failure.
 * @param[in]     min_size Desired minimal size of buffer `*ptr`
 * @see av_fast_malloc()
 */
//void av_fast_mallocz(void *ptr, unsigned int *size, size_t min_size);
//未测试
func AvFastMallocz(ptr ffcommon.FVoidP, size *ffcommon.FInt, min_size ffcommon.FSizeT) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_fast_mallocz").Call(
		ptr,
		uintptr(unsafe.Pointer(size)),
		uintptr(min_size),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Free a memory block which has been allocated with a function of av_malloc()
 * or av_realloc() family.
 *
 * @param ptr Pointer to the memory block which should be freed.
 *
 * @note `ptr = NULL` is explicitly allowed.
 * @note It is recommended that you use av_freep() instead, to prevent leaving
 *       behind dangling pointers.
 * @see av_freep()
 */
//void av_free(void *ptr);
//未测试
func AvFree(ptr ffcommon.FVoidP) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_free").Call(
		ptr,
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Free a memory block which has been allocated with a function of av_malloc()
 * or av_realloc() family, and set the pointer pointing to it to `NULL`.
 *
 * @code{.c}
 * uint8_t *buf = av_malloc(16);
 * av_free(buf);
 * // buf now contains a dangling pointer to freed memory, and accidental
 * // dereference of buf will result in a use-after-free, which may be a
 * // security risk.
 *
 * uint8_t *buf = av_malloc(16);
 * av_freep(&buf);
 * // buf is now NULL, and accidental dereference will only result in a
 * // NULL-pointer dereference.
 * @endcode
 *
 * @param ptr Pointer to the pointer to the memory block which should be freed
 * @note `*ptr = NULL` is safe and leads to no action.
 * @see av_free()
 */
//void av_freep(void *ptr);
//未测试
func AvFreep(ptr ffcommon.FVoidP) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_freep").Call(
		ptr,
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Duplicate a string.
 *
 * @param s String to be duplicated
 * @return Pointer to a newly-allocated string containing a
 *         copy of `s` or `NULL` if the string cannot be allocated
 * @see av_strndup()
 */
//char *av_strdup(const char *s) av_malloc_attrib;
//未测试
func AvStrdup(s ffcommon.FConstCharP) (res ffcommon.FCharP, err error) {
	var t uintptr
	var sp *byte
	sp, err = syscall.BytePtrFromString(s)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_strdup").Call(
		uintptr(unsafe.Pointer(sp)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.GoAStr(t)
	return
}

/**
 * Duplicate a substring of a string.
 *
 * @param s   String to be duplicated
 * @param len Maximum length of the resulting string (not counting the
 *            terminating byte)
 * @return Pointer to a newly-allocated string containing a
 *         substring of `s` or `NULL` if the string cannot be allocated
 */
//char *av_strndup(const char *s, size_t len) av_malloc_attrib;
//未测试
func AvStrndup(s ffcommon.FConstCharP, len0 ffcommon.FSizeT) (res ffcommon.FCharP, err error) {
	var t uintptr
	var sp *byte
	sp, err = syscall.BytePtrFromString(s)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_strndup").Call(
		uintptr(unsafe.Pointer(sp)),
		uintptr(len0),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.GoAStr(t)
	return
}

/**
 * Duplicate a buffer with av_malloc().
 *
 * @param p    Buffer to be duplicated
 * @param size Size in bytes of the buffer copied
 * @return Pointer to a newly allocated buffer containing a
 *         copy of `p` or `NULL` if the buffer cannot be allocated
 */
//void *av_memdup(const void *p, size_t size);
//未测试
func AvMemdup(p ffcommon.FVoidP, size ffcommon.FSizeT) (res ffcommon.FVoidP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_memdup").Call(
		uintptr(unsafe.Pointer(p)),
		uintptr(size),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = t
	return
}

/**
 * Overlapping memcpy() implementation.
 *
 * @param dst  Destination buffer
 * @param back Number of bytes back to start copying (i.e. the initial size of
 *             the overlapping window); must be > 0
 * @param cnt  Number of bytes to copy; must be >= 0
 *
 * @note `cnt > back` is valid, this will copy the bytes we just copied,
 *       thus creating a repeating pattern with a period length of `back`.
 */
//void av_memcpy_backptr(uint8_t *dst, int back, int cnt);
//未测试
func AvMemcpyBackptr(dst *ffcommon.FUint8T, back ffcommon.FInt, cnt ffcommon.FInt) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_memcpy_backptr").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(back),
		uintptr(cnt),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * @}
 */

/**
 * @defgroup lavu_mem_dynarray Dynamic Array
 *
 * Utilities to make an array grow when needed.
 *
 * Sometimes, the programmer would want to have an array that can grow when
 * needed. The libavutil dynamic array utilities fill that need.
 *
 * libavutil supports two systems of appending elements onto a dynamically
 * allocated array, the first one storing the pointer to the value in the
 * array, and the second storing the value directly. In both systems, the
 * caller is responsible for maintaining a variable containing the length of
 * the array, as well as freeing of the array after use.
 *
 * The first system stores pointers to values in a block of dynamically
 * allocated memory. Since only pointers are stored, the function does not need
 * to know the size of the type. Both av_dynarray_add() and
 * av_dynarray_add_nofree() implement this system.
 *
 * @code
 * type **array = NULL; //< an array of pointers to values
 * int    nb    = 0;    //< a variable to keep track of the length of the array
 *
 * type to_be_added  = ...;
 * type to_be_added2 = ...;
 *
 * av_dynarray_add(&array, &nb, &to_be_added);
 * if (nb == 0)
 *     return AVERROR(ENOMEM);
 *
 * av_dynarray_add(&array, &nb, &to_be_added2);
 * if (nb == 0)
 *     return AVERROR(ENOMEM);
 *
 * // Now:
 * //  nb           == 2
 * // &to_be_added  == array[0]
 * // &to_be_added2 == array[1]
 *
 * av_freep(&array);
 * @endcode
 *
 * The second system stores the value directly in a block of memory. As a
 * result, the function has to know the size of the type. av_dynarray2_add()
 * implements this mechanism.
 *
 * @code
 * type *array = NULL; //< an array of values
 * int   nb    = 0;    //< a variable to keep track of the length of the array
 *
 * type to_be_added  = ...;
 * type to_be_added2 = ...;
 *
 * type *addr = av_dynarray2_add((void **)&array, &nb, sizeof(*array), NULL);
 * if (!addr)
 *     return AVERROR(ENOMEM);
 * memcpy(addr, &to_be_added, sizeof(to_be_added));
 *
 * // Shortcut of the above.
 * type *addr = av_dynarray2_add((void **)&array, &nb, sizeof(*array),
 *                               (const void *)&to_be_added2);
 * if (!addr)
 *     return AVERROR(ENOMEM);
 *
 * // Now:
 * //  nb           == 2
 * //  to_be_added  == array[0]
 * //  to_be_added2 == array[1]
 *
 * av_freep(&array);
 * @endcode
 *
 * @{
 */

/**
 * Add the pointer to an element to a dynamic array.
 *
 * The array to grow is supposed to be an array of pointers to
 * structures, and the element to add must be a pointer to an already
 * allocated structure.
 *
 * The array is reallocated when its size reaches powers of 2.
 * Therefore, the amortized cost of adding an element is ffconstant.
 *
 * In case of success, the pointer to the array is updated in order to
 * point to the new grown array, and the number pointed to by `nb_ptr`
 * is incremented.
 * In case of failure, the array is freed, `*tab_ptr` is set to `NULL` and
 * `*nb_ptr` is set to 0.
 *
 * @param[in,out] tab_ptr Pointer to the array to grow
 * @param[in,out] nb_ptr  Pointer to the number of elements in the array
 * @param[in]     elem    Element to add
 * @see av_dynarray_add_nofree(), av_dynarray2_add()
 */
//void av_dynarray_add(void *tab_ptr, int *nb_ptr, void *elem);
//未测试
func AvDynarrayAdd(tab_ptr ffcommon.FVoidP, nb_ptr *ffcommon.FInt, elem ffcommon.FVoidP) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_dynarray_add").Call(
		tab_ptr,
		uintptr(unsafe.Pointer(nb_ptr)),
		elem,
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Add an element to a dynamic array.
 *
 * Function has the same functionality as av_dynarray_add(),
 * but it doesn't free memory on fails. It returns error code
 * instead and leave current buffer untouched.
 *
 * @return >=0 on success, negative otherwise
 * @see av_dynarray_add(), av_dynarray2_add()
 */
//av_warn_unused_result
//int av_dynarray_add_nofree(void *tab_ptr, int *nb_ptr, void *elem);
//未测试
func AvDynarrayAddNofree(tab_ptr ffcommon.FVoidP, nb_ptr *ffcommon.FInt, elem ffcommon.FVoidP) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_dynarray_add_nofree").Call(
		tab_ptr,
		uintptr(unsafe.Pointer(nb_ptr)),
		elem,
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
 * Add an element of size `elem_size` to a dynamic array.
 *
 * The array is reallocated when its number of elements reaches powers of 2.
 * Therefore, the amortized cost of adding an element is ffconstant.
 *
 * In case of success, the pointer to the array is updated in order to
 * point to the new grown array, and the number pointed to by `nb_ptr`
 * is incremented.
 * In case of failure, the array is freed, `*tab_ptr` is set to `NULL` and
 * `*nb_ptr` is set to 0.
 *
 * @param[in,out] tab_ptr   Pointer to the array to grow
 * @param[in,out] nb_ptr    Pointer to the number of elements in the array
 * @param[in]     elem_size Size in bytes of an element in the array
 * @param[in]     elem_data Pointer to the data of the element to add. If
 *                          `NULL`, the space of the newly added element is
 *                          allocated but left uninitialized.
 *
 * @return Pointer to the data of the element to copy in the newly allocated
 *         space
 * @see av_dynarray_add(), av_dynarray_add_nofree()
 */
//void *av_dynarray2_add(void **tab_ptr, int *nb_ptr, size_t elem_size,
//const uint8_t *elem_data);
//未测试
func AvDynarray2Add(tab_ptr *ffcommon.FVoidP, nb_ptr *ffcommon.FInt, elem_size ffcommon.FSizeT,
	elem_data *ffcommon.FUint8T) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_dynarray2_add").Call(
		uintptr(unsafe.Pointer(tab_ptr)),
		uintptr(unsafe.Pointer(nb_ptr)),
		uintptr(elem_size),
		uintptr(unsafe.Pointer(elem_data)),
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
 * Set the maximum size that may be allocated in one block.
 *
 * The value specified with this function is effective for all libavutil's @ref
 * lavu_mem_funcs "heap management functions."
 *
 * By default, the max value is defined as `INT_MAX`.
 *
 * @param max Value to be set as the new maximum size
 *
 * @warning Exercise extreme caution when using this function. Don't touch
 *          this if you do not understand the full consequence of doing so.
 */
//void av_max_alloc(size_t max);
//未测试
func AvMaxAlloc(max ffcommon.FSizeT) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_max_alloc").Call(
		uintptr(max),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}
