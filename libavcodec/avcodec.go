package libavcodec

import (
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/ffmpeg-go/libavutil"
)

/*
* copyright (c) 2001 Fabrice Bellard
*
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

//#ifndef AVCODEC_AVCODEC_H
//#define AVCODEC_AVCODEC_H

/**
* @file
* @ingroup libavc
* Libavcodec external API header
 */

//#include <errno.h>
//#include "../libavutil/samplefmt.h"
//#include "../libavutil/attributes.h"
//#include "../libavutil/avutil.h"
//#include "../libavutil/buffer.h"
//#include "../libavutil/cpu.h"
//#include "../libavutil/channel_layout.h"
//#include "../libavutil/dict.h"
//#include "../libavutil/frame.h"
//#include "../libavutil/hwcontext.h"
//#include "../libavutil/log.h"
//#include "../libavutil/pixfmt.h"
//#include "../libavutil/rational.h"
//
//#include "bsf.h"
//#include "codec.h"
//#include "codec_desc.h"
//#include "codec_par.h"
//#include "codec_id.h"
//#include "packet.h"
//#include "version.h"

/**
* @defgroup libavc libavcodec
* Encoding/Decoding Library
*
* @{
*
* @defgroup lavc_decoding Decoding
* @{
* @}
*
* @defgroup lavc_encoding Encoding
* @{
* @}
*
* @defgroup lavc_codec Codecs
* @{
* @defgroup lavc_codec_native Native Codecs
* @{
* @}
* @defgroup lavc_codec_wrappers External library wrappers
* @{
* @}
* @defgroup lavc_codec_hwaccel Hardware Accelerators bridge
* @{
* @}
* @}
* @defgroup lavc_internal Internal
* @{
* @}
* @}
 */

/**
* @ingroup libavc
* @defgroup lavc_encdec send/receive encoding and decoding API overview
* @{
*
* The avcodec_send_packet()/avcodec_receive_frame()/avcodec_send_frame()/
* avcodec_receive_packet() functions provide an encode/decode API, which
* decouples input and output.
*
* The API is very similar for encoding/decoding and audio/video, and works as
* follows:
* - Set up and open the AVCodecContext as usual.
* - Send valid input:
*   - For decoding, call avcodec_send_packet() to give the decoder raw
*     compressed data in an AVPacket.
*   - For encoding, call avcodec_send_frame() to give the encoder an AVFrame
*     containing uncompressed audio or video.
*
*   In both cases, it is recommended that AVPackets and AVFrames are
*   refcounted, or libavcodec might have to copy the input data. (libavformat
*   always returns refcounted AVPackets, and av_frame_get_buffer() allocates
*   refcounted AVFrames.)
* - Receive output in a loop. Periodically call one of the avcodec_receive_*()
*   functions and process their output:
*   - For decoding, call avcodec_receive_frame(). On success, it will return
*     an AVFrame containing uncompressed audio or video data.
*   - For encoding, call avcodec_receive_packet(). On success, it will return
*     an AVPacket with a compressed frame.
*
*   Repeat this call until it returns AVERROR(EAGAIN) or an error. The
*   AVERROR(EAGAIN) return value means that new input data is required to
*   return new output. In this case, continue with sending input. For each
*   input frame/packet, the codec will typically return 1 output frame/packet,
*   but it can also be 0 or more than 1.
*
* At the beginning of decoding or encoding, the codec might accept multiple
* input frames/packets without returning a frame, until its internal buffers
* are filled. This situation is handled transparently if you follow the steps
* outlined above.
*
* In theory, sending input can result in EAGAIN - this should happen only if
* not all output was received. You can use this to structure alternative decode
* or encode loops other than the one suggested above. For example, you could
* try sending new input on each iteration, and try to receive output if that
* returns EAGAIN.
*
* End of stream situations. These require "flushing" (aka draining) the codec,
* as the codec might buffer multiple frames or packets internally for
* performance or out of necessity (consider B-frames).
* This is handled as follows:
* - Instead of valid input, send NULL to the avcodec_send_packet() (decoding)
*   or avcodec_send_frame() (encoding) functions. This will enter draining
*   mode.
* - Call avcodec_receive_frame() (decoding) or avcodec_receive_packet()
*   (encoding) in a loop until AVERROR_EOF is returned. The functions will
*   not return AVERROR(EAGAIN), unless you forgot to enter draining mode.
* - Before decoding can be resumed again, the codec has to be reset with
*   avcodec_flush_buffers().
*
* Using the API as outlined above is highly recommended. But it is also
* possible to call functions outside of this rigid schema. For example, you can
* call avcodec_send_packet() repeatedly without calling
* avcodec_receive_frame(). In this case, avcodec_send_packet() will succeed
* until the codec's internal buffer has been filled up (which is typically of
* size 1 per output frame, after initial input), and then reject input with
* AVERROR(EAGAIN). Once it starts rejecting input, you have no choice but to
* read at least some output.
*
* Not all codecs will follow a rigid and predictable dataflow; the only
* guarantee is that an AVERROR(EAGAIN) return value on a send/receive call on
* one end implies that a receive/send call on the other end will succeed, or
* at least will not fail with AVERROR(EAGAIN). In general, no codec will
* permit unlimited buffering of input or output.
*
* This API replaces the following legacy functions:
* - avcodec_decode_video2() and avcodec_decode_audio4():
*   Use avcodec_send_packet() to feed input to the decoder, then use
*   avcodec_receive_frame() to receive decoded frames after each packet.
*   Unlike with the old video decoding API, multiple frames might result from
*   a packet. For audio, splitting the input packet into frames by partially
*   decoding packets becomes transparent to the API user. You never need to
*   feed an AVPacket to the API twice (unless it is rejected with AVERROR(EAGAIN) - then
*   no data was read from the packet).
*   Additionally, sending a flush/draining packet is required only once.
* - avcodec_encode_video2()/avcodec_encode_audio2():
*   Use avcodec_send_frame() to feed input to the encoder, then use
*   avcodec_receive_packet() to receive encoded packets.
*   Providing user-allocated buffers for avcodec_receive_packet() is not
*   possible.
* - The new API does not handle subtitles yet.
*
* Mixing new and old function calls on the same AVCodecContext is not allowed,
* and will result in undefined behavior.
*
* Some codecs might require using the new API; using the old API will return
* an error when calling it. All codecs support the new API.
*
* A codec is not allowed to return AVERROR(EAGAIN) for both sending and receiving. This
* would be an invalid state, which could put the codec user into an endless
* loop. The API has no concept of time either: it cannot happen that trying to
* do avcodec_send_packet() results in AVERROR(EAGAIN), but a repeated call 1 second
* later accepts the packet (with no other receive/flush API calls involved).
* The API is a strict state machine, and the passage of time is not supposed
* to influence it. Some timing-dependent behavior might still be deemed
* acceptable in certain cases. But it must never result in both send/receive
* returning EAGAIN at the same time at any point. It must also absolutely be
* avoided that the current state is "unstable" and can "flip-flop" between
* the send/receive APIs allowing progress. For example, it's not allowed that
* the codec randomly decides that it actually wants to consume a packet now
* instead of returning a frame, after it just returned AVERROR(EAGAIN) on an
* avcodec_send_packet() call.
* @}
 */

/**
* @defgroup lavc_core Core functions/structures.
* @ingroup libavc
*
* Basic definitions, functions for querying libavcodec capabilities,
* allocating core structures, etc.
* @{
 */

/**
* @ingroup lavc_decoding
* Required number of additionally allocated bytes at the end of the input bitstream for decoding.
* This is mainly needed because some optimized bitstream readers read
* 32 or 64 bit at once and could read over the end.<br>
* Note: If the first 23 bits of the additional bytes are not 0, then damaged
* MPEG bitstreams could cause overread and segfault.
 */
const AV_INPUT_BUFFER_PADDING_SIZE = 64

/**
* @ingroup lavc_encoding
* minimum encoding buffer size
* Used to avoid some checks during header writing.
 */
const AV_INPUT_BUFFER_MIN_SIZE = 16384

/**
* @ingroup lavc_decoding
 */
type AVDiscard int32

const (
	/* We leave some space between them for extensions (drop some
	 * keyframes for intra-only or drop just some bidir frames). */
	AVDISCARD_NONE     = -16 ///< discard nothing
	AVDISCARD_DEFAULT  = 0   ///< discard useless packets like 0 size packets in avi
	AVDISCARD_NONREF   = 8   ///< discard all non reference
	AVDISCARD_BIDIR    = 16  ///< discard all bidirectional frames
	AVDISCARD_NONINTRA = 24  ///< discard all non intra frames
	AVDISCARD_NONKEY   = 32  ///< discard all frames except keyframes
	AVDISCARD_ALL      = 48  ///< discard all
)

type AVAudioServiceType int32

const (
	AV_AUDIO_SERVICE_TYPE_MAIN              = 0
	AV_AUDIO_SERVICE_TYPE_EFFECTS           = 1
	AV_AUDIO_SERVICE_TYPE_VISUALLY_IMPAIRED = 2
	AV_AUDIO_SERVICE_TYPE_HEARING_IMPAIRED  = 3
	AV_AUDIO_SERVICE_TYPE_DIALOGUE          = 4
	AV_AUDIO_SERVICE_TYPE_COMMENTARY        = 5
	AV_AUDIO_SERVICE_TYPE_EMERGENCY         = 6
	AV_AUDIO_SERVICE_TYPE_VOICE_OVER        = 7
	AV_AUDIO_SERVICE_TYPE_KARAOKE           = 8
	AV_AUDIO_SERVICE_TYPE_NB                = AV_AUDIO_SERVICE_TYPE_KARAOKE + 1 ///< Not part of ABI
)

/**
* @ingroup lavc_encoding
 */

type RcOverride struct {
	StartFrame    ffcommon.FInt
	EndFrame      ffcommon.FInt
	Qscale        ffcommon.FInt // If this is 0 then quality_factor will be used instead.
	QualityFactor ffcommon.FFloat
}

/* encoding support
These flags can be passed in AVCodecContext.flags before initialization.
Note: Not everything is supported yet.
*/

/**
* Allow decoders to produce frames with data planes that are not aligned
* to CPU requirements (e.g. due to cropping).
 */
const AV_CODEC_FLAG_UNALIGNED = (1 << 0)

/**
* Use fixed qscale.
 */
const AV_CODEC_FLAG_QSCALE = (1 << 1)

/**
* 4 MV per MB allowed / advanced prediction for H.263.
 */
const AV_CODEC_FLAG_4MV = (1 << 2)

/**
* Output even those frames that might be corrupted.
 */
const AV_CODEC_FLAG_OUTPUT_CORRUPT = (1 << 3)

/**
* Use qpel MC.
 */
const AV_CODEC_FLAG_QPEL = (1 << 4)

/**
* Don't output frames whose parameters differ from first
* decoded frame in stream.
 */
const AV_CODEC_FLAG_DROPCHANGED = (1 << 5)

/**
* Use internal 2pass ratecontrol in first pass mode.
 */
const AV_CODEC_FLAG_PASS1 = (1 << 9)

/**
* Use internal 2pass ratecontrol in second pass mode.
 */
const AV_CODEC_FLAG_PASS2 = (1 << 10)

/**
* loop filter.
 */
const AV_CODEC_FLAG_LOOP_FILTER = (1 << 11)

/**
* Only decode/encode grayscale.
 */
const AV_CODEC_FLAG_GRAY = (1 << 13)

/**
* error[?] variables will be set during encoding.
 */
const AV_CODEC_FLAG_PSNR = (1 << 15)

/**
* Input bitstream might be truncated at a random location
* instead of only at frame boundaries.
 */
const AV_CODEC_FLAG_TRUNCATED = (1 << 16)

/**
* Use interlaced DCT.
 */
const AV_CODEC_FLAG_INTERLACED_DCT = (1 << 18)

/**
* Force low delay.
 */
const AV_CODEC_FLAG_LOW_DELAY = (1 << 19)

/**
* Place global headers in extradata instead of every keyframe.
 */
const AV_CODEC_FLAG_GLOBAL_HEADER = (1 << 22)

/**
* Use only bitexact stuff (except (I)DCT).
 */
const AV_CODEC_FLAG_BITEXACT = (1 << 23)

/* Fx : Flag for H.263+ extra options */
/**
* H.263 advanced intra coding / MPEG-4 AC prediction
 */
const AV_CODEC_FLAG_AC_PRED = (1 << 24)

/**
* interlaced motion estimation
 */
const AV_CODEC_FLAG_INTERLACED_ME = (1 << 29)
const AV_CODEC_FLAG_CLOSED_GOP = (1 << 31)

/**
* Allow non spec compliant speedup tricks.
 */
const AV_CODEC_FLAG2_FAST = (1 << 0)

/**
* Skip bitstream encoding.
 */
const AV_CODEC_FLAG2_NO_OUTPUT = (1 << 2)

/**
* Place global headers at every keyframe instead of in extradata.
 */
const AV_CODEC_FLAG2_LOCAL_HEADER = (1 << 3)

/**
* timecode is in drop frame format. DEPRECATED!!!!
 */
const AV_CODEC_FLAG2_DROP_FRAME_TIMECODE = (1 << 13)

/**
* Input bitstream might be truncated at a packet boundaries
* instead of only at frame boundaries.
 */
const AV_CODEC_FLAG2_CHUNKS = (1 << 15)

/**
* Discard cropping information from SPS.
 */
const AV_CODEC_FLAG2_IGNORE_CROP = (1 << 16)

/**
* Show all frames before the first keyframe
 */
const AV_CODEC_FLAG2_SHOW_ALL = (1 << 22)

/**
* Export motion vectors through frame side data
 */
const AV_CODEC_FLAG2_EXPORT_MVS = (1 << 28)

/**
* Do not skip samples and export skip information as frame side data
 */
const AV_CODEC_FLAG2_SKIP_MANUAL = (1 << 29)

/**
* Do not reset ASS ReadOrder field on flush (subtitles decoding)
 */
const AV_CODEC_FLAG2_RO_FLUSH_NOOP = (1 << 30)

/* Unsupported options :
*              Syntax Arithmetic coding (SAC)
*              Reference Picture Selection
*              Independent Segment Decoding */
/* /Fx */
/* codec capabilities */

/* Exported side data.
These flags can be passed in AVCodecContext.export_side_data before initialization.
*/
/**
* Export motion vectors through frame side data
 */
const AV_CODEC_EXPORT_DATA_MVS = (1 << 0)

/**
* Export encoder Producer Reference Time through packet side data
 */
const AV_CODEC_EXPORT_DATA_PRFT = (1 << 1)

/**
* Decoding only.
* Export the AVVideoEncParams structure through frame side data.
 */
const AV_CODEC_EXPORT_DATA_VIDEO_ENC_PARAMS = (1 << 2)

/**
* Decoding only.
* Do not apply film grain, export it instead.
 */
const AV_CODEC_EXPORT_DATA_FILM_GRAIN = (1 << 3)

/**
* Pan Scan area.
* This specifies the area which should be displayed.
* Note there may be multiple such areas for one frame.
 */
type AVPanScan struct {

	/**
	 * id
	 * - encoding: Set by user.
	 * - decoding: Set by libavcodec.
	 */
	Id ffcommon.FInt

	/**
	 * width and height in 1/16 pel
	 * - encoding: Set by user.
	 * - decoding: Set by libavcodec.
	 */
	Width  ffcommon.FInt
	Height ffcommon.FInt

	/**
	 * position of the top left corner in 1/16 pel for up to 3 fields/frames
	 * - encoding: Set by user.
	 * - decoding: Set by libavcodec.
	 */
	Position [3][2]ffcommon.FInt16T
}

/**
* This structure describes the bitrate properties of an encoded bitstream. It
* roughly corresponds to a subset the VBV parameters for MPEG-2 or HRD
* parameters for H.264/HEVC.
 */
type AVCPBProperties struct {

	/**
	 * Maximum bitrate of the stream, in bits per second.
	 * Zero if unknown or unspecified.
	 */
	//#if FF_API_UNSANITIZED_BITRATES
	MaxBitrate ffcommon.FInt
	//#else
	//int64_t max_bitrate;
	//#endif
	/**
	 * Minimum bitrate of the stream, in bits per second.
	 * Zero if unknown or unspecified.
	 */
	//#if FF_API_UNSANITIZED_BITRATES
	MinBitrate ffcommon.FInt
	//#else
	//int64_t min_bitrate;
	//#endif
	/**
	 * Average bitrate of the stream, in bits per second.
	 * Zero if unknown or unspecified.
	 */
	//#if FF_API_UNSANITIZED_BITRATES
	AvgBitrate ffcommon.FInt
	//#else
	//int64_t avg_bitrate;
	//#endif

	/**
	 * The size of the buffer to which the ratecontrol is applied, in bits.
	 * Zero if unknown or unspecified.
	 */
	BufferSize ffcommon.FInt

	/**
	 * The delay between the time the packet this structure is associated with
	 * is received and the time when it should be decoded, in periods of a 27MHz
	 * clock.
	 *
	 * UINT64_MAX when unknown or unspecified.
	 */
	VbvDelay ffcommon.FUint64T
}

/**
* This structure supplies correlation between a packet timestamp and a wall clock
* production time. The definition follows the Producer Reference Time ('prft')
* as defined in ISO/IEC 14496-12
 */
type AVProducerReferenceTime struct {

	/**
	 * A UTC timestamp, in microseconds, since Unix epoch (e.g, av_gettime()).
	 */
	Wallclock ffcommon.FInt64T
	Flags     ffcommon.FInt
}

/**
* The decoder will keep a reference to the frame and may reuse it later.
 */
const AV_GET_BUFFER_FLAG_REF = (1 << 0)

/**
* The encoder will keep a reference to the packet and may reuse it later.
 */
const AV_GET_ENCODE_BUFFER_FLAG_REF = (1 << 0)

type AVCodecInternal struct {
}

const FF_COMPLIANCE_VERY_STRICT = 2 ///< Strictly conform to an older more strict version of the spec or reference software.
const FF_COMPLIANCE_STRICT = 1      ///< Strictly conform to all the things in the spec no matter what consequences.
const FF_COMPLIANCE_NORMAL = 0
const FF_COMPLIANCE_UNOFFICIAL = -1   ///< Allow unofficial extensions
const FF_COMPLIANCE_EXPERIMENTAL = -2 ///< Allow nonstandardized experimental things.

/**
* main external API structure.
* New fields can be added to the end with minor version bumps.
* Removal, reordering and changes to existing fields require a major
* version bump.
* You can use AVOptions (av_opt* / av_set/get*()) to access these fields from user
* applications.
* The name string for AVOptions options matches the associated command line
* parameter name and can be found in libavcodec/options_table.h
* The AVOption/command line parameter names differ in some cases from the C
* structure field names for historic reasons or brevity.
* sizeof(AVCodecContext) must not be used outside libav*.
 */
