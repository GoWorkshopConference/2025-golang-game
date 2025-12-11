package entity

import (
	"github.com/GoWorkshopConference/golang-game/internal/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/samber/lo"
)

var (
	sauceImageScale = 0.07
	sauceDelta      = 6.
)

var _ Entity = &Sauce{}

type Sauce struct {
	Shape  *Shape
	HitBox *Shape
	image  *ebiten.Image
	isHit  bool
}

func NewSauceFromLeftUpperPos(posX, posY float64) *Sauce {
	width := float64(assets.SauceImage.Bounds().Max.X) * sauceImageScale
	height := float64(assets.SauceImage.Bounds().Max.Y) * sauceImageScale

	return NewSauce(posX-width/2, posY-height/2)
}

func NewSauce(posX, posY float64) *Sauce {
	width := float64(assets.SauceImage.Bounds().Max.X) * sauceImageScale
	height := float64(assets.SauceImage.Bounds().Max.Y) * sauceImageScale

	return &Sauce{
		image: assets.SauceImage,
		Shape: &Shape{
			X:      posX,
			Y:      posY,
			Width:  width,
			Height: height,
		},
		HitBox: &Shape{
			X:      posX + 15,
			Y:      posY,
			Width:  width - 30,
			Height: height,
		},
		isHit: false,
	}
}

func (s *Sauce) Update() {
	s.Shape.Y -= sauceDelta
	s.HitBox.Y -= sauceDelta

	s.isHit = s.checkHitEbiFly()
}

func (s *Sauce) checkHitEbiFly() (isHit bool) {
	isHit = false
	addedEbiFlies := []EbiFlyLike{}

	ebiFliesLock.Lock()
	ebiFlies = lo.Filter(ebiFlies, func(ebiFly EbiFlyLike, _ int) bool {
		isHit = EntityHit(*s.HitBox, *ebiFly.GetHitBox())
		if !isHit {
			return true
		}

		if ebiFly, ok := ebiFly.(*EbiFly); ok {
			addedEbiFlies = append(addedEbiFlies, ebiFly.UpgradeEbiFly())
			return false
		}

		return true
	})
	ebiFlies = append(ebiFlies, addedEbiFlies...)
	ebiFliesLock.Unlock()

	return isHit
}

func (s *Sauce) Draw(screen *ebiten.Image) {
	draw(screen, s, s.image, lo.T2(sauceImageScale, sauceImageScale))
}

func (s *Sauce) GetShape() *Shape {
	return s.Shape
}

func (s *Sauce) GetHitBox() *Shape {
	return s.HitBox
}
