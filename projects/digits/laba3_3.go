package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)

	rs := []rune(str)
	var max = 0
	for i := range rs {
		if int(rs[i]) > max {
			max = int(rs[i])
		}
	}
	fmt.Println(string(max))
}
