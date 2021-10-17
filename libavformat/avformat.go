package libavformat

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"ffmpeg-go/libavutil"
)

//type AVFormatContext struct {
//}

type AVDeviceInfoList struct {
}
type AVDeviceCapabilitiesQuery struct {
}

/**
* @defgroup metadata_api Public Metadata API
* @{
* @ingroup libavf
* The metadata API allows libavformat to export metadata tags to a client
* application when demuxing. Conversely it allows a client application to
* set metadata when muxing.
*
* Metadata is exported or set as pairs of key/value strings in the 'metadata'
* fields of the AVFormatContext, AVStream, AVChapter and AVProgram structs
* using the @ref lavu_dict "AVDictionary" API. Like all strings in FFmpeg,
* metadata is assumed to be UTF-8 encoded Unicode. Note that metadata
* exported by demuxers isn't checked to be valid UTF-8 in most cases.
*
* Important concepts to keep in mind:
* -  Keys are unique; there can never be 2 tags with the same key. This is
*    also meant semantically, i.e., a demuxer should not knowingly produce
*    several keys that are literally different but semantically identical.
*    E.g., key=Author5, key=Author6. In this example, all authors must be
*    placed in the same tag.
* -  Metadata is flat, not hierarchical; there are no subtags. If you
*    want to store, e.g., the email address of the child of producer Alice
*    and actor Bob, that could have key=alice_and_bobs_childs_email_address.
* -  Several modifiers can be applied to the tag name. This is done by
*    appending a dash character ('-') and the modifier name in the order
*    they appear in the list below -- e.g. foo-eng-sort, not foo-sort-eng.
*    -  language -- a tag whose value is localized for a particular language
*       is appended with the ISO 639-2/B 3-letter language code.
*       For example: Author-ger=Michael, Author-eng=Mike
*       The original/default language is in the unqualified "Author" tag.
*       A demuxer should set a default if it sets any translated tag.
*    -  sorting  -- a modified version of a tag that should be used for
*       sorting will have '-sort' appended. E.g. artist="The Beatles",
*       artist-sort="Beatles, The".
* - Some protocols and demuxers support metadata updates. After a successful
*   call to av_read_frame(), AVFormatContext.event_flags or AVStream.event_flags
*   will be updated to indicate if metadata changed. In order to detect metadata
*   changes on a stream, you need to loop through all streams in the AVFormatContext
*   and check their individual event_flags.
*
* -  Demuxers attempt to export metadata in a generic format, however tags
*    with no generic equivalents are left as they are stored in the container.
*    Follows a list of generic tag names:
*
@verbatim
album        -- name of the set this work belongs to
album_artist -- main creator of the set/album, if different from artist.
                e.g. "Various Artists" for compilation albums.
artist       -- main creator of the work
comment      -- any additional description of the file.
composer     -- who composed the work, if different from artist.
copyright    -- name of copyright holder.
creation_time-- date when the file was created, preferably in ISO 8601.
date         -- date when the work was created, preferably in ISO 8601.
disc         -- number of a subset, e.g. disc in a multi-disc collection.
encoder      -- name/settings of the software/hardware that produced the file.
encoded_by   -- person/group who created the file.
filename     -- original name of the file.
genre        -- <self-evident>.
language     -- main language in which the work is performed, preferably
                in ISO 639-2 format. Multiple languages can be specified by
                separating them with commas.
performer    -- artist who performed the work, if different from artist.
                E.g for "Also sprach Zarathustra", artist would be "Richard
                Strauss" and performer "London Philharmonic Orchestra".
publisher    -- name of the label/publisher.
service_name     -- name of the service in broadcasting (channel name).
service_provider -- name of the service provider in broadcasting.
title        -- name of the work.
track        -- number of this work in the set, can be in form current/total.
variant_bitrate -- the total bitrate of the bitrate variant that the current stream is part of
@endverbatim
*
* Look in the examples section for an application example how to use the Metadata API.
*
* @}
*/

/* packet functions */

/**
 * Allocate and read the payload of a packet and initialize its
 * fields with default values.
 *
 * @param s    associated IO context
 * @param pkt packet
 * @param size desired payload size
 * @return >0 (read size) if OK, AVERROR_xxx otherwise
 */
//int av_get_packet(AVIOContext *s, AVPacket *pkt, int size);
func (s *AVIOContext) av_get_packet(pkt *AVPacket, size ffcommon.FInt) (res ffcommon.FInt) {
	return
}

/**
 * Read data and append it to the current content of the AVPacket.
 * If pkt->size is 0 this is identical to av_get_packet.
 * Note that this uses av_grow_packet and thus involves a realloc
 * which is inefficient. Thus this function should only be used
 * when there is no reasonable way to know (an upper bound of)
 * the final size.
 *
 * @param s    associated IO context
 * @param pkt packet
 * @param size amount of data to read
 * @return >0 (read size) if OK, AVERROR_xxx otherwise, previous data
 *         will not be lost even if an error occurs.
 */
//int av_append_packet(AVIOContext *s, AVPacket *pkt, int size);
func (s *AVIOContext) AvAppendPacket(pkt *AVPacket, size ffcommon.FInt) (res ffcommon.FInt) {
	return
}

/*************************************************/
/* input/output formats */

type AVCodecTag struct {
}

/**
 * This structure contains the data a format has to probe a file.
 */
type AVProbeData struct {
}

/**
 * @addtogroup lavf_encoding
 * @{
 */
type AVOutputFormat struct {
}

/**
 * @addtogroup lavf_decoding
 * @{
 */
type AVInputFormat struct {
}

type AVIndexEntry struct {
}

/**
 * Stream structure.
 * New fields can be added to the end with minor version bumps.
 * Removal, reordering and changes to existing fields require a major
 * version bump.
 * sizeof(AVStream) must not be used outside libav*.
 */
type AVStream struct {
}
type AVRational struct {
	Num ffcommon.FInt ///< Numerator
	Den ffcommon.FInt ///< Denominator
}

//#if FF_API_FORMAT_GET_SET
/**
 * Accessors for some AVStream fields. These used to be provided for ABI
 * compatibility, and do not need to be used anymore.
 */
//attribute_deprecated
//AVRational av_stream_get_r_frame_rate(const AVStream *s);
func (s *AVStream) AvStreamGetRFrameRate() (res AVRational) {
	return
}

//attribute_deprecated
//void       av_stream_set_r_frame_rate(AVStream *s, AVRational r);
func (s *AVStream) AvStreamSetRFrameRate(r AVRational) {
	return
}

//#if FF_API_LAVF_FFSERVER
//attribute_deprecated
//char* av_stream_get_recommended_encoder_configuration(const AVStream *s);
func (s *AVStream) AvStreamGetRecommendedEncoderConfiguration() (res ffcommon.FCharP) {
	return
}

//attribute_deprecated
//void  av_stream_set_recommended_encoder_configuration(AVStream *s, char *configuration);
func (s *AVStream) AvStreamSetRecommendedEncoderConfiguration(configuration ffcommon.FCharP) {
	return
}

//#endif
//#endif
type AVCodecParserContext struct {
}

//struct AVCodecParserContext *av_stream_get_parser(const AVStream *s);
func (s *AVStream) AvStreamGetParser() (res *AVCodecParserContext) {
	return
}

/**
 * Returns the pts of the last muxed packet + its duration
 *
 * the retuned value is undefined when used with a demuxer.
 */
//int64_t    av_stream_get_end_pts(const AVStream *st);
func (st *AVStream) AvStreamGetEndPts() (res ffcommon.FInt64T) {
	return
}

/**
 * New fields can be added to the end with minor version bumps.
 * Removal, reordering and changes to existing fields require a major
 * version bump.
 * sizeof(AVProgram) must not be used outside libav*.
 */
type AVProgram struct {
}
type AVChapter struct {
}

//typedef struct AVFormatInternal AVFormatInternal;

