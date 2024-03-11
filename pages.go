package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

func initNavButtons() (*Button, *Button, *Button, *Button, *Button) {
	const buttonWidth int = 200
	const buttonHeight int = 50
	const buttonStartY int = 5
	const buttonDefaultState State = idle
	var buttonFont font.Face = mpRegular
	var buttonBGColor color.RGBA = Crust
	var buttonTextColor color.RGBA = FullBlack
	buttonStartX := ((ScreenWidth - (5 * buttonWidth)) / 5) / 2
	buttonDistance := ScreenWidth / 5

	AboutButton := new(Button)
	AboutButton.posX = buttonStartX
	AboutButton.posY = buttonStartY
	AboutButton.width = buttonWidth
	AboutButton.height = buttonHeight
	AboutButton.text = aboutPageTitle
	AboutButton.font = buttonFont
	AboutButton.bgColor = buttonBGColor
	AboutButton.textColor = buttonTextColor
	AboutButton.state = buttonDefaultState
	AboutButton.handlerFunc = func() { CurrentPage = About }

	SetupButton := new(Button)
	SetupButton.posX = AboutButton.posX + buttonDistance
	SetupButton.posY = buttonStartY
	SetupButton.width = buttonWidth
	SetupButton.height = buttonHeight
	SetupButton.text = setupPageTitle
	SetupButton.font = buttonFont
	SetupButton.bgColor = buttonBGColor
	SetupButton.textColor = buttonTextColor
	SetupButton.state = buttonDefaultState
	SetupButton.handlerFunc = func() { CurrentPage = Setup }

	CalendarButton := new(Button)
	CalendarButton.posX = SetupButton.posX + buttonDistance
	CalendarButton.posY = buttonStartY
	CalendarButton.width = buttonWidth
	CalendarButton.height = buttonHeight
	CalendarButton.text = calendarPageTitle
	CalendarButton.font = buttonFont
	CalendarButton.bgColor = buttonBGColor
	CalendarButton.textColor = buttonTextColor
	CalendarButton.state = buttonDefaultState
	CalendarButton.handlerFunc = func() { CurrentPage = Calendar }

	ExportButton := new(Button)
	ExportButton.posX = CalendarButton.posX + buttonDistance
	ExportButton.posY = buttonStartY
	ExportButton.width = buttonWidth
	ExportButton.height = buttonHeight
	ExportButton.text = exportPageTitle
	ExportButton.font = buttonFont
	ExportButton.bgColor = buttonBGColor
	ExportButton.textColor = buttonTextColor
	ExportButton.state = buttonDefaultState
	ExportButton.handlerFunc = func() { CurrentPage = Export }

	SettingsButton := new(Button)
	SettingsButton.posX = ExportButton.posX + buttonDistance
	SettingsButton.posY = buttonStartY
	SettingsButton.width = buttonWidth
	SettingsButton.height = buttonHeight
	SettingsButton.text = settingsPageTitle
	SettingsButton.font = buttonFont
	SettingsButton.bgColor = buttonBGColor
	SettingsButton.textColor = buttonTextColor
	SettingsButton.state = buttonDefaultState
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
	txt := notImplementedYet + "export"
	text.Draw(screen, txt, mpBold, 300, 300, Red)
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
	txt := notImplementedYet + "calendar"
	text.Draw(screen, txt, mpBold, 300, 300, Red)
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
	txt := notImplementedYet + "setup"
	text.Draw(screen, txt, mpBold, 300, 300, Red)
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
	txt := notImplementedYet + "settings"
	text.Draw(screen, txt, mpBold, 300, 300, Red)
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
	heading := fmt.Sprintf("%s: %s", WindowTitle, Version)
	centerWindow := ScreenWidth / 2
	headingBounds, _ := font.BoundString(mpExtraBold, heading)
	centerHeading := ((headingBounds.Max.X - headingBounds.Min.X) / 2).Round()
	headingX := centerWindow - centerHeading

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.4, 0.4)
	op.GeoM.Translate(20, 40)
	screen.DrawImage(thanksImage, op)
	text.Draw(screen, heading, mpExtraBold, headingX, 120, FullBlack)
	text.Draw(screen, thankYouText, mpBold, 45, int(float32(ScreenHeight)/2.5), FullBlack)
	text.Draw(screen, thanksList, mpRegular, 55, int(math.Round(float64(ScreenHeight)/2.2)), FullBlack)

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
