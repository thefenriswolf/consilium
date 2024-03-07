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
	About Page = iota
	Setup
	Calendar
	Export
	Settings
)

func initNavButtons() (*Button, *Button, *Button, *Button, *Button) {
	AboutButton := new(Button)
	AboutButton.posX = 30
	AboutButton.posY = 5
	AboutButton.width = 210
	AboutButton.height = 50
	AboutButton.text = "About"
	AboutButton.font = mpRegular
	AboutButton.bgColor = Crust
	AboutButton.textColor = FullBlack
	AboutButton.state = idle
	AboutButton.handlerFunc = func() { CurrentPage = About }

	SetupButton := new(Button)
	SetupButton.posX = 245
	SetupButton.posY = 5
	SetupButton.width = 200
	SetupButton.height = 50
	SetupButton.text = "Setup"
	SetupButton.font = mpRegular
	SetupButton.bgColor = Crust
	SetupButton.textColor = FullBlack
	SetupButton.state = idle
	SetupButton.handlerFunc = func() { CurrentPage = Setup }

	CalendarButton := new(Button)
	CalendarButton.posX = 450
	CalendarButton.posY = 5
	CalendarButton.width = 200
	CalendarButton.height = 50
	CalendarButton.text = "Calendar"
	CalendarButton.font = mpRegular
	CalendarButton.bgColor = Crust
	CalendarButton.textColor = FullBlack
	CalendarButton.state = idle
	CalendarButton.handlerFunc = func() { CurrentPage = Calendar }

	ExportButton := new(Button)
	ExportButton.posX = 655
	ExportButton.posY = 5
	ExportButton.width = 200
	ExportButton.height = 50
	ExportButton.text = "Export"
	ExportButton.font = mpRegular
	ExportButton.bgColor = Crust
	ExportButton.textColor = FullBlack
	ExportButton.state = idle
	ExportButton.handlerFunc = func() { CurrentPage = Export }

	SettingsButton := new(Button)
	SettingsButton.posX = 850
	SettingsButton.posY = 5
	SettingsButton.width = 200
	SettingsButton.height = 50
	SettingsButton.text = "Settings"
	SettingsButton.font = mpRegular
	SettingsButton.bgColor = Crust
	SettingsButton.textColor = FullBlack
	SettingsButton.state = idle
	SettingsButton.handlerFunc = func() { CurrentPage = Settings }

	return AboutButton, SetupButton, CalendarButton, ExportButton, SettingsButton
}

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
	AboutButton, SetupButton, CalendarButton, ExportButton, SettingsButton := initNavButtons()
	AboutButton.drawButton(screen)
	CalendarButton.drawButton(screen)
	SettingsButton.drawButton(screen)
	SetupButton.drawButton(screen)
	ExportButton.drawButton(screen)
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		cursorX, cursorY := ebiten.CursorPosition()
		AboutButton.buttonClick(screen, cursorX, cursorY)
		CalendarButton.buttonClick(screen, cursorX, cursorY)
		SettingsButton.buttonClick(screen, cursorX, cursorY)
		SetupButton.buttonClick(screen, cursorX, cursorY)
		ExportButton.buttonClick(screen, cursorX, cursorY)
	}
}

func calendarPage(screen *ebiten.Image) {
	text.Draw(screen, "calendar not implemented yet!", mpBold, 300, 300, Red)
	AboutButton, SetupButton, CalendarButton, ExportButton, SettingsButton := initNavButtons()
	AboutButton.drawButton(screen)
	CalendarButton.drawButton(screen)
	SettingsButton.drawButton(screen)
	SetupButton.drawButton(screen)
	ExportButton.drawButton(screen)
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		cursorX, cursorY := ebiten.CursorPosition()
		AboutButton.buttonClick(screen, cursorX, cursorY)
		CalendarButton.buttonClick(screen, cursorX, cursorY)
		SettingsButton.buttonClick(screen, cursorX, cursorY)
		SetupButton.buttonClick(screen, cursorX, cursorY)
		ExportButton.buttonClick(screen, cursorX, cursorY)
	}
}

