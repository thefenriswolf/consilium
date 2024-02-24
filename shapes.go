package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// wrapper function around filled and outlined Rectangle
func (g *Game) drawRect(screen *ebiten.Image, x_pos int, y_pos int, size_x int, size_y int, strokeWidth int, clr color.Color, antialias bool, fill bool) {
	x := float32(x_pos)
	y := float32(y_pos)
	width := float32(size_x)
	height := float32(size_y)
	sw := float32(strokeWidth)
	if !fill {
		vector.StrokeRect(screen, x, y, width, height, sw, clr, antialias)
	}
	if fill {
		vector.DrawFilledRect(screen, x, y, width, height, clr, antialias)
	}
}

// wrapper function around fillend and outlined Circle
func (g *Game) drawCirc(dst *ebiten.Image, x_pos int, y_pos int, radius int, strokeWidth int, color color.Color, antialias bool, fill bool) {
	x := float32(x_pos)
	y := float32(y_pos)
	r := float32(radius)
	sw := float32(strokeWidth)
	if !fill {
		vector.StrokeCircle(dst, x, y, r, sw, color, antialias)
	}
	if fill {
		vector.DrawFilledCircle(dst, x, y, r, color, antialias)
	}
}
