package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3

	fmt.Println(m)

	m1 := map[string]int{"A": 27, "B": 28, "C": 29}
	fmt.Println(m1)

	var m2 = make(map[string]int)
	m2["AA"] = 11
	m2["BB"] = 22
	fmt.Println(m2)

	var m3 = map[int]string{
		1: "Raghu",
		2: "Aish",
		3: "Rajath",
		4: "Pops",
		5: "Moms",
	}
	fmt.Println(m3)

	for k, v := range m3 {
		fmt.Println("Key: ", k, " - ", "Value: ", v)
	}

	delete(m3, 2)

	if val, exists := m3[2]; exists {
		fmt.Println("Value: ", val)
		fmt.Println("Exists? ", exists)
	} else {
		fmt.Println("Does not exist")
		fmt.Println("Value: ", val)
		fmt.Println("Exists? ", exists)
	}

	fmt.Println(m3)
}
