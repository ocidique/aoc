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

// depth-first search for bag occurrence
func bagOccurrenceDFS(graph luggageGraph, rule luggageRule, bag string, visited, containsBag map[string]bool) bool {
	if rule.color == bag || containsBag[rule.color] {
		containsBag[rule.color] = true
		return true
	} else if visited[rule.color] {
		return false
	}

	visited[rule.color] = true

	for childColor := range rule.contains {
		childRule := graph.colors[childColor]
		found := bagOccurrenceDFS(graph, childRule, bag, visited, containsBag)

		if found {
			containsBag[rule.color] = true
			return true
		}
	}
	return false
}

// depth-first search for counting individual bags inside given bag
func bagCountDFS(graph luggageGraph, bag string, bagTypeCounts map[string]int) int {
	if len(graph.colors[bag].contains) == 0 {
		return 0
	} else if bagTypeCounts[bag] != 0 {
		return bagTypeCounts[bag]
	}

	total := 0
	for color, count := range graph.colors[bag].contains {
		subCount := bagCountDFS(graph, color, bagTypeCounts)
		bagTypeCounts[color] = subCount
		total += count + count*bagCountDFS(graph, color, bagTypeCounts)
	}
	return total
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

			graph.colors[parts[0]] = rule
		}
	}

	bag := "shiny gold"
	visited := map[string]bool{}
	containsBag := map[string]bool{}
	for _, rule := range graph.colors {
		bagOccurrenceDFS(graph, rule, bag, visited, containsBag)
	}

	count := 0
	for c := range containsBag {
		if c != bag {
			count++
		}
	}

	total := bagCountDFS(graph, bag, map[string]int{})

	fmt.Println("Part 1: ", count)
	fmt.Println("Part 2: ", total)
	fmt.Println("Execution time: ", time.Since(start))
}
