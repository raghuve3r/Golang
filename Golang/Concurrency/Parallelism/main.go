package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var w sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func main() {
	w.Add(2)
	go fun1()
	go fun2()
	w.Wait()
}

func fun1() {
	for i := 0; i < 100; i++ {
		fmt.Println("Function 1: ", i)
		time.Sleep(time.Duration(20 * time.Millisecond))
	}
	w.Done()
}
func fun2() {
	for i := 0; i < 100; i++ {
		fmt.Println("Funcrion 2: ", i)
		time.Sleep(time.Duration(20 * time.Millisecond))
	}
	w.Done()
}
