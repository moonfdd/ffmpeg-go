package libavutil

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

type AVSHA512 struct {
}

/**
 * Allocate an AVSHA512 context.
 */
//struct AVSHA512 *av_sha512_alloc(void);
//未测试
func AvSha512Alloc() (res *AVSHA512, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_sha512_alloc").Call()
	if err != nil {
		//return
	}
	res = (*AVSHA512)(unsafe.Pointer(t))
	return
}

/**
 * Initialize SHA-2 512 hashing.
 *
 * @param context pointer to the function context (of size av_sha512_size)
 * @param bits    number of bits in digest (224, 256, 384 or 512 bits)
 * @return        zero if initialization succeeded, -1 otherwise
 */
//int av_sha512_init(struct AVSHA512* context, int bits);
//未测试
func (context *AVSHA512) AvSha512Init(bits ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_sha512_init").Call(
		uintptr(unsafe.Pointer(context)),
		uintptr(bits),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Update hash value.
 *
 * @param context hash function context
 * @param data    input data to update hash with
 * @param len     input data length
 */
//#if FF_API_CRYPTO_SIZE_T
//void av_sha512_update(struct AVSHA512* context, const uint8_t* data, unsigned int len);
//#else
//void av_sha512_update(struct AVSHA512* context, const uint8_t* data, size_t len);
//#endif
//未测试
func (context *AVSHA512) AvSha512Update(data *ffcommon.FUint8T, len0 ffcommon.FSizeT) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_sha512_update").Call(
		uintptr(unsafe.Pointer(context)),
		uintptr(unsafe.Pointer(data)),
		uintptr(len0),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Finish hashing and output digest value.
 *
 * @param context hash function context
 * @param digest  buffer where output digest value is stored
 */
//void av_sha512_final(struct AVSHA512* context, uint8_t *digest);
//未测试
func (context *AVSHA512) AvSha512Final(digest *ffcommon.FUint8T) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_sha512_update").Call(
		uintptr(unsafe.Pointer(context)),
		uintptr(unsafe.Pointer(digest)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}
