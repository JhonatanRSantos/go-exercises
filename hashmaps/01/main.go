package main

import (
	"fmt"
	"strings"

	"github.com/JhonatanRSantos/go-exercises/common"
)

func main() {
	magazine := common.ReadNextLine()
	note := common.ReadNextLine()

	fmt.Println(
		solution(magazine, note),
	)
}

func solution(magazine string, note string) string {
	check := true
	magazineWords := map[string]int{}

	for _, v := range strings.Split(magazine, " ") {
		if _, ok := magazineWords[v]; ok {
			magazineWords[v] += 1
		} else {
			magazineWords[v] = 1
		}
	}

	for _, v := range strings.Split(note, " ") {
		if value, ok := magazineWords[v]; ok && value > 0 {
			magazineWords[v] -= 1
		} else {
			check = false
			break
		}
	}

	if check {
		return "Yes"
	} else {
		return "No"
	}
}
