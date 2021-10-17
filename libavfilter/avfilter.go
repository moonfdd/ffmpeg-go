package libavfilter

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"ffmpeg-go/libavutil"
	"syscall"
	"unsafe"
)

/**
 * Return the LIBAVFILTER_VERSION_INT ffconstant.
 */
//unsigned avfilter_version(void);
func AvfilterVersion() (res ffcommon.FUint, err error) {
	t, _, _ := ffcommon.GetAvfilterDll().NewProc("avfilter_version").Call()
	res = ffcommon.FUint(t)
	return
}

/**
 * Return the libavfilter build-time configuration.
 */
//const char *avfilter_configuration(void);
func AvfilterConfiguration() (res ffcommon.FConstCharP, err error) {
	t, _, _ := ffcommon.GetAvfilterDll().NewProc("avfilter_configuration").Call()
	res = ffcommon.GoAStr(t)
	return
}

/**
 * Return the libavfilter license.
 */
//const char *avfilter_license(void);
func AvfilterLicense() (res ffcommon.FConstCharP, err error) {
	t, _, _ := ffcommon.GetAvfilterDll().NewProc("avfilter_license").Call()
	res = ffcommon.GoAStr(t)
	return
}

//type  AVFilterContext struct{}
//type  AVFilterLink    struct{}
//type  AVFilterPad     struct{}
type AVFilterFormats struct{}
type AVFilterChannelLayouts struct{}

/**
 * Get the number of elements in a NULL-terminated array of AVFilterPads (e.g.
 * AVFilter.inputs/outputs).
 */
