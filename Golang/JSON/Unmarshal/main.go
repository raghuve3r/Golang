package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Guy struct {
		FirstName string
		LastName  string
		Age       int
	}
	var g1 Guy
	fmt.Println("First Name: ", g1.FirstName)
	fmt.Println("Last Name: ", g1.LastName)
	fmt.Println("Age: ", g1.Age)

	fmt.Println("\nAfter Initializing")
	fmt.Println("----------------------------")

	v := []byte(`{"FirstName":"Raghu","LastName":"Kanchibail","Age":27}`)
	json.Unmarshal(v, &g1)

	fmt.Println("First Name: ", g1.FirstName)
	fmt.Println("Last Name: ", g1.LastName)
	fmt.Println("Age: ", g1.Age)
	fmt.Printf("%T \n", g1)

}
