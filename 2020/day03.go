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
	curPosX := 0
	boundary := 0
	droidsTotal := 1

	// using anonymous struct
	slopes := []struct {
		x, y int
	}{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	for _, slope := range slopes {
		droidCount = 0
		for index, line := range lines {
			boundary = len(line)

			// skip line in case of index and yIncrement remainder > 0
			if index%slope.y > 0 {
				continue
			}

			curPosX = index * slope.x / slope.y % boundary
			if string(line[curPosX]) == droidWeAreLookingFor {
				droidCount++
			}
		}
		fmt.Println("Droids on a slope: ", slope, droidCount)
		droidsTotal *= droidCount

	}

	fmt.Println("Droids total: ", droidsTotal)
	fmt.Println("Execution time: ", time.Since(start))
}
