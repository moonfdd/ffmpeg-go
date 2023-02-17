package ffcommon

import (
	"sync"

	"github.com/ying32/dylib"
)

var avUtilDll *dylib.LazyDLL
var avUtilDllOnce sync.Once

func GetAvutilDll() (ans *dylib.LazyDLL) {
	avUtilDllOnce.Do(func() {
		avUtilDll = dylib.NewLazyDLL(avutilPath)
	})
	ans = avUtilDll
	return
}

var avutilPath = "avutil-56.dll"

func SetAvutilPath(path0 string) {
	avutilPath = path0
}

var avcodecDll *dylib.LazyDLL
var avcodecDllOnce sync.Once

func GetAvcodecDll() (ans *dylib.LazyDLL) {
	avcodecDllOnce.Do(func() {
		avcodecDll = dylib.NewLazyDLL(avcodecPath)
	})
	ans = avcodecDll
	return
}

var avcodecPath = "avcodec-56.dll"

func SetAvcodecPath(path0 string) {
	avcodecPath = path0
}

var avdeviceDll *dylib.LazyDLL
var avdeviceDllOnce sync.Once

func GetAvdeviceDll() (ans *dylib.LazyDLL) {
	avdeviceDllOnce.Do(func() {
		avdeviceDll = dylib.NewLazyDLL(avdevicePath)
	})
	ans = avdeviceDll
	return
}

var avdevicePath = "avdevice-56.dll"

func SetAvdevicePath(path0 string) {
	avdevicePath = path0
}

var avfilterDll *dylib.LazyDLL
var avfilterDllOnce sync.Once

func GetAvfilterDll() (ans *dylib.LazyDLL) {
	avfilterDllOnce.Do(func() {
		avfilterDll = dylib.NewLazyDLL(avfilterPath)
	})
	ans = avfilterDll
	return
}

var avfilterPath = "avfilter-56.dll"

func SetAvfilterPath(path0 string) {
	avfilterPath = path0
}

var avformatDll *dylib.LazyDLL
var avformatDllOnce sync.Once

func GetAvformatDll() (ans *dylib.LazyDLL) {
	avformatDllOnce.Do(func() {
		avformatDll = dylib.NewLazyDLL(avformatPath)
	})
	ans = avformatDll
	return
}

var avformatPath = "avformat-58.dll"

func SetAvformatPath(path0 string) {
	avformatPath = path0
}

var avpostprocDll *dylib.LazyDLL
var avpostprocDllOnce sync.Once

func GetAvpostprocDll() (ans *dylib.LazyDLL) {
	avpostprocDllOnce.Do(func() {
		avpostprocDll = dylib.NewLazyDLL(avpostprocPath)
	})
	ans = avpostprocDll
	return
}

var avpostprocPath = "postproc-55.dll"

func SetAvpostprocPath(path0 string) {
	avpostprocPath = path0
}

var avswresampleDll *dylib.LazyDLL
var avswresampleDllOnce sync.Once

func GetAvswresampleDll() (ans *dylib.LazyDLL) {
	avswresampleDllOnce.Do(func() {
		avswresampleDll = dylib.NewLazyDLL(avswresamplePath)
	})
	ans = avswresampleDll
	return
}

var avswresamplePath = "swresample-3.dll"

func SetAvswresamplePath(path0 string) {
	avswresamplePath = path0
}

var avswscaleDll *dylib.LazyDLL
var avswscaleDllOnce sync.Once

func GetAvswscaleDll() (ans *dylib.LazyDLL) {
	avswscaleDllOnce.Do(func() {
		avswscaleDll = dylib.NewLazyDLL(avswscalePath)
	})
	ans = avswscaleDll
	return
}

var avswscalePath = "swscale-5.dll"

func SetAvswscalePath(path0 string) {
	avswscalePath = path0
}
