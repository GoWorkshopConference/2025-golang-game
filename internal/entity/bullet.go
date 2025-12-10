package entity

import (
	"github.com/GoWorkshopConference/golang-game/internal/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	bulletImageScale = 0.05
	bulletDelta      = 10
)

type Bullet struct {
	x      float64
	y      float64
	width  float64
	height float64
}

func NewBullet(posX, posY float64) *Bullet {
	return &Bullet{
		x:      posX,
		y:      posY,
		width:  float64(assets.EbiFryImage.Bounds().Max.X) * bulletImageScale,
		height: float64(assets.EbiFryImage.Bounds().Max.Y) * bulletImageScale,
	}
}

func (b *Bullet) Update() {
	b.y -= bulletDelta
	return
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(bulletImageScale, bulletImageScale)
	op.GeoM.Translate(b.x, b.y)
	screen.DrawImage(assets.EbiFryImage, op)

	return
}
