package main

import (
	"fmt"
	"math"
	"time"

	"./utils"
)

func bSearch(ticket string, i int, min float64, max float64) float64 {
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
	return bSearch(ticket, i+1, min, max)
}

func hasTicket(seatID int, takenSeats []int) bool {
	for _, ticketID := range takenSeats {
		if seatID == ticketID {
			return true
		}
	}
	return false
}

func main() {
	start := time.Now()
	tickets := utils.LinesInFile("day05_input.txt")
	maxSeatID := 0
	takenSeats := []int{}

	for _, ticket := range tickets {
		row := bSearch(ticket[:7], 0, 0, 127)
		col := bSearch(ticket[7:], 0, 0, 7)
		seatID := int(row)*8 + int(col)
		if maxSeatID < seatID {
			maxSeatID = seatID
		}
		takenSeats = append(takenSeats, seatID)
	}
	fmt.Println("Part one, max seatID: ", maxSeatID)

	for i := 0; i < len(takenSeats); i++ {
		if x, y, z := hasTicket(i, takenSeats), hasTicket(i-1, takenSeats), hasTicket(i+1, takenSeats); !x && y && z {
			fmt.Println("Part two, my seat: ", i)
		}
	}
	fmt.Println("Execution time: ", time.Since(start))
}
