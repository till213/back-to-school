package main

import (
	"testing"
)

func TestCountRectanglesZeroPoints(t *testing.T) {
	points := []point{}
	count := countRectangles(points)
	if count != 0 {
		t.Errorf("Expected 0, received %v", count)
	}
}

func TestCountRectanglesOnePoint(t *testing.T) {
	points := []point{{0, 0}}
	count := countRectangles(points)
	if count != 0 {
		t.Errorf("Expected 0, received %v", count)
	}
}

func TestCountRectanglesTwoPoints(t *testing.T) {
	points := []point{{0, 0}, {0, 10}}
	count := countRectangles(points)
	if count != 0 {
		t.Errorf("Expected 0, received %v", count)
	}
}

func TestCountRectanglesThreePoints(t *testing.T) {
	points := []point{{0, 0}, {0, 10}, {10, 0}}
	count := countRectangles(points)
	if count != 0 {
		t.Errorf("Expected 0, received %v", count)
	}
}

func TestCountRectanglesFourPointsRectangle(t *testing.T) {
	points := []point{{0, 0}, {0, 10}, {10, 0}, {10, 10}}
	count := countRectangles(points)
	if count != 1 {
		t.Errorf("Expected 1, received %v", count)
	}
}

func TestCountRectanglesSixPointsThreeRectangles(t *testing.T) {
	points := []point{{0, 0}, {0, 10}, {10, 0}, {10, 10}, {20, 0}, {20, 10}}
	count := countRectangles(points)
	if count != 3 {
		t.Errorf("Expected 3, received %v", count)
	}
}

func TestCountRectanglesEightPointsTwoRectangles(t *testing.T) {
	points := []point{{0, 0}, {0, 10}, {10, 0}, {10, 10}, {5, 5}, {5, 15}, {15, 5}, {15, 15}}
	count := countRectangles(points)
	if count != 2 {
		t.Errorf("Expected 2, received %v", count)
	}
}

func TestCountRectanglesNinePointsNineRectangles(t *testing.T) {
	points := []point{{0, 0}, {0, 10}, {0, 20}, {10, 0}, {10, 10}, {10, 20}, {20, 0}, {20, 10}, {20, 20}}
	count := countRectangles(points)
	if count != 3 {
		t.Errorf("Expected 9, received %v", count)
	}
}

func TestCountRectanglesFourPointsXAligned(t *testing.T) {
	points := []point{{0, 0}, {0, 10}, {0, 20}, {0, 30}}
	count := countRectangles(points)
	if count != 0 {
		t.Errorf("Expected 0, received %v", count)
	}
}

func TestCountRectanglesFourPointsYAligned(t *testing.T) {
	points := []point{{0, 0}, {10, 0}, {20, 0}, {30, 0}}
	count := countRectangles(points)
	if count != 0 {
		t.Errorf("Expected 0, received %v", count)
	}
}
