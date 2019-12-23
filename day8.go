package main

import (
	"fmt"
)

func day8A() {
	image := readFile(8)[0]

	width, height := 25, 6
	numLayers := len(image) / width / height

	fewest0Digits := int(^uint(0) >> 1)
	product := 0

	i := 0
	for l := 0; l < numLayers; l++ {
		digitsInLayer := make(map[byte]int, 3)
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				digitsInLayer[image[i]]++
				i++
			}
		}

		if digitsInLayer['0'] < fewest0Digits {
			fewest0Digits = digitsInLayer['0']
			product = digitsInLayer['1'] * digitsInLayer['2']
		}
	}

	fmt.Printf("Result part 1: %v\n", product)
}

func day8B() {
	image := readFile(8)[0]

	width, height := 25, 6
	numLayers := len(image) / width / height
	layers := make([][]byte, numLayers)

	// Layer-ize image
	i := 0
	for l := 0; l < numLayers; l++ {
		layers[l] = make([]byte, 0)

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				layers[l] = append(layers[l], image[i])
				i++
			}
		}

	}

	// Final result
	finalResult := make([]byte, width*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			depth := 0
			for layers[depth][x+y*width] == '2' {
				depth++
			}
			finalResult[x+y*width] = layers[depth][x+y*width]
		}
	}

	fmt.Println("Result part 2:")
	for y := 0; y < height+2; y++ {
		for x := 0; x < width+2; x++ {
			// Border added by me
			if y == 0 || y == height+1 || x == 0 || x == width+1 {
				fmt.Printf("▓")

			} else if finalResult[(x-1)+(y-1)*width] == '0' {
				fmt.Printf("▓")
			} else {
				fmt.Printf("░")
			}
		}
		fmt.Printf("\n")
	}
}
