package libavdevice

import (
	"ffmpeg-go/ffcommon"
	"unsafe"
)

/**
 * Return the LIBAVDEVICE_VERSION_INT ffconstant.
 */
//unsigned avdevice_version(void);
func AvdeviceVersion() (res ffcommon.FUint, err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("avdevice_version").Call()
	res = ffcommon.FUint(t)
	return
}
/**
 * Return the libavdevice build-time configuration.
 */
//const char *avdevice_configuration(void);
func AvdeviceConfiguration() (res ffcommon.FConstCharP, err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("avdevice_configuration").Call()
	res = ffcommon.GoAStr(t)
	return
}
/**
 * Return the libavdevice license.
 */
//const char *avdevice_license(void);
func AvdeviceLicense() (res ffcommon.FConstCharP, err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("avdevice_license").Call()
	res = ffcommon.GoAStr(t)
	return
}
/**
 * Initialize libavdevice and register all the input and output devices.
 */
//void avdevice_register_all(void);
func AvdeviceRegisterAll() ( err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("avdevice_register_all").Call()
	if t==0{

	}
	return
}
/**
 * Audio input devices iterator.
 *
 * If d is NULL, returns the first registered input audio/video device,
 * if d is non-NULL, returns the next registered input audio/video device after d
 * or NULL if d is the last one.
 */
//AVInputFormat *av_input_audio_device_next(AVInputFormat  *d);
//未测试
func (d *AVInputFormat)AvInputAudioDeviceNext() (res *AVInputFormat, err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_input_audio_device_next").Call(
		uintptr(unsafe.Pointer(d)),
		)
	res = (*AVInputFormat)(unsafe.Pointer(t))
	return
}
/**
 * Video input devices iterator.
 *
 * If d is NULL, returns the first registered input audio/video device,
 * if d is non-NULL, returns the next registered input audio/video device after d
 * or NULL if d is the last one.
 */
//AVInputFormat *av_input_video_device_next(AVInputFormat  *d);
//未测试
func (d *AVInputFormat)AvInputVideoDeviceNext() (res *AVInputFormat, err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_input_video_device_next").Call(
		uintptr(unsafe.Pointer(d)),
	)
	res = (*AVInputFormat)(unsafe.Pointer(t))
	return
}
/**
 * Audio output devices iterator.
 *
 * If d is NULL, returns the first registered output audio/video device,
 * if d is non-NULL, returns the next registered output audio/video device after d
 * or NULL if d is the last one.
 */
//AVOutputFormat *av_output_audio_device_next(AVOutputFormat *d);
//未测试
func (d *AVOutputFormat)AvOutputAudioDeviceNext() (res *AVOutputFormat, err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_output_audio_device_next").Call(
		uintptr(unsafe.Pointer(d)),
	)
	res = (*AVOutputFormat)(unsafe.Pointer(t))
	return
}
/**
 * Video output devices iterator.
 *
 * If d is NULL, returns the first registered output audio/video device,
 * if d is non-NULL, returns the next registered output audio/video device after d
 * or NULL if d is the last one.
 */
//AVOutputFormat *av_output_video_device_next(AVOutputFormat *d);
//未测试
func (d *AVOutputFormat)AvOutputVideoDeviceNext() (res *AVOutputFormat, err error) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_output_video_device_next").Call(
		uintptr(unsafe.Pointer(d)),
	)
	res = (*AVOutputFormat)(unsafe.Pointer(t))
	return
}

type  AVDeviceRect struct {
//int x;      /**< x coordinate of top left corner */
//int y;      /**< y coordinate of top left corner */
//int width;  /**< width */
//int height; /**< height */
}

/**
 * Send control message from application to device.
 *
 * @param s         device context.
 * @param type      message type.
 * @param data      message data. Exact type depends on message type.
 * @param data_size size of message data.
 * @return >= 0 on success, negative on error.
 *         AVERROR(ENOSYS) when device doesn't implement handler of the message.
 */
