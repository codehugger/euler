package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func reverse(s string) string {
	cs := make([]rune, utf8.RuneCountInString(s))
	i := len(cs)
	for _, c := range s {
		i--
		cs[i] = c
	}
	return string(cs)
}

func main() {
	largest := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			p := i * j

			prod := strconv.Itoa(p)

			if prod == reverse(prod) {
				if largest < p {
					largest = p
				}
			}
		}
	}

	fmt.Println("Largest 3-digit product palindrome:", largest)
}
