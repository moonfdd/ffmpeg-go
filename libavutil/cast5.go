package libavutil

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

type AVCAST5 struct {
}

/**
 * Allocate an AVCAST5 context
 * To free the struct: av_free(ptr)
 */
//struct AVCAST5 *av_cast5_alloc(void);
//未测试
func AvCast5Alloc() (res *AVCAST5, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_cast5_alloc").Call()
	if err != nil {
		//return
	}
	res = (*AVCAST5)(unsafe.Pointer(t))
	return
}

/**
 * Initialize an AVCAST5 context.
 *
 * @param ctx an AVCAST5 context
 * @param key a key of 5,6,...16 bytes used for encryption/decryption
 * @param key_bits number of keybits: possible are 40,48,...,128
 * @return 0 on success, less than 0 on failure
 */
//int av_cast5_init(struct AVCAST5 *ctx, const uint8_t *key, int key_bits);
//未测试
func (ctx *AVCAST5) AvCast5Init(key *ffcommon.FUint8T, key_bits ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_cast5_init").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(key)),
		uintptr(key_bits),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Encrypt or decrypt a buffer using a previously initialized context, ECB mode only
 *
 * @param ctx an AVCAST5 context
 * @param dst destination array, can be equal to src
 * @param src source array, can be equal to dst
 * @param count number of 8 byte blocks
 * @param decrypt 0 for encryption, 1 for decryption
 */
//void av_cast5_crypt(struct AVCAST5 *ctx, uint8_t *dst, const uint8_t *src, int count, int decrypt);
//未测试
func (ctx *AVCAST5) AvCast5Crypt(dst *ffcommon.FUint8T, src *ffcommon.FUint8T, count ffcommon.FInt, decrypt ffcommon.FInt) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_cast5_crypt").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
		uintptr(count),
		uintptr(decrypt),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Encrypt or decrypt a buffer using a previously initialized context
 *
 * @param ctx an AVCAST5 context
 * @param dst destination array, can be equal to src
 * @param src source array, can be equal to dst
 * @param count number of 8 byte blocks
 * @param iv initialization vector for CBC mode, NULL for ECB mode
 * @param decrypt 0 for encryption, 1 for decryption
 */
//void av_cast5_crypt2(struct AVCAST5 *ctx, uint8_t *dst, const uint8_t *src, int count, uint8_t *iv, int decrypt);
//未测试
func (ctx *AVCAST5) AvCast5Crypt2(dst *ffcommon.FUint8T, src *ffcommon.FUint8T, count ffcommon.FInt, iv *ffcommon.FUint8T, decrypt ffcommon.FInt) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_cast5_crypt2").Call(
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
	if t == 0 {

	}
	return
}
