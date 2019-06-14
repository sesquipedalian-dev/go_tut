package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed        = 1.0
	textureHeight      = 120
	textureWidth       = 120
	playerShotCooldown = time.Millisecond * 250
)

// note: analogous to prefab in Unity?
func newPlayer(renderer *sdl.Renderer) *element {
	player := newElement()

	player.position = vector{
		x: screenWidth / 2,
		y: screenHeight - (1.5 * textureHeight)}
	player.active = true

	sr := newSpriteRenderer(player, renderer, "sprites/player.bmp")
	player.addComponent(sr)

	km := newKeyboardMover(player, playerSpeed)
	player.addComponent(km)

	ks := newKeyboardShooter(player, playerShotCooldown)
	player.addComponent(ks)

	return player
}
