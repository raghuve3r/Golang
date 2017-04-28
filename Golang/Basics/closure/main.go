package main

import "fmt"

var x = 0

func increment() int {
	x++
	return x
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(increment())
	}
}
