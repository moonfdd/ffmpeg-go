package ffconstant

/**
 * @file
 * audio downmix medatata
 */

/**
 * @addtogroup lavu_audio
 * @{
 */

/**
 * @defgroup downmix_info Audio downmix metadata
 * @{
 */

/**
 * Possible downmix types.
 */
type AVDownmixType int32

const (
	AV_DOWNMIX_TYPE_UNKNOWN = 0 /**< Not indicated. */
	AV_DOWNMIX_TYPE_LORO        /**< Lo/Ro 2-channel downmix (Stereo). */
	AV_DOWNMIX_TYPE_LTRT        /**< Lt/Rt 2-channel downmix, Dolby Surround compatible. */
	AV_DOWNMIX_TYPE_DPLII       /**< Lt/Rt 2-channel downmix, Dolby Pro Logic II compatible. */
	AV_DOWNMIX_TYPE_NB          /**< Number of downmix types. Not part of ABI. */
)
