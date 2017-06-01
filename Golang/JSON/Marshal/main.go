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
		RandomVar string
	}

	//using tags

	type AnotherGuy struct {
		FirstName string
		LastName  string `json:"-"`
		Age       int    `json:"age value"`
	}
	g1 := Guy{"Raghu", "Krishnamurthy", 27, "eeho"}
	g2 := AnotherGuy{"Rajath", "Krishnamurthy", 23}

	ms, _ := json.Marshal(g1)
	ms1, _ := json.Marshal(g2)

	fmt.Println(ms)
	fmt.Printf("%T \n", ms)
	fmt.Println(string(ms))
	fmt.Println("")
	fmt.Println(ms1)
	fmt.Printf("%T \n", ms1)
	fmt.Println(string(ms1))

}