/**
 * Format I/O context.
 * New fields can be added to the end with minor version bumps.
 * Removal, reordering and changes to existing fields require a major
 * version bump.
 * sizeof(AVFormatContext) must not be used outside libav*, use
 * avformat_alloc_context() to create an AVFormatContext.
 *
 * Fields can be accessed through AVOptions (av_opt*),
 * the name string used matches the associated command line parameter name and
 * can be found in libavformat/options_table.h.
 * The AVOption/command line parameter names differ in some cases from the C
 * structure field names for historic reasons or brevity.
 */
type AVFormatContext struct {
}
type AVCodec struct {
}

//#if FF_API_FORMAT_GET_SET
/**
 * Accessors for some AVFormatContext fields. These used to be provided for ABI
 * compatibility, and do not need to be used anymore.
 */
//attribute_deprecated
//int av_format_get_probe_score(const AVFormatContext *s);
func (s *AVFormatContext) AvFormatGetProbeScore() (res ffcommon.FInt) {
	return
}

//attribute_deprecated
//AVCodec * av_format_get_video_codec(const AVFormatContext *s);
func (s *AVFormatContext) AvFormatGetVideoCodec() (res *AVCodec) {
	return
}

//attribute_deprecated
//void      av_format_set_video_codec(AVFormatContext *s, AVCodec *c);
func (s *AVFormatContext) AvFormatSetVideoCodec(c *AVCodec) {
	return
}

//attribute_deprecated
//AVCodec * av_format_get_audio_codec(const AVFormatContext *s);
func (s *AVFormatContext) AvFormatGetAudioCodec() (res *AVCodec) {
	return
}

//attribute_deprecated
//void      av_format_set_audio_codec(AVFormatContext *s, AVCodec *c);
func (s *AVFormatContext) AvFormatSetAudioCodec(c *AVCodec) {
	return
}

//attribute_deprecated
//AVCodec * av_format_get_subtitle_codec(const AVFormatContext *s);
func (s *AVFormatContext) AvFormatGetSubtitleCodec() (res *AVCodec) {
	return
}

//attribute_deprecated
//void      av_format_set_subtitle_codec(AVFormatContext *s, AVCodec *c);
func (s *AVFormatContext) AvFormatSetSubtitleCodec(c *AVCodec) {
	return
}

//attribute_deprecated
//AVCodec * av_format_get_data_codec(const AVFormatContext *s);
func (s *AVFormatContext) AvFormatGetDataCodec() (res *AVCodec) {
	return
}

//attribute_deprecated
//void      av_format_set_data_codec(AVFormatContext *s, AVCodec *c);
func (s *AVFormatContext) AvFormatSetDataCodec(c *AVCodec) {
	return
}

//attribute_deprecated
//int       av_format_get_metadata_header_padding(const AVFormatContext *s);
func (s *AVFormatContext) AvFormatGetMetadataHeaderPadding() (res ffcommon.FInt) {
	return
}

//attribute_deprecated
//void      av_format_set_metadata_header_padding(AVFormatContext *s, int c);
func (s *AVFormatContext) AvFormatSetMetadataHeaderPadding() (c ffcommon.FInt) {
	return
}

//attribute_deprecated
//void *    av_format_get_opaque(const AVFormatContext *s);
func (s *AVFormatContext) AvFormatGetOpaque() (c ffcommon.FVoidP) {
	return
}

//attribute_deprecated
//void      av_format_set_opaque(AVFormatContext *s, void *opaque);
func (s *AVFormatContext) AvFormatSetOpaque(opaque ffcommon.FVoidP) {
	return
}

type callback22 struct {
}

//attribute_deprecated
//av_format_control_message av_format_get_control_message_cb(const AVFormatContext *s);
func (s *AVFormatContext) AvFormatGetControlMessageCb() (res callback22) {
	return
}

//attribute_deprecated
//void      av_format_set_control_message_cb(AVFormatContext *s, av_format_control_message callback);
func (s *AVFormatContext) AvFormatSetControlMessageCb() (res callback22) {
	return
}

//#if FF_API_OLD_OPEN_CALLBACKS
//attribute_deprecated AVOpenCallback av_format_get_open_cb(const AVFormatContext *s);
func (s *AVFormatContext) AvFormatGetOpenCb() (res callback22) {
	return
}

//attribute_deprecated void av_format_set_open_cb(AVFormatContext *s, AVOpenCallback callback);
func (s *AVFormatContext) AvFormatSetOpenCb(callback *callback22) {
	return
}

//#endif
//#endif

/**
 * This function will cause global side data to be injected in the next packet
 * of each stream as well as after any subsequent seek.
 */
//void av_format_inject_global_side_data(AVFormatContext *s);
func (s *AVFormatContext) AvFormatInjectGlobalSideData() {
	return
}

/**
 * Returns the method used to set ctx->duration.
 *
 * @return AVFMT_DURATION_FROM_PTS, AVFMT_DURATION_FROM_STREAM, or AVFMT_DURATION_FROM_BITRATE.
 */
//enum AVDurationEstimationMethod av_fmt_ctx_get_duration_estimation_method(const AVFormatContext* ctx);
func (ctx *AVFormatContext) AvFmtCtxGetDurationEstimationMethod() (res ffconstant.AVDurationEstimationMethod) {
	return
}

/**
 * @defgroup lavf_core Core functions
 * @ingroup libavf
 *
 * Functions for querying libavformat capabilities, allocating core structures,
 * etc.
 * @{
 */

/**
 * Return the LIBAVFORMAT_VERSION_INT ffconstant.
 */
//unsigned avformat_version(void);
func AvformatVersion() (res ffcommon.FUnsigned) {
	return
}

/**
 * Return the libavformat build-time configuration.
 */
//const char *avformat_configuration(void);
func AvformatConfiguration() (res ffcommon.FConstCharP) {
	return
}

/**
 * Return the libavformat license.
 */
//const char *avformat_license(void);
func AvformatLicense() (res ffcommon.FConstCharP) {
	return
}

//#if FF_API_NEXT
/**
 * Initialize libavformat and register all the muxers, demuxers and
 * protocols. If you do not call this function, then you can select
 * exactly which formats you want to support.
 *
 * @see av_register_input_format()
 * @see av_register_output_format()
 */
//attribute_deprecated
//void av_register_all(void);
func AvRegisterAll() {
	return
}

//attribute_deprecated
//void av_register_input_format(AVInputFormat *format);
func (format *AVInputFormat) AvRegisterInputFormat() {
	return
}

//attribute_deprecated
//void av_register_output_format(AVOutputFormat *format);
func (format *AVOutputFormat) AvRegisterOutputFormat() {
	return
}

//#endif

/**
 * Do global initialization of network libraries. This is optional,
 * and not recommended anymore.
 *
 * This functions only exists to work around thread-safety issues
 * with older GnuTLS or OpenSSL libraries. If libavformat is linked
 * to newer versions of those libraries, or if you do not use them,
 * calling this function is unnecessary. Otherwise, you need to call
 * this function before any other threads using them are started.
 *
 * This function will be deprecated once support for older GnuTLS and
 * OpenSSL libraries is removed, and this function has no purpose
 * anymore.
 */
//int avformat_network_init(void);
func AvformatNetworkInit() (res ffcommon.FInt) {
	return
}

/**
 * Undo the initialization done by avformat_network_init. Call it only
 * once for each time you called avformat_network_init.
 */
//int avformat_network_deinit(void);
func AvformatNetworkDeinit() (res ffcommon.FInt) {
	return
}

//#if FF_API_NEXT
/**
 * If f is NULL, returns the first registered input format,
 * if f is non-NULL, returns the next registered input format after f
 * or NULL if f is the last one.
 */
//attribute_deprecated
//AVInputFormat  *av_iformat_next(const AVInputFormat  *f);
func (f *AVInputFormat) AvIformatNext() (res *AVInputFormat) {
	return
}

/**
 * If f is NULL, returns the first registered output format,
 * if f is non-NULL, returns the next registered output format after f
 * or NULL if f is the last one.
 */
