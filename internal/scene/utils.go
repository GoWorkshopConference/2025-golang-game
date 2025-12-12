package scene

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/samber/lo"
)

var (
	DisplayInterval = time.Duration(1500 * time.Millisecond)
)

func sceneTransitionWithInterval(makeNewScene func() Scene, createdAt time.Time) {
	pressedKeys := inpututil.AppendPressedKeys([]ebiten.Key{})
	touchIds := ebiten.AppendTouchIDs([]ebiten.TouchID{})

	if time.Since(createdAt) > DisplayInterval &&
		(lo.Contains(pressedKeys, ebiten.KeySpace) || len(touchIds) > 0) {
		CurrentScene = makeNewScene()
	}
}
