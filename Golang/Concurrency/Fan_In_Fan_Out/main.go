package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3)
	//fan out
	c1 := sq(in)
	c2 := sq(in)
	for i := range merge(c1, c2) {
		fmt.Println(i)
	}
}

func gen(num ...int) <-chan int {
	res := make(chan int)
	go func() {
		for _, n := range num {
			res <- n
		}
		close(res)
	}()
	return res
}

func sq(in <-chan int) <-chan int {
	res := make(chan int)
	go func() {
		for i := range in {
			res <- i * i
		}
		close(res)
	}()
	return res
}

func merge(in ...<-chan int) <-chan int {
	res := make(chan int)
	var w sync.WaitGroup
	w.Add(len(in))
	for _, c := range in {
		go func(ch <-chan int) {
			for n := range ch {
				res <- n
			}
			w.Done()
		}(c)
	}
	go func() {
		w.Wait()
		close(res)
	}()
	return res
}
