package main

import (
	"bytes"
	"embed"
	"fmt"
	"image"
	"image/color"
	"log"
	"strings"
	"sync"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"gonum.org/v1/gonum/mat"
)

const (
	// program title/name
	WindowTitle = "consilium"

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

var (

	// baked files
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
	MP_N_Font font.Face // MPlus regular

	// bundled stuff
	//go:embed assets/*
	Resources embed.FS
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

// load fonts from resources
func fontLoader(size int, kind string) font.Face {
	zero := mat.NewDense(3, 5, nil)
	fmt.Println(*zero)
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

// boilerplate ebiten function: init stuff
func init() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		Logos = logoLoader()
		wg.Done()
	}()
	go func() {
		MP_N_Font = fontLoader(30, "ExtraLight")
		wg.Done()
	}()
	wg.Wait()
	//WriteDB("test/dev.db", "devbucket", "devkey", []byte("devdata"))
}
