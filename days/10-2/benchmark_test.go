package main

import "testing"

var anchor int

// goos: linux
// goarch: amd64
// pkg: github.com/robryanx/adventofcode2024/days/10-2
// cpu: AMD Ryzen 7 7800X3D 8-Core Processor
// BenchmarkSolution-16    	   37785	     30906 ns/op	   27251 B/op	     117 allocs/op
func BenchmarkSolution(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = solution()
	}
	anchor = result
}
