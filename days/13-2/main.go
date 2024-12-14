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

		steps[2][0] += 10000000000000
		steps[2][1] += 10000000000000

		solved, i, j := solveEquations([3]int{steps[0][0], steps[1][0], steps[2][0]}, [3]int{steps[0][1], steps[1][1], steps[2][1]})
		if solved {
			total += i*3 + j
		}
	}

	return total
}

// solve for x and y given 2 equations of the form
// 1x + 2y = 9
func solveEquations(equationA, equationB [3]int) (bool, int, int) {
	detCoef := equationA[0]*equationB[1] - equationA[1]*equationB[0]
	detX := equationA[2]*equationB[1] - equationB[2]*equationA[1]
	detY := equationB[2]*equationA[0] - equationA[2]*equationB[0]

	if detCoef == 0 {
		return false, 0, 0
	}

	x := float64(detX) / float64(detCoef)
	y := float64(detY) / float64(detCoef)

	if x != float64(int64(x)) || y != float64(int64(y)) {
		return false, 0, 0
	}

	return true, detX / detCoef, detY / detCoef
}
