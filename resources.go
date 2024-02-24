package main

import (
	"embed"
	"image/color"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed assets/*
var Resources embed.FS

const (
	DPI          = 72
	ScreenWidth  = 800
	ScreenHeight = 600
)

var (
	NFont font.Face
)

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
)

func init() {
	//WriteDB("test/dev.db", "devbucket", "devkey", []byte("devdata"))
	file, err := Resources.ReadFile("assets/fonts/otf/Mplus1-Regular.otf")
	if err != nil {
		log.Fatal(err)
	}
	s, err := opentype.Parse(file)
	if err != nil {
		log.Fatal(err)
	}
	NFont, err = opentype.NewFace(s, &opentype.FaceOptions{
		Size:    11,
		DPI:     DPI,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}
