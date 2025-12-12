package scene

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/samber/lo"
)

const (
	DisplayInterval = time.Duration(1500 * time.Millisecond)
)

var _ Scene = &MenuScene{}

type MenuScene struct {
	sceneCreatedAt time.Time
}

func NewMenuScene() *MenuScene {
	return &MenuScene{
		sceneCreatedAt: time.Now(),
	}
}

func (s *MenuScene) Update() {
	pressedKeys := inpututil.AppendPressedKeys([]ebiten.Key{})
	touchIds := ebiten.AppendTouchIDs([]ebiten.TouchID{})

	if time.Since(s.sceneCreatedAt) > DisplayInterval && (lo.Contains(pressedKeys, ebiten.KeySpace) || len(touchIds) > 0) {
		CurrentScene = NewGameScene()
	}
}

func (s *MenuScene) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Press Space Key or Tap to Start")
}
