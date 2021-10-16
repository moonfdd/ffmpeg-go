package libavcodec

/**
 * This structure is used to provides the necessary configurations and data
 * to the DXVA2 FFmpeg HWAccel implementation.
 *
 * The application must make it available as AVCodecContext.hwaccel_context.
 */
type DxvaContext struct {

	///**
	// * DXVA2 decoder object
	// */
	//IDirectXVideoDecoder *decoder;
	//
	///**
	// * DXVA2 configuration used to create the decoder
	// */
	//const DXVA2_ConfigPictureDecode *cfg;
	//
	///**
	// * The number of surface in the surface array
	// */
	//unsigned surface_count;
	//
	///**
	// * The array of Direct3D surfaces used to create the decoder
	// */
	//LPDIRECT3DSURFACE9 *surface;
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
}
