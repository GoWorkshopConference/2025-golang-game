package entity

var (
	score int = 0
)

type GameController struct{}

func NewGameController() *GameController {
	score = 0

	return &GameController{}
}

func (g *GameController) Update() {

}

func (g *GameController) GetScore() int {
	return score
}
