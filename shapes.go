package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
)

type State uint8

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
	font        font.Face
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
	boundsText, _ := font.BoundString(b.font, b.text)
	centerTextY := ((boundsText.Max.Y - boundsText.Min.Y) / 2).Round()
	centerTextX := ((boundsText.Max.X - boundsText.Min.X) / 2).Round()
	centerButtonX := ((b.posX + b.width) - b.posX) / 2.0
	centerButtonY := ((b.posY + b.height) - b.posY) / 2.0
	textX := centerButtonX - centerTextX + b.posX
	textY := int(float32(centerButtonY)*1.4) - centerTextY + b.posY
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
	text.Draw(screen, b.text, b.font, textX, textY, b.textColor)
}

// wrapper function around filled and outlined Rectangle
func (g *Game) drawRect(screen *ebiten.Image, posX int, posY int, sizeX int, sizeY int, strokeWidth int, clr color.Color, antialias bool, fill bool) {
	x := float32(posX)
	y := float32(posY)
	width := float32(sizeX)
	height := float32(sizeY)
	sw := float32(strokeWidth)
	if !fill {
		vector.StrokeRect(screen, x, y, width, height, sw, clr, antialias)
	}
	if fill {
		vector.DrawFilledRect(screen, x, y, width, height, clr, antialias)
	}
}

// wrapper function around fillend and outlined Circle
func (g *Game) drawCirc(dst *ebiten.Image, posX int, posY int, radius int, strokeWidth int, color color.Color, antialias bool, fill bool) {
	x := float32(posX)
	y := float32(posY)
	r := float32(radius)
	sw := float32(strokeWidth)
	if !fill {
		vector.StrokeCircle(dst, x, y, r, sw, color, antialias)
	}
	if fill {
		vector.DrawFilledCircle(dst, x, y, r, color, antialias)
	}
}
