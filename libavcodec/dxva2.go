package libavcodec

import "github.com/moonfdd/ffmpeg-go/ffcommon"

/*
 * DXVA2 HW acceleration
 *
 * copyright (c) 2009 Laurent Aimar
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

//#ifndef AVCODEC_DXVA2_H
//#define AVCODEC_DXVA2_H
//
///**
// * @file
// * @ingroup lavc_codec_hwaccel_dxva2
// * Public libavcodec DXVA2 header.
// */
//
//#if !defined(_WIN32_WINNT) || _WIN32_WINNT < 0x0602
//#undef _WIN32_WINNT
//#define _WIN32_WINNT 0x0602
//#endif
//
//#include <stdint.h>
//#include <d3d9.h>
//#include <dxva2api.h>

/**
 * @defgroup lavc_codec_hwaccel_dxva2 DXVA2
 * @ingroup lavc_codec_hwaccel
 *
 * @{
 */

//const FF_DXVA2_WORKAROUND_SCALING_LIST_ZIGZAG= 1 ///< Work around for DXVA2 and old UVD/UVD+ ATI video cards
//const FF_DXVA2_WORKAROUND_INTEL_CLEARVIDEO   = 2 ///< Work around for DXVA2 and old Intel GPUs with ClearVideo interface

/**
 * This structure is used to provides the necessary configurations and data
 * to the DXVA2 FFmpeg HWAccel implementation.
 *
 * The application must make it available as AVCodecContext.hwaccel_context.
 */
type dxva_context struct {

	/**
	 * DXVA2 decoder object
	 */
	//IDirectXVideoDecoder *decoder;
	decoder uintptr

	/**
	 * DXVA2 configuration used to create the decoder
	 */
	//const DXVA2_ConfigPictureDecode *cfg;
	cfg uintptr

	/**
	 * The number of surface in the surface array
	 */
	surface_count ffcommon.FUnsigned

	/**
	 * The array of Direct3D surfaces used to create the decoder
	 */
	//LPDIRECT3DSURFACE9 *surface;
	surface uintptr
	/**
	 * A bit field configuring the workarounds needed for using the decoder
	 */
	workaround ffcommon.FUint64T

	/**
	 * Private to the FFmpeg AVHWAccel implementation
	 */
	report_id ffcommon.FUnsigned
}

/**
 * @}
 */

//#endif /* AVCODEC_DXVA2_H */
