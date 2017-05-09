package main

import "fmt"

func complex(num []int, callback func(int)) {
	for _, n := range num {
		callback(n)
	}
}

func anotherCallback(numbers []int, callback func(int) bool) []int {
	numberList := []int{}
	for _, i := range numbers {
		if callback(i) {
			numberList = append(numberList, i)
		}
	}
	return numberList
}

func main() {
	complex([]int{1, 2, 3, 4, 5}, func(n int) {
		fmt.Println(n)
	})

	ans := anotherCallback([]int{1, 2, 3, 4, 5, 6}, func(n int) bool {
		return n > 3
	})

	fmt.Println(ans)
}
