package libavcodec

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

type AVVorbisParseContext struct {
}

/**
 * Allocate and initialize the Vorbis parser using headers in the extradata.
 */
//AVVorbisParseContext *av_vorbis_parse_init(const uint8_t *extradata,
//int extradata_size);
//未测试
func AvVorbisParseInit(extradata *ffcommon.FUint8T, extradata_size ffcommon.FInt) (res *AVVorbisParseContext, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_vorbis_parse_init").Call(
		uintptr(unsafe.Pointer(extradata)),
		uintptr(extradata_size),
	)
	if err != nil {
		//return
	}
	res = (*AVVorbisParseContext)(unsafe.Pointer(t))
	return
}

/**
 * Free the parser and everything associated with it.
 */
//void av_vorbis_parse_free(AVVorbisParseContext **s);
//未测试
func AvVorbisParseFree(s **AVVorbisParseContext) (err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_vorbis_parse_free").Call(
		uintptr(unsafe.Pointer(&s)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Get the duration for a Vorbis packet.
 *
 * If @p flags is @c NULL,
 * special frames are considered invalid.
 *
 * @param s        Vorbis parser context
 * @param buf      buffer containing a Vorbis frame
 * @param buf_size size of the buffer
 * @param flags    flags for special frames
 */
//int av_vorbis_parse_frame_flags(AVVorbisParseContext *s, const uint8_t *buf,
//int buf_size, int *flags);
//未测试
func (s *AVVorbisParseContext) AvVorbisParseFrameFlags(buf *ffcommon.FUint8T,
	buf_size ffcommon.FInt, flags *ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_vorbis_parse_frame_flags").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(buf)),
		uintptr(buf_size),
		uintptr(unsafe.Pointer(flags)),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Get the duration for a Vorbis packet.
 *
 * @param s        Vorbis parser context
 * @param buf      buffer containing a Vorbis frame
 * @param buf_size size of the buffer
 */
//int av_vorbis_parse_frame(AVVorbisParseContext *s, const uint8_t *buf,
//int buf_size);
//未测试
func (s *AVVorbisParseContext) AvVorbisParseFrame(buf *ffcommon.FUint8T,
	buf_size ffcommon.FInt) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_vorbis_parse_frame").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(buf)),
		uintptr(buf_size),
	)
	if err != nil {
		//return
	}
	res = ffcommon.FInt(t)
	return
}

//void av_vorbis_parse_reset(AVVorbisParseContext *s);
//未测试
func AvVorbisParseReset(s **AVVorbisParseContext) (err error) {
	var t uintptr
	t, _, err = ffcommon.GetAvutilDll().NewProc("av_vorbis_parse_reset").Call(
		uintptr(unsafe.Pointer(&s)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}
