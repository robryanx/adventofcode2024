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

		robots = append(robots, &robot{
			y:         y,
			x:         x,
			yVelocity: yVelocity,
			xVelocity: xVelocity,
		})
	}

	for i := 0; i < 100; i++ {
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
			robot.y = newY
			robot.x = newX
		}
	}

	var quads [4]int
	for _, robot := range robots {
		if robot.x < (xSize/2) && robot.y < (ySize/2) {
			quads[0]++
		} else if robot.x > (xSize/2) && robot.y < (ySize/2) {
			quads[1]++
		} else if robot.x < (xSize/2) && robot.y > (ySize/2) {
			quads[2]++
		} else if robot.x > (xSize/2) && robot.y > (ySize/2) {
			quads[3]++
		}
	}

	return quads[0] * quads[1] * quads[2] * quads[3]
}