//type AVClass=libavutil.AVClass
type AVCodecContext struct {

	/**
	 * information on struct for av_log
	 * - set by avcodec_alloc_context3
	 */
	AvClass        *AVClass
	LogLevelOffset ffcommon.FInt

	CodecType AVMediaType /* see AVMEDIA_TYPE_xxx */
	Codec     *AVCodec
	CodecId   AVCodecID /* see AV_CODEC_ID_xxx */

	/**
	 * fourcc (LSB first, so "ABCD" -> ('D'<<24) + ('C'<<16) + ('B'<<8) + 'A').
	 * This is used to work around some encoder bugs.
	 * A demuxer should set this to what is stored in the field used to identify the codec.
	 * If there are multiple such fields in a container then the demuxer should choose the one
	 * which maximizes the information about the used codec.
	 * If the codec tag field in a container is larger than 32 bits then the demuxer should
	 * remap the longer ID to 32 bits with a table or other structure. Alternatively a new
	 * extra_codec_tag + size could be added but for this a clear advantage must be demonstrated
	 * first.
	 * - encoding: Set by user, if not then the default based on codec_id will be used.
	 * - decoding: Set by user, will be converted to uppercase by libavcodec during init.
	 */
	CodecTag ffcommon.FUnsignedInt

	PrivData ffcommon.FVoidP

	/**
	 * Private context used for internal data.
	 *
	 * Unlike priv_data, this is not codec-specific. It is used in general
	 * libavcodec functions.
	 */
	Internal *AVCodecInternal

	/**
	 * Private data of the user, can be used to carry app specific stuff.
	 * - encoding: Set by user.
	 * - decoding: Set by user.
	 */
	Opaque ffcommon.FVoidP

	/**
	 * the average bitrate
	 * - encoding: Set by user; unused for constant quantizer encoding.
	 * - decoding: Set by user, may be overwritten by libavcodec
	 *             if this info is available in the stream
	 */
	BitRate ffcommon.FInt64T

	/**
	 * number of bits the bitstream is allowed to diverge from the reference.
	 *           the reference can be CBR (for CBR pass1) or VBR (for pass2)
	 * - encoding: Set by user; unused for constant quantizer encoding.
	 * - decoding: unused
	 */
	BitRateTolerance ffcommon.FInt

	/**
	 * Global quality for codecs which cannot change it per frame.
	 * This should be proportional to MPEG-1/2/4 qscale.
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	GlobalQuality ffcommon.FInt

	/**
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	CompressionLevel ffcommon.FInt
	//const FF_COMPRESSION_DEFAULT= -1

	/**
	 * AV_CODEC_FLAG_*.
	 * - encoding: Set by user.
	 * - decoding: Set by user.
	 */
	Flags ffcommon.FInt

	/**
	 * AV_CODEC_FLAG2_*
	 * - encoding: Set by user.
	 * - decoding: Set by user.
	 */
	Flags2 ffcommon.FInt

	/**
	 * some codecs need / can use extradata like Huffman tables.
	 * MJPEG: Huffman tables
	 * rv10: additional flags
	 * MPEG-4: global headers (they can be in the bitstream or here)
	 * The allocated memory should be AV_INPUT_BUFFER_PADDING_SIZE bytes larger
	 * than extradata_size to avoid problems if it is read with the bitstream reader.
	 * The bytewise contents of extradata must not depend on the architecture or CPU endianness.
	 * Must be allocated with the av_malloc() family of functions.
	 * - encoding: Set/allocated/freed by libavcodec.
	 * - decoding: Set/allocated/freed by user.
	 */
	Extradata     *ffcommon.FUint8T
	ExtradataSize ffcommon.FInt

	/**
	 * This is the fundamental unit of time (in seconds) in terms
	 * of which frame timestamps are represented. For fixed-fps content,
	 * timebase should be 1/framerate and timestamp increments should be
	 * identically 1.
	 * This often, but not always is the inverse of the frame rate or field rate
	 * for video. 1/time_base is not the average frame rate if the frame rate is not
	 * constant.
	 *
	 * Like containers, elementary streams also can store timestamps, 1/time_base
	 * is the unit in which these timestamps are specified.
	 * As example of such codec time base see ISO/IEC 14496-2:2001(E)
	 * vop_time_increment_resolution and fixed_vop_rate
	 * (fixed_vop_rate == 0 implies that it is different from the framerate)
	 *
	 * - encoding: MUST be set by user.
	 * - decoding: the use of this field for decoding is deprecated.
	 *             Use framerate instead.
	 */
	TimeBase AVRational

	/**
	 * For some codecs, the time base is closer to the field rate than the frame rate.
	 * Most notably, H.264 and MPEG-2 specify time_base as half of frame duration
	 * if no telecine is used ...
	 *
	 * Set to time_base ticks per frame. Default 1, e.g., H.264/MPEG-2 set it to 2.
	 */
	TicksPerFrame ffcommon.FInt

	/**
	 * Codec delay.
	 *
	 * Encoding: Number of frames delay there will be from the encoder input to
	 *           the decoder output. (we assume the decoder matches the spec)
	 * Decoding: Number of frames delay in addition to what a standard decoder
	 *           as specified in the spec would produce.
	 *
	 * Video:
	 *   Number of frames the decoded output will be delayed relative to the
	 *   encoded input.
	 *
	 * Audio:
	 *   For encoding, this field is unused (see initial_padding).
	 *
	 *   For decoding, this is the number of samples the decoder needs to
	 *   output before the decoder's output is valid. When seeking, you should
	 *   start decoding this many samples prior to your desired seek point.
	 *
	 * - encoding: Set by libavcodec.
	 * - decoding: Set by libavcodec.
	 */
	Delay ffcommon.FInt

	/* video only */
	/**
	 * picture width / height.
	 *
	 * @note Those fields may not match the values of the last
	 * AVFrame output by avcodec_decode_video2 due frame
	 * reordering.
	 *
	 * - encoding: MUST be set by user.
	 * - decoding: May be set by the user before opening the decoder if known e.g.
	 *             from the container. Some decoders will require the dimensions
	 *             to be set by the caller. During decoding, the decoder may
	 *             overwrite those values as required while parsing the data.
	 */
	Width, Height ffcommon.FInt

	/**
	 * Bitstream width / height, may be different from width/height e.g. when
	 * the decoded frame is cropped before being output or lowres is enabled.
	 *
	 * @note Those field may not match the value of the last
	 * AVFrame output by avcodec_receive_frame() due frame
	 * reordering.
	 *
	 * - encoding: unused
	 * - decoding: May be set by the user before opening the decoder if known
	 *             e.g. from the container. During decoding, the decoder may
	 *             overwrite those values as required while parsing the data.
	 */
	CodedWidth, CodedHeight ffcommon.FInt

	/**
	 * the number of pictures in a group of pictures, or 0 for intra_only
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	GopSize ffcommon.FInt

	/**
	 * Pixel format, see AV_PIX_FMT_xxx.
	 * May be set by the demuxer if known from headers.
	 * May be overridden by the decoder if it knows better.
	 *
	 * @note This field may not match the value of the last
	 * AVFrame output by avcodec_receive_frame() due frame
	 * reordering.
	 *
	 * - encoding: Set by user.
	 * - decoding: Set by user if known, overridden by libavcodec while
	 *             parsing the data.
	 */
	PixFmt AVPixelFormat

	/**
	 * If non NULL, 'draw_horiz_band' is called by the libavcodec
	 * decoder to draw a horizontal band. It improves cache usage. Not
	 * all codecs can do that. You must check the codec capabilities
	 * beforehand.
	 * When multithreading is used, it may be called from multiple threads
	 * at the same time; threads might draw different parts of the same AVFrame,
	 * or multiple AVFrames, and there is no guarantee that slices will be drawn
	 * in order.
	 * The function is also used by hardware acceleration APIs.
	 * It is called at least once during frame decoding to pass
	 * the data needed for hardware render.
	 * In that mode instead of pixel data, AVFrame points to
	 * a structure specific to the acceleration API. The application
	 * reads the structure and can change some fields to indicate progress
	 * or mark state.
	 * - encoding: unused
	 * - decoding: Set by user.
	 * @param height the height of the slice
	 * @param y the y position of the slice
	 * @param type 1->top field, 2->bottom field, 3->frame
	 * @param offset offset into the AVFrame.data from which the slice should be read
	 */
	//void (*draw_horiz_band)(struct AVCodecContext *s,
	//const AVFrame *src, int offset[AV_NUM_DATA_POINTERS],
	//int y, int type, int height);
	DrawHorizBand uintptr
	/**
	 * callback to negotiate the pixelFormat
	 * @param fmt is the list of formats which are supported by the codec,
	 * it is terminated by -1 as 0 is a valid format, the formats are ordered by quality.
	 * The first is always the native one.
	 * @note The callback may be called again immediately if initialization for
	 * the selected (hardware-accelerated) pixel format failed.
	 * @warning Behavior is undefined if the callback returns a value not
	 * in the fmt list of formats.
	 * @return the chosen format
	 * - encoding: unused
	 * - decoding: Set by user, if not set the native format will be chosen.
	 */
	//enum AVPixelFormat (*get_format)(struct AVCodecContext *s, const enum AVPixelFormat * fmt);
	GetFormat uintptr
	/**
	 * maximum number of B-frames between non-B-frames
	 * Note: The output will be delayed by max_b_frames+1 relative to the input.
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	MaxBFrames ffcommon.FInt

	/**
	 * qscale factor between IP and B-frames
	 * If > 0 then the last P-frame quantizer will be used (q= lastp_q*factor+offset).
	 * If < 0 then normal ratecontrol will be done (q= -normal_q*factor+offset).
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	BQuantFactor ffcommon.FFloat

	//#if FF_API_PRIVATE_OPT
	/** @deprecated use encoder private options instead */
	//attribute_deprecated
	BFrameStrategy ffcommon.FInt
	//#endif

	/**
	 * qscale offset between IP and B-frames
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	BQuantOffset ffcommon.FFloat

	/**
	 * Size of the frame reordering buffer in the decoder.
	 * For MPEG-2 it is 1 IPB or 0 low delay IP.
	 * - encoding: Set by libavcodec.
	 * - decoding: Set by libavcodec.
	 */
	HasBFrames ffcommon.FInt

	//#if FF_API_PRIVATE_OPT
	/** @deprecated use encoder private options instead */
	//attribute_deprecated
	MpegQuant ffcommon.FInt
	//#endif

	/**
	 * qscale factor between P- and I-frames
	 * If > 0 then the last P-frame quantizer will be used (q = lastp_q * factor + offset).
	 * If < 0 then normal ratecontrol will be done (q= -normal_q*factor+offset).
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	IQuantFactor ffcommon.FFloat

	/**
	 * qscale offset between P and I-frames
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	IQuantOffset ffcommon.FFloat

	/**
	 * luminance masking (0-> disabled)
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	LumiMasking ffcommon.FFloat

	/**
	 * temporary complexity masking (0-> disabled)
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	TemporalCplxMasking ffcommon.FFloat

	/**
	 * spatial complexity masking (0-> disabled)
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	SpatialCplxMasking ffcommon.FFloat

	/**
	 * p block masking (0-> disabled)
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	PMasking ffcommon.FFloat

	/**
	 * darkness masking (0-> disabled)
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	DarkMasking ffcommon.FFloat

	/**
	 * slice count
	 * - encoding: Set by libavcodec.
	 * - decoding: Set by user (or 0).
	 */
	SliceCount ffcommon.FInt

	//#if FF_API_PRIVATE_OPT
	/** @deprecated use encoder private options instead */
	//attribute_deprecated
	PredictionMethod ffcommon.FInt
	//const FF_PRED_LEFT  = 0
	//const FF_PRED_PLANE = 1
	//const FF_PRED_MEDIAN =2
	//#endif

	/**
	 * slice offsets in the frame in bytes
	 * - encoding: Set/allocated by libavcodec.
	 * - decoding: Set/allocated by user (or NULL).
	 */
	SliceOffset *ffcommon.FInt

	/**
	 * sample aspect ratio (0 if unknown)
	 * That is the width of a pixel divided by the height of the pixel.
	 * Numerator and denominator must be relatively prime and smaller than 256 for some video standards.
	 * - encoding: Set by user.
	 * - decoding: Set by libavcodec.
	 */
	SampleAspectRatio AVRational

	/**
	 * motion estimation comparison function
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	MeCmp ffcommon.FInt
	/**
	 * subpixel motion estimation comparison function
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	MeSubCmp ffcommon.FInt
	/**
	 * macroblock comparison function (not supported yet)
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	MbCmp ffcommon.FInt
	/**
	 * interlaced DCT comparison function
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	IldctCmp ffcommon.FInt
	//const FF_CMP_SAD   =       0
	//const FF_CMP_SSE     =     1
	//const FF_CMP_SATD     =    2
	//const FF_CMP_DCT     =     3
	//const FF_CMP_PSNR   =      4
	//const FF_CMP_BIT    =      5
	//const FF_CMP_RD     =      6
	//const FF_CMP_ZERO    =     7
	//const FF_CMP_VSAD    =     8
	//const FF_CMP_VSSE    =     9
	//const FF_CMP_NSSE    =     10
	//const FF_CMP_W53     =     11
	//const FF_CMP_W97     =     12
	//const FF_CMP_DCTMAX     =  13
	//const FF_CMP_DCT264      = 14
	//const FF_CMP_MEDIAN_SAD =  15
	//const FF_CMP_CHROMA     =  256

	/**
	 * ME diamond size & shape
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	DiaSize ffcommon.FInt

	/**
	 * amount of previous MV predictors (2a+1 x 2a+1 square)
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	LastPredictorCount ffcommon.FInt

	//#if FF_API_PRIVATE_OPT
	///** @deprecated use encoder private options instead */
	//attribute_deprecated
	PreMe ffcommon.FInt
	//#endif

	/**
	 * motion estimation prepass comparison function
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	MePreCmp ffcommon.FInt

	/**
	 * ME prepass diamond size & shape
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	PreDiaSize ffcommon.FInt

	/**
	 * subpel ME quality
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	MeSubpelQuality ffcommon.FInt

	/**
	 * maximum motion estimation search range in subpel units
	 * If 0 then no limit.
	 *
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	MeRange ffcommon.FInt

	/**
	 * slice flags
	 * - encoding: unused
	 * - decoding: Set by user.
	 */
	SliceFlags ffcommon.FInt
	//const SLICE_FLAG_CODED_ORDER =   0x0001 ///< draw_horiz_band() is called in coded order instead of display
	//const SLICE_FLAG_ALLOW_FIELD   = 0x0002 ///< allow draw_horiz_band() with field slices (MPEG-2 field pics)
	//const SLICE_FLAG_ALLOW_PLANE   = 0x0004 ///< allow draw_horiz_band() with 1 component at a time (SVQ1)

	/**
	 * macroblock decision mode
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	MbDecision ffcommon.FInt
	//const FF_MB_DECISION_SIMPLE= 0        ///< uses mb_cmp
	//const FF_MB_DECISION_BITS =  1        ///< chooses the one which needs the fewest bits
	//const FF_MB_DECISION_RD   =  2        ///< rate distortion

	/**
	 * custom intra quantization matrix
	 * Must be allocated with the av_malloc() family of functions, and will be freed in
	 * avcodec_free_context().
	 * - encoding: Set/allocated by user, freed by libavcodec. Can be NULL.
	 * - decoding: Set/allocated/freed by libavcodec.
	 */
	IntraMatrix *ffcommon.FUint16T

	/**
	 * custom inter quantization matrix
	 * Must be allocated with the av_malloc() family of functions, and will be freed in
	 * avcodec_free_context().
	 * - encoding: Set/allocated by user, freed by libavcodec. Can be NULL.
	 * - decoding: Set/allocated/freed by libavcodec.
	 */
	InterMatrix *ffcommon.FUint16T

	//#if FF_API_PRIVATE_OPT
	/** @deprecated use encoder private options instead */
	//attribute_deprecated
	ScenechangeThreshold ffcommon.FInt

	/** @deprecated use encoder private options instead */
	//attribute_deprecated
	NoiseReduction ffcommon.FInt
	//#endif

	/**
	 * precision of the intra DC coefficient - 8
	 * - encoding: Set by user.
	 * - decoding: Set by libavcodec
	 */
	IntraDcPrecision ffcommon.FInt

	/**
	 * Number of macroblock rows at the top which are skipped.
	 * - encoding: unused
	 * - decoding: Set by user.
	 */
	SkipTop ffcommon.FInt

	/**
	 * Number of macroblock rows at the bottom which are skipped.
	 * - encoding: unused
	 * - decoding: Set by user.
	 */
	SkipBottom ffcommon.FInt

	/**
	 * minimum MB Lagrange multiplier
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	MbLmin ffcommon.FInt

	/**
	 * maximum MB Lagrange multiplier
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	MbLmax ffcommon.FInt

	//#if FF_API_PRIVATE_OPT
	/**
	 * @deprecated use encoder private options instead
	 */
	//attribute_deprecated
	MePenaltyCompensation ffcommon.FInt
	//#endif

	/**
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	BidirRefine ffcommon.FInt

	//#if FF_API_PRIVATE_OPT
	/** @deprecated use encoder private options instead */
	//attribute_deprecated
	BrdScale ffcommon.FInt
	//#endif

	/**
	 * minimum GOP size
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	KeyintMin ffcommon.FInt

	/**
	 * number of reference frames
	 * - encoding: Set by user.
	 * - decoding: Set by lavc.
	 */
	Refs ffcommon.FInt

	//#if FF_API_PRIVATE_OPT
	/** @deprecated use encoder private options instead */
	//attribute_deprecated
	Chromaoffset ffcommon.FInt
	//#endif

	/**
	 * Note: Value depends upon the compare function used for fullpel ME.
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	Mv0Threshold ffcommon.FInt

	//#if FF_API_PRIVATE_OPT
	/** @deprecated use encoder private options instead */
	//attribute_deprecated
	BSensitivity ffcommon.FInt
	//#endif

	/**
	 * Chromaticity coordinates of the source primaries.
	 * - encoding: Set by user
	 * - decoding: Set by libavcodec
	 */
	ColorPrimaries AVColorPrimaries

	/**
	 * Color Transfer Characteristic.
	 * - encoding: Set by user
	 * - decoding: Set by libavcodec
	 */
	ColorTrc AVColorTransferCharacteristic

	/**
	 * YUV colorspace type.
	 * - encoding: Set by user
	 * - decoding: Set by libavcodec
	 */
	Colorspace AVColorSpace

	/**
	 * MPEG vs JPEG YUV range.
	 * - encoding: Set by user
	 * - decoding: Set by libavcodec
	 */
	ColorRange AVColorRange

	/**
	 * This defines the location of chroma samples.
	 * - encoding: Set by user
	 * - decoding: Set by libavcodec
	 */
	ChromaSampleLocation AVChromaLocation

	/**
	 * Number of slices.
	 * Indicates number of picture subdivisions. Used for parallelized
	 * decoding.
	 * - encoding: Set by user
	 * - decoding: unused
	 */
	Slices ffcommon.FInt

	/** Field order
	 * - encoding: set by libavcodec
	 * - decoding: Set by user.
	 */
	FieldOrder AVFieldOrder

	/* audio only */
	SampleRate ffcommon.FInt ///< samples per second
	Channels   ffcommon.FInt ///< number of audio channels

	/**
	 * audio sample format
	 * - encoding: Set by user.
	 * - decoding: Set by libavcodec.
	 */
	SampleFmt AVSampleFormat ///< sample format

	/* The following data should not be initialized. */
	/**
	 * Number of samples per channel in an audio frame.
	 *
	 * - encoding: set by libavcodec in avcodec_open2(). Each submitted frame
	 *   except the last must contain exactly frame_size samples per channel.
	 *   May be 0 when the codec has AV_CODEC_CAP_VARIABLE_FRAME_SIZE set, then the
	 *   frame size is not restricted.
	 * - decoding: may be set by some decoders to indicate constant frame size
	 */
	FrameSize ffcommon.FInt

	/**
	 * Frame counter, set by libavcodec.
	 *
	 * - decoding: total number of frames returned from the decoder so far.
	 * - encoding: total number of frames passed to the encoder so far.
	 *
	 *   @note the counter is not incremented if encoding/decoding resulted in
	 *   an error.
	 */
	FrameNumber ffcommon.FInt

	/**
	 * number of bytes per packet if constant and known or 0
	 * Used by some WAV based audio codecs.
	 */
	BlockAlign ffcommon.FInt

	/**
	 * Audio cutoff bandwidth (0 means "automatic")
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	Cutoff ffcommon.FInt

	/**
	 * Audio channel layout.
	 * - encoding: set by user.
	 * - decoding: set by user, may be overwritten by libavcodec.
	 */
	ChannelLayout ffcommon.FUint64T

	/**
	 * Request decoder to use this channel layout if it can (0 for default)
	 * - encoding: unused
	 * - decoding: Set by user.
	 */
	RequestChannelLayout ffcommon.FUint64T

	/**
	 * Type of service that the audio stream conveys.
	 * - encoding: Set by user.
	 * - decoding: Set by libavcodec.
	 */
	AudioServiceType AVAudioServiceType

	/**
	 * desired sample format
	 * - encoding: Not used.
	 * - decoding: Set by user.
	 * Decoder will decode to this format if it can.
	 */
	RequestSampleFmt AVSampleFormat

	/**
	 * This callback is called at the beginning of each frame to get data
	 * buffer(s) for it. There may be one contiguous buffer for all the data or
	 * there may be a buffer per each data plane or anything in between. What
	 * this means is, you may set however many entries in buf[] you feel necessary.
	 * Each buffer must be reference-counted using the AVBuffer API (see description
	 * of buf[] below).
	 *
	 * The following fields will be set in the frame before this callback is
	 * called:
	 * - format
	 * - width, height (video only)
	 * - sample_rate, channel_layout, nb_samples (audio only)
	 * Their values may differ from the corresponding values in
	 * AVCodecContext. This callback must use the frame values, not the codec
	 * context values, to calculate the required buffer size.
	 *
	 * This callback must fill the following fields in the frame:
	 * - data[]
	 * - linesize[]
	 * - extended_data:
	 *   * if the data is planar audio with more than 8 channels, then this
	 *     callback must allocate and fill extended_data to contain all pointers
	 *     to all data planes. data[] must hold as many pointers as it can.
	 *     extended_data must be allocated with av_malloc() and will be freed in
	 *     av_frame_unref().
	 *   * otherwise extended_data must point to data
	 * - buf[] must contain one or more pointers to AVBufferRef structures. Each of
	 *   the frame's data and extended_data pointers must be contained in these. That
	 *   is, one AVBufferRef for each allocated chunk of memory, not necessarily one
	 *   AVBufferRef per data[] entry. See: av_buffer_create(), av_buffer_alloc(),
	 *   and av_buffer_ref().
	 * - extended_buf and nb_extended_buf must be allocated with av_malloc() by
	 *   this callback and filled with the extra buffers if there are more
	 *   buffers than buf[] can hold. extended_buf will be freed in
	 *   av_frame_unref().
	 *
	 * If AV_CODEC_CAP_DR1 is not set then get_buffer2() must call
	 * avcodec_default_get_buffer2() instead of providing buffers allocated by
	 * some other means.
	 *
	 * Each data plane must be aligned to the maximum required by the target
	 * CPU.
	 *
	 * @see avcodec_default_get_buffer2()
	 *
	 * Video:
	 *
	 * If AV_GET_BUFFER_FLAG_REF is set in flags then the frame may be reused
	 * (read and/or written to if it is writable) later by libavcodec.
	 *
	 * avcodec_align_dimensions2() should be used to find the required width and
	 * height, as they normally need to be rounded up to the next multiple of 16.
	 *
	 * Some decoders do not support linesizes changing between frames.
	 *
	 * If frame multithreading is used, this callback may be called from a
	 * different thread, but not from more than one at once. Does not need to be
	 * reentrant.
	 *
	 * @see avcodec_align_dimensions2()
	 *
	 * Audio:
	 *
	 * Decoders request a buffer of a particular size by setting
	 * AVFrame.nb_samples prior to calling get_buffer2(). The decoder may,
	 * however, utilize only part of the buffer by setting AVFrame.nb_samples
	 * to a smaller value in the output frame.
	 *
	 * As a convenience, av_samples_get_buffer_size() and
	 * av_samples_fill_arrays() in libavutil may be used by custom get_buffer2()
	 * functions to find the required data size and to fill data pointers and
	 * linesize. In AVFrame.linesize, only linesize[0] may be set for audio
	 * since all planes must be the same size.
	 *
	 * @see av_samples_get_buffer_size(), av_samples_fill_arrays()
	 *
	 * - encoding: unused
	 * - decoding: Set by libavcodec, user can override.
	 */
	//int (*get_buffer2)(struct AVCodecContext *s, AVFrame *frame, int flags);
	GetBuffer2 uintptr
	//#if FF_API_OLD_ENCDEC
	/**
	 * If non-zero, the decoded audio and video frames returned from
	 * avcodec_decode_video2() and avcodec_decode_audio4() are reference-counted
	 * and are valid indefinitely. The caller must free them with
	 * av_frame_unref() when they are not needed anymore.
	 * Otherwise, the decoded frames must not be freed by the caller and are
	 * only valid until the next decode call.
	 *
	 * This is always automatically enabled if avcodec_receive_frame() is used.
	 *
	 * - encoding: unused
	 * - decoding: set by the caller before avcodec_open2().
	 */
	//attribute_deprecated
	RefcountedFrames ffcommon.FInt
	//#endif

	/* - encoding parameters */
	Qcompress ffcommon.FFloat ///< amount of qscale change between easy & hard scenes (0.0-1.0)
	Qblur     ffcommon.FFloat ///< amount of qscale smoothing over time (0.0-1.0)

	/**
	 * minimum quantizer
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	Qmin ffcommon.FInt

	/**
	 * maximum quantizer
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	Qmax ffcommon.FInt

	/**
	 * maximum quantizer difference between frames
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	MaxQdiff ffcommon.FInt

	/**
	 * decoder bitstream buffer size
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	RcBufferSize ffcommon.FInt

	/**
	 * ratecontrol override, see RcOverride
	 * - encoding: Allocated/set/freed by user.
	 * - decoding: unused
	 */
	RcOverrideCount ffcommon.FInt
	COverride       *RcOverride

	/**
	 * maximum bitrate
	 * - encoding: Set by user.
	 * - decoding: Set by user, may be overwritten by libavcodec.
	 */
	RcMaxRate ffcommon.FInt64T

	/**
	 * minimum bitrate
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	RcMinRate ffcommon.FInt64T

	/**
	 * Ratecontrol attempt to use, at maximum, <value> of what can be used without an underflow.
	 * - encoding: Set by user.
	 * - decoding: unused.
	 */
	RcMaxAvailableVbvUse ffcommon.FFloat

	/**
	 * Ratecontrol attempt to use, at least, <value> times the amount needed to prevent a vbv overflow.
	 * - encoding: Set by user.
	 * - decoding: unused.
	 */
	RcMinVbvOverflowUse ffcommon.FFloat

	/**
	 * Number of bits which should be loaded into the rc buffer before decoding starts.
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	RcInitialBufferOccupancy ffcommon.FInt

	//#if FF_API_CODER_TYPE
	//const FF_CODER_TYPE_VLC     =  0
	//const FF_CODER_TYPE_AC    =    1
	//const FF_CODER_TYPE_RAW    =   2
	//const FF_CODER_TYPE_RLE    =   3
	/**
	 * @deprecated use encoder private options instead
	 */
	//attribute_deprecated
	CoderType ffcommon.FInt
	//#endif /* FF_API_CODER_TYPE */

	//#if FF_API_PRIVATE_OPT
	/** @deprecated use encoder private options instead */
	//attribute_deprecated
	ContextModel ffcommon.FInt
	//#endif

	//#if FF_API_PRIVATE_OPT
	/** @deprecated use encoder private options instead */
	//attribute_deprecated
	FrameSkipThreshold ffcommon.FInt

	/** @deprecated use encoder private options instead */
	//attribute_deprecated
	FrameSkipFactor ffcommon.FInt

	/** @deprecated use encoder private options instead */
	//attribute_deprecated
	FrameSkipExp ffcommon.FInt

	/** @deprecated use encoder private options instead */
	//attribute_deprecated
	FrameSkipCmp ffcommon.FInt
	//#endif /* FF_API_PRIVATE_OPT */

	/**
	 * trellis RD quantization
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	Trellis ffcommon.FInt

	//#if FF_API_PRIVATE_OPT
	/** @deprecated use encoder private options instead */
	//attribute_deprecated
	MinPredictionOrder ffcommon.FInt

	/** @deprecated use encoder private options instead */
	//attribute_deprecated
	MaxPredictionOrder ffcommon.FInt

	/** @deprecated use encoder private options instead */
	//attribute_deprecated
	TimecodeFrameStart ffcommon.FInt64T
	//#endif

	//#if FF_API_RTP_CALLBACK
	/**
	 * @deprecated unused
	 */
	/* The RTP callback: This function is called    */
	/* every time the encoder has a packet to send. */
	/* It depends on the encoder if the data starts */
	/* with a Start Code (it should). H.263 does.   */
	/* mb_nb contains the number of macroblocks     */
	/* encoded in the RTP payload.                  */
	//attribute_deprecated
	//void (*rtp_callback)(struct AVCodecContext *avctx, void *data, int size, int mb_nb);
	RtpCallback uintptr
	//#endif

	//#if FF_API_PRIVATE_OPT
	///** @deprecated use encoder private options instead */
	//attribute_deprecated
	RtpPayload_Size ffcommon.FInt /* The size of the RTP payload: the coder will  */
	/* do its best to deliver a chunk with size     */
	/* below rtp_payload_size, the chunk will start */
	/* with a start code on some codecs like H.263. */
	/* This doesn't take account of any particular  */
	/* headers inside the transmitted RTP payload.  */
	//#endif

	//#if FF_API_STAT_BITS
	///* statistics, used for 2-pass encoding */
	//attribute_deprecated
	MvBits ffcommon.FInt
	//attribute_deprecated
	HeaderBits ffcommon.FInt
	//attribute_deprecated
	ITexBits ffcommon.FInt
	//attribute_deprecated
	PTexBits ffcommon.FInt
	//attribute_deprecated
	ICount ffcommon.FInt
	//attribute_deprecated
	PCount ffcommon.FInt
	//attribute_deprecated
	SkipCount ffcommon.FInt
	//attribute_deprecated
	MiscBits ffcommon.FInt

	/** @deprecated this field is unused */
	//attribute_deprecated
	FrameBits ffcommon.FInt
	//#endif

	/**
	 * pass1 encoding statistics output buffer
	 * - encoding: Set by libavcodec.
	 * - decoding: unused
	 */
	StatsOut ffcommon.FCharPStruct

	/**
	 * pass2 encoding statistics input buffer
	 * Concatenated stuff from stats_out of pass1 should be placed here.
	 * - encoding: Allocated/set/freed by user.
	 * - decoding: unused
	 */
	StatsIn ffcommon.FCharPStruct

	/**
	 * Work around bugs in encoders which sometimes cannot be detected automatically.
	 * - encoding: Set by user
	 * - decoding: Set by user
	 */
	WorkaroundBugs ffcommon.FInt
	//const FF_BUG_AUTODETECT   =    1  ///< autodetection
	//const FF_BUG_XVID_ILACE   =    4
	//const FF_BUG_UMP4         =    8
	//const FF_BUG_NO_PADDING    =   16
	//const FF_BUG_AMV            =  32
	//const FF_BUG_QPEL_CHROMA     = 64
	//const FF_BUG_STD_QPEL       =  128
	//const FF_BUG_QPEL_CHROMA2   =  256
	//const FF_BUG_DIRECT_BLOCKSIZE =512
	//const FF_BUG_EDGE          =   1024
	//const FF_BUG_HPEL_CHROMA    =  2048
	//const FF_BUG_DC_CLIP       =   4096
	//const FF_BUG_MS             =  8192 ///< Work around various bugs in Microsoft's broken decoders.
	//const FF_BUG_TRUNCATED     =  16384
	//const FF_BUG_IEDGE        =   32768

	/**
	 * strictly follow the standard (MPEG-4, ...).
	 * - encoding: Set by user.
	 * - decoding: Set by user.
	 * Setting this to STRICT or higher means the encoder and decoder will
	 * generally do stupid things, whereas setting it to unofficial or lower
	 * will mean the encoder might produce output that is not supported by all
	 * spec-compliant decoders. Decoders don't differentiate between normal,
	 * unofficial and experimental (that is, they always try to decode things
	 * when they can) unless they are explicitly asked to behave stupidly
	 * (=strictly conform to the specs)
	 */
	StrictStdCompliance ffcommon.FInt
	//const FF_COMPLIANCE_VERY_STRICT =  2 ///< Strictly conform to an older more strict version of the spec or reference software.
	//const FF_COMPLIANCE_STRICT    =    1 ///< Strictly conform to all the things in the spec no matter what consequences.
	//const FF_COMPLIANCE_NORMAL     =   0
	//const FF_COMPLIANCE_UNOFFICIAL =  -1 ///< Allow unofficial extensions
	//const FF_COMPLIANCE_EXPERIMENTAL= -2 ///< Allow nonstandardized experimental things.

	/**
	 * error concealment flags
	 * - encoding: unused
	 * - decoding: Set by user.
	 */
	ErrorConcealment ffcommon.FInt
	//const FF_EC_GUESS_MVS  = 1
	//const FF_EC_DEBLOCK    = 2
	//const FF_EC_FAVOR_INTER =256

	/**
	 * debug
	 * - encoding: Set by user.
	 * - decoding: Set by user.
	 */
	Debug ffcommon.FInt
	//const FF_DEBUG_PICT_INFO =  1
	//const FF_DEBUG_RC        =  2
	//const FF_DEBUG_BITSTREAM  = 4
	//const FF_DEBUG_MB_TYPE   =  8
	//const FF_DEBUG_QP        =  16
	//const FF_DEBUG_DCT_COEFF =  0x00000040
	//const FF_DEBUG_SKIP       = 0x00000080
	//const FF_DEBUG_STARTCODE  = 0x00000100
	//const FF_DEBUG_ER        =  0x00000400
	//const FF_DEBUG_MMCO      =  0x00000800
	//const FF_DEBUG_BUGS      =  0x00001000
	//const FF_DEBUG_BUFFERS  =   0x00008000
	//const FF_DEBUG_THREADS   =  0x00010000
	//const FF_DEBUG_GREEN_MD   = 0x00800000
	//const FF_DEBUG_NOMC      =  0x01000000

	/**
	 * Error recognition; may misdetect some more or less valid parts as errors.
	 * - encoding: Set by user.
	 * - decoding: Set by user.
	 */
	ErrRecognition ffcommon.FInt

	/**
	 * Verify checksums embedded in the bitstream (could be of either encoded or
	 * decoded data, depending on the codec) and print an error message on mismatch.
	 * If AV_EF_EXPLODE is also set, a mismatching checksum will result in the
	 * decoder returning an error.
	 */
	//const AV_EF_CRCCHECK  =(1<<0)
	//const AV_EF_BITSTREAM =(1<<1)          ///< detect bitstream specification deviations
	//const AV_EF_BUFFER  =  (1<<2)          ///< detect improper bitstream length
	//const AV_EF_EXPLODE  = (1<<3)          ///< abort decoding on minor error detection
	//
	//const AV_EF_IGNORE_ERR =(1<<15)        ///< ignore errors and continue
	//const AV_EF_CAREFUL   = (1<<16)        ///< consider things that violate the spec, are fast to calculate and have not been seen in the wild as errors
	//const AV_EF_COMPLIANT = (1<<17)        ///< consider all spec non compliances as errors
	//const AV_EF_AGGRESSIVE =(1<<18)        ///< consider things that a sane encoder should not do as an error

	/**
	 * opaque 64-bit number (generally a PTS) that will be reordered and
	 * output in AVFrame.reordered_opaque
	 * - encoding: Set by libavcodec to the reordered_opaque of the input
	 *             frame corresponding to the last returned packet. Only
	 *             supported by encoders with the
	 *             AV_CODEC_CAP_ENCODER_REORDERED_OPAQUE capability.
	 * - decoding: Set by user.
	 */
	ReorderedOpaque ffcommon.FInt64T

	/**
	 * Hardware accelerator in use
	 * - encoding: unused.
	 * - decoding: Set by libavcodec
	 */
	Hwaccel *AVHWAccel

	/**
	 * Hardware accelerator context.
	 * For some hardware accelerators, a global context needs to be
	 * provided by the user. In that case, this holds display-dependent
	 * data FFmpeg cannot instantiate itself. Please refer to the
	 * FFmpeg HW accelerator documentation to know how to fill this
	 * is. e.g. for VA API, this is a struct vaapi_context.
	 * - encoding: unused
	 * - decoding: Set by user
	 */
	HwaccelContext ffcommon.FVoidP

	/**
	 * error
	 * - encoding: Set by libavcodec if flags & AV_CODEC_FLAG_PSNR.
	 * - decoding: unused
	 */
	Error [libavutil.AV_NUM_DATA_POINTERS]ffcommon.FUint64T

	/**
	 * DCT algorithm, see FF_DCT_* below
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	DctAlgo ffcommon.FInt
	//const FF_DCT_AUTO   = 0
	//const FF_DCT_FASTINT= 1
	//const FF_DCT_INT    = 2
	//const FF_DCT_MMX    = 3
	//const FF_DCT_ALTIVEC= 5
	//const FF_DCT_FAAN    =6

	/**
	 * IDCT algorithm, see FF_IDCT_* below.
	 * - encoding: Set by user.
	 * - decoding: Set by user.
	 */
	IdctAlgo ffcommon.FInt
	//const FF_IDCT_AUTO       =   0
	//const FF_IDCT_INT        =   1
	//const FF_IDCT_SIMPLE      =  2
	//const FF_IDCT_SIMPLEMMX   =  3
	//const FF_IDCT_ARM        =   7
	//const FF_IDCT_ALTIVEC     =  8
	//const FF_IDCT_SIMPLEARM    = 10
	//const FF_IDCT_XVID        =  14
	//const FF_IDCT_SIMPLEARMV5TE =16
	//const FF_IDCT_SIMPLEARMV6 =  17
	//const FF_IDCT_FAAN        =  20
	//const FF_IDCT_SIMPLENEON   = 22
	//const FF_IDCT_NONE         = 24 /* Used by XvMC to extract IDCT coefficients with FF_IDCT_PERM_NONE */
	//const FF_IDCT_SIMPLEAUTO  =  128

	/**
	 * bits per sample/pixel from the demuxer (needed for huffyuv).
	 * - encoding: Set by libavcodec.
	 * - decoding: Set by user.
	 */
	BitsPerCodedSample ffcommon.FInt

	/**
	 * Bits per sample/pixel of internal libavcodec pixel/sample format.
	 * - encoding: set by user.
	 * - decoding: set by libavcodec.
	 */
	BitsPerRawSample ffcommon.FInt

	/**
	 * low resolution decoding, 1-> 1/2 size, 2->1/4 size
	 * - encoding: unused
	 * - decoding: Set by user.
	 */
	Lowres ffcommon.FInt

	//#if FF_API_CODED_FRAME
	/**
	 * the picture in the bitstream
	 * - encoding: Set by libavcodec.
	 * - decoding: unused
	 *
	 * @deprecated use the quality factor packet side data instead
	 */
	//attribute_deprecated AVFrame *
	CodedFrame *AVFrame
	//#endif

	/**
	 * thread count
	 * is used to decide how many independent tasks should be passed to execute()
	 * - encoding: Set by user.
	 * - decoding: Set by user.
	 */
	ThreadCount ffcommon.FInt

	/**
	 * Which multithreading methods to use.
	 * Use of FF_THREAD_FRAME will increase decoding delay by one frame per thread,
	 * so clients which cannot provide future frames should not use it.
	 *
	 * - encoding: Set by user, otherwise the default is used.
	 * - decoding: Set by user, otherwise the default is used.
	 */
	ThreadType ffcommon.FInt
	//const FF_THREAD_FRAME   =1 ///< Decode more than one frame at once
	//const FF_THREAD_SLICE  = 2 ///< Decode more than one part of a single frame at once

	/**
	 * Which multithreading methods are in use by the codec.
	 * - encoding: Set by libavcodec.
	 * - decoding: Set by libavcodec.
	 */
	ActiveThreadType ffcommon.FInt

	//#if FF_API_THREAD_SAFE_CALLBACKS
	/**
	 * Set by the client if its custom get_buffer() callback can be called
	 * synchronously from another thread, which allows faster multithreaded decoding.
	 * draw_horiz_band() will be called from other threads regardless of this setting.
	 * Ignored if the default get_buffer() is used.
	 * - encoding: Set by user.
	 * - decoding: Set by user.
	 *
	 * @deprecated the custom get_buffer2() callback should always be
	 *   thread-safe. Thread-unsafe get_buffer2() implementations will be
	 *   invalid starting with LIBAVCODEC_VERSION_MAJOR=60; in other words,
	 *   libavcodec will behave as if this field was always set to 1.
	 *   Callers that want to be forward compatible with future libavcodec
	 *   versions should wrap access to this field in
	 *     #if LIBAVCODEC_VERSION_MAJOR < 60
	 */
	//attribute_deprecated
	ThreadSafeCallbacks ffcommon.FInt
	//#endif

	/**
	 * The codec may call this to execute several independent things.
	 * It will return only after finishing all tasks.
	 * The user may replace this with some multithreaded implementation,
	 * the default implementation will execute the parts serially.
	 * @param count the number of things to execute
	 * - encoding: Set by libavcodec, user can override.
	 * - decoding: Set by libavcodec, user can override.
	 */
	//int (*execute)(struct AVCodecContext *c, int (*func)(struct AVCodecContext *c2, void *arg), void *arg2, int *ret, int count, int size);
	Execute uintptr
	/**
	 * The codec may call this to execute several independent things.
	 * It will return only after finishing all tasks.
	 * The user may replace this with some multithreaded implementation,
	 * the default implementation will execute the parts serially.
	 * Also see avcodec_thread_init and e.g. the --enable-pthread configure option.
	 * @param c context passed also to func
	 * @param count the number of things to execute
	 * @param arg2 argument passed unchanged to func
	 * @param ret return values of executed functions, must have space for "count" values. May be NULL.
	 * @param func function that will be called count times, with jobnr from 0 to count-1.
	 *             threadnr will be in the range 0 to c->thread_count-1 < MAX_THREADS and so that no
	 *             two instances of func executing at the same time will have the same threadnr.
	 * @return always 0 currently, but code should handle a future improvement where when any call to func
	 *         returns < 0 no further calls to func may be done and < 0 is returned.
	 * - encoding: Set by libavcodec, user can override.
	 * - decoding: Set by libavcodec, user can override.
	 */
	//int (*execute2)(struct AVCodecContext *c, int (*func)(struct AVCodecContext *c2, void *arg, int jobnr, int threadnr), void *arg2, int *ret, int count);
	Execute2 uintptr
	/**
	 * noise vs. sse weight for the nsse comparison function
	 * - encoding: Set by user.
	 * - decoding: unused
	 */
	NsseWeight ffcommon.FInt

	/**
	 * profile
	 * - encoding: Set by user.
	 * - decoding: Set by libavcodec.
	 */
	Profile ffcommon.FInt
	//const FF_PROFILE_UNKNOWN =-99
	//const FF_PROFILE_RESERVED =-100
	//
	//const FF_PROFILE_AAC_MAIN =0
	//const FF_PROFILE_AAC_LOW = 1
	//const FF_PROFILE_AAC_SSR  =2
	//const FF_PROFILE_AAC_LTP = 3
	//const FF_PROFILE_AAC_HE =  4
	//const FF_PROFILE_AAC_HE_V2 =28
	//const FF_PROFILE_AAC_LD =  22
	//const FF_PROFILE_AAC_ELD = 38
	//const FF_PROFILE_MPEG2_AAC_LOW= 128
	//const FF_PROFILE_MPEG2_AAC_HE  =131

	//const FF_PROFILE_DNXHD      =   0
	//const FF_PROFILE_DNXHR_LB   =   1
	//const FF_PROFILE_DNXHR_SQ   =   2
	//const FF_PROFILE_DNXHR_HQ    =  3
	//const FF_PROFILE_DNXHR_HQX  =   4
	//const FF_PROFILE_DNXHR_444   =  5

	//const FF_PROFILE_DTS      =   20
	//const FF_PROFILE_DTS_ES    =  30
	//const FF_PROFILE_DTS_96_24  = 40
	//const FF_PROFILE_DTS_HD_HRA  =50
	//const FF_PROFILE_DTS_HD_MA  = 60
	//const FF_PROFILE_DTS_EXPRESS =70

	//const FF_PROFILE_MPEG2_422 =   0
	//const FF_PROFILE_MPEG2_HIGH =  1
	//const FF_PROFILE_MPEG2_SS  =   2
	//const FF_PROFILE_MPEG2_SNR_SCALABLE = 3
	//const FF_PROFILE_MPEG2_MAIN  = 4
	//const FF_PROFILE_MPEG2_SIMPLE =5

	//const FF_PROFILE_H264_CONSTRAINED = (1<<9)  // 8+1; constraint_set1_flag
	//const FF_PROFILE_H264_INTRA       = (1<<11) // 8+3; constraint_set3_flag
	//
	//const FF_PROFILE_H264_BASELINE        =     66
	//const FF_PROFILE_H264_CONSTRAINED_BASELINE =(66|FF_PROFILE_H264_CONSTRAINED)
	//const FF_PROFILE_H264_MAIN             =    77
	//const FF_PROFILE_H264_EXTENDED         =    88
	//const FF_PROFILE_H264_HIGH             =    100
	//const FF_PROFILE_H264_HIGH_10           =  110
	//const FF_PROFILE_H264_HIGH_10_INTRA     =   (110|FF_PROFILE_H264_INTRA)
	//const FF_PROFILE_H264_MULTIVIEW_HIGH    =   118
	//const FF_PROFILE_H264_HIGH_422         =    122
	//const FF_PROFILE_H264_HIGH_422_INTRA  =     (122|FF_PROFILE_H264_INTRA)
	//const FF_PROFILE_H264_STEREO_HIGH      =    128
	//const FF_PROFILE_H264_HIGH_444           =  144
	//const FF_PROFILE_H264_HIGH_444_PREDICTIVE = 244
	//const FF_PROFILE_H264_HIGH_444_INTRA    =   (244|FF_PROFILE_H264_INTRA)
	//const FF_PROFILE_H264_CAVLC_444          =  44
	//
	//const FF_PROFILE_VC1_SIMPLE =  0
	//const FF_PROFILE_VC1_MAIN    = 1
	//const FF_PROFILE_VC1_COMPLEX = 2
	//const FF_PROFILE_VC1_ADVANCED =3

	//const FF_PROFILE_MPEG4_SIMPLE              =       0
	//const FF_PROFILE_MPEG4_SIMPLE_SCALABLE      =      1
	//const FF_PROFILE_MPEG4_CORE                 =      2
	//const FF_PROFILE_MPEG4_MAIN                =       3
	//const FF_PROFILE_MPEG4_N_BIT                =      4
	//const FF_PROFILE_MPEG4_SCALABLE_TEXTURE     =      5
	//const FF_PROFILE_MPEG4_SIMPLE_FACE_ANIMATION  =    6
	//const FF_PROFILE_MPEG4_BASIC_ANIMATED_TEXTURE  =   7
	//const FF_PROFILE_MPEG4_HYBRID                =     8
	//const FF_PROFILE_MPEG4_ADVANCED_REAL_TIME   =      9
	//const FF_PROFILE_MPEG4_CORE_SCALABLE        =     10
	//const FF_PROFILE_MPEG4_ADVANCED_CODING     =      11
	//const FF_PROFILE_MPEG4_ADVANCED_CORE       =      12
	//const FF_PROFILE_MPEG4_ADVANCED_SCALABLE_TEXTURE =13
	//const FF_PROFILE_MPEG4_SIMPLE_STUDIO      =       14
	//const FF_PROFILE_MPEG4_ADVANCED_SIMPLE      =     15

	//const FF_PROFILE_JPEG2000_CSTREAM_RESTRICTION_0 =  1
	//const FF_PROFILE_JPEG2000_CSTREAM_RESTRICTION_1  = 2
	//const FF_PROFILE_JPEG2000_CSTREAM_NO_RESTRICTION = 32768
	//const FF_PROFILE_JPEG2000_DCINEMA_2K        =      3
	//const FF_PROFILE_JPEG2000_DCINEMA_4K        =      4
	//
	//const FF_PROFILE_VP9_0                  =          0
	//const FF_PROFILE_VP9_1                   =         1
	//const FF_PROFILE_VP9_2                   =         2
	//const FF_PROFILE_VP9_3                   =         3

	//const FF_PROFILE_HEVC_MAIN                 =       1
	//const FF_PROFILE_HEVC_MAIN_10               =      2
	//const FF_PROFILE_HEVC_MAIN_STILL_PICTURE    =      3
	//const FF_PROFILE_HEVC_REXT                  =      4
	//
	//const FF_PROFILE_VVC_MAIN_10                =      1
	//const FF_PROFILE_VVC_MAIN_10_444             =    33
	//
	//const FF_PROFILE_AV1_MAIN                 =        0
	//const FF_PROFILE_AV1_HIGH                  =       1
	//const FF_PROFILE_AV1_PROFESSIONAL          =       2

	//const FF_PROFILE_MJPEG_HUFFMAN_BASELINE_DCT         =   0xc0
	//const FF_PROFILE_MJPEG_HUFFMAN_EXTENDED_SEQUENTIAL_DCT =0xc1
	//const FF_PROFILE_MJPEG_HUFFMAN_PROGRESSIVE_DCT      =   0xc2
	//const FF_PROFILE_MJPEG_HUFFMAN_LOSSLESS            =    0xc3
	//const FF_PROFILE_MJPEG_JPEG_LS                     =    0xf7
	//
	//const FF_PROFILE_SBC_MSBC            =             1
	//
	//const FF_PROFILE_PRORES_PROXY   =  0
	//const FF_PROFILE_PRORES_LT      =  1
	//const FF_PROFILE_PRORES_STANDARD = 2
	//const FF_PROFILE_PRORES_HQ      =  3
	//const FF_PROFILE_PRORES_4444    =  4
	//const FF_PROFILE_PRORES_XQ     =   5

	//const FF_PROFILE_ARIB_PROFILE_A= 0
	//const FF_PROFILE_ARIB_PROFILE_C= 1
	//
	//const FF_PROFILE_KLVA_SYNC =0
	//const FF_PROFILE_KLVA_ASYNC =1

	/**
	 * level
	 * - encoding: Set by user.
	 * - decoding: Set by libavcodec.
	 */
	Level ffcommon.FInt
	//const FF_LEVEL_UNKNOWN =-99

	/**
	 * Skip loop filtering for selected frames.
	 * - encoding: unused
	 * - decoding: Set by user.
	 */
	SkipLoopFilter AVDiscard

	/**
	 * Skip IDCT/dequantization for selected frames.
	 * - encoding: unused
	 * - decoding: Set by user.
	 */
	SkipIdct AVDiscard

	/**
	 * Skip decoding for selected frames.
	 * - encoding: unused
	 * - decoding: Set by user.
	 */
	SkipFrame AVDiscard

	/**
	 * Header containing style information for text subtitles.
	 * For SUBTITLE_ASS subtitle type, it should contain the whole ASS
	 * [Script Info] and [V4+ Styles] section, plus the [Events] line and
	 * the Format line following. It shouldn't include any Dialogue line.
	 * - encoding: Set/allocated/freed by user (before avcodec_open2())
	 * - decoding: Set/allocated/freed by libavcodec (by avcodec_open2())
	 */
	SubtitleHeader      *ffcommon.FUint8T
	SubtitleHeaderSizev ffcommon.FInt

	//#if FF_API_VBV_DELAY
	/**
	 * VBV delay coded in the last frame (in periods of a 27 MHz clock).
	 * Used for compliant TS muxing.
	 * - encoding: Set by libavcodec.
	 * - decoding: unused.
	 * @deprecated this value is now exported as a part of
	 * AV_PKT_DATA_CPB_PROPERTIES packet side data
	 */
	//attribute_deprecated
	VbvDelay ffcommon.FUint64T
	//#endif

	//#if FF_API_SIDEDATA_ONLY_PKT
	/**
	 * Encoding only and set by default. Allow encoders to output packets
	 * that do not contain any encoded data, only side data.
	 *
	 * Some encoders need to output such packets, e.g. to update some stream
	 * parameters at the end of encoding.
	 *
	 * @deprecated this field disables the default behaviour and
	 *             it is kept only for compatibility.
	 */
	//attribute_deprecated
	SideDataOnlyPackets ffcommon.FInt
	//#endif

	/**
	 * Audio only. The number of "priming" samples (padding) inserted by the
	 * encoder at the beginning of the audio. I.e. this number of leading
	 * decoded samples must be discarded by the caller to get the original audio
	 * without leading padding.
	 *
	 * - decoding: unused
	 * - encoding: Set by libavcodec. The timestamps on the output packets are
	 *             adjusted by the encoder so that they always refer to the
	 *             first sample of the data actually contained in the packet,
	 *             including any added padding.  E.g. if the timebase is
	 *             1/samplerate and the timestamp of the first input sample is
	 *             0, the timestamp of the first output packet will be
	 *             -initial_padding.
	 */
	InitialPadding ffcommon.FInt

	/**
	 * - decoding: For codecs that store a framerate value in the compressed
	 *             bitstream, the decoder may export it here. { 0, 1} when
	 *             unknown.
	 * - encoding: May be used to signal the framerate of CFR content to an
	 *             encoder.
	 */
	Framerate AVRational

	/**
	 * Nominal unaccelerated pixel format, see AV_PIX_FMT_xxx.
	 * - encoding: unused.
	 * - decoding: Set by libavcodec before calling get_format()
	 */
	SwPixFmt AVPixelFormat

	/**
	 * Timebase in which pkt_dts/pts and AVPacket.dts/pts are.
	 * - encoding unused.
	 * - decoding set by user.
	 */
	PktTimebase AVRational

	/**
	 * AVCodecDescriptor
	 * - encoding: unused.
	 * - decoding: set by libavcodec.
	 */
	CodecDescriptor *AVCodecDescriptor

	/**
	 * Current statistics for PTS correction.
	 * - decoding: maintained and used by libavcodec, not intended to be used by user apps
	 * - encoding: unused
	 */
	PtsCorrectionNumFaultyPts ffcommon.FInt64T /// Number of incorrect PTS values so far
	PtsCorrectionNumFaultyDts ffcommon.FInt64T /// Number of incorrect DTS values so far
	PtsCorrectionLastPts      ffcommon.FInt64T /// PTS of the last frame
	PtsCorrectionLastDts      ffcommon.FInt64T /// DTS of the last frame

	/**
	 * Character encoding of the input subtitles file.
	 * - decoding: set by user
	 * - encoding: unused
	 */
	SubCharenc ffcommon.FCharPStruct

	/**
	 * Subtitles character encoding mode. Formats or codecs might be adjusting
	 * this setting (if they are doing the conversion themselves for instance).
	 * - decoding: set by libavcodec
	 * - encoding: unused
	 */
	SubCharencMode ffcommon.FInt
	//const FF_SUB_CHARENC_MODE_DO_NOTHING=  -1  ///< do nothing (demuxer outputs a stream supposed to be already in UTF-8, or the codec is bitmap for instance)
	//const FF_SUB_CHARENC_MODE_AUTOMATIC   = 0  ///< libavcodec will select the mode itself
	//const FF_SUB_CHARENC_MODE_PRE_DECODER = 1  ///< the AVPacket data needs to be recoded to UTF-8 before being fed to the decoder, requires iconv
	//const FF_SUB_CHARENC_MODE_IGNORE     =  2  ///< neither convert the subtitles, nor check them for valid UTF-8

	/**
	 * Skip processing alpha if supported by codec.
	 * Note that if the format uses pre-multiplied alpha (common with VP6,
	 * and recommended due to better video quality/compression)
	 * the image will look as if alpha-blended onto a black background.
	 * However for formats that do not use pre-multiplied alpha
	 * there might be serious artefacts (though e.g. libswscale currently
	 * assumes pre-multiplied alpha anyway).
	 *
	 * - decoding: set by user
	 * - encoding: unused
	 */
	SkipAlpha ffcommon.FInt

	/**
	 * Number of samples to skip after a discontinuity
	 * - decoding: unused
	 * - encoding: set by libavcodec
	 */
	SeekPreroll ffcommon.FInt

	//#if FF_API_DEBUG_MV
	/**
	 * @deprecated unused
	 */
	//attribute_deprecated
	DebugMv ffcommon.FInt
	//const FF_DEBUG_VIS_MV_P_FOR = 0x00000001 //visualize forward predicted MVs of P frames
	//const FF_DEBUG_VIS_MV_B_FOR = 0x00000002 //visualize forward predicted MVs of B frames
	//const FF_DEBUG_VIS_MV_B_BACK= 0x00000004 //visualize backward predicted MVs of B frames
	//#endif

	/**
	 * custom intra quantization matrix
	 * - encoding: Set by user, can be NULL.
	 * - decoding: unused.
	 */
	ChromaIntraMatrix *ffcommon.FUint16T

	/**
	 * dump format separator.
	 * can be ", " or "\n      " or anything else
	 * - encoding: Set by user.
	 * - decoding: Set by user.
	 */
	DumpSeparator *ffcommon.FUint8T

	/**
	 * ',' separated list of allowed decoders.
	 * If NULL then all are allowed
	 * - encoding: unused
	 * - decoding: set by user
	 */
	CodecWhitelist ffcommon.FCharPStruct

	/**
	 * Properties of the stream that gets decoded
	 * - encoding: unused
	 * - decoding: set by libavcodec
	 */
	Properties ffcommon.FUnsigned
	//const FF_CODEC_PROPERTY_LOSSLESS    =    0x00000001
	//const FF_CODEC_PROPERTY_CLOSED_CAPTIONS= 0x00000002

	/**
	 * Additional data associated with the entire coded stream.
	 *
	 * - decoding: unused
	 * - encoding: may be set by libavcodec after avcodec_open2().
	 */
	CodedSideData   *AVPacketSideData
	NbCodedSideData ffcommon.FInt

	/**
	 * A reference to the AVHWFramesContext describing the input (for encoding)
	 * or output (decoding) frames. The reference is set by the caller and
	 * afterwards owned (and freed) by libavcodec - it should never be read by
	 * the caller after being set.
	 *
	 * - decoding: This field should be set by the caller from the get_format()
	 *             callback. The previous reference (if any) will always be
	 *             unreffed by libavcodec before the get_format() call.
	 *
	 *             If the default get_buffer2() is used with a hwaccel pixel
	 *             format, then this AVHWFramesContext will be used for
	 *             allocating the frame buffers.
	 *
	 * - encoding: For hardware encoders configured to use a hwaccel pixel
	 *             format, this field should be set by the caller to a reference
	 *             to the AVHWFramesContext describing input frames.
	 *             AVHWFramesContext.format must be equal to
	 *             AVCodecContext.pix_fmt.
	 *
	 *             This field should be set before avcodec_open2() is called.
	 */
	HwFramesCtx *AVBufferRef

	/**
	 * Control the form of AVSubtitle.rects[N]->ass
	 * - decoding: set by user
	 * - encoding: unused
	 */
	SubTextFormat ffcommon.FInt
	//const FF_SUB_TEXT_FMT_ASS    =          0
	//#if FF_API_ASS_TIMING
	//const FF_SUB_TEXT_FMT_ASS_WITH_TIMINGS= 1
	//#endif

	/**
	 * Audio only. The amount of padding (in samples) appended by the encoder to
	 * the end of the audio. I.e. this number of decoded samples must be
	 * discarded by the caller from the end of the stream to get the original
	 * audio without any trailing padding.
	 *
	 * - decoding: unused
	 * - encoding: unused
	 */
	TrailingPadding ffcommon.FInt

	/**
	 * The number of pixels per image to maximally accept.
	 *
	 * - decoding: set by user
	 * - encoding: set by user
	 */
	MaxPixels ffcommon.FInt64T

	/**
	 * A reference to the AVHWDeviceContext describing the device which will
	 * be used by a hardware encoder/decoder.  The reference is set by the
	 * caller and afterwards owned (and freed) by libavcodec.
	 *
	 * This should be used if either the codec device does not require
	 * hardware frames or any that are used are to be allocated internally by
	 * libavcodec.  If the user wishes to supply any of the frames used as
	 * encoder input or decoder output then hw_frames_ctx should be used
	 * instead.  When hw_frames_ctx is set in get_format() for a decoder, this
	 * field will be ignored while decoding the associated stream segment, but
	 * may again be used on a following one after another get_format() call.
	 *
	 * For both encoders and decoders this field should be set before
	 * avcodec_open2() is called and must not be written to thereafter.
	 *
	 * Note that some decoders may require this field to be set initially in
	 * order to support hw_frames_ctx at all - in that case, all frames
	 * contexts used must be created on the same device.
	 */
	HwDeviceCtx *AVBufferRef

	/**
	 * Bit set of AV_HWACCEL_FLAG_* flags, which affect hardware accelerated
	 * decoding (if active).
	 * - encoding: unused
	 * - decoding: Set by user (either before avcodec_open2(), or in the
	 *             AVCodecContext.get_format callback)
	 */
	HwaccelFlags ffcommon.FInt

	/**
	 * Video decoding only. Certain video codecs support cropping, meaning that
	 * only a sub-rectangle of the decoded frame is intended for display.  This
	 * option controls how cropping is handled by libavcodec.
	 *
	 * When set to 1 (the default), libavcodec will apply cropping internally.
	 * I.e. it will modify the output frame width/height fields and offset the
	 * data pointers (only by as much as possible while preserving alignment, or
	 * by the full amount if the AV_CODEC_FLAG_UNALIGNED flag is set) so that
	 * the frames output by the decoder refer only to the cropped area. The
	 * crop_* fields of the output frames will be zero.
	 *
	 * When set to 0, the width/height fields of the output frames will be set
	 * to the coded dimensions and the crop_* fields will describe the cropping
	 * rectangle. Applying the cropping is left to the caller.
	 *
	 * @warning When hardware acceleration with opaque output frames is used,
	 * libavcodec is unable to apply cropping from the top/left border.
	 *
	 * @note when this option is set to zero, the width/height fields of the
	 * AVCodecContext and output AVFrames have different meanings. The codec
	 * context fields store display dimensions (with the coded dimensions in
	 * coded_width/height), while the frame fields store the coded dimensions
	 * (with the display dimensions being determined by the crop_* fields).
	 */
	ApplyCropping ffcommon.FInt

	/*
	 * Video decoding only.  Sets the number of extra hardware frames which
	 * the decoder will allocate for use by the caller.  This must be set
	 * before avcodec_open2() is called.
	 *
	 * Some hardware decoders require all frames that they will use for
	 * output to be defined in advance before decoding starts.  For such
	 * decoders, the hardware frame pool must therefore be of a fixed size.
	 * The extra frames set here are on top of any number that the decoder
	 * needs internally in order to operate normally (for example, frames
	 * used as reference pictures).
	 */
	ExtraHwFrames ffcommon.FInt

	/**
	 * The percentage of damaged samples to discard a frame.
	 *
	 * - decoding: set by user
	 * - encoding: unused
	 */
	DiscardDamagedPercentage ffcommon.FInt

	/**
	 * The number of samples per frame to maximally accept.
	 *
	 * - decoding: set by user
	 * - encoding: set by user
	 */
	MaxSamples ffcommon.FInt64T

	/**
	 * Bit set of AV_CODEC_EXPORT_DATA_* flags, which affects the kind of
	 * metadata exported in frame, packet, or coded stream side data by
	 * decoders and encoders.
	 *
	 * - decoding: set by user
	 * - encoding: set by user
	 */
	ExportSideData ffcommon.FInt

	/**
	 * This callback is called at the beginning of each packet to get a data
	 * buffer for it.
	 *
	 * The following field will be set in the packet before this callback is
	 * called:
	 * - size
	 * This callback must use the above value to calculate the required buffer size,
	 * which must padded by at least AV_INPUT_BUFFER_PADDING_SIZE bytes.
	 *
	 * This callback must fill the following fields in the packet:
	 * - data: alignment requirements for AVPacket apply, if any. Some architectures and
	 *   encoders may benefit from having aligned data.
	 * - buf: must contain a pointer to an AVBufferRef structure. The packet's
	 *   data pointer must be contained in it. See: av_buffer_create(), av_buffer_alloc(),
	 *   and av_buffer_ref().
	 *
	 * If AV_CODEC_CAP_DR1 is not set then get_encode_buffer() must call
	 * avcodec_default_get_encode_buffer() instead of providing a buffer allocated by
	 * some other means.
	 *
	 * The flags field may contain a combination of AV_GET_ENCODE_BUFFER_FLAG_ flags.
	 * They may be used for example to hint what use the buffer may get after being
	 * created.
	 * Implementations of this callback may ignore flags they don't understand.
	 * If AV_GET_ENCODE_BUFFER_FLAG_REF is set in flags then the packet may be reused
	 * (read and/or written to if it is writable) later by libavcodec.
	 *
	 * This callback must be thread-safe, as when frame threading is used, it may
	 * be called from multiple threads simultaneously.
	 *
	 * @see avcodec_default_get_encode_buffer()
	 *
	 * - encoding: Set by libavcodec, user can override.
	 * - decoding: unused
	 */
	//int (*get_encode_buffer)(struct AVCodecContext *s, AVPacket *pkt, int flags);
	GetEncodeBuffer uintptr
}

