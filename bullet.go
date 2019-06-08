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

func newBullet(renderer *sdl.Renderer) (bul bullet) {
	bul.tex = textureFromBMP(renderer, "sprites/bullet.bmp")
	return bul
}

func (bul *bullet) draw(renderer *sdl.Renderer) {
	if !bul.active {
		return
	}

	x := bul.x - bulletSize/2
	y := bul.y - bulletSize/2

	renderer.Copy(bul.tex,
		&sdl.Rect{X: 0, Y: 0, W: bulletSize, H: bulletSize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: bulletSize, H: bulletSize})
}

func (bul *bullet) update() {
	if !bul.active {
		return
	}

	bul.x += bulletSpeed * math.Cos(bul.angle)
	bul.y += bulletSpeed * math.Sin(bul.angle)

	if bul.x > screenWidth || bul.x < 0 || bul.y > screenHeight || bul.y < 0 {
		bul.active = false
	}
}

var bulletPool []*bullet

const bulletPoolSize = 30

func initBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < bulletPoolSize; i++ {
		bul := newBullet(renderer)
		bulletPool = append(bulletPool, &bul)
	}
}

func bulletFromPool() (*bullet, bool) {
	for _, bul := range bulletPool {
		if !bul.active {
			bul.active = true
			return bul, true
		}
	}

	return nil, false
}
