package ffconstant

/**
 * @addtogroup lavc_core
 * @{
 */

/**
 * Decoder can use draw_horiz_band callback.
 */
const AV_CODEC_CAP_DRAW_HORIZ_BAND = (1 << 0)

/**
 * Codec uses get_buffer() or get_encode_buffer() for allocating buffers and
 * supports custom allocators.
 * If not set, it might not use get_buffer() or get_encode_buffer() at all, or
 * use operations that assume the buffer was allocated by
 * avcodec_default_get_buffer2 or avcodec_default_get_encode_buffer.
 */
const AV_CODEC_CAP_DR1 = (1 << 1)
const AV_CODEC_CAP_TRUNCATED = (1 << 3)

/**
 * Encoder or decoder requires flushing with NULL input at the end in order to
 * give the complete and correct output.
 *
 * NOTE: If this flag is not set, the codec is guaranteed to never be fed with
 *       with NULL data. The user can still send NULL data to the public encode
 *       or decode function, but libavcodec will not pass it along to the codec
 *       unless this flag is set.
 *
 * Decoders:
 * The decoder has a non-zero delay and needs to be fed with avpkt->data=NULL,
 * avpkt->size=0 at the end to get the delayed data until the decoder no longer
 * returns frames.
 *
 * Encoders:
 * The encoder needs to be fed with NULL data at the end of encoding until the
 * encoder no longer returns data.
 *
 * NOTE: For encoders implementing the AVCodec.encode2() function, setting this
 *       flag also means that the encoder must set the pts and duration for
 *       each output packet. If this flag is not set, the pts and duration will
 *       be determined by libavcodec from the input frame.
 */
const AV_CODEC_CAP_DELAY = (1 << 5)

/**
 * Codec can be fed a final frame with a smaller size.
 * This can be used to prevent truncation of the last audio samples.
 */
const AV_CODEC_CAP_SMALL_LAST_FRAME = (1 << 6)

/**
 * Codec can output multiple frames per AVPacket
 * Normally demuxers return one frame at a time, demuxers which do not do
 * are connected to a parser to split what they return into proper frames.
 * This flag is reserved to the very rare category of codecs which have a
 * bitstream that cannot be split into frames without timeconsuming
 * operations like full decoding. Demuxers carrying such bitstreams thus
 * may return multiple frames in a packet. This has many disadvantages like
 * prohibiting stream copy in many cases thus it should only be considered
 * as a last resort.
 */
const AV_CODEC_CAP_SUBFRAMES = (1 << 8)

/**
 * Codec is experimental and is thus avoided in favor of non experimental
 * encoders
 */
const AV_CODEC_CAP_EXPERIMENTAL = (1 << 9)

/**
 * Codec should fill in channel configuration and samplerate instead of container
 */
const AV_CODEC_CAP_CHANNEL_CONF = (1 << 10)

/**
 * Codec supports frame-level multithreading.
 */
const AV_CODEC_CAP_FRAME_THREADS = (1 << 12)

/**
 * Codec supports slice-based (or partition-based) multithreading.
 */
const AV_CODEC_CAP_SLICE_THREADS = (1 << 13)

/**
 * Codec supports changed parameters at any point.
 */
const AV_CODEC_CAP_PARAM_CHANGE = (1 << 14)

/**
 * Codec supports multithreading through a method other than slice- or
 * frame-level multithreading. Typically this marks wrappers around
 * multithreading-capable external libraries.
 */
const AV_CODEC_CAP_OTHER_THREADS = (1 << 15)

//#if FF_API_AUTO_THREADS
const AV_CODEC_CAP_AUTO_THREADS = AV_CODEC_CAP_OTHER_THREADS

//#endif
/**
 * Audio encoder supports receiving a different number of samples in each call.
 */
const AV_CODEC_CAP_VARIABLE_FRAME_SIZE = (1 << 16)

/**
 * Decoder is not a preferred choice for probing.
 * This indicates that the decoder is not a good choice for probing.
 * It could for example be an expensive to spin up hardware decoder,
 * or it could simply not provide a lot of useful information about
 * the stream.
 * A decoder marked with this flag should only be used as last resort
 * choice for probing.
 */
const AV_CODEC_CAP_AVOID_PROBING = (1 << 17)

//#if FF_API_UNUSED_CODEC_CAPS
/**
 * Deprecated and unused. Use AVCodecDescriptor.props instead
 */
const AV_CODEC_CAP_INTRA_ONLY = 0x40000000

/**
 * Deprecated and unused. Use AVCodecDescriptor.props instead
 */
const AV_CODEC_CAP_LOSSLESS = 0x80000000

//#endif

/**
 * Codec is backed by a hardware implementation. Typically used to
 * identify a non-hwaccel hardware decoder. For information about hwaccels, use
 * avcodec_get_hw_config() instead.
 */
const AV_CODEC_CAP_HARDWARE = (1 << 18)

/**
 * Codec is potentially backed by a hardware implementation, but not
 * necessarily. This is used instead of AV_CODEC_CAP_HARDWARE, if the
 * implementation provides some sort of internal fallback.
 */
const AV_CODEC_CAP_HYBRID = (1 << 19)

/**
 * This codec takes the reordered_opaque field from input AVFrames
 * and returns it in the corresponding field in AVCodecContext after
 * encoding.
 */
const AV_CODEC_CAP_ENCODER_REORDERED_OPAQUE = (1 << 20)

/**
 * This encoder can be flushed using avcodec_flush_buffers(). If this flag is
 * not set, the encoder must be closed and reopened to ensure that no frames
 * remain pending.
 */
const AV_CODEC_CAP_ENCODER_FLUSH = (1 << 21)

const (
	/**
	 * The codec supports this format via the hw_device_ctx interface.
	 *
	 * When selecting this format, AVCodecContext.hw_device_ctx should
	 * have been set to a device of the specified type before calling
	 * avcodec_open2().
	 */
	AV_CODEC_HW_CONFIG_METHOD_HW_DEVICE_CTX = 0x01
	/**
	 * The codec supports this format via the hw_frames_ctx interface.
	 *
	 * When selecting this format for a decoder,
	 * AVCodecContext.hw_frames_ctx should be set to a suitable frames
	 * context inside the get_format() callback.  The frames context
	 * must have been created on a device of the specified type.
	 *
	 * When selecting this format for an encoder,
	 * AVCodecContext.hw_frames_ctx should be set to the context which
	 * will be used for the input frames before calling avcodec_open2().
	 */
	AV_CODEC_HW_CONFIG_METHOD_HW_FRAMES_CTX = 0x02
	/**
	 * The codec supports this format by some internal method.
	 *
	 * This format can be selected without any additional configuration -
	 * no device or frames context is required.
	 */
	AV_CODEC_HW_CONFIG_METHOD_INTERNAL = 0x04
	/**
	 * The codec supports this format by some ad-hoc method.
	 *
	 * Additional settings and/or function calls are required.  See the
	 * codec-specific documentation for details.  (Methods requiring
	 * this sort of configuration are deprecated and others should be
	 * used in preference.)
	 */
	AV_CODEC_HW_CONFIG_METHOD_AD_HOC = 0x08
)
