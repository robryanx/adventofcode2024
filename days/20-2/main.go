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
	for startLength, startPos := range path {
		counts := distances(grid, lookup, startLength, len(path), startPos)
		total += counts
	}

	return total
}

func distances(grid [][]byte, lookup map[[2]int]int, startLength, pathLength int, current [2]int) int {
	savedCount := 0
	radius := 20
	for yOffset := -radius; yOffset <= radius; yOffset++ {
		for xOffset := -radius; xOffset <= radius; xOffset++ {
			cheatLength := abs(yOffset) + abs(xOffset)
			if cheatLength <= radius {
				if ((yOffset <= 0 && current[0]+yOffset >= 0) || (yOffset > 0 && current[0]+yOffset < len(grid))) &&
					((xOffset <= 0 && current[1]+xOffset >= 0) || (xOffset > 0 && current[1]+xOffset < len(grid[0]))) &&
					grid[current[0]+yOffset][current[1]+xOffset] == '.' {

					remaining := pathLength - lookup[[2]int{current[0] + yOffset, current[1] + xOffset}]
					saved := pathLength - (startLength + cheatLength) - remaining

					if saved >= 100 {
						savedCount++
					}
				}
			}
		}
	}

	return savedCount
}

func abs(num int) int {
	if num < 0 {
		num *= -1
	}

	return num
}
