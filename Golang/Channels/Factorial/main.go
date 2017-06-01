package main

import "fmt"

var fact int

func main() {
	ans := factorial(5)
	fmt.Println("Recursion: ", ans)

	//using channels
	c := factChannel(5)
	for n := range c {
		fmt.Println("Channels: ", n)
	}
}
func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	fact = factorial(n-1) * n
	return fact
}

func factChannel(n int) chan int {
	res := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		res <- total
		close(res)
	}()
	return res
}
