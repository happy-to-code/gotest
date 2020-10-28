package main

import "testing"

func TestDemo(t *testing.T) {
	Demo()
}

func BenchmarkDemo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Demo()
	}
}
