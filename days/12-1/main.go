package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2024/util"
)

func main() {
	fmt.Println(solution())
}

func solution() int {
	rows, err := util.ReadStrings("12", false, "\n")
	if err != nil {
		panic(err)
	}

	var grid [][]byte
	for row := range rows {
		grid = append(grid, []byte(row))
	}

	total := 0
	pos := map[[2]int]struct{}{}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if _, ok := pos[[2]int{y, x}]; !ok {
				pos[[2]int{y, x}] = struct{}{}
				shape := [][2]int{{y, x}}
				fillShape(grid, y, x, pos, &shape)

				ch := grid[y][x]
				sides := 0
				for _, pos := range shape {
					if pos[0] == 0 || grid[pos[0]-1][pos[1]] != ch {
						sides++
					}
					if pos[0] == len(grid)-1 || grid[pos[0]+1][pos[1]] != ch {
						sides++
					}
					if pos[1] == 0 || grid[pos[0]][pos[1]-1] != ch {
						sides++
					}
					if pos[1] == len(grid[0])-1 || grid[pos[0]][pos[1]+1] != ch {
						sides++
					}
				}

				total += len(shape) * sides
			}
		}
	}

	return total
}

func fillShape(grid [][]byte, y, x int, pos map[[2]int]struct{}, shape *[][2]int) {
	ch := grid[y][x]

	if y-1 >= 0 && grid[y-1][x] == ch {
		testPos := [2]int{y - 1, x}
		if _, ok := pos[testPos]; !ok {
			pos[testPos] = struct{}{}
			*shape = append(*shape, testPos)
			fillShape(grid, y-1, x, pos, shape)
		}
	}
	if y+1 < len(grid) && grid[y+1][x] == ch {
		testPos := [2]int{y + 1, x}
		if _, ok := pos[testPos]; !ok {
			pos[testPos] = struct{}{}
			*shape = append(*shape, testPos)
			fillShape(grid, y+1, x, pos, shape)
		}
	}
	if x-1 >= 0 && grid[y][x-1] == ch {
		testPos := [2]int{y, x - 1}
		if _, ok := pos[testPos]; !ok {
			pos[testPos] = struct{}{}
			*shape = append(*shape, testPos)
			fillShape(grid, y, x-1, pos, shape)
		}
	}
	if x+1 < len(grid[0]) && grid[y][x+1] == ch {
		testPos := [2]int{y, x + 1}
		if _, ok := pos[testPos]; !ok {
			pos[testPos] = struct{}{}
			*shape = append(*shape, testPos)
			fillShape(grid, y, x+1, pos, shape)
		}
	}
}
