package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2024/util"
)

type space struct {
	start  int
	length int
	empty  bool
	id     uint16
	spaces []*space
}

func main() {
	fmt.Println(solution())
}

func solution() int {
	values, err := util.ReadBytes("9", false)
	if err != nil {
		panic(err)
	}

	spaces := []*space{}

	id := uint16(0)
	pos := 0
	empty := false
	for _, value := range values {

		if empty {
			empty = false

			spaces = append(spaces, &space{
				start:  pos,
				length: int(value - '0'),
				empty:  true,
			})
		} else {
			spaces = append(spaces, &space{
				start:  pos,
				length: int(value - '0'),
				empty:  false,
				id:     id,
			})

			id++
			empty = true
		}

		pos += int(value - '0')
	}

	for i := len(spaces) - 1; i > 0; i-- {
		if !spaces[i].empty {
			sp := nextSpace(spaces, spaces[i].length)
			if sp != nil {
				if sp.start > spaces[i].start {
					continue
				}

				sp.empty = false
				sp.id = spaces[i].id
				spaces[i].empty = true
			}
		}
	}

	return totalSpaces(spaces)
}

func totalSpaces(spaces []*space) int {
	total := 0
	for _, sp := range spaces {
		if len(sp.spaces) > 0 {
			total += totalSpaces(sp.spaces)
		} else {
			if !sp.empty {
				for i := 0; i < sp.length; i++ {
					total += int(sp.id) * (sp.start + i)
				}
			}
		}
	}
	return total
}

func nextSpace(spaces []*space, length int) *space {
	for _, sp := range spaces {
		if len(sp.spaces) > 0 {
			sp := nextSpace(sp.spaces, length)
			if sp != nil {
				return sp
			}
		} else {
			if sp.empty && sp.length >= length {
				remaining := sp.length - length
				if remaining > 0 {
					retSpace := &space{
						start:  sp.start,
						length: length,
						empty:  true,
					}
					sp.spaces = append(sp.spaces, retSpace)
					sp.spaces = append(sp.spaces, &space{
						start:  sp.start + length,
						length: remaining,
						empty:  true,
					})

					return retSpace
				} else {
					return sp
				}
			}
		}
	}

	return nil
}
