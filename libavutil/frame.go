package libavutil

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
)

/**
 * This structure describes decoded (raw) audio or video data.
 *
 * AVFrame must be allocated using av_frame_alloc(). Note that this only
 * allocates the AVFrame itself, the buffers for the data must be managed
 * through other means (see below).
 * AVFrame must be freed with av_frame_free().
 *
 * AVFrame is typically allocated once and then reused multiple times to hold
 * different data (e.g. a single AVFrame to hold frames received from a
 * decoder). In such a case, av_frame_unref() will free any references held by
 * the frame and reset it to its original clean state before it
 * is reused again.
 *
 * The data described by an AVFrame is usually reference counted through the
 * AVBuffer API. The underlying buffer references are stored in AVFrame.buf /
 * AVFrame.extended_buf. An AVFrame is considered to be reference counted if at
 * least one reference is set, i.e. if AVFrame.buf[0] != NULL. In such a case,
 * every single data plane must be contained in one of the buffers in
 * AVFrame.buf or AVFrame.extended_buf.
 * There may be a single buffer for all the data, or one separate buffer for
 * each plane, or anything in between.
 *
 * sizeof(AVFrame) is not a part of the public ABI, so new fields may be added
 * to the end with a minor bump.
 *
 * Fields can be accessed through AVOptions, the name string used, matches the
 * C structure field name for fields accessible through AVOptions. The AVClass
 * for AVFrame can be obtained from avcodec_get_frame_class()
 */
