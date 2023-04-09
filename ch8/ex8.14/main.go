package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type client struct {
	msgChannel chan<- string
	name       string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:

			for cli := range clients {
				cli.msgChannel <- msg
			}

		case cli := <-entering:
			var names []string

			if len(clients) > 0 {
				for c := range clients {
					names = append(names, c.name)
				}

				cli.msgChannel <- fmt.Sprintf("online users: %s", strings.Join(names, ", "))
			} else {
				cli.msgChannel <- "you are the first one to connect"
			}

			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.msgChannel)
		}
	}
}

func handleConn(conn net.Conn) {
	chOut := make(chan string)
	chIn := make(chan string)
	go clientWriter(conn, chOut)
	go clientReader(conn, chIn)

	chOut <- "Insert your name:\n"
	who := <-chIn
	chOut <- "You are " + who

	messages <- who + " has arrived"
	entering <- client{chOut, who}

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- client{chOut, who}
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func clientReader(conn net.Conn, ch chan<- string) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		ch <- input.Text()
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
