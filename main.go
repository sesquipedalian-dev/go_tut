package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

func main() {
	fmt.Println("Hello, world!")
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow(
		"Gaming in Go Ep 2",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer error", err)
		return
	}
	defer renderer.Destroy()

	p, err := newPlayer(renderer)
	if err != nil {
		fmt.Println("initializing player error", err)
		return
	}

	var enemies []basicEnemy

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			x := float64(i)/5*screenWidth + (basicEnemySize / 2.0)
			y := float64(j) * (basicEnemySize + 20)

			enemy, err := newBasicEnemy(renderer, x, y)
			if err != nil {
				fmt.Println("initializing enemy error", err)
				return
			}

			enemies = append(enemies, enemy)
		}
	}

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(0, 0, 0, 255) // RGBA White
		renderer.Clear()

		p.draw(renderer)
		p.update()

		for _, enemy := range enemies {
			enemy.draw(renderer)
		}

		renderer.Present()
	}
}
