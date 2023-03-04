package libavutil

import (
	"math"
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
)

/*
 * AVOptions
 * copyright (c) 2005 Michael Niedermayer <michaelni@gmx.at>
 *
 * This file is part of FFmpeg.
 *
 * FFmpeg is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * FFmpeg is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with FFmpeg; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA
 */

//#ifndef AVUTIL_OPT_H
//#define AVUTIL_OPT_H

/**
 * @file
 * AVOptions
 */

//#include "rational.h"
//#include "avutil.h"
//#include "dict.h"
//#include "log.h"
//#include "pixfmt.h"
//#include "samplefmt.h"
//#include "version.h"

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
 * @subsection avoptions_implement_named_constants Named constants
 *      It is possible to create named constants for options. Simply set the unit
 *      field of the option the constants should apply to a string and
 *      create the constants themselves as options of type AV_OPT_TYPE_CONST
 *      with their unit field set to the same string.
 *      Their default_val field should contain the value of the named
 *      constant.
 *      For example, to add some named constants for the test_flags option
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
	AV_OPT_TYPE_FLAGS = iota
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

/**
 * AVOption
 */
type AVOption struct {
	Name ffcommon.FCharPStruct

	/**
	 * short English help text
	 * @todo What about other languages?
	 */
	Help ffcommon.FCharPStruct

	/**
	 * The offset relative to the context structure where the option
	 * value is stored. It should be 0 for named constants.
	 */
	Offset ffcommon.FInt
	Type   AVOptionType

	/**
	 * the default value for scalar options
	 */
	//union {
	//int64_t i64;
	//double dbl;
	//const char *str;
	///* TODO those are unused now */
	//AVRational q;
	//} default_val;
	DefaultVal AVRational
	Min        ffcommon.FDouble ///< minimum valid value for the option
	Max        ffcommon.FDouble ///< maximum valid value for the option

	Flags ffcommon.FInt
	//const AV_OPT_FLAG_ENCODING_PARAM  =1   ///< a generic parameter which can be set by the user for muxing or encoding
	//	const AV_OPT_FLAG_DECODING_PARAM = 2   ///< a generic parameter which can be set by the user for demuxing or decoding
	//	const AV_OPT_FLAG_AUDIO_PARAM   =  8
	//	const AV_OPT_FLAG_VIDEO_PARAM   =  16
	//	const AV_OPT_FLAG_SUBTITLE_PARAM = 32
	///**
	// * The option is intended for exporting values to the caller.
	// */
	//	const AV_OPT_FLAG_EXPORT      =    64
	///**
	// * The option may not be set through the AVOptions API, only read.
	// * This flag only makes sense when AV_OPT_FLAG_EXPORT is also set.
	// */
	//	const AV_OPT_FLAG_READONLY     =   128
	//	const AV_OPT_FLAG_BSF_PARAM     =  (1<<8) ///< a generic parameter which can be set by the user for bit stream filtering
	//	const AV_OPT_FLAG_RUNTIME_PARAM  = (1<<15) ///< a generic parameter which can be set by the user at runtime
	//	const AV_OPT_FLAG_FILTERING_PARAM= (1<<16) ///< a generic parameter which can be set by the user for filtering
	//	const AV_OPT_FLAG_DEPRECATED    =  (1<<17) ///< set if option is deprecated, users should refer to AVOption.help text for more information
	//	const AV_OPT_FLAG_CHILD_CONSTS  =  (1<<18) ///< set if option constants can also reside in child objects
	//FIXME think about enc-audio, ... style flags

	/**
	 * The logical unit to which the option belongs. Non-constant
	 * options and corresponding named constants share the same
	 * unit. May be NULL.
	 */
	Unit ffcommon.FCharPStruct
}

/**
 * A single allowed range of values, or a single allowed value.
 */
type AVOptionRange struct {
	Str ffcommon.FCharPStruct
	/**
	 * Value range.
	 * For string ranges this represents the min/max length.
	 * For dimensions this represents the min/max pixel count or width/height in multi-component case.
	 */
	ValueMin, ValueMax ffcommon.FDouble
	/**
	 * Value's component range.
	 * For string this represents the unicode range for chars, 0-127 limits to ASCII.
	 */
	ComponentMin, ComponentMax ffcommon.FDouble
	/**
	 * Range flag.
	 * If set to 1 the struct encodes a range, if set to 0 a single value.
	 */
	IsRange ffcommon.FInt
}

/**
 * List of AVOptionRange structs.
 */
type AVOptionRanges struct {

	/**
	 * Array of option ranges.
	 *
	 * Most of option types use just one component.
	 * Following describes multi-component option types:
	 *
	 * AV_OPT_TYPE_IMAGE_SIZE:
	 * component index 0: range of pixel count (width * height).
	 * component index 1: range of width.
	 * component index 2: range of height.
	 *
	 * @note To obtain multi-component version of this structure, user must
	 *       provide AV_OPT_MULTI_COMPONENT_RANGE to av_opt_query_ranges or
	 *       av_opt_query_ranges_default function.
	 *
	 * Multi-component range can be read as in following example:
	 *
	 * @code
	 * int range_index, component_index;
	 * AVOptionRanges *ranges;
	 * AVOptionRange *range[3]; //may require more than 3 in the future.
	 * av_opt_query_ranges(&ranges, obj, key, AV_OPT_MULTI_COMPONENT_RANGE);
	 * for (range_index = 0; range_index < ranges->nb_ranges; range_index++) {
	 *     for (component_index = 0; component_index < ranges->nb_components; component_index++)
	 *         range[component_index] = ranges->range[ranges->nb_ranges * component_index + range_index];
	 *     //do something with range here.
	 * }
	 * av_opt_freep_ranges(&ranges);
	 * @endcode
	 */
	Range **AVOptionRange
	/**
	 * Number of ranges per component.
	 */
	NbRanges ffcommon.FInt
	/**
	 * Number of componentes.
	 */
	NbComponents ffcommon.FInt
}

