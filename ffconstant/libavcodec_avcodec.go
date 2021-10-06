package ffconstant

/**
 * @defgroup libavc libavcodec
 * Encoding/Decoding Library
 *
 * @{
 *
 * @defgroup lavc_decoding Decoding
 * @{
 * @}
 *
 * @defgroup lavc_encoding Encoding
 * @{
 * @}
 *
 * @defgroup lavc_codec Codecs
 * @{
 * @defgroup lavc_codec_native Native Codecs
 * @{
 * @}
 * @defgroup lavc_codec_wrappers External library wrappers
 * @{
 * @}
 * @defgroup lavc_codec_hwaccel Hardware Accelerators bridge
 * @{
 * @}
 * @}
 * @defgroup lavc_internal Internal
 * @{
 * @}
 * @}
 */

/**
 * @ingroup libavc
 * @defgroup lavc_encdec send/receive encoding and decoding API overview
 * @{
 *
 * The avcodec_send_packet()/avcodec_receive_frame()/avcodec_send_frame()/
 * avcodec_receive_packet() functions provide an encode/decode API, which
 * decouples input and output.
 *
 * The API is very similar for encoding/decoding and audio/video, and works as
 * follows:
 * - Set up and open the AVCodecContext as usual.
 * - Send valid input:
 *   - For decoding, call avcodec_send_packet() to give the decoder raw
 *     compressed data in an AVPacket.
 *   - For encoding, call avcodec_send_frame() to give the encoder an AVFrame
 *     containing uncompressed audio or video.
 *
 *   In both cases, it is recommended that AVPackets and AVFrames are
 *   refcounted, or libavcodec might have to copy the input data. (libavformat
 *   always returns refcounted AVPackets, and av_frame_get_buffer() allocates
 *   refcounted AVFrames.)
 * - Receive output in a loop. Periodically call one of the avcodec_receive_*()
 *   functions and process their output:
 *   - For decoding, call avcodec_receive_frame(). On success, it will return
 *     an AVFrame containing uncompressed audio or video data.
 *   - For encoding, call avcodec_receive_packet(). On success, it will return
 *     an AVPacket with a compressed frame.
 *
 *   Repeat this call until it returns AVERROR(EAGAIN) or an error. The
 *   AVERROR(EAGAIN) return value means that new input data is required to
 *   return new output. In this case, continue with sending input. For each
 *   input frame/packet, the codec will typically return 1 output frame/packet,
 *   but it can also be 0 or more than 1.
 *
 * At the beginning of decoding or encoding, the codec might accept multiple
 * input frames/packets without returning a frame, until its internal buffers
 * are filled. This situation is handled transparently if you follow the steps
 * outlined above.
 *
 * In theory, sending input can result in EAGAIN - this should happen only if
 * not all output was received. You can use this to structure alternative decode
 * or encode loops other than the one suggested above. For example, you could
 * try sending new input on each iteration, and try to receive output if that
 * returns EAGAIN.
 *
 * End of stream situations. These require "flushing" (aka draining) the codec,
 * as the codec might buffer multiple frames or packets internally for
 * performance or out of necessity (consider B-frames).
 * This is handled as follows:
 * - Instead of valid input, send NULL to the avcodec_send_packet() (decoding)
 *   or avcodec_send_frame() (encoding) functions. This will enter draining
 *   mode.
 * - Call avcodec_receive_frame() (decoding) or avcodec_receive_packet()
 *   (encoding) in a loop until AVERROR_EOF is returned. The functions will
 *   not return AVERROR(EAGAIN), unless you forgot to enter draining mode.
 * - Before decoding can be resumed again, the codec has to be reset with
 *   avcodec_flush_buffers().
 *
 * Using the API as outlined above is highly recommended. But it is also
 * possible to call functions outside of this rigid schema. For example, you can
 * call avcodec_send_packet() repeatedly without calling
 * avcodec_receive_frame(). In this case, avcodec_send_packet() will succeed
 * until the codec's internal buffer has been filled up (which is typically of
 * size 1 per output frame, after initial input), and then reject input with
 * AVERROR(EAGAIN). Once it starts rejecting input, you have no choice but to
 * read at least some output.
 *
 * Not all codecs will follow a rigid and predictable dataflow; the only
 * guarantee is that an AVERROR(EAGAIN) return value on a send/receive call on
 * one end implies that a receive/send call on the other end will succeed, or
 * at least will not fail with AVERROR(EAGAIN). In general, no codec will
 * permit unlimited buffering of input or output.
 *
 * This API replaces the following legacy functions:
 * - avcodec_decode_video2() and avcodec_decode_audio4():
 *   Use avcodec_send_packet() to feed input to the decoder, then use
 *   avcodec_receive_frame() to receive decoded frames after each packet.
 *   Unlike with the old video decoding API, multiple frames might result from
 *   a packet. For audio, splitting the input packet into frames by partially
 *   decoding packets becomes transparent to the API user. You never need to
 *   feed an AVPacket to the API twice (unless it is rejected with AVERROR(EAGAIN) - then
 *   no data was read from the packet).
 *   Additionally, sending a flush/draining packet is required only once.
 * - avcodec_encode_video2()/avcodec_encode_audio2():
 *   Use avcodec_send_frame() to feed input to the encoder, then use
 *   avcodec_receive_packet() to receive encoded packets.
 *   Providing user-allocated buffers for avcodec_receive_packet() is not
 *   possible.
 * - The new API does not handle subtitles yet.
 *
 * Mixing new and old function calls on the same AVCodecContext is not allowed,
 * and will result in undefined behavior.
 *
 * Some codecs might require using the new API; using the old API will return
 * an error when calling it. All codecs support the new API.
 *
 * A codec is not allowed to return AVERROR(EAGAIN) for both sending and receiving. This
 * would be an invalid state, which could put the codec user into an endless
 * loop. The API has no concept of time either: it cannot happen that trying to
 * do avcodec_send_packet() results in AVERROR(EAGAIN), but a repeated call 1 second
 * later accepts the packet (with no other receive/flush API calls involved).
 * The API is a strict state machine, and the passage of time is not supposed
 * to influence it. Some timing-dependent behavior might still be deemed
 * acceptable in certain cases. But it must never result in both send/receive
 * returning EAGAIN at the same time at any point. It must also absolutely be
 * avoided that the current state is "unstable" and can "flip-flop" between
 * the send/receive APIs allowing progress. For example, it's not allowed that
 * the codec randomly decides that it actually wants to consume a packet now
 * instead of returning a frame, after it just returned AVERROR(EAGAIN) on an
 * avcodec_send_packet() call.
 * @}
 */

