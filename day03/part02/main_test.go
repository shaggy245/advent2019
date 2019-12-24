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

func TestMain(t *testing.T) {
	wire1 := strings.Split("R75,D30,R83,U83,L12,D49,R71,U7,L72", ",")
	wire2 := strings.Split("U62,R66,U55,R34,D71,R55,D58,R83", ",")
	wire1Coords := createCoordinates(wire1, coord{x: 0, y: 0})
	wire2Coords := createCoordinates(wire2, coord{x: 0, y: 0})
	result := findShortest(findIntersectionSteps(wire1Coords, wire2Coords))
	if result != 610 {
		t.Errorf("\nGot  %d\nWant 610", result)
	}

	wire3 := strings.Split("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", ",")
	wire4 := strings.Split("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", ",")
	wire3Coords := createCoordinates(wire3, coord{x: 0, y: 0})
	wire4Coords := createCoordinates(wire4, coord{x: 0, y: 0})
	result = findShortest(findIntersectionSteps(wire3Coords, wire4Coords))
	if result != 410 {
		t.Errorf("\nGot  %d\nWant 410", result)
	}
}
