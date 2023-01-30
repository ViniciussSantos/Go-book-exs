package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(countElementNames(make(map[string]int), doc))
}

func countElementNames(elementCount map[string]int, n *html.Node) map[string]int {

	if n.Type == html.ElementNode {
		elementCount[n.Data]++
	}

	if n.FirstChild != nil {
		elementCount = countElementNames(elementCount, n.FirstChild)
	}
	if n.NextSibling != nil {
		elementCount = countElementNames(elementCount, n.NextSibling)
	}

	return elementCount

}
