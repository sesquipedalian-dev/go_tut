package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type spriteRenderer struct {
	container     *element
	tex           *sdl.Texture
	width, height float64
}

func textureFromBMP(renderer *sdl.Renderer, filename string) (tex *sdl.Texture) {
	// loads BMP only
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		panic(fmt.Errorf("can't load sprite: %v", err))
	}
	defer img.Free()

	tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("creating texture from %v %v", filename, err))
	}

	return tex
}

func newSpriteRenderer(container *element, renderer *sdl.Renderer, filename string) *spriteRenderer {
	tex := textureFromBMP(renderer, filename)

	_, _, width, height, err := tex.Query()
	if err != nil {
		panic(fmt.Errorf("can't query tex %v", err))
	}

	return &spriteRenderer{
		container: container,
		tex:       tex,
		width:     float64(width),
		height:    float64(height),
	}
}

func (sr *spriteRenderer) onDraw(renderer *sdl.Renderer) error {
	// convert coord to top left of sprite
	x := sr.container.position.x - sr.width/2
	y := sr.container.position.y - sr.height/2

	// draw sprite
	w := int32(sr.width)
	h := int32(sr.height)
	renderer.CopyEx(sr.tex,
		&sdl.Rect{X: 0, Y: 0, W: w, H: h},
		&sdl.Rect{X: int32(x), Y: int32(y), W: w, H: h},
		sr.container.rotation, &sdl.Point{X: w / 2, Y: h / 2},
		sdl.FLIP_NONE)

	return nil
}

func (sr *spriteRenderer) onUpdate() error {
	// do nothing
	return nil
}

func (sr *spriteRenderer) onCollision(other *element) error {
	// do nothing
	return nil
}
