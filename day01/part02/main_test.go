package main

import "testing"

func TestCalculateFuel(t *testing.T) {
	result := calculateFuel([]int{14})
	if result != 2 {
		t.Errorf("calculateFuel([14]) = %d; want 2", result)
	}
	result = calculateFuel([]int{1969})
	if result != 966 {
		t.Errorf("calculateFuel([14]) = %d; want 966", result)
	}
	result = calculateFuel([]int{100756})
	if result != 50346 {
		t.Errorf("calculateFuel([14]) = %d; want 50346", result)
	}
}
