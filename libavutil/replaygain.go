package libavutil

import "ffmpeg-go/ffcommon"

/**
 * ReplayGain information (see
 * http://wiki.hydrogenaudio.org/index.php?title=ReplayGain_1.0_specification).
 * The size of this struct is a part of the public ABI.
 */
type AVReplayGain struct {
	/**
	 * Track replay gain in microbels (divide by 100000 to get the value in dB).
	 * Should be set to INT32_MIN when unknown.
	 */
	TrackGain ffcommon.FInt32T
	/**
	 * Peak track amplitude, with 100000 representing full scale (but values
	 * may overflow). 0 when unknown.
	 */
	TrackPeak ffcommon.FUint32T
	/**
	 * Same as track_gain, but for the whole album.
	 */
	AlbumGain ffcommon.FInt32T
	/**
	 * Same as track_peak, but for the whole album,
	 */
	AlbumPeak ffcommon.FUint32T
}
