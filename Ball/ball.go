package Ball

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Ball struct {
	texture  *sdl.Texture
	position position
	move     move
	width    int32
	height   int32
}

type position struct {
	x int32
	y int32
}

type move struct {
	x int32
	y int32
}

func NewBall(renderer *sdl.Renderer) (*Ball, error) {
	texture, err := img.LoadTexture(renderer, "assets/ball.png")

	if err != nil {
		return nil, fmt.Errorf("texture load error : %v", err)
	}

	position := position{x: 190, y: 540}
	move := move{x: 0, y: 5}

	return &Ball{texture: texture, position: position, move: move, width: 20, height: 20}, nil
}

func (ball *Ball) Paint(renderer *sdl.Renderer) error {
	rect := &sdl.Rect{X: ball.position.x, Y: ball.position.y, W: ball.width, H: ball.height}
	err := renderer.CopyEx(ball.texture, nil, rect, 0, nil, sdl.FLIP_NONE)

	if err != nil {
		return fmt.Errorf("renderer error : %v", err)
	}

	return nil
}

func (ball *Ball) Destroy() {
	ball.texture.Destroy()
}

func (ball *Ball) Update(breakerPosition int32) {

	if ball.position.x < 0 {
		ball.move.x = -5
	}

	if ball.position.x > 400 {
		ball.move.x = 5
	}

	if ball.position.y < 0 {
		ball.move.y = -5
	}

	if ball.position.y == 550 {
		ballPosition := ball.position.x + 10
		breakerPositionA := breakerPosition - 100
		breakerPositionB := breakerPosition + 100

		if ballPosition > breakerPositionA && ballPosition < breakerPositionB {
			ball.move.y = 5
		}
	}

	ball.position.x -= ball.move.x
	ball.position.y -= ball.move.y
}
