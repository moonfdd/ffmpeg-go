package libavutil

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

/**
 * @defgroup lavu_des DES
 * @ingroup lavu_crypto
 * @{
 */

type AVDES struct {
	round_keys [3][16]ffcommon.FUint32T
	triple_des ffcommon.FInt
}

/**
 * Allocate an AVDES context.
 */
//AVDES *av_des_alloc(void);
//未测试
func AvDesAlloc() (res *AVDES, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_des_alloc").Call()
	if err != nil {
		//return
	}
	//res = &AVAES{}
	res = (*AVDES)(unsafe.Pointer(t))
	//res.instance = t
	//res.ptr = unsafe.Pointer(t)
	return
}

/**
 * @brief Initializes an AVDES context.
 *
 * @param key_bits must be 64 or 192
 * @param decrypt 0 for encryption/CBC-MAC, 1 for decryption
 * @return zero on success, negative value otherwise
 */
//int av_des_init(struct AVDES *d, const uint8_t *key, int key_bits, int decrypt);
//未测试
func (d *AVDES) AvDesInit(key *ffcommon.FUint8T, key_bits ffcommon.FInt, decrypt ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_des_init").Call(
		uintptr(unsafe.Pointer(d)),
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
 * @brief Encrypts / decrypts using the DES algorithm.
 *
 * @param count number of 8 byte blocks
 * @param dst destination array, can be equal to src, must be 8-byte aligned
 * @param src source array, can be equal to dst, must be 8-byte aligned, may be NULL
 * @param iv initialization vector for CBC mode, if NULL then ECB will be used,
 *           must be 8-byte aligned
 * @param decrypt 0 for encryption, 1 for decryption
 */
//void av_des_crypt(struct AVDES *d, uint8_t *dst, const uint8_t *src, int count, uint8_t *iv, int decrypt);
//未测试
func (d *AVDES) AvDesCrypt(dst *ffcommon.FUint8T, src *ffcommon.FUint8T, count ffcommon.FInt, iv *ffcommon.FUint8T, decrypt ffcommon.FInt) (err error) {
	_, _, err = ffcommon.GetAvutilDll().NewProc("av_des_crypt").Call(
		uintptr(unsafe.Pointer(d)),
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

/**
 * @brief Calculates CBC-MAC using the DES algorithm.
 *
 * @param count number of 8 byte blocks
 * @param dst destination array, can be equal to src, must be 8-byte aligned
 * @param src source array, can be equal to dst, must be 8-byte aligned, may be NULL
 */
//void av_des_mac(struct AVDES *d, uint8_t *dst, const uint8_t *src, int count);
//未测试
func (d *AVDES) AvDesMac(dst *ffcommon.FUint8T, src *ffcommon.FUint8T, count ffcommon.FInt) (err error) {
	_, _, err = ffcommon.GetAvutilDll().NewProc("av_des_mac").Call(
		uintptr(unsafe.Pointer(d)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
		uintptr(count),
	)
	if err != nil {
		//return
	}
	return
}