int avdevice_app_to_dev_control_message(struct AVFormatContext *s,
enum AVAppToDevMessageType type,
void *data, size_t data_size);

/**
 * Send control message from device to application.
 *
 * @param s         device context.
 * @param type      message type.
 * @param data      message data. Can be NULL.
 * @param data_size size of message data.
 * @return >= 0 on success, negative on error.
 *         AVERROR(ENOSYS) when application doesn't implement handler of the message.
 */
int avdevice_dev_to_app_control_message(struct AVFormatContext *s,
enum AVDevToAppMessageType type,
void *data, size_t data_size);

#if FF_API_DEVICE_CAPABILITIES
/**
 * Following API allows user to probe device capabilities (supported codecs,
 * pixel formats, sample formats, resolutions, channel counts, etc).
 * It is build on top op AVOption API.
 * Queried capabilities make it possible to set up converters of video or audio
 * parameters that fit to the device.
 *
 * List of capabilities that can be queried:
 *  - Capabilities valid for both audio and video devices:
 *    - codec:          supported audio/video codecs.
 *                      type: AV_OPT_TYPE_INT (AVCodecID value)
 *  - Capabilities valid for audio devices:
 *    - sample_format:  supported sample formats.
 *                      type: AV_OPT_TYPE_INT (AVSampleFormat value)
 *    - sample_rate:    supported sample rates.
 *                      type: AV_OPT_TYPE_INT
 *    - channels:       supported number of channels.
 *                      type: AV_OPT_TYPE_INT
 *    - channel_layout: supported channel layouts.
 *                      type: AV_OPT_TYPE_INT64
 *  - Capabilities valid for video devices:
 *    - pixel_format:   supported pixel formats.
 *                      type: AV_OPT_TYPE_INT (AVPixelFormat value)
 *    - window_size:    supported window sizes (describes size of the window size presented to the user).
 *                      type: AV_OPT_TYPE_IMAGE_SIZE
 *    - frame_size:     supported frame sizes (describes size of provided video frames).
 *                      type: AV_OPT_TYPE_IMAGE_SIZE
 *    - fps:            supported fps values
 *                      type: AV_OPT_TYPE_RATIONAL
 *
 * Value of the capability may be set by user using av_opt_set() function
 * and AVDeviceCapabilitiesQuery object. Following queries will
 * limit results to the values matching already set capabilities.
 * For example, setting a codec may impact number of formats or fps values
 * returned during next query. Setting invalid value may limit results to zero.
 *
 * Example of the usage basing on opengl output device:
 *
 * @code
 *  AVFormatContext *oc = NULL;
 *  AVDeviceCapabilitiesQuery *caps = NULL;
 *  AVOptionRanges *ranges;
 *  int ret;
 *
 *  if ((ret = avformat_alloc_output_context2(&oc, NULL, "opengl", NULL)) < 0)
 *      goto fail;
 *  if (avdevice_capabilities_create(&caps, oc, NULL) < 0)
 *      goto fail;
 *
 *  //query codecs
 *  if (av_opt_query_ranges(&ranges, caps, "codec", AV_OPT_MULTI_COMPONENT_RANGE)) < 0)
 *      goto fail;
 *  //pick codec here and set it
 *  av_opt_set(caps, "codec", AV_CODEC_ID_RAWVIDEO, 0);
 *
 *  //query format
 *  if (av_opt_query_ranges(&ranges, caps, "pixel_format", AV_OPT_MULTI_COMPONENT_RANGE)) < 0)
 *      goto fail;
 *  //pick format here and set it
 *  av_opt_set(caps, "pixel_format", AV_PIX_FMT_YUV420P, 0);
 *
 *  //query and set more capabilities
 *
 * fail:
 *  //clean up code
 *  avdevice_capabilities_free(&query, oc);
 *  avformat_free_context(oc);
 * @endcode
 */

