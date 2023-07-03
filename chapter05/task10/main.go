package main

import (
	"fmt"
)

// prereqs отображает каждый курс на список курсов, которые
// должны быть прочитаны раньше него.
var prereqs = map[string]map[string]struct{}{
	"algorithms": {"data structures": {}},
	"calculus":   {"linear algebra": {}},
	"compilers": {
		"data structures":       {},
		"formal languages":      {},
		"computer organization": {},
	},
	"data structures":  {"discrete math": {}},
	"databases":        {"data structures": {}},
	"discrete math":    {"intro to programming": {}},
	"formal languages": {"discrete math": {}},
	"networks":         {"operating systems": {}},
	"operating systems": {"data structures": {},
		"computer organization": {}},
	"programming languages": {"data structures": {},
		"computer organization": {}},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string]map[string]struct{}) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string]struct{})
	visitAll = func(items map[string]struct{}) {
		for key := range items {
			if !seen[key] {
				seen[key] = true
				visitAll(m[key])
				order = append(order, key)
			}
		}
	}

	for key := range m {
		if !seen[key] {
			seen[key] = true
			visitAll(m[key])
			order = append(order, key)
		}
	}
	return order
}
