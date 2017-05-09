package main

import "fmt"

func main() {
	s := printing("Raghu", "Kanchibail")
	fmt.Println(s)
	fmt.Println(printMultiple("Rajath", "Kanchibail"))
}

func printing(fname, lname string) string {
	return fmt.Sprint(fname, " ", lname)
}

//multliple return values
func printMultiple(fname, lname string) (string, string) {
	return fmt.Sprint(fname, " ", lname, "\n"), fmt.Sprint(lname, " ", fname)
}
