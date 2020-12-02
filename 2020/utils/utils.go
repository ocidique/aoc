package utils

import (
	"bufio"
	"os"
	"strconv"
)

// Check error helper
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// StrToInt for converting string number to int
func StrToInt(n string) int {
	x, err := strconv.Atoi(n)
	Check(err)
	return x
}

// LinesInFile Open file and return string array
func LinesInFile(fileName string) []string {
	f, err := os.Open(fileName)
	Check(err)

	scanner := bufio.NewScanner(f)
	result := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	f.Close()
	return result
}
