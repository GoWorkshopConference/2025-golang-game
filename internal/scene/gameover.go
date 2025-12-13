package scene

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/GoWorkshopConference/golang-game/internal"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	gameOverBaseLineY = 200.
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
	drawText(screen, "GameOver...", 100, gameOverBaseLineY, 48.0, internal.EbitenColor)
	drawText(screen, fmt.Sprintf("スコア： %d 点", s.score), 100, gameOverBaseLineY+70, 24.0, color.White)
	drawText(screen, fmt.Sprintf("生存時間： %.2f 秒", s.lifeTime.Seconds()), 100, gameOverBaseLineY+100, 24.0, color.White)

	if time.Since(s.sceneCreatedAt) > DisplayInterval {
		drawText(screen, "タップするかエンターキーで再挑戦！", baseLineX-20, gameOverBaseLineY+200, 20.0, internal.EbitenColor)
	}
}
