package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type bullet struct {
	tex    *sdl.Texture
	x, y   float64
	active bool

	angle float64
}

const (
	bulletSize  = 35
	bulletSpeed = 0.15
)

func newBullet(renderer *sdl.Renderer) *element {
	bul := newElement()
	bul.active = false

	sr := newSpriteRenderer(bul, renderer, "sprites/bullet.bmp")
	bul.addComponent(sr)

	mover := newBulletMover(bul)
	bul.addComponent(mover)

	return bul
}

type bulletMover struct {
	speed     float64
	angle     float64
	container *element
}

func newBulletMover(container *element) *bulletMover {
	return &bulletMover{
		container: container,
		speed:     bulletSpeed,
	}
}

func (bul *bulletMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (bul *bulletMover) onUpdate() error {
	if !bul.container.active {
		return nil
	}

	bul.container.position.x += bul.speed * math.Cos(bul.angle)
	bul.container.position.y += bul.speed * math.Sin(bul.angle)

	if bul.container.position.x > screenWidth || bul.container.position.x < 0 ||
		bul.container.position.y > screenHeight || bul.container.position.y < 0 {
		bul.container.active = false
	}

	return nil
}

var bulletPool []*element

const bulletPoolSize = 30

func initBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < bulletPoolSize; i++ {
		bul := newBullet(renderer)
		bulletPool = append(bulletPool, bul)
	}
}

func bulletFromPool() (*element, bool) {
	for _, bul := range bulletPool {
		if !bul.active {
			bul.active = true
			return bul, true
		}
	}

	return nil, false
}
