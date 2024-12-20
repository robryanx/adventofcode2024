package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2024/util"
)

func main() {
	strs, err := util.ReadStrings("1", false, "\n")
	if err != nil {
		panic(err)
	}

	var listA []int
	var listB []int
	for str := range strs {
		pair := strings.Split(str, "   ")

		intA, err := strconv.Atoi(pair[0])
		if err != nil {
			panic(err)
		}
		listA = append(listA, intA)

		intB, err := strconv.Atoi(pair[1])
		if err != nil {
			panic(err)
		}
		listB = append(listB, intB)
	}

	slices.Sort(listA)
	slices.Sort(listB)

	total := 0
	for i := 0; i < len(listA); i++ {
		total += abs(listA[i] - listB[i])
	}

	fmt.Println(total)
}

func abs(num int) int {
	if num < 0 {
		num *= -1
	}

	return num
}
