package libavformat

import (
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/ffmpeg-go/libavcodec"
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

//#ifndef AVFORMAT_AVFORMAT_H
//#define AVFORMAT_AVFORMAT_H

/**
 * @file
 * @ingroup libavf
 * Main libavformat public API header
 */

/**
 * @defgroup libavf libavformat
 * I/O and Muxing/Demuxing Library
 *
 * Libavformat (lavf) is a library for dealing with various media container
 * formats. Its main two purposes are demuxing - i.e. splitting a media file
 * into component streams, and the reverse process of muxing - writing supplied
 * data in a specified container format. It also has an @ref lavf_io
 * "I/O module" which supports a number of protocols for accessing the data (e.g.
 * file, tcp, http and others).
 * Unless you are absolutely sure you won't use libavformat's network
 * capabilities, you should also call avformat_network_init().
 *
 * A supported input format is described by an AVInputFormat struct, conversely
 * an output format is described by AVOutputFormat. You can iterate over all
 * input/output formats using the  av_demuxer_iterate / av_muxer_iterate() functions.
 * The protocols layer is not part of the public API, so you can only get the names
 * of supported protocols with the avio_enum_protocols() function.
 *
 * Main lavf structure used for both muxing and demuxing is AVFormatContext,
 * which exports all information about the file being read or written. As with
 * most Libavformat structures, its size is not part of public ABI, so it cannot be
 * allocated on stack or directly with av_malloc(). To create an
 * AVFormatContext, use avformat_alloc_context() (some functions, like
 * avformat_open_input() might do that for you).
 *
 * Most importantly an AVFormatContext contains:
 * @li the @ref AVFormatContext.iformat "input" or @ref AVFormatContext.oformat
 * "output" format. It is either autodetected or set by user for input;
 * always set by user for output.
 * @li an @ref AVFormatContext.streams "array" of AVStreams, which describe all
 * elementary streams stored in the file. AVStreams are typically referred to
 * using their index in this array.
 * @li an @ref AVFormatContext.pb "I/O context". It is either opened by lavf or
 * set by user for input, always set by user for output (unless you are dealing
 * with an AVFMT_NOFILE format).
 *
 * @section lavf_options Passing options to (de)muxers
 * It is possible to configure lavf muxers and demuxers using the @ref avoptions
 * mechanism. Generic (format-independent) libavformat options are provided by
 * AVFormatContext, they can be examined from a user program by calling
 * av_opt_next() / av_opt_find() on an allocated AVFormatContext (or its AVClass
 * from avformat_get_class()). Private (format-specific) options are provided by
 * AVFormatContext.priv_data if and only if AVInputFormat.priv_class /
 * AVOutputFormat.priv_class of the corresponding format struct is non-NULL.
 * Further options may be provided by the @ref AVFormatContext.pb "I/O context",
 * if its AVClass is non-NULL, and the protocols layer. See the discussion on
 * nesting in @ref avoptions documentation to learn how to access those.
 *
 * @section urls
 * URL strings in libavformat are made of a scheme/protocol, a ':', and a
 * scheme specific string. URLs without a scheme and ':' used for local files
 * are supported but deprecated. "file:" should be used for local files.
 *
 * It is important that the scheme string is not taken from untrusted
 * sources without checks.
 *
 * Note that some schemes/protocols are quite powerful, allowing access to
 * both local and remote files, parts of them, concatenations of them, local
 * audio and video devices and so on.
 *
 * @{
 *
 * @defgroup lavf_decoding Demuxing
 * @{
 * Demuxers read a media file and split it into chunks of data (@em packets). A
 * @ref AVPacket "packet" contains one or more encoded frames which belongs to a
 * single elementary stream. In the lavf API this process is represented by the
 * avformat_open_input() function for opening a file, av_read_frame() for
 * reading a single packet and finally avformat_close_input(), which does the
 * cleanup.
 *
 * @section lavf_decoding_open Opening a media file
 * The minimum information required to open a file is its URL, which
 * is passed to avformat_open_input(), as in the following code:
 * @code
 * const char    *url = "file:in.mp3";
 * AVFormatContext *s = NULL;
 * int ret = avformat_open_input(&s, url, NULL, NULL);
 * if (ret < 0)
 *     abort();
 * @endcode
 * The above code attempts to allocate an AVFormatContext, open the
 * specified file (autodetecting the format) and read the header, exporting the
 * information stored there into s. Some formats do not have a header or do not
 * store enough information there, so it is recommended that you call the
 * avformat_find_stream_info() function which tries to read and decode a few
 * frames to find missing information.
 *
 * In some cases you might want to preallocate an AVFormatContext yourself with
 * avformat_alloc_context() and do some tweaking on it before passing it to
 * avformat_open_input(). One such case is when you want to use custom functions
 * for reading input data instead of lavf internal I/O layer.
 * To do that, create your own AVIOContext with avio_alloc_context(), passing
 * your reading callbacks to it. Then set the @em pb field of your
 * AVFormatContext to newly created AVIOContext.
 *
 * Since the format of the opened file is in general not known until after
 * avformat_open_input() has returned, it is not possible to set demuxer private
 * options on a preallocated context. Instead, the options should be passed to
 * avformat_open_input() wrapped in an AVDictionary:
 * @code
 * AVDictionary *options = NULL;
 * av_dict_set(&options, "video_size", "640x480", 0);
 * av_dict_set(&options, "pixel_format", "rgb24", 0);
 *
 * if (avformat_open_input(&s, url, NULL, &options) < 0)
 *     abort();
 * av_dict_free(&options);
 * @endcode
 * This code passes the private options 'video_size' and 'pixel_format' to the
 * demuxer. They would be necessary for e.g. the rawvideo demuxer, since it
 * cannot know how to interpret raw video data otherwise. If the format turns
 * out to be something different than raw video, those options will not be
 * recognized by the demuxer and therefore will not be applied. Such unrecognized
 * options are then returned in the options dictionary (recognized options are
 * consumed). The calling program can handle such unrecognized options as it
 * wishes, e.g.
 * @code
 * AVDictionaryEntry *e;
 * if (e = av_dict_get(options, "", NULL, AV_DICT_IGNORE_SUFFIX)) {
 *     fprintf(stderr, "Option %s not recognized by the demuxer.\n", e->key);
 *     abort();
 * }
 * @endcode
 *
 * After you have finished reading the file, you must close it with
 * avformat_close_input(). It will free everything associated with the file.
 *
 * @section lavf_decoding_read Reading from an opened file
 * Reading data from an opened AVFormatContext is done by repeatedly calling
 * av_read_frame() on it. Each call, if successful, will return an AVPacket
 * containing encoded data for one AVStream, identified by
 * AVPacket.stream_index. This packet may be passed straight into the libavcodec
 * decoding functions avcodec_send_packet() or avcodec_decode_subtitle2() if the
 * caller wishes to decode the data.
 *
 * AVPacket.pts, AVPacket.dts and AVPacket.duration timing information will be
 * set if known. They may also be unset (i.e. AV_NOPTS_VALUE for
 * pts/dts, 0 for duration) if the stream does not provide them. The timing
 * information will be in AVStream.time_base units, i.e. it has to be
 * multiplied by the timebase to convert them to seconds.
 *
 * A packet returned by av_read_frame() is always reference-counted,
 * i.e. AVPacket.buf is set and the user may keep it indefinitely.
 * The packet must be freed with av_packet_unref() when it is no
 * longer needed.
 *
 * @section lavf_decoding_seek Seeking
 * @}
 *
 * @defgroup lavf_encoding Muxing
 * @{
 * Muxers take encoded data in the form of @ref AVPacket "AVPackets" and write
 * it into files or other output bytestreams in the specified container format.
 *
 * The main API functions for muxing are avformat_write_header() for writing the
 * file header, av_write_frame() / av_interleaved_write_frame() for writing the
 * packets and av_write_trailer() for finalizing the file.
 *
 * At the beginning of the muxing process, the caller must first call
 * avformat_alloc_context() to create a muxing context. The caller then sets up
 * the muxer by filling the various fields in this context:
 *
 * - The @ref AVFormatContext.oformat "oformat" field must be set to select the
 *   muxer that will be used.
 * - Unless the format is of the AVFMT_NOFILE type, the @ref AVFormatContext.pb
 *   "pb" field must be set to an opened IO context, either returned from
 *   avio_open2() or a custom one.
 * - Unless the format is of the AVFMT_NOSTREAMS type, at least one stream must
 *   be created with the avformat_new_stream() function. The caller should fill
 *   the @ref AVStream.codecpar "stream codec parameters" information, such as the
 *   codec @ref AVCodecParameters.codec_type "type", @ref AVCodecParameters.codec_id
 *   "id" and other parameters (e.g. width / height, the pixel or sample format,
 *   etc.) as known. The @ref AVStream.time_base "stream timebase" should
 *   be set to the timebase that the caller desires to use for this stream (note
 *   that the timebase actually used by the muxer can be different, as will be
 *   described later).
 * - It is advised to manually initialize only the relevant fields in
 *   AVCodecParameters, rather than using @ref avcodec_parameters_copy() during
 *   remuxing: there is no guarantee that the codec context values remain valid
 *   for both input and output format contexts.
 * - The caller may fill in additional information, such as @ref
 *   AVFormatContext.metadata "global" or @ref AVStream.metadata "per-stream"
 *   metadata, @ref AVFormatContext.chapters "chapters", @ref
 *   AVFormatContext.programs "programs", etc. as described in the
 *   AVFormatContext documentation. Whether such information will actually be
 *   stored in the output depends on what the container format and the muxer
 *   support.
 *
 * When the muxing context is fully set up, the caller must call
 * avformat_write_header() to initialize the muxer internals and write the file
 * header. Whether anything actually is written to the IO context at this step
 * depends on the muxer, but this function must always be called. Any muxer
 * private options must be passed in the options parameter to this function.
 *
 * The data is then sent to the muxer by repeatedly calling av_write_frame() or
 * av_interleaved_write_frame() (consult those functions' documentation for
 * discussion on the difference between them; only one of them may be used with
 * a single muxing context, they should not be mixed). Do note that the timing
 * information on the packets sent to the muxer must be in the corresponding
 * AVStream's timebase. That timebase is set by the muxer (in the
 * avformat_write_header() step) and may be different from the timebase
 * requested by the caller.
 *
 * Once all the data has been written, the caller must call av_write_trailer()
 * to flush any buffered packets and finalize the output file, then close the IO
 * context (if any) and finally free the muxing context with
 * avformat_free_context().
 * @}
 *
 * @defgroup lavf_io I/O Read/Write
 * @{
 * @section lavf_io_dirlist Directory listing
 * The directory listing API makes it possible to list files on remote servers.
 *
 * Some of possible use cases:
 * - an "open file" dialog to choose files from a remote location,
 * - a recursive media finder providing a player with an ability to play all
 * files from a given directory.
 *
 * @subsection lavf_io_dirlist_open Opening a directory
 * At first, a directory needs to be opened by calling avio_open_dir()
 * supplied with a URL and, optionally, ::AVDictionary containing
 * protocol-specific parameters. The function returns zero or positive
 * integer and allocates AVIODirContext on success.
 *
 * @code
 * AVIODirContext *ctx = NULL;
 * if (avio_open_dir(&ctx, "smb://example.com/some_dir", NULL) < 0) {
 *     fprintf(stderr, "Cannot open directory.\n");
 *     abort();
 * }
 * @endcode
 *
 * This code tries to open a sample directory using smb protocol without
 * any additional parameters.
 *
 * @subsection lavf_io_dirlist_read Reading entries
 * Each directory's entry (i.e. file, another directory, anything else
 * within ::AVIODirEntryType) is represented by AVIODirEntry.
 * Reading consecutive entries from an opened AVIODirContext is done by
 * repeatedly calling avio_read_dir() on it. Each call returns zero or
 * positive integer if successful. Reading can be stopped right after the
 * NULL entry has been read -- it means there are no entries left to be
 * read. The following code reads all entries from a directory associated
 * with ctx and prints their names to standard output.
 * @code
 * AVIODirEntry *entry = NULL;
 * for (;;) {
 *     if (avio_read_dir(ctx, &entry) < 0) {
 *         fprintf(stderr, "Cannot list directory.\n");
 *         abort();
 *     }
 *     if (!entry)
 *         break;
 *     printf("%s\n", entry->name);
 *     avio_free_directory_entry(&entry);
 * }
 * @endcode
 * @}
 *
 * @defgroup lavf_codec Demuxers
 * @{
 * @defgroup lavf_codec_native Native Demuxers
 * @{
 * @}
 * @defgroup lavf_codec_wrappers External library wrappers
 * @{
 * @}
 * @}
 * @defgroup lavf_protos I/O Protocols
 * @{
 * @}
 * @defgroup lavf_internal Internal
 * @{
 * @}
 * @}
 */

//#include <time.h>
//#include <stdio.h>  /* FILE */
//#include "../libavcodec/avcodec.h"
//#include "../libavutil/dict.h"
//#include "../libavutil/log.h"
//
//#include "avio.h"
//#include "../libavformat/version.h"

//struct AVFormatContext;
//struct AVDeviceInfoList;
// type AVDeviceInfoList = libavdevice.AVDeviceInfoList

//struct AVDeviceCapabilitiesQuery;
// type AVDeviceCapabilitiesQuery = libavdevice.AVDeviceCapabilitiesQuery
type AVCodecID = libavcodec.AVCodecID
type AVCodecContext = libavcodec.AVCodecContext
type AVPacket = libavcodec.AVPacket
type AVDiscard = libavcodec.AVDiscard

/**
* @defgroup metadata_api Public Metadata API
* @{
* @ingroup libavf
* The metadata API allows libavformat to export metadata tags to a client
* application when demuxing. Conversely it allows a client application to
* set metadata when muxing.
*
* Metadata is exported or set as pairs of key/value strings in the 'metadata'
* fields of the AVFormatContext, AVStream, AVChapter and AVProgram structs
* using the @ref lavu_dict "AVDictionary" API. Like all strings in FFmpeg,
* metadata is assumed to be UTF-8 encoded Unicode. Note that metadata
* exported by demuxers isn't checked to be valid UTF-8 in most cases.
*
* Important concepts to keep in mind:
* -  Keys are unique; there can never be 2 tags with the same key. This is
*    also meant semantically, i.e., a demuxer should not knowingly produce
*    several keys that are literally different but semantically identical.
*    E.g., key=Author5, key=Author6. In this example, all authors must be
*    placed in the same tag.
* -  Metadata is flat, not hierarchical; there are no subtags. If you
*    want to store, e.g., the email address of the child of producer Alice
*    and actor Bob, that could have key=alice_and_bobs_childs_email_address.
* -  Several modifiers can be applied to the tag name. This is done by
*    appending a dash character ('-') and the modifier name in the order
*    they appear in the list below -- e.g. foo-eng-sort, not foo-sort-eng.
*    -  language -- a tag whose value is localized for a particular language
*       is appended with the ISO 639-2/B 3-letter language code.
*       For example: Author-ger=Michael, Author-eng=Mike
*       The original/default language is in the unqualified "Author" tag.
*       A demuxer should set a default if it sets any translated tag.
*    -  sorting  -- a modified version of a tag that should be used for
*       sorting will have '-sort' appended. E.g. artist="The Beatles",
*       artist-sort="Beatles, The".
* - Some protocols and demuxers support metadata updates. After a successful
*   call to av_read_frame(), AVFormatContext.event_flags or AVStream.event_flags
*   will be updated to indicate if metadata changed. In order to detect metadata
*   changes on a stream, you need to loop through all streams in the AVFormatContext
*   and check their individual event_flags.
*
* -  Demuxers attempt to export metadata in a generic format, however tags
*    with no generic equivalents are left as they are stored in the container.
*    Follows a list of generic tag names:
*
@verbatim
album        -- name of the set this work belongs to
album_artist -- main creator of the set/album, if different from artist.
                e.g. "Various Artists" for compilation albums.
artist       -- main creator of the work
comment      -- any additional description of the file.
composer     -- who composed the work, if different from artist.
copyright    -- name of copyright holder.
creation_time-- date when the file was created, preferably in ISO 8601.
date         -- date when the work was created, preferably in ISO 8601.
disc         -- number of a subset, e.g. disc in a multi-disc collection.
encoder      -- name/settings of the software/hardware that produced the file.
encoded_by   -- person/group who created the file.
filename     -- original name of the file.
genre        -- <self-evident>.
language     -- main language in which the work is performed, preferably
                in ISO 639-2 format. Multiple languages can be specified by
                separating them with commas.
performer    -- artist who performed the work, if different from artist.
                E.g for "Also sprach Zarathustra", artist would be "Richard
                Strauss" and performer "London Philharmonic Orchestra".
publisher    -- name of the label/publisher.
service_name     -- name of the service in broadcasting (channel name).
service_provider -- name of the service provider in broadcasting.
title        -- name of the work.
track        -- number of this work in the set, can be in form current/total.
variant_bitrate -- the total bitrate of the bitrate variant that the current stream is part of
@endverbatim
*
* Look in the examples section for an application example how to use the Metadata API.
*
* @}
*/

/* packet functions */

/**
 * Allocate and read the payload of a packet and initialize its
 * fields with default values.
 *
 * @param s    associated IO context
 * @param pkt packet
 * @param size desired payload size
 * @return >0 (read size) if OK, AVERROR_xxx otherwise
 */
//int av_get_packet(AVIOContext *s, AVPacket *pkt, int size);
func (s *AVIOContext) AvGetPacket(pkt *AVPacket, size ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_get_packet").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(pkt)),
		uintptr(size),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Read data and append it to the current content of the AVPacket.
 * If pkt->size is 0 this is identical to av_get_packet.
 * Note that this uses av_grow_packet and thus involves a realloc
 * which is inefficient. Thus this function should only be used
 * when there is no reasonable way to know (an upper bound of)
 * the final size.
 *
 * @param s    associated IO context
 * @param pkt packet
 * @param size amount of data to read
 * @return >0 (read size) if OK, AVERROR_xxx otherwise, previous data
 *         will not be lost even if an error occurs.
 */
//int av_append_packet(AVIOContext *s, AVPacket *pkt, int size);
func (s *AVIOContext) AvAppendPacket(pkt *AVPacket, size ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_append_packet").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(pkt)),
		uintptr(size),
	)
	res = ffcommon.FInt(t)
	return
}

/*************************************************/
/* input/output formats */

//struct AVCodecTag;
type AVCodecTag struct {
}

/**
 * This structure contains the data a format has to probe a file.
 */
type AVProbeData struct {
	Filename ffcommon.FCharPStruct
	Buf      ffcommon.FUnsignedCharPStruct /**< Buffer must have AVPROBE_PADDING_SIZE of extra allocated bytes filled with zero. */
	BufSize  ffcommon.FInt                 /**< Size of buf except extra allocated bytes */
	MimeType ffcommon.FUnsignedCharPStruct /**< mime_type, when known. */
}

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
 * @addtogroup lavf_encoding
 * @{
 */
type AVOutputFormat struct {
	Name ffcommon.FCharPStruct
	/**
	 * Descriptive name for the format, meant to be more human-readable
	 * than name. You should use the NULL_IF_CONFIG_SMALL() macro
	 * to define it.
	 */
	LongName   ffcommon.FCharPStruct
	MimeType   ffcommon.FCharPStruct
	Extensions ffcommon.FCharPStruct /**< comma-separated filename extensions */
	/* output support */
	AudioCodec    AVCodecID /**< default audio codec */
	VideoCodec    AVCodecID /**< default video codec */
	SubtitleCodec AVCodecID /**< default subtitle codec */
	/**
	 * can use flags: AVFMT_NOFILE, AVFMT_NEEDNUMBER,
	 * AVFMT_GLOBALHEADER, AVFMT_NOTIMESTAMPS, AVFMT_VARIABLE_FPS,
	 * AVFMT_NODIMENSIONS, AVFMT_NOSTREAMS, AVFMT_ALLOW_FLUSH,
	 * AVFMT_TS_NONSTRICT, AVFMT_TS_NEGATIVE
	 */
	Flags ffcommon.FInt

	/**
	 * List of supported codec_id-codec_tag pairs, ordered by "better
	 * choice first". The arrays are all terminated by AV_CODEC_ID_NONE.
	 */
	//const struct AVCodecTag * const *codec_tag;
	CodecTag *AVCodecTag

	PrivClass *AVClass ///< AVClass for the private context

	/*****************************************************************
	 * No fields below this line are part of the public API. They
	 * may not be used outside of libavformat and can be changed and
	 * removed at will.
	 * New public fields should be added right above.
	 *****************************************************************
	 */
	/**
	 * The ff_const59 define is not part of the public API and will
	 * be removed without further warning.
	 */
	//#if FF_API_AVIOFORMAT
	//const ff_const59
	//#else
	//const ff_const59 const
	//#endif
	//#if FF_API_NEXT
	//ff_const59 struct AVOutputFormat *next;
	Next *AVOutputFormat
	//#endif
	/**
	 * size of private data so that it can be allocated in the wrapper
	 */
	PrivDataSize ffcommon.FInt

	//int (*write_header)(struct AVFormatContext *);
	WriteHeader uintptr
	/**
	 * Write a packet. If AVFMT_ALLOW_FLUSH is set in flags,
	 * pkt can be NULL in order to flush data buffered in the muxer.
	 * When flushing, return 0 if there still is more data to flush,
	 * or 1 if everything was flushed and there is no more buffered
	 * data.
	 */
	//int (*write_packet)(struct AVFormatContext *, AVPacket *pkt);
	WritePacket uintptr
	//int (*write_trailer)(struct AVFormatContext *);
	WriteTrailer uintptr
	/**
	 * A format-specific function for interleavement.
	 * If unset, packets will be interleaved by dts.
	 */
	//int (*interleave_packet)(struct AVFormatContext *, AVPacket *out,
	//AVPacket *in, int flush);
	InterleavePacket uintptr
	/**
	 * Test if the given codec can be stored in this container.
	 *
	 * @return 1 if the codec is supported, 0 if it is not.
	 *         A negative number if unknown.
	 *         MKTAG('A', 'P', 'I', 'C') if the codec is only supported as AV_DISPOSITION_ATTACHED_PIC
	 */
	//int (*query_codec)(enum AVCodecID id, int std_compliance);
	QueryCodec uintptr

	//void (*get_output_timestamp)(struct AVFormatContext *s, int stream,
	//int64_t *dts, int64_t *wall);
	GetOutputTimestamp uintptr
	/**
	 * Allows sending messages from application to device.
	 */
	//int (*control_message)(struct AVFormatContext *s, int type,
	//void *data, size_t data_size);
	ControlMessage uintptr

	/**
	 * Write an uncoded AVFrame.
	 *
	 * See av_write_uncoded_frame() for details.
	 *
	 * The library will free *frame afterwards, but the muxer can prevent it
	 * by setting the pointer to NULL.
	 */
	//int (*write_uncoded_frame)(struct AVFormatContext *, int stream_index,
	//AVFrame **frame, unsigned flags);
	WriteUncodedFrame uintptr
	/**
	 * Returns device list with it properties.
	 * @see avdevice_list_devices() for more details.
	 */
	//int (*get_device_list)(struct AVFormatContext *s, struct AVDeviceInfoList *device_list);
	GetDeviceList uintptr
	//#if LIBAVFORMAT_VERSION_MAJOR < 59
	/**
	 * Initialize device capabilities submodule.
	 * @see avdevice_capabilities_create() for more details.
	 */
	//int (*create_device_capabilities)(struct AVFormatContext *s, struct AVDeviceCapabilitiesQuery *caps);
	CreateDeviceCapabilities uintptr
	/**
	 * Free device capabilities submodule.
	 * @see avdevice_capabilities_free() for more details.
	 */
	//int (*free_device_capabilities)(struct AVFormatContext *s, struct AVDeviceCapabilitiesQuery *caps);
	FreeDeviceCapabilities uintptr
	//#endif
	DataCodec AVCodecID /**< default data codec */
	/**
	 * Initialize format. May allocate data here, and set any AVFormatContext or
	 * AVStream parameters that need to be set before packets are sent.
	 * This method must not write output.
	 *
	 * Return 0 if streams were fully configured, 1 if not, negative AVERROR on failure
	 *
	 * Any allocations made here must be freed in deinit().
	 */
	//int (*init)(struct AVFormatContext *);
	Init uintptr
	/**
	 * Deinitialize format. If present, this is called whenever the muxer is being
	 * destroyed, regardless of whether or not the header has been written.
	 *
	 * If a trailer is being written, this is called after write_trailer().
	 *
	 * This is called if init() fails as well.
	 */
	//void (*deinit)(struct AVFormatContext *);
	Deinit uintptr
	/**
	 * Set up any necessary bitstream filtering and extract any extra data needed
	 * for the global header.
	 * Return 0 if more packets from this stream must be checked; 1 if not.
	 */
	//int (*check_bitstream)(struct AVFormatContext *, const AVPacket *pkt);
	CheckBitstream uintptr
}

/**
 * @}
 */

/**
 * @addtogroup lavf_decoding
 * @{
 */
type AVInputFormat struct {

	/**
	 * A comma separated list of short names for the format. New names
	 * may be appended with a minor bump.
	 */
	Name ffcommon.FCharPStruct

	/**
	 * Descriptive name for the format, meant to be more human-readable
	 * than name. You should use the NULL_IF_CONFIG_SMALL() macro
	 * to define it.
	 */
	LongName ffcommon.FCharPStruct

	/**
	 * Can use flags: AVFMT_NOFILE, AVFMT_NEEDNUMBER, AVFMT_SHOW_IDS,
	 * AVFMT_NOTIMESTAMPS, AVFMT_GENERIC_INDEX, AVFMT_TS_DISCONT, AVFMT_NOBINSEARCH,
	 * AVFMT_NOGENSEARCH, AVFMT_NO_BYTE_SEEK, AVFMT_SEEK_TO_PTS.
	 */
	Flags ffcommon.FInt

	/**
	 * If extensions are defined, then no probe is done. You should
	 * usually not use extension format guessing because it is not
	 * reliable enough
	 */
	Extensions ffcommon.FCharPStruct

	//const struct AVCodecTag * const *codec_tag;
	CodecTag  *AVCodecTag
	PrivClass *AVClass ///< AVClass for the private context

	/**
	 * Comma-separated list of mime types.
	 * It is used check for matching mime types while probing.
	 * @see av_probe_input_format2
	 */
	MimeType ffcommon.FCharPStruct

	/*****************************************************************
	 * No fields below this line are part of the public API. They
	 * may not be used outside of libavformat and can be changed and
	 * removed at will.
	 * New public fields should be added right above.
	 *****************************************************************
	 */
	//#if FF_API_NEXT
	//ff_const59 struct AVInputFormat *next;
	Next *AVInputFormat
	//#endif

	/**
	 * Raw demuxers store their codec ID here.
	 */
	RawCodecId ffcommon.FInt

	/**
	 * Size of private data so that it can be allocated in the wrapper.
	 */
	PrivDataSize ffcommon.FInt

	/**
	 * Tell if a given file has a chance of being parsed as this format.
	 * The buffer provided is guaranteed to be AVPROBE_PADDING_SIZE bytes
	 * big so you do not have to check for that unless you need more.
	 */
	//int (*read_probe)(const AVProbeData *);
	ReadProbe uintptr
	/**
	 * Read the format header and initialize the AVFormatContext
	 * structure. Return 0 if OK. 'avformat_new_stream' should be
	 * called to create new streams.
	 */
	//int (*read_header)(struct AVFormatContext *);
	ReadHeader uintptr
	/**
	 * Read one packet and put it in 'pkt'. pts and flags are also
	 * set. 'avformat_new_stream' can be called only if the flag
	 * AVFMTCTX_NOHEADER is used and only in the calling thread (not in a
	 * background thread).
	 * @return 0 on success, < 0 on error.
	 *         Upon returning an error, pkt must be unreferenced by the caller.
	 */
	//int (*read_packet)(struct AVFormatContext *, AVPacket *pkt);
	ReadPacket uintptr
	/**
	 * Close the stream. The AVFormatContext and AVStreams are not
	 * freed by this function
	 */
	//int (*read_close)(struct AVFormatContext *);
	ReadClose uintptr
	/**
	 * Seek to a given timestamp relative to the frames in
	 * stream component stream_index.
	 * @param stream_index Must not be -1.
	 * @param flags Selects which direction should be preferred if no exact
	 *              match is available.
	 * @return >= 0 on success (but not necessarily the new offset)
	 */
	//int (*read_seek)(struct AVFormatContext *,
	//int stream_index, int64_t timestamp, int flags);
	ReadSeek uintptr
	/**
	 * Get the next timestamp in stream[stream_index].time_base units.
	 * @return the timestamp or AV_NOPTS_VALUE if an error occurred
	 */
	//int64_t (*read_timestamp)(struct AVFormatContext *s, int stream_index,
	//int64_t *pos, int64_t pos_limit);
	ReadTimestamp uintptr
	/**
	 * Start/resume playing - only meaningful if using a network-based format
	 * (RTSP).
	 */
	//int (*read_play)(struct AVFormatContext *);
	ReadPlay uintptr
	/**
	 * Pause playing - only meaningful if using a network-based format
	 * (RTSP).
	 */
	//int (*read_pause)(struct AVFormatContext *);
	ReadPause uintptr
	/**
	 * Seek to timestamp ts.
	 * Seeking will be done so that the point from which all active streams
	 * can be presented successfully will be closest to ts and within min/max_ts.
	 * Active streams are all streams that have AVStream.discard < AVDISCARD_ALL.
	 */
	//int (*read_seek2)(struct AVFormatContext *s, int stream_index, int64_t min_ts, int64_t ts, int64_t max_ts, int flags);
	ReadSeek2 uintptr
	/**
	 * Returns device list with it properties.
	 * @see avdevice_list_devices() for more details.
	 */
	//int (*get_device_list)(struct AVFormatContext *s, struct AVDeviceInfoList *device_list);
	GetDeviceList uintptr
	//#if LIBAVFORMAT_VERSION_MAJOR < 59
	/**
	 * Initialize device capabilities submodule.
	 * @see avdevice_capabilities_create() for more details.
	 */
	//int (*create_device_capabilities)(struct AVFormatContext *s, struct AVDeviceCapabilitiesQuery *caps);
	CreateDeviceCapabilities uintptr
	/**
	 * Free device capabilities submodule.
	 * @see avdevice_capabilities_free() for more details.
	 */
	//int (*free_device_capabilities)(struct AVFormatContext *s, struct AVDeviceCapabilitiesQuery *caps);
	FreeDeviceCapabilities uintptr
	//#endif
}

/**
 * @}
 */
type AVStreamParseType int32

const (
	AVSTREAM_PARSE_NONE       = iota
	AVSTREAM_PARSE_FULL       /**< full parsing and repack */
	AVSTREAM_PARSE_HEADERS    /**< Only parse headers, do not repack. */
	AVSTREAM_PARSE_TIMESTAMPS /**< full parsing and interpolation of timestamps for frames not starting on a packet boundary */
	AVSTREAM_PARSE_FULL_ONCE  /**< full parsing and repack of the first frame only, only implemented for H.264 currently */
	AVSTREAM_PARSE_FULL_RAW   /**< full parsing and repack with timestamp and position generation by parser for raw
	  this assumes that each packet in the file contains no demuxer level headers and
	  just codec level data, otherwise position generation would fail */
)

type AVIndexEntry struct {
	Pos       ffcommon.FInt64T
	Timestamp ffcommon.FInt64T /**<
	 * Timestamp in AVStream.time_base units, preferably the time from which on correctly decoded frames are available
	 * when seeking to this entry. That means preferable PTS on keyframe based formats.
	 * But demuxers can choose to store a different timestamp, if it is more convenient for the implementation or nothing better
	 * is known
	 */
	/*const AVINDEX_KEYFRAME =0x0001
	  const AVINDEX_DISCARD_FRAME = 0x0002*/ /**
	 * Flag is used to indicate which frame should be discarded after decoding.
	 */
	//int flags:2;
	//int size:30; //Yeah, trying to keep the size of this small to reduce memory requirements (it is 24 vs. 32 bytes due to possible 8-byte alignment).
	FlagsPlusSize int32
	MinDistance   ffcommon.FInt /**< Minimum distance between this and the previous keyframe, used to avoid unneeded searching. */
}

const AV_DISPOSITION_DEFAULT = 0x0001
const AV_DISPOSITION_DUB = 0x0002
const AV_DISPOSITION_ORIGINAL = 0x0004
const AV_DISPOSITION_COMMENT = 0x0008
const AV_DISPOSITION_LYRICS = 0x0010
const AV_DISPOSITION_KARAOKE = 0x0020

/**
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

//typedef struct AVStreamInternal AVStreamInternal;

type AVStreamInternal struct {
}

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
 * Stream structure.
 * New fields can be added to the end with minor version bumps.
 * Removal, reordering and changes to existing fields require a major
 * version bump.
 * sizeof(AVStream) must not be used outside libav*.
 */
type AVRational = libavutil.AVRational
type AVDictionary = libavutil.AVDictionary
type AVPacketSideData = libavcodec.AVPacketSideData
type AVCodecParserContext = libavcodec.AVCodecParserContext
type AVCodecParameters = libavcodec.AVCodecParameters
type AVStream struct {
	Index ffcommon.FInt /**< stream index in AVFormatContext */
	/**
	 * Format-specific stream ID.
	 * decoding: set by libavformat
	 * encoding: set by the user, replaced by libavformat if left unset
	 */
	Id ffcommon.FInt
	//#if FF_API_LAVF_AVCTX
	/**
	 * @deprecated use the codecpar struct instead
	 */
	//attribute_deprecated
	//AVCodecContext *codec;
	Codec *AVCodecContext
	//#endif
	PrivData ffcommon.FVoidP

	/**
	 * This is the fundamental unit of time (in seconds) in terms
	 * of which frame timestamps are represented.
	 *
	 * decoding: set by libavformat
	 * encoding: May be set by the caller before avformat_write_header() to
	 *           provide a hint to the muxer about the desired timebase. In
	 *           avformat_write_header(), the muxer will overwrite this field
	 *           with the timebase that will actually be used for the timestamps
	 *           written into the file (which may or may not be related to the
	 *           user-provided one, depending on the format).
	 */
	TimeBase AVRational

	/**
	 * Decoding: pts of the first frame of the stream in presentation order, in stream time base.
	 * Only set this if you are absolutely 100% sure that the value you set
	 * it to really is the pts of the first frame.
	 * This may be undefined (AV_NOPTS_VALUE).
	 * @note The ASF header does NOT contain a correct start_time the ASF
	 * demuxer must NOT set this.
	 */
	StartTime ffcommon.FInt64T

	/**
	 * Decoding: duration of the stream, in stream time base.
	 * If a source file does not specify a duration, but does specify
	 * a bitrate, this value will be estimated from bitrate and file size.
	 *
	 * Encoding: May be set by the caller before avformat_write_header() to
	 * provide a hint to the muxer about the estimated duration.
	 */
	Duration ffcommon.FInt64T

	NbFrames ffcommon.FInt64T ///< number of frames in this stream if known or 0

	Disposition ffcommon.FInt /**< AV_DISPOSITION_* bit field */

	Discard AVDiscard ///< Selects which packets can be discarded at will and do not need to be demuxed.

	/**
	 * sample aspect ratio (0 if unknown)
	 * - encoding: Set by user.
	 * - decoding: Set by libavformat.
	 */
	SampleAspectRatio AVRational

	Metadata *AVDictionary

	/**
	 * Average framerate
	 *
	 * - demuxing: May be set by libavformat when creating the stream or in
	 *             avformat_find_stream_info().
	 * - muxing: May be set by the caller before avformat_write_header().
	 */
	AvgFrameRate AVRational

	/**
	 * For streams with AV_DISPOSITION_ATTACHED_PIC disposition, this packet
	 * will contain the attached picture.
	 *
	 * decoding: set by libavformat, must not be modified by the caller.
	 * encoding: unused
	 */
	AttachedPic AVPacket

	/**
	 * An array of side data that applies to the whole stream (i.e. the
	 * container does not allow it to change between packets).
	 *
	 * There may be no overlap between the side data in this array and side data
	 * in the packets. I.e. a given side data is either exported by the muxer
	 * (demuxing) / set by the caller (muxing) in this array, then it never
	 * appears in the packets, or the side data is exported / sent through
	 * the packets (always in the first packet where the value becomes known or
	 * changes), then it does not appear in this array.
	 *
	 * - demuxing: Set by libavformat when the stream is created.
	 * - muxing: May be set by the caller before avformat_write_header().
	 *
	 * Freed by libavformat in avformat_free_context().
	 *
	 * @see av_format_inject_global_side_data()
	 */
	SideData *AVPacketSideData
	/**
	 * The number of elements in the AVStream.side_data array.
	 */
	NbSideData ffcommon.FInt

	/**
	 * Flags indicating events happening on the stream, a combination of
	 * AVSTREAM_EVENT_FLAG_*.
	 *
	 * - demuxing: may be set by the demuxer in avformat_open_input(),
	 *   avformat_find_stream_info() and av_read_frame(). Flags must be cleared
	 *   by the user once the event has been handled.
	 * - muxing: may be set by the user after avformat_write_header(). to
	 *   indicate a user-triggered event.  The muxer will clear the flags for
	 *   events it has handled in av_[interleaved]_write_frame().
	 */
	EventFlags ffcommon.FInt
	/**
	 * - demuxing: the demuxer read new metadata from the file and updated
	 *     AVStream.metadata accordingly
	 * - muxing: the user updated AVStream.metadata and wishes the muxer to write
	 *     it into the file
	 */
	//const AVSTREAM_EVENT_FLAG_METADATA_UPDATED 0x0001//todo
	/**
	 * - demuxing: new packets for this stream were read from the file. This
	 *   event is informational only and does not guarantee that new packets
	 *   for this stream will necessarily be returned from av_read_frame().
	 */
	//const AVSTREAM_EVENT_FLAG_NEW_PACKETS (1 << 1)

	/**
	 * Real base framerate of the stream.
	 * This is the lowest framerate with which all timestamps can be
	 * represented accurately (it is the least common multiple of all
	 * framerates in the stream). Note, this value is just a guess!
	 * For example, if the time base is 1/90000 and all frames have either
	 * approximately 3600 or 1800 timer ticks, then r_frame_rate will be 50/1.
	 */
	RFrameRate AVRational

	//#if FF_API_LAVF_FFSERVER
	/**
	 * String containing pairs of key and values describing recommended encoder configuration.
	 * Pairs are separated by ','.
	 * Keys are separated from values by '='.
	 *
	 * @deprecated unused
	 */
	//attribute_deprecated
	//char *recommended_encoder_configuration;
	RecommendedEncoderConfiguration ffcommon.FCharPStruct
	//#endif

	/**
	 * Codec parameters associated with this stream. Allocated and freed by
	 * libavformat in avformat_new_stream() and avformat_free_context()
	 * respectively.
	 *
	 * - demuxing: filled by libavformat on stream creation or in
	 *             avformat_find_stream_info()
	 * - muxing: filled by the caller before avformat_write_header()
	 */
	Codecpar *AVCodecParameters

	/*****************************************************************
	 * All fields below this line are not part of the public API. They
	 * may not be used outside of libavformat and can be changed and
	 * removed at will.
	 * Internal note: be aware that physically removing these fields
	 * will break ABI. Replace removed fields with dummy fields, and
	 * add new fields to AVStreamInternal.
	 *****************************************************************
	 */

	//#if LIBAVFORMAT_VERSION_MAJOR < 59
	// kept for ABI compatibility only, do not access in any way
	Unused ffcommon.FVoidP
	//#endif

	PtsWrapBits ffcommon.FInt /**< number of bits in pts (used for wrapping control) */

	// Timestamp generation support:
	/**
	 * Timestamp corresponding to the last dts sync point.
	 *
	 * Initialized when AVCodecParserContext.dts_sync_point >= 0 and
	 * a DTS is received from the underlying container. Otherwise set to
	 * AV_NOPTS_VALUE by default.
	 */
	FirstDts       ffcommon.FInt64T
	CurDts         ffcommon.FInt64T
	LastIpPts      ffcommon.FInt64T
	LastIpDuration ffcommon.FInt

	/**
	 * Number of packets to buffer for codec probing
	 */
	ProbePackets ffcommon.FInt

	/**
	 * Number of frames that have been demuxed during avformat_find_stream_info()
	 */
	CodecInfoNbFrames ffcommon.FInt

	/* av_read_frame() support */
	NeedParsing AVStreamParseType
	Parser      *AVCodecParserContext

	//#if LIBAVFORMAT_VERSION_MAJOR < 59
	// kept for ABI compatibility only, do not access in any way
	Unused7 ffcommon.FVoidP
	Unused6 AVProbeData
	Unused5 [16 + 1]ffcommon.FInt64T
	//#endif
	IndexEntries *AVIndexEntry /**< Only used if the format does not
	  support seeking natively. */
	Nb_indexEntries           ffcommon.FInt
	IndexEntriesAllocatedSize ffcommon.FUnsignedInt

	/**
	 * Stream Identifier
	 * This is the MPEG-TS stream identifier +1
	 * 0 means unknown
	 */
	StreamIdentifier ffcommon.FInt

	//#if LIBAVFORMAT_VERSION_MAJOR < 59
	// kept for ABI compatibility only, do not access in any way
	Unused8  ffcommon.FInt
	Unused9  ffcommon.FInt
	Unused10 ffcommon.FInt
	//#endif

	/**
	 * An opaque field for libavformat internal usage.
	 * Must not be accessed in any way by callers.
	 */
	Internal *AVStreamInternal
}

//#if FF_API_FORMAT_GET_SET
/**
 * Accessors for some AVStream fields. These used to be provided for ABI
 * compatibility, and do not need to be used anymore.
 */
//attribute_deprecated
//AVRational av_stream_get_r_frame_rate(const AVStream *s);
func (s *AVStream) AvStreamGetRFrameRate() (res AVRational) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_stream_get_r_frame_rate").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = *(*AVRational)(unsafe.Pointer(t))
	return
}

//attribute_deprecated
//void       av_stream_set_r_frame_rate(AVStream *s, AVRational r);
func (s *AVStream) AvStreamSetRFrameRate(r AVRational) {
	ffcommon.GetAvformatDll().NewProc("av_stream_set_r_frame_rate").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(&r)),
	)
}

