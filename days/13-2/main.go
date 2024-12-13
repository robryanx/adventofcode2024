package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"

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

		A := mat.NewDense(2, 2, []float64{float64(steps[0][0]), float64(steps[1][0]), float64(steps[0][1]), float64(steps[1][1])})
		b := mat.NewVecDense(2, []float64{float64(steps[2][0]), float64(steps[2][1])})

		var x mat.VecDense
		if err := x.SolveVec(A, b); err != nil {
			panic(err)
		}

		roundedI := math.Round(x.RawVector().Data[0]*10) / 10
		roundedJ := math.Round(x.RawVector().Data[1]*10) / 10

		if roundedI == float64(int64(roundedI)) &&
			roundedJ == float64(int64(roundedJ)) {
			total += int(math.Round(x.RawVector().Data[0]))*3 + int(math.Round(x.RawVector().Data[1]))
		}
	}

	return total
}
