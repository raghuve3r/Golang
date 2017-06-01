package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] > p[j] }

func main() {

	studyGroup := people{"Zoe", "John", "Al", "Jenny"}
	p := []string{"Zoe", "John", "Al", "Jenny"}
	fmt.Println("Before sorting: \n", studyGroup)
	sort.Sort(studyGroup)
	fmt.Println("After sorting: \n", studyGroup)
	fmt.Println("Before sorting: \n", p)
	sort.Strings(p)
	fmt.Println("After sorting: \n", p)
	i := []int{1, 4, 2, 5, 7, 8, 11, 9}
	fmt.Println("Before sorting: \n", i)
	sort.Ints(i)
	fmt.Println("Before sorting: \n", i)
	i1 := []int{1, 4, 2, 5, 7, 8, 11, 9}
	fmt.Println("Before sorting: \n", i1)
	sort.Sort(sort.Reverse(sort.IntSlice(i1)))
	fmt.Println("Before sorting: \n", i1)
}
