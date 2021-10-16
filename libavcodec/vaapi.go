package libavcodec

import "ffmpeg-go/ffcommon"

/**
 * @defgroup lavc_codec_hwaccel_vaapi VA API Decoding
 * @ingroup lavc_codec_hwaccel
 * @{
 */

/**
 * This structure is used to share data between the FFmpeg library and
 * the client video application.
 * This shall be zero-allocated and available as
 * AVCodecContext.hwaccel_context. All user members can be set once
 * during initialization or through each AVCodecContext.get_buffer()
 * function call. In any case, they must be valid prior to calling
 * decoding functions.
 *
 * Deprecated: use AVCodecContext.hw_frames_ctx instead.
 */
type VaapiContext struct {
	/**
	 * Window system dependent data
	 *
	 * - encoding: unused
	 * - decoding: Set by user
	 */
	Display ffcommon.FVoidP

	/**
	 * Configuration ID
	 *
	 * - encoding: unused
	 * - decoding: Set by user
	 */
	ConfigId ffcommon.FUint32T

	/**
	 * Context ID (video decode pipeline)
	 *
	 * - encoding: unused
	 * - decoding: Set by user
	 */
	ContextId ffcommon.FUint32T
}
