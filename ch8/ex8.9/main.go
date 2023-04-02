package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

// !+
func main() {
	// ...determine roots...

	//!-
	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	c := make([]chan int64, len(roots))
	for i := 0; i < len(c); i++ {
		c[i] = make(chan int64)
	}

	var n sync.WaitGroup
	for i, root := range roots {
		n.Add(1)
		go walkDir(root, &n, c[i])
	}

	go func() {
		n.Wait()
		for i := 0; i < len(c); i++ {
			close(c[i])
		}
	}()
	//!-

	// Print the results periodically.
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}

	nbytes := make([]int64, len(c))
	nfiles := make([]int64, len(c))

	var wg sync.WaitGroup

	for i := 0; i < len(c); i++ {
		wg.Add(1)

		go func(j int) {
			defer wg.Done()

		loop:
			for {
				select {
				case size, ok := <-c[j]:
					if !ok {
						break loop
					}
					nfiles[j]++
					nbytes[j] += size
				case <-tick:
					printDiskUsage(roots[j], nfiles[j], nbytes[j])
				}
			}

		}(i)
	}
	wg.Wait()
	for i := 0; i < len(roots); i++ {
		printDiskUsage(roots[i], nfiles[i], nbytes[i])

	}

}

func printDiskUsage(root string, nfiles, nbytes int64) {
	fmt.Printf("%s root %d files  %.1f GB\n", root, nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
