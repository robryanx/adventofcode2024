package main

import "testing"

// goos: linux
// goarch: amd64
// pkg: github.com/robryanx/adventofcode2024/days/8-2
// cpu: AMD Ryzen 7 7800X3D 8-Core Processor
// BenchmarkSolution-16    	   11938	     98318 ns/op	  109220 B/op	     315 allocs/op
func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = solution()
	}
}