/**
 * Show the obj options.
 *
 * @param req_flags requested flags for the options to show. Show only the
 * options for which it is opt->flags & req_flags.
 * @param rej_flags rejected flags for the options to show. Show only the
 * options for which it is !(opt->flags & req_flags).
 * @param av_log_obj log context to use for showing the options
 */
//int av_opt_show2(void *obj, void *av_log_obj, int req_flags, int rej_flags);
func AvOptShow2(obj, av_log_obj ffcommon.FVoidP, req_flags, rej_flags ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_show2").Call(
		obj,
		av_log_obj,
		uintptr(req_flags),
		uintptr(rej_flags),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Set the values of all AVOption fields to their default values.
 *
 * @param s an AVOption-enabled struct (its first member must be a pointer to AVClass)
 */
//void av_opt_set_defaults(void *s);
func AvOptSetDefaults(s ffcommon.FVoidP) {
	ffcommon.GetAvutilDll().NewProc("av_opt_set_defaults").Call(
		s,
	)
}

/**
 * Set the values of all AVOption fields to their default values. Only these
 * AVOption fields for which (opt->flags & mask) == flags will have their
 * default applied to s.
 *
 * @param s an AVOption-enabled struct (its first member must be a pointer to AVClass)
 * @param mask combination of AV_OPT_FLAG_*
 * @param flags combination of AV_OPT_FLAG_*
 */
//void av_opt_set_defaults2(void *s, int mask, int flags);
func AvOptSetDefaults2(s ffcommon.FVoidP, mask, flags ffcommon.FInt) {
	ffcommon.GetAvutilDll().NewProc("av_opt_set_defaults2").Call(
		s,
		uintptr(mask),
		uintptr(flags),
	)
}

/**
 * Parse the key/value pairs list in opts. For each key/value pair
 * found, stores the value in the field in ctx that is named like the
 * key. ctx must be an AVClass context, storing is done using
 * AVOptions.
 *
 * @param opts options string to parse, may be NULL
 * @param key_val_sep a 0-terminated list of characters used to
 * separate key from value
 * @param pairs_sep a 0-terminated list of characters used to separate
 * two pairs from each other
 * @return the number of successfully set key/value pairs, or a negative
 * value corresponding to an AVERROR code in case of error:
 * AVERROR(EINVAL) if opts cannot be parsed,
 * the error code issued by av_opt_set() if a key/value pair
 * cannot be set
 */
//int av_set_options_string(void *ctx, const char *opts,
//const char *key_val_sep, const char *pairs_sep);
func AvSetOptionsString(ctx ffcommon.FVoidP, opts,
	key_val_sep, pairs_sep ffcommon.FConstCharP) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_set_options_string").Call(
		ctx,
		ffcommon.UintPtrFromString(opts),
		ffcommon.UintPtrFromString(key_val_sep),
		ffcommon.UintPtrFromString(pairs_sep),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Parse the key-value pairs list in opts. For each key=value pair found,
 * set the value of the corresponding option in ctx.
 *
 * @param ctx          the AVClass object to set options on
 * @param opts         the options string, key-value pairs separated by a
 *                     delimiter
 * @param shorthand    a NULL-terminated array of options names for shorthand
 *                     notation: if the first field in opts has no key part,
 *                     the key is taken from the first element of shorthand;
 *                     then again for the second, etc., until either opts is
 *                     finished, shorthand is finished or a named option is
 *                     found; after that, all options must be named
 * @param key_val_sep  a 0-terminated list of characters used to separate
 *                     key from value, for example '='
 * @param pairs_sep    a 0-terminated list of characters used to separate
 *                     two pairs from each other, for example ':' or ','
 * @return  the number of successfully set key=value pairs, or a negative
 *          value corresponding to an AVERROR code in case of error:
 *          AVERROR(EINVAL) if opts cannot be parsed,
 *          the error code issued by av_set_string3() if a key/value pair
 *          cannot be set
 *
 * Options names must use only the following characters: a-z A-Z 0-9 - . / _
 * Separators must use characters distinct from option names and from each
 * other.
 */
//int av_opt_set_from_string(void *ctx, const char *opts,
//const char *const *shorthand,
//const char *key_val_sep, const char *pairs_sep);
func AvOptSetFromString(ctx ffcommon.FVoidP, opts ffcommon.FConstCharP,
	shorthand *ffcommon.FBuf,
	key_val_sep, pairs_sep ffcommon.FConstCharP) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_set_from_string").Call(
		ctx,
		ffcommon.UintPtrFromString(opts),
		uintptr(unsafe.Pointer(shorthand)),
		ffcommon.UintPtrFromString(key_val_sep),
		ffcommon.UintPtrFromString(pairs_sep),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Free all allocated objects in obj.
 */
//void av_opt_free(void *obj);
func AvOptFree(obj ffcommon.FVoidP) {
	ffcommon.GetAvutilDll().NewProc("av_opt_free").Call(
		obj,
	)
}

/**
 * Check whether a particular flag is set in a flags field.
 *
 * @param field_name the name of the flag field option
 * @param flag_name the name of the flag to check
 * @return non-zero if the flag is set, zero if the flag isn't set,
 *         isn't of the right type, or the flags field doesn't exist.
 */
//int av_opt_flag_is_set(void *obj, const char *field_name, const char *flag_name);
func AvOptFlagIsSet(obj ffcommon.FVoidP, field_name, flag_name ffcommon.FConstCharP) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_flag_is_set").Call(
		obj,
		ffcommon.UintPtrFromString(field_name),
		ffcommon.UintPtrFromString(flag_name),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Set all the options from a given dictionary on an object.
 *
 * @param obj a struct whose first element is a pointer to AVClass
 * @param options options to process. This dictionary will be freed and replaced
 *                by a new one containing all options not found in obj.
 *                Of course this new dictionary needs to be freed by caller
 *                with av_dict_free().
 *
 * @return 0 on success, a negative AVERROR if some option was found in obj,
 *         but could not be set.
 *
 * @see av_dict_copy()
 */
//int av_opt_set_dict(void *obj, struct AVDictionary **options);
func AvOptSetDict(obj ffcommon.FVoidP, options **AVDictionary) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_set_dict").Call(
		obj,
		uintptr(unsafe.Pointer(options)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Set all the options from a given dictionary on an object.
 *
 * @param obj a struct whose first element is a pointer to AVClass
 * @param options options to process. This dictionary will be freed and replaced
 *                by a new one containing all options not found in obj.
 *                Of course this new dictionary needs to be freed by caller
 *                with av_dict_free().
 * @param search_flags A combination of AV_OPT_SEARCH_*.
 *
 * @return 0 on success, a negative AVERROR if some option was found in obj,
 *         but could not be set.
 *
 * @see av_dict_copy()
 */
//int av_opt_set_dict2(void *obj, struct AVDictionary **options, int search_flags);
func AvOptSetDict2(obj ffcommon.FVoidP, options **AVDictionary, search_flags ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_set_dict2").Call(
		obj,
		uintptr(unsafe.Pointer(options)),
		uintptr(search_flags),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Extract a key-value pair from the beginning of a string.
 *
 * @param ropts        pointer to the options string, will be updated to
 *                     point to the rest of the string (one of the pairs_sep
 *                     or the final NUL)
 * @param key_val_sep  a 0-terminated list of characters used to separate
 *                     key from value, for example '='
 * @param pairs_sep    a 0-terminated list of characters used to separate
 *                     two pairs from each other, for example ':' or ','
 * @param flags        flags; see the AV_OPT_FLAG_* values below
 * @param rkey         parsed key; must be freed using av_free()
 * @param rval         parsed value; must be freed using av_free()
 *
 * @return  >=0 for success, or a negative value corresponding to an
 *          AVERROR code in case of error; in particular:
 *          AVERROR(EINVAL) if no key is present
 *
 */
//int av_opt_get_key_value(const char **ropts,
//const char *key_val_sep, const char *pairs_sep,
//unsigned flags,
//char **rkey, char **rval);
func AvOptGetKeyValue(ropts *ffcommon.FBuf,
	key_val_sep, pairs_sep ffcommon.FConstCharP,
	flags ffcommon.FUnsigned,
	rkey, rval *ffcommon.FBuf) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_get_key_value").Call(
		uintptr(unsafe.Pointer(ropts)),
		ffcommon.UintPtrFromString(key_val_sep),
		ffcommon.UintPtrFromString(pairs_sep),
		uintptr(flags),
		uintptr(unsafe.Pointer(rkey)),
		uintptr(unsafe.Pointer(rval)),
	)
	res = ffcommon.FInt(t)
	return
}

//enum {

/**
 * Accept to parse a value without a key; the key will then be returned
 * as NULL.
 */
const AV_OPT_FLAG_IMPLICIT_KEY = 1

//};

/**
 * @defgroup opt_eval_funcs Evaluating option strings
 * @{
 * This group of functions can be used to evaluate option strings
 * and get numbers out of them. They do the same thing as av_opt_set(),
 * except the result is written into the caller-supplied pointer.
 *
 * @param obj a struct whose first element is a pointer to AVClass.
 * @param o an option for which the string is to be evaluated.
 * @param val string to be evaluated.
 * @param *_out value of the string will be written here.
 *
 * @return 0 on success, a negative number on failure.
 */
//int av_opt_eval_flags (void *obj, const AVOption *o, const char *val, int        *flags_out);
func AvOptEvalFlags(obj ffcommon.FVoidP, o *AVOption, val ffcommon.FConstCharP, flags_out *ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_eval_flags").Call(
		obj,
		uintptr(unsafe.Pointer(o)),
		ffcommon.UintPtrFromString(val),
		uintptr(unsafe.Pointer(flags_out)),
	)
	res = ffcommon.FInt(t)
	return
}

