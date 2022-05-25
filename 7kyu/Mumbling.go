package main

import "strings"

func main() {
	println(Accum("Some"))
	println(Accum("text"))
	println(Accum("aaa"))
}

func Accum(s string) string {
	var r []string
	for i, c := range s {
		r = append(r, strings.ToUpper(string(c))+strings.Repeat(strings.ToLower(string(c)), i))
	}
	return strings.Join(r, "-")
}
