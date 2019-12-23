package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day2A() {
	s := strings.Split(readFile(2)[0], ",")

	var memory []int = make([]int, len(s))

	for i, val := range s {
		memory[i], _ = strconv.Atoi(val)
	}

	// restore the gravity assist program (your puzzle input) to the "1202 program alarm"
	memory[1] = 12
	memory[2] = 2

	runProgram(memory)

	fmt.Printf("Result part 1: %v\n", memory[0])
}

func day2B() {
	s := strings.Split(readFile(2)[0], ",")

	var memory []int = make([]int, len(s))

	for i, val := range s {
		memory[i], _ = strconv.Atoi(val)
	}

	memoryInitialState := make([]int, len(memory))
	copy(memoryInitialState, memory)

	foundSolution := false

	for noun := 0; noun < 100 && !foundSolution; noun++ {
		for verb := 0; verb < 100 && !foundSolution; verb++ {
			copy(memory, memoryInitialState)

			memory[1] = noun
			memory[2] = verb

			runProgram(memory)

			if memory[0] == 19690720 {
				fmt.Printf("Result part 2: %v\n", 100*noun+verb)
				foundSolution = true
			}
		}
	}
}

func runProgram(memory []int) {
	for instructionPointer := 0; instructionPointer < len(memory); {
		opcode := memory[instructionPointer]

		// halt
		if opcode == 99 {
			instructionPointer++
			break
		}

		// add
		if opcode == 1 {
			memory[memory[instructionPointer+3]] = memory[memory[instructionPointer+1]] + memory[memory[instructionPointer+2]]
			instructionPointer += 4
		}

		// multiply
		if opcode == 2 {
			memory[memory[instructionPointer+3]] = memory[memory[instructionPointer+1]] * memory[memory[instructionPointer+2]]
			instructionPointer += 4
		}
	}
}
