package ffconstant

const AVPROBE_SCORE_RETRY = (AVPROBE_SCORE_MAX / 4)
const AVPROBE_SCORE_STREAM_RETRY = (AVPROBE_SCORE_MAX/4 - 1)

const AVPROBE_SCORE_EXTENSION = 50 ///< score for file extension
const AVPROBE_SCORE_MIME = 75      ///< score for file mime type
const AVPROBE_SCORE_MAX = 100      ///< maximum score

const AVPROBE_PADDING_SIZE = 32 ///< extra allocated bytes at the end of the probe buffer

/// Demuxer will use avio_open, no opened file should be provided by the caller.
const AVFMT_NOFILE = 0x0001
const AVFMT_NEEDNUMBER = 0x0002    /**< Needs '%d' in filename. */
const AVFMT_SHOW_IDS = 0x0008      /**< Show format stream IDs numbers. */
const AVFMT_GLOBALHEADER = 0x0040  /**< Format wants global header. */
const AVFMT_NOTIMESTAMPS = 0x0080  /**< Format does not need / have any timestamps. */
const AVFMT_GENERIC_INDEX = 0x0100 /**< Use generic index building code. */
const AVFMT_TS_DISCONT = 0x0200    /**< Format allows timestamp discontinuities. Note, muxers always require valid (monotone) timestamps */
const AVFMT_VARIABLE_FPS = 0x0400  /**< Format allows variable fps. */
const AVFMT_NODIMENSIONS = 0x0800  /**< Format does not need width/height */
const AVFMT_NOSTREAMS = 0x1000     /**< Format does not require any streams */
const AVFMT_NOBINSEARCH = 0x2000   /**< Format does not allow to fall back on binary search via read_timestamp */
const AVFMT_NOGENSEARCH = 0x4000   /**< Format does not allow to fall back on generic search */
const AVFMT_NO_BYTE_SEEK = 0x8000  /**< Format does not allow seeking by bytes */
const AVFMT_ALLOW_FLUSH = 0x10000  /**< Format allows flushing. If not set, the muxer will not receive a NULL packet in the write_packet function. */
const AVFMT_TS_NONSTRICT = 0x20000 /**< Format does not require strictly
  increasing timestamps, but they must
  still be monotonic */
const AVFMT_TS_NEGATIVE = 0x40000 /**< Format allows muxing negative
  timestamps. If not set the timestamp
  will be shifted in av_write_frame and
  av_interleaved_write_frame so they
  start from 0.
  The user or muxer can override this through
  AVFormatContext.avoid_negative_ts
*/

const AVFMT_SEEK_TO_PTS = 0x4000000 /**< Seeking is based on PTS */

/**
 * @}
 */
type AVStreamParseType int32

const (
	AVSTREAM_PARSE_NONE       = 0
	AVSTREAM_PARSE_FULL       /**< full parsing and repack */
	AVSTREAM_PARSE_HEADERS    /**< Only parse headers, do not repack. */
	AVSTREAM_PARSE_TIMESTAMPS /**< full parsing and interpolation of timestamps for frames not starting on a packet boundary */
	AVSTREAM_PARSE_FULL_ONCE  /**< full parsing and repack of the first frame only, only implemented for H.264 currently */
	AVSTREAM_PARSE_FULL_RAW   /**< full parsing and repack with timestamp and position generation by parser for raw
	  this assumes that each packet in the file contains no demuxer level headers and
	  just codec level data, otherwise position generation would fail */
)

/* Timestamp in AVStream.time_base units, preferably the time from which on correctly decoded frames are available
 * when seeking to this entry. That means preferable PTS on keyframe based formats.
 * But demuxers can choose to store a different timestamp, if it is more convenient for the implementation or nothing better
 * is known
 */
const AVINDEX_KEYFRAME = 0x0001
const AVINDEX_DISCARD_FRAME = 0x0002

const AV_DISPOSITION_DEFAULT = 0x0001
const AV_DISPOSITION_DUB = 0x0002
const AV_DISPOSITION_ORIGINAL = 0x0004
const AV_DISPOSITION_COMMENT = 0x0008
const AV_DISPOSITION_LYRICS = 0x0010
const AV_DISPOSITION_KARAOKE = 0x0020

