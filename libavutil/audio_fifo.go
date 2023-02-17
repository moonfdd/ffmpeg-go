package libavutil

import (
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
)

/*
 * Audio FIFO
 * Copyright (c) 2012 Justin Ruggles <justin.ruggles@gmail.com>
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

/**
 * @file
 * Audio FIFO Buffer
 */

//#ifndef AVUTIL_AUDIO_FIFO_H
//#define AVUTIL_AUDIO_FIFO_H
//
//#include "avutil.h"
//#include "fifo.h"
//#include "samplefmt.h"

/**
 * @addtogroup lavu_audio
 * @{
 *
 * @defgroup lavu_audiofifo Audio FIFO Buffer
 * @{
 */

/**
 * Context for an Audio FIFO Buffer.
 *
 * - Operates at the sample level rather than the byte level.
 * - Supports multiple channels with either planar or packed sample format.
 * - Automatic reallocation when writing to a full buffer.
 */
//typedef struct AVAudioFifo AVAudioFifo;
type AVAudioFifo struct {
}

/**
 * Free an AVAudioFifo.
 *
 * @param af  AVAudioFifo to free
 */
//void av_audio_fifo_free(AVAudioFifo *af);
func (af *AVAudioFifo) AvAudioFifoFree() {
	ffcommon.GetAvutilDll().NewProc("av_audio_fifo_free").Call(
		uintptr(unsafe.Pointer(af)),
	)
}

/**
 * Allocate an AVAudioFifo.
 *
 * @param sample_fmt  sample format
 * @param channels    number of channels
 * @param nb_samples  initial allocation size, in samples
 * @return            newly allocated AVAudioFifo, or NULL on error
 */
//AVAudioFifo *av_audio_fifo_alloc(enum AVSampleFormat sample_fmt, int channels,
//int nb_samples);
func AvAudioFifoAlloc(sample_fmt AVSampleFormat, channels, nb_samples ffcommon.FInt) (res *AVAudioFifo) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_audio_fifo_alloc").Call(
		uintptr(sample_fmt),
		uintptr(channels),
		uintptr(nb_samples),
	)
	res = (*AVAudioFifo)(unsafe.Pointer(t))
	return
}

/**
 * Reallocate an AVAudioFifo.
 *
 * @param af          AVAudioFifo to reallocate
 * @param nb_samples  new allocation size, in samples
 * @return            0 if OK, or negative AVERROR code on failure
 */
