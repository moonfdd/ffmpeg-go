package ffconstant

/**
 * @defgroup lavu_frame AVFrame
 * @ingroup lavu_data
 *
 * @{
 * AVFrame is an abstraction for reference-counted raw multimedia data.
 */
type AVFrameSideDataType int32

const (
	/**
	 * The data is the AVPanScan struct defined in libavcodec.
	 */
	AV_FRAME_DATA_PANSCAN = 0
	/**
	 * ATSC A53 Part 4 Closed Captions.
	 * A53 CC bitstream is stored as uint8_t in AVFrameSideData.data.
	 * The number of bytes of CC data is AVFrameSideData.size.
	 */
	AV_FRAME_DATA_A53_CC
	/**
	 * Stereoscopic 3d metadata.
	 * The data is the AVStereo3D struct defined in libavutil/stereo3d.h.
	 */
	AV_FRAME_DATA_STEREO3D
	/**
	 * The data is the AVMatrixEncoding enum defined in libavutil/channel_layout.h.
	 */
	AV_FRAME_DATA_MATRIXENCODING
	/**
	 * Metadata relevant to a downmix procedure.
	 * The data is the AVDownmixInfo struct defined in libavutil/downmix_info.h.
	 */
	AV_FRAME_DATA_DOWNMIX_INFO
	/**
	 * ReplayGain information in the form of the AVReplayGain struct.
	 */
	AV_FRAME_DATA_REPLAYGAIN
	/**
	 * This side data contains a 3x3 transformation matrix describing an affine
	 * transformation that needs to be applied to the frame for correct
	 * presentation.
	 *
	 * See libavutil/display.h for a detailed description of the data.
	 */
	AV_FRAME_DATA_DISPLAYMATRIX
	/**
	 * Active Format Description data consisting of a single byte as specified
	 * in ETSI TS 101 154 using AVActiveFormatDescription enum.
	 */
	AV_FRAME_DATA_AFD
	/**
	 * Motion vectors exported by some codecs (on demand through the export_mvs
	 * flag set in the libavcodec AVCodecContext flags2 option).
	 * The data is the AVMotionVector struct defined in
	 * libavutil/motion_vector.h.
	 */
	AV_FRAME_DATA_MOTION_VECTORS
	/**
	 * Recommmends skipping the specified number of samples. This is exported
	 * only if the "skip_manual" AVOption is set in libavcodec.
	 * This has the same format as AV_PKT_DATA_SKIP_SAMPLES.
	 * @code
	 * u32le number of samples to skip from start of this packet
	 * u32le number of samples to skip from end of this packet
	 * u8    reason for start skip
	 * u8    reason for end   skip (0=padding silence, 1=convergence)
	 * @endcode
	 */
	AV_FRAME_DATA_SKIP_SAMPLES
	/**
	 * This side data must be associated with an audio frame and corresponds to
	 * enum AVAudioServiceType defined in avcodec.h.
	 */
	AV_FRAME_DATA_AUDIO_SERVICE_TYPE
	/**
	 * Mastering display metadata associated with a video frame. The payload is
	 * an AVMasteringDisplayMetadata type and contains information about the
	 * mastering display color volume.
	 */
	AV_FRAME_DATA_MASTERING_DISPLAY_METADATA
	/**
	 * The GOP timecode in 25 bit timecode format. Data format is 64-bit integer.
	 * This is set on the first frame of a GOP that has a temporal reference of 0.
	 */
	AV_FRAME_DATA_GOP_TIMECODE

	/**
	 * The data represents the AVSphericalMapping structure defined in
	 * libavutil/spherical.h.
	 */
	AV_FRAME_DATA_SPHERICAL

	/**
	 * Content light level (based on CTA-861.3). This payload contains data in
	 * the form of the AVContentLightMetadata struct.
	 */
	AV_FRAME_DATA_CONTENT_LIGHT_LEVEL

	/**
	 * The data contains an ICC profile as an opaque octet buffer following the
	 * format described by ISO 15076-1 with an optional name defined in the
	 * metadata key entry "name".
	 */
	AV_FRAME_DATA_ICC_PROFILE

	//#if FF_API_FRAME_QP
	///**
	// * Implementation-specific description of the format of AV_FRAME_QP_TABLE_DATA.
	// * The contents of this side data are undocumented and internal; use
	// * av_frame_set_qp_table() and av_frame_get_qp_table() to access this in a
	// * meaningful way instead.
	// */
	//AV_FRAME_DATA_QP_TABLE_PROPERTIES,
	//
	///**
	// * Raw QP table data. Its format is described by
	// * AV_FRAME_DATA_QP_TABLE_PROPERTIES. Use av_frame_set_qp_table() and
	// * av_frame_get_qp_table() to access this instead.
	// */
	//AV_FRAME_DATA_QP_TABLE_DATA,
	//#endif

	/**
	 * Timecode which conforms to SMPTE ST 12-1. The data is an array of 4 uint32_t
	 * where the first uint32_t describes how many (1-3) of the other timecodes are used.
	 * The timecode format is described in the documentation of av_timecode_get_smpte_from_framenum()
	 * function in libavutil/timecode.h.
	 */
	AV_FRAME_DATA_S12M_TIMECODE

	/**
	 * HDR dynamic metadata associated with a video frame. The payload is
	 * an AVDynamicHDRPlus type and contains information for color
	 * volume transform - application 4 of SMPTE 2094-40:2016 standard.
	 */
	AV_FRAME_DATA_DYNAMIC_HDR_PLUS

	/**
	 * Regions Of Interest, the data is an array of AVRegionOfInterest type, the number of
	 * array element is implied by AVFrameSideData.size / AVRegionOfInterest.self_size.
	 */
	AV_FRAME_DATA_REGIONS_OF_INTEREST

	/**
	 * Encoding parameters for a video frame, as described by AVVideoEncParams.
	 */
	AV_FRAME_DATA_VIDEO_ENC_PARAMS

	/**
	 * User data unregistered metadata associated with a video frame.
	 * This is the H.26[45] UDU SEI message, and shouldn't be used for any other purpose
	 * The data is stored as uint8_t in AVFrameSideData.data which is 16 bytes of
	 * uuid_iso_iec_11578 followed by AVFrameSideData.size - 16 bytes of user_data_payload_byte.
	 */
	AV_FRAME_DATA_SEI_UNREGISTERED

	/**
	 * Film grain parameters for a frame, described by AVFilmGrainParams.
	 * Must be present for every frame which should have film grain applied.
	 */
	AV_FRAME_DATA_FILM_GRAIN_PARAMS
)

