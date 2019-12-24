package main

import (
	"fmt"
	"strconv"
)

func isPassword(guess string) bool {
	consec := false
	//finalConsec := false
	charsInRow := 1
	for i := 1; i < len(guess); i++ {
		if guess[i-1] == guess[i] {
			charsInRow++
		} else if charsInRow == 2 && guess[i-1] != guess[i] {
			consec = true
			charsInRow = 1
		} else {
			charsInRow = 1
		}
		if guess[i-1] > guess[i] {
			return false
		}
	}
	if consec == false && charsInRow != 2 {
		return false
	}
	return true
}

func main() {
	minIn := 265275
	maxIn := 781584
	if minIn < 100000 {
		minIn = 100000
	}
	if maxIn > 999999 {
		maxIn = 999999
	}
	passwordCnt := 0
	for i := minIn; i <= maxIn; i++ {
		if isPassword(strconv.Itoa(i)) {
			passwordCnt++
		}
	}
	fmt.Println(passwordCnt)
}
