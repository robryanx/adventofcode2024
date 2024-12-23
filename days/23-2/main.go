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

func solution() string {
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

	largestLen := 0
	largest := []string{}
	for node, nextConnections := range connections {
		nextSet := commonSet(connections, nextConnections, 0)
		if largestLen < len(nextSet) {
			largestLen = len(nextSet)
			nextSet = append(nextSet, node)
			largest = nextSet
		}
	}

	slices.Sort(largest)

	return strings.Join(largest, ",")
}

func commonSet(connections map[string][]string, currentSet []string, pos int) []string {
	if pos >= len(currentSet) {
		return currentSet
	}

	nextSet := slices.Clone(currentSet)
	nextSet = slices.Delete(nextSet, pos, pos+1)
	retSet := commonSet(connections, nextSet, pos)

	nextSet = slices.Clone(currentSet)
	checkConnections := connections[nextSet[pos]]
	for i := len(nextSet) - 1; i >= 0; i-- {
		if nextSet[pos] != nextSet[i] && !slices.Contains(checkConnections, nextSet[i]) {
			nextSet = slices.Delete(nextSet, i, i+1)
		}
	}

	altRetSet := commonSet(connections, nextSet, pos+1)
	if len(altRetSet) > len(retSet) {
		return altRetSet
	}

	return retSet
}
