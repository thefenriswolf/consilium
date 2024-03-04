package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
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
	default:
		aboutPage(screen)
	}
}

// About Page
func aboutPage(screen *ebiten.Image) {
	boundsX, boundsY := ebiten.WindowSize()
	screen.Fill(FullWhite)
	heading := fmt.Sprintf("About %s", WindowTitle)
	thanksImage := imageLoader()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.4, 0.4)
	op.GeoM.Translate(20, 20)
	screen.DrawImage(thanksImage, op)
	text.Draw(screen, heading, mpExtraBold, (boundsX/2)-250, 80, FullBlack)
	text.Draw(screen, "Special Thanks to:\n", mpBold, 15, int(float32(boundsY)/2.5), FullBlack)
	thanksList := "bbolt: https://github.com/etcd-io/bbolt\nEbitengine: https://github.com/hajimehoshi/ebiten\nM+ Fonts: https://github.com/coz-m/MPLUS_FONTS\n"
	text.Draw(screen, thanksList, mpRegular, 25, int(float32(boundsY)/2.2), FullBlack)

}
