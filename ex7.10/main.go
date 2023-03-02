package main

import (
	"fmt"
	"sort"
	"strings"
)

func IsPalindrome(s sort.Interface) bool {

	for i := 0; i < s.Len()/2; i++ {
		if !(!s.Less(i, s.Len()-1-i) && !s.Less(s.Len()-1-i, i)) {
			return false
		}
	}
	return true
}

func main() {

	fmt.Println(IsPalindrome(sort.StringSlice(strings.Split("arara", ""))))

}
