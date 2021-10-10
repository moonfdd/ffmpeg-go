package libavutil

import (
	"ffmpeg-go/ffcommon"
	"syscall"
	"unsafe"
)

/**
 * @defgroup lavu_hash Hash Functions
 * @ingroup lavu_crypto
 * Hash functions useful in multimedia.
 *
 * Hash functions are widely used in multimedia, from error checking and
 * concealment to internal regression testing. libavutil has efficient
 * implementations of a variety of hash functions that may be useful for
 * FFmpeg and other multimedia applications.
 *
 * @{
 *
 * @defgroup lavu_hash_generic Generic Hashing API
 * An abstraction layer for all hash functions supported by libavutil.
 *
 * If your application needs to support a wide range of different hash
 * functions, then the Generic Hashing API is for you. It provides a generic,
 * reusable API for @ref lavu_hash "all hash functions" implemented in libavutil.
 * If you just need to use one particular hash function, use the @ref lavu_hash
 * "individual hash" directly.
 *
 * @section Sample Code
 *
 * A basic template for using the Generic Hashing API follows:
 *
 * @code
 * struct AVHashContext *ctx = NULL;
 * const char *hash_name = NULL;
 * uint8_t *output_buf = NULL;
 *
 * // Select from a string returned by av_hash_names()
 * hash_name = ...;
 *
 * // Allocate a hash context
 * ret = av_hash_alloc(&ctx, hash_name);
 * if (ret < 0)
 *     return ret;
 *
 * // Initialize the hash context
 * av_hash_init(ctx);
 *
 * // Update the hash context with data
 * while (data_left) {
 *     av_hash_update(ctx, data, size);
 * }
 *
 * // Now we have no more data, so it is time to finalize the hash and get the
 * // output. But we need to first allocate an output buffer. Note that you can
 * // use any memory allocation function, including malloc(), not just
 * // av_malloc().
 * output_buf = av_malloc(av_hash_get_size(ctx));
 * if (!output_buf)
 *     return AVERROR(ENOMEM);
 *
 * // Finalize the hash context.
 * // You can use any of the av_hash_final*() functions provided, for other
 * // output formats. If you do so, be sure to adjust the memory allocation
 * // above. See the function documentation below for the exact amount of extra
 * // memory needed.
 * av_hash_final(ctx, output_buffer);
 *
 * // Free the context
 * av_hash_freep(&ctx);
 * @endcode
 *
 * @section Hash Function-Specific Information
 * If the CRC32 hash is selected, the #AV_CRC_32_IEEE polynomial will be
 * used.
 *
 * If the Murmur3 hash is selected, the default seed will be used. See @ref
 * lavu_murmur3_seedinfo "Murmur3" for more information.
 *
 * @{
 */

/**
 * @example ffhash.c
 * This example is a simple command line application that takes one or more
 * arguments. It demonstrates a typical use of the hashing API with allocation,
 * initialization, updating, and finalizing.
 */

type AVHashContext struct {
}

/**
 * Allocate a hash context for the algorithm specified by name.
 *
 * @return  >= 0 for success, a negative error code for failure
 *
 * @note The context is not initialized after a call to this function; you must
 * call av_hash_init() to do so.
 */
//int av_hash_alloc(struct AVHashContext **ctx, const char *name);
//未测试
func AvHashAlloc(ctx **AVHashContext, name ffcommon.FConstCharP) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_hash_alloc").Call(
		uintptr(unsafe.Pointer(&ctx)),
		uintptr(unsafe.Pointer(namep)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Get the names of available hash algorithms.
 *
 * This function can be used to enumerate the algorithms.
 *
 * @param[in] i  Index of the hash algorithm, starting from 0
 * @return       Pointer to a static string or `NULL` if `i` is out of range
 */
//const char *av_hash_names(int i);
//未测试
func AvHashNames(i ffcommon.FInt) (res ffcommon.FConstCharP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_hash_names").Call(
		uintptr(i),
	)
	if err != nil {
		//return
	}
	res = ffcommon.GoAStr(t)
	return
}

/**
 * Get the name of the algorithm corresponding to the given hash context.
 */
