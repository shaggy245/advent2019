package main

import "testing"

func Equal(a, b []int) bool {
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

func TestCompute(t *testing.T) {
	result := compute([]int{1, 0, 0, 0, 99})
	if Equal(result, []int{2, 0, 0, 0, 99}) == false {
		t.Errorf("compute([]int{1,0,0,0,99}) = %d; want []int{2,0,0,0,99}", result)
	}
	result = compute([]int{2, 3, 0, 3, 99})
	if Equal(result, []int{2, 3, 0, 6, 99}) == false {
		t.Errorf("compute([]int{2,3,0,3,99}) = %d; want []int{2,3,0,6,99}", result)
	}
	result = compute([]int{2, 4, 4, 5, 99, 0})
	if Equal(result, []int{2, 4, 4, 5, 99, 9801}) == false {
		t.Errorf("compute([]int{2,4,4,5,99,0}) = %d; want []int{2,4,4,5,99,9801}", result)
	}
	result = compute([]int{1, 1, 1, 4, 99, 5, 6, 0, 99})
	if Equal(result, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}) == false {
		t.Errorf("compute([]int{1,1,1,4,99,5,6,0,99}) = %d; want []int{30,1,1,4,2,5,6,0,99}", result)
	}
}

func TestResetMemory(t *testing.T) {
	result := resetMemory(1, 2)
	if Equal(result, []int{1, 1, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 10, 1, 19, 1, 6, 19, 23, 2, 23, 6, 27, 1, 5, 27, 31, 1, 31, 9, 35, 2, 10, 35, 39, 1, 5, 39, 43, 2, 43, 10, 47, 1, 47, 6, 51, 2, 51, 6, 55, 2, 55, 13, 59, 2, 6, 59, 63, 1, 63, 5, 67, 1, 6, 67, 71, 2, 71, 9, 75, 1, 6, 75, 79, 2, 13, 79, 83, 1, 9, 83, 87, 1, 87, 13, 91, 2, 91, 10, 95, 1, 6, 95, 99, 1, 99, 13, 103, 1, 13, 103, 107, 2, 107, 10, 111, 1, 9, 111, 115, 1, 115, 10, 119, 1, 5, 119, 123, 1, 6, 123, 127, 1, 10, 127, 131, 1, 2, 131, 135, 1, 135, 10, 0, 99, 2, 14, 0, 0}) == false {
		t.Errorf("resetMemory(1, 2) = %d, want []int{1, 1, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 10, 1, 19, 1, 6, 19, 23, 2, 23, 6, 27, 1, 5, 27, 31, 1, 31, 9, 35, 2, 10, 35, 39, 1, 5, 39, 43, 2, 43, 10, 47, 1, 47, 6, 51, 2, 51, 6, 55, 2, 55, 13, 59, 2, 6, 59, 63, 1, 63, 5, 67, 1, 6, 67, 71, 2, 71, 9, 75, 1, 6, 75, 79, 2, 13, 79, 83, 1, 9, 83, 87, 1, 87, 13, 91, 2, 91, 10, 95, 1, 6, 95, 99, 1, 99, 13, 103, 1, 13, 103, 107, 2, 107, 10, 111, 1, 9, 111, 115, 1, 115, 10, 119, 1, 5, 119, 123, 1, 6, 123, 127, 1, 10, 127, 131, 1, 2, 131, 135, 1, 135, 10, 0, 99, 2, 14, 0, 0}", result)
	}
}
