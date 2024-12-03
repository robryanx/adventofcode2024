package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/robryanx/adventofcode2024/util"
)

var mulRegex = regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)

func main() {
	bytes, err := util.ReadBytes("3", false)
	if err != nil {
		panic(err)
	}

	total := 0
	mulList := mulRegex.FindAllSubmatch(bytes, -1)
	for _, mul := range mulList {
		numA, err := strconv.Atoi(string(mul[1]))
		if err != nil {
			panic(err)
		}

		numB, err := strconv.Atoi(string(mul[2]))
		if err != nil {
			panic(err)
		}

		total += numA * numB
	}

	fmt.Println(total)
}
