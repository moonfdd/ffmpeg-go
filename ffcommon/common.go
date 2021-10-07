package ffcommon

import (
	"syscall"
	"unsafe"
)

func UintptrToString(sptr uintptr) (res string) {
	if sptr <= 0 {
		return
	}
	i := 0
	buf := make([]byte, 0)
	for *(*byte)(unsafe.Pointer(sptr + uintptr(i))) != '0' {
		i++
		buf = append(buf, *(*byte)(unsafe.Pointer(sptr + uintptr(i))))
	}
	res = string(buf)
	return
}

func GoBool(val uintptr) bool {
	if val != 0 {
		return true
	}
	return false
}

func CBool(val bool) uintptr {
	if val {
		return 1
	}
	return 0
}

var kernel32dll = syscall.NewLazyDLL("kernel32.dll")
var _lstrlenW = kernel32dll.NewProc("lstrlenW")

func LstrlenW(lpString uintptr) int32 {
	r, _, _ := _lstrlenW.Call(lpString)
	return int32(r)
}

var msvcrtdll = syscall.NewLazyDLL("msvcrt.dll")
var _memcpy = msvcrtdll.NewProc("memcpy")

func Memcpy(dest, src uintptr, count uintptr) uintptr {
	r, _, _ := _memcpy.Call(dest, src, count)
	return r
}
func GoWStr(str uintptr) string {
	if str == 0 {
		return ""
	}
	l := LstrlenW(str)
	if l == 0 {
		return ""
	}
	buff := make([]uint16, l)
	Memcpy(uintptr(unsafe.Pointer(&buff[0])), str, uintptr(l*2))
	return syscall.UTF16ToString(buff)
}

var _lstrlen = kernel32dll.NewProc("lstrlenA")

func Lstrlen(lpString uintptr) int32 {
	r, _, _ := _lstrlen.Call(lpString)
	return int32(r)
}

//全英文，目前正常
func GoAStr(str uintptr) string {
	if str == 0 {
		return ""
	}
	l := Lstrlen(str)
	if l == 0 {
		return ""
	}
	buff := make([]byte, l)
	Memcpy(uintptr(unsafe.Pointer(&buff[0])), str, uintptr(l))
	return string(buff)
}

//
//func CWStr(str string) uintptr {
//	if str == "" {
//		return 0
//	}
//	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str)))
//}
//
//func CAStr(str string) uintptr {
//	if str == "" {
//		return 0
//	}
//	return uintptr(unsafe.Pointer(&([]byte(str + "\x00"))[0]))
//}
//
//// 下面两个函数主要是给 386平台使用的，amd64不需要使用
//func ToUInt64(r1, r2 uintptr) uint64 {
//	ret := uint64(r2)
//	ret = uint64(ret<<32) + uint64(r1)
//	return ret
//}
//
//func UInt64To(val uint64) (uintptr, uintptr) {
//	return uintptr(uint32(val)), uintptr(uint32(val >> 32))
//}
