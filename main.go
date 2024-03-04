package main

import (
	"fmt"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/colornames"
)

// boilerplate ebiten game struct
type Game struct {
	keys []ebiten.Key
}

func NewInstance() *Game {
	return &Game{}
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
	screen.Fill(colornames.White)
	newButton := &Button{
		posX:        800,
		posY:        800,
		width:       100,
		height:      100,
		text:        "test",
		bgColor:     FullWhite,
		textColor:   FullBlack,
		state:       idle,
		handlerFunc: func() { text.Draw(screen, "clicked", MP_N_Font, 500, 500, FullBlack) },
	}
	newButton.drawButton(screen)

	g.drawRect(screen, 200, 100, 300, 200, 14, Lavender, Antialias, true)
	g.drawRect(screen, 300, 200, 300, 200, 14, Lavender, Antialias, false)
	g.drawCirc(screen, 100, 100, 50, 2, Peach, Antialias, true)
	g.drawCirc(screen, 400, 400, 50, 2, Mauve, Antialias, false)
	var keys []string
	for _, k := range g.keys {
		keys = append(keys, k.String())
	}
	cursorX, cursorY := ebiten.CursorPosition()
	var msg string
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		newButton.buttonClick(screen)
		msg = fmt.Sprintf("TPS: %0.2f\nKey: \n%s\nX: %d, Y: %d\n", ebiten.ActualTPS(), keys, cursorX, cursorY)
	} else {
		newButton.state = idle
		msg = fmt.Sprintf("TPS: %0.2f\nKey: \n%s\n", ebiten.ActualTPS(), keys)
	}
	text.Draw(screen, msg, MP_N_Font, 40, 40, Sky)
}

// setup window
func windowSetup() {
	ebiten.SetWindowSize(ScreenWidth/2, ScreenHeight/2)            // set initial screen size
	ebiten.SetCursorMode(ebiten.CursorModeVisible)                 // make sure cursor is visible
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled) // enable window resizing by the user
	ebiten.SetWindowIcon(Logos)                                    // set window icon
	ebiten.SetTPS(TPS)                                             //set window update rate, default: 60
	ebiten.SetWindowClosingHandled(false)                          // do stuff when window is about to be closed
	ebiten.SetWindowFloating(true)                                 // set default window state
	ebiten.SetWindowTitle(WindowTitle)                             // set window title
	ebiten.MaximizeWindow()
}

func main() {
	windowSetup()
	if err := ebiten.RunGame(NewInstance()); err != nil {
		panic(err)
	}
}
