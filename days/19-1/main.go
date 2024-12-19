package main

import (
	"cmp"
	"fmt"
	"slices"
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
			slices.SortFunc(patterns, func(a, b string) int {
				return cmp.Compare(len(a), len(b))
			})
		} else {
			for _, design := range strings.Split(row, "\n") {
				notFoundCache := map[string]struct{}{}
				if matchPattern(patterns, design, 0, notFoundCache) {
					total++
				}
			}
		}
	}

	return total
}

func matchPattern(patterns []string, design string, pos int, notFoundCache map[string]struct{}) bool {
	for _, pattern := range patterns {
		if pos+len(pattern) <= len(design) && pattern == design[pos:pos+len(pattern)] {
			if _, ok := notFoundCache[design[pos+len(pattern):]]; !ok {
				if pos+len(pattern) == len(design) {
					return true
				} else if matchPattern(patterns, design, pos+len(pattern), notFoundCache) {
					return true
				} else {
					notFoundCache[design[pos+len(pattern):]] = struct{}{}
				}
			}
		}
	}

	return false
}
