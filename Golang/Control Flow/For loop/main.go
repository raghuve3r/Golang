package main

import "fmt"

func main() {
	sum := 20
	for i := 0; i < 10; i++ {
		if i != 9 {
			fmt.Print(sum+i, " , ")
		} else {
			fmt.Print(sum + i)
		}

	}
	fmt.Println()
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(i, " - ", j)
		}
	}
	x := 0
	for {
		x++
		if x > 10 {
			break
		} else if x%2 == 0 {
			continue
		}
		fmt.Println(x)
	}
}
