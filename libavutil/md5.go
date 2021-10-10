package libavutil

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

type AVMD5 struct {
}

/**
 * Allocate an AVMD5 context.
 */
//struct AVMD5 *av_md5_alloc(void);
//未测试
func AvMd5Alloc() (res *AVMD5, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_md5_alloc").Call()
	if err != nil {
		//return
	}
	res = (*AVMD5)(unsafe.Pointer(t))
	return
}

/**
 * Initialize MD5 hashing.
 *
 * @param ctx pointer to the function context (of size av_md5_size)
 */
//void av_md5_init(struct AVMD5 *ctx);
//未测试
func (res *AVMD5) AvMd5Init() (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_md5_init").Call()
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Update hash value.
 *
 * @param ctx hash function context
 * @param src input data to update hash with
 * @param len input data length
 */
//#if FF_API_CRYPTO_SIZE_T
//void av_md5_update(struct AVMD5 *ctx, const uint8_t *src, int len);
//#else
//void av_md5_update(struct AVMD5 *ctx, const uint8_t *src, size_t len);
//未测试
func (ctx *AVMD5) AvMd5Update(src *ffcommon.FUint8T, len0 ffcommon.FSizeT) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_md5_update").Call(
		uintptr(unsafe.Pointer(ctx)),
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
 * @param ctx hash function context
 * @param dst buffer where output digest value is stored
 */
//void av_md5_final(struct AVMD5 *ctx, uint8_t *dst);
//未测试
func (ctx *AVMD5) AvMd5Final(dst *ffcommon.FUint8T) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_md5_final").Call(
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
 * Hash an array of data.
 *
 * @param dst The output buffer to write the digest into
 * @param src The data to hash
 * @param len The length of the data, in bytes
 */
//#if FF_API_CRYPTO_SIZE_T
//void av_md5_sum(uint8_t *dst, const uint8_t *src, const int len);
//#else
//void av_md5_sum(uint8_t *dst, const uint8_t *src, size_t len);
//#endif
//未测试
func AvMd5Sum(dst *ffcommon.FUint8T, src *ffcommon.FUint8T, len0 ffcommon.FSizeT) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_md5_sum").Call(
		uintptr(unsafe.Pointer(dst)),
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
