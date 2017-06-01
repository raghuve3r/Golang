package main

import "fmt"

func main() {
	c := make(chan int)
	b := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		b <- true
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		b <- true
	}()

	go func() {
		<-b
		<-b
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
