package main

import (
	"fmt"
	"strconv"
	"strings"
)

type chemical struct {
	Name   string
	Amount int
}

type chemicalRequirement struct {
	Needed   int
	Produced int
}

func day14A() {
	entrada := readFile(14)

	reactions := make(map[string][]chemical)
	chemicals := make(map[string]*chemicalRequirement)

	for _, reaction := range entrada {
		r := strings.Split(reaction, " => ")
		outputChemical := newChemical(r[1])
		if _, ok := reactions[outputChemical.Name]; !ok {
			reactions[outputChemical.Name] = make([]chemical, 0)
			reactions[outputChemical.Name] = append(reactions[outputChemical.Name], outputChemical)
		}

		for _, inputChemical := range strings.Split(r[0], ", ") {
			reactions[outputChemical.Name] = append(reactions[outputChemical.Name], newChemical(inputChemical))
		}

		if outputChemical.Name == "FUEL" {
			chemicals[outputChemical.Name] = &chemicalRequirement{1, 0}
		} else {
			chemicals[outputChemical.Name] = &chemicalRequirement{0, 0}
		}
	}

	chemicals["ORE"] = &chemicalRequirement{0, 1<<(32<<(^uint(0)>>32&1)-1) - 1}

	reactionsNeeded := true
	for reactionsNeeded {
		reactionsNeeded = false
		for k, v := range chemicals {
			if v.needsReaction() {
				reactionsNeeded = true

				// Get reaction that produces chemical
				outputChemical := reactions[k][0]
				inputChemicals := reactions[k][1:]

				o := chemicals[k]
				o.Produced += outputChemical.Amount

				for _, inputNeeded := range inputChemicals {
					ch := chemicals[inputNeeded.Name]
					ch.Needed += inputNeeded.Amount
				}

				break
			}
		}
	}

	fmt.Printf("Result part 1: %v\n", chemicals["ORE"].Needed)
}

func day14B() {
	entrada := readFile(14)

	reactions := make(map[string][]chemical)
	chemicals := make(map[string]*chemicalRequirement)

	for _, reaction := range entrada {
		r := strings.Split(reaction, " => ")
		outputChemical := newChemical(r[1])
		if _, ok := reactions[outputChemical.Name]; !ok {
			reactions[outputChemical.Name] = make([]chemical, 0)
			reactions[outputChemical.Name] = append(reactions[outputChemical.Name], outputChemical)
		}

		for _, inputChemical := range strings.Split(r[0], ", ") {
			reactions[outputChemical.Name] = append(reactions[outputChemical.Name], newChemical(inputChemical))
		}

		chemicals[outputChemical.Name] = &chemicalRequirement{0, 0}
	}

	chemicals["ORE"] = &chemicalRequirement{0, 1000000000000}
	// Min amount we can produce, based on answer from part 1 (ORE-per-FUEL)
	chemicals["FUEL"].Needed = 1000000000000 / 1037742

	var overLimit bool

	for !overLimit {
		chemicals["FUEL"].Needed++

		reactionsNeeded := true
		for reactionsNeeded {
			reactionsNeeded = false
			for k, v := range chemicals {
				if v.needsReaction() {
					if k == "ORE" {
						overLimit = true
						break
					}
					reactionsNeeded = true

					// Get reaction that produces chemical
					outputChemical := reactions[k][0]
					inputChemicals := reactions[k][1:]

					// Calculate how many times this reaction needs to happen
					timesReaction := (v.Needed - v.Produced) / outputChemical.Amount
					if (v.Needed-v.Produced)%outputChemical.Amount != 0 {
						timesReaction++
					}

					// Output chemicals
					o := chemicals[k]
					o.Produced += outputChemical.Amount * timesReaction

					// Input chemicals
					for _, inputNeeded := range inputChemicals {
						ch := chemicals[inputNeeded.Name]
						ch.Needed += inputNeeded.Amount * timesReaction
					}

					break
				}
			}
		}
	}

	fmt.Printf("Result part 2: %v\n", chemicals["FUEL"].Produced-1)
}

func newChemical(s string) chemical {
	name := strings.Split(s, " ")[1]
	amount, _ := strconv.Atoi(strings.Split(s, " ")[0])
	return chemical{name, amount}
}

func (cr chemicalRequirement) needsReaction() bool {
	return cr.Needed > cr.Produced
}
