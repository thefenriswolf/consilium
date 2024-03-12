package main

// // wrapper function around filled and outlined Rectangle
// func (g *Game) drawRect(screen *ebiten.Image, posX int, posY int, sizeX int, sizeY int, strokeWidth int, clr color.Color, antialias bool, fill bool) {
// 	x := float32(posX)
// 	y := float32(posY)
// 	width := float32(sizeX)
// 	height := float32(sizeY)
// 	sw := float32(strokeWidth)
// 	if !fill {
// 		vector.StrokeRect(screen, x, y, width, height, sw, clr, antialias)
// 	}
// 	if fill {
// 		vector.DrawFilledRect(screen, x, y, width, height, clr, antialias)
// 	}
// }

// // wrapper function around fillend and outlined Circle
// func (g *Game) drawCirc(dst *ebiten.Image, posX int, posY int, radius int, strokeWidth int, color color.Color, antialias bool, fill bool) {
// 	x := float32(posX)
// 	y := float32(posY)
// 	r := float32(radius)
// 	sw := float32(strokeWidth)
// 	if !fill {
// 		vector.StrokeCircle(dst, x, y, r, sw, color, antialias)
// 	}
// 	if fill {
// 		vector.DrawFilledCircle(dst, x, y, r, color, antialias)
// 	}
// }
