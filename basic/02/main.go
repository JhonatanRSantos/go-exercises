package main

import (
	"fmt"

	"github.com/JhonatanRSantos/go-exercices/common"
)

func main() {
	fmt.Println(solution())
}

func solution() int {
	common.ReadNextLine()
	rawPath := common.ReadNextLine()

	valleys := 0
	altitude := 0
	canAddValley := true

	for _, p := range rawPath {
		if string(p) == "D" {
			altitude--
		} else {
			altitude++
		}

		if canAddValley && altitude < 0 {
			valleys++
			canAddValley = false
		} else if altitude >= 0 {
			canAddValley = true
		}
	}
	return valleys
}
