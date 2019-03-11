package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	shipSpeed        = 0.2
	shipWidth        = 60
	shipHeight       = 48
	shipShotCooldown = time.Millisecond * 250 // 0.25 sec
)

type ship struct {
	tex      *sdl.Texture
	x, y     float64
	lastShot time.Time
}

// newShip - создаем новый корабль
func newShip(renderer *sdl.Renderer) (s ship) {
	s.tex = textureFromBMP(renderer, "images/ship.bmp")
	// Стартовое положение корабля
	s.x = (screenWidth - shipWidth) / 2.0
	s.y = screenHeight - shipHeight

	return s
}

func (s *ship) draw(renderer *sdl.Renderer) {
	renderer.Copy(s.tex,
		&sdl.Rect{W: shipWidth, H: shipHeight},
		&sdl.Rect{X: int32(s.x), Y: int32(s.y - 10), W: shipWidth, H: shipHeight})
	// s.y - 10 Просто для красоты
}

// Обновляем положение корабля
func (s *ship) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		// Перемещение в лево
		if s.x > 0 {
			s.x -= shipSpeed
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		// Перемещение в право
		if s.x < screenWidth-shipWidth {
			s.x += shipSpeed
		}
	}

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(s.lastShot) >= shipShotCooldown { // задержка появления новой пули
			s.shoot(s.x+16, s.y-16)
			s.shoot(s.x+42, s.y-16)
			s.lastShot = time.Now() // время появления пули
		}
	}
}

func (s *ship) shoot(x, y float64) {
	if b, ok := bulletFromPool(); ok {
		b.active = true
		b.x = x
		b.y = y
		b.angle = 270 * (math.Pi / 180) // Переводим в радианы
	}
}
