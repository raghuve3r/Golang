package main

import "fmt"

type student struct {
	fname string
	lname string
	age   int
}

func (s student) fullName() string {
	return s.fname + " " + s.lname
}

type computer struct {
	student
	owns bool
}

func (s student) Greet() string {
	return "Boo"
}

func (c computer) Greet() string {
	return "Booyeah"
}

func main() {
	var raghu student
	raghu.age = 27
	raghu.fname = "Raghuveer"
	raghu.lname = "Krishnamurthy"

	rajath := student{"Rajath", "Kanchibail", 23}
	aish := computer{
		student: student{
			fname: "Aishwarya",
			lname: "Prasad",
			age:   28,
		},
		owns: true,
	}

	fmt.Printf("%T \n", raghu)
	fmt.Println(rajath)
	fmt.Println(raghu)
	fmt.Println(aish)
	fmt.Println(rajath.fullName())
	fmt.Println(raghu.fullName())
	fmt.Println(aish.fullName())
	fmt.Println(rajath.Greet())
	fmt.Println(raghu.Greet())
	fmt.Println(aish.Greet())

}
