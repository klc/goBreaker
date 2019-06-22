package Scene

import (
	"fmt"
	"github.com/mkilic91/goBreaker/Ball"
	"github.com/mkilic91/goBreaker/Block"
	"github.com/mkilic91/goBreaker/Breaker"
	"github.com/mkilic91/goBreaker/Print"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

type Scene struct {
	frameRate  int
	background *sdl.Texture

	breaker *Breaker.Breaker
	ball    *Ball.Ball
	blocks  *Block.Blocks
	score   *Print.Print
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

	ball, err := Ball.NewBall(renderer)
	if err != nil {
		return nil, fmt.Errorf("new ball error : %v", err)
	}

	blocks, err := Block.NewBlocks(renderer)
	if err != nil {
		return nil, fmt.Errorf("new blocks error : %v", err)
	}

	score, err := Print.NewPrint("0", 20, 30, 60, 40)
	if err != nil {
		return nil, fmt.Errorf("new print error : %v", err)
	}

	return &Scene{frameRate: 60, background: background, breaker: breaker, ball: ball, blocks: blocks, score: score}, nil
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
				_, y := scene.ball.GetPosition()

				if y > 600 {
					scene.restart()
				}
			}
		}
	}()

	return nil
}

func (scene *Scene) paint(renderer *sdl.Renderer) error {
	var err error
	renderer.Clear()

	err = scene.breaker.Paint(renderer)
	if err != nil {
		return fmt.Errorf("breaker paint error : %v", err)
	}

	err = scene.ball.Paint(renderer)
	if err != nil {
		return fmt.Errorf("ball paint error : %v", err)
	}

	err = scene.blocks.Paint(renderer)
	if err != nil {
		return fmt.Errorf("blocks paint error : %v", err)
	}

	err = scene.score.Paint(renderer)
	if err != nil {
		return fmt.Errorf("print score error : %v", err)
	}

	renderer.Present()

	return nil
}

func (scene *Scene) handleEvent(event sdl.Event) bool {
	switch e := event.(type) {
	case *sdl.QuitEvent:
		return true
	case *sdl.MouseMotionEvent:
		scene.breaker.NewPosition(e.X)
		/*default:
		fmt.Printf("event %T :", e)*/
	}
	return false
}

func (scene *Scene) update() {
	scene.ball.Update(scene.breaker.GetPosition())
	scene.blocks.Update(scene.ball, scene.score)
}

func (scene *Scene) destroy() {
	scene.breaker.Destroy()
	scene.score.Destroy()
}

func (scene *Scene) restart() {
	scene.ball.Restart()
	scene.score.Update("0")
	scene.blocks.Restart()
}
