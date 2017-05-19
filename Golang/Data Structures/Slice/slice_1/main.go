package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5}
	fmt.Println(sl)
	fmt.Printf("%T \n", sl)
	fmt.Println(sl[2:len(sl)])
	fmt.Println(cap(sl))

	s2 := make([]int, 0, 50)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
	for i := 0; i < 80; i++ {
		s2 = append(s2, i)
		fmt.Println("Len: ", len(s2), "Capacity: ", cap(s2), "Data: ", s2[i])
	}
}
