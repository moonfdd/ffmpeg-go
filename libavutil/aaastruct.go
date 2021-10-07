package libavutil

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

type AVRational struct {
	Num ffcommon.FInt ///< Numerator
	Den ffcommon.FInt ///< Denominator
}

type AVAESCTR struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

type AVAudioFifo struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

type AVBlowfish struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

type AVBPrint struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * @defgroup lavu_buffer AVBuffer
 * @ingroup lavu_data
 *
 * @{
 * AVBuffer is an API for reference-counted data buffers.
 *
 * There are two core objects in this API -- AVBuffer and AVBufferRef. AVBuffer
 * represents the data buffer itself; it is opaque and not meant to be accessed
 * by the caller directly, but only through AVBufferRef. However, the caller may
 * e.g. compare two AVBuffer pointers to check whether two different references
 * are describing the same data buffer. AVBufferRef represents a single
 * reference to an AVBuffer and it is the object that may be manipulated by the
 * caller directly.
 *
 * There are two functions provided for creating a new AVBuffer with a single
 * reference -- av_buffer_alloc() to just allocate a new buffer, and
 * av_buffer_create() to wrap an existing array in an AVBuffer. From an existing
 * reference, additional references may be created with av_buffer_ref().
 * Use av_buffer_unref() to free a reference (this will automatically free the
 * data once all the references are freed).
 *
 * The convention throughout this API and the rest of FFmpeg is such that the
 * buffer is considered writable if there exists only one reference to it (and
 * it has not been marked as read-only). The av_buffer_is_writable() function is
 * provided to check whether this is true and av_buffer_make_writable() will
 * automatically create a new writable buffer when necessary.
 * Of course nothing prevents the calling code from violating this convention,
 * however that is safe only when all the existing references are under its
 * control.
 *
 * @note Referencing and unreferencing the buffers is thread-safe and thus
 * may be done from multiple threads simultaneously without any need for
 * additional locking.
 *
 * @note Two different references to the same buffer can point to different
 * parts of the buffer (i.e. their AVBufferRef.data will not be equal).
 */

/**
 * A reference counted buffer type. It is opaque and is meant to be used through
 * references (AVBufferRef).
 */
type AVBuffer struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * A reference to a data buffer.
 *
 * The size of this struct is not a part of the public ABI and it is not meant
 * to be allocated directly.
 */
type AVBufferRef struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}
type AVCAMELLIA struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}
type AVCAST5 struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * @defgroup lavu_crc32 CRC
 * @ingroup lavu_hash
 * CRC (Cyclic Redundancy Check) hash function implementation.
 *
 * This module supports numerous CRC polynomials, in addition to the most
 * widely used CRC-32-IEEE. See @ref AVCRCId for a list of available
 * polynomials.
 *
 * @{
 */
type AVCRC ffcommon.FUint32T

/**
 * @defgroup lavu_des DES
 * @ingroup lavu_crypto
 * @{
 */
type AVDES struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}
type AVDictionaryEntry struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/*
 * DOVI configuration
 * ref: dolby-vision-bitstreams-within-the-iso-base-media-file-format-v2.1.2
        dolby-vision-bitstreams-in-mpeg-2-transport-stream-multiplex-v1.2
 * @code
 * uint8_t  dv_version_major, the major version number that the stream complies with
 * uint8_t  dv_version_minor, the minor version number that the stream complies with
 * uint8_t  dv_profile, the Dolby Vision profile
 * uint8_t  dv_level, the Dolby Vision level
 * uint8_t  rpu_present_flag
 * uint8_t  el_present_flag
 * uint8_t  bl_present_flag
 * uint8_t  dv_bl_signal_compatibility_id
 * @endcode
 *
 * @note The struct must be allocated with av_dovi_alloc() and
 *       its size is not a part of the public ABI.
*/
type AVDOVIDecoderConfigurationRecord struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * This structure describes optional metadata relevant to a downmix procedure.
 *
 * All fields are set by the decoder to the value indicated in the audio
 * bitstream (if present), or to a "sane" default otherwise.
 */
type AVDownmixInfo struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

type AVSubsampleEncryptionInfo struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * This describes encryption info for a packet.  This contains frame-specific
 * info for how to decrypt the packet before passing it to the decoder.
 *
 * The size of this struct is not part of the public ABI.
 */
type AVEncryptionInfo struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * This describes info used to initialize an encryption key system.
 *
 * The size of this struct is not part of the public ABI.
 */
type AVEncryptionInitInfo struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

type AVExpr struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

type AVFifoBuffer struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * This structure describes how to handle film grain synthesis for AOM codecs.
 *
 * @note The struct must be allocated as part of AVFilmGrainParams using
 *       av_film_grain_params_alloc(). Its size is not a part of the public ABI.
 */
