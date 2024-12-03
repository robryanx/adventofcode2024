package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/robryanx/adventofcode2024/util"
)

var mulRegex = regexp.MustCompile(`(mul\(([0-9]{1,3}),([0-9]{1,3})\)|do\(\)|don't\(\))`)

var do = []byte("do()")
var dont = []byte("don't()")

func main() {
	bytes, err := util.ReadBytes("3", false)
	if err != nil {
		panic(err)
	}

	accept := true
	total := 0
	for {
		nextMatch := mulRegex.FindSubmatchIndex(bytes)
		if nextMatch == nil {
			break
		}

		if cmp(bytes[nextMatch[0]:nextMatch[1]], do) {
			accept = true
		} else if cmp(bytes[nextMatch[0]:nextMatch[1]], dont) {
			accept = false
		} else if accept {
			numA, err := strconv.Atoi(string(bytes[nextMatch[4]:nextMatch[5]]))
			if err != nil {
				panic(err)
			}

			numB, err := strconv.Atoi(string(bytes[nextMatch[6]:nextMatch[7]]))
			if err != nil {
				panic(err)
			}

			total += numA * numB
		}

		bytes = bytes[nextMatch[1]:]
	}

	fmt.Println(total)
}

func cmp(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
