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

func solution() int {
	rows, err := util.ReadStrings("17", false, "\n\n")
	if err != nil {
		panic(err)
	}

	var registers []int
	var instructions [][2]int
	var instructionPointer int
	var rawInstructions string
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
			rawInstructions = row[strings.Index(row, " ")+1:]

			var prev int
			for i, instruction := range strings.Split(rawInstructions, ",") {
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

	// A % 8 == B
	// B ^ 1 == B
	// A / B == C
	// A / 8 == A
	// B ^ 4 == B
	// B ^ C == B
	// B % 8 == 2

	// count := 0
	// for {
	// 	registers[0] = count
	// 	if count%10000 == 0 {
	// 		fmt.Println(count)
	// 	}

	currentVal := 36880208374314
	endVal := 109037984852522 + (1073741824 * 100000)

	// cycles := []int{524283, 1572859, 2097147, 1048571, 16777211, 16252923, 11010043}
	cycles := []int{68719476731, 1030792151035}
	offsets := []int{0, 3, 5}

	// 2,4,1,1,7,5,0,3,1,4,4,0,5,5,3,0

	for {
		found := false
		for i := 0; i < len(cycles); i++ {
			for j := 0; j < len(offsets); j++ {
				testVal := currentVal + cycles[i] + 5 + offsets[j]
				registers[0] = testVal
				output := runProgram(registers, instructions, instructionPointer, false)
				if output[0] == 2 && output[1] == 4 && output[2] == 1 && output[3] == 1 && output[4] == 7 &&
					output[5] == 5 && output[6] == 0 && output[7] == 3 && output[8] == 1 && output[9] == 4 && output[10] == 4 {
					found = true
					if output[11] == 0 && output[12] == 5 && output[13] == 5 && output[14] == 3 && output[15] == 0 {
						return testVal
					}

				}
			}
			if found {
				currentVal += cycles[i] + 5
				break
			}
		}

		if currentVal > endVal || !found {
			break
		}
	}

	return 0
}

func runProgram(registers []int, instructions [][2]int, instructionPointer int, debug bool) []int {
	var output []int
	for {
		if instructionPointer > len(instructions)-1 {
			break
		}

		instructon := instructions[instructionPointer]

		switch instructon[0] {
		case 0:
			prev := registers[0]
			registers[0] = registers[0] / int(math.Pow(2, float64(comboVar(registers, instructon[1]))))
			if debug {
				fmt.Printf("adv operand: %d, combo: %d, register A %d->%d\n", instructon[1], comboVar(registers, instructon[1]), prev, registers[0])
			}
		case 1:
			prev := registers[1]
			registers[1] = registers[1] ^ instructon[1]
			if debug {
				fmt.Printf("bxl register B: %d, operand: %d -> register B: %d\n", prev, instructon[1], registers[1])
			}
		case 2:
			registers[1] = comboVar(registers, instructon[1]) % 8
			if debug {
				fmt.Printf("bst operand: %d, combo: %d, register B: %d\n", instructon[1], comboVar(registers, instructon[1]), registers[1])
			}
		case 3:
			if registers[0] != 0 {
				instructionPointer = instructon[1] / 2
				if debug {
					fmt.Printf("jnz operand: %d, instruction: %d\n", instructon[1], instructon[1]/2)
				}
			}
		case 4:
			prev := registers[1]
			registers[1] = registers[1] ^ registers[2]
			if debug {
				fmt.Printf("bxc register B: %d, register C: %d -> register B: %d\n", prev, registers[2], registers[1])
			}
		case 5:
			outVal := comboVar(registers, instructon[1]) % 8
			output = append(output, outVal)
			if debug {
				fmt.Printf("out operand: %d, combo: %d\n", instructon[1], outVal)
			}
		case 6:
			registers[1] = registers[0] / int(math.Pow(2, float64(comboVar(registers, instructon[1]))))
			if debug {
				fmt.Printf("bdv operand: %d, combo: %d, register A: %d, register B: %d\n", instructon[1], comboVar(registers, instructon[1]), registers[0], registers[1])
			}
		case 7:
			registers[2] = registers[0] / int(math.Pow(2, float64(comboVar(registers, instructon[1]))))
			if debug {
				fmt.Printf("cdv operand: %d, combo: %d, register A: %d, register C: %d\n", instructon[1], comboVar(registers, instructon[1]), registers[0], registers[2])
			}
		}

		if instructon[0] != 3 || registers[0] == 0 {
			instructionPointer++
		}
	}

	return output
}

func comboVar(registers []int, val int) int {
	if val < 4 {
		return val
	}
	return registers[val-4]
}
