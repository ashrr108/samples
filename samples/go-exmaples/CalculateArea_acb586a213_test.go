package main

import (
	"testing"
)

func TestCalculateArea_acb586a213(t *testing.T) {
	// Test case 1: Normal scenario
	rect1 := &Rectangle{width: 5, height: 4}
	rect1.CalculateArea()
	if rect1.area != 20 {
		t.Errorf("Failed TestCalculateArea_acb586a213: Expected area: %.2f, but got: %.2f", 20.0, rect1.area)
	}

	// Test case 2: Rectangle with width and height as 0
	rect2 := &Rectangle{width: 0, height: 0}
	rect2.CalculateArea()
	if rect2.area != 0 {
		t.Errorf("Failed TestCalculateArea_acb586a213: Expected area: %.2f, but got: %.2f", 0.0, rect2.area)
	}

	// Test case 3: Rectangle with negative width and height
	rect3 := &Rectangle{width: -5, height: -4}
	rect3.CalculateArea()
	if rect3.area != 20 {
		t.Errorf("Failed TestCalculateArea_acb586a213: Expected area: %.2f, but got: %.2f", 20.0, rect3.area)
	}

	// Test case 4: Rectangle with width as positive and height as negative
	rect4 := &Rectangle{width: 5, height: -4}
	rect4.CalculateArea()
	if rect4.area != -20 {
		t.Errorf("Failed TestCalculateArea_acb586a213: Expected area: %.2f, but got: %.2f", -20.0, rect4.area)
	}
}
