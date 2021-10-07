package libavutil

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

type AVTEA struct {
}

/**
 * Allocate an AVTEA context
 * To free the struct: av_free(ptr)
 */
//struct AVTEA *av_tea_alloc(void);
//未测试
func AvTeaAlloc() (res *AVTEA, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_tea_alloc").Call()
	if err != nil {
		//return
	}
	res = (*AVTEA)(unsafe.Pointer(t))
	return
}

/**
 * Initialize an AVTEA context.
 *
 * @param ctx an AVTEA context
 * @param key a key of 16 bytes used for encryption/decryption
 * @param rounds the number of rounds in TEA (64 is the "standard")
 */
//void av_tea_init(struct AVTEA *ctx, const uint8_t key[16], int rounds);
//未测试
func (ctx *AVTEA) av_tea_init(key [16]ffcommon.FUint8T, rounds ffcommon.FInt) (err error) {
	_, _, _ = ffcommon.GetAvutilDll().NewProc("av_tea_init").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(&key)),
		uintptr(rounds),
	)
	if err != nil {
		//return
	}
	return
}

/**
 * Encrypt or decrypt a buffer using a previously initialized context.
 *
 * @param ctx an AVTEA context
 * @param dst destination array, can be equal to src
 * @param src source array, can be equal to dst
 * @param count number of 8 byte blocks
 * @param iv initialization vector for CBC mode, if NULL then ECB will be used
 * @param decrypt 0 for encryption, 1 for decryption
 */
//void av_tea_crypt(struct AVTEA *ctx, uint8_t *dst, const uint8_t *src,
//int count, uint8_t *iv, int decrypt);
//未测试
func (ctx *AVTEA) av_tea_crypt(dst *ffcommon.FUint8T, src *ffcommon.FUint8T,
	count ffcommon.FInt, iv *ffcommon.FUint8T, decrypt ffcommon.FInt) (err error) {
	_, _, _ = ffcommon.GetAvutilDll().NewProc("av_tea_crypt").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
		uintptr(count),
		uintptr(unsafe.Pointer(iv)),
		uintptr(decrypt),
	)
	if err != nil {
		//return
	}
	return
}
