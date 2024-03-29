package main

import (
	"flag"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func startServer(port string) {
	listener, err := net.Listen("tcp", strings.Join([]string{"localhost:", port}, ""))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
}

func main() {
	var port = flag.String("port", "8000", "The port binded to the server")
	flag.Parse()
	startServer(*port)
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
