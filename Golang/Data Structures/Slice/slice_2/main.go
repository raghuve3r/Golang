package main

import "fmt"

func main() {
	records := make([][]string, 0)
	student1 := make([]string, 4)
	student1[0] = "Raghu"
	student1[1] = "Krishnamurthy"
	student1[2] = "4"
	student1[3] = "3.48"

	records = append(records, student1)

	student2 := make([]string, 4)
	student2[0] = "Supreeth"
	student2[1] = "Keragodu"
	student2[2] = "4"
	student2[3] = "3.8"

	records = append(records, student2)

	fmt.Println(records)

	rows := make([][]int, 0, 2)

	for i := 0; i < 2; i++ {
		row := make([]int, 0, 2)
		for j := 0; j < 3; j++ {
			row = append(row, j)
		}
		rows = append(rows, row)
	}
	fmt.Println(rows)

	var slc = make([]string, 1)
	slc[0] = "boo"
	fmt.Println(slc)
	slc = append(slc, "hoo")
	fmt.Println(slc)

	var a []string
	//a[0] = "bling"
	fmt.Println(a == nil)
}
