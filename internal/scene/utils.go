package scene

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/samber/lo"
)

var (
	DisplayInterval = time.Duration(1500 * time.Millisecond)
)

func sceneTransitionWithIntervalAndDisplayInterval(makeNewScene func() Scene, createdAt time.Time, displayInterval time.Duration) {
	pressedKeys := inpututil.AppendPressedKeys([]ebiten.Key{})
	touchIds := ebiten.AppendTouchIDs([]ebiten.TouchID{})

	if time.Since(createdAt) > displayInterval &&
		(lo.Contains(pressedKeys, ebiten.KeySpace) || len(touchIds) > 0) {
		CurrentScene = makeNewScene()
	}
}

func sceneTransitionWithInterval(makeNewScene func() Scene, createdAt time.Time) {
	sceneTransitionWithIntervalAndDisplayInterval(makeNewScene, createdAt, DisplayInterval)
}

func drawText(screen *ebiten.Image, str string, x, y float64, fontSize float64, color color.Color) {
	op := &text.DrawOptions{}
	op.ColorScale.ScaleWithColor(color)
	op.LineSpacing = fontSize * 1.2
	op.GeoM.Translate(x, y)

	text.Draw(screen, str, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   fontSize,
	}, op)

}
