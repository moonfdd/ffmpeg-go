package libavutil

import (
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
)

/*
 * This file is part of FFmpeg.
 *
 * FFmpeg is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public License
 * as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * FFmpeg is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with FFmpeg; if not, write to the Free Software Foundation, Inc.,
 * 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA
 */

//#ifndef AVUTIL_THREADMESSAGE_H
//#define AVUTIL_THREADMESSAGE_H

//typedef struct AVThreadMessageQueue AVThreadMessageQueue;
type AVThreadMessageQueue struct {
}
type AVThreadMessageFlags int32

const (

	/**
	 * Perform non-blocking operation.
	 * If this flag is set, send and recv operations are non-blocking and
	 * return AVERROR(EAGAIN) immediately if they can not proceed.
	 */
	AV_THREAD_MESSAGE_NONBLOCK = 1
)

/**
 * Allocate a new message queue.
 *
 * @param mq      pointer to the message queue
 * @param nelem   maximum number of elements in the queue
 * @param elsize  size of each element in the queue
 * @return  >=0 for success; <0 for error, in particular AVERROR(ENOSYS) if
 *          lavu was built without thread support
 */
//int av_thread_message_queue_alloc(AVThreadMessageQueue **mq,
//unsigned nelem,
//unsigned elsize);
func AvThreadMessageQueueAlloc(mq **AVThreadMessageQueue, nelem, elsize ffcommon.FUnsigned) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_thread_message_queue_alloc").Call(
		uintptr(unsafe.Pointer(mq)),
		uintptr(nelem),
		uintptr(elsize),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Free a message queue.
 *
 * The message queue must no longer be in use by another thread.
 */
//void av_thread_message_queue_free(AVThreadMessageQueue **mq);
func AvThreadMessageQueueFree(mq **AVThreadMessageQueue) {
	ffcommon.GetAvutilDll().NewProc("av_thread_message_queue_free").Call(
		uintptr(unsafe.Pointer(mq)),
	)
}

/**
 * Send a message on the queue.
 */
//int av_thread_message_queue_send(AVThreadMessageQueue *mq,
//void *msg,
//unsigned flags);
func (mq *AVThreadMessageQueue) AvThreadMessageQueueSend(msg ffcommon.FVoidP, flags ffcommon.FUnsigned) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_thread_message_queue_send").Call(
		uintptr(unsafe.Pointer(mq)),
		msg,
		uintptr(flags),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Receive a message from the queue.
 */
//int av_thread_message_queue_recv(AVThreadMessageQueue *mq,
//void *msg,
//unsigned flags);
func (mq *AVThreadMessageQueue) AvThreadMessageQueueRecv(msg ffcommon.FVoidP, flags ffcommon.FUnsigned) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_thread_message_queue_recv").Call(
		uintptr(unsafe.Pointer(mq)),
		msg,
		uintptr(flags),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Set the sending error code.
 *
 * If the error code is set to non-zero, av_thread_message_queue_send() will
 * return it immediately. Conventional values, such as AVERROR_EOF or
 * AVERROR(EAGAIN), can be used to cause the sending thread to stop or
 * suspend its operation.
 */
//void av_thread_message_queue_set_err_send(AVThreadMessageQueue *mq,
//int err);
func (mq *AVThreadMessageQueue) AvThreadMessageQueueSetErrSend(err ffcommon.FInt) {
	ffcommon.GetAvutilDll().NewProc("av_thread_message_queue_set_err_send").Call(
		uintptr(unsafe.Pointer(mq)),
		uintptr(err),
	)
}

/**
 * Set the receiving error code.
 *
 * If the error code is set to non-zero, av_thread_message_queue_recv() will
 * return it immediately when there are no longer available messages.
 * Conventional values, such as AVERROR_EOF or AVERROR(EAGAIN), can be used
 * to cause the receiving thread to stop or suspend its operation.
 */
//void av_thread_message_queue_set_err_recv(AVThreadMessageQueue *mq,
//int err);
func (mq *AVThreadMessageQueue) AvThreadMessageQueueSetErrRecv(err ffcommon.FInt) {
	ffcommon.GetAvutilDll().NewProc("av_thread_message_queue_set_err_recv").Call(
		uintptr(unsafe.Pointer(mq)),
		uintptr(err),
	)
}

/**
 * Set the optional free message callback function which will be called if an
 * operation is removing messages from the queue.
 */
//void av_thread_message_queue_set_free_func(AVThreadMessageQueue *mq,
//void (*free_func)(void *msg));
func (mq *AVThreadMessageQueue) AvThreadMessageQueueSetFreeFunc(free_func func(msg ffcommon.FVoidP) uintptr) {
	ffcommon.GetAvutilDll().NewProc("av_thread_message_queue_set_free_func").Call(
		uintptr(unsafe.Pointer(mq)),
		ffcommon.NewCallback(free_func),
	)
}

/**
 * Return the current number of messages in the queue.
 *
 * @return the current number of messages or AVERROR(ENOSYS) if lavu was built
 *         without thread support
 */
//int av_thread_message_queue_nb_elems(AVThreadMessageQueue *mq);
func (mq *AVThreadMessageQueue) AvThreadMessageQueueNbElems() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_thread_message_queue_nb_elems").Call(
		uintptr(unsafe.Pointer(mq)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Flush the message queue
 *
 * This function is mostly equivalent to reading and free-ing every message
 * except that it will be done in a single operation (no lock/unlock between
 * reads).
 */
//void av_thread_message_flush(AVThreadMessageQueue *mq);
func (mq *AVThreadMessageQueue) AvThreadMessageFlush() {
	ffcommon.GetAvutilDll().NewProc("av_thread_message_flush").Call(
		uintptr(unsafe.Pointer(mq)),
	)
}

//#endif /* AVUTIL_THREADMESSAGE_H */
