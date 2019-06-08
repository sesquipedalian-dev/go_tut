package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed        = 1.0
	textureHeight      = 120
	textureWidth       = 120
	playerShotCooldown = time.Millisecond * 250
)

type player struct {
	tex      *sdl.Texture
	x, y     float64
	lastShot time.Time
}

func newPlayer(renderer *sdl.Renderer) (p player) {
	p.tex = textureFromBMP(renderer, "sprites/player.bmp")

	// initial position
	p.x = screenWidth / 2                      // centered
	p.y = screenHeight - (1.5 * textureHeight) // 1.5 player lengths from the bottom

	return p
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

	// move player by dx and dy, constrained by screen size
	p.x = math.Max(math.Min(p.x+dx, screenWidth-textureWidth/2), 0+textureWidth/2)
	p.y = math.Max(math.Min(p.y+dy, screenHeight-textureHeight/2), 0+textureHeight/2)

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(p.lastShot) >= playerShotCooldown {
			p.shoot(0)
			p.shoot(2)

			p.lastShot = time.Now()
		}
	}
}

func (p *player) shoot(rel int32) {
	if bul, ok := bulletFromPool(); ok {
		bul.x = p.x - textureWidth/2 + textureWidth/2*float64(rel)
		bul.y = p.y - textureHeight/2
		bul.angle = 270 * (math.Pi / 180)
	}
}
