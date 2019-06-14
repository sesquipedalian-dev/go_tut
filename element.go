package main

import (
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type vector struct {
	x, y float64
}

type component interface {
	onUpdate() error // TODO game time since last frame?
	onDraw(renderer *sdl.Renderer) error
}

type element struct {
	position   vector
	rotation   float64
	active     bool
	components []component
}

func newElement() *element {
	elem := &element{}
	elements = append(elements, elem)
	return elem
}

func (elem *element) addComponent(new component) {
	typ := reflect.TypeOf(new)
	// only one component of a type
	for _, existing := range elem.components {
		if typ == reflect.TypeOf(existing) {
			panic(fmt.Sprintf("added component of already existing type %v",
				reflect.TypeOf(new)))
		}
	}

	elem.components = append(elem.components, new)
}

func (elem *element) getComponent(withType component) component {
	typ := reflect.TypeOf(withType)
	for _, existing := range elem.components {
		if typ == reflect.TypeOf(existing) {
			return existing
		}
	}

	// TODO maybe by default we return the component they passed in?
	panic(fmt.Sprintf("getComponent of type %v not found",
		reflect.TypeOf(withType)))
}

func (elem *element) draw(renderer *sdl.Renderer) error {
	if !elem.active {
		return nil
	}

	for _, comp := range elem.components {
		err := comp.onDraw(renderer)
		if err != nil {
			return err
		}
	}

	return nil
}

func (elem *element) update() error {
	if !elem.active {
		return nil
	}

	for _, comp := range elem.components {
		err := comp.onUpdate()
		if err != nil {
			return err
		}
	}

	return nil
}

var elements []*element
