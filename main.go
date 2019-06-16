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

	newPlayer(renderer)
	initBulletPool(renderer)

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			x := float64(i)/5*screenWidth + (120 / 2.0)
			y := float64(j) * (120 + 20)

			newBasicEnemy(renderer, vector{x: x, y: y})
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

		for _, elem := range elements {
			err = elem.update()
			if err != nil {
				fmt.Printf("couldn't update elem %v", err)
				return
			}

			err = elem.draw(renderer)
			if err != nil {
				fmt.Printf("couldn't draw elem %v", err)
				return
			}
		}

		err = checkCollisions()
		if err != nil {
			fmt.Printf("couldn't collide %v", err)
			return
		}

		renderer.Present()
	}
}