//attribute_deprecated
//AVOutputFormat *av_oformat_next(const AVOutputFormat *f);
func (f *AVOutputFormat) AvOformatNext() (res *AVOutputFormat) {
	return
}

//#endif

/**
 * Iterate over all registered muxers.
 *
 * @param opaque a pointer where libavformat will store the iteration state. Must
 *               point to NULL to start the iteration.
 *
 * @return the next registered muxer or NULL when the iteration is
 *         finished
 */
//const AVOutputFormat *av_muxer_iterate(void **opaque);
func AvMuxerIterate(opaque *ffcommon.FVoidP) (res *AVOutputFormat) {
	return
}

/**
 * Iterate over all registered demuxers.
 *
 * @param opaque a pointer where libavformat will store the iteration state. Must
 *               point to NULL to start the iteration.
 *
 * @return the next registered demuxer or NULL when the iteration is
 *         finished
 */
//const AVInputFormat *av_demuxer_iterate(void **opaque);
func AvDemuxerIterate(opaque *ffcommon.FVoidP) (res *AVInputFormat) {
	return
}

/**
 * Allocate an AVFormatContext.
 * avformat_free_context() can be used to free the context and everything
 * allocated by the framework within it.
 */
//AVFormatContext *avformat_alloc_context(void);
func AvformatAllocContext() (res *AVFormatContext) {
	return
}

/**
 * Free an AVFormatContext and all its streams.
 * @param s context to free
 */
//void avformat_free_context(AVFormatContext *s);
func (s *AVFormatContext) AvformatFreeContext() {
	return
}

/**
 * Get the AVClass for AVFormatContext. It can be used in combination with
 * AV_OPT_SEARCH_FAKE_OBJ for examining options.
 *
 * @see av_opt_find().
 */
//const AVClass *avformat_get_class(void);
func AvformatGetClass() (res *libavutil.AVClass) {
	return
}

/**
 * Add a new stream to a media file.
 *
 * When demuxing, it is called by the demuxer in read_header(). If the
 * flag AVFMTCTX_NOHEADER is set in s.ctx_flags, then it may also
 * be called in read_packet().
 *
 * When muxing, should be called by the user before avformat_write_header().
 *
 * User is required to call avcodec_close() and avformat_free_context() to
 * clean up the allocation by avformat_new_stream().
 *
 * @param s media file handle
 * @param c If non-NULL, the AVCodecContext corresponding to the new stream
 * will be initialized to use this codec. This is needed for e.g. codec-specific
 * defaults to be set, so codec should be provided if it is known.
 *
 * @return newly created stream or NULL on error.
 */
//AVStream *avformat_new_stream(AVFormatContext *s, const AVCodec *c);
func (s *AVFormatContext) AvformatNewStream(c *AVCodec) (ans *AVStream) {
	return
}

/**
 * Wrap an existing array as stream side data.
 *
 * @param st stream
 * @param type side information type
 * @param data the side data array. It must be allocated with the av_malloc()
 *             family of functions. The ownership of the data is transferred to
 *             st.
 * @param size side information size
 * @return zero on success, a negative AVERROR code on failure. On failure,
 *         the stream is unchanged and the data remains owned by the caller.
 */
//int av_stream_add_side_data(AVStream *st, enum AVPacketSideDataType type,
//uint8_t *data, size_t size);
func (st *AVStream) AvStreamAddSideData(type0 ffconstant.AVPacketSideDataType,
	data *ffcommon.FUint8T, size ffcommon.FSizeT) (ans ffcommon.FInt) {
	return
}

/**
 * Allocate new information from stream.
 *
 * @param stream stream
 * @param type desired side information type
 * @param size side information size
 * @return pointer to fresh allocated data or NULL otherwise
 */
//uint8_t *av_stream_new_side_data(AVStream *stream,
//#if FF_API_BUFFER_SIZE_T
//enum AVPacketSideDataType type, int size);
//#else
//enum AVPacketSideDataType type, size_t size);
//#endif
func (stream *AVStream) AvStreamNewSideData(type0 ffconstant.AVPacketSideDataType, size ffcommon.FSizeT) (ans ffcommon.FUint8T) {
	return
}

/**
 * Get side information from stream.
 *
 * @param stream stream
 * @param type desired side information type
 * @param size If supplied, *size will be set to the size of the side data
 *             or to zero if the desired side data is not present.
 * @return pointer to data if present or NULL otherwise
 */
//uint8_t *av_stream_get_side_data(const AVStream *stream,
//#if FF_API_BUFFER_SIZE_T
//enum AVPacketSideDataType type, int *size);
//#else
//enum AVPacketSideDataType type, size_t *size);
//#endif
func (stream *AVStream) AvStreamGetSideData(type0 ffconstant.AVPacketSideDataType, size ffcommon.FSizeT) (ans ffcommon.FUint8T) {
	return
}

//AVProgram *av_new_program(AVFormatContext *s, int id);
func (s *AVFormatContext) AvNewProgram(id ffcommon.FInt) (ans *AVProgram) {
	return
}

/**
 * @}
 */

/**
 * Allocate an AVFormatContext for an output format.
 * avformat_free_context() can be used to free the context and
 * everything allocated by the framework within it.
 *
 * @param *ctx is set to the created format context, or to NULL in
 * case of failure
 * @param oformat format to use for allocating the context, if NULL
 * format_name and filename are used instead
 * @param format_name the name of output format to use for allocating the
 * context, if NULL filename is used instead
 * @param filename the name of the filename to use for allocating the
 * context, may be NULL
 * @return >= 0 in case of success, a negative AVERROR code in case of
 * failure
 */
//int avformat_alloc_output_context2(AVFormatContext **ctx, ff_const59 AVOutputFormat *oformat,
//const char *format_name, const char *filename);

func AvformatAllocOutputContext2(ctx **AVFormatContext, oformat *AVOutputFormat,
	format_name ffcommon.FConstCharP, filename ffcommon.FConstCharP) (ans ffcommon.FInt) {
	return
}

/**
 * @addtogroup lavf_decoding
 * @{
 */

/**
 * Find AVInputFormat based on the short name of the input format.
 */
//ff_const59 AVInputFormat *av_find_input_format(const char *short_name);
func AvFindInputFormat(ctx **AVFormatContext, short_name ffcommon.FConstCharP) (ans *AVInputFormat) {
	return
}

/**
 * Guess the file format.
 *
 * @param pd        data to be probed
 * @param is_opened Whether the file is already opened; determines whether
 *                  demuxers with or without AVFMT_NOFILE are probed.
 */
//ff_const59 AVInputFormat *av_probe_input_format(ff_const59 AVProbeData *pd, int is_opened);
func AvProbeInputFormat(pd *AVProbeData, is_opened ffcommon.FInt) (ans *AVInputFormat) {
	return
}

/**
 * Guess the file format.
 *
 * @param pd        data to be probed
 * @param is_opened Whether the file is already opened; determines whether
 *                  demuxers with or without AVFMT_NOFILE are probed.
 * @param score_max A probe score larger that this is required to accept a
 *                  detection, the variable is set to the actual detection
 *                  score afterwards.
 *                  If the score is <= AVPROBE_SCORE_MAX / 4 it is recommended
 *                  to retry with a larger probe buffer.
 */
//ff_const59 AVInputFormat *av_probe_input_format2(ff_const59 AVProbeData *pd, int is_opened, int *score_max);
func AvProbeInputFormat2(pd *AVProbeData, is_opened ffcommon.FInt, score_max *ffcommon.FInt) (ans *AVInputFormat) {
	return
}

/**
 * Guess the file format.
 *
 * @param is_opened Whether the file is already opened; determines whether
 *                  demuxers with or without AVFMT_NOFILE are probed.
 * @param score_ret The score of the best detection.
 */
//ff_const59 AVInputFormat *av_probe_input_format3(ff_const59 AVProbeData *pd, int is_opened, int *score_ret);
func AvProbeInputFormat3(pd *AVProbeData, is_opened ffcommon.FInt, score_ret *ffcommon.FInt) (ans *AVInputFormat) {
	return
}

