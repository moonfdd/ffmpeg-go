package libavutil

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

type AVAES struct {
	//instance uintptr
	//// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	//ptr unsafe.Pointer
}

/**
 * Allocate an AVAES context.
 */
//struct AVAES *av_aes_alloc(void);
func AvAesAlloc() (res *AVAES, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_aes_alloc").Call()
	if err != nil {
		//return
	}
	//res = &AVAES{}
	res = (*AVAES)(unsafe.Pointer(t))
	//res.instance = t
	//res.ptr = unsafe.Pointer(t)
	return
}

/**
* Initialize an AVAES context.
* @param key_bits 128, 192 or 256
* @param decrypt 0 for encryption, 1 for decryption
 */
//int av_aes_init(struct AVAES *a, const uint8_t *key, int key_bits, int decrypt);
func (a *AVAES) AvAesInit(key *ffcommon.FUint8T, key_bits ffcommon.FInt, decrypt ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_aes_init").Call(
		uintptr(unsafe.Pointer(a)),
		uintptr(unsafe.Pointer(key)),
		uintptr(key_bits),
		uintptr(decrypt),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
* Encrypt or decrypt a buffer using a previously initialized context.
* @param count number of 16 byte blocks
* @param dst destination array, can be equal to src
* @param src source array, can be equal to dst
* @param iv initialization vector for CBC mode, if NULL then ECB will be used
* @param decrypt 0 for encryption, 1 for decryption
 */
//void av_aes_crypt(struct AVAES *a, uint8_t *dst, const uint8_t *src, int count, uint8_t *iv, int decrypt);
//测试失败
func (a *AVAES) AvAesCrypt(dst *ffcommon.FUint8T, src *ffcommon.FUint8T, count ffcommon.FInt, iv *ffcommon.FUint8T, decrypt ffcommon.FInt) (err error) {
	_, _, err = ffcommon.GetAvutilDll().NewProc("av_aes_crypt").Call(
		uintptr(unsafe.Pointer(a)),
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
