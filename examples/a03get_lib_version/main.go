package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/ffmpeg-go/libavcodec"
	"github.com/moonfdd/ffmpeg-go/libavutil"
)

func main() {
	os.Setenv("Path", os.Getenv("Path")+";./lib")
	ffcommon.SetAvutilPath("./lib/avutil-56.dll")
	ffcommon.SetAvcodecPath("./lib/avcodec-58.dll")
	codecVer := libavcodec.AvcodecVersion()
	ver_major := (codecVer >> 16) & 0xff
	ver_minor := (codecVer >> 8) & 0xff
	ver_micro := (codecVer) & 0xff
	fmt.Printf("FFmpeg version is: %s .\navcodec version is: %d=%d.%d.%d.\n", libavutil.FFMPEG_VERSION, codecVer, ver_major, ver_minor, ver_micro)

	fmt.Println("---------------------------------")
	data, err := exec.Command("./lib/ffmpeg", "-version").Output()
	if err != nil {
		fmt.Println("ffmpeg err = ", err)
	}
	fmt.Println(string(data))
}
