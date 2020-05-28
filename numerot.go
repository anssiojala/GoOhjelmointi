package main

import (
	"fmt"
	"sort"
)

var (
	a, b, c, i, x int
)

func main() {

	fmt.Println("Anna ensimmäinen numero: ")
	fmt.Scan(&a)
	fmt.Println("Anna toinen numero: ")
	fmt.Scan(&b)
	fmt.Println("Anna kolmas numero: ")
	fmt.Scan(&c)

	fmt.Println("Numerot nousevassa järjestyksessä: 1")
	fmt.Println("Numerot laskevassa järjestyksessä: 2")
	fmt.Scan(&x)

	switch x {
	case 1:
		numerot := []int{a, b, c}
		sort.Ints(numerot)
		fmt.Println("Numerosi ovat: ", numerot)

	case 2:
		numerot := []int{a, b, c}
		sort.Sort(sort.Reverse(sort.IntSlice(numerot)))
		fmt.Println("Numerosi ovat: ", numerot)
	}
}
