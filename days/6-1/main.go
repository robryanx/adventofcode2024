package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2024/util"
)

type Direction int

const (
	North Direction = 1
	East  Direction = 2
	South Direction = 3
	West  Direction = 4
)

var charDirectionMapping = map[byte]Direction{
	'^': North,
	'>': East,
	'v': South,
	'<': West,
}

var rotateMapping = map[Direction]Direction{
	North: East,
	East:  South,
	South: West,
	West:  North,
}

var nextMapping = map[Direction][2]int{
	North: {-1, 0},
	East:  {0, 1},
	South: {1, 0},
	West:  {0, -1},
}

func main() {
	rows, err := util.ReadStrings("6", false, "\n")
	if err != nil {
		panic(err)
	}

	var grid [][]byte
	for row := range rows {
		grid = append(grid, []byte(row))
	}

	currentDirection := North
	currentY := -1
	currentX := -1

	// find the starting position and direction
loop:
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if _, ok := charDirectionMapping[grid[y][x]]; ok {
				currentDirection = charDirectionMapping[grid[y][x]]
				currentX = x
				currentY = y
				grid[y][x] = '.'
				break loop
			}
		}
	}

	positions := 1
	for {
		next := nextMapping[currentDirection]
		nextY := currentY + next[0]
		if nextY < 0 || nextY >= len(grid) {
			break
		}

		nextX := currentX + next[1]
		if nextX < 0 || nextX >= len(grid[0]) {
			break
		}

		if grid[nextY][nextX] == '#' {
			currentDirection = rotateMapping[currentDirection]
		} else {
			if grid[nextY][nextX] != 'X' {
				grid[nextY][nextX] = 'X'
				positions++
			}

			currentY = nextY
			currentX = nextX
		}
	}

	fmt.Println(positions)
}
