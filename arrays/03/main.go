package main

import (
	"fmt"
	"strings"

	"github.com/JhonatanRSantos/go-exercises/common"
)

func main() {
	totalTests := common.StringSlice(
		strings.Split(common.ReadNextLine(), " "),
	)[0]

	for i := 0; i < common.ParseInt(totalTests); i++ {
		slice := common.StringSlice(
			strings.Split(common.ReadNextLine(), " "),
		).ParseInt()

		fmt.Println(solution(slice))
	}
}

func solution(slice []int) string {
	count := 0
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] != (i + 1) {
			if (i-1) >= 0 && slice[i-1] == (i+1) {
				count++
				slice[i], slice[i-1] = slice[i-1], slice[i]
			} else if (i-2) >= 0 && slice[i-2] == (i+1) {
				count += 2
				slice[i-2] = slice[i-1]
				slice[i-1] = slice[i]
				slice[i] = i + 1
			} else {
				count = -1
				break
			}
		}
	}

	if count == -1 {
		return "Too chaotic"
	} else {
		return fmt.Sprint(count)
	}
}
