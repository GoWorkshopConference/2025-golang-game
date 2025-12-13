package scene

import (
	"image/color"
	"time"

	"github.com/GoWorkshopConference/golang-game/internal"
	"github.com/GoWorkshopConference/golang-game/internal/entity"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/samber/lo"
)

var _ Scene = &DescriptionScene{}

var (
	descriptionBaseLineY = 120.
)

type DescriptionScene struct {
	sceneCreatedAt time.Time
	entities       []entity.Entity
}

func NewDescriptionScene() Scene {
	return &DescriptionScene{
		sceneCreatedAt: time.Now(),
		entities: []entity.Entity{
			entity.NewEbiFly(baseLineX+300, descriptionBaseLineY+45),
			entity.NewSauce(baseLineX+305, descriptionBaseLineY+125),
			entity.NewVirus(baseLineX+295, descriptionBaseLineY+185),
			entity.NewVirusComputer(baseLineX+310, descriptionBaseLineY+255),
		},
	}
}

func (s *DescriptionScene) Update() {
	sceneTransitionWithIntervalAndDisplayInterval(NewGameScene, s.sceneCreatedAt, time.Duration(500*time.Millisecond))
}

func (s *DescriptionScene) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "DescriptionScene")

	drawText(screen, "エビフライをいっぱい食べよう", baseLineX-40, descriptionBaseLineY+70, 16.0, color.White)
	drawText(screen, "ソースをかけると得点アップ！", baseLineX-40, descriptionBaseLineY+140, 16.0, color.White)
	drawText(screen, "ウイルスにぶつかるとライフ-1", baseLineX-40, descriptionBaseLineY+210, 16.0, color.White)
	drawText(screen, "コンピュータウイルスに当たると即アウト", baseLineX-40, descriptionBaseLineY+280, 16.0, color.White)

	if time.Since(s.sceneCreatedAt) > time.Duration(500*time.Millisecond) {
		drawText(screen, "スタート", baseLineX+120, descriptionBaseLineY+340, 16.0, internal.EbitenColor)
	}

	lo.ForEach(s.entities, func(entity entity.Entity, _ int) {
		entity.Draw(screen)
	})
}
