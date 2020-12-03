package main

import (
	"fmt"
	"time"

	"./utils"
)

func main() {
	start := time.Now()
	lines := utils.LinesInFile("day03_input.txt")
	droidWeAreLookingFor := "#"
	droidCount := 0
	xIncrement := 3
	curPosX := 0
	boundary := 0

	for index, line := range lines {
		boundary = len(line)
		curPosX = index * xIncrement % boundary

		if string(line[curPosX]) == droidWeAreLookingFor {
			droidCount++
		}
	}

	fmt.Println("Droids found: ", droidCount)
	fmt.Println("Execution time: ", time.Since(start))
}
