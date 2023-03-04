package libavformat

import (
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/ffmpeg-go/libavutil"
)

/*
 * copyright (c) 2001 Fabrice Bellard
 *
 * This file is part of FFmpeg.
 *
 * FFmpeg is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * FFmpeg is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with FFmpeg; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA
 */
//#ifndef AVFORMAT_AVIO_H
//const AVFORMAT_AVIO_H

/**
 * @file
 * @ingroup lavf_io
 * Buffered I/O operations
 */

//#include <stdint.h>
//
//#include "../libavutil/common.h"
//#include "../libavutil/dict.h"
//#include "../libavutil/log.h"
//
//#include "../libavformat/version.h"

/**
 * Seeking works like for a local file.
 */
const AVIO_SEEKABLE_NORMAL = (1 << 0)

/**
 * Seeking by timestamp with avio_seek_time() is possible.
 */
const AVIO_SEEKABLE_TIME = (1 << 1)

/**
 * Callback for checking whether to abort blocking functions.
 * AVERROR_EXIT is returned in this case by the interrupted
 * function. During blocking operations, callback is called with
 * opaque as parameter. If the callback returns 1, the
 * blocking operation will be aborted.
 *
 * No members can be added to this struct without a major bump, if
 * new elements have been added after this struct in AVFormatContext
 * or AVIOContext.
 */
type AVIOInterruptCB struct {

	//int (*callback)(void*);
	Callback uintptr
	//void *opaque;
	Opaque ffcommon.FVoidP
}

/**
 * Directory entry types.
 */
type AVIODirEntryType int32

