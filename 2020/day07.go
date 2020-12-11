package main

import (
	"fmt"
	"strings"
	"time"

	"./utils"
)

type luggageRule struct {
	color    string
	contains map[string]int
}

type luggageGraph struct {
	colors map[string]luggageRule
}

func main() {
	start := time.Now()
	luggages := utils.LinesInFile("day07_input.txt")
	graph := luggageGraph{colors: map[string]luggageRule{}}

	for _, luggage := range luggages {
		// remove useless stuff
		luggage = luggage[:len(luggage)-1]
		replacer := strings.NewReplacer(" bags", "", " bag", "")
		luggage = replacer.Replace(luggage)

		parts := strings.Split(luggage, " contain ")
		if strings.Contains(parts[1], "no other") {
			rule := luggageRule{color: parts[0], contains: map[string]int{}}
			graph.colors[parts[0]] = rule
		} else {
			children := strings.Split(parts[1], ", ")
			childCounts := make([]int, len(children))
			childColors := make([]string, len(children))

			for i, child := range children {
				// split child into count of bag type and bag type name
				split := strings.SplitN(child, " ", 2)
				// count of bag type
				childCounts[i] = utils.StrToInt(split[0])
				// color of bag type
				childColors[i] = split[1]
			}

			rule := luggageRule{color: parts[0], contains: map[string]int{}}
			for i, childColor := range childColors {
				rule.contains[childColor] = childCounts[i]
			}
			// fmt.Println(rule)
			graph.colors[parts[0]] = rule
		}
	}

	for _, g := range graph.colors {
		// fmt.Println(g)
		if value, ok := g.contains["shiny gold"]; ok {
			fmt.Println("color:", g.color)
			fmt.Println("rule:", g.contains)
			fmt.Println("shiny gold bags: ", value)
		}
	}

	fmt.Println("Execution time: ", time.Since(start))
}
