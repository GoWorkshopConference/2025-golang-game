package scene

import (
	"time"

	"github.com/GoWorkshopConference/golang-game/internal/entity"
	"github.com/hajimehoshi/ebiten/v2"
)

type DebugScene struct {
	ebiFly        *entity.EbiFly
	player        *entity.Player
	ebiFlyRich    *entity.EbiFlyRich
	sauce         *entity.Sauce
	virus         *entity.Virus
	virusComputer *entity.VirusComputer

	sceneCreatedAt time.Time
}

func NewDebugScene() *DebugScene {
	return &DebugScene{
		ebiFly:        entity.NewEbiFly(10, 10),
		player:        entity.NewPlayerWithPos(10, 100),
		ebiFlyRich:    entity.NewEbiFlyRich(100, 10),
		sauce:         entity.NewSauce(100, 100),
		virus:         entity.NewVirus(200, 10),
		virusComputer: entity.NewVirusComputer(300, 10),

		sceneCreatedAt: time.Now(),
	}
}

func (s *DebugScene) Update() {
	sceneTransitionWithInterval(NewGameScene, s.sceneCreatedAt)
}

func (s *DebugScene) Draw(screen *ebiten.Image) {
	s.ebiFly.Draw(screen)
	s.player.Draw(screen)
	s.ebiFlyRich.Draw(screen)
	s.sauce.Draw(screen)
	s.virus.Draw(screen)
	s.virusComputer.Draw(screen)

	return
}