//#if FF_API_LAVF_FFSERVER
//attribute_deprecated
//char* av_stream_get_recommended_encoder_configuration(const AVStream *s);
func (s *AVStream) AvStreamGetRecommendedEncoderConfiguration() (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_stream_get_recommended_encoder_configuration").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = ffcommon.StringFromPtr(t)
	return
}

//attribute_deprecated
//void  av_stream_set_recommended_encoder_configuration(AVStream *s, char *configuration);
func (s *AVStream) AvStreamSetRecommendedEncoderConfiguration(configuration ffcommon.FCharP) {
	ffcommon.GetAvformatDll().NewProc("av_stream_set_recommended_encoder_configuration").Call(
		uintptr(unsafe.Pointer(s)),
		ffcommon.UintPtrFromString(configuration),
	)
}

//#endif
//#endif

//struct AVCodecParserContext *av_stream_get_parser(const AVStream *s);
func (s *AVStream) AvStreamGetParser() (res *AVCodecParserContext) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_stream_get_parser").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = (*AVCodecParserContext)(unsafe.Pointer(t))
	return
}

/**
 * Returns the pts of the last muxed packet + its duration
 *
 * the retuned value is undefined when used with a demuxer.
 */
//int64_t    av_stream_get_end_pts(const AVStream *st);
func (st *AVStream) AvStreamGetEndPts() (res ffcommon.FInt64T) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_stream_get_end_pts").Call(
		uintptr(unsafe.Pointer(st)),
	)
	res = ffcommon.FInt64T(t)
	return
}