type AVFilmGrainAOMParams struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * This structure describes how to handle film grain synthesis in video
 * for specific codecs. Must be present on every frame where film grain is
 * meant to be synthesised for correct presentation.
 *
 * @note The struct must be allocated with av_film_grain_params_alloc() and
 *       its size is not a part of the public ABI.
 */
type AVFilmGrainParams struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * Structure to hold side data for an AVFrame.
 *
 * sizeof(AVFrameSideData) is not a part of the public ABI, so new fields may be added
 * to the end with a minor bump.
 */
type AVFrameSideData struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * Structure describing a single Region Of Interest.
 *
 * When multiple regions are defined in a single side-data block, they
 * should be ordered from most to least important - some encoders are only
 * capable of supporting a limited number of distinct regions, so will have
 * to truncate the list.
 *
 * When overlapping regions are defined, the first region containing a given
 * area of the frame applies.
 */
type AVRegionOfInterest struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * @example ffhash.c
 * This example is a simple command line application that takes one or more
 * arguments. It demonstrates a typical use of the hashing API with allocation,
 * initialization, updating, and finalizing.
 */
type AVHashContext struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * Represents the percentile at a specific percentage in
 * a distribution.
 */
type AVHDRPlusPercentile struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * Color transform parameters at a processing window in a dynamic metadata for
 * SMPTE 2094-40.
 */
type AVHDRPlusColorTransformParams struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * This struct represents dynamic metadata for color volume transform -
 * application 4 of SMPTE 2094-40:2016 standard.
 *
 * To be used as payload of a AVFrameSideData or AVPacketSideData with the
 * appropriate type.
 *
 * @note The struct should be allocated with
 * av_dynamic_hdr_plus_alloc() and its size is not a part of
 * the public ABI.
 */
type AVDynamicHDRPlus struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

type AVHMAC struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

type AVHWDeviceInternal struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * This struct aggregates all the (hardware/vendor-specific) "high-level" state,
 * i.e. state that is not tied to a concrete processing configuration.
 * E.g., in an API that supports hardware-accelerated encoding and decoding,
 * this struct will (if possible) wrap the state that is common to both encoding
 * and decoding and from which specific instances of encoders or decoders can be
 * derived.
 *
 * This struct is reference-counted with the AVBuffer mechanism. The
 * av_hwdevice_ctx_alloc() constructor yields a reference, whose data field
 * points to the actual AVHWDeviceContext. Further objects derived from
 * AVHWDeviceContext (such as AVHWFramesContext, describing a frame pool with
 * specific properties) will hold an internal reference to it. After all the
 * references are released, the AVHWDeviceContext itself will be freed,
 * optionally invoking a user-specified callback for uninitializing the hardware
 * state.
 */
type AVHWDeviceContext struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

type AVHWFramesInternal struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * This struct describes a set or pool of "hardware" frames (i.e. those with
 * data not located in normal system memory). All the frames in the pool are
 * assumed to be allocated in the same way and interchangeable.
 *
 * This struct is reference-counted with the AVBuffer mechanism and tied to a
 * given AVHWDeviceContext instance. The av_hwframe_ctx_alloc() constructor
 * yields a reference, whose data field points to the actual AVHWFramesContext
 * struct.
 */
type AVHWFramesContext struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

type AVCUDADeviceContextInternal struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

type AVCUDADeviceContext struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * This struct is allocated as AVHWDeviceContext.hwctx
 */
type AVD3D11VADeviceContext struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * D3D11 frame descriptor for pool allocation.
 *
 * In user-allocated pools, AVHWFramesContext.pool must return AVBufferRefs
 * with the data pointer pointing at an object of this type describing the
 * planes of the frame.
 *
 * This has no use outside of custom allocation, and AVFrame AVBufferRef do not
 * necessarily point to an instance of this struct.
 */
type AVD3D11FrameDescriptor struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * This struct is allocated as AVHWFramesContext.hwctx
 */
type AVD3D11VAFramesContext struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * DRM object descriptor.
 *
 * Describes a single DRM object, addressing it as a PRIME file
 * descriptor.
 */
type AVDRMObjectDescriptor struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * DRM plane descriptor.
 *
 * Describes a single plane of a layer, which is contained within
 * a single object.
 */
type AVDRMPlaneDescriptor struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * DRM layer descriptor.
 *
 * Describes a single layer within a frame.  This has the structure
 * defined by its format, and will contain one or more planes.
 */
type AVDRMLayerDescriptor struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
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
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * DRM device.
 *
 * Allocated as AVHWDeviceContext.hwctx.
 */
type AVDRMDeviceContext struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * This struct is allocated as AVHWDeviceContext.hwctx
 */
type AVDXVA2DeviceContext struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * This struct is allocated as AVHWFramesContext.hwctx
 */
type AVDXVA2FramesContext struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * MediaCodec details.
 *
 * Allocated as AVHWDeviceContext.hwctx
 */
