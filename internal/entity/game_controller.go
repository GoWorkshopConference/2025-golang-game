package entity

const (
	InitialLife = 3
)

var (
	score int = 0
	life  int = InitialLife
)

type GameController struct{}

func NewGameController() *GameController {
	score = 0
	life = InitialLife

	return &GameController{}
}

func (g *GameController) Update() {
}

func (g *GameController) GetScore() int {
	return score
}

func (g *GameController) GetLife() int {
	return life
}
