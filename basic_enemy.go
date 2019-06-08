package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type basicEnemy struct {
	tex  *sdl.Texture
	x, y float64
}

const basicEnemySize = 120

func newBasicEnemy(renderer *sdl.Renderer, initialX, initialY float64) (be basicEnemy) {
	be.tex = textureFromBMP(renderer, "sprites/baddie1.bmp")

	// initial position
	be.x = initialX
	be.y = initialY

	return be
}

func (be *basicEnemy) draw(renderer *sdl.Renderer) {
	// convert basic enemy coord to top left of sprite
	x := be.x - (basicEnemySize / 2)
	y := be.y - (basicEnemySize / 2)
	renderer.CopyEx(be.tex,
		&sdl.Rect{X: 0, Y: 0, W: basicEnemySize, H: basicEnemySize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: basicEnemySize, H: basicEnemySize},
		180, &sdl.Point{X: basicEnemySize / 2, Y: basicEnemySize / 2},
		sdl.FLIP_NONE)
}
