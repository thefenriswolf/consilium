package main

import (
	"image/color"

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

func newButton() *Button {
	btn := new(Button)
	btn.posX = 200
	btn.posY = 200
	btn.width = 200
	btn.height = 50
	btn.border = 3
	btn.text = ""
	btn.font = mpRegular
	btn.bgColor = Mauve
	btn.textColor = FullBlack
	btn.idleColor = Sky
	btn.clickedColor = Purple
	btn.state = idle
	return btn
}

func (b *Button) Click(screen *ebiten.Image, cursorX int, cursorY int) {
	if withinBounds(cursorX, cursorY, b.posX, b.posY, b.width, b.height) {
		b.state = clicked
		b.handlerFunc()
		b.Update(screen)
	} else {
		if b.state != idle {
			b.state = idle
		}
	}
}

func (b *Button) Update(screen *ebiten.Image) {
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
func (b *Button) Draw(screen *ebiten.Image) {
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

// image Button object struct
type IMGButton struct {
	posX        int
	posY        int
	width       int
	height      int
	border      int
	image       *ebiten.Image
	imgScalingX float64
	imgScalingY float64
	borderColor color.Color
	handlerFunc func()
}

// image button constructor function
func newIMGButton(image *ebiten.Image, width int, height int) *IMGButton {
	btn := new(IMGButton)
	btn.posX = 400
	btn.posY = 400
	btn.width = width
	btn.height = height
	btn.border = 3
	btn.image = image
	imgheight := float64(btn.image.Bounds().Dy())
	imgwidth := float64(btn.image.Bounds().Dx())
	btn.imgScalingX = float64(btn.width) / imgwidth
	btn.imgScalingY = float64(btn.height) / imgheight
	btn.borderColor = FullBlack
	return btn
}

// custom image button function with state
func (b *IMGButton) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(b.imgScalingX, b.imgScalingY)
	op.GeoM.Translate(float64(b.posX), float64(b.posY))
	screen.DrawImage(thanksImage, op)
	vector.StrokeRect(screen, float32(b.posX), float32(b.posY), float32(b.width), float32(b.height), float32(b.border), b.borderColor, true)
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
	posX        int
	posY        int
	width       int
	height      int
	border      int
	focused     bool
	text        string
	font        font.Face
	borderColor color.Color
	textColor   color.Color
	//	handlerFunc func()
}

func newTextbox() *TextBox {
	tbx := new(TextBox)
	tbx.posX = 400
	tbx.posY = 400
	tbx.width = 200
	tbx.height = 100
	tbx.border = 2
	tbx.focused = false
	tbx.text = "text"
	tbx.font = mpRegular
	tbx.borderColor = Mauve
	tbx.textColor = FullBlack
	//	handlerFunc func()
	return tbx
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

func (tb *TextBox) Draw(screen *ebiten.Image) {
	textX, textY := centerText(tb.posX, tb.posY, tb.width, tb.height, tb.text, tb.font)
	vector.StrokeRect(screen, float32(tb.posX), float32(tb.posY), float32(tb.width), float32(tb.height), float32(tb.border), tb.borderColor, true)
	text.Draw(screen, tb.text, tb.font, textX, textY, tb.textColor)
}

func (tb *TextBox) Update(screen *ebiten.Image, msg string) {
	tb.text = msg
	textX, textY := centerText(tb.posX, tb.posY, tb.width, tb.height, tb.text, tb.font)
	text.Draw(screen, tb.text, tb.font, textX, textY, tb.textColor)

}