// int av_opt_eval_int   (void *obj, const AVOption *o, const char *val, int        *int_out);
func AvOptEvalInt(obj ffcommon.FVoidP, o *AVOption, val ffcommon.FConstCharP, int_out *ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_eval_int").Call(
		obj,
		uintptr(unsafe.Pointer(o)),
		ffcommon.UintPtrFromString(val),
		uintptr(unsafe.Pointer(int_out)),
	)
	res = ffcommon.FInt(t)
	return
}

// int av_opt_eval_int64 (void *obj, const AVOption *o, const char *val, int64_t    *int64_out);
func AvOptEvalInt64(obj ffcommon.FVoidP, o *AVOption, val ffcommon.FConstCharP, int64_out *ffcommon.FInt64T) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_eval_int64").Call(
		obj,
		uintptr(unsafe.Pointer(o)),
		ffcommon.UintPtrFromString(val),
		uintptr(unsafe.Pointer(int64_out)),
	)
	res = ffcommon.FInt(t)
	return
}

// int av_opt_eval_float (void *obj, const AVOption *o, const char *val, float      *float_out);
func AvOptEvalFloat(obj ffcommon.FVoidP, o *AVOption, val ffcommon.FConstCharP, float_out *ffcommon.FFloat) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_eval_float").Call(
		obj,
		uintptr(unsafe.Pointer(o)),
		ffcommon.UintPtrFromString(val),
		uintptr(unsafe.Pointer(float_out)),
	)
	res = ffcommon.FInt(t)
	return
}

