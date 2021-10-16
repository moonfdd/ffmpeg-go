package libavcodec

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"syscall"
	"unsafe"
)

/**
 * AVProfile.
 */
type AVProfile struct {
	//int profile;
	//const char *name; ///< short name for the profile
}

type AVCodecDefault struct {
}

//type AVCodecContext struct {
//}
type AVSubtitle3 struct {
}

//type AVPacket struct {
//}

/**
 * AVCodec.
 */
type AVCodec struct {
	//
	///**
	// * Name of the codec implementation.
	// * The name is globally unique among encoders and among decoders (but an
	// * encoder and a decoder can share the same name).
	// * This is the primary way to find a codec from the user perspective.
	// */
	//const char *name;
	///**
	// * Descriptive name for the codec, meant to be more human readable than name.
	// * You should use the NULL_IF_CONFIG_SMALL() macro to define it.
	// */
	//const char *long_name;
	//enum AVMediaType type;
	//enum AVCodecID id;
	///**
	// * Codec capabilities.
	// * see AV_CODEC_CAP_*
	// */
	//int capabilities;
	//const AVRational *supported_framerates; ///< array of supported framerates, or NULL if any, array is terminated by {0,0}
	//const enum AVPixelFormat *pix_fmts;     ///< array of supported pixel formats, or NULL if unknown, array is terminated by -1
	//const int *supported_samplerates;       ///< array of supported audio samplerates, or NULL if unknown, array is terminated by 0
	//const enum AVSampleFormat *sample_fmts; ///< array of supported sample formats, or NULL if unknown, array is terminated by -1
	//const uint64_t *channel_layouts;         ///< array of support channel layouts, or NULL if unknown. array is terminated by 0
	//uint8_t max_lowres;                     ///< maximum value for lowres supported by the decoder
	//const AVClass *priv_class;              ///< AVClass for the private context
	//const AVProfile *profiles;              ///< array of recognized profiles, or NULL if unknown, array is terminated by {FF_PROFILE_UNKNOWN}
	//
	///**
	// * Group name of the codec implementation.
	// * This is a short symbolic name of the wrapper backing this codec. A
	// * wrapper uses some kind of external implementation for the codec, such
	// * as an external library, or a codec implementation provided by the OS or
	// * the hardware.
	// * If this field is NULL, this is a builtin, libavcodec native codec.
	// * If non-NULL, this will be the suffix in AVCodec.name in most cases
	// * (usually AVCodec.name will be of the form "<codec_name>_<wrapper_name>").
	// */
	//const char *wrapper_name;
	//
	///*****************************************************************
	// * No fields below this line are part of the public API. They
	// * may not be used outside of libavcodec and can be changed and
	// * removed at will.
	// * New public fields should be added right above.
	// *****************************************************************
	// */
	//int priv_data_size;
	//#if FF_API_NEXT
	//struct AVCodec *next;
	//#endif
	///**
	// * @name Frame-level threading support functions
	// * @{
	// */
	///**
	// * Copy necessary context variables from a previous thread context to the current one.
	// * If not defined, the next thread will start automatically; otherwise, the codec
	// * must call ff_thread_finish_setup().
	// *
	// * dst and src will (rarely) point to the same context, in which case memcpy should be skipped.
	// */
	//int (*update_thread_context)(struct AVCodecContext *dst, const struct AVCodecContext *src);
	///** @} */
	//
	///**
	// * Private codec-specific defaults.
	// */
	//const AVCodecDefault *defaults;
	//
	///**
	// * Initialize codec static data, called from av_codec_iterate().
	// *
	// * This is not intended for time consuming operations as it is
	// * run for every codec regardless of that codec being used.
	// */
	//void (*init_static_data)(struct AVCodec *codec);
	//
	//int (*init)(struct AVCodecContext *);
	//int (*encode_sub)(struct AVCodecContext *, uint8_t *buf, int buf_size,
	//const struct AVSubtitle *sub);
	///**
	// * Encode data to an AVPacket.
	// *
	// * @param      avctx          codec context
	// * @param      avpkt          output AVPacket
	// * @param[in]  frame          AVFrame containing the raw data to be encoded
	// * @param[out] got_packet_ptr encoder sets to 0 or 1 to indicate that a
	// *                            non-empty packet was returned in avpkt.
	// * @return 0 on success, negative error code on failure
	// */
	//int (*encode2)(struct AVCodecContext *avctx, struct AVPacket *avpkt,
	//const struct AVFrame *frame, int *got_packet_ptr);
	///**
	// * Decode picture or subtitle data.
	// *
	// * @param      avctx          codec context
	// * @param      outdata        codec type dependent output struct
	// * @param[out] got_frame_ptr  decoder sets to 0 or 1 to indicate that a
	// *                            non-empty frame or subtitle was returned in
	// *                            outdata.
	// * @param[in]  avpkt          AVPacket containing the data to be decoded
	// * @return amount of bytes read from the packet on success, negative error
	// *         code on failure
	// */
	//int (*decode)(struct AVCodecContext *avctx, void *outdata,
	//int *got_frame_ptr, struct AVPacket *avpkt);
	//int (*close)(struct AVCodecContext *);
	///**
	// * Encode API with decoupled frame/packet dataflow. This function is called
	// * to get one output packet. It should call ff_encode_get_frame() to obtain
	// * input data.
	// */
	//int (*receive_packet)(struct AVCodecContext *avctx, struct AVPacket *avpkt);
	//
	///**
	// * Decode API with decoupled packet/frame dataflow. This function is called
	// * to get one output frame. It should call ff_decode_get_packet() to obtain
	// * input data.
	// */
	//int (*receive_frame)(struct AVCodecContext *avctx, struct AVFrame *frame);
	///**
	// * Flush buffers.
	// * Will be called when seeking
	// */
	//void (*flush)(struct AVCodecContext *);
	///**
	// * Internal codec capabilities.
	// * See FF_CODEC_CAP_* in internal.h
	// */
	//int caps_internal;
	//
	///**
	// * Decoding only, a comma-separated list of bitstream filters to apply to
	// * packets before decoding.
	// */
	//const char *bsfs;
	//
	///**
	// * Array of pointers to hardware configurations supported by the codec,
	// * or NULL if no hardware supported.  The array is terminated by a NULL
	// * pointer.
	// *
	// * The user can only access this field via avcodec_get_hw_config().
	// */
	//const struct AVCodecHWConfigInternal *const *hw_configs;
	//
	///**
	// * List of supported codec_tags, terminated by FF_CODEC_TAGS_END.
	// */
	//const uint32_t *codec_tags;
}

