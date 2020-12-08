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

func removeYesGroup(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	start := time.Now()
	// actually byte[] to string is easier to handle than string[] with some of these data sets
	input, err := ioutil.ReadFile("day06_input.txt")
	utils.Check(err)
	lines := removeEmptyLines(string(input))
	total := 0
	totalPartTwo := 0

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

	for _, line := range lines {
		split := strings.Split(line, "\n")
		chars := split[0]
		split = removeYesGroup(split, 0)
		for _, s := range split {
			for _, c := range chars {
				chunk := string(c)
				if !strings.Contains(s, chunk) {
					chars = strings.Replace(chars, chunk, "", -1)
				}
			}

		}
		totalPartTwo += len(chars)
	}

	fmt.Println("Total: ", total)
	fmt.Println("Total part two: ", totalPartTwo)
	fmt.Println("Execution time: ", time.Since(start))
}