//const char *av_hash_get_name(const struct AVHashContext *ctx);
//未测试
func (ctx *AVHashContext) AvHashGetName() (res ffcommon.FConstCharP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_hash_get_name").Call(
		uintptr(unsafe.Pointer(ctx)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.GoAStr(t)
	return
}

/**
 * Get the size of the resulting hash value in bytes.
 *
 * The maximum value this function will currently return is available as macro
 * #AV_HASH_MAX_SIZE.
 *
 * @param[in]     ctx Hash context
 * @return            Size of the hash value in bytes
 */
//int av_hash_get_size(const struct AVHashContext *ctx);
//未测试
func (ctx *AVHashContext) AvHashGetSize() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_hash_get_size").Call(
		uintptr(unsafe.Pointer(ctx)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Initialize or reset a hash context.
 *
 * @param[in,out] ctx Hash context
 */
//void av_hash_init(struct AVHashContext *ctx);
//未测试
func (ctx *AVHashContext) AvHashInit() (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_hash_init").Call(
		uintptr(unsafe.Pointer(ctx)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Update a hash context with additional data.
 *
 * @param[in,out] ctx Hash context
 * @param[in]     src Data to be added to the hash context
 * @param[in]     len Size of the additional data
 */
//#if FF_API_CRYPTO_SIZE_T
//void av_hash_update(struct AVHashContext *ctx, const uint8_t *src, int len);
//#else
//void av_hash_update(struct AVHashContext *ctx, const uint8_t *src, size_t len);
//未测试
func (ctx *AVHashContext) AvHashUpdate(src *ffcommon.FUint8T, len0 ffcommon.FSizeT) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_hash_update").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(src)),
		uintptr(len0),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

//#endif

/**
 * Finalize a hash context and compute the actual hash value.
 *
 * The minimum size of `dst` buffer is given by av_hash_get_size() or
 * #AV_HASH_MAX_SIZE. The use of the latter macro is discouraged.
 *
 * It is not safe to update or finalize a hash context again, if it has already
 * been finalized.
 *
 * @param[in,out] ctx Hash context
 * @param[out]    dst Where the final hash value will be stored
 *
 * @see av_hash_final_bin() provides an alternative API
 */
//void av_hash_final(struct AVHashContext *ctx, uint8_t *dst);
//未测试
func (ctx *AVHashContext) AvHashFinal(dst *ffcommon.FUint8T) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_hash_final").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(dst)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Finalize a hash context and store the actual hash value in a buffer.
 *
 * It is not safe to update or finalize a hash context again, if it has already
 * been finalized.
 *
 * If `size` is smaller than the hash size (given by av_hash_get_size()), the
 * hash is truncated; if size is larger, the buffer is padded with 0.
 *
 * @param[in,out] ctx  Hash context
 * @param[out]    dst  Where the final hash value will be stored
 * @param[in]     size Number of bytes to write to `dst`
 */
//void av_hash_final_bin(struct AVHashContext *ctx, uint8_t *dst, int size);
//未测试
func (ctx *AVHashContext) AvHashFinalBin(dst *ffcommon.FUint8T, size ffcommon.FInt) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_hash_final_bin").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(dst)),
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
 * Finalize a hash context and store the hexadecimal representation of the
 * actual hash value as a string.
 *
 * It is not safe to update or finalize a hash context again, if it has already
 * been finalized.
 *
 * The string is always 0-terminated.
 *
 * If `size` is smaller than `2 * hash_size + 1`, where `hash_size` is the
 * value returned by av_hash_get_size(), the string will be truncated.
 *
 * @param[in,out] ctx  Hash context
 * @param[out]    dst  Where the string will be stored
 * @param[in]     size Maximum number of bytes to write to `dst`
 */
//void av_hash_final_hex(struct AVHashContext *ctx, uint8_t *dst, int size);
//未测试
func (ctx *AVHashContext) AvHashFinalHex(dst *ffcommon.FUint8T, size ffcommon.FInt) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_hash_final_hex").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(dst)),
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
 * Finalize a hash context and store the Base64 representation of the
 * actual hash value as a string.
 *
 * It is not safe to update or finalize a hash context again, if it has already
 * been finalized.
 *
 * The string is always 0-terminated.
 *
 * If `size` is smaller than AV_BASE64_SIZE(hash_size), where `hash_size` is
 * the value returned by av_hash_get_size(), the string will be truncated.
 *
 * @param[in,out] ctx  Hash context
 * @param[out]    dst  Where the final hash value will be stored
 * @param[in]     size Maximum number of bytes to write to `dst`
 */
//void av_hash_final_b64(struct AVHashContext *ctx, uint8_t *dst, int size);
//未测试
func (ctx *AVHashContext) AvHashFinalB64(dst *ffcommon.FUint8T, size ffcommon.FInt) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_hash_final_b64").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(dst)),
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
 * Free hash context and set hash context pointer to `NULL`.
 *
 * @param[in,out] ctx  Pointer to hash context
 */
//void av_hash_freep(struct AVHashContext **ctx);
//未测试
func AvHashFreep(ctx **AVHashContext) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_hash_freep").Call(
		uintptr(unsafe.Pointer(&ctx)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}