/**
 * Iterate over all registered codecs.
 *
 * @param opaque a pointer where libavcodec will store the iteration state. Must
 *               point to NULL to start the iteration.
 *
 * @return the next registered codec or NULL when the iteration is
 *         finished
 */
//const AVCodec *av_codec_iterate(void **opaque);
//未测试
func AvCodecIterate(opaque *ffcommon.FVoidP) (res *AVCodec, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_codec_iterate").Call(
		uintptr(unsafe.Pointer(opaque)),
	)
	if err != nil {
		//return
	}
	res = (*AVCodec)(unsafe.Pointer(t))
	return
}

/**
 * Find a registered decoder with a matching codec ID.
 *
 * @param id AVCodecID of the requested decoder
 * @return A decoder if one was found, NULL otherwise.
 */
//AVCodec *avcodec_find_decoder(enum AVCodecID id);
//未测试
func AvcodecFindDecoder(id ffconstant.AVCodecID) (res *AVCodec, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("avcodec_find_decoder").Call(
		uintptr(id),
	)
	if err != nil {
		//return
	}
	res = (*AVCodec)(unsafe.Pointer(t))
	return
}

/**
 * Find a registered decoder with the specified name.
 *
 * @param name name of the requested decoder
 * @return A decoder if one was found, NULL otherwise.
 */
