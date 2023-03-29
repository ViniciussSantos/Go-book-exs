package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func scanner(c net.Conn, text chan<- string) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		text <- input.Text()
	}

}

//!+
func handleConn(c net.Conn) {
	var wg sync.WaitGroup

	defer func() {
		wg.Wait()
		c.Close()
	}()

	text := make(chan string)

	go scanner(c, text)

	timer :=	time.NewTimer(10*time.Second)

	for  {
		select {
		case s := <-text:
			wg.Add(1)
			timer.Reset(10*time.Second)
			go func() {
				defer wg.Done()
				echo(c, s, 1*time.Second)		
			}()		
		case <-timer.C:
			fmt.Fprintln(c, "connection timeout")
			return
		}
	}
	
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
