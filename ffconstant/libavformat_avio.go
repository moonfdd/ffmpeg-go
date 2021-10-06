package ffconstant

/**
 * Seeking works like for a local file.
 */
const AVIO_SEEKABLE_NORMAL = (1 << 0)

/**
 * Seeking by timestamp with avio_seek_time() is possible.
 */
const AVIO_SEEKABLE_TIME = (1 << 1)

/**
 * Directory entry types.
 */
type AVIODirEntryType int32

const (
	AVIO_ENTRY_UNKNOWN = 0
	AVIO_ENTRY_BLOCK_DEVICE
	AVIO_ENTRY_CHARACTER_DEVICE
	AVIO_ENTRY_DIRECTORY
	AVIO_ENTRY_NAMED_PIPE
	AVIO_ENTRY_SYMBOLIC_LINK
	AVIO_ENTRY_SOCKET
	AVIO_ENTRY_FILE
	AVIO_ENTRY_SERVER
	AVIO_ENTRY_SHARE
	AVIO_ENTRY_WORKGROUP
)

/**
 * Different data types that can be returned via the AVIO
 * write_data_type callback.
 */
type AVIODataMarkerType int32

const (
	/**
	 * Header data; this needs to be present for the stream to be decodeable.
	 */
	AVIO_DATA_MARKER_HEADER = 0
	/**
	 * A point in the output bytestream where a decoder can start decoding
	 * (i.e. a keyframe). A demuxer/decoder given the data flagged with
	 * AVIO_DATA_MARKER_HEADER, followed by any AVIO_DATA_MARKER_SYNC_POINT,
	 * should give decodeable results.
	 */
	AVIO_DATA_MARKER_SYNC_POINT
	/**
	 * A point in the output bytestream where a demuxer can start parsing
	 * (for non self synchronizing bytestream formats). That is, any
	 * non-keyframe packet start point.
	 */
	AVIO_DATA_MARKER_BOUNDARY_POINT
	/**
	 * This is any, unlabelled data. It can either be a muxer not marking
	 * any positions at all, it can be an actual boundary/sync point
	 * that the muxer chooses not to mark, or a later part of a packet/fragment
	 * that is cut into multiple write callbacks due to limited IO buffer size.
	 */
	AVIO_DATA_MARKER_UNKNOWN
	/**
	 * Trailer data, which doesn't contain actual content, but only for
	 * finalizing the output file.
	 */
	AVIO_DATA_MARKER_TRAILER
	/**
	 * A point in the output bytestream where the underlying AVIOContext might
	 * flush the buffer depending on latency or buffering requirements. Typically
	 * means the end of a packet.
	 */
	AVIO_DATA_MARKER_FLUSH_POINT
)

/**
 * @name URL open modes
 * The flags argument to avio_open must be one of the following
 * ffconstants, optionally ORed with other flags.
 * @{
 */
const AVIO_FLAG_READ = 1                                        /**< read-only */
const AVIO_FLAG_WRITE = 2                                       /**< write-only */
const AVIO_FLAG_READ_WRITE = (AVIO_FLAG_READ | AVIO_FLAG_WRITE) /**< read-write pseudo flag */
/**
 * @}
 */

/**
 * Use non-blocking mode.
 * If this flag is set, operations on the context will return
 * AVERROR(EAGAIN) if they can not be performed immediately.
 * If this flag is not set, operations on the context will never return
 * AVERROR(EAGAIN).
 * Note that this flag does not affect the opening/connecting of the
 * context. Connecting a protocol will always block if necessary (e.g. on
 * network protocols) but never hang (e.g. on busy devices).
 * Warning: non-blocking protocols is work-in-progress; this flag may be
 * silently ignored.
 */
const AVIO_FLAG_NONBLOCK = 8

/**
 * Use direct mode.
 * avio_read and avio_write should if possible be satisfied directly
 * instead of going through a buffer, and avio_seek will always
 * call the underlying seek function directly.
 */
const AVIO_FLAG_DIRECT = 0x8000
