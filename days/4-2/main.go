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
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 'A' && isXmas(grid, y, x) {
				total++
			}
		}
	}

	fmt.Println(total)
}

func isXmas(grid [][]byte, y, x int) bool {
	if y-1 < 0 || y+1 > len(grid)-1 {
		return false
	}
	if x-1 < 0 || x+1 > len(grid[0])-1 {
		return false
	}
	if !((grid[y-1][x-1] == 'M' && grid[y+1][x+1] == 'S') ||
		(grid[y-1][x-1] == 'S' && grid[y+1][x+1] == 'M')) {
		return false
	}

	return (grid[y-1][x+1] == 'M' && grid[y+1][x-1] == 'S') ||
		(grid[y-1][x+1] == 'S' && grid[y+1][x-1] == 'M')
}
