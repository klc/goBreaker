package Block

import (
	"fmt"
	"github.com/mkilic91/goBreaker/Ball"
	"github.com/mkilic91/goBreaker/Print"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
	"strconv"
)

const w = 50
const h = 25

type Blocks struct {
	blocks   []*Block
	textures []*sdl.Texture
	size     size
}

type Block struct {
	textureKey int32
	position   position
	score      int
}

type position struct {
	x int32
	y int32
}

type size struct {
	w int32
	h int32
}

var images = []string{
	"assets/b1.png",
	"assets/b2.png",
	"assets/b3.png",
	"assets/b4.png",
	"assets/b5.png",
}

var ball *Ball.Ball
var score *Print.Print

//var xPositions = []int32{0, 50, 100, 150, 200, 250, 300, 350}
var xPositions = []int32{0, 50, 100, 150, 200, 250, 300, 350}
var yPositions = []int32{150, 175, 200, 225}

func NewBlocks(renderer *sdl.Renderer) (*Blocks, error) {
	blocks := &Blocks{
		size: size{w: w, h: h},
	}

	for _, image := range images {
		texture, err := img.LoadTexture(renderer, image)

		if err != nil {
			return blocks, fmt.Errorf("blocks texture load error :%v", err)
		}

		blocks.textures = append(blocks.textures, texture)
	}

	for _, yposition := range yPositions {
		for _, xposition := range xPositions {
			blocks.blocks = append(blocks.blocks, newBlock(xposition, yposition))
		}
	}

	return blocks, nil
}

func newBlock(x int32, y int32) *Block {
	textureKey := int32(rand.Intn(len(images)))

	return &Block{textureKey: textureKey, position: position{x: x, y: y}, score: 1}
}

func (blocks *Blocks) Paint(renderer *sdl.Renderer) error {
	for _, block := range blocks.blocks {
		err := block.Paint(renderer, blocks.textures, blocks.size.w, blocks.size.h)
		if err != nil {

			return fmt.Errorf("block paint error : %v", err)
		}
	}

	return nil
}

func (block *Block) Paint(renderer *sdl.Renderer, textures []*sdl.Texture, w int32, h int32) error {
	rect := &sdl.Rect{X: block.position.x, Y: block.position.y, W: w, H: h}
	err := renderer.CopyEx(textures[block.textureKey], nil, rect, 0, nil, sdl.FLIP_NONE)

	if err != nil {

		return fmt.Errorf("block renderer error : %v", err)
	}

	return nil
}

func (blocks *Blocks) Update(b *Ball.Ball, s *Print.Print) {
	ball = b
	score = s

	var remainingBlocks []*Block

	for _, block := range blocks.blocks {
		t := block.Update()
		if !t {
			remainingBlocks = append(remainingBlocks, block)
		}
	}

	blocks.blocks = remainingBlocks
}

func (block *Block) Update() bool {
	return block.Touch()
}

func (block *Block) Touch() bool {
	ballPositionX, ballPositionY := ball.GetPosition()
	//fmt.Println(ballPositionY)
	blockBottom := block.position.y + h
	blockTop := block.position.y
	blockRight := block.position.x + w
	blockLeft := block.position.x

	ballTop := ballPositionY
	ballBottom := ballPositionY + 20
	ballLeft := ballPositionX
	ballRight := ballPositionX + 20

	var moveX int32 = 0
	var moveY int32 = 0
	var touch bool = false

	switch true {

	case ballTop <= blockBottom && ballBottom > blockBottom:
		if ballLeft <= blockRight && ballRight >= blockLeft {
			moveX = 0
			moveY = -5
			touch = true
		}
		break
	case ballBottom >= blockTop && ballTop < blockTop:
		if ballLeft <= blockRight && ballRight >= blockLeft {
			moveX = 0
			moveY = 5
			touch = true
		}
		break
	case ballLeft <= blockRight && ballRight > blockRight:
		if ballBottom > blockTop && ballTop < blockBottom {
			moveX = -5
			moveY = 0
			touch = true
		}
		break
	case ballRight >= blockLeft && ballLeft < blockLeft:
		if ballBottom > blockTop && ballTop < blockBottom {
			moveX = 5
			moveY = 0
			touch = true
		}
		break
	}

	ball.SetMove(moveX, moveY)

	if moveX != 0 || moveY != 0 {
		s, _ := strconv.Atoi(score.GetContent())
		s += block.score
		score.Update(strconv.Itoa(s))
	}

	return touch
}

func (blocks *Blocks) Restart() {
	for _, yposition := range yPositions {
		for _, xposition := range xPositions {
			blocks.blocks = append(blocks.blocks, newBlock(xposition, yposition))
		}
	}
}
