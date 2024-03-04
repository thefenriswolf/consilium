package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

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
	pageSelector(screen, About)
	//g.testContent(screen)
}

func main() {
	windowSetup()
	if err := ebiten.RunGame(NewInstance()); err != nil {
		panic(err)
	}
}
