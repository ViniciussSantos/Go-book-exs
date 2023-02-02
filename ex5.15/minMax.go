package main

import "fmt"

func max(vals ...int) int {

	if len(vals) == 0 {
		return 0
	}

	if len(vals) == 1 {
		return vals[0]
	}

	max := vals[0]

	for i := 1; i < len(vals); i++ {
		if vals[i] > max {
			max = vals[i]
		}
	}
	return max
}

func min(vals ...int) int {

	if len(vals) == 0 {
		return 0
	}

	if len(vals) == 1 {
		return vals[0]
	}

	min := vals[0]

	for i := 1; i < len(vals); i++ {
		if vals[i] < min {
			min = vals[i]
		}
	}
	return min

}

func max1(n int, vals ...int) int {

	if len(vals) == 0 {
		return n
	}

	max := n

	for i := 1; i < len(vals); i++ {
		if vals[i] > max {
			max = vals[i]
		}
	}
	return max
}

func min1(n int, vals ...int) int {
	if len(vals) == 0 {
		return n
	}

	min := n

	for i := 1; i < len(vals); i++ {
		if vals[i] < min {
			min = vals[i]
		}
	}
	return min

}

func main() {

	fmt.Printf("Max: %d\n Min: %d\n", max(1, 2, 3, 4, 5, 6, 7, 8, 9, 100), min(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -1))
	fmt.Printf("Max1: %d\n Min1: %d\n", max1(1, 2, 3, 4, 5, 6, 7, 8, 9, 100), min1(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -1))

}
