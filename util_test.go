package main

import (
	"testing"
)

type objs struct {
	cursorX int
	cursorY int
	posX    int
	posY    int
	width   int
	height  int
	success bool
}

var objects = []objs{
	{200, 200, 200, 200, 100, 100, true},  // mouse in top left corner
	{200, 300, 200, 200, 100, 100, true},  // mouse in bottom left corner
	{200, 305, 200, 200, 100, 100, false}, // mouse out of bounds in y direction

}

func TestWithinBounds(t *testing.T) {
	for _, e := range objects {
		want := e.success
		got := withinBounds(e.cursorX, e.cursorY, e.posX, e.posY, e.width, e.height)
		if got != want {
			t.Fatalf("Got: %t, wanted: %t", got, want)
		}
	}

}

func BenchmarkImageLoader(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = imageLoader(thanks)
		_ = imageLoader(logo)
	}
}

// func BenchmarkFontLoading(b *testing.B) {
// 	for _, v := range fontNames {
// 		b.Run(fmt.Sprintf("font: %s", v), func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				_ = fontLoader(30, v)
// 			}
// 		})
// 	}
// }

// func BenchmarkResourceloadingSequential(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		_ = logoLoader()
// 		_ = fontLoader(30, "ExtraLight")
// 	}
// }
