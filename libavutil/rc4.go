package libavutil

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

/**
 * @defgroup lavu_rc4 RC4
 * @ingroup lavu_crypto
 * @{
 */

type AVRC4 struct {
	state [256]ffcommon.FUint8T
	X, Y  ffcommon.FInt
}

/**
 * Allocate an AVRC4 context.
 */
//AVRC4 *av_rc4_alloc(void);
//未测试
func AvRc4Alloc() (res *AVRC4, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_rc4_alloc").Call()
	if err != nil {
		//return
	}
	res = (*AVRC4)(unsafe.Pointer(t))
	return
}

/**
 * @brief Initializes an AVRC4 context.
 *
 * @param key_bits must be a multiple of 8
 * @param decrypt 0 for encryption, 1 for decryption, currently has no effect
 * @return zero on success, negative value otherwise
 */
//int av_rc4_init(struct AVRC4 *d, const uint8_t *key, int key_bits, int decrypt);
//未测试
func (d *AVRC4) AvRc4Init(key *ffcommon.FUint8T, key_bits ffcommon.FInt, decrypt ffcommon.FInt) (res *AVRC4, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_rc4_init").Call(
		uintptr(unsafe.Pointer(d)),
		uintptr(unsafe.Pointer(key)),
		uintptr(key_bits),
		uintptr(decrypt),
	)
	if err != nil {
		//return
	}
	res = (*AVRC4)(unsafe.Pointer(t))
	return
}

/**
 * @brief Encrypts / decrypts using the RC4 algorithm.
 *
 * @param count number of bytes
 * @param dst destination array, can be equal to src
 * @param src source array, can be equal to dst, may be NULL
 * @param iv not (yet) used for RC4, should be NULL
 * @param decrypt 0 for encryption, 1 for decryption, not (yet) used
 */
//void av_rc4_crypt(struct AVRC4 *d, uint8_t *dst, const uint8_t *src, int count, uint8_t *iv, int decrypt);
//未测试
func (d *AVRC4) AvRc4Crypt(dst *ffcommon.FUint8T, src *ffcommon.FUint8T, count ffcommon.FInt, iv *ffcommon.FUint8T, decrypt ffcommon.FInt) (res *AVRC4, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_rc4_crypt").Call(
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
	res = (*AVRC4)(unsafe.Pointer(t))
	return
}
