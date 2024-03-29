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
	const (
		buttonWidth        int   = 200
		buttonHeight       int   = 50
		buttonStartY       int   = 5
		buttonBorder       int   = 3
		buttonDefaultState State = idle
	)
	var (
		buttonFont         font.Face  = mpRegular
		buttonBGColor      color.RGBA = Crust
		buttonTextColor    color.RGBA = FullBlack
		buttonIdleColor    color.RGBA = Overlay2
		buttonClickedColor color.RGBA = Blue

		buttonStartX   = ((ScreenWidth - (5 * buttonWidth)) / 5) / 2
		buttonDistance = ScreenWidth / 5
	)
	AboutButton := newButton()
	AboutButton.posX = buttonStartX
	AboutButton.posY = buttonStartY
	AboutButton.width = buttonWidth
	AboutButton.height = buttonHeight
	AboutButton.border = buttonBorder
	AboutButton.text = aboutPageTitle
	AboutButton.font = buttonFont
	AboutButton.bgColor = buttonBGColor
	AboutButton.textColor = buttonTextColor
	AboutButton.idleColor = buttonIdleColor
	AboutButton.clickedColor = buttonClickedColor
	AboutButton.state = buttonDefaultState
	AboutButton.handlerFunc = func() { CurrentPage = About }

	SetupButton := newButton()
	SetupButton.posX = AboutButton.posX + buttonDistance
	SetupButton.posY = buttonStartY
	SetupButton.width = buttonWidth
	SetupButton.height = buttonHeight
	SetupButton.border = buttonBorder
	SetupButton.text = setupPageTitle
	SetupButton.font = buttonFont
	SetupButton.bgColor = buttonBGColor
	SetupButton.textColor = buttonTextColor
	SetupButton.idleColor = buttonIdleColor
	SetupButton.clickedColor = buttonClickedColor
	SetupButton.state = buttonDefaultState
	SetupButton.handlerFunc = func() { CurrentPage = Setup }

	CalendarButton := newButton()
	CalendarButton.posX = SetupButton.posX + buttonDistance
	CalendarButton.posY = buttonStartY
	CalendarButton.width = buttonWidth
	CalendarButton.height = buttonHeight
	CalendarButton.border = buttonBorder
	CalendarButton.text = calendarPageTitle
	CalendarButton.font = buttonFont
	CalendarButton.bgColor = buttonBGColor
	CalendarButton.textColor = buttonTextColor
	CalendarButton.idleColor = buttonIdleColor
	CalendarButton.clickedColor = buttonClickedColor
	CalendarButton.state = buttonDefaultState
	CalendarButton.handlerFunc = func() { CurrentPage = Calendar }

	ExportButton := newButton()
	ExportButton.posX = CalendarButton.posX + buttonDistance
	ExportButton.posY = buttonStartY
	ExportButton.width = buttonWidth
	ExportButton.height = buttonHeight
	ExportButton.border = buttonBorder
	ExportButton.text = exportPageTitle
	ExportButton.font = buttonFont
	ExportButton.bgColor = buttonBGColor
	ExportButton.textColor = buttonTextColor
	ExportButton.idleColor = buttonIdleColor
	ExportButton.clickedColor = buttonClickedColor
	ExportButton.state = buttonDefaultState
	ExportButton.handlerFunc = func() { CurrentPage = Export }

	SettingsButton := newButton()
	SettingsButton.posX = ExportButton.posX + buttonDistance
	SettingsButton.posY = buttonStartY
	SettingsButton.width = buttonWidth
	SettingsButton.height = buttonHeight
	SettingsButton.border = buttonBorder
	SettingsButton.text = settingsPageTitle
	SettingsButton.font = buttonFont
	SettingsButton.bgColor = buttonBGColor
	SettingsButton.textColor = buttonTextColor
	SettingsButton.idleColor = buttonIdleColor
	SettingsButton.clickedColor = buttonClickedColor
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

	imgButton := newIMGButton(imageLoader(logo), 100, 100)
	imgButton.Draw(screen)

	txt := notImplementedYet + "export"
	text.Draw(screen, txt, mpBold, 300, 300, Red)
	AboutButton, SetupButton, CalendarButton, ExportButton, SettingsButton := initNavButtons()
	AboutButton.Draw(screen)
	CalendarButton.Draw(screen)
	SettingsButton.Draw(screen)
	SetupButton.Draw(screen)
	ExportButton.Draw(screen)
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		cursorX, cursorY := ebiten.CursorPosition()
		AboutButton.Click(screen, cursorX, cursorY)
		CalendarButton.Click(screen, cursorX, cursorY)
		SettingsButton.Click(screen, cursorX, cursorY)
		SetupButton.Click(screen, cursorX, cursorY)
		ExportButton.Click(screen, cursorX, cursorY)
	}
}

