package libavcodec

import (
	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"unsafe"
)

/*
 * Direct3D11 HW acceleration
 *
 * copyright (c) 2009 Laurent Aimar
 * copyright (c) 2015 Steve Lhomme
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

//#ifndef AVCODEC_D3D11VA_H
//#define AVCODEC_D3D11VA_H
//
///**
// * @file
// * @ingroup lavc_codec_hwaccel_d3d11va
// * Public libavcodec D3D11VA header.
// */
//
//#if !defined(_WIN32_WINNT) || _WIN32_WINNT < 0x0602
//#undef _WIN32_WINNT
//#define _WIN32_WINNT 0x0602
//#endif
//
//#include <stdint.h>
//#include <d3d11.h>

/**
 * @defgroup lavc_codec_hwaccel_d3d11va Direct3D11
 * @ingroup lavc_codec_hwaccel
 *
 * @{
 */

const FF_DXVA2_WORKAROUND_SCALING_LIST_ZIGZAG = 1 ///< Work around for Direct3D11 and old UVD/UVD+ ATI video cards
const FF_DXVA2_WORKAROUND_INTEL_CLEARVIDEO = 2    ///< Work around for Direct3D11 and old Intel GPUs with ClearVideo interface

/**
 * This structure is used to provides the necessary configurations and data
 * to the Direct3D11 FFmpeg HWAccel implementation.
 *
 * The application must make it available as AVCodecContext.hwaccel_context.
 *
 * Use av_d3d11va_alloc_context() exclusively to allocate an AVD3D11VAContext.
 */
type AVD3D11VAContext struct {
	/**
	 * D3D11 decoder object
	 */
	//decoder *ID3D11VideoDecoder
	decoder uintptr

	/**
	 * D3D11 VideoContext
	 */
	//ID3D11VideoContext *video_context;
	video_context uintptr

	/**
	 * D3D11 configuration used to create the decoder
	 */
	//D3D11_VIDEO_DECODER_CONFIG *cfg;
	cfg uintptr

	/**
	 * The number of surface in the surface array
	 */
	surface_count ffcommon.FUnsigned

	/**
	 * The array of Direct3D surfaces used to create the decoder
	 */
	//ID3D11VideoDecoderOutputView **surface;
	surface *uintptr
	/**
	 * A bit field configuring the workarounds needed for using the decoder
	 */
	workaround ffcommon.FUint64T

	/**
	 * Private to the FFmpeg AVHWAccel implementation
	 */
	report_id ffcommon.FUnsigned

	/**
	 * Mutex to access video_context
	 */
	//HANDLE  context_mutex;
	context_mutex uintptr
}

/**
 * Allocate an AVD3D11VAContext.
 *
 * @return Newly-allocated AVD3D11VAContext or NULL on failure.
 */
//AVD3D11VAContext *av_d3d11va_alloc_context(void);
func AvD3d11vaAllocContext() (res *AVD3D11VAContext) {
	t, _, _ := ffcommon.GetAvcodecDll().NewProc("av_d3d11va_alloc_context").Call()
	if t == 0 {

	}
	res = (*AVD3D11VAContext)(unsafe.Pointer(t))
	return
}

/**
 * @}
 */

//#endif /* AVCODEC_D3D11VA_H */