const AV_PROGRAM_RUNNING = 1

/**
 * New fields can be added to the end with minor version bumps.
 * Removal, reordering and changes to existing fields require a major
 * version bump.
 * sizeof(AVProgram) must not be used outside libav*.
 */
type AVProgram struct {
	Id              ffcommon.FInt
	Flags           ffcommon.FInt
	Discard         AVDiscard ///< selects which program to discard and which to feed to the caller
	StreamIndex     *ffcommon.FUnsignedInt
	NbStreamIndexes ffcommon.FUnsignedInt
	Metadata        *AVDictionary

	ProgramNum ffcommon.FInt
	PmtPid     ffcommon.FInt
	PcrPid     ffcommon.FInt
	PmtVersion ffcommon.FInt

	/*****************************************************************
	 * All fields below this line are not part of the public API. They
	 * may not be used outside of libavformat and can be changed and
	 * removed at will.
	 * New public fields should be added right above.
	 *****************************************************************
	 */
	StartTime ffcommon.FInt64T
	EndTime   ffcommon.FInt64T

	PtsWrapReference ffcommon.FInt64T ///< reference dts for wrap detection
	PtsWrapBehavior  ffcommon.FInt    ///< behavior on wrap detection
}

const AVFMTCTX_NOHEADER = 0x0001 /**< signal that no header is present
  (streams are added dynamically) */
const AVFMTCTX_UNSEEKABLE = 0x0002 /**< signal that the stream is definitely
  not seekable, and attempts to call the
  seek function will fail. For some
  network protocols (e.g. HLS), this can
  change dynamically at runtime. */

type AVChapter struct {

	//#if FF_API_CHAPTER_ID_INT
	//int id;                 ///< unique ID to identify the chapter
	//#else
	//int64_t id;             ///< unique ID to identify the chapter
	//#endif
	Id         ffcommon.FIntOrInt64
	TimeBase   AVRational       ///< time base in which the start/end timestamps are specified
	Start, End ffcommon.FInt64T ///< chapter start/end time in time_base units
	Metadata   *AVDictionary
}

/**
 * Callback used by devices to communicate with application.
 */
//typedef int (*av_format_control_message)(struct AVFormatContext *s, int type,
//void *data, size_t data_size);
type AvFormatControlMessage = func(s AVFormatContext, type0 ffcommon.FInt, data ffcommon.FVoidP, data_size ffcommon.FSizeT) uintptr

//typedef int (*AVOpenCallback)(struct AVFormatContext *s, AVIOContext **pb, const char *url, int flags,
//const AVIOInterruptCB *int_cb, AVDictionary **options);
type AVOpenCallback = func(s *AVFormatContext, pb **AVIOContext, url ffcommon.FCharPStruct, flags ffcommon.FInt, int_cb *AVIOInterruptCB, options **AVDictionary) uintptr

/**
 * The duration of a video can be estimated through various ways, and this enum can be used
 * to know how the duration was estimated.
 */
type AVDurationEstimationMethod = int32

const (
	AVFMT_DURATION_FROM_PTS     = iota ///< Duration accurately estimated from PTSes
	AVFMT_DURATION_FROM_STREAM         ///< Duration estimated from a stream with a known duration
	AVFMT_DURATION_FROM_BITRATE        ///< Duration estimated from bitrate (less accurate)
)

//typedef struct AVFormatInternal AVFormatInternal;
type AVFormatInternal struct {
}

/**
 * Format I/O context.
 * New fields can be added to the end with minor version bumps.
 * Removal, reordering and changes to existing fields require a major
 * version bump.
 * sizeof(AVFormatContext) must not be used outside libav*, use
 * avformat_alloc_context() to create an AVFormatContext.
 *
 * Fields can be accessed through AVOptions (av_opt*),
 * the name string used matches the associated command line parameter name and
 * can be found in libavformat/options_table.h.
 * The AVOption/command line parameter names differ in some cases from the C
 * structure field names for historic reasons or brevity.
 */
