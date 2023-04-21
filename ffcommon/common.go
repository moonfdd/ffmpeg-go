package ffcommon

import (
	"syscall"
	"unsafe"
)

func BytePtrFromString(str string) (res *byte) {
	res, _ = syscall.BytePtrFromString(str)
	return
}

func UintPtrFromString(str string) uintptr {
	if str == "" {
		return uintptr(0)
	}
	return uintptr(unsafe.Pointer(BytePtrFromString(str)))
}

func UintPtrFromContainsEmptyString(str string) uintptr {
	return uintptr(unsafe.Pointer(BytePtrFromString(str)))
}

//func BoolFromUintptr(ptr uintptr) bool {
//	if ptr == 0 {
//		return false
//	}
//	return true
//}

func StringFromPtr(sptr uintptr) (res string) {
	if sptr <= 0 {
		return
	}
	buf := make([]byte, 0)
	for i := 0; *(*byte)(unsafe.Pointer(sptr + uintptr(i))) != 0; i++ {
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

func ByteSliceFromByteP(data *byte, len0 int) []byte {
	if data == nil {
		return nil
	}
	if len0 == 0 {
		return []byte{}
	}

	var sliceHeader sliceHeader
	sliceHeader.Data = data
	sliceHeader.Len = len0
	sliceHeader.Cap = len0
	return *(*[]byte)(unsafe.Pointer(&sliceHeader))
}

type sliceHeader struct {
	Data *byte
	Len  int
	Cap  int
}
