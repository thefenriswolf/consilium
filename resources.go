package main

import (
	"embed"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	bolt "go.etcd.io/bbolt"
	"golang.org/x/image/font"
)

// IMG type for iota
type IMG uint8

const (
	logo IMG = iota
	thanks
)

// LANG type for iota
type LANG uint8

const (
	// EN : English
	EN LANG = iota
	// GER : German
	GER
)

// Page type for iota
type Page uint8

const (
	About Page = iota
	Setup
	Calendar
	Export
	Settings
)

// CurrentPage tracker
// defaults to 0 = About
var CurrentPage Page

const (
	// WindowTitle defines the program title or name
	WindowTitle = "consilium"

	// Version defines the release string
	Version = "v20240311-pre-alpha"

	// TPS or FPS which are desired
	TPS = 30

	// Antialias enables/disables antialiasing
	Antialias = true

	// fontDPI recommended by golang opentype package
	fontDPI = 72

	// LANGUAGE to use for the entire program
	// currently supported: (EN, GER)
	LANGUAGE = EN
)

// internationalization
const (
	// Unilang
	thanksList = "- bbolt: https://github.com/etcd-io/bbolt\n" +
		"- Ebitengine: https://github.com/hajimehoshi/ebiten\n" +
		"- M+ Fonts: https://github.com/coz-m/MPLUS_FONTS\n" +
		"- Gonum: https://github.com/gonum/gonum"

	// ENGLISH
	// pages
	aboutPageEN         = "About"
	setupPageEN         = "Setup"
	calendarPageEN      = "Calendar"
	exportPageEN        = "Export"
	settingsPageEN      = "Settings"
	notImplementedYetEN = "Not implemented yet: "
	thankYouEN          = "Special Thanks to:\n"

	// GERMAN
	//pages
	aboutPageGER         = "Über"
	setupPageGER         = "Einrichtung"
	calendarPageGER      = "Kalender"
	exportPageGER        = "Export"
	settingsPageGER      = "Einstellungen"
	notImplementedYetGER = "Noch nicht verfügbar: "
	thankYouGER          = "Wir bedanken uns bei:\n"
)

var (
	// Pages
	aboutPageTitle    string
	setupPageTitle    string
	calendarPageTitle string
	exportPageTitle   string
	settingsPageTitle string
	thankYouText      string
	notImplementedYet string
)

var (
	// baked files

	// images used throughout the program
	// keep list in sync with IMG iota
	images = []string{
		"assets/consilium_logo.png",
		"assets/consilium_thanks.png",
	}
	thanksImage *ebiten.Image

	// listOfLogos: logos that ebiten uses automatically
	listOfLogos = []string{
		"assets/logo16x16.png",
		"assets/logo32x32.png",
		"assets/logo64x64.png"}
	// Logos: variable to hold all logo PNG files
	Logos []image.Image

	// Fonts
	mplus2Fonts = []string{
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
	//mpBlack     font.Face
	mpBold      font.Face
	mpExtraBold font.Face
	// mpExtraLight font.Face
	// mpLight      font.Face
	// mpMedium     font.Face
	// mpSemiBold   font.Face
	// mpThin       font.Face

	// Resources: bundled assets
	//go:embed assets/*
	Resources embed.FS

	// DB: database entry point
	DB *bolt.DB
)

// colors used throughout iplan
var (
	Purple color.RGBA = color.RGBA{255, 0, 255, 255}
	// catppuccin latte
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

var (
	// ScreenWidth: initial screen size: x-axis
	ScreenWidth = 1920
	// ScreenHeight: initial screen size: y-axis
	ScreenHeight = 1080
)
