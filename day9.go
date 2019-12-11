package main

import (
	"fmt"
	"strconv"
	"strings"
)

var relativeBase int = 0

func day9A() {
	entrada := strings.Split(readFile(9)[0], ",")

	memory := make([]int, len(entrada)*2)
	for index := 0; index < len(entrada); index++ {
		memory[index], _ = strconv.Atoi(entrada[index])
	}

	out, _ := runProgramDay9(memory, 0, 1)
	fmt.Printf("Result part 1: %v\n", out)
}

func day9B() {
	relativeBase = 0
	entrada := strings.Split(readFile(9)[0], ",")

	memory := make([]int, len(entrada)*2)
	for index := 0; index < len(entrada); index++ {
		memory[index], _ = strconv.Atoi(entrada[index])
	}

	out, _ := runProgramDay9(memory, 0, 2)
	fmt.Printf("Result part 2: %v\n", out)
}

func runProgramDay9(memory []int, instructionPointer, input int) (outputBeforeHalt int, PC int) {
	for instructionPointer < len(memory) {
		opcode := memory[instructionPointer] % 100
		parametersMode := memory[instructionPointer] / 100

		// halt
		if opcode == 99 {
			instructionPointer++
			break
		}

		// add
		if opcode == 1 {
			params := getParams(memory, getModes(memory[instructionPointer]), instructionPointer, 3)
			memory[params[2]] = memory[params[0]] + memory[params[1]]
			instructionPointer += 4
		}

		// multiply
		if opcode == 2 {
			params := getParams(memory, getModes(memory[instructionPointer]), instructionPointer, 3)
			memory[params[2]] = memory[params[0]] * memory[params[1]]
			instructionPointer += 4
		}

		// save input
		if opcode == 3 {
			params := getParams(memory, getModes(memory[instructionPointer]), instructionPointer, 1)
			memory[params[0]] = input
			instructionPointer += 2
		}

		// output value
		if opcode == 4 {
			parameter1 := getParameter(memory, instructionPointer, parametersMode, 1)
			outputBeforeHalt = parameter1
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
			params := getParams(memory, getModes(memory[instructionPointer]), instructionPointer, 3)
			if memory[params[0]] < memory[params[1]] {
				memory[params[2]] = 1
			} else {
				memory[params[2]] = 0
			}
			instructionPointer += 4
		}

		// less than
		if opcode == 8 {
			params := getParams(memory, getModes(memory[instructionPointer]), instructionPointer, 3)
			if memory[params[0]] == memory[params[1]] {
				memory[params[2]] = 1
			} else {
				memory[params[2]] = 0
			}
			instructionPointer += 4
		}

		// relative base offset
		if opcode == 9 {
			params := getParams(memory, getModes(memory[instructionPointer]), instructionPointer, 1)
			relativeBase += memory[params[0]]
			instructionPointer += 2
		}
	}

	return outputBeforeHalt, instructionPointer
}

func getModes(instruction int) []int {
	pModes := instruction / 100
	modes := make([]int, 3)
	for index := 0; index < 3; index++ {
		modes[index] = pModes % 10
		pModes /= 10
	}
	return modes
}

func getParams(memory, modes []int, PC, size int) []int {
	values := make([]int, size)
	for i := 0; i < size; i++ {
		if modes[i] == 0 {
			values[i] = memory[PC+1+i] // Position Mode
		} else if modes[i] == 1 {
			values[i] = PC + 1 + i
		} else if modes[i] == 2 {
			values[i] = memory[PC+1+i] + relativeBase // Relative Mode
		}
	}
	return values
}