/**
 * Probe a bytestream to determine the input format. Each time a probe returns
 * with a score that is too low, the probe buffer size is increased and another
 * attempt is made. When the maximum probe size is reached, the input format
 * with the highest score is returned.
 *
 * @param pb the bytestream to probe
 * @param fmt the input format is put here
 * @param url the url of the stream
 * @param logctx the log context
 * @param offset the offset within the bytestream to probe from
 * @param max_probe_size the maximum probe buffer size (zero for default)
 * @return the score in case of success, a negative value corresponding to an
 *         the maximal score is AVPROBE_SCORE_MAX
 * AVERROR code otherwise
 */
//int av_probe_input_buffer2(AVIOContext *pb, ff_const59 AVInputFormat **fmt,
//const char *url, void *logctx,
//unsigned int offset, unsigned int max_probe_size);
func (pb *AVIOContext) AvProbeInputBuffer2(fmt0 **AVInputFormat,
	url ffcommon.FConstCharP, logctx ffcommon.FVoidP,
	offset ffcommon.FUnsignedInt, max_probe_size ffcommon.FUnsignedInt) (ans ffcommon.FInt) {
	return
}

/**
 * Like av_probe_input_buffer2() but returns 0 on success
 */
//int av_probe_input_buffer(AVIOContext *pb, ff_const59 AVInputFormat **fmt,
//const char *url, void *logctx,
//unsigned int offset, unsigned int max_probe_size);
func (pb *AVIOContext) AvProbeInputBuffer(fmt0 **AVInputFormat,
	url ffcommon.FConstCharP, logctx ffcommon.FVoidP,
	offset ffcommon.FUnsignedInt, max_probe_size ffcommon.FUnsignedInt) (ans ffcommon.FInt) {
	return
}

/**
 * Open an input stream and read the header. The codecs are not opened.
 * The stream must be closed with avformat_close_input().
 *
 * @param ps Pointer to user-supplied AVFormatContext (allocated by avformat_alloc_context).
 *           May be a pointer to NULL, in which case an AVFormatContext is allocated by this
 *           function and written into ps.
 *           Note that a user-supplied AVFormatContext will be freed on failure.
 * @param url URL of the stream to open.
 * @param fmt If non-NULL, this parameter forces a specific input format.
 *            Otherwise the format is autodetected.
 * @param options  A dictionary filled with AVFormatContext and demuxer-private options.
 *                 On return this parameter will be destroyed and replaced with a dict containing
 *                 options that were not found. May be NULL.
 *
 * @return 0 on success, a negative AVERROR on failure.
 *
 * @note If you want to use custom IO, preallocate the format context and set its pb field.
 */
//int avformat_open_input(AVFormatContext **ps, const char *url, ff_const59 AVInputFormat *fmt, AVDictionary **options);
func AvformatOpenInput(ps **AVFormatContext, url ffcommon.FConstCharP, fmt0 *AVInputFormat, options **AVDictionary) (ans ffcommon.FInt) {
	return
}

type AVDictionary struct {
}

//#if FF_API_DEMUXER_OPEN
/**
 * @deprecated Use an AVDictionary to pass options to a demuxer.
 */
//attribute_deprecated
//int av_demuxer_open(AVFormatContext *ic);
func (ic *AVFormatContext) AvDemuxerOpen() (ans ffcommon.FInt) {
	return
}

//#endif

/**
 * Read packets of a media file to get stream information. This
 * is useful for file formats with no headers such as MPEG. This
 * function also computes the real framerate in case of MPEG-2 repeat
 * frame mode.
 * The logical file position is not changed by this function;
 * examined packets may be buffered for later processing.
 *
 * @param ic media file handle
 * @param options  If non-NULL, an ic.nb_streams long array of pointers to
 *                 dictionaries, where i-th member contains options for
 *                 codec corresponding to i-th stream.
 *                 On return each dictionary will be filled with options that were not found.
 * @return >=0 if OK, AVERROR_xxx on error
 *
 * @note this function isn't guaranteed to open all the codecs, so
 *       options being non-empty at return is a perfectly normal behavior.
 *
 * @todo Let the user decide somehow what information is needed so that
 *       we do not waste time getting stuff the user does not need.
 */
//int avformat_find_stream_info(AVFormatContext *ic, AVDictionary **options);
func (ic *AVFormatContext) AvformatFindStreamInfo(options **AVDictionary) (ans ffcommon.FInt) {
	return
}

/**
 * Find the programs which belong to a given stream.
 *
 * @param ic    media file handle
 * @param last  the last found program, the search will start after this
 *              program, or from the beginning if it is NULL
 * @param s     stream index
 * @return the next program which belongs to s, NULL if no program is found or
 *         the last program is not among the programs of ic.
 */
//AVProgram *av_find_program_from_stream(AVFormatContext *ic, AVProgram *last, int s);
func (ic *AVFormatContext) AvFindProgramFromStream(last *AVProgram, s ffcommon.FInt) (ans *AVProgram) {
	return
}

//void av_program_add_stream_index(AVFormatContext *ac, int progid, unsigned int idx);
func (ac *AVFormatContext) AvProgramAddStreamIndex(progid ffcommon.FInt, idx ffcommon.FUnsignedInt) {
	return
}

/**
 * Find the "best" stream in the file.
 * The best stream is determined according to various heuristics as the most
 * likely to be what the user expects.
 * If the decoder parameter is non-NULL, av_find_best_stream will find the
 * default decoder for the stream's codec; streams for which no decoder can
 * be found are ignored.
 *
 * @param ic                media file handle
 * @param type              stream type: video, audio, subtitles, etc.
 * @param wanted_stream_nb  user-requested stream number,
 *                          or -1 for automatic selection
 * @param related_stream    try to find a stream related (eg. in the same
 *                          program) to this one, or -1 if none
 * @param decoder_ret       if non-NULL, returns the decoder for the
 *                          selected stream
 * @param flags             flags; none are currently defined
 * @return  the non-negative stream number in case of success,
 *          AVERROR_STREAM_NOT_FOUND if no stream with the requested type
 *          could be found,
 *          AVERROR_DECODER_NOT_FOUND if streams were found but no decoder
 * @note  If av_find_best_stream returns successfully and decoder_ret is not
 *        NULL, then *decoder_ret is guaranteed to be set to a valid AVCodec.
 */
//int av_find_best_stream(AVFormatContext *ic,
//enum AVMediaType type,
//int wanted_stream_nb,
//int related_stream,
//AVCodec **decoder_ret,
//int flags);
func (ic *AVFormatContext) av_find_best_stream(
	type0 ffconstant.AVMediaType,
	wanted_stream_nb ffcommon.FInt,
	related_stream ffcommon.FInt,
	decoder_ret **AVCodec,
	flags ffcommon.FInt) (ans ffcommon.FInt) {
	return
}

/**
 * Return the next frame of a stream.
 * This function returns what is stored in the file, and does not validate
 * that what is there are valid frames for the decoder. It will split what is
 * stored in the file into frames and return one for each call. It will not
 * omit invalid data between valid frames so as to give the decoder the maximum
 * information possible for decoding.
 *
 * On success, the returned packet is reference-counted (pkt->buf is set) and
 * valid indefinitely. The packet must be freed with av_packet_unref() when
 * it is no longer needed. For video, the packet contains exactly one frame.
 * For audio, it contains an integer number of frames if each frame has
 * a known fixed size (e.g. PCM or ADPCM data). If the audio frames have
 * a variable size (e.g. MPEG audio), then it contains one frame.
 *
 * pkt->pts, pkt->dts and pkt->duration are always set to correct
 * values in AVStream.time_base units (and guessed if the format cannot
 * provide them). pkt->pts can be AV_NOPTS_VALUE if the video format
 * has B-frames, so it is better to rely on pkt->dts if you do not
 * decompress the payload.
 *
 * @return 0 if OK, < 0 on error or end of file. On error, pkt will be blank
 *         (as if it came from av_packet_alloc()).
 *
 * @note pkt will be initialized, so it may be uninitialized, but it must not
 *       contain data that needs to be freed.
 */
