package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2024/util"
)

type direction struct {
	x int
	y int
}

var directions = []direction{
	{
		x: 1,
		y: 0,
	},
	{
		x: 1,
		y: 1,
	},
	{
		x: 0,
		y: 1,
	},
	{
		x: -1,
		y: 1,
	},
	{
		x: -1,
		y: 0,
	},
	{
		x: -1,
		y: -1,
	},
	{
		x: 0,
		y: -1,
	},
	{
		x: 1,
		y: -1,
	},
}

func main() {
	rows, err := util.ReadStrings("4", false, "\n")
	if err != nil {
		panic(err)
	}

	var grid [][]byte
	for row := range rows {
		grid = append(grid, []byte(row))
	}

	total := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 'X' {
				total += isXmas(grid, y, x)
			}
		}
	}

	fmt.Println(total)
}

var xmas = []byte("XMAS")

func isXmas(grid [][]byte, y, x int) int {
	count := 0
loop:
	for _, direction := range directions {
		testY := y
		testX := x
		for pos := 1; pos < 4; pos++ {
			testY += direction.y
			if testY < 0 || testY > len(grid)-1 {
				continue loop
			}

			testX += direction.x
			if testX < 0 || testX > len(grid[y])-1 {
				continue loop
			}

			if grid[testY][testX] != xmas[pos] {
				continue loop
			}
		}

		count++
	}

	return count
}
