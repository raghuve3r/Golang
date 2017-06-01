package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen()

	c1 := fact(in)
	c2 := fact(in)
	c3 := fact(in)
	c4 := fact(in)
	c5 := fact(in)
	c6 := fact(in)
	c7 := fact(in)
	c8 := fact(in)
	c9 := fact(in)
	c10 := fact(in)

	var i int
	for n := range merge(c1, c2, c3, c4, c5, c6, c7, c8, c9, c10) {
		i++
		fmt.Println(i, "\t", n)
	}
}

func gen() <-chan int {
	res := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			for j := 1; j < 11; j++ {
				res <- j
			}
		}
		close(res)
	}()
	return res
}

func fact(in <-chan int) <-chan int {
	res := make(chan int)
	go func() {
		for i := range in {
			res <- factorial(i)
		}
		close(res)
	}()
	return res
}

func factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(in ...<-chan int) <-chan int {
	var w sync.WaitGroup
	res := make(chan int)

	out := func(c <-chan int) {
		for i := range c {
			res <- i
		}
		w.Done()
	}

	w.Add(len(in))
	for _, c := range in {
		go out(c)
	}

	go func() {
		w.Wait()
		close(res)
	}()
	return res
}
