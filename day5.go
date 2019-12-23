package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func day5A() {
	lines := readFile(5)
	s := strings.Split(lines[0], ",")

	var memory []int = make([]int, len(s))

	for i, val := range s {
		memory[i], _ = strconv.Atoi(val)
	}

	runProgramDay5(memory, false)
}

func day5B() {
	lines := readFile(5)
	s := strings.Split(lines[0], ",")

	var memory []int = make([]int, len(s))

	for i, val := range s {
		memory[i], _ = strconv.Atoi(val)
	}

	runProgramDay5(memory, true)
}

func runProgramDay5(memory []int, partB bool) {
	outputBeforeHalt := 0

	for instructionPointer := 0; instructionPointer < len(memory); {
		opcode := memory[instructionPointer] % 100
		parametersMode := memory[instructionPointer] / 100

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
			if partB {
				input = 5
			}
			memory[memory[instructionPointer+1]] = input
			instructionPointer += 2
		}

		// output value
		if opcode == 4 {
			outputBeforeHalt = memory[memory[instructionPointer+1]]
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
	}

	if partB {
		fmt.Printf("Result part 2: %v\n", outputBeforeHalt)
	} else {
		fmt.Printf("Result part 1: %v\n", outputBeforeHalt)
	}
}

func getParameter(memory []int, pointer, parametersMode, position int) int {
	var parameter int
	if getDigit(parametersMode, position) == 0 {
		// Position Mode
		parameter = memory[memory[pointer+position]]
	} else if getDigit(parametersMode, position) == 1 {
		parameter = memory[pointer+position]
	} else if getDigit(parametersMode, position) == 2 {
		// Relative Mode
		parameter = memory[memory[pointer+position]+relativeBase]
	}
	return parameter
}

func getDigit(number, position int) int {
	return (number / int(math.Pow10(position-1))) % int(math.Pow10(position))
}
