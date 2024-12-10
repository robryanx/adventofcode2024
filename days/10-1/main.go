package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2024/util"
)

func main() {
	fmt.Println(solution())
}

func solution() int {
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
				reached := map[int]struct{}{}
				pathsReached(grid, y, x, reached)
				total += len(reached)
			}
		}
	}

	return total
}

func pathsReached(grid [][]int, y, x int, reached map[int]struct{}) {
	nextVal := grid[y][x] + 1

	if y-1 >= 0 && grid[y-1][x] == nextVal {
		if nextVal == 9 {
			reached[((y-1)*len(grid[0]))+x] = struct{}{}
		} else {
			pathsReached(grid, y-1, x, reached)
		}
	}
	if x-1 >= 0 && grid[y][x-1] == nextVal {
		if nextVal == 9 {
			reached[(y*len(grid[0]))+x-1] = struct{}{}
		} else {
			pathsReached(grid, y, x-1, reached)
		}
	}
	if y+1 < len(grid) && grid[y+1][x] == nextVal {
		if nextVal == 9 {
			reached[((y+1)*len(grid[0]))+x] = struct{}{}
		} else {
			pathsReached(grid, y+1, x, reached)
		}
	}
	if x+1 < len(grid[0]) && grid[y][x+1] == nextVal {
		if nextVal == 9 {
			reached[(y*len(grid[0]))+x+1] = struct{}{}
		} else {
			pathsReached(grid, y, x+1, reached)
		}
	}
}
