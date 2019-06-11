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

func (scene *Scene) Run(renderer *sdl.Renderer, events <-chan sdl.Event) <-chan error {
	errc := make(chan error)

	err := renderer.Copy(scene.background, nil, nil)
	if err != nil {
		errc <- fmt.Errorf("backgound render error : %v", err)
	}

	go func() {
		defer close(errc)
		tick := time.Tick(10 * time.Millisecond)

		for {
			select {
			case e := <-events:
				if done := scene.handleEvent(e); done {
					return
				}
			case <-tick:
				scene.update()
				err := scene.paint(renderer)
				if err != nil {
					errc <- err
				}
			}
		}
	}()

	return nil
}

func (scene *Scene) paint(renderer *sdl.Renderer) error {
	renderer.Clear()

	err := scene.breaker.Paint(renderer)
	if err != nil {
		return fmt.Errorf("breaker paint error : %v", err)
	}

	renderer.Present()

	return nil
}

func (scene *Scene) handleEvent(event sdl.Event) bool {
	switch e := event.(type) {
	case *sdl.QuitEvent:
		return true
	case *sdl.KeyboardEvent:
		if event.GetType() == sdl.KEYDOWN {
			switch e.Keysym.Sym {
			case sdl.K_LEFT:
				scene.breaker.NewPosition(0)
				break
			case sdl.K_RIGHT:
				scene.breaker.NewPosition(1)
				break
			}
		}

	}
	return false
}

func (scene *Scene) update() {

}

func (scene *Scene) destroy() {
	scene.breaker.Destroy()
}

func (scene *Scene) restart() {

}
