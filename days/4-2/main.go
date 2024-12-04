package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2024/util"
)

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
	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[0])-1; x++ {
			if grid[y][x] == 'A' && isXmas(grid, y, x) {
				total++
			}
		}
	}

	fmt.Println(total)
}

func isXmas(grid [][]byte, y, x int) bool {
	if !((grid[y-1][x-1] == 'M' && grid[y+1][x+1] == 'S') ||
		(grid[y-1][x-1] == 'S' && grid[y+1][x+1] == 'M')) {
		return false
	}

	return (grid[y-1][x+1] == 'M' && grid[y+1][x-1] == 'S') ||
		(grid[y-1][x+1] == 'S' && grid[y+1][x-1] == 'M')
}
