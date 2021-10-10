package libavutil

import "ffmpeg-go/ffcommon"

/**
 * MediaCodec details.
 *
 * Allocated as AVHWDeviceContext.hwctx
 */
type AVMediaCodecDeviceContext struct {
	/**
	 * android/view/Surface handle, to be filled by the user.
	 *
	 * This is the default surface used by decoders on this device.
	 */
	surface ffcommon.FVoidP
}
