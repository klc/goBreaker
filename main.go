package main

import (
	"fmt"
	"github.com/mkilic91/goBreaker/Scene"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"runtime"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}
}

func run() error {
	var err error
	var window *sdl.Window
	var renderer *sdl.Renderer
	var events chan sdl.Event

	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return fmt.Errorf("init error : %v", err)
	}
	defer sdl.Quit()

	window, renderer, err = sdl.CreateWindowAndRenderer(400, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return fmt.Errorf("window renderer error : %v", err)
	}
	defer window.Destroy()
	defer renderer.Destroy()

	scene, err := Scene.NewScene(renderer)
	if err != nil {
		return fmt.Errorf("new scene error : %v", err)
	}

	events = make(chan sdl.Event)
	errc := scene.Run(renderer, events)

	runtime.LockOSThread()

	for {
		select {
		case events <- sdl.WaitEvent():
		case err := <-errc:
			return err
		}
	}
}
