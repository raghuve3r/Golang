package main

import "fmt"

func main() {
	c := incrementor()
	c1 := incrementor()
	c2 := puller(c)
	c3 := puller(c1)
	fmt.Println("Final Answer: ", <-c2+<-c3)

}

func incrementor() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out2 := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
			fmt.Println(sum)
		}
		out2 <- sum
		close(out2)
	}()
	return out2
}
