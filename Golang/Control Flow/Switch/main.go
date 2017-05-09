package main

import "fmt"

func main() {
	fmt.Println("Enter your name")
	var a string
	fmt.Scan(&a)
	switch a {
	case "Raghu":
		fmt.Println("Your name is: ", a)

	case "Rajath":
		fmt.Println("Your name is: ", a)

	default:
		fmt.Println("Your name is: Radha")

	}
}
