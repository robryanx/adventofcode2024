package main

import (
	"fmt"
	"strings"

	"github.com/robryanx/adventofcode2024/util"
)

func main() {
	fmt.Println(solution())
}

func solution() int {
	parts, err := util.ReadStrings("15", false, "\n\n")
	if err != nil {
		panic(err)
	}

	var grid [][]byte
	var moves string
	for part := range parts {
		if len(grid) == 0 {
			for _, row := range strings.Split(part, "\n") {
				grid = append(grid, []byte(row))
			}
		} else {
			moves = strings.Replace(part, "\n", "", -1)
		}
	}

	var currentY, currentX int
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == '@' {
				currentY = y
				currentX = x
				grid[y][x] = '.'
			}
		}
	}

	for _, move := range moves {
		switch move {
		case '<':
			if grid[currentY][currentX-1] == '.' {
				currentX = currentX - 1
			} else if grid[currentY][currentX-1] == 'O' {
				for x := currentX - 2; x > 0; x-- {
					if grid[currentY][x] == '#' {
						break
					} else if grid[currentY][x] == '.' {
						currentX = currentX - 1
						grid[currentY][currentX] = '.'
						grid[currentY][x] = 'O'
						break
					}
				}
			}
		case '>':
			if grid[currentY][currentX+1] == '.' {
				currentX = currentX + 1
			} else if grid[currentY][currentX+1] == 'O' {
				for x := currentX + 2; x < len(grid[0])-1; x++ {
					if grid[currentY][x] == '#' {
						break
					} else if grid[currentY][x] == '.' {
						currentX = currentX + 1
						grid[currentY][currentX] = '.'
						grid[currentY][x] = 'O'
						break
					}
				}
			}
		case '^':
			if grid[currentY-1][currentX] == '.' {
				currentY = currentY - 1
			} else if grid[currentY-1][currentX] == 'O' {
				for y := currentY - 2; y > 0; y-- {
					if grid[y][currentX] == '#' {
						break
					} else if grid[y][currentX] == '.' {
						currentY = currentY - 1
						grid[currentY][currentX] = '.'
						grid[y][currentX] = 'O'
						break
					}
				}
			}
		case 'v':
			if grid[currentY+1][currentX] == '.' {
				currentY = currentY + 1
			} else if grid[currentY+1][currentX] == 'O' {
				for y := currentY + 2; y < len(grid)-1; y++ {
					if grid[y][currentX] == '#' {
						break
					} else if grid[y][currentX] == '.' {
						currentY = currentY + 1
						grid[currentY][currentX] = '.'
						grid[y][currentX] = 'O'
						break
					}
				}
			}
		}
	}

	total := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 'O' {
				total += y*100 + x
			}
		}
	}

	return total
}
