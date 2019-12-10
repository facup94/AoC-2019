package main

import (
	"fmt"
	"strconv"
	"strings"
)

var relativeBase int = 0

func day9A() {
	entrada := strings.Split(readFile(9)[0], ",")

	memory := make([]int, len(entrada)*3)
	for index := 0; index < len(entrada); index++ {
		memory[index], _ = strconv.Atoi(entrada[index])
	}

	createFile("/Users/fparra/go/src/github.com/facup94/AoC2019-go/output.txt")

	out, _ := runProgramDay9(memory, 0)

	fmt.Printf("Result part 1: %v\n", out)
}

func runProgramDay9(memory []int, instructionPointer int) (outputBeforeHalt int, PC int) {
	for instructionPointer < len(memory) {
		opcode := memory[instructionPointer] % 100
		parametersMode := memory[instructionPointer] / 100

		writeFile("/Users/fparra/go/src/github.com/facup94/AoC2019-go/output.txt", fmt.Sprintf("%v\n", memory))

		// halt
		if opcode == 99 {
			instructionPointer++
			break
		}

		// add
		if opcode == 1 {
			parameter1 := getParameter(memory, instructionPointer, parametersMode, 1)
			parameter2 := getParameter(memory, instructionPointer, parametersMode, 2)
			memory[memory[instructionPointer+3]] = parameter1 + parameter2
			instructionPointer += 4
		}

		// multiply
		if opcode == 2 {
			parameter1 := getParameter(memory, instructionPointer, parametersMode, 1)
			parameter2 := getParameter(memory, instructionPointer, parametersMode, 2)
			memory[memory[instructionPointer+3]] = parameter1 * parameter2
			instructionPointer += 4
		}

		// save input
		if opcode == 3 {
			input := 1
			memory[memory[instructionPointer+1]] = input
			instructionPointer += 2
		}

		// output value
		if opcode == 4 {
			parameter1 := getParameter(memory, instructionPointer, parametersMode, 1)
			outputBeforeHalt = parameter1
			fmt.Println(outputBeforeHalt)
			instructionPointer += 2
		}

		// jump-if-true
		if opcode == 5 {
			parameter1 := getParameter(memory, instructionPointer, parametersMode, 1)
			parameter2 := getParameter(memory, instructionPointer, parametersMode, 2)
			if parameter1 != 0 {
				instructionPointer = parameter2
			} else {
				instructionPointer += 3
			}
		}

		// jump-if-false
		if opcode == 6 {
			parameter1 := getParameter(memory, instructionPointer, parametersMode, 1)
			parameter2 := getParameter(memory, instructionPointer, parametersMode, 2)
			if parameter1 == 0 {
				instructionPointer = parameter2
			} else {
				instructionPointer += 3
			}
		}

		// less than
		if opcode == 7 {
			parameter1 := getParameter(memory, instructionPointer, parametersMode, 1)
			parameter2 := getParameter(memory, instructionPointer, parametersMode, 2)
			if parameter1 < parameter2 {
				memory[memory[instructionPointer+3]] = 1
			} else {
				memory[memory[instructionPointer+3]] = 0
			}
			instructionPointer += 4
		}

		// less than
		if opcode == 8 {
			parameter1 := getParameter(memory, instructionPointer, parametersMode, 1)
			parameter2 := getParameter(memory, instructionPointer, parametersMode, 2)
			if parameter1 == parameter2 {
				memory[memory[instructionPointer+3]] = 1
			} else {
				memory[memory[instructionPointer+3]] = 0
			}
			instructionPointer += 4
		}

		// relative base offset
		if opcode == 9 {
			parameter1 := getParameter(memory, instructionPointer, parametersMode, 1)
			relativeBase += parameter1
			instructionPointer += 2
		}
	}

	return outputBeforeHalt, instructionPointer
}