//int av_read_frame(AVFormatContext *s, AVPacket *pkt);
func (s *AVFormatContext) AvReadFrame(pkt *AVPacket) (res ffcommon.FInt) {
	return
}

/**
 * Seek to the keyframe at timestamp.
 * 'timestamp' in 'stream_index'.
 *
 * @param s media file handle
 * @param stream_index If stream_index is (-1), a default
 * stream is selected, and timestamp is automatically converted
 * from AV_TIME_BASE units to the stream specific time_base.
 * @param timestamp Timestamp in AVStream.time_base units
 *        or, if no stream is specified, in AV_TIME_BASE units.
 * @param flags flags which select direction and seeking mode
 * @return >= 0 on success
 */
//int av_seek_frame(AVFormatContext *s, int stream_index, int64_t timestamp,
//int flags);
func (s *AVFormatContext) av_seek_frame(stream_index ffcommon.FInt, timestamp ffcommon.FInt64T,
	flags ffcommon.FInt) (res ffcommon.FInt) {
	return
}

/**
 * Seek to timestamp ts.
 * Seeking will be done so that the point from which all active streams
 * can be presented successfully will be closest to ts and within min/max_ts.
 * Active streams are all streams that have AVStream.discard < AVDISCARD_ALL.
 *
 * If flags contain AVSEEK_FLAG_BYTE, then all timestamps are in bytes and
 * are the file position (this may not be supported by all demuxers).
 * If flags contain AVSEEK_FLAG_FRAME, then all timestamps are in frames
 * in the stream with stream_index (this may not be supported by all demuxers).
 * Otherwise all timestamps are in units of the stream selected by stream_index
 * or if stream_index is -1, in AV_TIME_BASE units.
 * If flags contain AVSEEK_FLAG_ANY, then non-keyframes are treated as
 * keyframes (this may not be supported by all demuxers).
 * If flags contain AVSEEK_FLAG_BACKWARD, it is ignored.
 *
 * @param s media file handle
 * @param stream_index index of the stream which is used as time base reference
 * @param min_ts smallest acceptable timestamp
 * @param ts target timestamp
 * @param max_ts largest acceptable timestamp
 * @param flags flags
 * @return >=0 on success, error code otherwise
 *
 * @note This is part of the new seek API which is still under construction.
 */
//int avformat_seek_file(AVFormatContext *s, int stream_index, int64_t min_ts, int64_t ts, int64_t max_ts, int flags);
func (s *AVFormatContext) AvformatSeekFile(stream_index ffcommon.FInt, min_ts ffcommon.FUint64T, ts ffcommon.FUint64T, max_ts ffcommon.FUint64T, flags ffcommon.FInt) (res ffcommon.FInt) {
	return
}

/**
 * Discard all internally buffered data. This can be useful when dealing with
 * discontinuities in the byte stream. Generally works only with formats that
 * can resync. This includes headerless formats like MPEG-TS/TS but should also
 * work with NUT, Ogg and in a limited way AVI for example.
 *
 * The set of streams, the detected duration, stream parameters and codecs do
 * not change when calling this function. If you want a complete reset, it's
 * better to open a new AVFormatContext.
 *
 * This does not flush the AVIOContext (s->pb). If necessary, call
 * avio_flush(s->pb) before calling this function.
 *
 * @param s media file handle
 * @return >=0 on success, error code otherwise
 */
//int avformat_flush(AVFormatContext *s);
func (s *AVFormatContext) AvformatFlush() (res ffcommon.FInt) {
	return
}

/**
 * Start playing a network-based stream (e.g. RTSP stream) at the
 * current position.
 */
//int av_read_play(AVFormatContext *s);
func (s *AVFormatContext) AvReadPlay() (res ffcommon.FInt) {
	return
}

/**
 * Pause a network-based stream (e.g. RTSP stream).
 *
 * Use av_read_play() to resume it.
 */
//int av_read_pause(AVFormatContext *s);
func (s *AVFormatContext) AvReadPause() (res ffcommon.FInt) {
	return
}

/**
 * Close an opened input AVFormatContext. Free it and all its contents
 * and set *s to NULL.
 */
//void avformat_close_input(AVFormatContext **s);
func AvformatCloseInput(s **AVFormatContext) {
	return
}

/**
 * Allocate the stream private data and write the stream header to
 * an output media file.
 *
 * @param s Media file handle, must be allocated with avformat_alloc_context().
 *          Its oformat field must be set to the desired output format;
 *          Its pb field must be set to an already opened AVIOContext.
 * @param options  An AVDictionary filled with AVFormatContext and muxer-private options.
 *                 On return this parameter will be destroyed and replaced with a dict containing
 *                 options that were not found. May be NULL.
 *
 * @return AVSTREAM_INIT_IN_WRITE_HEADER on success if the codec had not already been fully initialized in avformat_init,
 *         AVSTREAM_INIT_IN_INIT_OUTPUT  on success if the codec had already been fully initialized in avformat_init,
 *         negative AVERROR on failure.
 *
 * @see av_opt_find, av_dict_set, avio_open, av_oformat_next, avformat_init_output.
 */
//av_warn_unused_result
//int avformat_write_header(AVFormatContext *s, AVDictionary **options);
func (s *AVFormatContext) AvformatWriteHeader(options **AVDictionary) (res ffcommon.FInt) {
	return
}

/**
 * Allocate the stream private data and initialize the codec, but do not write the header.
 * May optionally be used before avformat_write_header to initialize stream parameters
 * before actually writing the header.
 * If using this function, do not pass the same options to avformat_write_header.
 *
 * @param s Media file handle, must be allocated with avformat_alloc_context().
 *          Its oformat field must be set to the desired output format;
 *          Its pb field must be set to an already opened AVIOContext.
 * @param options  An AVDictionary filled with AVFormatContext and muxer-private options.
 *                 On return this parameter will be destroyed and replaced with a dict containing
 *                 options that were not found. May be NULL.
 *
 * @return AVSTREAM_INIT_IN_WRITE_HEADER on success if the codec requires avformat_write_header to fully initialize,
 *         AVSTREAM_INIT_IN_INIT_OUTPUT  on success if the codec has been fully initialized,
 *         negative AVERROR on failure.
 *
 * @see av_opt_find, av_dict_set, avio_open, av_oformat_next, avformat_write_header.
 */
//av_warn_unused_result
//int avformat_init_output(AVFormatContext *s, AVDictionary **options);
func (s *AVFormatContext) AvformatInitOutput(options **AVDictionary) (res ffcommon.FInt) {
	return
}

/**
 * Write a packet to an output media file.
 *
 * This function passes the packet directly to the muxer, without any buffering
 * or reordering. The caller is responsible for correctly interleaving the
 * packets if the format requires it. Callers that want libavformat to handle
 * the interleaving should call av_interleaved_write_frame() instead of this
 * function.
 *
 * @param s media file handle
 * @param pkt The packet containing the data to be written. Note that unlike
 *            av_interleaved_write_frame(), this function does not take
 *            ownership of the packet passed to it (though some muxers may make
 *            an internal reference to the input packet).
 *            <br>
 *            This parameter can be NULL (at any time, not just at the end), in
 *            order to immediately flush data buffered within the muxer, for
 *            muxers that buffer up data internally before writing it to the
 *            output.
 *            <br>
 *            Packet's @ref AVPacket.stream_index "stream_index" field must be
 *            set to the index of the corresponding stream in @ref
 *            AVFormatContext.streams "s->streams".
 *            <br>
 *            The timestamps (@ref AVPacket.pts "pts", @ref AVPacket.dts "dts")
 *            must be set to correct values in the stream's timebase (unless the
 *            output format is flagged with the AVFMT_NOTIMESTAMPS flag, then
 *            they can be set to AV_NOPTS_VALUE).
 *            The dts for subsequent packets passed to this function must be strictly
 *            increasing when compared in their respective timebases (unless the
 *            output format is flagged with the AVFMT_TS_NONSTRICT, then they
 *            merely have to be nondecreasing).  @ref AVPacket.duration
 *            "duration") should also be set if known.
 * @return < 0 on error, = 0 if OK, 1 if flushed and there is no more data to flush
 *
 * @see av_interleaved_write_frame()
 */