//#if FF_API_CODEC_GET_SET
/**
* Accessors for some AVCodecContext fields. These used to be provided for ABI
* compatibility, and do not need to be used anymore.
 */
//attribute_deprecated
//AVRational av_codec_get_pkt_timebase         (const AVCodecContext *avctx);
func (avctx *AVCodecContext) AvCodecGetPktTimebase() (res AVRational) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_codec_get_pkt_timebase").Call(
		uintptr(unsafe.Pointer(avctx)),
	)
	res = *(*AVRational)(unsafe.Pointer(t))
	return
}

// attribute_deprecated
// void       av_codec_set_pkt_timebase         (AVCodecContext *avctx, AVRational val);
func (avctx *AVCodecContext) AvCodecSetPktTimebase(val AVRational) {
	ffcommon.GetAvcodecDll().NewProc("av_codec_set_pkt_timebase").Call(
		uintptr(unsafe.Pointer(avctx)),
		uintptr(unsafe.Pointer(&val)),
	)
}

// attribute_deprecated
// const AVCodecDescriptor *av_codec_get_codec_descriptor(const AVCodecContext *avctx);
func (avctx *AVCodecContext) AvCodecGetCodecDescriptor() (res *AVCodecDescriptor) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_codec_get_codec_descriptor").Call(
		uintptr(unsafe.Pointer(avctx)),
	)
	res = (*AVCodecDescriptor)(unsafe.Pointer(t))
	return
}