const (
	AVIO_ENTRY_UNKNOWN = iota
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
 * Describes single entry of the directory.
 *
 * Only name and type fields are guaranteed be set.
 * Rest of fields are protocol or/and platform dependent and might be unknown.
 */
type AVIODirEntry struct {
	Name ffcommon.FCharPStruct /**< Filename */
	Type ffcommon.FInt         /**< Type of the entry */
	Utf8 ffcommon.FInt         /**< Set to 1 when name is encoded with UTF-8, 0 otherwise.
	  Name can be encoded with UTF-8 even though 0 is set. */
	Size                  ffcommon.FInt64T /**< File size in bytes, -1 if unknown. */
	ModificationTimestamp ffcommon.FInt64T /**< Time of last modification in microseconds since unix
	  epoch, -1 if unknown. */
	AccessTimestamp ffcommon.FInt64T /**< Time of last access in microseconds since unix epoch,
	  -1 if unknown. */
	StatusChangeTimestamp ffcommon.FInt64T /**< Time of last status change in microseconds since unix
	  epoch, -1 if unknown. */
	UserId   ffcommon.FInt64T /**< User ID of owner, -1 if unknown. */
	GroupId  ffcommon.FInt64T /**< Group ID of owner, -1 if unknown. */
	Filemode ffcommon.FInt64T /**< Unix file mode, -1 if unknown. */
}

type AVIODirContext struct {
	UrlContext uintptr //*URLContext;

}

/**
 * Different data types that can be returned via the AVIO
 * write_data_type callback.
 */
type AVIODataMarkerType int32

const (
	/**
	 * Header data; this needs to be present for the stream to be decodeable.
	 */
	AVIO_DATA_MARKER_HEADER = iota
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
 * Bytestream IO Context.
 * New fields can be added to the end with minor version bumps.
 * Removal, reordering and changes to existing fields require a major
 * version bump.
 * sizeof(AVIOContext) must not be used outside libav*.
 *
 * @note None of the function pointers in AVIOContext should be called
 *       directly, they should only be set by the client application
 *       when implementing custom I/O. Normally these are set to the
 *       function pointers specified in avio_alloc_context()
 */
type AVClass = libavutil.AVClass
type AVIOContext struct {

	/**
	 * A class for private options.
	 *
	 * If this AVIOContext is created by avio_open2(), av_class is set and
	 * passes the options down to protocols.
	 *
	 * If this AVIOContext is manually allocated, then av_class may be set by
	 * the caller.
	 *
	 * warning -- this field can be NULL, be sure to not pass this AVIOContext
	 * to any av_opt_* functions in that case.
	 */
	AvClass *AVClass

	/*
	 * The following shows the relationship between buffer, buf_ptr,
	 * buf_ptr_max, buf_end, buf_size, and pos, when reading and when writing
	 * (since AVIOContext is used for both):
	 *
	 **********************************************************************************
	 *                                   READING
	 **********************************************************************************
	 *
	 *                            |              buffer_size              |
	 *                            |---------------------------------------|
	 *                            |                                       |
	 *
	 *                         buffer          buf_ptr       buf_end
	 *                            +---------------+-----------------------+
	 *                            |/ / / / / / / /|/ / / / / / /|         |
	 *  read buffer:              |/ / consumed / | to be read /|         |
	 *                            |/ / / / / / / /|/ / / / / / /|         |
	 *                            +---------------+-----------------------+
	 *
	 *                                                         pos
	 *              +-------------------------------------------+-----------------+
	 *  input file: |                                           |                 |
	 *              +-------------------------------------------+-----------------+
	 *
	 *
	 **********************************************************************************
	 *                                   WRITING
	 **********************************************************************************
	 *
	 *                             |          buffer_size                 |
	 *                             |--------------------------------------|
	 *                             |                                      |
	 *
	 *                                                buf_ptr_max
	 *                          buffer                 (buf_ptr)       buf_end
	 *                             +-----------------------+--------------+
	 *                             |/ / / / / / / / / / / /|              |
	 *  write buffer:              | / / to be flushed / / |              |
	 *                             |/ / / / / / / / / / / /|              |
	 *                             +-----------------------+--------------+
	 *                               buf_ptr can be in this
	 *                               due to a backward seek
	 *
	 *                            pos
	 *               +-------------+----------------------------------------------+
	 *  output file: |             |                                              |
	 *               +-------------+----------------------------------------------+
	 *
	 */
	Buffer     ffcommon.FUnsignedCharPStruct /**< Start of the buffer. */
	BufferSize ffcommon.FInt                 /**< Maximum buffer size */
	BufPtr     ffcommon.FUnsignedCharPStruct /**< Current position in the buffer */
	BufEnd     ffcommon.FUnsignedCharPStruct /**< End of the data, may be less than
	  buffer+buffer_size if the read function returned
	  less data than requested, e.g. for streams where
	  no more data has been received yet. */
	Opaque ffcommon.FVoidP /**< A private pointer, passed to the read/write/seek/...
	  functions. */
	//int (*read_packet)(void *opaque, uint8_t *buf, int buf_size);
	ReadPacket uintptr
	//int (*write_packet)(void *opaque, uint8_t *buf, int buf_size);
	WritePacket uintptr
	//int64_t (*seek)(void *opaque, int64_t offset, int whence);
	Seek          uintptr
	Pos           ffcommon.FInt64T /**< position in the file of the current buffer */
	EofReached    ffcommon.FInt    /**< true if was unable to read due to error or eof */
	WriteFlag     ffcommon.FInt    /**< true if open for writing */
	MaxPacketSize ffcommon.FInt
	Checksum      ffcommon.FUnsignedLong
	ChecksumPtr   ffcommon.FUnsignedCharPStruct
	//unsigned long (*update_checksum)(unsigned long checksum, const uint8_t *buf, unsigned int size);
	UpdateChecksum uintptr
	Error          ffcommon.FInt /**< contains the error code or 0 if no error happened */
	/**
	 * Pause or resume playback for network streaming protocols - e.g. MMS.
	 */
	//int (*read_pause)(void *opaque, int pause);
	ReadPause uintptr
	/**
	 * Seek to a given timestamp in stream with the specified stream_index.
	 * Needed for some network streaming protocols which don't support seeking
	 * to byte position.
	 */
	//int64_t (*read_seek)(void *opaque, int stream_index,
	//int64_t timestamp, int flags);
	ReadSeek uintptr
	/**
	 * A combination of AVIO_SEEKABLE_ flags or 0 when the stream is not seekable.
	 */
	Seekable ffcommon.FInt

	/**
	 * max filesize, used to limit allocations
	 * This field is internal to libavformat and access from outside is not allowed.
	 */
	Maxsize ffcommon.FInt64T

	/**
	 * avio_read and avio_write should if possible be satisfied directly
	 * instead of going through a buffer, and avio_seek will always
	 * call the underlying seek function directly.
	 */
	Direct ffcommon.FInt

	/**
	 * Bytes read statistic
	 * This field is internal to libavformat and access from outside is not allowed.
	 */
	BytesRead ffcommon.FInt64T

	/**
	 * seek statistic
	 * This field is internal to libavformat and access from outside is not allowed.
	 */
	SeekCount ffcommon.FInt

	/**
	 * writeout statistic
	 * This field is internal to libavformat and access from outside is not allowed.
	 */
	WriteoutCount ffcommon.FInt

	/**
	 * Original buffer size
	 * used internally after probing and ensure seekback to reset the buffer size
	 * This field is internal to libavformat and access from outside is not allowed.
	 */
	OrigBufferSize ffcommon.FInt

	/**
	 * Threshold to favor readahead over seek.
	 * This is current internal only, do not use from outside.
	 */
	ShortSeekThreshold ffcommon.FInt

	/**
	 * ',' separated list of allowed protocols.
	 */
	ProtocolWhitelist ffcommon.FCharPStruct

	/**
	 * ',' separated list of disallowed protocols.
	 */
	ProtocolBlacklist ffcommon.FCharPStruct

	/**
	 * A callback that is used instead of write_packet.
	 */
	//int (*write_data_type)(void *opaque, uint8_t *buf, int buf_size,
	//enum AVIODataMarkerType type, int64_t time);
	WriteDataType uintptr
	/**
	 * If set, don't call write_data_type separately for AVIO_DATA_MARKER_BOUNDARY_POINT,
	 * but ignore them and treat them as AVIO_DATA_MARKER_UNKNOWN (to avoid needlessly
	 * small chunks of data returned from the callback).
	 */
	IgnoreBoundaryPoint ffcommon.FInt

	/**
	 * Internal, not meant to be used from outside of AVIOContext.
	 */
	CurrentType AVIODataMarkerType
	LastTime    ffcommon.FInt64T

	/**
	 * A callback that is used instead of short_seek_threshold.
	 * This is current internal only, do not use from outside.
	 */
	//int (*short_seek_get)(void *opaque);
	ShortSeekGet uintptr
	Written      ffcommon.FInt64T

	/**
	 * Maximum reached position before a backward seek in the write buffer,
	 * used keeping track of already written data for a later flush.
	 */
	BufPtrMax ffcommon.FUnsignedCharPStruct

	/**
	 * Try to buffer at least this amount of data before flushing it
	 */
	MinPacketSize ffcommon.FInt
}

/**
 * Return the name of the protocol that will handle the passed URL.
 *
 * NULL is returned if no protocol could be found for the given URL.
 *
 * @return Name of the protocol or NULL.
 */
//const char *avio_find_protocol_name(const char *url);
func AvioFindProtocolName(url ffcommon.FConstCharP) (res ffcommon.FConstCharP) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_find_protocol_name").Call(
		ffcommon.UintPtrFromString(url),
	)
	res = ffcommon.StringFromPtr(t)
	return
}