/**
 * Structure describes device capabilities.
 *
 * It is used by devices in conjunction with av_device_capabilities AVOption table
 * to implement capabilities probing API based on AVOption API. Should not be used directly.
 */
type  AVDeviceCapabilitiesQuery  struct{
//const AVClass *av_class;
//AVFormatContext *device_context;
//enum AVCodecID codec;
//enum AVSampleFormat sample_format;
//enum AVPixelFormat pixel_format;
//int sample_rate;
//int channels;
//int64_t channel_layout;
//int window_width;
//int window_height;
//int frame_width;
//int frame_height;
//AVRational fps;
}

/**
 * AVOption table used by devices to implement device capabilities API. Should not be used by a user.
 */
attribute_deprecated
extern const AVOption av_device_capabilities[];

/**
 * Initialize capabilities probing API based on AVOption API.
 *
 * avdevice_capabilities_free() must be called when query capabilities API is
 * not used anymore.
 *
 * @param[out] caps      Device capabilities data. Pointer to a NULL pointer must be passed.
 * @param s              Context of the device.
 * @param device_options An AVDictionary filled with device-private options.
 *                       On return this parameter will be destroyed and replaced with a dict
 *                       containing options that were not found. May be NULL.
 *                       The same options must be passed later to avformat_write_header() for output
 *                       devices or avformat_open_input() for input devices, or at any other place
 *                       that affects device-private options.
 *
 * @return >= 0 on success, negative otherwise.
 */
attribute_deprecated
int avdevice_capabilities_create(AVDeviceCapabilitiesQuery **caps, AVFormatContext *s,
AVDictionary **device_options);

/**
 * Free resources created by avdevice_capabilities_create()
 *
 * @param caps Device capabilities data to be freed.
 * @param s    Context of the device.
 */
attribute_deprecated
void avdevice_capabilities_free(AVDeviceCapabilitiesQuery **caps, AVFormatContext *s);
#endif

/**
 * Structure describes basic parameters of the device.
 */
type AVDeviceInfo struct{
//char *device_name;                   /**< device name, format depends on device */
//char *device_description;            /**< human friendly name */
}

/**
 * List of devices.
 */
type AVDeviceInfoList struct {
//AVDeviceInfo **devices;              /**< list of autodetected devices */
//int nb_devices;                      /**< number of autodetected devices */
//int default_device;                  /**< index of default device or -1 if no default */
}

/**
 * List devices.
 *
 * Returns available device names and their parameters.
 *
 * @note: Some devices may accept system-dependent device names that cannot be
 *        autodetected. The list returned by this function cannot be assumed to
 *        be always completed.
 *
 * @param s                device context.
 * @param[out] device_list list of autodetected devices.
 * @return count of autodetected devices, negative on error.
 */
int avdevice_list_devices(struct AVFormatContext *s, AVDeviceInfoList **device_list);

/**
 * Convenient function to free result of avdevice_list_devices().
 *
 * @param devices device list to be freed.
 */
void avdevice_free_list_devices(AVDeviceInfoList **device_list);

/**
 * List devices.
 *
 * Returns available device names and their parameters.
 * These are convinient wrappers for avdevice_list_devices().
 * Device context is allocated and deallocated internally.
 *
 * @param device           device format. May be NULL if device name is set.
 * @param device_name      device name. May be NULL if device format is set.
 * @param device_options   An AVDictionary filled with device-private options. May be NULL.
 *                         The same options must be passed later to avformat_write_header() for output
 *                         devices or avformat_open_input() for input devices, or at any other place
 *                         that affects device-private options.
 * @param[out] device_list list of autodetected devices
 * @return count of autodetected devices, negative on error.
 * @note device argument takes precedence over device_name when both are set.
 */
int avdevice_list_input_sources(struct AVInputFormat *device, const char *device_name,
AVDictionary *device_options, AVDeviceInfoList **device_list);
int avdevice_list_output_sinks(struct AVOutputFormat *device, const char *device_name,
AVDictionary *device_options, AVDeviceInfoList **device_list);

