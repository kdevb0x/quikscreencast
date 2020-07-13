// +build !windows

package main

import "io"

type screen struct {
	out       io.Writer
	outFormat Format
}

func newScreen() *screen {
	return new(screen)
}

// OutOutput sets the output where image data is written.
func (s *screen) Output(w io.Writer, format Format) {
	s.out = w
	s.outFormat = format
}
