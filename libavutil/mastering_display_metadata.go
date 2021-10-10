package libavutil

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

/**
 * Mastering display metadata capable of representing the color volume of
 * the display used to master the content (SMPTE 2086:2014).
 *
 * To be used as payload of a AVFrameSideData or AVPacketSideData with the
 * appropriate type.
 *
 * @note The struct should be allocated with av_mastering_display_metadata_alloc()
 *       and its size is not a part of the public ABI.
 */
type AVMasteringDisplayMetadata struct {
	/**
	 * CIE 1931 xy chromaticity coords of color primaries (r, g, b order).
	 */
	DisplayPrimaries [3][2]AVRational

	/**
	 * CIE 1931 xy chromaticity coords of white point.
	 */
	WhitePoint [2]AVRational

	/**
	 * Min luminance of mastering display (cd/m^2).
	 */
	MinLuminance AVRational

	/**
	 * Max luminance of mastering display (cd/m^2).
	 */
	MaxLuminance AVRational

	/**
	 * Flag indicating whether the display primaries (and white point) are set.
	 */
	HasPrimaries ffcommon.FInt

	/**
	 * Flag indicating whether the luminance (min_ and max_) have been set.
	 */
	HasLuminance ffcommon.FInt
}

/**
 * Allocate an AVMasteringDisplayMetadata structure and set its fields to
 * default values. The resulting struct can be freed using av_freep().
 *
 * @return An AVMasteringDisplayMetadata filled with default values or NULL
 *         on failure.
 */
//AVMasteringDisplayMetadata *av_mastering_display_metadata_alloc(void);
//未测试
func AvMasteringDisplayMetadataAlloc() (res *AVMasteringDisplayMetadata, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_mastering_display_metadata_alloc").Call()
	if err != nil {
		//return
	}
	res = (*AVMasteringDisplayMetadata)(unsafe.Pointer(t))
	return
}

/**
 * Allocate a complete AVMasteringDisplayMetadata and add it to the frame.
 *
 * @param frame The frame which side data is added to.
 *
 * @return The AVMasteringDisplayMetadata structure to be filled by caller.
 */
//AVMasteringDisplayMetadata *av_mastering_display_metadata_create_side_data(AVFrame *frame);
//未测试
func (frame *AVFrame) AvMasteringDisplayMetadataCreateSideData() (res *AVMasteringDisplayMetadata, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_mastering_display_metadata_create_side_data").Call(
		uintptr(unsafe.Pointer(frame)),
	)
	if err != nil {
		//return
	}
	res = (*AVMasteringDisplayMetadata)(unsafe.Pointer(t))
	return
}

/**
 * Content light level needed by to transmit HDR over HDMI (CTA-861.3).
 *
 * To be used as payload of a AVFrameSideData or AVPacketSideData with the
 * appropriate type.
 *
 * @note The struct should be allocated with av_content_light_metadata_alloc()
 *       and its size is not a part of the public ABI.
 */
type AVContentLightMetadata struct {

	/**
	 * Max content light level (cd/m^2).
	 */
	MaxCLL ffcommon.FUnsigned

	/**
	 * Max average light level per frame (cd/m^2).
	 */
	MaxFALL ffcommon.FUnsigned
}

/**
 * Allocate an AVContentLightMetadata structure and set its fields to
 * default values. The resulting struct can be freed using av_freep().
 *
 * @return An AVContentLightMetadata filled with default values or NULL
 *         on failure.
 */
//AVContentLightMetadata *av_content_light_metadata_alloc(size_t *size);
//未测试
func AvContentLightMetadataAlloc() (res *AVContentLightMetadata, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_content_light_metadata_alloc").Call()
	if err != nil {
		//return
	}
	res = (*AVContentLightMetadata)(unsafe.Pointer(t))
	return
}

/**
 * Allocate a complete AVContentLightMetadata and add it to the frame.
 *
 * @param frame The frame which side data is added to.
 *
 * @return The AVContentLightMetadata structure to be filled by caller.
 */
//AVContentLightMetadata *av_content_light_metadata_create_side_data(AVFrame *frame);
//未测试
func (frame *AVFrame) AvContentLightMetadataCreateSideData() (res *AVMasteringDisplayMetadata, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_content_light_metadata_create_side_data").Call(
		uintptr(unsafe.Pointer(frame)),
	)
	if err != nil {
		//return
	}
	res = (*AVMasteringDisplayMetadata)(unsafe.Pointer(t))
	return
}