//av_warn_unused_result
//int av_audio_fifo_realloc(AVAudioFifo *af, int nb_samples);
func (af *AVAudioFifo) AvAudioFifoRealloc(nb_samples ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_audio_fifo_realloc").Call(
		uintptr(unsafe.Pointer(af)),
		uintptr(nb_samples),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Write data to an AVAudioFifo.
 *
 * The AVAudioFifo will be reallocated automatically if the available space
 * is less than nb_samples.
 *
 * @see enum AVSampleFormat
 * The documentation for AVSampleFormat describes the data layout.
 *
 * @param af          AVAudioFifo to write to
 * @param data        audio data plane pointers
 * @param nb_samples  number of samples to write
 * @return            number of samples actually written, or negative AVERROR
 *                    code on failure. If successful, the number of samples
 *                    actually written will always be nb_samples.
 */
//int av_audio_fifo_write(AVAudioFifo *af, void **data, int nb_samples);
func (af *AVAudioFifo) AvAudioFifoWrite(data *ffcommon.FVoidP, nb_samples ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_audio_fifo_write").Call(
		uintptr(unsafe.Pointer(af)),
		uintptr(unsafe.Pointer(data)),
		uintptr(nb_samples),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Peek data from an AVAudioFifo.
 *
 * @see enum AVSampleFormat
 * The documentation for AVSampleFormat describes the data layout.
 *
 * @param af          AVAudioFifo to read from
 * @param data        audio data plane pointers
 * @param nb_samples  number of samples to peek
 * @return            number of samples actually peek, or negative AVERROR code
 *                    on failure. The number of samples actually peek will not
 *                    be greater than nb_samples, and will only be less than
 *                    nb_samples if av_audio_fifo_size is less than nb_samples.
 */
//int av_audio_fifo_peek(AVAudioFifo *af, void **data, int nb_samples);
func (af *AVAudioFifo) AvAudioFifoPeek(data *ffcommon.FVoidP, nb_samples ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_audio_fifo_peek").Call(
		uintptr(unsafe.Pointer(af)),
		uintptr(unsafe.Pointer(data)),
		uintptr(nb_samples),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Peek data from an AVAudioFifo.
 *
 * @see enum AVSampleFormat
 * The documentation for AVSampleFormat describes the data layout.
 *
 * @param af          AVAudioFifo to read from
 * @param data        audio data plane pointers
 * @param nb_samples  number of samples to peek
 * @param offset      offset from current read position
 * @return            number of samples actually peek, or negative AVERROR code
 *                    on failure. The number of samples actually peek will not
 *                    be greater than nb_samples, and will only be less than
 *                    nb_samples if av_audio_fifo_size is less than nb_samples.
 */
//int av_audio_fifo_peek_at(AVAudioFifo *af, void **data, int nb_samples, int offset);
func (af *AVAudioFifo) AvAudioFifoPeekAt(data *ffcommon.FVoidP, nb_samples, offset ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_audio_fifo_peek_at").Call(
		uintptr(unsafe.Pointer(af)),
		uintptr(unsafe.Pointer(data)),
		uintptr(nb_samples),
		uintptr(offset),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Read data from an AVAudioFifo.
 *
 * @see enum AVSampleFormat
 * The documentation for AVSampleFormat describes the data layout.
 *
 * @param af          AVAudioFifo to read from
 * @param data        audio data plane pointers
 * @param nb_samples  number of samples to read
 * @return            number of samples actually read, or negative AVERROR code
 *                    on failure. The number of samples actually read will not
 *                    be greater than nb_samples, and will only be less than
 *                    nb_samples if av_audio_fifo_size is less than nb_samples.
 */
//int av_audio_fifo_read(AVAudioFifo *af, void **data, int nb_samples);
func (af *AVAudioFifo) AvAudioFifoRead(data *ffcommon.FVoidP, nb_samples ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_audio_fifo_read").Call(
		uintptr(unsafe.Pointer(af)),
		uintptr(unsafe.Pointer(data)),
		uintptr(nb_samples),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Drain data from an AVAudioFifo.
 *
 * Removes the data without reading it.
 *
 * @param af          AVAudioFifo to drain
 * @param nb_samples  number of samples to drain
 * @return            0 if OK, or negative AVERROR code on failure
 */
//int av_audio_fifo_drain(AVAudioFifo *af, int nb_samples);
func (af *AVAudioFifo) AvAudioFifoDrain(nb_samples ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_audio_fifo_drain").Call(
		uintptr(unsafe.Pointer(af)),
		uintptr(nb_samples),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Reset the AVAudioFifo buffer.
 *
 * This empties all data in the buffer.
 *
 * @param af  AVAudioFifo to reset
 */
//void av_audio_fifo_reset(AVAudioFifo *af);
func (af *AVAudioFifo) AvAudioFifoReset() {
	ffcommon.GetAvutilDll().NewProc("av_audio_fifo_reset").Call(
		uintptr(unsafe.Pointer(af)),
	)
}

/**
 * Get the current number of samples in the AVAudioFifo available for reading.
 *
 * @param af  the AVAudioFifo to query
 * @return    number of samples available for reading
 */
//int av_audio_fifo_size(AVAudioFifo *af);
func (af *AVAudioFifo) AvAudioFifoSize() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_audio_fifo_size").Call(
		uintptr(unsafe.Pointer(af)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Get the current number of samples in the AVAudioFifo available for writing.
 *
 * @param af  the AVAudioFifo to query
 * @return    number of samples available for writing
 */
//int av_audio_fifo_space(AVAudioFifo *af);
func (af *AVAudioFifo) AvAudioFifoSpace() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_audio_fifo_space").Call(
		uintptr(unsafe.Pointer(af)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * @}
 * @}
 */

//#endif /* AVUTIL_AUDIO_FIFO_H */
