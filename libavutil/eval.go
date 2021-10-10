package libavutil

import (
	"ffmpeg-go/ffcommon"
	"syscall"
	"unsafe"
)

type AVExpr struct {
}

/**
 * Parse and evaluate an expression.
 * Note, this is significantly slower than av_expr_eval().
 *
 * @param res a pointer to a double where is put the result value of
 * the expression, or NAN in case of error
 * @param s expression as a zero terminated string, for example "1+2^3+5*5+sin(2/3)"
 * @param const_names NULL terminated array of zero terminated strings of ffconstant identifiers, for example {"PI", "E", 0}
 * @param const_values a zero terminated array of values for the identifiers from const_names
 * @param func1_names NULL terminated array of zero terminated strings of funcs1 identifiers
 * @param funcs1 NULL terminated array of function pointers for functions which take 1 argument
 * @param func2_names NULL terminated array of zero terminated strings of funcs2 identifiers
 * @param funcs2 NULL terminated array of function pointers for functions which take 2 arguments
 * @param opaque a pointer which will be passed to all functions from funcs1 and funcs2
 * @param log_ctx parent logging context
 * @return >= 0 in case of success, a negative value corresponding to an
 * AVERROR code otherwise
 */
//int av_expr_parse_and_eval(double *res, const char *s,
//const char * const *const_names, const double *const_values,
//const char * const *func1_names, double (* const *funcs1)(void *, double),
//const char * const *func2_names, double (* const *funcs2)(void *, double, double),
//void *opaque, int log_offset, void *log_ctx);
//未测试
func AvExprParseAndEval(res0 *ffcommon.FDouble, s ffcommon.FConstCharP,
	const_names *ffcommon.FBuf,
	func1_names *ffcommon.FBuf, funcs1 func(ffcommon.FVoidP, ffcommon.FDouble) ffcommon.FDouble,
	func2_names *ffcommon.FBuf, funcs2 func(ffcommon.FVoidP, ffcommon.FDouble, ffcommon.FDouble) ffcommon.FDouble,
	opaque ffcommon.FVoidP, log_offset ffcommon.FInt, log_ctx ffcommon.FVoidP) (res ffcommon.FInt, err error) {
	var t uintptr
	var sp *byte
	sp, err = syscall.BytePtrFromString(s)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_expr_parse_and_eval").Call(
		uintptr(unsafe.Pointer(res0)),
		uintptr(unsafe.Pointer(sp)),
		uintptr(unsafe.Pointer(&const_names)),
		uintptr(unsafe.Pointer(&func1_names)),
		uintptr(unsafe.Pointer(&funcs1)),
		uintptr(unsafe.Pointer(&func2_names)),
		uintptr(unsafe.Pointer(&funcs2)),
		uintptr(unsafe.Pointer(opaque)),
		uintptr(log_offset),
		uintptr(unsafe.Pointer(log_ctx)),
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
 * Parse an expression.
 *
 * @param expr a pointer where is put an AVExpr containing the parsed
 * value in case of successful parsing, or NULL otherwise.
 * The pointed to AVExpr must be freed with av_expr_free() by the user
 * when it is not needed anymore.
 * @param s expression as a zero terminated string, for example "1+2^3+5*5+sin(2/3)"
 * @param const_names NULL terminated array of zero terminated strings of ffconstant identifiers, for example {"PI", "E", 0}
 * @param func1_names NULL terminated array of zero terminated strings of funcs1 identifiers
 * @param funcs1 NULL terminated array of function pointers for functions which take 1 argument
 * @param func2_names NULL terminated array of zero terminated strings of funcs2 identifiers
 * @param funcs2 NULL terminated array of function pointers for functions which take 2 arguments
 * @param log_ctx parent logging context
 * @return >= 0 in case of success, a negative value corresponding to an
 * AVERROR code otherwise
 */
//int av_expr_parse(AVExpr **expr, const char *s,
//const char * const *const_names,
//const char * const *func1_names, double (* const *funcs1)(void *, double),
//const char * const *func2_names, double (* const *funcs2)(void *, double, double),
//int log_offset, void *log_ctx);
//未测试
func AvExprarse(expr **AVExpr, s ffcommon.FConstCharP,
	const_names *ffcommon.FBuf,
	func1_names *ffcommon.FBuf, funcs1 func(ffcommon.FVoidP, ffcommon.FDouble) ffcommon.FDouble,
	func2_names *ffcommon.FBuf, funcs2 func(ffcommon.FVoidP, ffcommon.FDouble, ffcommon.FDouble) ffcommon.FDouble,
	log_offset ffcommon.FInt, log_ctx ffcommon.FVoidP) (res ffcommon.FInt, err error) {
	var t uintptr
	var sp *byte
	sp, err = syscall.BytePtrFromString(s)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_expr_parse").Call(
		uintptr(unsafe.Pointer(expr)),
		uintptr(unsafe.Pointer(sp)),
		uintptr(unsafe.Pointer(&const_names)),
		uintptr(unsafe.Pointer(&func1_names)),
		uintptr(unsafe.Pointer(&funcs1)),
		uintptr(unsafe.Pointer(&func2_names)),
		uintptr(unsafe.Pointer(&funcs2)),
		uintptr(log_offset),
		uintptr(unsafe.Pointer(log_ctx)),
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
 * Evaluate a previously parsed expression.
 *
 * @param const_values a zero terminated array of values for the identifiers from av_expr_parse() const_names
 * @param opaque a pointer which will be passed to all functions from funcs1 and funcs2
 * @return the value of the expression
 */
//double av_expr_eval(AVExpr *e, const double *const_values, void *opaque);
//未测试
func (e *AVExpr) AvExprEval(const_values *ffcommon.FDouble, opaque ffcommon.FVoidP) (res ffcommon.FDouble, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_expr_eval").Call(
		uintptr(unsafe.Pointer(e)),
		uintptr(unsafe.Pointer(const_values)),
		opaque,
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FDouble(t)
	return
}

/**
 * Track the presence of variables and their number of occurrences in a parsed expression
 *
 * @param counter a zero-initialized array where the count of each variable will be stored
 * @param size size of array
 * @return 0 on success, a negative value indicates that no expression or array was passed
 * or size was zero
 */
//int av_expr_count_vars(AVExpr *e, unsigned *counter, int size);
//未测试
func (e *AVExpr) AvExprCountVars(counter *ffcommon.FUnsigned, size ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_expr_count_vars").Call(
		uintptr(unsafe.Pointer(e)),
		uintptr(unsafe.Pointer(counter)),
		uintptr(size),
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
 * Track the presence of user provided functions and their number of occurrences
 * in a parsed expression.
 *
 * @param counter a zero-initialized array where the count of each function will be stored
 *                if you passed 5 functions with 2 arguments to av_expr_parse()
 *                then for arg=2 this will use upto 5 entries.
 * @param size size of array
 * @param arg number of arguments the counted functions have
 * @return 0 on success, a negative value indicates that no expression or array was passed
 * or size was zero
 */
//int av_expr_count_func(AVExpr *e, unsigned *counter, int size, int arg);
//未测试
func (e *AVExpr) AvExprCountFunc(counter *ffcommon.FUnsigned, size ffcommon.FInt, arg ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_expr_count_func").Call(
		uintptr(unsafe.Pointer(e)),
		uintptr(unsafe.Pointer(counter)),
		uintptr(size),
		uintptr(arg),
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
 * Free a parsed expression previously created with av_expr_parse().
 */
//void av_expr_free(AVExpr *e);
//未测试
func (e *AVExpr) AvExprFree() (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_expr_free").Call(
		uintptr(unsafe.Pointer(e)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Parse the string in numstr and return its value as a double. If
 * the string is empty, contains only whitespaces, or does not contain
 * an initial substring that has the expected syntax for a
 * floating-point number, no conversion is performed. In this case,
 * returns a value of zero and the value returned in tail is the value
 * of numstr.
 *
 * @param numstr a string representing a number, may contain one of
 * the International System number postfixes, for example 'K', 'M',
 * 'G'. If 'i' is appended after the postfix, powers of 2 are used
 * instead of powers of 10. The 'B' postfix multiplies the value by
 * 8, and can be appended after another postfix or used alone. This
 * allows using for example 'KB', 'MiB', 'G' and 'B' as postfix.
 * @param tail if non-NULL puts here the pointer to the char next
 * after the last parsed character
 */
//double av_strtod(const char *numstr, char **tail);
//未测试
func AvStrtod(numstr ffcommon.FConstCharP, tail **byte) (res ffcommon.FDouble, err error) {
	var t uintptr
	var numstrp *byte
	numstrp, err = syscall.BytePtrFromString(numstr)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_strtod").Call(
		uintptr(unsafe.Pointer(numstrp)),
		uintptr(unsafe.Pointer(&tail)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FDouble(t)
	return
}
