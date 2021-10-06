package ffconstant

type AVClassCategory int32

const (
	AV_CLASS_CATEGORY_NA = 0
	AV_CLASS_CATEGORY_INPUT
	AV_CLASS_CATEGORY_OUTPUT
	AV_CLASS_CATEGORY_MUXER
	AV_CLASS_CATEGORY_DEMUXER
	AV_CLASS_CATEGORY_ENCODER
	AV_CLASS_CATEGORY_DECODER
	AV_CLASS_CATEGORY_FILTER
	AV_CLASS_CATEGORY_BITSTREAM_FILTER
	AV_CLASS_CATEGORY_SWSCALER
	AV_CLASS_CATEGORY_SWRESAMPLER
	AV_CLASS_CATEGORY_DEVICE_VIDEO_OUTPUT = 40
	AV_CLASS_CATEGORY_DEVICE_VIDEO_INPUT
	AV_CLASS_CATEGORY_DEVICE_AUDIO_OUTPUT
	AV_CLASS_CATEGORY_DEVICE_AUDIO_INPUT
	AV_CLASS_CATEGORY_DEVICE_OUTPUT
	AV_CLASS_CATEGORY_DEVICE_INPUT
	AV_CLASS_CATEGORY_NB ///< not part of ABI/API
)

//#define AV_IS_INPUT_DEVICE(category) \
//(((category) == AV_CLASS_CATEGORY_DEVICE_VIDEO_INPUT) || \
//((category) == AV_CLASS_CATEGORY_DEVICE_AUDIO_INPUT) || \
//((category) == AV_CLASS_CATEGORY_DEVICE_INPUT))
//
//#define AV_IS_OUTPUT_DEVICE(category) \
//(((category) == AV_CLASS_CATEGORY_DEVICE_VIDEO_OUTPUT) || \
//((category) == AV_CLASS_CATEGORY_DEVICE_AUDIO_OUTPUT) || \
//((category) == AV_CLASS_CATEGORY_DEVICE_OUTPUT))

/**
 * Print no output.
 */
const AV_LOG_QUIET = -8

/**
 * Something went really wrong and we will crash now.
 */
const AV_LOG_PANIC = 0

/**
 * Something went wrong and recovery is not possible.
 * For example, no header was found for a format which depends
 * on headers or an illegal combination of parameters is used.
 */
const AV_LOG_FATAL = 8

/**
 * Something went wrong and cannot losslessly be recovered.
 * However, not all future data is affected.
 */
const AV_LOG_ERROR = 16

/**
 * Something somehow does not look correct. This may or may not
 * lead to problems. An example would be the use of '-vstrict -2'.
 */
const AV_LOG_WARNING = 24

/**
 * Standard information.
 */
const AV_LOG_INFO = 32

/**
 * Detailed information.
 */
const AV_LOG_VERBOSE = 40

/**
 * Stuff which is only useful for libav* developers.
 */
const AV_LOG_DEBUG = 48

/**
 * Extremely verbose debugging, useful for libav* development.
 */
const AV_LOG_TRACE = 56

const AV_LOG_MAX_OFFSET = (AV_LOG_TRACE - AV_LOG_QUIET)

/**
 * @}
 */

/**
  * Sets additional colors for extended debugging sessions.
  * @code
    av_log(ctx, AV_LOG_DEBUG|AV_LOG_C(134), "Message in purple\n");
    @endcode
  * Requires 256color terminal support. Uses outside debugging is not
  * recommended.
*/
//#define AV_LOG_C(x) ((x) << 8)

/**
 * Skip repeated messages, this requires the user app to use av_log() instead of
 * (f)printf as the 2 would otherwise interfere and lead to
 * "Last message repeated x times" messages below (f)printf messages with some
 * bad luck.
 * Also to receive the last, "last repeated" line if any, the user app must
 * call av_log(NULL, AV_LOG_QUIET, "%s", ""); at the end
 */
const AV_LOG_SKIP_REPEATED = 1

/**
 * Include the log severity in messages originating from codecs.
 *
 * Results in messages such as:
 * [rawvideo @ 0xDEADBEEF] [error] encode did not produce valid pts
 */
const AV_LOG_PRINT_LEVEL = 2
