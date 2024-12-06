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

	visits := make(map[[2]int][]Direction, 0)

	startingDirection := North
	startingY := -1
	startingX := -1

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
				startingDirection = charDirectionMapping[grid[y][x]]
				startingX = x
				startingY = y
				grid[y][x] = '.'
				break loop
			}
		}
	}

	baseGrid := util.CopyGrid(grid, true)

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
			visits[[2]int{nextY, nextX}] = append(visits[[2]int{nextY, nextX}], currentDirection)
			currentY = nextY
			currentX = nextX
		}
	}

	cycles := 0
	for pos := range visits {
		newVisits := map[[2]int][]Direction{
			{startingY, startingX}: {startingDirection},
		}

		testGrid := util.CopyGrid(baseGrid, true)
		testGrid[pos[0]][pos[1]] = 'O'
		currentDirection = startingDirection
		currentY = startingY
		currentX = startingX
		start := false
	loop2:
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

			if testGrid[nextY][nextX] == '#' || testGrid[nextY][nextX] == 'O' {
				currentDirection = rotateMapping[currentDirection]

				if nextY == pos[0] && nextX == pos[1] {
					start = true
				}
			} else {
				if start {
					if dirs, ok := newVisits[[2]int{nextY, nextX}]; ok {
						for _, dir := range dirs {
							if dir == currentDirection {
								cycles++
								break loop2
							}
						}
					}
				}

				currentY = nextY
				currentX = nextX

				newVisits[[2]int{currentY, currentX}] = append(newVisits[[2]int{currentY, currentX}], currentDirection)
			}
		}
	}

	fmt.Println(cycles)
}
