package scene

import (
	"github.com/GoWorkshopConference/golang-game/internal/entity"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/samber/lo"
)

type DebugScene struct {
	ebiFly     *entity.EbiFly
	player     *entity.Player
	ebiFlyRich *entity.EbiFlyRich
	sauce      *entity.Sauce
}

func NewDebugScene() *DebugScene {
	return &DebugScene{
		ebiFly:     entity.NewEbiFly(10, 10),
		player:     entity.NewPlayerWithPos(10, 100),
		ebiFlyRich: entity.NewEbiFlyRich(100, 10),
		sauce:      entity.NewSauce(100, 100),
	}
}

func (s *DebugScene) Update() {
	pressedKeys := inpututil.AppendPressedKeys([]ebiten.Key{})
	if lo.Contains(pressedKeys, ebiten.KeySpace) {
		CurrentScene = NewGameScene()
	}
}

func (s *DebugScene) Draw(screen *ebiten.Image) {
	s.ebiFly.Draw(screen)
	s.player.Draw(screen)
	s.ebiFlyRich.Draw(screen)
	s.sauce.Draw(screen)

	return
}
