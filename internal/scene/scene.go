package scene

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
}

type SceneType string

var (
	SceneTypeDebug SceneType = "DEBUG_SCENE"
	SceneTypeGame  SceneType = "GAME_SCENE"
)

var CurrentScene Scene
