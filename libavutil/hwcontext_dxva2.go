package libavutil

/**
 * This struct is allocated as AVHWDeviceContext.hwctx
 */
type AVDXVA2DeviceContext struct {
	//IDirect3DDeviceManager9 *devmgr;
}

/**
 * This struct is allocated as AVHWFramesContext.hwctx
 */
type AVDXVA2FramesContext struct {
	///**
	// * The surface type (e.g. DXVA2_VideoProcessorRenderTarget or
	// * DXVA2_VideoDecoderRenderTarget). Must be set by the caller.
	// */
	//DWORD               surface_type;
	//
	///**
	// * The surface pool. When an external pool is not provided by the caller,
	// * this will be managed (allocated and filled on init, freed on uninit) by
	// * libavutil.
	// */
	//IDirect3DSurface9 **surfaces;
	//int              nb_surfaces;
	//
	///**
	// * Certain drivers require the decoder to be destroyed before the surfaces.
	// * To allow internally managed pools to work properly in such cases, this
	// * field is provided.
	// *
	// * If it is non-NULL, libavutil will call IDirectXVideoDecoder_Release() on
	// * it just before the internal surface pool is freed.
	// *
	// * This is for convenience only. Some code uses other methods to manage the
	// * decoder reference.
	// */
	//IDirectXVideoDecoder *decoder_to_release;
}
