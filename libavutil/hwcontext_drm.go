package libavutil

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
)

/**
 * DRM object descriptor.
 *
 * Describes a single DRM object, addressing it as a PRIME file
 * descriptor.
 */
type AVDRMObjectDescriptor struct {
	/**
	 * DRM PRIME fd for the object.
	 */
	Fd ffcommon.FInt
	/**
	 * Total size of the object.
	 *
	 * (This includes any parts not which do not contain image data.)
	 */
	Size ffcommon.FSizeT
	/**
	 * Format modifier applied to the object (DRM_FORMAT_MOD_*).
	 *
	 * If the format modifier is unknown then this should be set to
	 * DRM_FORMAT_MOD_INVALID.
	 */
	FormatModifier ffcommon.FUint64T
}

/**
 * DRM plane descriptor.
 *
 * Describes a single plane of a layer, which is contained within
 * a single object.
 */
type AVDRMPlaneDescriptor struct {
	/**
	 * Index of the object containing this plane in the objects
	 * array of the enclosing frame descriptor.
	 */
	ObjectIndex ffcommon.FInt
	/**
	 * Offset within that object of this plane.
	 */
	Offset ffcommon.FPtrdiffT
	/**
	 * Pitch (linesize) of this plane.
	 */
	Pitch ffcommon.FPtrdiffT
}

/**
 * DRM layer descriptor.
 *
 * Describes a single layer within a frame.  This has the structure
 * defined by its format, and will contain one or more planes.
 */
type AVDRMLayerDescriptor struct {
	/**
	 * Format of the layer (DRM_FORMAT_*).
	 */
	format ffcommon.FUint32T
	/**
	 * Number of planes in the layer.
	 *
	 * This must match the number of planes required by format.
	 */
	nb_planes ffcommon.FInt
	/**
	 * Array of planes in this layer.
	 */
	planes [ffconstant.AV_DRM_MAX_PLANES]AVDRMPlaneDescriptor
}

/**
 * DRM frame descriptor.
 *
 * This is used as the data pointer for AV_PIX_FMT_DRM_PRIME frames.
 * It is also used by user-allocated frame pools - allocating in
 * AVHWFramesContext.pool must return AVBufferRefs which contain
 * an object of this type.
 *
 * The fields of this structure should be set such it can be
 * imported directly by EGL using the EGL_EXT_image_dma_buf_import
 * and EGL_EXT_image_dma_buf_import_modifiers extensions.
 * (Note that the exact layout of a particular format may vary between
 * platforms - we only specify that the same platform should be able
 * to import it.)
 *
 * The total number of planes must not exceed AV_DRM_MAX_PLANES, and
 * the order of the planes by increasing layer index followed by
 * increasing plane index must be the same as the order which would
 * be used for the data pointers in the equivalent software format.
 */
type AVDRMFrameDescriptor struct {
	/**
	 * Number of DRM objects making up this frame.
	 */
	NbObjects ffcommon.FInt
	/**
	 * Array of objects making up the frame.
	 */
	Objects [ffconstant.AV_DRM_MAX_PLANES]AVDRMObjectDescriptor
	/**
	 * Number of layers in the frame.
	 */
	NbLayers ffcommon.FInt
	/**
	 * Array of layers in the frame.
	 */
	Layers [ffconstant.AV_DRM_MAX_PLANES]AVDRMLayerDescriptor
}

/**
 * DRM device.
 *
 * Allocated as AVHWDeviceContext.hwctx.
 */
type AVDRMDeviceContext struct {
	/**
	 * File descriptor of DRM device.
	 *
	 * This is used as the device to create frames on, and may also be
	 * used in some derivation and mapping operations.
	 *
	 * If no device is required, set to -1.
	 */
	Fd ffcommon.FInt
}
