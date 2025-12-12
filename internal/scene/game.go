package scene

import (
	"fmt"
	"time"

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
	virusController  *entity.VirusController
	gameController   *entity.GameController

	sceneCreatedAt time.Time
}

func NewGameScene() Scene {
	return &GameScene{
		player:           entity.NewPlayer(),
		controllerButton: entity.NewControllerButton(),
		shootButton:      entity.NewShootButton(),
		enemyController:  entity.NewEbiFlyController(),
		virusController:  entity.NewVirusController(),
		gameController:   entity.NewGameController(),

		sceneCreatedAt: time.Now(),
	}
}

func (s *GameScene) Update() {
	controllerButtonTouchEvent := s.controllerButton.Update()
	shootButtonTouchEvent := s.shootButton.Update()
	s.player.Update(controllerButtonTouchEvent, shootButtonTouchEvent)

	s.enemyController.Update()
	s.virusController.Update()

	if s.gameController.GetLife() <= 0 {
		CurrentScene = NewGameOverScene(s.gameController.GetScore())
	}
}

func (s *GameScene) Draw(screen *ebiten.Image) {
	s.enemyController.Draw(screen)
	s.virusController.Draw(screen)
	s.player.Draw(screen)
	s.controllerButton.Draw(screen)
	s.shootButton.Draw(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d, Life: %d", s.gameController.GetScore(), s.gameController.GetLife()))
}
