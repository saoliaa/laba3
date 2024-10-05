package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	rs := []rune(str)
	out := make([]rune, 0, 0)
	for i := range rs {
		if i == len(rs)-1 {
			out = append(out, rs[i])
		} else {
			out = append(out, rs[i], '*')
		}
	}
	fmt.Println(string(out))
}