/**
 * Return AVIO_FLAG_* access flags corresponding to the access permissions
 * of the resource in url, or a negative value corresponding to an
 * AVERROR code in case of failure. The returned access flags are
 * masked by the value in flags.
 *
 * @note This function is intrinsically unsafe, in the sense that the
 * checked resource may change its existence or permission status from
 * one call to another. Thus you should not trust the returned value,
 * unless you are sure that no other processes are accessing the
 * checked resource.
 */
//int avio_check(const char *url, int flags);
func AvioCheck(url ffcommon.FConstCharP, flags ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_check").Call(
		ffcommon.UintPtrFromString(url),
		uintptr(flags),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Move or rename a resource.
 *
 * @note url_src and url_dst should share the same protocol and authority.
 *
 * @param url_src url to resource to be moved
 * @param url_dst new url to resource if the operation succeeded
 * @return >=0 on success or negative on error.
 */
//int avpriv_io_move(const char *url_src, const char *url_dst);
func AvprivIoMove(url_src, url_dst ffcommon.FConstCharP) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avpriv_io_move").Call(
		ffcommon.UintPtrFromString(url_src),
		ffcommon.UintPtrFromString(url_dst),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Delete a resource.
 *
 * @param url resource to be deleted.
 * @return >=0 on success or negative on error.
 */
//int avpriv_io_delete(const char *url);
func AvprivIoDelete(url ffcommon.FConstCharP) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avpriv_io_delete").Call(
		ffcommon.UintPtrFromString(url),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Open directory for reading.
 *
 * @param s       directory read context. Pointer to a NULL pointer must be passed.
 * @param url     directory to be listed.
 * @param options A dictionary filled with protocol-private options. On return
 *                this parameter will be destroyed and replaced with a dictionary
 *                containing options that were not found. May be NULL.
 * @return >=0 on success or negative on error.
 */
//int avio_open_dir(AVIODirContext **s, const char *url, AVDictionary **options);
func AvioOpenDir(s **AVIODirContext, url ffcommon.FConstCharP, options **AVDictionary) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_open_dir").Call(
		uintptr(unsafe.Pointer(s)),
		ffcommon.UintPtrFromString(url),
		uintptr(unsafe.Pointer(options)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Get next directory entry.
 *
 * Returned entry must be freed with avio_free_directory_entry(). In particular
 * it may outlive AVIODirContext.
 *
 * @param s         directory read context.
 * @param[out] next next entry or NULL when no more entries.
 * @return >=0 on success or negative on error. End of list is not considered an
 *             error.
 */
//int avio_read_dir(AVIODirContext *s, AVIODirEntry **next);
func (s *AVIODirContext) AvioReadDir(next **AVIODirEntry) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_read_dir").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(next)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Close directory.
 *
 * @note Entries created using avio_read_dir() are not deleted and must be
 * freeded with avio_free_directory_entry().
 *
 * @param s         directory read context.
 * @return >=0 on success or negative on error.
 */
//int avio_close_dir(AVIODirContext **s);
func AvioCloseDir(s **AVIODirContext) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_close_dir").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Free entry allocated by avio_read_dir().
 *
 * @param entry entry to be freed.
 */
//void avio_free_directory_entry(AVIODirEntry **entry);
func AvioFreeDirectoryEntry(entry **AVIODirEntry) {
	ffcommon.GetAvformatDll().NewProc("avio_free_directory_entry").Call(
		uintptr(unsafe.Pointer(entry)),
	)
}

/**
 * Allocate and initialize an AVIOContext for buffered I/O. It must be later
 * freed with avio_context_free().
 *
 * @param buffer Memory block for input/output operations via AVIOContext.
 *        The buffer must be allocated with av_malloc() and friends.
 *        It may be freed and replaced with a new buffer by libavformat.
 *        AVIOContext.buffer holds the buffer currently in use,
 *        which must be later freed with av_free().
 * @param buffer_size The buffer size is very important for performance.
 *        For protocols with fixed blocksize it should be set to this blocksize.
 *        For others a typical size is a cache page, e.g. 4kb.
 * @param write_flag Set to 1 if the buffer should be writable, 0 otherwise.
 * @param opaque An opaque pointer to user-specific data.
 * @param read_packet  A function for refilling the buffer, may be NULL.
 *                     For stream protocols, must never return 0 but rather
 *                     a proper AVERROR code.
 * @param write_packet A function for writing the buffer contents, may be NULL.
 *        The function may not change the input buffers content.
 * @param seek A function for seeking to specified byte position, may be NULL.
 *
 * @return Allocated AVIOContext or NULL on failure.
 */
//AVIOContext *avio_alloc_context(
//unsigned char *buffer,
//int buffer_size,
//int write_flag,
//void *opaque,
//int (*read_packet)(void *opaque, uint8_t *buf, int buf_size),
//int (*write_packet)(void *opaque, uint8_t *buf, int buf_size),
//int64_t (*seek)(void *opaque, int64_t offset, int whence));
func AvioAllocContext(buffer ffcommon.FBuf,
	buffer_size,
	write_flag ffcommon.FInt,
	opaque ffcommon.FVoidP,
	read_packet func(opaque ffcommon.FVoidP, buf *ffcommon.FUint8T, buf_size ffcommon.FInt) uintptr,
	write_packet func(opaque ffcommon.FVoidP, buf *ffcommon.FUint8T, buf_size ffcommon.FInt) uintptr,
	seek func(opaque ffcommon.FVoidP, offset ffcommon.FInt64T, whence ffcommon.FInt) uintptr) (res *AVIOContext) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_alloc_context").Call(
		uintptr(unsafe.Pointer(buffer)),
		uintptr(buffer_size),
		uintptr(write_flag),
		opaque,
		ffcommon.NewCallback(read_packet),
		ffcommon.NewCallback(write_packet),
		ffcommon.NewCallback(seek),
	)
	res = (*AVIOContext)(unsafe.Pointer(t))
	return
}

