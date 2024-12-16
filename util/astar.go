package util

import (
	"container/heap"
)

type priorityQueue []*node

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].rank < pq[j].rank
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	no := x.(*node)
	no.index = n
	*pq = append(*pq, no)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	no := old[n-1]
	no.index = -1
	*pq = old[0 : n-1]
	return no
}

type NodePos struct {
	Y              int
	X              int
	Dir            Direction
	BlockDirection Direction
	Cost           int
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}

	return num
}

// node is a wrapper to store A* data for a Pather node.
type node struct {
	pos    NodePos
	rank   int
	parent *node
	open   bool
	closed bool
	index  int
}

func (p node) parentDirection() Direction {
	if p.parent == nil {
		return Unknown
	}

	if p.pos.X == p.parent.pos.X && p.pos.Y < p.parent.pos.Y {
		return North
	}

	if p.pos.X == p.parent.pos.X && p.pos.Y > p.parent.pos.Y {
		return South
	}

	if p.pos.Y == p.parent.pos.Y && p.pos.X > p.parent.pos.X {
		return East
	}

	if p.pos.Y == p.parent.pos.Y && p.pos.X < p.parent.pos.X {
		return West
	}

	return Unknown
}

func rotation(currentDir, newDirection Direction) int {
	return abs(int(currentDir)-int(newDirection)) % 2
}

func (p node) neighbors(grid [][]byte, initialDirection Direction) []NodePos {
	list := make([]NodePos, 0, 4)

	parentDirection := p.parentDirection()
	if parentDirection == Unknown {
		parentDirection = initialDirection
	}

	if p.pos.Y > 0 && grid[p.pos.Y-1][p.pos.X] != '#' && p.pos.BlockDirection != North {
		list = append(list, NodePos{
			Y:    p.pos.Y - 1,
			X:    p.pos.X,
			Dir:  North,
			Cost: 1 + rotation(parentDirection, North)*1000,
		})
	}

	if p.pos.Y < len(grid)-1 && grid[p.pos.Y+1][p.pos.X] != '#' && p.pos.BlockDirection != South {
		list = append(list, NodePos{
			Y:    p.pos.Y + 1,
			X:    p.pos.X,
			Dir:  South,
			Cost: 1 + rotation(parentDirection, South)*1000,
		})
	}

	if p.pos.X > 0 && grid[p.pos.Y][p.pos.X-1] != '#' && p.pos.BlockDirection != West {
		list = append(list, NodePos{
			Y:    p.pos.Y,
			X:    p.pos.X - 1,
			Dir:  West,
			Cost: 1 + rotation(parentDirection, West)*1000,
		})
	}

	if p.pos.X < len(grid[0])-1 && grid[p.pos.Y][p.pos.X+1] != '#' && p.pos.BlockDirection != East {
		list = append(list, NodePos{
			Y:    p.pos.Y,
			X:    p.pos.X + 1,
			Dir:  East,
			Cost: 1 + rotation(parentDirection, East)*1000,
		})
	}

	return list
}

func Pathfind(grid [][]byte, initalDirection Direction, from NodePos, to NodePos) ([]NodePos, int, bool, map[int]int) {
	nm := nodeMap{}
	nq := &priorityQueue{}
	heap.Init(nq)

	fromNode := nm.get(from)
	fromNode.open = true
	heap.Push(nq, fromNode)
	for {
		if nq.Len() == 0 {
			// There's no path, return found false.
			return nil, 0, false, nil
		}
		current := heap.Pop(nq).(*node)
		current.open = false
		current.closed = true

		if current == nm.get(to) {
			// Found a path to the goal.
			p := []NodePos{}
			curr := current
			for curr != nil {
				p = append(p, curr.pos)
				curr = curr.parent
			}

			costs := make(map[int]int)
			for _, n := range nm {
				costs[n.pos.Y*len(grid[0])+n.pos.X] = n.pos.Cost
			}

			return p, current.pos.Cost, true, costs
		}

		for _, neighbor := range current.neighbors(grid, initalDirection) {
			cost := current.pos.Cost + neighbor.Cost
			neighborNode := nm.get(neighbor)

			if cost <= neighborNode.pos.Cost {
				if neighborNode.open {
					heap.Remove(nq, neighborNode.index)
				}
				neighborNode.open = false
				neighborNode.closed = false
			}
			if !neighborNode.open && !neighborNode.closed {
				neighborNode.pos.Cost = cost
				neighborNode.open = true
				neighborNode.rank = cost
				neighborNode.parent = current

				heap.Push(nq, neighborNode)
			}
		}
	}
}

type Direction int

const (
	Unknown Direction = iota
	North
	East
	South
	West
)

type nodeMap map[[2]int]*node

func (nm nodeMap) get(pos NodePos) *node {
	key := [2]int{pos.Y, pos.X}
	n, ok := nm[key]
	if !ok {
		n = &node{
			pos: pos,
		}
		nm[key] = n
	}
	return n
}
