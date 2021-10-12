package libavutil

import (
	"ffmpeg-go/ffcommon"
	"syscall"
	"unsafe"
)

/**
 * Return a channel layout id that matches name, or 0 if no match is found.
 *
 * name can be one or several of the following notations,
 * separated by '+' or '|':
 * - the name of an usual channel layout (mono, stereo, 4.0, quad, 5.0,
 *   5.0(side), 5.1, 5.1(side), 7.1, 7.1(wide), downmix);
 * - the name of a single channel (FL, FR, FC, LFE, BL, BR, FLC, FRC, BC,
 *   SL, SR, TC, TFL, TFC, TFR, TBL, TBC, TBR, DL, DR);
 * - a number of channels, in decimal, followed by 'c', yielding
 *   the default channel layout for that number of channels (@see
 *   av_get_default_channel_layout);
 * - a channel layout mask, in hexadecimal starting with "0x" (see the
 *   AV_CH_* macros).
 *
 * Example: "stereo+FC" = "2c+FC" = "2c+1c" = "0x7"
 */
//uint64_t av_get_channel_layout(const char *name);
//未测试
func AvGetChannelLayout(name ffcommon.FConstCharP) (res ffcommon.FUint64T, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_get_channel_layout").Call(
		uintptr(unsafe.Pointer(namep)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FUint64T(t)
	return
}

/**
 * Return a channel layout and the number of channels based on the specified name.
 *
 * This function is similar to (@see av_get_channel_layout), but can also parse
 * unknown channel layout specifications.
 *
 * @param[in]  name             channel layout specification string
 * @param[out] channel_layout   parsed channel layout (0 if unknown)
 * @param[out] nb_channels      number of channels
 *
 * @return 0 on success, AVERROR(EINVAL) if the parsing fails.
 */
//int av_get_extended_channel_layout(const char *name, uint64_t* channel_layout, int* nb_channels);
//未测试
func AvGetExtendedChannelLayout(name ffcommon.FConstCharP, channel_layout *ffcommon.FUint64T, nb_channels *ffcommon.FInt) (res ffcommon.FUint64T, err error) {
	var t uintptr
	var namep *byte
	namep, err = syscall.BytePtrFromString(name)
	if err != nil {
		return
	}
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_get_extended_channel_layout").Call(
		uintptr(unsafe.Pointer(namep)),
		uintptr(unsafe.Pointer(channel_layout)),
		uintptr(unsafe.Pointer(nb_channels)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FUint64T(t)
	return
}

/**
 * Return a description of a channel layout.
 * If nb_channels is <= 0, it is guessed from the channel_layout.
 *
 * @param buf put here the string containing the channel layout
 * @param buf_size size in bytes of the buffer
 */
//void av_get_channel_layout_string(char *buf, int buf_size, int nb_channels, uint64_t channel_layout);
//未测试
func AvGetChannelLayoutString(buf ffcommon.FBuf, buf_size ffcommon.FInt, nb_channels ffcommon.FInt, channel_layout ffcommon.FUint64T) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_get_channel_layout_string").Call(
		uintptr(unsafe.Pointer(buf)),
		uintptr(buf_size),
		uintptr(nb_channels),
		uintptr(channel_layout),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Append a description of a channel layout to a bprint buffer.
 */
//void av_bprint_channel_layout(struct AVBPrint *bp, int nb_channels, uint64_t channel_layout);
//未测试
func (bp *AVBPrint) AvBprintChannelLayout(nb_channels ffcommon.FInt, channel_layout ffcommon.FUint64T) (err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_bprint_channel_layout").Call(
		uintptr(nb_channels),
		uintptr(channel_layout),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	return
}

/**
 * Return the number of channels in the channel layout.
 */
//int av_get_channel_layout_nb_channels(uint64_t channel_layout);
//未测试
func AvGetChannelLayoutNbChannels(channel_layout ffcommon.FUint64T) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_get_channel_layout_nb_channels").Call(
		uintptr(channel_layout),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Return default channel layout for a given number of channels.
 */
//int64_t av_get_default_channel_layout(int nb_channels);
//未测试
func AvGetDefaultChannelLayout(nb_channels ffcommon.FInt) (res ffcommon.FInt64T, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_get_default_channel_layout").Call(
		uintptr(nb_channels),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt64T(t)
	return
}

/**
 * Get the index of a channel in channel_layout.
 *
 * @param channel a channel layout describing exactly one channel which must be
 *                present in channel_layout.
 *
 * @return index of channel in channel_layout on success, a negative AVERROR
 *         on error.
 */
//int av_get_channel_layout_channel_index(uint64_t channel_layout,
//uint64_t channel);
//未测试
func AvGetChannelLayoutChannelIndex(channel_layout ffcommon.FUint64T, channel ffcommon.FUint64T) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_get_channel_layout_channel_index").Call(
		uintptr(channel_layout),
		uintptr(channel),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Get the channel with the given index in channel_layout.
 */
//uint64_t av_channel_layout_extract_channel(uint64_t channel_layout, int index);
//未测试
func AvChannelLayoutExtractChannel(channel_layout ffcommon.FUint64T, index ffcommon.FInt) (res ffcommon.FInt64T, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_channel_layout_extract_channel").Call(
		uintptr(channel_layout),
		uintptr(index),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt64T(t)
	return
}

/**
 * Get the name of a given channel.
 *
 * @return channel name on success, NULL on error.
 */
//const char *av_get_channel_name(uint64_t channel);
//未测试
func AvGetChannelName(channel ffcommon.FUint64T) (res ffcommon.FConstCharP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_get_channel_name").Call(
		uintptr(channel),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.GoAStr(t)
	return
}

/**
 * Get the description of a given channel.
 *
 * @param channel  a channel layout with a single channel
 * @return  channel description on success, NULL on error
 */
//const char *av_get_channel_description(uint64_t channel);
//未测试
func AvGetChannelDescription(channel ffcommon.FUint64T) (res ffcommon.FConstCharP, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_get_channel_description").Call(
		uintptr(channel),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.GoAStr(t)
	return
}

/**
 * Get the value and name of a standard channel layout.
 *
 * @param[in]  index   index in an internal list, starting at 0
 * @param[out] layout  channel layout mask
 * @param[out] name    name of the layout
 * @return  0  if the layout exists,
 *          <0 if index is beyond the limits
 */
//int av_get_standard_channel_layout(unsigned index, uint64_t *layout,
//const char **name);
//未测试
func AvGetStandardChannelLayout(index ffcommon.FUnsigned, layout *ffcommon.FUint64T, name *ffcommon.FBuf) (res ffcommon.FInt, err error) {
	var t uintptr
	t, _, _ = ffcommon.GetAvutilDll().NewProc("av_get_standard_channel_layout").Call(
		uintptr(index),
		uintptr(unsafe.Pointer(layout)),
		uintptr(unsafe.Pointer(&name)),
	)
	if err != nil {
		//return
	}
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}
