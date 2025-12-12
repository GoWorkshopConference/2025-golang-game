package entity

import (
	"math/rand/v2"
	"sync"

	"github.com/GoWorkshopConference/golang-game/internal"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/samber/lo"
)

const (
	virusSpawnRate    = 0.015
	computerSpawnRate = 0.1
)

var (
	virusLock             = sync.Mutex{}
	viruses   []VirusLike = []VirusLike{}
)

type VirusLike interface {
	Update()
	Draw(screen *ebiten.Image)
	GetShape() *Shape
	GetHitBox() *Shape
	HitPlayer()
	GetIsRemoved() bool
}

type VirusController struct{}

func NewVirusController() *VirusController {
	return &VirusController{}
}

func addVirus(virus VirusLike) {
	virusLock.Lock()
	viruses = append(viruses, virus)
	virusLock.Unlock()
}

func (v *VirusController) Update() {
	// ランダムに敵を生成
	if rand.Float64() < virusSpawnRate {
		// ランダムなX座標で敵を生成（画面幅内）
		randX := rand.Float64() * (internal.WindowWidth - 50)

		if rand.Float64() < computerSpawnRate {
			addVirus(NewVirusComputer(randX, -10.0))
		} else {
			addVirus(NewVirus(randX, -10.0))
		}
	}

	// ウイルス達の更新
	lo.ForEach(viruses, func(virus VirusLike, _ int) {
		virus.Update()
	})

	// 画面外に出たウイルスを削除
	ebiFliesLock.Lock()
	ebiFlies = lo.Filter(ebiFlies, func(enemy EbiFlyLike, _ int) bool {
		return enemy.GetShape().Y < internal.WindowHeight+50
	})
	ebiFliesLock.Unlock()
}

func (v *VirusController) Draw(screen *ebiten.Image) {
	lo.ForEach(viruses, func(virus VirusLike, _ int) {
		virus.Draw(screen)
	})
}
