package main

import (
	"fmt"
	"image/color"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// boilerplate ebiten game struct
type Game struct {
	keys []ebiten.Key
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

// boilerplate ebiten function: returns screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

// boilerplate ebiten function: runs every tick/frame
func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	return nil
}

// boilerplate ebiten function: draws stuff
func (g *Game) Draw(screen *ebiten.Image) {
	g.drawRect(screen, 200, 100, 300, 200, 14, Lavender, AA, true)
	g.drawRect(screen, 300, 200, 300, 200, 14, Lavender, AA, false)
	g.drawCirc(screen, 100, 100, 50, 2, Peach, AA, true)
	g.drawCirc(screen, 400, 400, 50, 2, Mauve, AA, false)
	var keys []string
	for _, k := range g.keys {
		keys = append(keys, k.String())
	}
	msg := fmt.Sprintf("TPS: %0.2f\nKey: \n%s", ebiten.ActualTPS(), keys)
	ebitenutil.DebugPrint(screen, msg)
	text.Draw(screen, msg, MP_N_Font, 40, 40, FullWhite)
}

// setup window
func windowSetup() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)                // set initial screen size
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled) // enable window resizing by the user
	ebiten.SetWindowIcon(Logos)                                    // set window icon
	ebiten.SetTPS(TPS)                                             //set window update rate, default: 60
	ebiten.SetWindowClosingHandled(false)                          // do stuff when window is about to be closed
	ebiten.SetWindowFloating(true)                                 // set default window state
	ebiten.SetWindowTitle(WindowTitle)                             // set window title
}

func main() {
	windowSetup()
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
