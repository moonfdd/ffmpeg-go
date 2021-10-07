package ffcommon

import "syscall"

var avUtilDll *syscall.LazyDLL

func GetAvutilDll() (ans *syscall.LazyDLL) {
	avUtilDll = syscall.NewLazyDLL("F:/BaiduNetdiskDownload/ffmpeg-4.4-full_build-shared/bin/avutil-56.dll")
	ans = avUtilDll
	return
}
