package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"./utils"
)

func getParams(regEx, input string) (paramsMap map[string]string) {

	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(input)

	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return
}

func main() {
	start := time.Now()
	inputs := utils.LinesInFile("day02_input.txt")
	regex := `(?P<min>\d+)-(?P<max>\d+) (?P<key>\w): (?P<pwd>\w+)`
	validPwds := 0
	validPwdsPartTwo := 0

	for _, input := range inputs {
		line := getParams(regex, input)
		min := utils.StrToInt(line["min"])
		max := utils.StrToInt(line["max"])

		// part one check if key is between min max values
		if strings.Count(line["pwd"], line["key"]) >= min &&
			strings.Count(line["pwd"], line["key"]) <= max {
			validPwds++
		}

		// part two check if key is either min-1 or max-1 index position, but not on both
		if (string(line["pwd"][min-1]) == line["key"] && !(string(line["pwd"][max-1]) == line["key"])) || (!(string(line["pwd"][min-1]) == line["key"]) && string(line["pwd"][max-1]) == line["key"]) {
			validPwdsPartTwo++
		}
	}

	fmt.Println("Valid passwords: ", validPwds)
	fmt.Println("Valid passwords part two: ", validPwdsPartTwo)
	fmt.Println("Execution time: ", time.Since(start))
}
