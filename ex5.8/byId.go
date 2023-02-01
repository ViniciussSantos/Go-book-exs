package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(*elementById(doc, "footer"))

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	if pre != nil {
		if !pre(n) {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		n = forEachNode(c, pre, nil)
		if n != nil {
			return n
		}
	}

	return nil

}

func elementById(doc *html.Node, id string) *html.Node {

	traversal := func(n *html.Node) bool {
		if n.Type != html.ElementNode || len(n.Attr) == 0 {
			return true
		}

		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				return false
			}
		}
		return true
	}

	return forEachNode(doc, traversal, nil)
}
