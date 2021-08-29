package main

import (
	"fmt"
	"strings"

	"github.com/JhonatanRSantos/go-exercices/common"
)

func main() {
	fmt.Println(solution())
}

func solution() int {
	common.ParseInt(common.ReadNextLine())
	sliceOfSocks := common.ReadNextLine()

	totalPairs := 0
	socksPairs := map[int]bool{}

	for _, sock := range strings.Split(sliceOfSocks, " ") {
		currentSock := common.ParseInt(sock)
		if _, ok := socksPairs[currentSock]; ok {
			totalPairs++
			delete(socksPairs, currentSock)
		} else {
			socksPairs[currentSock] = true
		}
	}
	return totalPairs
}
