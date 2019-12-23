package main

import (
	"fmt"
	"strconv"
	"strings"
)

type coordinate struct {
	X int
	Y int
}

type coordinateWithDistance struct {
	X              int
	Y              int
	distanceCable1 int
	distanceCable2 int
}

func day3A() {
	lines := readFile(3)

	// Cable 1
	visited := make(map[coordinate]int)
	currentPosition := coordinate{0, 0}
	steps := strings.Split(lines[0], ",")
	for _, step := range steps {
		direction := step[0]
		distance, _ := strconv.Atoi(step[1:])

		if direction == 'U' {
			for i := 0; i < distance; i++ {
				currentPosition.X--
				visited[currentPosition] = 1
			}
		}

		if direction == 'D' {
			for i := 0; i < distance; i++ {
				currentPosition.X++
				visited[currentPosition] = 1
			}
		}

		if direction == 'L' {
			for i := 0; i < distance; i++ {
				currentPosition.Y--
				visited[currentPosition] = 1
			}
		}

		if direction == 'R' {
			for i := 0; i < distance; i++ {
				currentPosition.Y++
				visited[currentPosition] = 1
			}
		}
	}

	// Cable 2
	currentPosition = coordinate{0, 0}
	crossCoordinates := make([]coordinate, len(visited))
	steps = strings.Split(lines[1], ",")
	for _, step := range steps {
		direction := step[0]
		distance, _ := strconv.Atoi(step[1:])

		if direction == 'U' {
			for i := 0; i < distance; i++ {
				currentPosition.X--
				if _, ok := visited[currentPosition]; ok {
					crossCoordinates = append(crossCoordinates, currentPosition)
				}
			}
		}

		if direction == 'D' {
			for i := 0; i < distance; i++ {
				currentPosition.X++
				if _, ok := visited[currentPosition]; ok {
					crossCoordinates = append(crossCoordinates, currentPosition)
				}
			}
		}

		if direction == 'L' {
			for i := 0; i < distance; i++ {
				currentPosition.Y--
				if _, ok := visited[currentPosition]; ok {
					crossCoordinates = append(crossCoordinates, currentPosition)
				}
			}
		}

		if direction == 'R' {
			for i := 0; i < distance; i++ {
				currentPosition.Y++
				if _, ok := visited[currentPosition]; ok {
					crossCoordinates = append(crossCoordinates, currentPosition)
				}
			}
		}
	}

	// Look for closest
	minDistance := -1

	for _, crossPoint := range crossCoordinates {
		if crossPoint.X != 0 || crossPoint.Y != 0 {
			if manhattanDistance(crossPoint.X, crossPoint.Y, 0, 0) < minDistance || minDistance == -1 {
				minDistance = manhattanDistance(crossPoint.X, crossPoint.Y, 0, 0)
			}
		}
	}

	fmt.Printf("Result part 1: %v\n", minDistance)
}

func day3B() {
	lines := readFile(3)

	// Cable 1
	visited := make(map[coordinate]int)
	currentPosition := coordinate{0, 0}
	distanceCable1 := 0
	steps := strings.Split(lines[0], ",")
	for _, step := range steps {
		direction := step[0]
		distance, _ := strconv.Atoi(step[1:])

		if direction == 'U' {
			for i := 0; i < distance; i++ {
				currentPosition.X--
				distanceCable1++
				if _, ok := visited[currentPosition]; !ok {
					visited[currentPosition] = distanceCable1
				}
			}
		}

		if direction == 'D' {
			for i := 0; i < distance; i++ {
				currentPosition.X++
				distanceCable1++
				if _, ok := visited[currentPosition]; !ok {
					visited[currentPosition] = distanceCable1
				}
			}
		}

		if direction == 'L' {
			for i := 0; i < distance; i++ {
				currentPosition.Y--
				distanceCable1++
				if _, ok := visited[currentPosition]; !ok {
					visited[currentPosition] = distanceCable1
				}
			}
		}

		if direction == 'R' {
			for i := 0; i < distance; i++ {
				currentPosition.Y++
				distanceCable1++
				if _, ok := visited[currentPosition]; !ok {
					visited[currentPosition] = distanceCable1
				}
			}
		}
	}

	// Cable 2
	currentPosition = coordinate{0, 0}
	crossCoordinates := make([]coordinateWithDistance, 100)
	distanceCable2 := 0
	steps = strings.Split(lines[1], ",")
	for _, step := range steps {
		direction := step[0]
		distance, _ := strconv.Atoi(step[1:])

		if direction == 'U' {
			for i := 0; i < distance; i++ {
				currentPosition.X--
				distanceCable2++
				if distanceCable1, ok := visited[currentPosition]; ok {
					crossCoordinates = append(crossCoordinates, coordinateWithDistance{currentPosition.X, currentPosition.Y, distanceCable1, distanceCable2})
				}
			}
		}

		if direction == 'D' {
			for i := 0; i < distance; i++ {
				currentPosition.X++
				distanceCable2++
				if distanceCable1, ok := visited[currentPosition]; ok {
					crossCoordinates = append(crossCoordinates, coordinateWithDistance{currentPosition.X, currentPosition.Y, distanceCable1, distanceCable2})
				}
			}
		}

		if direction == 'L' {
			for i := 0; i < distance; i++ {
				currentPosition.Y--
				distanceCable2++
				if distanceCable1, ok := visited[currentPosition]; ok {
					crossCoordinates = append(crossCoordinates, coordinateWithDistance{currentPosition.X, currentPosition.Y, distanceCable1, distanceCable2})
				}
			}
		}

		if direction == 'R' {
			for i := 0; i < distance; i++ {
				currentPosition.Y++
				distanceCable2++
				if distanceCable1, ok := visited[currentPosition]; ok {
					crossCoordinates = append(crossCoordinates, coordinateWithDistance{currentPosition.X, currentPosition.Y, distanceCable1, distanceCable2})
				}
			}
		}
	}

	// Look for closest
	minDistance := -1

	for _, crossPoint := range crossCoordinates {
		if crossPoint.X != 0 || crossPoint.Y != 0 {
			if crossPoint.distanceCable1+crossPoint.distanceCable2 < minDistance || minDistance == -1 {
				minDistance = crossPoint.distanceCable1 + crossPoint.distanceCable2
			}
		}
	}

	fmt.Printf("Result part 2: %v\n", minDistance)
}
