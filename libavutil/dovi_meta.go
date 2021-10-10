package libavutil

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

/*
   * DOVI configuration
   * ref: dolby-vision-bitstreams-within-the-iso-base-media-file-format-v2.1.2
          dolby-vision-bitstreams-in-mpeg-2-transport-stream-multiplex-v1.2
   * @code
   * uint8_t  dv_version_major, the major version number that the stream complies with
   * uint8_t  dv_version_minor, the minor version number that the stream complies with
   * uint8_t  dv_profile, the Dolby Vision profile
   * uint8_t  dv_level, the Dolby Vision level
   * uint8_t  rpu_present_flag
   * uint8_t  el_present_flag
   * uint8_t  bl_present_flag
   * uint8_t  dv_bl_signal_compatibility_id
   * @endcode
   *
   * @note The struct must be allocated with av_dovi_alloc() and
   *       its size is not a part of the public ABI.
*/
type AVDOVIDecoderConfigurationRecord struct {
	dv_version_major              ffcommon.FUint8T
	dv_version_minor              ffcommon.FUint8T
	dv_profile                    ffcommon.FUint8T
	dv_level                      ffcommon.FUint8T
	rpu_present_flag              ffcommon.FUint8T
	el_present_flag               ffcommon.FUint8T
	bl_present_flag               ffcommon.FUint8T
	dv_bl_signal_compatibility_id ffcommon.FUint8T
}

/**
 * Allocate a AVDOVIDecoderConfigurationRecord structure and initialize its
 * fields to default values.
 *
 * @return the newly allocated struct or NULL on failure
 */
//AVDOVIDecoderConfigurationRecord *av_dovi_alloc(size_t *size);
//未测试
func AvDoviAlloc(size *ffcommon.FSizeT) (res *AVDOVIDecoderConfigurationRecord, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_dovi_alloc").Call(
		uintptr(unsafe.Pointer(t)),
	)
	if err != nil {
		//return
	}
	res = (*AVDOVIDecoderConfigurationRecord)(unsafe.Pointer(t))
	return
}
