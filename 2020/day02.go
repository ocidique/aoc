package main

import (
	"fmt"
	"regexp"
	"strings"

	"./utils"
)

func getParams(regEx, url string) (paramsMap map[string]string) {

	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(url)

	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return
}

func countStringOccurences(str, key string) int {
	return strings.Count(str, key)
}

func main() {
	inputs := utils.LinesInFile("day02_input.txt")
	regex := `(?P<min>\d+)-(?P<max>\d+) (?P<key>\w): (?P<pwd>\w+)`
	validPwds := 0

	for _, input := range inputs {
		line := getParams(regex, input)
		if strings.Count(line["pwd"], line["key"]) >= utils.StrToInt(line["min"]) &&
			strings.Count(line["pwd"], line["key"]) <= utils.StrToInt(line["max"]) {
			validPwds++
		}
	}

	fmt.Println("Valid passwords: ", validPwds)
}
