package main

import (
	"fmt"
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
	listBCount := make(map[int]int, 0)
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

		if _, ok := listBCount[intB]; !ok {
			listBCount[intB] = 1
		} else {
			listBCount[intB]++
		}
	}

	total := 0
	for i := 0; i < len(listA); i++ {
		if count, ok := listBCount[listA[i]]; ok {
			total += listA[i] * count
		}
	}

	fmt.Println(total)
}
