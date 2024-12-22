package main

import (
	"fmt"
	"math"

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

	count := 0
	sequences := make(map[[4]int]int, 50000)
	for secret := range rows {
		changes := [][2]int{}

		prev := lastDigits(secret, 1)
		secretSequence := secret
		for i := 0; i < 2000; i++ {
			secretSequence = ((secretSequence * 64) ^ secretSequence) % 16777216
			secretSequence = ((secretSequence / 32) ^ secretSequence) % 16777216
			secretSequence = ((secretSequence * 2048) ^ secretSequence) % 16777216

			current := lastDigits(secretSequence, 1)
			changes = append(changes, [2]int{current, current - prev})
			prev = current
		}

		seen := make(map[[4]int]struct{}, 2000)
		for i := 3; i < len(changes); i++ {
			sequence := [4]int{changes[i-3][1], changes[i-2][1], changes[i-1][1], changes[i][1]}
			if _, ok := seen[sequence]; !ok {
				sequences[sequence] += changes[i][0]
				seen[sequence] = struct{}{}
			}
		}

		count++
	}

	best := 0
	for _, total := range sequences {
		if total > best {
			best = total
		}
	}

	return best
}

func lastDigits(num int, pos int) int {
	mask := math.Pow(10, float64(pos))
	lower := num % int(mask)
	return lower
}
