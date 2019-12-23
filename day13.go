package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day13A() {
	entrada := strings.Split(readFile(13)[0], ",")

	memory := make([]int, len(entrada)*2)
	for index := 0; index < len(entrada); index++ {
		memory[index], _ = strconv.Atoi(entrada[index])
	}

	blockTiles := make(map[string]bool)
	var PC int
	var out int
	for {
		out, PC = runProgramDay11(memory, PC, 0)
		xPos := out
		if out == -1 && PC == len(memory) {
			break
		}

		out, PC = runProgramDay11(memory, PC, 0)
		yPos := out

		out, PC = runProgramDay11(memory, PC, 0)
		tileType := out

		if tileType == 2 {
			blockTiles[fmt.Sprint(xPos, yPos)] = true
		}
	}
	fmt.Printf("Result part 1: %v\n", len(blockTiles))
}

func day13B() {
	entrada := strings.Split(readFile(13)[0], ",")

	memory := make([]int, len(entrada)*5)
	for index := 0; index < len(entrada); index++ {
		memory[index], _ = strconv.Atoi(entrada[index])
	}
	memory[0] = 2 // Play for free
	// var count int // 3240

	var PC int
	var out int
	var ballX, paddleX, joystick, score int
	joystick = -1
	for {
		out, PC = runProgramDay11(memory, PC, joystick)
		xPos := out
		if out == -1 && PC == len(memory) {
			break
		}

		out, PC = runProgramDay11(memory, PC, joystick)
		yPos := out

		out, PC = runProgramDay11(memory, PC, joystick)
		tileType := out

		if xPos == -1 && yPos == 0 {
			score = tileType
		}

		if tileType == 3 {
			paddleX = xPos
			joystick = calculateJoystick(ballX, paddleX)
		}
		if tileType == 4 {
			ballX = xPos
			joystick = calculateJoystick(ballX, paddleX)
		}

	}
	fmt.Printf("Result part 2: %v\n", score)
}

func calculateJoystick(ball, paddle int) int {
	if paddle < ball {
		return 1
	} else if paddle > ball {
		return -1
	} else {
		return 0
	}
}