/**
 * @defgroup lavc_core Core functions/structures.
 * @ingroup libavc
 *
 * Basic definitions, functions for querying libavcodec capabilities,
 * allocating core structures, etc.
 * @{
 */

/**
 * @ingroup lavc_decoding
 * Required number of additionally allocated bytes at the end of the input bitstream for decoding.
 * This is mainly needed because some optimized bitstream readers read
 * 32 or 64 bit at once and could read over the end.<br>
 * Note: If the first 23 bits of the additional bytes are not 0, then damaged
 * MPEG bitstreams could cause overread and segfault.
 */
const AV_INPUT_BUFFER_PADDING_SIZE = 64

/**
 * @ingroup lavc_encoding
 * minimum encoding buffer size
 * Used to avoid some checks during header writing.
 */
const AV_INPUT_BUFFER_MIN_SIZE = 16384

/**
 * @ingroup lavc_decoding
 */
type AVDiscard int32

const (
	/* We leave some space between them for extensions (drop some
	 * keyframes for intra-only or drop just some bidir frames). */
	AVDISCARD_NONE     = -16 ///< discard nothing
	AVDISCARD_DEFAULT  = 0   ///< discard useless packets like 0 size packets in avi
	AVDISCARD_NONREF   = 8   ///< discard all non reference
	AVDISCARD_BIDIR    = 16  ///< discard all bidirectional frames
	AVDISCARD_NONINTRA = 24  ///< discard all non intra frames
	AVDISCARD_NONKEY   = 32  ///< discard all frames except keyframes
	AVDISCARD_ALL      = 48  ///< discard all
)

