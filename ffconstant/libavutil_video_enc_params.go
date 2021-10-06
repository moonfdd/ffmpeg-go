package ffconstant

type AVVideoEncParamsType int32

const (
	AV_VIDEO_ENC_PARAMS_NONE = -1
	/**
	 * VP9 stores:
	 * - per-frame base (luma AC) quantizer index, exported as AVVideoEncParams.qp
	 * - deltas for luma DC, chroma AC and chroma DC, exported in the
	 *   corresponding entries in AVVideoEncParams.delta_qp
	 * - per-segment delta, exported as for each block as AVVideoBlockParams.delta_qp
	 *
	 * To compute the resulting quantizer index for a block:
	 * - for luma AC, add the base qp and the per-block delta_qp, saturating to
	 *   unsigned 8-bit.
	 * - for luma DC and chroma AC/DC, add the corresponding
	 *   AVVideoBlockParams.delta_qp to the luma AC index, again saturating to
	 *   unsigned 8-bit.
	 */
	AV_VIDEO_ENC_PARAMS_VP9

	/**
	 * H.264 stores:
	 * - in PPS (per-picture):
	 *   * initial QP_Y (luma) value, exported as AVVideoEncParams.qp
	 *   * delta(s) for chroma QP values (same for both, or each separately),
	 *     exported as in the corresponding entries in AVVideoEncParams.delta_qp
	 * - per-slice QP delta, not exported directly, added to the per-MB value
	 * - per-MB delta; not exported directly; the final per-MB quantizer
	 *   parameter - QP_Y - minus the value in AVVideoEncParams.qp is exported
	 *   as AVVideoBlockParams.qp_delta.
	 */
	AV_VIDEO_ENC_PARAMS_H264

	/*
	 * MPEG-2-compatible quantizer.
	 *
	 * Summing the frame-level qp with the per-block delta_qp gives the
	 * resulting quantizer for the block.
	 */
	AV_VIDEO_ENC_PARAMS_MPEG2
)