type AVCodec = libavcodec.AVCodec
type AVFormatContext struct {

	/**
	 * A class for logging and @ref avoptions. Set by avformat_alloc_context().
	 * Exports (de)muxer private options if they exist.
	 */
	AvClass *AVClass

	/**
	 * The input container format.
	 *
	 * Demuxing only, set by avformat_open_input().
	 */
	Iformat *AVInputFormat

	/**
	 * The output container format.
	 *
	 * Muxing only, must be set by the caller before avformat_write_header().
	 */
	Oformat *AVOutputFormat

	/**
	 * Format private data. This is an AVOptions-enabled struct
	 * if and only if iformat/oformat.priv_class is not NULL.
	 *
	 * - muxing: set by avformat_write_header()
	 * - demuxing: set by avformat_open_input()
	 */
	PrivData ffcommon.FVoidP

	/**
	 * I/O context.
	 *
	 * - demuxing: either set by the user before avformat_open_input() (then
	 *             the user must close it manually) or set by avformat_open_input().
	 * - muxing: set by the user before avformat_write_header(). The caller must
	 *           take care of closing / freeing the IO context.
	 *
	 * Do NOT set this field if AVFMT_NOFILE flag is set in
	 * iformat/oformat.flags. In such a case, the (de)muxer will handle
	 * I/O in some other way and this field will be NULL.
	 */
	Pb *AVIOContext

	/* stream info */
	/**
	 * Flags signalling stream properties. A combination of AVFMTCTX_*.
	 * Set by libavformat.
	 */
	CtxFlags ffcommon.FInt

	/**
	 * Number of elements in AVFormatContext.streams.
	 *
	 * Set by avformat_new_stream(), must not be modified by any other code.
	 */
	NbStreams ffcommon.FUnsignedInt
	/**
	 * A list of all streams in the file. New streams are created with
	 * avformat_new_stream().
	 *
	 * - demuxing: streams are created by libavformat in avformat_open_input().
	 *             If AVFMTCTX_NOHEADER is set in ctx_flags, then new streams may also
	 *             appear in av_read_frame().
	 * - muxing: streams are created by the user before avformat_write_header().
	 *
	 * Freed by libavformat in avformat_free_context().
	 */
	Streams **AVStream

	//#if FF_API_FORMAT_FILENAME
	/**
	 * input or output filename
	 *
	 * - demuxing: set by avformat_open_input()
	 * - muxing: may be set by the caller before avformat_write_header()
	 *
	 * @deprecated Use url instead.
	 */
	//attribute_deprecated
	//char filename[1024];
	Filename [1024]ffcommon.FChar
	//#endif

	/**
	 * input or output URL. Unlike the old filename field, this field has no
	 * length restriction.
	 *
	 * - demuxing: set by avformat_open_input(), initialized to an empty
	 *             string if url parameter was NULL in avformat_open_input().
	 * - muxing: may be set by the caller before calling avformat_write_header()
	 *           (or avformat_init_output() if that is called first) to a string
	 *           which is freeable by av_free(). Set to an empty string if it
	 *           was NULL in avformat_init_output().
	 *
	 * Freed by libavformat in avformat_free_context().
	 */
	Url ffcommon.FCharPStruct

	/**
	 * Position of the first frame of the component, in
	 * AV_TIME_BASE fractional seconds. NEVER set this value directly:
	 * It is deduced from the AVStream values.
	 *
	 * Demuxing only, set by libavformat.
	 */
	StartTime ffcommon.FInt64T

	/**
	 * Duration of the stream, in AV_TIME_BASE fractional
	 * seconds. Only set this value if you know none of the individual stream
	 * durations and also do not set any of them. This is deduced from the
	 * AVStream values if not set.
	 *
	 * Demuxing only, set by libavformat.
	 */
	Duration ffcommon.FInt64T

	/**
	 * Total stream bitrate in bit/s, 0 if not
	 * available. Never set it directly if the file_size and the
	 * duration are known as FFmpeg can compute it automatically.
	 */
	BitRate ffcommon.FInt64T

	PacketSize ffcommon.FUnsignedInt
	MaxDelay   ffcommon.FInt

	/**
	 * Flags modifying the (de)muxer behaviour. A combination of AVFMT_FLAG_*.
	 * Set by the user before avformat_open_input() / avformat_write_header().
	 */
	Flags ffcommon.FInt

	// //todo
	//const AVFMT_FLAG_GENPTS  =     0x0001 ///< Generate missing pts even if it requires parsing future frames.
	//const AVFMT_FLAG_IGNIDX   =    0x0002 ///< Ignore index.
	//const AVFMT_FLAG_NONBLOCK  =   0x0004 ///< Do not block when reading packets from input.
	//const AVFMT_FLAG_IGNDTS     =  0x0008 ///< Ignore DTS on frames that contain both DTS & PTS
	//const AVFMT_FLAG_NOFILLIN   =  0x0010 ///< Do not infer any values from other values, just return what is stored in the container
	//const AVFMT_FLAG_NOPARSE    =  0x0020 ///< Do not use AVParsers, you also must set AVFMT_FLAG_NOFILLIN as the fillin code works on frames and no parsing -> no frames. Also seeking to frames can not work if parsing to find frame boundaries has been disabled
	//const AVFMT_FLAG_NOBUFFER   =  0x0040 ///< Do not buffer frames when possible
	//const AVFMT_FLAG_CUSTOM_IO  =  0x0080 ///< The caller has supplied a custom AVIOContext, don't avio_close() it.
	//const AVFMT_FLAG_DISCARD_CORRUPT = 0x0100 ///< Discard frames marked corrupted
	//const AVFMT_FLAG_FLUSH_PACKETS  =  0x0200 ///< Flush the AVIOContext every packet.
	///**
	// * When muxing, try to avoid writing any random/volatile data to the output.
	// * This includes any random IDs, real-time timestamps/dates, muxer version, etc.
	// *
	// * This flag is mainly intended for testing.
	// */
	//const AVFMT_FLAG_BITEXACT   =      0x0400
	//#if FF_API_LAVF_MP4A_LATM
	//const AVFMT_FLAG_MP4A_LATM =   0x8000 ///< Deprecated, does nothing.
	//#endif
	//const AVFMT_FLAG_SORT_DTS   = 0x10000 ///< try to interleave outputted packets by dts (using this flag can slow demuxing down)
	//#if FF_API_LAVF_PRIV_OPT
	//const AVFMT_FLAG_PRIV_OPT  =  0x20000 ///< Enable use of private options by delaying codec open (deprecated, will do nothing once av_demuxer_open() is removed)
	//#endif
	//#if FF_API_LAVF_KEEPSIDE_FLAG
	//const AVFMT_FLAG_KEEP_SIDE_DATA= 0x40000 ///< Deprecated, does nothing.
	//#endif
	//const AVFMT_FLAG_FAST_SEEK  = 0x80000 ///< Enable fast, but inaccurate seeks for some formats
	//const AVFMT_FLAG_SHORTEST  = 0x100000 ///< Stop muxing when the shortest stream stops.
	//const AVFMT_FLAG_AUTO_BSF  = 0x200000 ///< Add bitstream filters as requested by the muxer

	/**
	 * Maximum size of the data read from input for determining
	 * the input container format.
	 * Demuxing only, set by the caller before avformat_open_input().
	 */
	Probesize ffcommon.FInt64T

	/**
	 * Maximum duration (in AV_TIME_BASE units) of the data read
	 * from input in avformat_find_stream_info().
	 * Demuxing only, set by the caller before avformat_find_stream_info().
	 * Can be set to 0 to let avformat choose using a heuristic.
	 */
	MaxAnalyzeDuration ffcommon.FInt64T

	Key    *ffcommon.FUint8T
	Keylen ffcommon.FInt

	NbPrograms ffcommon.FUnsignedInt
	Programs   **AVProgram

	/**
	 * Forced video codec_id.
	 * Demuxing: Set by user.
	 */
	VideoCodecId AVCodecID

	/**
	 * Forced audio codec_id.
	 * Demuxing: Set by user.
	 */
	AudioCodecId AVCodecID

	/**
	 * Forced subtitle codec_id.
	 * Demuxing: Set by user.
	 */
	SubtitleCodecId AVCodecID

	/**
	 * Maximum amount of memory in bytes to use for the index of each stream.
	 * If the index exceeds this size, entries will be discarded as
	 * needed to maintain a smaller size. This can lead to slower or less
	 * accurate seeking (depends on demuxer).
	 * Demuxers for which a full in-memory index is mandatory will ignore
	 * this.
	 * - muxing: unused
	 * - demuxing: set by user
	 */
	MaxIndexSize ffcommon.FUnsignedInt

	/**
	 * Maximum amount of memory in bytes to use for buffering frames
	 * obtained from realtime capture devices.
	 */
	MaxPictureBuffer ffcommon.FUnsignedInt

	/**
	 * Number of chapters in AVChapter array.
	 * When muxing, chapters are normally written in the file header,
	 * so nb_chapters should normally be initialized before write_header
	 * is called. Some muxers (e.g. mov and mkv) can also write chapters
	 * in the trailer.  To write chapters in the trailer, nb_chapters
	 * must be zero when write_header is called and non-zero when
	 * write_trailer is called.
	 * - muxing: set by user
	 * - demuxing: set by libavformat
	 */
	NbChapters ffcommon.FUnsignedInt
	Chapters   **AVChapter

	/**
	 * Metadata that applies to the whole file.
	 *
	 * - demuxing: set by libavformat in avformat_open_input()
	 * - muxing: may be set by the caller before avformat_write_header()
	 *
	 * Freed by libavformat in avformat_free_context().
	 */
	Metadata *AVDictionary

	/**
	 * Start time of the stream in real world time, in microseconds
	 * since the Unix epoch (00:00 1st January 1970). That is, pts=0 in the
	 * stream was captured at this real world time.
	 * - muxing: Set by the caller before avformat_write_header(). If set to
	 *           either 0 or AV_NOPTS_VALUE, then the current wall-time will
	 *           be used.
	 * - demuxing: Set by libavformat. AV_NOPTS_VALUE if unknown. Note that
	 *             the value may become known after some number of frames
	 *             have been received.
	 */
	StartTimeRealtime ffcommon.FInt64T

	/**
	 * The number of frames used for determining the framerate in
	 * avformat_find_stream_info().
	 * Demuxing only, set by the caller before avformat_find_stream_info().
	 */
	FpsProbeSize ffcommon.FInt

	/**
	 * Error recognition; higher values will detect more errors but may
	 * misdetect some more or less valid parts as errors.
	 * Demuxing only, set by the caller before avformat_open_input().
	 */
	ErrorRecognition ffcommon.FInt

	/**
	 * Custom interrupt callbacks for the I/O layer.
	 *
	 * demuxing: set by the user before avformat_open_input().
	 * muxing: set by the user before avformat_write_header()
	 * (mainly useful for AVFMT_NOFILE formats). The callback
	 * should also be passed to avio_open2() if it's used to
	 * open the file.
	 */
	InterruptCallback AVIOInterruptCB

	/**
	 * Flags to enable debugging.
	 */
	Debug ffcommon.FInt
	//const FF_FDEBUG_TS    =    0x0001

	/**
	 * Maximum buffering duration for interleaving.
	 *
	 * To ensure all the streams are interleaved correctly,
	 * av_interleaved_write_frame() will wait until it has at least one packet
	 * for each stream before actually writing any packets to the output file.
	 * When some streams are "sparse" (i.e. there are large gaps between
	 * successive packets), this can result in excessive buffering.
	 *
	 * This field specifies the maximum difference between the timestamps of the
	 * first and the last packet in the muxing queue, above which libavformat
	 * will output a packet regardless of whether it has queued a packet for all
	 * the streams.
	 *
	 * Muxing only, set by the caller before avformat_write_header().
	 */
	MaxInterleaveDelta ffcommon.FInt64T

	/**
	 * Allow non-standard and experimental extension
	 * @see AVCodecContext.strict_std_compliance
	 */
	StrictStdCompliance ffcommon.FInt

	/**
	 * Flags indicating events happening on the file, a combination of
	 * AVFMT_EVENT_FLAG_*.
	 *
	 * - demuxing: may be set by the demuxer in avformat_open_input(),
	 *   avformat_find_stream_info() and av_read_frame(). Flags must be cleared
	 *   by the user once the event has been handled.
	 * - muxing: may be set by the user after avformat_write_header() to
	 *   indicate a user-triggered event.  The muxer will clear the flags for
	 *   events it has handled in av_[interleaved]_write_frame().
	 */
	EventFlags ffcommon.FInt
	/**
	 * - demuxing: the demuxer read new metadata from the file and updated
	 *   AVFormatContext.metadata accordingly
	 * - muxing: the user updated AVFormatContext.metadata and wishes the muxer to
	 *   write it into the file
	 */
	//const AVFMT_EVENT_FLAG_METADATA_UPDATED= 0x0001

	/**
	 * Maximum number of packets to read while waiting for the first timestamp.
	 * Decoding only.
	 */
	MaxTsProbe ffcommon.FInt

	/**
	 * Avoid negative timestamps during muxing.
	 * Any value of the AVFMT_AVOID_NEG_TS_* constants.
	 * Note, this only works when using av_interleaved_write_frame. (interleave_packet_per_dts is in use)
	 * - muxing: Set by user
	 * - demuxing: unused
	 */
	AvoidNegativeTs ffcommon.FInt
	//const AVFMT_AVOID_NEG_TS_AUTO      =       -1 ///< Enabled when required by target format
	//const AVFMT_AVOID_NEG_TS_MAKE_NON_NEGATIVE =1 ///< Shift timestamps so they are non negative
	//const AVFMT_AVOID_NEG_TS_MAKE_ZERO     =    2 ///< Shift timestamps so that they start at 0

	/**
	 * Transport stream id.
	 * This will be moved into demuxer private options. Thus no API/ABI compatibility
	 */
	TsId ffcommon.FInt

	/**
	 * Audio preload in microseconds.
	 * Note, not all formats support this and unpredictable things may happen if it is used when not supported.
	 * - encoding: Set by user
	 * - decoding: unused
	 */
	AudioPreload ffcommon.FInt

	/**
	 * Max chunk time in microseconds.
	 * Note, not all formats support this and unpredictable things may happen if it is used when not supported.
	 * - encoding: Set by user
	 * - decoding: unused
	 */
	MaxChunkDuration ffcommon.FInt

	/**
	 * Max chunk size in bytes
	 * Note, not all formats support this and unpredictable things may happen if it is used when not supported.
	 * - encoding: Set by user
	 * - decoding: unused
	 */
	MaxChunkSize ffcommon.FInt

	/**
	 * forces the use of wallclock timestamps as pts/dts of packets
	 * This has undefined results in the presence of B frames.
	 * - encoding: unused
	 * - decoding: Set by user
	 */
	UseWallclockAsTimestamps ffcommon.FInt

	/**
	 * avio flags, used to force AVIO_FLAG_DIRECT.
	 * - encoding: unused
	 * - decoding: Set by user
	 */
	AvioFlags ffcommon.FInt

	/**
	 * The duration field can be estimated through various ways, and this field can be used
	 * to know how the duration was estimated.
	 * - encoding: unused
	 * - decoding: Read by user
	 */
	DurationEstimationMethod AVDurationEstimationMethod

	/**
	 * Skip initial bytes when opening stream
	 * - encoding: unused
	 * - decoding: Set by user
	 */
	SkipInitialBytes ffcommon.FInt64T

	/**
	 * Correct single timestamp overflows
	 * - encoding: unused
	 * - decoding: Set by user
	 */
	CorrectTsOverflow ffcommon.FUnsignedInt

	/**
	 * Force seeking to any (also non key) frames.
	 * - encoding: unused
	 * - decoding: Set by user
	 */
	Seek2any ffcommon.FInt

	/**
	 * Flush the I/O context after each packet.
	 * - encoding: Set by user
	 * - decoding: unused
	 */
	FlushPackets ffcommon.FInt

	/**
	 * format probing score.
	 * The maximal score is AVPROBE_SCORE_MAX, its set when the demuxer probes
	 * the format.
	 * - encoding: unused
	 * - decoding: set by avformat, read by user
	 */
	ProbeScore ffcommon.FInt

	/**
	 * number of bytes to read maximally to identify format.
	 * - encoding: unused
	 * - decoding: set by user
	 */
	FormatProbesize ffcommon.FInt

	/**
	 * ',' separated list of allowed decoders.
	 * If NULL then all are allowed
	 * - encoding: unused
	 * - decoding: set by user
	 */
	CodecWhitelist ffcommon.FCharPStruct

	/**
	 * ',' separated list of allowed demuxers.
	 * If NULL then all are allowed
	 * - encoding: unused
	 * - decoding: set by user
	 */
	FormatWhitelist ffcommon.FCharPStruct

	/**
	 * An opaque field for libavformat internal usage.
	 * Must not be accessed in any way by callers.
	 */
	Internal *AVFormatInternal

	/**
	 * IO repositioned flag.
	 * This is set by avformat when the underlaying IO context read pointer
	 * is repositioned, for example when doing byte based seeking.
	 * Demuxers can use the flag to detect such changes.
	 */
	IoRepositioned ffcommon.FInt

	/**
	 * Forced video codec.
	 * This allows forcing a specific decoder, even when there are multiple with
	 * the same codec_id.
	 * Demuxing: Set by user
	 */
	VideoCodec *AVCodec

	/**
	 * Forced audio codec.
	 * This allows forcing a specific decoder, even when there are multiple with
	 * the same codec_id.
	 * Demuxing: Set by user
	 */
	AudioCodec *AVCodec

	/**
	 * Forced subtitle codec.
	 * This allows forcing a specific decoder, even when there are multiple with
	 * the same codec_id.
	 * Demuxing: Set by user
	 */
	SubtitleCodec *AVCodec

	/**
	 * Forced data codec.
	 * This allows forcing a specific decoder, even when there are multiple with
	 * the same codec_id.
	 * Demuxing: Set by user
	 */
	DataCodec *AVCodec

	/**
	 * Number of bytes to be written as padding in a metadata header.
	 * Demuxing: Unused.
	 * Muxing: Set by user via av_format_set_metadata_header_padding.
	 */
	MetadataHeaderPadding ffcommon.FInt

	/**
	 * User data.
	 * This is a place for some private data of the user.
	 */
	Opaque ffcommon.FVoidP

	/**
	 * Callback used by devices to communicate with application.
	 */
	//av_format_control_message control_message_cb;
	ControlMessageCb uintptr
	/**
	 * Output timestamp offset, in microseconds.
	 * Muxing: set by user
	 */
	OutputTsOffset ffcommon.FInt64T

	/**
	 * dump format separator.
	 * can be ", " or "\n      " or anything else
	 * - muxing: Set by user.
	 * - demuxing: Set by user.
	 */
	DumpSeparator *ffcommon.FUint8T

	/**
	 * Forced Data codec_id.
	 * Demuxing: Set by user.
	 */
	DataCodecId AVCodecID

	//#if FF_API_OLD_OPEN_CALLBACKS
	/**
	 * Called to open further IO contexts when needed for demuxing.
	 *
	 * This can be set by the user application to perform security checks on
	 * the URLs before opening them.
	 * The function should behave like avio_open2(), AVFormatContext is provided
	 * as contextual information and to reach AVFormatContext.opaque.
	 *
	 * If NULL then some simple checks are used together with avio_open2().
	 *
	 * Must not be accessed directly from outside avformat.
	 * @See av_format_set_open_cb()
	 *
	 * Demuxing: Set by user.
	 *
	 * @deprecated Use io_open and io_close.
	 */
	//attribute_deprecated
	//int (*open_cb)(struct AVFormatContext *s, AVIOContext **p, const char *url, int flags, const AVIOInterruptCB *int_cb, AVDictionary **options);
	OpenCb uintptr
	//#endif

	/**
	 * ',' separated list of allowed protocols.
	 * - encoding: unused
	 * - decoding: set by user
	 */
	//char *protocol_whitelist;
	ProtocolWhitelist ffcommon.FCharPStruct
	/**
	 * A callback for opening new IO streams.
	 *
	 * Whenever a muxer or a demuxer needs to open an IO stream (typically from
	 * avformat_open_input() for demuxers, but for certain formats can happen at
	 * other times as well), it will call this callback to obtain an IO context.
	 *
	 * @param s the format context
	 * @param pb on success, the newly opened IO context should be returned here
	 * @param url the url to open
	 * @param flags a combination of AVIO_FLAG_*
	 * @param options a dictionary of additional options, with the same
	 *                semantics as in avio_open2()
	 * @return 0 on success, a negative AVERROR code on failure
	 *
	 * @note Certain muxers and demuxers do nesting, i.e. they open one or more
	 * additional internal format contexts. Thus the AVFormatContext pointer
	 * passed to this callback may be different from the one facing the caller.
	 * It will, however, have the same 'opaque' field.
	 */
	//int (*io_open)(struct AVFormatContext *s, AVIOContext **pb, const char *url,
	//int flags, AVDictionary **options);
	IoOpen uintptr
	/**
	 * A callback for closing the streams opened with AVFormatContext.io_open().
	 */
	//void (*io_close)(struct AVFormatContext *s, AVIOContext *pb);
	IoClose uintptr
	/**
	 * ',' separated list of disallowed protocols.
	 * - encoding: unused
	 * - decoding: set by user
	 */
	ProtocolBlacklist ffcommon.FCharPStruct

	/**
	 * The maximum number of streams.
	 * - encoding: unused
	 * - decoding: set by user
	 */
	MaxStreams ffcommon.FInt

	/**
	 * Skip duration calcuation in estimate_timings_from_pts.
	 * - encoding: unused
	 * - decoding: set by user
	 */
	SkipEstimateDurationFromPts ffcommon.FInt

	/**
	 * Maximum number of packets that can be probed
	 * - encoding: unused
	 * - decoding: set by user
	 */
	MaxProbePackets ffcommon.FInt
}

