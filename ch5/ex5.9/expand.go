package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {

	f := func(s string) string {
		return strings.Split(s, "$")[1]
	}

	fmt.Println(expand("$arroz $arros $feijao", f))

}

var regexPattern = regexp.MustCompile(`\$\w+`)

func expand(s string, f func(string) string) string {
	wrapperFn := func(s string) string {
		return f(s)
	}
	return regexPattern.ReplaceAllStringFunc(s, wrapperFn)
}