//int avfilter_pad_count(const AVFilterPad *pads);
//未测试
func (pads *AVFilterPad) AvfilterPadCount() (res ffcommon.FInt, err error) {
	t, _, _ := ffcommon.GetAvfilterDll().NewProc("avfilter_pad_count").Call(
		uintptr(unsafe.Pointer(pads)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Get the name of an AVFilterPad.
 *
 * @param pads an array of AVFilterPads
 * @param pad_idx index of the pad in the array; it is the caller's
 *                responsibility to ensure the index is valid
 *
 * @return name of the pad_idx'th pad in pads
 */
//const char *avfilter_pad_get_name(const AVFilterPad *pads, int pad_idx);
//未测试
func (pads *AVFilterPad) AvfilterPadGetName(pad_idx ffcommon.FInt) (res ffcommon.FConstCharP, err error) {
	t, _, _ := ffcommon.GetAvfilterDll().NewProc("avfilter_pad_get_name").Call(
		uintptr(unsafe.Pointer(pads)),
		uintptr(pad_idx),
	)
	res = ffcommon.GoAStr(t)
	return
}

/**
 * Get the type of an AVFilterPad.
 *
 * @param pads an array of AVFilterPads
 * @param pad_idx index of the pad in the array; it is the caller's
 *                responsibility to ensure the index is valid
 *
 * @return type of the pad_idx'th pad in pads
 */
//enum AVMediaType avfilter_pad_get_type(const AVFilterPad *pads, int pad_idx);
//未测试
func (pads *AVFilterPad) AvfilterPadGetType(pad_idx ffcommon.FInt) (res ffconstant.AVMediaType, err error) {
	t, _, _ := ffcommon.GetAvfilterDll().NewProc("avfilter_pad_get_type").Call(
		uintptr(unsafe.Pointer(pads)),
		uintptr(pad_idx),
	)
	res = ffconstant.AVMediaType(t)
	return
}

/**
 * Filter definition. This defines the pads a filter contains, and all the
 * callback functions used to interact with the filter.
 */
type AVFilter struct {
	///**
	// * Filter name. Must be non-NULL and unique among filters.
	// */
	//const char *name;
	//
	///**
	// * A description of the filter. May be NULL.
	// *
	// * You should use the NULL_IF_CONFIG_SMALL() macro to define it.
	// */
	//const char *description;
	//
	///**
	// * List of inputs, terminated by a zeroed element.
	// *
	// * NULL if there are no (static) inputs. Instances of filters with
	// * AVFILTER_FLAG_DYNAMIC_INPUTS set may have more inputs than present in
	// * this list.
	// */
	//const AVFilterPad *inputs;
	///**
	// * List of outputs, terminated by a zeroed element.
	// *
	// * NULL if there are no (static) outputs. Instances of filters with
	// * AVFILTER_FLAG_DYNAMIC_OUTPUTS set may have more outputs than present in
	// * this list.
	// */
	//const AVFilterPad *outputs;
	//
	///**
	// * A class for the private data, used to declare filter private AVOptions.
	// * This field is NULL for filters that do not declare any options.
	// *
	// * If this field is non-NULL, the first member of the filter private data
	// * must be a pointer to AVClass, which will be set by libavfilter generic
	// * code to this class.
	// */
	//const AVClass *priv_class;
	//
	///**
	// * A combination of AVFILTER_FLAG_*
	// */
	//int flags;
	//
	///*****************************************************************
	// * All fields below this line are not part of the public API. They
	// * may not be used outside of libavfilter and can be changed and
	// * removed at will.
	// * New public fields should be added right above.
	// *****************************************************************
	// */
	//
	///**
	// * Filter pre-initialization function
	// *
	// * This callback will be called immediately after the filter context is
	// * allocated, to allow allocating and initing sub-objects.
	// *
	// * If this callback is not NULL, the uninit callback will be called on
	// * allocation failure.
	// *
	// * @return 0 on success,
	// *         AVERROR code on failure (but the code will be
	// *           dropped and treated as ENOMEM by the calling code)
	// */
	//int (*preinit)(AVFilterContext *ctx);
	//
	///**
	// * Filter initialization function.
	// *
	// * This callback will be called only once during the filter lifetime, after
	// * all the options have been set, but before links between filters are
	// * established and format negotiation is done.
	// *
	// * Basic filter initialization should be done here. Filters with dynamic
	// * inputs and/or outputs should create those inputs/outputs here based on
	// * provided options. No more changes to this filter's inputs/outputs can be
	// * done after this callback.
	// *
	// * This callback must not assume that the filter links exist or frame
	// * parameters are known.
	// *
	// * @ref AVFilter.uninit "uninit" is guaranteed to be called even if
	// * initialization fails, so this callback does not have to clean up on
	// * failure.
	// *
	// * @return 0 on success, a negative AVERROR on failure
	// */
	//int (*init)(AVFilterContext *ctx);
	//
	///**
	// * Should be set instead of @ref AVFilter.init "init" by the filters that
	// * want to pass a dictionary of AVOptions to nested contexts that are
	// * allocated during init.
	// *
	// * On return, the options dict should be freed and replaced with one that
	// * contains all the options which could not be processed by this filter (or
	// * with NULL if all the options were processed).
	// *
	// * Otherwise the semantics is the same as for @ref AVFilter.init "init".
	// */
	//int (*init_dict)(AVFilterContext *ctx, AVDictionary **options);
	//
	///**
	// * Filter uninitialization function.
	// *
	// * Called only once right before the filter is freed. Should deallocate any
	// * memory held by the filter, release any buffer references, etc. It does
	// * not need to deallocate the AVFilterContext.priv memory itself.
	// *
	// * This callback may be called even if @ref AVFilter.init "init" was not
	// * called or failed, so it must be prepared to handle such a situation.
	// */
	//void (*uninit)(AVFilterContext *ctx);
	//
	///**
	// * Query formats supported by the filter on its inputs and outputs.
	// *
	// * This callback is called after the filter is initialized (so the inputs
	// * and outputs are fixed), shortly before the format negotiation. This
	// * callback may be called more than once.
	// *
	// * This callback must set AVFilterLink.outcfg.formats on every input link and
	// * AVFilterLink.incfg.formats on every output link to a list of pixel/sample
	// * formats that the filter supports on that link. For audio links, this
	// * filter must also set @ref AVFilterLink.incfg.samplerates "in_samplerates" /
	// * @ref AVFilterLink.outcfg.samplerates "out_samplerates" and
	// * @ref AVFilterLink.incfg.channel_layouts "in_channel_layouts" /
	// * @ref AVFilterLink.outcfg.channel_layouts "out_channel_layouts" analogously.
	// *
	// * This callback may be NULL for filters with one input, in which case
	// * libavfilter assumes that it supports all input formats and preserves
	// * them on output.
	// *
	// * @return zero on success, a negative value corresponding to an
	// * AVERROR code otherwise
	// */
	//int (*query_formats)(AVFilterContext *);
	//
	//int priv_size;      ///< size of private data to allocate for the filter
	//
	//int flags_internal; ///< Additional flags for avfilter internal use only.
	//
	//#if FF_API_NEXT
	///**
	// * Used by the filter registration system. Must not be touched by any other
	// * code.
	// */
	//struct AVFilter *next;
	//#endif
	//
	///**
	// * Make the filter instance process a command.
	// *
	// * @param cmd    the command to process, for handling simplicity all commands must be alphanumeric only
	// * @param arg    the argument for the command
	// * @param res    a buffer with size res_size where the filter(s) can return a response. This must not change when the command is not supported.
	// * @param flags  if AVFILTER_CMD_FLAG_FAST is set and the command would be
	// *               time consuming then a filter should treat it like an unsupported command
	// *
	// * @returns >=0 on success otherwise an error code.
	// *          AVERROR(ENOSYS) on unsupported commands
	// */
	//int (*process_command)(AVFilterContext *, const char *cmd, const char *arg, char *res, int res_len, int flags);
	//
	///**
	// * Filter initialization function, alternative to the init()
	// * callback. Args contains the user-supplied parameters, opaque is
	// * used for providing binary data.
	// */
	//int (*init_opaque)(AVFilterContext *ctx, void *opaque);
	//
	///**
	// * Filter activation function.
	// *
	// * Called when any processing is needed from the filter, instead of any
	// * filter_frame and request_frame on pads.
	// *
	// * The function must examine inlinks and outlinks and perform a single
	// * step of processing. If there is nothing to do, the function must do
	// * nothing and not return an error. If more steps are or may be
	// * possible, it must use ff_filter_set_ready() to schedule another
	// * activation.
	// */
	//int (*activate)(AVFilterContext *ctx);
}

type AVFilterInternal struct {
}

/** An instance of a filter */
type AVFilterContext struct {
	//const AVClass *av_class;        ///< needed for av_log() and filters common options
	//
	//const AVFilter *filter;         ///< the AVFilter of which this is an instance
	//
	//char *name;                     ///< name of this filter instance
	//
	//AVFilterPad   *input_pads;      ///< array of input pads
	//AVFilterLink **inputs;          ///< array of pointers to input links
	//unsigned    nb_inputs;          ///< number of input pads
	//
	//AVFilterPad   *output_pads;     ///< array of output pads
	//AVFilterLink **outputs;         ///< array of pointers to output links
	//unsigned    nb_outputs;         ///< number of output pads
	//
	//void *priv;                     ///< private data for use by the filter
	//
	//struct AVFilterGraph *graph;    ///< filtergraph this filter belongs to
	//
	///**
	// * Type of multithreading being allowed/used. A combination of
	// * AVFILTER_THREAD_* flags.
	// *
	// * May be set by the caller before initializing the filter to forbid some
	// * or all kinds of multithreading for this filter. The default is allowing
	// * everything.
	// *
	// * When the filter is initialized, this field is combined using bit AND with
	// * AVFilterGraph.thread_type to get the final mask used for determining
	// * allowed threading types. I.e. a threading type needs to be set in both
	// * to be allowed.
	// *
	// * After the filter is initialized, libavfilter sets this field to the
	// * threading type that is actually used (0 for no multithreading).
	// */
	//int thread_type;
	//
	///**
	// * An opaque struct for libavfilter internal use.
	// */
	//AVFilterInternal *internal;
	//
	//struct AVFilterCommand *command_queue;
	//
	//char *enable_str;               ///< enable expression string
	//void *enable;                   ///< parsed expression (AVExpr*)
	//double *var_values;             ///< variable values for the enable expression
	//int is_disabled;                ///< the enabled state from the last expression evaluation
	//
	///**
	// * For filters which will create hardware frames, sets the device the
	// * filter should create them in.  All other filters will ignore this field:
	// * in particular, a filter which consumes or processes hardware frames will
	// * instead use the hw_frames_ctx field in AVFilterLink to carry the
	// * hardware context information.
	// */
	//AVBufferRef *hw_device_ctx;
	//
	///**
	// * Max number of threads allowed in this filter instance.
	// * If <= 0, its value is ignored.
	// * Overrides global number of threads set per filter graph.
	// */
	//int nb_threads;
	//
	///**
	// * Ready status of the filter.
	// * A non-0 value means that the filter needs activating;
	// * a higher value suggests a more urgent activation.
	// */
	//unsigned ready;
	//
	///**
	// * Sets the number of extra hardware frames which the filter will
	// * allocate on its output links for use in following filters or by
	// * the caller.
	// *
	// * Some hardware filters require all frames that they will use for
	// * output to be defined in advance before filtering starts.  For such
	// * filters, any hardware frame pools used for output must therefore be
	// * of fixed size.  The extra frames set here are on top of any number
	// * that the filter needs internally in order to operate normally.
	// *
	// * This field must be set before the graph containing this filter is
	// * configured.
	// */
	//int extra_hw_frames;
}

/**
 * Lists of formats / etc. supported by an end of a link.
 *
 * This structure is directly part of AVFilterLink, in two copies:
 * one for the source filter, one for the destination filter.

 * These lists are used for negotiating the format to actually be used,
 * which will be loaded into the format and channel_layout members of
 * AVFilterLink, when chosen.
 */
type AVFilterFormatsConfig struct {

	///**
	// * List of supported formats (pixel or sample).
	// */
	//AVFilterFormats *formats;
	//
	///**
	// * Lists of supported sample rates, only for audio.
	// */
	//AVFilterFormats  *samplerates;
	//
	///**
	// * Lists of supported channel layouts, only for audio.
	// */
	//AVFilterChannelLayouts  *channel_layouts;

}

/**
 * A link between two filters. This contains pointers to the source and
 * destination filters between which this link exists, and the indexes of
 * the pads involved. In addition, this link also contains the parameters
 * which have been negotiated and agreed upon between the filter, such as
 * image dimensions, format, etc.
 *
 * Applications must not normally access the link structure directly.
 * Use the buffersrc and buffersink API instead.
 * In the future, access to the header may be reserved for filters
 * implementation.
 */
type AVFilterLink struct {
	//AVFilterContext *src;       ///< source filter
	//AVFilterPad *srcpad;        ///< output pad on the source filter
	//
	//AVFilterContext *dst;       ///< dest filter
	//AVFilterPad *dstpad;        ///< input pad on the dest filter
	//
	//enum AVMediaType type;      ///< filter media type
	//
	///* These parameters apply only to video */
	//int w;                      ///< agreed upon image width
	//int h;                      ///< agreed upon image height
	//AVRational sample_aspect_ratio; ///< agreed upon sample aspect ratio
	///* These parameters apply only to audio */
	//uint64_t channel_layout;    ///< channel layout of current buffer (see libavutil/channel_layout.h)
	//int sample_rate;            ///< samples per second
	//
	//int format;                 ///< agreed upon media format
	//
	///**
	// * Define the time base used by the PTS of the frames/samples
	// * which will pass through this link.
	// * During the configuration stage, each filter is supposed to
	// * change only the output timebase, while the timebase of the
	// * input link is assumed to be an unchangeable property.
	// */
	//AVRational time_base;
	//
	///*****************************************************************
	// * All fields below this line are not part of the public API. They
	// * may not be used outside of libavfilter and can be changed and
	// * removed at will.
	// * New public fields should be added right above.
	// *****************************************************************
	// */
	//
	///**
	// * Lists of supported formats / etc. supported by the input filter.
	// */
	//AVFilterFormatsConfig incfg;
	//
	///**
	// * Lists of supported formats / etc. supported by the output filter.
	// */
	//AVFilterFormatsConfig outcfg;
	//
	///** stage of the initialization of the link properties (dimensions, etc) */
	//enum {
	//AVLINK_UNINIT = 0,      ///< not started
	//AVLINK_STARTINIT,       ///< started, but incomplete
	//AVLINK_INIT             ///< complete
	//} init_state;
	//
	///**
	// * Graph the filter belongs to.
	// */
	//struct AVFilterGraph *graph;
	//
	///**
	// * Current timestamp of the link, as defined by the most recent
	// * frame(s), in link time_base units.
	// */
	//int64_t current_pts;
	//
	///**
	// * Current timestamp of the link, as defined by the most recent
	// * frame(s), in AV_TIME_BASE units.
	// */
	//int64_t current_pts_us;
	//
	///**
	// * Index in the age array.
	// */
	//int age_index;
	//
	///**
	// * Frame rate of the stream on the link, or 1/0 if unknown or variable;
	// * if left to 0/0, will be automatically copied from the first input
	// * of the source filter if it exists.
	// *
	// * Sources should set it to the best estimation of the real frame rate.
	// * If the source frame rate is unknown or variable, set this to 1/0.
	// * Filters should update it if necessary depending on their function.
	// * Sinks can use it to set a default output frame rate.
	// * It is similar to the r_frame_rate field in AVStream.
	// */
	//AVRational frame_rate;
	//
	///**
	// * Buffer partially filled with samples to achieve a fixed/minimum size.
	// */
	//AVFrame *partial_buf;
	//
	///**
	// * Size of the partial buffer to allocate.
	// * Must be between min_samples and max_samples.
	// */
	//int partial_buf_size;
	//
	///**
	// * Minimum number of samples to filter at once. If filter_frame() is
	// * called with fewer samples, it will accumulate them in partial_buf.
	// * This field and the related ones must not be changed after filtering
	// * has started.
	// * If 0, all related fields are ignored.
	// */
	//int min_samples;
	//
	///**
	// * Maximum number of samples to filter at once. If filter_frame() is
	// * called with more samples, it will split them.
	// */
	//int max_samples;
	//
	///**
	// * Number of channels.
	// */
	//int channels;
	//
	///**
	// * Number of past frames sent through the link.
	// */
	//int64_t frame_count_in, frame_count_out;
	//
	///**
	// * A pointer to a FFFramePool struct.
	// */
	//void *frame_pool;
	//
	///**
	// * True if a frame is currently wanted on the output of this filter.
	// * Set when ff_request_frame() is called by the output,
	// * cleared when a frame is filtered.
	// */
	//int frame_wanted_out;
	//
	///**
	// * For hwaccel pixel formats, this should be a reference to the
	// * AVHWFramesContext describing the frames.
	// */
	//AVBufferRef *hw_frames_ctx;
	//
	//#ifndef FF_INTERNAL_FIELDS
	//
	///**
	// * Internal structure members.
	// * The fields below this limit are internal for libavfilter's use
	// * and must in no way be accessed by applications.
	// */
	//char reserved[0xF000];
	//
	//#else /* FF_INTERNAL_FIELDS */
	//
	///**
	// * Queue of frames waiting to be filtered.
	// */
	//FFFrameQueue fifo;
	//
	///**
	// * If set, the source filter can not generate a frame as is.
	// * The goal is to avoid repeatedly calling the request_frame() method on
	// * the same link.
	// */
	//int frame_blocked_in;
	//
	///**
	// * Link input status.
	// * If not zero, all attempts of filter_frame will fail with the
	// * corresponding code.
	// */
	//int status_in;
	//
	///**
	// * Timestamp of the input status change.
	// */
	//int64_t status_in_pts;
	//
	///**
	// * Link output status.
	// * If not zero, all attempts of request_frame will fail with the
	// * corresponding code.
	// */
	//int status_out;
	//
	//#endif /* FF_INTERNAL_FIELDS */

}

/**
 * Link two filters together.
 *
 * @param src    the source filter
 * @param srcpad index of the output pad on the source filter
 * @param dst    the destination filter
 * @param dstpad index of the input pad on the destination filter
 * @return       zero on success
 */
//int avfilter_link(AVFilterContext *src, unsigned srcpad,
//AVFilterContext *dst, unsigned dstpad);
//未测试
func AvfilterLink(src *AVFilterContext, srcpad ffcommon.FUnsigned,
	dst *AVFilterContext, dstpad ffcommon.FUnsigned) (res ffcommon.FInt, err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("avfilter_link").Call(
		uintptr(unsafe.Pointer(src)),
		uintptr(srcpad),
		uintptr(unsafe.Pointer(dst)),
		uintptr(dstpad),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Free the link in *link, and set its pointer to NULL.
 */
//void avfilter_link_free(AVFilterLink **link);
//未测试
func AvfilterLinkFree(link **AVFilterLink) (err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("avfilter_link_free").Call(
		uintptr(unsafe.Pointer(&link)),
	)
	if t == 0 {

	}
	return
}

//#if FF_API_FILTER_GET_SET
///**
// * Get the number of channels of a link.
// * @deprecated Use av_buffersink_get_channels()
// */
//attribute_deprecated
//int avfilter_link_get_channels(AVFilterLink *link);
//未测试
func (link *AVFilterLink) AvfilterLinkGetChannels() (res ffcommon.FInt, err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("avfilter_link_get_channels").Call(
		uintptr(unsafe.Pointer(link)),
	)
	res = ffcommon.FInt(t)
	return
}

//#endif
//#if FF_API_FILTER_LINK_SET_CLOSED
///**
// * Set the closed field of a link.
// * @deprecated applications are not supposed to mess with links, they should
// * close the sinks.
// */
//attribute_deprecated
//void avfilter_link_set_closed(AVFilterLink *link, int closed);
//未测试
func (link *AVFilterLink) AvfilterLinkSetClosed(closed ffcommon.FInt) (res ffcommon.FInt, err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("avfilter_link_set_closed").Call(
		uintptr(unsafe.Pointer(link)),
		uintptr(closed),
	)
	res = ffcommon.FInt(t)
	return
}

//#endif
/**
 * Negotiate the media format, dimensions, etc of all inputs to a filter.
 *
 * @param filter the filter to negotiate the properties for its inputs
 * @return       zero on successful negotiation
 */
//int avfilter_config_links(AVFilterContext *filter);
//未测试
func (link *AVFilterLink) AvfilterConfigLinks() (res ffcommon.FInt, err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("avfilter_config_links").Call(
		uintptr(unsafe.Pointer(link)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Make the filter instance process a command.
 * It is recommended to use avfilter_graph_send_command().
 */
//int avfilter_process_command(AVFilterContext *filter, const char *cmd, const char *arg, char *res, int res_len, int flags);
//未测试
func (link *AVFilterLink) AvfilterProcessCommand(cmd ffcommon.FConstCharP, arg ffcommon.FConstCharP, res0 ffcommon.FConstCharP, res_len ffcommon.FInt, flags ffcommon.FInt) (res ffcommon.FInt, err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("avfilter_process_command").Call(
		uintptr(unsafe.Pointer(link)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Iterate over all registered filters.
 *
 * @param opaque a pointer where libavfilter will store the iteration state. Must
 *               point to NULL to start the iteration.
 *
 * @return the next registered filter or NULL when the iteration is
 *         finished
 */
//const AVFilter *av_filter_iterate(void **opaque);
//未测试
func AvFilterIterate(opaque *ffcommon.FVoidP) (res *AVFilter, err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_filter_iterate").Call(
		uintptr(unsafe.Pointer(opaque)),
	)
	res = (*AVFilter)(unsafe.Pointer(t))
	return
}

//#if FF_API_NEXT
///** Initialize the filter system. Register all builtin filters. */
//attribute_deprecated
//void avfilter_register_all(void);
//未测试
func AvfilterRegisterAll() (err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("avfilter_register_all").Call()
	if t == 0 {

	}
	return
}

/**
 * Register a filter. This is only needed if you plan to use
 * avfilter_get_by_name later to lookup the AVFilter structure by name. A
 * filter can still by instantiated with avfilter_graph_alloc_filter even if it
 * is not registered.
 *
 * @param filter the filter to register
 * @return 0 if the registration was successful, a negative value
 * otherwise
 */
//attribute_deprecated
//int avfilter_register(AVFilter *filter);
//未测试
func (filter *AVFilter) AvfilterRegister() (res ffcommon.FInt, err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("avfilter_register_all").Call(
		uintptr(unsafe.Pointer(filter)),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Iterate over all registered filters.
 * @return If prev is non-NULL, next registered filter after prev or NULL if
 * prev is the last filter. If prev is NULL, return the first registered filter.
 */
//attribute_deprecated
//const AVFilter *avfilter_next(const AVFilter *prev);
//#endif
//未测试
func (prev *AVFilter) AvfilterNext() (res *AVFilter, err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("avfilter_next").Call(
		uintptr(unsafe.Pointer(prev)),
	)
	if t == 0 {

	}
	res = (*AVFilter)(unsafe.Pointer(t))
	return
}

/**
 * Get a filter definition matching the given name.
 *
 * @param name the filter name to find
 * @return     the filter definition, if any matching one is registered.
 *             NULL if none found.
 */
//const AVFilter *avfilter_get_by_name(const char *name);
//未测试
func AvfilterGetByName(name ffcommon.FConstCharP) (res *AVFilter, err error) {
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ := ffcommon.GetAvutilDll().NewProc("avfilter_get_by_name").Call(
		uintptr(unsafe.Pointer(namep)),
	)
	if t == 0 {

	}
	res = (*AVFilter)(unsafe.Pointer(t))
	return
}

/**
 * Initialize a filter with the supplied parameters.
 *
 * @param ctx  uninitialized filter context to initialize
 * @param args Options to initialize the filter with. This must be a
 *             ':'-separated list of options in the 'key=value' form.
 *             May be NULL if the options have been set directly using the
 *             AVOptions API or there are no options that need to be set.
 * @return 0 on success, a negative AVERROR on failure
 */
//int avfilter_init_str(AVFilterContext *ctx, const char *args);
//未测试
func (ctx *AVFilterContext) AvfilterInitStr(args *ffcommon.FBuf) (res ffcommon.FInt, err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("avfilter_init_str").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(&args)),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Initialize a filter with the supplied dictionary of options.
 *
 * @param ctx     uninitialized filter context to initialize
 * @param options An AVDictionary filled with options for this filter. On
 *                return this parameter will be destroyed and replaced with
 *                a dict containing options that were not found. This dictionary
 *                must be freed by the caller.
 *                May be NULL, then this function is equivalent to
 *                avfilter_init_str() with the second parameter set to NULL.
 * @return 0 on success, a negative AVERROR on failure
 *
 * @note This function and avfilter_init_str() do essentially the same thing,
 * the difference is in manner in which the options are passed. It is up to the
 * calling code to choose whichever is more preferable. The two functions also
 * behave differently when some of the provided options are not declared as
 * supported by the filter. In such a case, avfilter_init_str() will fail, but
 * this function will leave those extra options in the options AVDictionary and
 * continue as usual.
 */
//int avfilter_init_dict(AVFilterContext *ctx, AVDictionary **options);
//未测试
func (ctx *AVFilterContext) AvfilterInitDict(options **libavutil.AVDictionary) (res ffcommon.FInt, err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("avfilter_init_dict").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(&options)),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Free a filter context. This will also remove the filter from its
 * filtergraph's list of filters.
 *
 * @param filter the filter to free
 */
//void avfilter_free(AVFilterContext *filter);
//未测试
func (filter *AVFilterContext) AvfilterFree() (err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("avfilter_free").Call(
		uintptr(unsafe.Pointer(filter)),
	)
	if t == 0 {

	}
	return
}

/**
 * Insert a filter in the middle of an existing link.
 *
 * @param link the link into which the filter should be inserted
 * @param filt the filter to be inserted
 * @param filt_srcpad_idx the input pad on the filter to connect
 * @param filt_dstpad_idx the output pad on the filter to connect
 * @return     zero on success
 */
//int avfilter_insert_filter(AVFilterLink *link, AVFilterContext *filt,
//unsigned filt_srcpad_idx, unsigned filt_dstpad_idx);
//未测试
func (link *AVFilterLink) AvfilterInsertFilter(filt *AVFilterContext,
	filt_srcpad_idx ffcommon.FUnsigned, filt_dstpad_idx ffcommon.FUnsigned) (res ffcommon.FInt, err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("avfilter_insert_filter").Call(
		uintptr(unsafe.Pointer(link)),
		uintptr(unsafe.Pointer(filt)),
		uintptr(filt_srcpad_idx),
		uintptr(filt_dstpad_idx),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * @return AVClass for AVFilterContext.
 *
 * @see av_opt_find().
 */
//const AVClass *avfilter_get_class(void);
//未测试
func AvfilterGetClass() (res *libavutil.AVClass, err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("avfilter_get_class").Call()
	if t == 0 {

	}
	res = (*libavutil.AVClass)(unsafe.Pointer(t))
	return
}

type AVFilterGraphInternal struct {
}

/**
 * A function pointer passed to the @ref AVFilterGraph.execute callback to be
 * executed multiple times, possibly in parallel.
 *
 * @param ctx the filter context the job belongs to
 * @param arg an opaque parameter passed through from @ref
 *            AVFilterGraph.execute
 * @param jobnr the index of the job being executed
 * @param nb_jobs the total number of jobs
 *
 * @return 0 on success, a negative AVERROR on error
 */
type AvfilterActionFunc = func(ctx *AVFilterContext, arg ffcommon.FVoidP, jobnr ffcommon.FInt, nb_jobs ffcommon.FInt) ffcommon.FInt

/**
 * A function executing multiple jobs, possibly in parallel.
 *
 * @param ctx the filter context to which the jobs belong
 * @param func the function to be called multiple times
 * @param arg the argument to be passed to func
 * @param ret a nb_jobs-sized array to be filled with return values from each
 *            invocation of func
 * @param nb_jobs the number of jobs to execute
 *
 * @return 0 on success, a negative AVERROR on error
 */
var AvfilterExecuteFunc func(ctx *AVFilterContext, f AvfilterActionFunc,
	arg ffcommon.FVoidP, ret *ffcommon.FInt, nb_jobs ffcommon.FInt) ffcommon.FInt

type AVFilterGraph struct {
	//const AVClass *av_class;
	//AVFilterContext **filters;
	//unsigned nb_filters;
	//
	//char *scale_sws_opts; ///< sws options to use for the auto-inserted scale filters
	//#if FF_API_LAVR_OPTS
	//	attribute_deprecated char *resample_lavr_opts;   ///< libavresample options to use for the auto-inserted resample filters
	//#endif
	//
	///**
	// * Type of multithreading allowed for filters in this graph. A combination
	// * of AVFILTER_THREAD_* flags.
	// *
	// * May be set by the caller at any point, the setting will apply to all
	// * filters initialized after that. The default is allowing everything.
	// *
	// * When a filter in this graph is initialized, this field is combined using
	// * bit AND with AVFilterContext.thread_type to get the final mask used for
	// * determining allowed threading types. I.e. a threading type needs to be
	// * set in both to be allowed.
	// */
	//int thread_type;
	//
	///**
	// * Maximum number of threads used by filters in this graph. May be set by
	// * the caller before adding any filters to the filtergraph. Zero (the
	// * default) means that the number of threads is determined automatically.
	// */
	//int nb_threads;
	//
	///**
	// * Opaque object for libavfilter internal use.
	// */
	//AVFilterGraphInternal *internal;
	//
	///**
	// * Opaque user data. May be set by the caller to an arbitrary value, e.g. to
	// * be used from callbacks like @ref AVFilterGraph.execute.
	// * Libavfilter will not touch this field in any way.
	// */
	//void *opaque;
	//
	///**
	// * This callback may be set by the caller immediately after allocating the
	// * graph and before adding any filters to it, to provide a custom
	// * multithreading implementation.
	// *
	// * If set, filters with slice threading capability will call this callback
	// * to execute multiple jobs in parallel.
	// *
	// * If this field is left unset, libavfilter will use its internal
	// * implementation, which may or may not be multithreaded depending on the
	// * platform and build options.
	// */
	//avfilter_execute_func *execute;
	//
	//char *aresample_swr_opts; ///< swr options to use for the auto-inserted aresample filters, Access ONLY through AVOptions
	//
	///**
	// * Private fields
	// *
	// * The following fields are for internal use only.
	// * Their type, offset, number and semantic can change without notice.
	// */
	//
	//AVFilterLink **sink_links;
	//int sink_links_count;
	//
	//unsigned disable_auto_convert;
}

/**
 * Allocate a filter graph.
 *
 * @return the allocated filter graph on success or NULL.
 */
//AVFilterGraph *avfilter_graph_alloc(void);
//未测试
func AvfilterGraphAlloc() (res *AVFilterGraph, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("avfilter_graph_alloc").Call()
	if err != nil {
		//return
	}
	//res = &AVAES{}
	res = (*AVFilterGraph)(unsafe.Pointer(t))
	//res.instance = t
	//res.ptr = unsafe.Pointer(t)
	return
}

/**
 * Create a new filter instance in a filter graph.
 *
 * @param graph graph in which the new filter will be used
 * @param filter the filter to create an instance of
 * @param name Name to give to the new instance (will be copied to
 *             AVFilterContext.name). This may be used by the caller to identify
 *             different filters, libavfilter itself assigns no semantics to
 *             this parameter. May be NULL.
 *
 * @return the context of the newly created filter instance (note that it is
 *         also retrievable directly through AVFilterGraph.filters or with
 *         avfilter_graph_get_filter()) on success or NULL on failure.
 */
//AVFilterContext *avfilter_graph_alloc_filter(AVFilterGraph *graph,
//const AVFilter *filter,
//const char *name);
//未测试
func (graph *AVFilterGraph) AvfilterGraphAllocFilter(filter *AVFilter,
	name ffcommon.FConstCharP) (res *AVFilterContext, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, err = ffcommon.GetAvutilDll().NewProc("avfilter_graph_alloc_filter").Call(
		uintptr(unsafe.Pointer(graph)),
		uintptr(unsafe.Pointer(filter)),
		uintptr(unsafe.Pointer(namep)),
	)
	if err != nil {
		//return
	}
	res = (*AVFilterContext)(unsafe.Pointer(t))
	return
}

/**
 * Get a filter instance identified by instance name from graph.
 *
 * @param graph filter graph to search through.
 * @param name filter instance name (should be unique in the graph).
 * @return the pointer to the found filter instance or NULL if it
 * cannot be found.
 */
//AVFilterContext *avfilter_graph_get_filter(AVFilterGraph *graph, const char *name);
//未测试
func (graph *AVFilterGraph) AvfilterGraphGetFilter(
	name ffcommon.FConstCharP) (res *AVFilterContext, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, err = ffcommon.GetAvutilDll().NewProc("avfilter_graph_get_filter").Call(
		uintptr(unsafe.Pointer(graph)),
		uintptr(unsafe.Pointer(namep)),
	)
	if err != nil {
		//return
	}
	res = (*AVFilterContext)(unsafe.Pointer(t))
	return
}

/**
 * Create and add a filter instance into an existing graph.
 * The filter instance is created from the filter filt and inited
 * with the parameter args. opaque is currently ignored.
 *
 * In case of success put in *filt_ctx the pointer to the created
 * filter instance, otherwise set *filt_ctx to NULL.
 *
 * @param name the instance name to give to the created filter instance
 * @param graph_ctx the filter graph
 * @return a negative AVERROR error code in case of failure, a non
 * negative value otherwise
 */
//int avfilter_graph_create_filter(AVFilterContext **filt_ctx, const AVFilter *filt,
//const char *name, const char *args, void *opaque,
//AVFilterGraph *graph_ctx);
//未测试
func AvfilterGraphCreateFilter(filt_ctx **AVFilterContext, filt *AVFilter,
	name ffcommon.FConstCharP, args ffcommon.FConstCharP, opaque ffcommon.FVoidP,
	graph_ctx *AVFilterGraph) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	var argsp *byte
	argsp, err = syscall.BytePtrFromString(args)
	if err != nil {
		return
	}
	t, _, err = ffcommon.GetAvutilDll().NewProc("avfilter_graph_create_filter").Call(
		uintptr(unsafe.Pointer(&filt_ctx)),
		uintptr(unsafe.Pointer(filt)),
		uintptr(unsafe.Pointer(namep)),
		uintptr(unsafe.Pointer(argsp)),
		opaque,
		uintptr(unsafe.Pointer(graph_ctx)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Enable or disable automatic format conversion inside the graph.
 *
 * Note that format conversion can still happen inside explicitly inserted
 * scale and aresample filters.
 *
 * @param flags  any of the AVFILTER_AUTO_CONVERT_* ffconstants
 */
//void avfilter_graph_set_auto_convert(AVFilterGraph *graph, unsigned flags);
//未测试
func (graph *AVFilterGraph) AvfilterGraphSetAutoConvert(flags ffcommon.FUnsigned) (err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("avfilter_graph_set_auto_convert").Call(
		uintptr(unsafe.Pointer(graph)),
		uintptr(flags),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Check validity and configure all the links and formats in the graph.
 *
 * @param graphctx the filter graph
 * @param log_ctx context used for logging
 * @return >= 0 in case of success, a negative AVERROR code otherwise
 */
//int avfilter_graph_config(AVFilterGraph *graphctx, void *log_ctx);
//未测试
func (graphctx *AVFilterGraph) AvfilterGraphConfig(log_ctx ffcommon.FVoidP) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("avfilter_graph_config").Call(
		uintptr(unsafe.Pointer(graphctx)),
		log_ctx,
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
 * Free a graph, destroy its links, and set *graph to NULL.
 * If *graph is NULL, do nothing.
 */
//void avfilter_graph_free(AVFilterGraph **graph);
//未测试
func AvfilterGraphFree(graphctx **AVFilterGraph) (err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("avfilter_graph_free").Call(
		uintptr(unsafe.Pointer(&graphctx)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * A linked-list of the inputs/outputs of the filter chain.
 *
 * This is mainly useful for avfilter_graph_parse() / avfilter_graph_parse2(),
 * where it is used to communicate open (unlinked) inputs and outputs from and
 * to the caller.
 * This struct specifies, per each not connected pad contained in the graph, the
 * filter context and the pad index required for establishing a link.
 */
type AVFilterInOut struct {
	///** unique name for this input/output in the list */
	//char *name;
	//
	///** filter context associated to this input/output */
	//AVFilterContext *filter_ctx;
	//
	///** index of the filt_ctx pad to use for linking */
	//int pad_idx;
	//
	///** next input/input in the list, NULL if this is the last */
	//struct AVFilterInOut *next;
}

/**
 * Allocate a single AVFilterInOut entry.
 * Must be freed with avfilter_inout_free().
 * @return allocated AVFilterInOut on success, NULL on failure.
 */
//AVFilterInOut *avfilter_inout_alloc(void);
//未测试
func AvfilterInoutAlloc() (res *AVFilterInOut, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("avfilter_inout_alloc").Call()
	if err != nil {
		//return
	}
	//res = &AVAES{}
	res = (*AVFilterInOut)(unsafe.Pointer(t))
	//res.instance = t
	//res.ptr = unsafe.Pointer(t)
	return
}

/**
 * Free the supplied list of AVFilterInOut and set *inout to NULL.
 * If *inout is NULL, do nothing.
 */
//void avfilter_inout_free(AVFilterInOut **inout);
//未测试
func AvfilterInoutFree(inout **AVFilterInOut) (err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("avfilter_inout_free").Call(
		uintptr(unsafe.Pointer(&inout)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Add a graph described by a string to a graph.
 *
 * @note The caller must provide the lists of inputs and outputs,
 * which therefore must be known before calling the function.
 *
 * @note The inputs parameter describes inputs of the already existing
 * part of the graph; i.e. from the point of view of the newly created
 * part, they are outputs. Similarly the outputs parameter describes
 * outputs of the already existing filters, which are provided as
 * inputs to the parsed filters.
 *
 * @param graph   the filter graph where to link the parsed graph context
 * @param filters string to be parsed
 * @param inputs  linked list to the inputs of the graph
 * @param outputs linked list to the outputs of the graph
 * @return zero on success, a negative AVERROR code on error
 */
//int avfilter_graph_parse(AVFilterGraph *graph, const char *filters,
//AVFilterInOut *inputs, AVFilterInOut *outputs,
//void *log_ctx);
//未测试
func (graph *AVFilterGraph) AvfilterGraphParse(filters ffcommon.FConstCharP,
	inputs *AVFilterInOut, outputs *AVFilterInOut,
	log_ctx ffcommon.FVoidP) (res ffcommon.FInt, err error) {
	var t uintptr
	var filtersp *byte
	filtersp, err = syscall.BytePtrFromString(filters)
	if err != nil {
		return
	}
	t, _, err = ffcommon.GetAvutilDll().NewProc("avfilter_graph_parse").Call(
		uintptr(unsafe.Pointer(graph)),
		uintptr(unsafe.Pointer(filtersp)),
		uintptr(unsafe.Pointer(inputs)),
		uintptr(unsafe.Pointer(outputs)),
		log_ctx,
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
 * Add a graph described by a string to a graph.
 *
 * In the graph filters description, if the input label of the first
 * filter is not specified, "in" is assumed; if the output label of
 * the last filter is not specified, "out" is assumed.
 *
 * @param graph   the filter graph where to link the parsed graph context
 * @param filters string to be parsed
 * @param inputs  pointer to a linked list to the inputs of the graph, may be NULL.
 *                If non-NULL, *inputs is updated to contain the list of open inputs
 *                after the parsing, should be freed with avfilter_inout_free().
 * @param outputs pointer to a linked list to the outputs of the graph, may be NULL.
 *                If non-NULL, *outputs is updated to contain the list of open outputs
 *                after the parsing, should be freed with avfilter_inout_free().
 * @return non negative on success, a negative AVERROR code on error
 */
//int avfilter_graph_parse_ptr(AVFilterGraph *graph, const char *filters,
//AVFilterInOut **inputs, AVFilterInOut **outputs,
//void *log_ctx);
//未测试
func (graph *AVFilterGraph) AvfilterGraphParsePtr(filters ffcommon.FConstCharP,
	inputs *AVFilterInOut, outputs *AVFilterInOut,
	log_ctx ffcommon.FVoidP) (res ffcommon.FInt, err error) {
	var t uintptr
	var filtersp *byte
	filtersp, err = syscall.BytePtrFromString(filters)
	if err != nil {
		return
	}
	t, _, err = ffcommon.GetAvutilDll().NewProc("avfilter_graph_parse_ptr").Call(
		uintptr(unsafe.Pointer(graph)),
		uintptr(unsafe.Pointer(filtersp)),
		uintptr(unsafe.Pointer(inputs)),
		uintptr(unsafe.Pointer(outputs)),
		log_ctx,
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
 * Add a graph described by a string to a graph.
 *
 * @param[in]  graph   the filter graph where to link the parsed graph context
 * @param[in]  filters string to be parsed
 * @param[out] inputs  a linked list of all free (unlinked) inputs of the
 *                     parsed graph will be returned here. It is to be freed
 *                     by the caller using avfilter_inout_free().
 * @param[out] outputs a linked list of all free (unlinked) outputs of the
 *                     parsed graph will be returned here. It is to be freed by the
 *                     caller using avfilter_inout_free().
 * @return zero on success, a negative AVERROR code on error
 *
 * @note This function returns the inputs and outputs that are left
 * unlinked after parsing the graph and the caller then deals with
 * them.
 * @note This function makes no reference whatsoever to already
 * existing parts of the graph and the inputs parameter will on return
 * contain inputs of the newly parsed part of the graph.  Analogously
 * the outputs parameter will contain outputs of the newly created
 * filters.
 */
//int avfilter_graph_parse2(AVFilterGraph *graph, const char *filters,
//AVFilterInOut **inputs,
//AVFilterInOut **outputs);
//未测试
func (graph *AVFilterGraph) AvfilterGraphParse2(filters ffcommon.FConstCharP,
	inputs **AVFilterInOut, outputs **AVFilterInOut) (res ffcommon.FInt, err error) {
	var t uintptr
	var filtersp *byte
	filtersp, err = syscall.BytePtrFromString(filters)
	if err != nil {
		return
	}
	t, _, err = ffcommon.GetAvutilDll().NewProc("avfilter_graph_parse2").Call(
		uintptr(unsafe.Pointer(graph)),
		uintptr(unsafe.Pointer(filtersp)),
		uintptr(unsafe.Pointer(inputs)),
		uintptr(unsafe.Pointer(outputs)),
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
 * Send a command to one or more filter instances.
 *
 * @param graph  the filter graph
 * @param target the filter(s) to which the command should be sent
 *               "all" sends to all filters
 *               otherwise it can be a filter or filter instance name
 *               which will send the command to all matching filters.
 * @param cmd    the command to send, for handling simplicity all commands must be alphanumeric only
 * @param arg    the argument for the command
 * @param res    a buffer with size res_size where the filter(s) can return a response.
 *
 * @returns >=0 on success otherwise an error code.
 *              AVERROR(ENOSYS) on unsupported commands
 */
//int avfilter_graph_send_command(AVFilterGraph *graph, const char *target, const char *cmd, const char *arg, char *res, int res_len, int flags);
//未测试
func (graph *AVFilterGraph) AvfilterGraphSendCommand(target ffcommon.FConstCharP,
	cmd ffcommon.FConstCharP,
	arg ffcommon.FCharP,
	res0 ffcommon.FBuf,
	res0_len ffcommon.FInt,
	flags ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	var targetp *byte
	targetp, err = syscall.BytePtrFromString(target)
	if err != nil {
		return
	}
	var cmdp *byte
	cmdp, err = syscall.BytePtrFromString(cmd)
	if err != nil {
		return
	}
	var argp *byte
	argp, err = syscall.BytePtrFromString(arg)
	if err != nil {
		return
	}
	t, _, err = ffcommon.GetAvutilDll().NewProc("avfilter_graph_send_command").Call(
		uintptr(unsafe.Pointer(graph)),
		uintptr(unsafe.Pointer(targetp)),
		uintptr(unsafe.Pointer(cmdp)),
		uintptr(unsafe.Pointer(argp)),
		uintptr(unsafe.Pointer(res0)),
		uintptr(res0_len),
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
 * Queue a command for one or more filter instances.
 *
 * @param graph  the filter graph
 * @param target the filter(s) to which the command should be sent
 *               "all" sends to all filters
 *               otherwise it can be a filter or filter instance name
 *               which will send the command to all matching filters.
 * @param cmd    the command to sent, for handling simplicity all commands must be alphanumeric only
 * @param arg    the argument for the command
 * @param ts     time at which the command should be sent to the filter
 *
 * @note As this executes commands after this function returns, no return code
 *       from the filter is provided, also AVFILTER_CMD_FLAG_ONE is not supported.
 */
//int avfilter_graph_queue_command(AVFilterGraph *graph, const char *target, const char *cmd, const char *arg, int flags, double ts);
//未测试
func (graph *AVFilterGraph) AvfilterGraphQueueCommand(target ffcommon.FConstCharP,
	cmd ffcommon.FConstCharP,
	arg ffcommon.FCharP,
	flags ffcommon.FInt,
	ts ffcommon.FDouble) (res ffcommon.FInt, err error) {
	var t uintptr
	var targetp *byte
	targetp, err = syscall.BytePtrFromString(target)
	if err != nil {
		return
	}
	var cmdp *byte
	cmdp, err = syscall.BytePtrFromString(cmd)
	if err != nil {
		return
	}
	var argp *byte
	argp, err = syscall.BytePtrFromString(arg)
	if err != nil {
		return
	}
	t, _, err = ffcommon.GetAvutilDll().NewProc("avfilter_graph_queue_command").Call(
		uintptr(unsafe.Pointer(graph)),
		uintptr(unsafe.Pointer(targetp)),
		uintptr(unsafe.Pointer(cmdp)),
		uintptr(unsafe.Pointer(argp)),
		uintptr(flags),
		uintptr(ts),
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
 * Dump a graph into a human-readable string representation.
 *
 * @param graph    the graph to dump
 * @param options  formatting options; currently ignored
 * @return  a string, or NULL in case of memory allocation failure;
 *          the string must be freed using av_free
 */
//char *avfilter_graph_dump(AVFilterGraph *graph, const char *options);
//未测试
func (graph *AVFilterGraph) AvfilterGraphDump(options ffcommon.FConstCharP) (res ffcommon.FInt, err error) {
	var t uintptr
	var targetp *byte
	targetp, err = syscall.BytePtrFromString(options)
	if err != nil {
		return
	}
	t, _, err = ffcommon.GetAvutilDll().NewProc("avfilter_graph_dump").Call(
		uintptr(unsafe.Pointer(graph)),
		uintptr(unsafe.Pointer(targetp)),
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
 * Request a frame on the oldest sink link.
 *
 * If the request returns AVERROR_EOF, try the next.
 *
 * Note that this function is not meant to be the sole scheduling mechanism
 * of a filtergraph, only a convenience function to help drain a filtergraph
 * in a balanced way under normal circumstances.
 *
 * Also note that AVERROR_EOF does not mean that frames did not arrive on
 * some of the sinks during the process.
 * When there are multiple sink links, in case the requested link
 * returns an EOF, this may cause a filter to flush pending frames
 * which are sent to another sink link, although unrequested.
 *
 * @return  the return value of ff_request_frame(),
 *          or AVERROR_EOF if all links returned AVERROR_EOF
 */
//int avfilter_graph_request_oldest(AVFilterGraph *graph);
//未测试
func (graph *AVFilterGraph) AvfilterGraphRequestOldest() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("avfilter_graph_request_oldest").Call(
		uintptr(unsafe.Pointer(graph)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}
