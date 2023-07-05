```package main

import (
	"testing"
)

// Shape is a struct
type Shape struct {
	length float64
	breadth float64
}

// Area returns area of square
func (sq *Shape) Area() float64 {
	return sq.length * sq.breadth
}

// TestArea_7903ff2d66 tests Area function
func TestArea_7903ff2d66(t *testing.T) {
	// create a new square
	sq := Shape{length: 10, breadth: 10}

	// expected area of square
	expectedArea := 100.0

	// actual area of square
	actualArea := sq.Area()

	// compare expected and actual area
	if expectedArea != actualArea {
		t.Errorf("Expected area %f, got %f", expectedArea, actualArea)
	}
}
```