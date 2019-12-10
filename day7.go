package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day7A() {
	lines := readFile(7)
	s := strings.Split(lines[0], ",")

	largestOutput := 0

	var originalMemory []int = make([]int, len(s))
	var memory []int = make([]int, len(s))

	for i, val := range s {
		originalMemory[i], _ = strconv.Atoi(val)
	}

	for _, perm := range permutations([]int{0, 1, 2, 3, 4}) {
		ioVar := 0
		for i := 0; i < 5; i++ {
			copy(memory, originalMemory)
			ioVar, _, _, _ = runProgramDay7(memory, 0, perm[i], ioVar, false)
		}
		if ioVar > largestOutput {
			largestOutput = ioVar
		}
	}
	fmt.Printf("Result part 1: %v\n", largestOutput)

}

func day7B() {
	lines := readFile(7)
	s := strings.Split(lines[0], ",")

	var originalMemory []int = make([]int, len(s))
	for i, val := range s {
		originalMemory[i], _ = strconv.Atoi(val)
	}

	largestOutput := 0

	for _, perm := range permutations([]int{5, 6, 7, 8, 9}) {
		// initialize memories
		var memories [5][]int
		var ampPCs [5]int
		var ampPhaseGiven [5]bool
		for i := 0; i < len(memories); i++ {
			memories[i] = make([]int, len(s))
			copy(memories[i], originalMemory)
		}

		var ioVar int
		var halted = false
		var lastAmpEOutput int

		for !halted {
			for i := 0; i < 5; i++ {
				halted = false
				ioVar, ampPCs[i], halted, ampPhaseGiven[i] = runProgramDay7(memories[i], ampPCs[i], perm[i], ioVar, ampPhaseGiven[i])
				
				if i == 4 && !halted {
					lastAmpEOutput = ioVar
				}
			}
		}

		if lastAmpEOutput > largestOutput {
			largestOutput = lastAmpEOutput
		}

	}
	fmt.Printf("Result part 2: %v\n", largestOutput)

}

func runProgramDay7(memory []int, PC, phase, input int, askedForPhase bool) (int, int, bool, bool) {
	output := 0
	halted := false

	for PC < len(memory) {
		opcode := memory[PC] % 100
		parametersMode := memory[PC] / 100

		// halt
		if opcode == 99 {
			halted = true
			PC++
			break
		}

		// add
		if opcode == 1 {
			parameter1 := getParameter(memory, PC, parametersMode, 1)
			parameter2 := getParameter(memory, PC, parametersMode, 2)
			memory[memory[PC+3]] = parameter1 + parameter2
			PC += 4
		}

		// multiply
		if opcode == 2 {
			parameter1 := getParameter(memory, PC, parametersMode, 1)
			parameter2 := getParameter(memory, PC, parametersMode, 2)
			memory[memory[PC+3]] = parameter1 * parameter2
			PC += 4
		}

		// save input
		if opcode == 3 {
			if askedForPhase {
				memory[memory[PC+1]] = input
			} else {
				memory[memory[PC+1]] = phase
				askedForPhase = true
			}
			PC += 2
		}

		// output value
		if opcode == 4 {
			output = memory[memory[PC+1]]
			PC += 2
			break
		}

		// jump-if-true
		if opcode == 5 {
			parameter1 := getParameter(memory, PC, parametersMode, 1)
			parameter2 := getParameter(memory, PC, parametersMode, 2)
			if parameter1 != 0 {
				PC = parameter2
			} else {
				PC += 3
			}
		}

		// jump-if-false
		if opcode == 6 {
			parameter1 := getParameter(memory, PC, parametersMode, 1)
			parameter2 := getParameter(memory, PC, parametersMode, 2)
			if parameter1 == 0 {
				PC = parameter2
			} else {
				PC += 3
			}
		}

		// less than
		if opcode == 7 {
			parameter1 := getParameter(memory, PC, parametersMode, 1)
			parameter2 := getParameter(memory, PC, parametersMode, 2)
			if parameter1 < parameter2 {
				memory[memory[PC+3]] = 1
			} else {
				memory[memory[PC+3]] = 0
			}
			PC += 4
		}

		// less than
		if opcode == 8 {
			parameter1 := getParameter(memory, PC, parametersMode, 1)
			parameter2 := getParameter(memory, PC, parametersMode, 2)
			if parameter1 == parameter2 {
				memory[memory[PC+3]] = 1
			} else {
				memory[memory[PC+3]] = 0
			}
			PC += 4
		}

	}

	return output, PC, halted, askedForPhase
}
