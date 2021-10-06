package ffconstant

/**
 * @defgroup lavu_hmac HMAC
 * @ingroup lavu_crypto
 * @{
 */
type AVHMACType int32

const (
	AV_HMAC_MD5 = 0
	AV_HMAC_SHA1
	AV_HMAC_SHA224
	AV_HMAC_SHA256
	AV_HMAC_SHA384
	AV_HMAC_SHA512
)
