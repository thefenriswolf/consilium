package main

import (
	"fmt"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

// boilerplate ebiten game struct
type Game struct {
	keys []ebiten.Key
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
