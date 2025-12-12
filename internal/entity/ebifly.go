package entity

import (
	"time"

	"github.com/GoWorkshopConference/golang-game/internal/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/samber/lo"
)

const (
	ebiFlySpawnRate     = 0.015
	ebiFlyImageScale    = 0.09
	ebiFlySpawnInterval = time.Duration(1 * time.Second)
	ebiFlyDelta         = 3
	ebiFlyInitialY      = -50
)

var _ Entity = &EbiFly{}
var _ EbiFlyLike = &EbiFly{}

type EbiFly struct {
	Shape     *Shape
	HitBox    *Shape
	isRemoved bool
}

func NewEbiFly(posX, posY float64) *EbiFly {
	width := float64(assets.EbiFryImage.Bounds().Max.X) * ebiFlyImageScale
	height := float64(assets.EbiFryImage.Bounds().Max.Y) * ebiFlyImageScale

	return &EbiFly{
		Shape: &Shape{
			X:      posX,
			Y:      posY,
			Width:  width,
			Height: height,
		},
		HitBox: &Shape{
			X:      posX + 2,
			Y:      posY + 3,
			Width:  width - 2,
			Height: height - 8,
		},
		isRemoved: false,
	}
}

func (e *EbiFly) UpgradeEbiFly() *EbiFlyRich {
	return NewEbiFlyRichFromCenterPos(
		e.Shape.X+e.Shape.Width/2,
		e.Shape.Y+e.Shape.Height/2,
	)
}

func (e *EbiFly) Update() {
	e.Shape.Y += ebiFlyDelta
	e.HitBox.Y += ebiFlyDelta
}

func (e *EbiFly) Draw(screen *ebiten.Image) {
	draw(screen, e, assets.EbiFryImage, lo.T2(ebiFlyImageScale, ebiFlyImageScale))
}

func (e *EbiFly) GetShape() *Shape {
	return e.Shape
}

func (e *EbiFly) GetHitBox() *Shape {
	return e.HitBox
}

func (e *EbiFly) HitPlayer() {
	e.isRemoved = true
	score += 100
}

func (e *EbiFly) GetIsRemoved() bool {
	return e.isRemoved
}
