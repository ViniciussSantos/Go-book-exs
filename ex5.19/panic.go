package main

import "fmt"

func panicAndRecover() (returnValue int) {

	defer func() {
		recover()
		returnValue = 1
	}()

	panic("PANIC")

}

func main() {
	fmt.Println(panicAndRecover())
}
