package main

import (
	"fmt"
	"slices"
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
			numsRaw := strings.Split(row, ",")
			nums := make([]int, 0, len(numsRaw))
			for pos, numRaw := range numsRaw {
				num, err := strconv.Atoi(numRaw)
				if err != nil {
					panic(err)
				}

				lookup[num] = pos
				nums = append(nums, num)
			}

			if nums, violation := runRules(rules, lookup, nums); violation {
				total += nums[len(nums)/2]
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

func runRules(rules map[int][]int, lookup map[int]int, nums []int) ([]int, bool) {
	violation := false
loop:
	for before, afterList := range rules {
		if posBefore, ok := lookup[before]; ok {
			for _, after := range afterList {
				if posAfter, ok := lookup[after]; ok && posBefore > posAfter {
					nums = slices.Insert(nums, posAfter, nums[posBefore])
					nums = slices.Delete(nums, posBefore+1, posBefore+2)
					lookup[before] = posAfter

					for i := posAfter; i <= posBefore; i++ {
						lookup[nums[i]] = i
					}

					violation = true
					break loop
				}
			}
		}
	}

	if violation {
		nums, _ = runRules(rules, lookup, nums)
		return nums, true
	}

	return nums, violation
}
