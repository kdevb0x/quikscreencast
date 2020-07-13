// +build linux

package main

import (
	"context"
	"io"
)

// Screen represents a view of the screen.
type Screen interface {
	Output(w io.Writer, f Format)
	Start(ctx context.Context) (pause, stop func())
	Paused() bool
	Stopped() bool
}

type Format int

const (
	_ Format = iota
	RGBA
	MJPEG
)

func main() {

}
