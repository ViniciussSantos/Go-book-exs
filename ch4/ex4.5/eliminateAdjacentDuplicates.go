package main

import "fmt"

func main() {
	duplicates := []string{"a", "a", "b", "c", "c", "c", "d", "d"}

	eliminateAdjacentDuplicates(&duplicates)
	fmt.Println(duplicates)
}

// similar to this algorithm: https://www.javatpoint.com/remove-duplicate-elements-from-an-array-in-c
func eliminateAdjacentDuplicates(s *[]string) {

	for i := 0; i < len(*s)-1; i++ {

		if (*s)[i] == (*s)[i+1] {
			copy((*s)[i:], (*s)[i+1:])
			(*s) = (*s)[:len(*s)-1]
			i--
		}
	}
}
