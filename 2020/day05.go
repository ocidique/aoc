package main

import (
	"fmt"
	"math"
	"time"

	"./utils"
)

func bsearch(ticket string, i int, min float64, max float64) float64 {
	for _, chunk := range ticket {
		if chunk == 'F' || chunk == 'L' {
			max -= math.Ceil((max - min) / 2)
		} else {
			min += math.Ceil((max - min) / 2)
		}
		if max == min {
			return max
		}
	}
	return bsearch(ticket, i+1, min, max)
}

func main() {
	start := time.Now()
	tickets := utils.LinesInFile("day05_input.txt")
	maxSeatID := 0

	for _, ticket := range tickets {
		row := bsearch(ticket[:7], 0, 0, 127)
		col := bsearch(ticket[7:], 0, 0, 7)
		seatID := int(row)*8 + int(col)
		if maxSeatID < seatID {
			maxSeatID = seatID
		}
	}

	// part one maxSeatID
	fmt.Println("maxSeatID: ", maxSeatID)
	fmt.Println("Execution time: ", time.Since(start))
}
