package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(day int) []string {
	file, err := os.Open(fmt.Sprintf("./inputs/%d.txt", day))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// This is our buffer now
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func manhattanDistance(x1, y1, x2, y2 int) int {
	dy := y2 - y1
	if dy < 0 {
		dy = -dy
	}

	dx := x2 - x1
	if dx < 0 {
		dx = -dx
	}

	return dx + dy
}

func customMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func customMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func createFile(path string) {
	// delete file
	err := os.Remove(path)
	if err != nil {
		fmt.Println(err.Error())
	}

	// create file
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()
}

func writeFile(path, text string) {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	// write some text line-by-line to file
	_, err = file.WriteString(fmt.Sprintf("%v", text))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// save changes
	err = file.Sync()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func permutations(arr []int) [][]int {
	result := [][]int{}

	var helper func([]int, int)
	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			result = append(result, tmp)

		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}

	helper(arr, len(arr))
	return result
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
