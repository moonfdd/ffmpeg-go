package libavcodec

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"ffmpeg-go/libavutil"
	"unsafe"
)

type AVPacketSideData struct {
	data *ffcommon.FUint8T
	//#if FF_API_BUFFER_SIZE_T
	size ffcommon.FInt
	//#else
	//size_t   size;
	//#endif
	type0 ffconstant.AVPacketSideDataType
}

/**
 * This structure stores compressed data. It is typically exported by demuxers
 * and then passed as input to decoders, or received as output from encoders and
 * then passed to muxers.
 *
 * For video, it should typically contain one compressed frame. For audio it may
 * contain several compressed frames. Encoders are allowed to output empty
 * packets, with no compressed data, containing only side data
 * (e.g. to update some stream parameters at the end of encoding).
 *
 * The semantics of data ownership depends on the buf field.
 * If it is set, the packet data is dynamically allocated and is
 * valid indefinitely until a call to av_packet_unref() reduces the
 * reference count to 0.
 *
 * If the buf field is not set av_packet_ref() would make a copy instead
 * of increasing the reference count.
 *
 * The side data is always allocated with av_malloc(), copied by
 * av_packet_ref() and freed by av_packet_unref().
 *
 * sizeof(AVPacket) being a part of the public ABI is deprecated. once
 * av_init_packet() is removed, new packets will only be able to be allocated
 * with av_packet_alloc(), and new fields may be added to the end of the struct
 * with a minor bump.
 *
 * @see av_packet_alloc
 * @see av_packet_ref
 * @see av_packet_unref
 */
type AVPacket struct {
	/**
	 * A reference to the reference-counted buffer where the packet data is
	 * stored.
	 * May be NULL, then the packet data is not reference-counted.
	 */
	buf *libavutil.AVBufferRef
	/**
	 * Presentation timestamp in AVStream->time_base units; the time at which
	 * the decompressed packet will be presented to the user.
	 * Can be AV_NOPTS_VALUE if it is not stored in the file.
	 * pts MUST be larger or equal to dts as presentation cannot happen before
	 * decompression, unless one wants to view hex dumps. Some formats misuse
	 * the terms dts and pts/cts to mean something different. Such timestamps
	 * must be converted to true pts/dts before they are stored in AVPacket.
	 */
	pts ffcommon.FInt64T
	/**
	 * Decompression timestamp in AVStream->time_base units; the time at which
	 * the packet is decompressed.
	 * Can be AV_NOPTS_VALUE if it is not stored in the file.
	 */
	dts          ffcommon.FInt64T
	data         ffcommon.FUint8T
	size         ffcommon.FUint
	stream_index ffcommon.FUint
	/**
	 * A combination of AV_PKT_FLAG values
	 */
	flags ffcommon.FUint
	/**
	 * Additional packet data that can be provided by the container.
	 * Packet can contain several types of side information.
	 */
	side_data       *AVPacketSideData
	side_data_elems ffcommon.FUint

	/**
	 * Duration of this packet in AVStream->time_base units, 0 if unknown.
	 * Equals next_pts - this_pts in presentation order.
	 */
	duration ffcommon.FInt64T

	pos ffcommon.FInt64T ///< byte position in stream, -1 if unknown

	/**
	 * @deprecated Same as the duration field, but as int64_t. This was required
	 * for Matroska subtitles, whose duration values could overflow when the
	 * duration field was still an int.
	 */
	convergence_duration ffcommon.FInt64T
}

type AVPacketList struct {
	pkt  AVPacket
	next *AVPacketList
}

/**
* Allocate an AVPacket and set its fields to default values.  The resulting
* struct must be freed using av_packet_free().
*
* @return An AVPacket filled with default values or NULL on failure.
*
* @note this only allocates the AVPacket itself, not the data buffers. Those
* must be allocated through other means such as av_new_packet.
*
* @see av_new_packet
 */
