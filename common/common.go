package common

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// Reads the next line
func ReadNextLine() string {
	line, err := reader.ReadString('\n')

	if err != nil {
		panic("Failed to read line")
	}

	return strings.ReplaceAll(line, "\n", "")
}

// Parse a string to int
func ParseInt(value string) int {
	result, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		panic("Failed parsing to int")
	}
	return int(result)
}