type AVAudioServiceType int32

const (
	AV_AUDIO_SERVICE_TYPE_MAIN              = 0
	AV_AUDIO_SERVICE_TYPE_EFFECTS           = 1
	AV_AUDIO_SERVICE_TYPE_VISUALLY_IMPAIRED = 2
	AV_AUDIO_SERVICE_TYPE_HEARING_IMPAIRED  = 3
	AV_AUDIO_SERVICE_TYPE_DIALOGUE          = 4
	AV_AUDIO_SERVICE_TYPE_COMMENTARY        = 5
	AV_AUDIO_SERVICE_TYPE_EMERGENCY         = 6
	AV_AUDIO_SERVICE_TYPE_VOICE_OVER        = 7
	AV_AUDIO_SERVICE_TYPE_KARAOKE           = 8
	AV_AUDIO_SERVICE_TYPE_NB                ///< Not part of ABI
)

/* encoding support
   These flags can be passed in AVCodecContext.flags before initialization.
   Note: Not everything is supported yet.
*/

/**
 * Allow decoders to produce frames with data planes that are not aligned
 * to CPU requirements (e.g. due to cropping).
 */
const AV_CODEC_FLAG_UNALIGNED = (1 << 0)

/**
 * Use fixed qscale.
 */
const AV_CODEC_FLAG_QSCALE = (1 << 1)

/**
 * 4 MV per MB allowed / advanced prediction for H.263.
 */
const AV_CODEC_FLAG_4MV = (1 << 2)

/**
 * Output even those frames that might be corrupted.
 */
const AV_CODEC_FLAG_OUTPUT_CORRUPT = (1 << 3)

/**
 * Use qpel MC.
 */
const AV_CODEC_FLAG_QPEL = (1 << 4)

/**
 * Don't output frames whose parameters differ from first
 * decoded frame in stream.
 */
const AV_CODEC_FLAG_DROPCHANGED = (1 << 5)

/**
 * Use internal 2pass ratecontrol in first pass mode.
 */
const AV_CODEC_FLAG_PASS1 = (1 << 9)

/**
 * Use internal 2pass ratecontrol in second pass mode.
 */
const AV_CODEC_FLAG_PASS2 = (1 << 10)

/**
 * loop filter.
 */
const AV_CODEC_FLAG_LOOP_FILTER = (1 << 11)

/**
 * Only decode/encode grayscale.
 */
const AV_CODEC_FLAG_GRAY = (1 << 13)

/**
 * error[?] variables will be set during encoding.
 */
const AV_CODEC_FLAG_PSNR = (1 << 15)

/**
 * Input bitstream might be truncated at a random location
 * instead of only at frame boundaries.
 */
const AV_CODEC_FLAG_TRUNCATED = (1 << 16)

/**
 * Use interlaced DCT.
 */
const AV_CODEC_FLAG_INTERLACED_DCT = (1 << 18)

/**
 * Force low delay.
 */
const AV_CODEC_FLAG_LOW_DELAY = (1 << 19)

/**
 * Place global headers in extradata instead of every keyframe.
 */
const AV_CODEC_FLAG_GLOBAL_HEADER = (1 << 22)

/**
 * Use only bitexact stuff (except (I)DCT).
 */
const AV_CODEC_FLAG_BITEXACT = (1 << 23)

/* Fx : Flag for H.263+ extra options */
/**
 * H.263 advanced intra coding / MPEG-4 AC prediction
 */
const AV_CODEC_FLAG_AC_PRED = (1 << 24)

/**
 * interlaced motion estimation
 */
const AV_CODEC_FLAG_INTERLACED_ME = (1 << 29)
const AV_CODEC_FLAG_CLOSED_GOP = (1 << 31)

/**
 * Allow non spec compliant speedup tricks.
 */
const AV_CODEC_FLAG2_FAST = (1 << 0)

/**
 * Skip bitstream encoding.
 */
const AV_CODEC_FLAG2_NO_OUTPUT = (1 << 2)

