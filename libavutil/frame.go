package libavutil

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"unsafe"
)

/**
 * Structure to hold side data for an AVFrame.
 *
 * sizeof(AVFrameSideData) is not a part of the public ABI, so new fields may be added
 * to the end with a minor bump.
 */
type AVFrameSideData struct {
	type0 ffconstant.AVFrameSideDataType
	data  *ffcommon.FUint8T
	//#if FF_API_BUFFER_SIZE_T
	//int      size;
	//#else
	size ffcommon.FSizeT
	//#endif
	metadata *AVDictionary
	buf      *AVBufferRef
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

	/**
	 * Must be set to the size of this data structure (that is,
	 * sizeof(AVRegionOfInterest)).
	 */
	self_size ffcommon.FUint32T
	/**
	 * Distance in pixels from the top edge of the frame to the top and
	 * bottom edges and from the left edge of the frame to the left and
	 * right edges of the rectangle defining this region of interest.
	 *
	 * The constraints on a region are encoder dependent, so the region
	 * actually affected may be slightly larger for alignment or other
	 * reasons.
	 */
	top    ffcommon.FInt
	bottom ffcommon.FInt
	left   ffcommon.FInt
	right  ffcommon.FInt
	/**
	 * Quantisation offset.
	 *
	 * Must be in the range -1 to +1.  A value of zero indicates no quality
	 * change.  A negative value asks for better quality (less quantisation),
	 * while a positive value asks for worse quality (greater quantisation).
	 *
	 * The range is calibrated so that the extreme values indicate the
	 * largest possible offset - if the rest of the frame is encoded with the
	 * worst possible quality, an offset of -1 indicates that this region
	 * should be encoded with the best possible quality anyway.  Intermediate
	 * values are then interpolated in some codec-dependent way.
	 *
	 * For example, in 10-bit H.264 the quantisation parameter varies between
	 * -12 and 51.  A typical qoffset value of -1/10 therefore indicates that
	 * this region should be encoded with a QP around one-tenth of the full
	 * range better than the rest of the frame.  So, if most of the frame
	 * were to be encoded with a QP of around 30, this region would get a QP
	 * of around 24 (an offset of approximately -1/10 * (51 - -12) = -6.3).
	 * An extreme value of -1 would indicate that this region should be
	 * encoded with the best possible quality regardless of the treatment of
	 * the rest of the frame - that is, should be encoded at a QP of -12.
	 */
	qoffset AVRational
}

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

//#if FF_API_FRAME_GET_SET
/**
 * Accessors for some AVFrame fields. These used to be provided for ABI
 * compatibility, and do not need to be used anymore.
 */
