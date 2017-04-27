package main

import "fmt"

func main() {
	a := 89
	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a //pointer pointing to address of 'a' : referencing
	fmt.Println(b)
	fmt.Println(*b) //value of the pointer pointing at a address : dereferencing

	*b = 90        // changing the value where the pointer b is pointing
	fmt.Println(a) // value of 'a' is changed since 'b' points to 'a'
}
