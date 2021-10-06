package ffconstant

/**
 * @defgroup lavfi_buffersrc Buffer source API
 * @ingroup lavfi
 * @{
 */

const (

	/**
	 * Do not check for format changes.
	 */
	AV_BUFFERSRC_FLAG_NO_CHECK_FORMAT = 1

	/**
	 * Immediately push the frame to the output.
	 */
	AV_BUFFERSRC_FLAG_PUSH = 4

	/**
	 * Keep a reference to the frame.
	 * If the frame if reference-counted, create a new reference; otherwise
	 * copy the frame data.
	 */
	AV_BUFFERSRC_FLAG_KEEP_REF = 8
)
