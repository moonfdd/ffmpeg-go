package ffconstant

/** @name Error flags returned by av_lzo1x_decode
 * @{ */
/// end of the input buffer reached before decoding finished
const AV_LZO_INPUT_DEPLETED = 1

/// decoded data did not fit into output buffer
const AV_LZO_OUTPUT_FULL = 2

/// a reference to previously decoded data was wrong
const AV_LZO_INVALID_BACKPTR = 4

/// a non-specific error in the compressed bitstream
const AV_LZO_ERROR = 8

/** @} */

const AV_LZO_INPUT_PADDING = 8
const AV_LZO_OUTPUT_PADDING = 12
