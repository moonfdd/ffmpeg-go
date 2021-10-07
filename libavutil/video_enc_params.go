package libavutil

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"unsafe"
)

/**
 * Video encoding parameters for a given frame. This struct is allocated along
 * with an optional array of per-block AVVideoBlockParams descriptors.
 * Must be allocated with av_video_enc_params_alloc().
 */
type AVVideoEncParams struct {
	/**
	 * Number of blocks in the array.
	 *
	 * May be 0, in which case no per-block information is present. In this case
	 * the values of blocks_offset / block_size are unspecified and should not
	 * be accessed.
	 */
	NbBlocks ffcommon.FUnsignedInt
	/**
	 * Offset in bytes from the beginning of this structure at which the array
	 * of blocks starts.
	 */
	BlocksOffset ffcommon.FSizeT
	/*
	 * Size of each block in bytes. May not match sizeof(AVVideoBlockParams).
	 */
	BlockSize ffcommon.FSizeT

	/**
	 * Type of the parameters (the codec they are used with).
	 */
	Type ffconstant.AVVideoEncParamsType

	/**
	 * Base quantisation parameter for the frame. The final quantiser for a
	 * given block in a given plane is obtained from this value, possibly
	 * combined with {@code delta_qp} and the per-block delta in a manner
	 * documented for each type.
	 */
	Qp ffcommon.FInt32T

	/**
	 * Quantisation parameter offset from the base (per-frame) qp for a given
	 * plane (first index) and AC/DC coefficients (second index).
	 */
	DeltaQp [4][2]ffcommon.FInt32T
}

/**
 * Data structure for storing block-level encoding information.
 * It is allocated as a part of AVVideoEncParams and should be retrieved with
 * av_video_enc_params_block().
 *
 * sizeof(AVVideoBlockParams) is not a part of the ABI and new fields may be
 * added to it.
 */
type AVVideoBlockParams struct {
	/**
	 * Distance in luma pixels from the top-left corner of the visible frame
	 * to the top-left corner of the block.
	 * Can be negative if top/right padding is present on the coded frame.
	 */
	SrcX, SrcY ffcommon.FInt
	/**
	 * Width and height of the block in luma pixels.
	 */
	W, H ffcommon.FInt

	/**
	 * Difference between this block's final quantization parameter and the
	 * corresponding per-frame value.
	 */
	DeltaQp ffcommon.FInt32T
}

/**
* Allocates memory for AVVideoEncParams of the given type, plus an array of
* {@code nb_blocks} AVVideoBlockParams and initializes the variables. Can be
* freed with a normal av_free() call.
*
* @param out_size if non-NULL, the size in bytes of the resulting data array is
* written here.
 */
//AVVideoEncParams *av_video_enc_params_alloc(enum AVVideoEncParamsType type,
//unsigned int nb_blocks, size_t *out_size);
//未测试
func AvVideoEncParamsAlloc(type0 ffconstant.AVVideoEncParamsType,
	nb_blocks ffcommon.FUnsignedInt, out_size *ffcommon.FSizeT) (res *AVVideoEncParams, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_video_enc_params_alloc").Call(
		uintptr(type0),
		uintptr(nb_blocks),
		uintptr(unsafe.Pointer(out_size)),
	)
	if err != nil {
		//return
	}
	res = (*AVVideoEncParams)(unsafe.Pointer(t))
	return
}

//
///**
// * Allocates memory for AVEncodeInfoFrame plus an array of
// * {@code nb_blocks} AVEncodeInfoBlock in the given AVFrame {@code frame}
// * as AVFrameSideData of type AV_FRAME_DATA_VIDEO_ENC_PARAMS
// * and initializes the variables.
// */
//AVVideoEncParams*
//av_video_enc_params_create_side_data(AVFrame *frame, enum AVVideoEncParamsType type,
//unsigned int nb_blocks);
//未测试
func (frame *AVFrame) av_video_enc_params_create_side_data(type0 ffconstant.AVVideoEncParamsType,
	nb_blocks ffcommon.FUnsignedInt, out_size *ffcommon.FSizeT) (res *AVVideoEncParams, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_video_enc_params_create_side_data").Call(
		uintptr(unsafe.Pointer(frame)),
		uintptr(type0),
		uintptr(nb_blocks),
		uintptr(unsafe.Pointer(out_size)),
	)
	if err != nil {
		//return
	}
	res = (*AVVideoEncParams)(unsafe.Pointer(t))
	return
}
