package libavcodec

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"ffmpeg-go/libavutil"
	"unsafe"
)

type DiracVersionInfo struct {
	major ffcommon.FInt
	minor ffcommon.FInt
}

type AVDiracSeqHeader struct {
	width         ffcommon.FUnsigned
	height        ffcommon.FUnsigned
	chroma_format ffcommon.FUint8T ///< 0: 444  1: 422  2: 420

	interlaced      ffcommon.FUint8T
	top_field_first ffcommon.FUint8T

	frame_rate_index   ffcommon.FUint8T ///< index into dirac_frame_rate[]
	aspect_ratio_index ffcommon.FUint8T ///< index into dirac_aspect_ratio[]

	clean_width        ffcommon.FUint16T
	clean_height       ffcommon.FUint16T
	clean_left_offset  ffcommon.FUint16T
	clean_right_offset ffcommon.FUint16T

	pixel_range_index ffcommon.FUint8T ///< index into dirac_pixel_range_presets[]
	color_spec_index  ffcommon.FUint8T ///< index into dirac_color_spec_presets[]

	profile ffcommon.FInt
	level   ffcommon.FInt

	framerate           libavutil.AVRational
	sample_aspect_ratio libavutil.AVRational

	pix_fmt         ffconstant.AVPixelFormat
	color_range     ffconstant.AVColorRange
	color_primaries ffconstant.AVColorPrimaries
	color_trc       ffconstant.AVColorTransferCharacteristic
	colorspace      ffconstant.AVColorSpace

	version   DiracVersionInfo
	bit_depth ffcommon.FInt
}

/**
 * Parse a Dirac sequence header.
 *
 * @param dsh this function will allocate and fill an AVDiracSeqHeader struct
 *            and write it into this pointer. The caller must free it with
 *            av_free().
 * @param buf the data buffer
 * @param buf_size the size of the data buffer in bytes
 * @param log_ctx if non-NULL, this function will log errors here
 * @return 0 on success, a negative AVERROR code on failure
 */
//int av_dirac_parse_sequence_header(AVDiracSeqHeader **dsh,
//const uint8_t *buf, size_t buf_size,
//void *log_ctx);
//未测试
func AvDiracParseSequenceHeader(dsh **AVDiracSeqHeader,
	buf *ffcommon.FUint8T, buf_size ffcommon.FSizeT,
	log_ctx ffcommon.FVoidP) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_dirac_parse_sequence_header").Call(
		uintptr(unsafe.Pointer(&dsh)),
		uintptr(unsafe.Pointer(buf)),
		uintptr(buf_size),
		uintptr(log_ctx),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}
