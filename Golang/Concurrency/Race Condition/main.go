package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var w sync.WaitGroup
var counter int

func main() {
	w.Add(2)
	go incrementor("Func 1")
	go incrementor("Func 2")
	w.Wait()
	fmt.Println("Final counter value : ", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, i, "Counter: ", counter)
	}
	w.Done()
}