// int av_opt_eval_double(void *obj, const AVOption *o, const char *val, double     *double_out);
func AvOptEvalDouble(obj ffcommon.FVoidP, o *AVOption, val ffcommon.FConstCharP, double_out *ffcommon.FDouble) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_eval_double").Call(
		obj,
		uintptr(unsafe.Pointer(o)),
		ffcommon.UintPtrFromString(val),
		uintptr(unsafe.Pointer(double_out)),
	)
	res = ffcommon.FInt(t)
	return
}

// int av_opt_eval_q     (void *obj, const AVOption *o, const char *val, AVRational *q_out);
func AvOptEvalQ(obj ffcommon.FVoidP, o *AVOption, val ffcommon.FConstCharP, q_out *AVRational) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_eval_q").Call(
		obj,
		uintptr(unsafe.Pointer(o)),
		ffcommon.UintPtrFromString(val),
		uintptr(unsafe.Pointer(q_out)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * @}
 */

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

/**
 * Look for an option in an object. Consider only options which
 * have all the specified flags set.
 *
 * @param[in] obj A pointer to a struct whose first element is a
 *                pointer to an AVClass.
 *                Alternatively a double pointer to an AVClass, if
 *                AV_OPT_SEARCH_FAKE_OBJ search flag is set.
 * @param[in] name The name of the option to look for.
 * @param[in] unit When searching for named constants, name of the unit
 *                 it belongs to.
 * @param opt_flags Find only options with all the specified flags set (AV_OPT_FLAG).
 * @param search_flags A combination of AV_OPT_SEARCH_*.
 *
 * @return A pointer to the option found, or NULL if no option
 *         was found.
 *
 * @note Options found with AV_OPT_SEARCH_CHILDREN flag may not be settable
 * directly with av_opt_set(). Use special calls which take an options
 * AVDictionary (e.g. avformat_open_input()) to set options found with this
 * flag.
 */
//const AVOption *av_opt_find(void *obj, const char *name, const char *unit,
//int opt_flags, int search_flags);
func AvOptFind(obj ffcommon.FVoidP, name, unit ffcommon.FConstCharP,
	opt_flags, search_flags ffcommon.FInt) (res *AVOption) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_find").Call(
		obj,
		ffcommon.UintPtrFromString(name),
		ffcommon.UintPtrFromString(unit),
		uintptr(opt_flags),
		uintptr(search_flags),
	)
	res = (*AVOption)(unsafe.Pointer(t))
	return
}

/**
 * Look for an option in an object. Consider only options which
 * have all the specified flags set.
 *
 * @param[in] obj A pointer to a struct whose first element is a
 *                pointer to an AVClass.
 *                Alternatively a double pointer to an AVClass, if
 *                AV_OPT_SEARCH_FAKE_OBJ search flag is set.
 * @param[in] name The name of the option to look for.
 * @param[in] unit When searching for named constants, name of the unit
 *                 it belongs to.
 * @param opt_flags Find only options with all the specified flags set (AV_OPT_FLAG).
 * @param search_flags A combination of AV_OPT_SEARCH_*.
 * @param[out] target_obj if non-NULL, an object to which the option belongs will be
 * written here. It may be different from obj if AV_OPT_SEARCH_CHILDREN is present
 * in search_flags. This parameter is ignored if search_flags contain
 * AV_OPT_SEARCH_FAKE_OBJ.
 *
 * @return A pointer to the option found, or NULL if no option
 *         was found.
 */
//const AVOption *av_opt_find2(void *obj, const char *name, const char *unit,
//int opt_flags, int search_flags, void **target_obj);
func AvOptFind2(obj ffcommon.FVoidP, name, unit ffcommon.FConstCharP,
	opt_flags, search_flags ffcommon.FInt, target_obj *ffcommon.FVoidP) (res *AVOption) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_find2").Call(
		obj,
		ffcommon.UintPtrFromString(name),
		ffcommon.UintPtrFromString(unit),
		uintptr(opt_flags),
		uintptr(search_flags),
		uintptr(unsafe.Pointer(target_obj)),
	)
	res = (*AVOption)(unsafe.Pointer(t))
	return
}

/**
 * Iterate over all AVOptions belonging to obj.
 *
 * @param obj an AVOptions-enabled struct or a double pointer to an
 *            AVClass describing it.
 * @param prev result of the previous call to av_opt_next() on this object
 *             or NULL
 * @return next AVOption or NULL
 */
//const AVOption *av_opt_next(const void *obj, const AVOption *prev);
func AvOptNext(obj ffcommon.FVoidP, prev *AVOption) (res *AVOption) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_next").Call(
		obj,
		uintptr(unsafe.Pointer(prev)),
	)
	res = (*AVOption)(unsafe.Pointer(t))
	return
}

/**
 * Iterate over AVOptions-enabled children of obj.
 *
 * @param prev result of a previous call to this function or NULL
 * @return next AVOptions-enabled child or NULL
 */
