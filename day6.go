package main

import (
	"fmt"
	"strings"
)

type planet struct {
	name         string
	closePlanets []*planet
}

var orbits map[string]string = make(map[string]string)

func day6A() {
	lines := readFile(6)
	for _, line := range lines {
		planets := strings.Split(line, ")")
		orbits[planets[1]] = planets[0]
	}

	totalOrbits := 0
	for _, value := range orbits {
		v := value
		for {
			totalOrbits++
			if x, found := orbits[v]; found {
				v = x
			} else {
				break
			}
		}
	}

	fmt.Printf("Result part 1: %v\n", totalOrbits)
}

func day6B() {
	// I know should be solved with graphs
	lines := readFile(6)
	for _, line := range lines {
		planets := strings.Split(line, ")")
		orbits[planets[1]] = planets[0]
	}

	santaOrbits := make([]string, 0)
	var v string = orbits["SAN"]
	for {
		santaOrbits = append(santaOrbits, v)
		if x, found := orbits[v]; found {
			v = x
		} else {
			break
		}
	}

	myOrbits := make([]string, 0)
	v = orbits["YOU"]
	for {
		myOrbits = append(myOrbits, v)
		if x, found := orbits[v]; found {
			v = x
		} else {
			break
		}
	}

	var commonAncestor string
	exitLoop := false
	for _, planet1 := range myOrbits {
		for _, planet2 := range santaOrbits {
			if planet1 == planet2 {
				commonAncestor = planet1
				exitLoop = true
				break
			}
		}
		if exitLoop {
			break
		}
	}

	distanceFromMe, _ := Find(myOrbits, commonAncestor)
	distanceFromSanta, _ := Find(santaOrbits, commonAncestor)
	fmt.Printf("Result part 2: %v\n", distanceFromMe+distanceFromSanta)
}
