package libavutil

/**
 * @file
 * An API-specific header for AV_HWDEVICE_TYPE_QSV.
 *
 * This API does not support dynamic frame pools. AVHWFramesContext.pool must
 * contain AVBufferRefs whose data pointer points to an mfxFrameSurface1 struct.
 */

/**
 * This struct is allocated as AVHWDeviceContext.hwctx
 */
type AVQSVDeviceContext struct {
	//mfxSession session;
}

/**
 * This struct is allocated as AVHWFramesContext.hwctx
 */
type AVQSVFramesContext struct {
	//mfxFrameSurface1 *surfaces;
	//int            nb_surfaces;
	//
	///**
	// * A combination of MFX_MEMTYPE_* describing the frame pool.
	// */
	//int frame_type;
}
