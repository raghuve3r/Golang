package main

import "fmt"

var kgToPounds float64 = 2.20462

func main() {
	var kg float64
	fmt.Println("Enter in kg: ")
	fmt.Scan(&kg)
	pounds := kgToPounds * kg
	fmt.Printf("%f kg in pounds is %f \n", kg, pounds)
}
