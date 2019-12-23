package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day15A() {
	entrada := strings.Split(readFile(15)[0], ",")

	memory := make([]int, len(entrada)*2)
	for index := 0; index < len(entrada); index++ {
		memory[index], _ = strconv.Atoi(entrada[index])
	}

	var visited [41][41]bool
	var maze [41][41]rune
	for y := 0; y < len(maze); y++ {
		for x := 0; x < len(maze[y]); x++ {
			maze[y][x] = ' '
		}
	}
	pos := &position{21, 21}
	var PC int
	route := make([]position, 0)
	route = append(route, *pos)
	var status int
	var input int

	for status != 2 {

		maze[pos.y][pos.x] = '.'
		visited[pos.y][pos.x] = true

		surroundingBlocks := pos.getSurroundingBlocks()
		visitableBlocks := filterVisitableBlocks(surroundingBlocks, maze, visited)

		var nextPosition position
		if len(visitableBlocks) == 0 {
			route = route[:len(route)-1]

			nextPosition = route[len(route)-1]
			route = route[:len(route)-1]
		} else {
			nextPosition = visitableBlocks[0]
		}

		input = pos.calculateRequiredCommand(nextPosition)

		status, PC = runProgramDay15(memory, PC, input)

		if status == 0 {
			maze[nextPosition.y][nextPosition.x] = '#'
		} else {
			pos = &nextPosition
			route = append(route, nextPosition)
		}
	}

	fmt.Printf("Result part 1: %v\n", len(route)-1)
}

func day15B() {
	entrada := strings.Split(readFile(15)[0], ",")

	memory := make([]int, len(entrada)*2)
	for index := 0; index < len(entrada); index++ {
		memory[index], _ = strconv.Atoi(entrada[index])
	}

	var visited [41][41]bool
	var maze [41][41]rune
	for y := 0; y < len(maze); y++ {
		for x := 0; x < len(maze[y]); x++ {
			maze[y][x] = ' '
		}
	}
	pos := &position{21, 21}
	var PC int
	route := make([]position, 0)
	route = append(route, *pos)
	var status int
	var input int
	var oxyPos position

	for {

		maze[pos.y][pos.x] = '.'
		visited[pos.y][pos.x] = true

		surroundingBlocks := pos.getSurroundingBlocks()
		visitableBlocks := filterVisitableBlocks(surroundingBlocks, maze, visited)

		var nextPosition position
		if len(visitableBlocks) == 0 {
			route = route[:len(route)-1]

			if len(route) == 0 {
				break
			}

			nextPosition = route[len(route)-1]
			route = route[:len(route)-1]
		} else {
			nextPosition = visitableBlocks[0]
		}

		input = pos.calculateRequiredCommand(nextPosition)

		status, PC = runProgramDay15(memory, PC, input)

		if status == 0 {
			maze[nextPosition.y][nextPosition.x] = '#'
		} else {
			pos = &nextPosition
			route = append(route, nextPosition)
		}

		if status == 2 {
			oxyPos = *pos
		}
	}

	maze[oxyPos.y][oxyPos.x] = 'O'

	stepsToFillO2 := -1
	growed := true
	for growed {
		stepsToFillO2++
		growed = false
		var nextMaze [len(maze)][len(maze[0])]rune

		for y := 0; y < len(maze); y++ {
			for x := 0; x < len(maze[y]); x++ {
				nextMaze[y][x] = maze[y][x]
				if maze[y][x] == '#' || maze[y][x] == 'O' {
					continue
				}

				cPos := position{x, y}
				surroundingBlocks := cPos.getSurroundingBlocks()

				for _, b := range surroundingBlocks {
					if maze[b.y][b.x] == 'O' {
						nextMaze[y][x] = 'O'
						growed = true
						break
					}
				}
			}
		}
		maze = nextMaze

	}

	fmt.Printf("Result part 2: %v\n", stepsToFillO2)

}

type position struct {
	x int
	y int
}

func (pos position) calculateRequiredCommand(nextPosition position) int {
	// fmt.Println("calculateRequiredCommand -", "from:", pos, "- to:", nextPosition)
	if pos.y > nextPosition.y && pos.x == nextPosition.x {
		return 1 // north
	}
	if pos.y < nextPosition.y && pos.x == nextPosition.x {
		return 2 // south
	}
	if pos.y == nextPosition.y && pos.x > nextPosition.x {
		return 3 // west
	}
	if pos.y == nextPosition.y && pos.x < nextPosition.x {
		return 4 // east
	}
	fmt.Println("SUPER ERROR EN calculateRequiredCommand")
	return 0
}

func (pos position) getSurroundingBlocks() []position {
	p := make([]position, 0)
	if pos.x > 0 {
		p = append(p, position{pos.x - 1, pos.y})
	}
	if pos.y < 40 {
		p = append(p, position{pos.x, pos.y + 1})
	}
	if pos.x < 40 {
		p = append(p, position{pos.x + 1, pos.y})
	}
	if pos.y > 0 {
		p = append(p, position{pos.x, pos.y - 1})
	}
	return p
}

func filterVisitableBlocks(posibles []position, maze [41][41]rune, visited [41][41]bool) []position {
	pp := make([]position, 0)

	for _, p := range posibles {
		if maze[p.y][p.x] == '#' {
			continue
		}
		if visited[p.y][p.x] {
			continue
		}
		pp = append(pp, p)
	}

	return pp
}

func bfsDay15(startingNode int, nodes, relations []int) {
	visited := make([]bool, len(nodes))
	q := make([]int, 0)
	q = append(q, startingNode)

	for len(q) > 0 {
		n := q[0]
		fmt.Println(n)
		q = q[1:]
		visited[n] = true

		// x, y := n%100, n/100
		for i := nodes[n]; i < nodes[n+1]; i++ {
			if !visited[relations[i]] {
				q = append(q, relations[i])
			}
		}
	}
}

func runProgramDay15(memory []int, instructionPointer, input int) (int, int) {
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
