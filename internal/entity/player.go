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
	playerDelta         = 5
	playerImageScale    = 0.1
	playerInitialYRatio = 0.95
	bulletInterval      = time.Duration(500 * time.Millisecond)
)

var (
	bullets []*Bullet
)

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

func (p *Player) Update() {
	keys := inpututil.AppendPressedKeys([]ebiten.Key{})

	if lo.Contains(keys, ebiten.KeyA) {
		p.x = max(0, p.x-playerDelta)
	}
	if lo.Contains(keys, ebiten.KeyD) {
		p.x = min(internal.WindowWidth-p.width, p.x+playerDelta)
	}
	if lo.Contains(keys, ebiten.KeyW) {
		p.y = max(0, p.y-playerDelta)
	}
	if lo.Contains(keys, ebiten.KeyS) {
		p.y = min(internal.WindowHeight-p.height, p.y+playerDelta)
	}

	if lo.Contains(keys, ebiten.KeySpace) &&
		time.Since(p.lastBulletTime) >= bulletInterval {
		p.lastBulletTime = time.Now()
		bulletX := p.x + p.width/2
		bulletY := p.y - internal.WindowHeight*0.01
		bullets = append(bullets, NewBullet(bulletX, bulletY))
	}

	lo.ForEach(bullets, func(bullet *Bullet, _ int) {
		bullet.Update()
	})
	bullets = lo.Filter(bullets, func(bullet *Bullet, _ int) bool {
		return bullet.y > 0
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
