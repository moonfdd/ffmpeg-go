package ffcommon

import "syscall"

var avUtilDll *syscall.LazyDLL

func GetAvcodecDll() (ans *syscall.LazyDLL) {
	avUtilDll = syscall.NewLazyDLL("F:/BaiduNetdiskDownload/ffmpeg-4.4-full_build-shared/bin/avcodec-56.dll")
	ans = avUtilDll
	return
}

func GetAvdeviceDll() (ans *syscall.LazyDLL) {
	avUtilDll = syscall.NewLazyDLL("F:/BaiduNetdiskDownload/ffmpeg-4.4-full_build-shared/bin/avdevice-56.dll")
	ans = avUtilDll
	return
}

func GetAvfilterDll() (ans *syscall.LazyDLL) {
	avUtilDll = syscall.NewLazyDLL("F:/BaiduNetdiskDownload/ffmpeg-4.4-full_build-shared/bin/avfilter-56.dll")
	ans = avUtilDll
	return
}
func GetAvformatDll() (ans *syscall.LazyDLL) {
	avUtilDll = syscall.NewLazyDLL("F:/BaiduNetdiskDownload/ffmpeg-4.4-full_build-shared/bin/avformat-56.dll")
	ans = avUtilDll
	return
}
func GetAvutilDll() (ans *syscall.LazyDLL) {
	avUtilDll = syscall.NewLazyDLL("F:/BaiduNetdiskDownload/ffmpeg-4.4-full_build-shared/bin/avutil-56.dll")
	ans = avUtilDll
	return
}
func GetAvpostprocDll() (ans *syscall.LazyDLL) {
	avUtilDll = syscall.NewLazyDLL("F:/BaiduNetdiskDownload/ffmpeg-4.4-full_build-shared/bin/avutil-56.dll")
	ans = avUtilDll
	return
}
func GetAvswresampleDll() (ans *syscall.LazyDLL) {
	avUtilDll = syscall.NewLazyDLL("F:/BaiduNetdiskDownload/ffmpeg-4.4-full_build-shared/bin/avutil-56.dll")
	ans = avUtilDll
	return
}
func GetAvswscaleDll() (ans *syscall.LazyDLL) {
	avUtilDll = syscall.NewLazyDLL("F:/BaiduNetdiskDownload/ffmpeg-4.4-full_build-shared/bin/avutil-56.dll")
	ans = avUtilDll
	return
}