type AVActiveFormatDescription int32

const (
	AV_AFD_SAME         = 8
	AV_AFD_4_3          = 9
	AV_AFD_16_9         = 10
	AV_AFD_14_9         = 11
	AV_AFD_4_3_SP_14_9  = 13
	AV_AFD_16_9_SP_14_9 = 14
	AV_AFD_SP_4_3       = 15
)

const AV_NUM_DATA_POINTERS = 8

/**
 * The frame data may be corrupted, e.g. due to decoding errors.
 */
const AV_FRAME_FLAG_CORRUPT = (1 << 0)

/**
 * A flag to mark the frames which need to be decoded, but shouldn't be output.
 */
const AV_FRAME_FLAG_DISCARD = (1 << 2)

const FF_DECODE_ERROR_INVALID_BITSTREAM = 1
const FF_DECODE_ERROR_MISSING_REFERENCE = 2
const FF_DECODE_ERROR_CONCEALMENT_ACTIVE = 4
const FF_DECODE_ERROR_DECODE_SLICES = 8

/**
 * Flags for frame cropping.
 */
/**
 * Apply the maximum possible cropping, even if it requires setting the
 * AVFrame.data[] entries to unaligned pointers. Passing unaligned data
 * to FFmpeg API is generally not allowed, and causes undefined behavior
 * (such as crashes). You can pass unaligned data only to FFmpeg APIs that
 * are explicitly documented to accept it. Use this flag only if you
 * absolutely know what you are doing.
 */
const AV_FRAME_CROP_UNALIGNED = 1 << 0
