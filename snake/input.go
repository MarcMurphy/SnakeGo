package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Input struct {
}

func NewInput() *Input {
	return &Input{}
}

func (d Direction) ToVelocity() (x, y int) {
	switch d {
	case Up:
		return 0, -1
	case Down:
		return 0, 1
	case Left:
		return -1, 0
	case Right:
		return 1, 0
	default:
		return 0, 0
	}
}

func (i *Input) GetNewDirection() (Direction, bool) {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		return Up, true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		return Down, true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		return Left, true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		return Right, true
	}
	return -1, false
}
