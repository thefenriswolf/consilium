package main

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Page uint8

const (
	Settings Page = iota
	Setup
	Calendar
	Export
	About
)

func pageSelector(screen *ebiten.Image, page Page) {
	switch {
	case page == About:
		aboutPage(screen)
	case page == Export:
		exportPage(screen)
	case page == Calendar:
		calendarPage(screen)
	case page == Setup:
		setupPage(screen)
	case page == Settings:
		settingsPage(screen)
	default:
		aboutPage(screen)
	}
}

func exportPage(screen *ebiten.Image) {
	text.Draw(screen, "export not implemented yet!", mpBold, 300, 300, Red)
}

func calendarPage(screen *ebiten.Image) {
	text.Draw(screen, "calendar not implemented yet!", mpBold, 300, 300, Red)
}

func setupPage(screen *ebiten.Image) {
	text.Draw(screen, "setup not implemented yet!", mpBold, 300, 300, Red)
}

func settingsPage(screen *ebiten.Image) {
	text.Draw(screen, "settings not implemented yet!", mpBold, 300, 300, Red)
}

// About Page
func aboutPage(screen *ebiten.Image) {
	screen.Fill(FullWhite)

	heading := fmt.Sprintf("%s: %s", WindowTitle, Version)
	centerWindow := ScreenWidth / 2
	headingBounds, _ := font.BoundString(mpExtraBold, heading)
	centerHeading := ((headingBounds.Max.X - headingBounds.Min.X) / 2).Round()
	headingX := centerWindow - centerHeading

	thanksImage := imageLoader()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.4, 0.4)
	op.GeoM.Translate(20, 20)
	screen.DrawImage(thanksImage, op)
	text.Draw(screen, heading, mpExtraBold, headingX, 80, FullBlack)
	text.Draw(screen, "Special Thanks to:\n", mpBold, 15, int(float32(ScreenHeight)/2.5), FullBlack)
	thanksList := "- bbolt: https://github.com/etcd-io/bbolt\n" +
		"- Ebitengine: https://github.com/hajimehoshi/ebiten\n" +
		"- M+ Fonts: https://github.com/coz-m/MPLUS_FONTS\n" +
		"- Gonum: https://github.com/gonum/gonum"
	text.Draw(screen, thanksList, mpRegular, 25, int(math.Round(ScreenHeight/2.2)), FullBlack)
	// newButton := &Button{
	// 	posX:        800,
	// 	posY:        800,
	// 	width:       150,
	// 	height:      100,
	// 	text:        "button",
	// 	font:        mpRegular,
	// 	bgColor:     Crust,
	// 	textColor:   FullBlack,
	// 	state:       idle,
	// 	handlerFunc: func() { fmt.Println("clicked") },
	// }
	// newButton.drawButton(screen)
}
