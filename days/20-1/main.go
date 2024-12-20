package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2024/util"
)

func main() {
	fmt.Println(solution())
}

func solution() int {
	rows, err := util.ReadStrings("20", false, "\n")
	if err != nil {
		panic(err)
	}

	var grid [][]byte
	var path [][2]int
	prev := [2]int{-1, -1}
	var end [2]int
	y := 0
	for row := range rows {
		rowBuild := make([]byte, 0, len(row))
		for x, ch := range row {
			if ch == 'S' {
				ch = '.'
				path = append(path, [2]int{y, x})
			} else if ch == 'E' {
				ch = '.'
				end = [2]int{y, x}
			}
			rowBuild = append(rowBuild, byte(ch))
		}

		grid = append(grid, rowBuild)
		y++
	}

	lookup := map[[2]int]int{}
	var current [2]int
	for {
		current = path[len(path)-1]
		if current[0] == end[0] && current[1] == end[1] {
			break
		}

		var location [2]int
		if current[0]-1 >= 0 && grid[current[0]-1][current[1]] == '.' && current[0]-1 != prev[0] {
			location = [2]int{current[0] - 1, current[1]}
			path = append(path, location)
		} else if current[0]+1 < len(grid) && grid[current[0]+1][current[1]] == '.' && current[0]+1 != prev[0] {
			location = [2]int{current[0] + 1, current[1]}
			path = append(path, location)
		} else if current[1]-1 >= 0 && grid[current[0]][current[1]-1] == '.' && current[1]-1 != prev[1] {
			location = [2]int{current[0], current[1] - 1}
			path = append(path, location)
		} else if current[1]+1 < len(grid[0]) && grid[current[0]][current[1]+1] == '.' && current[1]+1 != prev[1] {
			location = [2]int{current[0], current[1] + 1}
			path = append(path, location)
		} else {
			panic("something went wrong")
		}

		lookup[location] = len(path) - 1
		prev = current
	}

	total := 0
	for i, pos := range path {
		if pos[0]-2 >= 0 && grid[pos[0]-1][pos[1]] == '#' && grid[pos[0]-2][pos[1]] == '.' {
			length := i + 2 + len(path) - lookup[[2]int{pos[0] - 2, pos[1]}]
			if len(path)-length > 99 {
				total++
			}
		}
		if pos[0]+2 < len(grid) && grid[pos[0]+1][pos[1]] == '#' && grid[pos[0]+2][pos[1]] == '.' {
			length := i + 2 + len(path) - lookup[[2]int{pos[0] + 2, pos[1]}]
			if len(path)-length > 99 {
				total++
			}
		}
		if pos[1]-2 >= 0 && grid[pos[0]][pos[1]-1] == '#' && grid[pos[0]][pos[1]-2] == '.' {
			length := i + 2 + len(path) - lookup[[2]int{pos[0], pos[1] - 2}]
			if len(path)-length > 99 {
				total++
			}
		}
		if pos[1]+2 < len(grid[0]) && grid[pos[0]][pos[1]+1] == '#' && grid[pos[0]][pos[1]+2] == '.' {
			length := i + 2 + len(path) - lookup[[2]int{pos[0], pos[1] + 2}]
			if len(path)-length > 99 {
				total++
			}
		}
	}

	return total
}
