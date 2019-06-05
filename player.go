package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const playerSpeed = 1.0
const textureHeight = 120
const textureWidth = 120

type player struct {
	tex  *sdl.Texture
	x, y float64
}

func newPlayer(renderer *sdl.Renderer) (p player, err error) {
	// loads BMP only
	img, err := sdl.LoadBMP("sprites/player.bmp")
	if err != nil {
		return player{}, fmt.Errorf("can't load sprite: %v", err)
	}
	defer img.Free()

	p.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return player{}, fmt.Errorf("can't create player texture: %v", err)
	}

	// initial position
	p.x = screenWidth / 2                      // centered
	p.y = screenHeight - (1.5 * textureHeight) // 1.5 player lengths from the bottom

	return p, nil
}

func (p *player) draw(renderer *sdl.Renderer) {
	// convert player coord to top left of sprite
	x := p.x - (textureWidth / 2)
	y := p.y - (textureHeight / 2)
	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: textureWidth, H: textureHeight},
		&sdl.Rect{X: int32(x), Y: int32(y), W: textureWidth, H: textureHeight})
}

func (p *player) update() {
	keys := sdl.GetKeyboardState()
	var dx = 0.0
	var dy = 0.0
	if keys[sdl.SCANCODE_LEFT] == 1 {
		// Move player Left
		dx = -playerSpeed
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		// move right
		dx = playerSpeed
	} else if keys[sdl.SCANCODE_DOWN] == 1 {
		// move down
		dy = playerSpeed
	} else if keys[sdl.SCANCODE_UP] == 1 {
		// move up
		dy = -playerSpeed
	}

	p.x = p.x + dx
	p.y = p.y + dy
}
