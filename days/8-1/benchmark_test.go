package main

import "testing"

// goos: linux
// goarch: amd64
// pkg: github.com/robryanx/adventofcode2024/days/8-1
// cpu: AMD Ryzen 7 7800X3D 8-Core Processor
// BenchmarkSolution-16    	   24702	     46903 ns/op	   44145 B/op	     279 allocs/op
func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = solution()
	}
}
