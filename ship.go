package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

// ship - Структура корабля, текстура.
type ship struct {
	tex *sdl.Texture
}

// Newship - создаем новый корабль
func newShip(renderer *sdl.Renderer) (ship, error) {
	var s ship
	var err error

	img, err := sdl.LoadBMP("images/ship.bmp")
	if err != nil {
		return ship{}, fmt.Errorf("Load images ship: %v", err)
	}
	defer img.Free()

	s.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return ship{}, fmt.Errorf("creating ship texture: %v", err)
	}

	return s, nil
}

func (s *ship) draw(renderer *sdl.Renderer) {
	var src, dst sdl.Rect

	src = sdl.Rect{W: 60, H: 48}
	dst = sdl.Rect{X: 570, Y: 600, W: 60, H: 48}

	renderer.Copy(s.tex, &src, &dst)
}