//AVCodec *avcodec_find_decoder_by_name(const char *name);
//未测试
func AvcodecFindDecoderByName(name ffcommon.FConstCharP) (res *AVCodec, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("avcodec_find_decoder_by_name").Call(
		uintptr(unsafe.Pointer(namep)),
	)
	if err != nil {
		//return
	}
	res = (*AVCodec)(unsafe.Pointer(t))
	return
}

/**
 * Find a registered encoder with a matching codec ID.
 *
 * @param id AVCodecID of the requested encoder
 * @return An encoder if one was found, NULL otherwise.
 */
//AVCodec *avcodec_find_encoder(enum AVCodecID id);
//未测试
func AvcodecFindEncoder(id ffconstant.AVCodecID) (res *AVCodec, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("avcodec_find_encoder").Call(
		uintptr(id),
	)
	if err != nil {
		//return
	}
	res = (*AVCodec)(unsafe.Pointer(t))
	return
}

/**
 * Find a registered encoder with the specified name.
 *
 * @param name name of the requested encoder
 * @return An encoder if one was found, NULL otherwise.
 */
//AVCodec *avcodec_find_encoder_by_name(const char *name);
//未测试
func AvcodecFindEncoderByName(name ffcommon.FConstCharP) (res *AVCodec, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("avcodec_find_encoder_by_name").Call(
		uintptr(unsafe.Pointer(namep)),
	)
	if err != nil {
		//return
	}
	res = (*AVCodec)(unsafe.Pointer(t))
	return
}

/**
 * @return a non-zero number if codec is an encoder, zero otherwise
 */
//int av_codec_is_encoder(const AVCodec *codec);
//未测试
func (codec *AVCodec) av_codec_is_encoder() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_codec_is_encoder").Call(
		uintptr(unsafe.Pointer(codec)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * @return a non-zero number if codec is a decoder, zero otherwise
 */
//int av_codec_is_decoder(const AVCodec *codec);
//未测试
func (codec *AVCodec) AvCodecIsDecoder() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_codec_is_decoder").Call(
		uintptr(unsafe.Pointer(codec)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

type AVCodecHWConfig struct {
	/**
	  // * For decoders, a hardware pixel format which that decoder may be
	  // * able to decode to if suitable hardware is available.
	  // *
	  // * For encoders, a pixel format which the encoder may be able to
	  // * accept.  If set to AV_PIX_FMT_NONE, this applies to all pixel
	  // * formats supported by the codec.
	  // */
	//enum AVPixelFormat pix_fmt;
	///**
	// * Bit set of AV_CODEC_HW_CONFIG_METHOD_* flags, describing the possible
	// * setup methods which can be used with this configuration.
	// */
	//int methods;
	///**
	// * The device type associated with the configuration.
	// *
	// * Must be set for AV_CODEC_HW_CONFIG_METHOD_HW_DEVICE_CTX and
	// * AV_CODEC_HW_CONFIG_METHOD_HW_FRAMES_CTX, otherwise unused.
	// */
	//enum AVHWDeviceType device_type;
}

/**
 * Retrieve supported hardware configurations for a codec.
 *
 * Values of index from zero to some maximum return the indexed configuration
 * descriptor; all other values return NULL.  If the codec does not support
 * any hardware configurations then it will always return NULL.
 */
//const AVCodecHWConfig *avcodec_get_hw_config(const AVCodec *codec, int index);
//未测试
func (codec *AVCodec) AvcodecGetHwConfig(index ffcommon.FInt) (res *AVCodecHWConfig, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("avcodec_get_hw_config").Call(
		uintptr(unsafe.Pointer(codec)),
		uintptr(index),
	)
	if err != nil {
		//return
	}
	res = (*AVCodecHWConfig)(unsafe.Pointer(t))
	return
}