//attribute_deprecated
//int64_t av_frame_get_best_effort_timestamp(const AVFrame *frame);
//未测试
func (frame *AVFrame) AvFrameGetBestEffortTimestamp() (res ffcommon.FInt64T, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_get_best_effort_timestamp").Call(
		uintptr(unsafe.Pointer(frame)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt64T(t)
	return
}

//attribute_deprecated
//void    av_frame_set_best_effort_timestamp(AVFrame *frame, int64_t val);
//未测试
func (frame *AVFrame) AvFrameSetBestEffortTimestamp(val ffcommon.FInt64T) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_set_best_effort_timestamp").Call(
		uintptr(unsafe.Pointer(frame)),
		uintptr(val),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

//attribute_deprecated
//int64_t av_frame_get_pkt_duration         (const AVFrame *frame);
//未测试
func (frame *AVFrame) AvFrameGetPktDuration() (res ffcommon.FInt64T, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_get_pkt_duration").Call(
		uintptr(unsafe.Pointer(frame)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt64T(t)
	return
}

//attribute_deprecated
//void    av_frame_set_pkt_duration         (AVFrame *frame, int64_t val);
//未测试
func (frame *AVFrame) AvFrameSetPktDuration(val ffcommon.FInt64T) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_set_pkt_duration").Call(
		uintptr(unsafe.Pointer(frame)),
		uintptr(val),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

//attribute_deprecated
//int64_t av_frame_get_pkt_pos              (const AVFrame *frame);
//未测试
func (frame *AVFrame) AvFrameGetPktPos() (res ffcommon.FInt64T, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_get_pkt_pos").Call(
		uintptr(unsafe.Pointer(frame)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt64T(t)
	return
}

//attribute_deprecated
//void    av_frame_set_pkt_pos              (AVFrame *frame, int64_t val);
//未测试
func (frame *AVFrame) AvFrameSetPktPos(val ffcommon.FInt64T) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_set_pkt_pos").Call(
		uintptr(unsafe.Pointer(frame)),
		uintptr(val),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

//attribute_deprecated
//int64_t av_frame_get_channel_layout       (const AVFrame *frame);
//未测试
func (frame *AVFrame) AvFrameGetChannelLayout() (res ffcommon.FInt64T, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_get_channel_layout").Call(
		uintptr(unsafe.Pointer(frame)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt64T(t)
	return
}

//attribute_deprecated
//void    av_frame_set_channel_layout       (AVFrame *frame, int64_t val);
//未测试
func (frame *AVFrame) AvFrameSetChannelLayout(val ffcommon.FInt64T) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_set_channel_layout").Call(
		uintptr(unsafe.Pointer(frame)),
		uintptr(val),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

//attribute_deprecated
//int     av_frame_get_channels             (const AVFrame *frame);
//未测试
func (frame *AVFrame) AvFrameGetChannels() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_get_channels").Call(
		uintptr(unsafe.Pointer(frame)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

//attribute_deprecated
//void    av_frame_set_channels             (AVFrame *frame, int     val);
//未测试
func (frame *AVFrame) AvFrameSetChannels(val ffcommon.FInt) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_set_channels").Call(
		uintptr(unsafe.Pointer(frame)),
		uintptr(val),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

//attribute_deprecated
//int     av_frame_get_sample_rate          (const AVFrame *frame);
//未测试
func (frame *AVFrame) AvFrameGetSampleRate() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_get_sample_rate").Call(
		uintptr(unsafe.Pointer(frame)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

//attribute_deprecated
//void    av_frame_set_sample_rate          (AVFrame *frame, int     val);
//未测试
func (frame *AVFrame) AvFrameSetSampleRate(val ffcommon.FInt) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_set_sample_rate").Call(
		uintptr(unsafe.Pointer(frame)),
		uintptr(val),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

//attribute_deprecated
//AVDictionary *av_frame_get_metadata       (const AVFrame *frame);
//未测试
func (frame *AVFrame) AvFrameGetMetadata() (res *AVDictionary, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_get_metadata").Call(
		uintptr(unsafe.Pointer(frame)),
	)
	if err != nil {
		//return
	}
	res = (*AVDictionary)(unsafe.Pointer(t))
	return
}

//attribute_deprecated
//void          av_frame_set_metadata       (AVFrame *frame, AVDictionary *val);
//未测试
func (frame *AVFrame) AvFrameSetMetadata(val *AVDictionary) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_set_metadata").Call(
		uintptr(unsafe.Pointer(frame)),
		uintptr(unsafe.Pointer(val)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

//attribute_deprecated
//int     av_frame_get_decode_error_flags   (const AVFrame *frame);
//未测试
func (frame *AVFrame) AvFrameGetDecodeErrorFlags() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_get_decode_error_flags").Call(
		uintptr(unsafe.Pointer(frame)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

//attribute_deprecated
//void    av_frame_set_decode_error_flags   (AVFrame *frame, int     val);
//未测试
func (frame *AVFrame) AvFrameSetDecodeErrorFlags(val ffcommon.FInt) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_set_decode_error_flags").Call(
		uintptr(unsafe.Pointer(frame)),
		uintptr(val),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

//attribute_deprecated
//int     av_frame_get_pkt_size(const AVFrame *frame);
//未测试
func (frame *AVFrame) AvFrameGet_pPktSize() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_get_pkt_size").Call(
		uintptr(unsafe.Pointer(frame)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

//attribute_deprecated
//void    av_frame_set_pkt_size(AVFrame *frame, int val);
//未测试
func (frame *AVFrame) AvFrameSetPktSize(val ffcommon.FInt) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_set_pkt_size").Call(
		uintptr(unsafe.Pointer(frame)),
		uintptr(val),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

//#if FF_API_FRAME_QP
//attribute_deprecated
//int8_t *av_frame_get_qp_table(AVFrame *f, int *stride, int *type);
//未测试
func (f *AVFrame) AvFrameGetQpTable(stride *ffcommon.FInt, type0 *ffcommon.FInt) (res ffcommon.FUint8T, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_get_qp_table").Call(
		uintptr(unsafe.Pointer(f)),
		uintptr(unsafe.Pointer(stride)),
		uintptr(unsafe.Pointer(type0)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FUint8T(t)
	return
}

//attribute_deprecated
//int av_frame_set_qp_table(AVFrame *f, AVBufferRef *buf, int stride, int type);
//未测试
func (f *AVFrame) AvFrameSetQpTable(buf *AVBufferRef, stride ffcommon.FInt, type0 ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_set_qp_table").Call(
		uintptr(unsafe.Pointer(f)),
		uintptr(unsafe.Pointer(buf)),
		uintptr(stride),
		uintptr(type0),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//#endif
//attribute_deprecated
//enum AVColorSpace av_frame_get_colorspace(const AVFrame *frame);
//未测试
func (f *AVFrame) AvFrameGetColorspace() (res ffconstant.AVColorSpace, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_get_colorspace").Call(
		uintptr(unsafe.Pointer(f)),
	)
	if err != nil {
		//return
	}
	res = ffconstant.AVColorSpace(t)
	return
}

//attribute_deprecated
//void    av_frame_set_colorspace(AVFrame *frame, enum AVColorSpace val);
//未测试
func (f *AVFrame) AvFrameSetColorspace(val ffconstant.AVColorSpace) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_set_colorspace").Call(
		uintptr(unsafe.Pointer(f)),
		uintptr(val),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

//attribute_deprecated
//enum AVColorRange av_frame_get_color_range(const AVFrame *frame);
//未测试
func (frame *AVFrame) AvFrameGetColorRange() (res ffconstant.AVColorRange, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_get_color_range").Call(
		uintptr(unsafe.Pointer(frame)),
	)
	if err != nil {
		//return
	}
	res = ffconstant.AVColorRange(t)
	return
}

//attribute_deprecated
//void    av_frame_set_color_range(AVFrame *frame, enum AVColorRange val);
//未测试
func (f *AVFrame) AvFrameSetColorRange(val ffconstant.AVColorRange) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_set_color_range").Call(
		uintptr(unsafe.Pointer(f)),
		uintptr(val),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

//#endif

/**
 * Get the name of a colorspace.
 * @return a static string identifying the colorspace; can be NULL.
 */
//const char *av_get_colorspace_name(enum AVColorSpace val);
//未测试
func AvGetColorspaceName(val ffconstant.AVColorSpace) (res ffcommon.FConstCharP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_get_colorspace_name").Call(
		uintptr(val),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.GoAStr(t)
	return
}

/**
 * Allocate an AVFrame and set its fields to default values.  The resulting
 * struct must be freed using av_frame_free().
 *
 * @return An AVFrame filled with default values or NULL on failure.
 *
 * @note this only allocates the AVFrame itself, not the data buffers. Those
 * must be allocated through other means, e.g. with av_frame_get_buffer() or
 * manually.
 */
//AVFrame *av_frame_alloc(void);
//未测试
func AvFrameAlloc() (res *AVFrame, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_alloc").Call()
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = (*AVFrame)(unsafe.Pointer(t))
	return
}

/**
 * Free the frame and any dynamically allocated objects in it,
 * e.g. extended_data. If the frame is reference counted, it will be
 * unreferenced first.
 *
 * @param frame frame to be freed. The pointer will be set to NULL.
 */
//void av_frame_free(AVFrame **frame);
//未测试
func AvFrameFree(frame **AVFrame) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_free").Call(
		uintptr(unsafe.Pointer(&frame)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Set up a new reference to the data described by the source frame.
 *
 * Copy frame properties from src to dst and create a new reference for each
 * AVBufferRef from src.
 *
 * If src is not reference counted, new buffers are allocated and the data is
 * copied.
 *
 * @warning: dst MUST have been either unreferenced with av_frame_unref(dst),
 *           or newly allocated with av_frame_alloc() before calling this
 *           function, or undefined behavior will occur.
 *
 * @return 0 on success, a negative AVERROR on error
 */
//int av_frame_ref(AVFrame *dst, const AVFrame *src);
//未测试
func AvFrameRef(dst *AVFrame, src *AVFrame) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_ref").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Create a new frame that references the same data as src.
 *
 * This is a shortcut for av_frame_alloc()+av_frame_ref().
 *
 * @return newly created AVFrame on success, NULL on error.
 */
//AVFrame *av_frame_clone(const AVFrame *src);
//未测试
func (src *AVFrame) AvFrameClone() (res *AVFrame, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_clone").Call(
		uintptr(unsafe.Pointer(src)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = (*AVFrame)(unsafe.Pointer(t))
	return
}

/**
 * Unreference all the buffers referenced by frame and reset the frame fields.
 */
//void av_frame_unref(AVFrame *frame);
//未测试
func (frame *AVFrame) AvFrameUnref() (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_unref").Call(
		uintptr(unsafe.Pointer(frame)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Move everything contained in src to dst and reset src.
 *
 * @warning: dst is not unreferenced, but directly overwritten without reading
 *           or deallocating its contents. Call av_frame_unref(dst) manually
 *           before calling this function to ensure that no memory is leaked.
 */
//void av_frame_move_ref(AVFrame *dst, AVFrame *src);
//未测试
func AvFrameMoveRef(dst *AVFrame, src *AVFrame) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_move_ref").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Allocate new buffer(s) for audio or video data.
 *
 * The following fields must be set on frame before calling this function:
 * - format (pixel format for video, sample format for audio)
 * - width and height for video
 * - nb_samples and channel_layout for audio
 *
 * This function will fill AVFrame.data and AVFrame.buf arrays and, if
 * necessary, allocate and fill AVFrame.extended_data and AVFrame.extended_buf.
 * For planar formats, one buffer will be allocated for each plane.
 *
 * @warning: if frame already has been allocated, calling this function will
 *           leak memory. In addition, undefined behavior can occur in certain
 *           cases.
 *
 * @param frame frame in which to store the new buffers.
 * @param align Required buffer size alignment. If equal to 0, alignment will be
 *              chosen automatically for the current CPU. It is highly
 *              recommended to pass 0 here unless you know what you are doing.
 *
 * @return 0 on success, a negative AVERROR on error.
 */
//int av_frame_get_buffer(AVFrame *frame, int align);
//未测试
func (frame *AVFrame) AvFrameGetBuffer(align ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_get_buffer").Call(
		uintptr(unsafe.Pointer(frame)),
		uintptr(align),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Check if the frame data is writable.
 *
 * @return A positive value if the frame data is writable (which is true if and
 * only if each of the underlying buffers has only one reference, namely the one
 * stored in this frame). Return 0 otherwise.
 *
 * If 1 is returned the answer is valid until av_buffer_ref() is called on any
 * of the underlying AVBufferRefs (e.g. through av_frame_ref() or directly).
 *
 * @see av_frame_make_writable(), av_buffer_is_writable()
 */
//int av_frame_is_writable(AVFrame *frame);
//未测试
func (frame *AVFrame) AvFrameIsWritable() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_is_writable").Call(
		uintptr(unsafe.Pointer(frame)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Ensure that the frame data is writable, avoiding data copy if possible.
 *
 * Do nothing if the frame is writable, allocate new buffers and copy the data
 * if it is not.
 *
 * @return 0 on success, a negative AVERROR on error.
 *
 * @see av_frame_is_writable(), av_buffer_is_writable(),
 * av_buffer_make_writable()
 */
//int av_frame_make_writable(AVFrame *frame);
//未测试
func (frame *AVFrame) AvFrameMakeWritable() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_make_writable").Call(
		uintptr(unsafe.Pointer(frame)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Copy the frame data from src to dst.
 *
 * This function does not allocate anything, dst must be already initialized and
 * allocated with the same parameters as src.
 *
 * This function only copies the frame data (i.e. the contents of the data /
 * extended data arrays), not any other properties.
 *
 * @return >= 0 on success, a negative AVERROR on error.
 */
//int av_frame_copy(AVFrame *dst, const AVFrame *src);
//未测试
func AvFrameCopy(dst *AVFrame, src *AVFrame) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_copy").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Copy only "metadata" fields from src to dst.
 *
 * Metadata for the purpose of this function are those fields that do not affect
 * the data layout in the buffers.  E.g. pts, sample rate (for audio) or sample
 * aspect ratio (for video), but not width/height or channel layout.
 * Side data is also copied.
 */
//int av_frame_copy_props(AVFrame *dst, const AVFrame *src);
//未测试
func (frame *AVFrame) AvFrameCopyProps(dst *AVFrame, src *AVFrame) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_copy_props").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Get the buffer reference a given data plane is stored in.
 *
 * @param plane index of the data plane of interest in frame->extended_data.
 *
 * @return the buffer reference that contains the plane or NULL if the input
 * frame is not valid.
 */
//AVBufferRef *av_frame_get_plane_buffer(AVFrame *frame, int plane);
//未测试
func (frame *AVFrame) AvFrameGetPlaneBuffer(plane ffcommon.FInt) (res *AVBufferRef, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_get_plane_buffer").Call(
		uintptr(unsafe.Pointer(frame)),
		uintptr(plane),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = (*AVBufferRef)(unsafe.Pointer(t))
	return
}

/**
 * Add a new side data to a frame.
 *
 * @param frame a frame to which the side data should be added
 * @param type type of the added side data
 * @param size size of the side data
 *
 * @return newly added side data on success, NULL on error
 */
//AVFrameSideData *av_frame_new_side_data(AVFrame *frame,
//enum AVFrameSideDataType type,
//#if FF_API_BUFFER_SIZE_T
//int size);
//#else
//size_t size);
//#endif
//未测试
func (frame *AVFrame) AvFrameNewSideData(type0 ffconstant.AVFrameSideDataType, size ffcommon.FSizeT) (res *AVFrameSideData, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_new_side_data").Call(
		uintptr(unsafe.Pointer(frame)),
		uintptr(type0),
		uintptr(size),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = (*AVFrameSideData)(unsafe.Pointer(t))
	return
}

/**
 * Add a new side data to a frame from an existing AVBufferRef
 *
 * @param frame a frame to which the side data should be added
 * @param type  the type of the added side data
 * @param buf   an AVBufferRef to add as side data. The ownership of
 *              the reference is transferred to the frame.
 *
 * @return newly added side data on success, NULL on error. On failure
 *         the frame is unchanged and the AVBufferRef remains owned by
 *         the caller.
 */
//AVFrameSideData *av_frame_new_side_data_from_buf(AVFrame *frame,
//enum AVFrameSideDataType type,
//AVBufferRef *buf);
//未测试
func (frame *AVFrame) AvFrameNewSideDataFromBuf(type0 ffconstant.AVFrameSideDataType, buf *AVBufferRef) (res *AVFrameSideData, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_new_side_data_from_buf").Call(
		uintptr(unsafe.Pointer(frame)),
		uintptr(type0),
		uintptr(unsafe.Pointer(buf)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = (*AVFrameSideData)(unsafe.Pointer(t))
	return
}

/**
 * @return a pointer to the side data of a given type on success, NULL if there
 * is no side data with such type in this frame.
 */
//AVFrameSideData *av_frame_get_side_data(const AVFrame *frame,
//enum AVFrameSideDataType type);
//未测试
func (frame *AVFrame) AvFrameGetSideData(type0 ffconstant.AVFrameSideDataType) (res *AVFrameSideData, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_get_side_data").Call(
		uintptr(unsafe.Pointer(frame)),
		uintptr(type0),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = (*AVFrameSideData)(unsafe.Pointer(t))
	return
}

/**
 * Remove and free all side data instances of the given type.
 */
//void av_frame_remove_side_data(AVFrame *frame, enum AVFrameSideDataType type);
//未测试
func (frame *AVFrame) AvFrameRemoveSideData(type0 ffconstant.AVFrameSideDataType) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_remove_side_data").Call(
		uintptr(unsafe.Pointer(frame)),
		uintptr(type0),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Crop the given video AVFrame according to its crop_left/crop_top/crop_right/
 * crop_bottom fields. If cropping is successful, the function will adjust the
 * data pointers and the width/height fields, and set the crop fields to 0.
 *
 * In all cases, the cropping boundaries will be rounded to the inherent
 * alignment of the pixel format. In some cases, such as for opaque hwaccel
 * formats, the left/top cropping is ignored. The crop fields are set to 0 even
 * if the cropping was rounded or ignored.
 *
 * @param frame the frame which should be cropped
 * @param flags Some combination of AV_FRAME_CROP_* flags, or 0.
 *
 * @return >= 0 on success, a negative AVERROR on error. If the cropping fields
 * were invalid, AVERROR(ERANGE) is returned, and nothing is changed.
 */
//int av_frame_apply_cropping(AVFrame *frame, int flags);
//未测试
func (frame *AVFrame) AvFrameApplyCropping(flags ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_apply_cropping").Call(
		uintptr(unsafe.Pointer(frame)),
		uintptr(flags),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * @return a string identifying the side data type
 */
//const char *av_frame_side_data_name(enum AVFrameSideDataType type);
//未测试
func AvFrameSideDataName(type0 ffconstant.AVFrameSideDataType) (res ffcommon.FConstCharP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_frame_side_data_name").Call(
		uintptr(type0),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.GoAStr(t)
	return
}