/**
 * Place global headers at every keyframe instead of in extradata.
 */
const AV_CODEC_FLAG2_LOCAL_HEADER = (1 << 3)

/**
 * timecode is in drop frame format. DEPRECATED!!!!
 */
const AV_CODEC_FLAG2_DROP_FRAME_TIMECODE = (1 << 13)

/**
 * Input bitstream might be truncated at a packet boundaries
 * instead of only at frame boundaries.
 */
const AV_CODEC_FLAG2_CHUNKS = (1 << 15)

/**
 * Discard cropping information from SPS.
 */
const AV_CODEC_FLAG2_IGNORE_CROP = (1 << 16)

/**
 * Show all frames before the first keyframe
 */
const AV_CODEC_FLAG2_SHOW_ALL = (1 << 22)

/**
 * Export motion vectors through frame side data
 */
const AV_CODEC_FLAG2_EXPORT_MVS = (1 << 28)

/**
 * Do not skip samples and export skip information as frame side data
 */
const AV_CODEC_FLAG2_SKIP_MANUAL = (1 << 29)

/**
 * Do not reset ASS ReadOrder field on flush (subtitles decoding)
 */
const AV_CODEC_FLAG2_RO_FLUSH_NOOP = (1 << 30)

/* Unsupported options :
 *              Syntax Arithmetic coding (SAC)
 *              Reference Picture Selection
 *              Independent Segment Decoding */
/* /Fx */
/* codec capabilities */

/* Exported side data.
   These flags can be passed in AVCodecContext.export_side_data before initialization.
*/
/**
 * Export motion vectors through frame side data
 */
const AV_CODEC_EXPORT_DATA_MVS = (1 << 0)

/**
 * Export encoder Producer Reference Time through packet side data
 */
const AV_CODEC_EXPORT_DATA_PRFT = (1 << 1)

/**
 * Decoding only.
 * Export the AVVideoEncParams structure through frame side data.
 */
const AV_CODEC_EXPORT_DATA_VIDEO_ENC_PARAMS = (1 << 2)

/**
 * Decoding only.
 * Do not apply film grain, export it instead.
 */
const AV_CODEC_EXPORT_DATA_FILM_GRAIN = (1 << 3)

/**
 * The decoder will keep a reference to the frame and may reuse it later.
 */
const AV_GET_BUFFER_FLAG_REF = (1 << 0)

/**
 * The encoder will keep a reference to the packet and may reuse it later.
 */
const AV_GET_ENCODE_BUFFER_FLAG_REF = (1 << 0)

const FF_COMPRESSION_DEFAULT = -1

const FF_PRED_LEFT = 0
const FF_PRED_PLANE = 1
const FF_PRED_MEDIAN = 2

const FF_CMP_SAD = 0
const FF_CMP_SSE = 1
const FF_CMP_SATD = 2
const FF_CMP_DCT = 3
const FF_CMP_PSNR = 4
const FF_CMP_BIT = 5
const FF_CMP_RD = 6
const FF_CMP_ZERO = 7
const FF_CMP_VSAD = 8
const FF_CMP_VSSE = 9
const FF_CMP_NSSE = 10
const FF_CMP_W53 = 11
const FF_CMP_W97 = 12
const FF_CMP_DCTMAX = 13
const FF_CMP_DCT264 = 14
const FF_CMP_MEDIAN_SAD = 15
const FF_CMP_CHROMA = 256

const SLICE_FLAG_CODED_ORDER = 0x0001 ///< draw_horiz_band() is called in coded order instead of display
const SLICE_FLAG_ALLOW_FIELD = 0x0002 ///< allow draw_horiz_band() with field slices (MPEG-2 field pics)
const SLICE_FLAG_ALLOW_PLANE = 0x0004 ///< allow draw_horiz_band() with 1 component at a time (SVQ1)

const FF_MB_DECISION_SIMPLE = 0 ///< uses mb_cmp
const FF_MB_DECISION_BITS = 1   ///< chooses the one which needs the fewest bits
const FF_MB_DECISION_RD = 2     ///< rate distortion

