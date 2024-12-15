package main

import (
	"fmt"
	"slices"
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
				var buildRow []byte
				for _, pos := range row {
					if pos == '@' {
						buildRow = append(buildRow, byte(pos), '.')
					} else {
						buildRow = append(buildRow, byte(pos), byte(pos))
					}
				}
				grid = append(grid, buildRow)
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
				impacted := []map[int]struct{}{
					{
						currentX: {},
					},
				}
				impacted = push(grid, currentY-1, -1, impacted)

				if moveImpacted(grid, currentY, -1, impacted) {
					currentY--
				}
			}
		case 'v':
			if grid[currentY+1][currentX] == '.' {
				currentY = currentY + 1
			} else if grid[currentY+1][currentX] == 'O' {
				impacted := []map[int]struct{}{
					{
						currentX: {},
					},
				}
				impacted = push(grid, currentY+1, 1, impacted)

				if moveImpacted(grid, currentY, 1, impacted) {
					currentY++
				}
			}
		}
	}

	total := 0
	for y := 0; y < len(grid); y++ {
		count := 1
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 'O' {
				if count%2 != 0 {
					total += y*100 + x
				}

				count++
			}
		}
	}

	return total
}

func moveImpacted(grid [][]byte, startY, dir int, impacted []map[int]struct{}) bool {
	if dir == -1 {
		// test if we can move
		for i := 0; i < len(impacted); i++ {
			for x := range impacted[i] {
				if grid[startY-1-len(impacted)+i][x] == '#' {
					return false
				}
			}
		}

		// do the move
		for i := 0; i < len(impacted); i++ {
			for x := range impacted[i] {
				grid[startY-1-len(impacted)+i][x] = 'O'
				grid[startY-len(impacted)+i][x] = '.'
			}
		}
	} else {
		// test if we can move
		for i := 0; i < len(impacted); i++ {
			for x := range impacted[i] {
				if grid[startY+1+len(impacted)-i][x] == '#' {
					return false
				}
			}
		}

		// do the move
		for i := 0; i < len(impacted); i++ {
			for x := range impacted[i] {
				grid[startY+1+len(impacted)-i][x] = 'O'
				grid[startY+len(impacted)-i][x] = '.'
			}
		}
	}

	return true
}

func push(grid [][]byte, y, dir int, impacted []map[int]struct{}) []map[int]struct{} {
	// expand impacted on the current row
	posLookup := map[int]int{}
	count := 1
	for x := 2; x < len(grid[0])-2; x++ {
		if grid[y][x] == 'O' {
			posLookup[x] = count
			count++
		}
	}

	for x := range impacted[0] {
		if grid[y][x] == 'O' {
			if posLookup[x]%2 == 0 {
				impacted[0][x-1] = struct{}{}
			} else {
				impacted[0][x+1] = struct{}{}
			}
		}
	}

	clear := true
	nextImpacted := map[int]struct{}{}
	for x := range impacted[0] {
		if grid[y+dir][x] == 'O' {
			nextImpacted[x] = struct{}{}
			clear = false
		}
	}
	if clear {
		return impacted
	}

	impacted = slices.Concat([]map[int]struct{}{nextImpacted}, impacted)

	return push(grid, y+dir, dir, impacted)
}
