package main

import "fmt"

func helloWorld() func() string {
	return func() string {
		return "Heya!"
	}
}

func main() {
	hello := helloWorld()
	fmt.Println(hello())
	fmt.Printf("%T \n", hello)
}
