package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day17A() {
	entrada := strings.Split(readFile(17)[0], ",")

	memory := make([]int, len(entrada)*10)
	for index := 0; index < len(entrada); index++ {
		memory[index], _ = strconv.Atoi(string(entrada[index]))
	}

	scaffolding := make(map[coordinate]int)

	var curX, curY int
	var PC, output int
	for {
		output, PC = runProgramDay17(memory, PC, 0)
		if output == -1 {
			break
		}

		if output == 35 {
			scaffolding[coordinate{curX, curY}]++
		}

		curX++
		if output == 10 {
			curY++
			curX = 0
		}
	}

	alignmentParameters := 0
	for scaffold := range scaffolding {
		if countAdjacentScaffolding(scaffolding, scaffold) == 4 {
			alignmentParameters += scaffold.X * scaffold.Y
		}
	}

	fmt.Printf("Result part 1: %v\n", alignmentParameters)
}

func day17B() {
	entrada := strings.Split(readFile(17)[0], ",")

	memory := make([]int, len(entrada)*15)
	for index := 0; index < len(entrada); index++ {
		memory[index], _ = strconv.Atoi(string(entrada[index]))
	}
	memory[0] = 2 // Wake up the vacuum robot

	// All these were obtained by trial and error. Hand calculated (?)
	// sSequence := "R6,L10,R8,R8,R12,L8,L8,R6,L10,R8,R8,R12,L8,L8,L10,R6,R6,L8,R6,L10,R8,R8,R12,L8,L8,L10,R6,R6,L8,R6,L10,R8,L10,R6,R6,L8"
	//               R6 L10 R8 R8 R12 L8 L8 R6 L10 R8 R8 R12 L8 L8 L10 R6 R6 L8 R6 L10 R8 R8 R12 L8 L8 L10 R6 R6 L8 R6 L10 R8 L10 R6 R6 L8
	// A B A B C A B C A C
	// A = R6 L10 R8
	// B = R8 R12 L8 L8
	// C = L10 R6 R6 L8

	mainMovementRoutine := "A,B,A,B,C,A,B,C,A,C\n"
	movementFunctionA := "R,6,L,10,R,8\n"
	movementFunctionB := "R,8,R,12,L,8,L,8\n"
	movementFunctionC := "L,10,R,6,R,6,L,8\n"
	videoFeed := "n\n"
	allInputs := mainMovementRoutine + movementFunctionA + movementFunctionB + movementFunctionC + videoFeed

	inputList := make([]int, 0)

	for _, c := range allInputs {
		inputList = append(inputList, int(c))
	}
	inputList = append(inputList, 0)

	var PC, output, inputIndex int
	for {
		output, PC = runProgramDay17(memory, PC, inputList[inputIndex])
		if output == -2 {
			inputIndex++
			continue
		}
		if output > 255 || output < 1 {
			break
		}
	}

	fmt.Printf("Result part 2: %v\n", output)
}

func runProgramDay17(memory []int, instructionPointer, input int) (int, int) {
	for instructionPointer < len(memory) {
		// fmt.Println(instructionPointer, memory[instructionPointer])
		opcode := memory[instructionPointer] % 100
		parametersMode := memory[instructionPointer] / 100

		// halt
		if opcode == 99 {
			instructionPointer++
			return -1, len(memory)
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
			return -2, instructionPointer
		}

		// output value
		if opcode == 4 {
			parameter1 := getParameter(memory, instructionPointer, parametersMode, 1)
			instructionPointer += 2
			return parameter1, instructionPointer
		}

		// jump-if-true
		if opcode == 5 {
			params := getParams(memory, getModes(memory[instructionPointer]), instructionPointer, 2)
			if memory[params[0]] != 0 {
				instructionPointer = memory[params[1]]
			} else {
				instructionPointer += 3
			}
		}

		// jump-if-false
		if opcode == 6 {
			params := getParams(memory, getModes(memory[instructionPointer]), instructionPointer, 2)
			if memory[params[0]] == 0 {
				instructionPointer = memory[params[1]]
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

	return -1, -1
}

func countAdjacentScaffolding(scaffolding map[coordinate]int, pos coordinate) int {
	amount := 0
	if _, ok := scaffolding[coordinate{pos.X, pos.Y - 1}]; ok {
		amount++
	}
	if _, ok := scaffolding[coordinate{pos.X + 1, pos.Y}]; ok {
		amount++
	}
	if _, ok := scaffolding[coordinate{pos.X, pos.Y + 1}]; ok {
		amount++
	}
	if _, ok := scaffolding[coordinate{pos.X - 1, pos.Y}]; ok {
		amount++
	}
	return amount
}
