package main

import (
	"fmt"
	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/ffmpeg-go/libavutil"
)

func main() {
	d := int32(-1)
	u := uintptr(d)
	d2 := uint32(u)
	fmt.Println(d2)
	return
	ffcommon.SetAvutilPath("F:/BaiduNetdiskDownload/ffmpeg-4.4-full_build-shared/bin/avutil-56.dll")
	if true {
		fmt.Println(libavutil.AvCpuMaxAlign())
	}
	return
	if false {
		ans := libavutil.AvAdler32Update(111, nil, 0)
		fmt.Println(ans)
	}
	if false {
		ans := libavutil.AvAesAlloc()
		fmt.Println(ans)
	}
	if false {
		fmt.Println(libavutil.AV_MATRIX_ENCODING_DOLBYHEADPHONE)
	}
	if true {
		fmt.Println(libavutil.AvutilVersion())
	}
	if true {
		fmt.Println(libavutil.AvVersionInfo())
	}
	if true {
		fmt.Println(libavutil.AvutilConfiguration())
	}
	if true {
		fmt.Println(libavutil.AvutilLicense())
	}
	if true {
		fmt.Println(libavutil.AVMEDIA_TYPE_VIDEO)
		fmt.Println(libavutil.AvGetMediaTypeString(libavutil.AVMEDIA_TYPE_AUDIO))
	}
	if true {
		fmt.Println(libavutil.AvGetTimeBaseQ())
		libavutil.AvGetTimeBaseQ()
	}
}
