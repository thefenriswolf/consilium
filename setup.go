package main

import (
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Game is the boilerplate ebiten game struct
type Game struct {
	keys                 []ebiten.Key
	windowClosingHandled bool
}

// NewInstance is the constructor for the game object
func newInstance() *Game {
	return &Game{}
}

func setLanguageStrings(LANGUAGE LANG) {
	if LANGUAGE == EN {
		aboutPageTitle = aboutPageEN
		setupPageTitle = setupPageEN
		calendarPageTitle = calendarPageEN
		exportPageTitle = exportPageEN
		settingsPageTitle = settingsPageEN
		thankYouText = thankYouEN
		notImplementedYet = notImplementedYetEN
	}
	if LANGUAGE == GER {
		aboutPageTitle = aboutPageGER
		setupPageTitle = setupPageGER
		calendarPageTitle = calendarPageGER
		exportPageTitle = exportPageGER
		settingsPageTitle = settingsPageGER
		thankYouText = thankYouGER
		notImplementedYet = notImplementedYetGER
	}
}

// boilerplate ebiten function: init stuff
func init() {
	KBDC = newKBDCursor()

	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		setLanguageStrings(LANGUAGE)
		wg.Done()
	}()
	go func() {
		Logos = logoLoader()
		wg.Done()
	}()
	go func() {
		thanksImage = imageLoader(thanks)
		wg.Done()
	}()
	go func() {
		mpRegular = fontLoader(30, "Regular")
		//mpBlack = fontLoader(30, "Black")
		mpBold = fontLoader(30, "Bold")
		mpExtraBold = fontLoader(70, "ExtraBold")
		// mpExtraLight = fontLoader(30, "ExtraLight")
		// mpLight = fontLoader(30, "Light")
		// mpMedium = fontLoader(30, "Medium")
		// mpSemiBold = fontLoader(30, "SemiBold")
		// mpThin = fontLoader(30, "Thin")
		wg.Done()
	}()
	wg.Wait()
	WriteDB("test/dev.DB", "devbucket", "devkey", []byte("devdata"))
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
	ebiten.MaximizeWindow()                                        // maximize window on opening
	ebiten.SetWindowClosingHandled(true)                           // interrupt window closing
}

// game object
var (
	g = newInstance()
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

// // a bunch of test content to try drawing
// func (g *Game) testContent(screen *ebiten.Image) {
// 	screen.Fill(colornames.White)
// 	newButton := &Button{
// 		posX:        800,
// 		posY:        800,
// 		width:       100,
// 		height:      100,
// 		text:        "test",
// 		font:        mpRegular,
// 		bgColor:     FullWhite,
// 		textColor:   FullBlack,
// 		state:       idle,
// 		handlerFunc: func() { text.Draw(screen, "clicked", mpRegular, 500, 500, FullBlack) },
// 	}
// 	newButton.drawButton(screen)

// 	g.drawRect(screen, 200, 100, 300, 200, 14, Lavender, Antialias, true)
// 	g.drawRect(screen, 300, 200, 300, 200, 14, Lavender, Antialias, false)
// 	g.drawCirc(screen, 100, 100, 50, 2, Peach, Antialias, true)
// 	g.drawCirc(screen, 400, 400, 50, 2, Mauve, Antialias, false)
// 	var keys []string
// 	for _, k := range g.keys {
// 		keys = append(keys, k.String())
// 	}
// 	cursorX, cursorY := ebiten.CursorPosition()
// 	var msg string
// 	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
// 		newButton.buttonClick(screen, cursorX, cursorY)
// 		msg = fmt.Sprintf("TPS: %0.2f\nKey: \n%s\nX: %d, Y: %d\n", ebiten.ActualTPS(), keys, cursorX, cursorY)
// 	} else {
// 		newButton.state = idle
// 		msg = fmt.Sprintf("TPS: %0.2f\nKey: \n%s\n", ebiten.ActualTPS(), keys)
// 	}
// 	text.Draw(screen, msg, mpRegular, 40, 40, Sky)
// }
