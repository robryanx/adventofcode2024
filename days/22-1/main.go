package main

import (
	"fmt"
	"iter"

	"github.com/robryanx/adventofcode2024/util"
)

func main() {
	fmt.Println(solution())
}

func solution() int {
	rows, err := util.ReadInts("22", false, "\n")
	if err != nil {
		panic(err)
	}

	total := 0
	for secret := range rows {
		seq := sequence(secret)
		next, _ := iter.Pull(seq)

		var secretSequence int
		for i := 0; i < 2000; i++ {
			secretSequence, _ = next()
		}
		total += secretSequence
	}

	return total
}

func sequence(start int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for {
			start = ((start * 64) ^ start) % 16777216
			start = ((start / 32) ^ start) % 16777216
			start = ((start * 2048) ^ start) % 16777216

			if !yield(start) {
				return
			}
		}
	}
}