func (s *AVFormatContext) GetStream(index ffcommon.FUnsignedInt) (res *AVStream) {
	t := uintptr(unsafe.Pointer(s.Streams)) + 8*uintptr(index)
	t = *(*uintptr)(unsafe.Pointer(t))
	res = (*AVStream)(unsafe.Pointer(t))
	return
}

//#if FF_API_FORMAT_GET_SET
/**
 * Accessors for some AVFormatContext fields. These used to be provided for ABI
 * compatibility, and do not need to be used anymore.
 */
//attribute_deprecated
//int av_format_get_probe_score(const AVFormatContext *s);
func (s *AVFormatContext) AvFormatGetProbeScore() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_format_get_probe_score").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = ffcommon.FInt(t)
	return
}

//attribute_deprecated
//AVCodec * av_format_get_video_codec(const AVFormatContext *s);
func (s *AVFormatContext) AvFormatGetVideoCodec() (res *AVCodec) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_format_get_video_codec").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = (*AVCodec)(unsafe.Pointer(t))
	return
}

//attribute_deprecated
//void      av_format_set_video_codec(AVFormatContext *s, AVCodec *c);
func (s *AVFormatContext) AvFormatSetVideoCodec(c *AVCodec) {
	ffcommon.GetAvformatDll().NewProc("av_format_set_video_codec").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(c)),
	)
}

//attribute_deprecated
//AVCodec * av_format_get_audio_codec(const AVFormatContext *s);
func (s *AVFormatContext) AvFormatGetAudioCodec() (res *AVCodec) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_format_get_audio_codec").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = (*AVCodec)(unsafe.Pointer(t))
	return
}

//attribute_deprecated
//void      av_format_set_audio_codec(AVFormatContext *s, AVCodec *c);
func (s *AVFormatContext) AvFormatSetAudioCodec(c *AVCodec) {
	ffcommon.GetAvformatDll().NewProc("av_format_set_audio_codec").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(c)),
	)
}

//attribute_deprecated
//AVCodec * av_format_get_subtitle_codec(const AVFormatContext *s);
func (s *AVFormatContext) AvFormatGetSubtitleCodec() (res *AVCodec) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_format_get_subtitle_codec").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = (*AVCodec)(unsafe.Pointer(t))
	return
}

//attribute_deprecated
//void      av_format_set_subtitle_codec(AVFormatContext *s, AVCodec *c);
func (s *AVFormatContext) AvFormatSetSubtitleCodec(c *AVCodec) {
	ffcommon.GetAvformatDll().NewProc("av_format_set_subtitle_codec").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(c)),
	)
}

//attribute_deprecated
//AVCodec * av_format_get_data_codec(const AVFormatContext *s);
func (s *AVFormatContext) AvFormatGetDataCodec() (res *AVCodec) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_format_get_data_codec").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = (*AVCodec)(unsafe.Pointer(t))
	return
}

//attribute_deprecated
//void      av_format_set_data_codec(AVFormatContext *s, AVCodec *c);
func (s *AVFormatContext) AvFormatSetDataCodec(c *AVCodec) {
	ffcommon.GetAvformatDll().NewProc("av_format_set_data_codec").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(c)),
	)
}

//attribute_deprecated
//int       av_format_get_metadata_header_padding(const AVFormatContext *s);
func (s *AVFormatContext) AvFormatGetMetadataHeaderPadding() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_format_get_metadata_header_padding").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = ffcommon.FInt(t)
	return
}

//attribute_deprecated
//void      av_format_set_metadata_header_padding(AVFormatContext *s, int c);
func (s *AVFormatContext) AvFormatSetMetadataHeaderPadding(c ffcommon.FInt) {
	ffcommon.GetAvformatDll().NewProc("av_format_set_metadata_header_padding").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(c),
	)
}

//attribute_deprecated
//void *    av_format_get_opaque(const AVFormatContext *s);
func (s *AVFormatContext) AvFormatGetOpaque() (res ffcommon.FVoidP) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_format_get_opaque").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = t
	return
}

//attribute_deprecated
//void      av_format_set_opaque(AVFormatContext *s, void *opaque);
func (s *AVFormatContext) AvFormatSetOpaque(opaque ffcommon.FVoidP) {
	ffcommon.GetAvformatDll().NewProc("av_format_set_opaque").Call(
		uintptr(unsafe.Pointer(s)),
		opaque,
	)
}

//attribute_deprecated
//av_format_control_message av_format_get_control_message_cb(const AVFormatContext *s);
func (s *AVFormatContext) AvFormatGetControlMessageCb() (res uintptr) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_format_get_control_message_cb").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = t
	return
}

//attribute_deprecated
//void      av_format_set_control_message_cb(AVFormatContext *s, av_format_control_message callback);
func (s *AVFormatContext) AvFormatSetControlMessageCb(callback uintptr) {
	ffcommon.GetAvformatDll().NewProc("av_format_set_control_message_cb").Call(
		uintptr(unsafe.Pointer(s)),
		callback,
	)
}

//#if FF_API_OLD_OPEN_CALLBACKS
//attribute_deprecated AVOpenCallback av_format_get_open_cb(const AVFormatContext *s);
func (s *AVFormatContext) AvFormatGetOpenCb() (res uintptr) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_format_get_open_cb").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = t
	return
}

//attribute_deprecated void av_format_set_open_cb(AVFormatContext *s, AVOpenCallback callback);
func (s *AVFormatContext) AvFormatSetOpenCb(callback AVOpenCallback) {
	ffcommon.GetAvformatDll().NewProc("av_format_set_open_cb").Call(
		uintptr(unsafe.Pointer(s)),
		ffcommon.NewCallback(callback),
	)
}

//#endif
//#endif

/**
 * This function will cause global side data to be injected in the next packet
 * of each stream as well as after any subsequent seek.
 */
//void av_format_inject_global_side_data(AVFormatContext *s);
func (s *AVFormatContext) AvFormatInjectGlobalSideData() {
	ffcommon.GetAvformatDll().NewProc("av_format_inject_global_side_data").Call(
		uintptr(unsafe.Pointer(s)),
	)
}

/**
 * Returns the method used to set ctx->duration.
 *
 * @return AVFMT_DURATION_FROM_PTS, AVFMT_DURATION_FROM_STREAM, or AVFMT_DURATION_FROM_BITRATE.
 */
//enum AVDurationEstimationMethod av_fmt_ctx_get_duration_estimation_method(const AVFormatContext* ctx);
func (ctx *AVFormatContext) AvFmtCtxGetDurationEstimationMethod() (res AVDurationEstimationMethod) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_fmt_ctx_get_duration_estimation_method").Call(
		uintptr(unsafe.Pointer(ctx)),
	)
	res = AVDurationEstimationMethod(t)
	return
}

/**
 * @defgroup lavf_core Core functions
 * @ingroup libavf
 *
 * Functions for querying libavformat capabilities, allocating core structures,
 * etc.
 * @{
 */

/**
 * Return the LIBAVFORMAT_VERSION_INT constant.
 */
//unsigned avformat_version(void);
func AvformatVersion() (res ffcommon.FUnsigned) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avformat_version").Call()
	res = ffcommon.FUnsigned(t)
	return
}

/**
 * Return the libavformat build-time configuration.
 */
//const char *avformat_configuration(void);
func AvformatConfiguration() (res ffcommon.FConstCharP) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avformat_configuration").Call()
	res = ffcommon.StringFromPtr(t)
	return
}

/**
 * Return the libavformat license.
 */
//const char *avformat_license(void);
func AvformatLicense() (res ffcommon.FConstCharP) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avformat_license").Call()
	res = ffcommon.StringFromPtr(t)
	return
}

//#if FF_API_NEXT
/**
 * Initialize libavformat and register all the muxers, demuxers and
 * protocols. If you do not call this function, then you can select
 * exactly which formats you want to support.
 *
 * @see av_register_input_format()
 * @see av_register_output_format()
 */
//attribute_deprecated
//void av_register_all(void);
func AvRegisterAll() {
	ffcommon.GetAvformatDll().NewProc("av_register_all").Call()
}

//attribute_deprecated
//void av_register_input_format(AVInputFormat *format);
func (format *AVInputFormat) AvRegisterInputFormat() {
	ffcommon.GetAvformatDll().NewProc("av_register_input_format").Call(
		uintptr(unsafe.Pointer(format)),
	)
}

//attribute_deprecated
//void av_register_output_format(AVOutputFormat *format);
func (format *AVOutputFormat) AvRegisterOutputFormat() {
	ffcommon.GetAvformatDll().NewProc("av_register_output_format").Call(
		uintptr(unsafe.Pointer(format)),
	)
}

//#endif

/**
 * Do global initialization of network libraries. This is optional,
 * and not recommended anymore.
 *
 * This functions only exists to work around thread-safety issues
 * with older GnuTLS or OpenSSL libraries. If libavformat is linked
 * to newer versions of those libraries, or if you do not use them,
 * calling this function is unnecessary. Otherwise, you need to call
 * this function before any other threads using them are started.
 *
 * This function will be deprecated once support for older GnuTLS and
 * OpenSSL libraries is removed, and this function has no purpose
 * anymore.
 */
//int avformat_network_init(void);
func AvformatNetworkInit() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avformat_network_init").Call()
	res = ffcommon.FInt(t)
	return
}

/**
 * Undo the initialization done by avformat_network_init. Call it only
 * once for each time you called avformat_network_init.
 */
//int avformat_network_deinit(void);
func AvformatNetworkDeinit() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avformat_network_deinit").Call()
	res = ffcommon.FInt(t)
	return
}

//#if FF_API_NEXT
/**
 * If f is NULL, returns the first registered input format,
 * if f is non-NULL, returns the next registered input format after f
 * or NULL if f is the last one.
 */
//attribute_deprecated
//AVInputFormat  *av_iformat_next(const AVInputFormat  *f);
func (f *AVInputFormat) AvIformatNext() (res *AVInputFormat) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_iformat_next").Call()
	res = (*AVInputFormat)(unsafe.Pointer(t))
	return
}

/**
 * If f is NULL, returns the first registered output format,
 * if f is non-NULL, returns the next registered output format after f
 * or NULL if f is the last one.
 */
//attribute_deprecated
//AVOutputFormat *av_oformat_next(const AVOutputFormat *f);
func (f *AVOutputFormat) AvOformatNext() (res *AVOutputFormat) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_oformat_next").Call()
	res = (*AVOutputFormat)(unsafe.Pointer(t))
	return
}

//#endif

/**
 * Iterate over all registered muxers.
 *
 * @param opaque a pointer where libavformat will store the iteration state. Must
 *               point to NULL to start the iteration.
 *
 * @return the next registered muxer or NULL when the iteration is
 *         finished
 */
//const AVOutputFormat *av_muxer_iterate(void **opaque);
func AvMuxerIterate(opaque *ffcommon.FVoidP) (res *AVOutputFormat) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_muxer_iterate").Call(
		uintptr(unsafe.Pointer(opaque)),
	)
	res = (*AVOutputFormat)(unsafe.Pointer(t))
	return
}

/**
 * Iterate over all registered demuxers.
 *
 * @param opaque a pointer where libavformat will store the iteration state. Must
 *               point to NULL to start the iteration.
 *
 * @return the next registered demuxer or NULL when the iteration is
 *         finished
 */
//const AVInputFormat *av_demuxer_iterate(void **opaque);
func AvDemuxerIterate(opaque *ffcommon.FVoidP) (res *AVInputFormat) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_demuxer_iterate").Call(
		uintptr(unsafe.Pointer(opaque)),
	)
	res = (*AVInputFormat)(unsafe.Pointer(t))
	return
}

/**
 * Allocate an AVFormatContext.
 * avformat_free_context() can be used to free the context and everything
 * allocated by the framework within it.
 */
//AVFormatContext *avformat_alloc_context(void);
func AvformatAllocContext() (res *AVFormatContext) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avformat_alloc_context").Call()
	res = (*AVFormatContext)(unsafe.Pointer(t))
	return
}

/**
 * Free an AVFormatContext and all its streams.
 * @param s context to free
 */
//void avformat_free_context(AVFormatContext *s);
func (s *AVFormatContext) AvformatFreeContext() {
	ffcommon.GetAvformatDll().NewProc("avformat_free_context").Call(
		uintptr(unsafe.Pointer(s)),
	)
}

/**
 * Get the AVClass for AVFormatContext. It can be used in combination with
 * AV_OPT_SEARCH_FAKE_OBJ for examining options.
 *
 * @see av_opt_find().
 */
//const AVClass *avformat_get_class(void);
func AvformatGetClass() (res *AVClass) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avformat_get_class").Call()
	res = (*AVClass)(unsafe.Pointer(t))
	return
}

/**
 * Add a new stream to a media file.
 *
 * When demuxing, it is called by the demuxer in read_header(). If the
 * flag AVFMTCTX_NOHEADER is set in s.ctx_flags, then it may also
 * be called in read_packet().
 *
 * When muxing, should be called by the user before avformat_write_header().
 *
 * User is required to call avcodec_close() and avformat_free_context() to
 * clean up the allocation by avformat_new_stream().
 *
 * @param s media file handle
 * @param c If non-NULL, the AVCodecContext corresponding to the new stream
 * will be initialized to use this codec. This is needed for e.g. codec-specific
 * defaults to be set, so codec should be provided if it is known.
 *
 * @return newly created stream or NULL on error.
 */
//AVStream *avformat_new_stream(AVFormatContext *s, const AVCodec *c);
func (s *AVFormatContext) AvformatNewStream(c *AVCodec) (res *AVStream) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avformat_new_stream").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(c)),
	)
	res = (*AVStream)(unsafe.Pointer(t))
	return
}

/**
 * Wrap an existing array as stream side data.
 *
 * @param st stream
 * @param type side information type
 * @param data the side data array. It must be allocated with the av_malloc()
 *             family of functions. The ownership of the data is transferred to
 *             st.
 * @param size side information size
 * @return zero on success, a negative AVERROR code on failure. On failure,
 *         the stream is unchanged and the data remains owned by the caller.
 */
//int av_stream_add_side_data(AVStream *st, enum AVPacketSideDataType type,
//uint8_t *data, size_t size);
type AVPacketSideDataType = libavcodec.AVPacketSideDataType

func (st *AVStream) AvStreamAddSideData(type0 AVPacketSideDataType, data *ffcommon.FUint8T, size ffcommon.FSizeT) (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_stream_add_side_data").Call(
		uintptr(unsafe.Pointer(st)),
		uintptr(type0),
		uintptr(unsafe.Pointer(data)),
		uintptr(size),
	)
	res = ffcommon.StringFromPtr(t)
	return
}

/**
 * Allocate new information from stream.
 *
 * @param stream stream
 * @param type desired side information type
 * @param size side information size
 * @return pointer to fresh allocated data or NULL otherwise
 */
//uint8_t *av_stream_new_side_data(AVStream *stream,
//#if FF_API_BUFFER_SIZE_T
//enum AVPacketSideDataType type, int size);
//#else
//enum AVPacketSideDataType type, size_t size);
//#endif
func (stream *AVStream) AvStreamNewSideData(type0 AVPacketSideDataType, size ffcommon.FIntOrSizeT) (res *ffcommon.FUint8T) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_stream_new_side_data").Call(
		uintptr(unsafe.Pointer(stream)),
		uintptr(type0),
		uintptr(size),
	)
	res = (*ffcommon.FUint8T)(unsafe.Pointer(t))
	return
}

