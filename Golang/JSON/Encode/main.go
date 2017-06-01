package main

import (
	"encoding/json"
	"os"
)

func main() {
	type Person struct {
		First string
		Last  string
		Age   int
	}

	p1 := Person{"Raghu", "Kanchibail", 27}
	json.NewEncoder(os.Stdout).Encode(p1)
}
