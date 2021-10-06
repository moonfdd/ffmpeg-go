package ffconstant

/**
 * The number of the filter inputs is not determined just by AVFilter.inputs.
 * The filter might add additional inputs during initialization depending on the
 * options supplied to it.
 */
const AVFILTER_FLAG_DYNAMIC_INPUTS = (1 << 0)

/**
 * The number of the filter outputs is not determined just by AVFilter.outputs.
 * The filter might add additional outputs during initialization depending on
 * the options supplied to it.
 */
const AVFILTER_FLAG_DYNAMIC_OUTPUTS = (1 << 1)

/**
 * The filter supports multithreading by splitting frames into multiple parts
 * and processing them concurrently.
 */
const AVFILTER_FLAG_SLICE_THREADS = (1 << 2)

/**
 * Some filters support a generic "enable" expression option that can be used
 * to enable or disable a filter in the timeline. Filters supporting this
 * option have this flag set. When the enable expression is false, the default
 * no-op filter_frame() function is called in place of the filter_frame()
 * callback defined on each input pad, thus the frame is passed unchanged to
 * the next filters.
 */
const AVFILTER_FLAG_SUPPORT_TIMELINE_GENERIC = (1 << 16)

/**
 * Same as AVFILTER_FLAG_SUPPORT_TIMELINE_GENERIC, except that the filter will
 * have its filter_frame() callback(s) called as usual even when the enable
 * expression is false. The filter will disable filtering within the
 * filter_frame() callback(s) itself, for example executing code depending on
 * the AVFilterContext->is_disabled value.
 */
const AVFILTER_FLAG_SUPPORT_TIMELINE_INTERNAL = (1 << 17)

/**
 * Handy mask to test whether the filter supports or no the timeline feature
 * (internally or generically).
 */
const AVFILTER_FLAG_SUPPORT_TIMELINE = (AVFILTER_FLAG_SUPPORT_TIMELINE_GENERIC | AVFILTER_FLAG_SUPPORT_TIMELINE_INTERNAL)

/**
 * Process multiple parts of the frame concurrently.
 */
const AVFILTER_THREAD_SLICE = (1 << 0)

/** stage of the initialization of the link properties (dimensions, etc) */
type InitState int32

const (
	AVLINK_UNINIT    = 0 ///< not started
	AVLINK_STARTINIT     ///< started, but incomplete
	AVLINK_INIT          ///< complete
)

const AVFILTER_CMD_FLAG_ONE = 1  ///< Stop once a filter understood the command (for target=all for example), fast filters are favored automatically
const AVFILTER_CMD_FLAG_FAST = 2 ///< Only execute command when its fast (like a video out that supports contrast adjustment in hw)

const (
	AVFILTER_AUTO_CONVERT_ALL  = 0  /**< all automatic conversions enabled */
	AVFILTER_AUTO_CONVERT_NONE = -1 /**< all automatic conversions disabled */
)
