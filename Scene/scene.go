package Scene

import (
	"fmt"
	"github.com/mkilic91/goBreaker/Breaker"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

type Scene struct {
	frameRate  int
	background *sdl.Texture

	breaker *Breaker.Breaker
}

func NewScene(renderer *sdl.Renderer) (*Scene, error) {

	s, err := sdl.CreateRGBSurface(0, 400, 600, 32, 0, 0, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("background surface error : %v", err)
	}

	background, err := renderer.CreateTextureFromSurface(s)
	if err != nil {
		return nil, fmt.Errorf("backgound texture error : %v", err)
	}

	breaker, err := Breaker.NewBreaker(renderer)
	if err != nil {
		return nil, fmt.Errorf("new breaker error : %v", err)
	}

	return &Scene{frameRate: 60, background: background, breaker: breaker}, nil
}

func (scene *Scene) Run(renderer *sdl.Renderer) error {
	renderer.Clear()

	err := renderer.Copy(scene.background, nil, nil)
	if err != nil {
		return fmt.Errorf("backgound render error : %v", err)
	}

	go func() {
		for {
			scene.Paint(renderer)

			time.Sleep(time.Second)
		}
	}()

	renderer.Present()

	return nil
}

func (scene *Scene) Paint(renderer *sdl.Renderer) error {
	err := scene.breaker.Paint(renderer)
	if err != nil {
		return fmt.Errorf("breaker paint error : %v", err)
	}

	return nil
}
