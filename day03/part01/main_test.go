package main

import (
	"strings"
	"testing"
)

func Equal(a, b []coord) bool {
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

func TestCreateCoordinates(t *testing.T) {
	//wire1 := strings.Split("R75,D30,R83,U83,L12,D49,R71,U7,L72", ",")
	//wire2 := strings.Split("U62,R66,U55,R34,D71,R55,D58,R83", ",")
	wire := strings.Split("R4,U3,D4,L1", ",")
	var answer = []coord{
		coord{
			x: 1,
			y: 0,
		},
		coord{
			x: 2,
			y: 0,
		},
		coord{
			x: 3,
			y: 0,
		},
		coord{
			x: 4,
			y: 0,
		},
		coord{
			x: 4,
			y: 1,
		},
		coord{
			x: 4,
			y: 2,
		},
		coord{
			x: 4,
			y: 3,
		},
		coord{
			x: 4,
			y: 2,
		},
		coord{
			x: 4,
			y: 1,
		},
		coord{
			x: 4,
			y: 0,
		},
		coord{
			x: 4,
			y: -1,
		},
		coord{
			x: 3,
			y: -1,
		},
	}
	result := createCoordinates(wire, coord{x: 0, y: 0})
	if Equal(result, answer) == false {
		t.Errorf("\nGot  %d\nWant %d", result, answer)
	}
}
