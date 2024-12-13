package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2024/util"
)

func main() {
	fmt.Println(solution())
}

var buttonRegex = regexp.MustCompile(`Button [AB]{1}: X\+(\d+), Y\+(\d+)`)
var prizeRegex = regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

func solution() int {
	groups, err := util.ReadStrings("13", false, "\n\n")
	if err != nil {
		panic(err)
	}

	total := 0
	for group := range groups {
		lines := strings.Split(group, "\n")

		var steps [3][2]int
		for i := 0; i < 3; i++ {
			var matches []string
			if i != 2 {
				matches = buttonRegex.FindStringSubmatch(lines[i])
			} else {
				matches = prizeRegex.FindStringSubmatch(lines[i])
			}

			y, err := strconv.Atoi(matches[2])
			if err != nil {
				panic(err)
			}

			x, err := strconv.Atoi(matches[1])
			if err != nil {
				panic(err)
			}

			steps[i] = [2]int{x, y}
		}

		score := -1
	loop:
		for i := 0; i <= 100; i++ {
			for j := 0; j <= 100; j++ {
				if steps[0][0]*i+steps[1][0]*j == steps[2][0] &&
					steps[0][1]*i+steps[1][1]*j == steps[2][1] {
					score = i*3 + j
					break loop
				}
			}
		}

		if score != -1 {
			total += score
		}
	}

	return total
}