func calendarPage(screen *ebiten.Image) {
	txt := notImplementedYet + "calendar"
	text.Draw(screen, txt, mpBold, 300, 300, Red)
	AboutButton, SetupButton, CalendarButton, ExportButton, SettingsButton := initNavButtons()
	AboutButton.Draw(screen)
	CalendarButton.Draw(screen)
	SettingsButton.Draw(screen)
	SetupButton.Draw(screen)
	ExportButton.Draw(screen)
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		cursorX, cursorY := ebiten.CursorPosition()
		AboutButton.Click(screen, cursorX, cursorY)
		CalendarButton.Click(screen, cursorX, cursorY)
		SettingsButton.Click(screen, cursorX, cursorY)
		SetupButton.Click(screen, cursorX, cursorY)
		ExportButton.Click(screen, cursorX, cursorY)
	}
}

func setupPage(screen *ebiten.Image) {
	txt := notImplementedYet + "setup"
	text.Draw(screen, txt, mpBold, 300, 300, Red)
	AboutButton, SetupButton, CalendarButton, ExportButton, SettingsButton := initNavButtons()
	AboutButton.Draw(screen)
	CalendarButton.Draw(screen)
	SettingsButton.Draw(screen)
	SetupButton.Draw(screen)
	ExportButton.Draw(screen)
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		cursorX, cursorY := ebiten.CursorPosition()
		AboutButton.Click(screen, cursorX, cursorY)
		CalendarButton.Click(screen, cursorX, cursorY)
		SettingsButton.Click(screen, cursorX, cursorY)
		SetupButton.Click(screen, cursorX, cursorY)
		ExportButton.Click(screen, cursorX, cursorY)
	}
}

func settingsPage(screen *ebiten.Image) {
	txt := notImplementedYet + "settings"
	text.Draw(screen, txt, mpBold, 300, 300, Red)
	AboutButton, SetupButton, CalendarButton, ExportButton, SettingsButton := initNavButtons()
	AboutButton.Draw(screen)
	CalendarButton.Draw(screen)
	SettingsButton.Draw(screen)
	SetupButton.Draw(screen)
	ExportButton.Draw(screen)

	TB := newTextbox()
	TB.posX = 400
	TB.posY = 400
	TB.width = 500
	TB.height = 200
	TB.border = 2
	TB.focused = false
	TB.borderColor = Mauve
	TB.Draw(screen)

	buttons := []*Button{
		AboutButton,
		SetupButton,
		CalendarButton,
		ExportButton,
		SettingsButton}

	textfields := []*TextBox{TB}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		inputEventHandler(screen, true, false, buttons, textfields)
	}
}

// About Page
func aboutPage(screen *ebiten.Image) {
	heading := fmt.Sprintf("%s: %s", WindowTitle, Version)
	headingX, _ := centerText(0, 0, ScreenWidth, ScreenHeight, heading, mpExtraBold)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.4, 0.4)
	op.GeoM.Translate(20, 40)
	screen.DrawImage(thanksImage, op)
	text.Draw(screen, heading, mpExtraBold, headingX, 120, FullBlack)
	text.Draw(screen, thankYouText, mpBold, 45, int(float32(ScreenHeight)/2.5), FullBlack)
	text.Draw(screen, thanksList, mpRegular, 55, int(math.Round(float64(ScreenHeight)/2.2)), FullBlack)

	AboutButton, SetupButton, CalendarButton, ExportButton, SettingsButton := initNavButtons()
	AboutButton.Draw(screen)
	CalendarButton.Draw(screen)
	SettingsButton.Draw(screen)
	SetupButton.Draw(screen)
	ExportButton.Draw(screen)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		cursorX, cursorY := ebiten.CursorPosition()
		AboutButton.Click(screen, cursorX, cursorY)
		CalendarButton.Click(screen, cursorX, cursorY)
		SettingsButton.Click(screen, cursorX, cursorY)
		SetupButton.Click(screen, cursorX, cursorY)
		ExportButton.Click(screen, cursorX, cursorY)
	}
}
