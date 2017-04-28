package main

import "fmt"

var a = 5
var b = 2

func max(p int, q int) int {
	if p > q {
		return p
	}
	return q
}
func main() {
	x := max(a, b)
	fmt.Println(x)
}
