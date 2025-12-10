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
	playerImageScale      = 0.1
	playerInitialYRatio   = 0.95
	bulletInterval        = time.Duration(500 * time.Millisecond)
)

var (
	bullets []*Bullet
)

var _ Entity = &Player{}

type Player struct {
	x              float64
	y              float64
	width          float64
	height         float64
	lastBulletTime time.Time
}

func NewPlayer() *Player {
	width := float64(assets.PlayerImage.Bounds().Max.X) * playerImageScale
	height := float64(assets.PlayerImage.Bounds().Max.Y) * playerImageScale

	return &Player{
		x:              (internal.WindowWidth - width) / 2,
		y:              (internal.WindowHeight - height) * playerInitialYRatio,
		width:          width,
		height:         height,
		lastBulletTime: time.Now().Add(-bulletInterval),
	}
}

func (p *Player) move(newPoint lo.Tuple2[float64, float64]) {
	if newPoint.A < 0 {
		newPoint.A = 0
	}
	if newPoint.A > internal.WindowWidth-p.width {
		newPoint.A = internal.WindowWidth - p.width
	}
	if newPoint.B < 0 {
		newPoint.B = 0
	}
	if newPoint.B > internal.WindowHeight-p.height {
		newPoint.B = internal.WindowHeight - p.height
	}

	p.x = newPoint.A
	p.y = newPoint.B
}

func (p *Player) shoot() {
	if time.Since(p.lastBulletTime) >= bulletInterval {
		p.lastBulletTime = time.Now()
		bulletX := p.x + p.width/2
		bulletY := p.y - internal.WindowHeight*0.01
		bullets = append(bullets, NewBullet(bulletX, bulletY, 1))
	}
}

func (p *Player) Update(
	controllerButtonTouchEvent *ControllerButtonTouchEvent,
	shootButtonTouchEvent *ShootButtonTouchEvent,
) {
	keys := inpututil.AppendPressedKeys([]ebiten.Key{})

	if lo.Contains(keys, ebiten.KeyW) {
		p.move(lo.T2(p.x, p.y-playerDelta))
	}
	if lo.Contains(keys, ebiten.KeyA) {
		p.move(lo.T2(p.x-playerDelta, p.y))
	}
	if lo.Contains(keys, ebiten.KeyS) {
		p.move(lo.T2(p.x, p.y+playerDelta))
	}
	if lo.Contains(keys, ebiten.KeyD) {
		p.move(lo.T2(p.x+playerDelta, p.y))
	}

	if controllerButtonTouchEvent != nil {
		p.move(lo.T2(
			p.x+controllerButtonTouchEvent.Direction.A*playerControllerDelta,
			p.y+controllerButtonTouchEvent.Direction.B*playerControllerDelta,
		))
	}

	if lo.Contains(keys, ebiten.KeySpace) || shootButtonTouchEvent != nil {
		p.shoot()
	}

	lo.ForEach(bullets, func(bullet *Bullet, _ int) {
		bullet.Update()
	})
	bullets = lo.Filter(bullets, func(bullet *Bullet, _ int) bool {
		return bullet.y+bullet.height+5 > 0
	})
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(playerImageScale, playerImageScale)
	op.GeoM.Translate(p.x, p.y)
	screen.DrawImage(assets.PlayerImage, op)

	lo.ForEach(bullets, func(bullet *Bullet, _ int) {
		bullet.Draw(screen)
	})
}
