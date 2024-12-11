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

type wrappedNum struct {
	num    int
	blinks int
	next   *wrappedNum
}

var baseExpansions = map[[2]int]int{}

func solution() int {
	nums, err := util.ReadBytes("11", false)
	if err != nil {
		panic(err)
	}

	var start *wrappedNum
	var last *wrappedNum

	for _, num := range bytes.Split(nums, []byte{' '}) {
		n, err := strconv.Atoi(string(num))
		if err != nil {
			panic(err)
		}

		if start == nil {
			start = &wrappedNum{
				num:    n,
				blinks: 0,
			}
			last = start
		} else {
			last.next = &wrappedNum{
				num: n,
			}
			last = last.next
		}
	}

	getBaseExpansions()

	for i := 0; i < 75; i++ {
		blink(start, true)
	}

	return countWithExpansions(start)
}

func getBaseExpansions() {
	for i := 0; i < 10; i++ {
		start := &wrappedNum{
			num: i,
		}
		for j := 0; j < 40; j++ {
			blink(start, false)
			baseExpansions[[2]int{i, j + 1}] = countNums(start)
		}
	}
}

func countWithExpansions(start *wrappedNum) int {
	current := start
	count := 0
	for current != nil {
		if current.blinks == 75 {
			count++
		} else {
			count += baseExpansions[[2]int{current.num, 75 - current.blinks}]
		}

		current = current.next
	}

	return count
}

func blink(start *wrappedNum, earlyExit bool) {
	current := start

	for current != nil {
		d := digits(current.num)
		if d == 1 && current.blinks > 35 && earlyExit {
			current = current.next
			continue
		}

		if current.num == 0 {
			current.blinks++
			current.num = 1
		} else if d%2 == 0 {
			mask := math.Pow(10, float64(d/2))
			lower := current.num % int(mask)
			upper := current.num - lower

			next := current.next
			current.num = upper / int(mask)
			current.blinks++
			current.next = &wrappedNum{
				num:    lower,
				blinks: current.blinks,
				next:   next,
			}
			current = current.next
		} else {
			current.blinks++
			current.num *= 2024
		}
		current = current.next
	}
}

func printNums(start *wrappedNum) {
	current := start
	for current != nil {
		fmt.Printf("%d ", current.num)
		current = current.next
	}
	fmt.Println()
}

func countNums(start *wrappedNum) int {
	current := start
	count := 0
	for current != nil {
		count++
		current = current.next
	}

	return count
}

func digits(num int) int {
	count := 0
	for num > 0 {
		num = num / 10
		count++
	}

	return count
}
