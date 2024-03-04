package main

import (
	"bytes"
	"fmt"
	"image"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// boilerplate ebiten game struct
type Game struct {
	keys []ebiten.Key
}

// game struct constructor
func NewInstance() *Game {
	return &Game{}
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

// general image loader
func imageLoader() *ebiten.Image {
	var imageBuffer image.Image
	file, err := Resources.ReadFile(images[0])
	if err != nil {
		log.Fatal(err)
	}
	imageBuffer, _, err = image.Decode(bytes.NewReader(file))
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(imageBuffer)
}

// load logos from resources
func logoLoader() []image.Image {
	var logoBuffer []image.Image
	for i := range listOfLogos {
		file, err := Resources.ReadFile(listOfLogos[i])
		if err != nil {
			log.Fatal(err)
		}
		logo, _, err := image.Decode(bytes.NewReader(file))
		if err != nil {
			log.Fatal(err)
		}
		logoBuffer = append(logoBuffer, logo)
	}
	return logoBuffer
}

// load fonts from resources
func fontLoader(size int, kind string) font.Face {
	var loadedFont font.Face
	for i := range mplus2Fonts {
		file, err := Resources.ReadFile(mplus2Fonts[i])
		if err != nil {
			log.Fatal(err)
		}
		s, err := opentype.Parse(file)
		if err != nil {
			log.Fatal(err)
		}
		if strings.Contains(mplus2Fonts[i], kind) {
			loadedFont, err = opentype.NewFace(s, &opentype.FaceOptions{
				Size:    float64(size),
				DPI:     fontDPI,
				Hinting: font.HintingFull,
			})
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return loadedFont
}

// a bunch of test content to try drawing
func (g *Game) testContent(screen *ebiten.Image) {
	screen.Fill(colornames.White)
	newButton := &Button{
		posX:        800,
		posY:        800,
		width:       100,
		height:      100,
		text:        "test",
		font:        mpRegular,
		bgColor:     FullWhite,
		textColor:   FullBlack,
		state:       idle,
		handlerFunc: func() { text.Draw(screen, "clicked", mpRegular, 500, 500, FullBlack) },
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
	text.Draw(screen, msg, mpRegular, 40, 40, Sky)
}
