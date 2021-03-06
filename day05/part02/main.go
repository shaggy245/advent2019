package main

// aoc 2019 day 5 part 1
import (
	"fmt"
	"strconv"
	"strings"
)

func compute(intcode []int) []int {
	p := 0
	for p < len(intcode) {
		opcode, paramModes := parseInstruction(intcode[p])
		switch opcode {
		case "01":
			pinc := 4
			params := getParams(paramModes, pinc, intcode, p)
			intcode[intcode[p+3]] = params[0] + params[1]
			p += pinc
		case "02":
			pinc := 4
			params := getParams(paramModes, pinc, intcode, p)
			intcode[intcode[p+3]] = params[0] * params[1]
			p += 4
		case "03":
			intcode[intcode[p+1]] = opcode3()
			p += 2
		case "04":
			fmt.Println(intcode[intcode[p+1]])
			p += 2
		case "05":
			pinc := 3
			params := getParams(paramModes, pinc, intcode, p)
			if params[0] > 0 {
				p = params[1]
			} else {
				p += pinc
			}
		case "06":
			pinc := 3
			params := getParams(paramModes, pinc, intcode, p)
			if params[0] == 0 {
				p = params[1]
			} else {
				p += 3
			}
		case "07":
			pinc := 4
			params := getParams(paramModes, pinc, intcode, p)
			if params[0] < params[1] {
				intcode[intcode[p+3]] = 1
			} else {
				intcode[intcode[p+3]] = 0
			}
			p += pinc
		case "08":
			pinc := 4
			params := getParams(paramModes, pinc, intcode, p)
			if params[0] == params[1] {
				intcode[intcode[p+3]] = 1
			} else {
				intcode[intcode[p+3]] = 0
			}
			p += pinc
		case "99":
			return intcode
		}
	}
	return intcode
}

func parseInstruction(instruction int) (string, []string) {
	// when given code "1020", the last two chars "02" are the opcode, and the remaining
	// chars reversed "01" are the paramModes.
	instring := strconv.Itoa(instruction)
	inChars := strings.Split(instring, "")
	var padding []string
	if len(instring) != 6 {
		for i := 0; i < 6-len(instring); i++ {
			padding = append(padding, "0")
		}
	}
	instrChars := append(padding, inChars...)
	opcode := strings.Join(instrChars[len(instrChars)-2:], "")
	paramModes := instrChars[:len(instrChars)-2]
	// paramModes needs to be a reverse of the input
	for i := len(paramModes)/2 - 1; i >= 0; i-- {
		opp := len(paramModes) - 1 - i
		paramModes[i], paramModes[opp] = paramModes[opp], paramModes[i]
	}
	return opcode, paramModes
}

func getParams(paramModes []string, pinc int, intcode []int, p int) []int {
	var params []int
	for i := 1; i < pinc; i++ {
		if paramModes[i-1] == "0" {
			params = append(params, intcode[intcode[p+i]])
		} else if paramModes[i-1] == "1" {
			params = append(params, intcode[p+i])
		}
	}
	return params
}

func opcode3() int {
	return 5
}

func resetMemory() []int {
	return []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1001, 92, 74, 224, 1001, 224, -85, 224, 4, 224, 1002, 223, 8, 223, 101, 1, 224, 224, 1, 223, 224, 223, 1101, 14, 63, 225, 102, 19, 83, 224, 101, -760, 224, 224, 4, 224, 102, 8, 223, 223, 101, 2, 224, 224, 1, 224, 223, 223, 1101, 21, 23, 224, 1001, 224, -44, 224, 4, 224, 102, 8, 223, 223, 101, 6, 224, 224, 1, 223, 224, 223, 1102, 40, 16, 225, 1102, 6, 15, 225, 1101, 84, 11, 225, 1102, 22, 25, 225, 2, 35, 96, 224, 1001, 224, -350, 224, 4, 224, 102, 8, 223, 223, 101, 6, 224, 224, 1, 223, 224, 223, 1101, 56, 43, 225, 101, 11, 192, 224, 1001, 224, -37, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 4, 224, 1, 223, 224, 223, 1002, 122, 61, 224, 1001, 224, -2623, 224, 4, 224, 1002, 223, 8, 223, 101, 7, 224, 224, 1, 223, 224, 223, 1, 195, 87, 224, 1001, 224, -12, 224, 4, 224, 1002, 223, 8, 223, 101, 5, 224, 224, 1, 223, 224, 223, 1101, 75, 26, 225, 1101, 6, 20, 225, 1102, 26, 60, 224, 101, -1560, 224, 224, 4, 224, 102, 8, 223, 223, 101, 3, 224, 224, 1, 223, 224, 223, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 108, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 329, 1001, 223, 1, 223, 1108, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 344, 101, 1, 223, 223, 7, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 359, 1001, 223, 1, 223, 1007, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 374, 1001, 223, 1, 223, 1108, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 389, 1001, 223, 1, 223, 107, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 404, 101, 1, 223, 223, 1107, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 419, 1001, 223, 1, 223, 1007, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 434, 101, 1, 223, 223, 1107, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 449, 101, 1, 223, 223, 107, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 464, 1001, 223, 1, 223, 1008, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 479, 101, 1, 223, 223, 1007, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 494, 1001, 223, 1, 223, 8, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 509, 1001, 223, 1, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 524, 1001, 223, 1, 223, 1008, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 539, 1001, 223, 1, 223, 7, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 554, 101, 1, 223, 223, 1108, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 569, 101, 1, 223, 223, 107, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 584, 101, 1, 223, 223, 8, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 599, 101, 1, 223, 223, 108, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 614, 1001, 223, 1, 223, 7, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 629, 1001, 223, 1, 223, 1107, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 644, 101, 1, 223, 223, 8, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 659, 1001, 223, 1, 223, 1008, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 674, 1001, 223, 1, 223, 4, 223, 99, 226}
}

func guesser() int {
	output := compute(resetMemory())[0]
	return output
}

func main() {
	guesser()
}
