package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func linesInFile(fileName string) []string {
	f, err := os.Open(fileName)
	check(err)
	scanner := bufio.NewScanner(f)
	result := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	f.Close()
	return result
}

func strToInt(n string) int {
	x, err := strconv.Atoi(n)
	check(err)
	return x
}

func checkSum(sum int, array []int) bool {
	result := 0
	for _, v := range array {
		result += v
	}
	if result == sum {
		return true
	}
	return false
}

func multiplyNumbers(array []int) int {
	result := 1
	for _, v := range array {
		result = result * v
	}
	return result
}

func main() {
	inputs := linesInFile(`day01_input.txt`)
	sum := 2020
	hasAnswer := false
	for _, line := range inputs {
		for _, lineParaller := range inputs {
			if checkSum(sum, []int{strToInt(line), strToInt(lineParaller)}) {
				fmt.Printf("Answer: %v\n", multiplyNumbers([]int{strToInt(line), strToInt(lineParaller)}))
				break
			}
			for _, lineParaller2 := range inputs {
				if checkSum(sum, []int{strToInt(line), strToInt(lineParaller), strToInt(lineParaller2)}) {
					fmt.Printf("Answer b: %v\n", multiplyNumbers([]int{strToInt(line), strToInt(lineParaller), strToInt(lineParaller2)}))
					hasAnswer = true
					break
				}
			}
		}
		if hasAnswer {
			break
		}
	}
}
