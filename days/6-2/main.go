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

type order struct {
	pos int
	dir Direction
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

	path := [][3]int{}
	visits := make(map[[2]int][]order, 0)

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
				path = append(path, [3]int{y, x, int(currentDirection)})
				break loop
			}
		}
	}

	posCount := 0
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
			yx := [2]int{nextY, nextX}
			visits[yx] = append(visits[yx], order{
				pos: posCount,
				dir: currentDirection,
			})
			path = append(path, [3]int{nextY, nextX, int(currentDirection)})
			currentY = nextY
			currentX = nextX
			posCount++
		}
	}

	cycles := 0
	for pos, order := range visits {
		if order[0].pos == 0 {
			continue
		}
		newVisits := make(map[[3]int]struct{}, order[0].pos)
		for i := 0; i < order[0].pos-1; i++ {
			newVisits[path[i]] = struct{}{}
		}

		grid[pos[0]][pos[1]] = '#'

		currentY = path[order[0].pos-1][0]
		currentX = path[order[0].pos-1][1]
		currentDirection = Direction(path[order[0].pos-1][2])
		var nextY, nextX int
	loop2:
		for {
			switch currentDirection {
			case North:
				nextY = currentY - 1
				nextX = currentX
				if nextY < 0 {
					break loop2
				}
			case South:
				nextY = currentY + 1
				nextX = currentX
				if nextY >= len(grid) {
					break loop2
				}
			case East:
				nextX = currentX + 1
				nextY = currentY
				if nextX >= len(grid[0]) {
					break loop2
				}
			case West:
				nextX = currentX - 1
				nextY = currentY
				if nextX < 0 {
					break loop2
				}
			}

			if grid[nextY][nextX] == '#' {
				currentDirection = rotateMapping[currentDirection]
			} else {
				nextYx := [3]int{nextY, nextX, int(currentDirection)}
				if _, ok := newVisits[nextYx]; ok {
					cycles++
					break
				}
				newVisits[nextYx] = struct{}{}

				currentY = nextY
				currentX = nextX
			}
		}

		grid[pos[0]][pos[1]] = '.'
	}

	fmt.Println(cycles)
}
