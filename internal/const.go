package internal

import "image/color"

const (
	WindowWidth  = 480
	WindowHeight = 640
)

var (
	DebugColor  = color.RGBA{0xff, 0x20, 0xAA, 0xff}
	EbitenColor = color.RGBA{219, 87, 31, 0xff}
)

// 環境によって変わる変数
var (
	IsDebugMode = false
)
