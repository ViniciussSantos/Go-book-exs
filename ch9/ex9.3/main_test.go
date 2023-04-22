package memo

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestCancel(t *testing.T) {
	m := New(httpGetBody)
	url := "https://golang.org"
	defer m.Close()
	var n sync.WaitGroup
	n.Add(1)
	go func() {
		defer n.Done()
		start := time.Now()
		value, err := m.Get(url, nil)

		if err != nil {
			log.Print(err)
			return
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}()

	n.Add(1)
	go func() {
		cancel := make(chan struct{})
		close(cancel)
		defer n.Done()
		start := time.Now()
		value, err := m.Get(url, cancel)

		if err != nil {
			log.Print(err)
			return
		}

		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}()

}

// !+httpRequestBody
func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

//!-httpRequestBody

var HTTPGetBody = httpGetBody

type M interface {
	Get(key string) (interface{}, error)
}
