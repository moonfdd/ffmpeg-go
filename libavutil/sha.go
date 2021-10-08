package libavutil

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

type AVSHA struct {
}

/**
 * Allocate an AVSHA context.
 */
//struct AVSHA *av_sha_alloc(void);
//未测试
func AvShaAlloc() (res *AVSHA, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_sha_alloc").Call()
	if err != nil {
		//return
	}
	res = (*AVSHA)(unsafe.Pointer(t))
	return
}

/**
 * Initialize SHA-1 or SHA-2 hashing.
 *
 * @param context pointer to the function context (of size av_sha_size)
 * @param bits    number of bits in digest (SHA-1 - 160 bits, SHA-2 224 or 256 bits)
 * @return        zero if initialization succeeded, -1 otherwise
 */
//int av_sha_init(struct AVSHA* context, int bits);
//未测试
func (context *AVSHA) AvShaInit(bits ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_sha_init").Call(
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
 * @param ctx     hash function context
 * @param data    input data to update hash with
 * @param len     input data length
 */
//#if FF_API_CRYPTO_SIZE_T
//void av_sha_update(struct AVSHA *ctx, const uint8_t *data, unsigned int len);
//#else
//void av_sha_update(struct AVSHA *ctx, const uint8_t *data, size_t len);
//#endif
//未测试
func (context *AVSHA) AvShaUpdate(data *ffcommon.FUint8T, len0 ffcommon.FSizeT) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_sha_update").Call(
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
//void av_sha_final(struct AVSHA* context, uint8_t *digest);
//未测试
func (context *AVSHA) AvShaFinal(digest *ffcommon.FUint8T) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_sha_final").Call(
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
