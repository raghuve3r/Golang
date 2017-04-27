package main

import "fmt"

const p string = "Blah!"
const (
	x  = iota
	y  = 1 << (iota * 10)
	z  = 1 << (iota * 10)
	xy = 1 << (iota * 10)
)

func main() {
	const q int = 42
	fmt.Println("p -", p)
	a := 10 + q

	//q = a + 10 //constants cannot be changed

	fmt.Println("q -", q)
	fmt.Println(a)

	fmt.Println(x)
	fmt.Printf("%b = \t", y)
	fmt.Printf("%d = 1KB \n", y)
	fmt.Printf("%b = \t", z)
	fmt.Printf("%d = 1MB \n", z)
	fmt.Printf("%b = \t", xy)
	fmt.Printf("%d = 1GB \n", xy)
}
