//go:build windows
// +build windows

package ffcommon

import (
	"syscall"
)

func NewCallback(fn interface{}) uintptr {
	u := syscall.NewCallback(fn)
	return u
}
