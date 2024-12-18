package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2024/util"
)

func main() {
	fmt.Println(solution())
}

func solution() int {
	rows, err := util.ReadStrings("18", false, "\n")
	if err != nil {
		panic(err)
	}

	gridX := 71
	gridY := 71
	corrupted := 1024

	grid := make([][]byte, 0, gridY)
	for y := 0; y < gridY; y++ {
		gridRow := make([]byte, 0, gridX)
		for x := 0; x < gridX; x++ {
			gridRow = append(gridRow, '.')
		}
		grid = append(grid, gridRow)
	}

	count := 0
	for row := range rows {
		nums := strings.Split(row, ",")
		x, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}

		grid[y][x] = '#'
		count++
		if count >= corrupted {
			break
		}
	}

	start := util.NodePos{
		Y: 0,
		X: 0,
	}
	end := util.NodePos{
		Y: 70,
		X: 70,
	}

	_, distance, _, _ := util.Pathfind(grid, util.East, 0, start, end)

	return distance
}
