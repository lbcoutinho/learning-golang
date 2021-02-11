package main

import (
	"testing"
)

func TestTriangleGetArea(t *testing.T) {
	tri := triangle{height: 10, base: 5}
	expectedArea := 25.
	area := tri.getArea()
	if area != expectedArea {
		t.Errorf("Incorrect triangle area calculation, expected %f, got %f", expectedArea, area)
	}
}

func TestSquareGetArea(t *testing.T) {
	s := square{width: 2}
	expectedArea := 4.
	area := s.getArea()
	if area != expectedArea {
		t.Errorf("Incorrect square area calculation, expected %f, got %f", expectedArea, area)
	}
}