// attribute_deprecated
// void                     av_codec_set_codec_descriptor(AVCodecContext *avctx, const AVCodecDescriptor *desc);
func (avctx *AVCodecContext) AvCodecSetCodecDescriptor(val *AVCodecDescriptor) {
	ffcommon.GetAvcodecDll().NewProc("av_codec_set_codec_descriptor").Call(
		uintptr(unsafe.Pointer(avctx)),
		uintptr(unsafe.Pointer(val)),
	)
}

// attribute_deprecated
// unsigned av_codec_get_codec_properties(const AVCodecContext *avctx);
func (avctx *AVCodecContext) AvCodecGetCodecProperties() (res ffcommon.FUnsigned) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_codec_get_codec_properties").Call(
		uintptr(unsafe.Pointer(avctx)),
	)
	res = ffcommon.FUnsigned(t)
	return
}

// attribute_deprecated
// int  av_codec_get_lowres(const AVCodecContext *avctx);
func (avctx *AVCodecContext) AvCodecGetLowres() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_codec_get_lowres").Call(
		uintptr(unsafe.Pointer(avctx)),
	)
	res = ffcommon.FInt(t)
	return
}

// attribute_deprecated
// void av_codec_set_lowres(AVCodecContext *avctx, int val);
func (avctx *AVCodecContext) AvCodecSetLowres(val ffcommon.FInt) {
	ffcommon.GetAvcodecDll().NewProc("av_codec_set_lowres").Call(
		uintptr(unsafe.Pointer(avctx)),
		uintptr(val),
	)
}

// attribute_deprecated
// int  av_codec_get_seek_preroll(const AVCodecContext *avctx);
func (avctx *AVCodecContext) AvCodecGetSeekPreroll() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_codec_get_seek_preroll").Call(
		uintptr(unsafe.Pointer(avctx)),
	)
	res = ffcommon.FInt(t)
	return
}

// attribute_deprecated
// void av_codec_set_seek_preroll(AVCodecContext *avctx, int val);
func (avctx *AVCodecContext) AvCodecSetSeekPreroll(val ffcommon.FInt) {
	ffcommon.GetAvcodecDll().NewProc("av_codec_set_seek_preroll").Call(
		uintptr(unsafe.Pointer(avctx)),
		uintptr(val),
	)
}

// attribute_deprecated
// uint16_t *av_codec_get_chroma_intra_matrix(const AVCodecContext *avctx);
func (avctx *AVCodecContext) AvCodecGetChromaIntraMatrix() (res *ffcommon.FUint16T) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_codec_get_chroma_intra_matrix").Call(
		uintptr(unsafe.Pointer(avctx)),
	)
	res = (*ffcommon.FUint16T)(unsafe.Pointer(t))
	return
}

// attribute_deprecated
// void av_codec_set_chroma_intra_matrix(AVCodecContext *avctx, uint16_t *val);
func (avctx *AVCodecContext) AvCodecSetChromaIntraMatrix(val *ffcommon.FUint16T) {
	ffcommon.GetAvcodecDll().NewProc("av_codec_set_chroma_intra_matrix").Call(
		uintptr(unsafe.Pointer(avctx)),
		uintptr(unsafe.Pointer(val)),
	)
}

//#endif

//type AVSubtitle struct {
//
//}

