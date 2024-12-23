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

var keypadButtons = []byte{'A', '>', '^', 'v', '<'}
var expansions = map[byte]map[byte][]string{}

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
	keypadExpansions()

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

		minMoves := -1
		for _, list := range lists {
			count := expandMem(string(list), 25)
			if minMoves == -1 || count < minMoves {
				minMoves = count
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

func pathMoves(path [][2]int) string {
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
	return string(moves)
}

func keypadExpansions() {
	for _, currentCh := range keypadButtons {
		expansions[currentCh] = map[byte][]string{}

		for _, nextCh := range keypadButtons {
			paths := []pathStore{}
			dfs(dirKeypad, [][2]int{dirKeypadLookup[currentCh]}, dirKeypadLookup[nextCh], 0, &paths)

			slices.SortFunc(paths, func(a, b pathStore) int {
				return cmp.Compare(a.cost, b.cost)
			})

			minCost := paths[0].cost
			for _, path := range paths {
				if path.cost < minCost+5 {
					expansions[currentCh][nextCh] = append(expansions[currentCh][nextCh], pathMoves(path.p))
				}
			}
		}
	}
}

var cache = map[string]int{}

func expandMem(row string, depth int) int {
	if count, ok := cache[fmt.Sprintf("%s%d", row, depth)]; ok {
		return count
	}

	prev := byte('A')
	total := 0
	for _, move := range []byte(row) {
		minMoves := -1
		for _, path := range expansions[prev][move] {
			if depth == 1 {
				moves := len(path)
				if minMoves == -1 || moves < minMoves {
					minMoves = moves
				}
			} else {
				moves := expandMem(path, depth-1)
				if minMoves == -1 || moves < minMoves {
					minMoves = moves
				}
			}
		}

		prev = move
		total += minMoves
	}

	cache[fmt.Sprintf("%s%d", row, depth)] = total

	return total
}

func keypadMoves(pin []byte) [][][]byte {
	currentPos := keypadLookup['A']

	movesList := make([][][]byte, len(pin))
	for i, p := range pin {
		nextPos := keypadLookup[byte(p)]

		paths := []pathStore{}
		dfs(keypad, [][2]int{currentPos}, nextPos, 0, &paths)

		slices.SortFunc(paths, func(a, b pathStore) int {
			return cmp.Compare(a.cost, b.cost)
		})

		minPath := paths[0].cost
		for _, path := range paths {
			if path.cost != minPath {
				break
			}

			moves := []byte{}
			for i := 0; i < len(path.p)-1; i++ {
				if path.p[i][0] > path.p[i+1][0] {
					moves = append(moves, '^')
				} else if path.p[i][0] < path.p[i+1][0] {
					moves = append(moves, 'v')
				} else if path.p[i][1] > path.p[i+1][1] {
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

type pathStore struct {
	p    [][2]int
	cost int
}

func dfs(grid [][]byte, path [][2]int, destination [2]int, cost int, paths *[]pathStore) {
	current := path[len(path)-1]
	var prev *[2]int
	if len(path) > 1 {
		prev = &path[len(path)-2]
	}

	if current[0] == destination[0] && current[1] == destination[1] {
		*paths = append(*paths, pathStore{
			p:    path,
			cost: cost,
		})
		return
	}

	if current[0]-1 >= 0 && grid[current[0]-1][current[1]] != ' ' && !onPath(path, [2]int{current[0] - 1, current[1]}) {
		newPath := append(slices.Clone(path), [2]int{current[0] - 1, current[1]})
		nextCost := 9
		if prev != nil && prev[0] > current[0] {
			nextCost = 8
		}

		dfs(grid, newPath, destination, cost+nextCost, paths)
	}

	if current[0]+1 < len(grid) && grid[current[0]+1][current[1]] != ' ' && !onPath(path, [2]int{current[0] + 1, current[1]}) {
		newPath := append(slices.Clone(path), [2]int{current[0] + 1, current[1]})
		nextCost := 10
		if prev != nil && prev[0] < current[0] {
			nextCost = 8
		}

		dfs(grid, newPath, destination, cost+nextCost, paths)
	}

	if current[1]-1 >= 0 && grid[current[0]][current[1]-1] != ' ' && !onPath(path, [2]int{current[0], current[1] - 1}) {
		newPath := append(slices.Clone(path), [2]int{current[0], current[1] - 1})
		nextCost := 10
		if prev != nil && prev[1] > current[1] {
			nextCost = 8
		}

		dfs(grid, newPath, destination, cost+nextCost, paths)
	}

	if current[1]+1 < len(grid[0]) && grid[current[0]][current[1]+1] != ' ' && !onPath(path, [2]int{current[0], current[1] + 1}) {
		newPath := append(slices.Clone(path), [2]int{current[0], current[1] + 1})
		nextCost := 10
		if prev != nil && prev[1] < current[1] {
			nextCost = 8
		}

		dfs(grid, newPath, destination, cost+nextCost, paths)
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