package main

import "testing"

func TestSolution(t *testing.T) {
	currentSlice := []int{2, 1, 5, 3, 4}
	result := solution(currentSlice)

	if result != "3" {
		t.Errorf("Expected %v, but got %v", "3", result)
	}

	currentSlice = []int{2, 5, 1, 3, 4}
	result = solution(currentSlice)

	if result != "Too chaotic" {
		t.Errorf("Expected %v, but got %v", "Too chaotic", result)
	}

	currentSlice = []int{5, 1, 2, 3, 7, 8, 6, 4}
	result = solution(currentSlice)

	if result != "Too chaotic" {
		t.Errorf("Expected %v, but got %v", "Too chaotic", result)
	}

	currentSlice = []int{1, 2, 5, 3, 7, 8, 6, 4}
	result = solution(currentSlice)

	if result != "7" {
		t.Errorf("Expected %v, but got %v", "7", result)
	}
}