type AVFrame struct {
	//#define AV_NUM_DATA_POINTERS 8
	/**
	 * pointer to the picture/channel planes.
	 * This might be different from the first allocated byte
	 *
	 * Some decoders access areas outside 0,0 - width,height, please
	 * see avcodec_align_dimensions2(). Some filters and swscale can read
	 * up to 16 bytes beyond the planes, if these filters are to be used,
	 * then 16 extra bytes must be allocated.
	 *
	 * NOTE: Except for hwaccel formats, pointers not needed by the format
	 * MUST be set to NULL.
	 */
	Data *[ffconstant.AV_NUM_DATA_POINTERS]ffcommon.FInt

	/**
	 * For video, size in bytes of each picture line.
	 * For audio, size in bytes of each plane.
	 *
	 * For audio, only linesize[0] may be set. For planar audio, each channel
	 * plane must be the same size.
	 *
	 * For video the linesizes should be multiples of the CPUs alignment
	 * preference, this is 16 or 32 for modern desktop CPUs.
	 * Some code requires such alignment other code can be slower without
	 * correct alignment, for yet other it makes no difference.
	 *
	 * @note The linesize may be larger than the size of usable data -- there
	 * may be extra padding present for performance reasons.
	 */
	Linesize [ffconstant.AV_NUM_DATA_POINTERS]ffcommon.FInt

	/**
	 * pointers to the data planes/channels.
	 *
	 * For video, this should simply point to data[].
	 *
	 * For planar audio, each channel has a separate data pointer, and
	 * linesize[0] contains the size of each channel buffer.
	 * For packed audio, there is just one data pointer, and linesize[0]
	 * contains the total size of the buffer for all channels.
	 *
	 * Note: Both data and extended_data should always be set in a valid frame,
	 * but for planar audio with more channels that can fit in data,
	 * extended_data must be used in order to access all channels.
	 */
	ExtendedData **ffcommon.FUint8T

	/**
	 * @name Video dimensions
	 * Video frames only. The coded dimensions (in pixels) of the video frame,
	 * i.e. the size of the rectangle that contains some well-defined values.
	 *
	 * @note The part of the frame intended for display/presentation is further
	 * restricted by the @ref cropping "Cropping rectangle".
	 * @{
	 */
	Width, Height ffcommon.FInt
	/**
	 * @}
	 */

	/**
	 * number of audio samples (per channel) described by this frame
	 */
	NbSamples ffcommon.FInt

	/**
	 * format of the frame, -1 if unknown or unset
	 * Values correspond to enum AVPixelFormat for video frames,
	 * enum AVSampleFormat for audio)
	 */
	Format ffcommon.FInt

	/**
	 * 1 -> keyframe, 0-> not
	 */
	KeyFrame ffcommon.FInt

	/**
	 * Picture type of the frame.
	 */
	PictType ffconstant.AVPictureType

	/**
	 * Sample aspect ratio for the video frame, 0/1 if unknown/unspecified.
	 */
	SampleAspectRatio AVRational

	/**
	 * Presentation timestamp in time_base units (time when frame should be shown to user).
	 */
	Pts ffcommon.FInt64T

	//#if FF_API_PKT_PTS
	///**
	// * PTS copied from the AVPacket that was decoded to produce this frame.
	// * @deprecated use the pts field instead
	// */
	//attribute_deprecated
	PktPts ffcommon.FInt64T
	//#endif

	/**
	 * DTS copied from the AVPacket that triggered returning this frame. (if frame threading isn't used)
	 * This is also the Presentation time of this AVFrame calculated from
	 * only AVPacket.dts values without pts values.
	 */
	PktDts ffcommon.FInt64T

	/**
	 * picture number in bitstream order
	 */
	CodedPictureNumber ffcommon.FInt
	/**
	 * picture number in display order
	 */
	DisplayPictureNumber ffcommon.FInt

	/**
	 * quality (between 1 (good) and FF_LAMBDA_MAX (bad))
	 */
	Quality ffcommon.FInt

	/**
	 * for some private data of the user
	 */
	Opaque ffcommon.FVoidP

	//#if FF_API_ERROR_FRAME
	/**
	 * @deprecated unused
	 */
	//attribute_deprecated
	Error0 [ffconstant.AV_NUM_DATA_POINTERS]ffcommon.FUint64T
	//#endif

	/**
	 * When decoding, this signals how much the picture must be delayed.
	 * extra_delay = repeat_pict / (2*fps)
	 */
	RepeatPict ffcommon.FInt

	/**
	 * The content of the picture is interlaced.
	 */
	InterlacedFrame ffcommon.FInt

	/**
	 * If the content is interlaced, is top field displayed first.
	 */
	TopFieldFirst ffcommon.FInt

	/**
	 * Tell user application that palette has changed from previous frame.
	 */
	PaletteHasChanged ffcommon.FInt

	/**
	 * reordered opaque 64 bits (generally an integer or a double precision float
	 * PTS but can be anything).
	 * The user sets AVCodecContext.reordered_opaque to represent the input at
	 * that time,
	 * the decoder reorders values as needed and sets AVFrame.reordered_opaque
	 * to exactly one of the values provided by the user through AVCodecContext.reordered_opaque
	 */
	ReorderedOpaque ffcommon.FInt64T

	/**
	 * Sample rate of the audio data.
	 */
	SampleRate ffcommon.FInt

	/**
	 * Channel layout of the audio data.
	 */
	ChannelLayout ffcommon.FUint64T

	/**
	 * AVBuffer references backing the data for this frame. If all elements of
	 * this array are NULL, then this frame is not reference counted. This array
	 * must be filled contiguously -- if buf[i] is non-NULL then buf[j] must
	 * also be non-NULL for all j < i.
	 *
	 * There may be at most one AVBuffer per data plane, so for video this array
	 * always contains all the references. For planar audio with more than
	 * AV_NUM_DATA_POINTERS channels, there may be more buffers than can fit in
	 * this array. Then the extra AVBufferRef pointers are stored in the
	 * extended_buf array.
	 */
	Buf *[ffconstant.AV_NUM_DATA_POINTERS]*AVBufferRef

	/**
	 * For planar audio which requires more than AV_NUM_DATA_POINTERS
	 * AVBufferRef pointers, this array will hold all the references which
	 * cannot fit into AVFrame.buf.
	 *
	 * Note that this is different from AVFrame.extended_data, which always
	 * contains all the pointers. This array only contains the extra pointers,
	 * which cannot fit into AVFrame.buf.
	 *
	 * This array is always allocated using av_malloc() by whoever constructs
	 * the frame. It is freed in av_frame_unref().
	 */
	ExtendedBuf **AVBufferRef
	/**
	 * Number of elements in extended_buf.
	 */
	NbExtendedBuf ffcommon.FInt

	Side_Data  **AVFrameSideData
	NbSideData ffcommon.FInt

	/**
	 * @defgroup lavu_frame_flags AV_FRAME_FLAGS
	 * @ingroup lavu_frame
	 * Flags describing additional frame properties.
	 *
	 * @{
	 */

	/**
	 * The frame data may be corrupted, e.g. due to decoding errors.
	 */
	//#define AV_FRAME_FLAG_CORRUPT       (1 << 0)
	/**
	 * A flag to mark the frames which need to be decoded, but shouldn't be output.
	 */
	//#define AV_FRAME_FLAG_DISCARD   (1 << 2)
	/**
	 * @}
	 */

	/**
	 * Frame flags, a combination of @ref lavu_frame_flags
	 */
	Flags ffcommon.FInt

	/**
	 * MPEG vs JPEG YUV range.
	 * - encoding: Set by user
	 * - decoding: Set by libavcodec
	 */
	ColorRange ffconstant.AVColorRange

	ColorPrimaries ffconstant.AVColorPrimaries

	ColorTrc ffconstant.AVColorTransferCharacteristic

	/**
	 * YUV colorspace type.
	 * - encoding: Set by user
	 * - decoding: Set by libavcodec
	 */
	Colorspace ffconstant.AVColorSpace

	ChromaLocation ffconstant.AVChromaLocation

	/**
	 * frame timestamp estimated using various heuristics, in stream time base
	 * - encoding: unused
	 * - decoding: set by libavcodec, read by user.
	 */
	BestEffortTimestamp ffcommon.FInt64T

	/**
	 * reordered pos from the last AVPacket that has been input into the decoder
	 * - encoding: unused
	 * - decoding: Read by user.
	 */
	PktPos ffcommon.FInt64T

	/**
	 * duration of the corresponding packet, expressed in
	 * AVStream->time_base units, 0 if unknown.
	 * - encoding: unused
	 * - decoding: Read by user.
	 */
	PktDuration ffcommon.FInt64T

	/**
	 * metadata.
	 * - encoding: Set by user.
	 * - decoding: Set by libavcodec.
	 */
	Metadata *AVDictionary

	/**
	 * decode error flags of the frame, set to a combination of
	 * FF_DECODE_ERROR_xxx flags if the decoder produced a frame, but there
	 * were errors during the decoding.
	 * - encoding: unused
	 * - decoding: set by libavcodec, read by user.
	 */
	DecodeErrorFlags ffcommon.FInt
	//#define FF_DECODE_ERROR_INVALID_BITSTREAM   1
	//#define FF_DECODE_ERROR_MISSING_REFERENCE   2
	//#define FF_DECODE_ERROR_CONCEALMENT_ACTIVE  4
	//#define FF_DECODE_ERROR_DECODE_SLICES       8

	/**
	 * number of audio channels, only used for audio.
	 * - encoding: unused
	 * - decoding: Read by user.
	 */
	Channels ffcommon.FInt

	/**
	 * size of the corresponding packet containing the compressed
	 * frame.
	 * It is set to a negative value if unknown.
	 * - encoding: unused
	 * - decoding: set by libavcodec, read by user.
	 */
	PktSize ffcommon.FInt

	//#if FF_API_FRAME_QP
	/**
	 * QP table
	 */
	//attribute_deprecated
	QscaleTable *ffcommon.FUint8T
	/**
	 * QP store stride
	 */
	//attribute_deprecated
	Qstride ffcommon.FInt

	//attribute_deprecated
	QscaleType ffcommon.FInt

	//attribute_deprecated
	QpTableBuf *AVBufferRef
	//#endif
	/**
	 * For hwaccel-format frames, this should be a reference to the
	 * AVHWFramesContext describing the frame.
	 */
	HwFramesCtx *AVBufferRef

	/**
	 * AVBufferRef for free use by the API user. FFmpeg will never check the
	 * contents of the buffer ref. FFmpeg calls av_buffer_unref() on it when
	 * the frame is unreferenced. av_frame_copy_props() calls create a new
	 * reference with av_buffer_ref() for the target frame's opaque_ref field.
	 *
	 * This is unrelated to the opaque field, although it serves a similar
	 * purpose.
	 */
	OpaqueRef *AVBufferRef

	/**
	 * @anchor cropping
	 * @name Cropping
	 * Video frames only. The number of pixels to discard from the the
	 * top/bottom/left/right border of the frame to obtain the sub-rectangle of
	 * the frame intended for presentation.
	 * @{
	 */
	CropTop    ffcommon.FSizeT
	CropBottom ffcommon.FSizeT
	CropLeft   ffcommon.FSizeT
	CropRight  ffcommon.FSizeT
	/**
	 * @}
	 */

	/**
	 * AVBufferRef for internal use by a single libav* library.
	 * Must not be used to transfer data between libraries.
	 * Has to be NULL when ownership of the frame leaves the respective library.
	 *
	 * Code outside the FFmpeg libs should never check or change the contents of the buffer ref.
	 *
	 * FFmpeg calls av_buffer_unref() on it when the frame is unreferenced.
	 * av_frame_copy_props() calls create a new reference with av_buffer_ref()
	 * for the target frame's private_ref field.
	 */
	PrivateRef *AVBufferRef
}
