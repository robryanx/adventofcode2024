package main

import "testing"

// goos: linux
// goarch: amd64
// pkg: github.com/robryanx/adventofcode2024/days/9-2
// cpu: AMD Ryzen 7 7800X3D 8-Core Processor
// BenchmarkSolution-16    	      13	  81443772 ns/op	 1790049 B/op	   21507 allocs/op
func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = solution()
	}
}
