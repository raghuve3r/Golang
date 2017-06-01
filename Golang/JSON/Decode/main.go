package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	type Person struct {
		First string
		Last  string
		Age   int
	}

	var p1 Person
	read := strings.NewReader(`{"First":"Raghu", "Last":"Kanchibail","Age":27}`)
	json.NewDecoder(read).Decode(&p1)

	fmt.Println(p1.Last)
	fmt.Println(p1.First)
	fmt.Println(p1.Age)
}
