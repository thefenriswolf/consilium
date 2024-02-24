package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	keys []ebiten.Key
}

func (g *Game) drawRect(screen *ebiten.Image, x_pos int, y_pos int, size_x int, size_y int, clr color.Color, fill bool) {
	if fill {
		for x := x_pos; x < size_x+x_pos; x++ {
			for y := y_pos; y < size_y+y_pos; y++ {
				screen.Set(x, y, clr)
			}
		}
	}
}
func (g *Game) drawCirc(dst *ebiten.Image, x_pos int, y_pos int, radius int, strokeWidth int, color color.Color, antialias bool, fill bool) {
	x := float32(x_pos)
	y := float32(y_pos)
	r := float32(radius)
	sw := float32(strokeWidth)
	if !fill {
		vector.StrokeCircle(dst, x, y, r, sw, color, antialias)
	}
	if fill {
		vector.DrawFilledCircle(dst, x, y, r, color, antialias)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawRect(screen, 200, 100, 300, 200, Lavender, true)
	g.drawCirc(screen, 100, 100, 50, 2, Peach, true, true)
	g.drawCirc(screen, 400, 400, 50, 2, Mauve, true, false)
	var keys []string
	for _, k := range g.keys {
		keys = append(keys, k.String())
	}
	msg := fmt.Sprintf("TPS: %0.2f\nKey: \n%s", ebiten.ActualTPS(), keys)
	ebitenutil.DebugPrint(screen, msg)
	text.Draw(screen, msg, NFont, 20, 40, color.White)
}

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	file, err := Resources.ReadFile("assets/logo32x32.png")
	if err != nil {
		log.Fatal(err)
	}
	logofile, _, err := image.Decode(bytes.NewReader(file))
	if err != nil {
		log.Fatal(err)
	}
	var logos []image.Image
	logos = append(logos, logofile)
	ebiten.SetWindowIcon(logos)
	ebiten.SetTPS(30)                     //set window update rate, default: 60
	ebiten.SetWindowClosingHandled(false) // do stuff when window is about to be closed
	ebiten.SetWindowFloating(true)
	ebiten.SetWindowTitle("insomniplan")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
