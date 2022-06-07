package main

import "C"
import (
	"fmt"
	"github.com/moonfdd/ffmpeg-go/libavcodec"
	"github.com/moonfdd/ffmpeg-go/libavdevice"
	"github.com/moonfdd/ffmpeg-go/libavfilter"
	"github.com/moonfdd/ffmpeg-go/libavformat"
	"github.com/moonfdd/ffmpeg-go/libavutil"
	"github.com/moonfdd/ffmpeg-go/libswscale"
	"reflect"
	"unsafe"
)

//https://blog.csdn.net/u010824081/article/details/79427676
func main() {
	if true {
		var a [2]libavformat.AVIODirEntry
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavformat.AVIODirContext
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavformat.AVIOContext
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavformat.AVIOContext
		b := uintptr(unsafe.Pointer(&a[0].BytesRead)) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	return
	if true {
		var a [2]libavutil.AVFrameSideData
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVRegionOfInterest
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVFrame
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVRational
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVBlowfish
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVBPrint
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVBufferRef
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVDES
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVDictionaryEntry
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVDOVIDecoderConfigurationRecord
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVDownmixInfo
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVSubsampleEncryptionInfo
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVEncryptionInfo
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVEncryptionInitInfo
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVFifoBuffer
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVFilmGrainAOMParams
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVFilmGrainParams
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVHDRPlusPercentile
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVHDRPlusColorTransformParams
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
		var c libavutil.AVHDRPlusColorTransformParams
		b = uintptr(unsafe.Pointer(&c.ColorSaturationMappingFlag)) - uintptr(unsafe.Pointer(&c))
		fmt.Println("b = ", b)
	}
	if true {
		var a [2]libavutil.AVDynamicHDRPlus
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVHWDeviceContext
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVHWFramesContext
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVLFG
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVClass
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}

	if true {
		var a [2]libavutil.AVMasteringDisplayMetadata
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVContentLightMetadata
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVMotionVector
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}

	if true {
		var a [2]libavutil.AVOption
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVOptionRange
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVOptionRanges
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}

	if true {
		var a [2]libavutil.AVComponentDescriptor
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVPixFmtDescriptor
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}

	if true {
		var a [2]libavutil.AVRC4
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}

	if true {
		var a [2]libavutil.AVReplayGain
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}

	if true {
		var a [2]libavutil.AVSphericalMapping
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}

	if true {
		var a [2]libavutil.AVStereo3D
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}

	if true {
		var a [2]libavutil.AVTimecode
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}

	if true {
		var a [2]libavutil.AVComplexFloat
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVComplexDouble
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVComplexInt32
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}

	if true {
		var a [2]libavutil.AVVideoEncParams
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavutil.AVVideoBlockParams
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}

	if true {
		var a [2]libavutil.AVXTEA
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}

	if true {
		var a [2]libswscale.SwsVector
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libswscale.SwsFilter
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}

	if true {
		var a [2]libavformat.AVIOInterruptCB
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	fmt.Println("-------------------")
	if true {
		var a [2]libavformat.AVProbeData
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavformat.AVOutputFormat
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavformat.AVInputFormat
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavformat.AVIndexEntry
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavformat.AVStream
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavformat.AVProgram
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavformat.AVChapter
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavformat.AVFormatContext
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}

	if true {
		var a [2]libavfilter.AVBufferSrcParameters
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}

	if true {
		var a [2]libavfilter.AVBufferSinkParams
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavfilter.AVABufferSinkParams
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	fmt.Println("-------------------")
	if true {
		var a [2]libavfilter.AVFilter
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavfilter.AVFilterContext
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavfilter.AVFilterFormatsConfig
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavfilter.AVFilterLink
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavfilter.AVFilterGraph
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
		b = uintptr(unsafe.Pointer(&a[0].SinkLinksCount)) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println("b2 = ", b)
	}
	if true {
		var a [2]libavfilter.AVFilterInOut
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	fmt.Println("-------------------")
	if true {
		var a [2]libavdevice.AVDeviceRect
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavdevice.AVDeviceCapabilitiesQuery
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavdevice.AVDeviceInfo
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavdevice.AVDeviceInfoList
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	fmt.Println("-------------------")
	if true {
		var a [2]libavcodec.AVPacketSideData
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavcodec.AVPacket
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavcodec.AVPacketList
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}

	if true {
		var a [2]libavcodec.AVMediaCodecContext
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}

	if true {
		var a [2]libavcodec.AVDVProfile
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}

	if true {
		var a [2]libavcodec.DiracVersionInfo
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavcodec.AVDiracSeqHeader
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}

	if true {
		var a [2]libavcodec.AVCodecParameters
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}

	if true {
		var a [2]libavcodec.AVCodecDescriptor
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	fmt.Println("-------------------")
	if true {
		var a [2]libavcodec.AVProfile
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavcodec.AVCodec
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavcodec.AVCodecHWConfig
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}

	if true {
		var a [2]libavcodec.AVBSFContext
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavcodec.AVBitStreamFilter
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}

	if true {
		var a [2]libavcodec.FFTComplex
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}

	if true {
		var a [2]libavcodec.AVDCT
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	fmt.Println("-------------------")
	if true {
		var a [2]libavcodec.RcOverride
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavcodec.AVPanScan
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavcodec.AVCPBProperties
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavcodec.AVProducerReferenceTime
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavcodec.AVCodecContext
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}

	if true {
		var a [2]libavcodec.AVHWAccel
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavcodec.AVPicture
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavcodec.AVSubtitleRect
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavcodec.AVSubtitle
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavcodec.AVCodecParserContext
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}

	if true {
		var a [2]libavcodec.AVCodecParser
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
	if true {
		var a [2]libavcodec.AVBitStreamFilterContext
		b := uintptr(unsafe.Pointer(&a[1])) - uintptr(unsafe.Pointer(&a[0]))
		fmt.Println(b)
	}
}
func SizeStruct(data interface{}) int {
	return sizeof(reflect.ValueOf(data))
}

func sizeof(v reflect.Value) int {
	switch v.Kind() {
	case reflect.Map:
		sum := 0
		keys := v.MapKeys()
		for i := 0; i < len(keys); i++ {
			mapkey := keys[i]
			s := sizeof(mapkey)
			if s < 0 {
				return -1
			}
			sum += s
			s = sizeof(v.MapIndex(mapkey))
			if s < 0 {
				return -1
			}
			sum += s
		}
		return sum
	case reflect.Slice, reflect.Array:
		sum := 0
		for i, n := 0, v.Len(); i < n; i++ {
			s := sizeof(v.Index(i))
			if s < 0 {
				return -1
			}
			sum += s
		}
		return sum

	case reflect.String:
		sum := 0
		for i, n := 0, v.Len(); i < n; i++ {
			s := sizeof(v.Index(i))
			if s < 0 {
				return -1
			}
			sum += s
		}
		return sum

	case reflect.Ptr, reflect.Interface:
		p := (*[]byte)(unsafe.Pointer(v.Pointer()))
		if p == nil {
			return 0
		}
		return sizeof(v.Elem())
	case reflect.Struct:
		sum := 0
		for i, n := 0, v.NumField(); i < n; i++ {
			s := sizeof(v.Field(i))
			if s < 0 {
				return -1
			}
			sum += s
		}
		return sum

	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128,
		reflect.Int:
		return int(v.Type().Size())

	default:
		fmt.Println("t.Kind() no found:", v.Kind())
	}

	return -1
}
