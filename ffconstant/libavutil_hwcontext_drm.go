package ffconstant

/**
 * @file
 * API-specific header for AV_HWDEVICE_TYPE_DRM.
 *
 * Internal frame allocation is not currently supported - all frames
 * must be allocated by the user.  Thus AVHWFramesContext is always
 * NULL, though this may change if support for frame allocation is
 * added in future.
 */
const AV_DRM_MAX_PLANES = 4
