package main

import "fmt"

func main() {
	isOdd := func(x int) bool {
		if (x % 2) == 1 {
			return true
		}
		return false
	}

	for i := 1; i < 20; i++ {
		fmt.Println(i, " is Odd? : ", isOdd(i))
	}
}
