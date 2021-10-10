package libavutil

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"syscall"
	"unsafe"
)

//type AVOptionRanges struct {
//}

/**
 * Describe the class of an AVClass context structure. That is an
 * arbitrary struct of which the first field is a pointer to an
 * AVClass struct (e.g. AVCodecContext, AVFormatContext etc.).
 */
type AVClass struct {
	/**
	 * The name of the class; usually it is the same name as the
	 * context structure type to which the AVClass is associated.
	 */
	//const char* class_name;
	class_name *ffcommon.FUint8T
	/**
	 * A pointer to a function which returns the name of a context
	 * instance ctx associated with the class.
	 */
	//const char* (*item_name)(void* ctx);
	item_name func(ctx ffcommon.FVoidP)

	/**
	 * a pointer to the first option specified in the class if any or NULL
	 *
	 * @see av_set_default_options()
	 */
	//const struct AVOption *option;
	option *AVOption

	/**
	 * LIBAVUTIL_VERSION with which this structure was created.
	 * This is used to allow fields to be added without requiring major
	 * version bumps everywhere.
	 */

	version ffcommon.FInt

	/**
	 * Offset in the structure where log_level_offset is stored.
	 * 0 means there is no such variable
	 */
	log_level_offset_offset ffcommon.FInt

	/**
	 * Offset in the structure where a pointer to the parent context for
	 * logging is stored. For example a decoder could pass its AVCodecContext
	 * to eval as such a parent context, which an av_log() implementation
	 * could then leverage to display the parent context.
	 * The offset can be NULL.
	 */
	parent_log_context_offset ffcommon.FInt

	/**
	 * Return next AVOptions-enabled child or NULL
	 */
	//void* (*child_next)(void *obj, void *prev);
	child_next func(obj ffcommon.FVoidP, prev ffcommon.FVoidP) (ans ffcommon.FVoidP)

	//#if FF_API_CHILD_CLASS_NEXT
	///**
	// * Return an AVClass corresponding to the next potential
	// * AVOptions-enabled child.
	// *
	// * The difference between child_next and this is that
	// * child_next iterates over _already existing_ objects, while
	// * child_class_next iterates over _all possible_ children.
	// */
	//attribute_deprecated
	//const struct AVClass* (*child_class_next)(const struct AVClass *prev);
	child_class_next func(prev *AVClass) (res *AVClass)
	//#endif

	/**
	 * Category used for visualization (like color)
	 * This is only set if the category is equal for all objects using this class.
	 * available since version (51 << 16 | 56 << 8 | 100)
	 */
	category ffconstant.AVClassCategory

	/**
	 * Callback to return the category.
	 * available since version (51 << 16 | 59 << 8 | 100)
	 */
	//AVClassCategory (*get_category)(void* ctx);
	get_category func(ctx ffcommon.FVoidP) ffconstant.AVClassCategory

	/**
	 * Callback to return the supported/allowed ranges.
	 * available since version (52.12)
	 */
	//int (*query_ranges)(struct AVOptionRanges **, void *obj, const char *key, int flags);
	query_ranges func(a **AVOptionRanges, obj ffcommon.FVoidP, key ffcommon.FBuf, flags ffcommon.FInt) ffcommon.FInt
	/**
	 * Iterate over the AVClasses corresponding to potential AVOptions-enabled
	 * children.
	 *
	 * @param iter pointer to opaque iteration state. The caller must initialize
	 *             *iter to NULL before the first call.
	 * @return AVClass for the next AVOptions-enabled child or NULL if there are
	 *         no more such children.
	 *
	 * @note The difference between child_next and this is that child_next
	 *       iterates over _already existing_ objects, while child_class_iterate
	 *       iterates over _all possible_ children.
	 */
	//const struct AVClass* (*child_class_iterate)(void **iter);
	child_class_iterate func(iter *ffcommon.FVoidP) *AVClass
}

/**
 * Send the specified message to the log if the level is less than or equal
 * to the current av_log_level. By default, all logging messages are sent to
 * stderr. This behavior can be altered by setting a different logging callback
 * function.
 * @see av_log_set_callback
 *
 * @param avcl A pointer to an arbitrary struct of which the first field is a
 *        pointer to an AVClass struct or NULL if general log.
 * @param level The importance level of the message expressed using a @ref
 *        lavu_log_ffconstants "Logging ffconstant".
 * @param fmt The format string (printf-compatible) that specifies how
 *        subsequent arguments are converted to output.
 */
