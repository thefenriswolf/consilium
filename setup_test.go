package main

import "testing"

func BenchmarkInitA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InitA()
	}
}

func BenchmarkInitB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InitB()
	}
}
