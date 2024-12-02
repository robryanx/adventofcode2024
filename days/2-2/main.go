package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2024/util"
)

type Direction int

const (
	Increasing Direction = 1
	Decreasing Direction = 2
)

func main() {
	strs, err := util.ReadStrings("2", false, "\n")
	if err != nil {
		panic(err)
	}

	safe := 0
	for str := range strs {
		nums := make([]int, 0)
		for _, numRaw := range strings.Split(str, " ") {
			num, err := strconv.Atoi(numRaw)
			if err != nil {
				panic(err)
			}

			nums = append(nums, num)
		}

		if checkDirection(nums, true) {
			safe++
		}
	}

	fmt.Println(safe)
}

func checkDirection(nums []int, canRemove bool) bool {
	direction := Increasing
	if nums[0] > nums[len(nums)-1] {
		direction = Decreasing
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			if !canRemove {
				return false
			}

			return checkDirection(slices.Concat(nums[:i], nums[i+1:]), false)
		}

		if direction == Increasing && nums[i] > nums[i+1] ||
			direction == Decreasing && nums[i] < nums[i+1] ||
			abs(nums[i]-nums[i+1]) > 3 {

			if !canRemove {
				return false
			}

			return checkDirection(slices.Concat(nums[:i], nums[i+1:]), false) ||
				checkDirection(slices.Concat(nums[:i+1], nums[i+2:]), false)
		}
	}

	return true
}

func abs(num int) int {
	if num < 0 {
		num *= -1
	}

	return num
}
