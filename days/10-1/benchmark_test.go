package main

import "testing"

var anchor int

// goos: linux
// goarch: amd64
// pkg: github.com/robryanx/adventofcode2024/days/10-1
// cpu: AMD Ryzen 7 7800X3D 8-Core Processor
// BenchmarkSolution-16    	   29666	     37301 ns/op	   27250 B/op	     117 allocs/op
func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = solution()
	}
}
