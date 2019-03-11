package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 1366
	screenHeight = 688
)

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

	ship := newShip(renderer)

	var alienPool []alien
	// Вычисляем количество пришельцев
	numAlienX := (screenWidth - alienWidth*2.0) / (alienWidth * 2.0)
	numAlienY := (screenHeight - alienWidth*3) / (alienWidth * 2)

	for i := 0; i < int(numAlienX); i++ {
		for j := 0; j < numAlienY; j++ {
			x := float64(i)/numAlienX*screenWidth + alienWidth*1.5
			y := float64(j)*alienHeight*3.0/2 + alienHeight*1.5

			alien := newAlien(renderer, x, y)
			alienPool = append(alienPool, alien)
		}
	}

	initBulletPool(renderer)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(230, 230, 230, 230)
		renderer.Clear()

		// Помещаем объекты в окно
		ship.draw(renderer)

		for _, alien := range alienPool {
			alien.draw(renderer)
		}

		for _, b := range bulletPool {
			b.draw(renderer)
			b.update()
		}

		// Обновляем положение объектов
		ship.update()

		renderer.Present()
	}
}
