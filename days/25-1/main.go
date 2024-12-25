package main

import (
	"fmt"
	"strings"

	"github.com/robryanx/adventofcode2024/util"
)

func main() {
	fmt.Println(solution())
}

func solution() int {
	schematics, err := util.ReadStrings("25", false, "\n\n")
	if err != nil {
		panic(err)
	}

	var locks [][5]int
	var keys [][5]int

	for schematic := range schematics {
		lines := strings.Split(schematic, "\n")
		if lines[0][0] == '#' {
			lock := [5]int{-1, -1, -1, -1, -1}
			for y := 1; y < len(lines); y++ {
				for x := 0; x < len(lines[0]); x++ {
					if lines[y][x] == '.' && lock[x] == -1 {
						lock[x] = y - 1
					}
				}
			}
			locks = append(locks, lock)
		} else {
			key := [5]int{-1, -1, -1, -1, -1}
			for y := len(lines) - 2; y >= 0; y-- {
				for x := 0; x < len(lines[0]); x++ {
					if lines[y][x] == '.' && key[x] == -1 {
						key[x] = len(lines) - y - 2
					}
				}
			}
			keys = append(keys, key)
		}
	}

	count := 0

	for _, lock := range locks {
	loop:
		for _, key := range keys {
			for i := 0; i < 5; i++ {
				if lock[i]+key[i] > 5 {
					continue loop
				}
			}
			count++
		}
	}

	return count
}
