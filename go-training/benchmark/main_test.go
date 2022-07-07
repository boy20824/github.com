package main

import "testing"

func BenchmarkPrint01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		print01(100)
	}
}
