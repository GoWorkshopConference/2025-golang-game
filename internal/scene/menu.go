package scene

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var _ Scene = &MenuScene{}

type MenuScene struct {
	sceneCreatedAt time.Time
}

func NewMenuScene() Scene {
	return &MenuScene{
		sceneCreatedAt: time.Now(),
	}
}

func (s *MenuScene) Update() {
	sceneTransitionWithInterval(NewGameScene, s.sceneCreatedAt)
}

func (s *MenuScene) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Press Space Key or Tap to Start")
}
