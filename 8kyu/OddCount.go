package main

func main() {
	println(OddCount(15))    // => 7
	println(OddCount(15023)) // => 7511
}

func OddCount(n int) int {
	return n / 2
}
