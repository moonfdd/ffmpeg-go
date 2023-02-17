package libavcodec

import (
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
)

/*
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

//#ifndef AVCODEC_DV_PROFILE_H
//#define AVCODEC_DV_PROFILE_H
//
//#include <stdint.h>
//
//#include "../libavutil/pixfmt.h"
//#include "../libavutil/rational.h"
//#include "avcodec.h"

/* minimum number of bytes to read from a DV stream in order to
 * determine the profile */
const DV_PROFILE_BYTES = (6 * 80) /* 6 DIF blocks */

/*
 * AVDVProfile is used to express the differences between various
 * DV flavors. For now it's primarily used for differentiating
 * 525/60 and 625/50, but the plans are to use it for various
 * DV specs as well (e.g. SMPTE314M vs. IEC 61834).
 */
type AVDVProfile struct {
	Dsf             ffcommon.FInt     /* value of the dsf in the DV header */
	VideoStype      ffcommon.FInt     /* stype for VAUX source pack */
	FrameSize       ffcommon.FInt     /* total size of one frame in bytes */
	DifsegSize      ffcommon.FInt     /* number of DIF segments per DIF channel */
	NDifchan        ffcommon.FInt     /* number of DIF channels per frame */
	TimeBase        AVRational        /* 1/framerate */
	LtcDivisor      ffcommon.FInt     /* FPS from the LTS standpoint */
	Height          ffcommon.FInt     /* picture height in pixels */
	Width           ffcommon.FInt     /* picture width in pixels */
	Sar             [2]AVRational     /* sample aspect ratios for 4:3 and 16:9 */
	PixFmt          AVPixelFormat     /* picture pixel format */
	Bpm             ffcommon.FInt     /* blocks per macroblock */
	BlockSizes      *ffcommon.FUint8T /* AC block sizes, in bits */
	AudioStride     ffcommon.FInt     /* size of audio_shuffle table */
	AudioMinSamples [3]ffcommon.FInt  /* min amount of audio samples */
	/* for 48kHz, 44.1kHz and 32kHz */
	AudioSamplesDist [5]ffcommon.FInt /* how many samples are supposed to be */
	/* in each frame in a 5 frames window */
	//const uint8_t  (*audio_shuffle)[9];     /* PCM shuffling table */
	AdioShuffle *[9]ffcommon.FUint8T
}

/**
 * Get a DV profile for the provided compressed frame.
 *
 * @param sys the profile used for the previous frame, may be NULL
 * @param frame the compressed data buffer
 * @param buf_size size of the buffer in bytes
 * @return the DV profile for the supplied data or NULL on failure
 */
//const AVDVProfile *av_dv_frame_profile(const AVDVProfile *sys,
//const uint8_t *frame, unsigned buf_size);
func (sys *AVDVProfile) AvDvFrameProfile(frame *ffcommon.FUint8T, buf_size ffcommon.FUnsigned) (res *AVDVProfile) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_dv_frame_profile").Call(
		uintptr(unsafe.Pointer(sys)),
		uintptr(unsafe.Pointer(frame)),
		uintptr(buf_size),
	)
	res = (*AVDVProfile)(unsafe.Pointer(t))
	return
}

/**
 * Get a DV profile for the provided stream parameters.
 */
//const AVDVProfile *av_dv_codec_profile(int width, int height, enum AVPixelFormat pix_fmt);
func AvDvCodecProfile(width, height ffcommon.FInt, pix_fmt AVPixelFormat) (res *AVDVProfile) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_dv_codec_profile").Call(
		uintptr(width),
		uintptr(height),
		uintptr(pix_fmt),
	)
	res = (*AVDVProfile)(unsafe.Pointer(t))
	return
}

/**
 * Get a DV profile for the provided stream parameters.
 * The frame rate is used as a best-effort parameter.
 */
//const AVDVProfile *av_dv_codec_profile2(int width, int height, enum AVPixelFormat pix_fmt, AVRational frame_rate);
func AvDvCodecProfile2(width, height ffcommon.FInt, pix_fmt AVPixelFormat, frame_rate AVRational) (res *AVDVProfile) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_dv_codec_profile2").Call(
		uintptr(width),
		uintptr(height),
		uintptr(pix_fmt),
		uintptr(unsafe.Pointer(&frame_rate)),
	)
	res = (*AVDVProfile)(unsafe.Pointer(t))
	return
}

//#endif /* AVCODEC_DV_PROFILE_H */
