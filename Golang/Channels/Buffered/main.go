package main

import "fmt"

func main() {
	n := 10
	c := make(chan int)
	b := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			b <- true
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			<-b
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