//void *av_opt_child_next(void *obj, void *prev);
func AvOptChildNext(obj ffcommon.FVoidP, prev *AVOption) (res ffcommon.FVoidP) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_child_next").Call(
		obj,
		uintptr(unsafe.Pointer(prev)),
	)
	res = t
	return
}

//#if FF_API_CHILD_CLASS_NEXT
/**
 * Iterate over potential AVOptions-enabled children of parent.
 *
 * @param prev result of a previous call to this function or NULL
 * @return AVClass corresponding to next potential child or NULL
 *
 * @deprecated use av_opt_child_class_iterate
 */
//attribute_deprecated
//const AVClass *av_opt_child_class_next(const AVClass *parent, const AVClass *prev);
func (parent *AVClass) AvOptChildClassNext(prev *AVClass) (res *AVClass) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_child_class_next").Call(
		uintptr(unsafe.Pointer(parent)),
		uintptr(unsafe.Pointer(prev)),
	)
	res = (*AVClass)(unsafe.Pointer(t))
	return
}

//#endif

/**
 * Iterate over potential AVOptions-enabled children of parent.
 *
 * @param iter a pointer where iteration state is stored.
 * @return AVClass corresponding to next potential child or NULL
 */
//const AVClass *av_opt_child_class_iterate(const AVClass *parent, void **iter);
func (parent *AVClass) AvOptChildClassIterate(iter *ffcommon.FVoidP) (res *AVClass) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_child_class_iterate").Call(
		uintptr(unsafe.Pointer(parent)),
		uintptr(unsafe.Pointer(iter)),
	)
	res = (*AVClass)(unsafe.Pointer(t))
	return
}

/**
 * @defgroup opt_set_funcs Option setting functions
 * @{
 * Those functions set the field of obj with the given name to value.
 *
 * @param[in] obj A struct whose first element is a pointer to an AVClass.
 * @param[in] name the name of the field to set
 * @param[in] val The value to set. In case of av_opt_set() if the field is not
 * of a string type, then the given string is parsed.
 * SI postfixes and some named scalars are supported.
 * If the field is of a numeric type, it has to be a numeric or named
 * scalar. Behavior with more than one scalar and +- infix operators
 * is undefined.
 * If the field is of a flags type, it has to be a sequence of numeric
 * scalars or named flags separated by '+' or '-'. Prefixing a flag
 * with '+' causes it to be set without affecting the other flags;
 * similarly, '-' unsets a flag.
 * If the field is of a dictionary type, it has to be a ':' separated list of
 * key=value parameters. Values containing ':' special characters must be
 * escaped.
 * @param search_flags flags passed to av_opt_find2. I.e. if AV_OPT_SEARCH_CHILDREN
 * is passed here, then the option may be set on a child of obj.
 *
 * @return 0 if the value has been set, or an AVERROR code in case of
 * error:
 * AVERROR_OPTION_NOT_FOUND if no matching option exists
 * AVERROR(ERANGE) if the value is out of range
 * AVERROR(EINVAL) if the value is not valid
 */
//int av_opt_set         (void *obj, const char *name, const char *val, int search_flags);
func AvOptSet(obj ffcommon.FVoidP, name, val ffcommon.FConstCharP, search_flags ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_set").Call(
		obj,
		ffcommon.UintPtrFromString(name),
		ffcommon.UintPtrFromString(val),
		uintptr(search_flags),
	)
	res = ffcommon.FInt(t)
	return
}

// int av_opt_set_int     (void *obj, const char *name, int64_t     val, int search_flags);
func AvOptSetInt(obj ffcommon.FVoidP, name ffcommon.FConstCharP, val ffcommon.FInt64T, search_flags ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_set_int").Call(
		obj,
		ffcommon.UintPtrFromString(name),
		uintptr(val),
		uintptr(search_flags),
	)
	res = ffcommon.FInt(t)
	return
}

// int av_opt_set_double  (void *obj, const char *name, double      val, int search_flags);
func AvOptSetDouble(obj ffcommon.FVoidP, name ffcommon.FConstCharP, val ffcommon.FDouble, search_flags ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_set_double").Call(
		obj,
		ffcommon.UintPtrFromString(name),
		uintptr(unsafe.Pointer(&val)),
		uintptr(search_flags),
	)
	res = ffcommon.FInt(t)
	return
}

// int av_opt_set_q       (void *obj, const char *name, AVRational  val, int search_flags);
func AvOptSetQ(obj ffcommon.FVoidP, name ffcommon.FConstCharP, val AVRational, search_flags ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_set_q").Call(
		obj,
		ffcommon.UintPtrFromString(name),
		*(*uintptr)(unsafe.Pointer(&val)),
		uintptr(search_flags),
	)
	res = ffcommon.FInt(t)
	return
}

// int av_opt_set_bin     (void *obj, const char *name, const uint8_t *val, int size, int search_flags);
func AvOptSetBin(obj ffcommon.FVoidP, name ffcommon.FConstCharP, val *ffcommon.FUint8T, size ffcommon.FInt, search_flags ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_set_bin").Call(
		obj,
		ffcommon.UintPtrFromString(name),
		uintptr(unsafe.Pointer(val)),
		uintptr(size),
		uintptr(search_flags),
	)
	res = ffcommon.FInt(t)
	return
}

// int av_opt_set_image_size(void *obj, const char *name, int w, int h, int search_flags);
func AvOptSetImageSize(obj ffcommon.FVoidP, name ffcommon.FConstCharP, w, h ffcommon.FInt, search_flags ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_set_image_size").Call(
		obj,
		ffcommon.UintPtrFromString(name),
		uintptr(w),
		uintptr(h),
		uintptr(search_flags),
	)
	res = ffcommon.FInt(t)
	return
}

