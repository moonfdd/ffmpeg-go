package libavcodec

import "fmt"

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

//#ifndef AVCODEC_VERSION_H
//#define AVCODEC_VERSION_H
//
///**
// * @file
// * @ingroup libavc
// * Libavcodec version macros.
// */
//
//#include "../libavutil/version.h"

const LIBAVCODEC_VERSION_MAJOR = 58
const LIBAVCODEC_VERSION_MINOR = 134
const LIBAVCODEC_VERSION_MICRO = 100

//const LIBAVCODEC_VERSION_INT = AV_VERSION_INT(LIBAVCODEC_VERSION_MAJOR, \
//LIBAVCODEC_VERSION_MINOR, \
//LIBAVCODEC_VERSION_MICRO)
const LIBAVCODEC_VERSION_INT = LIBAVCODEC_VERSION_MAJOR<<16 | LIBAVCODEC_VERSION_MINOR<<8 | LIBAVCODEC_VERSION_MICRO

//const LIBAVCODEC_VERSION   =   AV_VERSION(LIBAVCODEC_VERSION_MAJOR,    \
//LIBAVCODEC_VERSION_MINOR,    \
//LIBAVCODEC_VERSION_MICRO)
//const LIBAVCODEC_BUILD     =   LIBAVCODEC_VERSION_INT
//
//const LIBAVCODEC_IDENT     =   "Lavc" AV_STRINGIFY(LIBAVCODEC_VERSION)
var LIBAVCODEC_IDENT = fmt.Sprintf("Lavc%d.%d.%d", LIBAVCODEC_VERSION_MAJOR, LIBAVCODEC_VERSION_MINOR, LIBAVCODEC_VERSION_MICRO)

/**
 * FF_API_* defines may be placed below to indicate public API that will be
 * dropped at a future version bump. The defines themselves are not part of
 * the public API and may change, break or disappear at any time.
 *
 * @note, when bumping the major version it is recommended to manually
 * disable each FF_API_* in its own commit instead of disabling them all
 * at once through the bump. This improves the git bisect-ability of the change.
 */

//#ifndef FF_API_AVCTX_TIMEBASE
const FF_API_AVCTX_TIMEBASE = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_CODED_FRAME
const FF_API_CODED_FRAME = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_SIDEDATA_ONLY_PKT
const FF_API_SIDEDATA_ONLY_PKT = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_VDPAU_PROFILE
const FF_API_VDPAU_PROFILE = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_CONVERGENCE_DURATION
const FF_API_CONVERGENCE_DURATION = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_AVPICTURE
const FF_API_AVPICTURE = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_AVPACKET_OLD_API
const FF_API_AVPACKET_OLD_API = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_RTP_CALLBACK
const FF_API_RTP_CALLBACK = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_VBV_DELAY
const FF_API_VBV_DELAY = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_CODER_TYPE
const FF_API_CODER_TYPE = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_STAT_BITS
const FF_API_STAT_BITS = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_PRIVATE_OPT
const FF_API_PRIVATE_OPT = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_ASS_TIMING
const FF_API_ASS_TIMING = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_OLD_BSF
const FF_API_OLD_BSF = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_COPY_CONTEXT
const FF_API_COPY_CONTEXT = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_GET_CONTEXT_DEFAULTS
const FF_API_GET_CONTEXT_DEFAULTS = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_NVENC_OLD_NAME
const FF_API_NVENC_OLD_NAME = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_STRUCT_VAAPI_CONTEXT
const FF_API_STRUCT_VAAPI_CONTEXT = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_MERGE_SD_API
const FF_API_MERGE_SD_API = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_TAG_STRING
const FF_API_TAG_STRING = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_GETCHROMA
const FF_API_GETCHROMA = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_CODEC_GET_SET
const FF_API_CODEC_GET_SET = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_USER_VISIBLE_AVHWACCEL
const FF_API_USER_VISIBLE_AVHWACCEL = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_LOCKMGR
const FF_API_LOCKMGR = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_NEXT
const FF_API_NEXT = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_UNSANITIZED_BITRATES
const FF_API_UNSANITIZED_BITRATES = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_OPENH264_SLICE_MODE
const FF_API_OPENH264_SLICE_MODE = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_OPENH264_CABAC
const FF_API_OPENH264_CABAC = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_UNUSED_CODEC_CAPS
const FF_API_UNUSED_CODEC_CAPS = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_AVPRIV_PUT_BITS
const FF_API_AVPRIV_PUT_BITS = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_OLD_ENCDEC
const FF_API_OLD_ENCDEC = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_AVCODEC_PIX_FMT
const FF_API_AVCODEC_PIX_FMT = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_MPV_RC_STRATEGY
const FF_API_MPV_RC_STRATEGY = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_PARSER_CHANGE
const FF_API_PARSER_CHANGE = (LIBAVCODEC_VERSION_MAJOR < 59)

//#endif
//#ifndef FF_API_THREAD_SAFE_CALLBACKS
const FF_API_THREAD_SAFE_CALLBACKS = (LIBAVCODEC_VERSION_MAJOR < 60)

//#endif
//#ifndef FF_API_DEBUG_MV
const FF_API_DEBUG_MV = (LIBAVCODEC_VERSION_MAJOR < 60)

//#endif
//#ifndef FF_API_GET_FRAME_CLASS
const FF_API_GET_FRAME_CLASS = (LIBAVCODEC_VERSION_MAJOR < 60)

//#endif
//#ifndef FF_API_AUTO_THREADS
const FF_API_AUTO_THREADS = (LIBAVCODEC_VERSION_MAJOR < 60)

//#endif
//#ifndef FF_API_INIT_PACKET
const FF_API_INIT_PACKET = (LIBAVCODEC_VERSION_MAJOR < 60)

//#endif
//
//#endif /* AVCODEC_VERSION_H */