/**
 * Free the supplied IO context and everything associated with it.
 *
 * @param s Double pointer to the IO context. This function will write NULL
 * into s.
 */
//void avio_context_free(AVIOContext **s);
func AvioContextFree(s **AVIOContext) {
	ffcommon.GetAvformatDll().NewProc("avio_context_free").Call(
		uintptr(unsafe.Pointer(s)),
	)
}

// void avio_w8(AVIOContext *s, int b);
func (s *AVIOContext) AvioW8(b ffcommon.FInt) {
	ffcommon.GetAvformatDll().NewProc("avio_w8").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(b),
	)
}

// void avio_write(AVIOContext *s, const unsigned char *buf, int size);
func (s *AVIOContext) AvioWrite(buf ffcommon.FUnsignedCharP, size ffcommon.FInt) {
	ffcommon.GetAvformatDll().NewProc("avio_write").Call(
		uintptr(unsafe.Pointer(s)),
		ffcommon.UintPtrFromString(buf),
		uintptr(size),
	)
}

// void avio_wl64(AVIOContext *s, uint64_t val);
func (s *AVIOContext) AvioWl64(val ffcommon.FUint64T) {
	ffcommon.GetAvformatDll().NewProc("avio_wl64").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(val),
	)
}

// void avio_wb64(AVIOContext *s, uint64_t val);
func (s *AVIOContext) AvioWb64(val ffcommon.FUint64T) {
	ffcommon.GetAvformatDll().NewProc("avio_wb64").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(val),
	)
}

// void avio_wl32(AVIOContext *s, unsigned int val);
func (s *AVIOContext) AvioWl32(val ffcommon.FUnsignedInt) {
	ffcommon.GetAvformatDll().NewProc("avio_wl32").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(val),
	)
}

// void avio_wb32(AVIOContext *s, unsigned int val);
func (s *AVIOContext) AvioWb32(val ffcommon.FUnsignedInt) {
	ffcommon.GetAvformatDll().NewProc("avio_wb32").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(val),
	)
}

// void avio_wl24(AVIOContext *s, unsigned int val);
func (s *AVIOContext) AvioWl24(val ffcommon.FUnsignedInt) {
	ffcommon.GetAvformatDll().NewProc("avio_wl24").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(val),
	)
}

// void avio_wb24(AVIOContext *s, unsigned int val);
func (s *AVIOContext) AvioWb24(val ffcommon.FUnsignedInt) {
	ffcommon.GetAvformatDll().NewProc("avio_wb24").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(val),
	)
}

// void avio_wl16(AVIOContext *s, unsigned int val);
func (s *AVIOContext) AvioWl16(val ffcommon.FUnsignedInt) {
	ffcommon.GetAvformatDll().NewProc("avio_wl16").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(val),
	)
}

// void avio_wb16(AVIOContext *s, unsigned int val);
func (s *AVIOContext) AvioWb16(val ffcommon.FUnsignedInt) {
	ffcommon.GetAvformatDll().NewProc("avio_wb16").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(val),
	)
}

/**
 * Write a NULL-terminated string.
 * @return number of bytes written.
 */
//int avio_put_str(AVIOContext *s, const char *str);
func (s *AVIOContext) AvioPutStr(str ffcommon.FConstCharP) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_put_str").Call(
		uintptr(unsafe.Pointer(s)),
		ffcommon.UintPtrFromString(str),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Convert an UTF-8 string to UTF-16LE and write it.
 * @param s the AVIOContext
 * @param str NULL-terminated UTF-8 string
 *
 * @return number of bytes written.
 */
