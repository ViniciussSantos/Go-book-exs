package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	channel := make(chan int)
	var i int64
	start := time.Now()
	go func() {
		channel <- 1
		for {
			i++
			channel <- <-channel
		}
	}()
	go func() {
		for {
			channel <- <-channel
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Printf("\n%f rounds per second\n", float64(i)/float64(time.Since(start))*1e9)
}