// int av_opt_set_pixel_fmt (void *obj, const char *name, enum AVPixelFormat fmt, int search_flags);
func AvOptSetPixelFmt(obj ffcommon.FVoidP, name ffcommon.FConstCharP, fmt0 AVPixelFormat, search_flags ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_set_pixel_fmt").Call(
		obj,
		ffcommon.UintPtrFromString(name),
		uintptr(fmt0),
		uintptr(search_flags),
	)
	res = ffcommon.FInt(t)
	return
}

// int av_opt_set_sample_fmt(void *obj, const char *name, enum AVSampleFormat fmt, int search_flags);
func AvOptSetSampleFmt(obj ffcommon.FVoidP, name ffcommon.FConstCharP, fmt0 AVSampleFormat, search_flags ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_set_sample_fmt").Call(
		obj,
		ffcommon.UintPtrFromString(name),
		uintptr(fmt0),
		uintptr(search_flags),
	)
	res = ffcommon.FInt(t)
	return
}

// int av_opt_set_video_rate(void *obj, const char *name, AVRational val, int search_flags);
func AvOptSetVideoRate(obj ffcommon.FVoidP, name ffcommon.FConstCharP, val AVRational, search_flags ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_set_video_rate").Call(
		obj,
		ffcommon.UintPtrFromString(name),
		uintptr(unsafe.Pointer(&val)),
		uintptr(search_flags),
	)
	res = ffcommon.FInt(t)
	return
}

// int av_opt_set_channel_layout(void *obj, const char *name, int64_t ch_layout, int search_flags);
func AvOptSetChannelLayout(obj ffcommon.FVoidP, name ffcommon.FConstCharP, ch_layout ffcommon.FInt64T, search_flags ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_set_channel_layout").Call(
		obj,
		ffcommon.UintPtrFromString(name),
		uintptr(ch_layout),
		uintptr(search_flags),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * @note Any old dictionary present is discarded and replaced with a copy of the new one. The
 * caller still owns val is and responsible for freeing it.
 */
//int av_opt_set_dict_val(void *obj, const char *name, const AVDictionary *val, int search_flags);
func AvOptSetDictVal(obj ffcommon.FVoidP, name ffcommon.FConstCharP, val *AVDictionary, search_flags ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_set_dict_val").Call(
		obj,
		ffcommon.UintPtrFromString(name),
		uintptr(unsafe.Pointer(val)),
		uintptr(search_flags),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Set a binary option to an integer list.
 *
 * @param obj    AVClass object to set options on
 * @param name   name of the binary option
 * @param val    pointer to an integer list (must have the correct type with
 *               regard to the contents of the list)
 * @param term   list terminator (usually 0 or -1)
 * @param flags  search flags
 */
//#define av_opt_set_int_list(obj, name, val, term, flags) \
//(av_int_list_length(val, term) > INT_MAX / sizeof(*(val)) ? \
//AVERROR(EINVAL) : \
//av_opt_set_bin(obj, name, (const uint8_t *)(val), \
//av_int_list_length(val, term) * sizeof(*(val)), flags))
//todo可能有问题，暂时别用
func AvOptSetIntList(obj ffcommon.FVoidP, name ffcommon.FConstCharP, val uintptr, size ffcommon.FInt /*sizeof(*(val))*/, term ffcommon.FUint64T, flags ffcommon.FInt) (res ffcommon.FInt) {
	//a := AvIntListLength(val, term)
	a := AvIntListLengthForSize(uint32(size), val, term)
	if a > uint32(math.MaxInt32/size) {
		res = -EINVAL
	} else {
		res = AvOptSetBin(obj, name, (*byte)(unsafe.Pointer(val)), size, flags)
	}
	return
}

/**
 * @}
 */

/**
 * @defgroup opt_get_funcs Option getting functions
 * @{
 * Those functions get a value of the option with the given name from an object.
 *
 * @param[in] obj a struct whose first element is a pointer to an AVClass.
 * @param[in] name name of the option to get.
 * @param[in] search_flags flags passed to av_opt_find2. I.e. if AV_OPT_SEARCH_CHILDREN
 * is passed here, then the option may be found in a child of obj.
 * @param[out] out_val value of the option will be written here
 * @return >=0 on success, a negative error code otherwise
 */
/**
 * @note the returned string will be av_malloc()ed and must be av_free()ed by the caller
 *
 * @note if AV_OPT_ALLOW_NULL is set in search_flags in av_opt_get, and the
 * option is of type AV_OPT_TYPE_STRING, AV_OPT_TYPE_BINARY or AV_OPT_TYPE_DICT
 * and is set to NULL, *out_val will be set to NULL instead of an allocated
 * empty string.
 */
//int av_opt_get         (void *obj, const char *name, int search_flags, uint8_t   **out_val);
func AvOptGet(obj ffcommon.FVoidP, name ffcommon.FConstCharP, search_flags ffcommon.FInt, out_val **ffcommon.FUint8T) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_get").Call(
		obj,
		ffcommon.UintPtrFromString(name),
		uintptr(search_flags),
		uintptr(unsafe.Pointer(out_val)),
	)
	res = ffcommon.FInt(t)
	return
}

// int av_opt_get_int     (void *obj, const char *name, int search_flags, int64_t    *out_val);
func AvOptGetInt(obj ffcommon.FVoidP, name ffcommon.FConstCharP, search_flags ffcommon.FInt, out_val *ffcommon.FInt64T) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_get_int").Call(
		obj,
		ffcommon.UintPtrFromString(name),
		uintptr(search_flags),
		uintptr(unsafe.Pointer(out_val)),
	)
	res = ffcommon.FInt(t)
	return
}