//int avio_put_str16le(AVIOContext *s, const char *str);
func (s *AVIOContext) AvioPutStr16le(str ffcommon.FConstCharP) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_put_str16le").Call(
		uintptr(unsafe.Pointer(s)),
		ffcommon.UintPtrFromString(str),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Convert an UTF-8 string to UTF-16BE and write it.
 * @param s the AVIOContext
 * @param str NULL-terminated UTF-8 string
 *
 * @return number of bytes written.
 */
//int avio_put_str16be(AVIOContext *s, const char *str);
func (s *AVIOContext) AvioPutStr16be(str ffcommon.FConstCharP) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_put_str16be").Call(
		uintptr(unsafe.Pointer(s)),
		ffcommon.UintPtrFromString(str),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Mark the written bytestream as a specific type.
 *
 * Zero-length ranges are omitted from the output.
 *
 * @param time the stream time the current bytestream pos corresponds to
 *             (in AV_TIME_BASE units), or AV_NOPTS_VALUE if unknown or not
 *             applicable
 * @param type the kind of data written starting at the current pos
 */
//void avio_write_marker(AVIOContext *s, int64_t time, enum AVIODataMarkerType type);
func (s *AVIOContext) AvioWriteMarker(time ffcommon.FInt64T, type0 AVIODataMarkerType) {
	ffcommon.GetAvformatDll().NewProc("avio_write_marker").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(time),
		uintptr(type0),
	)
}

/**
 * ORing this as the "whence" parameter to a seek function causes it to
 * return the filesize without seeking anywhere. Supporting this is optional.
 * If it is not supported then the seek function will return <0.
 */
const AVSEEK_SIZE = 0x10000

/**
 * Passing this flag as the "whence" parameter to a seek function causes it to
 * seek by any means (like reopening and linear reading) or other normally unreasonable
 * means that can be extremely slow.
 * This may be ignored by the seek code.
 */
const AVSEEK_FORCE = 0x20000

/**
 * fseek() equivalent for AVIOContext.
 * @return new position or AVERROR.
 */
//int64_t avio_seek(AVIOContext *s, int64_t offset, int whence);
func (s *AVIOContext) AvioSeek(offset ffcommon.FInt64T, whence ffcommon.FInt) (res ffcommon.FInt64T) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_seek").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(offset),
		uintptr(whence),
	)
	res = ffcommon.FInt64T(t)
	return
}

/**
 * Skip given number of bytes forward
 * @return new position or AVERROR.
 */
//int64_t avio_skip(AVIOContext *s, int64_t offset);
func (s *AVIOContext) AvioSkip(offset ffcommon.FInt64T) (res ffcommon.FInt64T) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_skip").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(offset),
	)
	res = ffcommon.FInt64T(t)
	return
}

/**
 * ftell() equivalent for AVIOContext.
 * @return position or AVERROR.
 */
//static av_always_inline int64_t avio_tell(AVIOContext *s)
//{
//return avio_seek(s, 0, SEEK_CUR);
//}
const SEEK_CUR = 1

func (s *AVIOContext) AvioTell() (res ffcommon.FInt64T) {
	res = s.AvioSeek(0, SEEK_CUR)
	return
}

/**
 * Get the filesize.
 * @return filesize or AVERROR
 */
//int64_t avio_size(AVIOContext *s);
func (s *AVIOContext) AvioSize() (res ffcommon.FInt64T) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_size").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = ffcommon.FInt64T(t)
	return
}

/**
 * Similar to feof() but also returns nonzero on read errors.
 * @return non zero if and only if at end of file or a read error happened when reading.
 */
//int avio_feof(AVIOContext *s);
func (s *AVIOContext) AvioFeof() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_feof").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Writes a formatted string to the context.
 * @return number of bytes written, < 0 on error.
 */
//int avio_printf(AVIOContext *s, const char *fmt, ...) av_printf_format(2, 3);
func (s *AVIOContext) AvioPrintf(fmt0 ...ffcommon.FConstCharP) (res ffcommon.FInt) {
	uintptrs := make([]uintptr, 0)
	uintptrs = append(uintptrs, uintptr(unsafe.Pointer(s)))
	for i := 0; i < len(fmt0); i++ {
		uintptrs = append(uintptrs, ffcommon.UintPtrFromString(fmt0[i]))
	}
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_printf").Call(
		uintptrs...,
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Write a NULL terminated array of strings to the context.
 * Usually you don't need to use this function directly but its macro wrapper,
 * avio_print.
 */
//void avio_print_string_array(AVIOContext *s, const char *strings[]);
func (s *AVIOContext) AvioPrintStringArray(strings *ffcommon.FBuf) {
	ffcommon.GetAvformatDll().NewProc("avio_print_string_array").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(strings)),
	)
}

/**
 * Write strings (const char *) to the context.
 * This is a convenience macro around avio_print_string_array and it
 * automatically creates the string array from the variable argument list.
 * For simple string concatenations this function is more performant than using
 * avio_printf since it does not need a temporary buffer.
 */
//const avio_print(s, ...) \
//avio_print_string_array(s, (const char*[]){__VA_ARGS__, NULL})

