package libavutil

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"unsafe"
)

/**
 * This structure describes optional metadata relevant to a downmix procedure.
 *
 * All fields are set by the decoder to the value indicated in the audio
 * bitstream (if present), or to a "sane" default otherwise.
 */
type AVDownmixInfo struct {
	/**
	 * Type of downmix preferred by the mastering engineer.
	 */
	preferred_downmix_type ffconstant.AVDownmixType

	/**
	 * Absolute scale factor representing the nominal level of the center
	 * channel during a regular downmix.
	 */
	center_mix_level ffcommon.FDouble

	/**
	 * Absolute scale factor representing the nominal level of the center
	 * channel during an Lt/Rt compatible downmix.
	 */
	center_mix_level_ltrt ffcommon.FDouble

	/**
	 * Absolute scale factor representing the nominal level of the surround
	 * channels during a regular downmix.
	 */
	surround_mix_level ffcommon.FDouble

	/**
	 * Absolute scale factor representing the nominal level of the surround
	 * channels during an Lt/Rt compatible downmix.
	 */
	surround_mix_level_ltrt ffcommon.FDouble

	/**
	 * Absolute scale factor representing the level at which the LFE data is
	 * mixed into L/R channels during downmixing.
	 */
	lfe_mix_level ffcommon.FDouble
}

/**
 * Get a frame's AV_FRAME_DATA_DOWNMIX_INFO side data for editing.
 *
 * If the side data is absent, it is created and added to the frame.
 *
 * @param frame the frame for which the side data is to be obtained or created
 *
 * @return the AVDownmixInfo structure to be edited by the caller, or NULL if
 *         the structure cannot be allocated.
 */
//AVDownmixInfo *av_downmix_info_update_side_data(AVFrame *frame);
//未测试
func (frame *AVFrame) AvDownmixInfoUpdateSideData() (res *AVDownmixInfo, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_downmix_info_update_side_data").Call(
		uintptr(unsafe.Pointer(frame)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = (*AVDownmixInfo)(unsafe.Pointer(t))
	return
}
