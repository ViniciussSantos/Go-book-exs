package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	url := os.Args[1]

	words, images, err := CountWordsAndImages(url)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("words: %d \n images: %d\n", words, images)

}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {

	if n.Type == html.TextNode {
		words += countWords(n.Data)
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		wordsChild, imagesChild := countWordsAndImages(c)
		words += wordsChild
		images += imagesChild
	}

	return words, images

}

func countWords(words string) int {
	var n int

	input := bufio.NewScanner(strings.NewReader(words))

	input.Split(bufio.ScanWords)

	for input.Scan() {
		n++
	}

	return n
}
