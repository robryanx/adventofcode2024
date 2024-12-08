package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2024/util"
)

type antenna struct {
	y int
	x int
}

func main() {
	rows, err := util.ReadStrings("8", false, "\n")
	if err != nil {
		panic(err)
	}

	var grid [][]byte
	for row := range rows {
		grid = append(grid, []byte(row))
	}

	pos := map[[2]int]struct{}{}
	frequency := map[byte][]antenna{}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid); x++ {
			if grid[y][x] != '.' {
				frequency[grid[y][x]] = append(frequency[grid[y][x]], antenna{
					y: y,
					x: x,
				})
				pos[[2]int{y, x}] = struct{}{}
			}
		}
	}

	for _, antennas := range frequency {
		calc := map[[2]int]struct{}{}
		for i := 0; i < len(antennas); i++ {
			for j := 0; j < len(antennas); j++ {
				if i == j {
					continue
				}

				if _, ok := calc[[2]int{j, i}]; ok {
					continue
				}
				calc[[2]int{i, j}] = struct{}{}

				diffY := antennas[i].y - antennas[j].y
				diffX := antennas[i].x - antennas[j].x

				count := 1
				for {
					if antennas[i].y+diffY*count >= 0 &&
						antennas[i].y+diffY*count < len(grid) &&
						antennas[i].x+diffX*count >= 0 &&
						antennas[i].x+diffX*count < len(grid[0]) {
						pos[[2]int{antennas[i].y + diffY*count, antennas[i].x + diffX*count}] = struct{}{}
					} else {
						break
					}
					count++
				}

				count = 1
				for {
					if antennas[j].y-diffY*count >= 0 &&
						antennas[j].y-diffY*count < len(grid) &&
						antennas[j].x-diffX*count >= 0 &&
						antennas[j].x-diffX*count < len(grid[0]) {
						pos[[2]int{antennas[j].y - diffY*count, antennas[j].x - diffX*count}] = struct{}{}
					} else {
						break
					}
					count++
				}

			}
		}
	}

	fmt.Println(len(pos))
}
