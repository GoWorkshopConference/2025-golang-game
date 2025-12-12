package scene

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/samber/lo"
)

var _ Scene = &GameOverScene{}

type GameOverScene struct {
	score int
}

func NewGameOverScene(score int) *GameOverScene {
	return &GameOverScene{
		score: score,
	}
}

func (s *GameOverScene) Update() {
	pressedKeys := inpututil.AppendPressedKeys([]ebiten.Key{})
	if lo.Contains(pressedKeys, ebiten.KeySpace) {
		CurrentScene = NewMenuScene()
	}

	touchIds := ebiten.AppendTouchIDs([]ebiten.TouchID{})
	if len(touchIds) > 0 {
		CurrentScene = NewMenuScene()
	}
}

func (s *GameOverScene) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("GameOver... \nScore: %d\nPress Enter or Touch to Continue", s.score))
}
