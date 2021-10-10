package libavutil

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
)

/**
 * @file
 * An API-specific header for AV_HWDEVICE_TYPE_VIDEOTOOLBOX.
 *
 * This API currently does not support frame allocation, as the raw VideoToolbox
 * API does allocation, and FFmpeg itself never has the need to allocate frames.
 *
 * If the API user sets a custom pool, AVHWFramesContext.pool must return
 * AVBufferRefs whose data pointer is a CVImageBufferRef or CVPixelBufferRef.
 *
 * Currently AVHWDeviceContext.hwctx and AVHWFramesContext.hwctx are always
 * NULL.
 */

/**
 * Convert a VideoToolbox (actually CoreVideo) format to AVPixelFormat.
 * Returns AV_PIX_FMT_NONE if no known equivalent was found.
 */
//enum AVPixelFormat av_map_videotoolbox_format_to_pixfmt(uint32_t cv_fmt);
//未测试
func AvMapVideotoolboxFormatToPixfmt(cv_fmt ffcommon.FUint32T) (res ffconstant.AVPixelFormat, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_map_videotoolbox_format_to_pixfmt").Call(
		uintptr(cv_fmt),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffconstant.AVPixelFormat(t)
	return
}

/**
 * Convert an AVPixelFormat to a VideoToolbox (actually CoreVideo) format.
 * Returns 0 if no known equivalent was found.
 */
//uint32_t av_map_videotoolbox_format_from_pixfmt(enum AVPixelFormat pix_fmt);
//未测试
func AvMapVideotoolboxFormatFromPixfmt(pix_fmt ffconstant.AVPixelFormat) (res ffcommon.FUint32T, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_map_videotoolbox_format_from_pixfmt").Call(
		uintptr(pix_fmt),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FUint32T(t)
	return
}

/**
 * Same as av_map_videotoolbox_format_from_pixfmt function, but can map and
 * return full range pixel formats via a flag.
 */
//uint32_t av_map_videotoolbox_format_from_pixfmt2(enum AVPixelFormat pix_fmt, bool full_range);
//未测试
func AvMapVideotoolboxFormatFromPixfmt2(pix_fmt ffconstant.AVPixelFormat, full_range bool) (res ffcommon.FUint32T, err error) {
	var t uintptr
	full := 0
	if full_range {
		full = 1
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_map_videotoolbox_format_from_pixfmt2").Call(
		uintptr(pix_fmt),
		uintptr(full),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FUint32T(t)
	return
}
