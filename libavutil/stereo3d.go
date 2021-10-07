package libavutil

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"unsafe"
)

/**
 * Stereo 3D type: this structure describes how two videos are packed
 * within a single video surface, with additional information as needed.
 *
 * @note The struct must be allocated with av_stereo3d_alloc() and
 *       its size is not a part of the public ABI.
 */
type AVStereo3D struct {
	/**
	 * How views are packed within the video.
	 */
	Type0 ffconstant.AVStereo3DType

	/**
	 * Additional information about the frame packing.
	 */
	Flags ffcommon.FInt

	/**
	 * Determines which views are packed.
	 */
	View ffconstant.AVStereo3DView
}

/**
 * Allocate an AVStereo3D structure and set its fields to default values.
 * The resulting struct can be freed using av_freep().
 *
 * @return An AVStereo3D filled with default values or NULL on failure.
 */
//AVStereo3D *av_stereo3d_alloc(void);
//未测试
func AvStereo3dAlloc() (res *AVStereo3D, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_stereo3d_alloc").Call()
	if err != nil {
		//return
	}
	res = (*AVStereo3D)(unsafe.Pointer(t))
	return
}

//
///**
// * Allocate a complete AVFrameSideData and add it to the frame.
// *
// * @param frame The frame which side data is added to.
// *
// * @return The AVStereo3D structure to be filled by caller.
// */
//AVStereo3D *av_stereo3d_create_side_data(AVFrame *frame);
//
///**
// * Provide a human-readable name of a given stereo3d type.
// *
// * @param type The input stereo3d type value.
// *
// * @return The name of the stereo3d value, or "unknown".
// */
//const char *av_stereo3d_type_name(unsigned int type);
//
///**
// * Get the AVStereo3DType form a human-readable name.
// *
// * @param name The input string.
// *
// * @return The AVStereo3DType value, or -1 if not found.
// */
//int av_stereo3d_from_name(const char *name);
