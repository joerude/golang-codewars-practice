package main

import "fmt"

func main() {
	println(countSheep(15))
	println(countSheep(5))
}

func countSheep(num int) string {
	outputString := ""
	for i := 0; i < num; i++ {
		outputString += fmt.Sprintf("%v sheep...", i+1)
	}
	return outputString
}
