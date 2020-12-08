package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"./utils"
)

func removeEmptyLines(input string) []string {
	return strings.Split(input, "\n\n")
}

func contains(chars []string, char string) bool {
	for _, chunk := range chars {
		if chunk == char {
			return true
		}
	}
	return false
}

func main() {
	start := time.Now()
	// actually byte[] to string is easier to handle than string[] with some of these data sets
	input, err := ioutil.ReadFile("day06_input.txt")
	utils.Check(err)
	lines := removeEmptyLines(string(input))
	total := 0

	for _, line := range lines {
		line = strings.Replace(line, "\n", "", -1)
		chars := []string{}
		for _, l := range line {
			chunk := string(l)
			if !contains(chars, chunk) {
				chars = append(chars, chunk)
			}
		}
		total += len(chars)
	}

	fmt.Println("Total: ", total)
	fmt.Println("Execution time: ", time.Since(start))
}
