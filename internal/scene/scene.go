package scene

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
}

type SceneType string

var (
	SceneTypeDebug SceneType = "DEBUG_SCENE"
	SceneTypeGame  SceneType = "GAME_SCENE"
)

var CurrentScene Scene

var (
	mplusFaceSource *text.GoTextFaceSource
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s
}
