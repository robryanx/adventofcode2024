package main

import (
	"fmt"
	"strings"

	"github.com/robryanx/adventofcode2024/util"
)

func main() {
	fmt.Println(solution())
}

func solution() int {
	rows, err := util.ReadStrings("19", false, "\n\n")
	if err != nil {
		panic(err)
	}

	patterns := []string{}

	total := 0
	for row := range rows {
		if len(patterns) == 0 {
			patterns = strings.Split(row, ", ")
		} else {
			for _, design := range strings.Split(row, "\n") {
				foundCache := map[string]int{}
				count := backtrackMatchPattern(patterns, design, len(design), foundCache)
				total += count
			}
		}
	}

	return total
}

func backtrackMatchPattern(patterns []string, design string, pos int, foundCache map[string]int) int {
	if pos == 0 {
		return 1
	}
	if count, found := foundCache[design[:pos]]; found {
		return count
	}

	count := 0
	for _, pattern := range patterns {
		if pos-len(pattern) >= 0 && pattern == design[pos-len(pattern):pos] {
			newCount := backtrackMatchPattern(patterns, design, pos-len(pattern), foundCache)
			foundCache[design[:pos-len(pattern)]] = newCount
			count += newCount
		}
	}

	return count
}
