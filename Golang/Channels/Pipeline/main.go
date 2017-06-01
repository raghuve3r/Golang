package main

import "fmt"

func main() {
	for i := range (1,100) {
		fmt.Println(i)
	}
}

func gen(n ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range n {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(num chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range num {
			out <- n * n
		}
		close(out)
	}()
	return out
}
