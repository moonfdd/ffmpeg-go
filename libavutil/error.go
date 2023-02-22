package libavutil

import (
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
)

/*
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

/**
 * @file
 * error code definitions
 */

//#ifndef AVUTIL_ERROR_H
//#define AVUTIL_ERROR_H
//
//#include <errno.h>
//#include <stddef.h>

/**
 * @addtogroup lavu_error
 *
 * @{
 */

/* error handling */
//#if EDOM > 0
//#define AVERROR(e) (-(e))   ///< Returns a negative error code from a POSIX error code, to return from library functions.
//#define AVUNERROR(e) (-(e)) ///< Returns a POSIX error code from a library function error return value.
//#else
///* Some platforms have E* and errno already negated. */
//#define AVERROR(e) (e)
//#define AVUNERROR(e) (e)
//#endif
//
//#define FFERRTAG(a, b, c, d) (-(int)MKTAG(a, b, c, d))
//
//#define AVERROR_BSF_NOT_FOUND      FFERRTAG(0xF8,'B','S','F') ///< Bitstream filter not found
const AVERROR_BSF_NOT_FOUND = -(0xF8 | ('B' << 8) | ('S' << 16) | ('F' << 24))

//#define AVERROR_BUG                FFERRTAG( 'B','U','G','!') ///< Internal bug, also see AVERROR_BUG2
const AVERROR_BUG = -('B' | ('U' << 8) | ('G' << 16) | ('!' << 24))

//#define AVERROR_BUFFER_TOO_SMALL   FFERRTAG( 'B','U','F','S') ///< Buffer too small
const AVERROR_BUFFER_TOO_SMALL = -('B' | ('U' << 8) | ('F' << 16) | ('S' << 24))

//#define AVERROR_DECODER_NOT_FOUND  FFERRTAG(0xF8,'D','E','C') ///< Decoder not found
const AVERROR_DECODER_NOT_FOUND = -(0xF8 | ('D' << 8) | ('E' << 16) | ('C' << 24))

//#define AVERROR_DEMUXER_NOT_FOUND  FFERRTAG(0xF8,'D','E','M') ///< Demuxer not found
const AVERROR_DEMUXER_NOT_FOUND = -(0xF8 | ('D' << 8) | ('E' << 16) | ('M' << 24))

//#define AVERROR_ENCODER_NOT_FOUND  FFERRTAG(0xF8,'E','N','C') ///< Encoder not found
const AVERROR_ENCODER_NOT_FOUND = -(0xF8 | ('E' << 8) | ('N' << 16) | ('C' << 24))

//#define AVERROR_EOF                FFERRTAG( 'E','O','F',' ') ///< End of file
const AVERROR_EOF = -('E' | ('O' << 8) | ('F' << 16) | (' ' << 24))

//#define AVERROR_EXIT               FFERRTAG( 'E','X','I','T') ///< Immediate exit was requested; the called function should not be restarted
const AVERROR_EXIT = -('E' | ('X' << 8) | ('I' << 16) | ('T' << 24))

//#define AVERROR_EXTERNAL           FFERRTAG( 'E','X','T',' ') ///< Generic error in an external library
const AVERROR_EXTERNAL = -('E' | ('X' << 8) | ('T' << 16) | (' ' << 24))

//#define AVERROR_FILTER_NOT_FOUND   FFERRTAG(0xF8,'F','I','L') ///< Filter not found
const AVERROR_FILTER_NOT_FOUND = -(0xF8 | 'F'<<8 | 'I'<<16 | 'L'<<24)

//#define AVERROR_INVALIDDATA        FFERRTAG( 'I','N','D','A') ///< Invalid data found when processing input
const AVERROR_INVALIDDATA = -('I' | 'N'<<8 | 'D'<<16 | 'A'<<24)

//#define AVERROR_MUXER_NOT_FOUND    FFERRTAG(0xF8,'M','U','X') ///< Muxer not found
const AVERROR_MUXER_NOT_FOUND = -(0xF8 | 'M'<<8 | 'U'<<16 | 'X'<<24)

//#define AVERROR_OPTION_NOT_FOUND   FFERRTAG(0xF8,'O','P','T') ///< Option not found
const AVERROR_OPTION_NOT_FOUND = -(0xF8 | 'O'<<8 | 'P'<<16 | 'T'<<24)

//#define AVERROR_PATCHWELCOME       FFERRTAG( 'P','A','W','E') ///< Not yet implemented in FFmpeg, patches welcome
const AVERROR_PATCHWELCOME = -('P' | 'A'<<8 | 'W'<<16 | 'E'<<24)

//#define AVERROR_PROTOCOL_NOT_FOUND FFERRTAG(0xF8,'P','R','O') ///< Protocol not found
const AVERROR_PROTOCOL_NOT_FOUND = -(0xF8 | 'P'<<8 | 'R'<<16 | 'O'<<24)

//
//#define AVERROR_STREAM_NOT_FOUND   FFERRTAG(0xF8,'S','T','R') ///< Stream not found
const AVERROR_STREAM_NOT_FOUND = -(0xF8 | 'S'<<8 | 'T'<<16 | 'R'<<24)

///**
// * This is semantically identical to AVERROR_BUG
// * it has been introduced in Libav after our AVERROR_BUG and with a modified value.
// */
//#define AVERROR_BUG2               FFERRTAG( 'B','U','G',' ')
const AVERROR_BUG2 = -('B' | 'U'<<8 | 'G'<<16 | ' '<<24)

