package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day16A() {
	entrada := readFile(16)[0]

	signal := make([]int, len(entrada))
	for index := 0; index < len(entrada); index++ {
		signal[index], _ = strconv.Atoi(string(entrada[index]))
	}

	for phase := 0; phase < 100; phase++ {
		newSignal := make([]int, len(signal))

		for iNS := 0; iNS < len(newSignal); iNS++ {

			// Sumas
			index := iNS
			i1 := 0
			for index < len(signal) {
				for i2 := 0; i2 < iNS+1; i2++ {
					if index+i2 >= len(signal) {
						break
					}
					newSignal[iNS] += signal[index+i2]
				}
				i1++
				index = iNS + i1*4*(iNS+1)
			}

			// Restas
			index = 2 + 3*iNS
			i1 = 0
			for index < len(signal) {
				for i2 := 0; i2 < iNS+1; i2++ {
					if index+i2 >= len(signal) {
						break
					}
					newSignal[iNS] -= signal[index+i2]
				}
				i1++
				index = 2 + 3*iNS + i1*4*(iNS+1)
			}

			newSignal[iNS] = intAbs(newSignal[iNS] % 10)
		}

		signal = newSignal
	}

	v := strings.Trim(strings.Join(strings.Split(fmt.Sprint(signal[:8]), " "), ""), "[]")
	fmt.Printf("Result part 1: %v\n", v)
}

func day16B() {
	entrada := readFile(16)[0]

	signal := make([]int, len(entrada)*10000)
	for index := 0; index < len(entrada); index++ {
		for j := 0; j < 10000; j++ {
			signal[index+len(entrada)*j], _ = strconv.Atoi(string(entrada[index]))
		}
	}

	offset, _ := strconv.Atoi(entrada[:7])
	
	for phase := 0; phase < 100; phase++ {
		newSignal := make([]int, len(signal))

		// Had to check reddit for this
		// The biggest single clue was that offset > len(signal)/2
		newSignal[len(signal)-1] = signal[len(signal)-1]

		for iNS := len(signal) - 2; iNS >= offset; iNS-- {
			newSignal[iNS] = newSignal[iNS+1] + signal[iNS]
			newSignal[iNS] = newSignal[iNS] % 10
		}

		signal = newSignal
	}

	v := strings.Trim(strings.Join(strings.Split(fmt.Sprint(signal[offset:offset+8]), " "), ""), "[]")
	fmt.Printf("Result part 2: %v\n", v)
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
