package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2024/util"
)

func main() {
	rows, err := util.ReadStrings("7", false, "\n")
	if err != nil {
		panic(err)
	}

	total := 0
	for row := range rows {
		parts := strings.Split(row, ":")
		target, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		var nums []int
		for _, numsRaw := range strings.Split(parts[1][1:], " ") {
			num, err := strconv.Atoi(numsRaw)
			if err != nil {
				panic(err)
			}

			nums = append(nums, num)
		}

		if evalLine(nums[1:], nums[0], target) {
			total += target
		}
	}

	fmt.Println(total)
}

func evalLine(nums []int, runningTotal, target int) bool {
	if len(nums) == 1 {
		return runningTotal*nums[0] == target || runningTotal+nums[0] == target
	}

	return evalLine(nums[1:], runningTotal*nums[0], target) ||
		evalLine(nums[1:], runningTotal+nums[0], target)
}
