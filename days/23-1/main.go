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
	rows, err := util.ReadStrings("23", false, "\n")
	if err != nil {
		panic(err)
	}

	connections := map[string][]string{}
	for row := range rows {
		nodes := strings.Split(row, "-")
		connections[nodes[0]] = append(connections[nodes[0]], nodes[1])
		connections[nodes[1]] = append(connections[nodes[1]], nodes[0])
	}

	connectedSets := map[string]struct{}{}
	for node := range connections {
		recurseConnections(connections, node, []string{node}, 0, connectedSets)
	}

	return len(connectedSets)
}

func recurseConnections(connections map[string][]string, start string, path []string, depth int, connectedSets map[string]struct{}) {
	if depth == 3 {
		if path[len(path)-1] == start {
			hasT := false
			for _, p := range path {
				if p[0] == 't' {
					hasT = true
					break
				}
			}

			if hasT {
				path = path[:len(path)-1]
				slices.Sort(path)
				connectedSets[strings.Join(path, ",")] = struct{}{}
			}
		}
		return
	}

	for _, nextConnection := range connections[path[len(path)-1]] {
		nextPath := append(slices.Clone(path), nextConnection)

		recurseConnections(connections, start, nextPath, depth+1, connectedSets)
	}
}
