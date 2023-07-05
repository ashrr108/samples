package main

import (
	"fmt"
	"testing"
)

type Rectangle struct {
	width  float64
	height float64
	area   float64
}

func (rect *Rectangle) CalculateArea() {
	rect.area = rect.width * rect.height //area of rectangle
}

func TestCalculateArea_acb586a213(t *testing.T) {
	// TODO: Add test cases
	rect := Rectangle{width: 10, height: 5}
	rect.CalculateArea()
	if rect.area != 50 {
		t.Error("CalculateArea() failed. Expected 50, got", rect.area)
	}
}