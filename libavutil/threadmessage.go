package libavutil

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

type AVThreadMessageQueue struct {
}

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
//未测试
func AvThreadMessageQueueAlloc(mq **AVThreadMessageQueue,
	nelem ffcommon.FUnsigned,
	elsize ffcommon.FUnsigned) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_thread_message_queue_alloc").Call(
		uintptr(unsafe.Pointer(&mq)),
		uintptr(nelem),
		uintptr(elsize),
	)
	if err != nil {
		return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Free a message queue.
 *
 * The message queue must no longer be in use by another thread.
 */
//void av_thread_message_queue_free(AVThreadMessageQueue **mq);
//未测试
func AvThreadMessageQueueFree(mq **AVThreadMessageQueue) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_thread_message_queue_free").Call(
		uintptr(unsafe.Pointer(&mq)),
	)
	if err != nil {
		return
	}
	if t == 0 {

	}
	return
}

/**
 * Send a message on the queue.
 */
//int av_thread_message_queue_send(AVThreadMessageQueue *mq,
//void *msg,
//unsigned flags);
//未测试
func (mq *AVThreadMessageQueue) AvThreadMessageQueueSend(
	msg ffcommon.FVoidP,
	flags ffcommon.FUnsigned) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_thread_message_queue_alloc").Call(
		uintptr(unsafe.Pointer(mq)),
		uintptr(unsafe.Pointer(msg)),
		uintptr(flags),
	)
	if err != nil {
		return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Receive a message from the queue.
 */
//int av_thread_message_queue_recv(AVThreadMessageQueue *mq,
//void *msg,
//unsigned flags);
//未测试
func (mq *AVThreadMessageQueue) AvThreadMessageQueueRecv(
	msg ffcommon.FVoidP,
	flags ffcommon.FUnsigned) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_thread_message_queue_recv").Call(
		uintptr(unsafe.Pointer(mq)),
		uintptr(unsafe.Pointer(msg)),
		uintptr(flags),
	)
	if err != nil {
		return
	}
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
//未测试
func (mq *AVThreadMessageQueue) AvThreadMessageQueueSetErrSend(
	err0 ffcommon.FInt) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_thread_message_queue_set_err_send").Call(
		uintptr(unsafe.Pointer(mq)),
		uintptr(err0),
	)
	if err != nil {
		return
	}
	if t == 0 {

	}
	return
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
//未测试
func (mq *AVThreadMessageQueue) AvThreadMessageQueueSetErrRecv(
	err0 ffcommon.FInt) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_thread_message_queue_set_err_recv").Call(
		uintptr(unsafe.Pointer(mq)),
		uintptr(err0),
	)
	if err != nil {
		return
	}
	if t == 0 {

	}
	return
}

/**
 * Set the optional free message callback function which will be called if an
 * operation is removing messages from the queue.
 */
//void av_thread_message_queue_set_free_func(AVThreadMessageQueue *mq,
//void (*free_func)(void *msg));
//未测试
func (mq *AVThreadMessageQueue) AvThreadMessageQueueSetFreeFunc(
	free_func func(msg ffcommon.FVoidP)) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_thread_message_queue_set_free_func").Call(
		uintptr(unsafe.Pointer(mq)),
		uintptr(unsafe.Pointer(&free_func)),
	)
	if err != nil {
		return
	}
	if t == 0 {

	}
	return
}

/**
 * Return the current number of messages in the queue.
 *
 * @return the current number of messages or AVERROR(ENOSYS) if lavu was built
 *         without thread support
 */
//int av_thread_message_queue_nb_elems(AVThreadMessageQueue *mq);
//未测试
func (mq *AVThreadMessageQueue) AvThreadMessageQueueNbElems() (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_thread_message_queue_nb_elems").Call(
		uintptr(unsafe.Pointer(mq)),
	)
	if err != nil {
		return
	}
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
//未测试
func (mq *AVThreadMessageQueue) AvThreadMessageFlush() (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_thread_message_flush").Call(
		uintptr(unsafe.Pointer(mq)),
	)
	if err != nil {
		return
	}
	if t == 0 {

	}
	return
}
