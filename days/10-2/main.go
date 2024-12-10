package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2024/util"
)

func main() {
	rows, err := util.ReadStrings("10", false, "\n")
	if err != nil {
		panic(err)
	}

	var grid [][]int
	for row := range rows {
		buildRow := make([]int, 0, len(row))
		for _, ch := range row {
			buildRow = append(buildRow, int(ch-'0'))
		}
		grid = append(grid, buildRow)
	}

	total := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 0 {
				total += pathsCount(grid, y, x)
			}
		}
	}

	fmt.Println(total)
}

func pathsCount(grid [][]int, y, x int) int {
	count := 0
	nextVal := grid[y][x] + 1

	if y-1 >= 0 && grid[y-1][x] == nextVal {
		if nextVal == 9 {
			count++
		} else {
			count += pathsCount(grid, y-1, x)
		}
	}
	if x-1 >= 0 && grid[y][x-1] == nextVal {
		if nextVal == 9 {
			count++
		} else {
			count += pathsCount(grid, y, x-1)
		}
	}
	if y+1 < len(grid) && grid[y+1][x] == nextVal {
		if nextVal == 9 {
			count++
		} else {
			count += pathsCount(grid, y+1, x)
		}
	}
	if x+1 < len(grid[0]) && grid[y][x+1] == nextVal {
		if nextVal == 9 {
			count++
		} else {
			count += pathsCount(grid, y, x+1)
		}
	}

	return count
}
