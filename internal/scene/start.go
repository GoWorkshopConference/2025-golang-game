package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/samber/lo"
)

type StartScene struct{}

func NewStartScene() *StartScene {
	return &StartScene{}
}

func (s *StartScene) Update() {
	pressedKeys := inpututil.AppendPressedKeys([]ebiten.Key{})
	if lo.Contains(pressedKeys, ebiten.KeySpace) {
		CurrentScene = NewGameScene()
	}
}

func (s *StartScene) Draw(screen *ebiten.Image) {

}
