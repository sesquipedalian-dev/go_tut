package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type keyboardMover struct {
	container *element
	speed     float64
	sr        *spriteRenderer
}

func newKeyboardMover(container *element, speed float64) *keyboardMover {
	return &keyboardMover{
		container: container,
		speed:     speed,
		sr:        container.getComponent(&spriteRenderer{}).(*spriteRenderer),
	}
}

func (mover *keyboardMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover *keyboardMover) onCollision(other *element) error {
	return nil
}

func (mover *keyboardMover) onUpdate() error {
	keys := sdl.GetKeyboardState()
	var dx = 0.0
	var dy = 0.0
	if keys[sdl.SCANCODE_LEFT] == 1 {
		// Move player Left
		dx = -mover.speed
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		// move right
		dx = mover.speed
	} else if keys[sdl.SCANCODE_DOWN] == 1 {
		// move down
		dy = mover.speed
	} else if keys[sdl.SCANCODE_UP] == 1 {
		// move up
		dy = -mover.speed
	}

	// move player by dx and dy, constrained by screen size
	mover.container.position.x = math.Max(math.Min(
		mover.container.position.x+dx, screenWidth-mover.sr.width/2), 0)
	mover.container.position.y = math.Max(math.Min(
		mover.container.position.y+dy, screenHeight-mover.sr.height/2), 0)

	return nil
}

type keyboardShooter struct {
	container *element
	cooldown  time.Duration
	lastShot  time.Time
	sr        *spriteRenderer
}

func newKeyboardShooter(container *element, cooldown time.Duration) *keyboardShooter {
	return &keyboardShooter{
		container: container,
		cooldown:  cooldown,
		sr:        container.getComponent(&spriteRenderer{}).(*spriteRenderer),
	}
}

func (shooter *keyboardShooter) onCollision(other *element) error {
	return nil
}

func (shooter *keyboardShooter) onUpdate() error {
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(shooter.lastShot) >= shooter.cooldown {
			shooter.shoot(0)
			shooter.shoot(2)

			shooter.lastShot = time.Now()
		}
	}

	return nil
}

func (shooter *keyboardShooter) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (shooter *keyboardShooter) shoot(rel int32) {
	if bul, ok := bulletFromPool(); ok {
		bul.position.x = shooter.container.position.x - shooter.sr.width/2 + shooter.sr.width/2*float64(rel)
		bul.position.y = shooter.container.position.y - shooter.sr.height/2
		bul.getComponent(&bulletMover{}).(*bulletMover).angle = 270 * (math.Pi / 180)
	}
}