//int av_write_frame(AVFormatContext *s, AVPacket *pkt);
func (s *AVFormatContext) AvWriteFrame(pkt *AVPacket) (res ffcommon.FInt) {
	return
}

/**
 * Write a packet to an output media file ensuring correct interleaving.
 *
 * This function will buffer the packets internally as needed to make sure the
 * packets in the output file are properly interleaved in the order of
 * increasing dts. Callers doing their own interleaving should call
 * av_write_frame() instead of this function.
 *
 * Using this function instead of av_write_frame() can give muxers advance
 * knowledge of future packets, improving e.g. the behaviour of the mp4
 * muxer for VFR content in fragmenting mode.
 *
 * @param s media file handle
 * @param pkt The packet containing the data to be written.
 *            <br>
 *            If the packet is reference-counted, this function will take
 *            ownership of this reference and unreference it later when it sees
 *            fit.
 *            The caller must not access the data through this reference after
 *            this function returns. If the packet is not reference-counted,
 *            libavformat will make a copy.
 *            <br>
 *            This parameter can be NULL (at any time, not just at the end), to
 *            flush the interleaving queues.
 *            <br>
 *            Packet's @ref AVPacket.stream_index "stream_index" field must be
 *            set to the index of the corresponding stream in @ref
 *            AVFormatContext.streams "s->streams".
 *            <br>
 *            The timestamps (@ref AVPacket.pts "pts", @ref AVPacket.dts "dts")
 *            must be set to correct values in the stream's timebase (unless the
 *            output format is flagged with the AVFMT_NOTIMESTAMPS flag, then
 *            they can be set to AV_NOPTS_VALUE).
 *            The dts for subsequent packets in one stream must be strictly
 *            increasing (unless the output format is flagged with the
 *            AVFMT_TS_NONSTRICT, then they merely have to be nondecreasing).
 *            @ref AVPacket.duration "duration") should also be set if known.
 *
 * @return 0 on success, a negative AVERROR on error. Libavformat will always
 *         take care of freeing the packet, even if this function fails.
 *
 * @see av_write_frame(), AVFormatContext.max_interleave_delta
 */
//int av_interleaved_write_frame(AVFormatContext *s, AVPacket *pkt);
func (s *AVFormatContext) AvInterleavedWriteFrame(pkt *AVPacket) (res ffcommon.FInt) {
	return
}

/**
 * Write an uncoded frame to an output media file.
 *
 * The frame must be correctly interleaved according to the container
 * specification; if not, av_interleaved_write_uncoded_frame() must be used.
 *
 * See av_interleaved_write_uncoded_frame() for details.
 */
//int av_write_uncoded_frame(AVFormatContext *s, int stream_index,
//AVFrame *frame);
func (s *AVFormatContext) AvWriteUncodedFrame(stream_index ffcommon.FInt, frame *libavutil.AVFrame) (res ffcommon.FInt) {
	return
}

/**
 * Write an uncoded frame to an output media file.
 *
 * If the muxer supports it, this function makes it possible to write an AVFrame
 * structure directly, without encoding it into a packet.
 * It is mostly useful for devices and similar special muxers that use raw
 * video or PCM data and will not serialize it into a byte stream.
 *
 * To test whether it is possible to use it with a given muxer and stream,
 * use av_write_uncoded_frame_query().
 *
 * The caller gives up ownership of the frame and must not access it
 * afterwards.
 *
 * @return  >=0 for success, a negative code on error
 */
//int av_interleaved_write_uncoded_frame(AVFormatContext *s, int stream_index,
//AVFrame *frame);
func (s *AVFormatContext) AvInterleavedWriteUncodedFrame(stream_index ffcommon.FInt, frame *libavutil.AVFrame) (res ffcommon.FInt) {
	return
}

/**
 * Test whether a muxer supports uncoded frame.
 *
 * @return  >=0 if an uncoded frame can be written to that muxer and stream,
 *          <0 if not
 */
//int av_write_uncoded_frame_query(AVFormatContext *s, int stream_index);
func (s *AVFormatContext) AvWriteUncodedFrameQuery(stream_index ffcommon.FInt) (res ffcommon.FInt) {
	return
}

/**
 * Write the stream trailer to an output media file and free the
 * file private data.
 *
 * May only be called after a successful call to avformat_write_header.
 *
 * @param s media file handle
 * @return 0 if OK, AVERROR_xxx on error
 */
//int av_write_trailer(AVFormatContext *s);
func (s *AVFormatContext) AvWriteTrailer() (res ffcommon.FInt) {
	return
}

/**
 * Return the output format in the list of registered output formats
 * which best matches the provided parameters, or return NULL if
 * there is no match.
 *
 * @param short_name if non-NULL checks if short_name matches with the
 * names of the registered formats
 * @param filename if non-NULL checks if filename terminates with the
 * extensions of the registered formats
 * @param mime_type if non-NULL checks if mime_type matches with the
 * MIME type of the registered formats
 */
//ff_const59 AVOutputFormat *av_guess_format(const char *short_name,
//const char *filename,
//const char *mime_type);
func AvGuessFormat(short_name ffcommon.FConstCharP,
	filename ffcommon.FConstCharP,
	mime_type ffcommon.FConstCharP) (ans *AVOutputFormat) {
	return
}

/**
 * Guess the codec ID based upon muxer and filename.
 */
//enum AVCodecID av_guess_codec(ff_const59 AVOutputFormat *fmt, const char *short_name,
//const char *filename, const char *mime_type,
//enum AVMediaType type);

func AvGuessCodec(fmt0 *AVOutputFormat, short_name ffcommon.FConstCharP,
	filename ffcommon.FConstCharP, mime_type ffcommon.FConstCharP,
	type0 ffconstant.AVMediaType) (ans ffconstant.AVCodecID) {
	return
}

/**
 * Get timing information for the data currently output.
 * The exact meaning of "currently output" depends on the format.
 * It is mostly relevant for devices that have an internal buffer and/or
 * work in real time.
 * @param s          media file handle
 * @param stream     stream in the media file
 * @param[out] dts   DTS of the last packet output for the stream, in stream
 *                   time_base units
 * @param[out] wall  absolute time when that packet whas output,
 *                   in microsecond
 * @return  0 if OK, AVERROR(ENOSYS) if the format does not support it
 * Note: some formats or devices may not allow to measure dts and wall
 * atomically.
 */
//int av_get_output_timestamp(struct AVFormatContext *s, int stream,
//int64_t *dts, int64_t *wall);
func (s *AVFormatContext) AvGetOutputTimestamp(stream ffcommon.FInt,
	dts *ffcommon.FInt64T, wall *ffcommon.FInt64T) (ans ffcommon.FInt) {
	return
}

/**
 * @}
 */

/**
 * @defgroup lavf_misc Utility functions
 * @ingroup libavf
 * @{
 *
 * Miscellaneous utility functions related to both muxing and demuxing
 * (or neither).
 */

/**
 * Send a nice hexadecimal dump of a buffer to the specified file stream.
 *
 * @param f The file stream pointer where the dump should be sent to.
 * @param buf buffer
 * @param size buffer size
 *
 * @see av_hex_dump_log, av_pkt_dump2, av_pkt_dump_log2
 */
//void av_hex_dump(FILE *f, const uint8_t *buf, int size);
type File struct {
}

func AvHexDump(f *File, buf ffcommon.FUnsigned, size ffcommon.FInt) {
	return
}

/**
 * Send a nice hexadecimal dump of a buffer to the log.
 *
 * @param avcl A pointer to an arbitrary struct of which the first field is a
 * pointer to an AVClass struct.
 * @param level The importance level of the message, lower values signifying
 * higher importance.
 * @param buf buffer
 * @param size buffer size
 *
 * @see av_hex_dump, av_pkt_dump2, av_pkt_dump_log2
 */
