package main

import (
	"context"
	"io"
)

// Screen represents a view of the screen.
type ScreenView interface {
	Output(w io.Writer, f Format)
	Start(ctx context.Context) (pause, stop func())
	Stopped() bool
}

type PausableScreen interface {
	ScreenView
	Pause()
	Paused() bool
}
type Format int

const (
	_ Format = iota
	RGBA
	MJPEG
)

func main() {

}
