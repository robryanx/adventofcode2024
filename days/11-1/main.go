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
	num  int
	next *wrappedNum
}

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
				num: n,
			}
			last = start
		} else {
			last.next = &wrappedNum{
				num: n,
			}
			last = last.next
		}
	}

	for i := 0; i < 25; i++ {
		blink(start)
	}

	return countNums(start)
}

func blink(start *wrappedNum) {
	current := start

	for current != nil {
		if current.num == 0 {
			current.num = 1
		} else if d := digits(current.num); d%2 == 0 {
			mask := math.Pow(10, float64(d/2))
			lower := current.num % int(mask)
			upper := current.num - lower

			next := current.next
			current.num = upper / int(mask)
			current.next = &wrappedNum{
				num:  lower,
				next: next,
			}
			current = current.next
		} else {
			current.num *= 2024
		}
		current = current.next
	}
}

func printNums(start *wrappedNum) {
	current := start
	for current != nil {
		fmt.Println(current.num)
		current = current.next
	}
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
