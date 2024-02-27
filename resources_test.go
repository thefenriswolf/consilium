package main

import (
	"fmt"
	"sync"
	"testing"
)

var fontNames = []string{
	"Black",
	"Bold",
	"ExtraBold",
	"ExtraLight",
	"Light",
	"Medium",
	"Regular",
	"SemiBold",
	"Thin"}

func BenchmarkFontLoading(b *testing.B) {
	for _, v := range fontNames {
		b.Run(fmt.Sprintf("font: %s", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = fontLoader(30, v)
			}
		})
	}
}

func BenchmarkResourceloadingSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = logoLoader()
		_ = fontLoader(30, "ExtraLight")
	}
}

func BenchmarkResourceloadingGoroutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			_ = logoLoader()
			wg.Done()
		}()
		go func() {
			_ = fontLoader(40, "Regular")
			wg.Done()
		}()
		wg.Wait()
	}
}
