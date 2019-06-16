package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func newBasicEnemy(renderer *sdl.Renderer, position vector) *element {
	be := newElement()

	be.active = true
	be.position = position
	be.rotation = 180.
	be.tag = "enemy"

	sr := newSpriteRenderer(be, renderer, "sprites/baddie1.bmp")
	be.addComponent(sr)

	col := circle{
		center: be.position,
		radius: 36.,
	}
	be.collisions = append(be.collisions, col)

	vtb := newVulnerableToBullets(be)
	be.addComponent(vtb)

	return be
}
