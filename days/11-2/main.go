package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"

	"github.com/robryanx/adventofcode2024/util"
)

func main() {
	fmt.Println(solution())
}

func solution() int {
	nums, err := util.ReadBytes("11", false)
	if err != nil {
		panic(err)
	}

	line := make(map[int]int, 5000)

	for _, num := range bytes.Split(nums, []byte{' '}) {
		n, err := strconv.Atoi(string(num))
		if err != nil {
			panic(err)
		}

		line[n] = 1
	}

	for i := 0; i < 75; i++ {
		line = blinkMap(line)
	}

	return totalMap(line)
}

func totalMap(line map[int]int) int {
	total := 0
	for _, count := range line {
		total += count
	}
	return total
}

func blinkMap(line map[int]int) map[int]int {
	nextLine := make(map[int]int, 5000)

	for num, count := range line {
		d := digits(num)

		if num == 0 {
			nextLine[1] += count
		} else if d%2 == 0 {
			mask := math.Pow(10, float64(d/2))
			lower := num % int(mask)
			upper := num - lower

			nextLine[upper/int(mask)] += count
			nextLine[lower] += count
		} else {
			num *= 2024
			nextLine[num] += count
		}
	}

	return nextLine
}

func digits(num int) int {
	count := 0
	for num > 0 {
		num = num / 10
		count++
	}

	return count
}
