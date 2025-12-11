package entity

import (
	"github.com/GoWorkshopConference/golang-game/internal/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/samber/lo"
)

var (
	ebiFlyRichImageScale = 0.15
)

var _ Entity = &EbiFlyRich{}
var _ EbiFlyLike = &EbiFlyRich{}

type EbiFlyRich struct {
	Shape     *Shape
	HitBox    *Shape
	isRemoved bool
}

func NewEbiFlyRichFromCenterPos(posX, posY float64) *EbiFlyRich {
	width := float64(assets.EbiFryRichImage.Bounds().Max.X) * ebiFlyRichImageScale
	height := float64(assets.EbiFryRichImage.Bounds().Max.Y) * ebiFlyRichImageScale

	return NewEbiFlyRich(posX-width/2, posY-height/2)
}

func NewEbiFlyRich(posX, posY float64) *EbiFlyRich {
	width := float64(assets.EbiFryRichImage.Bounds().Max.X) * ebiFlyRichImageScale
	height := float64(assets.EbiFryRichImage.Bounds().Max.Y) * ebiFlyRichImageScale

	return &EbiFlyRich{
		Shape: &Shape{
			X:      posX,
			Y:      posY,
			Width:  width,
			Height: height,
		},
		HitBox: &Shape{
			X:      posX,
			Y:      posY,
			Width:  width,
			Height: height,
		},
	}
}

func (e *EbiFlyRich) Update() {
	e.Shape.Y += ebiFlyDelta
	e.HitBox.Y += ebiFlyDelta
}

func (e *EbiFlyRich) Draw(screen *ebiten.Image) {
	draw(screen, e, assets.EbiFryRichImage, lo.T2(ebiFlyRichImageScale, ebiFlyRichImageScale))
}

func (e *EbiFlyRich) GetShape() *Shape {
	return e.Shape
}

func (e *EbiFlyRich) GetHitBox() *Shape {
	return e.HitBox
}

func (e *EbiFlyRich) HitPlayer() {
	e.isRemoved = true
	score += 500
}

func (e *EbiFlyRich) GetIsRemoved() bool {
	return e.isRemoved
}
