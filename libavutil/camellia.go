package libavutil

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

type AVCAMELLIA struct {
}

/**
 * Allocate an AVCAMELLIA context
 * To free the struct: av_free(ptr)
 */
//struct AVCAMELLIA *av_camellia_alloc(void);
//未测试
func AvCamelliaAlloc() (res *AVCAMELLIA, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_camellia_alloc").Call()
	if err != nil {
		//return
	}
	res = (*AVCAMELLIA)(unsafe.Pointer(t))
	return
}

/**
 * Initialize an AVCAMELLIA context.
 *
 * @param ctx an AVCAMELLIA context
 * @param key a key of 16, 24, 32 bytes used for encryption/decryption
 * @param key_bits number of keybits: possible are 128, 192, 256
 */
//int av_camellia_init(struct AVCAMELLIA *ctx, const uint8_t *key, int key_bits);
//未测试
func (ctx *AVCAMELLIA) AvCamelliaInit(key *ffcommon.FUint8T, key_bits ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_camellia_init").Call(
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
 * Encrypt or decrypt a buffer using a previously initialized context
 *
 * @param ctx an AVCAMELLIA context
 * @param dst destination array, can be equal to src
 * @param src source array, can be equal to dst
 * @param count number of 16 byte blocks
 * @paran iv initialization vector for CBC mode, NULL for ECB mode
 * @param decrypt 0 for encryption, 1 for decryption
 */
//void av_camellia_crypt(struct AVCAMELLIA *ctx, uint8_t *dst, const uint8_t *src, int count, uint8_t* iv, int decrypt);
//未测试
func (ctx *AVCAST5) AvCamelliaCrypt(dst *ffcommon.FUint8T, src *ffcommon.FUint8T, count ffcommon.FInt, iv *ffcommon.FUint8T, decrypt ffcommon.FInt) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_camellia_crypt").Call(
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
