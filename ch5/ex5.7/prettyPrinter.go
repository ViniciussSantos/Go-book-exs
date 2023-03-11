package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		Outline(url)
	}
}

func Outline(url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {

	if n.Type == html.ElementNode {
		if n.FirstChild == nil {
			fmt.Printf("%*s<%s %s\\>\n", depth*2, "", n.Data, formatAtrributes(n.Attr))
			return
		}
		fmt.Printf("%*s<%s %s>\n", depth*2, "", n.Data, formatAtrributes(n.Attr))
		depth++
	}

	if n.Type == html.TextNode {
		fmt.Printf("%*s %s\n", depth*2, "", n.Data)
	}

}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode && n.FirstChild != nil {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func formatAtrributes(attributes []html.Attribute) (formatedAttr string) {
	if len(attributes) == 0 {
		return
	}

	for _, attr := range attributes {

		formatedAttr += fmt.Sprintf("%s = \"%s\" ", attr.Key, attr.Val)
	}

	return strings.TrimSpace(formatedAttr)

}
