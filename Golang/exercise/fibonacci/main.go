package main

import "fmt"

func main() {
	var total int
	var fib int
	var i = 0
	var j = 1
	for {
		total = i + j
		//fmt.Println(total)
		i = j
		j = total
		if j > 4000000 {
			break
		}
		if total%2 == 0 {
			fib += total
		}
	}
	fmt.Println(fib)
}
