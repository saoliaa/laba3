package main

import "fmt"

func main() {
	var a, k int
	fmt.Scan(&a)
	out := make([]int, 0, 0)

	for {
		k = a % 10
		out = append(out, k)
		a = a / 10
		if a == 0 {
			break
		}
	}
	for i := len(out) - 1; i >= 0; i-- {
		fmt.Print(out[i] * out[i])
	}
}
