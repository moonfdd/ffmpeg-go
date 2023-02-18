package main

import (
	"fmt"
	"os"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/ffmpeg-go/libavutil"
)

func main() {
	os.Setenv("Path", os.Getenv("Path")+";./lib")
	ffcommon.SetAvutilPath("./lib/avutil-56.dll")
	ffcommon.SetAvcodecPath("./lib/avcodec-58.dll")
	ffcommon.SetAvdevicePath("./lib/avdevice-56.dll")
	ffcommon.SetAvfilterPath("./lib/avfilter-56.dll")
	ffcommon.SetAvformatPath("./lib/avformat-58.dll")
	ffcommon.SetAvpostprocPath("./lib/postproc-55.dll")
	ffcommon.SetAvswresamplePath("./lib/swresample-3.dll")
	ffcommon.SetAvswscalePath("./lib/swscale-5.dll")
	if true {
		ret := libavutil.AvFrameAlloc()
		fmt.Println(ret)
		fmt.Println(libavutil.AV_NUM_DATA_POINTERS)
	}
	// if true {
	// 	libavutil.AvLog(0, 0, "a", "b")
	// }
	// if true {
	// 	ret := libavcodec.AvcodecVersion()
	// 	fmt.Println(ret)
	// }
	// if true {
	// 	fmt.Println(libavutil.AvVersionInfo())
	// }
	// if true {
	// 	fmt.Println(libavcodec.AvcodecLicense())
	// }
	// if true {
	// 	fmt.Println(libavcodec.AvcodecConfiguration())
	// }
	// if true {
	// 	ans := libavutil.AvAdler32Update(111, nil, 0)
	// 	fmt.Println(ans)
	// }
	// if true {
	// 	ans := libavutil.AvAesAlloc()
	// 	fmt.Println(ans)
	// }
	// if true {
	// 	fmt.Println(libavutil.AV_MATRIX_ENCODING_DOLBYHEADPHONE)
	// }
	// if true {
	// 	fmt.Println(libavutil.AvutilVersion())
	// }
	// if true {
	// 	fmt.Println(libavutil.AvVersionInfo())
	// }
	// if true {
	// 	fmt.Println(libavutil.AvutilConfiguration())
	// }
	// if true {
	// 	fmt.Println(libavutil.AvutilLicense())
	// }
	// if true {
	// 	fmt.Println(libavutil.AVMEDIA_TYPE_VIDEO)
	// 	fmt.Println(libavutil.AvGetMediaTypeString(libavutil.AVMEDIA_TYPE_AUDIO))
	// }
	// if true {
	// 	fmt.Println(libavutil.AvGetTimeBaseQ())
	// 	libavutil.AvGetTimeBaseQ()
	// }
}
