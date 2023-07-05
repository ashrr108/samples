package main

import (
	"testing"
)

func TestCalculateArea_8b82ffea81(t *testing.T) {
	// Test case 1: Positive scenario
	sq1 := &Square{side: 5}
	sq1.CalculateArea()
	if sq1.area != 25 {
		t.Errorf("Failed: Expected area to be 25, got %f", sq1.area)
	} else {
		t.Logf("Success: Expected area to be 25, got %f", sq1.area)
	}

	// Test case 2: Scenario when side is zero
	sq2 := &Square{side: 0}
	sq2.CalculateArea()
	if sq2.area != 0 {
		t.Errorf("Failed: Expected area to be 0, got %f", sq2.area)
	} else {
		t.Logf("Success: Expected area to be 0, got %f", sq2.area)
	}

	// Test case 3: Scenario when side is negative
	sq3 := &Square{side: -3}
	sq3.CalculateArea()
	if sq3.area != 9 {
		t.Errorf("Failed: Expected area to be 9, got %f", sq3.area)
	} else {
		t.Logf("Success: Expected area to be 9, got %f", sq3.area)
	}
}
