package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 1200
	screenHeight = 650
)

var bgColor = []uint8{230, 230, 230, 230}
var window *sdl.Window
var renderer *sdl.Renderer
var err error

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		println("initializing SDL:", err)
		return
	}

	window, err = sdl.CreateWindow(
		"Alien Invasion",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		println("initializing Window:", err)
		return
	}
	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		println("initializing Renderer:", err)
		return
	}
	defer renderer.Destroy()

	ship, err := newShip(renderer)
	if err != nil {
		println("load newShip:", err)
		return
	}

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit.")
				return
			}
		}
		renderer.SetDrawColor(230, 230, 230, 230)
		renderer.Clear()

		ship.draw(renderer)

		renderer.Present()
	}
}
