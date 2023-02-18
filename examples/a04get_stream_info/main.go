package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/ffmpeg-go/libavformat"
)

func main() {
	os.Setenv("Path", os.Getenv("Path")+";./lib")
	ffcommon.SetAvformatPath("./lib/avformat-58.dll")
	fmt_ctx := libavformat.AvformatAllocContext() //创建对象并初始化
	ret := int32(0)
	fileName := "./resources/big_buck_bunny.mp4" //文件地址

	for {
		//打开文件
		ret = libavformat.AvformatOpenInput(&fmt_ctx, fileName, nil, nil)
		if ret < 0 {
			fmt.Printf("Cannot open video file\n")
			break //Cannot open video file
		}

		//查找流信息（音频流和视频流）
		ret = fmt_ctx.AvformatFindStreamInfo(nil)
		if ret < 0 {
			fmt.Printf("Cannot find stream information\n")
			break
		}

		fmt_ctx.AvDumpFormat(0, fileName, 0) //输出视频信息
		break
	}

	libavformat.AvformatCloseInput(&fmt_ctx) //关闭文件

	fmt.Println("---------------------------------")
	cmd := exec.Command("./lib/ffprobe", fileName)
	data, err2 := cmd.CombinedOutput()
	if err2 != nil {
		fmt.Println("ffprobe err = ", err2)
		return
	}
	fmt.Println(string(data))
}
