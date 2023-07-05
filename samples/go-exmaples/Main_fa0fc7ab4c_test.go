package main

import (
	"testing"
)

// TestMain_fa0fc7ab4c tests the area calculation of Rectangle and Square
func TestMain_fa0fc7ab4c(t *testing.T) {
	// Test case for Rectangle
	rect := &Rectangle{width: 16, height: 6}
	rect.CalculateArea()
	if rect.Area() != 96 {
		t.Error("Expected Area of Rectangle: 96, but got ", rect.Area())
	} else {
		t.Log("Success: Expected Area of Rectangle: 96 and got ", rect.Area())
	}

	// Test case for Square
	sq := &Square{side: 8}
	sq.CalculateArea()
	if sq.Area() != 64 {
		t.Error("Expected Area of Square: 64, but got ", sq.Area())
	} else {
		t.Log("Success: Expected Area of Square: 64 and got ", sq.Area())
	}
}