/**
 * Get side information from stream.
 *
 * @param stream stream
 * @param type desired side information type
 * @param size If supplied, *size will be set to the size of the side data
 *             or to zero if the desired side data is not present.
 * @return pointer to data if present or NULL otherwise
 */
//uint8_t *av_stream_get_side_data(const AVStream *stream,
//#if FF_API_BUFFER_SIZE_T
//enum AVPacketSideDataType type, int *size);
//#else
//enum AVPacketSideDataType type, size_t *size);
//#endif
func (stream *AVStream) AvStreamGetSideData(type0 AVPacketSideDataType, size ffcommon.FIntOrSizeT) (res *ffcommon.FUint8T) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_stream_get_side_data").Call(
		uintptr(unsafe.Pointer(stream)),
		uintptr(type0),
		uintptr(size),
	)
	res = (*ffcommon.FUint8T)(unsafe.Pointer(t))
	return
}

//AVProgram *av_new_program(AVFormatContext *s, int id);
func (s *AVFormatContext) AvNewProgram(id ffcommon.FInt) (res *AVProgram) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_new_program").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(id),
	)
	res = (*AVProgram)(unsafe.Pointer(t))
	return
}

/**
 * @}
 */

/**
 * Allocate an AVFormatContext for an output format.
 * avformat_free_context() can be used to free the context and
 * everything allocated by the framework within it.
 *
 * @param *ctx is set to the created format context, or to NULL in
 * case of failure
 * @param oformat format to use for allocating the context, if NULL
 * format_name and filename are used instead
 * @param format_name the name of output format to use for allocating the
 * context, if NULL filename is used instead
 * @param filename the name of the filename to use for allocating the
 * context, may be NULL
 * @return >= 0 in case of success, a negative AVERROR code in case of
 * failure
 */
//int avformat_alloc_output_context2(AVFormatContext **ctx, ff_const59 AVOutputFormat *oformat,
//const char *format_name, const char *filename);
func AvformatAllocOutputContext2(ctx **AVFormatContext, oformat *AVOutputFormat,
	format_name, filename ffcommon.FConstCharP) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avformat_alloc_output_context2").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(oformat)),
		ffcommon.UintPtrFromString(format_name),
		ffcommon.UintPtrFromString(filename),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * @addtogroup lavf_decoding
 * @{
 */

/**
 * Find AVInputFormat based on the short name of the input format.
 */
//ff_const59 AVInputFormat *av_find_input_format(const char *short_name);
func AvFindInputFormat(short_name ffcommon.FConstCharP) (res *AVInputFormat) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_find_input_format").Call(
		ffcommon.UintPtrFromString(short_name),
	)
	res = (*AVInputFormat)(unsafe.Pointer(t))
	return
}

/**
 * Guess the file format.
 *
 * @param pd        data to be probed
 * @param is_opened Whether the file is already opened; determines whether
 *                  demuxers with or without AVFMT_NOFILE are probed.
 */
//ff_const59 AVInputFormat *av_probe_input_format(ff_const59 AVProbeData *pd, int is_opened);
func AvProbeInputFormat(short_name ffcommon.FConstCharP) (res *AVInputFormat) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_probe_input_format").Call(
		ffcommon.UintPtrFromString(short_name),
	)
	res = (*AVInputFormat)(unsafe.Pointer(t))
	return
}

/**
 * Guess the file format.
 *
 * @param pd        data to be probed
 * @param is_opened Whether the file is already opened; determines whether
 *                  demuxers with or without AVFMT_NOFILE are probed.
 * @param score_max A probe score larger that this is required to accept a
 *                  detection, the variable is set to the actual detection
 *                  score afterwards.
 *                  If the score is <= AVPROBE_SCORE_MAX / 4 it is recommended
 *                  to retry with a larger probe buffer.
 */
//ff_const59 AVInputFormat *av_probe_input_format2(ff_const59 AVProbeData *pd, int is_opened, int *score_max);
func (pd *AVProbeData) AvProbeInputFormat2(is_opened ffcommon.FInt, score_max *ffcommon.FInt) (res *AVInputFormat) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_probe_input_format2").Call(
		uintptr(unsafe.Pointer(pd)),
		uintptr(is_opened),
		uintptr(unsafe.Pointer(score_max)),
	)
	res = (*AVInputFormat)(unsafe.Pointer(t))
	return
}

/**
 * Guess the file format.
 *
 * @param is_opened Whether the file is already opened; determines whether
 *                  demuxers with or without AVFMT_NOFILE are probed.
 * @param score_ret The score of the best detection.
 */
//ff_const59 AVInputFormat *av_probe_input_format3(ff_const59 AVProbeData *pd, int is_opened, int *score_ret);
func (pd *AVProbeData) AvProbeInputFormat3(is_opened ffcommon.FInt, score_ret *ffcommon.FInt) (res *AVInputFormat) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_probe_input_format3").Call(
		uintptr(unsafe.Pointer(pd)),
		uintptr(is_opened),
		uintptr(unsafe.Pointer(score_ret)),
	)
	res = (*AVInputFormat)(unsafe.Pointer(t))
	return
}

/**
 * Probe a bytestream to determine the input format. Each time a probe returns
 * with a score that is too low, the probe buffer size is increased and another
 * attempt is made. When the maximum probe size is reached, the input format
 * with the highest score is returned.
 *
 * @param pb the bytestream to probe
 * @param fmt the input format is put here
 * @param url the url of the stream
 * @param logctx the log context
 * @param offset the offset within the bytestream to probe from
 * @param max_probe_size the maximum probe buffer size (zero for default)
 * @return the score in case of success, a negative value corresponding to an
 *         the maximal score is AVPROBE_SCORE_MAX
 * AVERROR code otherwise
 */
//int av_probe_input_buffer2(AVIOContext *pb, ff_const59 AVInputFormat **fmt,
//const char *url, void *logctx,
//unsigned int offset, unsigned int max_probe_size);
func (pb *AVIOContext) AvProbeInputBuffer2(fmt0 AVInputFormat,
	url ffcommon.FConstCharP, logctx ffcommon.FVoidP,
	offset, max_probe_size ffcommon.FUnsignedInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_probe_input_buffer2").Call(
		uintptr(unsafe.Sizeof(pb)),
		uintptr(unsafe.Sizeof(fmt0)),
		ffcommon.UintPtrFromString(url),
		logctx,
		uintptr(offset),
		uintptr(max_probe_size),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Like av_probe_input_buffer2() but returns 0 on success
 */
//int av_probe_input_buffer(AVIOContext *pb, ff_const59 AVInputFormat **fmt,
//const char *url, void *logctx,
//unsigned int offset, unsigned int max_probe_size);
func (pb *AVIOContext) AvProbeInputBuffer(fmt0 AVInputFormat,
	url ffcommon.FConstCharP, logctx ffcommon.FVoidP,
	offset, max_probe_size ffcommon.FUnsignedInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_probe_input_buffer").Call(
		uintptr(unsafe.Sizeof(pb)),
		uintptr(unsafe.Sizeof(fmt0)),
		ffcommon.UintPtrFromString(url),
		logctx,
		uintptr(offset),
		uintptr(max_probe_size),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Open an input stream and read the header. The codecs are not opened.
 * The stream must be closed with avformat_close_input().
 *
 * @param ps Pointer to user-supplied AVFormatContext (allocated by avformat_alloc_context).
 *           May be a pointer to NULL, in which case an AVFormatContext is allocated by this
 *           function and written into ps.
 *           Note that a user-supplied AVFormatContext will be freed on failure.
 * @param url URL of the stream to open.
 * @param fmt If non-NULL, this parameter forces a specific input format.
 *            Otherwise the format is autodetected.
 * @param options  A dictionary filled with AVFormatContext and demuxer-private options.
 *                 On return this parameter will be destroyed and replaced with a dict containing
 *                 options that were not found. May be NULL.
 *
 * @return 0 on success, a negative AVERROR on failure.
 *
 * @note If you want to use custom IO, preallocate the format context and set its pb field.
 */
//int avformat_open_input(AVFormatContext **ps, const char *url, ff_const59 AVInputFormat *fmt, AVDictionary **options);
func AvformatOpenInput(ps **AVFormatContext, url ffcommon.FConstCharP, fmt0 *AVInputFormat, options **AVDictionary) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avformat_open_input").Call(
		uintptr(unsafe.Pointer(ps)),
		ffcommon.UintPtrFromString(url),
		uintptr(unsafe.Pointer(fmt0)),
		uintptr(unsafe.Pointer(options)),
	)
	res = ffcommon.FInt(t)
	return
}

//#if FF_API_DEMUXER_OPEN
/**
 * @deprecated Use an AVDictionary to pass options to a demuxer.
 */
//attribute_deprecated
//int av_demuxer_open(AVFormatContext *ic);
//todo
func av_demuxer_open() (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_demuxer_open").Call()
	res = ffcommon.StringFromPtr(t)
	return
}

//#endif

/**
 * Read packets of a media file to get stream information. This
 * is useful for file formats with no headers such as MPEG. This
 * function also computes the real framerate in case of MPEG-2 repeat
 * frame mode.
 * The logical file position is not changed by this function;
 * examined packets may be buffered for later processing.
 *
 * @param ic media file handle
 * @param options  If non-NULL, an ic.nb_streams long array of pointers to
 *                 dictionaries, where i-th member contains options for
 *                 codec corresponding to i-th stream.
 *                 On return each dictionary will be filled with options that were not found.
 * @return >=0 if OK, AVERROR_xxx on error
 *
 * @note this function isn't guaranteed to open all the codecs, so
 *       options being non-empty at return is a perfectly normal behavior.
 *
 * @todo Let the user decide somehow what information is needed so that
 *       we do not waste time getting stuff the user does not need.
 */
//int avformat_find_stream_info(AVFormatContext *ic, AVDictionary **options);
func (ic *AVFormatContext) AvformatFindStreamInfo(options **AVDictionary) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avformat_find_stream_info").Call(
		uintptr(unsafe.Pointer(ic)),
		uintptr(unsafe.Pointer(options)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Find the programs which belong to a given stream.
 *
 * @param ic    media file handle
 * @param last  the last found program, the search will start after this
 *              program, or from the beginning if it is NULL
 * @param s     stream index
 * @return the next program which belongs to s, NULL if no program is found or
 *         the last program is not among the programs of ic.
 */
//AVProgram *av_find_program_from_stream(AVFormatContext *ic, AVProgram *last, int s);
func (ic *AVFormatContext) AvFindProgramFromStream(last *AVProgram, s ffcommon.FInt) (res *AVProgram) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_find_program_from_stream").Call(
		uintptr(unsafe.Pointer(ic)),
		uintptr(unsafe.Pointer(last)),
		uintptr(s),
	)
	res = (*AVProgram)(unsafe.Pointer(t))
	return
}

//void av_program_add_stream_index(AVFormatContext *ac, int progid, unsigned int idx);
func (ac *AVFormatContext) AvProgramAddStreamIndex(progid ffcommon.FInt, idx ffcommon.FUnsignedInt) {
	ffcommon.GetAvformatDll().NewProc("av_program_add_stream_index").Call(
		uintptr(unsafe.Pointer(ac)),
		uintptr(progid),
		uintptr(idx),
	)
}

/**
 * Find the "best" stream in the file.
 * The best stream is determined according to various heuristics as the most
 * likely to be what the user expects.
 * If the decoder parameter is non-NULL, av_find_best_stream will find the
 * default decoder for the stream's codec; streams for which no decoder can
 * be found are ignored.
 *
 * @param ic                media file handle
 * @param type              stream type: video, audio, subtitles, etc.
 * @param wanted_stream_nb  user-requested stream number,
 *                          or -1 for automatic selection
 * @param related_stream    try to find a stream related (eg. in the same
 *                          program) to this one, or -1 if none
 * @param decoder_ret       if non-NULL, returns the decoder for the
 *                          selected stream
 * @param flags             flags; none are currently defined
 * @return  the non-negative stream number in case of success,
 *          AVERROR_STREAM_NOT_FOUND if no stream with the requested type
 *          could be found,
 *          AVERROR_DECODER_NOT_FOUND if streams were found but no decoder
 * @note  If av_find_best_stream returns successfully and decoder_ret is not
 *        NULL, then *decoder_ret is guaranteed to be set to a valid AVCodec.
 */
//int av_find_best_stream(AVFormatContext *ic,
//enum AVMediaType type,
//int wanted_stream_nb,
//int related_stream,
//AVCodec **decoder_ret,
//int flags);
type AVMediaType = libavutil.AVMediaType

