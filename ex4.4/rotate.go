package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	rotate(s, 2)
	fmt.Println(s)

}

func rotate(array []int, numberOfRotations int) {

	for i := 0; i < numberOfRotations; i++ {
		first := array[0]
		copy(array, array[1:])
		array[len(array)-1] = first
	}

}
