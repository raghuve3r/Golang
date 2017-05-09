package main

import "fmt"

func main() {
	num := make([]string, 1, 20)
	fmt.Println(num)
	modify(num)
	fmt.Println(num)
}

func modify(l []string) []string {
	l[0] = "Raghu"
	return l
}