func (ic *AVFormatContext) AvFindBestStream(type0 AVMediaType, wanted_stream_nb,
	related_stream ffcommon.FInt,
	decoder_ret **AVCodec,
	flags ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_find_best_stream").Call(
		uintptr(unsafe.Pointer(ic)),
		uintptr(type0),
		uintptr(wanted_stream_nb),
		uintptr(related_stream),
		uintptr(unsafe.Pointer(decoder_ret)),
		uintptr(flags),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Return the next frame of a stream.
 * This function returns what is stored in the file, and does not validate
 * that what is there are valid frames for the decoder. It will split what is
 * stored in the file into frames and return one for each call. It will not
 * omit invalid data between valid frames so as to give the decoder the maximum
 * information possible for decoding.
 *
 * On success, the returned packet is reference-counted (pkt->buf is set) and
 * valid indefinitely. The packet must be freed with av_packet_unref() when
 * it is no longer needed. For video, the packet contains exactly one frame.
 * For audio, it contains an integer number of frames if each frame has
 * a known fixed size (e.g. PCM or ADPCM data). If the audio frames have
 * a variable size (e.g. MPEG audio), then it contains one frame.
 *
 * pkt->pts, pkt->dts and pkt->duration are always set to correct
 * values in AVStream.time_base units (and guessed if the format cannot
 * provide them). pkt->pts can be AV_NOPTS_VALUE if the video format
 * has B-frames, so it is better to rely on pkt->dts if you do not
 * decompress the payload.
 *
 * @return 0 if OK, < 0 on error or end of file. On error, pkt will be blank
 *         (as if it came from av_packet_alloc()).
 *
 * @note pkt will be initialized, so it may be uninitialized, but it must not
 *       contain data that needs to be freed.
 */
//int av_read_frame(AVFormatContext *s, AVPacket *pkt);
func (s *AVFormatContext) AvReadFrame(pkt *AVPacket) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_read_frame").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(pkt)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Seek to the keyframe at timestamp.
 * 'timestamp' in 'stream_index'.
 *
 * @param s media file handle
 * @param stream_index If stream_index is (-1), a default
 * stream is selected, and timestamp is automatically converted
 * from AV_TIME_BASE units to the stream specific time_base.
 * @param timestamp Timestamp in AVStream.time_base units
 *        or, if no stream is specified, in AV_TIME_BASE units.
 * @param flags flags which select direction and seeking mode
 * @return >= 0 on success
 */
//int av_seek_frame(AVFormatContext *s, int stream_index, int64_t timestamp,
//int flags);
func (s *AVFormatContext) AvSeekFrame(stream_index ffcommon.FInt, timestamp ffcommon.FInt64T, flags ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_seek_frame").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(stream_index),
		uintptr(timestamp),
		uintptr(flags),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Seek to timestamp ts.
 * Seeking will be done so that the point from which all active streams
 * can be presented successfully will be closest to ts and within min/max_ts.
 * Active streams are all streams that have AVStream.discard < AVDISCARD_ALL.
 *
 * If flags contain AVSEEK_FLAG_BYTE, then all timestamps are in bytes and
 * are the file position (this may not be supported by all demuxers).
 * If flags contain AVSEEK_FLAG_FRAME, then all timestamps are in frames
 * in the stream with stream_index (this may not be supported by all demuxers).
 * Otherwise all timestamps are in units of the stream selected by stream_index
 * or if stream_index is -1, in AV_TIME_BASE units.
 * If flags contain AVSEEK_FLAG_ANY, then non-keyframes are treated as
 * keyframes (this may not be supported by all demuxers).
 * If flags contain AVSEEK_FLAG_BACKWARD, it is ignored.
 *
 * @param s media file handle
 * @param stream_index index of the stream which is used as time base reference
 * @param min_ts smallest acceptable timestamp
 * @param ts target timestamp
 * @param max_ts largest acceptable timestamp
 * @param flags flags
 * @return >=0 on success, error code otherwise
 *
 * @note This is part of the new seek API which is still under construction.
 */
//int avformat_seek_file(AVFormatContext *s, int stream_index, int64_t min_ts, int64_t ts, int64_t max_ts, int flags);
func (s *AVFormatContext) AvformatSeekFile(stream_index ffcommon.FInt, min_ts, ts, max_ts ffcommon.FInt64T, flags ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avformat_seek_file").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(stream_index),
		uintptr(min_ts),
		uintptr(ts),
		uintptr(max_ts),
		uintptr(flags),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Discard all internally buffered data. This can be useful when dealing with
 * discontinuities in the byte stream. Generally works only with formats that
 * can resync. This includes headerless formats like MPEG-TS/TS but should also
 * work with NUT, Ogg and in a limited way AVI for example.
 *
 * The set of streams, the detected duration, stream parameters and codecs do
 * not change when calling this function. If you want a complete reset, it's
 * better to open a new AVFormatContext.
 *
 * This does not flush the AVIOContext (s->pb). If necessary, call
 * avio_flush(s->pb) before calling this function.
 *
 * @param s media file handle
 * @return >=0 on success, error code otherwise
 */
//int avformat_flush(AVFormatContext *s);
func (s *AVFormatContext) AvformatFlush() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avformat_flush").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Start playing a network-based stream (e.g. RTSP stream) at the
 * current position.
 */
//int av_read_play(AVFormatContext *s);
func (s *AVFormatContext) AvReadPlay() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_read_play").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Pause a network-based stream (e.g. RTSP stream).
 *
 * Use av_read_play() to resume it.
 */
//int av_read_pause(AVFormatContext *s);
func (s *AVFormatContext) AvReadPause() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_read_pause").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Close an opened input AVFormatContext. Free it and all its contents
 * and set *s to NULL.
 */
//void avformat_close_input(AVFormatContext **s);
func AvformatCloseInput(s **AVFormatContext) {
	ffcommon.GetAvformatDll().NewProc("avformat_close_input").Call(
		uintptr(unsafe.Pointer(s)),
	)
}

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

/**
 * Allocate the stream private data and write the stream header to
 * an output media file.
 *
 * @param s Media file handle, must be allocated with avformat_alloc_context().
 *          Its oformat field must be set to the desired output format;
 *          Its pb field must be set to an already opened AVIOContext.
 * @param options  An AVDictionary filled with AVFormatContext and muxer-private options.
 *                 On return this parameter will be destroyed and replaced with a dict containing
 *                 options that were not found. May be NULL.
 *
 * @return AVSTREAM_INIT_IN_WRITE_HEADER on success if the codec had not already been fully initialized in avformat_init,
 *         AVSTREAM_INIT_IN_INIT_OUTPUT  on success if the codec had already been fully initialized in avformat_init,
 *         negative AVERROR on failure.
 *
 * @see av_opt_find, av_dict_set, avio_open, av_oformat_next, avformat_init_output.
 */
//av_warn_unused_result
//int avformat_write_header(AVFormatContext *s, AVDictionary **options);
func (s *AVFormatContext) AvformatWriteHeader(options **AVDictionary) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avformat_write_header").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(options)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Allocate the stream private data and initialize the codec, but do not write the header.
 * May optionally be used before avformat_write_header to initialize stream parameters
 * before actually writing the header.
 * If using this function, do not pass the same options to avformat_write_header.
 *
 * @param s Media file handle, must be allocated with avformat_alloc_context().
 *          Its oformat field must be set to the desired output format;
 *          Its pb field must be set to an already opened AVIOContext.
 * @param options  An AVDictionary filled with AVFormatContext and muxer-private options.
 *                 On return this parameter will be destroyed and replaced with a dict containing
 *                 options that were not found. May be NULL.
 *
 * @return AVSTREAM_INIT_IN_WRITE_HEADER on success if the codec requires avformat_write_header to fully initialize,
 *         AVSTREAM_INIT_IN_INIT_OUTPUT  on success if the codec has been fully initialized,
 *         negative AVERROR on failure.
 *
 * @see av_opt_find, av_dict_set, avio_open, av_oformat_next, avformat_write_header.
 */
//av_warn_unused_result
//int avformat_init_output(AVFormatContext *s, AVDictionary **options);
func (s *AVFormatContext) AvformatInitOutput(options **AVDictionary) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avformat_init_output").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(options)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Write a packet to an output media file.
 *
 * This function passes the packet directly to the muxer, without any buffering
 * or reordering. The caller is responsible for correctly interleaving the
 * packets if the format requires it. Callers that want libavformat to handle
 * the interleaving should call av_interleaved_write_frame() instead of this
 * function.
 *
 * @param s media file handle
 * @param pkt The packet containing the data to be written. Note that unlike
 *            av_interleaved_write_frame(), this function does not take
 *            ownership of the packet passed to it (though some muxers may make
 *            an internal reference to the input packet).
 *            <br>
 *            This parameter can be NULL (at any time, not just at the end), in
 *            order to immediately flush data buffered within the muxer, for
 *            muxers that buffer up data internally before writing it to the
 *            output.
 *            <br>
 *            Packet's @ref AVPacket.stream_index "stream_index" field must be
 *            set to the index of the corresponding stream in @ref
 *            AVFormatContext.streams "s->streams".
 *            <br>
 *            The timestamps (@ref AVPacket.pts "pts", @ref AVPacket.dts "dts")
 *            must be set to correct values in the stream's timebase (unless the
 *            output format is flagged with the AVFMT_NOTIMESTAMPS flag, then
 *            they can be set to AV_NOPTS_VALUE).
 *            The dts for subsequent packets passed to this function must be strictly
 *            increasing when compared in their respective timebases (unless the
 *            output format is flagged with the AVFMT_TS_NONSTRICT, then they
 *            merely have to be nondecreasing).  @ref AVPacket.duration
 *            "duration") should also be set if known.
 * @return < 0 on error, = 0 if OK, 1 if flushed and there is no more data to flush
 *
 * @see av_interleaved_write_frame()
 */
//int av_write_frame(AVFormatContext *s, AVPacket *pkt);
func (s *AVFormatContext) AvWriteFrame(pkt *AVPacket) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_write_frame").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(pkt)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Write a packet to an output media file ensuring correct interleaving.
 *
 * This function will buffer the packets internally as needed to make sure the
 * packets in the output file are properly interleaved in the order of
 * increasing dts. Callers doing their own interleaving should call
 * av_write_frame() instead of this function.
 *
 * Using this function instead of av_write_frame() can give muxers advance
 * knowledge of future packets, improving e.g. the behaviour of the mp4
 * muxer for VFR content in fragmenting mode.
 *
 * @param s media file handle
 * @param pkt The packet containing the data to be written.
 *            <br>
 *            If the packet is reference-counted, this function will take
 *            ownership of this reference and unreference it later when it sees
 *            fit.
 *            The caller must not access the data through this reference after
 *            this function returns. If the packet is not reference-counted,
 *            libavformat will make a copy.
 *            <br>
 *            This parameter can be NULL (at any time, not just at the end), to
 *            flush the interleaving queues.
 *            <br>
 *            Packet's @ref AVPacket.stream_index "stream_index" field must be
 *            set to the index of the corresponding stream in @ref
 *            AVFormatContext.streams "s->streams".
 *            <br>
 *            The timestamps (@ref AVPacket.pts "pts", @ref AVPacket.dts "dts")
 *            must be set to correct values in the stream's timebase (unless the
 *            output format is flagged with the AVFMT_NOTIMESTAMPS flag, then
 *            they can be set to AV_NOPTS_VALUE).
 *            The dts for subsequent packets in one stream must be strictly
 *            increasing (unless the output format is flagged with the
 *            AVFMT_TS_NONSTRICT, then they merely have to be nondecreasing).
 *            @ref AVPacket.duration "duration") should also be set if known.
 *
 * @return 0 on success, a negative AVERROR on error. Libavformat will always
 *         take care of freeing the packet, even if this function fails.
 *
 * @see av_write_frame(), AVFormatContext.max_interleave_delta
 */
//int av_interleaved_write_frame(AVFormatContext *s, AVPacket *pkt);
func (s *AVFormatContext) AvInterleavedWriteFrame(pkt *AVPacket) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_interleaved_write_frame").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(pkt)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Write an uncoded frame to an output media file.
 *
 * The frame must be correctly interleaved according to the container
 * specification; if not, av_interleaved_write_uncoded_frame() must be used.
 *
 * See av_interleaved_write_uncoded_frame() for details.
 */
//int av_write_uncoded_frame(AVFormatContext *s, int stream_index,
//AVFrame *frame);
type AVFrame = libavutil.AVFrame

func (s *AVFormatContext) AvWriteUncodedFrame(stream_index ffcommon.FInt, frame *AVFrame) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_write_uncoded_frame").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(stream_index),
		uintptr(unsafe.Pointer(frame)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Write an uncoded frame to an output media file.
 *
 * If the muxer supports it, this function makes it possible to write an AVFrame
 * structure directly, without encoding it into a packet.
 * It is mostly useful for devices and similar special muxers that use raw
 * video or PCM data and will not serialize it into a byte stream.
 *
 * To test whether it is possible to use it with a given muxer and stream,
 * use av_write_uncoded_frame_query().
 *
 * The caller gives up ownership of the frame and must not access it
 * afterwards.
 *
 * @return  >=0 for success, a negative code on error
 */
//int av_interleaved_write_uncoded_frame(AVFormatContext *s, int stream_index,
//AVFrame *frame);
func (s *AVFormatContext) AvInterleavedWriteUncodedFrame(stream_index ffcommon.FInt, frame *AVFrame) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_interleaved_write_uncoded_frame").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(stream_index),
		uintptr(unsafe.Pointer(frame)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Test whether a muxer supports uncoded frame.
 *
 * @return  >=0 if an uncoded frame can be written to that muxer and stream,
 *          <0 if not
 */
//int av_write_uncoded_frame_query(AVFormatContext *s, int stream_index);
func (s *AVFormatContext) AvWriteUncodedFrameQuery(stream_index ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_write_uncoded_frame_query").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(stream_index),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Write the stream trailer to an output media file and free the
 * file private data.
 *
 * May only be called after a successful call to avformat_write_header.
 *
 * @param s media file handle
 * @return 0 if OK, AVERROR_xxx on error
 */
//int av_write_trailer(AVFormatContext *s);
func (s *AVFormatContext) AvWriteTrailer() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_write_trailer").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Return the output format in the list of registered output formats
 * which best matches the provided parameters, or return NULL if
 * there is no match.
 *
 * @param short_name if non-NULL checks if short_name matches with the
 * names of the registered formats
 * @param filename if non-NULL checks if filename terminates with the
 * extensions of the registered formats
 * @param mime_type if non-NULL checks if mime_type matches with the
 * MIME type of the registered formats
 */
//ff_const59 AVOutputFormat *av_guess_format(const char *short_name,
//const char *filename,
//const char *mime_type);
func AvGuessFormat(short_name, filename, mime_type ffcommon.FConstCharP) (res *AVOutputFormat) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_guess_format").Call(
		ffcommon.UintPtrFromString(short_name),
		ffcommon.UintPtrFromString(filename),
		ffcommon.UintPtrFromString(mime_type),
	)
	res = (*AVOutputFormat)(unsafe.Pointer(t))
	return
}

/**
 * Guess the codec ID based upon muxer and filename.
 */
//enum AVCodecID av_guess_codec(ff_const59 AVOutputFormat *fmt, const char *short_name,
//const char *filename, const char *mime_type,
//enum AVMediaType type);
func (fmt0 *AVOutputFormat) AvGuessCodec(short_name, filename, mime_type ffcommon.FCharP, type0 AVMediaType) (res AVCodecID) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_guess_codec").Call(
		uintptr(unsafe.Pointer(fmt0)),
		ffcommon.UintPtrFromString(short_name),
		ffcommon.UintPtrFromString(filename),
		ffcommon.UintPtrFromString(mime_type),
		uintptr(type0),
	)
	res = AVCodecID(t)
	return
}

/**
 * Get timing information for the data currently output.
 * The exact meaning of "currently output" depends on the format.
 * It is mostly relevant for devices that have an internal buffer and/or
 * work in real time.
 * @param s          media file handle
 * @param stream     stream in the media file
 * @param[out] dts   DTS of the last packet output for the stream, in stream
 *                   time_base units
 * @param[out] wall  absolute time when that packet whas output,
 *                   in microsecond
 * @return  0 if OK, AVERROR(ENOSYS) if the format does not support it
 * Note: some formats or devices may not allow to measure dts and wall
 * atomically.
 */
//int av_get_output_timestamp(struct AVFormatContext *s, int stream,
//int64_t *dts, int64_t *wall);
func (s *AVFormatContext) AvGetOutputTimestamp(stream ffcommon.FInt, dts, wall *ffcommon.FInt64T) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_get_output_timestamp").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(stream),
		uintptr(unsafe.Pointer(dts)),
		uintptr(unsafe.Pointer(wall)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * @}
 */

/**
 * @defgroup lavf_misc Utility functions
 * @ingroup libavf
 * @{
 *
 * Miscellaneous utility functions related to both muxing and demuxing
 * (or neither).
 */

/**
 * Send a nice hexadecimal dump of a buffer to the specified file stream.
 *
 * @param f The file stream pointer where the dump should be sent to.
 * @param buf buffer
 * @param size buffer size
 *
 * @see av_hex_dump_log, av_pkt_dump2, av_pkt_dump_log2
 */
//void av_hex_dump(FILE *f, const uint8_t *buf, int size);
func AvHexDump(f ffcommon.FFileP, buf *ffcommon.FUint8T, size ffcommon.FInt) {
	ffcommon.GetAvformatDll().NewProc("av_hex_dump").Call(
		f,
		uintptr(unsafe.Pointer(buf)),
		uintptr(size),
	)
}

/**
 * Send a nice hexadecimal dump of a buffer to the log.
 *
 * @param avcl A pointer to an arbitrary struct of which the first field is a
 * pointer to an AVClass struct.
 * @param level The importance level of the message, lower values signifying
 * higher importance.
 * @param buf buffer
 * @param size buffer size
 *
 * @see av_hex_dump, av_pkt_dump2, av_pkt_dump_log2
 */
//void av_hex_dump_log(void *avcl, int level, const uint8_t *buf, int size);
func AvHexDumpLog(avcl ffcommon.FVoidP, level ffcommon.FInt, buf *ffcommon.FUint8T, size ffcommon.FInt) {
	ffcommon.GetAvformatDll().NewProc("av_hex_dump_log").Call(
		avcl,
		uintptr(level),
		uintptr(unsafe.Pointer(buf)),
		uintptr(size),
	)
}

/**
 * Send a nice dump of a packet to the specified file stream.
 *
 * @param f The file stream pointer where the dump should be sent to.
 * @param pkt packet to dump
 * @param dump_payload True if the payload must be displayed, too.
 * @param st AVStream that the packet belongs to
 */
//void av_pkt_dump2(FILE *f, const AVPacket *pkt, int dump_payload, const AVStream *st);
func AvPktDump2(f ffcommon.FFileP, pkt *AVPacket, dump_payload ffcommon.FInt, st *AVStream) {
	ffcommon.GetAvformatDll().NewProc("av_pkt_dump2").Call(
		f,
		uintptr(unsafe.Pointer(pkt)),
		uintptr(dump_payload),
		uintptr(unsafe.Pointer(st)),
	)
}

/**
 * Send a nice dump of a packet to the log.
 *
 * @param avcl A pointer to an arbitrary struct of which the first field is a
 * pointer to an AVClass struct.
 * @param level The importance level of the message, lower values signifying
 * higher importance.
 * @param pkt packet to dump
 * @param dump_payload True if the payload must be displayed, too.
 * @param st AVStream that the packet belongs to
 */
//void av_pkt_dump_log2(void *avcl, int level, const AVPacket *pkt, int dump_payload,
//const AVStream *st);
func AvPktDumpLog2(avcl ffcommon.FVoidP, level ffcommon.FInt, pkt *AVPacket, dump_payload ffcommon.FInt, st *AVStream) {
	ffcommon.GetAvformatDll().NewProc("av_pkt_dump_log2").Call(
		avcl,
		uintptr(level),
		uintptr(unsafe.Pointer(pkt)),
		uintptr(dump_payload),
		uintptr(unsafe.Pointer(st)),
	)
}

/**
 * Get the AVCodecID for the given codec tag tag.
 * If no codec id is found returns AV_CODEC_ID_NONE.
 *
 * @param tags list of supported codec_id-codec_tag pairs, as stored
 * in AVInputFormat.codec_tag and AVOutputFormat.codec_tag
 * @param tag  codec tag to match to a codec ID
 */
//enum AVCodecID av_codec_get_id(const struct AVCodecTag * const *tags, unsigned int tag);
func AvCodecGetId(tags **AVCodecTag, tag ffcommon.FUnsignedInt) (res AVCodecID) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_codec_get_id").Call(
		uintptr(unsafe.Pointer(tags)),
		uintptr(tag),
	)
	res = AVCodecID(t)
	return
}

/**
 * Get the codec tag for the given codec id id.
 * If no codec tag is found returns 0.
 *
 * @param tags list of supported codec_id-codec_tag pairs, as stored
 * in AVInputFormat.codec_tag and AVOutputFormat.codec_tag
 * @param id   codec ID to match to a codec tag
 */
