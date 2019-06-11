package Breaker

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Breaker struct {
	texture *sdl.Texture
	postion int32
}

func NewBreaker(renderer *sdl.Renderer) (*Breaker, error) {
	texture, err := img.LoadTexture(renderer, "Assets/breaker.png")

	if err != nil {
		return nil, fmt.Errorf("texture load error : %v", err)
	}

	return &Breaker{texture: texture, postion: 150}, nil
}

func (breaker *Breaker) Paint(renderer *sdl.Renderer) error {

	rect := &sdl.Rect{X: breaker.postion, Y: 550, W: 100, H: 20}
	err := renderer.CopyEx(breaker.texture, nil, rect, 0, nil, sdl.FLIP_NONE)

	if err != nil {

		return fmt.Errorf("renderer error : %v", err)
	}

	return nil
}

func (breaker *Breaker) NewPosition(position int8) {
	if position == 0 {
		breaker.postion -= 10
	} else if position == 1 {
		breaker.postion += 10
	}
}

func (breaker *Breaker) Destroy() {
	breaker.texture.Destroy()
}
