package entity

import (
	_ "embed"
	_ "image/png"
	"time"

	"github.com/GoWorkshopConference/golang-game/internal"
	"github.com/GoWorkshopConference/golang-game/internal/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/samber/lo"
)

const (
	playerDelta           = 5
	playerControllerDelta = 8
	playerImageScale      = 0.15
	playerInitialYRatio   = 0.95
	shootInterval         = time.Duration(1000 * time.Millisecond)

	hitBoxOffset = 20.0
)

var (
	sauces []*Sauce
)

var _ Entity = &Player{}

type Player struct {
	Shape         *Shape
	HitBox        *Shape
	lastShootTime time.Time
}

func NewPlayer() *Player {
	width := float64(assets.PlayerImage.Bounds().Max.X) * playerImageScale
	height := float64(assets.PlayerImage.Bounds().Max.Y) * playerImageScale

	return NewPlayerWithPos(
		(internal.WindowWidth-width)/2,
		(internal.WindowHeight-height)*playerInitialYRatio)
}

func NewPlayerWithPos(posX, posY float64) *Player {
	width := float64(assets.PlayerImage.Bounds().Max.X) * playerImageScale
	height := float64(assets.PlayerImage.Bounds().Max.Y) * playerImageScale

	return &Player{
		Shape: &Shape{
			X:      posX,
			Y:      posY,
			Width:  width,
			Height: height,
		},
		HitBox: &Shape{
			X:      posX + hitBoxOffset,
			Y:      posY,
			Width:  width - (hitBoxOffset * 2),
			Height: height,
		},
		lastShootTime: time.Now().Add(-shootInterval),
	}
}

func (p *Player) move(newPoint lo.Tuple2[float64, float64]) {
	if newPoint.A < 0 {
		newPoint.A = 0
	}
	if newPoint.A > internal.WindowWidth-p.Shape.Width {
		newPoint.A = internal.WindowWidth - p.Shape.Width
	}
	if newPoint.B < 0 {
		newPoint.B = 0
	}
	if newPoint.B > internal.WindowHeight-p.Shape.Height {
		newPoint.B = internal.WindowHeight - p.Shape.Height
	}

	p.Shape.X = newPoint.A
	p.Shape.Y = newPoint.B
	p.HitBox.X = newPoint.A + hitBoxOffset
	p.HitBox.Y = newPoint.B
}

func (p *Player) shoot() {
	if time.Since(p.lastShootTime) >= shootInterval {
		p.lastShootTime = time.Now()
		sauceX := p.HitBox.X + p.HitBox.Width/2
		sauceY := p.Shape.Y - internal.WindowHeight*0.01
		sauces = append(sauces, NewSauceFromLeftUpperPos(sauceX, sauceY))
	}
}

func (p *Player) Update(
	controllerButtonTouchEvent *ControllerButtonTouchEvent,
	shootButtonTouchEvent *ShootButtonTouchEvent,
) {
	keys := inpututil.AppendPressedKeys([]ebiten.Key{})

	if lo.Contains(keys, ebiten.KeyW) {
		p.move(lo.T2(p.Shape.X, p.Shape.Y-playerDelta))
	}
	if lo.Contains(keys, ebiten.KeyA) {
		p.move(lo.T2(p.Shape.X-playerDelta, p.Shape.Y))
	}
	if lo.Contains(keys, ebiten.KeyS) {
		p.move(lo.T2(p.Shape.X, p.Shape.Y+playerDelta))
	}
	if lo.Contains(keys, ebiten.KeyD) {
		p.move(lo.T2(p.Shape.X+playerDelta, p.Shape.Y))
	}

	if controllerButtonTouchEvent != nil {
		p.move(lo.T2(
			p.Shape.X+controllerButtonTouchEvent.Direction.A*playerControllerDelta,
			p.Shape.Y+controllerButtonTouchEvent.Direction.B*playerControllerDelta,
		))
	}

	if lo.Contains(keys, ebiten.KeySpace) || shootButtonTouchEvent != nil {
		p.shoot()
	}

	lo.ForEach(sauces, func(sauce *Sauce, _ int) {
		sauce.Update()
	})
	sauces = lo.Filter(sauces, func(sauce *Sauce, _ int) bool {
		return sauce.Shape.Y+sauce.Shape.Height+5 > 0 && !sauce.isHit
	})

	p.checkHitEbiFly()
}

func (p *Player) checkHitEbiFly() {
	ebiFliesLock.Lock()
	for _, ebiFly := range ebiFlies {
		isHit := EntityHit(*p.HitBox, *ebiFly.GetHitBox())
		if isHit {
			p.HitEbiFly()
			ebiFly.HitPlayer()
		}
	}
	ebiFlies = lo.Filter(ebiFlies, func(ebiFly EbiFlyLike, _ int) bool {
		return !ebiFly.GetIsRemoved()
	})
	ebiFliesLock.Unlock()
}

func (p *Player) HitEbiFly() {}

func (p *Player) Draw(screen *ebiten.Image) {
	draw(screen, p, assets.PlayerImage, lo.T2(playerImageScale, playerImageScale))

	lo.ForEach(sauces, func(sauce *Sauce, _ int) {
		sauce.Draw(screen)
	})
}

func (p *Player) GetShape() *Shape {
	return p.Shape
}

func (p *Player) GetHitBox() *Shape {
	return p.HitBox
}
