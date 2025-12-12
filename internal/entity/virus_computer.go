package entity

import (
	"github.com/GoWorkshopConference/golang-game/internal/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/samber/lo"
)

const (
	virusComputerDelta      = 3
	virusComputerImageScale = 0.06
)

var _ Entity = &VirusComputer{}
var _ VirusLike = &VirusComputer{}

type VirusComputer struct {
	Shape     *Shape
	HitBox    *Shape
	isRemoved bool
}

func NewVirusComputer(posX, posY float64) *VirusComputer {
	width := float64(assets.VirusComputerImage.Bounds().Max.X) * virusComputerImageScale
	height := float64(assets.VirusComputerImage.Bounds().Max.Y) * virusComputerImageScale

	return &VirusComputer{
		Shape: &Shape{
			X:      posX,
			Y:      posY,
			Width:  width,
			Height: height,
		},
		HitBox: &Shape{
			X:      posX + 5,
			Y:      posY + 7,
			Width:  width - 12,
			Height: height - 10,
		},
		isRemoved: false,
	}
}

func (v *VirusComputer) Update() {
	v.Shape.Y += virusComputerDelta
	v.HitBox.Y += virusComputerDelta
}

func (v *VirusComputer) Draw(screen *ebiten.Image) {
	draw(screen, v, assets.VirusComputerImage, lo.T2(virusComputerImageScale, virusComputerImageScale))
}

func (v *VirusComputer) GetShape() *Shape {
	return v.Shape
}

func (v *VirusComputer) GetHitBox() *Shape {
	return v.HitBox
}

func (v *VirusComputer) HitPlayer() {
	v.isRemoved = true
}

func (v *VirusComputer) GetIsRemoved() bool {
	return v.isRemoved
}