//#define AVERROR_UNKNOWN            FFERRTAG( 'U','N','K','N') ///< Unknown error, typically from an external library
const AVERROR_UNKNOWN = -('U' | 'N'<<8 | 'K'<<16 | 'N'<<24)

//#define AVERROR_EXPERIMENTAL       (-0x2bb2afa8) ///< Requested feature is flagged experimental. Set strict_std_compliance if you really want to use it.
const AVERROR_EXPERIMENTAL = (-0x2bb2afa8)

//#define AVERROR_INPUT_CHANGED      (-0x636e6701) ///< Input changed between calls. Reconfiguration is required. (can be OR-ed with AVERROR_OUTPUT_CHANGED)
const AVERROR_INPUT_CHANGED = (-0x636e6701)

//#define AVERROR_OUTPUT_CHANGED     (-0x636e6702) ///< Output changed between calls. Reconfiguration is required. (can be OR-ed with AVERROR_INPUT_CHANGED)
const AVERROR_OUTPUT_CHANGED = (-0x636e6702)

///* HTTP & RTSP errors */
//#define AVERROR_HTTP_BAD_REQUEST   FFERRTAG(0xF8,'4','0','0')
const AVERROR_HTTP_BAD_REQUEST = -(0xF8 | '4'<<8 | '0'<<16 | '0'<<24)

//#define AVERROR_HTTP_UNAUTHORIZED  FFERRTAG(0xF8,'4','0','1')
const AVERROR_HTTP_UNAUTHORIZED = -(0xF8 | '4'<<8 | '0'<<16 | '1'<<24)

//#define AVERROR_HTTP_FORBIDDEN     FFERRTAG(0xF8,'4','0','3')
const AVERROR_HTTP_FORBIDDEN = -(0xF8 | '4'<<8 | '0'<<16 | '3'<<24)

//#define AVERROR_HTTP_NOT_FOUND     FFERRTAG(0xF8,'4','0','4')
const AVERROR_HTTP_NOT_FOUND = -(0xF8 | '4'<<8 | '0'<<16 | '4'<<24)

//#define AVERROR_HTTP_OTHER_4XX     FFERRTAG(0xF8,'4','X','X')
const AVERROR_HTTP_OTHER_4XX = -(0xF8 | '4'<<8 | 'X'<<16 | 'X'<<24)

//#define AVERROR_HTTP_SERVER_ERROR  FFERRTAG(0xF8,'5','X','X')
const AVERROR_HTTP_SERVER_ERROR = -(0xF8 | '5'<<8 | 'X'<<16 | 'X'<<24)

const AV_ERROR_MAX_STRING_SIZE = 64

/**
 * Put a description of the AVERROR code errnum in errbuf.
 * In case of failure the global variable errno is set to indicate the
 * error. Even in case of failure av_strerror() will print a generic
 * error message indicating the errnum provided to errbuf.
 *
 * @param errnum      error code to describe
 * @param errbuf      buffer to which description is written
 * @param errbuf_size the size in bytes of errbuf
 * @return 0 on success, a negative value if a description for errnum
 * cannot be found
 */
//int av_strerror(int errnum, char *errbuf, size_t errbuf_size);
func AvStrerror(errnum ffcommon.FInt, errbuf ffcommon.FBuf, errbuf_size ffcommon.FSizeT) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_strerror").Call(
		uintptr(errnum),
		uintptr(unsafe.Pointer(errbuf)),
		uintptr(errbuf_size),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Fill the provided buffer with a string containing an error string
 * corresponding to the AVERROR code errnum.
 *
 * @param errbuf         a buffer
 * @param errbuf_size    size in bytes of errbuf
 * @param errnum         error code to describe
 * @return the buffer in input, filled with the error description
 * @see av_strerror()
 */
//static inline char *av_make_error_string(char *errbuf, size_t errbuf_size, int errnum)
//{
//    av_strerror(errnum, errbuf, errbuf_size);
//    return errbuf;
//}
func AvMakeErrorString(errbuf ffcommon.FBuf, errbuf_size ffcommon.FSizeT, errnum ffcommon.FInt) (res ffcommon.FCharP) {
	AvStrerror(errnum, errbuf, errbuf_size)
	res = ffcommon.StringFromPtr(uintptr(unsafe.Pointer(errbuf)))
	return

}

/**
 * Convenience macro, the return value should be used only directly in
 * function arguments but never stand-alone.
 */
//#define av_err2str(errnum) \
//    av_make_error_string((char[AV_ERROR_MAX_STRING_SIZE]){0}, AV_ERROR_MAX_STRING_SIZE, errnum)
func AvErr2str(errnum ffcommon.FInt) (res ffcommon.FCharP) {

	b := make([]byte, AV_ERROR_MAX_STRING_SIZE, AV_ERROR_MAX_STRING_SIZE)
	// AvStrerror(errnum, (*byte)(unsafe.Pointer(&b[0])), AV_ERROR_MAX_STRING_SIZE)
	// t, _, _ := ffcommon.GetAvutilDll().NewProc("av_err2str").Call()
	// res = ffcommon.StringFromPtr(t)

	AvMakeErrorString((*byte)(unsafe.Pointer(&b[0])), AV_ERROR_MAX_STRING_SIZE, errnum)
	res = ffcommon.StringFromPtr(uintptr(unsafe.Pointer(&b[0])))
	return
}

/**
 * @}
 */

//#endif /* AVUTIL_ERROR_H */