// #if FF_API_CODEC_GET_SET
// attribute_deprecated
// int av_codec_get_max_lowres(const AVCodec *codec);
// #endif
func (codec *AVCodec) AvCodecGetMaxLowres() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_codec_get_max_lowres").Call(
		uintptr(unsafe.Pointer(codec)),
	)
	res = ffcommon.FInt(t)
	return
}

type MpegEncContext struct {
}

/**
* @defgroup lavc_hwaccel AVHWAccel
*
* @note  Nothing in this structure should be accessed by the user.  At some
*        point in future it will not be externally visible at all.
*
* @{
 */
type AVHWAccel struct {

	/**
	 * Name of the hardware accelerated codec.
	 * The name is globally unique among encoders and among decoders (but an
	 * encoder and a decoder can share the same name).
	 */
	Name ffcommon.FCharPStruct

	/**
	 * Type of codec implemented by the hardware accelerator.
	 *
	 * See AVMEDIA_TYPE_xxx
	 */
	Type AVMediaType

	/**
	 * Codec implemented by the hardware accelerator.
	 *
	 * See AV_CODEC_ID_xxx
	 */
	Id AVCodecID

	/**
	 * Supported pixel format.
	 *
	 * Only hardware accelerated formats are supported here.
	 */
	PixFmt AVPixelFormat

	/**
	 * Hardware accelerated codec capabilities.
	 * see AV_HWACCEL_CODEC_CAP_*
	 */
	Capabilities ffcommon.FInt

	/*****************************************************************
	 * No fields below this line are part of the public API. They
	 * may not be used outside of libavcodec and can be changed and
	 * removed at will.
	 * New public fields should be added right above.
	 *****************************************************************
	 */

	/**
	 * Allocate a custom buffer
	 */
	//int (*alloc_frame)(AVCodecContext *avctx, AVFrame *frame);
	AllocFrame uintptr
	/**
	 * Called at the beginning of each frame or field picture.
	 *
	 * Meaningful frame information (codec specific) is guaranteed to
	 * be parsed at this point. This function is mandatory.
	 *
	 * Note that buf can be NULL along with buf_size set to 0.
	 * Otherwise, this means the whole frame is available at this point.
	 *
	 * @param avctx the codec context
	 * @param buf the frame data buffer base
	 * @param buf_size the size of the frame in bytes
	 * @return zero if successful, a negative value otherwise
	 */
	//int (*start_frame)(AVCodecContext *avctx, const uint8_t *buf, uint32_t buf_size);
	StartFrame uintptr
	/**
	 * Callback for parameter data (SPS/PPS/VPS etc).
	 *
	 * Useful for hardware decoders which keep persistent state about the
	 * video parameters, and need to receive any changes to update that state.
	 *
	 * @param avctx the codec context
	 * @param type the nal unit type
	 * @param buf the nal unit data buffer
	 * @param buf_size the size of the nal unit in bytes
	 * @return zero if successful, a negative value otherwise
	 */
	//int (*decode_params)(AVCodecContext *avctx, int type, const uint8_t *buf, uint32_t buf_size);
	DecodeParams uintptr
	/**
	 * Callback for each slice.
	 *
	 * Meaningful slice information (codec specific) is guaranteed to
	 * be parsed at this point. This function is mandatory.
	 * The only exception is XvMC, that works on MB level.
	 *
	 * @param avctx the codec context
	 * @param buf the slice data buffer base
	 * @param buf_size the size of the slice in bytes
	 * @return zero if successful, a negative value otherwise
	 */
	//int (*decode_slice)(AVCodecContext *avctx, const uint8_t *buf, uint32_t buf_size);
	DecodeSlice uintptr
	/**
	 * Called at the end of each frame or field picture.
	 *
	 * The whole picture is parsed at this point and can now be sent
	 * to the hardware accelerator. This function is mandatory.
	 *
	 * @param avctx the codec context
	 * @return zero if successful, a negative value otherwise
	 */
	//int (*end_frame)(AVCodecContext *avctx);
	EndFrame uintptr
	/**
	 * Size of per-frame hardware accelerator private data.
	 *
	 * Private data is allocated with av_mallocz() before
	 * AVCodecContext.get_buffer() and deallocated after
	 * AVCodecContext.release_buffer().
	 */
	FramePrivDataSize ffcommon.FInt

	/**
	 * Called for every Macroblock in a slice.
	 *
	 * XvMC uses it to replace the ff_mpv_reconstruct_mb().
	 * Instead of decoding to raw picture, MB parameters are
	 * stored in an array provided by the video driver.
	 *
	 * @param s the mpeg context
	 */
	//void (*decode_mb)(struct MpegEncContext *s);
	DecodeMb uintptr
	/**
	 * Initialize the hwaccel private data.
	 *
	 * This will be called from ff_get_format(), after hwaccel and
	 * hwaccel_context are set and the hwaccel private data in AVCodecInternal
	 * is allocated.
	 */
	//int (*init)(AVCodecContext *avctx);
	Init uintptr
	/**
	 * Uninitialize the hwaccel private data.
	 *
	 * This will be called from get_format() or avcodec_close(), after hwaccel
	 * and hwaccel_context are already uninitialized.
	 */
	//int (*uninit)(AVCodecContext *avctx);
	Uninit uintptr
	/**
	 * Size of the private data to allocate in
	 * AVCodecInternal.hwaccel_priv_data.
	 */
	PrivDataSize ffcommon.FInt

	/**
	 * Internal hwaccel capabilities.
	 */
	CapsInternal ffcommon.FInt

	/**
	 * Fill the given hw_frames context with current codec parameters. Called
	 * from get_format. Refer to avcodec_get_hw_frames_parameters() for
	 * details.
	 *
	 * This CAN be called before AVHWAccel.init is called, and you must assume
	 * that avctx->hwaccel_priv_data is invalid.
	 */
	//int (*frame_params)(AVCodecContext *avctx, AVBufferRef *hw_frames_ctx);
	FrameParams uintptr
}

/**
* HWAccel is experimental and is thus avoided in favor of non experimental
* codecs
 */
const AV_HWACCEL_CODEC_CAP_EXPERIMENTAL = 0x0200

/**
* Hardware acceleration should be used for decoding even if the codec level
* used is unknown or higher than the maximum supported level reported by the
* hardware driver.
*
* It's generally a good idea to pass this flag unless you have a specific
* reason not to, as hardware tends to under-report supported levels.
 */
const AV_HWACCEL_FLAG_IGNORE_LEVEL = (1 << 0)

/**
* Hardware acceleration can output YUV pixel formats with a different chroma
* sampling than 4:2:0 and/or other than 8 bits per component.
 */
const AV_HWACCEL_FLAG_ALLOW_HIGH_DEPTH = (1 << 1)

/**
* Hardware acceleration should still be attempted for decoding when the
* codec profile does not match the reported capabilities of the hardware.
*
* For example, this can be used to try to decode baseline profile H.264
* streams in hardware - it will often succeed, because many streams marked
* as baseline profile actually conform to constrained baseline profile.
*
* @warning If the stream is actually not supported then the behaviour is
*          undefined, and may include returning entirely incorrect output
*          while indicating success.
 */
const AV_HWACCEL_FLAG_ALLOW_PROFILE_MISMATCH = (1 << 2)

/**
* @}
 */

//#if FF_API_AVPICTURE
/**
* @defgroup lavc_picture AVPicture
*
* Functions for working with AVPicture
* @{
 */

/**
* Picture data structure.
*
* Up to four components can be stored into it, the last component is
* alpha.
* @deprecated use AVFrame or imgutils functions instead
 */
type AVPicture struct {

	//attribute_deprecated
	Data [libavutil.AV_NUM_DATA_POINTERS]*ffcommon.FUint8T ///< pointers to the image data planes
	//attribute_deprecated
	Linesize [libavutil.AV_NUM_DATA_POINTERS]ffcommon.FInt ///< number of bytes per line
}

/**
* @}
 */
//#endif
type AVSubtitleType int32

const (
	SUBTITLE_NONE = iota

	SUBTITLE_BITMAP ///< A bitmap, pict will be set

	/**
	 * Plain text, the text field must be set by the decoder and is
	 * authoritative. ass and pict fields may contain approximations.
	 */
	SUBTITLE_TEXT

	/**
	 * Formatted text, the ass field must be set by the decoder and is
	 * authoritative. pict and text fields may contain approximations.
	 */
	SUBTITLE_ASS
)

const AV_SUBTITLE_FLAG_FORCED = 0x00000001

type AVSubtitleRect struct {
	X        ffcommon.FInt ///< top left corner  of pict, undefined when pict is not set
	Y        ffcommon.FInt ///< top left corner  of pict, undefined when pict is not set
	W        ffcommon.FInt ///< width            of pict, undefined when pict is not set
	H        ffcommon.FInt ///< height           of pict, undefined when pict is not set
	NbColors ffcommon.FInt ///< number of colors in pict, undefined when pict is not set

	//#if FF_API_AVPICTURE
	///**
	//* @deprecated unused
	//*/
	//attribute_deprecated
	Pict AVPicture
	//#endif
	/**
	 * data+linesize for the bitmap of this subtitle.
	 * Can be set for text/ass as well once they are rendered.
	 */
	Data     [4]*ffcommon.FUint8T
	Linesize [4]ffcommon.FInt

	Type AVSubtitleType

	Text ffcommon.FCharPStruct ///< 0 terminated plain UTF-8 text

	/**
	 * 0 terminated ASS/SSA compatible event line.
	 * The presentation of this is unaffected by the other values in this
	 * struct.
	 */
	Ass ffcommon.FCharPStruct

	Flags ffcommon.FInt
}

type AVSubtitle struct {
	Format           ffcommon.FUint16T /* 0 = graphics */
	StartDisplayTime ffcommon.FUint32T /* relative to packet pts, in ms */
	EndDisplayTime   ffcommon.FUint32T /* relative to packet pts, in ms */
	NumRects         ffcommon.FUnsigned
	Rects            **AVSubtitleRect
	Pts              ffcommon.FInt64T ///< Same as packet pts, in AV_TIME_BASE
}

// #if FF_API_NEXT
// /**
// * If c is NULL, returns the first registered codec,
// * if c is non-NULL, returns the next registered codec after c,
// * or NULL if c is the last one.
// */
// attribute_deprecated
// AVCodec *av_codec_next(const AVCodec *c);
func (c *AVCodec) AvCodecNext() (res *AVCodec) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_codec_next").Call(
		uintptr(unsafe.Pointer(c)),
	)
	res = (*AVCodec)(unsafe.Pointer(t))
	return
}

//#endif

/**
* Return the LIBAVCODEC_VERSION_INT constant.
 */
//unsigned avcodec_version(void);
func AvcodecVersion() (res ffcommon.FUnsigned) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_version").Call()
	res = ffcommon.FUnsigned(t)
	return
}

/**
* Return the libavcodec build-time configuration.
 */
//const char *avcodec_configuration(void);
func AvcodecConfiguration() (res ffcommon.FConstCharP) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_configuration").Call()
	res = ffcommon.StringFromPtr(t)
	return
}

/**
* Return the libavcodec license.
 */
//const char *avcodec_license(void);
func AvcodecLicense() (res ffcommon.FConstCharP) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_license").Call()
	res = ffcommon.StringFromPtr(t)
	return
}

//#if FF_API_NEXT
/**
* @deprecated Calling this function is unnecessary.
 */
//attribute_deprecated
//void avcodec_register(AVCodec *codec);
func (codec *AVCodec) AvcodecRegister() {
	ffcommon.GetAvcodecDll().NewProc("avcodec_register").Call(
		uintptr(unsafe.Pointer(codec)),
	)
}

/**
* @deprecated Calling this function is unnecessary.
 */
//attribute_deprecated
//void avcodec_register_all(void);
func AvcodecRegisterAll() {
	ffcommon.GetAvcodecDll().NewProc("avcodec_register_all").Call()
}

//#endif

/**
* Allocate an AVCodecContext and set its fields to default values. The
* resulting struct should be freed with avcodec_free_context().
*
* @param codec if non-NULL, allocate private data and initialize defaults
*              for the given codec. It is illegal to then call avcodec_open2()
*              with a different codec.
*              If NULL, then the codec-specific defaults won't be initialized,
*              which may result in suboptimal default settings (this is
*              important mainly for encoders, e.g. libx264).
*
* @return An AVCodecContext filled with default values or NULL on failure.
 */
//AVCodecContext *avcodec_alloc_context3(const AVCodec *codec);
func (codec *AVCodec) AvcodecAllocContext3() (res *AVCodecContext) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_alloc_context3").Call(
		uintptr(unsafe.Pointer(codec)),
	)
	res = (*AVCodecContext)(unsafe.Pointer(t))
	return
}

/**
* Free the codec context and everything associated with it and write NULL to
* the provided pointer.
 */
//void avcodec_free_context(AVCodecContext **avctx);
func AvcodecFreeContext(avctx **AVCodecContext) {
	ffcommon.GetAvcodecDll().NewProc("avcodec_free_context").Call(
		uintptr(unsafe.Pointer(avctx)),
	)
}

//#if FF_API_GET_CONTEXT_DEFAULTS
/**
* @deprecated This function should not be used, as closing and opening a codec
* context multiple time is not supported. A new codec context should be
* allocated for each new use.
 */
//int avcodec_get_context_defaults3(AVCodecContext *s, const AVCodec *codec);
func (s *AVCodecContext) AvcodecGetContextDefaults3(codec *AVCodec) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_get_context_defaults3").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(codec)),
	)
	res = ffcommon.FInt(t)
	return
}

//#endif

/**
* Get the AVClass for AVCodecContext. It can be used in combination with
* AV_OPT_SEARCH_FAKE_OBJ for examining options.
*
* @see av_opt_find().
 */
//const AVClass *avcodec_get_class(void);
func AvcodecGetClass() (res *AVClass) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_get_class").Call()
	res = (*AVClass)(unsafe.Pointer(t))
	return
}

//#if FF_API_GET_FRAME_CLASS
/**
* @deprecated This function should not be used.
 */
//attribute_deprecated
//const AVClass *avcodec_get_frame_class(void);
func AvcodecGetFrameClass() (res *AVClass) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_get_frame_class").Call()
	res = (*AVClass)(unsafe.Pointer(t))
	return
}

//#endif

/**
* Get the AVClass for AVSubtitleRect. It can be used in combination with
* AV_OPT_SEARCH_FAKE_OBJ for examining options.
*
* @see av_opt_find().
 */
//const AVClass *avcodec_get_subtitle_rect_class(void);
func AvcodecGetSubtitleRectClass() (res *AVClass) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_get_subtitle_rect_class").Call()
	res = (*AVClass)(unsafe.Pointer(t))
	return
}

//#if FF_API_COPY_CONTEXT
/**
* Copy the settings of the source AVCodecContext into the destination
* AVCodecContext. The resulting destination codec context will be
* unopened, i.e. you are required to call avcodec_open2() before you
* can use this AVCodecContext to decode/encode video/audio data.
*
* @param dest target codec context, should be initialized with
*             avcodec_alloc_context3(NULL), but otherwise uninitialized
* @param src source codec context
* @return AVERROR() on error (e.g. memory allocation error), 0 on success
*
* @deprecated The semantics of this function are ill-defined and it should not
* be used. If you need to transfer the stream parameters from one codec context
* to another, use an intermediate AVCodecParameters instance and the
* avcodec_parameters_from_context() / avcodec_parameters_to_context()
* functions.
 */
//attribute_deprecated
//int avcodec_copy_context(AVCodecContext *dest, const AVCodecContext *src);
func AvcodecCopyContext(dest, src *AVCodecContext) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_copy_context").Call(
		uintptr(unsafe.Pointer(dest)),
		uintptr(unsafe.Pointer(src)),
	)
	res = ffcommon.FInt(t)
	return
}

//#endif

/**
* Fill the parameters struct based on the values from the supplied codec
* context. Any allocated fields in par are freed and replaced with duplicates
* of the corresponding fields in codec.
*
* @return >= 0 on success, a negative AVERROR code on failure
 */
//int avcodec_parameters_from_context(AVCodecParameters *par,
//const AVCodecContext *codec);
func (par *AVCodecParameters) AvcodecParametersFromContext(codec *AVCodecContext) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_parameters_from_context").Call(
		uintptr(unsafe.Pointer(par)),
		uintptr(unsafe.Pointer(codec)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* Fill the codec context based on the values from the supplied codec
* parameters. Any allocated fields in codec that have a corresponding field in
* par are freed and replaced with duplicates of the corresponding field in par.
* Fields in codec that do not have a counterpart in par are not touched.
*
* @return >= 0 on success, a negative AVERROR code on failure.
 */
//int avcodec_parameters_to_context(AVCodecContext *codec,
//const AVCodecParameters *par);
func (codec *AVCodecContext) AvcodecParametersToContext(par *AVCodecParameters) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_parameters_to_context").Call(
		uintptr(unsafe.Pointer(codec)),
		uintptr(unsafe.Pointer(par)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* Initialize the AVCodecContext to use the given AVCodec. Prior to using this
* function the context has to be allocated with avcodec_alloc_context3().
*
* The functions avcodec_find_decoder_by_name(), avcodec_find_encoder_by_name(),
* avcodec_find_decoder() and avcodec_find_encoder() provide an easy way for
* retrieving a codec.
*
* @warning This function is not thread safe!
*
* @note Always call this function before using decoding routines (such as
* @ref avcodec_receive_frame()).
*
* @code
* av_dict_set(&opts, "b", "2.5M", 0);
* codec = avcodec_find_decoder(AV_CODEC_ID_H264);
* if (!codec)
*     exit(1);
*
* context = avcodec_alloc_context3(codec);
*
* if (avcodec_open2(context, codec, opts) < 0)
*     exit(1);
* @endcode
*
* @param avctx The context to initialize.
* @param codec The codec to open this context for. If a non-NULL codec has been
*              previously passed to avcodec_alloc_context3() or
*              for this context, then this parameter MUST be either NULL or
*              equal to the previously passed codec.
* @param options A dictionary filled with AVCodecContext and codec-private options.
*                On return this object will be filled with options that were not found.
*
* @return zero on success, a negative value on error
* @see avcodec_alloc_context3(), avcodec_find_decoder(), avcodec_find_encoder(),
*      av_dict_set(), av_opt_find().
 */
type AVDictionary = libavutil.AVDictionary

// int avcodec_open2(AVCodecContext *avctx, const AVCodec *codec, AVDictionary **options);
func (avctx *AVCodecContext) AvcodecOpen2(codec *AVCodec, options **AVDictionary) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_open2").Call(
		uintptr(unsafe.Pointer(avctx)),
		uintptr(unsafe.Pointer(codec)),
		uintptr(unsafe.Pointer(options)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* Close a given AVCodecContext and free all the data associated with it
* (but not the AVCodecContext itself).
*
* Calling this function on an AVCodecContext that hasn't been opened will free
* the codec-specific data allocated in avcodec_alloc_context3() with a non-NULL
* codec. Subsequent calls will do nothing.
*
* @note Do not use this function. Use avcodec_free_context() to destroy a
* codec context (either open or closed). Opening and closing a codec context
* multiple times is not supported anymore -- use multiple codec contexts
* instead.
 */
//int avcodec_close(AVCodecContext *avctx);
func (avctx *AVCodecContext) AvcodecClose() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_close").Call(
		uintptr(unsafe.Pointer(avctx)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* Free all allocated data in the given subtitle struct.
*
* @param sub AVSubtitle to free.
 */
//void avsubtitle_free(AVSubtitle *sub);
func (sub *AVSubtitle) AvsubtitleFree() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avsubtitle_free").Call(
		uintptr(unsafe.Pointer(sub)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* @}
 */

/**
* @addtogroup lavc_decoding
* @{
 */

/**
* The default callback for AVCodecContext.get_buffer2(). It is made public so
* it can be called by custom get_buffer2() implementations for decoders without
* AV_CODEC_CAP_DR1 set.
 */
//int avcodec_default_get_buffer2(AVCodecContext *s, AVFrame *frame, int flags);
func (s *AVCodecContext) AvcodecDefaultGetBuffer2(frame *AVFrame, flags ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_default_get_buffer2").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(frame)),
		uintptr(flags),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* The default callback for AVCodecContext.get_encode_buffer(). It is made public so
* it can be called by custom get_encode_buffer() implementations for encoders without
* AV_CODEC_CAP_DR1 set.
 */
//int avcodec_default_get_encode_buffer(AVCodecContext *s, AVPacket *pkt, int flags);
func (s *AVCodecContext) AvcodecDefaultGetEncodeBuffer(pkt *AVPacket, flags ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_default_get_encode_buffer").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(pkt)),
		uintptr(flags),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* Modify width and height values so that they will result in a memory
* buffer that is acceptable for the codec if you do not use any horizontal
* padding.
*
* May only be used if a codec with AV_CODEC_CAP_DR1 has been opened.
 */
//void avcodec_align_dimensions(AVCodecContext *s, int *width, int *height);
func (s *AVCodecContext) AvcodecAlignDimensions(width, height *ffcommon.FInt) {
	ffcommon.GetAvcodecDll().NewProc("avcodec_align_dimensions").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(width)),
		uintptr(unsafe.Pointer(height)),
	)
}

/**
* Modify width and height values so that they will result in a memory
* buffer that is acceptable for the codec if you also ensure that all
* line sizes are a multiple of the respective linesize_align[i].
*
* May only be used if a codec with AV_CODEC_CAP_DR1 has been opened.
 */
//void avcodec_align_dimensions2(AVCodecContext *s, int *width, int *height,
//int linesize_align[AV_NUM_DATA_POINTERS]);
func (s *AVCodecContext) AvcodecAlignDimensions2(width, height *ffcommon.FInt,
	linesize_align *[libavutil.AV_NUM_DATA_POINTERS]ffcommon.FInt) {
	ffcommon.GetAvcodecDll().NewProc("avcodec_align_dimensions2").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(width)),
		uintptr(unsafe.Pointer(height)),
		uintptr(unsafe.Pointer(linesize_align)),
	)
}

/**
* Converts AVChromaLocation to swscale x/y chroma position.
*
* The positions represent the chroma (0,0) position in a coordinates system
* with luma (0,0) representing the origin and luma(1,1) representing 256,256
*
* @param xpos  horizontal chroma sample position
* @param ypos  vertical   chroma sample position
 */
//int avcodec_enum_to_chroma_pos(int *xpos, int *ypos, enum AVChromaLocation pos);
func AvcodecEnumToChromaPos(xpos, ypos *ffcommon.FInt, pos AVChromaLocation) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_enum_to_chroma_pos").Call(
		uintptr(unsafe.Pointer(xpos)),
		uintptr(unsafe.Pointer(ypos)),
		uintptr(pos),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* Converts swscale x/y chroma position to AVChromaLocation.
*
* The positions represent the chroma (0,0) position in a coordinates system
* with luma (0,0) representing the origin and luma(1,1) representing 256,256
*
* @param xpos  horizontal chroma sample position
* @param ypos  vertical   chroma sample position
 */
//enum AVChromaLocation avcodec_chroma_pos_to_enum(int xpos, int ypos);
func AvcodecChromaPosToEnum(xpos, ypos ffcommon.FInt) (res AVChromaLocation) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_chroma_pos_to_enum").Call(
		uintptr(xpos),
		uintptr(ypos),
	)
	res = AVChromaLocation(t)
	return
}

