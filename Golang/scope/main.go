package main

import "fmt"

var b int = 34 // package scope

func main() {
	a := "Woohoo!"
	fmt.Println(a)
	fmt.Println(c)
	print()
}

var c = 100

func print() {
	fmt.Println(b)
	//fmt.Println(a) //Cannot access : out of scope
}