// int av_opt_get_double  (void *obj, const char *name, int search_flags, double     *out_val);
func av_opt_get_double(obj ffcommon.FVoidP, name ffcommon.FConstCharP, search_flags ffcommon.FInt, out_val *ffcommon.FDouble) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_get_double").Call(
		obj,
		ffcommon.UintPtrFromString(name),
		uintptr(search_flags),
		uintptr(unsafe.Pointer(out_val)),
	)
	res = ffcommon.FInt(t)
	return
}

// int av_opt_get_q       (void *obj, const char *name, int search_flags, AVRational *out_val);
func AvOptGetQ(obj ffcommon.FVoidP, name ffcommon.FConstCharP, search_flags ffcommon.FInt, out_val *AVRational) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_get_q").Call(
		obj,
		ffcommon.UintPtrFromString(name),
		uintptr(search_flags),
		uintptr(unsafe.Pointer(out_val)),
	)
	res = ffcommon.FInt(t)
	return
}

// int av_opt_get_image_size(void *obj, const char *name, int search_flags, int *w_out, int *h_out);
func AvOptGetImageSize(obj ffcommon.FVoidP, name ffcommon.FConstCharP, search_flags ffcommon.FInt, w_out, h_out *ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_get_image_size").Call(
		obj,
		ffcommon.UintPtrFromString(name),
		uintptr(search_flags),
		uintptr(unsafe.Pointer(w_out)),
		uintptr(unsafe.Pointer(h_out)),
	)
	res = ffcommon.FInt(t)
	return
}

// int av_opt_get_pixel_fmt (void *obj, const char *name, int search_flags, enum AVPixelFormat *out_fmt);
func AvOptGetPixelFmt(obj ffcommon.FVoidP, name ffcommon.FConstCharP, search_flags ffcommon.FInt, out_fmt *AVPixelFormat) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_get_pixel_fmt").Call(
		obj,
		ffcommon.UintPtrFromString(name),
		uintptr(search_flags),
		uintptr(unsafe.Pointer(out_fmt)),
	)
	res = ffcommon.FInt(t)
	return
}

// int av_opt_get_sample_fmt(void *obj, const char *name, int search_flags, enum AVSampleFormat *out_fmt);
func AvOptGetSampleFmt(obj ffcommon.FVoidP, name ffcommon.FConstCharP, search_flags ffcommon.FInt, out_fmt *AVSampleFormat) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_get_sample_fmt").Call(
		obj,
		ffcommon.UintPtrFromString(name),
		uintptr(search_flags),
		uintptr(unsafe.Pointer(out_fmt)),
	)
	res = ffcommon.FInt(t)
	return
}

// int av_opt_get_video_rate(void *obj, const char *name, int search_flags, AVRational *out_val);
func AvOptGetVideoRate(obj ffcommon.FVoidP, name ffcommon.FConstCharP, search_flags ffcommon.FInt, out_fmt *AVRational) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_get_video_rate").Call(
		obj,
		ffcommon.UintPtrFromString(name),
		uintptr(search_flags),
		uintptr(unsafe.Pointer(out_fmt)),
	)
	res = ffcommon.FInt(t)
	return
}

// int av_opt_get_channel_layout(void *obj, const char *name, int search_flags, int64_t *ch_layout);
func AvOptGetChannelLayout(obj ffcommon.FVoidP, name ffcommon.FConstCharP, search_flags ffcommon.FInt, ch_layout *ffcommon.FInt64T) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_get_channel_layout").Call(
		obj,
		ffcommon.UintPtrFromString(name),
		uintptr(search_flags),
		uintptr(unsafe.Pointer(ch_layout)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * @param[out] out_val The returned dictionary is a copy of the actual value and must
 * be freed with av_dict_free() by the caller
 */
//int av_opt_get_dict_val(void *obj, const char *name, int search_flags, AVDictionary **out_val);
func AvOptGetDictVal(obj ffcommon.FVoidP, name ffcommon.FConstCharP, search_flags ffcommon.FInt, out_val **AVDictionary) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_get_dict_val").Call(
		obj,
		ffcommon.UintPtrFromString(name),
		uintptr(search_flags),
		uintptr(unsafe.Pointer(out_val)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * @}
 */
/**
 * Gets a pointer to the requested field in a struct.
 * This function allows accessing a struct even when its fields are moved or
 * renamed since the application making the access has been compiled,
 *
 * @returns a pointer to the field, it can be cast to the correct type and read
 *          or written to.
 */
//void *av_opt_ptr(const AVClass *avclass, void *obj, const char *name);
func (avclass *AVClass) AvOptPtr(obj ffcommon.FVoidP, name ffcommon.FConstCharP) (res ffcommon.FVoidP) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_ptr").Call(
		uintptr(unsafe.Pointer(avclass)),
		obj,
		ffcommon.UintPtrFromString(name),
	)
	res = t
	return
}

/**
 * Free an AVOptionRanges struct and set it to NULL.
 */
//void av_opt_freep_ranges(AVOptionRanges **ranges);
func AvOptFreepRanges(ranges **AVOptionRanges) {
	ffcommon.GetAvutilDll().NewProc("av_opt_freep_ranges").Call(
		uintptr(unsafe.Pointer(ranges)),
	)
}

