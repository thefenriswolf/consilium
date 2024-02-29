package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Button struct {
	posX      int
	posY      int
	width     int
	height    int
	text      string
	bgColor   color.Color
	textColor color.Color
}

func (b *Button) drawButton(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(b.posX), float32(b.posY), float32(b.width), float32(b.height), b.bgColor, true)
	vector.StrokeRect(screen, float32(b.posX), float32(b.posY), float32(b.width), float32(b.height), 20, Mantle, true)

	//ebitenutil.DrawRect(screen, float64(x+rect.Min.X), float64(y+rect.Min.Y), float64(rect.Max.X-rect.Min.X), float64(rect.Max.Y-rect.Min.Y), bgColor)
	text.Draw(screen, b.text, MP_N_Font, b.posX+int((float32(b.width)*0.1)), b.posY+(b.height/2), b.textColor)
}

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
