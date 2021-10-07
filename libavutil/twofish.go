package libavutil

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

type AVTWOFISH struct {
	//instance uintptr
	//// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	//ptr unsafe.Pointer
}

/**
 * Allocate an AVTWOFISH context
 * To free the struct: av_free(ptr)
 */
//struct AVTWOFISH *av_twofish_alloc(void);
//未测试
func AvTwofishAlloc() (res *AVTWOFISH, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_twofish_alloc").Call()
	if err != nil {
		//return
	}
	res = (*AVTWOFISH)(unsafe.Pointer(t))
	return
}

/**
 * Initialize an AVTWOFISH context.
 *
 * @param ctx an AVTWOFISH context
 * @param key a key of size ranging from 1 to 32 bytes used for encryption/decryption
 * @param key_bits number of keybits: 128, 192, 256 If less than the required, padded with zeroes to nearest valid value; return value is 0 if key_bits is 128/192/256, -1 if less than 0, 1 otherwise
 */
//int av_twofish_init(struct AVTWOFISH *ctx, const uint8_t *key, int key_bits);
//未测试
func (ctx *AVTWOFISH) av_twofish_init(key *ffcommon.FUint8T, key_bits ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_twofish_init").Call(
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
 * @param ctx an AVTWOFISH context
 * @param dst destination array, can be equal to src
 * @param src source array, can be equal to dst
 * @param count number of 16 byte blocks
 * @paran iv initialization vector for CBC mode, NULL for ECB mode
 * @param decrypt 0 for encryption, 1 for decryption
 */
//void av_twofish_crypt(struct AVTWOFISH *ctx, uint8_t *dst, const uint8_t *src, int count, uint8_t* iv, int decrypt);
//未测试
func (ctx *AVTWOFISH) av_twofish_crypt(dst *ffcommon.FUint8T, src *ffcommon.FUint8T, count ffcommon.FInt, iv *ffcommon.FUint8T, decrypt ffcommon.FInt) (err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_twofish_crypt").Call(
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