/*
 * Track should be used during playback by default.
 * Useful for subtitle track that should be displayed
 * even when user did not explicitly ask for subtitles.
 */
const AV_DISPOSITION_FORCED = 0x0040
const AV_DISPOSITION_HEARING_IMPAIRED = 0x0080 /**< stream for hearing impaired audiences */
const AV_DISPOSITION_VISUAL_IMPAIRED = 0x0100  /**< stream for visual impaired audiences */
const AV_DISPOSITION_CLEAN_EFFECTS = 0x0200    /**< stream without voice */
/**
 * The stream is stored in the file as an attached picture/"cover art" (e.g.
 * APIC frame in ID3v2). The first (usually only) packet associated with it
 * will be returned among the first few packets read from the file unless
 * seeking takes place. It can also be accessed at any time in
 * AVStream.attached_pic.
 */
const AV_DISPOSITION_ATTACHED_PIC = 0x0400

/**
 * The stream is sparse, and contains thumbnail images, often corresponding
 * to chapter markers. Only ever used with AV_DISPOSITION_ATTACHED_PIC.
 */
const AV_DISPOSITION_TIMED_THUMBNAILS = 0x0800

/**
 * To specify text track kind (different from subtitles default).
 */
const AV_DISPOSITION_CAPTIONS = 0x10000
const AV_DISPOSITION_DESCRIPTIONS = 0x20000
const AV_DISPOSITION_METADATA = 0x40000
const AV_DISPOSITION_DEPENDENT = 0x80000    ///< dependent audio stream (mix_type=0 in mpegts)
const AV_DISPOSITION_STILL_IMAGE = 0x100000 ///< still images in video stream (still_picture_flag=1 in mpegts)

/**
 * Options for behavior on timestamp wrap detection.
 */
const AV_PTS_WRAP_IGNORE = 0      ///< ignore the wrap
const AV_PTS_WRAP_ADD_OFFSET = 1  ///< add the format specific offset on wrap detection
const AV_PTS_WRAP_SUB_OFFSET = -1 ///< subtract the format specific offset on wrap detection

/**
 * - demuxing: the demuxer read new metadata from the file and updated
 *     AVStream.metadata accordingly
 * - muxing: the user updated AVStream.metadata and wishes the muxer to write
 *     it into the file
 */
const AVSTREAM_EVENT_FLAG_METADATA_UPDATED = 0x0001

/**
 * - demuxing: new packets for this stream were read from the file. This
 *   event is informational only and does not guarantee that new packets
 *   for this stream will necessarily be returned from av_read_frame().
 */
const AVSTREAM_EVENT_FLAG_NEW_PACKETS = (1 << 1)

const AV_PROGRAM_RUNNING = 1

const AVFMTCTX_NOHEADER = 0x0001 /**< signal that no header is present
  (streams are added dynamically) */
const AVFMTCTX_UNSEEKABLE = 0x0002 /**< signal that the stream is definitely
  not seekable, and attempts to call the
  seek function will fail. For some
  network protocols (e.g. HLS), this can
  change dynamically at runtime. */

/**
 * The duration of a video can be estimated through various ways, and this enum can be used
 * to know how the duration was estimated.
 */
type AVDurationEstimationMethod int32

const (
	AVFMT_DURATION_FROM_PTS     = 0 ///< Duration accurately estimated from PTSes
	AVFMT_DURATION_FROM_STREAM      ///< Duration estimated from a stream with a known duration
	AVFMT_DURATION_FROM_BITRATE     ///< Duration estimated from bitrate (less accurate)
)

