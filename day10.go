package main

import (
	"fmt"
	"math"
)

func day10A() {
	lines := readFile(10)

	width, height := len(lines[0]), len(lines)

	// Create maps
	mapa := make([][]rune, height)
	distancias := make([][]int, height)
	for i := 0; i < height; i++ {
		mapa[i] = []rune(lines[i])
		distancias[i] = make([]int, width)
	}

	for y0 := 0; y0 < height; y0++ {
		for x0 := 0; x0 < width; x0++ {
			if mapa[y0][x0] != '#' {
				continue
			}

			visibleAsteroids := 0

			for y1 := 0; y1 < height; y1++ {
				for x1 := 0; x1 < width; x1++ {
					if mapa[y1][x1] != '#' {
						continue
					}

					if y0 == y1 && x0 == x1 {
						continue
					}

					lineLen := distanceBetweenAsteroids(x1, y1, x0, y0)

					visible := true

					for y := intMin(y0, y1); y < intMax(y0, y1)+1; y++ {
						for x := intMin(x0, x1); x < intMax(x0, x1)+1; x++ {
							if mapa[y][x] != '#' {
								continue
							}

							if (x0 == x && y0 == y) || (x1 == x && y1 == y) {
								continue
							}

							d1 := distanceBetweenAsteroids(x, y, x0, y0)
							d2 := distanceBetweenAsteroids(x, y, x1, y1)

							if math.Abs(lineLen-(d1+d2)) < 0.00001 {
								visible = false
								break
							}
						}
					}

					if visible {
						visibleAsteroids++
					}
				}
			}

			distancias[y0][x0] = visibleAsteroids

		}
	}

	maxVisibleAsteroids, maxX, maxY := 0, 0, 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if distancias[y][x] > maxVisibleAsteroids {
				maxVisibleAsteroids = distancias[y][x]
				maxX = x
				maxY = y
			}
		}
	}
	fmt.Printf("Result part 1: %v (%v,%v)\n", maxVisibleAsteroids, maxX, maxY)
}
func day10B() {
	lines := readFile(10)

	width, height := len(lines[0]), len(lines)
	xStation, yStation := 8, 16

	// Create maps
	mapa := make([][]rune, height)
	angulos := make([][]float64, height)
	for i := 0; i < height; i++ {
		mapa[i] = []rune(lines[i])
		for j := 0; j < width; j++ {
			angulos[i] = append(angulos[i], -1)
		}
	}

	asteroidsVaporized := 0

	for asteroidsVaporized < 200 {

		// Calculate angle for each visible asteroid
		for y1 := 0; y1 < height; y1++ {
			for x1 := 0; x1 < width; x1++ {

				// Calculate angle only for asteroids
				if mapa[y1][x1] != '#' {
					continue
				}

				// Don't calculate for station's asteroid
				if y1 == yStation && x1 == xStation {
					continue
				}

				lineLen := distanceBetweenAsteroids(x1, y1, xStation, yStation)

				visible := true

				// For each posible asteroid between station and currently analized asteroid
				for y := intMin(y1, yStation); y < intMax(y1, yStation)+1; y++ {
					for x := intMin(x1, xStation); x < intMax(x1, xStation)+1; x++ {
						if mapa[y][x] != '#' {
							continue
						}

						if (xStation == x && yStation == y) || (x1 == x && y1 == y) {
							continue
						}

						d1 := distanceBetweenAsteroids(x, y, xStation, yStation)
						d2 := distanceBetweenAsteroids(x, y, x1, y1)

						// Check if it is in the middle
						if math.Abs(lineLen-(d1+d2)) < 0.00001 {
							visible = false
							break
						}
					}
				}

				// Calculate angle between station and currently analized asteroid
				if visible {
					angulos[y1][x1] = math.Atan2(float64(y1-yStation), float64(x1-xStation))*180/math.Pi + 90

					if angulos[y1][x1] < 0 {
						angulos[y1][x1] += 360
					}
				}
			}
		}

		// Once calculated all angles between station and visible asteroids
		// Delete one by one, clockwise

		asteroidVaporized := true
		for asteroidVaporized {
			asteroidVaporized = false

			minAngle, xToVaporize, yToVaporize := 360.0, -1, -1
			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {
					if angulos[y][x] == -1 {
						continue
					}

					if angulos[y][x] < minAngle {
						minAngle = angulos[y][x]
						xToVaporize, yToVaporize = x, y
						asteroidVaporized = true
					}
				}
			}

			if asteroidVaporized {
				mapa[yToVaporize][xToVaporize] = '.'
				angulos[yToVaporize][xToVaporize] = -1
				asteroidsVaporized++

				if asteroidsVaporized == 200 {
					fmt.Printf("Result part 2: %v\n", xToVaporize*100+yToVaporize)
					break
				}
			}
		}
	}
}

func distanceBetweenAsteroids(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2))
}

func intMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func intMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
