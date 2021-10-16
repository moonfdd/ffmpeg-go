package libavcodec

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/libavutil"
	"unsafe"
)

/**
 * This struct is used for communicating QSV parameters between libavcodec and
 * the caller. It is managed by the caller and must be assigned to
 * AVCodecContext.hwaccel_context.
 * - decoding: hwaccel_context must be set on return from the get_format()
 *             callback
 * - encoding: hwaccel_context must be set before avcodec_open2()
 */
type AVQSVContext struct {
	/**
	 * If non-NULL, the session to use for encoding or decoding.
	 * Otherwise, libavcodec will try to create an internal session.
	 */
	//mfxSession session;

	/**
	 * The IO pattern to use.
	 */
	iopattern ffcommon.FInt

	/**
	 * Extra buffers to pass to encoder or decoder initialization.
	 */
	//mfxExtBuffer **ext_buffers;
	NbExtBuffers ffcommon.FInt

	/**
	 * Encoding only. If this field is set to non-zero by the caller, libavcodec
	 * will create an mfxExtOpaqueSurfaceAlloc extended buffer and pass it to
	 * the encoder initialization. This only makes sense if iopattern is also
	 * set to MFX_IOPATTERN_IN_OPAQUE_MEMORY.
	 *
	 * The number of allocated opaque surfaces will be the sum of the number
	 * required by the encoder and the user-provided value nb_opaque_surfaces.
	 * The array of the opaque surfaces will be exported to the caller through
	 * the opaque_surfaces field.
	 */
	OpaqueAlloc ffcommon.FInt

	/**
	 * Encoding only, and only if opaque_alloc is set to non-zero. Before
	 * calling avcodec_open2(), the caller should set this field to the number
	 * of extra opaque surfaces to allocate beyond what is required by the
	 * encoder.
	 *
	 * On return from avcodec_open2(), this field will be set by libavcodec to
	 * the total number of allocated opaque surfaces.
	 */
	NbOpaqueSurfaces ffcommon.FInt

	/**
	 * Encoding only, and only if opaque_alloc is set to non-zero. On return
	 * from avcodec_open2(), this field will be used by libavcodec to export the
	 * array of the allocated opaque surfaces to the caller, so they can be
	 * passed to other parts of the pipeline.
	 *
	 * The buffer reference exported here is owned and managed by libavcodec,
	 * the callers should make their own reference with av_buffer_ref() and free
	 * it with av_buffer_unref() when it is no longer needed.
	 *
	 * The buffer data is an nb_opaque_surfaces-sized array of mfxFrameSurface1.
	 */
	OpaqueSurfaces *libavutil.AVBufferRef

	/**
	 * Encoding only, and only if opaque_alloc is set to non-zero. On return
	 * from avcodec_open2(), this field will be set to the surface type used in
	 * the opaque allocation request.
	 */
	OpaqueAllocType ffcommon.FInt
}

/**
 * Allocate a new context.
 *
 * It must be freed by the caller with av_free().
 */
//AVQSVContext *av_qsv_alloc_context(void);
//未测试
func AvQsvAllocContext() (res *AVQSVContext, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_qsv_alloc_context").Call()
	if err != nil {
		//return
	}
	//res = &AVAES{}
	res = (*AVQSVContext)(unsafe.Pointer(t))
	//res.instance = t
	//res.ptr = unsafe.Pointer(t)
	return
}
