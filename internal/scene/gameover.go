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
	lifeTime       time.Duration
}

func NewGameOverScene(score int, lifeTime time.Duration) *GameOverScene {
	allData := internal.GetAllParentLocalStorage()
	log.Printf("allData: %+v", allData)

	// プレイヤー名を取得
	username := "unknown"
	if playerName, ok := allData["playerName"]; ok && playerName != "" {
		username = playerName
	}

	// サーバーにスコアを送信（非同期）
	lifeTimeSeconds := int(lifeTime.Seconds())
	if err := internal.SendScoreToServer(username, score, lifeTimeSeconds); err != nil {
		log.Printf("スコアの送信に失敗しました: %v", err)
	} else {
		log.Printf("スコアを送信しました: username=%s, score=%d, lifeTime=%d", username, score, lifeTimeSeconds)
	}

	return &GameOverScene{
		score:          score,
		sceneCreatedAt: time.Now(),
		lifeTime:       lifeTime,
	}
}

func (s *GameOverScene) Update() {
	sceneTransitionWithInterval(NewMenuScene, s.sceneCreatedAt)
}

func (s *GameOverScene) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("GameOver... \nScore: %d,\nLifeTime: %.2f (s)\nPress Enter or Touch to Continue", s.score, s.lifeTime.Seconds()))
}
