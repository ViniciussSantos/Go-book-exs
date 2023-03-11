package main

import (
	"fmt"
	"io"
	"os"
)

type ByteCounter struct {
	counter int64
	writer  io.Writer
}

func (b *ByteCounter) Write(p []byte) (int, error) {

	b.counter += int64(len(p))

	return b.writer.Write(p)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	countingWriter := ByteCounter{0, w}
	return &countingWriter, &countingWriter.counter
}

func main() {
	coutingWriter, counter := CountingWriter(os.Stdout)
	fmt.Fprintf(coutingWriter, "Testing testing testing\n")
	fmt.Println(*counter)

}
