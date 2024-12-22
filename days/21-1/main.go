package main

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2024/util"
)

func main() {
	fmt.Println(solution())
}

var keypad = [][]byte{
	{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{' ', '0', 'A'},
}

var dirKeypad = [][]byte{
	{' ', '^', 'A'},
	{'<', 'v', '>'},
}

var keypadLookup = map[byte][2]int{
	'7': {0, 0},
	'8': {0, 1},
	'9': {0, 2},
	'4': {1, 0},
	'5': {1, 1},
	'6': {1, 2},
	'1': {2, 0},
	'2': {2, 1},
	'3': {2, 2},
	'0': {3, 1},
	'A': {3, 2},
}

var dirKeypadLookup = map[byte][2]int{
	'^': {0, 1},
	'A': {0, 2},
	'<': {1, 0},
	'v': {1, 1},
	'>': {1, 2},
}

func solution() int {
	rows, err := util.ReadStrings("21", false, "\n")
	if err != nil {
		panic(err)
	}

	total := 0
	for row := range rows {
		segmentMoves := keypadMoves([]byte(row))

		lists := segmentMoves[0]
		for i := 1; i < len(segmentMoves); i++ {
			newLists := [][]byte{}
			for j := 0; j < len(segmentMoves[i]); j++ {
				for _, list := range lists {
					newLists = append(newLists, slices.Concat(list, segmentMoves[i][j]))
				}
			}
			lists = newLists
		}

		minMoves := 100000
		for _, list := range lists {
			segmentMoves = dirKeypadMoves(list)

			nextLists := segmentMoves[0]
			for i := 1; i < len(segmentMoves); i++ {
				newLists := [][]byte{}
				for j := 0; j < len(segmentMoves[i]); j++ {
					for _, list := range nextLists {
						newLists = append(newLists, slices.Concat(list, segmentMoves[i][j]))
					}
				}
				nextLists = newLists
			}

			for _, nextList := range nextLists {
				moves := dirKeypadMovesTotal(nextList)
				if moves <= minMoves {
					minMoves = moves
				}
			}
		}

		numRaw := strings.Replace(row, "A", "", -1)
		num, err := strconv.Atoi(numRaw)
		if err != nil {
			panic(err)
		}

		total += minMoves * num
	}

	return total
}

func keypadMoves(pin []byte) [][][]byte {
	currentPos := keypadLookup['A']

	movesList := make([][][]byte, len(pin))
	for i, p := range pin {
		nextPos := keypadLookup[byte(p)]

		paths := [][][2]int{}
		dfs(keypad, [][2]int{currentPos}, nextPos, &paths)

		slices.SortFunc(paths, func(a, b [][2]int) int {
			return cmp.Compare(len(a), len(b))
		})

		minPath := len(paths[0])
		for _, path := range paths {
			if len(path) != minPath {
				break
			}

			moves := []byte{}
			for i := 0; i < len(path)-1; i++ {
				if path[i][0] > path[i+1][0] {
					moves = append(moves, '^')
				} else if path[i][0] < path[i+1][0] {
					moves = append(moves, 'v')
				} else if path[i][1] > path[i+1][1] {
					moves = append(moves, '<')
				} else {
					moves = append(moves, '>')
				}
			}
			moves = append(moves, 'A')
			movesList[i] = append(movesList[i], moves)
		}

		currentPos = nextPos
	}

	return movesList
}

func dirKeypadMovesTotal(prevMoves []byte) int {
	total := 0
	currentPos := dirKeypadLookup['A']
	for _, p := range prevMoves {
		nextPos := dirKeypadLookup[p]

		distance, _ := util.GetPath(dirKeypad, currentPos, nextPos)
		total += distance + 1

		currentPos = nextPos
	}

	return total
}

func dirKeypadMoves(prevMoves []byte) [][][]byte {
	currentPos := dirKeypadLookup['A']
	movesList := make([][][]byte, len(prevMoves))
	for i, p := range prevMoves {
		nextPos := dirKeypadLookup[p]

		paths := [][][2]int{}
		dfs(dirKeypad, [][2]int{currentPos}, nextPos, &paths)

		slices.SortFunc(paths, func(a, b [][2]int) int {
			return cmp.Compare(len(a), len(b))
		})

		minPath := len(paths[0])
		for _, path := range paths {
			if len(path) != minPath {
				break
			}

			moves := []byte{}
			for i := 0; i < len(path)-1; i++ {
				if path[i][0] > path[i+1][0] {
					moves = append(moves, '^')
				} else if path[i][0] < path[i+1][0] {
					moves = append(moves, 'v')
				} else if path[i][1] > path[i+1][1] {
					moves = append(moves, '<')
				} else {
					moves = append(moves, '>')
				}
			}
			moves = append(moves, 'A')
			movesList[i] = append(movesList[i], moves)
		}

		currentPos = nextPos
	}

	return movesList
}

func dfs(grid [][]byte, path [][2]int, destination [2]int, paths *[][][2]int) {
	current := path[len(path)-1]

	if current[0] == destination[0] && current[1] == destination[1] {
		*paths = append(*paths, path)
		return
	}

	if current[0]-1 >= 0 && grid[current[0]-1][current[1]] != ' ' && !onPath(path, [2]int{current[0] - 1, current[1]}) {
		newPath := append(slices.Clone(path), [2]int{current[0] - 1, current[1]})

		dfs(grid, newPath, destination, paths)
	}

	if current[0]+1 < len(grid) && grid[current[0]+1][current[1]] != ' ' && !onPath(path, [2]int{current[0] + 1, current[1]}) {
		newPath := append(slices.Clone(path), [2]int{current[0] + 1, current[1]})

		dfs(grid, newPath, destination, paths)
	}

	if current[1]-1 >= 0 && grid[current[0]][current[1]-1] != ' ' && !onPath(path, [2]int{current[0], current[1] - 1}) {
		newPath := append(slices.Clone(path), [2]int{current[0], current[1] - 1})

		dfs(grid, newPath, destination, paths)
	}

	if current[1]+1 < len(grid[0]) && grid[current[0]][current[1]+1] != ' ' && !onPath(path, [2]int{current[0], current[1] + 1}) {
		newPath := append(slices.Clone(path), [2]int{current[0], current[1] + 1})

		dfs(grid, newPath, destination, paths)
	}
}

func onPath(path [][2]int, pos [2]int) bool {
	for _, p := range path {
		if p[0] == pos[0] && p[1] == pos[1] {
			return true
		}
	}
	return false
}
