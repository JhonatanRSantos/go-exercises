package main

import (
	"testing"

	"github.com/JhonatanRSantos/go-exercises/common"
)

func TestSolution(t *testing.T) {
	response := solution("aaaabbcc")
	if response != "NO" {
		t.Errorf("Expected %v but got %v", "NO", response)
	}

	response = solution("aabbccddeefghi")
	if response != "NO" {
		t.Errorf("Expected %v but got %v", "NO", response)
	}

	response = solution("aabbccddeefghi")
	if response != "NO" {
		t.Errorf("Expected %v but got %v", "NO", response)
	}

	response = solution("a")
	if response != "YES" {
		t.Errorf("Expected %v but got %v", "YES	", response)
	}

	response = solution(string(common.ReadFile("./mock/test_input_01.txt")))
	if response != "YES" {
		t.Errorf("Expected %v but got %v", "YES	", response)
	}
}
