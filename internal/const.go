package internal

import "image/color"

const (
	WindowWidth  = 480
	WindowHeight = 640
)

var (
	DebugColor  = color.RGBA{0xff, 0x20, 0xAA, 0xff}
	IsDebugMode = false
)
