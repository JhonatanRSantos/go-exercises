package main

import (
	"fmt"
	"strings"

	"github.com/JhonatanRSantos/go-exercices/common"
)

func main() {
	matrix := [6][]int{}

	for index, _ := range matrix {
		sl := common.StringSlice(strings.Split(common.ReadNextLine(), " "))
		matrix[index] = sl.ParseInt()
	}
	fmt.Println(
		solution(matrix[:]),
	)
}

func solution(slice [][]int) int {
	hourglassSum := -100
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			hourglass := 0
			hourglass += slice[row][col] + slice[row][col+1] + slice[row][col+2]
			hourglass += slice[row+1][col+1]
			hourglass += slice[row+2][col] + slice[row+2][col+1] + slice[row+2][col+2]

			if hourglass > hourglassSum {
				hourglassSum = hourglass
			}
		}
	}
	return hourglassSum
}
