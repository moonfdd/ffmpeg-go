package ffconstant

type AVThreadMessageFlags int32

const (
	/**
	 * Perform non-blocking operation.
	 * If this flag is set, send and recv operations are non-blocking and
	 * return AVERROR(EAGAIN) immediately if they can not proceed.
	 */
	AV_THREAD_MESSAGE_NONBLOCK = 1
)
