package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/robryanx/adventofcode2024/util"
)

func main() {
	fmt.Println(solution())
}

var registerRegex = regexp.MustCompile(`Register [A-Z]{1}: ([0-9]+)`)

func solution() string {
	rows, err := util.ReadStrings("17", false, "\n\n")
	if err != nil {
		panic(err)
	}

	var registers []int
	var instructions [][2]int
	var instructionPointer int
	for row := range rows {
		if len(registers) == 0 {
			for _, rRow := range strings.Split(row, "\n") {
				matches := registerRegex.FindStringSubmatch(rRow)

				val, err := strconv.Atoi(matches[1])
				if err != nil {
					panic(err)
				}

				registers = append(registers, val)
			}
		} else {
			var prev int
			for i, instruction := range strings.Split(row[strings.Index(row, " ")+1:], ",") {
				instructionVal, err := strconv.Atoi(instruction)
				if err != nil {
					panic(err)
				}
				if (i+1)%2 == 1 {
					prev = instructionVal
				} else {
					instructions = append(instructions, [2]int{prev, instructionVal})
				}
			}
		}
	}

	var output []string
	for {
		if instructionPointer > len(instructions)-1 {
			break
		}

		instructon := instructions[instructionPointer]
		switch instructon[0] {
		case 0:
			registers[0] = registers[0] / int(math.Pow(2, float64(comboVar(registers, instructon[1]))))
			instructionPointer++
		case 1:
			registers[1] = registers[1] ^ instructon[1]
			instructionPointer++
		case 2:
			registers[1] = comboVar(registers, instructon[1]) % 8
			instructionPointer++
		case 3:
			if registers[0] != 0 {
				instructionPointer = instructon[1] / 2
			} else {
				instructionPointer++
			}
		case 4:
			registers[1] = registers[1] ^ registers[2]
			instructionPointer++
		case 5:
			outVal := comboVar(registers, instructon[1]) % 8
			output = append(output, strconv.Itoa(outVal))
			instructionPointer++
		case 6:
			registers[1] = registers[0] / int(math.Pow(2, float64(comboVar(registers, instructon[1]))))
			instructionPointer++
		case 7:
			registers[2] = registers[0] / int(math.Pow(2, float64(comboVar(registers, instructon[1]))))
			instructionPointer++
		}
	}

	return strings.Join(output, ",")

}

func comboVar(registers []int, val int) int {
	if val < 4 {
		return val
	}
	return registers[val-4]
}
