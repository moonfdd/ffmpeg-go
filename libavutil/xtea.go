package libavutil

import (
	"ffmpeg-go/ffcommon"
	"fmt"
	"unsafe"
)

/**
 * @file
 * @brief Public header for libavutil XTEA algorithm
 * @defgroup lavu_xtea XTEA
 * @ingroup lavu_crypto
 * @{
 */

type AVXTEA struct {
	Key [16]ffcommon.FUint32T
}

/**
 * Allocate an AVXTEA context.
 */
//AVXTEA *av_xtea_alloc(void);
func AvXteaAlloc() (res *AVXTEA, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_xtea_alloc").Call()
	if err != nil {
		//return
	}
	res = (*AVXTEA)(unsafe.Pointer(t))
	return
}

/**
* Initialize an AVXTEA context.
*
* @param ctx an AVXTEA context
* @param key a key of 16 bytes used for encryption/decryption,
*            interpreted as big endian 32 bit numbers
 */
//void av_xtea_init(struct AVXTEA *ctx, const uint8_t key[16]);
func (ctx *AVXTEA) AvXteaInit(key [16]ffcommon.FUint8T) (err error) {
	fmt.Println("key = ", key)
	_, _, err = ffcommon.GetAvutilDll().NewProc("av_xtea_init").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(&key)),
	)
	if err != nil {
		//return
	}
	return
}

/**
* Initialize an AVXTEA context.
*
* @param ctx an AVXTEA context
* @param key a key of 16 bytes used for encryption/decryption,
*            interpreted as little endian 32 bit numbers
 */
//void av_xtea_le_init(struct AVXTEA *ctx, const uint8_t key[16]);
func (ctx *AVXTEA) AvXteaLeInit(key [16]ffcommon.FUint8T) (err error) {
	fmt.Println("key = ", key)
	_, _, err = ffcommon.GetAvutilDll().NewProc("av_xtea_le_init").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(&key)),
	)
	if err != nil {
		//return
	}
	return
}

/**
* Encrypt or decrypt a buffer using a previously initialized context,
* in big endian format.
*
* @param ctx an AVXTEA context
* @param dst destination array, can be equal to src
* @param src source array, can be equal to dst
* @param count number of 8 byte blocks
* @param iv initialization vector for CBC mode, if NULL then ECB will be used
* @param decrypt 0 for encryption, 1 for decryption
 */
//void av_xtea_crypt(struct AVXTEA *ctx, uint8_t *dst, const uint8_t *src,
//int count, uint8_t *iv, int decrypt);
func (ctx *AVXTEA) AvXteaCrypt(dst *ffcommon.FUint8T, src *ffcommon.FUint8T,
	count ffcommon.FInt, iv *ffcommon.FUint8T, decrypt ffcommon.FInt) (err error) {
	_, _, err = ffcommon.GetAvutilDll().NewProc("av_xtea_crypt").Call(
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

//
///**
// * Encrypt or decrypt a buffer using a previously initialized context,
// * in little endian format.
// *
// * @param ctx an AVXTEA context
// * @param dst destination array, can be equal to src
// * @param src source array, can be equal to dst
// * @param count number of 8 byte blocks
// * @param iv initialization vector for CBC mode, if NULL then ECB will be used
// * @param decrypt 0 for encryption, 1 for decryption
// */
//void av_xtea_le_crypt(struct AVXTEA *ctx, uint8_t *dst, const uint8_t *src,
//int count, uint8_t *iv, int decrypt);
func (ctx *AVXTEA) AvXteaLeCrypt(dst *ffcommon.FUint8T, src *ffcommon.FUint8T,
	count ffcommon.FInt, iv *ffcommon.FUint8T, decrypt ffcommon.FInt) (err error) {
	_, _, err = ffcommon.GetAvutilDll().NewProc("av_xtea_le_crypt").Call(
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
