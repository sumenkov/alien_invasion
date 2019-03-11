package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletWidth  = 3
	bulletHeight = 15
	bulletSpeed  = 0.4
)

type bullet struct {
	tex    *sdl.Texture
	x, y   float64
	angle  float64
	active bool
}

func newBullet(renderer *sdl.Renderer) (b bullet) {
	b.tex = textureFromBMP(renderer, "images/bullet.bmp")
	return b
}

func (b *bullet) draw(renderer *sdl.Renderer) {
	if !b.active {
		return
	}
	x := int32(b.x)
	y := int32(b.y)
	renderer.Copy(b.tex,
		&sdl.Rect{W: bulletWidth, H: bulletHeight},
		&sdl.Rect{X: x, Y: y, W: bulletWidth, H: bulletHeight})
}

func (b *bullet) update() {
	// Направление движения пули
	b.x += bulletSpeed * math.Cos(b.angle)
	b.y += bulletSpeed * math.Sin(b.angle)

	// убираем пулю, если она вышла за размер окна
	// b.x > screenWidth || b.x < 0 || b.y > screenHeight || b.y < 0
	if b.y < 0 {
		b.active = false
	}
}

var bulletPool []*bullet

func initBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < 10; i++ {
		b := newBullet(renderer)
		bulletPool = append(bulletPool, &b)
	}
}

func bulletFromPool() (*bullet, bool) {
	for _, b := range bulletPool {
		if !b.active {
			return b, true
		}
	}

	return nil, false
}