//#if FF_API_CODER_TYPE
const FF_CODER_TYPE_VLC = 0
const FF_CODER_TYPE_AC = 1
const FF_CODER_TYPE_RAW = 2
const FF_CODER_TYPE_RLE = 3

//#endif /* FF_API_CODER_TYPE */

const FF_BUG_AUTODETECT = 1 ///< autodetection
const FF_BUG_XVID_ILACE = 4
const FF_BUG_UMP4 = 8
const FF_BUG_NO_PADDING = 16
const FF_BUG_AMV = 32
const FF_BUG_QPEL_CHROMA = 64
const FF_BUG_STD_QPEL = 128
const FF_BUG_QPEL_CHROMA2 = 256
const FF_BUG_DIRECT_BLOCKSIZE = 512
const FF_BUG_EDGE = 1024
const FF_BUG_HPEL_CHROMA = 2048
const FF_BUG_DC_CLIP = 4096
const FF_BUG_MS = 8192 ///< Work around various bugs in Microsoft's broken decoders.
const FF_BUG_TRUNCATED = 16384
const FF_BUG_IEDGE = 32768

const FF_COMPLIANCE_VERY_STRICT = 2 ///< Strictly conform to an older more strict version of the spec or reference software.
const FF_COMPLIANCE_STRICT = 1      ///< Strictly conform to all the things in the spec no matter what consequences.
const FF_COMPLIANCE_NORMAL = 0
const FF_COMPLIANCE_UNOFFICIAL = -1   ///< Allow unofficial extensions
const FF_COMPLIANCE_EXPERIMENTAL = -2 ///< Allow nonstandardized experimental things.

const FF_EC_GUESS_MVS = 1
const FF_EC_DEBLOCK = 2
const FF_EC_FAVOR_INTER = 256

const FF_DEBUG_PICT_INFO = 1
const FF_DEBUG_RC = 2
const FF_DEBUG_BITSTREAM = 4
const FF_DEBUG_MB_TYPE = 8
const FF_DEBUG_QP = 16
const FF_DEBUG_DCT_COEFF = 0x00000040
const FF_DEBUG_SKIP = 0x00000080
const FF_DEBUG_STARTCODE = 0x00000100
const FF_DEBUG_ER = 0x00000400
const FF_DEBUG_MMCO = 0x00000800
const FF_DEBUG_BUGS = 0x00001000
const FF_DEBUG_BUFFERS = 0x00008000
const FF_DEBUG_THREADS = 0x00010000
const FF_DEBUG_GREEN_MD = 0x00800000
const FF_DEBUG_NOMC = 0x01000000

/**
 * Verify checksums embedded in the bitstream (could be of either encoded or
 * decoded data, depending on the codec) and print an error message on mismatch.
 * If AV_EF_EXPLODE is also set, a mismatching checksum will result in the
 * decoder returning an error.
 */
const AV_EF_CRCCHECK = (1 << 0)
const AV_EF_BITSTREAM = (1 << 1) ///< detect bitstream specification deviations
const AV_EF_BUFFER = (1 << 2)    ///< detect improper bitstream length
const AV_EF_EXPLODE = (1 << 3)   ///< abort decoding on minor error detection

const AV_EF_IGNORE_ERR = (1 << 15) ///< ignore errors and continue
const AV_EF_CAREFUL = (1 << 16)    ///< consider things that violate the spec, are fast to calculate and have not been seen in the wild as errors
const AV_EF_COMPLIANT = (1 << 17)  ///< consider all spec non compliances as errors
const AV_EF_AGGRESSIVE = (1 << 18) ///< consider things that a sane encoder should not do as an error

const FF_DCT_AUTO = 0
const FF_DCT_FASTINT = 1
const FF_DCT_INT = 2
const FF_DCT_MMX = 3
const FF_DCT_ALTIVEC = 5
const FF_DCT_FAAN = 6

