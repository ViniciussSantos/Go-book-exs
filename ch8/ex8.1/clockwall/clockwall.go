package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type conn struct {
	timezone, host string
}

func main() {
	var args = os.Args[1:]

	if len(args) == 0 {
		log.Fatal("no ports passed. Usage: PLACE=HOST")
	}
	conns := make([]conn, 0)

	for _, v := range args {
		before, after, found := strings.Cut(v, "=")

		if !found {
			log.Fatal("error parsing time zones")
		}

		conns = append(conns, conn{before, after})
	}

	for _, c := range conns {

		conn, err := net.Dial("tcp", c.host)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		go c.logConn(os.Stdout, conn)
	}

	for {
		time.Sleep(time.Minute)
	}

}

func (c conn) logConn(dst io.Writer, src io.Reader) {
	s := bufio.NewScanner(src)
	for s.Scan() {
		fmt.Fprintf(dst, "%s: %s\n", c.timezone, s.Text())
	}
}
