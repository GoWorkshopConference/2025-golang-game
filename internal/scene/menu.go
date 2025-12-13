package scene

import (
	"bytes"
	"image/color"
	"log"
	"time"

	"github.com/GoWorkshopConference/golang-game/internal"
	"github.com/GoWorkshopConference/golang-game/internal/entity"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	logoY     = 230.
	baseLineX = 90.
)

var _ Scene = &MenuScene{}

type MenuScene struct {
	sceneCreatedAt time.Time

	logoEbiFly *entity.EbiFly
}

func NewMenuScene() Scene {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s

	return &MenuScene{
		sceneCreatedAt: time.Now(),
		logoEbiFly:     entity.NewEbiFly(baseLineX+240, logoY),
	}
}

func (s *MenuScene) Update() {
	sceneTransitionWithInterval(NewDescriptionScene, s.sceneCreatedAt)
}

func (s *MenuScene) Draw(screen *ebiten.Image) {
	drawText(screen, "えびフライ", baseLineX, logoY, 48.0, internal.EbitenColor)

	if time.Since(s.sceneCreatedAt) > time.Duration(500*time.Millisecond) {
		drawText(screen, "スペースキーを押すか", baseLineX+70, logoY+70, 16.0, color.White)

		drawText(screen, "画面をタップしてゲームスタート", baseLineX+30, logoY+90, 16.0, color.White)
	}

	s.logoEbiFly.Draw(screen)
}
