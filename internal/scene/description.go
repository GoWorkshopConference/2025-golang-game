package scene

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var _ Scene = &DescriptionScene{}

type DescriptionScene struct {
	sceneCreatedAt time.Time
}

func NewDescriptionScene() *DescriptionScene {
	return &DescriptionScene{
		sceneCreatedAt: time.Now(),
	}
}

func (s *DescriptionScene) Update() {
	sceneTransitionWithInterval(NewGameScene, s.sceneCreatedAt)
}

func (s *DescriptionScene) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "DescriptionScene")
}
