package main

import (
	"testing"
)

func TestArea_7903ff2d66(t *testing.T) {
    // Test case 1: Normal scenario
    shape1 := &Shape{area: 25.0}
    if shape1.Area() != 25.0 {
        t.Error(`Test case 1 failed: Expected 25.0 but got `, shape1.Area())
    } else {
        t.Log(`Test case 1 passed`)
    }

    // Test case 2: Zero area
    shape2 := &Shape{area: 0.0}
    if shape2.Area() != 0.0 {
        t.Error(`Test case 2 failed: Expected 0.0 but got `, shape2.Area())
    } else {
        t.Log(`Test case 2 passed`)
    }

    // Test case 3: Negative area
    shape3 := &Shape{area: -5.0}
    if shape3.Area() != -5.0 {
        t.Error(`Test case 3 failed: Expected -5.0 but got `, shape3.Area())
    } else {
        t.Log(`Test case 3 passed`)
    }
}
