package main

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/robryanx/adventofcode2024/util"
)

func main() {
	fmt.Println(solution())
}

type Direction int

const (
	North Direction = 0
	East  Direction = 1
	South Direction = 2
	West  Direction = 3
)

type side struct {
	y int
	x int
}

func solution() int {
	rows, err := util.ReadStrings("12", false, "\n")
	if err != nil {
		panic(err)
	}

	var grid [][]byte
	for row := range rows {
		grid = append(grid, []byte(row))
	}

	total := 0
	pos := map[[2]int]struct{}{}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if _, ok := pos[[2]int{y, x}]; !ok {
				pos[[2]int{y, x}] = struct{}{}
				shape := [][2]int{{y, x}}
				fillShape(grid, y, x, pos, &shape)

				ch := grid[y][x]
				sides := [4][]side{}
				for _, pos := range shape {
					if pos[0] == 0 || grid[pos[0]-1][pos[1]] != ch {
						sides[North] = append(sides[North], side{
							y: pos[0],
							x: pos[1],
						})
					}
					if pos[0] == len(grid)-1 || grid[pos[0]+1][pos[1]] != ch {
						sides[South] = append(sides[South], side{
							y: pos[0],
							x: pos[1],
						})
					}
					if pos[1] == 0 || grid[pos[0]][pos[1]-1] != ch {
						sides[West] = append(sides[West], side{
							y: pos[0],
							x: pos[1],
						})
					}
					if pos[1] == len(grid[0])-1 || grid[pos[0]][pos[1]+1] != ch {
						sides[East] = append(sides[East], side{
							y: pos[0],
							x: pos[1],
						})
					}
				}

				newSides := [4][]side{}
				sideCount := 0
				for i := 0; i < 4; i++ {
					slices.SortFunc(sides[i], func(a, b side) int {
						if i == int(North) || i == int(South) {
							if a.y == b.y {
								return cmp.Compare(a.x, b.x)
							}
							return cmp.Compare(a.y, b.y)
						}

						if a.x == b.x {
							return cmp.Compare(a.y, b.y)
						}
						return cmp.Compare(a.x, b.x)
					})

					prev := sides[i][0]
					for j := 1; j < len(sides[i]); j++ {
						if i == int(North) || i == int(South) {
							if !(prev.y == sides[i][j].y && prev.x+1 == sides[i][j].x) {
								newSides[i] = append(newSides[i], prev)
							}
						} else {
							if !(prev.x == sides[i][j].x && prev.y+1 == sides[i][j].y) {
								newSides[i] = append(newSides[i], prev)
							}
						}

						prev = sides[i][j]
					}
					newSides[i] = append(newSides[i], prev)
					sideCount += len(newSides[i])
				}

				total += sideCount * len(shape)
			}
		}
	}

	return total
}

func fillShape(grid [][]byte, y, x int, pos map[[2]int]struct{}, shape *[][2]int) {
	ch := grid[y][x]

	if y-1 >= 0 && grid[y-1][x] == ch {
		testPos := [2]int{y - 1, x}
		if _, ok := pos[testPos]; !ok {
			pos[testPos] = struct{}{}
			*shape = append(*shape, testPos)
			fillShape(grid, y-1, x, pos, shape)
		}
	}
	if y+1 < len(grid) && grid[y+1][x] == ch {
		testPos := [2]int{y + 1, x}
		if _, ok := pos[testPos]; !ok {
			pos[testPos] = struct{}{}
			*shape = append(*shape, testPos)
			fillShape(grid, y+1, x, pos, shape)
		}
	}
	if x-1 >= 0 && grid[y][x-1] == ch {
		testPos := [2]int{y, x - 1}
		if _, ok := pos[testPos]; !ok {
			pos[testPos] = struct{}{}
			*shape = append(*shape, testPos)
			fillShape(grid, y, x-1, pos, shape)
		}
	}
	if x+1 < len(grid[0]) && grid[y][x+1] == ch {
		testPos := [2]int{y, x + 1}
		if _, ok := pos[testPos]; !ok {
			pos[testPos] = struct{}{}
			*shape = append(*shape, testPos)
			fillShape(grid, y, x+1, pos, shape)
		}
	}
}
