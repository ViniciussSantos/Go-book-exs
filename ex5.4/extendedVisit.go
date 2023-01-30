package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func visit(links []string, node *html.Node) []string {
	switch node.Data {
	case "a":
		for _, a := range node.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	case "script":
		for _, script := range node.Attr {
			if script.Key == "src" {
				links = append(links, script.Val)
			}
		}
	case "link":
		for _, script := range node.Attr {
			if script.Key == "src" {
				links = append(links, script.Val)
			}
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		links = visit(links, child)
	}
	return links
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
