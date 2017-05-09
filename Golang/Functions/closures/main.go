package main

import "fmt"

func wrapper() func() int32 {
	var x int32

	//scope of a variable
	//x lies only in the scope of this function
	return func() int32 {
		x++
		return x
	}
}

func main() {
	inc := wrapper()
	fmt.Println(inc())
	fmt.Println(inc())
}
