package scene

import (
	"fmt"

	"github.com/GoWorkshopConference/golang-game/internal/entity"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var _ Scene = &GameScene{}

type GameScene struct {
	Score            int
	player           *entity.Player
	controllerButton *entity.ControllerButton
	shootButton      *entity.ShootButton
	enemyController  *entity.EbiFlyController
	gameController   *entity.GameController
}

func NewGameScene() *GameScene {
	return &GameScene{
		player:           entity.NewPlayer(),
		controllerButton: entity.NewControllerButton(),
		shootButton:      entity.NewShootButton(),
		enemyController:  entity.NewEbiFlyController(),
		gameController:   entity.NewGameController(),
	}
}

func (s *GameScene) Update() {
	controllerButtonTouchEvent := s.controllerButton.Update()
	shootButtonTouchEvent := s.shootButton.Update()
	s.player.Update(controllerButtonTouchEvent, shootButtonTouchEvent)

	s.enemyController.Update()
}

func (s *GameScene) Draw(screen *ebiten.Image) {
	s.enemyController.Draw(screen)
	s.player.Draw(screen)
	s.controllerButton.Draw(screen)
	s.shootButton.Draw(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", s.gameController.GetScore()))
}
