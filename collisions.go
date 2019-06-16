package main

import "math"

type circle struct {
	center vector
	radius float64
}

func collision(c1, c2 circle) bool {
	dist := math.Sqrt(math.Pow(c1.center.x-c2.center.x, 2) +
		math.Pow(c1.center.y-c2.center.y, 2))
	return dist <= c1.radius+c2.radius
}

func checkCollisions() error {
	for i := 0; i < len(elements)-1; i++ {
		for j := i + 1; j < len(elements); j++ {
			if !elements[i].active && !elements[j].active {
				continue
			}
			for _, c1 := range elements[i].collisions {
				for _, c2 := range elements[j].collisions {
					if collision(c1, c2) {
						err := elements[i].collide(elements[j])
						if err != nil {
							return err
						}

						err = elements[j].collide(elements[i])
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}

	return nil
}
