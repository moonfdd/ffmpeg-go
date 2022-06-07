package libavcodec

import (
	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/ffmpeg-go/libavutil"
	"unsafe"
)

/*
 * Codec parameters public API
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

//#ifndef AVCODEC_CODEC_PAR_H
//#define AVCODEC_CODEC_PAR_H
//
//#include <stdint.h>
//
//#include "../libavutil/avutil.h"
//#include "../libavutil/rational.h"
//#include "../libavutil/pixfmt.h"
//
//#include "codec_id.h"

/**
 * @addtogroup lavc_core
 */
type AVFieldOrder = int32

const (
	AV_FIELD_UNKNOWN = iota
	AV_FIELD_PROGRESSIVE
	AV_FIELD_TT //< Top coded_first, top displayed first
	AV_FIELD_BB //< Bottom coded first, bottom displayed first
	AV_FIELD_TB //< Top coded first, bottom displayed first
	AV_FIELD_BT //< Bottom coded first, top displayed first
)

type AVColorRange = libavutil.AVColorRange
type AVColorPrimaries = libavutil.AVColorPrimaries
type AVColorTransferCharacteristic = libavutil.AVColorTransferCharacteristic
type AVColorSpace = libavutil.AVColorSpace
type AVChromaLocation = libavutil.AVChromaLocation

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
	CodecType AVMediaType
	/**
	 * Specific type of the encoded data (the codec used).
	 */
	CodecId AVCodecID
	/**
	 * Additional information about the codec (corresponds to the AVI FOURCC).
	 */
	CodecTag ffcommon.FUint32T

	/**
	 * Extra binary data needed for initializing the decoder, codec-dependent.
	 *
	 * Must be allocated with av_malloc() and will be freed by
	 * avcodec_parameters_free(). The allocated size of extradata must be at
	 * least extradata_size + AV_INPUT_BUFFER_PADDING_SIZE, with the padding
	 * bytes zeroed.
	 */
	Extradata *ffcommon.FUint8T
	/**
	 * Size of the extradata content in bytes.
	 */
	ExtradataSize ffcommon.FInt

	/**
	 * - video: the pixel format, the value corresponds to enum AVPixelFormat.
	 * - audio: the sample format, the value corresponds to enum AVSampleFormat.
	 */
	Format ffcommon.FInt

	/**
	 * The average bitrate of the encoded data (in bits per second).
	 */
	BitRate ffcommon.FInt64T

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
	BitsPerCodedSample ffcommon.FInt

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
	BitsPerRawSample ffcommon.FInt

	/**
	 * Codec-specific bitstream restrictions that the stream conforms to.
	 */
	Profile ffcommon.FInt
	Level   ffcommon.FInt

	/**
	 * Video only. The dimensions of the video frame in pixels.
	 */
	Width  ffcommon.FInt
	Height ffcommon.FInt

	/**
	 * Video only. The aspect ratio (width / height) which a single pixel
	 * should have when displayed.
	 *
	 * When the aspect ratio is unknown / undefined, the numerator should be
	 * set to 0 (the denominator may have any value).
	 */
	SampleAspectRatio AVRational

	/**
	 * Video only. The order of the fields in interlaced video.
	 */
	FieldOrder AVFieldOrder

	/**
	 * Video only. Additional colorspace characteristics.
	 */
	ColorRange     AVColorRange
	ColorPrimaries AVColorPrimaries
	ColorTrc       AVColorTransferCharacteristic
	ColorSpace     AVColorSpace
	ChromaLocation AVChromaLocation

	/**
	 * Video only. Number of delayed frames.
	 */
	VideoDelay ffcommon.FInt

	/**
	 * Audio only. The channel layout bitmask. May be 0 if the channel layout is
	 * unknown or unspecified, otherwise the number of bits set must be equal to
	 * the channels field.
	 */
	ChannelLayout ffcommon.FUint64T
	/**
	 * Audio only. The number of audio channels.
	 */
	Channels ffcommon.FInt
	/**
	 * Audio only. The number of audio samples per second.
	 */
	SampleRate ffcommon.FInt
	/**
	 * Audio only. The number of bytes per coded audio frame, required by some
	 * formats.
	 *
	 * Corresponds to nBlockAlign in WAVEFORMATEX.
	 */
	BlockAlign ffcommon.FInt
	/**
	 * Audio only. Audio frame size, if known. Required by some formats to be static.
	 */
	FrameSize ffcommon.FInt

	/**
	 * Audio only. The amount of padding (in samples) inserted by the encoder at
	 * the beginning of the audio. I.e. this number of leading decoded samples
	 * must be discarded by the caller to get the original audio without leading
	 * padding.
	 */
	InitialPadding ffcommon.FInt
	/**
	 * Audio only. The amount of padding (in samples) appended by the encoder to
	 * the end of the audio. I.e. this number of decoded samples must be
	 * discarded by the caller from the end of the stream to get the original
	 * audio without any trailing padding.
	 */
	TrailingPadding ffcommon.FInt
	/**
	 * Audio only. Number of samples to skip after a discontinuity.
	 */
	SeekPreroll ffcommon.FInt
}

/**
 * Allocate a new AVCodecParameters and set its fields to default values
 * (unknown/invalid/0). The returned struct must be freed with
 * avcodec_parameters_free().
 */
//AVCodecParameters *avcodec_parameters_alloc(void);
func AvcodecParametersAlloc() (res *AVCodecParameters) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_parameters_alloc").Call()
	if t == 0 {

	}
	res = (*AVCodecParameters)(unsafe.Pointer(t))
	return
}

/**
 * Free an AVCodecParameters instance and everything associated with it and
 * write NULL to the supplied pointer.
 */
//void avcodec_parameters_free(AVCodecParameters **par);
func AvcodecParametersFree(par **AVCodecParameters) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_parameters_free").Call(
		uintptr(unsafe.Pointer(par)),
	)
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
func AvcodecParametersCopy(dst, src *AVCodecParameters) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("avcodec_parameters_copy").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * @}
 */

//#endif // AVCODEC_CODEC_PAR_H