/**
 * Get a list of allowed ranges for the given option.
 *
 * The returned list may depend on other fields in obj like for example profile.
 *
 * @param flags is a bitmask of flags, undefined flags should not be set and should be ignored
 *              AV_OPT_SEARCH_FAKE_OBJ indicates that the obj is a double pointer to a AVClass instead of a full instance
 *              AV_OPT_MULTI_COMPONENT_RANGE indicates that function may return more than one component, @see AVOptionRanges
 *
 * The result must be freed with av_opt_freep_ranges.
 *
 * @return number of compontents returned on success, a negative errro code otherwise
 */
//int av_opt_query_ranges(AVOptionRanges **, void *obj, const char *key, int flags);
func AvOptQueryRanges(a **AVOptionRanges, obj ffcommon.FVoidP, key ffcommon.FConstCharP, flags ffcommon.FInt) (res ffcommon.FInt64T) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_query_ranges").Call(
		uintptr(unsafe.Pointer(a)),
		obj,
		ffcommon.UintPtrFromString(key),
		uintptr(flags),
	)
	res = ffcommon.FInt64T(t)
	return
}

/**
 * Copy options from src object into dest object.
 *
 * Options that require memory allocation (e.g. string or binary) are malloc'ed in dest object.
 * Original memory allocated for such options is freed unless both src and dest options points to the same memory.
 *
 * @param dest Object to copy from
 * @param src  Object to copy into
 * @return 0 on success, negative on error
 */
//int av_opt_copy(void *dest, const void *src);
func AvOptCopy(dest ffcommon.FVoidP, src ffcommon.FConstVoidP) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_copy").Call()
	res = ffcommon.FInt(t)
	return
}

/**
 * Get a default list of allowed ranges for the given option.
 *
 * This list is constructed without using the AVClass.query_ranges() callback
 * and can be used as fallback from within the callback.
 *
 * @param flags is a bitmask of flags, undefined flags should not be set and should be ignored
 *              AV_OPT_SEARCH_FAKE_OBJ indicates that the obj is a double pointer to a AVClass instead of a full instance
 *              AV_OPT_MULTI_COMPONENT_RANGE indicates that function may return more than one component, @see AVOptionRanges
 *
 * The result must be freed with av_opt_free_ranges.
 *
 * @return number of compontents returned on success, a negative errro code otherwise
 */
//int av_opt_query_ranges_default(AVOptionRanges **, void *obj, const char *key, int flags);
func AvOptQueryRangesDefault(a **AVOptionRanges, obj ffcommon.FVoidP, key ffcommon.FConstCharP, flags ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_query_ranges_default").Call(
		uintptr(unsafe.Pointer(a)),
		obj,
		ffcommon.UintPtrFromString(key),
		uintptr(flags),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Check if given option is set to its default value.
 *
 * Options o must belong to the obj. This function must not be called to check child's options state.
 * @see av_opt_is_set_to_default_by_name().
 *
 * @param obj  AVClass object to check option on
 * @param o    option to be checked
 * @return     >0 when option is set to its default,
 *              0 when option is not set its default,
 *             <0 on error
 */
//int av_opt_is_set_to_default(void *obj, const AVOption *o);
func AvOptIsSetToDefault(obj ffcommon.FVoidP, o *AVOption) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_is_set_to_default").Call(
		obj,
		uintptr(unsafe.Pointer(o)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Check if given option is set to its default value.
 *
 * @param obj          AVClass object to check option on
 * @param name         option name
 * @param search_flags combination of AV_OPT_SEARCH_*
 * @return             >0 when option is set to its default,
 *                     0 when option is not set its default,
 *                     <0 on error
 */
//int av_opt_is_set_to_default_by_name(void *obj, const char *name, int search_flags);
func AvOptIsSetToDefaultByName(obj ffcommon.FVoidP, name ffcommon.FConstCharP, search_flags ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_is_set_to_default_by_name").Call(
		obj,
		ffcommon.UintPtrFromString(name),
		uintptr(search_flags),
	)
	res = ffcommon.FInt(t)
	return
}

const AV_OPT_SERIALIZE_SKIP_DEFAULTS = 0x00000001   ///< Serialize options that are not set to default values only.
const AV_OPT_SERIALIZE_OPT_FLAGS_EXACT = 0x00000002 ///< Serialize options that exactly match opt_flags only.

/**
 * Serialize object's options.
 *
 * Create a string containing object's serialized options.
 * Such string may be passed back to av_opt_set_from_string() in order to restore option values.
 * A key/value or pairs separator occurring in the serialized value or
 * name string are escaped through the av_escape() function.
 *
 * @param[in]  obj           AVClass object to serialize
 * @param[in]  opt_flags     serialize options with all the specified flags set (AV_OPT_FLAG)
 * @param[in]  flags         combination of AV_OPT_SERIALIZE_* flags
 * @param[out] buffer        Pointer to buffer that will be allocated with string containg serialized options.
 *                           Buffer must be freed by the caller when is no longer needed.
 * @param[in]  key_val_sep   character used to separate key from value
 * @param[in]  pairs_sep     character used to separate two pairs from each other
 * @return                   >= 0 on success, negative on error
 * @warning Separators cannot be neither '\\' nor '\0'. They also cannot be the same.
 */
//int av_opt_serialize(void *obj, int opt_flags, int flags, char **buffer,
//const char key_val_sep, const char pairs_sep);
func AvOptSerialize(obj ffcommon.FVoidP, opt_flags, flags ffcommon.FInt, buffer *ffcommon.FBuf,
	key_val_sep, pairs_sep ffcommon.FChar) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_opt_serialize").Call(
		obj,
		uintptr(opt_flags),
		uintptr(flags),
		uintptr(unsafe.Pointer(buffer)),
		uintptr(key_val_sep),
		uintptr(pairs_sep),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * @}
 */

//#endif /* AVUTIL_OPT_H */
