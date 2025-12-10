package main

import (
	"log"

	"github.com/GoWorkshopConference/golang-game/internal"
	"github.com/GoWorkshopConference/golang-game/internal/entity"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type entities interface {
	Update()
	Draw(screen *ebiten.Image)
}

type Game struct {
	player *entity.Player
}

func NewGame() *Game {
	return &Game{
		player: entity.NewPlayer(),
	}
}

func (g *Game) Update() error {
	g.player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(internal.WindowWidth, internal.WindowHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
