package main

import (
	"fmt"
	"math"

	"github.com/JhonatanRSantos/go-exercises/common"
)

func main() {
	fmt.Println(
		solution(
			common.ReadNextLine(),
		),
	)
}

func solution(s string) string {
	letterFrequency := map[string]int{}

	groupA, groupB := "", ""
	canRemoveLetter := true

	for _, v := range s {
		if _, ok := letterFrequency[string(v)]; ok {
			letterFrequency[string(v)] += 1
		} else {
			if groupA == "" {
				groupA = string(v)
			} else if groupB == "" {
				groupB = string(v)
			}

			letterFrequency[string(v)] = 1
		}
	}

	groupDiff := math.Abs(float64(letterFrequency[groupA] - letterFrequency[groupB]))

	if groupDiff > 1 {
		return "NO"
	}

	if groupDiff == 1 && canRemoveLetter {
		canRemoveLetter = false
	}

	for _, v := range letterFrequency {
		if canRemoveLetter && letterFrequency[groupA] != v && letterFrequency[groupB] != v {
			canRemoveLetter = false
		} else if !canRemoveLetter && letterFrequency[groupA] != v && letterFrequency[groupB] != v {
			return "NO"
		}
	}
	return "YES"
}
