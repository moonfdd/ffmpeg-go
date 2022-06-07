package ffcommon

import (
	"sync"
	"syscall"
)

var avUtilDll *syscall.LazyDLL
var avUtilDllOnce sync.Once

func GetAvutilDll() (ans *syscall.LazyDLL) {
	avUtilDllOnce.Do(func() {
		avUtilDll = syscall.NewLazyDLL(avutilPath)
	})
	ans = avUtilDll
	return
}

//F:/BaiduNetdiskDownload/ffmpeg-4.4-full_build-shared/bin/avutil-56.dll
var avutilPath = "avutil-56.dll"

func SetAvutilPath(path0 string) {
	avutilPath = path0
}

var avcodecDll *syscall.LazyDLL
var avcodecDllOnce sync.Once

func GetAvcodecDll() (ans *syscall.LazyDLL) {
	avcodecDllOnce.Do(func() {
		avcodecDll = syscall.NewLazyDLL(avcodecPath)
	})
	ans = avcodecDll
	return
}

//F:/BaiduNetdiskDownload/ffmpeg-4.4-full_build-shared/bin/avcodec-56.dll
var avcodecPath = "avcodec-56.dll"

func SetAvcodecPath(path0 string) {
	avcodecPath = path0
}

var avdeviceDll *syscall.LazyDLL
var avdeviceDllOnce sync.Once

func GetAvdeviceDll() (ans *syscall.LazyDLL) {
	avdeviceDllOnce.Do(func() {
		avdeviceDll = syscall.NewLazyDLL(avdevicePath)
	})
	ans = avdeviceDll
	return
}

//F:/BaiduNetdiskDownload/ffmpeg-4.4-full_build-shared/bin/avdevice-56.dll
var avdevicePath = "avdevice-56.dll"

func SetAvdevicePath(path0 string) {
	avdevicePath = path0
}

var avfilterDll *syscall.LazyDLL
var avfilterDllOnce sync.Once

func GetAvfilterDll() (ans *syscall.LazyDLL) {
	avfilterDllOnce.Do(func() {
		avfilterDll = syscall.NewLazyDLL(avfilterPath)
	})
	ans = avfilterDll
	return
}

//F:/BaiduNetdiskDownload/ffmpeg-4.4-full_build-shared/bin/avfilter-56.dll
var avfilterPath = "avfilter-56.dll"

func SetAvfilterPath(path0 string) {
	avfilterPath = path0
}

var avformatDll *syscall.LazyDLL
var avformatDllOnce sync.Once

func GetAvformatDll() (ans *syscall.LazyDLL) {
	avformatDllOnce.Do(func() {
		avformatDll = syscall.NewLazyDLL(avformatPath)
	})
	ans = avformatDll
	return
}

//F:/BaiduNetdiskDownload/ffmpeg-4.4-full_build-shared/bin/avformat-56.dll
var avformatPath = "avformat-56.dll"

func SetAvformatPath(path0 string) {
	avformatPath = path0
}

var avpostprocDll *syscall.LazyDLL
var avpostprocDllOnce sync.Once

func GetAvpostprocDll() (ans *syscall.LazyDLL) {
	avpostprocDllOnce.Do(func() {
		avpostprocDll = syscall.NewLazyDLL(avpostprocPath)
	})
	ans = avpostprocDll
	return
}

//F:/BaiduNetdiskDownload/ffmpeg-4.4-full_build-shared/bin/postproc-55.dll
var avpostprocPath = "postproc-55.dll"

func SetAvpostprocPath(path0 string) {
	avpostprocPath = path0
}

var avswresampleDll *syscall.LazyDLL
var avswresampleDllOnce sync.Once

func GetAvswresampleDll() (ans *syscall.LazyDLL) {
	avswresampleDllOnce.Do(func() {
		avswresampleDll = syscall.NewLazyDLL(avswresamplePath)
	})
	ans = avswresampleDll
	return
}

//F:/BaiduNetdiskDownload/ffmpeg-4.4-full_build-shared/bin/swresample-3.dll
var avswresamplePath = "swresample-3.dll"

func SetAvswresamplePath(path0 string) {
	avswresamplePath = path0
}

var avswscaleDll *syscall.LazyDLL
var avswscaleDllOnce sync.Once

func GetAvswscaleDll() (ans *syscall.LazyDLL) {
	avswscaleDllOnce.Do(func() {
		avswscaleDll = syscall.NewLazyDLL(avswscalePath)
	})
	ans = avswscaleDll
	return
}

//F:/BaiduNetdiskDownload/ffmpeg-4.4-full_build-shared/bin/swscale-5.dll
var avswscalePath = "swscale-5.dll"

func SetAvswscalePath(path0 string) {
	avswscalePath = path0
}
