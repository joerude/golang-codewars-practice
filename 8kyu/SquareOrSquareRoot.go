package main

// kata: https://www.codewars.com/kata/55685cd7ad70877c23000102/train/go

import "math"

func main() {
	print(SquareOrSquareRoot([]int{4, 3, 9, 7, 2, 1}))
	println(SquareOrSquareRoot([]int{100, 101, 5, 5, 1, 1}))
	println(SquareOrSquareRoot([]int{1, 2, 3, 4, 5, 6}))
}

func SquareOrSquareRoot(arr []int) []int {
	var m []int
	for _, item := range arr {
		root := math.Sqrt(float64(item))
		if root-math.Floor(root) != 0 {
			m = append(m, item*item)
		} else {
			m = append(m, int(root))
		}
	}
	return m
}
