package main

import "testing"

// goos: linux
// goarch: amd64
// pkg: github.com/robryanx/adventofcode2024/days/9-1
// cpu: AMD Ryzen 7 7800X3D 8-Core Processor
// BenchmarkSolution-16    	    4672	    236383 ns/op	  217954 B/op	      12 allocs/op
func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = solution()
	}
}
