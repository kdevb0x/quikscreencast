package main

import (
	"image"
)

type Capturer interface {
	Capture() ([]byte, error)
}

type ImageCapturer interface {
	CaptureImage() (image.Image, error)
}