/**
 * Force flushing of buffered data.
 *
 * For write streams, force the buffered data to be immediately written to the output,
 * without to wait to fill the internal buffer.
 *
 * For read streams, discard all currently buffered data, and advance the
 * reported file position to that of the underlying stream. This does not
 * read new data, and does not perform any seeks.
 */
//void avio_flush(AVIOContext *s);
func (s *AVIOContext) AvioFlush() {
	ffcommon.GetAvformatDll().NewProc("avio_flush").Call(
		uintptr(unsafe.Pointer(s)),
	)
}

/**
 * Read size bytes from AVIOContext into buf.
 * @return number of bytes read or AVERROR
 */
//int avio_read(AVIOContext *s, unsigned char *buf, int size);
func (s *AVIOContext) AvioRead(buf ffcommon.FUnsignedCharP, size ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_read").Call(
		uintptr(unsafe.Pointer(s)),
		ffcommon.UintPtrFromString(buf),
		uintptr(size),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Read size bytes from AVIOContext into buf. Unlike avio_read(), this is allowed
 * to read fewer bytes than requested. The missing bytes can be read in the next
 * call. This always tries to read at least 1 byte.
 * Useful to reduce latency in certain cases.
 * @return number of bytes read or AVERROR
 */
//int avio_read_partial(AVIOContext *s, unsigned char *buf, int size);
func (s *AVIOContext) AvioReadPartial(buf ffcommon.FUnsignedCharP, size ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_read_partial").Call(
		uintptr(unsafe.Pointer(s)),
		ffcommon.UintPtrFromString(buf),
		uintptr(size),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * @name Functions for reading from AVIOContext
 * @{
 *
 * @note return 0 if EOF, so you cannot use it if EOF handling is
 *       necessary
 */
//int          avio_r8  (AVIOContext *s);
func (s *AVIOContext) AvioR8() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_r8").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = ffcommon.FInt(t)
	return
}

// unsigned int avio_rl16(AVIOContext *s);
func (s *AVIOContext) AvioRl16() (res ffcommon.FUnsignedInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_rl16").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = ffcommon.FUnsignedInt(t)
	return
}

// unsigned int avio_rl24(AVIOContext *s);
func (s *AVIOContext) AvioRl24() (res ffcommon.FUnsignedInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_rl24").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = ffcommon.FUnsignedInt(t)
	return
}

// unsigned int avio_rl32(AVIOContext *s);
func (s *AVIOContext) AvioRl32() (res ffcommon.FUnsignedInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_rl32").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = ffcommon.FUnsignedInt(t)
	return
}

// uint64_t     avio_rl64(AVIOContext *s);
func (s *AVIOContext) AvioRl64() (res ffcommon.FUint64T) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_rl64").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = ffcommon.FUint64T(t)
	return
}

// unsigned int avio_rb16(AVIOContext *s);
func (s *AVIOContext) AvioRb16() (res ffcommon.FUnsignedInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_rb16").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = ffcommon.FUnsignedInt(t)
	return
}

// unsigned int avio_rb24(AVIOContext *s);
func (s *AVIOContext) AvioRb24() (res ffcommon.FUnsignedInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_rb24").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = ffcommon.FUnsignedInt(t)
	return
}

// unsigned int avio_rb32(AVIOContext *s);
func (s *AVIOContext) AvioRb32() (res ffcommon.FUnsignedInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_rb32").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = ffcommon.FUnsignedInt(t)
	return
}

// uint64_t     avio_rb64(AVIOContext *s);
func (s *AVIOContext) AvioRb64() (res ffcommon.FUint64T) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_rb64").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = ffcommon.FUint64T(t)
	return
}

/**
 * @}
 */

/**
 * Read a string from pb into buf. The reading will terminate when either
 * a NULL character was encountered, maxlen bytes have been read, or nothing
 * more can be read from pb. The result is guaranteed to be NULL-terminated, it
 * will be truncated if buf is too small.
 * Note that the string is not interpreted or validated in any way, it
 * might get truncated in the middle of a sequence for multi-byte encodings.
 *
 * @return number of bytes read (is always <= maxlen).
 * If reading ends on EOF or error, the return value will be one more than
 * bytes actually read.
 */
//int avio_get_str(AVIOContext *pb, int maxlen, char *buf, int buflen);
func (pb *AVIOContext) AvioGetStr(maxlen ffcommon.FInt, buf ffcommon.FCharP, buflen ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_get_str").Call(
		uintptr(unsafe.Pointer(pb)),
		uintptr(maxlen),
		ffcommon.UintPtrFromString(buf),
		uintptr(buflen),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Read a UTF-16 string from pb and convert it to UTF-8.
 * The reading will terminate when either a null or invalid character was
 * encountered or maxlen bytes have been read.
 * @return number of bytes read (is always <= maxlen)
 */
//int avio_get_str16le(AVIOContext *pb, int maxlen, char *buf, int buflen);
func (pb *AVIOContext) AvioGetStr16le(maxlen ffcommon.FInt, buf ffcommon.FCharP, buflen ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_get_str16le").Call(
		uintptr(unsafe.Pointer(pb)),
		uintptr(maxlen),
		ffcommon.UintPtrFromString(buf),
		uintptr(buflen),
	)
	res = ffcommon.FInt(t)
	return
}

