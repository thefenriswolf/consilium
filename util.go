package main

import (
	"bytes"
	"image"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func withinBounds(cursorX int, cursorY int, posX int, posY int, width int, height int) bool {
	if cursorX >= posX && cursorX <= posX+width && cursorY >= posY && cursorY <= posY+height {
		return true
	}
	return false
}

// centerText: center text in any frame of reference
func centerText(frameX int, frameY int, frameWidth int, frameHeight int, text string, fnt font.Face) (int, int) {
	boundsText, _ := font.BoundString(fnt, text)
	centerTextY := ((boundsText.Max.Y + boundsText.Min.Y) / 2).Round()
	centerTextX := ((boundsText.Max.X + boundsText.Min.X) / 2).Round()
	centerButtonX := int((frameX + frameWidth) / 2.0)
	centerButtonY := int((frameY + frameHeight) / 2.0)
	textX := centerButtonX - centerTextX + (frameX / 2)
	textY := centerButtonY - centerTextY + (frameY / 2)
	return textX, textY
}

// general image loader
func imageLoader(im IMG) *ebiten.Image {
	var imageBuffer image.Image
	file, err := Resources.ReadFile(images[im])
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
