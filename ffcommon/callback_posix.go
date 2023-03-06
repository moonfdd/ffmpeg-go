//go:build !windows
// +build !windows

package ffcommon

func NewCallback(fn interface{}) uintptr {
	if fn == nil {
		return uintptr(0)
	} else {
		panic("not support")
	}
}
