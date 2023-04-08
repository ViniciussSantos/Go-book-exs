package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var done = make(chan struct{})
var responses = make(chan *http.Response)

func query(url string, wg *sync.WaitGroup) {
	newReq, _ := http.NewRequest("GET", url, nil)
	newReq.Cancel = done
	resp, err := http.DefaultClient.Do(newReq)
	defer wg.Done()
	if err != nil {
		log.Fatalf("err: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatalf("getting %s: %s", url, resp.Status)
	}

	responses <- resp
}

func main() {

	wg := sync.WaitGroup{}

	links := os.Args[1:]

	for _, link := range links {
		wg.Add(1)
		go query(link, &wg)
	}

	res := <-responses
	close(done)

	fmt.Println(res.Request.URL)
	wg.Wait()
}
