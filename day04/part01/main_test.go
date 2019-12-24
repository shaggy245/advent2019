package main

import (
	"strconv"
	"testing"
)

func TestIsPassword(t *testing.T) {
	result := isPassword(strconv.Itoa(111111))
	if result != true {
		t.Errorf("For 111111, got %t, but want %t", result, true)
	}
	result = isPassword(strconv.Itoa(223450))
	if result != false {
		t.Errorf("For 223450, got %t, but want %t", result, false)
	}
	result = isPassword(strconv.Itoa(123789))
	if result != false {
		t.Errorf("For 123789, got %t, but want %t", result, false)
	}
}
