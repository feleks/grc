package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func getCursorPos() (int, int, error) {
	userDll := syscall.NewLazyDLL("user32.dll")
	getWindowRectProc := userDll.NewProc("GetCursorPos")
	type POINT struct {
		X, Y int32
	}
	var pt POINT
	_, _, eno := syscall.SyscallN(getWindowRectProc.Addr(), uintptr(unsafe.Pointer(&pt)))
	if eno != 0 {
		return 0, 0, fmt.Errorf("failed to get cursor pos: eno is 0")
	}

	return int(pt.X), int(pt.Y), nil
}
