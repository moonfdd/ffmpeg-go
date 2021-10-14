package libavutil

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"unsafe"
)

type AVBlowfish struct {
	p [ffconstant.AV_BF_ROUNDS + 2]ffcommon.FUint32T
	s [4][256]ffcommon.FUint32T
}

/**
 * Allocate an AVBlowfish context.
 */
//AVBlowfish *av_blowfish_alloc(void);
//未测试
func AvBlowfishAlloc() (res *AVBlowfish, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_blowfish_alloc").Call()
	if err != nil {
		//return
	}
	//res = &AVAES{}
	res = (*AVBlowfish)(unsafe.Pointer(t))
	//res.instance = t
	//res.ptr = unsafe.Pointer(t)
	return
}

/**
 * Initialize an AVBlowfish context.
 *
 * @param ctx an AVBlowfish context
 * @param key a key
 * @param key_len length of the key
 */
//void av_blowfish_init(struct AVBlowfish *ctx, const uint8_t *key, int key_len);
//未测试
func (d *AVBlowfish) AvBlowfishInit(key *ffcommon.FUint8T, key_bits ffcommon.FInt, key_len ffcommon.FInt) (err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_blowfish_init").Call(
		uintptr(unsafe.Pointer(d)),
		uintptr(unsafe.Pointer(key)),
		uintptr(key_bits),
		uintptr(key_len),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Encrypt or decrypt a buffer using a previously initialized context.
 *
 * @param ctx an AVBlowfish context
 * @param xl left four bytes halves of input to be encrypted
 * @param xr right four bytes halves of input to be encrypted
 * @param decrypt 0 for encryption, 1 for decryption
 */
//void av_blowfish_crypt_ecb(struct AVBlowfish *ctx, uint32_t *xl, uint32_t *xr,
//int decrypt);
//未测试
func (d *AVBlowfish) AvBlowfishCryptEcb(xl *ffcommon.FUint32T, xr *ffcommon.FUint32T,
	decrypt ffcommon.FInt) (err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_blowfish_crypt_ecb").Call(
		uintptr(unsafe.Pointer(d)),
		uintptr(unsafe.Pointer(xl)),
		uintptr(unsafe.Pointer(xr)),
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
 * Encrypt or decrypt a buffer using a previously initialized context.
 *
 * @param ctx an AVBlowfish context
 * @param dst destination array, can be equal to src
 * @param src source array, can be equal to dst
 * @param count number of 8 byte blocks
 * @param iv initialization vector for CBC mode, if NULL ECB will be used
 * @param decrypt 0 for encryption, 1 for decryption
 */
//void av_blowfish_crypt(struct AVBlowfish *ctx, uint8_t *dst, const uint8_t *src,
//int count, uint8_t *iv, int decrypt);
//未测试
func (d *AVBlowfish) AvBlowfishCrypt(dst *ffcommon.FUint8T, src *ffcommon.FUint8T,
	count ffcommon.FInt, iv *ffcommon.FUint8T, decrypt ffcommon.FInt) (err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_blowfish_crypt").Call(
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
	if t == 0 {

	}
	return
}
