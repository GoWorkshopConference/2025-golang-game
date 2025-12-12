package entity

import (
	"math/rand/v2"
	"sync"
	"time"

	"github.com/GoWorkshopConference/golang-game/internal"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/samber/lo"
)

type EbiFlyLike interface {
	Update()
	Draw(screen *ebiten.Image)
	GetShape() *Shape
	GetHitBox() *Shape
	HitPlayer()
	GetIsRemoved() bool
}

var (
	ebiFliesLock              = sync.Mutex{}
	ebiFlies     []EbiFlyLike = []EbiFlyLike{}
)

func addEbiFly(ebiFly EbiFlyLike) {
	ebiFliesLock.Lock()
	ebiFlies = append(ebiFlies, ebiFly)
	ebiFliesLock.Unlock()
}

type EbiFlyController struct {
	lastSpawnTime time.Time
}

func NewEbiFlyController() *EbiFlyController {
	ebiFlies = []EbiFlyLike{}

	return &EbiFlyController{
		lastSpawnTime: time.Now().Add(-ebiFlySpawnInterval),
	}
}

func (e *EbiFlyController) Update() {
	// ランダムに敵を生成
	if rand.Float64() < ebiFlySpawnRate {
		e.lastSpawnTime = time.Now()

		// ランダムなX座標で敵を生成（画面幅内）
		randX := rand.Float64() * (internal.WindowWidth - 50)
		addEbiFly(NewEbiFly(randX, ebiFlyInitialY))
	}

	// エビフライ達の更新
	lo.ForEach(ebiFlies, func(enemy EbiFlyLike, _ int) {
		enemy.Update()
	})

	// 画面外に出たエビフライを削除
	ebiFliesLock.Lock()
	ebiFlies = lo.Filter(ebiFlies, func(enemy EbiFlyLike, _ int) bool {
		return enemy.GetShape().Y < internal.WindowHeight+50
	})
	ebiFliesLock.Unlock()
}

func (e *EbiFlyController) Draw(screen *ebiten.Image) {
	lo.ForEach(ebiFlies, func(enemy EbiFlyLike, _ int) {
		enemy.Draw(screen)
	})
}
