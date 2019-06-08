package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

func textureFromBMP(renderer *sdl.Renderer, filename string) (tex *sdl.Texture) {
	// loads BMP only
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		panic(fmt.Errorf("can't load sprite: %v", err))
	}
	defer img.Free()

	tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("creating texture from %v %v", filename, err))
	}

	return tex
}

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

	p := newPlayer(renderer)

	var enemies []basicEnemy

	initBulletPool(renderer)

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			x := float64(i)/5*screenWidth + (basicEnemySize / 2.0)
			y := float64(j) * (basicEnemySize + 20)

			enemy := newBasicEnemy(renderer, x, y)

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

		for _, bullet := range bulletPool {
			bullet.update()
			bullet.draw(renderer)
		}

		renderer.Present()
	}
}
