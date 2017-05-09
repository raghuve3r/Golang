package main

import "fmt"

func main() {
	fmt.Println(average(10, 20, 30, 40, 50))
	points := []float64{60, 70, 80, 90, 100}

	//variadic arguments
	fmt.Println(average(points...))
}

func average(val ...float64) float64 {
	fmt.Println(val)
	fmt.Printf("%T \n", val)
	var total float64
	for _, v := range val {
		total += v
	}
	return total / float64(len(val))
}
