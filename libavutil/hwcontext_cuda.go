package libavutil

/**
 * @file
 * An API-specific header for AV_HWDEVICE_TYPE_CUDA.
 *
 * This API supports dynamic frame pools. AVHWFramesContext.pool must return
 * AVBufferRefs whose data pointer is a CUdeviceptr.
 */

type AVCUDADeviceContextInternal struct {
}

/**
 * This struct is allocated as AVHWDeviceContext.hwctx
 */
type AVCUDADeviceContext struct {
	//CUcontext                   cuda_ctx
	//CUstream                    stream
	//AVCUDADeviceContextInternal *internal
}