//void av_hex_dump_log(void *avcl, int level, const uint8_t *buf, int size);
func AvHexDumpLog(avcl ffcommon.FVoidP, level ffcommon.FInt, buf *ffcommon.FUint8T, size ffcommon.FInt) {
	return
}

/**
 * Send a nice dump of a packet to the specified file stream.
 *
 * @param f The file stream pointer where the dump should be sent to.
 * @param pkt packet to dump
 * @param dump_payload True if the payload must be displayed, too.
 * @param st AVStream that the packet belongs to
 */
//void av_pkt_dump2(FILE *f, const AVPacket *pkt, int dump_payload, const AVStream *st);
func AvPktDump2(f *File, pkt *AVPacket, dump_payload ffcommon.FInt, st *AVStream) {
	return
}

/**
 * Send a nice dump of a packet to the log.
 *
 * @param avcl A pointer to an arbitrary struct of which the first field is a
 * pointer to an AVClass struct.
 * @param level The importance level of the message, lower values signifying
 * higher importance.
 * @param pkt packet to dump
 * @param dump_payload True if the payload must be displayed, too.
 * @param st AVStream that the packet belongs to
 */
//void av_pkt_dump_log2(void *avcl, int level, const AVPacket *pkt, int dump_payload,
//const AVStream *st);
func AvPktDumpLog2(avcl ffcommon.FVoidP, level ffcommon.FInt, pkt *AVPacket, dump_payload ffcommon.FInt,
	st *AVStream) {
	return
}

/**
 * Get the AVCodecID for the given codec tag tag.
 * If no codec id is found returns AV_CODEC_ID_NONE.
 *
 * @param tags list of supported codec_id-codec_tag pairs, as stored
 * in AVInputFormat.codec_tag and AVOutputFormat.codec_tag
 * @param tag  codec tag to match to a codec ID
 */
//enum AVCodecID av_codec_get_id(const struct AVCodecTag * const *tags, unsigned int tag);
func AvCodecGetId(tags **AVCodecTag, tag ffcommon.FUnsignedInt) (res ffconstant.AVCodecID) {
	return
}

/**
 * Get the codec tag for the given codec id id.
 * If no codec tag is found returns 0.
 *
 * @param tags list of supported codec_id-codec_tag pairs, as stored
 * in AVInputFormat.codec_tag and AVOutputFormat.codec_tag
 * @param id   codec ID to match to a codec tag
 */
//unsigned int av_codec_get_tag(const struct AVCodecTag * const *tags, enum AVCodecID id);
func AvCodecGetTag(tags **AVCodecTag, id ffconstant.AVCodecID) (res ffcommon.FUnsignedInt) {
	return
}

/**
 * Get the codec tag for the given codec id.
 *
 * @param tags list of supported codec_id - codec_tag pairs, as stored
 * in AVInputFormat.codec_tag and AVOutputFormat.codec_tag
 * @param id codec id that should be searched for in the list
 * @param tag A pointer to the found tag
 * @return 0 if id was not found in tags, > 0 if it was found
 */
//int av_codec_get_tag2(const struct AVCodecTag * const *tags, enum AVCodecID id,
//unsigned int *tag);
func AvCodecGetTag2(tags **AVCodecTag, id ffconstant.AVCodecID, tag *ffcommon.FUnsignedInt) (res ffcommon.FInt) {
	return
}

//int av_find_default_stream_index(AVFormatContext *s);
func (s *AVFormatContext) AvFindDefaultStreamIndex() (res ffcommon.FInt) {
	return
}

/**
 * Get the index for a specific timestamp.
 *
 * @param st        stream that the timestamp belongs to
 * @param timestamp timestamp to retrieve the index for
 * @param flags if AVSEEK_FLAG_BACKWARD then the returned index will correspond
 *                 to the timestamp which is <= the requested one, if backward
 *                 is 0, then it will be >=
 *              if AVSEEK_FLAG_ANY seek to any frame, only keyframes otherwise
 * @return < 0 if no such timestamp could be found
 */
//int av_index_search_timestamp(AVStream *st, int64_t timestamp, int flags);
func (st *AVStream) AvIndexSearchTimestamp(timestamp ffcommon.FInt64T, flags ffcommon.FInt) (res ffcommon.FInt) {
	return
}

/**
 * Add an index entry into a sorted list. Update the entry if the list
 * already contains it.
 *
 * @param timestamp timestamp in the time base of the given stream
 */
//int av_add_index_entry(AVStream *st, int64_t pos, int64_t timestamp,
//int size, int distance, int flags);
func (st *AVStream) av_add_index_entry(pos ffcommon.FInt64T, timestamp ffcommon.FInt64T, size ffcommon.FInt, distance ffcommon.FInt, flags ffcommon.FInt) (res ffcommon.FInt) {
	return
}

/**
 * Split a URL string into components.
 *
 * The pointers to buffers for storing individual components may be null,
 * in order to ignore that component. Buffers for components not found are
 * set to empty strings. If the port is not found, it is set to a negative
 * value.
 *
 * @param proto the buffer for the protocol
 * @param proto_size the size of the proto buffer
 * @param authorization the buffer for the authorization
 * @param authorization_size the size of the authorization buffer
 * @param hostname the buffer for the host name
 * @param hostname_size the size of the hostname buffer
 * @param port_ptr a pointer to store the port number in
 * @param path the buffer for the path
 * @param path_size the size of the path buffer
 * @param url the URL to split
 */
//void av_url_split(char *proto,         int proto_size,
//char *authorization, int authorization_size,
//char *hostname,      int hostname_size,
//int *port_ptr,
//char *path,          int path_size,
//const char *url);
func AvUrlSplit(proto ffcommon.FCharP, proto_size ffcommon.FInt,
	authorization ffcommon.FCharP, authorization_size ffcommon.FInt,
	hostname ffcommon.FCharP, hostname_size ffcommon.FInt,
	port_ptr *ffcommon.FInt,
	path ffcommon.FCharP, path_size ffcommon.FInt,
	url ffcommon.FConstCharP) {
	return
}

/**
 * Print detailed information about the input or output format, such as
 * duration, bitrate, streams, container, programs, metadata, side data,
 * codec and time base.
 *
 * @param ic        the context to analyze
 * @param index     index of the stream to dump information about
 * @param url       the URL to print, such as source or destination file
 * @param is_output Select whether the specified context is an input(0) or output(1)
 */
//void av_dump_format(AVFormatContext *ic,
//int index,
//const char *url,
//int is_output);
func (ic *AVFormatContext) AvDumpFormat(
	index ffcommon.FInt,
	url ffcommon.FConstCharP,
	is_output ffcommon.FInt) {
	return
}

/**
 * Return in 'buf' the path with '%d' replaced by a number.
 *
 * Also handles the '%0nd' format where 'n' is the total number
 * of digits and '%%'.
 *
 * @param buf destination buffer
 * @param buf_size destination buffer size
 * @param path numbered sequence string
 * @param number frame number
 * @param flags AV_FRAME_FILENAME_FLAGS_*
 * @return 0 if OK, -1 on format error
 */
//int av_get_frame_filename2(char *buf, int buf_size,
//const char *path, int number, int flags);
func av_get_frame_filename2(buf ffcommon.FCharP, buf_size ffcommon.FInt,
	path0 ffcommon.FConstCharP, number ffcommon.FInt, flags ffcommon.FInt) (res ffcommon.FInt) {
	return
}

//int av_get_frame_filename(char *buf, int buf_size,
//const char *path, int number);
func av_get_frame_filename(buf ffcommon.FCharP, buf_size ffcommon.FInt,
	path ffcommon.FConstCharP, number ffcommon.FInt) (res ffcommon.FInt) {
	return
}

/**
 * Check whether filename actually is a numbered sequence generator.
 *
 * @param filename possible numbered sequence string
 * @return 1 if a valid numbered sequence string, 0 otherwise
 */
