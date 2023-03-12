//go:build !windows
// +build !windows

package ffcommon

import "reflect"

func NewCallback(fn interface{}) uintptr {
	if fn == nil {
		return uintptr(0)
	} else {
		//未测试，不一定行
		return reflect.ValueOf(fn).Pointer()
	}
}
