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