//int av_filename_number_test(const char *filename);
func av_filename_number_test(filename ffcommon.FConstCharP) (res ffcommon.FInt) {
	return
}

/**
 * Generate an SDP for an RTP session.
 *
 * Note, this overwrites the id values of AVStreams in the muxer contexts
 * for getting unique dynamic payload types.
 *
 * @param ac array of AVFormatContexts describing the RTP streams. If the
 *           array is composed by only one context, such context can contain
 *           multiple AVStreams (one AVStream per RTP stream). Otherwise,
 *           all the contexts in the array (an AVCodecContext per RTP stream)
 *           must contain only one AVStream.
 * @param n_files number of AVCodecContexts contained in ac
 * @param buf buffer where the SDP will be stored (must be allocated by
 *            the caller)
 * @param size the size of the buffer
 * @return 0 if OK, AVERROR_xxx on error
 */
//int av_sdp_create(AVFormatContext *ac[], int n_files, char *buf, int size);
func av_sdp_create(ac []*AVFormatContext, n_files ffcommon.FInt, buf ffcommon.FCharP, size ffcommon.FInt) (res ffcommon.FInt) {
	return
}

/**
 * Return a positive value if the given filename has one of the given
 * extensions, 0 otherwise.
 *
 * @param filename   file name to check against the given extensions
 * @param extensions a comma-separated list of filename extensions
 */
//int av_match_ext(const char *filename, const char *extensions);
func av_match_ext(filename ffcommon.FConstCharP, extensions ffcommon.FConstCharP) (res ffcommon.FInt) {
	return
}

/**
 * Test if the given container can store a codec.
 *
 * @param ofmt           container to check for compatibility
 * @param codec_id       codec to potentially store in container
 * @param std_compliance standards compliance level, one of FF_COMPLIANCE_*
 *
 * @return 1 if codec with ID codec_id can be stored in ofmt, 0 if it cannot.
 *         A negative number if this information is not available.
 */
//int avformat_query_codec(const AVOutputFormat *ofmt, enum AVCodecID codec_id,
//int std_compliance);
func (ofmt *AVOutputFormat) avformat_query_codec(codec_id ffconstant.AVCodecID, std_compliance ffcommon.FInt) (res ffcommon.FInt) {
	return
}

/**
 * @defgroup riff_fourcc RIFF FourCCs
 * @{
 * Get the tables mapping RIFF FourCCs to libavcodec AVCodecIDs. The tables are
 * meant to be passed to av_codec_get_id()/av_codec_get_tag() as in the
 * following code:
 * @code
 * uint32_t tag = MKTAG('H', '2', '6', '4');
 * const struct AVCodecTag *table[] = { avformat_get_riff_video_tags(), 0 };
 * enum AVCodecID id = av_codec_get_id(table, tag);
 * @endcode
 */
/**
 * @return the table mapping RIFF FourCCs for video to libavcodec AVCodecID.
 */
//const struct AVCodecTag *avformat_get_riff_video_tags(void);
func avformat_get_riff_video_tags() (res *AVCodecTag) {
	return
}

/**
 * @return the table mapping RIFF FourCCs for audio to AVCodecID.
 */
//const struct AVCodecTag *avformat_get_riff_audio_tags(void);
func avformat_get_riff_audio_tags() (res *AVCodecTag) {
	return
}

/**
 * @return the table mapping MOV FourCCs for video to libavcodec AVCodecID.
 */
//const struct AVCodecTag *avformat_get_mov_video_tags(void);
func avformat_get_mov_video_tags() (res *AVCodecTag) {
	return
}

/**
 * @return the table mapping MOV FourCCs for audio to AVCodecID.
 */
//const struct AVCodecTag *avformat_get_mov_audio_tags(void);
func avformat_get_mov_audio_tags() (res *AVCodecTag) {
	return
}

/**
 * @}
 */

/**
 * Guess the sample aspect ratio of a frame, based on both the stream and the
 * frame aspect ratio.
 *
 * Since the frame aspect ratio is set by the codec but the stream aspect ratio
 * is set by the demuxer, these two may not be equal. This function tries to
 * return the value that you should use if you would like to display the frame.
 *
 * Basic logic is to use the stream aspect ratio if it is set to something sane
 * otherwise use the frame aspect ratio. This way a container setting, which is
 * usually easy to modify can override the coded value in the frames.
 *
 * @param format the format context which the stream is part of
 * @param stream the stream which the frame is part of
 * @param frame the frame with the aspect ratio to be determined
 * @return the guessed (valid) sample_aspect_ratio, 0/1 if no idea
 */
//AVRational av_guess_sample_aspect_ratio(AVFormatContext *format, AVStream *stream, AVFrame *frame);
func (format *AVFormatContext) av_guess_sample_aspect_ratio(stream *AVStream, frame *libavutil.AVFrame) (res AVRational) {
	return
}

/**
 * Guess the frame rate, based on both the container and codec information.
 *
 * @param ctx the format context which the stream is part of
 * @param stream the stream which the frame is part of
 * @param frame the frame for which the frame rate should be determined, may be NULL
 * @return the guessed (valid) frame rate, 0/1 if no idea
 */
//AVRational av_guess_frame_rate(AVFormatContext *ctx, AVStream *stream, AVFrame *frame);
func (ctx *AVFormatContext) av_guess_frame_rate(stream *AVStream, frame *libavutil.AVFrame) (res AVRational) {
	return
}

/**
 * Check if the stream st contained in s is matched by the stream specifier
 * spec.
 *
 * See the "stream specifiers" chapter in the documentation for the syntax
 * of spec.
 *
 * @return  >0 if st is matched by spec;
 *          0  if st is not matched by spec;
 *          AVERROR code if spec is invalid
 *
 * @note  A stream specifier can match several streams in the format.
 */
//int avformat_match_stream_specifier(AVFormatContext *s, AVStream *st,
//const char *spec);
func (s *AVFormatContext) avformat_match_stream_specifier(st *AVStream, spec ffcommon.FConstCharP) (res AVRational) {
	return
}

//int avformat_queue_attached_pictures(AVFormatContext *s);
func (s *AVFormatContext) avformat_queue_attached_pictures() (res ffcommon.FInt) {
	return
}

//#if FF_API_OLD_BSF
/**
 * Apply a list of bitstream filters to a packet.
 *
 * @param codec AVCodecContext, usually from an AVStream
 * @param pkt the packet to apply filters to. If, on success, the returned
 *        packet has size == 0 and side_data_elems == 0, it indicates that
 *        the packet should be dropped
 * @param bsfc a NULL-terminated list of filters to apply
 * @return  >=0 on success;
 *          AVERROR code on failure
 */
//attribute_deprecated
//int av_apply_bitstream_filters(AVCodecContext *codec, AVPacket *pkt,
//AVBitStreamFilterContext *bsfc);
func (codec AVCodecContext) av_apply_bitstream_filters(pkt AVPacket,
	bsfc *AVBitStreamFilterContext) (res ffcommon.FInt) {
	return
}

//#endif

/**
 * Transfer internal timing information from one stream to another.
 *
 * This function is useful when doing stream copy.
 *
 * @param ofmt     target output format for ost
 * @param ost      output stream which needs timings copy and adjustments
 * @param ist      reference input stream to copy timings from
 * @param copy_tb  define from where the stream codec timebase needs to be imported
 */
//int avformat_transfer_internal_stream_timing_info(const AVOutputFormat *ofmt,
//AVStream *ost, const AVStream *ist,
//enum AVTimebaseSource copy_tb);

func (ofmt *AVOutputFormat) avformat_transfer_internal_stream_timing_info(
	ost *AVStream, ist *AVStream,
	copy_tb ffconstant.AVTimebaseSource) (res ffcommon.FInt) {
	return
}

/**
 * Get the internal codec timebase from a stream.
 *
 * @param st  input stream to extract the timebase from
 */
//AVRational av_stream_get_codec_timebase(const AVStream *st);
func (st *AVStream) av_stream_get_codec_timebase() (res AVRational) {
	return
}
