package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2024/util"
)

func main() {
	rows, err := util.ReadStrings("5", false, "\n")
	if err != nil {
		panic(err)
	}

	rules := make(map[int][]int)
	updates := false
	total := 0
	for row := range rows {
		if row == "" {
			updates = true
			continue
		}

		if updates {
			lookup := make(map[int]int)
			nums := strings.Split(row, ",")
			middle := -1
			for pos, numRaw := range nums {
				num, err := strconv.Atoi(numRaw)
				if err != nil {
					panic(err)
				}

				lookup[num] = pos
				if pos == len(nums)/2 {
					middle = num
				}
			}

			pass := true
		loop:
			for before, afterList := range rules {
				if posBefore, ok := lookup[before]; ok {
					for _, after := range afterList {
						if posAfter, ok := lookup[after]; ok && posBefore > posAfter {
							pass = false
							break loop
						}
					}
				}
			}

			if pass {
				total += middle
			}
		} else {
			pair := strings.Split(row, "|")
			numBefore, err := strconv.Atoi(pair[0])
			if err != nil {
				panic(err)
			}

			numAfter, err := strconv.Atoi(pair[1])
			if err != nil {
				panic(err)
			}

			rules[numBefore] = append(rules[numBefore], numAfter)
		}
	}

	fmt.Println(total)
}
