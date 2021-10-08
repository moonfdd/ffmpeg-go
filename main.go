package main

import "C"
import (
	"ffmpeg-go/ffcommon"
	"ffmpeg-go/libavutil"
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

//https://blog.csdn.net/u010824081/article/details/79427676
func main() {
	if true {
		ret, _ := libavutil.AvSphericalFromName("abc哈哈")
		fmt.Println("AvSphericalFromName = ", ret)
	}
	if false {
		type add = func(a, b int) int
		var a add
		a = func(a, b int) int {
			return a + b
		}
		ret := a(1, 2)
		fmt.Println(ret)
	}
	if false {
		ret, _ := libavutil.AvXteaAlloc()
		fmt.Println("AvXteaAlloc = ", ret)
		ret.AvXteaLeInit([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
		fmt.Println("AvXteaLeInit = ", ret)
		ret.AvXteaLeCrypt(nil, nil, 0, nil, 0)
		fmt.Println("AvXteaLeCrypt = ", ret)
	}

	if false {
		b := new(byte)
		*b = 12
		p := unsafe.Pointer(b)
		u := uintptr(p)
		p2 := unsafe.Pointer(u)
		c := (*byte)(p2)
		*c = 33

		p = unsafe.Pointer(b)
		u = uintptr(p)
		p2 = unsafe.Pointer(u)
		c = (*byte)(p2)
		*c = 44
		fmt.Println("*b = ", *b)
	}
	if false {
		aes, _ := libavutil.AvAesAlloc()
		fmt.Println("aes111222= ", aes)
		b := byte(33)
		var buf = new(byte)
		*buf = b
		c := new(byte)
		*c = 11
		err := aes.AvAesCrypt(buf, c, ffcommon.FInt(1), nil, ffcommon.FInt(1))
		fmt.Println("AvAesInit ret = ", err)
		return
	}
	if false {
		aes, _ := libavutil.AvAesAlloc()
		b := byte(33)
		var buf = new(byte)
		*buf = b
		ret, _ := aes.AvAesInit(buf, ffcommon.FInt(192), ffcommon.FInt(1))
		fmt.Println("AvAesInit ret = ", ret)
		return
	}

	if false {
		b := byte(33)
		var buf = new(byte)
		*buf = b
		ret, _ := libavutil.AvAesAlloc()
		fmt.Println("AvAesAlloc ret = ", ret)
		return
	}

	if false {
		b := byte(33)
		var buf = new(byte)
		*buf = b
		ret, _ := libavutil.AvAdler32Update(uint32(5), buf, 1)
		fmt.Println("AvAdler32Update ret = ", ret)
	}

	//if true {
	//	ret := libavutil.AvutilLicense()
	//	fmt.Println("AvutilLicense ret = ", ret)
	//}
	//if true {
	//	ret := libavutil.AvutilConfiguration()
	//	fmt.Println("AvutilConfiguration ret = ", ret)
	//}
	//if true {
	//	ret := libavutil.AvVersionInfo()
	//	fmt.Println("AvVersionInfo ret = ", ret)
	//}
	//if true {
	//	ret := libavutil.AvGetMediaTypeString(ffconstant.AVMEDIA_TYPE_VIDEO)
	//	fmt.Println("AvGetMediaTypeString ret = ", ret)
	//}
	//
	//if true {
	//	ret := libavutil.AvGetPictureTypeChar(ffconstant.AV_PICTURE_TYPE_I)
	//	fmt.Println("AvGetPictureTypeChar ret = ", ret)
	//}
	//
	//if true {
	//	ret := libavutil.AvIntListLengthForSize(5, uintptr(0), 0)
	//	fmt.Println("AvIntListLengthForSize ret = ", ret)
	//}
	//if true {
	//	ret := libavutil.AvFopenUtf8("F:/看视频进度.txt", "r+")
	//	fmt.Println("AvFopenUtf8 ret = ", ret)
	//}
	//if true {
	//	ret := libavutil.AvGetTimeBaseQ()
	//	fmt.Println("AvGetTimeBaseQ retden = ", ret.Den)
	//	fmt.Println("AvGetTimeBaseQ retnum = ", ret.Num)
	//}

	if false {

		ret, _ := libavutil.AvFourccMakeString((*byte)(unsafe.Pointer(&([]byte{'a', 'b', 'c'}))), 3)
		fmt.Println("AvFourccMakeString ret = ", ret)
	}

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
