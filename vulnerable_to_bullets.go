package main

import "github.com/veandco/go-sdl2/sdl"

type vulnerableToBullets struct {
	container *element
}

func newVulnerableToBullets(container *element) *vulnerableToBullets {
	return &vulnerableToBullets{container: container}
}

func (vul *vulnerableToBullets) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (vul *vulnerableToBullets) onUpdate() error {
	vul.container.collisions[0].center = vul.container.position
	return nil
}

func (vul *vulnerableToBullets) onCollision(other *element) error {
	if other.tag == "bullet" {
		vul.container.active = false
		other.active = false
	}
	return nil
}
