package ffconstant

/**
 * @file
 * API-specific header for AV_HWDEVICE_TYPE_VAAPI.
 *
 * Dynamic frame pools are supported, but note that any pool used as a render
 * target is required to be of fixed size in order to be be usable as an
 * argument to vaCreateContext().
 *
 * For user-allocated pools, AVHWFramesContext.pool must return AVBufferRefs
 * with the data pointer set to a VASurfaceID.
 */

const (
	/**
	 * The quirks field has been set by the user and should not be detected
	 * automatically by av_hwdevice_ctx_init().
	 */
	AV_VAAPI_DRIVER_QUIRK_USER_SET = (1 << 0)
	/**
	 * The driver does not destroy parameter buffers when they are used by
	 * vaRenderPicture().  Additional code will be required to destroy them
	 * separately afterwards.
	 */
	AV_VAAPI_DRIVER_QUIRK_RENDER_PARAM_BUFFERS = (1 << 1)

	/**
	 * The driver does not support the VASurfaceAttribMemoryType attribute,
	 * so the surface allocation code will not try to use it.
	 */
	AV_VAAPI_DRIVER_QUIRK_ATTRIB_MEMTYPE = (1 << 2)

	/**
	 * The driver does not support surface attributes at all.
	 * The surface allocation code will never pass them to surface allocation,
	 * and the results of the vaQuerySurfaceAttributes() call will be faked.
	 */
	AV_VAAPI_DRIVER_QUIRK_SURFACE_ATTRIBUTES = (1 << 3)
)
