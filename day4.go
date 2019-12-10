package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day4A() {
	values := strings.Split(readFile(4)[0], "-")
	min, _ := strconv.Atoi(values[0])
	max, _ := strconv.Atoi(values[1])

	var amount int64 = 0
	for i := min; i < max+1; i++ {
		r := []rune(strconv.Itoa(i))

		if !twoAdjacentDigitSame(r) {
			continue
		}

		if !incrementalDigits(r) {
			continue
		}

		amount++
	}

	fmt.Printf("Result part 1: %v\n", amount)
}

func day4B() {
	values := strings.Split(readFile(4)[0], "-")
	min, _ := strconv.Atoi(values[0])
	max, _ := strconv.Atoi(values[1])

	var amount int64 = 0
	for i := min; i < max+1; i++ {
		r := []rune(strconv.Itoa(i))

		if !twoAdjacentButNotTreeDigitSame(r) {
			continue
		}

		if !incrementalDigits(r) {
			continue
		}

		amount++
	}

	fmt.Printf("Result part 2: %v\n", amount)
}

func twoAdjacentDigitSame(number []rune) bool {
	for i := 0; i < 5; i++ {
		if number[i] == number[i+1] {
			return true
		}
	}
	return false
}

func twoAdjacentButNotTreeDigitSame(number []rune) bool {
	for i := 0; i < 5; i++ {
		if number[i] == number[i+1] {
			// Check for third before pair
			if i > 0 {
				if number[i-1] == number[i] {
					continue
				}
			}

			// Check for third after pair
			if i < 4 {
				if number[i+1] == number[i+2] {
					continue
				}
			}

			return true
		}
	}

	return false
}

func incrementalDigits(number []rune) bool {
	for i := 0; i < len(number)-1; i++ {
		if number[i] > number[i+1] {
			return false
		}
	}

	return true
}