// int avio_get_str16be(AVIOContext *pb, int maxlen, char *buf, int buflen);
func (pb *AVIOContext) AvioGetStr16be(maxlen ffcommon.FInt, buf ffcommon.FCharP, buflen ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_get_str16be").Call(
		uintptr(unsafe.Pointer(pb)),
		uintptr(maxlen),
		ffcommon.UintPtrFromString(buf),
		uintptr(buflen),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * @name URL open modes
 * The flags argument to avio_open must be one of the following
 * constants, optionally ORed with other flags.
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

/**
 * Create and initialize a AVIOContext for accessing the
 * resource indicated by url.
 * @note When the resource indicated by url has been opened in
 * read+write mode, the AVIOContext can be used only for writing.
 *
 * @param s Used to return the pointer to the created AVIOContext.
 * In case of failure the pointed to value is set to NULL.
 * @param url resource to access
 * @param flags flags which control how the resource indicated by url
 * is to be opened
 * @return >= 0 in case of success, a negative value corresponding to an
 * AVERROR code in case of failure
 */
//int avio_open(AVIOContext **s, const char *url, int flags);
func AvioOpen(s **AVIOContext, url ffcommon.FConstCharP, flags ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_open").Call(
		uintptr(unsafe.Pointer(s)),
		ffcommon.UintPtrFromString(url),
		uintptr(flags),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Create and initialize a AVIOContext for accessing the
 * resource indicated by url.
 * @note When the resource indicated by url has been opened in
 * read+write mode, the AVIOContext can be used only for writing.
 *
 * @param s Used to return the pointer to the created AVIOContext.
 * In case of failure the pointed to value is set to NULL.
 * @param url resource to access
 * @param flags flags which control how the resource indicated by url
 * is to be opened
 * @param int_cb an interrupt callback to be used at the protocols level
 * @param options  A dictionary filled with protocol-private options. On return
 * this parameter will be destroyed and replaced with a dict containing options
 * that were not found. May be NULL.
 * @return >= 0 in case of success, a negative value corresponding to an
 * AVERROR code in case of failure
 */
//int avio_open2(AVIOContext **s, const char *url, int flags,
//const AVIOInterruptCB *int_cb, AVDictionary **options);
func AvioOpen2(s **AVIOContext, url ffcommon.FConstCharP, flags ffcommon.FInt,
	int_cb *AVIOInterruptCB, options **AVDictionary) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_open2").Call(
		uintptr(unsafe.Pointer(s)),
		ffcommon.UintPtrFromString(url),
		uintptr(flags),
		uintptr(unsafe.Pointer(int_cb)),
		uintptr(unsafe.Pointer(options)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Close the resource accessed by the AVIOContext s and free it.
 * This function can only be used if s was opened by avio_open().
 *
 * The internal buffer is automatically flushed before closing the
 * resource.
 *
 * @return 0 on success, an AVERROR < 0 on error.
 * @see avio_closep
 */
//int avio_close(AVIOContext *s);
func (s *AVIOContext) AvioClose() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_close").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Close the resource accessed by the AVIOContext *s, free it
 * and set the pointer pointing to it to NULL.
 * This function can only be used if s was opened by avio_open().
 *
 * The internal buffer is automatically flushed before closing the
 * resource.
 *
 * @return 0 on success, an AVERROR < 0 on error.
 * @see avio_close
 */
//int avio_closep(AVIOContext **s);
func AvioClosep(s **AVIOContext) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_closep").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Open a write only memory stream.
 *
 * @param s new IO context
 * @return zero if no error.
 */
//int avio_open_dyn_buf(AVIOContext **s);
func AvioOpenDynBuf(s **AVIOContext) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_open_dyn_buf").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Return the written size and a pointer to the buffer.
 * The AVIOContext stream is left intact.
 * The buffer must NOT be freed.
 * No padding is added to the buffer.
 *
 * @param s IO context
 * @param pbuffer pointer to a byte buffer
 * @return the length of the byte buffer
 */
//int avio_get_dyn_buf(AVIOContext *s, uint8_t **pbuffer);
func (s *AVIOContext) AvioGetDynBuf(pbuffer **ffcommon.FUint8T) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_get_dyn_buf").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(pbuffer)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Return the written size and a pointer to the buffer. The buffer
 * must be freed with av_free().
 * Padding of AV_INPUT_BUFFER_PADDING_SIZE is added to the buffer.
 *
 * @param s IO context
 * @param pbuffer pointer to a byte buffer
 * @return the length of the byte buffer
 */
//int avio_close_dyn_buf(AVIOContext *s, uint8_t **pbuffer);
func (s *AVIOContext) AvioCloseDynBuf(pbuffer **ffcommon.FUint8T) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_close_dyn_buf").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(pbuffer)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Iterate through names of available protocols.
 *
 * @param opaque A private pointer representing current protocol.
 *        It must be a pointer to NULL on first iteration and will
 *        be updated by successive calls to avio_enum_protocols.
 * @param output If set to 1, iterate over output protocols,
 *               otherwise over input protocols.
 *
 * @return A static string containing the name of current protocol or NULL
 */
//const char *avio_enum_protocols(void **opaque, int output);
func AvioEnumProtocols(opaque *ffcommon.FVoidP, output ffcommon.FInt) (res ffcommon.FConstCharP) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_enum_protocols").Call(
		uintptr(unsafe.Pointer(opaque)),
		uintptr(output),
	)
	res = ffcommon.StringFromPtr(t)
	return
}

/**
 * Get AVClass by names of available protocols.
 *
 * @return A AVClass of input protocol name or NULL
 */
//const AVClass *avio_protocol_get_class(const char *name);
func AvioProtocolGetClass(name ffcommon.FConstCharP) (res *AVClass) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_protocol_get_class").Call(
		ffcommon.UintPtrFromString(name),
	)
	res = (*AVClass)(unsafe.Pointer(t))
	return
}

