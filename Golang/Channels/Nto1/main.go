package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	var w sync.WaitGroup
	w.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		w.Done()
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		w.Done()
	}()

	go func() {
		w.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
