package main

import "C"
import (
	"fmt"
	"syscall"
	"time"
	"unsafe"
)

var (
	user32, _     = syscall.LoadLibrary("user32.dll")
	messageBox, _ = syscall.GetProcAddress(user32, "MessageBoxW")
)

func IntPtr(n int) uintptr {
	return uintptr(n)
}

//func StrPtr(s string) uintptr {
//	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
//}

func StrPtr(s string) uintptr {
	aaa, _ := syscall.UTF16PtrFromString(s)
	return uintptr(unsafe.Pointer(aaa))
}

func ShowMessage2(title string, text string) {
	user32 := syscall.NewLazyDLL("user32.dll")
	MessageBoxW := user32.NewProc("MessageBoxW")
	MessageBoxW.Call(IntPtr(0), StrPtr(text), StrPtr(title), IntPtr(0))
}
func AvutilVersion() uint {
	d := syscall.NewLazyDLL("F:/BaiduNetdiskDownload/ffmpeg-4.4-full_build-shared/bin/avutil-56.dll")
	p := d.NewProc("avutil_version")
	d.NewProc("avutil_version2222")
	ret1, ret2, err := p.Call()
	fmt.Println(ret1, ret2, err)
	return 0
}
func main() {

	ret := AvutilVersion()

	fmt.Println("ret = ", ret)

	if true {
	}
	if false {
		var x struct {
			a bool
			b int16
			c []int
		}
		// NOTE: subtly incorrect!
		tmp := uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
		pb := (*int16)(unsafe.Pointer(tmp))
		*pb = 421
		fmt.Println(*pb)
	}
	if false {
		var x struct {
			a bool
			b int16
			c []int
		}

		/**
		  unsafe.Offsetof 函数的参数必须是一个字段 x.f, 然后返回 f 字段相对于 x 起始地址的偏移量, 包括可能的空洞.
		*/

		/**
		  uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
		  指针的运算
		*/
		// 和 pb := &x.b 等价
		pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
		*pb = 42
		fmt.Println(x.b) // "42"
	}
	if false {
		defer syscall.FreeLibrary(user32)
		ShowMessage2("提升", "哈哈")
		time.Sleep(3 * time.Second)
	}
}
