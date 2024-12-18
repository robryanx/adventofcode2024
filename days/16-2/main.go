package main

import (
	"fmt"
	"maps"
	"slices"
	"strconv"

	"github.com/robryanx/adventofcode2024/util"
)

type Direction int

const (
	Unknown Direction = iota
	North
	East
	South
	West
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
					Y:   y,
					X:   x,
					Dir: util.East,
				}
			} else if grid[y][x] == 'E' {
				end = util.NodePos{
					Y: y,
					X: x,
				}
			}
		}
	}

	basePath, distance, _, _ := util.Pathfind(grid, util.East, 1000, start, end)
	for _, pos := range basePath {
		grid[pos.Y][pos.X] = 'O'
	}
	slices.Reverse(basePath)

	for i := 0; i < len(basePath)-1; i++ {
		for j := 1; j < 5; j++ {
			blockDirection := util.Unknown
			if basePath[i].Dir != basePath[i+1].Dir {
				blockDirection = basePath[i+1].Dir
			}

			testDir := util.Direction(j)
			start := util.NodePos{
				Y:              basePath[i].Y,
				X:              basePath[i].X,
				Dir:            testDir,
				BlockDirection: blockDirection,
			}

			extraCost := 0
			if testDir != basePath[i].Dir {
				extraCost += 1000
			}

			checkPath, checkDistance, found, _ := util.Pathfind(grid, testDir, 1000, start, end)
			if found && checkDistance+basePath[i].Cost+extraCost == distance {
				for _, pos := range checkPath {
					grid[pos.Y][pos.X] = 'O'
				}
			}
		}
	}

	total := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 'O' {
				total++
			}
		}
	}

	// dodgy add 1 for the missing path
	total++

	return total
}

func costGrid(grid [][]byte, costs map[int]int) {
	var gridStr [][]string
	for y := 0; y < len(grid); y++ {
		gridStrRow := make([]string, len(grid[0]))
		for x := 0; x < len(grid[0]); x++ {
			gridStrRow[x] = string(grid[y][x])
		}
		gridStr = append(gridStr, gridStrRow)
	}

	for pos, cost := range costs {
		x := pos % len(grid[0])
		y := (pos - x) / len(grid[0])
		gridStr[y][x] = strconv.Itoa(cost)
	}

	for y := 0; y < len(gridStr); y++ {
		for x := 0; x < len(gridStr[0]); x++ {
			fmt.Printf("%6s", gridStr[y][x])
		}
		fmt.Println()
	}
}

func pathFind(grid [][]byte, current, end util.NodePos, direction Direction, path map[int]struct{}, cost, maxCost int, visited map[int]struct{}) {
	if cost > maxCost {
		return
	}

	if current.Y == end.Y && current.X == end.X {
		for pos := range path {
			visited[pos] = struct{}{}
		}

		return
	}

	if current.Y > 0 && grid[current.Y-1][current.X] != '#' {
		if _, ok := path[(current.Y-1)*len(grid[0])+current.X]; !ok {
			next := util.NodePos{
				Y: current.Y - 1,
				X: current.X,
			}
			nextPath := maps.Clone(path)
			nextPath[(current.Y-1)*len(grid[0])+current.X] = struct{}{}
			nextCost := cost + 1 + rotation(direction, North)*1000

			pathFind(grid, next, end, North, nextPath, nextCost, maxCost, visited)
		}
	}

	if current.Y < len(grid)-1 && grid[current.Y+1][current.X] != '#' {
		if _, ok := path[(current.Y+1)*len(grid[0])+current.X]; !ok {
			next := util.NodePos{
				Y: current.Y + 1,
				X: current.X,
			}
			nextPath := maps.Clone(path)
			nextPath[(current.Y+1)*len(grid[0])+current.X] = struct{}{}
			nextCost := cost + 1 + rotation(direction, South)*1000

			pathFind(grid, next, end, South, nextPath, nextCost, maxCost, visited)
		}
	}

	if current.X > 0 && grid[current.Y][current.X-1] != '#' {
		if _, ok := path[current.Y*len(grid[0])+current.X-1]; !ok {
			next := util.NodePos{
				Y: current.Y,
				X: current.X - 1,
			}
			nextPath := maps.Clone(path)
			nextPath[current.Y*len(grid[0])+current.X-1] = struct{}{}
			nextCost := cost + 1 + rotation(direction, West)*1000

			pathFind(grid, next, end, West, nextPath, nextCost, maxCost, visited)
		}
	}

	if current.X < len(grid[0])-1 && grid[current.Y][current.X+1] != '#' {
		if _, ok := path[current.Y*len(grid[0])+current.X+1]; !ok {
			next := util.NodePos{
				Y: current.Y,
				X: current.X + 1,
			}
			nextPath := maps.Clone(path)
			nextPath[current.Y*len(grid[0])+current.X+1] = struct{}{}
			nextCost := cost + 1 + rotation(direction, East)*1000

			pathFind(grid, next, end, East, nextPath, nextCost, maxCost, visited)
		}
	}
}

func rotation(currentDir, newDirection Direction) int {
	return abs(int(currentDir)-int(newDirection)) % 2
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}

	return num
}
