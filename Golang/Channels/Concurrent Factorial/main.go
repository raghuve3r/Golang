package main

import "fmt"

type fac struct {
	i uint64
	j uint64
}

func main() {
	c := gen()
	ans := factorial(c)
	for n := range ans {
		fmt.Println("Fact of: ", n.i, " is ", n.j)
	}
}

func gen() <-chan uint64 {
	out := make(chan uint64)
	go func() {
		for i := 1; i <= 20; i++ {
			out <- uint64(i)
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan uint64) <-chan fac {
	out := make(chan fac)
	go func() {
		for n := range in {
			out <- fac{n, fact(n)}
		}
		close(out)
	}()
	return out
}

func fact(n uint64) uint64 {
	var total uint64
	total = 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
