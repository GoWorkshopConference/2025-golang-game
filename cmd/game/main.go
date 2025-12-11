package main

import (
	"log"
	"os"

	"github.com/GoWorkshopConference/golang-game/internal"
	"github.com/GoWorkshopConference/golang-game/internal/scene"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/samber/lo"
)

type Game struct{}

func NewGame() *Game {
	if internal.IsDebugMode {
		scene.CurrentScene = scene.NewDebugScene()
	} else {
		scene.CurrentScene = scene.NewGameScene()
	}

	return &Game{}
}

func (g *Game) Update() error {
	if internal.IsDebugMode {
		keys := inpututil.AppendPressedKeys([]ebiten.Key{})
		if lo.Contains(keys, ebiten.KeyQ) {
			os.Exit(0)
		}
	}

	scene.CurrentScene.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	scene.CurrentScene.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return internal.WindowWidth, internal.WindowHeight
}

func main() {
	ebiten.SetWindowSize(internal.WindowWidth, internal.WindowHeight)
	ebiten.SetWindowTitle("えびてん！")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
