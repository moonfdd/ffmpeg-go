package libavcodec

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"ffmpeg-go/libavutil"
	"unsafe"
)

/*
 * AVDVProfile is used to express the differences between various
 * DV flavors. For now it's primarily used for differentiating
 * 525/60 and 625/50, but the plans are to use it for various
 * DV specs as well (e.g. SMPTE314M vs. IEC 61834).
 */
type AVDVProfile struct {
	dsf               ffcommon.FInt            /* value of the dsf in the DV header */
	video_stype       ffcommon.FInt            /* stype for VAUX source pack */
	frame_size        ffcommon.FInt            /* total size of one frame in bytes */
	difseg_size       ffcommon.FInt            /* number of DIF segments per DIF channel */
	n_difchan         ffcommon.FInt            /* number of DIF channels per frame */
	time_base         libavutil.AVRational     /* 1/framerate */
	ltc_divisor       ffcommon.FInt            /* FPS from the LTS standpoint */
	height            ffcommon.FInt            /* picture height in pixels */
	width             ffcommon.FInt            /* picture width in pixels */
	sar               [2]libavutil.AVRational  /* sample aspect ratios for 4:3 and 16:9 */
	pix_fmt           ffconstant.AVPixelFormat /* picture pixel format */
	bpm               ffcommon.FInt            /* blocks per macroblock */
	block_sizes       *ffcommon.FUint8T        /* AC block sizes, in bits */
	audio_stride      ffcommon.FInt            /* size of audio_shuffle table */
	audio_min_samples [3]ffcommon.FInt         /* min amount of audio samples */
	/* for 48kHz, 44.1kHz and 32kHz */
	audio_samples_dist [5]ffcommon.FInt /* how many samples are supposed to be */
	/* in each frame in a 5 frames window */
	//const uint8_t  (*audio_shuffle)[9];     /* PCM shuffling table */
}

/**
* Get a DV profile for the provided compressed frame.
*
* @param sys the profile used for the previous frame, may be NULL
* @param frame the compressed data buffer
* @param buf_size size of the buffer in bytes
* @return the DV profile for the supplied data or NULL on failure
 */
//const AVDVProfile *av_dv_frame_profile(const AVDVProfile *sys,
//const uint8_t *frame, unsigned buf_size);
//未测试
func (sys *AVDVProfile) AvDvFrameProfile(frame *ffcommon.FUint8T, buf_size ffcommon.FUnsigned) (res *AVDVProfile, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_dv_frame_profile").Call(
		uintptr(unsafe.Pointer(sys)),
		uintptr(unsafe.Pointer(frame)),
		uintptr(buf_size),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = (*AVDVProfile)(unsafe.Pointer(t))
	return
}

/**
* Get a DV profile for the provided stream parameters.
 */
//const AVDVProfile *av_dv_codec_profile(int width, int height, enum AVPixelFormat pix_fmt);
//未测试
func (sys *AVDVProfile) AvDvCodecProfile(width ffcommon.FInt, height ffcommon.FInt, pix_fmt ffconstant.AVPixelFormat) (res *AVDVProfile, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_dv_codec_profile").Call(
		uintptr(unsafe.Pointer(sys)),
		uintptr(width),
		uintptr(height),
		uintptr(pix_fmt),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = (*AVDVProfile)(unsafe.Pointer(t))
	return
}

/**
* Get a DV profile for the provided stream parameters.
* The frame rate is used as a best-effort parameter.
 */
//const AVDVProfile *av_dv_codec_profile2(int width, int height, enum AVPixelFormat pix_fmt, AVRational frame_rate);
//未测试
func (sys *AVDVProfile) av_dv_codec_profile2(width ffcommon.FInt, height ffcommon.FInt, pix_fmt ffconstant.AVPixelFormat, frame_rate libavutil.AVRational) (res *AVDVProfile, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_dv_codec_profile2").Call(
		uintptr(unsafe.Pointer(sys)),
		uintptr(width),
		uintptr(height),
		uintptr(pix_fmt),
		uintptr(unsafe.Pointer(&frame_rate)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = (*AVDVProfile)(unsafe.Pointer(t))
	return
}
