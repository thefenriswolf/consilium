package main

import (
	"image/color"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
)

// State defines the button state (idle, clicked)
type State uint8

const (
	idle State = iota
	clicked
)

// Button object struct
type Button struct {
	posX         int
	posY         int
	width        int
	height       int
	border       int
	text         string
	font         font.Face
	bgColor      color.Color
	textColor    color.Color
	idleColor    color.RGBA
	clickedColor color.RGBA
	state        State
	handlerFunc  func()
}

func withinBounds(cursorX int, cursorY int, posX int, posY int, width int, height int) bool {
	if cursorX >= posX && cursorX <= posX+width && cursorY >= posY && cursorY <= posY+height {
		return true
	}
	return false
}

func (b *Button) buttonClick(screen *ebiten.Image, cursorX int, cursorY int) {
	if withinBounds(cursorX, cursorY, b.posX, b.posY, b.width, b.height) {
		b.state = clicked
		b.handlerFunc()
		b.drawButtonState(screen)
	} else {
		if b.state != idle {
			b.state = idle
		}
	}
}

func (b *Button) drawButtonState(screen *ebiten.Image) {
	var stateColor color.RGBA
	if b.state == idle {
		//	stateColor = Crust
		stateColor = b.idleColor
	}
	if b.state == clicked {
		//	stateColor = FullBlack
		stateColor = b.clickedColor
	}
	vector.StrokeRect(screen, float32(b.posX), float32(b.posY), float32(b.width), float32(b.height), float32(b.border), stateColor, true)
}

// custom button function with state
func (b *Button) drawButton(screen *ebiten.Image) {
	var stateColor color.RGBA
	textX, textY := centerText(b.posX, b.posY, b.width, b.height, b.text, b.font)
	if b.state == idle {
		//	stateColor = Crust
		stateColor = b.idleColor
	}
	if b.state == clicked {
		//	stateColor = FullBlack
		stateColor = b.clickedColor
	}
	vector.DrawFilledRect(screen, float32(b.posX), float32(b.posY), float32(b.width), float32(b.height), b.bgColor, true)
	vector.StrokeRect(screen, float32(b.posX), float32(b.posY), float32(b.width), float32(b.height), float32(b.border), stateColor, true)
	text.Draw(screen, b.text, b.font, textX, textY, b.textColor)
}

type kbdCursor struct {
	X int
	Y int
}

func newKBDCursor() *kbdCursor {
	cursor := new(kbdCursor)
	cursor.X = 0
	cursor.Y = 0
	return cursor
}

type TextBox struct {
	posX    int
	posY    int
	width   int
	height  int
	border  int
	focused bool
	text    string
	//	font        font.Face
	borderColor color.Color
	//	textColor   color.Color
	//	handlerFunc func()
}

func (tb *TextBox) Focus(cursorX int, cursorY int) {
	if withinBounds(cursorX, cursorY, tb.posX, tb.posY, tb.width, tb.height) {
		tb.focused = true
	}
	if withinBounds(KBDC.X, KBDC.Y, tb.posX, tb.posY, tb.width, tb.height) {
		tb.focused = true
	} else {
		tb.focused = false
	}
}

func (tb *TextBox) drawTextbox(screen *ebiten.Image) {
	vector.StrokeRect(screen, float32(tb.posX), float32(tb.posY), float32(tb.width), float32(tb.height), float32(tb.border), tb.borderColor, true)
}

// centerText: center text in any frame of reference
func centerText(frameX int, frameY int, frameWidth int, frameHeight int, text string, fnt font.Face) (int, int) {
	boundsText, _ := font.BoundString(fnt, text)
	centerTextY := ((boundsText.Max.Y + boundsText.Min.Y) / 2).Round()
	centerTextX := ((boundsText.Max.X + boundsText.Min.X) / 2).Round()
	centerButtonX := (frameX + frameWidth) / 2.0
	centerButtonY := (frameY + frameHeight) / 2.0
	textX := centerButtonX - centerTextX + (frameX / 2)
	textY := centerButtonY - centerTextY + frameY
	return textX, textY
}

func inputEventHandler(screen *ebiten.Image, mouseInput bool, kbdInput bool, buttons []*Button, textboxes []*TextBox) {
	if mouseInput {
		cursorX, cursorY := ebiten.CursorPosition()
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			for _, button := range buttons {
				button.buttonClick(screen, cursorX, cursorY)
			}
			wg.Done()
		}()
		go func() {
			for _, textbox := range textboxes {
				if withinBounds(cursorX, cursorY, textbox.posX, textbox.posY, textbox.width, textbox.height) {
					// #+TODO: implement textbox input
					vector.DrawFilledRect(screen, float32(textbox.posX), float32(textbox.posY), float32(textbox.width), float32(textbox.height), Teal, true)
				}
			}
			wg.Done()
		}()
		wg.Wait()
	}
	if kbdInput {
		cursorX := KBDC.X
		cursorY := KBDC.Y
		for _, textbox := range textboxes {
			if withinBounds(cursorX, cursorY, textbox.posX, textbox.posY, textbox.width, textbox.height) {
				// #+TODO: implement textbox input
				// #+TODO: implement vim motions
				vector.DrawFilledRect(screen, float32(textbox.posX), float32(textbox.posY), float32(textbox.width), float32(textbox.height), Teal, true)
			}
		}
	}
}

// // wrapper function around filled and outlined Rectangle
// func (g *Game) drawRect(screen *ebiten.Image, posX int, posY int, sizeX int, sizeY int, strokeWidth int, clr color.Color, antialias bool, fill bool) {
// 	x := float32(posX)
// 	y := float32(posY)
// 	width := float32(sizeX)
// 	height := float32(sizeY)
// 	sw := float32(strokeWidth)
// 	if !fill {
// 		vector.StrokeRect(screen, x, y, width, height, sw, clr, antialias)
// 	}
// 	if fill {
// 		vector.DrawFilledRect(screen, x, y, width, height, clr, antialias)
// 	}
// }

// // wrapper function around fillend and outlined Circle
// func (g *Game) drawCirc(dst *ebiten.Image, posX int, posY int, radius int, strokeWidth int, color color.Color, antialias bool, fill bool) {
// 	x := float32(posX)
// 	y := float32(posY)
// 	r := float32(radius)
// 	sw := float32(strokeWidth)
// 	if !fill {
// 		vector.StrokeCircle(dst, x, y, r, sw, color, antialias)
// 	}
// 	if fill {
// 		vector.DrawFilledCircle(dst, x, y, r, color, antialias)
// 	}
// }
