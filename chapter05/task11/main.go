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
	"linear algebra": {"calculus": {}},
}

func main() {
	if list, err := topoSort(prereqs); err == nil {
		for i, course := range list {
			fmt.Printf("%d:\t%s\n", i+1, course)
		}
	} else {
		fmt.Printf("ошибка при сортировке: %q", err)
	}
}

func topoSort(m map[string]map[string]struct{}) (order []string, err error) {
	seen := make(map[string]bool)
	var parentKey string
	var visitAll func(map[string]struct{}) error
	visitAll = func(items map[string]struct{}) error {
		for key := range items {
			if key == parentKey {
				return fmt.Errorf("для курса %s обнаружен цикл", key)
			}
			if !seen[key] {
				seen[key] = true
				err = visitAll(m[key])
				if err != nil {
					return err
				}
				order = append(order, key)
			}
		}
		return nil
	}

	for key := range m {
		if !seen[key] {
			seen[key] = true
			parentKey = key
			err = visitAll(m[key])
			if err != nil {
				return
			}
			order = append(order, key)
		}
	}
	return
}
