package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var w sync.WaitGroup
var counter int64

func main() {
	w.Add(2)
	go incrementor("Func 1")
	go incrementor("Func 2")
	w.Wait()
	fmt.Println("Final counter value : ", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter: ", counter)
	}
	w.Done()
}
