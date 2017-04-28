package main

import "fmt"

func doZero(y *int) {
	*y = 0
}
func main() {
	x := 5
	doZero(&x)
	fmt.Println(x)
}
