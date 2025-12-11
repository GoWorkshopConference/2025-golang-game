package entity

import (
	"github.com/GoWorkshopConference/golang-game/internal"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/samber/lo"
)

type Entity interface {
	Draw(screen *ebiten.Image)
	GetShape() *Shape
	GetHitBox() *Shape
}

type Shape struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func draw(screen *ebiten.Image, entity Entity, image *ebiten.Image, scale lo.Tuple2[float64, float64]) {
	shape := entity.GetShape()
	hitBox := entity.GetHitBox()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale.A, scale.B)
	op.GeoM.Translate(shape.X, shape.Y)
	screen.DrawImage(image, op)

	if internal.IsDebugMode {
		vector.StrokeRect(screen,
			float32(hitBox.X), float32(hitBox.Y),
			float32(hitBox.Width), float32(hitBox.Height),
			2, internal.DebugColor, true)
	}
}