//unsigned int av_codec_get_tag(const struct AVCodecTag * const *tags, enum AVCodecID id);
func AvCodecGetTag(tags **AVCodecTag, id AVCodecID) (res ffcommon.FUnsignedInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_codec_get_tag").Call(
		uintptr(unsafe.Pointer(tags)),
		uintptr(id),
	)
	res = ffcommon.FUnsignedInt(t)
	return
}

/**
 * Get the codec tag for the given codec id.
 *
 * @param tags list of supported codec_id - codec_tag pairs, as stored
 * in AVInputFormat.codec_tag and AVOutputFormat.codec_tag
 * @param id codec id that should be searched for in the list
 * @param tag A pointer to the found tag
 * @return 0 if id was not found in tags, > 0 if it was found
 */
//int av_codec_get_tag2(const struct AVCodecTag * const *tags, enum AVCodecID id,
//unsigned int *tag);
func AvCodecGetTag2(tags **AVCodecTag, id AVCodecID, tag *ffcommon.FUnsignedInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_codec_get_tag2").Call(
		uintptr(unsafe.Pointer(tags)),
		uintptr(id),
		uintptr(unsafe.Pointer(tag)),
	)
	res = ffcommon.FInt(t)
	return
}

//int av_find_default_stream_index(AVFormatContext *s);
func (s *AVFormatContext) AvFindDefaultStreamIndex() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_find_default_stream_index").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Get the index for a specific timestamp.
 *
 * @param st        stream that the timestamp belongs to
 * @param timestamp timestamp to retrieve the index for
 * @param flags if AVSEEK_FLAG_BACKWARD then the returned index will correspond
 *                 to the timestamp which is <= the requested one, if backward
 *                 is 0, then it will be >=
 *              if AVSEEK_FLAG_ANY seek to any frame, only keyframes otherwise
 * @return < 0 if no such timestamp could be found
 */
//int av_index_search_timestamp(AVStream *st, int64_t timestamp, int flags);
func (s *AVStream) AvIndexSearchTimestamp(timestamp ffcommon.FInt64T, flags ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_index_search_timestamp").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(timestamp),
		uintptr(flags),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Add an index entry into a sorted list. Update the entry if the list
 * already contains it.
 *
 * @param timestamp timestamp in the time base of the given stream
 */
//int av_add_index_entry(AVStream *st, int64_t pos, int64_t timestamp,
//int size, int distance, int flags);
func (st *AVStream) AvAddIndexEntry(pos, timestamp ffcommon.FInt64T, size, distance, flags ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_add_index_entry").Call(
		uintptr(unsafe.Pointer(st)),
		uintptr(pos),
		uintptr(timestamp),
		uintptr(size),
		uintptr(distance),
		uintptr(flags),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Split a URL string into components.
 *
 * The pointers to buffers for storing individual components may be null,
 * in order to ignore that component. Buffers for components not found are
 * set to empty strings. If the port is not found, it is set to a negative
 * value.
 *
 * @param proto the buffer for the protocol
 * @param proto_size the size of the proto buffer
 * @param authorization the buffer for the authorization
 * @param authorization_size the size of the authorization buffer
 * @param hostname the buffer for the host name
 * @param hostname_size the size of the hostname buffer
 * @param port_ptr a pointer to store the port number in
 * @param path the buffer for the path
 * @param path_size the size of the path buffer
 * @param url the URL to split
 */
//void av_url_split(char *proto,         int proto_size,
//char *authorization, int authorization_size,
//char *hostname,      int hostname_size,
//int *port_ptr,
//char *path,          int path_size,
//const char *url);
func AvUrlSplit(proto ffcommon.FCharP, proto_size ffcommon.FInt,
	authorization ffcommon.FCharP, authorization_size ffcommon.FInt,
	hostname ffcommon.FCharP, hostname_size ffcommon.FInt,
	port_ptr *ffcommon.FInt,
	path0 ffcommon.FCharP, path_size ffcommon.FInt,
	url ffcommon.FCharP) {
	ffcommon.GetAvformatDll().NewProc("av_url_split").Call(
		ffcommon.UintPtrFromString(proto),
		uintptr(proto_size),
		ffcommon.UintPtrFromString(authorization),
		uintptr(authorization_size),
		ffcommon.UintPtrFromString(hostname),
		uintptr(hostname_size),
		uintptr(unsafe.Pointer(port_ptr)),
		ffcommon.UintPtrFromString(path0),
		uintptr(path_size),
		ffcommon.UintPtrFromString(url),
	)
}

/**
 * Print detailed information about the input or output format, such as
 * duration, bitrate, streams, container, programs, metadata, side data,
 * codec and time base.
 *
 * @param ic        the context to analyze
 * @param index     index of the stream to dump information about
 * @param url       the URL to print, such as source or destination file
 * @param is_output Select whether the specified context is an input(0) or output(1)
 */
//void av_dump_format(AVFormatContext *ic,
//int index,
//const char *url,
//int is_output);
func (ic *AVFormatContext) AvDumpFormat(index ffcommon.FInt, url ffcommon.FConstCharP, is_output ffcommon.FInt) {
	ffcommon.GetAvformatDll().NewProc("av_dump_format").Call(
		uintptr(unsafe.Pointer(ic)),
		uintptr(index),
		ffcommon.UintPtrFromString(url),
		uintptr(is_output),
	)
}

const AV_FRAME_FILENAME_FLAGS_MULTIPLE = 1 ///< Allow multiple %d

/**
 * Return in 'buf' the path with '%d' replaced by a number.
 *
 * Also handles the '%0nd' format where 'n' is the total number
 * of digits and '%%'.
 *
 * @param buf destination buffer
 * @param buf_size destination buffer size
 * @param path numbered sequence string
 * @param number frame number
 * @param flags AV_FRAME_FILENAME_FLAGS_*
 * @return 0 if OK, -1 on format error
 */
//int av_get_frame_filename2(char *buf, int buf_size,
//const char *path, int number, int flags);
func AvGetFrameFilename2(buf ffcommon.FCharP, buf_size ffcommon.FInt,
	path0 ffcommon.FCharP, number, flags ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_get_frame_filename2").Call(
		ffcommon.UintPtrFromString(buf),
		uintptr(buf_size),
		ffcommon.UintPtrFromString(path0),
		uintptr(number),
		uintptr(flags),
	)
	res = ffcommon.FInt(t)
	return
}

//int av_get_frame_filename(char *buf, int buf_size,
//const char *path, int number);
func AvGetFrameFilename(buf ffcommon.FCharP, buf_size ffcommon.FInt,
	path0 ffcommon.FCharP, number ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_get_frame_filename").Call(
		ffcommon.UintPtrFromString(buf),
		uintptr(buf_size),
		ffcommon.UintPtrFromString(path0),
		uintptr(number),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Check whether filename actually is a numbered sequence generator.
 *
 * @param filename possible numbered sequence string
 * @return 1 if a valid numbered sequence string, 0 otherwise
 */
//int av_filename_number_test(const char *filename);
func AvFlenameNumberTest(filename ffcommon.FCharP) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_filename_number_test").Call(
		ffcommon.UintPtrFromString(filename),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Generate an SDP for an RTP session.
 *
 * Note, this overwrites the id values of AVStreams in the muxer contexts
 * for getting unique dynamic payload types.
 *
 * @param ac array of AVFormatContexts describing the RTP streams. If the
 *           array is composed by only one context, such context can contain
 *           multiple AVStreams (one AVStream per RTP stream). Otherwise,
 *           all the contexts in the array (an AVCodecContext per RTP stream)
 *           must contain only one AVStream.
 * @param n_files number of AVCodecContexts contained in ac
 * @param buf buffer where the SDP will be stored (must be allocated by
 *            the caller)
 * @param size the size of the buffer
 * @return 0 if OK, AVERROR_xxx on error
 */
//int av_sdp_create(AVFormatContext *ac[], int n_files, char *buf, int size);
func AvSdpCreate(ac **AVFormatContext, n_files ffcommon.FInt, buf ffcommon.FCharP, size ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_sdp_create").Call(
		uintptr(unsafe.Pointer(ac)),
		uintptr(n_files),
		ffcommon.UintPtrFromString(buf),
		uintptr(size),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Return a positive value if the given filename has one of the given
 * extensions, 0 otherwise.
 *
 * @param filename   file name to check against the given extensions
 * @param extensions a comma-separated list of filename extensions
 */
//int av_match_ext(const char *filename, const char *extensions);
func AvMatchExt(filename, extensions ffcommon.FConstCharP) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_match_ext").Call(
		ffcommon.UintPtrFromString(filename),
		ffcommon.UintPtrFromString(extensions),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Test if the given container can store a codec.
 *
 * @param ofmt           container to check for compatibility
 * @param codec_id       codec to potentially store in container
 * @param std_compliance standards compliance level, one of FF_COMPLIANCE_*
 *
 * @return 1 if codec with ID codec_id can be stored in ofmt, 0 if it cannot.
 *         A negative number if this information is not available.
 */
//int avformat_query_codec(const AVOutputFormat *ofmt, enum AVCodecID codec_id,
//int std_compliance);
func (ofmt *AVOutputFormat) AvformatQueryCodec(codec_id AVCodecID, std_compliance ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avformat_query_codec").Call(
		uintptr(unsafe.Pointer(ofmt)),
		uintptr(codec_id),
		uintptr(std_compliance),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * @defgroup riff_fourcc RIFF FourCCs
 * @{
 * Get the tables mapping RIFF FourCCs to libavcodec AVCodecIDs. The tables are
 * meant to be passed to av_codec_get_id()/av_codec_get_tag() as in the
 * following code:
 * @code
 * uint32_t tag = MKTAG('H', '2', '6', '4');
 * const struct AVCodecTag *table[] = { avformat_get_riff_video_tags(), 0 };
 * enum AVCodecID id = av_codec_get_id(table, tag);
 * @endcode
 */
/**
 * @return the table mapping RIFF FourCCs for video to libavcodec AVCodecID.
 */
//const struct AVCodecTag *avformat_get_riff_video_tags(void);
func AvformatGetRiffVideoTags() (res *AVCodecTag) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avformat_get_riff_video_tags").Call()
	res = (*AVCodecTag)(unsafe.Pointer(t))
	return
}

/**
 * @return the table mapping RIFF FourCCs for audio to AVCodecID.
 */
//const struct AVCodecTag *avformat_get_riff_audio_tags(void);
func AvformatGetRiffAudioTags() (res *AVCodecTag) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avformat_get_riff_audio_tags").Call()
	res = (*AVCodecTag)(unsafe.Pointer(t))
	return
}

/**
 * @return the table mapping MOV FourCCs for video to libavcodec AVCodecID.
 */
//const struct AVCodecTag *avformat_get_mov_video_tags(void);
func AvformatGetMovVideoTags() (res *AVCodecTag) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avformat_get_mov_video_tags").Call()
	res = (*AVCodecTag)(unsafe.Pointer(t))
	return
}

/**
 * @return the table mapping MOV FourCCs for audio to AVCodecID.
 */
//const struct AVCodecTag *avformat_get_mov_audio_tags(void);
func AvformatGetMovAudioTags() (res *AVCodecTag) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avformat_get_mov_audio_tags").Call()
	res = (*AVCodecTag)(unsafe.Pointer(t))
	return
}

/**
 * @}
 */

/**
 * Guess the sample aspect ratio of a frame, based on both the stream and the
 * frame aspect ratio.
 *
 * Since the frame aspect ratio is set by the codec but the stream aspect ratio
 * is set by the demuxer, these two may not be equal. This function tries to
 * return the value that you should use if you would like to display the frame.
 *
 * Basic logic is to use the stream aspect ratio if it is set to something sane
 * otherwise use the frame aspect ratio. This way a container setting, which is
 * usually easy to modify can override the coded value in the frames.
 *
 * @param format the format context which the stream is part of
 * @param stream the stream which the frame is part of
 * @param frame the frame with the aspect ratio to be determined
 * @return the guessed (valid) sample_aspect_ratio, 0/1 if no idea
 */
//AVRational av_guess_sample_aspect_ratio(AVFormatContext *format, AVStream *stream, AVFrame *frame);
func (format *AVFormatContext) AvGuessSampleAspectRatio(stream *AVStream, frame *AVFrame) (res AVRational) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_guess_sample_aspect_ratio").Call(
		uintptr(unsafe.Pointer(format)),
		uintptr(unsafe.Pointer(stream)),
		uintptr(unsafe.Pointer(frame)),
	)
	res = *(*AVRational)(unsafe.Pointer(t))
	return
}

/**
 * Guess the frame rate, based on both the container and codec information.
 *
 * @param ctx the format context which the stream is part of
 * @param stream the stream which the frame is part of
 * @param frame the frame for which the frame rate should be determined, may be NULL
 * @return the guessed (valid) frame rate, 0/1 if no idea
 */
//AVRational av_guess_frame_rate(AVFormatContext *ctx, AVStream *stream, AVFrame *frame);
func (ctx *AVFormatContext) AvGuessFrameRate(stream *AVStream, frame *AVFrame) (res AVRational) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_guess_frame_rate").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(stream)),
		uintptr(unsafe.Pointer(frame)),
	)
	res = *(*AVRational)(unsafe.Pointer(t))
	return
}

/**
 * Check if the stream st contained in s is matched by the stream specifier
 * spec.
 *
 * See the "stream specifiers" chapter in the documentation for the syntax
 * of spec.
 *
 * @return  >0 if st is matched by spec;
 *          0  if st is not matched by spec;
 *          AVERROR code if spec is invalid
 *
 * @note  A stream specifier can match several streams in the format.
 */
//int avformat_match_stream_specifier(AVFormatContext *s, AVStream *st,
//const char *spec);
func (s *AVFormatContext) AvformatMatchStreamSpecifier(st *AVStream, spec ffcommon.FConstCharP) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avformat_match_stream_specifier").Call(
		uintptr(unsafe.Pointer(s)),
		uintptr(unsafe.Pointer(st)),
		ffcommon.UintPtrFromString(spec),
	)
	res = ffcommon.FInt(t)
	return
}

//int avformat_queue_attached_pictures(AVFormatContext *s);
func (s *AVFormatContext) AvformatQueueAttachedPictures() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avformat_queue_attached_pictures").Call(
		uintptr(unsafe.Pointer(s)),
	)
	res = ffcommon.FInt(t)
	return
}

//#if FF_API_OLD_BSF
/**
 * Apply a list of bitstream filters to a packet.
 *
 * @param codec AVCodecContext, usually from an AVStream
 * @param pkt the packet to apply filters to. If, on success, the returned
 *        packet has size == 0 and side_data_elems == 0, it indicates that
 *        the packet should be dropped
 * @param bsfc a NULL-terminated list of filters to apply
 * @return  >=0 on success;
 *          AVERROR code on failure
 */
//attribute_deprecated
//int av_apply_bitstream_filters(AVCodecContext *codec, AVPacket *pkt,
//AVBitStreamFilterContext *bsfc);
func AvApplyBitstreamFilters(codec *AVCodecContext, pkt *AVPacket, bsfc *libavcodec.AVBitStreamFilterContext) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_apply_bitstream_filters").Call(
		uintptr(unsafe.Pointer(codec)),
		uintptr(unsafe.Pointer(pkt)),
		uintptr(unsafe.Pointer(bsfc)),
	)
	res = ffcommon.FInt(t)
	return
}

//#endif
type AVTimebaseSource int32

const (
	AVFMT_TBCF_AUTO = iota - 1
	AVFMT_TBCF_DECODER
	AVFMT_TBCF_DEMUXER
	//#if FF_API_R_FRAME_RATE
	AVFMT_TBCF_R_FRAMERATE

//#endif
)

/**
 * Transfer internal timing information from one stream to another.
 *
 * This function is useful when doing stream copy.
 *
 * @param ofmt     target output format for ost
 * @param ost      output stream which needs timings copy and adjustments
 * @param ist      reference input stream to copy timings from
 * @param copy_tb  define from where the stream codec timebase needs to be imported
 */
//int avformat_transfer_internal_stream_timing_info(const AVOutputFormat *ofmt,
//AVStream *ost, const AVStream *ist,
//enum AVTimebaseSource copy_tb);
func (ofmt *AVOutputFormat) AvformatTransferInternalStreamTimingInfo(ost, ist *AVStream, copy_tb AVTimebaseSource) (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("avformat_transfer_internal_stream_timing_info").Call(
		uintptr(unsafe.Pointer(ofmt)),
		uintptr(unsafe.Pointer(ost)),
		uintptr(unsafe.Pointer(ist)),
		uintptr(copy_tb),
	)
	res = ffcommon.StringFromPtr(t)
	return
}

/**
 * Get the internal codec timebase from a stream.
 *
 * @param st  input stream to extract the timebase from
 */
//AVRational av_stream_get_codec_timebase(const AVStream *st);
func (st *AVStream) AvStreamGetCodecTimebase() (res AVRational) {
	t, _, _ := ffcommon.GetAvformatDll().NewProc("av_stream_get_codec_timebase").Call(
		uintptr(unsafe.Pointer(st)),
	)
	res = *(*AVRational)(unsafe.Pointer(t))
	return
}

/**
 * @}
 */

//#endif /* AVFORMAT_AVFORMAT_H */
