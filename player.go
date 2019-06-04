package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
	tex *sdl.Texture
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

	return p, nil
}

func (p *player) draw(renderer *sdl.Renderer) {
	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: 120, H: 120},
		&sdl.Rect{X: 50, Y: 50, W: 240, H: 60})
}
