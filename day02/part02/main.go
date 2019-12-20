package main

import (
	"fmt"
)

// aoc 2019 day 2 part 2
func compute(intcode []int) []int {
	for i, code := range intcode {
		if i%4 == 0 {
			switch code {
			case 1:
				intcode[intcode[i+3]] = intcode[intcode[i+1]] + intcode[intcode[i+2]]
			case 2:
				intcode[intcode[i+3]] = intcode[intcode[i+1]] * intcode[intcode[i+2]]
			case 99:
				break
			}
		}
	}
	return intcode
}

func resetMemory(noun int, verb int) []int {
	return []int{1, noun, verb, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 10, 1, 19, 1, 6, 19, 23, 2, 23, 6, 27, 1, 5, 27, 31, 1, 31, 9, 35, 2, 10, 35, 39, 1, 5, 39, 43, 2, 43, 10, 47, 1, 47, 6, 51, 2, 51, 6, 55, 2, 55, 13, 59, 2, 6, 59, 63, 1, 63, 5, 67, 1, 6, 67, 71, 2, 71, 9, 75, 1, 6, 75, 79, 2, 13, 79, 83, 1, 9, 83, 87, 1, 87, 13, 91, 2, 91, 10, 95, 1, 6, 95, 99, 1, 99, 13, 103, 1, 13, 103, 107, 2, 107, 10, 111, 1, 9, 111, 115, 1, 115, 10, 119, 1, 5, 119, 123, 1, 6, 123, 127, 1, 10, 127, 131, 1, 2, 131, 135, 1, 135, 10, 0, 99, 2, 14, 0, 0}
}

func guesser() int {
	for n := 0; n <= 90; n++ {
		for v := 0; v <= 90; v++ {
			output := compute(resetMemory(n, v))[0]
			if output == 19690720 {
				return 100*n + v
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(guesser())
}
