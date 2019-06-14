package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func newBasicEnemy(renderer *sdl.Renderer, position vector) *element {
	be := newElement()

	be.active = true
	be.position = position
	be.rotation = 180.

	sr := newSpriteRenderer(be, renderer, "sprites/baddie1.bmp")
	be.addComponent(sr)

	return be
}
