package main

import "testing"

func TestGCTest(t *testing.T) {
	GCTest()
}

func BenchmarkGCTest(b *testing.B) {
	for i := 0; i < 20; i++ {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				GCTest()
			}
		})
	}

}
