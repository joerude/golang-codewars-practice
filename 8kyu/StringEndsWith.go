package main

func main() {
	println(StringEndsWith("abc", "c"))
	println(StringEndsWith("banana", "na"))
	println(StringEndsWith("da", "a"))
	println(StringEndsWith("net", "e"))
}

func StringEndsWith(str, ending string) bool {
	if len(ending) > len(str) {
		return false
	}
	return str[len(str)-len(ending):] == ending
}
