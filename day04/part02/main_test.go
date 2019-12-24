package main

import (
	"strconv"
	"testing"
)

func TestIsPassword(t *testing.T) {
	result := isPassword(strconv.Itoa(112233))
	if result != true {
		t.Errorf("For 112233, got %t, but want %t", result, true)
	}
	result = isPassword(strconv.Itoa(123444))
	if result != false {
		t.Errorf("For 123444, got %t, but want %t", result, false)
	}
	result = isPassword(strconv.Itoa(111122))
	if result != true {
		t.Errorf("For 111122, got %t, but want %t", result, true)
	}
	result = isPassword(strconv.Itoa(114433))
	if result != false {
		t.Errorf("For 114433, got %t, but want %t", result, false)
	}
}