func setupPage(screen *ebiten.Image) {
	text.Draw(screen, "setup not implemented yet!", mpBold, 300, 300, Red)
	AboutButton, SetupButton, CalendarButton, ExportButton, SettingsButton := initNavButtons()
	AboutButton.drawButton(screen)
	CalendarButton.drawButton(screen)
	SettingsButton.drawButton(screen)
	SetupButton.drawButton(screen)
	ExportButton.drawButton(screen)
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		cursorX, cursorY := ebiten.CursorPosition()
		AboutButton.buttonClick(screen, cursorX, cursorY)
		CalendarButton.buttonClick(screen, cursorX, cursorY)
		SettingsButton.buttonClick(screen, cursorX, cursorY)
		SetupButton.buttonClick(screen, cursorX, cursorY)
		ExportButton.buttonClick(screen, cursorX, cursorY)
	}
}

func settingsPage(screen *ebiten.Image) {
	text.Draw(screen, "settings not implemented yet!", mpBold, 300, 300, Red)
	AboutButton, SetupButton, CalendarButton, ExportButton, SettingsButton := initNavButtons()
	AboutButton.drawButton(screen)
	CalendarButton.drawButton(screen)
	SettingsButton.drawButton(screen)
	SetupButton.drawButton(screen)
	ExportButton.drawButton(screen)
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		cursorX, cursorY := ebiten.CursorPosition()
		AboutButton.buttonClick(screen, cursorX, cursorY)
		CalendarButton.buttonClick(screen, cursorX, cursorY)
		SettingsButton.buttonClick(screen, cursorX, cursorY)
		SetupButton.buttonClick(screen, cursorX, cursorY)
		ExportButton.buttonClick(screen, cursorX, cursorY)
	}
}

// About Page
func aboutPage(screen *ebiten.Image) {
	screen.Fill(FullWhite)

	heading := fmt.Sprintf("%s: %s", WindowTitle, Version)
	centerWindow := ScreenWidth / 2
	headingBounds, _ := font.BoundString(mpExtraBold, heading)
	centerHeading := ((headingBounds.Max.X - headingBounds.Min.X) / 2).Round()
	headingX := centerWindow - centerHeading

	thanksImage := imageLoader(thanks)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.4, 0.4)
	op.GeoM.Translate(20, 40)
	screen.DrawImage(thanksImage, op)
	text.Draw(screen, heading, mpExtraBold, headingX, 120, FullBlack)
	text.Draw(screen, "Special Thanks to:\n", mpBold, 45, int(float32(ScreenHeight)/2.5), FullBlack)
	thanksList := "- bbolt: https://github.com/etcd-io/bbolt\n" +
		"- Ebitengine: https://github.com/hajimehoshi/ebiten\n" +
		"- M+ Fonts: https://github.com/coz-m/MPLUS_FONTS\n" +
		"- Gonum: https://github.com/gonum/gonum"
	text.Draw(screen, thanksList, mpRegular, 55, int(math.Round(ScreenHeight/2.2)), FullBlack)

	AboutButton, SetupButton, CalendarButton, ExportButton, SettingsButton := initNavButtons()
	AboutButton.drawButton(screen)
	CalendarButton.drawButton(screen)
	SettingsButton.drawButton(screen)
	SetupButton.drawButton(screen)
	ExportButton.drawButton(screen)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		cursorX, cursorY := ebiten.CursorPosition()
		AboutButton.buttonClick(screen, cursorX, cursorY)
		CalendarButton.buttonClick(screen, cursorX, cursorY)
		SettingsButton.buttonClick(screen, cursorX, cursorY)
		SetupButton.buttonClick(screen, cursorX, cursorY)
		ExportButton.buttonClick(screen, cursorX, cursorY)
	}
}
