package libavutil

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"unsafe"
)

type AVHMAC struct {
}

/**
 * Allocate an AVHMAC context.
 * @param type The hash function used for the HMAC.
 */
//AVHMAC *av_hmac_alloc(enum AVHMACType type);
//未测试
func AvHmacAlloc(type0 ffconstant.AVHMACType) (res *AVHMAC, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_hmac_alloc").Call(
		uintptr(type0),
	)
	if err != nil {
		//return
	}
	res = (*AVHMAC)(unsafe.Pointer(t))
	return
}

/**
 * Free an AVHMAC context.
 * @param ctx The context to free, may be NULL
 */
//void av_hmac_free(AVHMAC *ctx);
//未测试
func (res *AVMD5) AvHmacFree() (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_hmac_free").Call()
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Initialize an AVHMAC context with an authentication key.
 * @param ctx    The HMAC context
 * @param key    The authentication key
 * @param keylen The length of the key, in bytes
 */
//void av_hmac_init(AVHMAC *ctx, const uint8_t *key, unsigned int keylen);
//未测试
func (ctx *AVHMAC) AvHmacInit(key *ffcommon.FUint8T, keylen ffcommon.FUnsignedInt) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_hmac_init").Call(
		uintptr(unsafe.Pointer(key)),
		uintptr(keylen),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Hash data with the HMAC.
 * @param ctx  The HMAC context
 * @param data The data to hash
 * @param len  The length of the data, in bytes
 */
//void av_hmac_update(AVHMAC *ctx, const uint8_t *data, unsigned int len);
//未测试
func (ctx *AVHMAC) AvHmacUpdate(data *ffcommon.FUint8T, len0 ffcommon.FUnsignedInt) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_hmac_update").Call(
		uintptr(unsafe.Pointer(ctx)),
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
 * Finish hashing and output the HMAC digest.
 * @param ctx    The HMAC context
 * @param out    The output buffer to write the digest into
 * @param outlen The length of the out buffer, in bytes
 * @return       The number of bytes written to out, or a negative error code.
 */
//int av_hmac_final(AVHMAC *ctx, uint8_t *out, unsigned int outlen);
//未测试
func (ctx *AVHMAC) AvHmacFinal(out *ffcommon.FUint8T, outlen ffcommon.FUnsignedInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_hmac_final").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(out)),
		uintptr(outlen),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Hash an array of data with a key.
 * @param ctx    The HMAC context
 * @param data   The data to hash
 * @param len    The length of the data, in bytes
 * @param key    The authentication key
 * @param keylen The length of the key, in bytes
 * @param out    The output buffer to write the digest into
 * @param outlen The length of the out buffer, in bytes
 * @return       The number of bytes written to out, or a negative error code.
 */
//int av_hmac_calc(AVHMAC *ctx, const uint8_t *data, unsigned int len,
//const uint8_t *key, unsigned int keylen,
//uint8_t *out, unsigned int outlen);
//未测试
func (ctx *AVHMAC) AvHmacCalc(data *ffcommon.FUint8T, len0 ffcommon.FUnsignedInt,
	key *ffcommon.FUint8T, keylen ffcommon.FUnsignedInt,
	out *ffcommon.FUint8T, outlen ffcommon.FUnsignedInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_hmac_calc").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(data)),
		uintptr(len0),
		uintptr(unsafe.Pointer(key)),
		uintptr(keylen),
		uintptr(unsafe.Pointer(out)),
		uintptr(outlen),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}
