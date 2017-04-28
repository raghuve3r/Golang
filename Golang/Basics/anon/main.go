package main

import "fmt"

func main() {
	x := 0

	/*
	  Anonymous function: Assigning a function to a variable
	  Anonymous functions are within a function
	*/
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
}
