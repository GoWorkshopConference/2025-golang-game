package main

import (
	"fmt"
	"log"
	"os"

	"github.com/GoWorkshopConference/golang-game/internal"
	"github.com/GoWorkshopConference/golang-game/internal/entity"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/samber/lo"
)

type entities interface {
	Update()
	Draw(screen *ebiten.Image)
}

type Game struct {
	player           *entity.Player
	controllerButton *entity.ControllerButton
	shootButton      *entity.ShootButton
	text             string
}

func NewGame() *Game {
	return &Game{
		player:           entity.NewPlayer(),
		controllerButton: entity.NewControllerButton(),
		shootButton:      entity.NewShootButton(),
	}
}

func (g *Game) Update() error {
	keys := inpututil.AppendPressedKeys([]ebiten.Key{})
	if lo.Contains(keys, ebiten.KeyQ) {
		os.Exit(0)
	}

	controllerButtonTouchEvent := g.controllerButton.Update()
	shootButtonTouchEvent := g.shootButton.Update()
	g.player.Update(controllerButtonTouchEvent, shootButtonTouchEvent)

	touchIDs := ebiten.AppendTouchIDs([]ebiten.TouchID{})
	// log.Println(touchIDs)
	if len(touchIDs) > 0 {
		x, y := ebiten.TouchPosition(touchIDs[0])
		g.text = fmt.Sprintf("Touch ID: %d, Position: %v, %v", touchIDs[0], x, y)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
	g.controllerButton.Draw(screen)
	g.shootButton.Draw(screen)

	ebitenutil.DebugPrint(screen, "Hello, World!")
	ebitenutil.DebugPrint(screen, g.text)

	// lo.ForEach(g.buttons, func(button *entity.Button, _ int) {
	// 	button.Draw(screen)
	// })
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return internal.WindowWidth, internal.WindowHeight
}

func main() {
	ebiten.SetWindowSize(internal.WindowWidth, internal.WindowHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