//#if FF_API_OLD_ENCDEC
/**
 * Decode the audio frame of size avpkt->size from avpkt->data into frame.
 *
 * Some decoders may support multiple frames in a single AVPacket. Such
 * decoders would then just decode the first frame and the return value would be
 * less than the packet size. In this case, avcodec_decode_audio4 has to be
 * called again with an AVPacket containing the remaining data in order to
 * decode the second frame, etc...  Even if no frames are returned, the packet
 * needs to be fed to the decoder with remaining data until it is completely
 * consumed or an error occurs.
 *
 * Some decoders (those marked with AV_CODEC_CAP_DELAY) have a delay between input
 * and output. This means that for some packets they will not immediately
 * produce decoded output and need to be flushed at the end of decoding to get
 * all the decoded data. Flushing is done by calling this function with packets
 * with avpkt->data set to NULL and avpkt->size set to 0 until it stops
 * returning samples. It is safe to flush even those decoders that are not
 * marked with AV_CODEC_CAP_DELAY, then no samples will be returned.
 *
 * @warning The input buffer, avpkt->data must be AV_INPUT_BUFFER_PADDING_SIZE
 *          larger than the actual read bytes because some optimized bitstream
 *          readers read 32 or 64 bits at once and could read over the end.
 *
 * @note The AVCodecContext MUST have been opened with @ref avcodec_open2()
 * before packets may be fed to the decoder.
 *
 * @param      avctx the codec context
 * @param[out] frame The AVFrame in which to store decoded audio samples.
 *                   The decoder will allocate a buffer for the decoded frame by
 *                   calling the AVCodecContext.get_buffer2() callback.
 *                   When AVCodecContext.refcounted_frames is set to 1, the frame is
 *                   reference counted and the returned reference belongs to the
 *                   caller. The caller must release the frame using av_frame_unref()
 *                   when the frame is no longer needed. The caller may safely write
 *                   to the frame if av_frame_is_writable() returns 1.
 *                   When AVCodecContext.refcounted_frames is set to 0, the returned
 *                   reference belongs to the decoder and is valid only until the
 *                   next call to this function or until closing or flushing the
 *                   decoder. The caller may not write to it.
 * @param[out] got_frame_ptr Zero if no frame could be decoded, otherwise it is
 *                           non-zero. Note that this field being set to zero
 *                           does not mean that an error has occurred. For
 *                           decoders with AV_CODEC_CAP_DELAY set, no given decode
 *                           call is guaranteed to produce a frame.
 * @param[in]  avpkt The input AVPacket containing the input buffer.
 *                   At least avpkt->data and avpkt->size should be set. Some
 *                   decoders might also require additional fields to be set.
 * @return A negative error code is returned if an error occurred during
 *         decoding, otherwise the number of bytes consumed from the input
 *         AVPacket is returned.
 *
* @deprecated Use avcodec_send_packet() and avcodec_receive_frame().
*/
//attribute_deprecated
//int avcodec_decode_audio4(AVCodecContext *avctx, AVFrame *frame,
//int *got_frame_ptr, const AVPacket *avpkt);
func (avctx *AVCodecContext) AvcodecDecodeAudio4(frame *AVFrame, got_frame_ptr *ffcommon.FInt, avpkt *AVPacket) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_decode_audio4").Call(
		uintptr(unsafe.Pointer(avctx)),
		uintptr(unsafe.Pointer(frame)),
		uintptr(unsafe.Pointer(got_frame_ptr)),
		uintptr(unsafe.Pointer(avpkt)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* Decode the video frame of size avpkt->size from avpkt->data into picture.
* Some decoders may support multiple frames in a single AVPacket, such
* decoders would then just decode the first frame.
*
* @warning The input buffer must be AV_INPUT_BUFFER_PADDING_SIZE larger than
* the actual read bytes because some optimized bitstream readers read 32 or 64
* bits at once and could read over the end.
*
* @warning The end of the input buffer buf should be set to 0 to ensure that
* no overreading happens for damaged MPEG streams.
*
* @note Codecs which have the AV_CODEC_CAP_DELAY capability set have a delay
* between input and output, these need to be fed with avpkt->data=NULL,
* avpkt->size=0 at the end to return the remaining frames.
*
* @note The AVCodecContext MUST have been opened with @ref avcodec_open2()
* before packets may be fed to the decoder.
*
* @param avctx the codec context
* @param[out] picture The AVFrame in which the decoded video frame will be stored.
*             Use av_frame_alloc() to get an AVFrame. The codec will
*             allocate memory for the actual bitmap by calling the
*             AVCodecContext.get_buffer2() callback.
*             When AVCodecContext.refcounted_frames is set to 1, the frame is
*             reference counted and the returned reference belongs to the
*             caller. The caller must release the frame using av_frame_unref()
*             when the frame is no longer needed. The caller may safely write
*             to the frame if av_frame_is_writable() returns 1.
*             When AVCodecContext.refcounted_frames is set to 0, the returned
*             reference belongs to the decoder and is valid only until the
*             next call to this function or until closing or flushing the
*             decoder. The caller may not write to it.
*
* @param[in] avpkt The input AVPacket containing the input buffer.
*            You can create such packet with av_init_packet() and by then setting
*            data and size, some decoders might in addition need other fields like
*            flags&AV_PKT_FLAG_KEY. All decoders are designed to use the least
*            fields possible.
* @param[in,out] got_picture_ptr Zero if no frame could be decompressed, otherwise, it is nonzero.
* @return On error a negative value is returned, otherwise the number of bytes
* used or zero if no frame could be decompressed.
*
* @deprecated Use avcodec_send_packet() and avcodec_receive_frame().
 */
//attribute_deprecated
//int avcodec_decode_video2(AVCodecContext *avctx, AVFrame *picture,
//int *got_picture_ptr,
//const AVPacket *avpkt);
func (avctx *AVCodecContext) AvcodecDecodeVideo2(picture *AVFrame, got_picture_ptr *ffcommon.FInt, avpkt *AVPacket) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_decode_video2").Call(
		uintptr(unsafe.Pointer(avctx)),
		uintptr(unsafe.Pointer(got_picture_ptr)),
		uintptr(unsafe.Pointer(avpkt)),
	)
	res = ffcommon.FInt(t)
	return
}

//#endif

/**
* Decode a subtitle message.
* Return a negative value on error, otherwise return the number of bytes used.
* If no subtitle could be decompressed, got_sub_ptr is zero.
* Otherwise, the subtitle is stored in *sub.
* Note that AV_CODEC_CAP_DR1 is not available for subtitle codecs. This is for
* simplicity, because the performance difference is expected to be negligible
* and reusing a get_buffer written for video codecs would probably perform badly
* due to a potentially very different allocation pattern.
*
* Some decoders (those marked with AV_CODEC_CAP_DELAY) have a delay between input
* and output. This means that for some packets they will not immediately
* produce decoded output and need to be flushed at the end of decoding to get
* all the decoded data. Flushing is done by calling this function with packets
* with avpkt->data set to NULL and avpkt->size set to 0 until it stops
* returning subtitles. It is safe to flush even those decoders that are not
* marked with AV_CODEC_CAP_DELAY, then no subtitles will be returned.
*
* @note The AVCodecContext MUST have been opened with @ref avcodec_open2()
* before packets may be fed to the decoder.
*
* @param avctx the codec context
* @param[out] sub The preallocated AVSubtitle in which the decoded subtitle will be stored,
*                 must be freed with avsubtitle_free if *got_sub_ptr is set.
* @param[in,out] got_sub_ptr Zero if no subtitle could be decompressed, otherwise, it is nonzero.
* @param[in] avpkt The input AVPacket containing the input buffer.
 */
//int avcodec_decode_subtitle2(AVCodecContext *avctx, AVSubtitle *sub,
//int *got_sub_ptr,
//AVPacket *avpkt);
func (avctx *AVCodecContext) AvcodecDecodeSubtitle2(sub *AVSubtitle,
	got_sub_ptr *ffcommon.FInt,
	avpkt *AVPacket) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_decode_subtitle2").Call(
		uintptr(unsafe.Pointer(avctx)),
		uintptr(unsafe.Pointer(sub)),
		uintptr(unsafe.Pointer(got_sub_ptr)),
		uintptr(unsafe.Pointer(avpkt)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* Supply raw packet data as input to a decoder.
*
* Internally, this call will copy relevant AVCodecContext fields, which can
* influence decoding per-packet, and apply them when the packet is actually
* decoded. (For example AVCodecContext.skip_frame, which might direct the
* decoder to drop the frame contained by the packet sent with this function.)
*
* @warning The input buffer, avpkt->data must be AV_INPUT_BUFFER_PADDING_SIZE
*          larger than the actual read bytes because some optimized bitstream
*          readers read 32 or 64 bits at once and could read over the end.
*
* @warning Do not mix this API with the legacy API (like avcodec_decode_video2())
*          on the same AVCodecContext. It will return unexpected results now
*          or in future libavcodec versions.
*
* @note The AVCodecContext MUST have been opened with @ref avcodec_open2()
*       before packets may be fed to the decoder.
*
* @param avctx codec context
* @param[in] avpkt The input AVPacket. Usually, this will be a single video
*                  frame, or several complete audio frames.
*                  Ownership of the packet remains with the caller, and the
*                  decoder will not write to the packet. The decoder may create
*                  a reference to the packet data (or copy it if the packet is
*                  not reference-counted).
*                  Unlike with older APIs, the packet is always fully consumed,
*                  and if it contains multiple frames (e.g. some audio codecs),
*                  will require you to call avcodec_receive_frame() multiple
*                  times afterwards before you can send a new packet.
*                  It can be NULL (or an AVPacket with data set to NULL and
*                  size set to 0); in this case, it is considered a flush
*                  packet, which signals the end of the stream. Sending the
*                  first flush packet will return success. Subsequent ones are
*                  unnecessary and will return AVERROR_EOF. If the decoder
*                  still has frames buffered, it will return them after sending
*                  a flush packet.
*
* @return 0 on success, otherwise negative error code:
*      AVERROR(EAGAIN):   input is not accepted in the current state - user
*                         must read output with avcodec_receive_frame() (once
*                         all output is read, the packet should be resent, and
*                         the call will not fail with EAGAIN).
*      AVERROR_EOF:       the decoder has been flushed, and no new packets can
*                         be sent to it (also returned if more than 1 flush
*                         packet is sent)
*      AVERROR(EINVAL):   codec not opened, it is an encoder, or requires flush
*      AVERROR(ENOMEM):   failed to add packet to internal queue, or similar
*      other errors: legitimate decoding errors
 */
//int avcodec_send_packet(AVCodecContext *avctx, const AVPacket *avpkt);
func (avctx *AVCodecContext) AvcodecSendPacket(avpkt *AVPacket) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_send_packet").Call(
		uintptr(unsafe.Pointer(avctx)),
		uintptr(unsafe.Pointer(avpkt)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* Return decoded output data from a decoder.
*
* @param avctx codec context
* @param frame This will be set to a reference-counted video or audio
*              frame (depending on the decoder type) allocated by the
*              decoder. Note that the function will always call
*              av_frame_unref(frame) before doing anything else.
*
* @return
*      0:                 success, a frame was returned
*      AVERROR(EAGAIN):   output is not available in this state - user must try
*                         to send new input
*      AVERROR_EOF:       the decoder has been fully flushed, and there will be
*                         no more output frames
*      AVERROR(EINVAL):   codec not opened, or it is an encoder
*      AVERROR_INPUT_CHANGED:   current decoded frame has changed parameters
*                               with respect to first decoded frame. Applicable
*                               when flag AV_CODEC_FLAG_DROPCHANGED is set.
*      other negative values: legitimate decoding errors
 */
//int avcodec_receive_frame(AVCodecContext *avctx, AVFrame *frame);
func (avctx *AVCodecContext) AvcodecReceiveFrame(frame *AVFrame) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_receive_frame").Call(
		uintptr(unsafe.Pointer(avctx)),
		uintptr(unsafe.Pointer(frame)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* Supply a raw video or audio frame to the encoder. Use avcodec_receive_packet()
* to retrieve buffered output packets.
*
* @param avctx     codec context
* @param[in] frame AVFrame containing the raw audio or video frame to be encoded.
*                  Ownership of the frame remains with the caller, and the
*                  encoder will not write to the frame. The encoder may create
*                  a reference to the frame data (or copy it if the frame is
*                  not reference-counted).
*                  It can be NULL, in which case it is considered a flush
*                  packet.  This signals the end of the stream. If the encoder
*                  still has packets buffered, it will return them after this
*                  call. Once flushing mode has been entered, additional flush
*                  packets are ignored, and sending frames will return
*                  AVERROR_EOF.
*
*                  For audio:
*                  If AV_CODEC_CAP_VARIABLE_FRAME_SIZE is set, then each frame
*                  can have any number of samples.
*                  If it is not set, frame->nb_samples must be equal to
*                  avctx->frame_size for all frames except the last.
*                  The final frame may be smaller than avctx->frame_size.
* @return 0 on success, otherwise negative error code:
*      AVERROR(EAGAIN):   input is not accepted in the current state - user
*                         must read output with avcodec_receive_packet() (once
*                         all output is read, the packet should be resent, and
*                         the call will not fail with EAGAIN).
*      AVERROR_EOF:       the encoder has been flushed, and no new frames can
*                         be sent to it
*      AVERROR(EINVAL):   codec not opened, refcounted_frames not set, it is a
*                         decoder, or requires flush
*      AVERROR(ENOMEM):   failed to add packet to internal queue, or similar
*      other errors: legitimate encoding errors
 */
//int avcodec_send_frame(AVCodecContext *avctx, const AVFrame *frame);
func (avctx *AVCodecContext) AvcodecSendFrame(frame *AVFrame) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_send_frame").Call(
		uintptr(unsafe.Pointer(avctx)),
		uintptr(unsafe.Pointer(frame)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* Read encoded data from the encoder.
*
* @param avctx codec context
* @param avpkt This will be set to a reference-counted packet allocated by the
*              encoder. Note that the function will always call
*              av_packet_unref(avpkt) before doing anything else.
* @return 0 on success, otherwise negative error code:
*      AVERROR(EAGAIN):   output is not available in the current state - user
*                         must try to send input
*      AVERROR_EOF:       the encoder has been fully flushed, and there will be
*                         no more output packets
*      AVERROR(EINVAL):   codec not opened, or it is a decoder
*      other errors: legitimate encoding errors
 */
//int avcodec_receive_packet(AVCodecContext *avctx, AVPacket *avpkt);
func (avctx *AVCodecContext) AvcodecReceivePacket(avpkt *AVPacket) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_receive_packet").Call(
		uintptr(unsafe.Pointer(avctx)),
		uintptr(unsafe.Pointer(avpkt)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* Create and return a AVHWFramesContext with values adequate for hardware
* decoding. This is meant to get called from the get_format callback, and is
* a helper for preparing a AVHWFramesContext for AVCodecContext.hw_frames_ctx.
* This API is for decoding with certain hardware acceleration modes/APIs only.
*
* The returned AVHWFramesContext is not initialized. The caller must do this
* with av_hwframe_ctx_init().
*
* Calling this function is not a requirement, but makes it simpler to avoid
* codec or hardware API specific details when manually allocating frames.
*
* Alternatively to this, an API user can set AVCodecContext.hw_device_ctx,
* which sets up AVCodecContext.hw_frames_ctx fully automatically, and makes
* it unnecessary to call this function or having to care about
* AVHWFramesContext initialization at all.
*
* There are a number of requirements for calling this function:
*
* - It must be called from get_format with the same avctx parameter that was
*   passed to get_format. Calling it outside of get_format is not allowed, and
*   can trigger undefined behavior.
* - The function is not always supported (see description of return values).
*   Even if this function returns successfully, hwaccel initialization could
*   fail later. (The degree to which implementations check whether the stream
*   is actually supported varies. Some do this check only after the user's
*   get_format callback returns.)
* - The hw_pix_fmt must be one of the choices suggested by get_format. If the
*   user decides to use a AVHWFramesContext prepared with this API function,
*   the user must return the same hw_pix_fmt from get_format.
* - The device_ref passed to this function must support the given hw_pix_fmt.
* - After calling this API function, it is the user's responsibility to
*   initialize the AVHWFramesContext (returned by the out_frames_ref parameter),
*   and to set AVCodecContext.hw_frames_ctx to it. If done, this must be done
*   before returning from get_format (this is implied by the normal
*   AVCodecContext.hw_frames_ctx API rules).
* - The AVHWFramesContext parameters may change every time time get_format is
*   called. Also, AVCodecContext.hw_frames_ctx is reset before get_format. So
*   you are inherently required to go through this process again on every
*   get_format call.
* - It is perfectly possible to call this function without actually using
*   the resulting AVHWFramesContext. One use-case might be trying to reuse a
*   previously initialized AVHWFramesContext, and calling this API function
*   only to test whether the required frame parameters have changed.
* - Fields that use dynamically allocated values of any kind must not be set
*   by the user unless setting them is explicitly allowed by the documentation.
*   If the user sets AVHWFramesContext.free and AVHWFramesContext.user_opaque,
*   the new free callback must call the potentially set previous free callback.
*   This API call may set any dynamically allocated fields, including the free
*   callback.
*
* The function will set at least the following fields on AVHWFramesContext
* (potentially more, depending on hwaccel API):
*
* - All fields set by av_hwframe_ctx_alloc().
* - Set the format field to hw_pix_fmt.
* - Set the sw_format field to the most suited and most versatile format. (An
*   implication is that this will prefer generic formats over opaque formats
*   with arbitrary restrictions, if possible.)
* - Set the width/height fields to the coded frame size, rounded up to the
*   API-specific minimum alignment.
* - Only _if_ the hwaccel requires a pre-allocated pool: set the initial_pool_size
*   field to the number of maximum reference surfaces possible with the codec,
*   plus 1 surface for the user to work (meaning the user can safely reference
*   at most 1 decoded surface at a time), plus additional buffering introduced
*   by frame threading. If the hwaccel does not require pre-allocation, the
*   field is left to 0, and the decoder will allocate new surfaces on demand
*   during decoding.
* - Possibly AVHWFramesContext.hwctx fields, depending on the underlying
*   hardware API.
*
* Essentially, out_frames_ref returns the same as av_hwframe_ctx_alloc(), but
* with basic frame parameters set.
*
* The function is stateless, and does not change the AVCodecContext or the
* device_ref AVHWDeviceContext.
*
* @param avctx The context which is currently calling get_format, and which
*              implicitly contains all state needed for filling the returned
*              AVHWFramesContext properly.
* @param device_ref A reference to the AVHWDeviceContext describing the device
*                   which will be used by the hardware decoder.
* @param hw_pix_fmt The hwaccel format you are going to return from get_format.
* @param out_frames_ref On success, set to a reference to an _uninitialized_
*                       AVHWFramesContext, created from the given device_ref.
*                       Fields will be set to values required for decoding.
*                       Not changed if an error is returned.
* @return zero on success, a negative value on error. The following error codes
*         have special semantics:
*      AVERROR(ENOENT): the decoder does not support this functionality. Setup
*                       is always manual, or it is a decoder which does not
*                       support setting AVCodecContext.hw_frames_ctx at all,
*                       or it is a software format.
*      AVERROR(EINVAL): it is known that hardware decoding is not supported for
*                       this configuration, or the device_ref is not supported
*                       for the hwaccel referenced by hw_pix_fmt.
 */
//int avcodec_get_hw_frames_parameters(AVCodecContext *avctx,
//AVBufferRef *device_ref,
//enum AVPixelFormat hw_pix_fmt,
//AVBufferRef **out_frames_ref);
func (avctx *AVCodecContext) AvcodecGetHwFramesParameters(device_ref *AVBufferRef,
	hw_pix_fmt AVPixelFormat,
	out_frames_ref **AVBufferRef) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_get_hw_frames_parameters").Call(
		uintptr(unsafe.Pointer(avctx)),
		uintptr(unsafe.Pointer(device_ref)),
		uintptr(hw_pix_fmt),
		uintptr(unsafe.Pointer(out_frames_ref)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* @defgroup lavc_parsing Frame parsing
* @{
 */
type AVPictureStructure int32

const (
	AV_PICTURE_STRUCTURE_UNKNOWN      = iota //< unknown
	AV_PICTURE_STRUCTURE_TOP_FIELD           //< coded as top field
	AV_PICTURE_STRUCTURE_BOTTOM_FIELD        //< coded as bottom field
	AV_PICTURE_STRUCTURE_FRAME               //< coded as frame
)
const AV_PARSER_PTS_NB = 4

// type AVRational = libavutil.AVRational
type AVCodecParserContext struct {
	PrivData    ffcommon.FVoidP
	Parser      *AVCodecParser
	FrameOffset ffcommon.FInt64T /* offset of the current frame */
	CurOffset   ffcommon.FInt64T /* current offset
	(incremented by each av_parser_parse()) */
	NextFrameOffset ffcommon.FInt64T /* offset of the next frame */
	/* video info */
	PictType ffcommon.FInt /* XXX: Put it back in AVCodecContext. */
	/**
	 * This field is used for proper frame duration computation in lavf.
	 * It signals, how much longer the frame duration of the current frame
	 * is compared to normal frame duration.
	 *
	 * frame_duration = (1 + repeat_pict) * time_base
	 *
	 * It is used by codecs like H.264 to display telecined material.
	 */
	RepeatPict ffcommon.FInt    /* XXX: Put it back in AVCodecContext. */
	Pts        ffcommon.FInt64T /* pts of the current frame */
	Dts        ffcommon.FInt64T /* dts of the current frame */

	/* private data */
	LastPts        ffcommon.FInt64T
	LastDts        ffcommon.FInt64T
	FetchTimestamp ffcommon.FInt

	//const AV_PARSER_PTS_NB= 4
	CurFrameStartIndex ffcommon.FInt
	CurFrameOffset     [4]ffcommon.FInt64T
	CurFramePts        [4]ffcommon.FInt64T
	CurFrameDts        [AV_PARSER_PTS_NB]ffcommon.FInt64T

	Flags ffcommon.FInt
	//const PARSER_FLAG_COMPLETE_FRAMES      =     0x0001
	//const PARSER_FLAG_ONCE                 =     0x0002
	///// Set if the parser has a valid file offset
	//const PARSER_FLAG_FETCHED_OFFSET       =     0x0004
	//const PARSER_FLAG_USE_CODEC_TS         =     0x1000

	Offset      ffcommon.FInt64T ///< byte offset from starting packet start
	CurFrameEnd [AV_PARSER_PTS_NB]ffcommon.FInt64T

	/**
	 * Set by parser to 1 for key frames and 0 for non-key frames.
	 * It is initialized to -1, so if the parser doesn't set this flag,
	 * old-style fallback using AV_PICTURE_TYPE_I picture type as key frames
	 * will be used.
	 */
	KeyFrame ffcommon.FInt

	//#if FF_API_CONVERGENCE_DURATION
	///**
	//* @deprecated unused
	//*/
	//attribute_deprecated
	ConvergenceDuration ffcommon.FInt64T
	//#endif

	// Timestamp generation support:
	/**
	 * Synchronization point for start of timestamp generation.
	 *
	 * Set to >0 for sync point, 0 for no sync point and <0 for undefined
	 * (default).
	 *
	 * For example, this corresponds to presence of H.264 buffering period
	 * SEI message.
	 */
	DtsSyncPoint ffcommon.FInt

	/**
	 * Offset of the current timestamp against last timestamp sync point in
	 * units of AVCodecContext.time_base.
	 *
	 * Set to INT_MIN when dts_sync_point unused. Otherwise, it must
	 * contain a valid timestamp offset.
	 *
	 * Note that the timestamp of sync point has usually a nonzero
	 * dts_ref_dts_delta, which refers to the previous sync point. Offset of
	 * the next frame after timestamp sync point will be usually 1.
	 *
	 * For example, this corresponds to H.264 cpb_removal_delay.
	 */
	DtsRefDtsDelta ffcommon.FInt

	/**
	 * Presentation delay of current frame in units of AVCodecContext.time_base.
	 *
	 * Set to INT_MIN when dts_sync_point unused. Otherwise, it must
	 * contain valid non-negative timestamp delta (presentation time of a frame
	 * must not lie in the past).
	 *
	 * This delay represents the difference between decoding and presentation
	 * time of the frame.
	 *
	 * For example, this corresponds to H.264 dpb_output_delay.
	 */
	PtsDtsDelta ffcommon.FInt

	/**
	 * Position of the packet in file.
	 *
	 * Analogous to cur_frame_pts/dts
	 */
	CurFramePos [AV_PARSER_PTS_NB]ffcommon.FInt64T

	/**
	 * Byte position of currently parsed frame in stream.
	 */
	Pos ffcommon.FInt64T

	/**
	 * Previous frame byte position.
	 */
	LastPos ffcommon.FInt64T

	/**
	 * Duration of the current frame.
	 * For audio, this is in units of 1 / AVCodecContext.sample_rate.
	 * For all other types, this is in units of AVCodecContext.time_base.
	 */
	Duration ffcommon.FInt

	FieldOrder AVFieldOrder

	/**
	 * Indicate whether a picture is coded as a frame, top field or bottom field.
	 *
	 * For example, H.264 field_pic_flag equal to 0 corresponds to
	 * AV_PICTURE_STRUCTURE_FRAME. An H.264 picture with field_pic_flag
	 * equal to 1 and bottom_field_flag equal to 0 corresponds to
	 * AV_PICTURE_STRUCTURE_TOP_FIELD.
	 */
	PictureStructure AVPictureStructure

	/**
	 * Picture number incremented in presentation or output order.
	 * This field may be reinitialized at the first picture of a new sequence.
	 *
	 * For example, this corresponds to H.264 PicOrderCnt.
	 */
	OutputPictureNumber ffcommon.FInt

	/**
	 * Dimensions of the decoded video intended for presentation.
	 */
	Width  ffcommon.FInt
	Height ffcommon.FInt

	/**
	 * Dimensions of the coded video.
	 */
	CodedWidth  ffcommon.FInt
	CodedHeight ffcommon.FInt

	/**
	 * The format of the coded data, corresponds to enum AVPixelFormat for video
	 * and for enum AVSampleFormat for audio.
	 *
	 * Note that a decoder can have considerable freedom in how exactly it
	 * decodes the data, so the format reported here might be different from the
	 * one returned by a decoder.
	 */
	Format ffcommon.FInt
}

type AVCodecParser struct {
	CodecIds     [5]ffcommon.FInt /* several codec IDs are permitted */
	PrivDataSize ffcommon.FInt
	//int (*parser_init)(AVCodecParserContext *s);
	ParserInit uintptr
	/* This callback never returns an error, a negative value means that
	 * the frame start was in a previous packet. */
	//int (*parser_parse)(AVCodecParserContext *s,
	//AVCodecContext *avctx,
	//const uint8_t **poutbuf, int *poutbuf_size,
	//const uint8_t *buf, int buf_size);
	//void (*parser_close)(AVCodecParserContext *s);
	ParserParse uintptr
	ParserClose uintptr
	//int (*split)(AVCodecContext *avctx, const uint8_t *buf, int buf_size);
	Split uintptr
	//#if FF_API_NEXT
	//attribute_deprecated
	Next *AVCodecParser
	//#endif
}

/**
* Iterate over all registered codec parsers.
*
* @param opaque a pointer where libavcodec will store the iteration state. Must
*               point to NULL to start the iteration.
*
* @return the next registered codec parser or NULL when the iteration is
*         finished
 */
//const AVCodecParser *av_parser_iterate(void **opaque);
func AvParserIterate(opaque *ffcommon.FVoidP) (res *AVCodecParser) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_parser_iterate").Call(
		uintptr(unsafe.Pointer(opaque)),
	)
	res = (*AVCodecParser)(unsafe.Pointer(t))
	return
}

// #if FF_API_NEXT
// attribute_deprecated
// AVCodecParser *av_parser_next(const AVCodecParser *c);
func (c *AVCodecParser) AvParserNext() (res *AVCodecParser) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_parser_next").Call(
		uintptr(unsafe.Pointer(c)),
	)
	res = (*AVCodecParser)(unsafe.Pointer(t))
	return
}

// attribute_deprecated
// void av_register_codec_parser(AVCodecParser *parser);
func (parser *AVCodecParser) AvRegisterCodecParser() {
	ffcommon.GetAvcodecDll().NewProc("av_register_codec_parser").Call(
		uintptr(unsafe.Pointer(parser)),
	)
}

// #endif
// AVCodecParserContext *av_parser_init(int codec_id);
func AvParserInit(codec_id ffcommon.FInt) (res *AVCodecParserContext) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_parser_init").Call(
		uintptr(codec_id),
	)
	res = (*AVCodecParserContext)(unsafe.Pointer(t))
	return
}

/**
* Parse a packet.
*
* @param s             parser context.
* @param avctx         codec context.
* @param poutbuf       set to pointer to parsed buffer or NULL if not yet finished.
* @param poutbuf_size  set to size of parsed buffer or zero if not yet finished.
* @param buf           input buffer.
* @param buf_size      buffer size in bytes without the padding. I.e. the full buffer
                      size is assumed to be buf_size + AV_INPUT_BUFFER_PADDING_SIZE.
                      To signal EOF, this should be 0 (so that the last frame
                      can be output).
* @param pts           input presentation timestamp.
* @param dts           input decoding timestamp.
* @param pos           input byte position in stream.
* @return the number of bytes of the input bitstream used.
*
* Example:
* @code
*   while(in_len){
*       len = av_parser_parse2(myparser, AVCodecContext, &data, &size,
*                                        in_data, in_len,
*                                        pts, dts, pos);
*       in_data += len;
*       in_len  -= len;
*
*       if(size)
*          decode_frame(data, size);
*   }
* @endcode
*/
//int av_parser_parse2(AVCodecParserContext *s,
//AVCodecContext *avctx,
//uint8_t **poutbuf, int *poutbuf_size,
//const uint8_t *buf, int buf_size,
//int64_t pts, int64_t dts,
//int64_t pos);
func (s *AVCodecParserContext) AvParserParse2(avctx *AVCodecContext,
	poutbuf **ffcommon.FUint8T, poutbuf_size *ffcommon.FInt,
	buf *ffcommon.FUint8T, buf_size ffcommon.FInt,
	pts, dts,
	pos ffcommon.FInt64T) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_parser_parse2").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(avctx)),
		uintptr(unsafe.Pointer(poutbuf)),
		uintptr(unsafe.Pointer(poutbuf_size)),
		uintptr(unsafe.Pointer(buf)),
		uintptr(buf_size),
		uintptr(pts),
		uintptr(dts),
		uintptr(pos),
	)
	res = ffcommon.FInt(t)
	return
}

