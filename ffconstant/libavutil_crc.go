package ffconstant

type AVCRCId uint32

const (
	AV_CRC_8_ATM = 0
	AV_CRC_16_ANSI
	AV_CRC_16_CCITT
	AV_CRC_32_IEEE
	AV_CRC_32_IEEE_LE /*< reversed bitorder version of AV_CRC_32_IEEE */
	AV_CRC_16_ANSI_LE /*< reversed bitorder version of AV_CRC_16_ANSI */
	AV_CRC_24_IEEE
	AV_CRC_8_EBU
	AV_CRC_MAX /*< Not part of public API! Do not use outside libavutil. */
)
