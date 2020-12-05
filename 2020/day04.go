package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"./utils"
)

func main() {
	start := time.Now()
	papersPlease := utils.LinesInFile("day04_input.txt")
	validPassports := 0
	validPassportFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	// declare list of papers
	passports := []map[string]string{}
	// current papers
	p := make(map[string]string)

	for _, papers := range papersPlease {
		// if empty line, papers please
		if len(papers) > 0 {
			papersLines := strings.Fields(papers)
			for _, papersLine := range papersLines {
				paperField := strings.Split(papersLine, ":")
				p[paperField[0]] = paperField[1]
			}
		} else {
			// when empty line we append current papers to passports and clear current papers
			passports = append(passports, p)
			p = make(map[string]string)
		}
	}

	// append last missing line because empty lines are out in the loop
	passports = append(passports, p)

	for _, k := range passports {
		validPapers := true

		for _, validPassportField := range validPassportFields {
			// use golang special condition and initalization statement
			if _, ok := k[validPassportField]; !ok {
				validPapers = false
				continue
			}

			switch validPassportField {
			case "byr":
				// byr (Birth Year) - four digits; at least 1920 and at most 2002.
				byrValue := utils.StrToInt(k[validPassportField])
				if !(byrValue >= 1920 && byrValue <= 2002) {
					validPapers = false
					break
				}

			case "iyr":
				// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
				iyrValue := utils.StrToInt(k[validPassportField])
				if !(iyrValue >= 2010 && iyrValue <= 2020) {
					validPapers = false
					break
				}

			case "eyr":
				// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
				eyrValue := utils.StrToInt(k[validPassportField])
				if !(eyrValue >= 2020 && eyrValue <= 2030) {
					validPapers = false
					break
				}

			case "hgt":
				// hgt (Height) - a number followed by either cm or in:
				// If cm, the number must be at least 150 and at most 193.
				// If in, the number must be at least 59 and at most 76.
				hgtNumber := regexp.MustCompile("[0-9]+")
				hgtUnit := regexp.MustCompile("(cm|in)")
				numStr := hgtNumber.FindString(k[validPassportField])
				unit := hgtUnit.FindString(k[validPassportField])
				num := utils.StrToInt(numStr)

				if unit == "cm" {
					if !(num >= 150 && num <= 193) {
						validPapers = false
						break
					}

				} else if unit == "in" {
					if !(num >= 59 && num <= 76) {
						validPapers = false
						break
					}
				} else {
					validPapers = false
					break
				}

			case "hcl":
				// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
				hcl := regexp.MustCompile("^#([a-f|0-9]{6})$")
				hclValid := hcl.MatchString(k[validPassportField])
				if !hclValid {
					validPapers = false
					break
				}

			case "ecl":
				// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
				validEyeColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
				foundValidEyeColor := false
				for _, validEyeColor := range validEyeColors {
					if validEyeColor == k[validPassportField] {
						foundValidEyeColor = true
						continue
					}
				}

				if !foundValidEyeColor {
					validPapers = false
					break
				}

			case "pid":
				// pid (Passport ID) - a nine-digit number, including leading zeroes.
				pid := regexp.MustCompile("^([0-9]{9})$")
				pidValid := pid.MatchString(k[validPassportField])

				if !pidValid {
					validPapers = false
					break
				}
			}

		}

		if validPapers {
			validPassports++
		}
	}

	fmt.Println("Valid passports: ", validPassports)
	fmt.Println("Execution time: ", time.Since(start))
}
