package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Raghu"), boring("Rajath"))
	for n := 0; n < 10; n++ {
		fmt.Println(<-c)
	}
	fmt.Println("Time to snap")
}

func boring(in string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", in, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(ip1, ip2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-ip1
		}
	}()
	go func() {
		for {
			c <- <-ip2
		}
	}()
	return c
}
