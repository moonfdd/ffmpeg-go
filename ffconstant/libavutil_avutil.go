package ffconstant

/**
 * @}
 */

/**
 * @addtogroup lavu_media Media Type
 * @brief Media Type
 */
type AVMediaType int32

const (
	AVMEDIA_TYPE_UNKNOWN = -1 ///< Usually treated as AVMEDIA_TYPE_DATA
	AVMEDIA_TYPE_VIDEO
	AVMEDIA_TYPE_AUDIO
	AVMEDIA_TYPE_DATA ///< Opaque data information usually continuous
	AVMEDIA_TYPE_SUBTITLE
	AVMEDIA_TYPE_ATTACHMENT ///< Opaque data information usually sparse
	AVMEDIA_TYPE_NB
)

/**
 * @defgroup lavu_const ffconstants
 * @{
 *
 * @defgroup lavu_enc Encoding specific
 *
 * @note those definition should move to avcodec
 * @{
 */

const FF_LAMBDA_SHIFT = 7
const FF_LAMBDA_SCALE = (1 << FF_LAMBDA_SHIFT)
const FF_QP2LAMBDA = 118 ///< factor to convert from H.263 QP to lambda
const FF_LAMBDA_MAX = (256*128 - 1)

const FF_QUALITY_SCALE = FF_LAMBDA_SCALE //FIXME maybe remove

/**
 * @}
 * @defgroup lavu_time Timestamp specific
 *
 * FFmpeg internal timebase and timestamp definitions
 *
 * @{
 */

/**
 * @brief Undefined timestamp value
 *
 * Usually reported by demuxer that work on containers that do not provide
 * either pts or dts.
 */

const AV_NOPTS_VALUE = 0x8000000000000000

/**
 * Internal time base represented as integer
 */

const AV_TIME_BASE = 1000000

/**
 * Internal time base represented as fractional value
 */

//const AV_TIME_BASE_Q       =   (AVRational){1, AV_TIME_BASE}

/**
 * @}
 * @}
 * @defgroup lavu_picture Image related
 *
 * AVPicture types, pixel formats and basic image planes manipulation.
 *
 * @{
 */
type AVPictureType int32

const (
	AV_PICTURE_TYPE_NONE = 0 ///< Undefined
	AV_PICTURE_TYPE_I        ///< Intra
	AV_PICTURE_TYPE_P        ///< Predicted
	AV_PICTURE_TYPE_B        ///< Bi-dir predicted
	AV_PICTURE_TYPE_S        ///< S(GMC)-VOP MPEG-4
	AV_PICTURE_TYPE_SI       ///< Switching Intra
	AV_PICTURE_TYPE_SP       ///< Switching Predicted
	AV_PICTURE_TYPE_BI       ///< BI type
)

const AV_FOURCC_MAX_STRING_SIZE = 32
