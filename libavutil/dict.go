package libavutil

import (
	"ffmpeg-go/ffcommon"
	"syscall"
	"unsafe"
)

type AVDictionaryEntry struct {
	key   ffcommon.FBuf
	value ffcommon.FBuf
}

type AVDictionary struct {
}

/**
* Get a dictionary entry with matching key.
*
* The returned entry key or value must not be changed, or it will
* cause undefined behavior.
*
* To iterate through all the dictionary entries, you can set the matching key
* to the null string "" and set the AV_DICT_IGNORE_SUFFIX flag.
*
* @param prev Set to the previous matching element to find the next.
*             If set to NULL the first matching element is returned.
* @param key matching key
* @param flags a collection of AV_DICT_* flags controlling how the entry is retrieved
* @return found entry or NULL in case no matching entry was found in the dictionary
 */
//AVDictionaryEntry *av_dict_get(const AVDictionary *m, const char *key,
//const AVDictionaryEntry *prev, int flags);
//未测试
func (m *AVDictionary) AvDictGet(key ffcommon.FConstCharP, prev *AVDictionaryEntry, flags ffcommon.FInt) (res *AVDictionaryEntry, err error) {
	var t uintptr
	var keyp *byte
	keyp, err = syscall.BytePtrFromString(key)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_dict_get").Call(
		uintptr(unsafe.Pointer(m)),
		uintptr(unsafe.Pointer(keyp)),
		uintptr(unsafe.Pointer(prev)),
		uintptr(flags),
	)
	if err != nil {
		//return
	}
	res = (*AVDictionaryEntry)(unsafe.Pointer(t))
	return
}

/**
* Get number of entries in dictionary.
*
* @param m dictionary
* @return  number of entries in dictionary
 */
