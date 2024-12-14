package util

import (
	"fmt"
	"os"
	"os/exec"
)

func PrintGrid(grid [][]byte) {
	for y := 0; y < len(grid); y++ {
		fmt.Printf("%s\n", string(grid[y]))
	}
}

var Reset = "\033[0m"
var Red = "\033[31m"

func PrintUint8Grid(grid [][]uint8) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] > 0 {
				fmt.Printf("%s%d%s", Red, grid[y][x], Reset)
			} else {
				fmt.Printf("%d", grid[y][x])
			}
		}
		fmt.Println()
	}
}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func CopyGrid(grid [][]byte, populate bool) [][]byte {
	nextGrid := make([][]byte, len(grid))
	for y := 0; y < len(grid); y++ {
		nextGrid[y] = make([]byte, len(grid[0]))
		if populate {
			for x := 0; x < len(grid[0]); x++ {
				nextGrid[y][x] = grid[y][x]
			}
		}
	}

	return nextGrid
}

func AdjacentMatch(grid [][]byte, y, x int, incDiagonal bool, cb func(char byte, y, x int) bool) {
	if y-1 >= 0 {
		earlyExit := cb(grid[y-1][x], y-1, x)
		if earlyExit {
			return
		}

		if incDiagonal {
			if x-1 >= 0 {
				earlyExit := cb(grid[y-1][x-1], y-1, x-1)
				if earlyExit {
					return
				}
			}

			if x+1 < len(grid[0]) {
				earlyExit := cb(grid[y-1][x+1], y-1, x+1)
				if earlyExit {
					return
				}
			}
		}
	}

	if y+1 < len(grid) {
		earlyExit := cb(grid[y+1][x], y+1, x)
		if earlyExit {
			return
		}

		if incDiagonal {
			if x-1 >= 0 {
				earlyExit := cb(grid[y+1][x-1], y+1, x-1)
				if earlyExit {
					return
				}
			}

			if x+1 < len(grid[0]) {
				earlyExit := cb(grid[y+1][x+1], y+1, x+1)
				if earlyExit {
					return
				}
			}
		}
	}

	if x-1 >= 0 {
		earlyExit := cb(grid[y][x-1], y, x-1)
		if earlyExit {
			return
		}
	}

	if x+1 < len(grid[0]) {
		_ = cb(grid[y][x+1], y, x+1)
	}
}
