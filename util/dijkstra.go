package util

import (
	hp "container/heap"
	"slices"
)

type path struct {
	value int
	nodes [][2]int
}

type minPath []path

func (h minPath) Len() int           { return len(h) }
func (h minPath) Less(i, j int) bool { return h[i].value < h[j].value }
func (h minPath) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minPath) Push(x interface{}) {
	*h = append(*h, x.(path))
}

func (h *minPath) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type dheap struct {
	values *minPath
}

func newHeap() *dheap {
	return &dheap{values: &minPath{}}
}

func (h *dheap) push(p path) {
	hp.Push(h.values, p)
}

func (h *dheap) pop() path {
	i := hp.Pop(h.values)
	return i.(path)
}

func getEdges(graph [][]byte, node [2]int) [][3]int {
	edges := [][3]int{}

	if node[0]-1 >= 0 && graph[node[0]-1][node[1]] != ' ' {
		edges = append(edges, [3]int{node[0] - 1, node[1], 1})
	}
	if node[0]+1 < len(graph) && graph[node[0]+1][node[1]] != ' ' {
		edges = append(edges, [3]int{node[0] + 1, node[1], 1})
	}
	if node[1]-1 >= 0 && graph[node[0]][node[1]-1] != ' ' {
		edges = append(edges, [3]int{node[0], node[1] - 1, 1})
	}
	if node[1]+1 < len(graph[0]) && graph[node[0]][node[1]+1] != ' ' {
		edges = append(edges, [3]int{node[0], node[1] + 1, 1})
	}

	return edges
}

func GetPath(grid [][]byte, origin, destination [2]int) (int, [][2]int) {
	h := newHeap()
	h.push(path{value: 0, nodes: [][2]int{origin}})
	visited := make(map[[2]int]bool)

	for len(*h.values) > 0 {
		// Find the nearest yet to visit node
		p := h.pop()
		node := p.nodes[len(p.nodes)-1]

		if visited[node] {
			continue
		}

		if node == destination {
			return p.value, p.nodes
		}

		for _, e := range getEdges(grid, node) {
			eNode := [2]int{e[0], e[1]}
			if !visited[eNode] {
				newNodes := slices.Clone(p.nodes)
				newNodes = append(newNodes, eNode)

				h.push(path{
					value: p.value + e[2],
					nodes: newNodes,
				})
			}
		}

		visited[node] = true
	}

	return 0, nil
}
