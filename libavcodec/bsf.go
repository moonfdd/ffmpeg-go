package libavcodec

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/libavutil"
	"syscall"
	"unsafe"
)

/**
 * @addtogroup lavc_core
 * @{
 */

type AVBSFInternal struct {
}

/**
 * The bitstream filter state.
 *
 * This struct must be allocated with av_bsf_alloc() and freed with
 * av_bsf_free().
 *
 * The fields in the struct will only be changed (by the caller or by the
 * filter) as described in their documentation, and are to be considered
 * immutable otherwise.
 */
type AVBSFContext struct {
	///**
	// * A class for logging and AVOptions
	// */
	//const AVClass *av_class;
	//
	///**
	// * The bitstream filter this context is an instance of.
	// */
	//const struct AVBitStreamFilter *filter;
	//
	///**
	// * Opaque libavcodec internal data. Must not be touched by the caller in any
	// * way.
	// */
	//AVBSFInternal *internal;
	//
	///**
	// * Opaque filter-specific private data. If filter->priv_class is non-NULL,
	// * this is an AVOptions-enabled struct.
	// */
	//void *priv_data;
	//
	///**
	// * Parameters of the input stream. This field is allocated in
	// * av_bsf_alloc(), it needs to be filled by the caller before
	// * av_bsf_init().
	// */
	//AVCodecParameters *par_in;
	//
	///**
	// * Parameters of the output stream. This field is allocated in
	// * av_bsf_alloc(), it is set by the filter in av_bsf_init().
	// */
	//AVCodecParameters *par_out;
	//
	///**
	// * The timebase used for the timestamps of the input packets. Set by the
	// * caller before av_bsf_init().
	// */
	//AVRational time_base_in;
	//
	///**
	// * The timebase used for the timestamps of the output packets. Set by the
	// * filter in av_bsf_init().
	// */
	//AVRational time_base_out;
}

type AVBitStreamFilter struct {
	//const char *name;
	//
	///**
	// * A list of codec ids supported by the filter, terminated by
	// * AV_CODEC_ID_NONE.
	// * May be NULL, in that case the bitstream filter works with any codec id.
	// */
	//const enum AVCodecID *codec_ids;
	//
	///**
	// * A class for the private data, used to declare bitstream filter private
	// * AVOptions. This field is NULL for bitstream filters that do not declare
	// * any options.
	// *
	// * If this field is non-NULL, the first member of the filter private data
	// * must be a pointer to AVClass, which will be set by libavcodec generic
	// * code to this class.
	// */
	//const AVClass *priv_class;
	//
	///*****************************************************************
	// * No fields below this line are part of the public API. They
	// * may not be used outside of libavcodec and can be changed and
	// * removed at will.
	// * New public fields should be added right above.
	// *****************************************************************
	// */
	//
	//int priv_data_size;
	//int (*init)(AVBSFContext *ctx);
	//int (*filter)(AVBSFContext *ctx, AVPacket *pkt);
	//void (*close)(AVBSFContext *ctx);
	//void (*flush)(AVBSFContext *ctx);
}

/**
 * @return a bitstream filter with the specified name or NULL if no such
 *         bitstream filter exists.
 */
//const AVBitStreamFilter *av_bsf_get_by_name(const char *name);
//未测试
func AvBsfGetByName(name ffcommon.FConstCharP) (res *AVBitStreamFilter, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_bsf_get_by_name").Call(
		uintptr(unsafe.Pointer(namep)),
	)
	if err != nil {
		//return
	}
	res = (*AVBitStreamFilter)(unsafe.Pointer(t))
	return
}

/**
* Iterate over all registered bitstream filters.
*
* @param opaque a pointer where libavcodec will store the iteration state. Must
*               point to NULL to start the iteration.
*
* @return the next registered bitstream filter or NULL when the iteration is
*         finished
 */
//const AVBitStreamFilter *av_bsf_iterate(void **opaque);
//未测试
func AvBsfIterate(opaque *ffcommon.FVoidP) (res *AVBitStreamFilter, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_bsf_iterate").Call(
		uintptr(unsafe.Pointer(opaque)),
	)
	if err != nil {
		//return
	}
	res = (*AVBitStreamFilter)(unsafe.Pointer(t))
	return
}

/**
* Allocate a context for a given bitstream filter. The caller must fill in the
* context parameters as described in the documentation and then call
* av_bsf_init() before sending any data to the filter.
*
* @param filter the filter for which to allocate an instance.
* @param ctx a pointer into which the pointer to the newly-allocated context
*            will be written. It must be freed with av_bsf_free() after the
*            filtering is done.
*
* @return 0 on success, a negative AVERROR code on failure
 */