//AVPacket *av_packet_alloc(void);
//未测试
func AvPacketAlloc() (res *AVPacket, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_packet_alloc").Call()
	if err != nil {
		//return
	}
	//res = &AVAES{}
	res = (*AVPacket)(unsafe.Pointer(t))
	//res.instance = t
	//res.ptr = unsafe.Pointer(t)
	return
}

/**
* Create a new packet that references the same data as src.
*
* This is a shortcut for av_packet_alloc()+av_packet_ref().
*
* @return newly created AVPacket on success, NULL on error.
*
* @see av_packet_alloc
* @see av_packet_ref
 */
//AVPacket *av_packet_clone(const AVPacket *src);
//未测试
func (src *AVPacket) AvPacketClone() (res *AVPacket, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_packet_clone").Call(
		uintptr(unsafe.Pointer(src)),
	)
	if err != nil {
		//return
	}
	//res = &AVAES{}
	res = (*AVPacket)(unsafe.Pointer(t))
	//res.instance = t
	//res.ptr = unsafe.Pointer(t)
	return
}

/**
* Free the packet, if the packet is reference counted, it will be
* unreferenced first.
*
* @param pkt packet to be freed. The pointer will be set to NULL.
* @note passing NULL is a no-op.
 */
//void av_packet_free(AVPacket **pkt);
//未测试
func AvPacketFree(pkt **AVPacket) (err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_packet_free").Call(
		uintptr(unsafe.Pointer(&pkt)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

//#if FF_API_INIT_PACKET
/**
* Initialize optional fields of a packet with default values.
*
* Note, this does not touch the data and size members, which have to be
* initialized separately.
*
* @param pkt packet
*
* @see av_packet_alloc
* @see av_packet_unref
*
* @deprecated This function is deprecated. Once it's removed,
              sizeof(AVPacket) will not be a part of the ABI anymore.
*/
//attribute_deprecated
//void av_init_packet(AVPacket *pkt);
//未测试
func (pkt *AVPacket) AvInitPacket() (err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_init_packet").Call(
		uintptr(unsafe.Pointer(pkt)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

//#endif

/**
* Allocate the payload of a packet and initialize its fields with
* default values.
*
* @param pkt packet
* @param size wanted payload size
* @return 0 if OK, AVERROR_xxx otherwise
 */
//int av_new_packet(AVPacket *pkt, int size);
//未测试
func (pkt *AVPacket) AvNewPacket(size ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_new_packet").Call(
		uintptr(unsafe.Pointer(pkt)),
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
* Reduce packet size, correctly zeroing padding
*
* @param pkt packet
* @param size new size
 */
//void av_shrink_packet(AVPacket *pkt, int size);
//未测试
func (pkt *AVPacket) AvShrinkPacket(size ffcommon.FInt) (err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_shrink_packet").Call(
		uintptr(unsafe.Pointer(pkt)),
		uintptr(size),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
* Increase packet size, correctly zeroing padding
*
* @param pkt packet
* @param grow_by number of bytes by which to increase the size of the packet
 */
//int av_grow_packet(AVPacket *pkt, int grow_by);
//未测试
func (pkt *AVPacket) AvGrowPacket(grow_by ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_grow_packet").Call(
		uintptr(unsafe.Pointer(pkt)),
		uintptr(grow_by),
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
* Initialize a reference-counted packet from av_malloc()ed data.
*
* @param pkt packet to be initialized. This function will set the data, size,
*        and buf fields, all others are left untouched.
* @param data Data allocated by av_malloc() to be used as packet data. If this
*        function returns successfully, the data is owned by the underlying AVBuffer.
*        The caller may not access the data through other means.
* @param size size of data in bytes, without the padding. I.e. the full buffer
*        size is assumed to be size + AV_INPUT_BUFFER_PADDING_SIZE.
*
* @return 0 on success, a negative AVERROR on error
 */
//int av_packet_from_data(AVPacket *pkt, uint8_t *data, int size);
//未测试
func (pkt *AVPacket) AvPacketFromData(data *ffcommon.FUint8T, size ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_packet_from_data").Call(
		uintptr(unsafe.Pointer(pkt)),
		uintptr(unsafe.Pointer(data)),
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

//#if FF_API_AVPACKET_OLD_API
///**
//* @warning This is a hack - the packet memory allocation stuff is broken. The
//* packet is allocated if it was not really allocated.
//*
//* @deprecated Use av_packet_ref or av_packet_make_refcounted
//*/
//attribute_deprecated
//int av_dup_packet(AVPacket *pkt);
//未测试
func (pkt *AVPacket) AvDupPacket() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_dup_packet").Call(
		uintptr(unsafe.Pointer(pkt)),
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
* Copy packet, including contents
*
* @return 0 on success, negative AVERROR on fail
*
* @deprecated Use av_packet_ref
 */
//attribute_deprecated
//int av_copy_packet(AVPacket *dst, const AVPacket *src);
//未测试
func AvCopyPacket(dst *AVPacket, src *AVPacket) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_copy_packet").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
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
* Copy packet side data
*
* @return 0 on success, negative AVERROR on fail
*
* @deprecated Use av_packet_copy_props
 */
//attribute_deprecated
//int av_copy_packet_side_data(AVPacket *dst, const AVPacket *src);
//未测试
func AvCopyPacketSideData(dst *AVPacket, src *AVPacket) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_copy_packet_side_data").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
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
* Free a packet.
*
* @deprecated Use av_packet_unref
*
* @param pkt packet to free
 */
//attribute_deprecated
//void av_free_packet(AVPacket *pkt);
//#endif
//未测试
func (pkt *AVPacket) AvFreePacket() (err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_free_packet").Call(
		uintptr(unsafe.Pointer(pkt)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
* Allocate new information of a packet.
*
* @param pkt packet
* @param type side information type
* @param size side information size
* @return pointer to fresh allocated data or NULL otherwise
 */
//uint8_t* av_packet_new_side_data(AVPacket *pkt, enum AVPacketSideDataType type,
//#if FF_API_BUFFER_SIZE_T
//int size);
//#else
//size_t size);
//#endif
//未测试
func (pkt *AVPacket) AvPacketNewSideData(type0 ffconstant.AVPacketSideDataType,
	size ffcommon.FInt) (res *ffcommon.FUint8T, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_packet_new_side_data").Call(
		uintptr(unsafe.Pointer(pkt)),
		uintptr(type0),
		uintptr(size),
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
* Wrap an existing array as a packet side data.
*
* @param pkt packet
* @param type side information type
* @param data the side data array. It must be allocated with the av_malloc()
*             family of functions. The ownership of the data is transferred to
*             pkt.
* @param size side information size
* @return a non-negative number on success, a negative AVERROR code on
*         failure. On failure, the packet is unchanged and the data remains
*         owned by the caller.
 */
//int av_packet_add_side_data(AVPacket *pkt, enum AVPacketSideDataType type,
//uint8_t *data, size_t size);
//未测试
func (pkt *AVPacket) AvPacketAddSideData(type0 ffconstant.AVPacketSideDataType,
	data *ffcommon.FUint8T, size ffcommon.FSizeT) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_packet_add_side_data").Call(
		uintptr(unsafe.Pointer(pkt)),
		uintptr(type0),
		uintptr(unsafe.Pointer(data)),
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
* Shrink the already allocated side data buffer
*
* @param pkt packet
* @param type side information type
* @param size new side information size
* @return 0 on success, < 0 on failure
 */
//int av_packet_shrink_side_data(AVPacket *pkt, enum AVPacketSideDataType type,
//#if FF_API_BUFFER_SIZE_T
//int size);
//#else
//size_t size);
//#endif
//未测试
func (pkt *AVPacket) AvPacketShrinkSideData(type0 ffconstant.AVPacketSideDataType,
	size ffcommon.FSizeT) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_packet_shrink_side_data").Call(
		uintptr(unsafe.Pointer(pkt)),
		uintptr(type0),
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
* Get side information from packet.
*
* @param pkt packet
* @param type desired side information type
* @param size If supplied, *size will be set to the size of the side data
*             or to zero if the desired side data is not present.
* @return pointer to data if present or NULL otherwise
 */
//uint8_t* av_packet_get_side_data(const AVPacket *pkt, enum AVPacketSideDataType type,
//#if FF_API_BUFFER_SIZE_T
//int *size);
//#else
//size_t *size);
//#endif

//未测试
func (pkt *AVPacket) AvPacketGetSideData(type0 ffconstant.AVPacketSideDataType,
	size ffcommon.FSizeT) (res *ffcommon.FUint8T, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_packet_get_side_data").Call(
		uintptr(unsafe.Pointer(pkt)),
		uintptr(type0),
		uintptr(size),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = (*ffcommon.FUint8T)(unsafe.Pointer(t))
	return
}

//#if FF_API_MERGE_SD_API
//attribute_deprecated
//int av_packet_merge_side_data(AVPacket *pkt);
//未测试
func (pkt *AVPacket) AvPacketMergeSideData() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_packet_merge_side_data").Call(
		uintptr(unsafe.Pointer(pkt)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//attribute_deprecated
//int av_packet_split_side_data(AVPacket *pkt);
//#endif
//未测试
func (pkt *AVPacket) AvPacketSplitSideData() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_packet_split_side_data").Call(
		uintptr(unsafe.Pointer(pkt)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

//const char *av_packet_side_data_name(enum AVPacketSideDataType type);
//未测试
func AvPacketSideDataName(type0 ffconstant.AVPacketSideDataType) (res ffcommon.FConstCharP, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_packet_side_data_name").Call(
		uintptr(type0),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.GoAStr(t)
	return
}

/**
* Pack a dictionary for use in side_data.
*
* @param dict The dictionary to pack.
* @param size pointer to store the size of the returned data
* @return pointer to data if successful, NULL otherwise
 */
//#if FF_API_BUFFER_SIZE_T
//uint8_t *av_packet_pack_dictionary(AVDictionary *dict, int *size);
//#else
//uint8_t *av_packet_pack_dictionary(AVDictionary *dict, size_t *size);
//#endif
//未测试
func AvPacketPackDictionary(dict *libavutil.AVDictionary, size *ffcommon.FInt) (res *ffcommon.FUint8T, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_packet_pack_dictionary").Call(
		uintptr(unsafe.Pointer(dict)),
		uintptr(unsafe.Pointer(size)),
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
* Unpack a dictionary from side_data.
*
* @param data data from side_data
* @param size size of the data
* @param dict the metadata storage dictionary
* @return 0 on success, < 0 on failure
 */
//#if FF_API_BUFFER_SIZE_T
//int av_packet_unpack_dictionary(const uint8_t *data, int size, AVDictionary **dict);
//#else
//int av_packet_unpack_dictionary(const uint8_t *data, size_t size,
//AVDictionary **dict);
//#endif
//未测试
func AvPacketUnpackDictionary(data *ffcommon.FUint8T, size ffcommon.FInt, dict *libavutil.AVDictionary) (res *ffcommon.FUint8T, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_packet_unpack_dictionary").Call(
		uintptr(unsafe.Pointer(data)),
		uintptr(size),
		uintptr(unsafe.Pointer(dict)),
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
* Convenience function to free all the side data stored.
* All the other fields stay untouched.
*
* @param pkt packet
 */
//void av_packet_free_side_data(AVPacket *pkt);
//未测试
func (pkt *AVPacket) AvPacketFreeSideData() (err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_packet_free_side_data").Call(
		uintptr(unsafe.Pointer(pkt)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
* Setup a new reference to the data described by a given packet
*
* If src is reference-counted, setup dst as a new reference to the
* buffer in src. Otherwise allocate a new buffer in dst and copy the
* data from src into it.
*
* All the other fields are copied from src.
*
* @see av_packet_unref
*
* @param dst Destination packet. Will be completely overwritten.
* @param src Source packet
*
* @return 0 on success, a negative AVERROR on error. On error, dst
*         will be blank (as if returned by av_packet_alloc()).
 */
//int av_packet_ref(AVPacket *dst, const AVPacket *src);
//未测试
func AvPacketRef(dst *AVPacket, src *AVPacket) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_packet_ref").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
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
* Wipe the packet.
*
* Unreference the buffer referenced by the packet and reset the
* remaining packet fields to their default values.
*
* @param pkt The packet to be unreferenced.
 */
//void av_packet_unref(AVPacket *pkt);
//未测试
func AvPacketUnref(dst *AVPacket, src *AVPacket) (err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_packet_unref").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
* Move every field in src to dst and reset src.
*
* @see av_packet_unref
*
* @param src Source packet, will be reset
* @param dst Destination packet
 */
//void av_packet_move_ref(AVPacket *dst, AVPacket *src);
//未测试
func AvPacketMoveRef(dst *AVPacket, src *AVPacket) (err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_packet_move_ref").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
* Copy only "properties" fields from src to dst.
*
* Properties for the purpose of this function are all the fields
* beside those related to the packet data (buf, data, size)
*
* @param dst Destination packet
* @param src Source packet
*
* @return 0 on success AVERROR on failure.
 */
//int av_packet_copy_props(AVPacket *dst, const AVPacket *src);
//未测试
func AvPacketCopyProps(dst *AVPacket, src *AVPacket) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_packet_copy_props").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
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
* Ensure the data described by a given packet is reference counted.
*
* @note This function does not ensure that the reference will be writable.
*       Use av_packet_make_writable instead for that purpose.
*
* @see av_packet_ref
* @see av_packet_make_writable
*
* @param pkt packet whose data should be made reference counted.
*
* @return 0 on success, a negative AVERROR on error. On failure, the
*         packet is unchanged.
 */
//int av_packet_make_refcounted(AVPacket *pkt);
//未测试
func (pkt *AVPacket) AvPacketMakeRefcounted() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_packet_make_refcounted").Call(
		uintptr(unsafe.Pointer(pkt)),
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
* Create a writable reference for the data described by a given packet,
* avoiding data copy if possible.
*
* @param pkt Packet whose data should be made writable.
*
* @return 0 on success, a negative AVERROR on failure. On failure, the
*         packet is unchanged.
 */
//int av_packet_make_writable(AVPacket *pkt);
//未测试
func (pkt *AVPacket) AvPacketMakeWritable() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_packet_make_writable").Call(
		uintptr(unsafe.Pointer(pkt)),
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
* Convert valid timing fields (timestamps / durations) in a packet from one
* timebase to another. Timestamps with unknown values (AV_NOPTS_VALUE) will be
* ignored.
*
* @param pkt packet on which the conversion will be performed
* @param tb_src source timebase, in which the timing fields in pkt are
*               expressed
* @param tb_dst destination timebase, to which the timing fields will be
*               converted
 */
//void av_packet_rescale_ts(AVPacket *pkt, AVRational tb_src, AVRational tb_dst);
//未测试
func (pkt *AVPacket) AvPacketRescaleTs(tb_src libavutil.AVRational, tb_dst libavutil.AVRational) (err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_packet_rescale_ts").Call(
		uintptr(unsafe.Pointer(pkt)),
		uintptr(unsafe.Pointer(&tb_src)),
		uintptr(unsafe.Pointer(&tb_dst)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}
