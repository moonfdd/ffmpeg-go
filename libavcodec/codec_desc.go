package libavcodec

import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/ffconstant"
	"syscall"
	"unsafe"
)

/**
 * @addtogroup lavc_core
 * @{
 */

/**
 * This struct describes the properties of a single codec described by an
 * AVCodecID.
 * @see avcodec_descriptor_get()
 */
type AVCodecDescriptor struct {
	id    ffconstant.AVCodecID
	type0 ffconstant.AVMediaType
	/**
	 * Name of the codec described by this descriptor. It is non-empty and
	 * unique for each codec descriptor. It should contain alphanumeric
	 * characters and '_' only.
	 */
	name ffcommon.FBuf
	/**
	 * A more descriptive name for this codec. May be NULL.
	 */
	long_name ffcommon.FBuf
	/**
	 * Codec properties, a combination of AV_CODEC_PROP_* flags.
	 */
	props ffcommon.FInt
	/**
	 * MIME type(s) associated with the codec.
	 * May be NULL; if not, a NULL-terminated array of MIME types.
	 * The first item is always non-NULL and is the preferred MIME type.
	 */
	mime_types *ffcommon.FBuf
	/**
	 * If non-NULL, an array of profiles recognized for this codec.
	 * Terminated with FF_PROFILE_UNKNOWN.
	 */
	profiles *AVProfile
}

/**
 * @return descriptor for given codec ID or NULL if no descriptor exists.
 */
//const AVCodecDescriptor *avcodec_descriptor_get(enum AVCodecID id);
//未测试
func AvcodecDescriptorGet(codec_id ffconstant.AVCodecID) (res *AVCodecDescriptor, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("avcodec_descriptor_get").Call(
		uintptr(codec_id),
	)
	if err != nil {
		//return
	}
	res = (*AVCodecDescriptor)(unsafe.Pointer(t))
	return
}

/**
 * Iterate over all codec descriptors known to libavcodec.
 *
 * @param prev previous descriptor. NULL to get the first descriptor.
 *
 * @return next descriptor or NULL after the last descriptor
 */
//const AVCodecDescriptor *avcodec_descriptor_next(const AVCodecDescriptor *prev);
//未测试
func (prev *AVCodecDescriptor) avcodec_descriptor_next() (res *AVCodecDescriptor, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("avcodec_descriptor_next").Call(
		uintptr(unsafe.Pointer(prev)),
	)
	if err != nil {
		//return
	}
	res = (*AVCodecDescriptor)(unsafe.Pointer(t))
	return
}

/**
 * @return codec descriptor with the given name or NULL if no such descriptor
 *         exists.
 */
//const AVCodecDescriptor *avcodec_descriptor_get_by_name(const char *name);
//未测试
func AvcodecDescriptorGetByName(name ffcommon.FConstCharP) (res *AVCodecDescriptor, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("avcodec_descriptor_next").Call(
		uintptr(unsafe.Pointer(namep)),
	)
	if err != nil {
		//return
	}
	res = (*AVCodecDescriptor)(unsafe.Pointer(t))
	return
}
