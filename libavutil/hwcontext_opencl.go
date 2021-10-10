package libavutil

/**
 * @file
 * API-specific header for AV_HWDEVICE_TYPE_OPENCL.
 *
 * Pools allocated internally are always dynamic, and are primarily intended
 * to be used in OpenCL-only cases.  If interoperation is required, it is
 * typically required to allocate frames in the other API and then map the
 * frames context to OpenCL with av_hwframe_ctx_create_derived().
 */

/**
 * OpenCL frame descriptor for pool allocation.
 *
 * In user-allocated pools, AVHWFramesContext.pool must return AVBufferRefs
 * with the data pointer pointing at an object of this type describing the
 * planes of the frame.
 */
type AVOpenCLFrameDescriptor struct {
	///**
	// * Number of planes in the frame.
	// */
	//int nb_planes;
	///**
	// * OpenCL image2d objects for each plane of the frame.
	// */
	//cl_mem planes[AV_NUM_DATA_POINTERS];
}

/**
 * OpenCL device details.
 *
 * Allocated as AVHWDeviceContext.hwctx
 */
type AVOpenCLDeviceContext struct {
	///**
	// * The primary device ID of the device.  If multiple OpenCL devices
	// * are associated with the context then this is the one which will
	// * be used for all operations internal to FFmpeg.
	// */
	//cl_device_id device_id;
	///**
	// * The OpenCL context which will contain all operations and frames on
	// * this device.
	// */
	//cl_context context;
	///**
	// * The default command queue for this device, which will be used by all
	// * frames contexts which do not have their own command queue.  If not
	// * intialised by the user, a default queue will be created on the
	// * primary device.
	// */
	//cl_command_queue command_queue;
}

/**
 * OpenCL-specific data associated with a frame pool.
 *
 * Allocated as AVHWFramesContext.hwctx.
 */
type AVOpenCLFramesContext struct {
	///**
	// * The command queue used for internal asynchronous operations on this
	// * device (av_hwframe_transfer_data(), av_hwframe_map()).
	// *
	// * If this is not set, the command queue from the associated device is
	// * used instead.
	// */
	//cl_command_queue command_queue;
}
