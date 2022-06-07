package libavutil

import (
	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"unsafe"
)

/**
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

//#ifndef AVUTIL_ENCRYPTION_INFO_H
//#define AVUTIL_ENCRYPTION_INFO_H
//
//#include <stddef.h>
//#include <stdint.h>

type AVSubsampleEncryptionInfo struct {

	/** The number of bytes that are clear. */
	BytesOfClearData ffcommon.FUnsignedInt

	/**
	 * The number of bytes that are protected.  If using pattern encryption,
	 * the pattern applies to only the protected bytes; if not using pattern
	 * encryption, all these bytes are encrypted.
	 */
	BytesOfProtectedData ffcommon.FUnsignedInt
}

/**
 * This describes encryption info for a packet.  This contains frame-specific
 * info for how to decrypt the packet before passing it to the decoder.
 *
 * The size of this struct is not part of the public ABI.
 */
type AVEncryptionInfo struct {

	/** The fourcc encryption scheme, in big-endian byte order. */
	Scheme ffcommon.FUint32T

	/**
	 * Only used for pattern encryption.  This is the number of 16-byte blocks
	 * that are encrypted.
	 */
	CryptByteBlock ffcommon.FUint32T

	/**
	 * Only used for pattern encryption.  This is the number of 16-byte blocks
	 * that are clear.
	 */
	SkipByteBlock ffcommon.FUint32T

	/**
	 * The ID of the key used to encrypt the packet.  This should always be
	 * 16 bytes long, but may be changed in the future.
	 */
	KeyId     *ffcommon.FUint8T
	KeyIdSize ffcommon.FUint32T

	/**
	 * The initialization vector.  This may have been zero-filled to be the
	 * correct block size.  This should always be 16 bytes long, but may be
	 * changed in the future.
	 */
	Iv     *ffcommon.FUint8T
	IvSize ffcommon.FUint32T

	/**
	 * An array of subsample encryption info specifying how parts of the sample
	 * are encrypted.  If there are no subsamples, then the whole sample is
	 * encrypted.
	 */
	Subsamples     *AVSubsampleEncryptionInfo
	SubsampleCount ffcommon.FUint32T
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
	SystemId     *ffcommon.FUint8T
	SystemIdSize ffcommon.FUint32T

	/**
	 * An array of key IDs this initialization data is for.  All IDs are the
	 * same length.  Can be NULL if there are no known key IDs.
	 */
	KeyIds **ffcommon.FUint8T
	/** The number of key IDs. */
	NumKeyIds ffcommon.FUint32T
	/**
	 * The number of bytes in each key ID.  This should always be 16, but may
	 * change in the future.
	 */
	KeyIdSize ffcommon.FUint32T

	/**
	 * Key-system specific initialization data.  This data is copied directly
	 * from the file and the format depends on the specific key system.  This
	 * can be NULL if there is no initialization data; in that case, there
	 * will be at least one key ID.
	 */
	Data     *ffcommon.FUint8T
	DataSize ffcommon.FUint32T
	/**
	 * An optional pointer to the next initialization info in the list.
	 */
	Next *AVEncryptionInitInfo
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
func AvEncryptionInfoAlloc(subsample_count, key_id_size, iv_size ffcommon.FUint32T) (res *AVEncryptionInfo) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_encryption_info_alloc").Call(
		uintptr(subsample_count),
		uintptr(key_id_size),
		uintptr(iv_size),
	)
	if t == 0 {

	}
	res = (*AVEncryptionInfo)(unsafe.Pointer(t))
	return
}

/**
 * Allocates an AVEncryptionInfo structure with a copy of the given data.
 * @return The new AVEncryptionInfo structure, or NULL on error.
 */
//AVEncryptionInfo *av_encryption_info_clone(const AVEncryptionInfo *info);
func (info *AVEncryptionInfo) AvEncryptionInfoClone() (res *AVEncryptionInfo) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_encryption_info_clone").Call(
		uintptr(unsafe.Pointer(info)),
	)
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
func (info *AVEncryptionInfo) AvEncryptionInfoFree() {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_encryption_info_free").Call(
		uintptr(unsafe.Pointer(info)),
	)
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
func AvEncryptionInfoGetSideData(side_data ffcommon.FConstCharP, side_data_size ffcommon.FSizeT) (res *AVEncryptionInfo) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_encryption_info_get_side_data").Call(
		ffcommon.UintPtrFromString(side_data),
		uintptr(side_data_size),
	)
	if t == 0 {

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
func (info *AVEncryptionInfo) AvEncryptionInfo_AddSideData(side_data_size ffcommon.FSizeT) (res *ffcommon.FUint8T) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_encryption_info_add_side_data").Call(
		uintptr(unsafe.Pointer(info)),
		uintptr(side_data_size),
	)
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
func AvEncryptionInitInfoAlloc(system_id_size, num_key_ids, key_id_size, data_size ffcommon.FUint32T) (res *AVEncryptionInitInfo) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_encryption_init_info_alloc").Call(
		uintptr(system_id_size),
		uintptr(num_key_ids),
		uintptr(key_id_size),
		uintptr(data_size),
	)
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
func (info *AVEncryptionInfo) AvEncryptionInitInfoFree() {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_encryption_init_info_free").Call(
		uintptr(unsafe.Pointer(info)),
	)
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
func AvEncryptionInitInfoGetSideData(side_data *ffcommon.FUint8T, side_data_size ffcommon.FSizeT) (res *AVEncryptionInitInfo) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_encryption_init_info_get_side_data").Call(
		uintptr(unsafe.Pointer(side_data)),
		uintptr(side_data_size),
	)
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
func (info *AVEncryptionInfo) AvEncryptionInitInfoAddSideData(side_data_size *ffcommon.FSizeT) (res *ffcommon.FUint8T) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_encryption_init_info_add_side_data").Call(
		uintptr(unsafe.Pointer(info)),
		uintptr(unsafe.Pointer(side_data_size)),
	)
	if t == 0 {

	}
	res = (*ffcommon.FUint8T)(unsafe.Pointer(t))
	return
}

//#endif /* AVUTIL_ENCRYPTION_INFO_H */