//void av_log(void *avcl, int level, const char *fmt, ...) av_printf_format(3, 4);
//未测试
func AvLog(avcl ffcommon.FVoidP, level ffcommon.FInt, fmt0 ffcommon.FConstCharP, output ...[]uintptr) (res ffcommon.FInt, err error) {
	var t uintptr
	var fmt0p *byte
	fmt0p, err = syscall.BytePtrFromString(fmt0)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_log").Call(
		avcl,
		uintptr(level),
		uintptr(unsafe.Pointer(fmt0p)),
		uintptr(unsafe.Pointer(&output)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Send the specified message to the log once with the initial_level and then with
 * the subsequent_level. By default, all logging messages are sent to
 * stderr. This behavior can be altered by setting a different logging callback
 * function.
 * @see av_log
 *
 * @param avcl A pointer to an arbitrary struct of which the first field is a
 *        pointer to an AVClass struct or NULL if general log.
 * @param initial_level importance level of the message expressed using a @ref
 *        lavu_log_ffconstants "Logging ffconstant" for the first occurance.
 * @param subsequent_level importance level of the message expressed using a @ref
 *        lavu_log_ffconstants "Logging ffconstant" after the first occurance.
 * @param fmt The format string (printf-compatible) that specifies how
 *        subsequent arguments are converted to output.
 * @param state a variable to keep trak of if a message has already been printed
 *        this must be initialized to 0 before the first use. The same state
 *        must not be accessed by 2 Threads simultaneously.
 */
//void av_log_once(void* avcl, int initial_level, int subsequent_level, int *state, const char *fmt, ...) av_printf_format(5, 6);
//未测试
func AvLogOnce(avcl ffcommon.FVoidP, level ffcommon.FInt, subsequent_level ffcommon.FInt, state *ffcommon.FInt, fmt0 ffcommon.FConstCharP, output ...[]uintptr) (res ffcommon.FInt, err error) {
	var t uintptr
	var fmt0p *byte
	fmt0p, err = syscall.BytePtrFromString(fmt0)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_log_once").Call(
		avcl,
		uintptr(level),
		uintptr(subsequent_level),
		uintptr(unsafe.Pointer(state)),
		uintptr(unsafe.Pointer(fmt0p)),
		uintptr(unsafe.Pointer(&output)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Send the specified message to the log if the level is less than or equal
 * to the current av_log_level. By default, all logging messages are sent to
 * stderr. This behavior can be altered by setting a different logging callback
 * function.
 * @see av_log_set_callback
 *
 * @param avcl A pointer to an arbitrary struct of which the first field is a
 *        pointer to an AVClass struct.
 * @param level The importance level of the message expressed using a @ref
 *        lavu_log_ffconstants "Logging ffconstant".
 * @param fmt The format string (printf-compatible) that specifies how
 *        subsequent arguments are converted to output.
 * @param vl The arguments referenced by the format string.
 */
//void av_vlog(void *avcl, int level, const char *fmt, va_list vl);
//未测试
func AvVlog(avcl ffcommon.FVoidP, level ffcommon.FInt, fmt0 ffcommon.FConstCharP, vl ffcommon.FBuf, output ...[]uintptr) (err error) {
	var t uintptr
	var fmt0p *byte
	fmt0p, err = syscall.BytePtrFromString(fmt0)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_vlog").Call(
		avcl,
		uintptr(level),
		uintptr(unsafe.Pointer(fmt0p)),
		uintptr(unsafe.Pointer(&output)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Get the current log level
 *
 * @see lavu_log_ffconstants
 *
 * @return Current log level
 */
//int av_log_get_level(void);
//未测试
func AvLogGetLevel() (res ffcommon.FInt, err error) {
	var t uintptr

	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_log_get_level").Call()
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Set the log level
 *
 * @see lavu_log_ffconstants
 *
 * @param level Logging level
 */
//void av_log_set_level(int level);
//未测试
func AvLogSetLevel(level ffcommon.FInt) (err error) {
	var t uintptr

	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_log_set_level").Call(
		uintptr(level),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Set the logging callback
 *
 * @note The callback must be thread safe, even if the application does not use
 *       threads itself as some codecs are multithreaded.
 *
 * @see av_log_default_callback
 *
 * @param callback A logging function with a compatible signature.
 */
//void av_log_set_callback(void (*callback)(void*, int, const char*, va_list));
//未测试
func AvLogSetCallback(callback func(ffcommon.FVoidP, ffcommon.FInt, ffcommon.FBuf, ffcommon.FVaList)) (err error) {
	var t uintptr

	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_log_set_callback").Call(
		uintptr(unsafe.Pointer(&callback)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Default logging callback
 *
 * It prints the message to stderr, optionally colorizing it.
 *
 * @param avcl A pointer to an arbitrary struct of which the first field is a
 *        pointer to an AVClass struct.
 * @param level The importance level of the message expressed using a @ref
 *        lavu_log_ffconstants "Logging ffconstant".
 * @param fmt The format string (printf-compatible) that specifies how
 *        subsequent arguments are converted to output.
 * @param vl The arguments referenced by the format string.
 */
//void av_log_default_callback(void *avcl, int level, const char *fmt,
//va_list vl);
//未测试
func AvLogDefaultCallback(avcl ffcommon.FVoidP, level ffcommon.FInt, fmt0 ffcommon.FConstCharP, vl ffcommon.FVaList) (err error) {
	var t uintptr
	var fmt0p *byte
	fmt0p, err = syscall.BytePtrFromString(fmt0)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_log_default_callback").Call(
		uintptr(unsafe.Pointer(fmt0p)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Return the context name
 *
 * @param  ctx The AVClass context
 *
 * @return The AVClass class_name
 */
//const char* av_default_item_name(void* ctx);
//未测试
func AvDefaultItemName(ctx ffcommon.FVoidP) (res ffcommon.FConstCharP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_default_item_name").Call(
		uintptr(unsafe.Pointer(ctx)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.GoAStr(t)
	return
}

//AVClassCategory av_default_get_category(void *ptr);

/**
 * Format a line of log the same way as the default callback.
 * @param line          buffer to receive the formatted line
 * @param line_size     size of the buffer
 * @param print_prefix  used to store whether the prefix must be printed;
 *                      must point to a persistent integer initially set to 1
 */
//void av_log_format_line(void *ptr, int level, const char *fmt, va_list vl,
//char *line, int line_size, int *print_prefix);
//未测试
func AvLogFormatLine(ptr ffcommon.FVoidP, level ffcommon.FInt, fmt0 ffcommon.FConstCharP, vl ffcommon.FBuf,
	line ffcommon.FBuf, line_size ffcommon.FInt, print_prefix *ffcommon.FInt) (err error) {
	var t uintptr
	var fmt0p *byte
	fmt0p, err = syscall.BytePtrFromString(fmt0)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_log_format_line").Call(
		ptr,
		uintptr(level),
		uintptr(unsafe.Pointer(fmt0p)),
		uintptr(unsafe.Pointer(vl)),
		uintptr(unsafe.Pointer(line)),
		uintptr(line_size),
		uintptr(unsafe.Pointer(print_prefix)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Format a line of log the same way as the default callback.
 * @param line          buffer to receive the formatted line;
 *                      may be NULL if line_size is 0
 * @param line_size     size of the buffer; at most line_size-1 characters will
 *                      be written to the buffer, plus one null terminator
 * @param print_prefix  used to store whether the prefix must be printed;
 *                      must point to a persistent integer initially set to 1
 * @return Returns a negative value if an error occurred, otherwise returns
 *         the number of characters that would have been written for a
 *         sufficiently large buffer, not including the terminating null
 *         character. If the return value is not less than line_size, it means
 *         that the log message was truncated to fit the buffer.
 */
//int av_log_format_line2(void *ptr, int level, const char *fmt, va_list vl,
//char *line, int line_size, int *print_prefix);
//未测试
func AvLogFormatLine2(ptr ffcommon.FVoidP, level ffcommon.FInt, fmt0 ffcommon.FConstCharP, vl ffcommon.FBuf,
	line ffcommon.FBuf, line_size ffcommon.FInt, print_prefix *ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	var fmt0p *byte
	fmt0p, err = syscall.BytePtrFromString(fmt0)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_log_format_line2").Call(
		ptr,
		uintptr(level),
		uintptr(unsafe.Pointer(fmt0p)),
		uintptr(unsafe.Pointer(vl)),
		uintptr(unsafe.Pointer(line)),
		uintptr(line_size),
		uintptr(unsafe.Pointer(print_prefix)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//void av_log_set_flags(int arg);
//未测试
func AvLogSetFlags(arg ffcommon.FInt) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_log_set_flags").Call(
		uintptr(arg),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

//int av_log_get_flags(void);
//未测试
func AvLogGetFlags() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_log_get_flags").Call()
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}
