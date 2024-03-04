package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type State int64

const (
	idle State = iota
	clicked
)

type Button struct {
	posX        int
	posY        int
	width       int
	height      int
	text        string
	bgColor     color.Color
	textColor   color.Color
	state       State
	handlerFunc func()
}

func (b *Button) buttonClick(screen *ebiten.Image) {
	cursorX, cursorY := ebiten.CursorPosition()
	if cursorX >= b.posX && cursorX <= b.posX+b.width && cursorY >= b.posY && cursorY <= b.posY+b.height {
		b.state = clicked
		b.handlerFunc()
		b.drawButtonState(screen)
		go func() {
			time.Sleep(time.Second * 1)
			b.state = idle
			b.drawButtonState(screen)
		}()
	}
}

func (b *Button) drawButtonState(screen *ebiten.Image) {
	var stateColor color.RGBA
	if b.state == idle {
		//	stateColor = Crust
		stateColor = Purple
	}
	if b.state == clicked {
		//	stateColor = FullBlack
		stateColor = Blue
	}
	vector.StrokeRect(screen, float32(b.posX), float32(b.posY), float32(b.width), float32(b.height), 5, stateColor, true)
}

// custom button function with state
func (b *Button) drawButton(screen *ebiten.Image) {
	var stateColor color.RGBA
	if b.state == idle {
		//	stateColor = Crust
		stateColor = Purple
	}
	if b.state == clicked {
		//	stateColor = FullBlack
		stateColor = Blue
	}
	vector.DrawFilledRect(screen, float32(b.posX), float32(b.posY), float32(b.width), float32(b.height), b.bgColor, true)
	vector.StrokeRect(screen, float32(b.posX), float32(b.posY), float32(b.width), float32(b.height), 5, stateColor, true)
	text.Draw(screen, b.text, MP_N_Font, b.posX+int((float32(b.width)*0.2)), b.posY+int((float32(b.height)/1.7)), b.textColor)
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
