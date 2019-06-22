package Print

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Print struct {
	content  string
	texture  *sdl.Texture
	font     *ttf.Font
	color    sdl.Color
	position position
	size     size
}

type position struct {
	x int32
	y int32
}

type size struct {
	w int32
	h int32
}

func NewPrint(content string, x int32, y int32, w int32, h int32) (*Print, error) {
	color := sdl.Color{R: 255, G: 255, B: 255, A: 255}
	var position = position{x: x, y: y}
	var size = size{w: w, h: h}

	return &Print{content: content, color: color, position: position, size: size}, nil

}

func (print *Print) Paint(renderer *sdl.Renderer) error {
	font, err := ttf.OpenFont("assets/DigitalDream.ttf", 50)
	if err != nil {
		return fmt.Errorf("font load error : %v", err)
	}
	defer font.Close()

	surface, err := font.RenderUTF8Solid(print.content, print.color)

	if err != nil {
		return fmt.Errorf("font render error : %v", err)
	}
	defer surface.Free()

	print.texture, err = renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return fmt.Errorf("create texture error : %v", err)
	}
	defer print.texture.Destroy()

	rect := &sdl.Rect{X: print.position.x, Y: print.position.y, W: print.size.w, H: print.size.h}
	err = renderer.CopyEx(print.texture, nil, rect, 0, nil, sdl.FLIP_NONE)
	if err != nil {
		return fmt.Errorf("create texture error : %v", err)
	}

	return nil
}

func (print *Print) Destroy() {
	print.texture.Destroy()
}

func (print *Print) Update(content string) {
	print.content = content
}

func (print *Print) GetContent() string {
	return print.content
}
