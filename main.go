package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"log"
	// "github.com/hajimehoshi/ebiten/v2/text"
	"bytes"
	"embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image"
	"image/color"
	_ "image/png"
)

//go:embed assets/*
var resources embed.FS

const (
	screenWidth  = 800
	screenHeight = 600
)

var (
	purple color.RGBA = color.RGBA{255, 0, 255, 255}
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
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawRect(screen, 200, 100, 300, 200, purple, true)
	g.drawCirc(screen, 100, 100, 50, 2, Peach, true, true)
	g.drawCirc(screen, 400, 400, 50, 2, Mauve, true, false)
	var keys []string
	for _, k := range g.keys {
		keys = append(keys, k.String())
	}
	msg := fmt.Sprintf("TPS: %0.2f\nKey: \n%s", ebiten.ActualTPS(), keys)
	ebitenutil.DebugPrint(screen, msg)
}

func init() {
	WriteDB("dev.db", "devbucket", "devkey", []byte("devdata"))
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	file, err := resources.ReadFile("assets/logo32x32.png")
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
	// ebiten.SetTPS(30) //set window update rate, default: 60
	ebiten.SetWindowFloating(true)
	ebiten.SetWindowTitle("insomniplan")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
