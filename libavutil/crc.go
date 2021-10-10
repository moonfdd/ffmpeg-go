package libavutil

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"unsafe"
)

/**
 * @defgroup lavu_crc32 CRC
 * @ingroup lavu_hash
 * CRC (Cyclic Redundancy Check) hash function implementation.
 *
 * This module supports numerous CRC polynomials, in addition to the most
 * widely used CRC-32-IEEE. See @ref AVCRCId for a list of available
 * polynomials.
 *
 * @{
 */

type AVCRC ffcommon.FUint32T

/**
 * Initialize a CRC table.
 * @param ctx must be an array of size sizeof(AVCRC)*257 or sizeof(AVCRC)*1024
 * @param le If 1, the lowest bit represents the coefficient for the highest
 *           exponent of the corresponding polynomial (both for poly and
 *           actual CRC).
 *           If 0, you must swap the CRC parameter and the result of av_crc
 *           if you need the standard representation (can be simplified in
 *           most cases to e.g. bswap16):
 *           av_bswap32(crc << (32-bits))
 * @param bits number of bits for the CRC
 * @param poly generator polynomial without the x**bits coefficient, in the
 *             representation as specified by le
 * @param ctx_size size of ctx in bytes
 * @return <0 on failure
 */
//int av_crc_init(AVCRC *ctx, int le, int bits, uint32_t poly, int ctx_size);
//未测试
func (ctx *AVCRC) AvCrcInit(le ffcommon.FInt, bits ffcommon.FInt, poly ffcommon.FUint32T, ctx_size ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_crc_init").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(bits),
		uintptr(poly),
		uintptr(ctx_size),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
* Get an initialized standard CRC table.
* @param crc_id ID of a standard CRC
* @return a pointer to the CRC table or NULL on failure
 */
//const AVCRC *av_crc_get_table(AVCRCId crc_id);
//未测试
func AvCrcGetTable(crc_id ffconstant.AVCRCId) (res *AVCRC, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_crc_get_table").Call(
		uintptr(crc_id),
	)
	if err != nil {
		//return
	}
	res = (*AVCRC)(unsafe.Pointer(t))
	return
}

/**
* Calculate the CRC of a block.
* @param crc CRC of previous blocks if any or initial value for CRC
* @return CRC updated with the data from the given block
*
* @see av_crc_init() "le" parameter
 */
//uint32_t av_crc(const AVCRC *ctx, uint32_t crc,
//const uint8_t *buffer, size_t length) av_pure;
//未测试
func (ctx *AVCRC) AvCrc(crc ffcommon.FUint32T,
	buffer *ffcommon.FUint8T, length ffcommon.FSizeT) (res ffcommon.FUint32T, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_crc").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(crc),
		uintptr(unsafe.Pointer(buffer)),
		uintptr(length),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FUint32T(t)
	return
}
