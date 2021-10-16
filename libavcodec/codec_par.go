package libavcodec

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"ffmpeg-go/libavutil"
	"unsafe"
)

/**
 * This struct describes the properties of an encoded stream.
 *
 * sizeof(AVCodecParameters) is not a part of the public ABI, this struct must
 * be allocated with avcodec_parameters_alloc() and freed with
 * avcodec_parameters_free().
 */
type AVCodecParameters struct {

	/**
	 * General type of the encoded data.
	 */
	codec_type ffconstant.AVMediaType
	/**
	 * Specific type of the encoded data (the codec used).
	 */
	codec_id ffconstant.AVCodecID
	/**
	 * Additional information about the codec (corresponds to the AVI FOURCC).
	 */
	codec_tag ffcommon.FUint32T

	/**
	 * Extra binary data needed for initializing the decoder, codec-dependent.
	 *
	 * Must be allocated with av_malloc() and will be freed by
	 * avcodec_parameters_free(). The allocated size of extradata must be at
	 * least extradata_size + AV_INPUT_BUFFER_PADDING_SIZE, with the padding
	 * bytes zeroed.
	 */
	extradata *ffcommon.FUint8T
	/**
	 * Size of the extradata content in bytes.
	 */
	extradata_size ffcommon.FInt

	/**
	 * - video: the pixel format, the value corresponds to enum AVPixelFormat.
	 * - audio: the sample format, the value corresponds to enum AVSampleFormat.
	 */
	format ffcommon.FInt

	/**
	 * The average bitrate of the encoded data (in bits per second).
	 */
	bit_rate ffcommon.FUint64T

	/**
	 * The number of bits per sample in the codedwords.
	 *
	 * This is basically the bitrate per sample. It is mandatory for a bunch of
	 * formats to actually decode them. It's the number of bits for one sample in
	 * the actual coded bitstream.
	 *
	 * This could be for example 4 for ADPCM
	 * For PCM formats this matches bits_per_raw_sample
	 * Can be 0
	 */
	bits_per_coded_sample ffcommon.FInt

	/**
	 * This is the number of valid bits in each output sample. If the
	 * sample format has more bits, the least significant bits are additional
	 * padding bits, which are always 0. Use right shifts to reduce the sample
	 * to its actual size. For example, audio formats with 24 bit samples will
	 * have bits_per_raw_sample set to 24, and format set to AV_SAMPLE_FMT_S32.
	 * To get the original sample use "(int32_t)sample >> 8"."
	 *
	 * For ADPCM this might be 12 or 16 or similar
	 * Can be 0
	 */
	bits_per_raw_sample ffcommon.FInt

	/**
	 * Codec-specific bitstream restrictions that the stream conforms to.
	 */
	profile ffcommon.FInt
	level   ffcommon.FInt

	/**
	 * Video only. The dimensions of the video frame in pixels.
	 */
	width  ffcommon.FInt
	height ffcommon.FInt

	/**
	 * Video only. The aspect ratio (width / height) which a single pixel
	 * should have when displayed.
	 *
	 * When the aspect ratio is unknown / undefined, the numerator should be
	 * set to 0 (the denominator may have any value).
	 */
	sample_aspect_ratio libavutil.AVRational

	/**
	 * Video only. The order of the fields in interlaced video.
	 */
	field_order ffconstant.AVFieldOrder

	/**
	 * Video only. Additional colorspace characteristics.
	 */
	color_range     ffconstant.AVColorRange
	color_primaries ffconstant.AVColorPrimaries
	color_trc       ffconstant.AVColorTransferCharacteristic
	color_space     ffconstant.AVColorSpace
	chroma_location ffconstant.AVChromaLocation

	/**
	 * Video only. Number of delayed frames.
	 */
	video_delay ffcommon.FInt

	/**
	 * Audio only. The channel layout bitmask. May be 0 if the channel layout is
	 * unknown or unspecified, otherwise the number of bits set must be equal to
	 * the channels field.
	 */
	channel_layout ffcommon.FUint64T
	/**
	 * Audio only. The number of audio channels.
	 */
	channels ffcommon.FInt
	/**
	 * Audio only. The number of audio samples per second.
	 */
	sample_rate ffcommon.FInt
	/**
	 * Audio only. The number of bytes per coded audio frame, required by some
	 * formats.
	 *
	 * Corresponds to nBlockAlign in WAVEFORMATEX.
	 */
	block_align ffcommon.FInt
	/**
	 * Audio only. Audio frame size, if known. Required by some formats to be static.
	 */
	frame_size ffcommon.FInt

	/**
	 * Audio only. The amount of padding (in samples) inserted by the encoder at
	 * the beginning of the audio. I.e. this number of leading decoded samples
	 * must be discarded by the caller to get the original audio without leading
	 * padding.
	 */
	initial_padding ffcommon.FInt
	/**
	 * Audio only. The amount of padding (in samples) appended by the encoder to
	 * the end of the audio. I.e. this number of decoded samples must be
	 * discarded by the caller from the end of the stream to get the original
	 * audio without any trailing padding.
	 */
	trailing_padding ffcommon.FInt
	/**
	 * Audio only. Number of samples to skip after a discontinuity.
	 */
	seek_preroll ffcommon.FInt
}

/**
 * Allocate a new AVCodecParameters and set its fields to default values
 * (unknown/invalid/0). The returned struct must be freed with
 * avcodec_parameters_free().
 */
//AVCodecParameters *avcodec_parameters_alloc(void);
//未测试
func AvcodecParametersAlloc() (res *AVCodecParameters, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("avcodec_parameters_alloc").Call()
	if err != nil {
		//return
	}
	res = (*AVCodecParameters)(unsafe.Pointer(t))
	return
}

/**
 * Free an AVCodecParameters instance and everything associated with it and
 * write NULL to the supplied pointer.
 */
//void avcodec_parameters_free(AVCodecParameters **par);
//未测试
func AvcodecParametersFree(par **AVCodecParameters) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("avcodec_parameters_free").Call(
		uintptr(unsafe.Pointer(par)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Copy the contents of src to dst. Any allocated fields in dst are freed and
 * replaced with newly allocated duplicates of the corresponding fields in src.
 *
 * @return >= 0 on success, a negative AVERROR code on failure.
 */
//int avcodec_parameters_copy(AVCodecParameters *dst, const AVCodecParameters *src);
//未测试
func AvcodecParametersCopy(dst *AVCodecParameters, src *AVCodecParameters) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("avcodec_parameters_copy").Call(
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
