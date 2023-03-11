package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	printHTMLText(doc)

}

func printHTMLText(n *html.Node) {

	if n.Data != "script" && n.Data != "style" && n.Type == html.TextNode {
		if len(strings.TrimSpace(n.Data)) != 0 {
			fmt.Println(n.Data)
		}
	}

	if n.FirstChild != nil {
		printHTMLText(n.FirstChild)
	}
	if n.NextSibling != nil {
		printHTMLText(n.NextSibling)
	}

}
