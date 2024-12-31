package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"time"
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
	snake      []Coordinate
	input      *Input
	direction  Direction
	nextMove   Direction
	updateTick time.Duration
	lastUpdate time.Time
}

func NewGame() ebiten.Game {
	startX, startY := boardWidth/2, boardHeight/2
	snake := []Coordinate{
		{startX, startY},
		{startX - 1, startY},
		{startX - 2, startY},
	}
	return &Game{
		snake:      snake,
		input:      NewInput(),
		direction:  Up,
		updateTick: 200 * time.Millisecond,
		lastUpdate: time.Now(),
	}
}
func (g *Game) changeDirection() {
	if dir, keyPressed := g.input.GetNewDirection(); keyPressed {
		dx1, dy1 := g.direction.ToVelocity()
		dx2, dy2 := dir.ToVelocity()
		// Only change direction if it's NOT the opposite
		if !(dx1 == -dx2 && dy1 == -dy2) {
			g.nextMove = dir
		}
	}
}

func wrap(value, max int) int {
	//shorthand version of "if x < 0 then x = boardWidth OR if x > boardWidth then x = 0"
	return (value%max + max) % max
}

func (g *Game) moveSnake() {
	g.direction = g.nextMove
	head := g.snake[0]

	velocityX, velocityY := g.direction.ToVelocity()
	newHead := Coordinate{X: head.X + velocityX, Y: head.Y + velocityY}

	newHead.X = wrap(newHead.X, boardWidth)
	newHead.Y = wrap(newHead.Y, boardHeight)

	// Insert new head at the front
	g.snake = append([]Coordinate{newHead}, g.snake...)
	// Remove the tail
	g.snake = g.snake[:len(g.snake)-1]
}

func (g *Game) Update() error {
	// update direction every frame but only move snake every X frames, as the keyboard input is smoother this way
	g.changeDirection()
	if time.Since(g.lastUpdate) < g.updateTick {
		return nil
	}
	g.lastUpdate = time.Now()
	g.moveSnake()

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