/**
 * Pause and resume playing - only meaningful if using a network streaming
 * protocol (e.g. MMS).
 *
 * @param h     IO context from which to call the read_pause function pointer
 * @param pause 1 for pause, 0 for resume
 */
//int     avio_pause(AVIOContext *h, int pause);
func (h *AVIOContext) AvioPause(pause ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_pause").Call(
		uintptr(unsafe.Pointer(h)),
		uintptr(pause),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Seek to a given timestamp relative to some component stream.
 * Only meaningful if using a network streaming protocol (e.g. MMS.).
 *
 * @param h IO context from which to call the seek function pointers
 * @param stream_index The stream index that the timestamp is relative to.
 *        If stream_index is (-1) the timestamp should be in AV_TIME_BASE
 *        units from the beginning of the presentation.
 *        If a stream_index >= 0 is used and the protocol does not support
 *        seeking based on component streams, the call will fail.
 * @param timestamp timestamp in AVStream.time_base units
 *        or if there is no stream specified then in AV_TIME_BASE units.
 * @param flags Optional combination of AVSEEK_FLAG_BACKWARD, AVSEEK_FLAG_BYTE
 *        and AVSEEK_FLAG_ANY. The protocol may silently ignore
 *        AVSEEK_FLAG_BACKWARD and AVSEEK_FLAG_ANY, but AVSEEK_FLAG_BYTE will
 *        fail if used and not supported.
 * @return >= 0 on success
 * @see AVInputFormat::read_seek
 */
//int64_t avio_seek_time(AVIOContext *h, int stream_index,
//int64_t timestamp, int flags);
func (h *AVIOContext) AvioSeekTime(stream_index ffcommon.FInt, timestamp ffcommon.FInt64T, flags ffcommon.FInt) (res ffcommon.FInt64T) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_seek_time").Call(
		uintptr(unsafe.Pointer(h)),
		uintptr(stream_index),
		uintptr(timestamp),
		uintptr(flags),
	)
	res = ffcommon.FInt64T(t)
	return
}

/* Avoid a warning. The header can not be included because it breaks c++. */
//struct AVBPrint;

/**
 * Read contents of h into print buffer, up to max_size bytes, or up to EOF.
 *
 * @return 0 for success (max_size bytes read or EOF reached), negative error
 * code otherwise
 */
//int avio_read_to_bprint(AVIOContext *h, struct AVBPrint *pb, size_t max_size);
type AVBPrint = libavutil.AVBPrint

func (h *AVIOContext) AvioReadToBprint(pb *AVBPrint, max_size ffcommon.FSizeT) (res ffcommon.FInt64T) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_read_to_bprint").Call(
		uintptr(unsafe.Pointer(h)),
		uintptr(unsafe.Pointer(pb)),
		uintptr(max_size),
	)
	res = ffcommon.FInt64T(t)
	return
}

/**
 * Accept and allocate a client context on a server context.
 * @param  s the server context
 * @param  c the client context, must be unallocated
 * @return   >= 0 on success or a negative value corresponding
 *           to an AVERROR on failure
 */
//int avio_accept(AVIOContext *s, AVIOContext **c);
func (h *AVIOContext) AvioAccept(c **AVIOContext) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_accept").Call(
		uintptr(unsafe.Pointer(h)),
		uintptr(unsafe.Pointer(c)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Perform one step of the protocol handshake to accept a new client.
 * This function must be called on a client returned by avio_accept() before
 * using it as a read/write context.
 * It is separate from avio_accept() because it may block.
 * A step of the handshake is defined by places where the application may
 * decide to change the proceedings.
 * For example, on a protocol with a request header and a reply header, each
 * one can constitute a step because the application may use the parameters
 * from the request to change parameters in the reply; or each individual
 * chunk of the request can constitute a step.
 * If the handshake is already finished, avio_handshake() does nothing and
 * returns 0 immediately.
 *
 * @param  c the client context to perform the handshake on
 * @return   0   on a complete and successful handshake
 *           > 0 if the handshake progressed, but is not complete
 *           < 0 for an AVERROR code
 */
//int avio_handshake(AVIOContext *c);
func (c *AVIOContext) AvioHandshake() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avio_handshake").Call(
		uintptr(unsafe.Pointer(c)),
	)
	res = ffcommon.FInt(t)
	return
}

//#endif /* AVFORMAT_AVIO_H */
