package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480

	tileSize    = 16
	boardWidth  = ScreenWidth / tileSize
	boardHeight = ScreenHeight / tileSize
)

type Coordinate struct {
	X, Y int
}

type Game struct {
	snake     []Coordinate
	input     *Input
	direction Direction
}

func NewGame() ebiten.Game {
	startX, startY := boardWidth/2, boardHeight/2
	snake := []Coordinate{
		{startX, startY},
		{startX - 1, startY},
		{startX - 2, startY},
	}
	return &Game{
		snake:     snake,
		input:     NewInput(),
		direction: Up,
	}
}

func (g *Game) Update() error {
	dir, keyPressed := g.input.GetNewDirection()
	if keyPressed {
		g.direction = dir
	}
	var newHead Coordinate

	velocityX, velocityY := g.direction.ToVelocity()

	newHead.X = g.snake[0].X + velocityX
	newHead.Y = g.snake[0].Y + velocityY

	g.snake = append([]Coordinate{newHead}, g.snake...)
	g.snake = g.snake[:len(g.snake)-1]

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	for _, seg := range g.snake {
		vector.DrawFilledRect(
			screen,
			float32(seg.X*tileSize),
			float32(seg.Y*tileSize),
			tileSize,
			tileSize,
			color.RGBA{0, 255, 0, 255},
			false,
		)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
