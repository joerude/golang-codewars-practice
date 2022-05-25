package main

func main() {
	println(MakeNegative(4))
	println(MakeNegative(-10))
	println(MakeNegative(176))
}

func MakeNegative(x int) int {

	if x > 0 {
		return -x
	}
	return x

}
