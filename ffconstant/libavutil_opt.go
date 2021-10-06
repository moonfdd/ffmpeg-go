package ffconstant

/**
 * @defgroup avoptions AVOptions
 * @ingroup lavu_data
 * @{
 * AVOptions provide a generic system to declare options on arbitrary structs
 * ("objects"). An option can have a help text, a type and a range of possible
 * values. Options may then be enumerated, read and written to.
 *
 * @section avoptions_implement Implementing AVOptions
 * This section describes how to add AVOptions capabilities to a struct.
 *
 * All AVOptions-related information is stored in an AVClass. Therefore
 * the first member of the struct should be a pointer to an AVClass describing it.
 * The option field of the AVClass must be set to a NULL-terminated static array
 * of AVOptions. Each AVOption must have a non-empty name, a type, a default
 * value and for number-type AVOptions also a range of allowed values. It must
 * also declare an offset in bytes from the start of the struct, where the field
 * associated with this AVOption is located. Other fields in the AVOption struct
 * should also be set when applicable, but are not required.
 *
 * The following example illustrates an AVOptions-enabled struct:
 * @code
 * typedef struct test_struct {
 *     const AVClass *class;
 *     int      int_opt;
 *     char    *str_opt;
 *     uint8_t *bin_opt;
 *     int      bin_len;
 * } test_struct;
 *
 * static const AVOption test_options[] = {
 *   { "test_int", "This is a test option of int type.", offsetof(test_struct, int_opt),
 *     AV_OPT_TYPE_INT, { .i64 = -1 }, INT_MIN, INT_MAX },
 *   { "test_str", "This is a test option of string type.", offsetof(test_struct, str_opt),
 *     AV_OPT_TYPE_STRING },
 *   { "test_bin", "This is a test option of binary type.", offsetof(test_struct, bin_opt),
 *     AV_OPT_TYPE_BINARY },
 *   { NULL },
 * };
 *
 * static const AVClass test_class = {
 *     .class_name = "test class",
 *     .item_name  = av_default_item_name,
 *     .option     = test_options,
 *     .version    = LIBAVUTIL_VERSION_INT,
 * };
 * @endcode
 *
 * Next, when allocating your struct, you must ensure that the AVClass pointer
 * is set to the correct value. Then, av_opt_set_defaults() can be called to
 * initialize defaults. After that the struct is ready to be used with the
 * AVOptions API.
 *
 * When cleaning up, you may use the av_opt_free() function to automatically
 * free all the allocated string and binary options.
 *
 * Continuing with the above example:
 *
 * @code
 * test_struct *alloc_test_struct(void)
 * {
 *     test_struct *ret = av_mallocz(sizeof(*ret));
 *     ret->class = &test_class;
 *     av_opt_set_defaults(ret);
 *     return ret;
 * }
 * void free_test_struct(test_struct **foo)
 * {
 *     av_opt_free(*foo);
 *     av_freep(foo);
 * }
 * @endcode
 *
 * @subsection avoptions_implement_nesting Nesting
 *      It may happen that an AVOptions-enabled struct contains another
 *      AVOptions-enabled struct as a member (e.g. AVCodecContext in
 *      libavcodec exports generic options, while its priv_data field exports
 *      codec-specific options). In such a case, it is possible to set up the
 *      parent struct to export a child's options. To do that, simply
 *      implement AVClass.child_next() and AVClass.child_class_iterate() in the
 *      parent struct's AVClass.
 *      Assuming that the test_struct from above now also contains a
 *      child_struct field:
 *
 *      @code
 *      typedef struct child_struct {
 *          AVClass *class;
 *          int flags_opt;
 *      } child_struct;
 *      static const AVOption child_opts[] = {
 *          { "test_flags", "This is a test option of flags type.",
 *            offsetof(child_struct, flags_opt), AV_OPT_TYPE_FLAGS, { .i64 = 0 }, INT_MIN, INT_MAX },
 *          { NULL },
 *      };
 *      static const AVClass child_class = {
 *          .class_name = "child class",
 *          .item_name  = av_default_item_name,
 *          .option     = child_opts,
 *          .version    = LIBAVUTIL_VERSION_INT,
 *      };
 *
 *      void *child_next(void *obj, void *prev)
 *      {
 *          test_struct *t = obj;
 *          if (!prev && t->child_struct)
 *              return t->child_struct;
 *          return NULL
 *      }
 *      const AVClass child_class_iterate(void **iter)
 *      {
 *          const AVClass *c = *iter ? NULL : &child_class;
 *          *iter = (void*)(uintptr_t)c;
 *          return c;
 *      }
 *      @endcode
 *      Putting child_next() and child_class_iterate() as defined above into
 *      test_class will now make child_struct's options accessible through
 *      test_struct (again, proper setup as described above needs to be done on
 *      child_struct right after it is created).
 *
 *      From the above example it might not be clear why both child_next()
 *      and child_class_iterate() are needed. The distinction is that child_next()
 *      iterates over actually existing objects, while child_class_iterate()
 *      iterates over all possible child classes. E.g. if an AVCodecContext
 *      was initialized to use a codec which has private options, then its
 *      child_next() will return AVCodecContext.priv_data and finish
 *      iterating. OTOH child_class_iterate() on AVCodecContext.av_class will
 *      iterate over all available codecs with private options.
 *
 * @subsection avoptions_implement_named_ffconstants Named ffconstants
 *      It is possible to create named ffconstants for options. Simply set the unit
 *      field of the option the ffconstants should apply to a string and
 *      create the ffconstants themselves as options of type AV_OPT_TYPE_CONST
 *      with their unit field set to the same string.
 *      Their default_val field should contain the value of the named
 *      ffconstant.
 *      For example, to add some named ffconstants for the test_flags option
 *      above, put the following into the child_opts array:
 *      @code
 *      { "test_flags", "This is a test option of flags type.",
 *        offsetof(child_struct, flags_opt), AV_OPT_TYPE_FLAGS, { .i64 = 0 }, INT_MIN, INT_MAX, "test_unit" },
 *      { "flag1", "This is a flag with value 16", 0, AV_OPT_TYPE_CONST, { .i64 = 16 }, 0, 0, "test_unit" },
 *      @endcode
 *
 * @section avoptions_use Using AVOptions
 * This section deals with accessing options in an AVOptions-enabled struct.
 * Such structs in FFmpeg are e.g. AVCodecContext in libavcodec or
 * AVFormatContext in libavformat.
 *
 * @subsection avoptions_use_examine Examining AVOptions
 * The basic functions for examining options are av_opt_next(), which iterates
 * over all options defined for one object, and av_opt_find(), which searches
 * for an option with the given name.
 *
 * The situation is more complicated with nesting. An AVOptions-enabled struct
 * may have AVOptions-enabled children. Passing the AV_OPT_SEARCH_CHILDREN flag
 * to av_opt_find() will make the function search children recursively.
 *
 * For enumerating there are basically two cases. The first is when you want to
 * get all options that may potentially exist on the struct and its children
 * (e.g.  when constructing documentation). In that case you should call
 * av_opt_child_class_iterate() recursively on the parent struct's AVClass.  The
 * second case is when you have an already initialized struct with all its
 * children and you want to get all options that can be actually written or read
 * from it. In that case you should call av_opt_child_next() recursively (and
 * av_opt_next() on each result).
 *
 * @subsection avoptions_use_get_set Reading and writing AVOptions
 * When setting options, you often have a string read directly from the
 * user. In such a case, simply passing it to av_opt_set() is enough. For
 * non-string type options, av_opt_set() will parse the string according to the
 * option type.
 *
 * Similarly av_opt_get() will read any option type and convert it to a string
 * which will be returned. Do not forget that the string is allocated, so you
 * have to free it with av_free().
 *
 * In some cases it may be more convenient to put all options into an
 * AVDictionary and call av_opt_set_dict() on it. A specific case of this
 * are the format/codec open functions in lavf/lavc which take a dictionary
 * filled with option as a parameter. This makes it possible to set some options
 * that cannot be set otherwise, since e.g. the input file format is not known
 * before the file is actually opened.
 */
