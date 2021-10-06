package ffconstant

///* error handling */
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
//#define AVERROR_BUG                FFERRTAG( 'B','U','G','!') ///< Internal bug, also see AVERROR_BUG2
//#define AVERROR_BUFFER_TOO_SMALL   FFERRTAG( 'B','U','F','S') ///< Buffer too small
//#define AVERROR_DECODER_NOT_FOUND  FFERRTAG(0xF8,'D','E','C') ///< Decoder not found
//#define AVERROR_DEMUXER_NOT_FOUND  FFERRTAG(0xF8,'D','E','M') ///< Demuxer not found
//#define AVERROR_ENCODER_NOT_FOUND  FFERRTAG(0xF8,'E','N','C') ///< Encoder not found
//#define AVERROR_EOF                FFERRTAG( 'E','O','F',' ') ///< End of file
//#define AVERROR_EXIT               FFERRTAG( 'E','X','I','T') ///< Immediate exit was requested; the called function should not be restarted
//#define AVERROR_EXTERNAL           FFERRTAG( 'E','X','T',' ') ///< Generic error in an external library
//#define AVERROR_FILTER_NOT_FOUND   FFERRTAG(0xF8,'F','I','L') ///< Filter not found
//#define AVERROR_INVALIDDATA        FFERRTAG( 'I','N','D','A') ///< Invalid data found when processing input
//#define AVERROR_MUXER_NOT_FOUND    FFERRTAG(0xF8,'M','U','X') ///< Muxer not found
//#define AVERROR_OPTION_NOT_FOUND   FFERRTAG(0xF8,'O','P','T') ///< Option not found
//#define AVERROR_PATCHWELCOME       FFERRTAG( 'P','A','W','E') ///< Not yet implemented in FFmpeg, patches welcome
//#define AVERROR_PROTOCOL_NOT_FOUND FFERRTAG(0xF8,'P','R','O') ///< Protocol not found
//
//#define AVERROR_STREAM_NOT_FOUND   FFERRTAG(0xF8,'S','T','R') ///< Stream not found
///**
// * This is semantically identical to AVERROR_BUG
// * it has been introduced in Libav after our AVERROR_BUG and with a modified value.
// */
//#define AVERROR_BUG2               FFERRTAG( 'B','U','G',' ')
//#define AVERROR_UNKNOWN            FFERRTAG( 'U','N','K','N') ///< Unknown error, typically from an external library
const AVERROR_EXPERIMENTAL = (-0x2bb2afa8)   ///< Requested feature is flagged experimental. Set strict_std_compliance if you really want to use it.
const AVERROR_INPUT_CHANGED = (-0x636e6701)  ///< Input changed between calls. Reconfiguration is required. (can be OR-ed with AVERROR_OUTPUT_CHANGED)
const AVERROR_OUTPUT_CHANGED = (-0x636e6702) ///< Output changed between calls. Reconfiguration is required. (can be OR-ed with AVERROR_INPUT_CHANGED)
///* HTTP & RTSP errors */
//#define AVERROR_HTTP_BAD_REQUEST   FFERRTAG(0xF8,'4','0','0')
//#define AVERROR_HTTP_UNAUTHORIZED  FFERRTAG(0xF8,'4','0','1')
//#define AVERROR_HTTP_FORBIDDEN     FFERRTAG(0xF8,'4','0','3')
//#define AVERROR_HTTP_NOT_FOUND     FFERRTAG(0xF8,'4','0','4')
//#define AVERROR_HTTP_OTHER_4XX     FFERRTAG(0xF8,'4','X','X')
//#define AVERROR_HTTP_SERVER_ERROR  FFERRTAG(0xF8,'5','X','X')

const AV_ERROR_MAX_STRING_SIZE = 64
