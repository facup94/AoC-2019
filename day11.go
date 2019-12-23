package main

import (
	"fmt"
	"strconv"
	"strings"
)

var panels map[string]int

func day11A() {
	entrada := strings.Split(readFile(11)[0], ",")

	memory := make([]int, len(entrada)*5)
	for index := 0; index < len(entrada); index++ {
		memory[index], _ = strconv.Atoi(entrada[index])
	}

	x, y, direction, PC := 0, 0, 0, 0

	panels = make(map[string]int)

	for {
		var colorToPaint int
		colorToPaint, PC = runProgramDay11(memory, PC, panels[fmt.Sprintf("(%v,%v)", x, y)])
		if colorToPaint == -1 {
			break
		}

		if colorToPaint == 0 {
			panels[fmt.Sprintf("(%v,%v)", x, y)] = 0
		} else {
			panels[fmt.Sprintf("(%v,%v)", x, y)] = 1
		}

		var rotation int
		rotation, PC = runProgramDay11(memory, PC, panels[fmt.Sprintf("(%v,%v)", x, y)])
		if rotation == -1 {
			break
		}

		// Rotate
		if rotation == 0 {
			direction -= 90
			if direction < 0 {
				direction += 360
			}
		} else {
			direction += 90
			if direction >= 360 {
				direction -= 360
			}
		}

		// Move
		if direction == 0 {
			y--
		} else if direction == 90 {
			x++
		} else if direction == 180 {
			y++
		} else if direction == 270 {
			x--
		}

	}

	fmt.Printf("Result part 1: %v\n", len(panels))
}
func day11B() {
	entrada := strings.Split(readFile(11)[0], ",")

	memory := make([]int, len(entrada)*5)
	for index := 0; index < len(entrada); index++ {
		memory[index], _ = strconv.Atoi(entrada[index])
	}

	x, y, direction, PC := 0, 0, 0, 0

	panels = make(map[string]int)
	panels["(0,0)"] = 1 // Start on white panel

	for {
		var colorToPaint int
		colorToPaint, PC = runProgramDay11(memory, PC, panels[fmt.Sprintf("(%v,%v)", x, y)])
		if colorToPaint == -1 {
			break
		}

		if colorToPaint == 0 {
			panels[fmt.Sprintf("(%v,%v)", x, y)] = 0
		} else {
			panels[fmt.Sprintf("(%v,%v)", x, y)] = 1
		}

		var rotation int
		rotation, PC = runProgramDay11(memory, PC, panels[fmt.Sprintf("(%v,%v)", x, y)])
		if rotation == -1 {
			break
		}

		// Rotate
		if rotation == 0 {
			direction -= 90
			if direction < 0 {
				direction += 360
			}
		} else {
			direction += 90
			if direction >= 360 {
				direction -= 360
			}
		}

		// Move
		if direction == 0 {
			y--
		} else if direction == 90 {
			x++
		} else if direction == 180 {
			y++
		} else if direction == 270 {
			x--
		}

	}

	var minX, minY = 1000, 1000
	var maxX, maxY int
	for k := range panels {
		s := strings.Split(k[1:len(k)-1], ",")
		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])

		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}

	fmt.Println("Result part 2:")
	for y := minY - 1; y <= maxY+1; y++ {
		for x := minX; x <= maxX; x++ {
			if panels[fmt.Sprintf("(%v,%v)", x, y)] == 0 {
				fmt.Print("■")
			} else {
				fmt.Print("□")
			}
		}
		fmt.Println()
	}

}

func runProgramDay11(memory []int, instructionPointer, input int) (int, int) {
	for instructionPointer < len(memory) {
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
