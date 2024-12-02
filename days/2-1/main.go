package main

import (
	"fmt"
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
loop:
	for str := range strs {
		direction := Increasing
		nums := make([]int, 0)
		for _, numRaw := range strings.Split(str, " ") {
			num, err := strconv.Atoi(numRaw)
			if err != nil {
				panic(err)
			}

			nums = append(nums, num)
		}

		if nums[0] == nums[1] {
			continue
		} else if nums[0] > nums[1] {
			direction = Decreasing
		}

		for i := 0; i < len(nums)-1; i++ {
			if nums[i] == nums[i+1] {
				continue loop
			} else if direction == Increasing && nums[i] > nums[i+1] {
				continue loop
			} else if direction == Decreasing && nums[i] < nums[i+1] {
				continue loop
			} else if abs(nums[i]-nums[i+1]) > 3 {
				continue loop
			}
		}

		safe++
	}

	fmt.Println(safe)
}

func abs(num int) int {
	if num < 0 {
		num *= -1
	}

	return num
}
