package libavutil

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

/**
 * @defgroup lavu_murmur3 Murmur3
 * @ingroup lavu_hash
 * MurmurHash3 hash function implementation.
 *
 * MurmurHash3 is a non-cryptographic hash function, of which three
 * incompatible versions were created by its inventor Austin Appleby:
 *
 * - 32-bit output
 * - 128-bit output for 32-bit platforms
 * - 128-bit output for 64-bit platforms
 *
 * FFmpeg only implements the last variant: 128-bit output designed for 64-bit
 * platforms. Even though the hash function was designed for 64-bit platforms,
 * the function in reality works on 32-bit systems too, only with reduced
 * performance.
 *
 * @anchor lavu_murmur3_seedinfo
 * By design, MurmurHash3 requires a seed to operate. In response to this,
 * libavutil provides two functions for hash initiation, one that requires a
 * seed (av_murmur3_init_seeded()) and one that uses a fixed arbitrary integer
 * as the seed, and therefore does not (av_murmur3_init()).
 *
 * To make hashes comparable, you should provide the same seed for all calls to
 * this hash function -- if you are supplying one yourself, that is.
 *
 * @{
 */

/**
 * Allocate an AVMurMur3 hash context.
 *
 * @return Uninitialized hash context or `NULL` in case of error
 */
//struct AVMurMur3 *av_murmur3_alloc(void);
//未测试
func AvMurmur3Alloc() (res *AVMurMur3, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_murmur3_alloc").Call()
	if err != nil {
		//return
	}
	res = (*AVMurMur3)(unsafe.Pointer(t))
	return
}

/**
 * Initialize or reinitialize an AVMurMur3 hash context with a seed.
 *
 * @param[out] c    Hash context
 * @param[in]  seed Random seed
 *
 * @see av_murmur3_init()
 * @see @ref lavu_murmur3_seedinfo "Detailed description" on a discussion of
 * seeds for MurmurHash3.
 */
//void av_murmur3_init_seeded(struct AVMurMur3 *c, uint64_t seed);
//未测试
func (c *AVMurMur3) AvMurmur3InitSeeded(seed ffcommon.FUint64T) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_murmur3_init_seeded").Call(
		uintptr(unsafe.Pointer(c)),
		uintptr(seed),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Initialize or reinitialize an AVMurMur3 hash context.
 *
 * Equivalent to av_murmur3_init_seeded() with a built-in seed.
 *
 * @param[out] c    Hash context
 *
 * @see av_murmur3_init_seeded()
 * @see @ref lavu_murmur3_seedinfo "Detailed description" on a discussion of
 * seeds for MurmurHash3.
 */
//void av_murmur3_init(struct AVMurMur3 *c);
//未测试
func (c *AVMurMur3) AvMurmur3Init() (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_murmur3_init").Call(
		uintptr(unsafe.Pointer(c)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Update hash context with new data.
 *
 * @param[out] c    Hash context
 * @param[in]  src  Input data to update hash with
 * @param[in]  len  Number of bytes to read from `src`
 */
//#if FF_API_CRYPTO_SIZE_T
//void av_murmur3_update(struct AVMurMur3 *c, const uint8_t *src, int len);
//#else
//void av_murmur3_update(struct AVMurMur3 *c, const uint8_t *src, size_t len);
//未测试
func (c *AVMurMur3) AvMurmur3Update(src *ffcommon.FUint8T, len0 ffcommon.FSizeT) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_murmur3_update").Call(
		uintptr(unsafe.Pointer(c)),
		uintptr(unsafe.Pointer(src)),
		uintptr(len0),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

//#endif

/**
 * Finish hashing and output digest value.
 *
 * @param[in,out] c    Hash context
 * @param[out]    dst  Buffer where output digest value is stored
 */
//void av_murmur3_final(struct AVMurMur3 *c, uint8_t dst[16]);
//未测试
func (c *AVMurMur3) AvMurmur3Final(dst [16]ffcommon.FUint8T) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_murmur3_final").Call(
		uintptr(unsafe.Pointer(c)),
		uintptr(unsafe.Pointer(&dst)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}
