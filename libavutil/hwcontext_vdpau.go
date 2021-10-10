package libavutil

/**
 * @file
 * An API-specific header for AV_HWDEVICE_TYPE_VDPAU.
 *
 * This API supports dynamic frame pools. AVHWFramesContext.pool must return
 * AVBufferRefs whose data pointer is a VdpVideoSurface.
 */

/**
 * This struct is allocated as AVHWDeviceContext.hwctx
 */
type AVVDPAUDeviceContext struct {
	//VdpDevice         device
	//VdpGetProcAddress *get_proc_address
}

/**
 * AVHWFramesContext.hwctx is currently not used
 */
