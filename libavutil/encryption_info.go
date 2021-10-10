package libavutil

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

type AVSubsampleEncryptionInfo struct {
	/** The number of bytes that are clear. */
	bytes_of_clear_data ffcommon.FUnsignedInt

	/**
	 * The number of bytes that are protected.  If using pattern encryption,
	 * the pattern applies to only the protected bytes; if not using pattern
	 * encryption, all these bytes are encrypted.
	 */
	bytes_of_protected_data ffcommon.FUnsignedInt
}

/**
 * This describes encryption info for a packet.  This contains frame-specific
 * info for how to decrypt the packet before passing it to the decoder.
 *
 * The size of this struct is not part of the public ABI.
 */
type AVEncryptionInfo struct {
	/** The fourcc encryption scheme, in big-endian byte order. */
	scheme ffcommon.FUint32T

	/**
	 * Only used for pattern encryption.  This is the number of 16-byte blocks
	 * that are encrypted.
	 */
	crypt_byte_block ffcommon.FUint32T

	/**
	 * Only used for pattern encryption.  This is the number of 16-byte blocks
	 * that are clear.
	 */
	skip_byte_block ffcommon.FUint32T

	/**
	 * The ID of the key used to encrypt the packet.  This should always be
	 * 16 bytes long, but may be changed in the future.
	 */
	key_id      *ffcommon.FUint8T
	key_id_size ffcommon.FUint32T

	/**
	 * The initialization vector.  This may have been zero-filled to be the
	 * correct block size.  This should always be 16 bytes long, but may be
	 * changed in the future.
	 */
	iv      *ffcommon.FUint8T
	iv_size ffcommon.FUint32T

	/**
	 * An array of subsample encryption info specifying how parts of the sample
	 * are encrypted.  If there are no subsamples, then the whole sample is
	 * encrypted.
	 */
	subsamples      *AVSubsampleEncryptionInfo
	subsample_count ffcommon.FUint32T
}

/**
 * This describes info used to initialize an encryption key system.
 *
 * The size of this struct is not part of the public ABI.
 */
type AVEncryptionInitInfo struct {
	/**
	 * A unique identifier for the key system this is for, can be NULL if it
	 * is not known.  This should always be 16 bytes, but may change in the
	 * future.
	 */
	system_id      *ffcommon.FUint8T
	system_id_size ffcommon.FUint32T

	/**
	 * An array of key IDs this initialization data is for.  All IDs are the
	 * same length.  Can be NULL if there are no known key IDs.
	 */
	key_ids **ffcommon.FUint8T
	/** The number of key IDs. */
	num_key_ids ffcommon.FUint32T
	/**
	 * The number of bytes in each key ID.  This should always be 16, but may
	 * change in the future.
	 */
	key_id_size ffcommon.FUint32T

	/**
	 * Key-system specific initialization data.  This data is copied directly
	 * from the file and the format depends on the specific key system.  This
	 * can be NULL if there is no initialization data; in that case, there
	 * will be at least one key ID.
	 */
	data      *ffcommon.FUint8T
	data_size ffcommon.FUint32T

	/**
	 * An optional pointer to the next initialization info in the list.
	 */
	next *AVEncryptionInitInfo
}

/**
* Allocates an AVEncryptionInfo structure and sub-pointers to hold the given
* number of subsamples.  This will allocate pointers for the key ID, IV,
* and subsample entries, set the size members, and zero-initialize the rest.
*
* @param subsample_count The number of subsamples.
* @param key_id_size The number of bytes in the key ID, should be 16.
* @param iv_size The number of bytes in the IV, should be 16.
*
* @return The new AVEncryptionInfo structure, or NULL on error.
 */
//AVEncryptionInfo *av_encryption_info_alloc(uint32_t subsample_count, uint32_t key_id_size, uint32_t iv_size);
//未测试
func AvEncryptionInfoAlloc(subsample_count ffcommon.FUint32T, key_id_size ffcommon.FUint32T, iv_size ffcommon.FUint32T) (res *AVEncryptionInfo, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_encryption_info_alloc").Call(
		uintptr(subsample_count),
		uintptr(key_id_size),
		uintptr(iv_size),
	)
	if err != nil {
		//return
	}
	res = (*AVEncryptionInfo)(unsafe.Pointer(t))
	return
}

/**
* Allocates an AVEncryptionInfo structure with a copy of the given data.
* @return The new AVEncryptionInfo structure, or NULL on error.
 */
