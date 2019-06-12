package Breaker

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Breaker struct {
	texture  *sdl.Texture
	position int32

	width  int32
	height int32
}

func NewBreaker(renderer *sdl.Renderer) (*Breaker, error) {
	texture, err := img.LoadTexture(renderer, "Assets/breaker.png")

	if err != nil {
		return nil, fmt.Errorf("texture load error : %v", err)
	}

	return &Breaker{texture: texture, position: 150, width: 100, height: 20}, nil
}

func (breaker *Breaker) Paint(renderer *sdl.Renderer) error {

	rect := &sdl.Rect{X: breaker.position, Y: 550, W: breaker.width, H: breaker.height}
	err := renderer.CopyEx(breaker.texture, nil, rect, 0, nil, sdl.FLIP_NONE)

	if err != nil {

		return fmt.Errorf("renderer error : %v", err)
	}

	return nil
}

func (breaker *Breaker) NewPosition(position int8) {
	if breaker.position > 0 || breaker.position < 400 {
		if position == 0 {
			breaker.position -= 10
		} else if position == 1 {
			breaker.position += 10
		}
	}
}

func (breaker *Breaker) Destroy() {
	breaker.texture.Destroy()
}

func (breaker *Breaker) GetPosition() int32 {
	return breaker.position
}
