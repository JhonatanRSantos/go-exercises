package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	currentSlice := []int{1, 2, 3, 4, 5}
	rotatedSlice := solution(4, currentSlice)
	solutionSlice := []int{5, 1, 2, 3, 4}

	if !reflect.DeepEqual(rotatedSlice, solutionSlice) {
		t.Errorf("Expected %v, but got %v", solutionSlice, rotatedSlice)
	}

	currentSlice = []int{41, 73, 89, 7, 10, 1, 59, 58, 84, 77, 77, 97, 58, 1, 86, 58, 26, 10, 86, 51}
	rotatedSlice = solution(10, currentSlice)
	solutionSlice = []int{77, 97, 58, 1, 86, 58, 26, 10, 86, 51, 41, 73, 89, 7, 10, 1, 59, 58, 84, 77}

	if !reflect.DeepEqual(rotatedSlice, solutionSlice) {
		t.Errorf("Expected %v, but got %v", solutionSlice, rotatedSlice)
	}

	currentSlice = []int{33, 47, 70, 37, 8, 53, 13, 93, 71, 72, 51, 100, 60, 87, 97}
	rotatedSlice = solution(13, currentSlice)
	solutionSlice = []int{87, 97, 33, 47, 70, 37, 8, 53, 13, 93, 71, 72, 51, 100, 60}

	if !reflect.DeepEqual(rotatedSlice, solutionSlice) {
		t.Errorf("Expected %v, but got %v", solutionSlice, rotatedSlice)
	}
}