const AVFMT_FLAG_GENPTS = 0x0001          ///< Generate missing pts even if it requires parsing future frames.
const AVFMT_FLAG_IGNIDX = 0x0002          ///< Ignore index.
const AVFMT_FLAG_NONBLOCK = 0x0004        ///< Do not block when reading packets from input.
const AVFMT_FLAG_IGNDTS = 0x0008          ///< Ignore DTS on frames that contain both DTS & PTS
const AVFMT_FLAG_NOFILLIN = 0x0010        ///< Do not infer any values from other values, just return what is stored in the container
const AVFMT_FLAG_NOPARSE = 0x0020         ///< Do not use AVParsers, you also must set AVFMT_FLAG_NOFILLIN as the fillin code works on frames and no parsing -> no frames. Also seeking to frames can not work if parsing to find frame boundaries has been disabled
const AVFMT_FLAG_NOBUFFER = 0x0040        ///< Do not buffer frames when possible
const AVFMT_FLAG_CUSTOM_IO = 0x0080       ///< The caller has supplied a custom AVIOContext, don't avio_close() it.
const AVFMT_FLAG_DISCARD_CORRUPT = 0x0100 ///< Discard frames marked corrupted
const AVFMT_FLAG_FLUSH_PACKETS = 0x0200   ///< Flush the AVIOContext every packet.
/**
 * When muxing, try to avoid writing any random/volatile data to the output.
 * This includes any random IDs, real-time timestamps/dates, muxer version, etc.
 *
 * This flag is mainly intended for testing.
 */
const AVFMT_FLAG_BITEXACT = 0x0400

//#if FF_API_LAVF_MP4A_LATM
//const AVFMT_FLAG_MP4A_LATM   = 0x8000 ///< Deprecated, does nothing.
//#endif
//const AVFMT_FLAG_SORT_DTS   = 0x10000 ///< try to interleave outputted packets by dts (using this flag can slow demuxing down)
//#if FF_API_LAVF_PRIV_OPT
//const AVFMT_FLAG_PRIV_OPT  =  0x20000 ///< Enable use of private options by delaying codec open (deprecated, will do nothing once av_demuxer_open() is removed)
//#endif
//#if FF_API_LAVF_KEEPSIDE_FLAG
//const AVFMT_FLAG_KEEP_SIDE_DATA =0x40000 ///< Deprecated, does nothing.
//#endif
const AVFMT_FLAG_FAST_SEEK = 0x80000 ///< Enable fast, but inaccurate seeks for some formats
const AVFMT_FLAG_SHORTEST = 0x100000 ///< Stop muxing when the shortest stream stops.
const AVFMT_FLAG_AUTO_BSF = 0x200000 ///< Add bitstream filters as requested by the muxer

const FF_FDEBUG_TS = 0x0001

/**
 * - demuxing: the demuxer read new metadata from the file and updated
 *   AVFormatContext.metadata accordingly
 * - muxing: the user updated AVFormatContext.metadata and wishes the muxer to
 *   write it into the file
 */
const AVFMT_EVENT_FLAG_METADATA_UPDATED = 0x0001

const AVFMT_AVOID_NEG_TS_AUTO = -1             ///< Enabled when required by target format
const AVFMT_AVOID_NEG_TS_MAKE_NON_NEGATIVE = 1 ///< Shift timestamps so they are non negative
const AVFMT_AVOID_NEG_TS_MAKE_ZERO = 2         ///< Shift timestamps so that they start at 0

/**
 * @}
 */

const AVSEEK_FLAG_BACKWARD = 1 ///< seek backward
const AVSEEK_FLAG_BYTE = 2     ///< seeking based on position in bytes
const AVSEEK_FLAG_ANY = 4      ///< seek to any frame, even non-keyframes
const AVSEEK_FLAG_FRAME = 8    ///< seeking based on frame number

/**
 * @addtogroup lavf_encoding
 * @{
 */

const AVSTREAM_INIT_IN_WRITE_HEADER = 0 ///< stream parameters initialized in avformat_write_header
const AVSTREAM_INIT_IN_INIT_OUTPUT = 1  ///< stream parameters initialized in avformat_init_output

type AVTimebaseSource int32

const (
	AVFMT_TBCF_AUTO = -1
	AVFMT_TBCF_DECODER
	AVFMT_TBCF_DEMUXER
	//#if FF_API_R_FRAME_RATE
	AVFMT_TBCF_R_FRAMERATE

//#endif
)
