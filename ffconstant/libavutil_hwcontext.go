package ffconstant

type AVHWDeviceType int32

const (
	AV_HWDEVICE_TYPE_NONE = 0
	AV_HWDEVICE_TYPE_VDPAU
	AV_HWDEVICE_TYPE_CUDA
	AV_HWDEVICE_TYPE_VAAPI
	AV_HWDEVICE_TYPE_DXVA2
	AV_HWDEVICE_TYPE_QSV
	AV_HWDEVICE_TYPE_VIDEOTOOLBOX
	AV_HWDEVICE_TYPE_D3D11VA
	AV_HWDEVICE_TYPE_DRM
	AV_HWDEVICE_TYPE_OPENCL
	AV_HWDEVICE_TYPE_MEDIACODEC
	AV_HWDEVICE_TYPE_VULKAN
)

type AVHWFrameTransferDirection int32

const (
	/**
	 * Transfer the data from the queried hw frame.
	 */
	AV_HWFRAME_TRANSFER_DIRECTION_FROM = 0

	/**
	 * Transfer the data to the queried hw frame.
	 */
	AV_HWFRAME_TRANSFER_DIRECTION_TO
)

/**
 * Flags to apply to frame mappings.
 */
const (
	/**
	 * The mapping must be readable.
	 */
	AV_HWFRAME_MAP_READ = 1 << 0
	/**
	 * The mapping must be writeable.
	 */
	AV_HWFRAME_MAP_WRITE = 1 << 1
	/**
	 * The mapped frame will be overwritten completely in subsequent
	 * operations, so the current frame data need not be loaded.  Any values
	 * which are not overwritten are unspecified.
	 */
	AV_HWFRAME_MAP_OVERWRITE = 1 << 2
	/**
	 * The mapping must be direct.  That is, there must not be any copying in
	 * the map or unmap steps.  Note that performance of direct mappings may
	 * be much lower than normal memory.
	 */
	AV_HWFRAME_MAP_DIRECT = 1 << 3
)
