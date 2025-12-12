package entity

import (
	"github.com/GoWorkshopConference/golang-game/internal/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/samber/lo"
)

const (
	virusDelta      = 3
	virusImageScale = 0.1
)

var _ Entity = &Virus{}
var _ VirusLike = &Virus{}

type Virus struct {
	Shape     *Shape
	HitBox    *Shape
	isRemoved bool
}

func NewVirus(posX, posY float64) *Virus {
	width := float64(assets.VirusImage.Bounds().Max.X) * virusImageScale
	height := float64(assets.VirusImage.Bounds().Max.Y) * virusImageScale

	return &Virus{
		Shape: &Shape{
			X:      posX,
			Y:      posY,
			Width:  width,
			Height: height,
		},
		HitBox: &Shape{
			X:      posX + 18,
			Y:      posY + 12,
			Width:  width - 35,
			Height: height - 15,
		},
		isRemoved: false,
	}
}

func (v *Virus) Update() {
	v.Shape.Y += virusDelta
	v.HitBox.Y += virusDelta
}

func (v *Virus) Draw(screen *ebiten.Image) {
	draw(screen, v, assets.VirusImage, lo.T2(virusImageScale, virusImageScale))
}

func (v *Virus) GetShape() *Shape {
	return v.Shape
}

func (v *Virus) GetHitBox() *Shape {
	return v.HitBox
}

func (v *Virus) HitPlayer() {
	v.isRemoved = true
}

func (v *Virus) GetIsRemoved() bool {
	return v.isRemoved
}