type AVOptionType int32

const (
	AV_OPT_TYPE_FLAGS = 0
	AV_OPT_TYPE_INT
	AV_OPT_TYPE_INT64
	AV_OPT_TYPE_DOUBLE
	AV_OPT_TYPE_FLOAT
	AV_OPT_TYPE_STRING
	AV_OPT_TYPE_RATIONAL
	AV_OPT_TYPE_BINARY ///< offset must point to a pointer immediately followed by an int for the length
	AV_OPT_TYPE_DICT
	AV_OPT_TYPE_UINT64
	AV_OPT_TYPE_CONST
	AV_OPT_TYPE_IMAGE_SIZE ///< offset must point to two consecutive integers
	AV_OPT_TYPE_PIXEL_FMT
	AV_OPT_TYPE_SAMPLE_FMT
	AV_OPT_TYPE_VIDEO_RATE ///< offset must point to AVRational
	AV_OPT_TYPE_DURATION
	AV_OPT_TYPE_COLOR
	AV_OPT_TYPE_CHANNEL_LAYOUT
	AV_OPT_TYPE_BOOL
)

const AV_OPT_FLAG_ENCODING_PARAM = 1 ///< a generic parameter which can be set by the user for muxing or encoding
const AV_OPT_FLAG_DECODING_PARAM = 2 ///< a generic parameter which can be set by the user for demuxing or decoding
const AV_OPT_FLAG_AUDIO_PARAM = 8
const AV_OPT_FLAG_VIDEO_PARAM = 16
const AV_OPT_FLAG_SUBTITLE_PARAM = 32

