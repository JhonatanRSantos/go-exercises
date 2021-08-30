package main

import (
	"fmt"
	"strings"

	"github.com/JhonatanRSantos/go-exercices/common"
)

func main() {
	rotations := common.ParseInt(
		common.StringSlice(
			strings.Split(common.ReadNextLine(), " "),
		)[1],
	)

	slice := common.StringSlice(
		strings.Split(common.ReadNextLine(), " "),
	).ParseInt()

	fmt.Println(
		solution(rotations, slice),
	)
}

func solution(rotations int, slice []int) []int {
	return append(slice[rotations:], slice[0:rotations]...)
}
