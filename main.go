package main

import (
	"fmt"
	"github.com/mkilic91/goBreaker/Scene"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"runtime"
	"time"
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

	err = scene.Run(renderer)
	if err != nil {
		return fmt.Errorf("scene run error : %v", err)
	}

	runtime.LockOSThread()

	time.Sleep(time.Second * 3)

	return nil
}
