package libavcodec

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

/**
 * This structure is used to provides the necessary configurations and data
 * to the Direct3D11 FFmpeg HWAccel implementation.
 *
 * The application must make it available as AVCodecContext.hwaccel_context.
 *
 * Use av_d3d11va_alloc_context() exclusively to allocate an AVD3D11VAContext.
 */
type AVD3D11VAContext struct {
	///**
	// * D3D11 decoder object
	// */
	//ID3D11VideoDecoder *decoder;
	//
	///**
	// * D3D11 VideoContext
	// */
	//ID3D11VideoContext *video_context;
	//
	///**
	// * D3D11 configuration used to create the decoder
	// */
	//D3D11_VIDEO_DECODER_CONFIG *cfg;
	//
	///**
	// * The number of surface in the surface array
	// */
	//unsigned surface_count;
	//
	///**
	// * The array of Direct3D surfaces used to create the decoder
	// */
	//ID3D11VideoDecoderOutputView **surface;
	//
	///**
	// * A bit field configuring the workarounds needed for using the decoder
	// */
	//uint64_t workaround;
	//
	///**
	// * Private to the FFmpeg AVHWAccel implementation
	// */
	//unsigned report_id;
	//
	///**
	// * Mutex to access video_context
	// */
	//HANDLE  context_mutex;
}

/**
 * Allocate an AVD3D11VAContext.
 *
 * @return Newly-allocated AVD3D11VAContext or NULL on failure.
 */
//AVD3D11VAContext *av_d3d11va_alloc_context(void);
//未测试
func AvD3d11vaAllocContext() (res *AVD3D11VAContext, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_d3d11va_alloc_context").Call()
	if err != nil {
		//return
	}
	res = (*AVD3D11VAContext)(unsafe.Pointer(t))
	return
}