//int av_bsf_alloc(const AVBitStreamFilter *filter, AVBSFContext **ctx);
//未测试
func (filter *AVBitStreamFilter) AvBsfAlloc(ctx **AVBSFContext) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_bsf_alloc").Call(
		uintptr(unsafe.Pointer(filter)),
		uintptr(unsafe.Pointer(&ctx)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
* Prepare the filter for use, after all the parameters and options have been
* set.
 */
//int av_bsf_init(AVBSFContext *ctx);
//未测试
func (ctx *AVBSFContext) AvBsfInit() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_bsf_init").Call(
		uintptr(unsafe.Pointer(ctx)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
* Submit a packet for filtering.
*
* After sending each packet, the filter must be completely drained by calling
* av_bsf_receive_packet() repeatedly until it returns AVERROR(EAGAIN) or
* AVERROR_EOF.
*
* @param pkt the packet to filter. The bitstream filter will take ownership of
* the packet and reset the contents of pkt. pkt is not touched if an error occurs.
* If pkt is empty (i.e. NULL, or pkt->data is NULL and pkt->side_data_elems zero),
* it signals the end of the stream (i.e. no more non-empty packets will be sent;
* sending more empty packets does nothing) and will cause the filter to output
* any packets it may have buffered internally.
*
* @return 0 on success. AVERROR(EAGAIN) if packets need to be retrieved from the
* filter (using av_bsf_receive_packet()) before new input can be consumed. Another
* negative AVERROR value if an error occurs.
 */
//int av_bsf_send_packet(AVBSFContext *ctx, AVPacket *pkt);
//未测试
func (ctx *AVBSFContext) AvBsfSendPacket(pkt *AVPacket) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_bsf_send_packet").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(pkt)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
* Retrieve a filtered packet.
*
* @param[out] pkt this struct will be filled with the contents of the filtered
*                 packet. It is owned by the caller and must be freed using
*                 av_packet_unref() when it is no longer needed.
*                 This parameter should be "clean" (i.e. freshly allocated
*                 with av_packet_alloc() or unreffed with av_packet_unref())
*                 when this function is called. If this function returns
*                 successfully, the contents of pkt will be completely
*                 overwritten by the returned data. On failure, pkt is not
*                 touched.
*
* @return 0 on success. AVERROR(EAGAIN) if more packets need to be sent to the
* filter (using av_bsf_send_packet()) to get more output. AVERROR_EOF if there
* will be no further output from the filter. Another negative AVERROR value if
* an error occurs.
*
* @note one input packet may result in several output packets, so after sending
* a packet with av_bsf_send_packet(), this function needs to be called
* repeatedly until it stops returning 0. It is also possible for a filter to
* output fewer packets than were sent to it, so this function may return
* AVERROR(EAGAIN) immediately after a successful av_bsf_send_packet() call.
 */
//int av_bsf_receive_packet(AVBSFContext *ctx, AVPacket *pkt);
//未测试
func (ctx *AVBSFContext) AvBsfReceivePacket(pkt *AVPacket) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_bsf_receive_packet").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(pkt)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
* Reset the internal bitstream filter state. Should be called e.g. when seeking.
 */
//void av_bsf_flush(AVBSFContext *ctx);
//未测试
func (ctx *AVBSFContext) AvBsfFlush(pkt *AVPacket) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_bsf_flush").Call(
		uintptr(unsafe.Pointer(ctx)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
* Free a bitstream filter context and everything associated with it; write NULL
* into the supplied pointer.
 */
//void av_bsf_free(AVBSFContext **ctx);
//未测试
func AvBsfFree(ctx **AVBSFContext) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_bsf_free").Call(
		uintptr(unsafe.Pointer(&ctx)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
* Get the AVClass for AVBSFContext. It can be used in combination with
* AV_OPT_SEARCH_FAKE_OBJ for examining options.
*
* @see av_opt_find().
 */
//const AVClass *av_bsf_get_class(void);
//未测试
func AvBsfGetClass() (res *libavutil.AVClass, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_bsf_get_class").Call()
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
* Structure for chain/list of bitstream filters.
* Empty list can be allocated by av_bsf_list_alloc().
 */
type AVBSFList struct {
}

/**
* Allocate empty list of bitstream filters.
* The list must be later freed by av_bsf_list_free()
* or finalized by av_bsf_list_finalize().
*
* @return Pointer to @ref AVBSFList on success, NULL in case of failure
 */
//AVBSFList *av_bsf_list_alloc(void);
//未测试
func AvBsfListAlloc() (res *AVBSFList, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_bsf_list_alloc").Call()
	if err != nil {
		//return
	}
	res = (*AVBSFList)(unsafe.Pointer(t))
	return
}

/**
* Free list of bitstream filters.
*
* @param lst Pointer to pointer returned by av_bsf_list_alloc()
 */
//void av_bsf_list_free(AVBSFList **lst);
//未测试
func AvBsfListFree(lst **AVBSFList) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_bsf_list_free").Call(
		uintptr(unsafe.Pointer(&lst)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
* Append bitstream filter to the list of bitstream filters.
*
* @param lst List to append to
* @param bsf Filter context to be appended
*
* @return >=0 on success, negative AVERROR in case of failure
 */
//int av_bsf_list_append(AVBSFList *lst, AVBSFContext *bsf);
//未测试
func (lst *AVBSFList) AvBsfListAppend(bsf *AVBSFContext) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_bsf_list_append").Call(
		uintptr(unsafe.Pointer(lst)),
		uintptr(unsafe.Pointer(bsf)),
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
* Construct new bitstream filter context given it's name and options
* and append it to the list of bitstream filters.
*
* @param lst      List to append to
* @param bsf_name Name of the bitstream filter
* @param options  Options for the bitstream filter, can be set to NULL
*
* @return >=0 on success, negative AVERROR in case of failure
 */
//int av_bsf_list_append2(AVBSFList *lst, const char * bsf_name, AVDictionary **options);
//未测试
func (lst *AVBSFList) AvBsfListAppend2(bsf_name ffcommon.FConstCharP, options **libavutil.AVDictionary) (res ffcommon.FInt, err error) {
	var t uintptr
	var bsf_namep *byte
	bsf_namep, err = syscall.BytePtrFromString(bsf_name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_bsf_list_append").Call(
		uintptr(unsafe.Pointer(lst)),
		uintptr(unsafe.Pointer(bsf_namep)),
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
* Finalize list of bitstream filters.
*
* This function will transform @ref AVBSFList to single @ref AVBSFContext,
* so the whole chain of bitstream filters can be treated as single filter
* freshly allocated by av_bsf_alloc().
* If the call is successful, @ref AVBSFList structure is freed and lst
* will be set to NULL. In case of failure, caller is responsible for
* freeing the structure by av_bsf_list_free()
*
* @param      lst Filter list structure to be transformed
* @param[out] bsf Pointer to be set to newly created @ref AVBSFContext structure
*                 representing the chain of bitstream filters
*
* @return >=0 on success, negative AVERROR in case of failure
 */
//int av_bsf_list_finalize(AVBSFList **lst, AVBSFContext **bsf);
//未测试
func AvBsfListFinalize(lst **AVBSFList, bsf **AVBSFContext) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_bsf_list_finalize").Call(
		uintptr(unsafe.Pointer(lst)),
		uintptr(unsafe.Pointer(bsf)),
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
* Parse string describing list of bitstream filters and create single
* @ref AVBSFContext describing the whole chain of bitstream filters.
* Resulting @ref AVBSFContext can be treated as any other @ref AVBSFContext freshly
* allocated by av_bsf_alloc().
*
* @param      str String describing chain of bitstream filters in format
*                 `bsf1[=opt1=val1:opt2=val2][,bsf2]`
* @param[out] bsf Pointer to be set to newly created @ref AVBSFContext structure
*                 representing the chain of bitstream filters
*
* @return >=0 on success, negative AVERROR in case of failure
 */
//int av_bsf_list_parse_str(const char *str, AVBSFContext **bsf);
//未测试
func AvBsfListParseStr(str ffcommon.FConstCharP, bsf **AVBSFContext) (res ffcommon.FInt, err error) {
	var t uintptr
	var strp *byte
	strp, err = syscall.BytePtrFromString(str)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_bsf_list_parse_str").Call(
		uintptr(unsafe.Pointer(strp)),
		uintptr(unsafe.Pointer(bsf)),
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
* Get null/pass-through bitstream filter.
*
* @param[out] bsf Pointer to be set to new instance of pass-through bitstream filter
*
* @return
 */
//int av_bsf_get_null_filter(AVBSFContext **bsf);
//未测试
func AvBsfGetNullFilter(bsf **AVBSFContext) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_bsf_get_null_filter").Call(
		uintptr(unsafe.Pointer(bsf)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}
