package scene

import (
	"fmt"
	"log"
	"time"

	"github.com/GoWorkshopConference/golang-game/internal"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var _ Scene = &GameOverScene{}

type GameOverScene struct {
	sceneCreatedAt time.Time
	score          int
}

func NewGameOverScene(score int) *GameOverScene {
	allData := internal.GetAllParentLocalStorage()
	log.Printf("allData: %+v", allData)

	return &GameOverScene{
		score:          score,
		sceneCreatedAt: time.Now(),
	}
}

func (s *GameOverScene) Update() {
	sceneTransitionWithInterval(NewMenuScene, s.sceneCreatedAt)
}

func (s *GameOverScene) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("GameOver... \nScore: %d\nPress Enter or Touch to Continue", s.score))
}
