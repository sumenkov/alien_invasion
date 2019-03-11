package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	//alienSpeed  = 0.05
	alienWidth  = 60
	alienHeight = 58
)

type alien struct {
	tex  *sdl.Texture
	x, y float64
}

func newAlien(renderer *sdl.Renderer, x, y float64) (a alien) {
	a.tex = textureFromBMP(renderer, "images/alien.bmp")
	// Стартовое положение пришельца
	a.x = x
	a.y = y

	return a
}

func (a *alien) draw(renderer *sdl.Renderer) {
	x := a.x - alienWidth
	y := a.y - alienHeight

	renderer.CopyEx(a.tex,
		&sdl.Rect{W: alienWidth, H: alienHeight},
		&sdl.Rect{X: int32(x), Y: int32(y), W: alienWidth, H: alienHeight},
		180,
		&sdl.Point{X: alienWidth / 2.0, Y: alienHeight / 2.0},
		sdl.FLIP_NONE)
}
