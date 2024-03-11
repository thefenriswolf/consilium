package main

import (
	_ "image/png"
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// game object
var (
	g = NewInstance()
)

// Layout is a boilerplate ebiten function: returns screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

// Update is a boilerplate ebiten function: runs every tick/frame
func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	// handle window closing
	if ebiten.IsWindowBeingClosed() {
		g.windowClosingHandled = true
	}
	if g.windowClosingHandled {
		// make window closing wait for program save and cleanup
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			windowClosingHandler()
			wg.Done()
		}()
		wg.Wait()
		return ebiten.Termination
	}
	return nil
}

// Draw is a boilerplate ebiten function: draws stuff to screen once
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(FullWhite)
	pageSelector(screen, CurrentPage)
	//g.testContent(screen)
}

func main() {
	windowSetup()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