//int av_dict_count(const AVDictionary *m);
//未测试
func (m *AVDictionary) AvDictCount() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_dict_count").Call(
		uintptr(unsafe.Pointer(m)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
* Set the given entry in *pm, overwriting an existing entry.
*
* Note: If AV_DICT_DONT_STRDUP_KEY or AV_DICT_DONT_STRDUP_VAL is set,
* these arguments will be freed on error.
*
* Warning: Adding a new entry to a dictionary invalidates all existing entries
* previously returned with av_dict_get.
*
* @param pm pointer to a pointer to a dictionary struct. If *pm is NULL
* a dictionary struct is allocated and put in *pm.
* @param key entry key to add to *pm (will either be av_strduped or added as a new key depending on flags)
* @param value entry value to add to *pm (will be av_strduped or added as a new key depending on flags).
*        Passing a NULL value will cause an existing entry to be deleted.
* @return >= 0 on success otherwise an error code <0
 */
//int av_dict_set(AVDictionary **pm, const char *key, const char *value, int flags);
//未测试
func AvDictSet(pm **AVDictionary, key ffcommon.FConstCharP, value ffcommon.FConstCharP, flags ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr

	var keyp *byte
	keyp, err = syscall.BytePtrFromString(key)
	if err != nil {
		return
	}

	var valuep *byte
	valuep, err = syscall.BytePtrFromString(value)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_dict_set").Call(
		uintptr(unsafe.Pointer(&pm)),
		uintptr(unsafe.Pointer(keyp)),
		uintptr(unsafe.Pointer(valuep)),
		uintptr(flags),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
* Convenience wrapper for av_dict_set that converts the value to a string
* and stores it.
*
* Note: If AV_DICT_DONT_STRDUP_KEY is set, key will be freed on error.
 */
//int av_dict_set_int(AVDictionary **pm, const char *key, int64_t value, int flags);
//未测试
func aADictSetInt(pm **AVDictionary, key ffcommon.FConstCharP, value ffcommon.FConstCharP, flags ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr

	var keyp *byte
	keyp, err = syscall.BytePtrFromString(key)
	if err != nil {
		return
	}

	var valuep *byte
	valuep, err = syscall.BytePtrFromString(value)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_dict_set_int").Call(
		uintptr(unsafe.Pointer(&pm)),
		uintptr(unsafe.Pointer(keyp)),
		uintptr(unsafe.Pointer(valuep)),
		uintptr(flags),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
* Parse the key/value pairs list and add the parsed entries to a dictionary.
*
* In case of failure, all the successfully set entries are stored in
* *pm. You may need to manually free the created dictionary.
*
* @param key_val_sep  a 0-terminated list of characters used to separate
*                     key from value
* @param pairs_sep    a 0-terminated list of characters used to separate
*                     two pairs from each other
* @param flags        flags to use when adding to dictionary.
*                     AV_DICT_DONT_STRDUP_KEY and AV_DICT_DONT_STRDUP_VAL
*                     are ignored since the key/value tokens will always
*                     be duplicated.
* @return             0 on success, negative AVERROR code on failure
 */
//int av_dict_parse_string(AVDictionary **pm, const char *str,
//const char *key_val_sep, const char *pairs_sep,
//int flags);
//未测试
func AvDictParseString(pm **AVDictionary, key ffcommon.FConstCharP, key_val_sep ffcommon.FConstCharP, pairs_sep ffcommon.FConstCharP, flags ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr

	var keyp *byte
	keyp, err = syscall.BytePtrFromString(key)
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

	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_dict_parse_string").Call(
		uintptr(unsafe.Pointer(&pm)),
		uintptr(unsafe.Pointer(keyp)),
		uintptr(unsafe.Pointer(key_val_sepp)),
		uintptr(unsafe.Pointer(pairs_sepp)),
		uintptr(flags),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
* Copy entries from one AVDictionary struct into another.
* @param dst pointer to a pointer to a AVDictionary struct. If *dst is NULL,
*            this function will allocate a struct for you and put it in *dst
* @param src pointer to source AVDictionary struct
* @param flags flags to use when setting entries in *dst
* @note metadata is read using the AV_DICT_IGNORE_SUFFIX flag
* @return 0 on success, negative AVERROR code on failure. If dst was allocated
*           by this function, callers should free the associated memory.
 */
//int av_dict_copy(AVDictionary **dst, const AVDictionary *src, int flags);
//未测试
func AvDictCopy(dst **AVDictionary, src *AVDictionary, flags ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr

	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_dict_copy").Call(
		uintptr(unsafe.Pointer(&dst)),
		uintptr(unsafe.Pointer(src)),
		uintptr(flags),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
* Free all the memory allocated for an AVDictionary struct
* and all keys and values.
 */
//void av_dict_free(AVDictionary **m);
//未测试
func AvDictFree(m **AVDictionary) (err error) {
	var t uintptr

	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_dict_free").Call(
		uintptr(unsafe.Pointer(&m)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
* Get dictionary entries as a string.
*
* Create a string containing dictionary's entries.
* Such string may be passed back to av_dict_parse_string().
* @note String is escaped with backslashes ('\').
*
* @param[in]  m             dictionary
* @param[out] buffer        Pointer to buffer that will be allocated with string containg entries.
*                           Buffer must be freed by the caller when is no longer needed.
* @param[in]  key_val_sep   character used to separate key from value
* @param[in]  pairs_sep     character used to separate two pairs from each other
* @return                   >= 0 on success, negative on error
* @warning Separators cannot be neither '\\' nor '\0'. They also cannot be the same.
 */
//int av_dict_get_string(const AVDictionary *m, char **buffer,
//const char key_val_sep, const char pairs_sep);
//未测试
func (m *AVDictionary) AvDictGetString(buffer *ffcommon.FConstCharP,
	key_val_sep ffcommon.FConstCharP, pairs_sep ffcommon.FConstCharP) (res ffcommon.FInt, err error) {
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
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_dict_get_string").Call(
		uintptr(unsafe.Pointer(m)),
		uintptr(unsafe.Pointer(&buffer)),
		uintptr(unsafe.Pointer(key_val_sepp)),
		uintptr(unsafe.Pointer(pairs_sepp)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}
