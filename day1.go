package main

import (
	"fmt"
	"strconv"
)

func day1A() {
	var total int = 0

	for _, line := range readFile(1) {
		moduleMass, _ := strconv.Atoi(line)

		moduleFuel := calculateFuelPart1(moduleMass)

		total += moduleFuel
	}
	fmt.Printf("Result part 1: %v\n", total)
}

func day1B() {
	var total int = 0

	for _, line := range readFile(1) {
		moduleMass, _ := strconv.Atoi(line)

		moduleFuel := calculateFuelPart2(moduleMass)

		total += moduleFuel
	}
	fmt.Printf("Result part 2: %v\n", total)
}

func calculateFuelPart1(mass int) (fuel int) {
	fuel = mass/3 - 2
	return
}

func calculateFuelPart2(mass int) (fuel int) {
	fuel = 0
	partialFuel := calculateFuelPart1(mass)

	for partialFuel != 0 {
		fuel += partialFuel
		partialFuel = calculateFuelPart1(partialFuel)
		if partialFuel < 0 {
			partialFuel = 0
		}
	}
	return
}
