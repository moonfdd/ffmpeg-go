package libavutil

/**
 * VAAPI connection details.
 *
 * Allocated as AVHWDeviceContext.hwctx
 */
type AVVAAPIDeviceContext struct {
	///**
	// * The VADisplay handle, to be filled by the user.
	// */
	//VADisplay display;
	///**
	// * Driver quirks to apply - this is filled by av_hwdevice_ctx_init(),
	// * with reference to a table of known drivers, unless the
	// * AV_VAAPI_DRIVER_QUIRK_USER_SET bit is already present.  The user
	// * may need to refer to this field when performing any later
	// * operations using VAAPI with the same VADisplay.
	// */
	//unsigned int driver_quirks;
	//} AVVAAPIDeviceContext;
	//
	///**
	// * VAAPI-specific data associated with a frame pool.
	// *
	// * Allocated as AVHWFramesContext.hwctx.
	// */
	//typedef struct AVVAAPIFramesContext {
	///**
	// * Set by the user to apply surface attributes to all surfaces in
	// * the frame pool.  If null, default settings are used.
	// */
	//VASurfaceAttrib *attributes;
	//int           nb_attributes;
	///**
	// * The surfaces IDs of all surfaces in the pool after creation.
	// * Only valid if AVHWFramesContext.initial_pool_size was positive.
	// * These are intended to be used as the render_targets arguments to
	// * vaCreateContext().
	// */
	//VASurfaceID     *surface_ids;
	//int           nb_surfaces;
}

/**
 * VAAPI hardware pipeline configuration details.
 *
 * Allocated with av_hwdevice_hwconfig_alloc().
 */
type AVVAAPIHWConfig struct {
	///**
	// * ID of a VAAPI pipeline configuration.
	// */
	//VAConfigID config_id;
}
