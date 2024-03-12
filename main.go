package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	windowSetup()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
