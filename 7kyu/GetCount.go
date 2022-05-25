package main

func main() {
	println(GetCount("abracadabra"))
	println(GetCount("saminahri"))
	println(GetCount("vowels"))
	println(GetCount("dota"))
}

func GetCount(str string) (count int) {
	for _, c := range str {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		default:
		}
	}
	//println(str)
	return count
}
