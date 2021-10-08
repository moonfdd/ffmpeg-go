package libavutil

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

type AVRIPEMD struct {
}

/**
 * Allocate an AVRIPEMD context.
 */
//struct AVRIPEMD *av_ripemd_alloc(void);
//未测试
func AvRipemdAlloc() (res *AVRIPEMD, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_ripemd_alloc").Call()
	if err != nil {
		//return
	}
	res = (*AVRIPEMD)(unsafe.Pointer(t))
	return
}

/**
 * Initialize RIPEMD hashing.
 *
 * @param context pointer to the function context (of size av_ripemd_size)
 * @param bits    number of bits in digest (128, 160, 256 or 320 bits)
 * @return        zero if initialization succeeded, -1 otherwise
 */
//int av_ripemd_init(struct AVRIPEMD* context, int bits);
//未测试
func (context *AVRIPEMD) AvRipemdInit(bits ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_ripemd_init").Call(
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
//void av_ripemd_update(struct AVRIPEMD* context, const uint8_t* data, unsigned int len);
//#else
//void av_ripemd_update(struct AVRIPEMD* context, const uint8_t* data, size_t len);
//#endif
//未测试
func (context *AVRIPEMD) AvRipemdUpdate(data *ffcommon.FUint8T, len0 ffcommon.FSizeT) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_ripemd_update").Call(
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
//void av_ripemd_final(struct AVRIPEMD* context, uint8_t *digest);
//未测试
func (context *AVRIPEMD) AvRipemdFinal(digest *ffcommon.FUint8T) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_ripemd_final").Call(
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