const FF_IDCT_AUTO = 0
const FF_IDCT_INT = 1
const FF_IDCT_SIMPLE = 2
const FF_IDCT_SIMPLEMMX = 3
const FF_IDCT_ARM = 7
const FF_IDCT_ALTIVEC = 8
const FF_IDCT_SIMPLEARM = 10
const FF_IDCT_XVID = 14
const FF_IDCT_SIMPLEARMV5TE = 16
const FF_IDCT_SIMPLEARMV6 = 17
const FF_IDCT_FAAN = 20
const FF_IDCT_SIMPLENEON = 22
const FF_IDCT_NONE = 24 /* Used by XvMC to extract IDCT coefficients with FF_IDCT_PERM_NONE */
const FF_IDCT_SIMPLEAUTO = 128

const FF_THREAD_FRAME = 1 ///< Decode more than one frame at once
const FF_THREAD_SLICE = 2 ///< Decode more than one part of a single frame at once

const FF_PROFILE_UNKNOWN = -99
const FF_PROFILE_RESERVED = -100

const FF_PROFILE_AAC_MAIN = 0
const FF_PROFILE_AAC_LOW = 1
const FF_PROFILE_AAC_SSR = 2
const FF_PROFILE_AAC_LTP = 3
const FF_PROFILE_AAC_HE = 4
const FF_PROFILE_AAC_HE_V2 = 28
const FF_PROFILE_AAC_LD = 22
const FF_PROFILE_AAC_ELD = 38
const FF_PROFILE_MPEG2_AAC_LOW = 128
const FF_PROFILE_MPEG2_AAC_HE = 131

const FF_PROFILE_DNXHD = 0
const FF_PROFILE_DNXHR_LB = 1
const FF_PROFILE_DNXHR_SQ = 2
const FF_PROFILE_DNXHR_HQ = 3
const FF_PROFILE_DNXHR_HQX = 4
const FF_PROFILE_DNXHR_444 = 5

const FF_PROFILE_DTS = 20
const FF_PROFILE_DTS_ES = 30
const FF_PROFILE_DTS_96_24 = 40
const FF_PROFILE_DTS_HD_HRA = 50
const FF_PROFILE_DTS_HD_MA = 60
const FF_PROFILE_DTS_EXPRESS = 70

const FF_PROFILE_MPEG2_422 = 0
const FF_PROFILE_MPEG2_HIGH = 1
const FF_PROFILE_MPEG2_SS = 2
const FF_PROFILE_MPEG2_SNR_SCALABLE = 3
const FF_PROFILE_MPEG2_MAIN = 4
const FF_PROFILE_MPEG2_SIMPLE = 5

const FF_PROFILE_H264_CONSTRAINED = (1 << 9) // 8+1; constraint_set1_flag
const FF_PROFILE_H264_INTRA = (1 << 11)      // 8+3; constraint_set3_flag

const FF_PROFILE_H264_BASELINE = 66
const FF_PROFILE_H264_CONSTRAINED_BASELINE = (66 | FF_PROFILE_H264_CONSTRAINED)
const FF_PROFILE_H264_MAIN = 77
const FF_PROFILE_H264_EXTENDED = 88
const FF_PROFILE_H264_HIGH = 100
const FF_PROFILE_H264_HIGH_10 = 110
const FF_PROFILE_H264_HIGH_10_INTRA = (110 | FF_PROFILE_H264_INTRA)
const FF_PROFILE_H264_MULTIVIEW_HIGH = 118
const FF_PROFILE_H264_HIGH_422 = 122
const FF_PROFILE_H264_HIGH_422_INTRA = (122 | FF_PROFILE_H264_INTRA)
const FF_PROFILE_H264_STEREO_HIGH = 128
const FF_PROFILE_H264_HIGH_444 = 144
const FF_PROFILE_H264_HIGH_444_PREDICTIVE = 244
const FF_PROFILE_H264_HIGH_444_INTRA = (244 | FF_PROFILE_H264_INTRA)
const FF_PROFILE_H264_CAVLC_444 = 44

const FF_PROFILE_VC1_SIMPLE = 0
const FF_PROFILE_VC1_MAIN = 1
const FF_PROFILE_VC1_COMPLEX = 2
const FF_PROFILE_VC1_ADVANCED = 3

