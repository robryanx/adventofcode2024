package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/robryanx/adventofcode2024/util"
)

func main() {
	fmt.Println(solution())
}

var posRegex = regexp.MustCompile(`p=([0-9\-]+),([0-9\-]+) v=([0-9\-]+),([0-9\-]+)`)

type robot struct {
	y         int
	x         int
	yVelocity int
	xVelocity int
}

func solution() int {
	rows, err := util.ReadStrings("14", false, "\n")
	if err != nil {
		panic(err)
	}

	ySize := 103
	xSize := 101

	var grid [][]uint8
	for y := 0; y < ySize; y++ {
		grid = append(grid, make([]uint8, xSize))
	}

	robots := make([]*robot, 0, 100)
	for row := range rows {
		matches := posRegex.FindStringSubmatch(row)
		y, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}
		x, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}
		yVelocity, err := strconv.Atoi(matches[4])
		if err != nil {
			panic(err)
		}
		xVelocity, err := strconv.Atoi(matches[3])
		if err != nil {
			panic(err)
		}

		grid[y][x]++

		robots = append(robots, &robot{
			y:         y,
			x:         x,
			yVelocity: yVelocity,
			xVelocity: xVelocity,
		})
	}

	for i := 0; i < 10000; i++ {
		for _, robot := range robots {
			newY := robot.y + robot.yVelocity
			if newY >= ySize {
				newY %= ySize
			}
			if newY < 0 {
				newY = ySize + newY
			}

			newX := robot.x + robot.xVelocity
			if newX >= xSize {
				newX %= xSize
			}
			if newX < 0 {
				newX = xSize + newX
			}
			grid[robot.y][robot.x]--
			robot.y = newY
			robot.x = newX
			grid[robot.y][robot.x]++
		}

		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[0]); x++ {
				if grid[y][x] > 0 && x > xSize/3 {
					horizontalLength := 0
					for z := 0; z < xSize-x; z++ {
						if grid[y][x+z] > 0 {
							horizontalLength++
						} else {
							break
						}
					}
					if horizontalLength > 10 {
						return i + 1
					}
				}
			}
		}
	}

	return 0
}
