//go:build windows
// +build windows

package ffcommon

import (
	"syscall"
)

func NewCallback(fn interface{}) uintptr {
	if fn == nil {
		return uintptr(0)
	}
	u := syscall.NewCallback(fn)
	return u
}
