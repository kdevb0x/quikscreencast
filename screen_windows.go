// +build windows

package main

import (
	"unsafe"

	win "golang.org/x/sys/windows"
)

type HWND win.Handle

type HDC win.Handle

// global map of all the windows api system calls used
type wapi map[string]proc

type proc struct {
	dll     *win.LazyDLL
	p       *win.LazyProc
	cleanup func()
}

type screen struct {
	h   HWND
	dc  HDC
	err error

	// the windows api procedures
	wproc wapi
}

func NewScreen() *screen {
	return &screen{wproc: make(map[string]proc)}
}

func (s *screen) Error() string {
	return s.err.Error()
}

func (s *screen) Close() {
	for _, v := range s.wproc {
		v.cleanup()
	}
}

func (s *screen) openwWinHandle() error {

	lazyUser32 := win.NewLazySystemDLL("user32.dll")

	windowDC := lazyUser32.NewProc("GetDC")
	dc, _, _ := windowDC.Call(uintptr(unsafe.Pointer(new(HWND))))

	const null = 0
	if dc == null {
		panic("call to GetDC failed")
	}

	s.dc = *(*HDC)(unsafe.Pointer(dc))

	var cleanfunc = func() {
		r, _, err := s.wproc["GetDC"].dll.NewProc("ReleaseDC").Call(uintptr(unsafe.Pointer(&s.h)), uintptr(unsafe.Pointer(&s.dc)))

		if r == 0 {
			panic("Call to ReleaseDC failed: " + err.Error())
		}
	}
	s.wproc["GetDC"] = proc{lazyUser32, windowDC, cleanfunc}
	return nil
}

func free(p uintptr) {
	C.free(unsafe.Pointer(p))
}
