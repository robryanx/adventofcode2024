package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2024/util"
)

func main() {
	values, err := util.ReadBytes("9", false)
	if err != nil {
		panic(err)
	}

	length := 0
	for _, value := range values {
		length += int(value - '0')
	}

	line := make([]uint16, length)
	id := uint16(1)
	pos := 0
	space := false
	spacePos := 0
	for _, value := range values {
		if space {
			pos += int(value - '0')
			space = false
		} else {
			val := int(value - '0')
			for i := pos; i < pos+val; i++ {
				line[i] = uint16(id)
			}
			id++
			pos += val
			space = true
		}
	}

	idPos := len(line) - 1
	for idPos > spacePos {
		for line[idPos] == 0 {
			idPos--
		}
		for line[spacePos] != 0 {
			spacePos++
		}
		line[spacePos] = line[idPos]
		line[idPos] = 0
		idPos--
		spacePos++
	}

	total := 0
	for i := 0; i < idPos+1; i++ {
		total += (int(line[i]) - 1) * i
	}

	fmt.Println(total)
}
