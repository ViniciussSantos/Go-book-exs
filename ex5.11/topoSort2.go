package main

import (
	"fmt"
	"log"
	"sort"
)

// !+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {

	courses, cycleDetected := topoSort(prereqs)

	if cycleDetected {
		log.Fatalln("cycle detected")
	}

	for i, course := range courses {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, bool) {
	var order []string
	seen := make(map[string]bool)
	//the stack keeps track to see if a course has a cyclic dependency
	stack := make(map[string]bool)

	var visitAll func(items []string) bool

	visitAll = func(items []string) bool {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				stack[item] = false

				if visitAll(m[item]) {
					return true
				}
				stack[item] = true
				order = append(order, item)
			} else if !stack[item] {
				return true
			}
		}
		return false

	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	if found := visitAll(keys); found {
		return nil, true
	}
	return order, false
}
