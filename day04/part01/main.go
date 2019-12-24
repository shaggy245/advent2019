package main

import (
	"fmt"
	"strconv"
)

func isPassword(guess string) bool {
	//guessChars := strings.Split(guess, "")
	consec := false
	for i := 1; i < len(guess); i++ {
		//curr, _ := strconv.Atoi(guessChars[i])
		//prev, _ := strconv.Atoi(guessChars[i-1])
		if consec == false && guess[i-1] == guess[i] {
			consec = true
		}
		if guess[i-1] > guess[i] {
			return false
		}
	}
	if consec == false {
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
