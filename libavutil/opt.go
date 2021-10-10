package libavutil

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"syscall"
	"unsafe"
)

/**
 * AVOption
 */
type AVOption struct {
	Name ffcommon.FBuf

	/**
	 * short English help text
	 * @todo What about other languages?
	 */
	Help ffcommon.FBuf

	/**
	 * The offset relative to the context structure where the option
	 * value is stored. It should be 0 for named ffconstants.
	 */
	Offset ffcommon.FInt
	Type0  ffconstant.AVOptionType

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
	I64 ffcommon.FInt64T
	Min ffcommon.FDouble ///< minimum valid value for the option
	Max ffcommon.FDouble ///< maximum valid value for the option

	Flags ffcommon.FInt

	/**
	 * The logical unit to which the option belongs. Non-ffconstant
	 * options and corresponding named ffconstants share the same
	 * unit. May be NULL.
	 */
	Unit ffcommon.FBuf
}

/**
 * A single allowed range of values, or a single allowed value.
 */
type AVOptionRange struct {
	Str ffcommon.FBuf
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
//未测试
func AvOptShow2(obj ffcommon.FVoidP, av_log_obj ffcommon.FVoidP, req_flags ffcommon.FInt, rej_flags ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_show2").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(av_log_obj)),
		uintptr(req_flags),
		uintptr(rej_flags),
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
 * Set the values of all AVOption fields to their default values.
 *
 * @param s an AVOption-enabled struct (its first member must be a pointer to AVClass)
 */
//void av_opt_set_defaults(void *s);
//未测试
func AvOptSetDefaults(s ffcommon.FVoidP) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_set_defaults").Call(
		uintptr(unsafe.Pointer(s)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
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
//未测试
func AvOptSetDefaults2(s ffcommon.FVoidP, mask ffcommon.FInt, flags ffcommon.FInt) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_set_defaults2").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(mask),
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
//未测试
func AvSetOptionsString(ctx ffcommon.FVoidP, opts ffcommon.FCharP,
	key_val_sep ffcommon.FCharP, pairs_sep ffcommon.FCharP) (res ffcommon.FInt, err error) {
	var t uintptr
	var optsp *byte
	optsp, err = syscall.BytePtrFromString(opts)
	if err != nil {
		return
	}
	var key_val_sepp *byte
	key_val_sepp, err = syscall.BytePtrFromString(key_val_sep)
	if err != nil {
		return
	}
	var pairs_sepp *byte
	pairs_sepp, err = syscall.BytePtrFromString(pairs_sep)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_set_options_string").Call(
		uintptr(ctx),
		uintptr(unsafe.Pointer(optsp)),
		uintptr(unsafe.Pointer(key_val_sepp)),
		uintptr(unsafe.Pointer(pairs_sepp)),
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
//未测试
func AvOptSetFromString(ctx ffcommon.FVoidP, opts ffcommon.FCharP,
	shorthand *ffcommon.FBuf,
	key_val_sep ffcommon.FCharP, pairs_sep ffcommon.FCharP) (res ffcommon.FInt, err error) {
	var t uintptr
	var optsp *byte
	optsp, err = syscall.BytePtrFromString(opts)
	if err != nil {
		return
	}
	var key_val_sepp *byte
	key_val_sepp, err = syscall.BytePtrFromString(key_val_sep)
	if err != nil {
		return
	}
	var pairs_sepp *byte
	pairs_sepp, err = syscall.BytePtrFromString(pairs_sep)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_set_from_string").Call(
		uintptr(ctx),
		uintptr(unsafe.Pointer(optsp)),
		uintptr(unsafe.Pointer(&shorthand)),
		uintptr(unsafe.Pointer(key_val_sepp)),
		uintptr(unsafe.Pointer(pairs_sepp)),
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
 * Free all allocated objects in obj.
 */
//void av_opt_free(void *obj);
//未测试
func AvOptFree(obj ffcommon.FVoidP) (err error) {
	var t uintptr

	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_free").Call(
		uintptr(unsafe.Pointer(obj)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
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
//未测试
func AvOptFlagIsSet(ctx ffcommon.FVoidP, field_name ffcommon.FCharP, flag_name ffcommon.FCharP) (res ffcommon.FInt, err error) {
	var t uintptr
	var field_namep *byte
	field_namep, err = syscall.BytePtrFromString(field_name)
	if err != nil {
		return
	}
	var flag_namep *byte
	flag_namep, err = syscall.BytePtrFromString(flag_name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_flag_is_set").Call(
		uintptr(ctx),
		uintptr(unsafe.Pointer(field_namep)),
		uintptr(unsafe.Pointer(&flag_namep)),
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
//未测试
func AvOptSetDict(obj ffcommon.FVoidP, options **AVDictionary) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_set_dict").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&options)),
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
//未测试
func AvOptSetDict2(obj ffcommon.FVoidP, options **AVDictionary, search_flags ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_set_dict2").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&options)),
		uintptr(search_flags),
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
//未测试
func AvOptGetKeyValue(ropts *ffcommon.FBuf,
	key_val_sep ffcommon.FConstCharP, pairs_sep ffcommon.FConstCharP,
	flags ffcommon.FUnsigned,
	rkey *ffcommon.FBuf, rval *ffcommon.FBuf) (res ffcommon.FInt, err error) {
	var t uintptr
	var key_val_sepp *byte
	key_val_sepp, err = syscall.BytePtrFromString(key_val_sep)
	if err != nil {
		return
	}
	var pairs_sepp *byte
	pairs_sepp, err = syscall.BytePtrFromString(pairs_sep)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_get_key_value").Call(
		uintptr(unsafe.Pointer(&ropts)),
		uintptr(unsafe.Pointer(key_val_sepp)),
		uintptr(unsafe.Pointer(pairs_sepp)),
		uintptr(flags),
		uintptr(unsafe.Pointer(&rkey)),
		uintptr(unsafe.Pointer(&rval)),
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
//未测试
func AvOptEvalFlags(obj ffcommon.FVoidP, o *AVOption, val ffcommon.FConstCharP, flags_out *ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	var valp *byte
	valp, err = syscall.BytePtrFromString(val)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_eval_flags").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(o)),
		uintptr(unsafe.Pointer(valp)),
		uintptr(unsafe.Pointer(flags_out)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//int av_opt_eval_int   (void *obj, const AVOption *o, const char *val, int        *int_out);
//未测试
func AvOptEvalInt(obj ffcommon.FVoidP, o *AVOption, val ffcommon.FConstCharP, int_out *ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	var valp *byte
	valp, err = syscall.BytePtrFromString(val)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_eval_int").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(o)),
		uintptr(unsafe.Pointer(valp)),
		uintptr(unsafe.Pointer(int_out)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//int av_opt_eval_int64 (void *obj, const AVOption *o, const char *val, int64_t    *int64_out);
//未测试
func AvOptEvalInt64(obj ffcommon.FVoidP, o *AVOption, val ffcommon.FConstCharP, int_out *ffcommon.FInt64T) (res ffcommon.FInt, err error) {
	var t uintptr
	var valp *byte
	valp, err = syscall.BytePtrFromString(val)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_eval_int64").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(o)),
		uintptr(unsafe.Pointer(valp)),
		uintptr(unsafe.Pointer(int_out)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//int av_opt_eval_float (void *obj, const AVOption *o, const char *val, float      *float_out);
//未测试
func AvOptEvalFloat(obj ffcommon.FVoidP, o *AVOption, val ffcommon.FConstCharP, float_out *ffcommon.FFloat) (res ffcommon.FInt, err error) {
	var t uintptr
	var valp *byte
	valp, err = syscall.BytePtrFromString(val)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_eval_float").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(o)),
		uintptr(unsafe.Pointer(valp)),
		uintptr(unsafe.Pointer(float_out)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//int av_opt_eval_double(void *obj, const AVOption *o, const char *val, double     *double_out);
//未测试
func AvOptEvalDouble(obj ffcommon.FVoidP, o *AVOption, val ffcommon.FConstCharP, double_out *ffcommon.FDouble) (res ffcommon.FInt, err error) {
	var t uintptr
	var valp *byte
	valp, err = syscall.BytePtrFromString(val)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_eval_double").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(o)),
		uintptr(unsafe.Pointer(valp)),
		uintptr(unsafe.Pointer(double_out)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//int av_opt_eval_q     (void *obj, const AVOption *o, const char *val, AVRational *q_out);
//未测试
func AvOptEvalQ(obj ffcommon.FVoidP, o *AVOption, val ffcommon.FConstCharP, q_out *AVRational) (res ffcommon.FInt, err error) {
	var t uintptr
	var valp *byte
	valp, err = syscall.BytePtrFromString(val)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_eval_q").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(o)),
		uintptr(unsafe.Pointer(valp)),
		uintptr(unsafe.Pointer(q_out)),
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
 * Look for an option in an object. Consider only options which
 * have all the specified flags set.
 *
 * @param[in] obj A pointer to a struct whose first element is a
 *                pointer to an AVClass.
 *                Alternatively a double pointer to an AVClass, if
 *                AV_OPT_SEARCH_FAKE_OBJ search flag is set.
 * @param[in] name The name of the option to look for.
 * @param[in] unit When searching for named ffconstants, name of the unit
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
//未测试
func AvOptFind(obj ffcommon.FVoidP, name ffcommon.FConstCharP, unit ffcommon.FConstCharP,
	opt_flags ffcommon.FInt, search_flags ffcommon.FInt) (res *AVOption, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	var unitp *byte
	unitp, err = syscall.BytePtrFromString(unit)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_find").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(namep)),
		uintptr(unsafe.Pointer(unitp)),
		uintptr(opt_flags),
		uintptr(search_flags),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
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
 * @param[in] unit When searching for named ffconstants, name of the unit
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
//未测试
func AvOptFind2(obj ffcommon.FVoidP, name ffcommon.FConstCharP, unit ffcommon.FConstCharP,
	opt_flags ffcommon.FInt, search_flags ffcommon.FInt, target_obj *ffcommon.FVoidP) (res *AVOption, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	var unitp *byte
	unitp, err = syscall.BytePtrFromString(unit)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_find2").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(namep)),
		uintptr(unsafe.Pointer(unitp)),
		uintptr(opt_flags),
		uintptr(search_flags),
		uintptr(unsafe.Pointer(target_obj)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
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
//未测试
func AvOptNext(obj ffcommon.FConstVoidP, prev *AVOption) (res *AVOption, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_next").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(prev)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
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
//未测试
func AvOptChildNext(obj ffcommon.FVoidP, prev ffcommon.FVoidP) (res ffcommon.FVoidP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_child_next").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(prev)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
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
//未测试
func AvOptChildClassNext(parent *AVClass, prev *AVClass) (res *AVClass, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_child_class_next").Call(
		uintptr(unsafe.Pointer(parent)),
		uintptr(unsafe.Pointer(prev)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
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
//未测试
func AvOptChildClassIterate(parent *AVClass, prev *AVClass) (res *AVClass, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_child_class_iterate").Call(
		uintptr(unsafe.Pointer(parent)),
		uintptr(unsafe.Pointer(prev)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
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
//未测试
func AvOptSet(obj ffcommon.FVoidP, name ffcommon.FConstCharP, val ffcommon.FConstCharP, search_flags ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	var valp *byte
	valp, err = syscall.BytePtrFromString(val)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_set").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(namep)),
		uintptr(unsafe.Pointer(valp)),
		uintptr(search_flags),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//int av_opt_set_int     (void *obj, const char *name, int64_t     val, int search_flags);
//未测试
func AvOptSetInt(obj ffcommon.FVoidP, name ffcommon.FConstCharP, val ffcommon.FInt64T, search_flags ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_set_int").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(namep)),
		uintptr(val),
		uintptr(search_flags),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//int av_opt_set_double  (void *obj, const char *name, double      val, int search_flags);
//未测试
func AvOptSetDouble(obj ffcommon.FVoidP, name ffcommon.FConstCharP, val ffcommon.FDouble, search_flags ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_set_double").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(namep)),
		uintptr(val),
		uintptr(search_flags),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//int av_opt_set_q       (void *obj, const char *name, AVRational  val, int search_flags);
//未测试
func AvOptSetQ(obj ffcommon.FVoidP, name ffcommon.FConstCharP, val AVRational, search_flags ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_set_q").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(namep)),
		uintptr(unsafe.Pointer(&val)),
		uintptr(search_flags),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//int av_opt_set_bin     (void *obj, const char *name, const uint8_t *val, int size, int search_flags);
//未测试
func AvOptSetBin(obj ffcommon.FVoidP, name ffcommon.FConstCharP, val *ffcommon.FUint8T, size ffcommon.FInt, search_flags ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_set_bin").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(namep)),
		uintptr(unsafe.Pointer(val)),
		uintptr(size),
		uintptr(search_flags),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//int av_opt_set_image_size(void *obj, const char *name, int w, int h, int search_flags);
//未测试
func AvOptSetImageSize(obj ffcommon.FVoidP, name ffcommon.FConstCharP, w ffcommon.FInt, h ffcommon.FInt, search_flags ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_set_image_size").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(namep)),
		uintptr(w),
		uintptr(h),
		uintptr(search_flags),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//int av_opt_set_pixel_fmt (void *obj, const char *name, enum AVPixelFormat fmt, int search_flags);
//未测试
func AvOptSetPixelFmt(obj ffcommon.FVoidP, name ffcommon.FConstCharP, fmt0 ffconstant.AVPixelFormat, search_flags ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_set_pixel_fmt").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(namep)),
		uintptr(fmt0),
		uintptr(search_flags),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//int av_opt_set_sample_fmt(void *obj, const char *name, enum AVSampleFormat fmt, int search_flags);
//未测试
func AvOptSetSampleFmt(obj ffcommon.FVoidP, name ffcommon.FConstCharP, fmt0 ffconstant.AVSampleFormat, search_flags ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_set_sample_fmt").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(namep)),
		uintptr(fmt0),
		uintptr(search_flags),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//int av_opt_set_video_rate(void *obj, const char *name, AVRational val, int search_flags);
//未测试
func AvOptSetVideoRate(obj ffcommon.FVoidP, name ffcommon.FConstCharP, fmt0 AVRational, search_flags ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_set_video_rate").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(namep)),
		uintptr(unsafe.Pointer(&fmt0)),
		uintptr(search_flags),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//int av_opt_set_channel_layout(void *obj, const char *name, int64_t ch_layout, int search_flags);
//未测试
func AvOptSetChannelLayout(obj ffcommon.FVoidP, name ffcommon.FConstCharP, ch_layout ffcommon.FInt64T, search_flags ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_set_channel_layout").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(namep)),
		uintptr(ch_layout),
		uintptr(search_flags),
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
 * @note Any old dictionary present is discarded and replaced with a copy of the new one. The
 * caller still owns val is and responsible for freeing it.
 */
//int av_opt_set_dict_val(void *obj, const char *name, const AVDictionary *val, int search_flags);
//未测试
func AvOptSetDictVal(obj ffcommon.FVoidP, name ffcommon.FConstCharP, val *AVDictionary, search_flags ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_set_dict_val").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(namep)),
		uintptr(unsafe.Pointer(val)),
		uintptr(search_flags),
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
//未测试
func AvOptGet(obj ffcommon.FVoidP, name ffcommon.FConstCharP, search_flags ffcommon.FInt, out_val **ffcommon.FUint8T) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_get").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(namep)),
		uintptr(search_flags),
		uintptr(unsafe.Pointer(&out_val)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//int av_opt_get_int     (void *obj, const char *name, int search_flags, int64_t    *out_val);
//未测试
func AvOptGetInt(obj ffcommon.FVoidP, name ffcommon.FConstCharP, search_flags ffcommon.FInt, out_val *ffcommon.FInt64T) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_get_int").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(namep)),
		uintptr(search_flags),
		uintptr(unsafe.Pointer(out_val)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//int av_opt_get_double  (void *obj, const char *name, int search_flags, double     *out_val);
//未测试
func AvOptGetDouble(obj ffcommon.FVoidP, name ffcommon.FConstCharP, search_flags ffcommon.FInt, out_val *ffcommon.FDouble) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_get_double").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(namep)),
		uintptr(search_flags),
		uintptr(unsafe.Pointer(out_val)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//int av_opt_get_q       (void *obj, const char *name, int search_flags, AVRational *out_val);
//未测试
func AvOptGetQ(obj ffcommon.FVoidP, name ffcommon.FConstCharP, search_flags ffcommon.FInt, out_val *AVRational) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_get_q").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(namep)),
		uintptr(search_flags),
		uintptr(unsafe.Pointer(out_val)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//int av_opt_get_image_size(void *obj, const char *name, int search_flags, int *w_out, int *h_out);
//未测试
func AvOptGetImageSize(obj ffcommon.FVoidP, name ffcommon.FConstCharP, search_flags ffcommon.FInt, w_out *ffcommon.FInt, h_out *ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_get_image_size").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(namep)),
		uintptr(search_flags),
		uintptr(unsafe.Pointer(w_out)),
		uintptr(unsafe.Pointer(h_out)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//int av_opt_get_pixel_fmt (void *obj, const char *name, int search_flags, enum AVPixelFormat *out_fmt);
//未测试
func AvOptGetPixelFmt(obj ffcommon.FVoidP, name ffcommon.FConstCharP, search_flags ffcommon.FInt, out_fmt *ffconstant.AVPixelFormat) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_get_pixel_fmt").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(namep)),
		uintptr(search_flags),
		uintptr(unsafe.Pointer(out_fmt)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//int av_opt_get_sample_fmt(void *obj, const char *name, int search_flags, enum AVSampleFormat *out_fmt);
//未测试
func AvOptGetSampleFmt(obj ffcommon.FVoidP, name ffcommon.FConstCharP, search_flags ffcommon.FInt, out_fmt *ffconstant.AVPixelFormat) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_get_sample_fmt").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(namep)),
		uintptr(search_flags),
		uintptr(unsafe.Pointer(out_fmt)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//int av_opt_get_video_rate(void *obj, const char *name, int search_flags, AVRational *out_val);
//未测试
func AvOptGetVideoRate(obj ffcommon.FVoidP, name ffcommon.FConstCharP, search_flags ffcommon.FInt, out_fmt *AVRational) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_get_video_rate").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(namep)),
		uintptr(search_flags),
		uintptr(unsafe.Pointer(out_fmt)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//int av_opt_get_channel_layout(void *obj, const char *name, int search_flags, int64_t *ch_layout);
//未测试
func AvOptGetChannelLayout(obj ffcommon.FVoidP, name ffcommon.FConstCharP, search_flags ffcommon.FInt, ch_layout ffcommon.FInt64T) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_get_channel_layout").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(namep)),
		uintptr(search_flags),
		uintptr(ch_layout),
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
 * @param[out] out_val The returned dictionary is a copy of the actual value and must
 * be freed with av_dict_free() by the caller
 */
//int av_opt_get_dict_val(void *obj, const char *name, int search_flags, AVDictionary **out_val);
//未测试
func AvOptGetDictVal(obj ffcommon.FVoidP, name ffcommon.FConstCharP, search_flags ffcommon.FInt, out_val **AVDictionary) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_get_dict_val").Call(
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(namep)),
		uintptr(search_flags),
		uintptr(unsafe.Pointer(&out_val)),
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
//未测试
func AvOptPtr(avclass *AVClass, obj ffcommon.FVoidP, name ffcommon.FConstCharP) (res ffcommon.FVoidP, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_ptr").Call(
		uintptr(unsafe.Pointer(avclass)),
		obj,
		uintptr(unsafe.Pointer(namep)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FVoidP(t)
	return
}

/**
 * Free an AVOptionRanges struct and set it to NULL.
 */
//void av_opt_freep_ranges(AVOptionRanges **ranges);
//未测试
func AvOptFreepRanges(ranges **AVOptionRanges) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_freep_ranges").Call(
		uintptr(unsafe.Pointer(&ranges)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
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
//未测试
func AvOptQueryRanges(a **AVOptionRanges, obj ffcommon.FVoidP, key ffcommon.FConstCharP, flags ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	var keyp *byte
	keyp, err = syscall.BytePtrFromString(key)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_query_ranges").Call(
		uintptr(unsafe.Pointer(&a)),
		obj,
		uintptr(unsafe.Pointer(keyp)),
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
//未测试
func AvOptCopy(dest ffcommon.FVoidP, src ffcommon.FVoidP) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_copy").Call(
		dest,
		src,
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
//未测试
func AvOptQueryRangesDefault(a **AVOptionRanges, obj ffcommon.FVoidP, key ffcommon.FConstCharP, flags ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	var keyp *byte
	keyp, err = syscall.BytePtrFromString(key)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_query_ranges_default").Call(
		uintptr(unsafe.Pointer(&a)),
		obj,
		uintptr(unsafe.Pointer(keyp)),
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
//未测试
func AvOptIsSetToDefault(obj ffcommon.FVoidP, o *AVOption) (res ffcommon.FInt, err error) {
	var t uintptr

	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_is_set_to_default").Call(
		obj,
		uintptr(unsafe.Pointer(o)),
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
//未测试
func AvOptIsSetToDefaultByName(obj ffcommon.FVoidP, name ffcommon.FConstCharP, search_flags ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_is_set_to_default_by_name").Call(
		obj,
		uintptr(unsafe.Pointer(namep)),
		uintptr(search_flags),
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
//未测试
func AvOptSerialize(obj ffcommon.FVoidP, opt_flags ffcommon.FInt, flags ffcommon.FInt, buffer *ffcommon.FBuf,
	key_val_sep ffcommon.FUint8T, pairs_sep ffcommon.FUint8T) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_opt_serialize").Call(
		obj,
		uintptr(opt_flags),
		uintptr(flags),
		uintptr(unsafe.Pointer(&buffer)),
		uintptr(key_val_sep),
		uintptr(pairs_sep),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}