//AVEncryptionInfo *av_encryption_info_clone(const AVEncryptionInfo *info);
//未测试
func (info *AVEncryptionInfo) AvEncryptionInfoClone() (res *AVEncryptionInfo, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_encryption_info_clone").Call(
		uintptr(unsafe.Pointer(info)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = (*AVEncryptionInfo)(unsafe.Pointer(t))
	return
}

/**
* Frees the given encryption info object.  This MUST NOT be used to free the
* side-data data pointer, that should use normal side-data methods.
 */
//void av_encryption_info_free(AVEncryptionInfo *info);
//未测试
func (info *AVEncryptionInfo) AvEncryptionInfoFree() (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_encryption_info_free").Call(
		uintptr(unsafe.Pointer(info)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
* Creates a copy of the AVEncryptionInfo that is contained in the given side
* data.  The resulting object should be passed to av_encryption_info_free()
* when done.
*
* @return The new AVEncryptionInfo structure, or NULL on error.
 */
//AVEncryptionInfo *av_encryption_info_get_side_data(const uint8_t *side_data, size_t side_data_size);
//未测试
func AvEncryptionInfoGetSideData(side_data *ffcommon.FUint8T, side_data_size ffcommon.FSizeT) (res *AVEncryptionInfo, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_encryption_info_get_side_data").Call(
		uintptr(unsafe.Pointer(side_data)),
		uintptr(side_data_size),
	)
	if err != nil {
		//return
	}
	res = (*AVEncryptionInfo)(unsafe.Pointer(t))
	return
}

/**
* Allocates and initializes side data that holds a copy of the given encryption
* info.  The resulting pointer should be either freed using av_free or given
* to av_packet_add_side_data().
*
* @return The new side-data pointer, or NULL.
 */
//uint8_t *av_encryption_info_add_side_data(
//const AVEncryptionInfo *info, size_t *side_data_size);
//未测试
func (info *AVEncryptionInfo) AvEncryptionInfoAddSideData(side_data_size *ffcommon.FSizeT) (res *ffcommon.FUint8T, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_encryption_info_add_side_data").Call(
		uintptr(unsafe.Pointer(info)),
		uintptr(unsafe.Pointer(side_data_size)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = (*ffcommon.FUint8T)(unsafe.Pointer(t))
	return
}

/**
* Allocates an AVEncryptionInitInfo structure and sub-pointers to hold the
* given sizes.  This will allocate pointers and set all the fields.
*
* @return The new AVEncryptionInitInfo structure, or NULL on error.
 */
//AVEncryptionInitInfo *av_encryption_init_info_alloc(
//uint32_t system_id_size, uint32_t num_key_ids, uint32_t key_id_size, uint32_t data_size);
//未测试
func AvEncryptionInitInfoAlloc(system_id_size ffcommon.FUint32T, num_key_ids ffcommon.FUint32T, key_id_size ffcommon.FUint32T, data_size ffcommon.FUint32T) (res *AVEncryptionInitInfo, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_encryption_init_info_alloc").Call(
		uintptr(system_id_size),
		uintptr(num_key_ids),
		uintptr(key_id_size),
		uintptr(data_size),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = (*AVEncryptionInitInfo)(unsafe.Pointer(t))
	return
}

/**
* Frees the given encryption init info object.  This MUST NOT be used to free
* the side-data data pointer, that should use normal side-data methods.
 */
//void av_encryption_init_info_free(AVEncryptionInitInfo* info);
//未测试
func (info *AVEncryptionInitInfo) AvEncryptionInitInfoFree() (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_encryption_init_info_free").Call(
		uintptr(unsafe.Pointer(info)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
* Creates a copy of the AVEncryptionInitInfo that is contained in the given
* side data.  The resulting object should be passed to
* av_encryption_init_info_free() when done.
*
* @return The new AVEncryptionInitInfo structure, or NULL on error.
 */
//AVEncryptionInitInfo *av_encryption_init_info_get_side_data(
//const uint8_t* side_data, size_t side_data_size);
//未测试
func AvEncryptionInitInfoGetSideData(side_data *ffcommon.FUint8T, side_data_size ffcommon.FSizeT) (res *AVEncryptionInitInfo, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_encryption_init_info_get_side_data").Call(
		uintptr(unsafe.Pointer(side_data)),
		uintptr(side_data_size),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = (*AVEncryptionInitInfo)(unsafe.Pointer(t))
	return
}

/**
* Allocates and initializes side data that holds a copy of the given encryption
* init info.  The resulting pointer should be either freed using av_free or
* given to av_packet_add_side_data().
*
* @return The new side-data pointer, or NULL.
 */
//uint8_t *av_encryption_init_info_add_side_data(
//const AVEncryptionInitInfo *info, size_t *side_data_size);
//未测试
func (info *AVEncryptionInitInfo) AvEncryptionInitInfoAddSideData(side_data_size *ffcommon.FSizeT) (res *ffcommon.FUint8T, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_encryption_init_info_add_side_data").Call(
		uintptr(unsafe.Pointer(info)),
		uintptr(unsafe.Pointer(side_data_size)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = (*ffcommon.FUint8T)(unsafe.Pointer(t))
	return
}
