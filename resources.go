package main

import (
	"embed"
	"image"
	"image/color"
	"sync"

	"golang.org/x/image/font"
)

type IMG uint8

const (
	// program title or name
	WindowTitle = "consilium"

	// release version
	Version = "v20240307-pre-alpha"

	//TPS
	TPS = 30

	// enable/disable antialiasing
	Antialias = true

	// recommended by golang opentype package
	fontDPI = 72

	// initial screen size
	ScreenWidth  = 1920
	ScreenHeight = 1080
)

const (
	logo IMG = iota
	thanks
)

var (
	// baked files

	// images used throughout the program
	// keep list in sync with IMG iota
	images = []string{
		"assets/consilium_logo.png",
		"assets/consilium_thanks.png",
	}
	// LOGOS
	listOfLogos = [3]string{
		"assets/logo16x16.png",
		"assets/logo32x32.png",
		"assets/logo64x64.png"}
	// variable to hold all logo PNG files
	Logos []image.Image

	// Fonts
	mplus2Fonts = [9]string{
		"assets/fonts/otf/Mplus2-Black.otf",
		"assets/fonts/otf/Mplus2-Bold.otf",
		"assets/fonts/otf/Mplus2-ExtraBold.otf",
		"assets/fonts/otf/Mplus2-ExtraLight.otf",
		"assets/fonts/otf/Mplus2-Light.otf",
		"assets/fonts/otf/Mplus2-Medium.otf",
		"assets/fonts/otf/Mplus2-Regular.otf",
		"assets/fonts/otf/Mplus2-SemiBold.otf",
		"assets/fonts/otf/Mplus2-Thin.otf"}

	// globaly availabe fonts
	mpRegular font.Face // MPlus regular
	// mpBlack      font.Face
	mpBold      font.Face
	mpExtraBold font.Face
	// mpExtraLight font.Face
	// mpLight      font.Face
	// mpMedium     font.Face
	// mpSemiBold   font.Face
	// mpThin       font.Face

	// bundled stuff
	//go:embed assets/*
	Resources embed.FS

	// current page tracker
	CurrentPage Page
)

// colors used throughout iplan
var (
	Purple color.RGBA = color.RGBA{255, 0, 255, 255}
	// source: catppuccin latte
	Rosewater color.RGBA = color.RGBA{220, 138, 120, 255}
	Flamingo  color.RGBA = color.RGBA{221, 120, 120, 255}
	Pink      color.RGBA = color.RGBA{234, 118, 203, 255}
	Mauve     color.RGBA = color.RGBA{136, 57, 239, 255}
	Red       color.RGBA = color.RGBA{210, 15, 57, 255}
	Maroon    color.RGBA = color.RGBA{230, 69, 83, 255}
	Peach     color.RGBA = color.RGBA{254, 100, 11, 255}
	Yellow    color.RGBA = color.RGBA{223, 142, 29, 255}
	Green     color.RGBA = color.RGBA{64, 160, 43, 255}
	Teal      color.RGBA = color.RGBA{23, 146, 153, 255}
	Sky       color.RGBA = color.RGBA{4, 165, 229, 255}
	Sapphire  color.RGBA = color.RGBA{32, 159, 181, 255}
	Blue      color.RGBA = color.RGBA{30, 102, 245, 255}
	Lavender  color.RGBA = color.RGBA{114, 135, 253, 255}
	Text      color.RGBA = color.RGBA{76, 79, 105, 255}
	FullWhite color.RGBA = color.RGBA{255, 255, 255, 255}
	Subtext1  color.RGBA = color.RGBA{92, 95, 119, 255}
	Subtext0  color.RGBA = color.RGBA{108, 111, 133, 255}
	Overlay2  color.RGBA = color.RGBA{124, 127, 147, 255}
	Overlay1  color.RGBA = color.RGBA{140, 143, 161, 255}
	Overlay0  color.RGBA = color.RGBA{156, 160, 176, 255}
	Surface2  color.RGBA = color.RGBA{172, 176, 190, 255}
	Surface1  color.RGBA = color.RGBA{188, 192, 204, 255}
	Surface0  color.RGBA = color.RGBA{204, 208, 218, 255}
	Base      color.RGBA = color.RGBA{239, 241, 245, 255}
	Mantle    color.RGBA = color.RGBA{230, 233, 239, 255}
	Crust     color.RGBA = color.RGBA{220, 224, 232, 255}
	Black     color.RGBA = color.RGBA{17, 17, 27, 255}
	FullBlack color.RGBA = color.RGBA{0, 0, 0, 255}
)

// boilerplate ebiten function: init stuff
func init() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		Logos = logoLoader()
		wg.Done()
	}()
	go func() {
		mpRegular = fontLoader(30, "Regular")
		// mpBlack = fontLoader(30, "Black")
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
	//WriteDB("test/dev.db", "devbucket", "devkey", []byte("devdata"))
}
