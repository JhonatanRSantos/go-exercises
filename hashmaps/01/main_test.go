package main

import (
	"strings"
	"testing"

	"github.com/JhonatanRSantos/go-exercices/common"
)

func TestSolution(t *testing.T) {
	input := strings.Split(
		string(common.ReadFile("./mock/test_input_01.txt")), "\n",
	)

	result := solution(input[1], input[2])

	if result != "Yes" {
		t.Errorf("Expected %v, but got %v", "Yes", result)
	}

	input = strings.Split(
		string(common.ReadFile("./mock/test_input_02.txt")), "\n",
	)

	result = solution(input[1], input[2])

	if result != "No" {
		t.Errorf("Expected %v, but got %v", "No", result)
	}
}