//#if FF_API_PARSER_CHANGE
/**
* @return 0 if the output buffer is a subset of the input, 1 if it is allocated and must be freed
* @deprecated Use dump_extradata, remove_extra or extract_extradata
*             bitstream filters instead.
 */
//attribute_deprecated
//int av_parser_change(AVCodecParserContext *s,
//AVCodecContext *avctx,
//uint8_t **poutbuf, int *poutbuf_size,
//const uint8_t *buf, int buf_size, int keyframe);
//#endif
func (s *AVCodecParserContext) AvParserChange(avctx *AVCodecContext, poutbuf **ffcommon.FUint8T, poutbuf_size *ffcommon.FInt, buf *ffcommon.FUint8T, buf_size, keyframe ffcommon.FInt) (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_parser_change").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(avctx)),
		uintptr(unsafe.Pointer(poutbuf)),
		uintptr(unsafe.Pointer(poutbuf_size)),
		uintptr(unsafe.Pointer(buf)),
		uintptr(buf_size),
		uintptr(keyframe),
	)
	res = ffcommon.StringFromPtr(t)
	return
}

// void av_parser_close(AVCodecParserContext *s);
func (s *AVCodecParserContext) AvParserClose() {
	ffcommon.GetAvcodecDll().NewProc("av_parser_close").Call()
}

/**
* @}
* @}
 */

/**
* @addtogroup lavc_encoding
* @{
 */

//#if FF_API_OLD_ENCDEC
/**
* Encode a frame of audio.
*
* Takes input samples from frame and writes the next output packet, if
* available, to avpkt. The output packet does not necessarily contain data for
* the most recent frame, as encoders can delay, split, and combine input frames
* internally as needed.
*
* @param avctx     codec context
* @param avpkt     output AVPacket.
*                  The user can supply an output buffer by setting
*                  avpkt->data and avpkt->size prior to calling the
*                  function, but if the size of the user-provided data is not
*                  large enough, encoding will fail. If avpkt->data and
*                  avpkt->size are set, avpkt->destruct must also be set. All
*                  other AVPacket fields will be reset by the encoder using
*                  av_init_packet(). If avpkt->data is NULL, the encoder will
*                  allocate it. The encoder will set avpkt->size to the size
*                  of the output packet.
*
*                  If this function fails or produces no output, avpkt will be
*                  freed using av_packet_unref().
* @param[in] frame AVFrame containing the raw audio data to be encoded.
*                  May be NULL when flushing an encoder that has the
*                  AV_CODEC_CAP_DELAY capability set.
*                  If AV_CODEC_CAP_VARIABLE_FRAME_SIZE is set, then each frame
*                  can have any number of samples.
*                  If it is not set, frame->nb_samples must be equal to
*                  avctx->frame_size for all frames except the last.
*                  The final frame may be smaller than avctx->frame_size.
* @param[out] got_packet_ptr This field is set to 1 by libavcodec if the
*                            output packet is non-empty, and to 0 if it is
*                            empty. If the function returns an error, the
*                            packet can be assumed to be invalid, and the
*                            value of got_packet_ptr is undefined and should
*                            not be used.
* @return          0 on success, negative error code on failure
*
* @deprecated use avcodec_send_frame()/avcodec_receive_packet() instead.
*             If allowed and required, set AVCodecContext.get_encode_buffer to
*             a custom function to pass user supplied output buffers.
 */
//attribute_deprecated
//int avcodec_encode_audio2(AVCodecContext *avctx, AVPacket *avpkt,
//const AVFrame *frame, int *got_packet_ptr);
func (avctx *AVCodecContext) AvcodecEncodeAudio2(avpkt *AVPacket, frame *AVFrame, got_packet_ptr *ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_encode_audio2").Call(
		uintptr(unsafe.Pointer(avctx)),
		uintptr(unsafe.Pointer(avpkt)),
		uintptr(unsafe.Pointer(frame)),
		uintptr(unsafe.Pointer(got_packet_ptr)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* Encode a frame of video.
*
* Takes input raw video data from frame and writes the next output packet, if
* available, to avpkt. The output packet does not necessarily contain data for
* the most recent frame, as encoders can delay and reorder input frames
* internally as needed.
*
* @param avctx     codec context
* @param avpkt     output AVPacket.
*                  The user can supply an output buffer by setting
*                  avpkt->data and avpkt->size prior to calling the
*                  function, but if the size of the user-provided data is not
*                  large enough, encoding will fail. All other AVPacket fields
*                  will be reset by the encoder using av_init_packet(). If
*                  avpkt->data is NULL, the encoder will allocate it.
*                  The encoder will set avpkt->size to the size of the
*                  output packet. The returned data (if any) belongs to the
*                  caller, he is responsible for freeing it.
*
*                  If this function fails or produces no output, avpkt will be
*                  freed using av_packet_unref().
* @param[in] frame AVFrame containing the raw video data to be encoded.
*                  May be NULL when flushing an encoder that has the
*                  AV_CODEC_CAP_DELAY capability set.
* @param[out] got_packet_ptr This field is set to 1 by libavcodec if the
*                            output packet is non-empty, and to 0 if it is
*                            empty. If the function returns an error, the
*                            packet can be assumed to be invalid, and the
*                            value of got_packet_ptr is undefined and should
*                            not be used.
* @return          0 on success, negative error code on failure
*
* @deprecated use avcodec_send_frame()/avcodec_receive_packet() instead.
*             If allowed and required, set AVCodecContext.get_encode_buffer to
*             a custom function to pass user supplied output buffers.
 */
//attribute_deprecated
//int avcodec_encode_video2(AVCodecContext *avctx, AVPacket *avpkt,
//const AVFrame *frame, int *got_packet_ptr);
//#endif
func (avctx *AVCodecContext) AvcodecEncodeVideo2(avpkt *AVPacket, frame *AVFrame, got_packet_ptr *ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_encode_video2").Call(
		uintptr(unsafe.Pointer(avctx)),
		uintptr(unsafe.Pointer(avpkt)),
		uintptr(unsafe.Pointer(frame)),
		uintptr(unsafe.Pointer(got_packet_ptr)),
	)
	res = ffcommon.FInt(t)
	return
}

// int avcodec_encode_subtitle(AVCodecContext *avctx, uint8_t *buf, int buf_size,
// const AVSubtitle *sub);
func (avctx *AVCodecContext) AvcodecEncodeSubtitle(buf *ffcommon.FUint8T, buf_size ffcommon.FInt, sub *AVSubtitle) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_encode_subtitle").Call(
		uintptr(unsafe.Pointer(avctx)),
		uintptr(unsafe.Pointer(buf)),
		uintptr(buf_size),
		uintptr(unsafe.Pointer(sub)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* @}
 */

//#if FF_API_AVPICTURE
/**
* @addtogroup lavc_picture
* @{
 */

/**
* @deprecated unused
 */
//attribute_deprecated
//int avpicture_alloc(AVPicture *picture, enum AVPixelFormat pix_fmt, int width, int height);
func (picture *AVPicture) AvpictureAlloc(pix_fmt AVPixelFormat, width, height ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avpicture_alloc").Call(
		uintptr(unsafe.Pointer(picture)),
		uintptr(pix_fmt),
		uintptr(width),
		uintptr(height),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* @deprecated unused
 */
//attribute_deprecated
//void avpicture_free(AVPicture *picture);
func (picture *AVPicture) AvpictureFree() {
	ffcommon.GetAvcodecDll().NewProc("avpicture_free").Call(
		uintptr(unsafe.Pointer(picture)),
	)
}

/**
* @deprecated use av_image_fill_arrays() instead.
 */
//attribute_deprecated
//int avpicture_fill(AVPicture *picture, const uint8_t *ptr,
//enum AVPixelFormat pix_fmt, int width, int height);
func (picture *AVPicture) AvpictureFill(ptr *ffcommon.FUint8T, pix_fmt AVPixelFormat, width, height ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avpicture_fill").Call(
		uintptr(unsafe.Pointer(picture)),
		uintptr(unsafe.Pointer(ptr)),
		uintptr(pix_fmt),
		uintptr(width),
		uintptr(height),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* @deprecated use av_image_copy_to_buffer() instead.
 */
//attribute_deprecated
//int avpicture_layout(const AVPicture *src, enum AVPixelFormat pix_fmt,
//int width, int height,
//unsigned char *dest, int dest_size);
func (src *AVPicture) AvpictureLayout(pix_fmt AVPixelFormat, width, height ffcommon.FInt, dest ffcommon.FCharPStruct, dest_size ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avpicture_layout").Call(
		uintptr(unsafe.Pointer(src)),
		uintptr(pix_fmt),
		uintptr(width),
		uintptr(height),
		dest,
		uintptr(dest_size),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* @deprecated use av_image_get_buffer_size() instead.
 */
//attribute_deprecated
//int avpicture_get_size(enum AVPixelFormat pix_fmt, int width, int height);
func AvpictureGetSize(pix_fmt AVPixelFormat, width, height ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avpicture_get_size").Call(
		uintptr(pix_fmt),
		uintptr(width),
		uintptr(height),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* @deprecated av_image_copy() instead.
 */
//attribute_deprecated
//void av_picture_copy(AVPicture *dst, const AVPicture *src,
//enum AVPixelFormat pix_fmt, int width, int height);
func AvPictureCopy(dst, src *AVPicture, pix_fmt AVPixelFormat, width, height ffcommon.FInt) {
	ffcommon.GetAvcodecDll().NewProc("av_picture_copy").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
		uintptr(pix_fmt),
		uintptr(width),
		uintptr(height),
	)
}

/**
* @deprecated unused
 */
//attribute_deprecated
//int av_picture_crop(AVPicture *dst, const AVPicture *src,
//enum AVPixelFormat pix_fmt, int top_band, int left_band);
func AvPictureCrop(dst, src *AVPicture, pix_fmt AVPixelFormat, top_band, left_band ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_picture_crop").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
		uintptr(pix_fmt),
		uintptr(top_band),
		uintptr(left_band),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* @deprecated unused
 */
//attribute_deprecated
//int av_picture_pad(AVPicture *dst, const AVPicture *src, int height, int width, enum AVPixelFormat pix_fmt,
//int padtop, int padbottom, int padleft, int padright, int *color);
func AvPicturePad(dst, src *AVPicture, height, width, pix_fmt AVPixelFormat, padtop, padbottom, padleft, padright ffcommon.FInt, color *ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_picture_pad").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
		uintptr(height),
		uintptr(width),
		uintptr(pix_fmt),
		uintptr(padtop),
		uintptr(padbottom),
		uintptr(padleft),
		uintptr(padright),
		uintptr(unsafe.Pointer(color)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* @}
 */
//#endif

/**
* @defgroup lavc_misc Utility functions
* @ingroup libavc
*
* Miscellaneous utility functions related to both encoding and decoding
* (or neither).
* @{
 */

/**
* @defgroup lavc_misc_pixfmt Pixel formats
*
* Functions for working with pixel formats.
* @{
 */

//#if FF_API_GETCHROMA
/**
* @deprecated Use av_pix_fmt_get_chroma_sub_sample
 */

// attribute_deprecated
// void avcodec_get_chroma_sub_sample(enum AVPixelFormat pix_fmt, int *h_shift, int *v_shift);
// #endif
func AvcodecGetChromaSubSample(pix_fmt AVPixelFormat, h_shift, v_shift *ffcommon.FInt) {
	ffcommon.GetAvcodecDll().NewProc("avcodec_get_chroma_sub_sample").Call(
		uintptr(pix_fmt),
		uintptr(unsafe.Pointer(h_shift)),
		uintptr(unsafe.Pointer(v_shift)),
	)
}

/**
* Return a value representing the fourCC code associated to the
* pixel format pix_fmt, or 0 if no associated fourCC code can be
* found.
 */
//unsigned int avcodec_pix_fmt_to_codec_tag(enum AVPixelFormat pix_fmt);
func AvcodecPixFmtToCodecTag(pix_fmt AVPixelFormat) (res ffcommon.FUnsignedInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_pix_fmt_to_codec_tag").Call(
		uintptr(pix_fmt),
	)
	res = ffcommon.FUnsignedInt(t)
	return
}

/**
* Find the best pixel format to convert to given a certain source pixel
* format.  When converting from one pixel format to another, information loss
* may occur.  For example, when converting from RGB24 to GRAY, the color
* information will be lost. Similarly, other losses occur when converting from
* some formats to other formats. avcodec_find_best_pix_fmt_of_2() searches which of
* the given pixel formats should be used to suffer the least amount of loss.
* The pixel formats from which it chooses one, are determined by the
* pix_fmt_list parameter.
*
*
* @param[in] pix_fmt_list AV_PIX_FMT_NONE terminated array of pixel formats to choose from
* @param[in] src_pix_fmt source pixel format
* @param[in] has_alpha Whether the source pixel format alpha channel is used.
* @param[out] loss_ptr Combination of flags informing you what kind of losses will occur.
* @return The best pixel format to convert to or -1 if none was found.
 */
//enum AVPixelFormat avcodec_find_best_pix_fmt_of_list(const enum AVPixelFormat *pix_fmt_list,
//enum AVPixelFormat src_pix_fmt,
//int has_alpha, int *loss_ptr);
func AvcodecFindBestPixFmtOfList(pix_fmt_list *AVPixelFormat,
	src_pix_fmt AVPixelFormat,
	has_alpha ffcommon.FInt, loss_ptr *ffcommon.FInt) (res AVPixelFormat) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_find_best_pix_fmt_of_list").Call(
		uintptr(unsafe.Pointer(pix_fmt_list)),
		uintptr(src_pix_fmt),
		uintptr(has_alpha),
		uintptr(unsafe.Pointer(loss_ptr)),
	)
	res = AVPixelFormat(t)
	return
}

//#if FF_API_AVCODEC_PIX_FMT
/**
* @deprecated see av_get_pix_fmt_loss()
 */
//attribute_deprecated
//int avcodec_get_pix_fmt_loss(enum AVPixelFormat dst_pix_fmt, enum AVPixelFormat src_pix_fmt,
//int has_alpha);
func AvcodecGetPixFmtLoss(dst_pix_fmt, src_pix_fmt AVPixelFormat, has_alpha ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_get_pix_fmt_loss").Call(
		uintptr(dst_pix_fmt),
		uintptr(src_pix_fmt),
		uintptr(has_alpha),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* @deprecated see av_find_best_pix_fmt_of_2()
 */
// attribute_deprecated
// enum AVPixelFormat avcodec_find_best_pix_fmt_of_2(enum AVPixelFormat dst_pix_fmt1, enum AVPixelFormat dst_pix_fmt2,
// enum AVPixelFormat src_pix_fmt, int has_alpha, int *loss_ptr);
func AvCodecFindBestPixFmtOf2(dst_pix_fmt1, dst_pix_fmt2, src_pix_fmt AVPixelFormat, has_alpha ffcommon.FInt, loss_ptr *ffcommon.FInt) (res AVPixelFormat) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_find_best_pix_fmt_of_2").Call(
		uintptr(dst_pix_fmt1),
		uintptr(dst_pix_fmt2),
		uintptr(src_pix_fmt),
		uintptr(has_alpha),
		uintptr(unsafe.Pointer(loss_ptr)),
	)
	res = AVPixelFormat(t)
	return
}

// attribute_deprecated
// enum AVPixelFormat avcodec_find_best_pix_fmt2(enum AVPixelFormat dst_pix_fmt1, enum AVPixelFormat dst_pix_fmt2,
// enum AVPixelFormat src_pix_fmt, int has_alpha, int *loss_ptr);
func AvcodecFindBestPixFmt2(dst_pix_fmt1, dst_pix_fmt2, src_pix_fmt AVPixelFormat, has_alpha ffcommon.FInt, loss_ptr *ffcommon.FInt) (res AVPixelFormat) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_find_best_pix_fmt2").Call(
		uintptr(dst_pix_fmt1),
		uintptr(dst_pix_fmt2),
		uintptr(src_pix_fmt),
		uintptr(has_alpha),
		uintptr(unsafe.Pointer(loss_ptr)),
	)
	res = AVPixelFormat(t)
	return
}

