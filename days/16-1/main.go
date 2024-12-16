package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2024/util"
)

func main() {
	fmt.Println(solution())
}

func solution() int {
	rows, err := util.ReadStrings("16", false, "\n")
	if err != nil {
		panic(err)
	}

	var grid [][]byte
	for row := range rows {
		grid = append(grid, []byte(row))
	}

	var start util.NodePos
	var end util.NodePos

	// find the start and end coords
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 'S' {

				start = util.NodePos{
					Y: y,
					X: x,
				}
			} else if grid[y][x] == 'E' {
				end = util.NodePos{
					Y: y,
					X: x,
				}
			}
		}
	}

	_, distance, _, _ := util.Pathfind(grid, util.East, start, end)

	return distance
}