type AVMediaCodecDeviceContext struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

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
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * OpenCL device details.
 *
 * Allocated as AVHWDeviceContext.hwctx
 */
type AVOpenCLDeviceContext struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * OpenCL-specific data associated with a frame pool.
 *
 * Allocated as AVHWFramesContext.hwctx.
 */
type AVOpenCLFramesContext struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * This struct is allocated as AVHWDeviceContext.hwctx
 */
type AVQSVDeviceContext struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * This struct is allocated as AVHWFramesContext.hwctx
 */
type AVQSVFramesContext struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * VAAPI connection details.
 *
 * Allocated as AVHWDeviceContext.hwctx
 */
type AVVAAPIDeviceContext struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * VAAPI-specific data associated with a frame pool.
 *
 * Allocated as AVHWFramesContext.hwctx.
 */
type AVVAAPIFramesContext struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * VAAPI hardware pipeline configuration details.
 *
 * Allocated with av_hwdevice_hwconfig_alloc().
 */
type AVVAAPIHWConfig struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * This struct is allocated as AVHWDeviceContext.hwctx
 */
type AVVDPAUDeviceContext struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * Main Vulkan context, allocated as AVHWDeviceContext.hwctx.
 * All of these can be set before init to change what the context uses
 */
type AVVulkanDeviceContext struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * Allocated as AVHWFramesContext.hwctx, used to set pool-specific options
 */
type AVVulkanFramesContext struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/*
 * Frame structure, the VkFormat of the image will always match
 * the pool's sw_format.
 * All frames, imported or allocated, will be created with the
 * VK_IMAGE_CREATE_ALIAS_BIT flag set, so the memory may be aliased if needed.
 *
 * If all three queue family indices in the device context are the same,
 * images will be created with the EXCLUSIVE sharing mode. Otherwise, all images
 * will be created using the CONCURRENT sharing mode.
 *
 * @note the size of this structure is not part of the ABI, to allocate
 * you must use @av_vk_frame_alloc().
 */
type AVVkFrame struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * Context structure for the Lagged Fibonacci PRNG.
 * The exact layout, types and content of this struct may change and should
 * not be accessed directly. Only its sizeof() is guranteed to stay the same
 * to allow easy instanciation.
 */
type AVLFG struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * Describe the class of an AVClass context structure. That is an
 * arbitrary struct of which the first field is a pointer to an
 * AVClass struct (e.g. AVCodecContext, AVFormatContext etc.).
 */
type AVClass struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * Mastering display metadata capable of representing the color volume of
 * the display used to master the content (SMPTE 2086:2014).
 *
 * To be used as payload of a AVFrameSideData or AVPacketSideData with the
 * appropriate type.
 *
 * @note The struct should be allocated with av_mastering_display_metadata_alloc()
 *       and its size is not a part of the public ABI.
 */
type AVMasteringDisplayMetadata struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * Content light level needed by to transmit HDR over HDMI (CTA-861.3).
 *
 * To be used as payload of a AVFrameSideData or AVPacketSideData with the
 * appropriate type.
 *
 * @note The struct should be allocated with av_content_light_metadata_alloc()
 *       and its size is not a part of the public ABI.
 */
type AVContentLightMetadata struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

type AVMD5 struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

type AVMotionVector struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

type AVMurMur3 struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * AVOption
 */
type AVOption struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * A single allowed range of values, or a single allowed value.
 */
type AVOptionRange struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * List of AVOptionRange structs.
 */
type AVOptionRanges struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

type AVComponentDescriptor struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * Descriptor that unambiguously describes how the bits of a pixel are
 * stored in the up to 4 data planes of an image. It also stores the
 * subsampling factors and number of components.
 *
 * @note This is separate of the colorspace (RGB, YCbCr, YPbPr, JPEG-style YUV
 *       and all the YUV variants) AVPixFmtDescriptor just stores how values
 *       are stored not what these values represent.
 */
type AVPixFmtDescriptor struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * @defgroup lavu_rc4 RC4
 * @ingroup lavu_crypto
 * @{
 */
type AVRC4 struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * ReplayGain information (see
 * http://wiki.hydrogenaudio.org/index.php?title=ReplayGain_1.0_specification).
 * The size of this struct is a part of the public ABI.
 */
type AVReplayGain struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

type AVRIPEMD struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

type AVSHA struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

type AVSHA512 struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * This structure describes how to handle spherical videos, outlining
 * information about projection, initial layout, and any other view modifier.
 *
 * @note The struct must be allocated with av_spherical_alloc() and
 *       its size is not a part of the public ABI.
 */
type AVSphericalMapping struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}

/**
 * List of possible 3D Types
 */
type AVStereo3DType struct {
	instance uintptr
	// 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
	ptr unsafe.Pointer
}