//#endif

// enum AVPixelFormat avcodec_default_get_format(struct AVCodecContext *s, const enum AVPixelFormat * fmt);
func (s *AVCodecContext) AvcodecDefaultGetFormat(fmt0 AVPixelFormat) (res AVPixelFormat) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_default_get_format").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(fmt0),
	)
	res = AVPixelFormat(t)
	return
}

/**
* @}
 */

//#if FF_API_TAG_STRING
/**
* Put a string representing the codec tag codec_tag in buf.
*
* @param buf       buffer to place codec tag in
* @param buf_size size in bytes of buf
* @param codec_tag codec tag to assign
* @return the length of the string that would have been generated if
* enough space had been available, excluding the trailing null
*
* @deprecated see av_fourcc_make_string() and av_fourcc2str().
 */
//attribute_deprecated
//size_t av_get_codec_tag_string(char *buf, size_t buf_size, unsigned int codec_tag);
func AvGetCodecTagString(buf ffcommon.FCharPStruct, buf_size ffcommon.FSizeT, codec_tag ffcommon.FUnsignedInt) (res ffcommon.FSizeT) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_get_codec_tag_string").Call(
		buf,
		uintptr(buf_size),
		uintptr(codec_tag),
	)
	res = ffcommon.FSizeT(t)
	return
}

//#endif

// void avcodec_string(char *buf, int buf_size, AVCodecContext *enc, int encode);
func AvcodecString(buf ffcommon.FCharPStruct, buf_size ffcommon.FInt, enc *AVCodecContext, encode ffcommon.FInt) {
	ffcommon.GetAvcodecDll().NewProc("avcodec_string").Call(
		buf,
		uintptr(buf_size),
		uintptr(unsafe.Pointer(enc)),
		uintptr(encode),
	)
}

/**
* Return a name for the specified profile, if available.
*
* @param codec the codec that is searched for the given profile
* @param profile the profile value for which a name is requested
* @return A name for the profile if found, NULL otherwise.
 */
//const char *av_get_profile_name(const AVCodec *codec, int profile);
func (codec *AVCodec) AvGetProfileName(profile ffcommon.FInt) (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_get_profile_name").Call(
		uintptr(unsafe.Pointer(codec)),
		uintptr(profile),
	)
	res = ffcommon.StringFromPtr(t)
	return
}

/**
* Return a name for the specified profile, if available.
*
* @param codec_id the ID of the codec to which the requested profile belongs
* @param profile the profile value for which a name is requested
* @return A name for the profile if found, NULL otherwise.
*
* @note unlike av_get_profile_name(), which searches a list of profiles
*       supported by a specific decoder or encoder implementation, this
*       function searches the list of profiles from the AVCodecDescriptor
 */
//const char *avcodec_profile_name(enum AVCodecID codec_id, int profile);
func AvcodecProfileName(codec_id AVCodecID, profile ffcommon.FInt) (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_profile_name").Call(
		uintptr(codec_id),
		uintptr(profile),
	)
	res = ffcommon.StringFromPtr(t)
	return
}

// int avcodec_default_execute(AVCodecContext *c, int (*func)(AVCodecContext *c2, void *arg2),void *arg, int *ret, int count, int size);
func (c *AVCodecContext) AvcodecDefaultExecute(func0 func(c2 *AVCodecContext, arg2 ffcommon.FVoidP) uintptr, arg ffcommon.FVoidP, ret *ffcommon.FInt, count, size ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_default_execute").Call(
		uintptr(unsafe.Pointer(c)),
		ffcommon.NewCallback(func0),
		arg,
		uintptr(unsafe.Pointer(ret)),
		uintptr(count),
		uintptr(size),
	)
	res = ffcommon.FInt(t)
	return
}

// int avcodec_default_execute2(AVCodecContext *c, int (*func)(AVCodecContext *c2, void *arg2, int, int),void *arg, int *ret, int count);
func (c *AVCodecContext) AvcodecDefaultExecute2(func0 func(c2 *AVCodecContext, arg2 ffcommon.FVoidP, a, b ffcommon.FInt) uintptr, arg ffcommon.FVoidP, ret *ffcommon.FInt, count ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_default_execute2").Call(
		uintptr(unsafe.Pointer(c)),
		ffcommon.NewCallback(func0),
		arg,
		uintptr(unsafe.Pointer(ret)),
		uintptr(count),
	)
	res = ffcommon.FInt(t)
	return
}

//FIXME func typedef

/**
* Fill AVFrame audio data and linesize pointers.
*
* The buffer buf must be a preallocated buffer with a size big enough
* to contain the specified samples amount. The filled AVFrame data
* pointers will point to this buffer.
*
* AVFrame extended_data channel pointers are allocated if necessary for
* planar audio.
*
* @param frame       the AVFrame
*                    frame->nb_samples must be set prior to calling the
*                    function. This function fills in frame->data,
*                    frame->extended_data, frame->linesize[0].
* @param nb_channels channel count
* @param sample_fmt  sample format
* @param buf         buffer to use for frame data
* @param buf_size    size of buffer
* @param align       plane size sample alignment (0 = default)
* @return            >=0 on success, negative error code on failure
* @todo return the size in bytes required to store the samples in
* case of success, at the next libavutil bump
 */
//int avcodec_fill_audio_frame(AVFrame *frame, int nb_channels,
//enum AVSampleFormat sample_fmt, const uint8_t *buf,
//int buf_size, int align);

func AvcodecFillAudioFrame(frame *AVFrame, nb_channels ffcommon.FInt,
	sample_fmt AVSampleFormat, buf *ffcommon.FUint8T,
	buf_size, align ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_fill_audio_frame").Call(
		uintptr(unsafe.Pointer(frame)),
		uintptr(nb_channels),
		uintptr(sample_fmt),
		uintptr(unsafe.Pointer(buf)),
		uintptr(buf_size),
		uintptr(align),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* Reset the internal codec state / flush internal buffers. Should be called
* e.g. when seeking or when switching to a different stream.
*
* @note for decoders, when refcounted frames are not used
* (i.e. avctx->refcounted_frames is 0), this invalidates the frames previously
* returned from the decoder. When refcounted frames are used, the decoder just
* releases any references it might keep internally, but the caller's reference
* remains valid.
*
* @note for encoders, this function will only do something if the encoder
* declares support for AV_CODEC_CAP_ENCODER_FLUSH. When called, the encoder
* will drain any remaining packets, and can then be re-used for a different
* stream (as opposed to sending a null frame which will leave the encoder
* in a permanent EOF state after draining). This can be desirable if the
* cost of tearing down and replacing the encoder instance is high.
 */
//void avcodec_flush_buffers(AVCodecContext *avctx);
func (avctx *AVCodecContext) AvcodecFlushBuffers() {
	ffcommon.GetAvcodecDll().NewProc("avcodec_flush_buffers").Call(
		uintptr(unsafe.Pointer(avctx)),
	)
}

/**
* Return codec bits per sample.
*
* @param[in] codec_id the codec
* @return Number of bits per sample or zero if unknown for the given codec.
 */
//int av_get_bits_per_sample(enum AVCodecID codec_id);
func AvGetBitsPerSample(codec_id AVCodecID) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_get_bits_per_sample").Call(
		uintptr(codec_id),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* Return the PCM codec associated with a sample format.
* @param be  endianness, 0 for little, 1 for big,
*            -1 (or anything else) for native
* @return  AV_CODEC_ID_PCM_* or AV_CODEC_ID_NONE
 */
//enum AVCodecID av_get_pcm_codec(enum AVSampleFormat fmt, int be);
func AvGetPcmCodec(fmt0 AVSampleFormat, be ffcommon.FInt) (res AVCodecID) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_get_pcm_codec").Call(
		uintptr(fmt0),
		uintptr(be),
	)
	res = AVCodecID(t)
	return
}

/**
* Return codec bits per sample.
* Only return non-zero if the bits per sample is exactly correct, not an
* approximation.
*
* @param[in] codec_id the codec
* @return Number of bits per sample or zero if unknown for the given codec.
 */
//int av_get_exact_bits_per_sample(enum AVCodecID codec_id);
func AvGetExactBitsPerSample(codec_id AVCodecID) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_get_exact_bits_per_sample").Call(
		uintptr(codec_id),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* Return audio frame duration.
*
* @param avctx        codec context
* @param frame_bytes  size of the frame, or 0 if unknown
* @return             frame duration, in samples, if known. 0 if not able to
*                     determine.
 */
//int av_get_audio_frame_duration(AVCodecContext *avctx, int frame_bytes);
func (avctx *AVCodecContext) AvGetAudioFrameDuration(frame_bytes ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_get_audio_frame_duration").Call(
		uintptr(unsafe.Pointer(avctx)),
		uintptr(frame_bytes),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* This function is the same as av_get_audio_frame_duration(), except it works
* with AVCodecParameters instead of an AVCodecContext.
 */
//int av_get_audio_frame_duration2(AVCodecParameters *par, int frame_bytes);
func (par *AVCodecParameters) AvGetAudioFrameDuration2(frame_bytes ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_get_audio_frame_duration2").Call(
		uintptr(unsafe.Pointer(par)),
		uintptr(frame_bytes),
	)
	res = ffcommon.FInt(t)
	return
}

// #if FF_API_OLD_BSF
type AVBitStreamFilterContext struct {
	PrivData ffcommon.FVoidP
	Filter   *AVBitStreamFilter
	Parser   *AVCodecParserContext
	Next     *AVBitStreamFilterContext
	/**
	 * Internal default arguments, used if NULL is passed to av_bitstream_filter_filter().
	 * Not for access by library users.
	 */
	Args ffcommon.FCharPStruct
}

/**
* @deprecated the old bitstream filtering API (using AVBitStreamFilterContext)
* is deprecated. Use the new bitstream filtering API (using AVBSFContext).
 */
//attribute_deprecated
//void av_register_bitstream_filter(AVBitStreamFilter *bsf);
func (bsf *AVBitStreamFilter) AvRegisterBitstreamFilter() {
	ffcommon.GetAvcodecDll().NewProc("av_register_bitstream_filter").Call(
		uintptr(unsafe.Pointer(bsf)),
	)
}

/**
* @deprecated the old bitstream filtering API (using AVBitStreamFilterContext)
* is deprecated. Use av_bsf_get_by_name(), av_bsf_alloc(), and av_bsf_init()
* from the new bitstream filtering API (using AVBSFContext).
 */
//attribute_deprecated
//AVBitStreamFilterContext *av_bitstream_filter_init(const char *name);
func AvBitstreamFilterInit(name ffcommon.FCharP) (res *AVBitStreamFilterContext) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_bitstream_filter_init").Call(
		ffcommon.UintPtrFromString(name),
	)
	res = (*AVBitStreamFilterContext)(unsafe.Pointer(t))
	return
}

/**
* @deprecated the old bitstream filtering API (using AVBitStreamFilterContext)
* is deprecated. Use av_bsf_send_packet() and av_bsf_receive_packet() from the
* new bitstream filtering API (using AVBSFContext).
 */
//attribute_deprecated
//int av_bitstream_filter_filter(AVBitStreamFilterContext *bsfc,
//AVCodecContext *avctx, const char *args,
//uint8_t **poutbuf, int *poutbuf_size,
//const uint8_t *buf, int buf_size, int keyframe);
func (bsfc *AVBitStreamFilterContext) AvBitstreamFilterFilter(avctx *AVCodecContext, args ffcommon.FConstCharP,
	poutbuf **ffcommon.FUint8T, poutbuf_size *ffcommon.FInt,
	buf *ffcommon.FUint8T, buf_size, keyframe ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_bitstream_filter_filter").Call(
		uintptr(unsafe.Pointer(bsfc)),
		uintptr(unsafe.Pointer(avctx)),
		ffcommon.UintPtrFromString(args),
		uintptr(unsafe.Pointer(poutbuf)),
		uintptr(unsafe.Pointer(poutbuf_size)),
		uintptr(unsafe.Pointer(buf)),
		uintptr(buf_size),
		uintptr(keyframe),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* @deprecated the old bitstream filtering API (using AVBitStreamFilterContext)
* is deprecated. Use av_bsf_free() from the new bitstream filtering API (using
* AVBSFContext).
 */
//attribute_deprecated
//void av_bitstream_filter_close(AVBitStreamFilterContext *bsf);
func (bsf *AVBitStreamFilterContext) AvBitstreamFilterClose() {
	ffcommon.GetAvcodecDll().NewProc("av_bitstream_filter_close").Call(
		uintptr(unsafe.Pointer(bsf)),
	)
}

/**
* @deprecated the old bitstream filtering API (using AVBitStreamFilterContext)
* is deprecated. Use av_bsf_iterate() from the new bitstream filtering API (using
* AVBSFContext).
 */
//attribute_deprecated
//const AVBitStreamFilter *av_bitstream_filter_next(const AVBitStreamFilter *f);
func (bsf *AVBitStreamFilterContext) AvBitstreamFilterNext() (res *AVBitStreamFilter) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_bitstream_filter_next").Call(
		uintptr(unsafe.Pointer(bsf)),
	)
	res = (*AVBitStreamFilter)(unsafe.Pointer(t))
	return
}

//#endif

// #if FF_API_NEXT
// attribute_deprecated
// const AVBitStreamFilter *av_bsf_next(void **opaque);
func AvBsfNext(opaque *ffcommon.FVoidP) (res *AVBitStreamFilter) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_bsf_next").Call(
		uintptr(unsafe.Pointer(opaque)),
	)
	res = (*AVBitStreamFilter)(unsafe.Pointer(t))
	return
}

//#endif

/* memory */

/**
* Same behaviour av_fast_malloc but the buffer has additional
* AV_INPUT_BUFFER_PADDING_SIZE at the end which will always be 0.
*
* In addition the whole buffer will initially and after resizes
* be 0-initialized so that no uninitialized data will ever appear.
 */
//void av_fast_padded_malloc(void *ptr, unsigned int *size, size_t min_size);
func AvFastPaddedMalloc(ptr ffcommon.FVoidP, size *ffcommon.FUnsignedInt, min_size ffcommon.FSizeT) {
	ffcommon.GetAvcodecDll().NewProc("av_fast_padded_malloc").Call(
		ptr,
		uintptr(unsafe.Pointer(size)),
		uintptr(min_size),
	)
}

/**
* Same behaviour av_fast_padded_malloc except that buffer will always
* be 0-initialized after call.
 */
//void av_fast_padded_mallocz(void *ptr, unsigned int *size, size_t min_size);
func AvFastPaddedMallocz(ptr ffcommon.FVoidP, size *ffcommon.FUnsignedInt, min_size ffcommon.FSizeT) {
	ffcommon.GetAvcodecDll().NewProc("av_fast_padded_mallocz").Call(
		ptr,
		uintptr(unsafe.Pointer(size)),
		uintptr(min_size),
	)
}

/**
* Encode extradata length to a buffer. Used by xiph codecs.
*
* @param s buffer to write to; must be at least (v/255+1) bytes long
* @param v size of extradata in bytes
* @return number of bytes written to the buffer.
 */
//unsigned int av_xiphlacing(unsigned char *s, unsigned int v);
func AvXiphlacing(s ffcommon.FUnsignedCharP, v ffcommon.FUnsignedInt) (res ffcommon.FUnsignedInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_xiphlacing").Call(
		ffcommon.UintPtrFromString(s),
		uintptr(v),
	)
	res = ffcommon.FUnsignedInt(t)
	return
}

//#if FF_API_USER_VISIBLE_AVHWACCEL
/**
* Register the hardware accelerator hwaccel.
*
* @deprecated  This function doesn't do anything.
 */
//attribute_deprecated
//void av_register_hwaccel(AVHWAccel *hwaccel);
func (hwaccel *AVHWAccel) AvRegisterHwaccel() {
	ffcommon.GetAvcodecDll().NewProc("av_register_hwaccel").Call(
		uintptr(unsafe.Pointer(hwaccel)),
	)
}

/**
* If hwaccel is NULL, returns the first registered hardware accelerator,
* if hwaccel is non-NULL, returns the next registered hardware accelerator
* after hwaccel, or NULL if hwaccel is the last one.
*
* @deprecated  AVHWaccel structures contain no user-serviceable parts, so
*              this function should not be used.
 */
//attribute_deprecated
//AVHWAccel *av_hwaccel_next(const AVHWAccel *hwaccel);
func (hwaccel *AVHWAccel) AvHwaccelNext() (res *AVHWAccel) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_hwaccel_next").Call(
		uintptr(unsafe.Pointer(hwaccel)),
	)
	res = (*AVHWAccel)(unsafe.Pointer(t))
	return
}

//#endif

//#if FF_API_LOCKMGR
/**
* Lock operation used by lockmgr
*
* @deprecated Deprecated together with av_lockmgr_register().
 */
type AVLockOp int32

const (
	AV_LOCK_CREATE  = iota ///< Create a mutex
	AV_LOCK_OBTAIN         ///< Lock the mutex
	AV_LOCK_RELEASE        ///< Unlock the mutex
	AV_LOCK_DESTROY        ///< Free mutex resources
)

/**
* Register a user provided lock manager supporting the operations
* specified by AVLockOp. The "mutex" argument to the function points
* to a (void *) where the lockmgr should store/get a pointer to a user
* allocated mutex. It is NULL upon AV_LOCK_CREATE and equal to the
* value left by the last call for all other ops. If the lock manager is
* unable to perform the op then it should leave the mutex in the same
* state as when it was called and return a non-zero value. However,
* when called with AV_LOCK_DESTROY the mutex will always be assumed to
* have been successfully destroyed. If av_lockmgr_register succeeds
* it will return a non-negative value, if it fails it will return a
* negative value and destroy all mutex and unregister all callbacks.
* av_lockmgr_register is not thread-safe, it must be called from a
* single thread before any calls which make use of locking are used.
*
* @param cb User defined callback. av_lockmgr_register invokes calls
*           to this callback and the previously registered callback.
*           The callback will be used to create more than one mutex
*           each of which must be backed by its own underlying locking
*           mechanism (i.e. do not use a single static object to
*           implement your lock manager). If cb is set to NULL the
*           lockmgr will be unregistered.
*
* @deprecated This function does nothing, and always returns 0. Be sure to
*             build with thread support to get basic thread safety.
 */
//attribute_deprecated
//int av_lockmgr_register(int (*cb)(void **mutex, enum AVLockOp op));
//#endif
func AvLockmgrRegister(cb func(mutex *ffcommon.FVoidP, op AVLockOp) uintptr) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_lockmgr_register").Call(
		ffcommon.NewCallback(cb),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* @return a positive value if s is open (i.e. avcodec_open2() was called on it
* with no corresponding avcodec_close()), 0 otherwise.
 */
//int avcodec_is_open(AVCodecContext *s);
func (s *AVCodecContext) AvcodecIsOpen() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_is_open").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
* Allocate a CPB properties structure and initialize its fields to default
* values.
*
* @param size if non-NULL, the size of the allocated struct will be written
*             here. This is useful for embedding it in side data.
*
* @return the newly allocated struct or NULL on failure
 */
//AVCPBProperties *av_cpb_properties_alloc(size_t *size);
func AvCpbPropertiesAlloc(size *ffcommon.FSizeT) (res *AVCPBProperties) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_cpb_properties_alloc").Call(
		uintptr(unsafe.Pointer(size)),
	)
	res = (*AVCPBProperties)(unsafe.Pointer(t))
	return
}

/**
* @}
 */

//#endif /* AVCODEC_AVCODEC_H */