const FF_PROFILE_MPEG4_SIMPLE = 0
const FF_PROFILE_MPEG4_SIMPLE_SCALABLE = 1
const FF_PROFILE_MPEG4_CORE = 2
const FF_PROFILE_MPEG4_MAIN = 3
const FF_PROFILE_MPEG4_N_BIT = 4
const FF_PROFILE_MPEG4_SCALABLE_TEXTURE = 5
const FF_PROFILE_MPEG4_SIMPLE_FACE_ANIMATION = 6
const FF_PROFILE_MPEG4_BASIC_ANIMATED_TEXTURE = 7
const FF_PROFILE_MPEG4_HYBRID = 8
const FF_PROFILE_MPEG4_ADVANCED_REAL_TIME = 9
const FF_PROFILE_MPEG4_CORE_SCALABLE = 10
const FF_PROFILE_MPEG4_ADVANCED_CODING = 11
const FF_PROFILE_MPEG4_ADVANCED_CORE = 12
const FF_PROFILE_MPEG4_ADVANCED_SCALABLE_TEXTURE = 13
const FF_PROFILE_MPEG4_SIMPLE_STUDIO = 14
const FF_PROFILE_MPEG4_ADVANCED_SIMPLE = 15

const FF_PROFILE_JPEG2000_CSTREAM_RESTRICTION_0 = 1
const FF_PROFILE_JPEG2000_CSTREAM_RESTRICTION_1 = 2
const FF_PROFILE_JPEG2000_CSTREAM_NO_RESTRICTION = 32768
const FF_PROFILE_JPEG2000_DCINEMA_2K = 3
const FF_PROFILE_JPEG2000_DCINEMA_4K = 4

const FF_PROFILE_VP9_0 = 0
const FF_PROFILE_VP9_1 = 1
const FF_PROFILE_VP9_2 = 2
const FF_PROFILE_VP9_3 = 3

const FF_PROFILE_HEVC_MAIN = 1
const FF_PROFILE_HEVC_MAIN_10 = 2
const FF_PROFILE_HEVC_MAIN_STILL_PICTURE = 3
const FF_PROFILE_HEVC_REXT = 4

const FF_PROFILE_VVC_MAIN_10 = 1
const FF_PROFILE_VVC_MAIN_10_444 = 33

const FF_PROFILE_AV1_MAIN = 0
const FF_PROFILE_AV1_HIGH = 1
const FF_PROFILE_AV1_PROFESSIONAL = 2

const FF_PROFILE_MJPEG_HUFFMAN_BASELINE_DCT = 0xc0
const FF_PROFILE_MJPEG_HUFFMAN_EXTENDED_SEQUENTIAL_DCT = 0xc1
const FF_PROFILE_MJPEG_HUFFMAN_PROGRESSIVE_DCT = 0xc2
const FF_PROFILE_MJPEG_HUFFMAN_LOSSLESS = 0xc3
const FF_PROFILE_MJPEG_JPEG_LS = 0xf7

const FF_PROFILE_SBC_MSBC = 1

const FF_PROFILE_PRORES_PROXY = 0
const FF_PROFILE_PRORES_LT = 1
const FF_PROFILE_PRORES_STANDARD = 2
const FF_PROFILE_PRORES_HQ = 3
const FF_PROFILE_PRORES_4444 = 4
const FF_PROFILE_PRORES_XQ = 5

const FF_PROFILE_ARIB_PROFILE_A = 0
const FF_PROFILE_ARIB_PROFILE_C = 1

const FF_PROFILE_KLVA_SYNC = 0
const FF_PROFILE_KLVA_ASYNC = 1

const FF_LEVEL_UNKNOWN = -99

const FF_SUB_CHARENC_MODE_DO_NOTHING = -1 ///< do nothing (demuxer outputs a stream supposed to be already in UTF-8, or the codec is bitmap for instance)
const FF_SUB_CHARENC_MODE_AUTOMATIC = 0   ///< libavcodec will select the mode itself
const FF_SUB_CHARENC_MODE_PRE_DECODER = 1 ///< the AVPacket data needs to be recoded to UTF-8 before being fed to the decoder, requires iconv
const FF_SUB_CHARENC_MODE_IGNORE = 2      ///< neither convert the subtitles, nor check them for valid UTF-8

