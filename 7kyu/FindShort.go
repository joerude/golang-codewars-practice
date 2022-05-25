package main

import "strings"

func main() {
	println(FindShort("bitcoin take over the world maybe who knows perhaps"))
	println(FindShort("turns out random test cases are easier than writing out"))
	println(FindShort("Let's travel abroad shall we"))
}

func FindShort(s string) (c int) {

	words := strings.Split(s, " ")

	for _, word := range words {
		if c == 0 || len(word) < c {
			c = len(word)
		}
	}
	return
}
