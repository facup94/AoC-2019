package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type moon struct {
	x  int
	y  int
	z  int
	xV int
	yV int
	zV int
}

func day12A() {
	lines := readFile(12)

	// Initialize moons
	var moons [4]moon
	for i := 0; i < 4; i++ {
		var pos [3]int
		for j, s := range strings.Split(lines[i][1:len(lines[i])-1], ",") {
			pos[j], _ = strconv.Atoi(strings.Split(s, "=")[1])
		}
		moons[i] = moon{pos[0], pos[1], pos[2], 0, 0, 0}
	}

	// Simulate time-steps
	for step := 0; step < 1000; step++ {
		// Update velocity based on gravity
		for i := 0; i < 3; i++ {
			for j := i; j < 4; j++ {
				moons[i].xV += getGravity(moons[i].x, moons[j].x)
				moons[i].yV += getGravity(moons[i].y, moons[j].y)
				moons[i].zV += getGravity(moons[i].z, moons[j].z)

				moons[j].xV -= getGravity(moons[i].x, moons[j].x)
				moons[j].yV -= getGravity(moons[i].y, moons[j].y)
				moons[j].zV -= getGravity(moons[i].z, moons[j].z)
			}
		}

		// Update position based on velocity
		for i := 0; i < 4; i++ {
			moons[i].y += moons[i].yV
			moons[i].z += moons[i].zV
			moons[i].x += moons[i].xV
		}

		// Calculate energy
		var totalEnergy int
		for i := 0; i < 4; i++ {
			potentialEnergy, kineticEnergy := 0, 0
			potentialEnergy += int(math.Abs(float64(moons[i].x))) + int(math.Abs(float64(moons[i].y))) + int(math.Abs(float64(moons[i].z)))
			kineticEnergy += int(math.Abs(float64(moons[i].xV))) + int(math.Abs(float64(moons[i].yV))) + int(math.Abs(float64(moons[i].zV)))
			totalEnergy += potentialEnergy * kineticEnergy
		}

		if step == 999 {
			fmt.Printf("Result part 1: %v\n", totalEnergy)
		}
	}
}

func day12B() {
	lines := readFile(12)

	// Initialize moons
	var moons [4]moon
	for i := 0; i < 4; i++ {
		var pos [3]int
		for j, s := range strings.Split(lines[i][1:len(lines[i])-1], ",") {
			pos[j], _ = strconv.Atoi(strings.Split(s, "=")[1])
		}
		moons[i] = moon{pos[0], pos[1], pos[2], 0, 0, 0}
	}

	cycleTimeX, cycleTimeY, cycleTimeZ := 0, 0, 0

	visited := make(map[string]int)
	step := 0
	for {
		for i := 0; i < 3; i++ {
			for j := i + 1; j < 4; j++ {
				moons[i].xV += getGravity(moons[i].x, moons[j].x)
				moons[j].xV -= getGravity(moons[i].x, moons[j].x)
			}
		}

		for i := 0; i < 4; i++ {
			moons[i].x += moons[i].xV
		}

		k := fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v)", moons[0].x, moons[1].x, moons[2].x, moons[3].x, moons[0].xV, moons[1].xV, moons[2].xV, moons[3].xV)
		if v, ok := visited[k]; ok {
			cycleTimeX = step - v
			break
		} else {
			visited[k] = step
		}
		step++
	}

	visited = make(map[string]int)
	step = 0
	for {
		for i := 0; i < 3; i++ {
			for j := i + 1; j < 4; j++ {
				moons[i].yV += getGravity(moons[i].y, moons[j].y)
				moons[j].yV -= getGravity(moons[i].y, moons[j].y)
			}
		}

		for i := 0; i < 4; i++ {
			moons[i].y += moons[i].yV
		}

		k := fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v)", moons[0].y, moons[1].y, moons[2].y, moons[3].y, moons[0].yV, moons[1].yV, moons[2].yV, moons[3].yV)
		if v, ok := visited[k]; ok {
			cycleTimeY = step - v
			break
		} else {
			visited[k] = step
		}
		step++
	}

	visited = make(map[string]int)
	step = 0
	for {
		for i := 0; i < 3; i++ {
			for j := i + 1; j < 4; j++ {
				moons[i].zV += getGravity(moons[i].z, moons[j].z)
				moons[j].zV -= getGravity(moons[i].z, moons[j].z)
			}
		}

		for i := 0; i < 4; i++ {
			moons[i].z += moons[i].zV
		}

		k := fmt.Sprintf("(%v,%v,%v,%v,%v,%v,%v,%v)", moons[0].z, moons[1].z, moons[2].z, moons[3].z, moons[0].zV, moons[1].zV, moons[2].zV, moons[3].zV)
		if v, ok := visited[k]; ok {
			cycleTimeZ = step - v
			break
		} else {
			visited[k] = step
		}
		step++
	}

	// Prime Factorization
	cycleTimeXFactors := make(map[int]int)
	i := 2
	for i <= cycleTimeX {
		if cycleTimeX%i == 0 {
			cycleTimeXFactors[i]++
			cycleTimeX /= i
		} else {
			i++
		}
	}
	cycleTimeYFactors := make(map[int]int)
	i = 2
	for i <= cycleTimeY {
		if cycleTimeY%i == 0 {
			cycleTimeYFactors[i]++
			cycleTimeY /= i
		} else {
			i++
		}
	}
	cycleTimeZFactors := make(map[int]int)
	i = 2
	for i <= cycleTimeZ {
		if cycleTimeZ%i == 0 {
			cycleTimeZFactors[i]++
			cycleTimeZ /= i
		} else {
			i++
		}
	}

	// List of factors where it occurs most often
	finalFactors := make(map[int]int)
	for factor, times := range cycleTimeXFactors {
		if v, ok := finalFactors[factor]; ok {
			if times > v {
				finalFactors[factor] = times
			}
		} else {
			finalFactors[factor] = times
		}
	}
	for factor, times := range cycleTimeYFactors {
		if v, ok := finalFactors[factor]; ok {
			if times > v {
				finalFactors[factor] = times
			}
		} else {
			finalFactors[factor] = times
		}
	}
	for factor, times := range cycleTimeZFactors {
		if v, ok := finalFactors[factor]; ok {
			if times > v {
				finalFactors[factor] = times
			}
		} else {
			finalFactors[factor] = times
		}
	}

	// Final result
	result := 1
	for factor, times := range finalFactors {
		result *= int(math.Pow(float64(factor), float64(times)))
	}

	fmt.Printf("Result part 2: %v\n", result)
}

func getGravity(a, b int) int {
	if a < b {
		return 1
	} else if a > b {
		return -1
	} else {
		return 0
	}
}
