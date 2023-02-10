package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	url := os.Args[1]
	tags := os.Args[2:]
	outline(url, tags)

}

func outline(url string, tags []string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	for _, v := range ElementsByTagName(doc, tags...) {

		fmt.Printf("%+v\n", v)
	}

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {

	if pre != nil {
		if !pre(n) {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, nil)
	}

	return nil

}

func ElementsByTagName(doc *html.Node, tags ...string) []*html.Node {

	elements := make([]*html.Node, 0)

	traversal := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}

		for _, tag := range tags {
			if tag == n.Data {
				elements = append(elements, n)
			}
		}
		return true
	}

	forEachNode(doc, traversal, nil)

	return elements
}