const FF_DEBUG_VIS_MV_P_FOR = 0x00000001  //visualize forward predicted MVs of P frames
const FF_DEBUG_VIS_MV_B_FOR = 0x00000002  //visualize forward predicted MVs of B frames
const FF_DEBUG_VIS_MV_B_BACK = 0x00000004 //visualize backward predicted MVs of B frames

const FF_CODEC_PROPERTY_LOSSLESS = 0x00000001
const FF_CODEC_PROPERTY_CLOSED_CAPTIONS = 0x00000002

const FF_SUB_TEXT_FMT_ASS = 0

//#if FF_API_ASS_TIMING
const FF_SUB_TEXT_FMT_ASS_WITH_TIMINGS = 1

//#endif

/**
 * HWAccel is experimental and is thus avoided in favor of non experimental
 * codecs
 */
const AV_HWACCEL_CODEC_CAP_EXPERIMENTAL = 0x0200

/**
 * Hardware acceleration should be used for decoding even if the codec level
 * used is unknown or higher than the maximum supported level reported by the
 * hardware driver.
 *
 * It's generally a good idea to pass this flag unless you have a specific
 * reason not to, as hardware tends to under-report supported levels.
 */
const AV_HWACCEL_FLAG_IGNORE_LEVEL = (1 << 0)

/**
 * Hardware acceleration can output YUV pixel formats with a different chroma
 * sampling than 4:2:0 and/or other than 8 bits per component.
 */
const AV_HWACCEL_FLAG_ALLOW_HIGH_DEPTH = (1 << 1)

/**
 * Hardware acceleration should still be attempted for decoding when the
 * codec profile does not match the reported capabilities of the hardware.
 *
 * For example, this can be used to try to decode baseline profile H.264
 * streams in hardware - it will often succeed, because many streams marked
 * as baseline profile actually conform to constrained baseline profile.
 *
 * @warning If the stream is actually not supported then the behaviour is
 *          undefined, and may include returning entirely incorrect output
 *          while indicating success.
 */
const AV_HWACCEL_FLAG_ALLOW_PROFILE_MISMATCH = (1 << 2)

type AVSubtitleType int32

const (
	SUBTITLE_NONE = 0

	SUBTITLE_BITMAP ///< A bitmap, pict will be set

	/**
	 * Plain text, the text field must be set by the decoder and is
	 * authoritative. ass and pict fields may contain approximations.
	 */
	SUBTITLE_TEXT

	/**
	 * Formatted text, the ass field must be set by the decoder and is
	 * authoritative. pict and text fields may contain approximations.
	 */
	SUBTITLE_ASS
)

const AV_SUBTITLE_FLAG_FORCED = 0x00000001

/**
 * @defgroup lavc_parsing Frame parsing
 * @{
 */
type AVPictureStructure int32

const (
	AV_PICTURE_STRUCTURE_UNKNOWN      = 0 //< unknown
	AV_PICTURE_STRUCTURE_TOP_FIELD        //< coded as top field
	AV_PICTURE_STRUCTURE_BOTTOM_FIELD     //< coded as bottom field
	AV_PICTURE_STRUCTURE_FRAME            //< coded as frame
)

const AV_PARSER_PTS_NB = 4
const PARSER_FLAG_COMPLETE_FRAMES = 0x0001
const PARSER_FLAG_ONCE = 0x0002

/// Set if the parser has a valid file offset
const PARSER_FLAG_FETCHED_OFFSET = 0x0004
const PARSER_FLAG_USE_CODEC_TS = 0x1000

/**
 * Lock operation used by lockmgr
 *
 * @deprecated Deprecated together with av_lockmgr_register().
 */
type AVLockOp int32

const (
	AV_LOCK_CREATE  = 0 ///< Create a mutex
	AV_LOCK_OBTAIN      ///< Lock the mutex
	AV_LOCK_RELEASE     ///< Unlock the mutex
	AV_LOCK_DESTROY     ///< Free mutex resources
)
