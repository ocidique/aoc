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
  f, _ := os.Open(fileName)
  scanner := bufio.NewScanner(f)
  result := []string{}

  for scanner.Scan() {
    line := scanner.Text()
    result = append(result, line)
  }
  return result
}

func strToInt(n string) int {
  x, err := strconv.Atoi(n)
  check(err)	
  return x
}

func checkSum(sum int, n1 string, n2 string) bool {
  if strToInt(n1) + strToInt(n2) == sum {
    return true
  }
  return false
}

func multiplyStrLn(n1 string, n2 string) int {
  return strToInt(n1) * strToInt(n2)
}


func main() {

  inputs := linesInFile(`day01_input.txt`)
  sum := 2020
  hasAnswer := false
  for _, line := range inputs {
    for _, lineParaller := range inputs {
      if checkSum(sum, line, lineParaller) {
        fmt.Printf("Answer: %v", multiplyStrLn(line, lineParaller))
        hasAnswer = true
        break
      }
    }
    if hasAnswer {
      break
    }
  }
}