/**
 * The option is intended for exporting values to the caller.
 */
const AV_OPT_FLAG_EXPORT = 64

/**
 * The option may not be set through the AVOptions API, only read.
 * This flag only makes sense when AV_OPT_FLAG_EXPORT is also set.
 */
const AV_OPT_FLAG_READONLY = 128
const AV_OPT_FLAG_BSF_PARAM = (1 << 8)        ///< a generic parameter which can be set by the user for bit stream filtering
const AV_OPT_FLAG_RUNTIME_PARAM = (1 << 15)   ///< a generic parameter which can be set by the user at runtime
const AV_OPT_FLAG_FILTERING_PARAM = (1 << 16) ///< a generic parameter which can be set by the user for filtering
const AV_OPT_FLAG_DEPRECATED = (1 << 17)      ///< set if option is deprecated, users should refer to AVOption.help text for more information
const AV_OPT_FLAG_CHILD_CONSTS = (1 << 18)    ///< set if option ffconstants can also reside in child objects

const AV_OPT_FLAG_IMPLICIT_KEY = 1

const AV_OPT_SEARCH_CHILDREN = (1 << 0) /**< Search in possible children of the
  given object first. */
/**
 *  The obj passed to av_opt_find() is fake -- only a double pointer to AVClass
 *  instead of a required pointer to a struct containing AVClass. This is
 *  useful for searching for options without needing to allocate the corresponding
 *  object.
 */
const AV_OPT_SEARCH_FAKE_OBJ = (1 << 1)

/**
 *  In av_opt_get, return NULL if the option has a pointer type and is set to NULL,
 *  rather than returning an empty string.
 */
const AV_OPT_ALLOW_NULL = (1 << 2)

/**
 *  Allows av_opt_query_ranges and av_opt_query_ranges_default to return more than
 *  one component for certain option types.
 *  @see AVOptionRanges for details.
 */
const AV_OPT_MULTI_COMPONENT_RANGE = (1 << 12)

const AV_OPT_SERIALIZE_SKIP_DEFAULTS = 0x00000001   ///< Serialize options that are not set to default values only.
const AV_OPT_SERIALIZE_OPT_FLAGS_EXACT = 0x00000002 ///< Serialize options that exactly match opt_flags only.
